package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
	"ish/src"
	"ish/src/openai"
	"os"
)

func main() {
	// Initialize viper and read configurations
	viper.SetConfigFile("config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("failed to read config file: %s", err))
	}

	openai.Init()

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
	port := viper.GetString("port")
	if port == "" {
		port = ":8080"
	}

	e.Logger.Fatal(e.Start(port))
}
