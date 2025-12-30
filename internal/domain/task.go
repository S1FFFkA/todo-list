package domain

import (
	"time"
)

type Task struct {
	ID          int        `json:"id"`
	Headline    string     `json:"headline"`
	Description string     `json:"description"`
	Done        bool       `json:"done"`
	CreatedAt   time.Time  `json:"created_at"`
	CompletedAt *time.Time `json:"completed_at,omitempty"`
}

func NewTask(id int, headline string, description string) *Task {
	return &Task{
		ID:          id,
		Headline:    headline,
		Description: description,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: nil,
	}
}
