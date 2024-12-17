package repository

import (
	"Bot-or-Not/internal/domain/entity"
	"Bot-or-Not/internal/domain/repository"
	"Bot-or-Not/internal/infra/database"
	"context"
)

type playerRepository struct {
	db *database.DB
}

func NewPlayerRepository(db *database.DB) repository.IPlayerRepository {
	return &playerRepository{db: db}
}

func (r *playerRepository) CreatePlayer(ctx context.Context, player *entity.Player) (*entity.Player, error) {
	if err := r.db.WithContext(ctx).Create(player).Error; err != nil {
		return nil, err
	}
	return player, nil
}

func (r *playerRepository) GetPlayerByID(ctx context.Context, id uint) (*entity.Player, error) {
	var player entity.Player
	if err := r.db.WithContext(ctx).First(&player, id).Error; err != nil {
		return nil, err
	}
	return &player, nil
}

func (r *playerRepository) GetPlayersByGameID(ctx context.Context, gameID uint) ([]entity.Player, error) {
	var players []entity.Player
	if err := r.db.WithContext(ctx).Where("game_id = ?", gameID).Find(&players).Error; err != nil {
		return nil, err
	}
	return players, nil
}

func (r *playerRepository) UpdatePlayerScore(ctx context.Context, playerID uint, score int) error {
	if err := r.db.WithContext(ctx).Model(&entity.Player{}).Where("id = ?", playerID).Update("score", score).Error; err != nil {
		return err
	}
	return nil
}
