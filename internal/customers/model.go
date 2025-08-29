package customers

import (
	"perretes-api/internal/users"

	"github.com/google/uuid"
)

type Customer struct {
	ID            uuid.UUID `json:"id" db:"id"`
	Name          string    `json:"name" db:"name"`
	Surname       string    `json:"surname" db:"surname"`
	PhoneNumber   string    `json:"phone_number" db:"phone_number"`
	Email         string    `json:"email" db:"email"`		
	User users.User `json:"user"`
	IsActive	  bool      `json:"is_active" db:"is_active"`
}