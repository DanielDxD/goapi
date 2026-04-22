package handlers

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/DanielDxD/go-clean/internal/usecases"
)

type Handlers struct {
	useCases *usecases.UseCases
}

func New(useCases *usecases.UseCases) *Handlers {
	return &Handlers{
		useCases: useCases,
	}
}

func (h Handlers) Listen(port int) error {

	h.registerUserEndpoints()

	slog.Info(fmt.Sprintf("Server running on http://localhost:%v", port))
	return http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
}
