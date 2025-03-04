package http

import (
	"github.com/gin-gonic/gin"
	"github.com/vitovidale/TECH-CHALLENGE/internal/adapter/driver/handler/http/request"
	"github.com/vitovidale/TECH-CHALLENGE/internal/adapter/driver/handler/http/response"
	"github.com/vitovidale/TECH-CHALLENGE/internal/core/domain"
	"github.com/vitovidale/TECH-CHALLENGE/internal/core/port"
)

type CategoryHandler struct {
	service port.CategoryService
}

func NewCategoryHandler(service port.CategoryService) *CategoryHandler {
	return &CategoryHandler{service: service}
}

// GetByID godoc
//
//	@Summary	Get a category by its ID
//	@Tags		Categories
//	@Accept		json
//	@Produce	json
//	@Param		id	path		string						true	"Category ID"
//	@Success	200	{object}	response.CategoryResponse	"Category found"
//	@Failure	500	{object}	response.ErrorResponse		"Internal server error"
//	@Router		/categories/{id} [get]
func (h *CategoryHandler) GetByID(ctx *gin.Context) {
	var req request.GetCategoryRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		response.HandleError(ctx, err)
		return
	}

	id, _ := domain.ParseID(req.ID)
	category, err := h.service.GetByID(ctx, id)
	if err != nil {
		response.HandleError(ctx, err)
		return
	}

	response.HandleSuccess(ctx, response.NewCategoryResponse(category))
}

// GetAll godoc
//
//	@Summary		Get all categories
//	@Description	Get a complete list of all categories
//	@Tags			Categories
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	[]response.CategoryResponse	"List of categories"
//	@Failure		500	{object}	response.ErrorResponse		"Internal server error"
//	@Router			/categories [get]
func (h *CategoryHandler) GetAll(ctx *gin.Context) {
	categories, err := h.service.GetAll(ctx)
	if err != nil {
		response.HandleError(ctx, err)
		return
	}

	response.HandleSuccess(ctx, response.NewCategoryListResponse(categories))
}

// Create godoc
//
//	@Summary		Create a new category
//	@Description	Creates a new category with the given name
//	@Tags			Categories
//	@Accept			json
//	@Produce		json
//	@Param			CreateCategoryRequest	body		request.CreateCategoryRequest	true	"Create category request"
//	@Success		200						{object}	response.CategoryResponse		"Category created"
//	@Failure		400						{object}	response.ErrorResponse			"Validation errors"
//	@Failure		500						{object}	response.ErrorResponse			"Internal server error"
//	@Router			/categories [post]
func (h *CategoryHandler) Create(ctx *gin.Context) {
	var req request.CreateCategoryRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.HandleError(ctx, err)
		return
	}

	category, err := h.service.Create(ctx, domain.NewCategory(req.Name))
	if err != nil {
		response.HandleError(ctx, err)
		return
	}

	response.HandleSuccess(ctx, response.NewCategoryResponse(category))
}

// Update godoc
//
//	@Summary		Updates a category
//	@Description	Updates a category based on its ID
//	@Tags			Categories
//	@Accept			json
//	@Produce		json
//	@Param			id						path		string							true	"Category ID"
//	@Param			UpdateCategoryRequest	body		request.UpdateCategoryRequest	true	"Update category request"
//	@Success		200						{object}	response.CategoryResponse		"Category updated"
//	@Failure		400						{object}	response.ErrorResponse			"Bad Request error"
//	@Failure		404						{object}	response.ErrorResponse			"Not found error"
//	@Failure		500						{object}	response.ErrorResponse			"Internal server error"
//	@Router			/categories/{id} [put]
func (h *CategoryHandler) Update(ctx *gin.Context) {
	var req request.UpdateCategoryRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.HandleError(ctx, err)
		return
	}

	id, _ := domain.ParseID(ctx.Params.ByName("id"))
	category, err := h.service.Update(ctx, &domain.Category{
		ID:   id,
		Name: req.Name,
	})

	if err != nil {
		response.HandleError(ctx, err)
		return
	}

	response.HandleSuccess(ctx, response.NewCategoryResponse(category))
}

// Delete godoc
//
//	@Summary		Deletes a category
//	@Description	Deletes a category in our app base in its ID
//	@Tags			Categories
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string					true	"Category ID"
//	@Success		200	{object}	bool					"Category deleted"
//	@Failure		400	{object}	response.ErrorResponse	"Bad Request error"
//	@Failure		404	{object}	response.ErrorResponse	"Not found error"
//	@Failure		500	{object}	response.ErrorResponse	"Internal server error"
//	@Router			/categories/{id} [delete]
func (h *CategoryHandler) Delete(ctx *gin.Context) {
	id, _ := domain.ParseID(ctx.Params.ByName("id"))

	err := h.service.Delete(ctx, id)
	if err != nil {
		response.HandleError(ctx, err)
		return
	}
	response.HandleSuccess(ctx, true)
}
