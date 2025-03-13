package request

type CreateCustomerRequest struct {
	ID        uint64 `json:"id" binding:"required,min=1" example:"12345678910"`
	FirstName string `json:"firstName" binding:"required" example:"John"`
	LastName  string `json:"lastName" binding:"required" example:"Doe"`
	Email     string `json:"email" binding:"required,email" example:"john.doe@example.com"`
	Password  string `json:"password" binding:"required,min=8" example:"12345678"`
}

type UpdateCustomerRequest struct {
	FirstName string `json:"firstName" example:"John"`
	LastName  string `json:"lastName" example:"Doe"`
	Email     string `json:"email" example:"john.doe@example.com"`
}

type AuthCustomerRequest struct {
	ID       uint64 `json:"id" example:"12345678910"`
	Email    string `json:"email" example:"john.doe@example.com"`
	Password string `json:"password" binding:"required" example:"12345678"`
}

type GetCustomerByIDRequest struct {
	ID uint64 `uri:"id" binding:"required,min=1" example:"12345678910"`
}
