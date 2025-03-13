package response

import (
	"time"

	"github.com/vitovidale/fastfood-app/internal/core/domain"
)

type OrderResponse struct {
	ID             domain.ID              `json:"id" example:"1"`
	CustomerID     uint64                 `json:"customerId" example:"1"`
	Total          float64                `json:"total" example:"100"`
	Status         string                 `json:"status" example:"pending"`
	TrackingNumber *uint16                `json:"trackingNumber" example:"1"`
	CreatedAt      time.Time              `json:"createdAt" example:"1970-01-01T00:00:00Z"`
	StartedAt      *time.Time             `json:"startedAt" example:"1970-01-01T00:00:00Z"`
	ReadyAt        *time.Time             `json:"readyAt" example:"1970-01-01T00:00:00Z"`
	Products       []OrderProductResponse `json:"products"`
}

type OrderProductResponse struct {
	ID        string          `json:"id" example:"00000000-0000-0000-0000-000000000000"`
	ProductID string          `json:"productId" example:"00000000-0000-0000-0000-000000000000"`
	Product   ProductResponse `json:"product"`
	Quantity  uint16          `json:"quantity" example:"1"`
	Notes     string          `json:"notes" example:"notes"`
}

func NewOrderResponse(order *domain.Order) OrderResponse {
	return OrderResponse{
		ID:             order.ID,
		CustomerID:     order.CustomerID,
		Total:          order.Total,
		Status:         order.Status,
		TrackingNumber: order.TrackingNumber,
		CreatedAt:      order.CreatedAt,
		StartedAt:      order.StartedAt,
		ReadyAt:        order.ReadyAt,
	}
}
