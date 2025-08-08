package api

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"

	"github.com/s-union/canalia/internal/middleware"
	"github.com/s-union/canalia/internal/types"
	"github.com/s-union/canalia/internal/utils"
	"github.com/s-union/canalia/internal/utils/auth0"
)

// validateCreateUserRequest validates the CreateUserRequest
func validateCreateUserRequest(req *types.CreateUserRequest) error {
	if strings.TrimSpace(req.FamilyName) == "" {
		return errors.New("familyName is required")
	}
	if strings.TrimSpace(req.GivenName) == "" {
		return errors.New("givenName is required")
	}
	return nil
}

// validateUpdateUserRequest validates the UpdateUserRequest
func validateUpdateUserRequest(req *types.UpdateUserRequest) error {
	if strings.TrimSpace(req.FamilyName) == "" {
		return errors.New("familyName is required")
	}
	if strings.TrimSpace(req.GivenName) == "" {
		return errors.New("givenName is required")
	}
	return nil
}

func (s *Server) GetUser(c echo.Context) error {
	user, ok := c.Get(string(middleware.UserContextKey)).(*auth0.UserInfo)
	if !ok {
		return c.JSON(http.StatusInternalServerError, types.Error{
			Code:    http.StatusInternalServerError,
			Message: "Failed to get user info",
		})
	}

	ctx := context.Background()
	dbUser, err := s.queries.GetUserByEmail(ctx, user.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return c.JSON(http.StatusNotFound, types.Error{
				Code:    http.StatusNotFound,
				Message: "User not found",
			})
		}
		return c.JSON(http.StatusInternalServerError, types.Error{
			Code:    http.StatusInternalServerError,
			Message: "Failed to retrieve user",
		})
	}

	apiUser := utils.ConvertDBUserToAPIUser(dbUser)
	return c.JSON(http.StatusOK, apiUser)
}

func (s *Server) CreateUser(c echo.Context) error {
	user, ok := c.Get(string(middleware.UserContextKey)).(*auth0.UserInfo)
	if !ok {
		return c.JSON(http.StatusInternalServerError, types.Error{
			Code:    http.StatusInternalServerError,
			Message: "Failed to get user info",
		})
	}

	var req types.CreateUserRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, types.Error{
			Code:    http.StatusBadRequest,
			Message: "Invalid request body",
		})
	}

	// Validate request
	if err := validateCreateUserRequest(&req); err != nil {
		return c.JSON(http.StatusBadRequest, types.Error{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	ctx := context.Background()

	// Check if user already exists
	exists, err := s.queries.CheckUserExists(ctx, user.Email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, types.Error{
			Code:    http.StatusInternalServerError,
			Message: "Failed to check user existence",
		})
	}
	if exists {
		return c.JSON(http.StatusConflict, types.Error{
			Code:    http.StatusConflict,
			Message: "User already exists",
		})
	}

	// Create user
	params := utils.ConvertCreateUserRequestToDBParams(&req, user.Email)
	dbUser, err := s.queries.CreateUser(ctx, params)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, types.Error{
			Code:    http.StatusInternalServerError,
			Message: "Failed to create user",
		})
	}

	apiUser := utils.ConvertDBUserToAPIUser(dbUser)
	return c.JSON(http.StatusCreated, apiUser)
}

func (s *Server) UpdateUser(c echo.Context) error {
	user, ok := c.Get(string(middleware.UserContextKey)).(*auth0.UserInfo)
	if !ok {
		return c.JSON(http.StatusInternalServerError, types.Error{
			Code:    http.StatusInternalServerError,
			Message: "Failed to get user info",
		})
	}

	var req types.UpdateUserRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, types.Error{
			Code:    http.StatusBadRequest,
			Message: "Invalid request body",
		})
	}

	// Validate request
	if err := validateUpdateUserRequest(&req); err != nil {
		return c.JSON(http.StatusBadRequest, types.Error{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	ctx := context.Background()
	params := utils.ConvertUpdateUserRequestToDBParams(&req, user.Email)
	dbUser, err := s.queries.UpdateUser(ctx, params)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return c.JSON(http.StatusNotFound, types.Error{
				Code:    http.StatusNotFound,
				Message: "User not found",
			})
		}
		return c.JSON(http.StatusInternalServerError, types.Error{
			Code:    http.StatusInternalServerError,
			Message: "Failed to update user",
		})
	}

	apiUser := utils.ConvertDBUserToAPIUser(dbUser)
	return c.JSON(http.StatusOK, apiUser)
}
