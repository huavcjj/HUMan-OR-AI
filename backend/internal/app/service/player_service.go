package service

import (
	"Bot-or-Not/internal/app/dto"
	"Bot-or-Not/internal/domain/repository"
	"Bot-or-Not/pkg/config"
	"Bot-or-Not/pkg/openai"
	"context"
)

type IPlayerService interface {
	CreatePlayer(ctx context.Context, player *dto.Player) (*dto.Player, error)
	GetPlayerByID(ctx context.Context, id uint) (*dto.Player, error)
	FindAvailableOpponentByPasscode(ctx context.Context, id uint, passcode string) (*dto.Player, error)
	UpdateTopicAndAIAnswer(ctx context.Context, id uint, topic string) (*dto.Player, error)
	UpdateOpponentAnswer(ctx context.Context, id uint, opponentAnswer string) (*dto.Player, error)
	UpdateSelectAnswerIsPlayer(ctx context.Context, id uint, selectAnswerIsPlayer bool) (*dto.Player, error)
	DeletePlayerByID(ctx context.Context, id uint) error
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

func (ps *playerService) GetPlayerByID(ctx context.Context, id uint) (*dto.Player, error) {
	player, err := ps.pr.GetPlayerByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return dto.NewPlayerFromEntity(player), nil
}

func (ps *playerService) FindAvailableOpponentByPasscode(ctx context.Context, id uint, passcode string) (*dto.Player, error) {
	players, err := ps.pr.GetPlayersByPasscode(ctx, passcode)
	if err != nil {
		return nil, err
	}
	for _, player := range players {
		if player.ID != id {
			return dto.NewPlayerFromEntity(player), nil
		}
	}
	return nil, nil
}

func (ps *playerService) UpdateTopicAndAIAnswer(ctx context.Context, id uint, topic string) (*dto.Player, error) {
	player, err := ps.pr.GetPlayerByID(ctx, id)
	if err != nil {
		return nil, err
	}
	player.Topic = topic

	openAI := openai.NewOpenAI(config.APIKey)
	aiAnswer, err := openAI.GenerateAIAnswer(ctx, topic)
	if err != nil {
		return nil, err
	}
	player.AIAnswer = aiAnswer

	updatedPlayer, err := ps.pr.UpdatePlayer(ctx, player)
	if err != nil {
		return nil, err
	}
	return dto.NewPlayerFromEntity(updatedPlayer), nil
}

func (ps *playerService) UpdateOpponentAnswer(ctx context.Context, id uint, opponentAnswer string) (*dto.Player, error) {
	player, err := ps.pr.GetPlayerByID(ctx, id)
	if err != nil {
		return nil, err
	}
	player.OpponentAnswer = opponentAnswer

	updatedPlayer, err := ps.pr.UpdatePlayer(ctx, player)
	if err != nil {
		return nil, err
	}
	return dto.NewPlayerFromEntity(updatedPlayer), nil
}

func (ps *playerService) UpdateSelectAnswerIsPlayer(ctx context.Context, id uint, selectAnswerIsPlayer bool) (*dto.Player, error) {
	player, err := ps.pr.GetPlayerByID(ctx, id)
	if err != nil {
		return nil, err
	}
	player.SelectAnswerIsPlayer = selectAnswerIsPlayer

	updatedPlayer, err := ps.pr.UpdatePlayer(ctx, player)
	if err != nil {
		return nil, err
	}
	return dto.NewPlayerFromEntity(updatedPlayer), nil
}

func (ps *playerService) DeletePlayerByID(ctx context.Context, id uint) error {
	if err := ps.pr.DeletePlayerByID(ctx, id); err != nil {
		return err
	}
	return nil
}
