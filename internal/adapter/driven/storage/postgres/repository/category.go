package repository

import (
	"context"
	"time"

	"github.com/vitovidale/fastfood-app/internal/adapter/driven/storage/postgres"
	"github.com/vitovidale/fastfood-app/internal/core/domain"
)

type CategoryRepository struct {
	db *postgres.DB
}

func NewCategoryRepository(db *postgres.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (r *CategoryRepository) Create(ctx context.Context, c *domain.Category) error {
	result := r.db.WithContext(ctx).Create(&c)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *CategoryRepository) Update(ctx context.Context, c *domain.Category) error {
	result := r.db.WithContext(ctx).
		Where("id = ?", c.ID).
		Updates(&c)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *CategoryRepository) Delete(ctx context.Context, id domain.ID) error {
	var c domain.Category
	result := r.db.WithContext(ctx).First(&c, "id = ?", id).Update("deleted_at", time.Now())
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// Read operations on category
func (r *CategoryRepository) FindCategoryByID(ctx context.Context, id domain.ID) (*domain.Category, error) {
	var c domain.Category
	result := r.db.WithContext(ctx).First(&c, "id = ? AND deleted_at IS NULL", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &c, nil
}

func (r *CategoryRepository) FindAllCategories(ctx context.Context) ([]*domain.Category, error) {
	var categories []*domain.Category
	result := r.db.WithContext(ctx).Where("deleted_at IS NULL").Find(&categories)
	if result.Error != nil {
		return nil, result.Error
	}
	return categories, nil
}
