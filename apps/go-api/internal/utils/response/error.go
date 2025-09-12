package response

import (
	"net/http"

	"github.com/s-union/canalia/internal/generated"
)

// CreateError creates a standardized error response
func CreateError(code int, message string) generated.Error {
	return generated.Error{
		Code:    &code,
		Message: &message,
	}
}

// Common error response helpers
func BadRequest(message string) generated.Error {
	return CreateError(http.StatusBadRequest, message)
}

func Unauthorized(message string) generated.Error {
	return CreateError(http.StatusUnauthorized, message)
}

func NotFound(message string) generated.Error {
	return CreateError(http.StatusNotFound, message)
}

func InternalServerError(message string) generated.Error {
	return CreateError(http.StatusInternalServerError, message)
}

func ValidationError(message string) generated.Error {
	return CreateError(http.StatusBadRequest, "Validation failed: "+message)
}
