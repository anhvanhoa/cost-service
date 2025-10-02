package cost_tracking_service

import (
	"context"

	proto_cost_tracking "github.com/anhvanhoa/sf-proto/gen/cost_tracking/v1"
)

func (s *CostTrackingService) DeleteCostTracking(ctx context.Context, req *proto_cost_tracking.DeleteCostTrackingRequest) (*proto_cost_tracking.DeleteCostTrackingResponse, error) {
	err := s.CostTrackingUsecase.Delete(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &proto_cost_tracking.DeleteCostTrackingResponse{
		Message: "Cost tracking record deleted successfully",
	}, nil
}
