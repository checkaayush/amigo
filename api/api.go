package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/checkaayush/amigo/pkg/config"
	"github.com/checkaayush/amigo/pkg/mongodb"
)

// Start starts the API server
func Start(cfg *config.Configuration) error {
	// Echo instance
	e := echo.New()

	// Database
	_, err := mongodb.New(cfg)
	if err != nil {
		return err
	}

	// Middlewares
	initializeMiddlewares(e)

	// Routes
	e.GET("/", hello)

	// Start server
	address := ":" + cfg.Server.Port
	if err := e.Start(address); err != nil {
		return err
	}
	return nil
}

// initializeMiddlewares initializes built-in and custom middlewares
func initializeMiddlewares(e *echo.Echo) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
