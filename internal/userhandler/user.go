package internal

import (
	"context"
	config "echo-api-007/config/database"
	internal "echo-api-007/internal/userdto"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func Register(c echo.Context) error {
	var req internal.RegisterUser
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid Request"})
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Invalid generate password"})
	}

	query := "INSERT INTO users(name, email, password) VALUES ($1,$2, $3) RETURNING id"
	var userID int
	err = config.Pool.QueryRow(context.Background(), query, req.Name, req.Email, string(hashPassword)).Scan(&userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Register Failed"})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "User Registere!",
	})
}
