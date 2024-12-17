package repository

import (
	"Bot-or-Not/internal/domain/entity"
	"context"
)

type IAnswerRepository interface {
	CreateAnswer(ctx context.Context, answer *entity.Answer) (*entity.Answer, error)
	GetAnswersByGameID(ctx context.Context, gameID uint) ([]entity.Answer, error)
	GetAnswerByPlayerID(ctx context.Context, playerID uint, gameID uint) (*entity.Answer, error)
}
