package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/s-union/canalia/generated"
)

var _ generated.ServerInterface = (*Server)(nil)

type Server struct{}

func NewServer() Server {
	return Server{}
}

func (s *Server) GetUser(ctx echo.Context) error {
	email := "sample@example.com"
	name := "hogehoge"
	return ctx.JSON(http.StatusOK, generated.User{
		Email: &email,
		Name:  &name,
	})
}
