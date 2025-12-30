package domain

import (
	"testing"
)

func TestNewTask(t *testing.T) {
	id := 123
	headline := "Test Headline"
	description := "Test Description"

	task := NewTask(id, headline, description)

	if task == nil {
		t.Fatal("nil")
	}
	if task.ID != id {
		t.Errorf("ID: %d != %d", id, task.ID)
	}
	if task.Headline != headline {
		t.Errorf("headline: %s != %s", headline, task.Headline)
	}
	if task.Description != description {
		t.Errorf("description: %s != %s", description, task.Description)
	}
	if task.Done != false {
		t.Error("done != false")
	}
	if task.CreatedAt.IsZero() {
		t.Error("CreatedAt zero")
	}
	if task.CompletedAt != nil {
		t.Error("CompletedAt != nil")
	}
}
