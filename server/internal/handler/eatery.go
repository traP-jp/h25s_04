package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime/types"

	"github.com/traP-jp/h25s_04/server/internal/repository"
	"github.com/traP-jp/h25s_04/server/internal/schema"
)

// GetEateries implements schema.ServerInterface.
func (h *Handler) GetEateries(c echo.Context, params schema.GetEateriesParams) error {
	return c.JSON(http.StatusNotImplemented, schema.Error{
		Code:  "NOT_IMPLEMENTED",
		Error: "GetEateries endpoint is not implemented yet",
	})
}

// PostEateries implements schema.ServerInterface.
func (h *Handler) PostEateries(c echo.Context, params schema.PostEateriesParams) error {
	return c.JSON(http.StatusNotImplemented, schema.Error{
		Code:  "NOT_IMPLEMENTED",
		Error: "PostEateries endpoint is not implemented yet",
	})
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
	userID := getUserID(params.XForwardedUser)
	eatery := repository.Eatery{
		Name:        req.Name,
		Description: req.Description,
	}
	if err := h.repo.UpdateEatery(c.Request().Context(), eateryId, eatery, userID); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Failed to update eatery: %v", err))
	}
	return c.NoContent(http.StatusNoContent)
}

// GetEateriesEateryIdReviews implements schema.ServerInterface.
func (h *Handler) GetEateriesEateryIdReviews(c echo.Context, eateryId types.UUID, params schema.GetEateriesEateryIdReviewsParams) error {
	return c.JSON(http.StatusNotImplemented, schema.Error{
		Code:  "NOT_IMPLEMENTED",
		Error: "GetEateriesEateryIdReviews endpoint is not implemented yet",
	})
}

// PostEateriesEateryIdReviews implements schema.ServerInterface.
func (h *Handler) PostEateriesEateryIdReviews(c echo.Context, eateryId types.UUID, params schema.PostEateriesEateryIdReviewsParams) error {
	return c.JSON(http.StatusNotImplemented, schema.Error{
		Code:  "NOT_IMPLEMENTED",
		Error: "PostEateriesEateryIdReviews endpoint is not implemented yet",
	})
}
