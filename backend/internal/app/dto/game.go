package dto

import (
	"Bot-or-Not/internal/domain/entity"
	"time"
)

type Game struct {
	ID        uint      `json:"id" `
	Topic     string    `json:"topic"`
	AIAnswer  string    `json:"ai_answer"`
	CreatedAt time.Time `json:"created_at"`
}

func NewGameFromEntity(e *entity.Game) *Game {
	return &Game{
		ID:        e.ID,
		Topic:     e.Topic,
		AIAnswer:  e.AIAnswer,
		CreatedAt: e.CreatedAt,
	}
}
