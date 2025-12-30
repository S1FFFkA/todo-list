package app

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/S1FFFkA/todo-list/internal/handlers"
	"github.com/S1FFFkA/todo-list/internal/server"
	"github.com/S1FFFkA/todo-list/internal/service"
	"github.com/S1FFFkA/todo-list/pkg/logger"
)

func Run() {
	// Инициализация логгера
	if err := logger.InitLogger(); err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}

	taskService := service.NewTaskService()
	taskHandler := handlers.NewTaskHandler(taskService)
	router := server.NewRouter(taskHandler)

	port := ":8080"
	srv := &http.Server{
		Addr:    port,
		Handler: router,
	}

	logger.Logger.Info("starting server", "port", port)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Logger.Error("server failed to start", "error", err.Error())
			log.Fatalf("Server failed to start: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Logger.Info("shutting down server", "timeout", "5s")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Logger.Error("server forced to shutdown", "error", err.Error())
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	logger.Logger.Info("server exited gracefully")
}
