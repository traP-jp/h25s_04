package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) PostPing(c echo.Context) error {
	return c.NoContent(http.StatusNoContent)
}
