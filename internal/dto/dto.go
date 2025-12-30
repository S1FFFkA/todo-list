package dto

import (
	"encoding/json"
	"time"

	"github.com/S1FFFkA/todo-list/internal/domain"
)

// Request DTO
type CreateTaskReq struct {
	Headline    string `json:"headline"`
	Description string `json:"description"`
}

func (t CreateTaskReq) ValidateForCreate() error {
	if t.Headline == "" {
		return domain.ErrInvalidRequest
	}
	if t.Description == "" {
		return domain.ErrInvalidRequest
	}
	return nil
}

type UpdateTaskReq struct {
	Headline    string `json:"headline"`
	Description string `json:"description"`
}

func (t UpdateTaskReq) ValidateForUpdate() error {
	if t.Headline == "" {
		return domain.ErrInvalidRequest
	}
	return nil
}

// Response DTO
type TaskRes struct {
	ID          int        `json:"id"`
	Headline    string     `json:"headline"`
	Description string     `json:"description"`
	Done        bool       `json:"done"`
	CreatedAt   time.Time  `json:"created_at"`
	CompletedAt *time.Time `json:"completed_at,omitempty"`
}

type TaskListRes struct {
	Tasks []TaskRes `json:"tasks"`
}

type ErrorDTO struct {
	Message string    `json:"message"`
	Time    time.Time `json:"time"`
}

func (e ErrorDTO) ToString() string {
	b, err := json.MarshalIndent(e, "", "    ")
	if err != nil {
		panic(err)
	}
	return string(b)
}
