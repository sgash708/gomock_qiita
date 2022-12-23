package handler

import (
	"context"
	"server/api/application"
	"server/api/client/i18n"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	Application application.ApplicationInterface
	I18n        i18n.I18nClientInterface
}

func NewHandler(
	app application.ApplicationInterface,
	i18n i18n.I18nClientInterface,
) *Handler {
	return &Handler{
		Application: app,
		I18n:        i18n,
	}
}

func (h *Handler) AssignRoutes(e *echo.Echo) {
	v1g := e.Group("v1")
	{
		v1bg := v1g.Group("/books")
		{
			v1bg.GET("/:uuid", h.GetBook)
			v1bg.GET("", h.GetBooks)
			v1bg.POST("", h.CreateBook)
			v1bg.PUT("/:uuid", h.UpdateBook)
			v1bg.DELETE("/:uuid", h.DeleteBook)
		}
	}
}

func (h *Handler) GetCtx(ec echo.Context) context.Context {
	return ec.Request().Context()
}
