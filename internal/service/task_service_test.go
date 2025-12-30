package service

import (
	"testing"

	"github.com/S1FFFkA/todo-list/internal/domain"
)

func TestNewTaskService(t *testing.T) {
	service := NewTaskService()
	if service == nil {
		t.Fatal("nil")
	}
	if service.tasks == nil {
		t.Fatal("tasks nil")
	}
	if len(service.tasks) != 0 {
		t.Fatal("not empty")
	}
}

func TestCreateTaskSuccess(t *testing.T) {
	service := NewTaskService()
	headline := "Test Task"
	description := "Test Description"

	task := service.CreateTask(headline, description)

	if task == nil {
		t.Fatal("nil")
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
	if task.ID <= 0 {
		t.Errorf("ID <= 0: %d", task.ID)
	}
	if task.ID > 99999999 {
		t.Errorf("ID > 99999999: %d", task.ID)
	}
	if task.CreatedAt.IsZero() {
		t.Error("CreatedAt zero")
	}
	if task.CompletedAt != nil {
		t.Error("CompletedAt != nil")
	}

	allTasks := service.GetAllTasks()
	if len(allTasks) != 1 {
		t.Errorf("len != 1: %d", len(allTasks))
	}
}

func TestCreateTaskUniqueIDs(t *testing.T) {
	service := NewTaskService()
	ids := make(map[int]bool)

	for i := 0; i < 100; i++ {
		task := service.CreateTask("Task", "Description")
		if ids[task.ID] {
			t.Errorf("duplicate: %d", task.ID)
		}
		ids[task.ID] = true
	}

	if len(ids) != 100 {
		t.Errorf("len != 100: %d", len(ids))
	}
}

func TestGetAllTasksEmpty(t *testing.T) {
	service := NewTaskService()
	tasks := service.GetAllTasks()

	if tasks == nil {
		t.Fatal("nil")
	}
	if len(tasks) != 0 {
		t.Errorf("len != 0: %d", len(tasks))
	}
}

func TestGetAllTasksMultiple(t *testing.T) {
	service := NewTaskService()

	service.CreateTask("Task 1", "Description 1")
	service.CreateTask("Task 2", "Description 2")
	service.CreateTask("Task 3", "Description 3")

	tasks := service.GetAllTasks()
	if len(tasks) != 3 {
		t.Errorf("len != 3: %d", len(tasks))
	}
}

func TestGetTaskSuccess(t *testing.T) {
	service := NewTaskService()
	createdTask := service.CreateTask("Test Task", "Test Description")

	retrievedTask, err := service.GetTask(createdTask.ID)
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	if retrievedTask == nil {
		t.Fatal("nil")
	}
	if retrievedTask.ID != createdTask.ID {
		t.Errorf("ID: %d != %d", createdTask.ID, retrievedTask.ID)
	}
	if retrievedTask.Headline != createdTask.Headline {
		t.Errorf("headline: %s != %s", createdTask.Headline, retrievedTask.Headline)
	}
}

func TestGetTaskNotFound(t *testing.T) {
	service := NewTaskService()

	_, err := service.GetTask(99999)
	if err == nil {
		t.Fatal("no error")
	}
	if err != domain.ErrNotFound {
		t.Errorf("want ErrNotFound, got %v", err)
	}
}

func TestUpdateTaskSuccess(t *testing.T) {
	service := NewTaskService()
	createdTask := service.CreateTask("Test Task", "Test Description")

	updatedTask, err := service.UpdateTask(createdTask.ID)
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	if updatedTask == nil {
		t.Fatal("nil")
	}
	if !updatedTask.Done {
		t.Error("done != true")
	}
	if updatedTask.CompletedAt == nil {
		t.Error("CompletedAt nil")
	}
	if updatedTask.Headline != createdTask.Headline {
		t.Error("headline changed")
	}
	if updatedTask.Description != createdTask.Description {
		t.Error("description changed")
	}
}

func TestUpdateTaskNotFound(t *testing.T) {
	service := NewTaskService()

	_, err := service.UpdateTask(99999)
	if err == nil {
		t.Fatal("no error")
	}
	if err != domain.ErrNotFound {
		t.Errorf("want ErrNotFound, got %v", err)
	}
}

func TestUpdateContentSuccess(t *testing.T) {
	service := NewTaskService()
	createdTask := service.CreateTask("Old Headline", "Old Description")

	newHeadline := "New Headline"
	newDescription := "New Description"

	updatedTask, err := service.UpdateContent(createdTask.ID, newHeadline, newDescription)
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	if updatedTask == nil {
		t.Fatal("nil")
	}
	if updatedTask.Headline != newHeadline {
		t.Errorf("headline: %s != %s", newHeadline, updatedTask.Headline)
	}
	if updatedTask.Description != newDescription {
		t.Errorf("description: %s != %s", newDescription, updatedTask.Description)
	}
	if updatedTask.Done != createdTask.Done {
		t.Error("done changed")
	}
}

func TestUpdateContentNotFound(t *testing.T) {
	service := NewTaskService()

	_, err := service.UpdateContent(99999, "Headline", "Description")
	if err == nil {
		t.Fatal("no error")
	}
	if err != domain.ErrNotFound {
		t.Errorf("want ErrNotFound, got %v", err)
	}
}

func TestDeleteTaskSuccess(t *testing.T) {
	service := NewTaskService()
	createdTask := service.CreateTask("Test Task", "Test Description")

	err := service.DeleteTask(createdTask.ID)
	if err != nil {
		t.Fatalf("error: %v", err)
	}

	_, err = service.GetTask(createdTask.ID)
	if err == nil {
		t.Fatal("not deleted")
	}
	if err != domain.ErrNotFound {
		t.Errorf("want ErrNotFound, got %v", err)
	}

	tasks := service.GetAllTasks()
	if len(tasks) != 0 {
		t.Errorf("len != 0: %d", len(tasks))
	}
}

func TestDeleteTaskNotFound(t *testing.T) {
	service := NewTaskService()

	err := service.DeleteTask(99999)
	if err == nil {
		t.Fatal("no error")
	}
	if err != domain.ErrNotFound {
		t.Errorf("want ErrNotFound, got %v", err)
	}
}

func TestTaskServiceConcurrency(t *testing.T) {
	service := NewTaskService()
	done := make(chan bool)

	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 10; j++ {
				service.CreateTask("Task", "Description")
			}
			done <- true
		}()
	}

	for i := 0; i < 10; i++ {
		<-done
	}

	tasks := service.GetAllTasks()
	if len(tasks) != 100 {
		t.Errorf("len != 100: %d", len(tasks))
	}
}
