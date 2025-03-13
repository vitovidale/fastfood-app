package request

type GetOrderRequest struct {
	ID string `uri:"id" binding:"required,min=1" example:"00000000-0000-0000-0000-000000000000"`
}

type GetOrderByCustomerRequest struct {
	CustomerID uint64 `uri:"customerId" binding:"required,min=1" example:"1"`
}

type CreateOrderProductRequest struct {
	ProductID string `json:"productId" example:"00000000-0000-0000-0000-000000000000"`
	Quantity  uint16 `json:"quantity" example:"1"`
	Notes     string `json:"notes" example:"notes"`
}

type CreateOrderRequest struct {
	CustomerID uint64                      `uri:"customerId" binding:"required,min=1" example:"1"`
	Products   []CreateOrderProductRequest `json:"products" binding:"required"`
}

type RemoveProductRequest struct {
	OrderID        string `uri:"orderId" binding:"required,min=1" example:"00000000-0000-0000-0000-000000000000"`
	OrderProductID string `uri:"orderProductId" binding:"required,min=1" example:"00000000-0000-0000-0000-000000000000"`
}

type AddProductRequest struct {
	ProductID  string `json:"productId" binding:"required" example:"00000000-0000-0000-0000-000000000000"`
	CustomerID uint64 `json:"customerId" example:"00000000000"`
	Quantity   uint16 `json:"quantity" binding:"required" example:"1"`
	Notes      string `json:"notes" example:"notes"`
}
