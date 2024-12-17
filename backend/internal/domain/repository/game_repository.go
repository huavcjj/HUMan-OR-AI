package repository

import (
	"Bot-or-Not/internal/domain/entity"
	"context"
)

type IGameRepository interface {
	CreateGame(ctx context.Context, game *entity.Game) (*entity.Game, error)
	GetGameByID(ctx context.Context, id uint) (*entity.Game, error)
}
