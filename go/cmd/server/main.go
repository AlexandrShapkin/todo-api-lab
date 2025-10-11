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

	mux.Handle("/v1/auth/register", handlers.Logging(http.HandlerFunc(authHandler.Register)))
	mux.Handle("/v1/auth/login", handlers.Logging(http.HandlerFunc(authHandler.Login)))
	mux.Handle("/v1/auth/logout", handlers.Logging(http.HandlerFunc(authHandler.Logout)))
	mux.Handle("/v1/auth/me", handlers.Logging(http.HandlerFunc(authHandler.Me)))
	mux.Handle("/v1/tasks", handlers.Logging(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			taskHandler.List(w, r)
		case http.MethodPost:
			taskHandler.Create(w, r)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})))
	mux.Handle("/v1/tasks/", handlers.Logging(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
	})))

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
