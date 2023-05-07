package entity

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

type JSONB map[string]interface{}

func (j JSONB) Value() (driver.Value, error) {
	valueString, err := json.MarshalIndent(j, "", "\t")
	return string(valueString), err
}

func (j *JSONB) Scan(value interface{}) error {
	if err := json.Unmarshal(value.([]byte), &j); err != nil {
		return err
	}
	return nil
}

type Answer struct {
	ID          int    `json:"id"`
	Label       string `json:"label"`
	Description string `json:"description,omitempty"`
	IsTrue      bool   `json:"isTrue"`
}

type Question struct {
	ID          int      `json:"id"`
	Label       string   `json:"label"`
	Description string   `json:"description"`
	Answers     []Answer `json:"answers"`
}

type Quiz struct {
	ID        int       `json:"id" gorm:"id"`
	AuthorID  int       `json:"authorId" gorm:"author_id"`
	QuizHash  string    `json:"quizHash" gorm:"quiz_hash"`
	Title     string    `json:"title" gorm:"title"`
	Questions JSONB     `json:"questions" gorm:"questions,type:jsonb"`
	Active    bool      `json:"active" gorm:"active"`
	CreatedAt time.Time `json:"createdAt" gorm:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"updated_at"`
}

type QuizesResponse struct {
	Success     bool   `json:"success"`
	Description string `json:"description"`
	Quizes      []Quiz `json:"quizes"`
}

type QuizResponse struct {
	Success     bool   `json:"success"`
	Description string `json:"description"`
	Quiz        *Quiz  `json:"quiz"`
}
