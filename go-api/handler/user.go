package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/s-union/canalia/generated"
	"github.com/s-union/canalia/middleware"
	"github.com/s-union/canalia/utils/auth0"
)

func (s *Server) GetUser(c echo.Context) error {
	user, ok := c.Get(string(middleware.UserContextKey)).(*auth0.UserInfo)
	if !ok {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    http.StatusInternalServerError,
			"message": "Failed to get user info",
		})
	}
	email := user.Email
	name := "HogeHoge"

	return c.JSON(http.StatusOK, generated.User{
		Email: &email,
		Name:  &name,
	})
}
