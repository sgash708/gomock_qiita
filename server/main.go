package main

import (
	"fmt"
	"log"
	"server/config"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	err error
	cfg *config.ServerConfig
)

func main() {
	if cfg, err = config.LoadEnvConfig(); err != nil {
		panic(err)
	}
	log.Println("STARTING MY NICE SERVER...")

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())

	if err = assignRoutes(e, cfg); err != nil {
		panic(err)
	}

	echoPort := fmt.Sprintf(":%s", cfg.Port)
	if err = e.Start(echoPort); err != nil {
		panic(err)
	}
}

func assignRoutes(e *echo.Echo, cfg *config.ServerConfig) error {
	/* DI */

	// Repository

	// Application

	// Handler

	return nil
}
