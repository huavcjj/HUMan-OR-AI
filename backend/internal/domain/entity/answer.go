package entity

import "time"

type Answer struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	PlayerID  uint      `json:"player_id" gorm:"not null"`
	GameID    uint      `json:"game_id" gorm:"not null"`
	Answer    string    `json:"answer"`
	CreatedAt time.Time `json:"created_at"`
	Player    Player    `json:"player" gorm:"foreignKey:PlayerID"`
	Game      Game      `json:"game" gorm:"foreignKey:GameID"`
}

func NewAnswer(playerID uint, gameID uint, answer string) *Answer {
	return &Answer{
		PlayerID: playerID,
		GameID:   gameID,
		Answer:   answer,
	}
}
