package cost_tracking

import (
	"context"

	"cost_service/domain/entity"
	"cost_service/domain/repository"
)

type GetCostTrackingUsecaseI interface {
	Execute(ctx context.Context, id string) (*entity.CostTracking, error)
}

type GetCostTrackingUsecase struct {
	costTrackingRepo repository.CostTrackingRepository
}

func NewGetCostTrackingUsecase(costTrackingRepo repository.CostTrackingRepository) GetCostTrackingUsecaseI {
	return &GetCostTrackingUsecase{
		costTrackingRepo: costTrackingRepo,
	}
}

func (u *GetCostTrackingUsecase) Execute(ctx context.Context, id string) (*entity.CostTracking, error) {
	if id == "" {
		return nil, ErrInvalidID
	}

	costTracking, err := u.costTrackingRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if costTracking == nil {
		return nil, ErrCostTrackingNotFound
	}

	return costTracking, nil
}
