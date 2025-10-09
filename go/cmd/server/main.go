package main

import (
	"log"
	"net/http"

	"github.com/AlexandrShapkin/todo-api-lab/go/internal/handlers"
	"github.com/AlexandrShapkin/todo-api-lab/go/internal/services"
	"github.com/AlexandrShapkin/todo-api-lab/go/internal/storage"
)

func main() {
	st := storage.NewMemoryStorage()
	authService := services.NewAuthService(st)
	taskService := services.NewTaskService(st)

	authHandler := handlers.NewAuthHandler(authService)
	taskHandler := handlers.NewTaskHandler(taskService)

	mux := http.NewServeMux()

	mux.HandleFunc("/v1/auth/register", authHandler.Register)
	mux.HandleFunc("/v1/auth/login", authHandler.Login)
	mux.HandleFunc("/v1/auth/logout", authHandler.Logout)
	mux.HandleFunc("/v1/auth/me", authHandler.Me)
	mux.HandleFunc("/v1/tasks", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			taskHandler.List(w, r)
		case http.MethodPost:
			taskHandler.Create(w, r)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})
	mux.HandleFunc("/v1/tasks/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			taskHandler.GetByID(w, r)
		case http.MethodPatch:
			taskHandler.Update(w, r)
		case http.MethodPut:
			taskHandler.Replace(w, r)
		case http.MethodDelete:
			taskHandler.Delete(w, r)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
