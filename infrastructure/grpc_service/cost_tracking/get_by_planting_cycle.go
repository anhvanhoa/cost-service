package cost_tracking_service

import (
	"context"

	"github.com/anhvanhoa/service-core/common"
	common_proto "github.com/anhvanhoa/sf-proto/gen/common/v1"
	proto_cost_tracking "github.com/anhvanhoa/sf-proto/gen/cost_tracking/v1"
)

func (s *CostTrackingService) GetCostTrackingsByPlantingCycle(ctx context.Context, req *proto_cost_tracking.GetCostTrackingsByPlantingCycleRequest) (*proto_cost_tracking.GetCostTrackingsByPlantingCycleResponse, error) {
	pagination := common.Pagination{
		Page:     int(req.Pagination.Page),
		PageSize: int(req.Pagination.PageSize),
	}

	filter := s.convertCostTrackingFilter(req.Filter)
	costTrackings, total, err := s.CostTrackingUsecase.GetByPlantingCycle(ctx, req.PlantingCycleId, pagination, filter)
	if err != nil {
		return nil, err
	}

	var protoCostTrackings []*proto_cost_tracking.CostTracking
	for _, ct := range costTrackings {
		protoCostTrackings = append(protoCostTrackings, s.convertCostTracking(ct))
	}

	return &proto_cost_tracking.GetCostTrackingsByPlantingCycleResponse{
		CostTrackings: protoCostTrackings,
		Pagination: &common_proto.PaginationResponse{
			Page:       int32(pagination.Page),
			PageSize:   int32(pagination.PageSize),
			Total:      int32(total),
			TotalPages: int32((total + int64(pagination.PageSize) - 1) / int64(pagination.PageSize)),
		},
	}, nil
}
