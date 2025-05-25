package deliveryDto

type DeliverCampaignsResponse struct {
	Campaigns []CampaignDetail `json:"campaigns"`
}

type CampaignDetail struct {
	Code  string `json:"cid"`
	Image string `json:"img"`
	CTA   string `json:"cta"`
}
