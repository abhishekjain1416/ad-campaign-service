package campaignRepository

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/abhishekjain1416/ad-campaign-service/db"
	campaignModel "github.com/abhishekjain1416/ad-campaign-service/internal/campaign/model"
	"github.com/abhishekjain1416/ad-campaign-service/pkg/constants"
	"github.com/gin-gonic/gin"
)

type CampaignRepository interface {
	GetCampaignDetailsByIDs(ctx *gin.Context, campaignIDs []int64) ([]campaignModel.Campaign, error)
}

type campaignRepository struct {
}

func NewCampaignRepository() CampaignRepository {
	return &campaignRepository{}
}

func (r *campaignRepository) GetCampaignDetailsByIDs(ctx *gin.Context, campaignIDs []int64) (
	[]campaignModel.Campaign, error) {

	// Step 1: Prepare Redis keys
	keys := make([]string, len(campaignIDs))
	for index, id := range campaignIDs {
		key := fmt.Sprintf("%s%d", constants.RedisCampaignDetailKeyPrefix, id)
		keys[index] = key
	}

	// Step 2: Try to get from Redis
	cacheResults, err := db.Redis.MGet(ctx, keys...).Result()
	if err != nil {
		return nil, err
	}

	var cachedCampaigns []campaignModel.Campaign
	var missingIDs []int64

	for index, raw := range cacheResults {
		if raw == nil {
			// Cache miss
			missingIDs = append(missingIDs, campaignIDs[index])
			continue
		}
		var campaign campaignModel.Campaign
		if err := json.Unmarshal([]byte(raw.(string)), &campaign); err == nil {
			cachedCampaigns = append(cachedCampaigns, campaign)
		} else {
			// If unmarshal fails, refetch from DB
			missingIDs = append(missingIDs, campaignIDs[index])
		}
	}

	// Step 3: Fetch missing ones from DB
	var dbCampaigns []campaignModel.Campaign
	if len(missingIDs) > 0 {
		err := db.DB.
			Model(&campaignModel.Campaign{}).
			Where("id IN ? AND deleted = ? AND status = ?", missingIDs, false, constants.StatusActive).
			Find(&dbCampaigns).
			Error
		if err != nil {
			return nil, err
		}

		// Step 4: Cache the DB results
		pipe := db.Redis.Pipeline()
		for _, c := range dbCampaigns {
			data, _ := json.Marshal(c)
			key := fmt.Sprintf("%s%d", constants.RedisCampaignDetailKeyPrefix, c.ID)
			pipe.Set(ctx, key, data, constants.RedisCampaignDetailKeyExpiry)
		}
		_, _ = pipe.Exec(ctx)
	}

	log.Printf("Cache hit: %d campaigns found in Redis", len(cachedCampaigns))
	log.Printf("Cache miss: %d campaigns fetched from DB", len(dbCampaigns))

	// Step 5: Combine and return
	finalCampaigns := append(cachedCampaigns, dbCampaigns...)

	return finalCampaigns, nil
}
