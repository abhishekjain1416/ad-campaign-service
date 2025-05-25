package matchEngineRepository

import (
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

	err := db.DB.
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
	return campaignIDs, nil
}

func (r *targetRulesRepository) GetCampaignsWithNoRules(dimension string) ([]int64, error) {
	var campaignIDs []int64

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

	err := db.DB.
		Raw(query, dimension, false, false, constants.StatusActive).
		Scan(&campaignIDs).
		Error
	if err != nil {
		return nil, err
	}

	return campaignIDs, nil
}
