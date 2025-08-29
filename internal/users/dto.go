package users

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserRequest struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	IsCustomer bool   `json:"is_customer" binding:"required"`
}

type ChangePasswordRequest struct {
	ID       string `json:"id" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	User  User   `json:"user"`
	Token string `json:"token"`
}