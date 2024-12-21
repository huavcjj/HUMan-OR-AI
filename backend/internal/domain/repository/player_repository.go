package repository

import (
	"Bot-or-Not/internal/domain/entity"
	"context"
)

type IPlayerRepository interface {
	CreatePlayer(ctx context.Context, player *entity.Player) (*entity.Player, error)
	GetPlayerByID(ctx context.Context, id uint) (*entity.Player, error)
	GetPlayersByPasscode(ctx context.Context, passcode string) ([]*entity.Player, error)
	UpdatePlayer(ctx context.Context, player *entity.Player) (*entity.Player, error)
	DeletePlayerByID(ctx context.Context, id uint) error
}
