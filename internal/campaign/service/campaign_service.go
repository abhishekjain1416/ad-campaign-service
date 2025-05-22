package campaignService

import campaignRepository "github.com/abhishekjain1416/ad-campaign-service/internal/campaign/repository"

type CampaignService interface {
}

type campaignService struct {
	campaignRepository campaignRepository.CampaignRepository
}

func NewCampaignService(campaignRepository campaignRepository.CampaignRepository) CampaignService {
	return &campaignService{
		campaignRepository: campaignRepository,
	}
}
