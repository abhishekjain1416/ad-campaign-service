package deliveryService

import (
	"os"
	"strconv"
	"sync"

	campaignModel "github.com/abhishekjain1416/ad-campaign-service/internal/campaign/model"
	campaignService "github.com/abhishekjain1416/ad-campaign-service/internal/campaign/service"
	deliveryDto "github.com/abhishekjain1416/ad-campaign-service/internal/delivery/dto"
	matchEngineDto "github.com/abhishekjain1416/ad-campaign-service/internal/match_engine/dto"
	matchEngineService "github.com/abhishekjain1416/ad-campaign-service/internal/match_engine/service"
	"github.com/gin-gonic/gin"
)

type DeliveryService interface {
	FetchTargetedCampaigns(ctx *gin.Context, request deliveryDto.DeliverCampaignsRequest) (
		[]deliveryDto.CampaignDetail, error)
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

func (s *deliveryService) FetchTargetedCampaigns(ctx *gin.Context, request deliveryDto.DeliverCampaignsRequest) (
	[]deliveryDto.CampaignDetail, error) {

	// Call match engine to get targeted campaign IDs
	campaignIDs, err := s.matchEngineFilterService.GetTargetedCampaigns(
		matchEngineDto.GetTargetedCampaignsRequest{
			Country: request.Country,
			OS:      request.OS,
			App:     request.App,
		})
	if err != nil {
		return nil, err
	}
	if len(campaignIDs) == 0 {
		return nil, nil
	}

	// Batch the campaign IDs and fetch detail in parallel
	var batchSize, _ = strconv.Atoi(os.Getenv("DELIVERY_CAMPAIGN_BATCH_SIZE"))
	var wg sync.WaitGroup
	var mu sync.Mutex
	var allCampaigns []campaignModel.Campaign
	errCh := make(chan error, 1)

	for i := 0; i < len(campaignIDs); i += batchSize {
		end := i + batchSize
		if end > len(campaignIDs) {
			end = len(campaignIDs)
		}
		batch := campaignIDs[i:end]

		wg.Add(1)
		go func(ids []int64) {
			defer wg.Done()
			campaigns, err := s.campaignService.GetCampaignDetails(ctx, ids)
			if err != nil {
				select {
				case errCh <- err:
				default:
				}
				return
			}
			mu.Lock()
			allCampaigns = append(allCampaigns, campaigns...)
			mu.Unlock()
		}(batch)
	}

	wg.Wait()
	close(errCh)

	if err := <-errCh; err != nil {
		return nil, err
	}

	// Build response
	var response []deliveryDto.CampaignDetail = make([]deliveryDto.CampaignDetail, 0)
	for _, campaign := range allCampaigns {
		response = append(response, deliveryDto.CampaignDetail{
			Code:  campaign.Code,
			Image: campaign.Image,
			CTA:   campaign.CTA,
		})
	}

	return response, nil
}
