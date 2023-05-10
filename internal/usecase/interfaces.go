package usecase

import (
	"context"
	"quiz-mtuci-server/internal/entity"
)

type AnalyzeTaskIDType string

type (
	QuizRepo interface {
		GetAllQuiz(ctx context.Context) ([]*entity.QuizUI, error)
		GetQuizById(ctx context.Context, quizId int) (*entity.QuizUI, error)
		SaveQuiz(ctx context.Context, quiz *entity.QuizUI) (*entity.QuizUI, error)
	}
	AuthRepo interface {
		GetUser() error
	}
	UseCase interface {
		GetAllQuiz(ctx context.Context) ([]*entity.QuizUI, error)
		GetQuizById(ctx context.Context, quizId int) (*entity.QuizUI, error)
		SaveQuiz(ctx context.Context, quiz *entity.QuizUI) (*entity.QuizUI, error)
		GetUser() error
	}
)
