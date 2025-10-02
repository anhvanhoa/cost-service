package cost_tracking

import "errors"

// Cost tracking related errors
var (
	ErrCostTrackingNotFound   = errors.New("cost tracking record not found")
	ErrInvalidID              = errors.New("invalid ID")
	ErrInvalidPlantingCycleID = errors.New("invalid planting cycle ID")
	ErrInvalidCostCategory    = errors.New("invalid cost category")
	ErrInvalidCostType        = errors.New("invalid cost type")
	ErrInvalidPaymentMethod   = errors.New("invalid payment method")
	ErrInvalidPaymentStatus   = errors.New("invalid payment status")
	ErrInvalidQuantity        = errors.New("invalid quantity")
	ErrInvalidUnitCost        = errors.New("invalid unit cost")
	ErrInvalidTotalCost       = errors.New("invalid total cost")
	ErrInvalidTaxAmount       = errors.New("invalid tax amount")
	ErrInvalidDiscountAmount  = errors.New("invalid discount amount")
	ErrInvalidWarrantyPeriod  = errors.New("invalid warranty period")
)
