package domain

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

// TODO change the ID from CPF to UUID
type Customer struct {
	ID        uint64     `gorm:"type:bigint"`
	FirstName string     `gorm:"size:100;not null"`
	LastName  string     `gorm:"size:100;not null"`
	Email     string     `gorm:"size:255;unique;not null"`
	Password  string     `gorm:"size:64;not null"`
	CreatedAt time.Time  `gorm:"autoCreateTime;not null"`
	UpdatedAt *time.Time `gorm:"autoUpdateTime"`
	DeletedAt *time.Time
}

func (p *Customer) IsActive() bool {
	return p.DeletedAt.IsZero()
}

func (p *Customer) Deactivate() error {
	if p.DeletedAt != nil {
		return ErrorCustomerAlreadyInactive
	}
	deletedAt := time.Now()
	p.DeletedAt = &deletedAt
	return nil
}

func (p *Customer) Activate() error {
	if p.DeletedAt == nil {
		return ErrorCustomerAlreadyActive
	}
	p.DeletedAt = nil
	return nil
}

func (p *Customer) Authenticate(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(p.Password), []byte(password))
}
