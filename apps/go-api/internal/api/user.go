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
	"github.com/s-union/canalia/internal/generated"
	"github.com/s-union/canalia/internal/utils/auth0"
	"github.com/s-union/canalia/internal/utils/response"
)

const (
	defaultTimeout = 5 * time.Second
)

var validate = validator.New()

// Helper function to convert DB model to API response
func convertUserToAPIResponse(dbUser *db.Users) generated.User {
	email := openapi_types.Email(dbUser.Email)
	user := generated.User{
		Id:         func() *int { id := int(dbUser.ID); return &id }(),
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
		errResp := response.InternalServerError("Failed to get user info")
		return c.JSON(http.StatusInternalServerError, errResp)
	}

	// Query user by email from database
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	dbUser, err := s.queries.GetUserByEmail(ctx, userInfo.Email)
	if err != nil {
		if err == pgx.ErrNoRows {
			errResp := response.NotFound("User not found")
			return c.JSON(http.StatusNotFound, errResp)
		}
		errResp := response.InternalServerError("Database error")
		return c.JSON(http.StatusInternalServerError, errResp)
	}

	// Convert to API response format
	apiResponse := convertUserToAPIResponse(dbUser)
	return c.JSON(http.StatusOK, apiResponse)
}

func (s *Server) PostUser(c echo.Context) error {
	// Get authenticated user info from Auth0 middleware
	userInfo, ok := c.Get(string(middleware.UserContextKey)).(*auth0.UserInfo)
	if !ok {
		errResp := response.InternalServerError("Failed to get user info")
		return c.JSON(http.StatusInternalServerError, errResp)
	}

	// Parse request body using OpenAPI generated type with validation tags
	var req generated.UserRegistrationInput
	if err := c.Bind(&req); err != nil {
		errResp := response.BadRequest("Invalid request format")
		return c.JSON(http.StatusBadRequest, errResp)
	}

	// Validate request using go-playground/validator
	if err := validate.Struct(req); err != nil {
		errResp := response.ValidationError(err.Error())
		return c.JSON(http.StatusBadRequest, errResp)
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
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	dbUser, err := s.queries.UpsertUserByEmail(ctx, params)
	if err != nil {
		errResp := response.InternalServerError("Failed to save user")
		return c.JSON(http.StatusInternalServerError, errResp)
	}

	// Convert to API response format
	apiResponse := convertUserToAPIResponse(dbUser)
	return c.JSON(http.StatusOK, apiResponse)
}
