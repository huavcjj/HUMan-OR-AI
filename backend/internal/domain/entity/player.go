package entity

type Player struct {
	ID             uint   `gorm:"primaryKey"`
	Passcode       string `gorm:"not null"`
	Topic          string `gorm:"not null"`
	AIAnswer       string
	OpponentAnswer string
}

func NewPlayer(passcode, topic string) *Player {
	return &Player{
		Passcode: passcode,
		Topic:    topic,
	}
}
