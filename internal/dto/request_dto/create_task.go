package dto

import (
	"github.com/S1FFFkA/todo-list/internal/domain"
)

type CreateTaskReq struct {
	Id          int
	Headline    string
	Description string
}

func (t *CreateTaskReq) ValidateTask() error {

	if t.Id <= 0 {
		return domain.ErrInvalidRequest
	}

	if t.Headline == "" {
		return domain.ErrInvalidRequest
	}

	return nil
}
