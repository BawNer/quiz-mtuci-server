package usecase

import (
	"context"
	"quiz-mtuci-server/internal/entity"
)

type AnalyzeTaskIDType string

type (
	QuizRepo interface {
		GetAllQuiz(ctx context.Context, groupId int) ([]*entity.QuizEntityDB, error)
		GetQuizByHash(ctx context.Context, quizHash string) (*entity.QuizEntityDB, error)
		SaveReviewers(ctx context.Context, reviewer *entity.Reviewers) error
	}
	AuthRepo interface {
		GetUserByLoginWithPassword(ctx context.Context, login entity.UserLogin) (*entity.User, error)
		GetUserByID(ctx context.Context, id int) (*entity.User, error)
		GetGroupByID(ctx context.Context, groupID int) (*entity.Group, error)
	}
	UseCase interface {
		GetAllQuiz(ctx context.Context) ([]*entity.QuizUI, error)
		GetQuizByHash(ctx context.Context, quizHash string) (*entity.QuizUI, error)
		GetUserByLoginWithPassword(ctx context.Context, login entity.UserLogin) (*entity.User, error)
		SaveReviewers(ctx context.Context, reviewer *entity.Reviewers) error
	}
)
