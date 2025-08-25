package entity

import (
	"github.com/go-playground/validator/v10"
)

type Category struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func (m *Category) Validate() error {
	validate := validator.New()
	return validate.Struct(m)
}
