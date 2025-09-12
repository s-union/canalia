package api

import (
	"github.com/labstack/echo/v4"
	db "github.com/s-union/canalia/internal/db/generated"
)

type Server struct {
	queries db.Querier
}

func NewServer(queries db.Querier) *Server {
	return &Server{queries: queries}
}

// RegisterRoutes registers all API routes
func RegisterRoutes(e *echo.Echo, s *Server) {
	// User routes
	e.GET("/user", s.GetUser)
	e.POST("/user", s.PostUser)
}
