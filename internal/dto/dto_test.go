package dto

import (
	"strings"
	"testing"
	"time"

	"github.com/S1FFFkA/todo-list/internal/domain"
)

func TestCreateTaskReq_ValidateForCreate_Success(t *testing.T) {
	req := CreateTaskReq{
		Headline:    "Test Headline",
		Description: "Test Description",
	}

	err := req.ValidateForCreate()
	if err != nil {
		t.Errorf("error: %v", err)
	}
}

func TestCreateTaskReq_ValidateForCreate_EmptyHeadline(t *testing.T) {
	req := CreateTaskReq{
		Headline:    "",
		Description: "Test Description",
	}

	err := req.ValidateForCreate()
	if err == nil {
		t.Fatal("no error")
	}
	if err != domain.ErrInvalidRequest {
		t.Errorf("want ErrInvalidRequest, got %v", err)
	}
}

func TestCreateTaskReq_ValidateForCreate_EmptyDescription(t *testing.T) {
	req := CreateTaskReq{
		Headline:    "Test Headline",
		Description: "",
	}

	err := req.ValidateForCreate()
	if err == nil {
		t.Fatal("no error")
	}
	if err != domain.ErrInvalidRequest {
		t.Errorf("want ErrInvalidRequest, got %v", err)
	}
}

func TestCreateTaskReq_ValidateForCreate_BothEmpty(t *testing.T) {
	req := CreateTaskReq{
		Headline:    "",
		Description: "",
	}

	err := req.ValidateForCreate()
	if err == nil {
		t.Fatal("no error")
	}
	if err != domain.ErrInvalidRequest {
		t.Errorf("want ErrInvalidRequest, got %v", err)
	}
}

func TestUpdateTaskReq_ValidateForUpdate_Success(t *testing.T) {
	req := UpdateTaskReq{
		Headline:    "Test Headline",
		Description: "Test Description",
	}

	err := req.ValidateForUpdate()
	if err != nil {
		t.Errorf("error: %v", err)
	}
}

func TestUpdateTaskReq_ValidateForUpdate_EmptyHeadline(t *testing.T) {
	req := UpdateTaskReq{
		Headline:    "",
		Description: "Test Description",
	}

	err := req.ValidateForUpdate()
	if err == nil {
		t.Fatal("no error")
	}
	if err != domain.ErrInvalidRequest {
		t.Errorf("want ErrInvalidRequest, got %v", err)
	}
}

func TestUpdateTaskReq_ValidateForUpdate_EmptyDescription(t *testing.T) {
	req := UpdateTaskReq{
		Headline:    "Test Headline",
		Description: "",
	}

	err := req.ValidateForUpdate()
	if err != nil {
		t.Errorf("description can be empty, got error: %v", err)
	}
}

func TestUpdateTaskReq_ValidateForUpdate_BothEmpty(t *testing.T) {
	req := UpdateTaskReq{
		Headline:    "",
		Description: "",
	}

	err := req.ValidateForUpdate()
	if err == nil {
		t.Fatal("no error")
	}
	if err != domain.ErrInvalidRequest {
		t.Errorf("want ErrInvalidRequest, got %v", err)
	}
}

func TestErrorDTO_ToString(t *testing.T) {
	errDTO := ErrorDTO{
		Message: "Test error message",
		Time:    time.Now(),
	}

	result := errDTO.ToString()
	if result == "" {
		t.Fatal("empty")
	}

	if !strings.Contains(result, "Test error message") {
		t.Error("no message")
	}
}
