package entity

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type Topic struct {
	ID          int64     `json:"id"`
	CategoryID  int64     `json:"category_id" validate:"required"`
	Title       string    `json:"title" validate:"required,min=1,max=100"`
	Description string    `json:"description" validate:"max=500"`
	CreatedAt   time.Time `json:"created_at"`
}

func (t *Topic) Validate() error {
	validate := validator.New()
	return validate.Struct(t)
}
