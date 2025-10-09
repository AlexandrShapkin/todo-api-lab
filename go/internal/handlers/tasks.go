package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/AlexandrShapkin/todo-api-lab-go-shared/auth"
	"github.com/AlexandrShapkin/todo-api-lab-go-shared/models"
	"github.com/AlexandrShapkin/todo-api-lab-go-shared/utils"
	"github.com/AlexandrShapkin/todo-api-lab/go/internal/services"
)

type TaskHandler struct {
	service *services.TaskService
}

func NewTaskHandler(s *services.TaskService) *TaskHandler {
	return &TaskHandler{
		service: s,
	}
}

func (h *TaskHandler) Create(w http.ResponseWriter, r *http.Request) {
	token := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")
	userID, err := auth.ValidateToken(token)
	if err != nil {
		utils.WriteError(w, http.StatusUnauthorized, "invalid token")
		return
	}

	var body models.RawTask
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		utils.WriteError(w, http.StatusBadRequest, "invalid json")
		return
	}

	task := h.service.Create(userID, body)
	utils.WriteJSON(w, http.StatusCreated, task)
}

func (h *TaskHandler) List(w http.ResponseWriter, r *http.Request) {
	token := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")
	userID, err := auth.ValidateToken(token)
	if err != nil {
		utils.WriteError(w, http.StatusUnauthorized, "invalid token")
		return
	}

	tasks := h.service.GetAll(userID)
	utils.WriteJSON(w, http.StatusOK, tasks)
}

func (h *TaskHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/v1/tasks/")
	token := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")
	_, err := auth.ValidateToken(token)
	if err != nil {
		utils.WriteError(w, http.StatusUnauthorized, "invalid token")
		return
	}

	task := h.service.GetByID(id)
	if task == nil {
		utils.WriteError(w, http.StatusNotFound, "task not found")
		return
	}

	utils.WriteJSON(w, http.StatusOK, task)
}

func (h *TaskHandler) Update(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/v1/tasks/")
	token := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")
	_, err := auth.ValidateToken(token)
	if err != nil {
		utils.WriteError(w, http.StatusUnauthorized, "invalid token")
		return
	}

	var body models.OptionalTaskRows
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		utils.WriteError(w, http.StatusBadRequest, "invalid json")
		return
	}

	task, err := h.service.Update(id, body)
	if err != nil {
		utils.WriteError(w, http.StatusNotFound, err.Error())
		return
	}

	utils.WriteJSON(w, http.StatusOK, task)
}

func (h *TaskHandler) Replace(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/v1/tasks/")
	token := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")
	_, err := auth.ValidateToken(token)
	if err != nil {
		utils.WriteError(w, http.StatusUnauthorized, "invalid token")
		return
	}

	var body models.RawTask
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		utils.WriteError(w, http.StatusBadRequest, "invalid json")
		return
	}

	task, err := h.service.Update(id, models.OptionalTaskRows{
		Title:       &body.Title,
		Description: &body.Description,
		DueTime:     &body.DueTime,
		Completed:   &body.Completed,
	})
	if err != nil {
		utils.WriteError(w, http.StatusNotFound, err.Error())
		return
	}

	utils.WriteJSON(w, http.StatusOK, task)
}

func (h *TaskHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/v1/tasks/")
	token := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")
	_, err := auth.ValidateToken(token)
	if err != nil {
		utils.WriteError(w, http.StatusUnauthorized, "invalid token")
		return
	}

	h.service.Delete(id)
	w.WriteHeader(http.StatusNoContent)
}
