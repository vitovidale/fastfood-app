package service

import (
	"context"
	"time"

	"github.com/vitovidale/fastfood-app/internal/core/domain"
	"github.com/vitovidale/fastfood-app/internal/core/port"
)

type OrderService struct {
	orderRepository    port.OrderRepository
	productRepository  port.ProductRepository
	customerRepository port.CustomerRepository
}

func NewOrderService(
	orderRepository port.OrderRepository,
	productRepository port.ProductRepository,
	customerRepository port.CustomerRepository,
) *OrderService {
	return &OrderService{
		orderRepository:    orderRepository,
		productRepository:  productRepository,
		customerRepository: customerRepository,
	}
}

// GetByCustomerID returns an order by its customer ID.
func (s *OrderService) GetByCustomer(ctx context.Context, customerId uint64) (*domain.Order, error) {
	o, err := s.orderRepository.FindByCustomer(ctx, customerId)
	if err != nil {
		return nil, err
	}
	return o, nil
}

func (s *OrderService) List(ctx context.Context) ([]*domain.Order, error) {
	orders, err := s.orderRepository.List(ctx)
	if err != nil {
		return nil, err
	}
	return orders, nil
}

// AddProduct adds a product to an order, based on the order and product data.
func (s *OrderService) AddProduct(ctx context.Context, o *domain.Order, p *domain.OrderProduct) error {
	// TODO cache products for faster lookup
	product, err := s.productRepository.FindByID(ctx, p.ProductID)

	if err != nil {
		if err.Error() == domain.ErrorDataNotFound.Error() {
			return domain.ErrorProductNotFound
		}
		return err
	}

	p.ID = domain.NewID()
	p.OrderID = o.ID
	p.Total = product.Price * float64(p.Quantity)

	err = s.orderRepository.AddProduct(ctx, p)

	if err != nil {
		return err
	}

	// update the order values
	s.orderRepository.Patch(ctx, o.ID, &domain.Order{
		Total:          o.Total + p.Total,
		TrackingNumber: s.orderRepository.GetTrackingNumber(ctx, o.TrackingNumber),
	})

	return nil
}

// RemoveProduct removes a product from an order, based on the order ID and the order product ID.
func (s *OrderService) RemoveProduct(ctx context.Context, orderId domain.ID, orderProductId domain.ID) error {
	o, err := s.orderRepository.FindByID(ctx, orderId)

	if err != nil {
		return err
	}

	op, err := s.orderRepository.FindOrderProduct(ctx, orderProductId)

	if err != nil {
		return err
	}

	err = s.orderRepository.RemoveProduct(ctx, &domain.OrderProduct{ID: orderProductId})
	if err != nil {
		return err
	}

	// update the order values
	_ = s.orderRepository.Patch(ctx, o.ID, &domain.Order{
		Total: o.Total - op.Total,
	})

	return nil
}

func (s *OrderService) Pay(ctx context.Context, id domain.ID) error {
	o, err := s.orderRepository.FindByID(ctx, id)

	if err != nil {
		return err
	}

	if o.Status != domain.OrderStatusPending.String() {
		return domain.ErrorOrderAlreadyProcessing
	}

	err = s.orderRepository.Patch(ctx, id, &domain.Order{Status: domain.OrderStatusProcessing.String()})
	if err != nil {
		return err
	}

	// wait for 5 seconds to MOCK pay the order
	time.Sleep(time.Second * 5)

	err = s.orderRepository.Patch(ctx, id, &domain.Order{Status: domain.OrderStatusConfirmed.String()})
	if err != nil {
		return err
	}

	return nil
}

func (s *OrderService) Prepare(ctx context.Context, id domain.ID) error {
	o, err := s.orderRepository.FindByID(ctx, id)
	if err != nil {
		return err
	}

	if o.Status != domain.OrderStatusConfirmed.String() {
		return domain.ErrorOrderAlreadyStarted
	}

	startedAt := time.Now()

	err = s.orderRepository.Patch(ctx, id, &domain.Order{
		Status:    domain.OrderStatusStarted.String(),
		StartedAt: &startedAt,
	})

	if err != nil {
		return err
	}

	return nil
}

func (s *OrderService) Complete(ctx context.Context, id domain.ID) error {
	o, err := s.orderRepository.FindByID(ctx, id)
	if err != nil {
		return err
	}

	if o.Status != domain.OrderStatusStarted.String() {
		return domain.ErrorOrderAlreadyDone
	}

	readyAt := time.Now()
	err = s.orderRepository.Patch(ctx, id, &domain.Order{
		Status:  domain.OrderStatusDone.String(),
		ReadyAt: &readyAt,
	})

	if err != nil {
		return err
	}

	return nil
}

// TODO better type safe custom model handling
func (s *OrderService) GetNestedByID(ctx context.Context, id domain.ID) (any, error) {
	o, err := s.orderRepository.FindNestedByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return o, nil
}

func (s *OrderService) Create(ctx context.Context, customerId uint64, products []domain.OrderProduct) (*domain.ID, error) {
	_, err := s.customerRepository.FindByID(ctx, customerId)
	if err != nil {
		if err.Error() == domain.ErrorDataNotFound.Error() {
			return nil, domain.ErrorCustomerNotFound
		}
		return nil, err
	}

	order, err := s.orderRepository.FindByCustomer(ctx, customerId)
	if err != nil {
		if err.Error() != domain.ErrorDataNotFound.Error() {
			return nil, err
		}

		// order = domain.NewOrderWithCustomer(customerId)
		order, err = s.orderRepository.Save(ctx, domain.NewOrderWithCustomer(customerId))
	}

	for _, product := range products {
		err = s.AddProduct(ctx, order, &product)

		if err != nil {
			return nil, err
		}
	}

	return &order.ID, nil
}

func (s *OrderService) GetByID(ctx context.Context, id domain.ID) (*domain.Order, error) {
	o, err := s.orderRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return o, nil
}
