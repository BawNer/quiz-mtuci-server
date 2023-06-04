package usecase

import (
	"context"
	"quiz-mtuci-server/internal/entity"
	"quiz-mtuci-server/pkg/logger"
	"strconv"
	"strings"
	"time"
)

type ServiceUseCase struct {
	logger *logger.Logger
	jwt    JWT
	repo   QuizRepo
	auth   AuthRepo
}

func New(logger *logger.Logger, j JWT, r QuizRepo, a AuthRepo) *ServiceUseCase {
	return &ServiceUseCase{
		logger: logger,
		jwt:    j,
		repo:   r,
		auth:   a,
	}
}

func (s *ServiceUseCase) GetAllQuiz(ctx context.Context) ([]*entity.QuizUI, error) {
	var response []*entity.QuizUI

	groupID := s.jwt.Parse(ctx.Value("token").(map[string]interface{})).GroupID
	quizzes, err := s.repo.GetAllQuiz(ctx, groupID)
	if err != nil {
		return nil, err
	}

	for _, quiz := range quizzes {
		var groups []*entity.Group
		t := strings.Split(strings.ReplaceAll(quiz.AccessFor, " ", ""), ",")
		for _, v := range t {
			g, ok := strconv.Atoi(v)
			if ok != nil {
				return nil, ok
			}
			group, err := s.auth.GetGroupByID(ctx, g)
			if err != nil {
				return nil, err
			}
			groups = append(groups, group)
		}
		user, err := s.auth.GetUserByID(ctx, quiz.AuthorID)
		if err != nil {
			return nil, err
		}

		quiz.Author = user
		response = append(response, &entity.QuizUI{
			ID:        quiz.ID,
			AuthorID:  quiz.AuthorID,
			Author:    quiz.Author,
			AccessFor: groups,
			QuizHash:  quiz.QuizHash,
			Title:     quiz.Title,
			Questions: quiz.Questions,
			Active:    quiz.Active,
			CreatedAt: quiz.CreatedAt,
			UpdatedAt: quiz.UpdatedAt,
		})
	}

	return response, nil
}

func (s *ServiceUseCase) GetQuizByHash(ctx context.Context, quizHash string) (*entity.QuizUI, error) {
	var (
		response *entity.QuizUI
		groups   []*entity.Group
	)
	quiz, err := s.repo.GetQuizByHash(ctx, quizHash)
	if err != nil {
		return nil, err
	}
	t := strings.Split(strings.ReplaceAll(quiz.AccessFor, " ", ""), ",")
	for _, v := range t {
		g, ok := strconv.Atoi(v)
		if ok != nil {
			return nil, ok
		}
		group, err := s.auth.GetGroupByID(ctx, g)
		if err != nil {
			return nil, err
		}
		groups = append(groups, group)
	}
	user, err := s.auth.GetUserByID(ctx, quiz.AuthorID)
	if err != nil {
		return nil, err
	}

	quiz.Author = user
	response = &entity.QuizUI{
		ID:        quiz.ID,
		AuthorID:  quiz.AuthorID,
		Author:    quiz.Author,
		AccessFor: groups,
		QuizHash:  quiz.QuizHash,
		Title:     quiz.Title,
		Questions: quiz.Questions,
		Active:    quiz.Active,
		CreatedAt: quiz.CreatedAt,
		UpdatedAt: quiz.UpdatedAt,
	}

	return response, nil
}

func (s *ServiceUseCase) GetUserByLoginWithPassword(ctx context.Context, user entity.UserLogin) (*entity.User, error) {
	foundedUser, err := s.auth.GetUserByLoginWithPassword(ctx, user)
	if err != nil {
		return nil, err
	}
	token, err := s.jwt.Create(time.Hour*24, foundedUser)
	if err != nil {
		return nil, err
	}
	foundedUser.Token = token

	return foundedUser, nil
}

func (s *ServiceUseCase) SaveReviewers(ctx context.Context, reviewer *entity.Reviewers) error {
	reviewer.UserID = s.jwt.Parse(ctx.Value("token").(map[string]interface{})).ID
	return s.repo.SaveReviewers(ctx, reviewer)
}
