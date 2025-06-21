package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime/types"

	"github.com/traP-jp/h25s_04/server/internal/schema"
)

// PostImages implements schema.ServerInterface.
func (h *Handler) PostImages(c echo.Context) error {
	var req schema.PostImagesMultipartRequestBody
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "bad request")
	}

	reader, err := req.Image.Reader()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "failed to read image")
	}
	defer reader.Close()

	imageID, err := h.repo.UploadImage(c.Request().Context(), reader)

	if err != nil {
		c.Logger().Errorf("failed to upload image: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError).SetInternal(err)
	}

	resp := schema.ImageUploadResponse{
		Id: types.UUID(imageID),
	}

	return c.JSON(http.StatusCreated, resp)
}

// GetImagesImageId implements schema.ServerInterface.
func (h *Handler) GetImagesImageId(c echo.Context, imageId types.UUID) error {
	return c.JSON(http.StatusNotImplemented, schema.Error{
		Code:  "NOT_IMPLEMENTED",
		Error: "GetImagesImageId endpoint is not implemented yet",
	})
}
