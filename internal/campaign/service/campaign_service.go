package campaignService

import (
	"log"

	campaignModel "github.com/abhishekjain1416/ad-campaign-service/internal/campaign/model"
	campaignRepository "github.com/abhishekjain1416/ad-campaign-service/internal/campaign/repository"
	"github.com/gin-gonic/gin"
)

type CampaignService interface {
	GetCampaignDetails(ctx *gin.Context, campaignIDs []int64) ([]campaignModel.Campaign, error)
}

type campaignService struct {
	campaignRepository campaignRepository.CampaignRepository
}

func NewCampaignService(campaignRepository campaignRepository.CampaignRepository) CampaignService {
	return &campaignService{
		campaignRepository: campaignRepository,
	}
}

func (s *campaignService) GetCampaignDetails(ctx *gin.Context, campaignIDs []int64) ([]campaignModel.Campaign, error) {
	if len(campaignIDs) == 0 {
		return []campaignModel.Campaign{}, nil
	}

	campaigns, err := s.campaignRepository.GetCampaignDetailsByIDs(ctx, campaignIDs)
	if err != nil {
		log.Printf("Error fetching campaign details: %v", err)
		return nil, err
	}

	return campaigns, nil
}
