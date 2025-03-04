package http

import (
	"log/slog"

	"github.com/davecgh/go-spew/spew"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	sloggin "github.com/samber/slog-gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/vitovidale/TECH-CHALLENGE/docs"
	"github.com/vitovidale/TECH-CHALLENGE/internal/adapter/driven/config"
)

// Router is a struct that wraps all the routes for the app.
type Router struct {
	*gin.Engine
}

// NewRouter creates a new HTTP Router.
func NewRouter(
	config *config.HTTP,
	productHandler ProductHandler,
	categoryHandler CategoryHandler,
	customerHandler CustomerHandler,
	orderHandler OrderHandler,
	healthHandler HealthHandler,
) (*Router, error) {
	if config.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()

	// cors setup
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true

	// spew setup
	spew.Config.Indent = "    "
	spew.Config.MaxDepth = 1

	// TODO fix this afterwards
	// allowedOrigins := config.AllowedOrigins
	// originsList := strings.Split(allowedOrigins, ",")
	// corsConfig.AllowOrigins = originsList

	router.Use(sloggin.New(slog.Default()), gin.Recovery(), cors.New(corsConfig))
	v1 := router.Group("/v1")
	{
		products := v1.Group("/products")
		{
			products.GET("/category/:id", productHandler.GetByCategory)
			products.GET("/:id", productHandler.GetByID)
			products.PUT("/:id", productHandler.Update)
			products.DELETE("/:id", productHandler.Delete)
			products.GET("", productHandler.GetAll)
			products.POST("", productHandler.Create)
		}

		categories := v1.Group("/categories")
		{
			categories.GET("/:id", categoryHandler.GetByID)
			categories.PUT("/:id", categoryHandler.Update)
			categories.DELETE("/:id", categoryHandler.Delete)
			categories.GET("", categoryHandler.GetAll)
			categories.POST("", categoryHandler.Create)
		}

		customers := v1.Group("/customers")
		{
			customers.GET("/:id", customerHandler.GetByID)
			customers.PUT("/:id", customerHandler.Update)
			customers.DELETE("/:id", customerHandler.Delete)
			customers.POST("/auth", customerHandler.Auth)
			customers.POST("", customerHandler.Create)
		}

		orders := v1.Group("/orders")
		{
			orders.GET("/:id/status", orderHandler.GetStatus)
			orders.PATCH("/:id/pay", orderHandler.Pay)
			orders.PATCH("/:id/prepare", orderHandler.Prepare)
			orders.PATCH("/:id/complete", orderHandler.Complete)
			orders.POST("/products", orderHandler.AddProduct)
			orders.DELETE("/:orderId/products/:orderProductId", orderHandler.RemoveProduct)
			orders.GET("/customer/:customerId", orderHandler.GetByCustomerID)
			orders.GET("/:id", orderHandler.GetByID)
			orders.GET("", orderHandler.List)
			orders.POST("", orderHandler.Create)
		}

		health := v1.Group("/health")
		{
			health.GET("/readiness", healthHandler.Readiness)
			health.GET("/liveness", healthHandler.Liveness)
			health.GET("/start", healthHandler.Start)
		}
	}

	// swagger setup
	docs.SwaggerInfo.BasePath = "/v1"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return &Router{router}, nil
}

func (r *Router) Start(listenAddress string) error {
	return r.Run(listenAddress)
}
