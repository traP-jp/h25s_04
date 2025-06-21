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

// GetEateries implements schema.ServerInterface.
func (h *Handler) GetEateries(c echo.Context, params schema.GetEateriesParams) error {
	eateries, err := h.repo.GetEateries(c.Request().Context(), params)
	if err != nil {
		c.Logger().Errorf("failed to get eateries: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError).SetInternal(err)
	}

	res := make([]schema.Eatery, len(eateries))
	for i, eatery := range eateries {
		res[i] = schema.Eatery{
			Id:          types.UUID(eatery.ID),
			Name:        eatery.Name,
			Description: eatery.Description,
		}
	}
	return c.JSON(http.StatusOK, res)
}

// PostEateries implements schema.ServerInterface.
func (h *Handler) PostEateries(c echo.Context, params schema.PostEateriesParams) error {
	var req schema.EateryCreate
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "bad request")
	}

	createParams := repository.CreateEateryParams{
		Name:        req.Name,
		Description: req.Description,
	}
	eateryID, err := h.repo.CreateEateries(c.Request().Context(), createParams)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, "fail to create eatery")
	}

	res := schema.Eatery{
		Id:          eateryID,
		Name:        req.Name,
		Description: req.Description,
	}

	return c.JSON(http.StatusOK, res)
}

// GetEateriesEateryId implements schema.ServerInterface.a
func (h *Handler) GetEateriesEateryId(c echo.Context, eateryId types.UUID) error {
	//Id := types.UUID(eateryId)
	eatery, err := h.repo.GetEatery(c.Request().Context(), eateryId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Failed to get eatery: %v", err))
	}
	// TODO: Return the eatery as JSON or handle as needed
	return c.JSON(http.StatusOK, eatery)

}

// PutEateriesEateryId implements schema.ServerInterface.
func (h *Handler) PutEateriesEateryId(c echo.Context, eateryId types.UUID, params schema.PutEateriesEateryIdParams) error {
	var req schema.EateryCreate
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Failed to bind request: %v", err))
	}
	eatery := repository.Eatery{
		Name:        req.Name,
		Description: req.Description,
	}
	if err := h.repo.UpdateEatery(c.Request().Context(), eateryId, eatery); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Failed to update eatery: %v", err))
	}
	res := schema.Eatery{
		Id:          eateryId,
		Name:        req.Name,
		Description: req.Description,
	}
	return c.JSON(http.StatusOK, res)
}

// GetEateriesEateryIdReviews implements schema.ServerInterface.
func (h *Handler) GetEateriesEateryIdReviews(c echo.Context, eateryId types.UUID, params schema.GetEateriesEateryIdReviewsParams) error {
	reviews, err := h.repo.GetEateryEateryIDReviews(c.Request().Context(), eateryId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError).SetInternal(err)
	}

	res := make([]schema.ReviewDetail, len(reviews))
	for i, review := range reviews {
		res[i] = schema.ReviewDetail{
			Id:       review.Id,
			EateryId: review.EateryID,
			AuthorId: review.UserID,
			Content:  review.Content,
		}
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
