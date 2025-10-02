package cost_tracking

import (
	"context"

	"cost_service/domain/repository"
)

type DeleteCostTrackingUsecaseI interface {
	Execute(ctx context.Context, id string) error
}

type DeleteCostTrackingUsecase struct {
	costTrackingRepo repository.CostTrackingRepository
}

func NewDeleteCostTrackingUsecase(costTrackingRepo repository.CostTrackingRepository) DeleteCostTrackingUsecaseI {
	return &DeleteCostTrackingUsecase{
		costTrackingRepo: costTrackingRepo,
	}
}

func (u *DeleteCostTrackingUsecase) Execute(ctx context.Context, id string) error {
	if id == "" {
		return ErrInvalidID
	}

	_, err := u.costTrackingRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	err = u.costTrackingRepo.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
