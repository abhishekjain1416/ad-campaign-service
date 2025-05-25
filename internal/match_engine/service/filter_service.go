package matchEngineService

import (
	"sync"

	matchEngineDto "github.com/abhishekjain1416/ad-campaign-service/internal/match_engine/dto"
	matchEngineRepository "github.com/abhishekjain1416/ad-campaign-service/internal/match_engine/repository"
	"github.com/abhishekjain1416/ad-campaign-service/pkg/constants"
)

type FilterService interface {
	GetTargetedCampaigns(request matchEngineDto.GetTargetedCampaignsRequest) ([]int64, error)
}

type filterService struct {
	targetRulesRepository matchEngineRepository.TargetRulesRepository
}

func NewFilterAudienceService(targetRulesRepository matchEngineRepository.TargetRulesRepository) FilterService {
	return &filterService{
		targetRulesRepository: targetRulesRepository,
	}
}

func (s *filterService) GetTargetedCampaigns(request matchEngineDto.GetTargetedCampaignsRequest) ([]int64, error) {
	type qualifiedCampaigns struct {
		campaigns []int64
		err       error
	}

	var wg sync.WaitGroup
	var numberOfDimensions int = 3
	ch := make(chan qualifiedCampaigns, numberOfDimensions)

	// Define dimensions to be checked
	dimensions := []struct {
		key   string
		value string
	}{
		{constants.DimensionCountry, request.Country},
		{constants.DimensionOS, request.OS},
		{constants.DimensionApp, request.App},
	}

	// For each dimension, launch a goroutine to fetch both qualified and no-rule campaigns
	for _, dimension := range dimensions {
		wg.Add(1)
		go func(dimension, value string) {
			defer wg.Done()
			qualified, err1 := s.targetRulesRepository.GetQualifiedCampaigns(dimension, value)
			if err1 != nil {
				ch <- qualifiedCampaigns{nil, err1}
				return
			}

			noRules, err2 := s.targetRulesRepository.GetCampaignsWithNoRules(dimension)
			if err2 != nil {
				ch <- qualifiedCampaigns{nil, err2}
				return
			}

			// Union of qualified and no-rule campaigns
			unionMap := make(map[int64]bool)
			for _, id := range qualified {
				unionMap[id] = true
			}
			for _, id := range noRules {
				unionMap[id] = true
			}
			unionList := make([]int64, 0, len(unionMap))
			for id := range unionMap {
				unionList = append(unionList, id)
			}

			ch <- qualifiedCampaigns{unionList, nil}
		}(dimension.key, dimension.value)
	}

	// Wait for all goroutines to complete
	wg.Wait()
	close(ch)

	// Collect goroutine results
	allQualifiedCampaigns := make([][]int64, 0, numberOfDimensions)
	for result := range ch {
		if result.err != nil {
			return nil, result.err
		}
		allQualifiedCampaigns = append(allQualifiedCampaigns, result.campaigns)
	}

	// Build intersection of results
	campaignMatchCount := make(map[int64]int)
	for _, campaignsList := range allQualifiedCampaigns {
		for _, id := range campaignsList {
			campaignMatchCount[id]++
		}
	}

	var finalCampaigns []int64
	for campaignId, count := range campaignMatchCount {
		if count == numberOfDimensions {
			finalCampaigns = append(finalCampaigns, campaignId)
		}
	}

	return finalCampaigns, nil
}
