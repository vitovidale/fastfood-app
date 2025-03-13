package request

type CreateProductRequest struct {
	Name        string  `json:"name" binding:"required" example:"Potato Chips"`
	Price       float64 `json:"price" binding:"required,min=0" example:"10000"`
	Description string  `json:"description" example:"Potato chips with cheese flavor"`
	CategoryID  string  `json:"categoryId"`
}

type GetProductRequest struct {
	ID string `uri:"id"`
}
