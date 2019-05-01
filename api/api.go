package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/checkaayush/amigo/pkg/config"
)

// Start starts the API server
func Start(cfg *config.Configuration) error {
	// Echo instance
	e := echo.New()

	// Initialize connection to MongoDB
	// timeout, err := strconv.Atoi(cfg.MongodbTimeout)
	// if err != nil {
	// 	return err
	// }

	// _, err := mongodb.New(
	// 	cfg.MongodbHost,
	// 	cfg.MongodbName,
	// 	cfg.MongodbUsername,
	// 	cfg.MongodbPassword,
	// 	cfg.MongodbTimeout,
	// )
	// if err != nil {
	// 	return err
	// }

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", hello)

	// Start server
	address := ":" + cfg.ServerPort
	if err := e.Start(address); err != nil {
		return err
	}
	return nil
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
