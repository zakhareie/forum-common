package entity

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type Comment struct {
	ID        int64     `json:"id"`
	PostID    int64     `json:"post_id" validate:"required"`
	AuthorID  int64     `json:"author_id" validate:"required"`
	Content   string    `json:"content" validate:"required,min=1,max=1000"`
	CreatedAt time.Time `json:"created_at"`
}

func (c *Comment) Validate() error {
	validate := validator.New()
	return validate.Struct(c)
}
