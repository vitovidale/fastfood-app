package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/vitovidale/fastfood-app/internal/adapter/driven/config"
	"github.com/vitovidale/fastfood-app/internal/adapter/driven/storage/postgres"
	"github.com/vitovidale/fastfood-app/internal/adapter/driven/storage/postgres/repository"
	"github.com/vitovidale/fastfood-app/internal/adapter/driver/handler/http"
	"github.com/vitovidale/fastfood-app/internal/adapter/logger"
	"github.com/vitovidale/fastfood-app/internal/core/service"
)

//	@title			Tech Challenge API
//	@version		1.0
//	@description	This is an API for Tech Challenge from FIAP.
//
//	@license.name	MIT
//	@license.url	https://github.com/vitovidale/fastfood-app/blob/main/LICENSE
//
//	@host			127.0.0.1:8080
//	@BasePath		/v1
//	@schemes		http https

func main() {
	config, err := config.New()
	if err != nil {
		slog.Error("Error when reading configuration on main")
		os.Exit(1)
	}

	logger.Set(config.App)
	slog.Info("Starting the application", "app", config.App.Name, "env", config.App.Env)

	ctx := context.Background()
	db, err := postgres.New(ctx, config.DB)
	if err != nil {
		slog.Error("Error connecting to database")
		os.Exit(1)
	}
	slog.Info("Successfully connected to the database", "db", config.DB.Connection)

	// Dependency Injection initialization
	// Category
	categoryRepo := repository.NewCategoryRepository(db)
	categoryService := service.NewCategoryService(categoryRepo)
	categoryHandler := http.NewCategoryHandler(categoryService)

	// Product
	productRepo := repository.NewProductRepository(db)
	productService := service.NewProductService(categoryRepo, productRepo)
	productHandler := http.NewProductHandler(productService)

	// Customer
	customerRepo := repository.NewCustomerRepository(db)
	customerService := service.NewCustomerService(customerRepo)
	customerHandler := http.NewCustomerHandler(customerService)

	// Order
	orderRepo := repository.NewOrderRepository(db)
	orderService := service.NewOrderService(orderRepo, productRepo, customerRepo)
	orderHandler := http.NewOrderHandler(orderService)

	// Health
	healthHandler := http.NewHealthHandler()

	// Router initialization
	router, err := http.NewRouter(
		config.HTTP,
		*productHandler,
		*categoryHandler,
		*customerHandler,
		*orderHandler,
		*healthHandler,
	)

	if err != nil {
		slog.Error("Error initializing router")
		os.Exit(1)
	}
	slog.Info("Started the HTTP Router")

	listenAddress := fmt.Sprintf("%s:%s", config.HTTP.Url, config.HTTP.Port)

	go func() {
		err = router.Start(listenAddress)
		if err != nil {
			slog.Error("Error running server")
			os.Exit(1)
		}
	}()

	slog.Info("Starting the HTTP server", "listen_address", listenAddress)

	http.SetReady(true)
	http.SetStarted(true)

	select {}
}
