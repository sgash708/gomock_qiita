package main

import (
	"fmt"
	"log"
	"server/api/application"
	"server/api/client/i18n"
	"server/api/handler"
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
	log.Println("...STARTING MY NICE SERVER...")

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

	// Client
	i18nClient, err := i18n.NewI18nClient()
	if err != nil {
		return err
	}

	// Repository

	// Application
	app := application.NewApplication(
		&application.ApplicationBundle{
			ServerConfig: cfg,
		},
	)

	// Handler
	h := handler.NewHandler(app, i18nClient)
	h.AssignRoutes(e)

	return nil
}
