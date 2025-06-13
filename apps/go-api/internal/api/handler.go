package api

import (
	"github.com/labstack/echo/v4"
)

type Server struct{}

func NewServer() *Server {
	return &Server{}
}

// RegisterRoutes registers all API routes
func RegisterRoutes(e *echo.Echo, s *Server) {
	// User routes
	e.GET("/user", s.GetUser)
}
