package cost_tracking_service

import (
	"cost_service/domain/entity"
	"cost_service/domain/usecase/cost_tracking"
	"cost_service/infrastructure/repo"

	proto_cost_tracking "github.com/anhvanhoa/sf-proto/gen/cost_tracking/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type CostTrackingService struct {
	CostTrackingUsecase cost_tracking.CostTrackingUsecaseI
	proto_cost_tracking.UnsafeCostTrackingServiceServer
}

func NewCostTrackingService(repo *repo.RepositoryFactory) proto_cost_tracking.CostTrackingServiceServer {
	costTrackingUsecase := cost_tracking.NewCostTrackingUsecase(repo.GetCostTrackingRepository())
	return &CostTrackingService{
		CostTrackingUsecase: costTrackingUsecase,
	}
}

func (s *CostTrackingService) convertCostTracking(costTracking *entity.CostTracking) *proto_cost_tracking.CostTracking {
	ct := &proto_cost_tracking.CostTracking{
		Id:              costTracking.ID,
		PlantingCycleId: costTracking.PlantingCycleID,
		CostCategory:    costTracking.CostCategory,
		CostType:        costTracking.CostType,
		ItemName:        costTracking.ItemName,
		Description:     costTracking.Description,
		Quantity:        costTracking.Quantity,
		Unit:            costTracking.Unit,
		UnitCost:        costTracking.UnitCost,
		TotalCost:       costTracking.TotalCost,
		Currency:        costTracking.Currency,
		Supplier:        costTracking.Supplier,
		SupplierContact: costTracking.SupplierContact,
		InvoiceNumber:   costTracking.InvoiceNumber,
		PaymentMethod:   costTracking.PaymentMethod,
		PaymentStatus:   costTracking.PaymentStatus,
		TaxAmount:       costTracking.TaxAmount,
		DiscountAmount:  costTracking.DiscountAmount,
		WarrantyPeriod:  int32(costTracking.WarrantyPeriod),
		Notes:           costTracking.Notes,
		ReceiptImage:    costTracking.ReceiptImage,
		CreatedBy:       costTracking.CreatedBy,
		CreatedAt:       timestamppb.New(costTracking.CreatedAt),
	}

	if costTracking.PurchaseDate != nil {
		ct.PurchaseDate = timestamppb.New(*costTracking.PurchaseDate)
	}

	if costTracking.PaymentDueDate != nil {
		ct.PaymentDueDate = timestamppb.New(*costTracking.PaymentDueDate)
	}

	if costTracking.UpdatedAt != nil {
		ct.UpdatedAt = timestamppb.New(*costTracking.UpdatedAt)
	}

	return ct
}
