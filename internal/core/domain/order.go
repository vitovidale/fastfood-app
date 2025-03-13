package domain

import (
	"time"
)

type OrderStatus uint8

// OrderStatus is an enum that represents the status of an order.
//
// The status can be one of the following:
//
// - Pending: the order is being created by the user (adding products)
//
// - Processing: the order is waiting for payment
//
// - Confirmed: the order is paid
//
// - Started: the order is being prepared by the kitchen
//
// - Done: the order is done and sent to customer
//
// - Cancelled: the order is cancelled
const (
	OrderStatusPending    OrderStatus = 0
	OrderStatusProcessing OrderStatus = 1
	OrderStatusConfirmed  OrderStatus = 2
	OrderStatusStarted    OrderStatus = 3
	OrderStatusDone       OrderStatus = 4
	OrderStatusCancelled  OrderStatus = 5
)

func (s OrderStatus) String() string {
	switch s {
	case OrderStatusPending:
		return "pending"
	case OrderStatusProcessing:
		return "processing"
	case OrderStatusConfirmed:
		return "confirmed"
	case OrderStatusStarted:
		return "started"
	case OrderStatusDone:
		return "done"
	case OrderStatusCancelled:
		return "cancelled"
	}
	return "unknown"
}

type Order struct {
	ID             ID     `gorm:"size:36"`
	CustomerID     uint64 `gorm:"type:bigint"`
	Customer       Customer
	Status         string    `gorm:"size:20"`
	Products       []Product `gorm:"many2many:order_products;"`
	Total          float64   `gorm:"not null;precision:14;scale:2;"`
	TrackingNumber *uint16   ``
	CreatedAt      time.Time `gorm:"autoCreateTime;not null"`
	StartedAt      *time.Time
	ReadyAt        *time.Time
	DeletedAt      *time.Time
}

type OrderProduct struct {
	ID        ID      `gorm:"primary_key"`
	OrderID   ID      `gorm:"size:36;not null"`
	ProductID ID      `gorm:"size:36;not null"`
	Order     Order   `gorm:"foreignkey:OrderID"`
	Product   Product `gorm:"foreignkey:ProductID"`
	Quantity  uint16  `gorm:"not null"`
	Total     float64 `gorm:"not null;default:0;precision:14;scale:2;"`
	Notes     string  `gorm:"size:500"`
	CreatedAt time.Time
}

func NewOrderWithCustomer(customerId uint64) *Order {
	return &Order{
		ID:         NewID(),
		CustomerID: customerId,
		Total:      0,
		Status:     OrderStatusPending.String(),
		CreatedAt:  time.Now(),
	}
}

func (o *Order) IsActive() bool {
	return o.DeletedAt.IsZero()
}

func (o *Order) Pay() error {
	// o.Total =
	o.Status = OrderStatusProcessing.String()
	return nil
}

func (o *Order) Start() error {
	if o.StartedAt != nil {
		return ErrorOrderAlreadyStarted
	}

	startedAt := time.Now()
	o.StartedAt = &startedAt
	o.Status = OrderStatusStarted.String()

	return nil
}

func (o *Order) Cancel() error {
	if o.DeletedAt != nil {
		return ErrorOrderAlreadyCancelled
	}

	deletedAt := time.Now()
	o.DeletedAt = &deletedAt
	o.Status = OrderStatusCancelled.String()

	return nil
}

func (o *Order) Complete() error {
	if o.ReadyAt != nil {
		return ErrorOrderAlreadyDone
	}

	readyAt := time.Now()
	o.ReadyAt = &readyAt
	o.Status = OrderStatusDone.String()

	return nil
}
