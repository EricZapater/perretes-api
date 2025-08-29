package users

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID `json:"id" db:"id"`
	Username string    `json:"username" db:"username"`
	Password string    `json:"password" db:"password"`	
	IsActive bool `json:"is_active" db:"is_active"`	
	IsCustomer bool `json:"is_customer" db:"is_customer"`
	PasswordChangedAt *time.Time `json:"password_changed_at" db:"password_changed_at"`
}