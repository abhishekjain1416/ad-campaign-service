package audienceService

import audienceRepository "github.com/abhishekjain1416/ad-campaign-service/internal/audience/repository"

type FilterAudienceService interface {
}

type filterAudienceService struct {
	targetRulesRepository audienceRepository.TargetRulesRepository
}

func NewFilterAudienceService(targetRulesRepository audienceRepository.TargetRulesRepository) FilterAudienceService {
	return &filterAudienceService{
		targetRulesRepository: targetRulesRepository,
	}
}
