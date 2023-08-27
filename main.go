package main

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"

	"ish/src"
	"ish/src/oai"
)

func main() {
	oai.Init()

	// Create an Echo instance
	e := echo.New()

	// Set the logger to use a custom format
	e.Logger.SetLevel(log.INFO)
	e.Logger.SetOutput(os.Stdout)
	e.Logger.SetHeader("${time_rfc3339} ${level}")

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.Static("static"))

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	// Routes
	// completions
	e.POST("/rewrite", src.Rewrite)

	// Start the server
	e.Logger.Fatal(e.Start(":8080"))
}
