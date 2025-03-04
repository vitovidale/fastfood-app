package http

import (
	"sync/atomic"

	"github.com/gin-gonic/gin"
	"github.com/vitovidale/TECH-CHALLENGE/internal/adapter/driver/handler/http/response"
	"github.com/vitovidale/TECH-CHALLENGE/internal/core/domain"
)

type HealthHandler struct {
}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

var (
	isAppReady   atomic.Bool
	isAppStarted atomic.Bool
)

func SetReady(status bool) {
	isAppReady.Store(status)
}

func IsReady() bool {
	return isAppReady.Load()
}

func SetStarted(status bool) {
	isAppStarted.Store(status)
}

func IsStarted() bool {
	return isAppStarted.Load()
}

// Start godoc
//
//	@Summary		Start the app
//	@Description	Run a warmup request
//	@Tags			Health
//	@Produce		json
//	@Success		200	{object}	bool					"App started"
//	@Failure		500	{object}	response.ErrorResponse	"Internal server error"
//	@Router			/health/start [get]
func (h *HealthHandler) Start(ctx *gin.Context) {
	if IsStarted() {
		response.HandleSuccess(ctx, true)
	} else {
		response.HandleError(ctx, domain.ErrorAppNotStarted)
	}
}

// Readiness godoc
//
//	@Summary		Readiness check
//	@Description	Checks if the app is ready to serve requests
//	@Tags			Health
//	@Produce		json
//	@Success		200	{object}	bool					"App ready"
//	@Failure		500	{object}	response.ErrorResponse	"Internal server error"
//	@Router			/health/readiness [get]
func (h *HealthHandler) Readiness(ctx *gin.Context) {
	if IsReady() {
		response.HandleSuccess(ctx, true)
	} else {
		response.HandleError(ctx, domain.ErrorAppNotReady)
	}
}

// Liveness godoc
//
//	@Summary		Liveness check
//	@Description	Checks if the app is alive
//	@Tags			Health
//	@Produce		json
//	@Success		200	{object}	bool					"App alive"
//	@Failure		500	{object}	response.ErrorResponse	"Internal server error"
//	@Router			/health/liveness [get]
func (h *HealthHandler) Liveness(ctx *gin.Context) {
	response.HandleSuccess(ctx, true)
}
