package handler

import (
	"net/http"
	"server/api/handler/request"

	"github.com/labstack/echo/v4"
)

func (h *Handler) GetBook(ec echo.Context) error {
	var req request.GetBookRequest
	if err := ec.Bind(&req); err != nil {
		return h.NewErrorResponse(ec, err)
	}

	ctx := h.GetCtx(ec)
	res, err := h.Application.GetBook(ctx, req.UUID)
	if err != nil {
		return h.NewErrorResponse(ec, err)
	}

	return ec.JSON(http.StatusOK, res)
}

func (h *Handler) GetBooks(ec echo.Context) error {
	ctx := h.GetCtx(ec)
	res, err := h.Application.GetBooks(ctx)
	if err != nil {
		return h.NewErrorResponse(ec, err)
	}

	return ec.JSON(http.StatusOK, res)
}

func (h *Handler) CreateBook(ec echo.Context) error {
	return nil
}

func (h *Handler) UpdateBook(ec echo.Context) error {
	return nil
}

func (h *Handler) DeleteBook(ec echo.Context) error {
	return nil
}
