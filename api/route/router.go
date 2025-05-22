package route

import (
	audienceRepository "github.com/abhishekjain1416/ad-campaign-service/internal/audience/repository"
	audienceService "github.com/abhishekjain1416/ad-campaign-service/internal/audience/service"
	campaignRepository "github.com/abhishekjain1416/ad-campaign-service/internal/campaign/repository"
	campaignService "github.com/abhishekjain1416/ad-campaign-service/internal/campaign/service"
	deliveryHandler "github.com/abhishekjain1416/ad-campaign-service/internal/delivery/handler"
	deliveryService "github.com/abhishekjain1416/ad-campaign-service/internal/delivery/service"
	"github.com/gin-gonic/gin"
)

var (
	campaignRepositoryObj campaignRepository.CampaignRepository = campaignRepository.NewCampaignRepository()
	campaignServiceObj    campaignService.CampaignService       = campaignService.NewCampaignService(campaignRepositoryObj)

	targetRulesRepositoryObj audienceRepository.TargetRulesRepository = audienceRepository.NewTargetRulesRepository()
	filterAudienceServiceObj audienceService.FilterAudienceService    = audienceService.NewFilterAudienceService(targetRulesRepositoryObj)

	deliveryServiceObj deliveryService.DeliveryService = deliveryService.NewDeliveryService(campaignServiceObj, filterAudienceServiceObj)
	deliveryHandlerObj deliveryHandler.DeliveryHandler = deliveryHandler.NewDeliveryHandler(deliveryServiceObj)
)

func RegisterRoutes(router *gin.Engine) {

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	v1 := router.Group("/api/v1")
	{
		deliveryRoutes := v1.Group("/delivery")
		{
			deliveryRoutes.GET("", deliveryHandlerObj.DeliverCampaigns)
		}
	}
}
