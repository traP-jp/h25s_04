package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime/types"

	"github.com/traP-jp/h25s_04/server/internal/schema"
)

// GetEateries implements schema.ServerInterface.
func (h *Handler) GetEateries(c echo.Context, params schema.GetEateriesParams) error {
	eateries, err := h.repo.GetEateries(c.Request().Context())
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

	// クエリパラメータが指定されている場合は、フィルタリングを行う
	if params.Query != nil {
		filteredEateries := make([]schema.Eatery, 0)
		for _, eatery := range res {
			if eatery.Name == *params.Query {
				filteredEateries = append(filteredEateries, eatery)
			}
		}
		return c.JSON(http.StatusOK, filteredEateries)
	}

	// クエリパラメータが空の場合は、全ての飲食店を返す
	return c.JSON(http.StatusOK, res)
}

// PostEateries implements schema.ServerInterface.
func (h *Handler) PostEateries(c echo.Context) error {
	return c.JSON(http.StatusNotImplemented, schema.Error{
		Code:  "NOT_IMPLEMENTED",
		Error: "PostEateries endpoint is not implemented yet",
	})
}

// GetEateriesEateryId implements schema.ServerInterface.
func (h *Handler) GetEateriesEateryId(c echo.Context, eateryId types.UUID) error {
	return c.JSON(http.StatusNotImplemented, schema.Error{
		Code:  "NOT_IMPLEMENTED",
		Error: "GetEateriesEateryId endpoint is not implemented yet",
	})
}

// PutEateriesEateryId implements schema.ServerInterface.
func (h *Handler) PutEateriesEateryId(c echo.Context, eateryId types.UUID) error {
	return c.JSON(http.StatusNotImplemented, schema.Error{
		Code:  "NOT_IMPLEMENTED",
		Error: "PutEateriesEateryId endpoint is not implemented yet",
	})
}

// GetEateriesEateryIdReviews implements schema.ServerInterface.
func (h *Handler) GetEateriesEateryIdReviews(c echo.Context, eateryId types.UUID, params schema.GetEateriesEateryIdReviewsParams) error {
	return c.JSON(http.StatusNotImplemented, schema.Error{
		Code:  "NOT_IMPLEMENTED",
		Error: "GetEateriesEateryIdReviews endpoint is not implemented yet",
	})
}

// PostEateriesEateryIdReviews implements schema.ServerInterface.
func (h *Handler) PostEateriesEateryIdReviews(c echo.Context, eateryId types.UUID) error {
	return c.JSON(http.StatusNotImplemented, schema.Error{
		Code:  "NOT_IMPLEMENTED",
		Error: "PostEateriesEateryIdReviews endpoint is not implemented yet",
	})
}
