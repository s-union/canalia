package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	mockapi "github.com/s-union/canalia/internal/api/generated"
	db "github.com/s-union/canalia/internal/db/generated"
	"github.com/s-union/canalia/internal/middleware"
	"github.com/s-union/canalia/internal/generated"
	"github.com/s-union/canalia/internal/utils/auth0"
)

// Helper functions for test setup and common data

func createTestUserInfo() *auth0.UserInfo {
	return &auth0.UserInfo{
		Sub:           "auth0|123",
		Name:          "Test User",
		Email:         "test@example.com",
		EmailVerified: true,
	}
}

func createTestUser() *db.Users {
	return &db.Users{
		ID:         1,
		Email:      "test@example.com",
		FamilyName: "Test",
		GivenName:  "User",
		IsVerified: true,
		IsActive:   true,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
}

func setupTest(t *testing.T) (*mockapi.MockQuerier, *Server, *echo.Echo) {
	ctrl := gomock.NewController(t)
	mockQuerier := mockapi.NewMockQuerier(ctrl)
	server := NewServer(mockQuerier)
	e := echo.New()
	RegisterRoutes(e, server)

	t.Cleanup(ctrl.Finish)

	return mockQuerier, server, e
}

func assertUserResponse(t *testing.T, rec *httptest.ResponseRecorder, expectedID int, expectedEmail string) {
	assert.Equal(t, http.StatusOK, rec.Code)

	var response generated.User
	err := json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, expectedID, *response.Id)
	assert.Equal(t, expectedEmail, string(*response.Email))
}

func TestGetUser(t *testing.T) {
	mockQuerier, server, e := setupTest(t)

	// Mock user info
	userInfo := createTestUserInfo()

	// Mock DB response
	expectedUser := createTestUser()

	mockQuerier.EXPECT().GetUserByEmail(gomock.Any(), "test@example.com").Return(expectedUser, nil)

	req := httptest.NewRequest(http.MethodGet, "/user", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Set user info in context (simulating middleware)
	c.Set(string(middleware.UserContextKey), userInfo)

	err := server.GetUser(c)

	assert.NoError(t, err)
	assertUserResponse(t, rec, 1, "test@example.com")
}

func TestPostUser(t *testing.T) {
	mockQuerier, server, e := setupTest(t)

	// Mock user info
	userInfo := createTestUserInfo()

	// Mock request body
	reqBody := generated.UserRegistrationInput{
		FamilyName: "Test",
		GivenName:  "User",
	}

	// Mock DB response
	expectedUser := createTestUser()

	mockQuerier.EXPECT().UpsertUserByEmail(gomock.Any(), gomock.Any()).Return(expectedUser, nil)

	reqBodyBytes, _ := json.Marshal(reqBody)
	req := httptest.NewRequest(http.MethodPost, "/user", bytes.NewReader(reqBodyBytes))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Set user info in context (simulating middleware)
	c.Set(string(middleware.UserContextKey), userInfo)

	err := server.PostUser(c)

	assert.NoError(t, err)
	assertUserResponse(t, rec, 1, "test@example.com")
}
