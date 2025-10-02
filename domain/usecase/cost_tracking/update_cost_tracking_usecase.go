package cost_tracking

import (
	"context"
	"time"

	"cost_service/domain/entity"
	"cost_service/domain/repository"
)

type UpdateCostTrackingRequest struct {
	ID              string
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
}

type UpdateCostTrackingUsecaseI interface {
	Execute(ctx context.Context, req *UpdateCostTrackingRequest) (*entity.CostTracking, error)
}

type UpdateCostTrackingUsecase struct {
	costTrackingRepo repository.CostTrackingRepository
}

func NewUpdateCostTrackingUsecase(costTrackingRepo repository.CostTrackingRepository) UpdateCostTrackingUsecaseI {
	return &UpdateCostTrackingUsecase{
		costTrackingRepo: costTrackingRepo,
	}
}

func (u *UpdateCostTrackingUsecase) Execute(ctx context.Context, req *UpdateCostTrackingRequest) (*entity.CostTracking, error) {
	// Get existing record
	existingCostTracking, err := u.costTrackingRepo.GetByID(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	if existingCostTracking == nil {
		return nil, ErrCostTrackingNotFound
	}

	// Validate cost category if provided
	if req.CostCategory != "" && !entity.CostCategory(req.CostCategory).IsValid() {
		return nil, ErrInvalidCostCategory
	}

	// Validate cost type if provided
	if req.CostType != "" && !entity.CostType(req.CostType).IsValid() {
		return nil, ErrInvalidCostType
	}

	// Validate payment method if provided
	if req.PaymentMethod != "" && !entity.PaymentMethod(req.PaymentMethod).IsValid() {
		return nil, ErrInvalidPaymentMethod
	}

	// Validate payment status if provided
	if req.PaymentStatus != "" && !entity.PaymentStatus(req.PaymentStatus).IsValid() {
		return nil, ErrInvalidPaymentStatus
	}

	// Update timestamp
	now := time.Now()
	existingCostTracking.UpdatedAt = &now

	// Save to repository
	err = u.costTrackingRepo.Update(ctx, existingCostTracking)
	if err != nil {
		return nil, err
	}

	return existingCostTracking, nil
}
