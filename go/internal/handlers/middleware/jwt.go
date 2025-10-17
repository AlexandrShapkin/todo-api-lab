package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/AlexandrShapkin/todo-api-lab-go-shared/auth"
	"github.com/AlexandrShapkin/todo-api-lab-go-shared/utils"
)

type JWT struct {
	handler http.Handler
}

func (j *JWT) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	token := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")
	if token == "" {
		utils.WriteError(w, http.StatusUnauthorized, "missing token")
		return	
	}

	userID, err := auth.ValidateToken(token)
	if err != nil {
		utils.WriteError(w, http.StatusUnauthorized, "invalid token")
		return
	}

	ctx := context.WithValue(r.Context(), UserIDKey, userID)
	r = r.WithContext(ctx)

	j.handler.ServeHTTP(w, r)
}

func NewJWT(handler http.Handler) http.Handler {
	return &JWT{
		handler: handler,
	}
}
