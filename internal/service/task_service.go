package service

import (
	"crypto/rand"
	"encoding/binary"
	"sync"
	"time"

	"github.com/S1FFFkA/todo-list/internal/domain"
)

type TaskService struct {
	tasks map[int]*domain.Task
	mtx   sync.RWMutex
}

func NewTaskService() *TaskService {
	return &TaskService{
		tasks: make(map[int]*domain.Task),
	}
}

func (s *TaskService) generateID() int {
	for {
		var b [8]byte
		if _, err := rand.Read(b[:]); err != nil {
			continue
		}
		id := int(binary.BigEndian.Uint64(b[:]))
		if id < 0 {
			id = -id
		}
		// Генерируем 8-значное число от 1 до 99999999 (от 00000001 до 99999999)
		id = 1 + id%99999999
		if _, exists := s.tasks[id]; !exists {
			return id
		}
	}
}

func (s *TaskService) CreateTask(headline string, description string) *domain.Task {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	id := s.generateID()
	task := domain.NewTask(id, headline, description)

	s.tasks[id] = task
	return task
}

func (s *TaskService) GetAllTasks() []*domain.Task {
	s.mtx.RLock()
	defer s.mtx.RUnlock()

	tasks := make([]*domain.Task, 0, len(s.tasks))
	for _, task := range s.tasks {
		tasks = append(tasks, task)
	}

	return tasks
}

func (s *TaskService) GetTask(id int) (*domain.Task, error) {
	s.mtx.RLock()
	defer s.mtx.RUnlock()

	task, ok := s.tasks[id]
	if !ok {
		return nil, domain.ErrNotFound
	}

	return task, nil
}

func (s *TaskService) UpdateTask(id int) (*domain.Task, error) {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	task, ok := s.tasks[id]
	if !ok {
		return nil, domain.ErrNotFound
	}

	completeTime := time.Now()
	task.Done = true
	task.CompletedAt = &completeTime

	return task, nil
}

func (s *TaskService) UpdateContent(id int, headline string, description string) (*domain.Task, error) {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	task, ok := s.tasks[id]
	if !ok {
		return nil, domain.ErrNotFound
	}

	task.Headline = headline
	task.Description = description

	return task, nil
}

func (s *TaskService) DeleteTask(id int) error {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	_, ok := s.tasks[id]
	if !ok {
		return domain.ErrNotFound
	}

	delete(s.tasks, id)
	return nil
}
