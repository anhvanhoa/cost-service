package cost_tracking_service

import (
	"context"
	"cost_service/domain/entity"

	"github.com/anhvanhoa/service-core/common"
	common_proto "github.com/anhvanhoa/sf-proto/gen/common/v1"
	proto_cost_tracking "github.com/anhvanhoa/sf-proto/gen/cost_tracking/v1"
)

func (s *CostTrackingService) ListCostTrackings(ctx context.Context, req *proto_cost_tracking.ListCostTrackingsRequest) (*proto_cost_tracking.ListCostTrackingsResponse, error) {
	pagination := common.Pagination{
		Page:     int(req.Pagination.Page),
		PageSize: int(req.Pagination.PageSize),
	}

	filter := s.convertCostTrackingFilter(req.Filter)

	costTrackings, total, err := s.CostTrackingUsecase.List(ctx, pagination, filter)
	if err != nil {
		return nil, err
	}

	// Convert response
	var protoCostTrackings []*proto_cost_tracking.CostTracking
	for _, ct := range costTrackings {
		protoCostTrackings = append(protoCostTrackings, s.convertCostTracking(ct))
	}

	return &proto_cost_tracking.ListCostTrackingsResponse{
		CostTrackings: protoCostTrackings,
		Pagination: &common_proto.PaginationResponse{
			Page:       int32(pagination.Page),
			PageSize:   int32(pagination.PageSize),
			Total:      int32(total),
			TotalPages: int32((total + int64(pagination.PageSize) - 1) / int64(pagination.PageSize)),
		},
	}, nil
}

func (s *CostTrackingService) convertCostTrackingFilter(filter *proto_cost_tracking.CostTrackingFilter) entity.CostTrackingFilter {
	if filter == nil {
		return entity.CostTrackingFilter{}
	}

	return entity.CostTrackingFilter{
		Category:      filter.Category,
		CostType:      filter.CostType,
		PaymentStatus: filter.PaymentStatus,
		Supplier:      filter.Supplier,
		Search:        filter.Search,
	}
}
