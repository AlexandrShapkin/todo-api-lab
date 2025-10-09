package storage

import (
	"sync"

	"github.com/AlexandrShapkin/todo-api-lab-go-shared/models"
)

type MemoryStorage struct {
	mu        sync.RWMutex
	users     map[string]*models.User
	tasks     map[string]*models.Task
	userTasks map[string][]string // userID -> tasksIDs
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		users:     make(map[string]*models.User),
		tasks:     make(map[string]*models.Task),
		userTasks: make(map[string][]string),
	}
}

func (s *MemoryStorage) CreateUser(user *models.User) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.users[user.ID] = user
}

func (s *MemoryStorage) GetUserByUsername(username string) *models.User {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, u := range s.users {
		if u.Username == username {
			return u
		}
	}

	return nil
}

func (s *MemoryStorage) GetUserByID(id string) *models.User {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.users[id]
}

func (s *MemoryStorage) CreateTask(userID string, t *models.Task) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.tasks[t.ID] = t
	s.userTasks[userID] = append(s.userTasks[userID], t.ID)
}

func (s *MemoryStorage) GetTasksByUser(userID string) []*models.Task {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var result []*models.Task

	for _, id := range s.userTasks[userID] {
		if t, ok := s.tasks[id]; ok {
			result = append(result, t)
		}
	}

	return result
}

func (s *MemoryStorage) GetTaskByID(id string) *models.Task {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.tasks[id]
}

func (s *MemoryStorage) UpdateTask(t *models.Task) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.tasks[t.ID] = t
}

func (s *MemoryStorage) DeleteTask(id string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.tasks, id)
}
