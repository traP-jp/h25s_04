package handler

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime/types"

	"github.com/traP-jp/h25s_04/server/internal/repository"
	"github.com/traP-jp/h25s_04/server/internal/schema"
)

// GetEateriesEateryIdReviews implements schema.ServerInterface.
func (h *Handler) GetEateriesEateryIdReviews(c echo.Context, eateryId types.UUID, params schema.GetEateriesEateryIdReviewsParams) error {
	limit := 10 // Default limit
	if params.Limit != nil {
		limit = *params.Limit
	}
	pages := 1 // Default page
	if params.Page != nil {
		pages = *params.Page
	}
	offset := (pages - 1) * limit

	reviews, err := h.repo.GetEateryEateryIDReviews(c.Request().Context(), uuid.UUID(eateryId), limit, offset)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError).SetInternal(err)
	}

	resData := make([]schema.ReviewDetail, len(reviews))
	for i, review := range reviews {
		resData[i] = schema.ReviewDetail{
			Id:        review.Id,
			EateryId:  review.EateryID,
			AuthorId:  review.UserID,
			Content:   review.Content,
			CreatedAt: review.CreatedAt,
			UpdatedAt: review.UpdatedAt,
		}
	}

	res := schema.ReviewDetailListResponse{
		Data: resData,
		Pagination: schema.Pagination{
			Limit: limit,
			Page:  pages,
		},
	}

	return c.JSON(http.StatusOK, res)
}

// PostEateriesEateryIdReviews implements schema.ServerInterface.
func (h *Handler) PostEateriesEateryIdReviews(c echo.Context, eateryId types.UUID, params schema.PostEateriesEateryIdReviewsParams) error {
	var req schema.ReviewDetail
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body")
	}

	createParams := repository.CreateEateryReviewParams{
		ID:       uuid.New(),
		EateryID: uuid.UUID(eateryId),
		Content:  req.Content,
		UserID:   getUserID(params.XForwardedUser),
	}

	if _, err := h.repo.EateryExists(c.Request().Context(), createParams); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, fmt.Sprintf("Eatery with ID %s not found", eateryId))
	}

	reviewID, err := h.repo.PostEateryReview(c.Request().Context(), createParams)
	userID := getUserID(params.XForwardedUser)
	//reviewIDには、リポジトリから返された新しいレビューのIDが入る
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "fail to post eatery review")
	}

	res := schema.ReviewDetail{
		Id:       reviewID,
		Content:  req.Content,
		EateryId: eateryId,
		AuthorId: userID,
	}

	return c.JSON(http.StatusOK, res)
}

// GetReviews implements schema.ServerInterface.
func (h *Handler) GetReviews(c echo.Context, params schema.GetReviewsParams) error {
	return c.JSON(http.StatusNotImplemented, schema.Error{
		Code:  "NOT_IMPLEMENTED",
		Error: "GetReviews endpoint is not implemented yet",
	})
}

// DeleteReviewsReviewId implements schema.ServerInterface.
func (h *Handler) DeleteReviewsReviewId(c echo.Context, reviewId types.UUID, params schema.DeleteReviewsReviewIdParams) error {

	if _, err := h.repo.ReviewExists(c.Request().Context(), reviewId); err != nil {
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
