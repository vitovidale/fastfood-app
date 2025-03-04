package http

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vitovidale/TECH-CHALLENGE/internal/adapter/driver/handler/http/request"
	"github.com/vitovidale/TECH-CHALLENGE/internal/adapter/driver/handler/http/response"
	"github.com/vitovidale/TECH-CHALLENGE/internal/core/domain"
	"github.com/vitovidale/TECH-CHALLENGE/internal/core/port"
)

type CustomerHandler struct {
	service port.CustomerService
}

func NewCustomerHandler(service port.CustomerService) *CustomerHandler {
	return &CustomerHandler{service: service}
}

// Create godoc
//
//	@Summary		Create a new customer
//	@Description	Creates a new customer with first name, last name, email and password
//	@Tags			Customers
//	@Accept			json
//	@Produce		json
//	@Param			CreateCustomerRequest	body		request.CreateCustomerRequest	true	"Create customer request"
//	@Success		200						{object}	response.CustomerResponse		"Customer created"
//	@Failure		400						{object}	response.ErrorResponse			"Bad Request error"
//	@Router			/customers [post]
func (h *CustomerHandler) Create(ctx *gin.Context) {
	var req request.CreateCustomerRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.HandleError(ctx, err)
		return
	}

	customer := &domain.Customer{
		ID:        req.ID,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Password:  req.Password,
	}

	customer, err := h.service.Create(ctx, customer)

	if err != nil {
		response.HandleError(ctx, err)
		return
	}

	response.HandleSuccess(ctx, response.NewCustomerResponse(customer))
}

// GetByID godoc
//
//	@Summary		Get a single customer
//	@Description	Returns a single customer by its ID
//	@Tags			Customers
//	@Accept			json
//	@Produce		json
//	@Param			id	path		uint64						true	"Customer ID"
//	@Success		200	{object}	response.CustomerResponse	"Customer found"
//	@Failure		400	{object}	response.ErrorResponse		"Bad Request error"
//	@Failure		404	{object}	response.ErrorResponse		"Not found error"
//	@Router			/customers/{id} [get]
func (h *CustomerHandler) GetByID(c *gin.Context) {
	var req request.GetCustomerByIDRequest
	if err := c.ShouldBindUri(&req); err != nil {
		response.HandleError(c, err)
		return
	}

	customer, err := h.service.GetByID(c, req.ID)
	if err != nil {
		response.HandleError(c, err)
		return
	}

	response.HandleSuccess(c, response.NewCustomerResponse(customer))
}

// Auth godoc
//
//	@Summary		Authenticate a customer
//	@Description	Authenticates a customer with email or id, and password
//	@Tags			Customers
//	@Accept			json
//	@Produce		json
//	@Param			AuthCustomerRequest	body		request.AuthCustomerRequest	true	"Authenticate customer request"
//	@Success		200					{object}	response.CustomerResponse	"Customer authenticated"
//	@Failure		400					{object}	response.ErrorResponse		"Bad Request error"
//	@Router			/customers/auth [post]
func (h *CustomerHandler) Auth(ctx *gin.Context) {
	var req request.AuthCustomerRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.HandleError(ctx, err)
		return
	}

	customer, err := h.service.Authenticate(ctx,
		&domain.Customer{
			ID:       req.ID,
			Email:    req.Email,
			Password: req.Password,
		},
	)

	if err != nil {
		response.HandleError(ctx, err)
		return
	}

	response.HandleSuccess(ctx, response.NewCustomerResponse(customer))
}

// Update godoc
//
//	@Summary		Update a customer
//	@Description	Updates a single customer based on its ID
//	@Tags			Customers
//	@Accept			json
//	@Produce		json
//	@Param			id						path		uint64							true	"Customer ID"
//	@Param			UpdateCustomerRequest	body		request.UpdateCustomerRequest	true	"Update customer request"
//	@Success		200						{object}	response.CustomerResponse		"Customer updated"
//	@Failure		400						{object}	response.ErrorResponse			"Bad Request error"
//	@Failure		404						{object}	response.ErrorResponse			"Not found error"
//	@Router			/customers/{id} [put]
func (h *CustomerHandler) Update(ctx *gin.Context) {
	req := domain.Customer{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.HandleError(ctx, err)
		return
	}

	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		response.HandleError(ctx, err)
		return
	}

	req.ID = id

	c, err := h.service.Update(ctx, &req)
	if err != nil {
		response.HandleError(ctx, err)
		return
	}

	response.HandleSuccess(ctx, response.NewCustomerResponse(c))
}

// Delete godoc
//
//	@Summary		Deletes a customer
//	@Description	Deletes a customer by its ID
//	@Tags			Customers
//	@Accept			json
//	@Produce		json
//	@Param			id	path		uint64					true	"Customer ID"
//	@Success		200	{boolean}	bool					"Customer deleted"
//	@Failure		400	{object}	response.ErrorResponse	"Bad Request error"
//	@Router			/customers/{id} [delete]
func (h *CustomerHandler) Delete(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		response.HandleError(ctx, err)
		return
	}

	err = h.service.Delete(ctx, id)
	if err != nil {
		response.HandleError(ctx, err)
		return
	}
	response.HandleSuccess(ctx, true)
}
