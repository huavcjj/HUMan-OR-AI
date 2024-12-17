package dto

import (
	"Bot-or-Not/internal/domain/entity"
	"time"
)

type Player struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	GameID    uint      `json:"game_id" gorm:"not null"`
	Name      string    `json:"name" gorm:"unique"`
	Score     uint      `json:"score" `
	CreatedAt time.Time `json:"created_at"`
}

func NewPlayerFromEntity(e *entity.Player) *Player {
	return &Player{
		ID:        e.ID,
		GameID:    e.GameID,
		Name:      e.Name,
		Score:     e.Score,
		CreatedAt: e.CreatedAt,
	}
}

func (p *Player) ToEntity() *entity.Player {
	return &entity.Player{
		ID:        p.ID,
		GameID:    p.GameID,
		Name:      p.Name,
		Score:     p.Score,
		CreatedAt: p.CreatedAt,
	}
}
