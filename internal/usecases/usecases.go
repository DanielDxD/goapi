package usecases

import (
	"errors"
	"log/slog"

	"github.com/DanielDxD/go-clean/internal/models"
	"github.com/DanielDxD/go-clean/internal/repositories"
	"github.com/google/uuid"
)

type UseCases struct {
	repos *repositories.Repositories
}

func New(repos *repositories.Repositories) *UseCases {
	return &UseCases{
		repos: repos,
	}
}

func (u UseCases) GetAll() []models.User {
	users := u.repos.User.GetAll()

	return users
}

func (u UseCases) Add(newUser models.CreateUserRequest) (models.User, error) {
	emailExists := u.repos.User.EmailExists(newUser.Email)

	if emailExists {
		slog.Error("Email already exists", "email", newUser.Email)

		return models.User{}, errors.New("email already exists")
	}

	user := models.User{
		ID:    uuid.New(),
		Name:  newUser.Name,
		Email: newUser.Email,
	}

	u.repos.User.Add(user)

	return user, nil
}
