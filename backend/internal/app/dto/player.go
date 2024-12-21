package dto

import (
	"Bot-or-Not/internal/domain/entity"
)

type Player struct {
	ID             uint   `json:"id" gorm:"primaryKey"`
	Passcode       string `json:"passcode" gorm:"not null"`
	Topic          string `json:"topic" gorm:"not null"`
	AIAnswer       string `json:"ai_answer"`
	Answer         string `json:"answer"`
	OpponentAnswer string `json:"opponent_answer"`
	SelectAnswer   string `json:"select_answer"`
}

type AnswersResp struct {
	AIAnswer       string `json:"ai_answer"`
	OpponentAnswer string `json:"opponent_answer"`
}

func NewPlayer(passcode string) *Player {
	return &Player{
		Passcode: passcode,
	}
}

func NewTopic(topic string) *Player {
	return &Player{
		Topic: topic,
	}
}

func NewPlayerFromEntity(player *entity.Player) *Player {
	return &Player{
		ID:             player.ID,
		Passcode:       player.Passcode,
		Topic:          player.Topic,
		AIAnswer:       player.AIAnswer,
		Answer:         player.Answer,
		OpponentAnswer: player.OpponentAnswer,
	}
}

func (p *Player) ToEntity() *entity.Player {
	return &entity.Player{
		ID:             p.ID,
		Passcode:       p.Passcode,
		Topic:          p.Topic,
		AIAnswer:       p.AIAnswer,
		Answer:         p.Answer,
		OpponentAnswer: p.OpponentAnswer,
	}
}
