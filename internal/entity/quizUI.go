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
	QuestionID  int    `json:"questionId"`
	Label       string `json:"label"`
	Description string `json:"description,omitempty"`
}

type Question struct {
	ID          int    `json:"id"`
	QuizID      int    `json:"quizId"`
	Label       string `json:"label"`
	Description string `json:"description"`
}

type Quiz struct {
	ID        int       `json:"id" gorm:"id"`
	AuthorID  int       `json:"authorId" gorm:"author_id"`
	Type      string    `json:"type" gorm:"type"`
	QuizHash  string    `json:"quizHash" gorm:"quiz_hash"`
	Title     string    `json:"title" gorm:"title"`
	Active    bool      `json:"active" gorm:"active"`
	CreatedAt time.Time `json:"createdAt" gorm:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"updated_at"`
}

type QuizesResponse struct {
	Success     bool    `json:"success"`
	Description string  `json:"description"`
	Quizes      []*Quiz `json:"quizes"`
}

type QuizResponse struct {
	Success     bool   `json:"success"`
	Description string `json:"description"`
	Quiz        *Quiz  `json:"quiz"`
}
