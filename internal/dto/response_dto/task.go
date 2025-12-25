package dto

import "time"

type TaskRes struct {
	Id          int
	Headline    string
	Description string
	CreatedAt   time.Time
	CompletedAt *time.Time
}
