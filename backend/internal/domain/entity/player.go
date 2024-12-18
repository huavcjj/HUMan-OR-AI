package entity

import "time"

type Player struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	GameID    uint      `json:"game_id" gorm:"not null"`
	Name      string    `json:"name" gorm:"unique;not null"`
	Answer    string    `json:"answer" gorm:"not null"`
	Score     uint      `json:"score"`
	CreatedAt time.Time `json:"created_at"`
}

func NewPlayer(gameID uint, name string, score uint) *Player {
	return &Player{
		GameID: gameID,
		Name:   name,
		Score:  score,
	}
}
