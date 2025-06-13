package api

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/s-union/canalia/internal/middleware"
	"github.com/s-union/canalia/internal/types"
	"github.com/s-union/canalia/internal/utils/auth0"
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

	return c.JSON(http.StatusOK, types.User{
		Email: &email,
		Name:  &name,
	})
}
