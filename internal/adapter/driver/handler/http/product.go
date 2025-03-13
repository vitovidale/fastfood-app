package http

import (
	"github.com/gin-gonic/gin"
	"github.com/vitovidale/fastfood-app/internal/adapter/driver/handler/http/request"
	"github.com/vitovidale/fastfood-app/internal/adapter/driver/handler/http/response"
	"github.com/vitovidale/fastfood-app/internal/core/domain"
	"github.com/vitovidale/fastfood-app/internal/core/port"
)

type ProductHandler struct {
	service port.ProductService
}

func NewProductHandler(service port.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

// Create godoc
//
//	@Summary		Create a new product
//	@Description	Creates a new product with name, price and descripton
//	@Tags			Products
//	@Accept			json
//	@Produce		json
//	@Param			CreateProductRequest	body		request.CreateProductRequest	true	"Create product request"
//	@Success		200						{object}	response.ProductResponse		"Product created"
//	@Failure		400						{object}	response.ErrorResponse			"Bad Request error"
//	@Failure		404						{object}	response.ErrorResponse			"Not found error"
//	@Failure		409						{object}	response.ErrorResponse			"Conflict error"
//	@Failure		500						{object}	response.ErrorResponse			"Internal server error"
//	@Router			/products [post]
func (handler *ProductHandler) Create(ctx *gin.Context) {
	var request request.CreateProductRequest
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		response.HandleError(ctx, err)
		return
	}

	categoryId, _ := domain.ParseID(request.CategoryID)
	product := &domain.Product{
		Name:        request.Name,
		Price:       request.Price,
		Description: request.Description,
		CategoryID:  categoryId,
	}

	product, err = handler.service.Create(ctx, product)
	if err != nil {
		response.HandleError(ctx, err)
		return
	}

	response.HandleSuccess(ctx, response.NewProductResponse(product))
}

// GetByID godoc
//
//	@Summary		Get a product
//	@Description	Retrieves a product by its ID
//	@Tags			Products
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string						true	"Product ID"
//	@Success		200	{object}	response.ProductResponse	"Product found"
//	@Failure		400	{object}	response.ErrorResponse		"Bad Request error"
//	@Failure		404	{object}	response.ErrorResponse		"Not found error"
//	@Failure		500	{object}	response.ErrorResponse		"Internal server error"
//	@Router			/products/{id} [get]
func (handler *ProductHandler) GetByID(ctx *gin.Context) {
	var request request.GetProductRequest

	if err := ctx.ShouldBindUri(&request); err != nil {
		response.HandleError(ctx, err)
		return
	}

	id, _ := domain.ParseID(request.ID)
	product, err := handler.service.GetByID(ctx, id)
	if err != nil {
		response.HandleError(ctx, err)
		return
	}

	response.HandleSuccess(ctx, response.NewProductResponse(product))
}

// GetAll godoc
//
//	@Summary		Get all products
//	@Description	Returns a list of all products
//	@Tags			Products
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	[]response.ProductResponse	"Product list"
//	@Failure		400	{object}	response.ErrorResponse		"Bad Request error"
//	@Failure		404	{object}	response.ErrorResponse		"Not found error"
//	@Failure		500	{object}	response.ErrorResponse		"Internal server error"
//	@Router			/products [get]
func (handler *ProductHandler) GetAll(ctx *gin.Context) {
	products, err := handler.service.GetAll(ctx)
	if err != nil {
		response.HandleError(ctx, err)
		return
	}

	response.HandleSuccess(ctx, response.NewProductListResponse(products))
}

// GetByCategory godoc
//
//	@Summary		Get products by category
//	@Description	Returns all products by category
//	@Tags			Products
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string						true	"Category ID"
//	@Success		200	{object}	[]response.ProductResponse	"Product"
//	@Failure		400	{object}	response.ErrorResponse		"Bad Request error"
//	@Failure		404	{object}	response.ErrorResponse		"Not found error"
//	@Failure		500	{object}	response.ErrorResponse		"Internal server error"
//	@Router			/products/category/{id} [get]
func (h *ProductHandler) GetByCategory(ctx *gin.Context) {
	id, _ := domain.ParseID(ctx.Params.ByName("id"))
	products, err := h.service.GetByCategory(ctx, id)

	if err != nil {
		response.HandleError(ctx, err)
		return
	}

	response.HandleSuccess(ctx, response.NewProductListResponse(products))
}

// Update godoc
//
//	@Summary		Updates a product
//	@Description	Updates a product by its ID
//	@Tags			Products
//	@Accept			json
//	@Produce		json
//	@Param			id						path		string							true	"Product ID"
//	@Param			CreateProductRequest	body		request.CreateProductRequest	true	"Update product request"
//	@Success		200						{object}	response.ProductResponse		"Product updated"
//	@Failure		400						{object}	response.ErrorResponse			"Bad Request error"
//	@Failure		404						{object}	response.ErrorResponse			"Not found error"
//	@Failure		500						{object}	response.ErrorResponse			"Internal server error"
//	@Router			/products/{id} [put]
func (h *ProductHandler) Update(ctx *gin.Context) {
	req := domain.Product{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.HandleError(ctx, err)
		return
	}
	req.ID = domain.ParseIDOrNil(ctx.Param("id"))

	p, err := h.service.Update(ctx, &req)
	if err != nil {
		response.HandleError(ctx, err)
		return
	}

	response.HandleSuccess(ctx, response.NewProductResponse(p))
}

// Delete godoc
//
//	@Summary		Deletes a product
//	@Description	Deletes a single product by its ID
//	@Tags			Products
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string					true	"Product ID"
//	@Success		200	{object}	bool					"Product deleted"
//	@Failure		400	{object}	response.ErrorResponse	"Bad Request error"
//	@Failure		404	{object}	response.ErrorResponse	"Not found error"
//	@Failure		500	{object}	response.ErrorResponse	"Internal server error"
//	@Router			/products/{id} [delete]
func (h *ProductHandler) Delete(ctx *gin.Context) {
	id, _ := domain.ParseID(ctx.Params.ByName("id"))

	err := h.service.Delete(ctx, id)
	if err != nil {
		response.HandleError(ctx, err)
		return
	}
	response.HandleSuccess(ctx, true)
}
