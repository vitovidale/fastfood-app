package domain

import (
	"time"
)

type Category struct {
	ID        ID         `gorm:"size:36"`
	Name      string     `gorm:"size:60;not null"`
	CreatedAt time.Time  `gorm:"autoCreateTime;not null"`
	UpdatedAt *time.Time `gorm:"autoUpdateTime"`
	DeletedAt *time.Time
}

func NewCategoryWithID(id ID, name string) *Category {
	category := NewCategory(name)
	category.ID = id
	return category
}

func NewCategory(name string) *Category {
	return &Category{
		ID:   NewID(),
		Name: name,
	}
}

func (c *Category) IsActive() bool {
	return c.DeletedAt.IsZero()
}

func (c *Category) Inactivate() error {
	if c.DeletedAt != nil {
		return ErrorCategoryAlreadyInactive
	}
	deletedAt := time.Now()
	c.DeletedAt = &deletedAt
	return nil
}

func (c *Category) Activate() error {
	if c.DeletedAt == nil {
		return ErrorCategoryAlreadyActive
	}
	c.DeletedAt = nil
	return nil
}
