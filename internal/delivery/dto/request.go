package deliveryDto

type DeliverCampaignsRequest struct {
	Country string `form:"country" binding:"required"`
	OS      string `form:"os" binding:"required"`
	App     string `form:"app" binding:"required"`
}
