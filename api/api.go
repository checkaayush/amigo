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
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", hello)

	// Start server
	address := ":" + cfg.Server.Port
	if err := e.Start(address); err != nil {
		return err
	}
	return nil
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
