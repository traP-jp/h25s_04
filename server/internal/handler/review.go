package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime/types"

	"github.com/traP-jp/h25s_04/server/internal/schema"
)

// GetReviews implements schema.ServerInterface.
func (h *Handler) GetReviews(c echo.Context, params schema.GetReviewsParams) error {
	return c.JSON(http.StatusNotImplemented, schema.Error{
		Code:  "NOT_IMPLEMENTED",
		Error: "GetReviews endpoint is not implemented yet",
	})
}

// DeleteReviewsReviewId implements schema.ServerInterface.
func (h *Handler) DeleteReviewsReviewId(c echo.Context, reviewId types.UUID, params schema.DeleteReviewsReviewIdParams) error {
	return c.JSON(http.StatusNotImplemented, schema.Error{
		Code:  "NOT_IMPLEMENTED",
		Error: "DeleteReviewsReviewId endpoint is not implemented yet",
	})
}

// GetReviewsReviewId implements schema.ServerInterface.
func (h *Handler) GetReviewsReviewId(c echo.Context, reviewId types.UUID) error {
	return c.JSON(http.StatusNotImplemented, schema.Error{
		Code:  "NOT_IMPLEMENTED",
		Error: "GetReviewsReviewId endpoint is not implemented yet",
	})
}

// PutReviewsReviewId implements schema.ServerInterface.
func (h *Handler) PutReviewsReviewId(c echo.Context, reviewId types.UUID, params schema.PutReviewsReviewIdParams) error {
	return c.JSON(http.StatusNotImplemented, schema.Error{
		Code:  "NOT_IMPLEMENTED",
		Error: "PutReviewsReviewId endpoint is not implemented yet",
	})
}
