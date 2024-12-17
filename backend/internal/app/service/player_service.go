package service

import (
	"Bot-or-Not/internal/app/dto"
	"Bot-or-Not/internal/domain/repository"
	"context"
)

type IPlayerService interface {
	CreatePlayer(ctx context.Context, player *dto.Player) (*dto.Player, error)
	GetPlayersByGameID(ctx context.Context, gameID uint) ([]dto.Player, error)
}

type playerService struct {
	pr repository.IPlayerRepository
}

func NewPlayerService(pr repository.IPlayerRepository) IPlayerService {
	return &playerService{
		pr: pr,
	}
}

func (ps *playerService) CreatePlayer(ctx context.Context, player *dto.Player) (*dto.Player, error) {
	newPlayer, err := ps.pr.CreatePlayer(ctx, player.ToEntity())
	if err != nil {
		return nil, err
	}
	return dto.NewPlayerFromEntity(newPlayer), nil
}

func (ps *playerService) GetPlayersByGameID(ctx context.Context, gameID uint) ([]dto.Player, error) {
	players, err := ps.pr.GetPlayersByGameID(ctx, gameID)
	if err != nil {
		return nil, err
	}
	var dtoPlayers []dto.Player
	for _, player := range players {
		dtoPlayers = append(dtoPlayers, *dto.NewPlayerFromEntity(&player))
	}
	return dtoPlayers, nil
}
