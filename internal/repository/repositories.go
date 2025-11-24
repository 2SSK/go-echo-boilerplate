package repository

import "github.com/2SSK/boilerplate/internal/server"

type Repositories struct {
}

func NewRepositories(s *server.Server) *Repositories {
	return &Repositories{}
}
