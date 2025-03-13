package service

import (
	"context"

	"github.com/vitovidale/fastfood-app/internal/core/domain"
	"github.com/vitovidale/fastfood-app/internal/core/port"
)

type CategoryService struct {
	categoryRepository port.CategoryRepository
}

func NewCategoryService(categoryRepository port.CategoryRepository) *CategoryService {
	return &CategoryService{categoryRepository: categoryRepository}
}

func (s *CategoryService) GetByID(ctx context.Context, id domain.ID) (*domain.Category, error) {
	c, err := s.categoryRepository.FindCategoryByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (s *CategoryService) GetAll(ctx context.Context) ([]*domain.Category, error) {
	c, err := s.categoryRepository.FindAllCategories(ctx)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (s *CategoryService) Create(ctx context.Context, c *domain.Category) (*domain.Category, error) {
	err := s.categoryRepository.Create(ctx, c)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (s *CategoryService) Update(ctx context.Context, c *domain.Category) (*domain.Category, error) {
	err := s.categoryRepository.Update(ctx, c)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (s *CategoryService) Delete(ctx context.Context, id domain.ID) error {
	c, err := s.categoryRepository.FindCategoryByID(ctx, id)
	if err != nil {
		return err
	}
	err = c.Inactivate()
	if err != nil {
		return err
	}
	err = s.categoryRepository.Update(ctx, c)
	if err != nil {
		return err
	}
	return nil
}

func (s *CategoryService) Activate(ctx context.Context, id domain.ID) error {
	c, err := s.categoryRepository.FindCategoryByID(ctx, id)
	if err != nil {
		return err
	}

	err = c.Activate()
	if err != nil {
		return err
	}

	err = s.categoryRepository.Update(ctx, c)
	if err != nil {
		return err
	}

	return nil
}
