package cost_tracking

import (
	"context"

	"cost_service/domain/entity"
	"cost_service/domain/repository"

	"github.com/anhvanhoa/service-core/common"
)

type CostTrackingUsecaseI interface {
	Create(ctx context.Context, req *CreateCostTrackingRequest) (*entity.CostTracking, error)
	GetByID(ctx context.Context, id string) (*entity.CostTracking, error)
	Update(ctx context.Context, req *UpdateCostTrackingRequest) (*entity.CostTracking, error)
	Delete(ctx context.Context, id string) error
	List(ctx context.Context, pagination common.Pagination, filter entity.CostTrackingFilter) ([]*entity.CostTracking, int64, error)
	GetByPlantingCycle(ctx context.Context, plantingCycleID string, pagination common.Pagination, filter entity.CostTrackingFilter) ([]*entity.CostTracking, int64, error)
}

type CostTrackingUsecase struct {
	createUsecase             CreateCostTrackingUsecaseI
	getUsecase                GetCostTrackingUsecaseI
	updateUsecase             UpdateCostTrackingUsecaseI
	deleteUsecase             DeleteCostTrackingUsecaseI
	listUsecase               ListCostTrackingUsecaseI
	getByPlantingCycleUsecase GetCostTrackingsByPlantingCycleUsecaseI
}

func NewCostTrackingUsecase(costTrackingRepo repository.CostTrackingRepository) CostTrackingUsecaseI {
	return &CostTrackingUsecase{
		createUsecase:             NewCreateCostTrackingUsecase(costTrackingRepo),
		getUsecase:                NewGetCostTrackingUsecase(costTrackingRepo),
		updateUsecase:             NewUpdateCostTrackingUsecase(costTrackingRepo),
		deleteUsecase:             NewDeleteCostTrackingUsecase(costTrackingRepo),
		listUsecase:               NewListCostTrackingUsecase(costTrackingRepo),
		getByPlantingCycleUsecase: NewGetCostTrackingsByPlantingCycleUsecase(costTrackingRepo),
	}
}

func (u *CostTrackingUsecase) Create(ctx context.Context, req *CreateCostTrackingRequest) (*entity.CostTracking, error) {
	return u.createUsecase.Execute(ctx, req)
}

func (u *CostTrackingUsecase) GetByID(ctx context.Context, id string) (*entity.CostTracking, error) {
	return u.getUsecase.Execute(ctx, id)
}

func (u *CostTrackingUsecase) Update(ctx context.Context, req *UpdateCostTrackingRequest) (*entity.CostTracking, error) {
	return u.updateUsecase.Execute(ctx, req)
}

func (u *CostTrackingUsecase) Delete(ctx context.Context, id string) error {
	return u.deleteUsecase.Execute(ctx, id)
}

func (u *CostTrackingUsecase) List(ctx context.Context, pagination common.Pagination, filter entity.CostTrackingFilter) ([]*entity.CostTracking, int64, error) {
	return u.listUsecase.Execute(ctx, pagination, filter)
}

func (u *CostTrackingUsecase) GetByPlantingCycle(ctx context.Context, plantingCycleID string, pagination common.Pagination, filter entity.CostTrackingFilter) ([]*entity.CostTracking, int64, error) {
	return u.getByPlantingCycleUsecase.Execute(ctx, plantingCycleID, pagination, filter)
}
