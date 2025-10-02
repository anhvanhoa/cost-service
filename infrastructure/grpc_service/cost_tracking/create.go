package cost_tracking_service

import (
	"context"
	"cost_service/domain/usecase/cost_tracking"

	proto_cost_tracking "github.com/anhvanhoa/sf-proto/gen/cost_tracking/v1"
)

func (s *CostTrackingService) CreateCostTracking(ctx context.Context, req *proto_cost_tracking.CreateCostTrackingRequest) (*proto_cost_tracking.CreateCostTrackingResponse, error) {
	ct := s.convertRequestCreateCostTracking(req)
	costTracking, err := s.CostTrackingUsecase.Create(ctx, ct)
	if err != nil {
		return nil, err
	}
	return &proto_cost_tracking.CreateCostTrackingResponse{
		CostTracking: s.convertCostTracking(costTracking),
	}, nil
}

func (s *CostTrackingService) convertRequestCreateCostTracking(req *proto_cost_tracking.CreateCostTrackingRequest) *cost_tracking.CreateCostTrackingRequest {
	ct := &cost_tracking.CreateCostTrackingRequest{
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
		CreatedBy:       req.CreatedBy,
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
