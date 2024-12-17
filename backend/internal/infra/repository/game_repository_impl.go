package repository

import (
	"Bot-or-Not/internal/domain/entity"
	"Bot-or-Not/internal/domain/repository"
	"Bot-or-Not/internal/infra/database"
	"context"
)

type gameRepository struct {
	db *database.DB
}

func NewGameRepository(db *database.DB) repository.IGameRepository {
	return &gameRepository{db: db}
}

func (r *gameRepository) CreateGame(ctx context.Context, game *entity.Game) (*entity.Game, error) {
	if err := r.db.WithContext(ctx).Create(game).Error; err != nil {
		return nil, err
	}
	return game, nil
}

func (r *gameRepository) GetGameByID(ctx context.Context, id uint) (*entity.Game, error) {
	var game entity.Game
	if err := r.db.WithContext(ctx).First(&game, id).Error; err != nil {
		return nil, err
	}
	return &game, nil
}
