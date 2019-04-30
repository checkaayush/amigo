package api

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/checkaayush/amigo/pkg/config"
)

// Start starts the API server
func Start(cfg *config.Configuration) error {
	// Echo instance
	e := echo.New()

	// Middleware
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
