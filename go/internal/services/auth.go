package services

import (
	"errors"
	"fmt"

	"github.com/AlexandrShapkin/todo-api-lab-go-shared/auth"
	"github.com/AlexandrShapkin/todo-api-lab-go-shared/models"
	"github.com/AlexandrShapkin/todo-api-lab/go/internal/storage"
	"github.com/google/uuid"
)

type AuthService struct {
	storage *storage.MemoryStorage
}

func NewAuthService(s *storage.MemoryStorage) *AuthService {
	return &AuthService{
		storage: s,
	}
}

func (a *AuthService) Register(username, password string) (*models.AuthResponse, error) {
	if a.storage.GetUserByUsername(username) != nil {
		return nil, errors.New("user already exists")
	}

	user := &models.User{
		ID:       uuid.New().String(),
		Username: username,
		Password: password,
	}
	a.storage.CreateUser(user)

	access, aexp, refresh, rexp, err := auth.GenerateTokenPair(user.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to generate token pair: %v", err)
	}

	return &models.AuthResponse{
		Username:         username,
		UserID:           user.ID,
		AccessToken:      access,
		AccessExpiresIn:  aexp,
		RefreshToken:     refresh,
		RefreshExpiresIn: rexp,
	}, nil
}

func (a *AuthService) Login(username, password string) (*models.AuthResponse, error) {
	user := a.storage.GetUserByUsername(username)
	if user == nil || user.Password != password {
		return nil, errors.New("invalid credentials")
	}

	access, aexp, refresh, rexp, err := auth.GenerateTokenPair(user.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to generate token pair: %v", err)
	}

	return &models.AuthResponse{
		Username:         username,
		UserID:           user.ID,
		AccessToken:      access,
		AccessExpiresIn:  aexp,
		RefreshToken:     refresh,
		RefreshExpiresIn: rexp,
	}, nil
}

func (a *AuthService) GetMe(userID string) *models.User {
	return a.storage.GetUserByID(userID)
}
