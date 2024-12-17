package repository

import (
	"Bot-or-Not/internal/domain/entity"
	"context"
)

type IPlayerRepository interface {
	CreatePlayer(ctx context.Context, player *entity.Player) (*entity.Player, error)
	GetPlayerByID(ctx context.Context, id uint) (*entity.Player, error)
	GetPlayersByGameID(ctx context.Context, gameID uint) ([]entity.Player, error)
	UpdatePlayerScore(ctx context.Context, playerID uint, score int) error
}
