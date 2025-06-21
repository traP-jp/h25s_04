package handler

import (
	"bytes"
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

	b, err := req.Image.Bytes()
	if err != nil {
		c.Logger().Errorf("failed to read image bytes: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, "bad request")
	}

	UploadImage := http.DetectContentType(b)
	if UploadImage != "image/" {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid file format")
	}


	reader := bytes.NewReader(b)
	imageID, err := h.repo.UploadImage(c.Request().Context(), reader, UploadImage)

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
	image, contentType, err := h.repo.GetImage(c.Request().Context(), imageId)
	if err != nil {
		c.Logger().Errorf("failed to download image: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError).SetInternal(err)
	}

	defer image.Close()

	return c.Stream(http.StatusOK, contentType, image)
}
