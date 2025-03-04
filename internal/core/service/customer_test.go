package service

import (
	"context"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/vitovidale/TECH-CHALLENGE/internal/core/domain"
	mock_port "github.com/vitovidale/TECH-CHALLENGE/internal/core/port/mock"
	"go.uber.org/mock/gomock"
)

type CreateCustomerInput struct {
	customer *domain.Customer
}

type CreateCustomerOutput struct {
	customer *domain.Customer
	err      error
}

func TestCustomerService(t *testing.T) {
	ctx := context.Background()
	id := uint64(0)
	firstName := gofakeit.FirstName()
	email := gofakeit.Email()
	password := gofakeit.Password(true, true, true, true, false, 32)
	lastName := gofakeit.LastName()
	createdAt := time.Now()
	updatedAt := time.Now()

	customerInput := &domain.Customer{
		ID:        id,
		FirstName: firstName,
		Email:     email,
		Password:  password,
		LastName:  lastName,
		CreatedAt: createdAt,
		UpdatedAt: &updatedAt,
	}

	customerOutput := &domain.Customer{
		ID:        id,
		FirstName: firstName,
		Email:     email,
		Password:  password,
		LastName:  lastName,
		CreatedAt: createdAt,
		UpdatedAt: &updatedAt,
	}

	testCases := []struct {
		title string
		mocks func(
			customerRepository *mock_port.MockCustomerRepository,
		)
		input  CreateCustomerInput
		output CreateCustomerOutput
	}{
		{
			title: "Create a customer",
			mocks: func(
				customerRepository *mock_port.MockCustomerRepository,
			) {
				customerRepository.EXPECT().Create(ctx, customerInput).Return(nil)
			},
			input: CreateCustomerInput{
				customer: customerInput,
			},
			output: CreateCustomerOutput{
				customer: customerOutput,
				err:      nil,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.title, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			customerRepository := mock_port.NewMockCustomerRepository(ctrl)
			tc.mocks(customerRepository)

			// TODO regen customer mocks
			// customerService := NewCustomerService(customerRepository)
			// customer, err := customerService.Create(ctx, tc.input.customer)

			// assert.Equal(t, tc.output.customer, customer)
			// assert.Equal(t, tc.output.err, err)
		})
	}
}
