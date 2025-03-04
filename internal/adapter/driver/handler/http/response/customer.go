package response

import (
	"time"

	"github.com/vitovidale/TECH-CHALLENGE/internal/core/domain"
)

type CustomerResponse struct {
	ID        uint64 `json:"id"`
	FirstName string `json:"firstName" example:"John"`
	LastName  string `json:"lastName" example:"Doe"`
	Email     string `json:"email" example:"john.doe@example.com"`
	CreatedAt string `json:"createdAt" example:"1970-01-01T00:00:00Z"`
	UpdatedAt string `json:"updatedAt" example:"1970-01-01T00:00:00Z"`
}

func NewCustomerResponse(customer *domain.Customer) CustomerResponse {
	return CustomerResponse{
		ID:        customer.ID,
		FirstName: customer.FirstName,
		LastName:  customer.LastName,
		Email:     customer.Email,
		CreatedAt: customer.CreatedAt.Format(time.RFC3339),
		UpdatedAt: customer.UpdatedAt.Format(time.RFC3339),
	}
}
