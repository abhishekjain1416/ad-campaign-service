package deliveryService

import (
	campaignService "github.com/abhishekjain1416/ad-campaign-service/internal/campaign/service"
	matchEngineService "github.com/abhishekjain1416/ad-campaign-service/internal/match_engine/service"
)

type DeliveryService interface {
}

type deliveryService struct {
	campaignService          campaignService.CampaignService
	matchEngineFilterService matchEngineService.FilterService
}

func NewDeliveryService(campaignService campaignService.CampaignService,
	matchEngineFilterService matchEngineService.FilterService) DeliveryService {

	return &deliveryService{
		campaignService:          campaignService,
		matchEngineFilterService: matchEngineFilterService,
	}
}
