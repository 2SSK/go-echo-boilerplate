package handler

import (
	"github.com/2SSK/boilerplate/internal/server"
	"github.com/2SSK/boilerplate/internal/service"
)

type Handlers struct {
	Health  *HealthHandler
	OpenAPI *OpenAPIHandler
	Home    *HomeHandler
}

func NewHandlers(s *server.Server, services *service.Services) *Handlers {
	return &Handlers{
		Health:  NewHealthHandler(s),
		OpenAPI: NewOpenAPIHandler(s),
		Home:    NewHomeHandler(s),
	}
}
