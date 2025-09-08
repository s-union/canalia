package api

import (
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	openapi_types "github.com/oapi-codegen/runtime/types"

	"github.com/s-union/canalia/internal/types"
)

func TestUserRegistrationValidation(t *testing.T) {
	validator := validator.New()

	tests := []struct {
		name      string
		req       types.UserRegistrationInput
		wantError bool
	}{
		{
			name: "Valid request with all fields",
			req: types.UserRegistrationInput{
				FamilyName:   "田中",
				GivenName:    "太郎",
				ContactEmail: emailPtr("tanaka@example.com"),
				PhoneNumber:  stringPtr("090-1234-5678"),
			},
			wantError: false,
		},
		{
			name: "Valid request with required fields only",
			req: types.UserRegistrationInput{
				FamilyName: "田中",
				GivenName:  "太郎",
			},
			wantError: false,
		},
		{
			name: "Missing family name",
			req: types.UserRegistrationInput{
				GivenName: "太郎",
			},
			wantError: true,
		},
		{
			name: "Missing given name",
			req: types.UserRegistrationInput{
				FamilyName: "田中",
			},
			wantError: true,
		},
		{
			name: "Empty family name",
			req: types.UserRegistrationInput{
				FamilyName: "",
				GivenName:  "太郎",
			},
			wantError: true,
		},
		{
			name: "Empty given name",
			req: types.UserRegistrationInput{
				FamilyName: "田中",
				GivenName:  "",
			},
			wantError: true,
		},
		{
			name: "Family name too long",
			req: types.UserRegistrationInput{
				FamilyName: stringRepeat("あ", 101),
				GivenName:  "太郎",
			},
			wantError: true,
		},
		{
			name: "Given name too long",
			req: types.UserRegistrationInput{
				FamilyName: "田中",
				GivenName:  stringRepeat("あ", 101),
			},
			wantError: true,
		},
		{
			name: "Invalid email format",
			req: types.UserRegistrationInput{
				FamilyName:   "田中",
				GivenName:    "太郎",
				ContactEmail: emailPtr("invalid-email"),
			},
			wantError: true,
		},
		{
			name: "Phone number too long",
			req: types.UserRegistrationInput{
				FamilyName:  "田中",
				GivenName:   "太郎",
				PhoneNumber: stringPtr(stringRepeat("0", 21)),
			},
			wantError: true,
		},
		{
			name: "Valid email with special characters",
			req: types.UserRegistrationInput{
				FamilyName:   "田中",
				GivenName:    "太郎",
				ContactEmail: emailPtr("test+label@sub.domain.com"),
			},
			wantError: false,
		},
		{
			name: "Valid phone number with various formats",
			req: types.UserRegistrationInput{
				FamilyName:  "田中",
				GivenName:   "太郎",
				PhoneNumber: stringPtr("03-1234-5678"),
			},
			wantError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validator.Struct(tt.req)
			if tt.wantError {
				assert.Error(t, err, "Expected validation error but got none")
			} else {
				assert.NoError(t, err, "Expected no validation error but got: %v", err)
			}
		})
	}
}

// Helper function to create email pointer
func emailPtr(s string) *openapi_types.Email {
	email := openapi_types.Email(s)
	return &email
}

// Helper function to create string pointer  
func stringPtr(s string) *string {
	return &s
}

// Helper function to repeat string n times
func stringRepeat(s string, n int) string {
	result := ""
	for i := 0; i < n; i++ {
		result += s
	}
	return result
}