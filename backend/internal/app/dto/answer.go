package dto

import (
	"Bot-or-Not/internal/domain/entity"
	"time"
)

type Answer struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	PlayerID  uint      `json:"player_id" gorm:"not null"`
	GameID    uint      `json:"game_id" gorm:"not null"`
	Answer    string    `json:"answer" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
}

func NewAnswerFromEntity(e *entity.Answer) *Answer {
	return &Answer{
		ID:        e.ID,
		PlayerID:  e.PlayerID,
		GameID:    e.GameID,
		Answer:    e.Answer,
		CreatedAt: e.CreatedAt,
	}
}

func (a *Answer) ToEntity() *entity.Answer {
	return &entity.Answer{
		ID:        a.ID,
		PlayerID:  a.PlayerID,
		GameID:    a.GameID,
		Answer:    a.Answer,
		CreatedAt: a.CreatedAt,
	}
}
