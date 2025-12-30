package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/S1FFFkA/todo-list/internal/domain"
	"github.com/S1FFFkA/todo-list/internal/dto"
	"github.com/S1FFFkA/todo-list/internal/service"
	"github.com/S1FFFkA/todo-list/pkg/logger"
)

type TaskHandler struct {
	taskService *service.TaskService
}

func NewTaskHandler(taskService *service.TaskService) *TaskHandler {
	return &TaskHandler{
		taskService: taskService,
	}
}

func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	logger.Logger.Info("creating task")

	if r.Method != http.MethodPost {
		logger.Logger.Warn("method not allowed", "method", r.Method)
		h.sendError(w, domain.ErrMethodNotAllowed.Error(), http.StatusMethodNotAllowed)
		return
	}

	var req dto.CreateTaskReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		logger.Logger.Error("failed to decode JSON", "error", err.Error())
		h.sendError(w, domain.ErrFailedToDecodeJSON.Error(), http.StatusInternalServerError)
		return
	}

	if err := req.ValidateForCreate(); err != nil {
		logger.Logger.Warn("validation error", "error", err.Error())
		h.sendError(w, err.Error(), http.StatusBadRequest)
		return
	}

	task := h.taskService.CreateTask(req.Headline, req.Description)

	b, err := json.MarshalIndent(task, "", "    ")
	if err != nil {
		logger.Logger.Error("internal server error", "error", err.Error())
		h.sendError(w, domain.ErrInternalError.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	if _, err := w.Write(b); err != nil {
		logger.Logger.Error("failed to write response", "error", err.Error())
	}
}

func (h *TaskHandler) GetAllTasks(w http.ResponseWriter, r *http.Request) {
	logger.Logger.Info("getting all tasks")

	if r.Method != http.MethodGet {
		h.sendError(w, domain.ErrMethodNotAllowed.Error(), http.StatusMethodNotAllowed)
		return
	}

	tasks := h.taskService.GetAllTasks()

	b, err := json.MarshalIndent(tasks, "", "    ")
	if err != nil {
		logger.Logger.Error("internal server error", "error", err.Error())
		h.sendError(w, domain.ErrInternalError.Error(), http.StatusInternalServerError)
		return
	}

	if _, err := w.Write(b); err != nil {
		logger.Logger.Error("failed to write response", "error", err.Error())
	}
}

func (h *TaskHandler) GetTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		h.sendError(w, domain.ErrMethodNotAllowed.Error(), http.StatusMethodNotAllowed)
		return
	}

	id, err := h.extractID(r)
	if err != nil {
		logger.Logger.Warn("invalid task ID", "error", err.Error())
		h.sendError(w, domain.ErrInvalidRequest.Error(), http.StatusBadRequest)
		return
	}

	logger.Logger.Info("getting task", "task_id", id)

	task, err := h.taskService.GetTask(id)
	if err != nil {
		if err == domain.ErrNotFound {
			logger.Logger.Warn("task not found", "task_id", id)
			h.sendError(w, domain.ErrNotFound.Error(), http.StatusNotFound)
			return
		}
		logger.Logger.Error("internal server error", "error", err.Error())
		h.sendError(w, domain.ErrInternalError.Error(), http.StatusInternalServerError)
		return
	}

	b, err := json.MarshalIndent(task, "", "    ")
	if err != nil {
		h.sendError(w, domain.ErrInternalError.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(b)
}

func (h *TaskHandler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		h.sendError(w, domain.ErrMethodNotAllowed.Error(), http.StatusMethodNotAllowed)
		return
	}

	id, err := h.extractID(r)
	if err != nil {
		logger.Logger.Warn("invalid task ID", "error", err.Error())
		h.sendError(w, domain.ErrInvalidRequest.Error(), http.StatusBadRequest)
		return
	}

	logger.Logger.Info("updating task", "task_id", id)

	var req dto.UpdateTaskReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		logger.Logger.Error("failed to decode JSON", "error", err.Error())
		h.sendError(w, domain.ErrFailedToDecodeJSON.Error(), http.StatusInternalServerError)
		return
	}

	if err := req.ValidateForUpdate(); err != nil {
		logger.Logger.Warn("validation error", "error", err.Error())
		h.sendError(w, err.Error(), http.StatusBadRequest)
		return
	}

	task, err := h.taskService.UpdateContent(id, req.Headline, req.Description)
	if err != nil {
		if err == domain.ErrNotFound {
			logger.Logger.Warn("task not found", "task_id", id)
			h.sendError(w, domain.ErrNotFound.Error(), http.StatusNotFound)
			return
		}
		logger.Logger.Error("internal server error", "error", err.Error())
		h.sendError(w, domain.ErrInternalError.Error(), http.StatusInternalServerError)
		return
	}

	b, err := json.MarshalIndent(task, "", "    ")
	if err != nil {
		logger.Logger.Error("internal server error", "error", err.Error())
		h.sendError(w, domain.ErrInternalError.Error(), http.StatusInternalServerError)
		return
	}

	if _, err := w.Write(b); err != nil {
		logger.Logger.Error("failed to write response", "error", err.Error())
	}
}

func (h *TaskHandler) CompleteTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPatch {
		h.sendError(w, domain.ErrMethodNotAllowed.Error(), http.StatusMethodNotAllowed)
		return
	}

	id, err := h.extractID(r)
	if err != nil {
		logger.Logger.Warn("invalid task ID", "error", err.Error())
		h.sendError(w, domain.ErrInvalidRequest.Error(), http.StatusBadRequest)
		return
	}

	logger.Logger.Info("completing task", "task_id", id)

	task, err := h.taskService.UpdateTask(id)
	if err != nil {
		if err == domain.ErrNotFound {
			logger.Logger.Warn("task not found", "task_id", id)
			h.sendError(w, domain.ErrNotFound.Error(), http.StatusNotFound)
			return
		}
		logger.Logger.Error("internal server error", "error", err.Error())
		h.sendError(w, domain.ErrInternalError.Error(), http.StatusInternalServerError)
		return
	}

	b, err := json.MarshalIndent(task, "", "    ")
	if err != nil {
		logger.Logger.Error("internal server error", "error", err.Error())
		h.sendError(w, domain.ErrInternalError.Error(), http.StatusInternalServerError)
		return
	}

	if _, err := w.Write(b); err != nil {
		logger.Logger.Error("failed to write response", "error", err.Error())
	}
}

func (h *TaskHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		h.sendError(w, domain.ErrMethodNotAllowed.Error(), http.StatusMethodNotAllowed)
		return
	}

	id, err := h.extractID(r)
	if err != nil {
		logger.Logger.Warn("invalid task ID", "error", err.Error())
		h.sendError(w, domain.ErrInvalidRequest.Error(), http.StatusBadRequest)
		return
	}

	logger.Logger.Info("deleting task", "task_id", id)

	err = h.taskService.DeleteTask(id)
	if err != nil {
		if err == domain.ErrNotFound {
			logger.Logger.Warn("task not found", "task_id", id)
			h.sendError(w, domain.ErrNotFound.Error(), http.StatusNotFound)
			return
		}
		logger.Logger.Error("internal server error", "error", err.Error())
		h.sendError(w, domain.ErrInternalError.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *TaskHandler) extractID(r *http.Request) (int, error) {
	idStr := strings.TrimPrefix(r.URL.Path, "/todos/")
	return strconv.Atoi(idStr)
}

func (h *TaskHandler) sendError(w http.ResponseWriter, message string, statusCode int) {
	errDTO := dto.ErrorDTO{
		Message: message,
		Time:    time.Now(),
	}

	w.WriteHeader(statusCode)
	w.Write([]byte(errDTO.ToString()))
}
