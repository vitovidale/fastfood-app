package postgres

import (
	"context"
	"fmt"

	"github.com/vitovidale/fastfood-app/internal/adapter/driven/config"
	"github.com/vitovidale/fastfood-app/internal/core/domain"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	*gorm.DB
}

func CreateOrderTrackingNumberSequence(db *gorm.DB) error {
	return db.Exec(`
		CREATE SEQUENCE IF NOT EXISTS order_tracking_number_sequence
		START 1
		INCREMENT 1
		MINVALUE 1
		MAXVALUE 999
		CYCLE
	`).Error
}

func ResetOrderTrackingNumberSequence(db *gorm.DB) error {
	return db.Exec(`ALTER SEQUENCE order_tracking_number_sequence RESTART WITH 1`).Error
}

func New(ctx context.Context, config *config.DB) (*DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Host,
		config.Port,
		config.User,
		config.Password,
		config.Name,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// custom table definitions
	db.SetupJoinTable(&domain.Order{}, "Products", &domain.OrderProduct{})

	// run GORMs auto migration process
	db.AutoMigrate(
		&domain.Category{},
		&domain.Product{},
		&domain.Customer{},
		&domain.Order{},
	)

	CreateOrderTrackingNumberSequence(db)

	return &DB{db}, nil
}
