package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) GetPing(c echo.Context) error {
	return c.NoContent(http.StatusNoContent)
}
