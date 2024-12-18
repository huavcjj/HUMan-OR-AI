package dto

import (
	"Bot-or-Not/internal/domain/entity"
)

type Player struct {
	ID     uint   `json:"id"`
	GameID uint   `json:"game_id"`
	Name   string `json:"name"`
	Answer string `json:"answer"`
	// Score     uint      `json:"score"`
	// CreatedAt time.Time `json:"created_at"`
}

func NewPlayerFromEntity(e *entity.Player) *Player {
	return &Player{
		ID:     e.ID,
		GameID: e.GameID,
		Name:   e.Name,
		Answer: e.Answer,
		// Score:     e.Score,
		// CreatedAt: e.CreatedAt,
	}
}

func (p *Player) ToEntity() *entity.Player {
	return &entity.Player{
		ID:     p.ID,
		GameID: p.GameID,
		Name:   p.Name,
		Answer: p.Answer,
		// Score:     p.Score,
		// CreatedAt: p.CreatedAt,
	}
}
