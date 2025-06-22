package handler

import (
	"bytes"
	"net/http"

	"strings"

	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime/types"

	"github.com/traP-jp/h25s_04/server/internal/schema"
)

// PostImages implements schema.ServerInterface.
func (h *Handler) PostImages(c echo.Context) error {
	file, err := c.FormFile("image")
	if err != nil {
		c.Logger().Errorf("failed to get form file: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, "bad request")
	}

	f, err := file.Open()
	if err != nil {
		c.Logger().Errorf("failed to open form file: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, "bad request")
	}
	defer f.Close()

	b := make([]byte, file.Size)
	if _, err := f.Read(b); err != nil {
		c.Logger().Errorf("failed to read form file: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "internal server error")
	}

	contentType := http.DetectContentType(b)
	if !strings.HasPrefix(contentType, "image/") {
		c.Logger().Errorf("invalid image content type: %s", contentType)
		return echo.NewHTTPError(http.StatusBadRequest, "invalid file format")
	}

	reader := bytes.NewReader(b)
	imageID, err := h.repo.UploadImage(c.Request().Context(), reader, contentType)

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
