package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime/types"

	"github.com/traP-jp/h25s_04/server/internal/repository"
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
	var review schema.ReviewDetail
	if err := c.Bind(&review); err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("Invalid request body: %v", err))
	}

	createParams := repository.Review{
		Id:       reviewId,
		EateryID: review.EateryId,
		UserID:   getUserID(params.XForwardedUser),
		Content:  review.Content,
	}

	if _, err := h.repo.ReviewExists(c.Request().Context(), createParams); err != nil {
		return c.JSON(http.StatusNotFound, fmt.Sprintf("Review with ID %s does not exist", reviewId))

	}
	if err := h.repo.DeleteReview(c.Request().Context(), reviewId); err != nil {
		return c.JSON(http.StatusInternalServerError, schema.Error{
			Code:  "INTERNAL_SERVER_ERROR",
			Error: "Failed to delete review",
		})
	}

	return c.JSON(http.StatusOK, "Review deleted successfully")

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
