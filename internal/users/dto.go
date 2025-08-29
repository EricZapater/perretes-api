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

type UserResponse struct {
	ID                string  `json:"id" db:"id"`
	Username          string  `json:"username" db:"username"`
	Password          string  `json:"password" db:"password"`
	IsActive          bool    `json:"is_active" db:"is_active"`
	IsCustomer        bool    `json:"is_customer" db:"is_customer"`
	PasswordChangedAt *string `json:"password_changed_at" db:"password_changed_at"`
}

type LoginResponse struct {
	User  UserResponse `json:"user"`
	Token string       `json:"token"`
}