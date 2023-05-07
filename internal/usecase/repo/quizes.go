package repo

import (
	"context"
	"fmt"
	"quiz-mtuci-server/internal/entity"
	"quiz-mtuci-server/pkg/logger"
	"quiz-mtuci-server/pkg/postgres"
)

type ServiceRepo struct {
	*postgres.Postgres
	l *logger.Logger
}

func New(pg *postgres.Postgres, l *logger.Logger) *ServiceRepo {
	return &ServiceRepo{pg, l}
}

func (r *ServiceRepo) GetAllQuiz(ctx context.Context) ([]entity.Quiz, error) {
	var quizes []entity.Quiz
	result := r.DB.Table("quizes").Find(&quizes)
	if result.Error != nil {
		return nil, fmt.Errorf("quiz repo err %v", result.Error)
	}
	return quizes, nil
}

func (r *ServiceRepo) GetQuizById(ctx context.Context, quizId int) (*entity.Quiz, error) {
	var quiz *entity.Quiz
	result := r.DB.Table("quizes").Find(&quiz, quizId)
	if result.Error != nil {
		return nil, fmt.Errorf("quiz repo err %v", result.Error)
	}
	return quiz, nil
}
