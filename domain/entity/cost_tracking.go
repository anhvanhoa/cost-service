package entity

import (
	"time"
)

type CostTracking struct {
	tableName       struct{} `pg:"cost_trackings"`
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
	CreatedBy       string
	CreatedAt       time.Time
	UpdatedAt       *time.Time
}

func (c *CostTracking) TableName() any {
	return c.tableName
}

type CostCategory string

const (
	CostCategorySeed           CostCategory = "seed"
	CostCategoryFertilizer     CostCategory = "fertilizer"
	CostCategoryPesticide      CostCategory = "pesticide"
	CostCategoryLabor          CostCategory = "labor"
	CostCategoryUtilities      CostCategory = "utilities"
	CostCategoryEquipment      CostCategory = "equipment"
	CostCategoryPackaging      CostCategory = "packaging"
	CostCategoryTransportation CostCategory = "transportation"
)

type CostType string

const (
	CostTypeFixed     CostType = "fixed"
	CostTypeVariable  CostType = "variable"
	CostTypeOneTime   CostType = "one_time"
	CostTypeRecurring CostType = "recurring"
)

type PaymentMethod string

const (
	PaymentMethodCash PaymentMethod = "cash"
	PaymentMethodBank PaymentMethod = "bank"
	PaymentMethodCard PaymentMethod = "card"
)

type PaymentStatus string

const (
	PaymentStatusPending   PaymentStatus = "pending"
	PaymentStatusPaid      PaymentStatus = "paid"
	PaymentStatusOverdue   PaymentStatus = "overdue"
	PaymentStatusCancelled PaymentStatus = "cancelled"
)

func (c CostCategory) IsValid() bool {
	switch c {
	case CostCategorySeed, CostCategoryFertilizer, CostCategoryPesticide,
		CostCategoryLabor, CostCategoryUtilities, CostCategoryEquipment,
		CostCategoryPackaging, CostCategoryTransportation:
		return true
	default:
		return false
	}
}

func (c CostType) IsValid() bool {
	switch c {
	case CostTypeFixed, CostTypeVariable, CostTypeOneTime, CostTypeRecurring:
		return true
	default:
		return false
	}
}

func (p PaymentMethod) IsValid() bool {
	switch p {
	case PaymentMethodCash, PaymentMethodBank, PaymentMethodCard:
		return true
	default:
		return false
	}
}

func (p PaymentStatus) IsValid() bool {
	switch p {
	case PaymentStatusPending, PaymentStatusPaid, PaymentStatusOverdue, PaymentStatusCancelled:
		return true
	default:
		return false
	}
}

type CostTrackingFilter struct {
	Category      string
	CostType      string
	PaymentStatus string
	Supplier      string
	Search        string
}

type CostSummary struct {
	TotalCost              float64
	TotalRecords           int64
	CategoryBreakdown      []CategoryCostBreakdown
	PaymentStatusBreakdown []PaymentStatusBreakdown
	MonthlyBreakdown       []MonthlyCostBreakdown
}

type CategoryCostBreakdown struct {
	Category   string
	TotalCost  float64
	Count      int64
	Percentage float64
}

type PaymentStatusBreakdown struct {
	Status     string
	TotalCost  float64
	Count      int64
	Percentage float64
}

type MonthlyCostBreakdown struct {
	Month     string
	Year      int
	TotalCost float64
	Count     int64
}
