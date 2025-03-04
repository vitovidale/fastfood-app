package service

import (
	"context"
	"time"

	"github.com/vitovidale/TECH-CHALLENGE/internal/core/domain"
	"github.com/vitovidale/TECH-CHALLENGE/internal/core/port"
	"golang.org/x/crypto/bcrypt"
)

type CustomerService struct {
	customerRepository port.CustomerRepository
}

func NewCustomerService(customerRepository port.CustomerRepository) *CustomerService {
	return &CustomerService{customerRepository: customerRepository}
}

func (s *CustomerService) GetByID(ctx context.Context, id uint64) (*domain.Customer, error) {
	c, err := s.customerRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (s *CustomerService) GetByKeys(ctx context.Context, c *domain.Customer) (*domain.Customer, error) {
	c, err := s.customerRepository.FindByKeys(ctx, c.ID, c.Email)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (s *CustomerService) Create(ctx context.Context, c *domain.Customer) (*domain.Customer, error) {
	hash, _ := bcrypt.GenerateFromPassword([]byte(c.Password), bcrypt.DefaultCost)
	c.Password = string(hash)

	err := s.customerRepository.Create(ctx, c)

	if err != nil {
		return nil, err
	}

	return c, nil
}

func (s *CustomerService) Update(ctx context.Context, c *domain.Customer) (*domain.Customer, error) {
	data := domain.Customer{}

	if c.FirstName != "" {
		data.FirstName = c.FirstName
	}

	if c.LastName != "" {
		data.LastName = c.LastName
	}

	if c.Email != "" {
		data.Email = c.Email
	}

	err := s.customerRepository.Patch(ctx, c.ID, &data)
	if err != nil {
		return nil, err
	}

	return s.GetByID(ctx, c.ID)
}

func (s *CustomerService) Delete(ctx context.Context, id uint64) error {
	_, err := s.customerRepository.FindByID(ctx, id)

	if err != nil {
		return err
	}

	deletedAt := time.Now()
	err = s.customerRepository.Patch(ctx, id, &domain.Customer{DeletedAt: &deletedAt})

	if err != nil {
		return err
	}
	return nil
}

// func (s *CustomerService) Deactivate(ctx context.Context, id uint64) error {
// 	c, err := s.customerRepository.FindByID(ctx, id)

// 	if err != nil {
// 		return err
// 	}

// 	err = c.Deactivate()
// 	if err != nil {
// 		return err
// 	}

// 	err = s.customerRepository.Update(ctx, c)

// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (s *CustomerService) Activate(ctx context.Context, id uint64) error {
// 	c, err := s.customerRepository.FindByID(ctx, id)
// 	if err != nil {
// 		return err
// 	}

// 	err = c.Activate()

// 	if err != nil {
// 		return err
// 	}

// 	err = s.customerRepository.Update(ctx, c)

// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

func (s *CustomerService) Authenticate(ctx context.Context, c *domain.Customer) (*domain.Customer, error) {
	res, err := s.GetByKeys(ctx, c)
	if err != nil {
		return nil, err
	}

	err = res.Authenticate(c.Password)
	if err != nil {
		return nil, domain.ErrorCustomerWrongPassword
	}

	return res, nil
}
