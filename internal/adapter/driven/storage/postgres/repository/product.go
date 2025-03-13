package repository

import (
	"context"

	"github.com/vitovidale/fastfood-app/internal/adapter/driven/storage/postgres"
	"github.com/vitovidale/fastfood-app/internal/core/domain"
)

type ProductRepository struct {
	db *postgres.DB
}

func NewProductRepository(db *postgres.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) Create(ctx context.Context, p *domain.Product) error {
	result := r.db.WithContext(ctx).Create(&p)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *ProductRepository) Patch(ctx context.Context, id domain.ID, p *domain.Product) error {

	result := r.db.WithContext(ctx).
		Model(&domain.Product{ID: id}).
		Updates(p)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *ProductRepository) FindByID(ctx context.Context, id domain.ID) (*domain.Product, error) {
	p := &domain.Product{}

	result := r.db.WithContext(ctx).
		Where("deleted_at IS NULL").
		Preload("Category").
		First(&p, id)

	if result.Error != nil {
		return nil, result.Error
	}
	return p, nil
}

func (r *ProductRepository) FindAll(ctx context.Context) ([]*domain.Product, error) {
	var products []*domain.Product
	result := r.db.WithContext(ctx).
		Where("deleted_at IS NULL").
		Preload("Category").
		Find(&products)

	if result.Error != nil {
		return nil, result.Error
	}
	return products, nil
}

func (r *ProductRepository) FindByCategory(ctx context.Context, id domain.ID) ([]*domain.Product, error) {
	var products []*domain.Product
	result := r.db.WithContext(ctx).
		Where("category_id = ? AND deleted_at IS NULL", id).
		Preload("Category").
		Find(&products)

	if result.Error != nil {
		return nil, result.Error
	}
	return products, nil
}
