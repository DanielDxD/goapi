package repositories

import (
	"github.com/DanielDxD/go-clean/internal/models"
	"github.com/DanielDxD/go-clean/internal/repositories/users"
)

type Repositories struct {
	User interface {
		GetAll() []models.User
		Add(newUser models.User)
		EmailExists(email string) bool
	}
}

func New() *Repositories {
	return &Repositories{
		User: users.New(),
	}
}
