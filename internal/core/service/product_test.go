package service

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/vitovidale/TECH-CHALLENGE/internal/core/domain"
	mock_port "github.com/vitovidale/TECH-CHALLENGE/internal/core/port/mock"
	"go.uber.org/mock/gomock"
)

type createProductTestedInput struct {
	product *domain.Product
}

type createProductTestedOutput struct {
	product *domain.Product
	err     error
}

func TestProductService_CreateProduct(t *testing.T) {
	ctx := context.Background()
	categoryID, _ := domain.ParseID(gofakeit.UUID())
	categoryName := gofakeit.ProductCategory()
	categoyCreatedAt := gofakeit.Date()
	now := gofakeit.Date()

	category := &domain.Category{
		ID:        categoryID,
		Name:      categoryName,
		CreatedAt: categoyCreatedAt,
		UpdatedAt: &now,
	}

	productID, _ := domain.ParseID(gofakeit.UUID())
	productName := gofakeit.ProductName()
	productPrice := gofakeit.Price(10, 100)
	productCreatedAt := gofakeit.Date()

	productInput := &domain.Product{
		ID:         productID,
		Name:       productName,
		Price:      productPrice,
		CategoryID: categoryID,
		CreatedAt:  productCreatedAt,
		UpdatedAt:  &now,
	}

	productOutput := &domain.Product{
		ID:         productID,
		Name:       productName,
		Price:      productPrice,
		CategoryID: categoryID,
		Category:   category,
		CreatedAt:  productCreatedAt,
		UpdatedAt:  &now,
	}

	testCases := []struct {
		title string
		mocks func(
			productRepository *mock_port.MockProductRepository,
			categoryRepository *mock_port.MockCategoryRepository,
		)
		input  createProductTestedInput
		output createProductTestedOutput
	}{
		{
			title: "Create product successfully",
			mocks: func(
				productRepository *mock_port.MockProductRepository,
				categoryRepository *mock_port.MockCategoryRepository,
			) {
				categoryRepository.EXPECT().FindCategoryByID(gomock.Any(), gomock.Eq(categoryID)).Return(category, nil)
				productRepository.EXPECT().Create(gomock.Any(), gomock.Eq(productInput)).Return(nil)
				productRepository.EXPECT().FindByID(gomock.Any(), gomock.Eq(productID)).Return(productOutput, nil)
			},
			input: createProductTestedInput{
				product: productInput,
			},
			output: createProductTestedOutput{
				product: productOutput,
				err:     nil,
			},
		},
		{
			title: "Category not found",
			mocks: func(
				productRepository *mock_port.MockProductRepository,
				categoryRepository *mock_port.MockCategoryRepository,
			) {
				categoryRepository.EXPECT().FindCategoryByID(ctx, categoryID).Return(nil, domain.ErrorDataNotFound)
			},
			input: createProductTestedInput{
				product: productInput,
			},
			output: createProductTestedOutput{
				product: nil,
				err:     domain.ErrorCategoryNotFound,
			},
		},
		{
			title: "Internal error",
			mocks: func(
				productRepository *mock_port.MockProductRepository,
				categoryRepository *mock_port.MockCategoryRepository,
			) {
				categoryRepository.EXPECT().FindCategoryByID(ctx, categoryID).Return(nil, domain.ErrorInternal)
			},
			input: createProductTestedInput{
				product: productInput,
			},
			output: createProductTestedOutput{
				product: nil,
				err:     domain.ErrorInternal,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.title, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			productRepository := mock_port.NewMockProductRepository(ctrl)
			categoryRepository := mock_port.NewMockCategoryRepository(ctrl)

			tc.mocks(productRepository, categoryRepository)

			// service := NewProductService(categoryRepository, productRepository)
			// product, err := service.Create(ctx, tc.input.product)
			// if err != nil {
			// 	assert.Equal(t, tc.output.err, err, "Error mismatch")
			// 	return
			// }
			// product, err = service.GetByID(ctx, productID)
			// if err != nil {
			// 	assert.Equal(t, tc.output.err, err, "Error mismatch")
			// 	return
			// }
			// assert.Equal(t, tc.output.product, product, "Product mismatch")
			// assert.Equal(t, tc.output.err, err, "Error mismatch")
		})
	}
}
