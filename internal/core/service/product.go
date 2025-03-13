package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/vitovidale/fastfood-app/internal/core/domain"
	"github.com/vitovidale/fastfood-app//internal/core/port"
)

type ProductService struct {
	categoryRepository port.CategoryRepository
	productRepository  port.ProductRepository
}

func NewProductService(categoryRepository port.CategoryRepository, productRepository port.ProductRepository) *ProductService {
	return &ProductService{categoryRepository: categoryRepository, productRepository: productRepository}
}

func (s *ProductService) GetByID(ctx context.Context, id domain.ID) (*domain.Product, error) {
	p, err := s.productRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (s *ProductService) GetAll(ctx context.Context) ([]*domain.Product, error) {
	p, err := s.productRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (s *ProductService) GetByCategory(ctx context.Context, id domain.ID) ([]*domain.Product, error) {
	p, err := s.productRepository.FindByCategory(ctx, id)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (s *ProductService) Create(ctx context.Context, p *domain.Product) (*domain.Product, error) {
	p = domain.NewProduct(p.Name, p.Description, p.Price, p.CategoryID)

	if err := s.findAndSetCategory(ctx, p); err != nil {
		return nil, err
	}

	err := s.productRepository.Create(ctx, p)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (s *ProductService) Update(ctx context.Context, p *domain.Product) (*domain.Product, error) {
	product := domain.Product{}

	if p.Name != "" {
		product.Name = p.Name
	}

	if p.Description != "" {
		product.Description = p.Description
	}

	if p.Price != 0 {
		product.Price = p.Price
	}

	if err := s.findAndSetCategory(ctx, p); err != nil {
		return nil, err
	}

	err := s.productRepository.Patch(ctx, p.ID, &product)
	if err != nil {
		return nil, err
	}

	return s.GetByID(ctx, p.ID)
}

func (s *ProductService) Delete(ctx context.Context, id domain.ID) error {
	p, err := s.productRepository.FindByID(ctx, id)
	if err != nil {
		return err
	}
	err = p.Inactivate()
	if err != nil {
		return err
	}
	err = s.productRepository.Patch(ctx, id, &domain.Product{DeletedAt: p.DeletedAt})
	if err != nil {
		return err
	}
	return nil
}

func (s *ProductService) findAndSetCategory(ctx context.Context, p *domain.Product) error {
	if p.CategoryID == uuid.Nil {
		return nil
	}

	category, err := s.categoryRepository.FindCategoryByID(ctx, p.CategoryID)

	if err != nil {
		if err.Error() == domain.ErrorDataNotFound.Error() {
			return domain.ErrorCategoryNotFound
		}
		return domain.ErrorInternal
	}
	p.Category = category
	return nil
}
