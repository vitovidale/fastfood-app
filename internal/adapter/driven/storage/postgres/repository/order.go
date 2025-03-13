package repository

import (
	"context"
	"time"

	"github.com/vitovidale/fastfood-app/internal/adapter/driven/storage/postgres"
	"github.com/vitovidale/fastfood-app/internal/adapter/driven/storage/postgres/dtos"
	"github.com/vitovidale/fastfood-app/internal/core/domain"
)

type OrderRepository struct {
	db *postgres.DB
}

func NewOrderRepository(db *postgres.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (r *OrderRepository) Save(ctx context.Context, o *domain.Order) (*domain.Order, error) {
	result := r.db.WithContext(ctx).Save(&o)
	if result.Error != nil {
		return nil, result.Error
	}
	return o, nil
}

func (r *OrderRepository) Delete(ctx context.Context, id domain.ID) error {
	result := r.db.WithContext(ctx).
		Where("id = ?", id).
		Update("deleted_at", time.Now())

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *OrderRepository) FindByID(ctx context.Context, id domain.ID) (*domain.Order, error) {
	o := &domain.Order{}

	result := r.db.WithContext(ctx).
		Where("deleted_at IS NULL").
		First(&o, id)

	if result.Error != nil {
		return nil, result.Error
	}
	return o, nil
}

func (r *OrderRepository) List(ctx context.Context) ([]*domain.Order, error) {
	var orders []*domain.Order

	result := r.db.WithContext(ctx).
		Order("status ASC").
		Where("deleted_at IS NULL AND status NOT IN (?, ?)", domain.OrderStatusCancelled.String(), domain.OrderStatusDone.String()).
		Find(&orders)

	if result.Error != nil {
		return nil, result.Error
	}
	return orders, nil
}

func (r *OrderRepository) FindNestedByID(ctx context.Context, id domain.ID) (any, error) {
	data := dtos.Order{}

	result := r.db.WithContext(ctx).
		Table("orders").
		Preload("Products.Product").
		First(&data, id)

	if result.Error != nil {
		return nil, result.Error
	}

	return data, nil
}

func (r *OrderRepository) FindByCustomer(ctx context.Context, id uint64) (*domain.Order, error) {
	o := &domain.Order{}

	result := r.db.WithContext(ctx).
		First(&o, "customer_id = ? AND deleted_at IS NULL AND status NOT IN (?, ?)", id, domain.OrderStatusCancelled.String(), domain.OrderStatusDone.String())

	if result.Error != nil {
		return nil, result.Error
	}
	return o, nil
}

func (r *OrderRepository) FindOrderProduct(ctx context.Context, orderProductId domain.ID) (*domain.OrderProduct, error) {
	p := &domain.OrderProduct{}

	result := r.db.WithContext(ctx).
		First(&p, orderProductId)

	if result.Error != nil {
		return nil, result.Error
	}
	return p, nil
}

func (r *OrderRepository) AddProduct(ctx context.Context, p *domain.OrderProduct) error {
	result := r.db.WithContext(ctx).
		Create(&p)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *OrderRepository) RemoveProduct(ctx context.Context, p *domain.OrderProduct) error {
	result := r.db.WithContext(ctx).
		Delete(&domain.OrderProduct{}, p.ID)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *OrderRepository) Patch(ctx context.Context, id domain.ID, data *domain.Order) error {
	result := r.db.WithContext(ctx).
		Model(&domain.Order{}).
		Where("id = ?", id).
		Updates(data)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

// GetTrackingNumber returns the next tracking number available for an order or the existing one
// TODO create a reset heuristic
func (r *OrderRepository) GetTrackingNumber(ctx context.Context, num *uint16) *uint16 {
	if num != nil && *num > 0 {
		return num
	}

	num = new(uint16)

	result := r.db.WithContext(ctx).
		Raw(`SELECT nextval('order_tracking_number_sequence');`).
		Scan(num)

	if result.Error != nil {
		return nil
	}

	return num
}
