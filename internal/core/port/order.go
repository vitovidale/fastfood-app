package port

import (
	"context"

	"github.com/vitovidale/fastfood-app/internal/core/domain"
)

type OrderRepositoryReader interface {
	FindByID(ctx context.Context, id domain.ID) (*domain.Order, error)
	FindByCustomer(ctx context.Context, customerId uint64) (*domain.Order, error)
	List(ctx context.Context) ([]*domain.Order, error)

	// return a new tracking number or the existing one
	GetTrackingNumber(ctx context.Context, num *uint16) *uint16

	// return type detached from domain
	FindNestedByID(ctx context.Context, id domain.ID) (any, error)
	FindOrderProduct(ctx context.Context, orderProductId domain.ID) (*domain.OrderProduct, error)
}

type OrderRepositoryWriter interface {
	AddProduct(ctx context.Context, p *domain.OrderProduct) error
	RemoveProduct(ctx context.Context, p *domain.OrderProduct) error
	Save(ctx context.Context, o *domain.Order) (*domain.Order, error)
	Delete(ctx context.Context, id domain.ID) error
	Patch(ctx context.Context, id domain.ID, data *domain.Order) error
}

type OrderRepository interface {
	OrderRepositoryReader
	OrderRepositoryWriter
}

type OrderService interface {
	GetByCustomer(ctx context.Context, customerId uint64) (*domain.Order, error)
	GetNestedByID(ctx context.Context, id domain.ID) (any, error)
	GetByID(ctx context.Context, id domain.ID) (*domain.Order, error)

	// add a product and returns the new order ID or existing order ID
	AddProduct(ctx context.Context, o *domain.Order, p *domain.OrderProduct) error
	RemoveProduct(ctx context.Context, id domain.ID) error

	// create a new order for a customer with a list of products
	Create(ctx context.Context, customerId uint64, products []domain.OrderProduct) (*domain.ID, error)

	// order status movements
	Pay(ctx context.Context, id domain.ID) error
	Prepare(ctx context.Context, id domain.ID) error
	Complete(ctx context.Context, id domain.ID) error
}
