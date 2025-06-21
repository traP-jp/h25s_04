package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime/types"

	"github.com/traP-jp/h25s_04/server/internal/schema"
)

// GetReviews implements schema.ServerInterface.
func (h *Handler) GetReviews(ctx echo.Context, params schema.GetReviewsParams) error {
	return ctx.JSON(http.StatusNotImplemented, schema.Error{
		Code:  "NOT_IMPLEMENTED",
		Error: "GetReviews endpoint is not implemented yet",
	})
}

// DeleteReviewsReviewId implements schema.ServerInterface.
func (h *Handler) DeleteReviewsReviewId(ctx echo.Context, reviewId types.UUID) error {
	return ctx.JSON(http.StatusNotImplemented, schema.Error{
		Code:  "NOT_IMPLEMENTED",
		Error: "DeleteReviewsReviewId endpoint is not implemented yet",
	})
}

// GetReviewsReviewId implements schema.ServerInterface.
func (h *Handler) GetReviewsReviewId(ctx echo.Context, reviewId types.UUID) error {
	return ctx.JSON(http.StatusNotImplemented, schema.Error{
		Code:  "NOT_IMPLEMENTED",
		Error: "GetReviewsReviewId endpoint is not implemented yet",
	})
}

// PutReviewsReviewId implements schema.ServerInterface.
func (h *Handler) PutReviewsReviewId(ctx echo.Context, reviewId types.UUID) error {
	return ctx.JSON(http.StatusNotImplemented, schema.Error{
		Code:  "NOT_IMPLEMENTED",
		Error: "PutReviewsReviewId endpoint is not implemented yet",
	})
}
