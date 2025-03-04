package port

import (
	"context"

	"github.com/vitovidale/TECH-CHALLENGE/internal/core/domain"
)

// CategoryRepositoryReader is an interface that wraps all the reading operations for a category.
type CategoryRepositoryReader interface {
	FindCategoryByID(ctx context.Context, id domain.ID) (*domain.Category, error)
	FindAllCategories(ctx context.Context) ([]*domain.Category, error)
}

// CategoryRepositoryWriter is an interface that wraps all the writing operations for a category.
type CategoryRepositoryWriter interface {
	Create(ctx context.Context, c *domain.Category) error
	Update(ctx context.Context, c *domain.Category) error
	Delete(ctx context.Context, id domain.ID) error
}

// CategoryRepository is an interface that wraps all the reading and writing operations for a category.
type CategoryRepository interface {
	CategoryRepositoryReader
	CategoryRepositoryWriter
}

// CategoryService is an interface that wraps all the operations for a category.
type CategoryService interface {
	GetByID(ctx context.Context, id domain.ID) (*domain.Category, error)
	GetAll(ctx context.Context) ([]*domain.Category, error)
	Create(ctx context.Context, c *domain.Category) (*domain.Category, error)
	Update(ctx context.Context, c *domain.Category) (*domain.Category, error)
	Delete(ctx context.Context, id domain.ID) error
	Activate(ctx context.Context, id domain.ID) error
}
