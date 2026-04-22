package main

import (
	"github.com/DanielDxD/go-clean/internal/handlers"
	"github.com/DanielDxD/go-clean/internal/repositories"
	"github.com/DanielDxD/go-clean/internal/usecases"
)

func main() {
	repos := repositories.New()
	useCases := usecases.New(repos)

	handler := handlers.New(useCases)

	handler.Listen(3030)
}
