package middleware

import (
	"context"
	"log"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/auth0/go-jwt-middleware/v2/jwks"
	"github.com/auth0/go-jwt-middleware/v2/validator"

	"github.com/labstack/echo/v4"

	"github.com/s-union/canalia/internal/utils/auth0"
)

// ContextKey はコンテキストに値を設定/取得するためのキーの型です。
// 文字列リテラルよりも型安全性を高めるために使用します。
type ContextKey string

const (
	ClaimsContextKey      ContextKey = "claims"
	AccessTokenContextKey ContextKey = "user_access_token"
	UserContextKey        ContextKey = "user"
)

type CustomClaims struct {
	Scope string `json:"scope"`
}

func (c *CustomClaims) Validate(ctx context.Context) error {
	return nil
}

func JWTAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		issuerURL, err := url.Parse("https://" + os.Getenv("AUTH0_DOMAIN") + "/")
		if err != nil {
			log.Fatalf("Error parsing issuer URL: %v", err)
		}

		provider := jwks.NewCachingProvider(issuerURL, 5*time.Minute)

		jwtValidator, err := validator.New(
			provider.KeyFunc,
			validator.RS256,
			issuerURL.String(),
			[]string{os.Getenv("AUTH0_AUDIENCE")},
			validator.WithCustomClaims(
				func() validator.CustomClaims {
					return &CustomClaims{}
				},
			),
			validator.WithAllowedClockSkew(time.Minute),
		)
		if err != nil {
			log.Fatalf("Error creating JWT validator: %v", err)
		}

		authHeader := c.Request().Header.Get("Authorization")
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			log.Println("Authorization header format must be Bearer {token}")
			return echo.ErrUnauthorized
		}
		token := parts[1]

		validToken, err := jwtValidator.ValidateToken(c.Request().Context(), token)
		if err != nil {
			log.Printf("Error validating token: %v", err)
			return echo.ErrUnauthorized
		}
		claims := validToken.(*validator.ValidatedClaims)

		c.Set(string(AccessTokenContextKey), token)
		c.Set(string(ClaimsContextKey), claims)

		userContext, err := auth0.FetchUserInfo(token)
		if err != nil {
			log.Printf("Error fetching user info: %v", err)
		}
		c.Set(string(UserContextKey), userContext)

		return next(c)
	}
}
