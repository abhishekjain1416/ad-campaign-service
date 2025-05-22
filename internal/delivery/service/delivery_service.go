package deliveryService

import (
	audienceService "github.com/abhishekjain1416/ad-campaign-service/internal/audience/service"
	campaignService "github.com/abhishekjain1416/ad-campaign-service/internal/campaign/service"
)

type DeliveryService interface {
}

type deliveryService struct {
	campaignService       campaignService.CampaignService
	filterAudienceService audienceService.FilterAudienceService
}

func NewDeliveryService(campaignService campaignService.CampaignService,
	filterAudienceService audienceService.FilterAudienceService) DeliveryService {

	return &deliveryService{
		campaignService:       campaignService,
		filterAudienceService: filterAudienceService,
	}
}
