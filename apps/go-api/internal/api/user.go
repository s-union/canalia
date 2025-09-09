package api

import (
	"context"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo/v4"
	openapi_types "github.com/oapi-codegen/runtime/types"

	db "github.com/s-union/canalia/internal/db/generated"
	"github.com/s-union/canalia/internal/middleware"
	"github.com/s-union/canalia/internal/types"
	"github.com/s-union/canalia/internal/utils/auth0"
	"github.com/s-union/canalia/internal/utils/response"
)

var validate = validator.New()

// Helper function to convert DB model to API response
func convertUserToAPIResponse(dbUser *db.Users) types.User {
	email := openapi_types.Email(dbUser.Email)
	user := types.User{
		Id:         func(i int32) *int { v := int(i); return &v }(dbUser.ID),
		Email:      &email,
		FamilyName: &dbUser.FamilyName,
		GivenName:  &dbUser.GivenName,
		IsVerified: &dbUser.IsVerified,
		IsActive:   &dbUser.IsActive,
		CreatedAt:  &dbUser.CreatedAt,
		UpdatedAt:  &dbUser.UpdatedAt,
	}

	// Handle nullable fields
	if dbUser.ContactEmail != nil {
		contactEmail := openapi_types.Email(*dbUser.ContactEmail)
		user.ContactEmail = &contactEmail
	}
	if dbUser.PhoneNumber != nil {
		user.PhoneNumber = dbUser.PhoneNumber
	}

	return user
}

func (s *Server) GetUser(c echo.Context) error {
	// Get authenticated user info from Auth0 middleware
	userInfo, ok := c.Get(string(middleware.UserContextKey)).(*auth0.UserInfo)
	if !ok {
		err := response.InternalServerError("Failed to get user info")
		return c.JSON(http.StatusInternalServerError, err)
	}

	// Query user by email from database
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	dbUser, err := s.queries.GetUserByEmail(ctx, userInfo.Email)
	if err != nil {
		if err == pgx.ErrNoRows {
			err := response.NotFound("User not found")
			return c.JSON(http.StatusNotFound, err)
		}
		err := response.InternalServerError("Database error")
		return c.JSON(http.StatusInternalServerError, err)
	}

	// Convert to API response format
	apiResponse := convertUserToAPIResponse(dbUser)
	return c.JSON(http.StatusOK, apiResponse)
}

func (s *Server) PostUser(c echo.Context) error {
	// Get authenticated user info from Auth0 middleware
	userInfo, ok := c.Get(string(middleware.UserContextKey)).(*auth0.UserInfo)
	if !ok {
		err := response.InternalServerError("Failed to get user info")
		return c.JSON(http.StatusInternalServerError, err)
	}

	// Parse request body using OpenAPI generated type with validation tags
	var req types.UserRegistrationInput
	if err := c.Bind(&req); err != nil {
		err := response.BadRequest("Invalid request format")
		return c.JSON(http.StatusBadRequest, err)
	}

	// Validate request using go-playground/validator
	if err := validate.Struct(req); err != nil {
		err := response.ValidationError(err.Error())
		return c.JSON(http.StatusBadRequest, err)
	}

	// Prepare database parameters
	var contactEmail *string
	if req.ContactEmail != nil {
		emailStr := string(*req.ContactEmail)
		contactEmail = &emailStr
	}
	params := &db.UpsertUserByEmailParams{
		Email:        userInfo.Email,
		FamilyName:   req.FamilyName,
		GivenName:    req.GivenName,
		ContactEmail: contactEmail,
		PhoneNumber:  req.PhoneNumber,
	}

	// Execute upsert
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	dbUser, err := s.queries.UpsertUserByEmail(ctx, params)
	if err != nil {
		err := response.InternalServerError("Failed to save user")
		return c.JSON(http.StatusInternalServerError, err)
	}

	// Convert to API response format
	apiResponse := convertUserToAPIResponse(dbUser)
	return c.JSON(http.StatusOK, apiResponse)
}
