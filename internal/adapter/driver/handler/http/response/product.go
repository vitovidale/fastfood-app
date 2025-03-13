package response

import (
	"time"

	"github.com/vitovidale/fastfood-app/internal/core/domain"
)

type ProductResponse struct {
	ID          string           `json:"id" example:"1"`
	Name        string           `json:"name" example:"Potato Chips"`
	Price       float64          `json:"price" example:"10000"`
	Description string           `json:"description" example:"Potato chips with cheese flavor"`
	Category    CategoryResponse `json:"category"`
	CreatedAt   time.Time        `json:"createdAt" example:"1970-01-01T00:00:00Z"`
	UpdatedAt   *time.Time       `json:"updatedAt" example:"1970-01-01T00:00:00Z"`
}

func NewProductResponse(product *domain.Product) ProductResponse {
	return ProductResponse{
		ID:          product.ID.String(),
		Name:        product.Name,
		Price:       product.Price,
		Description: product.Description,
		Category:    NewCategoryResponse(product.Category),
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
	}
}

func NewProductListResponse(products []*domain.Product) []ProductResponse {
	var list []ProductResponse
	for _, product := range products {
		list = append(list, NewProductResponse(product))
	}
	return list
}
