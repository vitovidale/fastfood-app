package response

import (
	"time"

	"github.com/vitovidale/fastfood-app/internal/core/domain"
)

type CategoryResponse struct {
	ID        domain.ID  `json:"id"`
	Name      string     `json:"name" example:"Snacks"`
	CreatedAt time.Time  `json:"createdAt" example:"1970-01-01T00:00:00Z"`
	UpdatedAt *time.Time `json:"updatedAt" example:"1970-01-01T00:00:00Z"`
}

func NewCategoryResponse(category *domain.Category) CategoryResponse {
	return CategoryResponse{
		ID:        category.ID,
		Name:      category.Name,
		CreatedAt: category.CreatedAt,
		UpdatedAt: category.UpdatedAt,
	}
}

func NewCategoryListResponse(categories []*domain.Category) []CategoryResponse {
	var list []CategoryResponse
	for _, category := range categories {
		list = append(list, NewCategoryResponse(category))
	}
	return list
}
