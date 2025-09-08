package response

import (
	"net/http"

	"github.com/s-union/canalia/internal/types"
)

// CreateError creates a standardized error response
func CreateError(code int, message string) types.Error {
	return types.Error{
		Code:    &code,
		Message: &message,
	}
}

// Common error response helpers
func BadRequest(message string) types.Error {
	return CreateError(http.StatusBadRequest, message)
}

func Unauthorized(message string) types.Error {
	return CreateError(http.StatusUnauthorized, message)
}

func NotFound(message string) types.Error {
	return CreateError(http.StatusNotFound, message)
}

func InternalServerError(message string) types.Error {
	return CreateError(http.StatusInternalServerError, message)
}

func ValidationError(message string) types.Error {
	return CreateError(http.StatusBadRequest, "Validation failed: "+message)
}