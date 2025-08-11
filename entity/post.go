package entity

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type Post struct {
	ID        int64     `json:"id"`
	TopicID   int64     `json:"topic_id" validate:"required"`
	AuthorID  int64     `json:"author_id" validate:"required"`
	Title     string    `json:"title" validate:"required,min=1,max=100"`
	Content   string    `json:"content" validate:"required,min=1,max=2000"`
	CreatedAt time.Time `json:"created_at"`
}

func (p *Post) Validate() error {
	validate := validator.New()
	return validate.Struct(p)
}
