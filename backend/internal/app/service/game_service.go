package service

import (
	"Bot-or-Not/internal/app/dto"
	"Bot-or-Not/internal/domain/entity"
	"Bot-or-Not/internal/domain/repository"
	"Bot-or-Not/pkg/config"
	"Bot-or-Not/pkg/openai"
	"context"
)

type IGameService interface {
	CreateGame(ctx context.Context, id uint) (*dto.Game, error)
	GetGameByID(ctx context.Context, id uint) (*dto.Game, error)
}

type gameService struct {
	gr repository.IGameRepository
}

func NewGameService(gr repository.IGameRepository) IGameService {
	return &gameService{gr: gr}
}

func (s *gameService) CreateGame(ctx context.Context, id uint) (*dto.Game, error) {

	openAI := openai.NewOpenAI(config.APIKey)
	topic, aiAnswer, err := openAI.GenerateGameTopicAIAnswer(ctx)
	if err != nil {
		return nil, err
	}

	newGame := &entity.Game{
		Topic:    topic,
		AIAnswer: aiAnswer,
		ID:       id,
	}

	createdGame, err := s.gr.CreateGame(ctx, newGame)
	if err != nil {
		return nil, err
	}
	return dto.NewGameFromEntity(createdGame), nil
}

func (s *gameService) GetGameByID(ctx context.Context, id uint) (*dto.Game, error) {
	game, err := s.gr.GetGameByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return dto.NewGameFromEntity(game), nil
}
