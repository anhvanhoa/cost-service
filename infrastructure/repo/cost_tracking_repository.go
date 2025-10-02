package repo

import (
	"context"
	"fmt"

	"cost_service/domain/entity"
	"cost_service/domain/repository"

	"github.com/anhvanhoa/service-core/common"
	"github.com/anhvanhoa/service-core/utils"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

type costTrackingRepository struct {
	db *pg.DB
	h  utils.Helper
}

func NewCostTrackingRepository(db *pg.DB, h utils.Helper) repository.CostTrackingRepository {
	return &costTrackingRepository{
		db: db,
		h:  h,
	}
}

func (r *costTrackingRepository) Create(ctx context.Context, costTracking *entity.CostTracking) error {
	_, err := r.db.ModelContext(ctx, costTracking).Insert()
	return err
}

func (r *costTrackingRepository) GetByID(ctx context.Context, id string) (*entity.CostTracking, error) {
	costTracking := &entity.CostTracking{}
	err := r.db.ModelContext(ctx, costTracking).Where("id = ?", id).Select()
	if err != nil {
		return nil, err
	}
	return costTracking, nil
}

func (r *costTrackingRepository) Update(ctx context.Context, costTracking *entity.CostTracking) error {
	_, err := r.db.ModelContext(ctx, costTracking).Where("id = ?", costTracking.ID).UpdateNotZero()
	return err
}

func (r *costTrackingRepository) Delete(ctx context.Context, id string) error {
	_, err := r.db.ModelContext(ctx, &entity.CostTracking{}).Where("id = ?", id).Delete()
	return err
}

func (r *costTrackingRepository) List(ctx context.Context, pagination common.Pagination, filter entity.CostTrackingFilter) ([]*entity.CostTracking, int64, error) {
	var costTrackings []*entity.CostTracking
	query := r.db.ModelContext(ctx, &costTrackings)

	// Apply filters
	query = r.applyFilters(query, filter)

	// Get total count
	total, err := query.Count()
	if err != nil {
		return nil, 0, err
	}

	// Apply pagination
	query = r.applyPagination(query, pagination)

	// Execute query
	err = query.Select()
	if err != nil {
		return nil, 0, err
	}

	return costTrackings, int64(total), nil
}

func (r *costTrackingRepository) GetByPlantingCycle(ctx context.Context, plantingCycleID string, pagination common.Pagination, filter entity.CostTrackingFilter) ([]*entity.CostTracking, int64, error) {
	var costTrackings []*entity.CostTracking
	query := r.db.ModelContext(ctx, &costTrackings).Where("planting_cycle_id = ?", plantingCycleID)

	// Apply filters
	query = r.applyFilters(query, filter)

	// Get total count
	total, err := query.Count()
	if err != nil {
		return nil, 0, err
	}

	// Apply pagination
	r.applyPagination(query, pagination)

	// Execute query
	err = query.Select()
	if err != nil {
		return nil, 0, err
	}

	return costTrackings, int64(total), nil
}

func (r *costTrackingRepository) GetTotalCostByPlantingCycle(ctx context.Context, plantingCycleID string) (float64, error) {
	var total float64
	err := r.db.ModelContext(ctx, &entity.CostTracking{}).
		Where("planting_cycle_id = ?", plantingCycleID).
		ColumnExpr("COALESCE(SUM(total_cost), 0)").
		Select(&total)
	return total, err
}

func (r *costTrackingRepository) GetCostSummary(ctx context.Context, plantingCycleID *string) (*entity.CostSummary, error) {
	query := r.db.ModelContext(ctx, &entity.CostTracking{})

	if plantingCycleID != nil {
		query = query.Where("planting_cycle_id = ?", *plantingCycleID)
	}

	var totalCost float64
	var totalRecords int64

	err := query.ColumnExpr("COALESCE(SUM(total_cost), 0)").Select(&totalCost)
	if err != nil {
		return nil, err
	}

	err = query.ColumnExpr("COUNT(*)").Select(&totalRecords)
	if err != nil {
		return nil, err
	}

	var categoryBreakdown []entity.CategoryCostBreakdown
	err = query.ColumnExpr("cost_category as category, COALESCE(SUM(total_cost), 0) as total_cost, COUNT(*) as count").
		Group("cost_category").
		Select(&categoryBreakdown)
	if err != nil {
		return nil, err
	}

	for i := range categoryBreakdown {
		if totalCost > 0 {
			categoryBreakdown[i].Percentage = (categoryBreakdown[i].TotalCost / totalCost) * 100
		}
	}

	var paymentStatusBreakdown []entity.PaymentStatusBreakdown
	err = query.ColumnExpr("payment_status as status, COALESCE(SUM(total_cost), 0) as total_cost, COUNT(*) as count").
		Group("payment_status").
		Select(&paymentStatusBreakdown)
	if err != nil {
		return nil, err
	}

	for i := range paymentStatusBreakdown {
		if totalCost > 0 {
			paymentStatusBreakdown[i].Percentage = (paymentStatusBreakdown[i].TotalCost / totalCost) * 100
		}
	}

	var monthlyBreakdown []entity.MonthlyCostBreakdown
	err = query.ColumnExpr("TO_CHAR(purchase_date, 'YYYY-MM') as month, EXTRACT(YEAR FROM purchase_date) as year, COALESCE(SUM(total_cost), 0) as total_cost, COUNT(*) as count").
		Where("purchase_date IS NOT NULL").
		Group("TO_CHAR(purchase_date, 'YYYY-MM'), EXTRACT(YEAR FROM purchase_date)").
		Order("year DESC, month DESC").
		Select(&monthlyBreakdown)
	if err != nil {
		return nil, err
	}

	return &entity.CostSummary{
		TotalCost:              totalCost,
		TotalRecords:           totalRecords,
		CategoryBreakdown:      categoryBreakdown,
		PaymentStatusBreakdown: paymentStatusBreakdown,
		MonthlyBreakdown:       monthlyBreakdown,
	}, nil
}

func (r *costTrackingRepository) applyFilters(query *orm.Query, filter entity.CostTrackingFilter) *orm.Query {
	if filter.Category != "" {
		query = query.Where("cost_category = ?", filter.Category)
	}

	if filter.CostType != "" {
		query = query.Where("cost_type = ?", filter.CostType)
	}

	if filter.PaymentStatus != "" {
		query = query.Where("payment_status = ?", filter.PaymentStatus)
	}

	if filter.Supplier != "" {
		query = query.Where("supplier ILIKE ?", "%"+filter.Supplier+"%")
	}

	if filter.Search != "" {
		searchTerm := "%" + filter.Search + "%"
		query = query.Where("(item_name ILIKE ? OR description ILIKE ? OR supplier ILIKE ?)",
			searchTerm, searchTerm, searchTerm)
	}

	return query
}

func (r *costTrackingRepository) applyPagination(query *orm.Query, pagination common.Pagination) *orm.Query {
	if pagination.Page > 0 {
		offset := r.h.CalculateOffset(pagination.Page, pagination.PageSize)
		query = query.Offset(offset)
	}
	if pagination.PageSize > 0 {
		query = query.Limit(pagination.PageSize)
	}

	if pagination.SortBy != "" {
		sortOrder := "ASC"
		if pagination.SortOrder == "desc" {
			sortOrder = "DESC"
		}
		query = query.Order(fmt.Sprintf("%s %s", pagination.SortBy, sortOrder))
	} else {
		query = query.Order("created_at DESC")
	}
	return query
}
