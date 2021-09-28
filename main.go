package main

import (
	"dozn/app-server/logger"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {

	logger.LOG.Trace("Starting the application...")
	logger.LOG.Info("Something noteworthy happened...")
	logger.LOG.Warn("There is something you should know about...")
	logger.LOG.Error("Something went wrong...")

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", hello)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
