package entity

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type UserRole string

const (
	RoleGuest UserRole = "guest"
	RoleUser  UserRole = "user"
	RoleAdmin UserRole = "admin"
)

type User struct {
	ID        int64     `json:"id"`
	Username  string    `json:"username" validate:"required,min=3,max=32"`
	Password  string    `json:"password,omitempty" validate:"required,min=6"` // Hashed
	Role      UserRole  `json:"role" validate:"required,oneof=guest user admin"`
	CreatedAt time.Time `json:"created_at"`
}

func (u *User) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}
