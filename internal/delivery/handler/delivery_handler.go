package deliveryHandler

import (
	"net/http"

	deliveryDto "github.com/abhishekjain1416/ad-campaign-service/internal/delivery/dto"
	deliveryService "github.com/abhishekjain1416/ad-campaign-service/internal/delivery/service"
	"github.com/abhishekjain1416/ad-campaign-service/pkg/constants"
	"github.com/abhishekjain1416/ad-campaign-service/pkg/dto"
	"github.com/abhishekjain1416/ad-campaign-service/pkg/util"
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

	var request deliveryDto.DeliverCampaignsRequest
	var response deliveryDto.DeliverCampaignsResponse

	if err := ctx.ShouldBindQuery(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, util.
			ErrorResponse(constants.DeliverCampaignsMalformedRequest.GetErrorCode(),
				constants.DeliverCampaignsMalformedRequest.GetMessage(), []dto.Error{{ErrorMessage: err.Error()}}))
		return
	}
	campaigns, err := h.deliveryService.FetchTargetedCampaigns(ctx, request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, util.
			ErrorResponse(constants.DeliverCampaignsError.GetErrorCode(),
				constants.DeliverCampaignsError.GetMessage(), []dto.Error{{ErrorMessage: err.Error()}}))
		return
	}
	if len(campaigns) == 0 {
		ctx.JSON(http.StatusNoContent, nil)
	}
	response.Campaigns = campaigns

	ctx.JSON(http.StatusOK, util.SuccessResponse(constants.DeliverCampaignsSuccess.GetSuccessCode(),
		constants.DeliverCampaignsSuccess.GetMessage(), response))
}
