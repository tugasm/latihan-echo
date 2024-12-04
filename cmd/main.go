package main

import (
	config "echo-api-007/config/database"
	internal "echo-api-007/internal/middleware"
	handler "echo-api-007/internal/userhandler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	config.InitDB()
	defer config.CloseDB()

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	u := e.Group("/users")
	u.Use(internal.JwtMiddleware)
	u.POST("/", handler.CreateUser)
	u.GET("/:id", handler.GetUserById)

	// /users/
	// /users/1

	e.POST("/login", handler.LoginUser)
	e.POST("/register", handler.Register)

	e.Logger.Fatal(e.Start(":8080"))

}
