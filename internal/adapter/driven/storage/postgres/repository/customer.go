package repository

import (
	"context"

	"github.com/vitovidale/TECH-CHALLENGE/internal/adapter/driven/storage/postgres"
	"github.com/vitovidale/TECH-CHALLENGE/internal/core/domain"
)

type CustomerRepository struct {
	db *postgres.DB
}

func NewCustomerRepository(db *postgres.DB) *CustomerRepository {
	return &CustomerRepository{db: db}
}

func (r *CustomerRepository) Create(ctx context.Context, c *domain.Customer) error {
	result := r.db.WithContext(ctx).Create(&c)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *CustomerRepository) FindByID(ctx context.Context, id uint64) (*domain.Customer, error) {
	c := &domain.Customer{}
	result := r.db.WithContext(ctx).
		Where("deleted_at IS NULL").
		First(&c, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return c, nil
}

func (r *CustomerRepository) FindByKeys(ctx context.Context, id uint64, email string) (*domain.Customer, error) {
	c := &domain.Customer{}
	result := r.db.WithContext(ctx).
		Where("id = ? OR email = ?", id, email).
		Where("deleted_at IS NULL").
		First(&c)

	if result.Error != nil {
		return nil, result.Error
	}
	return c, nil
}

func (r *CustomerRepository) Patch(ctx context.Context, id uint64, data *domain.Customer) error {
	result := r.db.WithContext(ctx).
		Model(&domain.Customer{}).
		Where("id = ?", id).
		Updates(data)

	if result.Error != nil {
		return result.Error
	}
	return nil
}
