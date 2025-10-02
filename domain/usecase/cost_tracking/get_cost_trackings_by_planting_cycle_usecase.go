package cost_tracking

import (
	"context"

	"cost_service/domain/entity"
	"cost_service/domain/repository"

	"github.com/anhvanhoa/service-core/common"
)

type GetCostTrackingsByPlantingCycleUsecaseI interface {
	Execute(ctx context.Context, plantingCycleID string, pagination common.Pagination, filter entity.CostTrackingFilter) ([]*entity.CostTracking, int64, error)
}

type GetCostTrackingsByPlantingCycleUsecase struct {
	costTrackingRepo repository.CostTrackingRepository
}

func NewGetCostTrackingsByPlantingCycleUsecase(costTrackingRepo repository.CostTrackingRepository) GetCostTrackingsByPlantingCycleUsecaseI {
	return &GetCostTrackingsByPlantingCycleUsecase{
		costTrackingRepo: costTrackingRepo,
	}
}

func (u *GetCostTrackingsByPlantingCycleUsecase) Execute(ctx context.Context, plantingCycleID string, pagination common.Pagination, filter entity.CostTrackingFilter) ([]*entity.CostTracking, int64, error) {
	if plantingCycleID == "" {
		return nil, 0, ErrInvalidPlantingCycleID
	}

	if pagination.Page <= 0 {
		pagination.Page = 1
	}
	if pagination.PageSize <= 0 {
		pagination.PageSize = 10
	}

	costTrackings, total, err := u.costTrackingRepo.GetByPlantingCycle(ctx, plantingCycleID, pagination, filter)
	if err != nil {
		return nil, 0, err
	}

	return costTrackings, total, nil
}
