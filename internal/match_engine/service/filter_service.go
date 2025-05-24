package matchEngineService

import matchEngineRepository "github.com/abhishekjain1416/ad-campaign-service/internal/match_engine/repository"

type FilterService interface {
}

type filterService struct {
	targetRulesRepository matchEngineRepository.TargetRulesRepository
}

func NewFilterAudienceService(targetRulesRepository matchEngineRepository.TargetRulesRepository) FilterService {
	return &filterService{
		targetRulesRepository: targetRulesRepository,
	}
}
