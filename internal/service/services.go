package service

import (
	"github.com/2SSK/boilerplate/internal/repository"
	"github.com/2SSK/boilerplate/internal/server"
)

type Services struct {
}

func NewServices(s *server.Server, repos *repository.Repositories) (*Services, error) {

	return &Services{}, nil
}
