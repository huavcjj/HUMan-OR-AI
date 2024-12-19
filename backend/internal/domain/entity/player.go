package entity

type Player struct {
	ID             uint   `json:"id" gorm:"primaryKey"`
	Passcode       string `json:"passcode" gorm:"not null"`
	Topic          string `json:"topic" gorm:"not null"`
	AIAnswer       string `json:"ai_answer"`
	Answer         string `json:"answer"`
	OpponentAnswer string `json:"opponent_answer"`
	SelectAnswer   string `json:"select_answer"`
}

func NewPlayer(passcode, topic string) *Player {
	return &Player{
		Passcode: passcode,
		Topic:    topic,
	}
}
