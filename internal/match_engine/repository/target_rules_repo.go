package matchEngineRepository

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/abhishekjain1416/ad-campaign-service/db"
	matchEngineModel "github.com/abhishekjain1416/ad-campaign-service/internal/match_engine/model"
	"github.com/abhishekjain1416/ad-campaign-service/pkg/constants"
)

type TargetRulesRepository interface {
	GetQualifiedCampaigns(dimension, value string) ([]int64, error)
	GetCampaignsWithNoRules(dimension string) ([]int64, error)
}

type targetRulesRepository struct {
}

func NewTargetRulesRepository() TargetRulesRepository {
	return &targetRulesRepository{}
}

func (r *targetRulesRepository) GetQualifiedCampaigns(dimension, value string) ([]int64, error) {
	var campaignIDs []int64

	// Generate Redis key
	cacheKey := fmt.Sprintf("%s%s:%s", constants.RedisQualifiedCampaignsKeyPrefix, dimension, value)

	// Try fetching from Redis
	cached, err := db.Redis.Get(context.Background(), cacheKey).Result()
	if err == nil {
		// Cache hit
		log.Printf("[CACHE HIT] Key: %s", cacheKey)
		if unmarshalErr := json.Unmarshal([]byte(cached), &campaignIDs); unmarshalErr == nil {
			return campaignIDs, nil
		}
		// If unmarshal fails, fall back to DB and overwrite cache
	}

	// Cache miss or error, query DB
	log.Printf("[CACHE MISS] Key: %s", cacheKey)
	err = db.DB.
		Model(&matchEngineModel.TargetRule{}).
		Select("DISTINCT campaign_id").
		Where("dimension = ?", dimension).
		Where(`
			(rule_type = ? AND ? = ANY(values)) OR 
			(rule_type = ? AND ? <> ALL(values))`,
			constants.RuleTypeInclude, value, constants.RuleTypeExclude, value).
		Where("deleted = ?", false).
		Find(&campaignIDs).
		Error
	if err != nil {
		return campaignIDs, err
	}

	// Cache the result in Redis
	data, _ := json.Marshal(campaignIDs)
	_ = db.Redis.Set(context.Background(), cacheKey, data, constants.RedisQualifiedCampaignsKeyExpiry).Err()

	return campaignIDs, nil
}

func (r *targetRulesRepository) GetCampaignsWithNoRules(dimension string) ([]int64, error) {
	var campaignIDs []int64

	cacheKey := fmt.Sprintf("%s%s", constants.RedisCampaignsWithNoRulesKeyPrefix, dimension)

	// Try fetching from Redis
	cached, err := db.Redis.Get(context.Background(), cacheKey).Result()
	if err == nil {
		log.Printf("[CACHE HIT] Key: %s", cacheKey)
		if unmarshalErr := json.Unmarshal([]byte(cached), &campaignIDs); unmarshalErr == nil {
			return campaignIDs, nil
		}
	}

	// Cache miss or unmarshal failure, run DB query
	log.Printf("[CACHE MISS] Key: %s", cacheKey)
	query := `
		SELECT c.id
		FROM campaigns c
		LEFT JOIN target_rules tr
			ON c.id = tr.campaign_id
			AND tr.dimension = ?
			AND tr.deleted = ?
		WHERE
			c.deleted = ?
			AND c.status = ?
			AND tr.id IS NULL;
	`

	err = db.DB.
		Raw(query, dimension, false, false, constants.StatusActive).
		Scan(&campaignIDs).
		Error
	if err != nil {
		return nil, err
	}

	// Cache the result
	data, _ := json.Marshal(campaignIDs)
	_ = db.Redis.Set(context.Background(), cacheKey, data, constants.RedisCampaignsWithNoRulesKeyExpiry).Err()

	return campaignIDs, nil
}
