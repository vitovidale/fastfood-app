package domain

import (
	"time"
)

type Product struct {
	ID          ID      `gorm:"size:36"`
	Name        string  `gorm:"size:60;not null"`
	Description string  `gorm:"size:100"`
	Price       float64 `gorm:"not null"`
	CategoryID  ID      `gorm:"size:36;not null"`
	Category    *Category
	CreatedAt   time.Time  `gorm:"autoCreateTime;not null"`
	UpdatedAt   *time.Time `gorm:"autoUpdateTime"`
	DeletedAt   *time.Time
}

func NewProductWithID(id ID, name string, description string, price float64, categoryID ID) *Product {
	product := NewProduct(name, description, price, categoryID)
	product.ID = id
	return product
}

func NewProduct(name string, description string, price float64, categoryID ID) *Product {
	return &Product{
		ID:          NewID(),
		Name:        name,
		Description: description,
		Price:       price,
		CategoryID:  categoryID,
	}
}

func (p *Product) IsActive() bool {
	return p.DeletedAt.IsZero()
}

func (p *Product) GetPrice() float64 {
	return p.Price
}

func (p *Product) Inactivate() error {
	if p.DeletedAt != nil {
		return ErrorProductAlreadyInactive
	}
	deletedAt := time.Now()
	p.DeletedAt = &deletedAt
	return nil
}

func (p *Product) Activate() error {
	if p.DeletedAt == nil {
		return ErrorProductAlreadyActive
	}
	p.DeletedAt = nil
	return nil
}
