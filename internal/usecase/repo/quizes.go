package repo

import (
	"context"
	"fmt"
	"quiz-mtuci-server/internal/entity"
	"quiz-mtuci-server/pkg/logger"
	"quiz-mtuci-server/pkg/postgres"
)

type QuizRepo struct {
	*postgres.Postgres
	l *logger.Logger
}

func New(pg *postgres.Postgres, l *logger.Logger) *QuizRepo {
	return &QuizRepo{pg, l}
}

func (r *QuizRepo) GetAllQuiz(ctx context.Context, groupID int) ([]*entity.QuizEntityDB, error) {
	var (
		response []*entity.QuizEntityDB
		quizzes  []entity.Quiz
	)
	result := r.DB.Table("quizzes").Where("active = ?", true).Where("access_for LIKE ?", fmt.Sprintf("%%%d%%", groupID)).Find(&quizzes)
	if result.Error != nil {
		return nil, fmt.Errorf("quiz repo err %v", result.Error)
	}

	for _, quiz := range quizzes {
		var (
			questionsUI []entity.QuestionsUI
			questions   []entity.Question
			answers     []entity.AnswersOption
		)
		if err := r.DB.Table("questions").Where("quiz_id = ?", quiz.ID).Find(&questions); err.Error != nil {
			return nil, err.Error
		}

		for _, question := range questions {

			if err := r.DB.Table("answers_options").Where("question_id = ?", question.ID).Find(&answers); err.Error != nil {
				return nil, err.Error
			}

			questionsUI = append(questionsUI, entity.QuestionsUI{
				ID:             question.ID,
				Label:          question.Label,
				Description:    question.Description,
				AnswersOptions: answers,
			})
		}
		response = append(response, &entity.QuizEntityDB{
			ID:        quiz.ID,
			AuthorID:  quiz.AuthorID,
			AccessFor: quiz.AccessFor,
			QuizHash:  quiz.QuizHash,
			Title:     quiz.Title,
			Questions: questionsUI,
			Active:    quiz.Active,
			CreatedAt: quiz.CreatedAt,
			UpdatedAt: quiz.UpdatedAt,
		})
	}

	return response, nil
}

func (r *QuizRepo) GetQuizById(ctx context.Context, quizId int) (*entity.QuizUI, error) {
	var (
		questionsUI []entity.QuestionsUI
		quiz        entity.Quiz
		questions   []entity.Question
	)
	result := r.DB.Table("quizzes").Where("active = ?", true).First(&quiz, quizId)
	if result.Error != nil {
		return nil, fmt.Errorf("quiz repo err %v", result.Error)
	}

	if err := r.DB.Table("questions").Where("quiz_id = ?", quiz.ID).Find(&questions); err.Error != nil {
		return nil, err.Error
	}

	for _, question := range questions {
		var answers []entity.AnswersOption

		if err := r.DB.Table("answers_options").Where("question_id = ?", question.ID).Find(&answers); err.Error != nil {
			return nil, err.Error
		}

		questionsUI = append(questionsUI, entity.QuestionsUI{
			ID:             question.ID,
			Label:          question.Label,
			Description:    question.Description,
			AnswersOptions: answers,
		})
	}

	response := &entity.QuizUI{
		ID:        quiz.ID,
		AuthorID:  quiz.AuthorID,
		QuizHash:  quiz.QuizHash,
		Title:     quiz.Title,
		Questions: questionsUI,
		Active:    quiz.Active,
		CreatedAt: quiz.CreatedAt,
		UpdatedAt: quiz.UpdatedAt,
	}

	return response, nil
}

func (r *QuizRepo) GetQuizByHash(ctx context.Context, quizHash string) (*entity.QuizEntityDB, error) {
	var (
		questionsUI []entity.QuestionsUI
		quiz        entity.Quiz
		questions   []entity.Question
	)
	result := r.DB.Table("quizzes").Where("active = ?", true).Where("quiz_hash = ?", quizHash).First(&quiz)
	if result.Error != nil {
		return nil, fmt.Errorf("quiz repo err %v", result.Error)
	}

	if err := r.DB.Table("questions").Where("quiz_id = ?", quiz.ID).Find(&questions); err.Error != nil {
		return nil, err.Error
	}

	for _, question := range questions {
		var answers []entity.AnswersOption

		if err := r.DB.Table("answers_options").Where("question_id = ?", question.ID).Find(&answers); err.Error != nil {
			return nil, err.Error
		}

		questionsUI = append(questionsUI, entity.QuestionsUI{
			ID:             question.ID,
			Label:          question.Label,
			Description:    question.Description,
			AnswersOptions: answers,
		})
	}

	response := &entity.QuizEntityDB{
		ID:        quiz.ID,
		AuthorID:  quiz.AuthorID,
		QuizHash:  quiz.QuizHash,
		AccessFor: quiz.AccessFor,
		Title:     quiz.Title,
		Questions: questionsUI,
		Active:    quiz.Active,
		CreatedAt: quiz.CreatedAt,
		UpdatedAt: quiz.UpdatedAt,
	}

	return response, nil
}

func (r *QuizRepo) SaveReviewers(ctx context.Context, reviewer *entity.Reviewers) error {
	err := r.DB.Table("reviewers").Save(&reviewer)
	if err.Error != nil {
		return err.Error
	}

	return nil
}
