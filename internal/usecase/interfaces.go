package usecase

import (
	"context"
	"quiz-mtuci-server/internal/entity"
)

type AnalyzeTaskIDType string

type (
	QuizRepo interface {
		GetAllQuiz(ctx context.Context) ([]entity.Quiz, error)
		GetQuizById(ctx context.Context, quizId int) (*entity.Quiz, error)
	}
	AuthRepo interface {
		GetUser() error
	}
	UseCase interface {
		GetAllQuiz(ctx context.Context) ([]entity.Quiz, error)
		GetQuizById(ctx context.Context, quizId int) (*entity.Quiz, error)
		GetUser() error
	}
)
