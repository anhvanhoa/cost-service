package cost_tracking_service

import (
	"context"

	proto_cost_tracking "github.com/anhvanhoa/sf-proto/gen/cost_tracking/v1"
)

func (s *CostTrackingService) GetCostTracking(ctx context.Context, req *proto_cost_tracking.GetCostTrackingRequest) (*proto_cost_tracking.GetCostTrackingResponse, error) {
	costTracking, err := s.CostTrackingUsecase.GetByID(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &proto_cost_tracking.GetCostTrackingResponse{
		CostTracking: s.convertCostTracking(costTracking),
	}, nil
}
