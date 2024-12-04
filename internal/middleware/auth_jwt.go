package middleware

import (
	"net/http"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var JwtMiddleware = middleware.JWTWithConfig(middleware.JWTConfig{
	SigningKey:  []byte("12345"),
	Claims:      &jwt.RegisteredClaims{},
	TokenLookup: "header:Authorization",
	AuthScheme:  "Bearer",
	ErrorHandler: func(err error) error {
		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid Authentication")
	},
})
