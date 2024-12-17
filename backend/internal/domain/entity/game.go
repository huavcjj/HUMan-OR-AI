package entity

import "time"

type Game struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Topic     string    `json:"topic"`
	AIAnswer  string    `json:"ai_answer"`
	CreatedAt time.Time `json:"created_at"`
}

func NewGame(topic string, aiAnswer string) *Game {
	return &Game{
		Topic:    topic,
		AIAnswer: aiAnswer,
	}
}
