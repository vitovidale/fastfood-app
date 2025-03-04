package port

import (
	"context"

	"github.com/vitovidale/TECH-CHALLENGE/internal/core/domain"
)

// CustomerRepositoryReader is an interface that wraps all the reading operations for a customer.
type CustomerRepositoryReader interface {
	FindByID(ctx context.Context, id uint64) (*domain.Customer, error)
	FindByKeys(ctx context.Context, id uint64, email string) (*domain.Customer, error)
}

// CustomerRepositoryWriter is an interface that wraps all the writing operations for a customer.
type CustomerRepositoryWriter interface {
	Create(ctx context.Context, c *domain.Customer) error
	Patch(ctx context.Context, id uint64, data *domain.Customer) error
}

// CustomerRepository is an interface that wraps all the reading and writing operations for a customer.
type CustomerRepository interface {
	CustomerRepositoryReader
	CustomerRepositoryWriter
}

// CustomerService is an interface that wraps all the operations for a customer.
type CustomerService interface {
	GetByID(ctx context.Context, id uint64) (*domain.Customer, error)
	Create(ctx context.Context, c *domain.Customer) (*domain.Customer, error)
	Authenticate(ctx context.Context, c *domain.Customer) (*domain.Customer, error)
	Update(ctx context.Context, c *domain.Customer) (*domain.Customer, error)
	Delete(ctx context.Context, id uint64) error
}
