package repository

import (
	"Bot-or-Not/internal/domain/entity"
	"Bot-or-Not/internal/domain/repository"
	"Bot-or-Not/internal/infra/database"
	"context"
)

type answerRepository struct {
	db *database.DB
}

func NewAnswerRepository(db *database.DB) repository.IAnswerRepository {
	return &answerRepository{db: db}
}

func (r *answerRepository) CreateAnswer(ctx context.Context, answer *entity.Answer) (*entity.Answer, error) {
	if err := r.db.WithContext(ctx).Create(answer).Error; err != nil {
		return nil, err
	}
	return answer, nil
}

func (r *answerRepository) GetAnswersByGameID(ctx context.Context, gameID uint) ([]entity.Answer, error) {
	var answers []entity.Answer
	if err := r.db.WithContext(ctx).Where("game_id = ?", gameID).Find(&answers).Error; err != nil {
		return nil, err
	}
	return answers, nil
}

func (r *answerRepository) GetAnswerByPlayerID(ctx context.Context, playerID uint, gameID uint) (*entity.Answer, error) {
	var answer entity.Answer
	if err := r.db.WithContext(ctx).Where("player_id = ? AND game_id = ?", playerID, gameID).First(&answer).Error; err != nil {
		return nil, err
	}
	return &answer, nil
}
