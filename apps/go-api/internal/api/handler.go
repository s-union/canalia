package api

import (
	"database/sql"

	"github.com/labstack/echo/v4"
	"github.com/s-union/canalia/internal/db/generated"
)

type Server struct {
	db      *sql.DB
	queries *db.Queries
}

func NewServer(database *sql.DB) *Server {
	return &Server{
		db:      database,
		queries: db.New(database),
	}
}

// RegisterRoutes registers all API routes
func RegisterRoutes(e *echo.Echo, s *Server) {
	// User routes
	e.GET("/user", s.GetUser)
	e.POST("/user", s.CreateUser)
	e.PUT("/user", s.UpdateUser)
}
