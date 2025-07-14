package api

import (
	"rocketseat-go/internal/services"

	"github.com/go-chi/chi/v5"
)

type Application struct {
	Router      *chi.Mux
	TasksService services.TaskService
}
