package campaignRepository

type CampaignRepository interface {
}

type campaignRepository struct {
}

func NewCampaignRepository() CampaignRepository {
	return &campaignRepository{}
}
