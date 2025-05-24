package matchEngineRepository

type TargetRulesRepository interface {
}

type targetRulesRepository struct {
}

func NewTargetRulesRepository() TargetRulesRepository {
	return &targetRulesRepository{}
}
