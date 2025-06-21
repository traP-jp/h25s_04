package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime/types"

	"github.com/traP-jp/h25s_04/server/internal/schema"
)

// PostImages implements schema.ServerInterface.
func (h *Handler) PostImages(ctx echo.Context) error {
	return ctx.JSON(http.StatusNotImplemented, schema.Error{
		Code:  "NOT_IMPLEMENTED",
		Error: "PostImages endpoint is not implemented yet",
	})
}

// GetImagesImageId implements schema.ServerInterface.
func (h *Handler) GetImagesImageId(ctx echo.Context, imageId types.UUID) error {
	return ctx.JSON(http.StatusNotImplemented, schema.Error{
		Code:  "NOT_IMPLEMENTED",
		Error: "GetImagesImageId endpoint is not implemented yet",
	})
}
