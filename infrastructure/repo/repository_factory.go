package repo

import (
	"cost_service/domain/repository"

	"github.com/anhvanhoa/service-core/utils"
	"github.com/go-pg/pg/v10"
)

type RepositoryFactory struct {
	CostTrackingRepository repository.CostTrackingRepository
}

func NewRepositoryFactory(db *pg.DB, hepler utils.Helper) *RepositoryFactory {
	return &RepositoryFactory{
		CostTrackingRepository: NewCostTrackingRepository(db, hepler),
	}
}

func (f *RepositoryFactory) GetCostTrackingRepository() repository.CostTrackingRepository {
	return f.CostTrackingRepository
}
