package deliveryHandler

import (
	deliveryService "github.com/abhishekjain1416/ad-campaign-service/internal/delivery/service"
	"github.com/gin-gonic/gin"
)

type DeliveryHandler interface {
	DeliverCampaigns(ctx *gin.Context)
}

type deliveryHandler struct {
	deliveryService deliveryService.DeliveryService
}

func NewDeliveryHandler(deliveryService deliveryService.DeliveryService) DeliveryHandler {
	return &deliveryHandler{
		deliveryService: deliveryService,
	}
}

func (h *deliveryHandler) DeliverCampaigns(ctx *gin.Context) {
}
