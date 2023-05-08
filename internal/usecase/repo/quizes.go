package repo

import (
	"context"
	"fmt"
	"quiz-mtuci-server/internal/entity"
	"quiz-mtuci-server/pkg/logger"
	"quiz-mtuci-server/pkg/postgres"

	"github.com/google/uuid"
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
	result := r.DB.Table("quizes").First(&quiz, quizId)
	if result.Error != nil {
		return nil, fmt.Errorf("quiz repo err %v", result.Error)
	}
	return quiz, nil
}

func (r *ServiceRepo) SaveQuiz(ctx context.Context, quiz *entity.Quiz) (*entity.Quiz, error) {
	hash := uuid.New().String()

	newTerminal := map[string]interface{}{
		"author_id": 1,
		"quiz_hash": hash,
		"title":     quiz.Title,
		"questions": quiz.Questions,
		"active":    true,
	}

	if result := r.DB.Table("quizes").Create(newTerminal); result.Error != nil {
		return nil, result.Error
	}

	var createdQuiz *entity.Quiz
	if response := r.DB.Table("quizes").Where("quiz_hash = ?", hash).Find(&createdQuiz); response.Error != nil {
		return nil, response.Error
	}

	return createdQuiz, nil
}
