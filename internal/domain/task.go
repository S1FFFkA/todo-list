package domain

import (
	"time"
)

type Task struct {
	id          int
	headline    string
	description string
	done        bool
	createdAt   time.Time
	completedAt *time.Time
}
