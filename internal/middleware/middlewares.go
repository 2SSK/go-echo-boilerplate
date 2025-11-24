package middleware

import (
	"github.com/2SSK/boilerplate/internal/server"
)

type Middlewares struct {
	Global    *GlobalMiddlewares
	RateLimit *RateLimitMiddleware
	Context   *ContextEnhancer
}

func NewMiddlewares(s *server.Server) *Middlewares {

	return &Middlewares{
		Global:    NewGlobalMiddlewares(s),
		RateLimit: NewRateLimitMiddleware(s),
		Context:   NewContextEnhancer(s),
	}
}
