package http

import (
	"github.com/gin-gonic/gin"
	"github.com/vitovidale/TECH-CHALLENGE/internal/adapter/driver/handler/http/request"
	"github.com/vitovidale/TECH-CHALLENGE/internal/adapter/driver/handler/http/response"
	"github.com/vitovidale/TECH-CHALLENGE/internal/core/domain"
	"github.com/vitovidale/TECH-CHALLENGE/internal/core/service"
)

type OrderHandler struct {
	service *service.OrderService
}

func NewOrderHandler(service *service.OrderService) *OrderHandler {
	return &OrderHandler{service: service}
}

// GetByID godoc
//
//	@Summary		Get an order
//	@Description	Returns an order based on its ID
//	@Tags			Orders
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string					true	"Order ID"
//	@Success		200	{object}	response.OrderResponse	"Order found"
//	@Failure		400	{object}	response.ErrorResponse	"Bad Request error"
//	@Failure		404	{object}	response.ErrorResponse	"Not found error"
//	@Router			/orders/{id} [get]
func (h *OrderHandler) GetByID(ctx *gin.Context) {
	var req request.GetOrderRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		response.HandleError(ctx, err)
		return
	}

	id, _ := domain.ParseID(req.ID)
	order, err := h.service.GetNestedByID(ctx, id)
	if err != nil {
		response.HandleError(ctx, err)
		return
	}

	response.HandleSuccess(ctx, order)
}

// GetByCustomerID godoc
//
//	@Summary		Get an order by its customer ID
//	@Description	Gets an active order by its customer ID, used for incremental orders, allowing the addition and removal of products
//	@Tags			Orders
//	@Accept			json
//	@Produce		json
//	@Param			customerId	path		uint64					true	"Customer ID"
//	@Success		200			{object}	response.OrderResponse	"Order found"
//	@Failure		400			{object}	response.ErrorResponse	"Bad Request error"
//	@Failure		404			{object}	response.ErrorResponse	"Not found error"
//	@Router			/orders/customer/{customerId} [get]
func (h *OrderHandler) GetByCustomerID(ctx *gin.Context) {
	var req request.GetOrderByCustomerRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		response.HandleError(ctx, err)
		return
	}

	order, err := h.service.GetByCustomer(ctx, req.CustomerID)
	if err != nil {
		response.HandleError(ctx, err)
		return
	}

	o, _ := h.service.GetNestedByID(ctx, order.ID)
	response.HandleSuccess(ctx, o)
}

// AddProduct godoc
//
//	@Summary		Add a product to an order
//	@Description	Adds a product to an existing **active** order, based on its CustomerID. If the order doesn't exist, it will be created.
//	@Tags			Orders
//	@Accept			json
//	@Produce		json
//	@Param			AddProductRequest	body		request.AddProductRequest	true	"Add product request"
//	@Success		200					{object}	response.OrderResponse		"Product added"
//	@Failure		400					{object}	response.ErrorResponse		"Bad Request error"
//	@Failure		404					{object}	response.ErrorResponse		"Not found error"
//	@Failure		500					{object}	response.ErrorResponse		"Internal server error"
//	@Router			/orders/products [post]
func (h *OrderHandler) AddProduct(ctx *gin.Context) {
	var req request.AddProductRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.HandleError(ctx, err)
		return
	}

	id, err := h.service.Create(ctx, req.CustomerID, []domain.OrderProduct{{
		ProductID: domain.ParseIDOrNil(req.ProductID),
		Quantity:  req.Quantity,
		Notes:     req.Notes,
	}})

	if err != nil {
		response.HandleError(ctx, err)
		return
	}

	o, _ := h.service.GetNestedByID(ctx, *id)
	response.HandleSuccess(ctx, o)
}

// RemoveProduct godoc
//
//	@Summary		Remove a product from an order
//	@Description	Removes a product from an existing **active** order, based on its OrderID and OrderProductID.
//	@Tags			Orders
//	@Accept			json
//	@Produce		json
//	@Param			orderId			path		string					true	"Order ID"
//	@Param			orderProductId	path		string					true	"Product ID"
//	@Success		200				{object}	response.OrderResponse	"Order found"
//	@Failure		400				{object}	response.ErrorResponse	"Bad Request error"
//	@Failure		404				{object}	response.ErrorResponse	"Not found error"
//	@Failure		500				{object}	response.ErrorResponse	"Internal server error"
//	@Router			/orders/{orderId}/products/{orderProductId} [delete]
func (h *OrderHandler) RemoveProduct(ctx *gin.Context) {
	var req request.RemoveProductRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		response.HandleError(ctx, err)
		return
	}

	orderId := domain.ParseIDOrNil(req.OrderID)

	err := h.service.RemoveProduct(ctx, orderId, domain.ParseIDOrNil(req.OrderProductID))

	if err != nil {
		response.HandleError(ctx, err)
		return
	}

	o, _ := h.service.GetNestedByID(ctx, orderId)
	response.HandleSuccess(ctx, o)
}

// Pay godoc
//
//	@Summary		Pays an order
//	@Description	Marked the order as paid, setting its status to `confirmed`, based on the order ID
//	@Tags			Orders
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string					true	"Order ID"
//	@Success		200	{object}	response.OrderResponse	"Order found"
//	@Failure		400	{object}	response.ErrorResponse	"Bad Request error"
//	@Failure		404	{object}	response.ErrorResponse	"Not found error"
//	@Failure		500	{object}	response.ErrorResponse	"Internal server error"
//	@Router			/orders/{id}/pay [patch]
func (h *OrderHandler) Pay(ctx *gin.Context) {
	id, _ := domain.ParseID(ctx.Param("id"))
	err := h.service.Pay(ctx, id)

	if err != nil {
		response.HandleError(ctx, err)
		return
	}

	o, _ := h.service.GetNestedByID(ctx, id)
	response.HandleSuccess(ctx, o)
}

// Prepare godoc
//
//	@Summary		Prepare an order
//	@Description	Signal that the kitchen has started preparing the order, changing its status to `started`, based on the order ID
//	@Tags			Orders
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string					true	"Order ID"
//	@Success		200	{object}	response.OrderResponse	"Order found"
//	@Failure		400	{object}	response.ErrorResponse	"Bad Request error"
//	@Failure		404	{object}	response.ErrorResponse	"Not found error"
//	@Failure		500	{object}	response.ErrorResponse	"Internal server error"
//	@Router			/orders/{id}/prepare [patch]
func (h *OrderHandler) Prepare(ctx *gin.Context) {
	id, _ := domain.ParseID(ctx.Param("id"))
	err := h.service.Prepare(ctx, id)

	if err != nil {
		response.HandleError(ctx, err)
		return
	}

	o, _ := h.service.GetNestedByID(ctx, id)
	response.HandleSuccess(ctx, o)
}

// Complete godoc
//
//	@Summary		Complete an order
//	@Description	Changes the state of the order to `done`, where the customer needs to **take out** the order from the counter, based on the order ID
//	@Tags			Orders
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string					true	"Order ID"
//	@Success		200	{object}	response.OrderResponse	"Order found"
//	@Failure		400	{object}	response.ErrorResponse	"Bad Request error"
//	@Failure		404	{object}	response.ErrorResponse	"Not found error"
//	@Failure		500	{object}	response.ErrorResponse	"Internal server error"
//	@Router			/orders/{id}/complete [patch]
func (h *OrderHandler) Complete(ctx *gin.Context) {
	id, _ := domain.ParseID(ctx.Param("id"))
	err := h.service.Complete(ctx, id)

	if err != nil {
		response.HandleError(ctx, err)
		return
	}

	o, _ := h.service.GetNestedByID(ctx, id)
	response.HandleSuccess(ctx, o)
}

// List godoc
//
//	@Summary		List orders
//	@Description	Returns all orders, sorted by status ASC, and ignores inactive or completed orders
//	@Tags			Orders
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	[]response.OrderResponse	"Order found"
//	@Failure		400	{object}	response.ErrorResponse		"Bad Request error"
//	@Failure		404	{object}	response.ErrorResponse		"Not found error"
//	@Failure		500	{object}	response.ErrorResponse		"Internal server error"
//	@Router			/orders [get]
func (h *OrderHandler) List(ctx *gin.Context) {
	orders, err := h.service.List(ctx)

	if err != nil {
		response.HandleError(ctx, err)
		return
	}

	response.HandleSuccess(ctx, orders)
}

// Create godoc
//
//	@Summary		Create an order
//	@Description	Creates an order for a customer with a list of products in a single request
//	@Tags			Orders
//	@Accept			json
//	@Produce		json
//	@Param			CreateOrderRequest	body		request.CreateOrderRequest	true	"Payload with customer ID and products"
//	@Success		200					{object}	response.OrderResponse		"Order created"
//	@Failure		400					{object}	response.ErrorResponse		"Bad Request error"
//	@Failure		404					{object}	response.ErrorResponse		"Not found error"
//	@Failure		500					{object}	response.ErrorResponse		"Internal server error"
//	@Router			/orders [post]
func (h *OrderHandler) Create(ctx *gin.Context) {
	var req request.CreateOrderRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.HandleError(ctx, err)
		return
	}

	var products []domain.OrderProduct

	for _, p := range req.Products {
		products = append(products, domain.OrderProduct{
			ProductID: domain.ParseIDOrNil(p.ProductID),
			Quantity:  p.Quantity,
			Notes:     p.Notes,
		})
	}

	id, err := h.service.Create(ctx, req.CustomerID, products)

	if err != nil {
		response.HandleError(ctx, err)
		return
	}

	o, _ := h.service.GetNestedByID(ctx, *id)
	response.HandleSuccess(ctx, o)
}

// GetStatus godoc
//
//	@Summary		Get the order status
//	@Description	Returns the order status by its ID
//	@Tags			Orders
//	@Produce		json
//	@Param			id	path		string						true	"Order ID"
//	@Success		200	{object}	response.DefaultResponse	"Order status"
//	@Failure		400	{object}	response.ErrorResponse		"Bad Request error"
//	@Failure		404	{object}	response.ErrorResponse		"Not found error"
//	@Failure		500	{object}	response.ErrorResponse		"Internal server error"
//	@Router			/orders/{id}/status [get]
func (h *OrderHandler) GetStatus(ctx *gin.Context) {
	var req request.GetOrderRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		response.HandleError(ctx, err)
		return
	}

	order, err := h.service.GetByID(ctx, domain.ParseIDOrNil(req.ID))

	if err != nil {
		response.HandleError(ctx, err)
		return
	}

	response.HandleSuccess(ctx, order.Status)
}
