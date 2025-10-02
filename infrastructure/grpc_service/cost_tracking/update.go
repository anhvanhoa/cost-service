package cost_tracking_service

import (
	"context"
	"cost_service/domain/usecase/cost_tracking"

	proto_cost_tracking "github.com/anhvanhoa/sf-proto/gen/cost_tracking/v1"
)

func (s *CostTrackingService) UpdateCostTracking(ctx context.Context, req *proto_cost_tracking.UpdateCostTrackingRequest) (*proto_cost_tracking.UpdateCostTrackingResponse, error) {
	ct := s.convertRequestUpdateCostTracking(req)
	costTracking, err := s.CostTrackingUsecase.Update(ctx, ct)
	if err != nil {
		return nil, err
	}
	return &proto_cost_tracking.UpdateCostTrackingResponse{
		CostTracking: s.convertCostTracking(costTracking),
	}, nil
}

func (s *CostTrackingService) convertRequestUpdateCostTracking(req *proto_cost_tracking.UpdateCostTrackingRequest) *cost_tracking.UpdateCostTrackingRequest {
	ct := &cost_tracking.UpdateCostTrackingRequest{
		ID:              req.Id,
		PlantingCycleID: req.PlantingCycleId,
		CostCategory:    req.CostCategory,
		CostType:        req.CostType,
		ItemName:        req.ItemName,
		Description:     req.Description,
		Quantity:        req.Quantity,
		Unit:            req.Unit,
		UnitCost:        req.UnitCost,
		TotalCost:       req.TotalCost,
		Currency:        req.Currency,
		Supplier:        req.Supplier,
		SupplierContact: req.SupplierContact,
		InvoiceNumber:   req.InvoiceNumber,
		PaymentMethod:   req.PaymentMethod,
		PaymentStatus:   req.PaymentStatus,
		TaxAmount:       req.TaxAmount,
		DiscountAmount:  req.DiscountAmount,
		Notes:           req.Notes,
		ReceiptImage:    req.ReceiptImage,
		WarrantyPeriod:  int(req.WarrantyPeriod),
	}

	if req.PurchaseDate != nil {
		purchaseDate := req.PurchaseDate.AsTime()
		ct.PurchaseDate = &purchaseDate
	}

	if req.PaymentDueDate != nil {
		paymentDueDate := req.PaymentDueDate.AsTime()
		ct.PaymentDueDate = &paymentDueDate
	}

	return ct
}
