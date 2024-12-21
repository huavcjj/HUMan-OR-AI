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

func (r *playerRepository) UpdatePlayer(ctx context.Context, player *entity.Player) (*entity.Player, error) {
	if err := r.db.WithContext(ctx).Save(player).Error; err != nil {
		return nil, err
	}
	return player, nil
}

func (r *playerRepository) DeletePlayerByID(ctx context.Context, id uint) error {
	var player entity.Player

	if err := r.db.WithContext(ctx).First(&player, id).Error; err != nil {
		return err
	}

	if err := r.db.WithContext(ctx).Delete(&player).Error; err != nil {
		return err
	}
	return nil
}

func (r *playerRepository) GetPlayersByPasscode(ctx context.Context, passcode string) ([]*entity.Player, error) {
	var players []*entity.Player
	if err := r.db.WithContext(ctx).Where("passcode = ?", passcode).Find(&players).Error; err != nil {
		return nil, err
	}
	return players, nil
}
