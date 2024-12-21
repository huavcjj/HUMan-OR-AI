package dto

import (
	"Bot-or-Not/internal/domain/entity"
)

type Player struct {
	ID             uint   `json:"id"`
	Passcode       string `json:"passcode"`
	Topic          string `json:"topic"`
	AIAnswer       string `json:"ai_answer"`
	OpponentAnswer string `json:"opponent_answer"`
}

type PasscodeReq struct {
	Passcode string `json:"passcode"`
}

type TopicReq struct {
	ID    uint   `json:"id"`
	Topic string `json:"topic"`
}

type AnswerReq struct {
	ID       uint   `json:"id"`
	Passcode string `json:"passcode"`
	Answer   string `json:"answer"`
}

type AnswerIsPlayerReq struct {
	ID           uint   `json:"id"`
	SelectAnswer string `json:"select_answer"`
}

type IDReq struct {
	ID uint `json:"id"`
}

type AnswersResp struct {
	AIAnswer       string `json:"ai_answer"`
	OpponentAnswer string `json:"opponent_answer"`
}

type PasscodeResp struct {
	ID       uint   `json:"id"`
	Passcode string `json:"passcode"`
}

type TopicResp struct {
	Topic string `json:"topic"`
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
		OpponentAnswer: player.OpponentAnswer,
	}
}

func (p *Player) ToEntity() *entity.Player {
	return &entity.Player{
		ID:             p.ID,
		Passcode:       p.Passcode,
		Topic:          p.Topic,
		AIAnswer:       p.AIAnswer,
		OpponentAnswer: p.OpponentAnswer,
	}
}
