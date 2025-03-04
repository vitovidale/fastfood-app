package request

type CreateCategoryRequest struct {
	Name string `json:"name" binding:"required" example:"Snacks"`
}

type GetCategoryRequest struct {
	ID string `uri:"id" binding:"required"`
}

type UpdateCategoryRequest struct {
	Name string `json:"name" binding:"required" example:"New snack"`
}
