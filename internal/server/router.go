package server

import (
	"net/http"
	"strings"
	"time"

	"github.com/S1FFFkA/todo-list/internal/domain"
	"github.com/S1FFFkA/todo-list/internal/dto"
	"github.com/S1FFFkA/todo-list/internal/handlers"
)

func NewRouter(taskHandler *handlers.TaskHandler) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			taskHandler.CreateTask(w, r)
		case http.MethodGet:
			taskHandler.GetAllTasks(w, r)
		default:
			sendError(w, domain.ErrMethodNotAllowed.Error(), http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/todos/", func(w http.ResponseWriter, r *http.Request) {
		path := strings.TrimPrefix(r.URL.Path, "/todos/")
		if path == "" {
			sendError(w, domain.ErrInvalidRequest.Error(), http.StatusBadRequest)
			return
		}

		switch r.Method {
		case http.MethodGet:
			taskHandler.GetTask(w, r)
		case http.MethodPut:
			taskHandler.UpdateTask(w, r)
		case http.MethodPatch:
			taskHandler.CompleteTask(w, r)
		case http.MethodDelete:
			taskHandler.DeleteTask(w, r)
		default:
			sendError(w, domain.ErrMethodNotAllowed.Error(), http.StatusMethodNotAllowed)
		}
	})

	return mux
}

func sendError(w http.ResponseWriter, message string, statusCode int) {
	errDTO := dto.ErrorDTO{
		Message: message,
		Time:    time.Now(),
	}

	w.WriteHeader(statusCode)
	if _, err := w.Write([]byte(errDTO.ToString())); err != nil {
	}
}
