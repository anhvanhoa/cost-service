package repository

import (
	"context"

	"cost_service/domain/entity"

	"github.com/anhvanhoa/service-core/common"
)

type CostTrackingRepository interface {
	Create(ctx context.Context, costTracking *entity.CostTracking) error
	GetByID(ctx context.Context, id string) (*entity.CostTracking, error)
	Update(ctx context.Context, costTracking *entity.CostTracking) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context, pagination common.Pagination, filter entity.CostTrackingFilter) ([]*entity.CostTracking, int64, error)
	GetByPlantingCycle(ctx context.Context, plantingCycleID string, pagination common.Pagination, filter entity.CostTrackingFilter) ([]*entity.CostTracking, int64, error)
	GetTotalCostByPlantingCycle(ctx context.Context, plantingCycleID string) (float64, error)
	GetCostSummary(ctx context.Context, plantingCycleID *string) (*entity.CostSummary, error)
}
