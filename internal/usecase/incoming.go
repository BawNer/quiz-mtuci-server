package usecase

import (
	"context"
	"quiz-mtuci-server/internal/entity"
	"quiz-mtuci-server/pkg/logger"
)

type ServiceUseCase struct {
	logger *logger.Logger
	repo   QuizRepo
	auth   AuthRepo
}

func New(logger *logger.Logger, r QuizRepo, a AuthRepo) *ServiceUseCase {
	return &ServiceUseCase{
		logger: logger,
		repo:   r,
		auth:   a,
	}
}

func (s *ServiceUseCase) GetAllQuiz(ctx context.Context) ([]*entity.Quiz, error) {
	return s.repo.GetAllQuiz(ctx)
}

func (s *ServiceUseCase) GetQuizById(ctx context.Context, quizId int) (*entity.Quiz, error) {
	return s.repo.GetQuizById(ctx, quizId)
}

func (s *ServiceUseCase) SaveQuiz(ctx context.Context, quiz *entity.Quiz) (*entity.Quiz, error) {
	return s.repo.SaveQuiz(ctx, quiz)
}

func (s *ServiceUseCase) GetUser() error {
	return s.GetUser()
}
