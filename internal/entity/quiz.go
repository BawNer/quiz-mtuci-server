package entity

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

type JSONBQuestions []Question

func (j JSONBQuestions) Value() (driver.Value, error) {
	valueString, err := json.MarshalIndent(j, "", "\t")
	return string(valueString), err
}

func (j *JSONBQuestions) Scan(value interface{}) error {
	var data []byte
	switch v := value.(type) {
	case string:
		data = []byte(v)
	case []byte:
		data = v
	}
	return json.Unmarshal(data, &j)
}

type AnswerOption struct {
	ID          int    `json:"id"`
	QuestionID  int    `json:"-"`
	Label       string `json:"label"`
	Description string `json:"description,omitempty"`
}

type Question struct {
	ID          int    `json:"id"`
	QuizID      int    `json:"-"`
	Label       string `json:"label"`
	Description string `json:"description"`
}

type Quiz struct {
	ID        int       `json:"id"`
	AuthorID  int       `json:"authorId"`
	Type      string    `json:"type"`
	QuizHash  string    `json:"quizHash"`
	Title     string    `json:"title"`
	Active    bool      `json:"active"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type QuizEntityDB struct {
	ID                  int       `gorm:"id"`
	AuthorID            int       `gorm:"author_id"`
	Type                string    `gorm:"type"`
	QuizHash            string    `gorm:"quiz_hash"`
	Title               string    `gorm:"title"`
	Active              bool      `gorm:"active"`
	CreatedAt           time.Time `gorm:"created_at"`
	UpdatedAt           time.Time `gorm:"updated_at"`
	QuestionID          int       `gorm:"question_id"`
	QuestionQuizID      int       `gorm:"question_quiz_id"`
	QuestionLabel       string    `gorm:"question_label"`
	QuestionDescription string    `gorm:"question_description"`
	AnswerID            int       `gorm:"answer_id"`
	AnswerQuestionID    int       `gorm:"answer_question_id"`
	AnswerLabel         string    `gorm:"answer_label"`
	AnswerDescription   string    `gorm:"answer_description"`
}
