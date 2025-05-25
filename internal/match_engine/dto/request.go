package matchEngineDto

type GetTargetedCampaignsRequest struct {
	Country string `json:"country"`
	OS      string `json:"os"`
	App     string `json:"app"`
}
