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
	eateries, err := h.repo.GetEateries(c.Request().Context(), params)
	if err != nil {
		c.Logger().Errorf("failed to get eateries: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError).SetInternal(err)
	}

	resData := make([]schema.Eatery, len(eateries))
	for i, eatery := range eateries {
		resData[i] = schema.Eatery{
			Id:          types.UUID(eatery.ID),
			Name:        eatery.Name,
			Description: eatery.Description,
		}
	}

	res := schema.EateryListResponse{
		Data: resData,
		Pagination: schema.Pagination{
			Limit: 10,
			Page:  1,
		},
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
