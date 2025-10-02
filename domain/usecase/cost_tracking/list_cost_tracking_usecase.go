package cost_tracking

import (
	"context"

	"cost_service/domain/entity"
	"cost_service/domain/repository"

	"github.com/anhvanhoa/service-core/common"
)

type ListCostTrackingUsecaseI interface {
	Execute(ctx context.Context, pagination common.Pagination, filter entity.CostTrackingFilter) ([]*entity.CostTracking, int64, error)
}

type ListCostTrackingUsecase struct {
	costTrackingRepo repository.CostTrackingRepository
}

func NewListCostTrackingUsecase(costTrackingRepo repository.CostTrackingRepository) ListCostTrackingUsecaseI {
	return &ListCostTrackingUsecase{
		costTrackingRepo: costTrackingRepo,
	}
}

func (u *ListCostTrackingUsecase) Execute(ctx context.Context, pagination common.Pagination, filter entity.CostTrackingFilter) ([]*entity.CostTracking, int64, error) {
	if pagination.Page <= 0 {
		pagination.Page = 1
	}
	if pagination.PageSize <= 0 {
		pagination.PageSize = 10
	}

	costTrackings, total, err := u.costTrackingRepo.List(ctx, pagination, filter)
	if err != nil {
		return nil, 0, err
	}

	return costTrackings, total, nil
}
