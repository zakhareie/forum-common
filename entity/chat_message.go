package entity

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type ChatMessage struct {
	ID        int64     `json:"id"`
	UserID    int64     `json:"user_id" validate:"required"`
	Username  string    `json:"username"`
	Content   string    `json:"content" validate:"required,min=1,max=2000"`
	CreatedAt time.Time `json:"created_at"`
	Timestamp time.Time `json:"timestamp"`
}

func (m *ChatMessage) Validate() error {
	validate := validator.New()
	return validate.Struct(m)
}
