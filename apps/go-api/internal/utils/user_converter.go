package utils

import (
	"database/sql"

	"github.com/oapi-codegen/runtime/types"
	"github.com/s-union/canalia/internal/db/generated"
	apiTypes "github.com/s-union/canalia/internal/types"
)

// ConvertDBUserToAPIUser converts a database User to an API User
func ConvertDBUserToAPIUser(dbUser *db.Users) *apiTypes.User {
	user := &apiTypes.User{
		Id:         int(dbUser.ID),
		Email:      types.Email(dbUser.Email),
		IsVerified: dbUser.IsVerified,
		FamilyName: dbUser.FamilyName,
		GivenName:  dbUser.GivenName,
		IsActive:   dbUser.IsActive,
		CreatedAt:  dbUser.CreatedAt,
		UpdatedAt:  dbUser.UpdatedAt,
	}

	// Handle nullable fields
	if dbUser.ContactEmail.Valid {
		contactEmail := types.Email(dbUser.ContactEmail.String)
		user.ContactEmail = &contactEmail
	}

	if dbUser.PhoneNumber.Valid {
		user.PhoneNumber = &dbUser.PhoneNumber.String
	}

	return user
}

// ConvertCreateUserRequestToDBParams converts a CreateUserRequest to database parameters
func ConvertCreateUserRequestToDBParams(req *apiTypes.CreateUserRequest, email string) *db.CreateUserParams {
	params := &db.CreateUserParams{
		Email:      email,
		FamilyName: req.FamilyName,
		GivenName:  req.GivenName,
	}

	// Handle nullable fields
	if req.ContactEmail != nil {
		params.ContactEmail = sql.NullString{
			String: string(*req.ContactEmail),
			Valid:  true,
		}
	}

	if req.PhoneNumber != nil {
		params.PhoneNumber = sql.NullString{
			String: *req.PhoneNumber,
			Valid:  true,
		}
	}

	return params
}

// ConvertUpdateUserRequestToDBParams converts an UpdateUserRequest to database parameters
func ConvertUpdateUserRequestToDBParams(req *apiTypes.UpdateUserRequest, email string) *db.UpdateUserParams {
	params := &db.UpdateUserParams{
		Email:      email,
		FamilyName: req.FamilyName,
		GivenName:  req.GivenName,
	}

	// Handle nullable fields
	if req.ContactEmail != nil {
		params.ContactEmail = sql.NullString{
			String: string(*req.ContactEmail),
			Valid:  true,
		}
	}

	if req.PhoneNumber != nil {
		params.PhoneNumber = sql.NullString{
			String: *req.PhoneNumber,
			Valid:  true,
		}
	}

	return params
}