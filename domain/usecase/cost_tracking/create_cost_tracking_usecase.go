package cost_tracking

import (
	"context"
	"time"

	"cost_service/domain/entity"
	"cost_service/domain/repository"
)

type CreateCostTrackingRequest struct {
	PlantingCycleID string
	CostCategory    string
	CostType        string
	ItemName        string
	Description     string
	Quantity        float64
	Unit            string
	UnitCost        float64
	TotalCost       float64
	Currency        string
	PurchaseDate    *time.Time
	Supplier        string
	SupplierContact string
	InvoiceNumber   string
	PaymentMethod   string
	PaymentStatus   string
	PaymentDueDate  *time.Time
	TaxAmount       float64
	DiscountAmount  float64
	WarrantyPeriod  int
	Notes           string
	ReceiptImage    string
	CreatedBy       string
}

type CreateCostTrackingUsecaseI interface {
	Execute(ctx context.Context, req *CreateCostTrackingRequest) (*entity.CostTracking, error)
}

type CreateCostTrackingUsecase struct {
	costTrackingRepo repository.CostTrackingRepository
}

func NewCreateCostTrackingUsecase(costTrackingRepo repository.CostTrackingRepository) CreateCostTrackingUsecaseI {
	return &CreateCostTrackingUsecase{
		costTrackingRepo: costTrackingRepo,
	}
}

func (u *CreateCostTrackingUsecase) Execute(ctx context.Context, req *CreateCostTrackingRequest) (*entity.CostTracking, error) {
	if !entity.CostCategory(req.CostCategory).IsValid() {
		return nil, ErrInvalidCostCategory
	}

	if !entity.CostType(req.CostType).IsValid() {
		return nil, ErrInvalidCostType
	}

	if req.PaymentMethod != "" && !entity.PaymentMethod(req.PaymentMethod).IsValid() {
		return nil, ErrInvalidPaymentMethod
	}

	if req.PaymentStatus != "" && !entity.PaymentStatus(req.PaymentStatus).IsValid() {
		return nil, ErrInvalidPaymentStatus
	}

	if req.Currency == "" {
		req.Currency = "VND"
	}

	costTracking := &entity.CostTracking{
		PlantingCycleID: req.PlantingCycleID,
		CostCategory:    req.CostCategory,
		CostType:        req.CostType,
		ItemName:        req.ItemName,
		Description:     req.Description,
		Quantity:        req.Quantity,
		Unit:            req.Unit,
		UnitCost:        req.UnitCost,
		TotalCost:       req.TotalCost,
		Currency:        req.Currency,
		PurchaseDate:    req.PurchaseDate,
		Supplier:        req.Supplier,
		SupplierContact: req.SupplierContact,
		InvoiceNumber:   req.InvoiceNumber,
		PaymentMethod:   req.PaymentMethod,
		PaymentStatus:   req.PaymentStatus,
		PaymentDueDate:  req.PaymentDueDate,
		TaxAmount:       req.TaxAmount,
		DiscountAmount:  req.DiscountAmount,
		WarrantyPeriod:  req.WarrantyPeriod,
		Notes:           req.Notes,
		ReceiptImage:    req.ReceiptImage,
		CreatedBy:       req.CreatedBy,
		CreatedAt:       time.Now(),
	}

	// Save to repository
	err := u.costTrackingRepo.Create(ctx, costTracking)
	if err != nil {
		return nil, err
	}

	return costTracking, nil
}
