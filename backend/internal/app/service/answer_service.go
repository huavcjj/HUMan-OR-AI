package service

import (
	"Bot-or-Not/internal/app/dto"
	"Bot-or-Not/internal/domain/entity"
	"Bot-or-Not/internal/domain/repository"
	"context"
)

type IAnswerService interface {
	CreateAnswer(ctx context.Context, answer *dto.Answer) (*dto.Answer, error)
	GetAnswersByGameID(ctx context.Context, gameID uint) ([]dto.Answer, error)
}

type answerService struct {
	as repository.IAnswerRepository
}

func NewAnswerService(as repository.IAnswerRepository) IAnswerService {
	return &answerService{as: as}
}

func (s *answerService) CreateAnswer(ctx context.Context, answer *dto.Answer) (*dto.Answer, error) {

	entityAnswer := &entity.Answer{
		GameID:   answer.GameID,
		PlayerID: answer.PlayerID,
		Answer:   answer.Answer,
	}

	createdEntityAnswer, err := s.as.CreateAnswer(ctx, entityAnswer)
	if err != nil {
		return nil, err
	}

	return dto.NewAnswerFromEntity(createdEntityAnswer), nil
}

func (s *answerService) GetAnswersByGameID(ctx context.Context, gameID uint) ([]dto.Answer, error) {
	entityAnswers, err := s.as.GetAnswersByGameID(ctx, gameID)
	if err != nil {
		return nil, err
	}

	var answers []dto.Answer
	for _, entityAnswer := range entityAnswers {
		answers = append(answers, *dto.NewAnswerFromEntity(&entityAnswer))
	}

	return answers, nil
}
