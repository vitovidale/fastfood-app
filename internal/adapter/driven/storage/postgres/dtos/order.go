package dtos

import (
	"time"
)

type Order struct {
	ID             string         `json:"id" example:"00000000-0000-0000-0000-000000000000"`
	CustomerID     uint64         `json:"customerId" example:"1"`
	Status         string         `json:"status" example:"pending"`
	Total          float64        `json:"total" example:"100"`
	TrackingNumber *uint16        `json:"trackingNumber" example:"1"`
	CreatedAt      time.Time      `json:"createdAt" example:"1970-01-01T00:00:00Z"`
	Products       []OrderProduct `json:"products" gorm:"foreignKey:OrderID"`
}

type OrderProduct struct {
	ID        string  `json:"id" example:"00000000-0000-0000-0000-000000000000"`
	OrderID   string  `json:"orderId" example:"00000000-0000-0000-0000-000000000000"`
	ProductID string  `json:"productId" example:"00000000-0000-0000-0000-000000000000"`
	Product   Product `json:"product"`
	Quantity  uint16  `json:"quantity" example:"1"`
	Notes     string  `json:"notes" example:"notes"`
}

type Product struct {
	ID          string  `json:"id" example:"00000000-0000-0000-0000-000000000000"`
	Name        string  `json:"name" example:"product name"`
	Description string  `json:"description" example:"product description"`
	Price       float64 `json:"price" example:"100"`
}
