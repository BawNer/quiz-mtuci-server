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

func (r *ServiceRepo) GetAllQuiz(ctx context.Context) ([]*entity.QuizUI, error) {
	var (
		responseDB []entity.QuizEntityDB
		response   []*entity.QuizUI
	)

	err := r.DB.Table("answers_options").Select(`
		quizzes.id, quizzes.author_id, quizzes.type, quizzes.quiz_hash,
		quizzes.active, quizzes.created_at, quizzes.updated_at, quizzes.title,
		questions.id as question_id, questions.label as question_label, questions.description as question_description, questions.quiz_id as question_quiz_id,
		answers_options.id as answer_id, answers_options.label as answer_label, answers_options.description as answer_description, answers_options.question_id as answer_question_id
	`).Joins(`left join questions ON answers_options.question_id = questions.id`).Joins(`left join quizzes ON questions.quiz_id = quizzes.id`).Scan(&responseDB)

	if err.Error != nil {
		return nil, err.Error
	}

	iter := 0
	for _, row := range responseDB {
		if response == nil {
			var (
				question []entity.QuestionsUI
				answers  []*entity.AnswerOption
			)
			answers = append(answers, &entity.AnswerOption{
				ID:          row.AnswerID,
				QuestionID:  row.QuestionID,
				Label:       row.AnswerLabel,
				Description: row.AnswerDescription,
			})
			question = append(question, entity.QuestionsUI{
				ID:             row.QuestionID,
				Label:          row.QuestionLabel,
				Description:    row.QuestionDescription,
				AnswersOptions: answers,
			})
			response = append(response, &entity.QuizUI{
				ID:        row.ID,
				AuthorID:  row.AuthorID,
				Type:      row.Type,
				QuizHash:  row.QuizHash,
				Title:     row.Title,
				Questions: question,
				Active:    row.Active,
				CreatedAt: row.CreatedAt,
				UpdatedAt: row.UpdatedAt,
			})
		} else if response[iter].ID == row.QuestionQuizID {
			// добавляем в текущий вопрос ответы
			qiter := len(response[iter].Questions) - 1
			response[iter].Questions[qiter].AnswersOptions = append(response[iter].Questions[qiter].AnswersOptions, &entity.AnswerOption{
				ID:          row.AnswerID,
				QuestionID:  row.QuestionID,
				Label:       row.AnswerLabel,
				Description: row.AnswerDescription,
			})
		} else {
			// переходим на некст итерацию
			var (
				question []entity.QuestionsUI
				answers  []*entity.AnswerOption
			)
			answers = append(answers, &entity.AnswerOption{
				ID:          row.AnswerID,
				QuestionID:  row.QuestionID,
				Label:       row.AnswerLabel,
				Description: row.AnswerDescription,
			})
			question = append(question, entity.QuestionsUI{
				ID:             row.QuestionID,
				Label:          row.QuestionLabel,
				Description:    row.QuestionDescription,
				AnswersOptions: answers,
			})
			response = append(response, &entity.QuizUI{
				ID:        row.ID,
				AuthorID:  row.AuthorID,
				Type:      row.Type,
				QuizHash:  row.QuizHash,
				Title:     row.Title,
				Questions: question,
				Active:    row.Active,
				CreatedAt: row.CreatedAt,
				UpdatedAt: row.UpdatedAt,
			})
			iter++
		}
	}

	return response, nil
}

func (r *ServiceRepo) GetQuizById(ctx context.Context, quizId int) (*entity.QuizUI, error) {
	var (
		questionsUI []entity.QuestionsUI
		quiz        *entity.Quiz
		questions   []*entity.Question
	)
	result := r.DB.Table("quizzes").First(&quiz, quizId)
	if result.Error != nil {
		return nil, fmt.Errorf("quiz repo err %v", result.Error)
	}

	if err := r.DB.Table("questions").Where("quiz_id = ?", quiz.ID).Find(&questions); err.Error != nil {
		return nil, err.Error
	}

	for _, question := range questions {
		var answers []*entity.AnswerOption

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
		Type:      quiz.Type,
		QuizHash:  quiz.QuizHash,
		Title:     quiz.Title,
		Questions: questionsUI,
		Active:    quiz.Active,
		CreatedAt: quiz.CreatedAt,
		UpdatedAt: quiz.UpdatedAt,
	}

	return response, nil
}

func (r *ServiceRepo) SaveQuiz(ctx context.Context, quiz *entity.QuizUI) (*entity.QuizUI, error) {
	var (
		questions []entity.QuestionsUI
		answers   []*entity.AnswerOption
	)

	newQuiz := entity.Quiz{
		AuthorID: quiz.AuthorID,
		Type:     quiz.Type,
		QuizHash: uuid.New().String(),
		Title:    quiz.Title,
		Active:   quiz.Active,
	}
	if createQuiz := r.DB.Table("quizzes").Create(&newQuiz); createQuiz.Error != nil {
		return nil, createQuiz.Error
	}

	// добавляем вопросы к квизу
	for _, question := range quiz.Questions {
		newQuestions := entity.Question{
			QuizID:      newQuiz.ID,
			Label:       question.Label,
			Description: question.Description,
		}
		if createQuestion := r.DB.Table("questions").Create(&newQuestions); createQuestion.Error != nil {
			return nil, createQuestion.Error
		}
		// добавляем варианты ответа
		for _, answer := range question.AnswersOptions {
			newAnswerOption := entity.AnswerOption{
				QuestionID:  newQuestions.ID,
				Label:       answer.Label,
				Description: answer.Description,
			}
			if createAnswerOption := r.DB.Table("answers_options").Create(&newAnswerOption); createAnswerOption.Error != nil {
				return nil, createAnswerOption.Error
			}

			answers = append(answers, &entity.AnswerOption{
				ID:          newAnswerOption.ID,
				QuestionID:  newQuestions.ID,
				Label:       newAnswerOption.Label,
				Description: newAnswerOption.Description,
			})
		}

		questions = append(questions, entity.QuestionsUI{
			ID:             newQuestions.ID,
			Label:          newQuestions.Label,
			Description:    newQuestions.Description,
			AnswersOptions: answers,
		})
	}

	createdQuiz := &entity.QuizUI{
		ID:        newQuiz.ID,
		AuthorID:  newQuiz.AuthorID,
		Type:      newQuiz.Type,
		QuizHash:  newQuiz.QuizHash,
		Title:     newQuiz.Title,
		Questions: questions,
		Active:    newQuiz.Active,
		CreatedAt: newQuiz.CreatedAt,
		UpdatedAt: newQuiz.UpdatedAt,
	}

	return createdQuiz, nil
}
