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
		ImageIDs, err := h.repo.GetImageIDsByReviewID(c.Request().Context(), review.Id)
		if err != nil {
			echo.NewHTTPError(http.StatusInternalServerError).SetInternal(err)
		}

		resData[i] = schema.ReviewDetail{
			Id:        review.Id,
			EateryId:  review.EateryID,
			AuthorId:  review.UserID,
			Content:   review.Content,
			CreatedAt: review.CreatedAt,
			UpdatedAt: review.UpdatedAt,
			ImageIds:  ImageIDs,
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

	// 紐づけた画像をDBに追加する
	if err := h.repo.InsertImageToReview(c.Request().Context(), reviewID, req.ImageIds); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to insert image info to DB").SetInternal(err)
	}

	go func() {
		// traQにメッセージを投稿する処理を書く
	}()

	res := schema.ReviewDetail{
		Id:       reviewID,
		Content:  req.Content,
		EateryId: eateryId,
		AuthorId: userID,
		ImageIds: req.ImageIds,
	}
	return c.JSON(http.StatusCreated, res)
}

// GetReviews implements schema.ServerInterface.
func (h *Handler) GetReviews(c echo.Context, params schema.GetReviewsParams) error {
	reviews, err := h.repo.GetReviews(c.Request().Context())
	if err != nil {
		c.Logger().Errorf("failed to get reviews: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError).SetInternal(err)
	}

	resData := make([]schema.ReviewSummary, len(reviews))
	for i, reviews := range reviews {
		ImageIDs, err := h.repo.GetImageIDsByReviewID(c.Request().Context(), reviews.Id)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError).SetInternal(err)
		}

		resData[i] = schema.ReviewSummary{
			Id:        reviews.Id,
			EateryId:  reviews.EateryID,
			AuthorId:  reviews.UserID,
			CreatedAt: reviews.CreatedAt,
			UpdatedAt: reviews.UpdatedAt,
			ImageIds:  ImageIDs,
		}
	}
	res := schema.ReviewListResponse{
		Data: resData,
	}

	return c.JSON(http.StatusOK, res)
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
	ctx := c.Request().Context()

	// ユーザーID取得
	userID := getUserID(params.XForwardedUser)

	// リクエストボディをパース
	var req schema.ReviewUpdate
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request body")
	}

	// レビュー取得
	review, err := h.repo.GetReview(ctx, uuid.UUID(reviewId))
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "review not found")
	}

	// ユーザー自身か確認
	if review.UserID != userID {
		return echo.NewHTTPError(http.StatusForbidden, "you are not the author of this review")
	}

	// 内容を更新
	var content string
	if req.Content != nil {
		content = *req.Content
	} else {
		return echo.NewHTTPError(http.StatusBadRequest, "content cannot be null")
	}
	if err := h.repo.UpdateReviewContent(ctx, review.Id, content); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to update review").SetInternal(err)
	}

	// レスポンス生成（最新の内容で）
	res := schema.ReviewDetail{
		Id:       review.Id,
		EateryId: review.EateryID,
		AuthorId: review.UserID,
		Content:  *req.Content,
	}

	return c.JSON(http.StatusOK, res)
}
