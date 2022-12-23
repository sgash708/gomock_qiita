package handler

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type ErrorResponse struct {
	Errors []string `json:"errors"`
}

func (h *Handler) NewErrorResponse(ec echo.Context, err error) error {
	var errres ErrorResponse
	switch err := err.(type) {
	case validator.ValidationErrors:
		errs := err
		for _, err := range errs {
			errres.Errors = append(errres.Errors, h.I18n.EmbedT(err.Tag(), err.Field(), err.Param()))
		}
	default:
		errres.Errors = append(errres.Errors, h.I18n.T(err.Error()))
	}

	return ec.JSON(http.StatusBadRequest, errres)
}
