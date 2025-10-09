package services

import (
	"errors"
	"time"

	"github.com/AlexandrShapkin/todo-api-lab-go-shared/models"
	"github.com/AlexandrShapkin/todo-api-lab/go/internal/storage"
	"github.com/google/uuid"
)

type TaskService struct {
	storage *storage.MemoryStorage
}

func NewTaskService(s *storage.MemoryStorage) *TaskService {
	return &TaskService{
		storage: s,
	}
}

func (t *TaskService) Create(userID string, input models.RawTask) *models.Task {
	task := &models.Task{
		ID:          uuid.New().String(),
		Title:       input.Title,
		Description: input.Description,
		CreatedAt:   time.Now(),
		DueTime:     input.DueTime,
		Completed:   input.Completed,
	}
	
	t.storage.CreateTask(userID, task)

	return task
}

func (t *TaskService) GetAll(userID string) []*models.Task {
	return t.storage.GetTasksByUser(userID)
}

func (t *TaskService) GetByID(id string) *models.Task {
	return t.storage.GetTaskByID(id)
}

func (t *TaskService) Update(id string, input models.OptionalTaskRows) (*models.Task, error) {
	task := t.storage.GetTaskByID(id)
	if task == nil {
		return nil, errors.New("task not found")
	}

	if input.Title != nil {
		task.Title = *input.Title
	}

	if input.Description != nil {
		task.Description = *input.Description
	}

	if input.DueTime != nil {
		task.DueTime = *input.DueTime
	}

	if input.Completed != nil {
		task.Completed = *input.Completed
	}

	t.storage.UpdateTask(task)

	return task, nil
}

func (t *TaskService) Delete(id string) {
	t.storage.DeleteTask(id)
}
