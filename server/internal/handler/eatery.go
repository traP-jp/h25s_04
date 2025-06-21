package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime/types"

	"github.com/traP-jp/h25s_04/server/internal/schema"
)

// GetEateries implements schema.ServerInterface.
func (h *Handler) GetEateries(ctx echo.Context, params schema.GetEateriesParams) error {
	return ctx.JSON(http.StatusNotImplemented, schema.Error{
		Code:  "NOT_IMPLEMENTED",
		Error: "GetEateries endpoint is not implemented yet",
	})
}

// PostEateries implements schema.ServerInterface.
func (h *Handler) PostEateries(ctx echo.Context) error {
	return ctx.JSON(http.StatusNotImplemented, schema.Error{
		Code:  "NOT_IMPLEMENTED",
		Error: "PostEateries endpoint is not implemented yet",
	})
}

// GetEateriesEateryId implements schema.ServerInterface.
func (h *Handler) GetEateriesEateryId(ctx echo.Context, eateryId types.UUID) error {
	return ctx.JSON(http.StatusNotImplemented, schema.Error{
		Code:  "NOT_IMPLEMENTED",
		Error: "GetEateriesEateryId endpoint is not implemented yet",
	})
}

// PutEateriesEateryId implements schema.ServerInterface.
func (h *Handler) PutEateriesEateryId(ctx echo.Context, eateryId types.UUID) error {
	return ctx.JSON(http.StatusNotImplemented, schema.Error{
		Code:  "NOT_IMPLEMENTED",
		Error: "PutEateriesEateryId endpoint is not implemented yet",
	})
}

// GetEateriesEateryIdReviews implements schema.ServerInterface.
func (h *Handler) GetEateriesEateryIdReviews(ctx echo.Context, eateryId types.UUID, params schema.GetEateriesEateryIdReviewsParams) error {
	return ctx.JSON(http.StatusNotImplemented, schema.Error{
		Code:  "NOT_IMPLEMENTED",
		Error: "GetEateriesEateryIdReviews endpoint is not implemented yet",
	})
}

// PostEateriesEateryIdReviews implements schema.ServerInterface.
func (h *Handler) PostEateriesEateryIdReviews(ctx echo.Context, eateryId types.UUID) error {
	return ctx.JSON(http.StatusNotImplemented, schema.Error{
		Code:  "NOT_IMPLEMENTED",
		Error: "PostEateriesEateryIdReviews endpoint is not implemented yet",
	})
}
