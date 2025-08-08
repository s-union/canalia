package api

import (
	"context"
	"database/sql"
	"errors"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"

	"github.com/s-union/canalia/internal/middleware"
	"github.com/s-union/canalia/internal/types"
	"github.com/s-union/canalia/internal/utils"
	"github.com/s-union/canalia/internal/utils/auth0"
)

// CreateUserValidationRequest represents the validation structure for creating a user
type CreateUserValidationRequest struct {
	ContactEmail *string `json:"contactEmail" validate:"omitempty,email"`
	FamilyName   string  `json:"familyName" validate:"required,min=1"`
	GivenName    string  `json:"givenName" validate:"required,min=1"`
	PhoneNumber  *string `json:"phoneNumber"`
}

// UpdateUserValidationRequest represents the validation structure for updating a user
type UpdateUserValidationRequest struct {
	ContactEmail *string `json:"contactEmail" validate:"omitempty,email"`
	FamilyName   string  `json:"familyName" validate:"required,min=1"`
	GivenName    string  `json:"givenName" validate:"required,min=1"`
	PhoneNumber  *string `json:"phoneNumber"`
}

// validator instance
var validate = validator.New()

// validateCreateUserRequest validates the CreateUserRequest using go-playground/validator
func validateCreateUserRequest(req *types.CreateUserRequest) error {
	// Convert to validation struct
	validationReq := CreateUserValidationRequest{
		FamilyName:  req.FamilyName,
		GivenName:   req.GivenName,
		PhoneNumber: req.PhoneNumber,
	}
	
	// Convert email pointer if provided
	if req.ContactEmail != nil {
		emailStr := string(*req.ContactEmail)
		validationReq.ContactEmail = &emailStr
	}
	
	if err := validate.Struct(validationReq); err != nil {
		// Return the first validation error with a user-friendly message
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			for _, fieldError := range validationErrors {
				switch fieldError.Tag() {
				case "required":
					return errors.New(fieldError.Field() + " is required")
				case "min":
					return errors.New(fieldError.Field() + " cannot be empty")
				case "email":
					return errors.New(fieldError.Field() + " must be a valid email address")
				default:
					return errors.New(fieldError.Field() + " is invalid")
				}
			}
		}
		return err
	}
	return nil
}

// validateUpdateUserRequest validates the UpdateUserRequest using go-playground/validator
func validateUpdateUserRequest(req *types.UpdateUserRequest) error {
	// Convert to validation struct
	validationReq := UpdateUserValidationRequest{
		FamilyName:  req.FamilyName,
		GivenName:   req.GivenName,
		PhoneNumber: req.PhoneNumber,
	}
	
	// Convert email pointer if provided
	if req.ContactEmail != nil {
		emailStr := string(*req.ContactEmail)
		validationReq.ContactEmail = &emailStr
	}
	
	if err := validate.Struct(validationReq); err != nil {
		// Return the first validation error with a user-friendly message
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			for _, fieldError := range validationErrors {
				switch fieldError.Tag() {
				case "required":
					return errors.New(fieldError.Field() + " is required")
				case "min":
					return errors.New(fieldError.Field() + " cannot be empty")
				case "email":
					return errors.New(fieldError.Field() + " must be a valid email address")
				default:
					return errors.New(fieldError.Field() + " is invalid")
				}
			}
		}
		return err
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
