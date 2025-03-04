package service

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/assert"
	"github.com/vitovidale/TECH-CHALLENGE/internal/core/domain"
	mock_port "github.com/vitovidale/TECH-CHALLENGE/internal/core/port/mock"
	"go.uber.org/mock/gomock"
)

type createCategoryTestedInput struct {
	category *domain.Category
}

type createCategoryTestedOutput struct {
	category *domain.Category
	err      error
}

func TestCategoryService_CreateCategory(t *testing.T) {
	ctx := context.Background()
	categoryID, _ := domain.ParseID(gofakeit.UUID())
	categoryName := gofakeit.ProductCategory()
	categoryCreatedAt := gofakeit.Date()
	now := gofakeit.Date()

	categoryInput := &domain.Category{
		ID:        categoryID,
		Name:      categoryName,
		CreatedAt: categoryCreatedAt,
		UpdatedAt: &now,
	}

	categoryOutput := &domain.Category{
		ID:        categoryID,
		Name:      categoryName,
		CreatedAt: categoryCreatedAt,
		UpdatedAt: &now,
	}

	testCases := []struct {
		title string
		mocks func(
			categoryRepository *mock_port.MockCategoryRepository,
		)
		input  createCategoryTestedInput
		output createCategoryTestedOutput
	}{
		{
			title: "Success",
			mocks: func(
				categoryRepository *mock_port.MockCategoryRepository,
			) {
				categoryRepository.EXPECT().Create(ctx, categoryInput).Return(nil)
			},
			input: createCategoryTestedInput{
				category: categoryInput,
			},
			output: createCategoryTestedOutput{
				category: categoryOutput,
				err:      nil,
			},
		},
		{
			title: "Error",
			mocks: func(
				categoryRepository *mock_port.MockCategoryRepository,
			) {
				categoryRepository.EXPECT().Create(ctx, categoryInput).Return(domain.ErrorInternal)
			},
			input: createCategoryTestedInput{
				category: categoryInput,
			},
			output: createCategoryTestedOutput{
				category: nil,
				err:      domain.ErrorInternal,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.title, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			categoryRepository := mock_port.NewMockCategoryRepository(ctrl)
			tc.mocks(categoryRepository)

			service := NewCategoryService(categoryRepository)

			category, err := service.Create(ctx, tc.input.category)

			assert.Equal(t, tc.output.category, category)
			assert.Equal(t, tc.output.err, err)
		})
	}
}
