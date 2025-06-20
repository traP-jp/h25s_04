// Package schema provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package schema

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

// Defines values for GetEateriesEateryIdReviewsParamsSortBy.
const (
	GetEateriesEateryIdReviewsParamsSortByNewest GetEateriesEateryIdReviewsParamsSortBy = "newest"
	GetEateriesEateryIdReviewsParamsSortByOldest GetEateriesEateryIdReviewsParamsSortBy = "oldest"
)

// Defines values for GetReviewsParamsSortBy.
const (
	GetReviewsParamsSortByNewest GetReviewsParamsSortBy = "newest"
	GetReviewsParamsSortByOldest GetReviewsParamsSortBy = "oldest"
)

// Eatery defines model for Eatery.
type Eatery struct {
	// CreatedAt Creation timestamp
	CreatedAt time.Time `json:"createdAt"`

	// Description Description of the eatery
	Description *string `json:"description,omitempty"`

	// Id Unique identifier for the eatery
	Id openapi_types.UUID `json:"id"`

	// Latitude 緯度
	Latitude float64 `json:"latitude"`

	// Longitude 経度
	Longitude float64 `json:"longitude"`

	// Name Name of the eatery
	Name string `json:"name"`

	// UpdatedAt Last update timestamp
	UpdatedAt time.Time `json:"updatedAt"`
}

// EateryCreate defines model for EateryCreate.
type EateryCreate struct {
	// Description Description of the eatery
	Description *string `json:"description,omitempty"`

	// Latitude 緯度
	Latitude float64 `json:"latitude"`

	// Longitude 経度
	Longitude float64 `json:"longitude"`

	// Name Name of the eatery
	Name string `json:"name"`
}

// EateryListResponse defines model for EateryListResponse.
type EateryListResponse struct {
	Data       []Eatery   `json:"data"`
	Pagination Pagination `json:"pagination"`
}

// EateryUpdate defines model for EateryUpdate.
type EateryUpdate struct {
	// Description Description of the eatery
	Description *string `json:"description,omitempty"`

	// Latitude 緯度
	Latitude *float64 `json:"latitude,omitempty"`

	// Longitude 経度
	Longitude *float64 `json:"longitude,omitempty"`

	// Name Name of the eatery
	Name *string `json:"name,omitempty"`
}

// Error defines model for Error.
type Error struct {
	// Code Error code
	Code string `json:"code"`

	// Details Additional error details
	Details *map[string]interface{} `json:"details,omitempty"`

	// Error Error message
	Error string `json:"error"`
}

// ImageUploadResponse defines model for ImageUploadResponse.
type ImageUploadResponse struct {
	// Id Unique identifier for the uploaded image
	Id openapi_types.UUID `json:"id"`

	// Url URL to access the uploaded image
	Url string `json:"url"`
}

// Pagination defines model for Pagination.
type Pagination struct {
	// Limit Number of items per page
	Limit int `json:"limit"`

	// Page Current page number
	Page int `json:"page"`

	// Total Total number of items
	Total int `json:"total"`

	// TotalPages Total number of pages
	TotalPages int `json:"totalPages"`
}

// ReviewCreate defines model for ReviewCreate.
type ReviewCreate struct {
	// AuthorId ID of the review author
	AuthorId string `json:"authorId"`

	// Content Content of the review
	Content string `json:"content"`

	// ImageIds Array of image IDs to attach to the review
	ImageIds *[]openapi_types.UUID `json:"imageIds,omitempty"`
}

// ReviewDetail Detailed review information
type ReviewDetail struct {
	// AuthorId ID of the review author
	AuthorId string `json:"authorId"`

	// Content Full content of the review
	Content string `json:"content"`

	// CreatedAt Creation timestamp
	CreatedAt time.Time `json:"createdAt"`

	// EateryId ID of the eatery being reviewed
	EateryId openapi_types.UUID `json:"eateryId"`

	// EateryName Name of the eatery being reviewed
	EateryName string `json:"eateryName"`

	// Id Unique identifier for the review
	Id openapi_types.UUID `json:"id"`

	// ImageIds Array of image IDs attached to the review
	ImageIds []openapi_types.UUID `json:"imageIds"`

	// UpdatedAt Last update timestamp
	UpdatedAt time.Time `json:"updatedAt"`
}

// ReviewDetailListResponse defines model for ReviewDetailListResponse.
type ReviewDetailListResponse struct {
	Data       []ReviewDetail `json:"data"`
	Pagination Pagination     `json:"pagination"`
}

// ReviewListResponse defines model for ReviewListResponse.
type ReviewListResponse struct {
	Data       []ReviewSummary `json:"data"`
	Pagination Pagination      `json:"pagination"`
}

// ReviewSummary Summary of a review for list views (without full content)
type ReviewSummary struct {
	// AuthorId ID of the review author
	AuthorId string `json:"authorId"`

	// CreatedAt Creation timestamp
	CreatedAt time.Time `json:"createdAt"`

	// EateryId ID of the eatery being reviewed
	EateryId openapi_types.UUID `json:"eateryId"`

	// EateryName Name of the eatery being reviewed
	EateryName string `json:"eateryName"`

	// Id Unique identifier for the review
	Id openapi_types.UUID `json:"id"`

	// ImageIds Array of image IDs attached to the review
	ImageIds []openapi_types.UUID `json:"imageIds"`

	// Summary Short summary or first few lines of the review
	Summary *string `json:"summary,omitempty"`

	// UpdatedAt Last update timestamp
	UpdatedAt time.Time `json:"updatedAt"`
}

// ReviewUpdate defines model for ReviewUpdate.
type ReviewUpdate struct {
	// Content Content of the review
	Content *string `json:"content,omitempty"`

	// ImageIds Array of image IDs to attach to the review
	ImageIds *[]openapi_types.UUID `json:"imageIds,omitempty"`
}

// XForwardedUser defines model for X-Forwarded-User.
type XForwardedUser = string

// GetEateriesParams defines parameters for GetEateries.
type GetEateriesParams struct {
	// Query Search query to filter eateries by name or description
	Query *string `form:"query,omitempty" json:"query,omitempty"`

	// Page Page number for pagination
	Page *int `form:"page,omitempty" json:"page,omitempty"`

	// Limit Number of items per page
	Limit *int `form:"limit,omitempty" json:"limit,omitempty"`
}

// PostEateriesParams defines parameters for PostEateries.
type PostEateriesParams struct {
	// XForwardedUser ログインしているユーザーのtraQ ID（NeoShowcaseが自動で付与）
	XForwardedUser *XForwardedUser `json:"X-Forwarded-User,omitempty"`
}

// PutEateriesEateryIdParams defines parameters for PutEateriesEateryId.
type PutEateriesEateryIdParams struct {
	// XForwardedUser ログインしているユーザーのtraQ ID（NeoShowcaseが自動で付与）
	XForwardedUser *XForwardedUser `json:"X-Forwarded-User,omitempty"`
}

// GetEateriesEateryIdReviewsParams defines parameters for GetEateriesEateryIdReviews.
type GetEateriesEateryIdReviewsParams struct {
	// SortBy Sort order for reviews
	SortBy *GetEateriesEateryIdReviewsParamsSortBy `form:"sortBy,omitempty" json:"sortBy,omitempty"`

	// Page Page number for pagination
	Page *int `form:"page,omitempty" json:"page,omitempty"`

	// Limit Number of items per page
	Limit *int `form:"limit,omitempty" json:"limit,omitempty"`
}

// GetEateriesEateryIdReviewsParamsSortBy defines parameters for GetEateriesEateryIdReviews.
type GetEateriesEateryIdReviewsParamsSortBy string

// PostEateriesEateryIdReviewsParams defines parameters for PostEateriesEateryIdReviews.
type PostEateriesEateryIdReviewsParams struct {
	// XForwardedUser ログインしているユーザーのtraQ ID（NeoShowcaseが自動で付与）
	XForwardedUser *XForwardedUser `json:"X-Forwarded-User,omitempty"`
}

// PostImagesMultipartBody defines parameters for PostImages.
type PostImagesMultipartBody struct {
	// Image Image file to upload
	Image openapi_types.File `json:"image"`
}

// GetReviewsParams defines parameters for GetReviews.
type GetReviewsParams struct {
	// SortBy Sort order for reviews
	SortBy *GetReviewsParamsSortBy `form:"sortBy,omitempty" json:"sortBy,omitempty"`

	// Page Page number for pagination
	Page *int `form:"page,omitempty" json:"page,omitempty"`

	// Limit Number of items per page
	Limit *int `form:"limit,omitempty" json:"limit,omitempty"`
}

// GetReviewsParamsSortBy defines parameters for GetReviews.
type GetReviewsParamsSortBy string

// DeleteReviewsReviewIdParams defines parameters for DeleteReviewsReviewId.
type DeleteReviewsReviewIdParams struct {
	// XForwardedUser ログインしているユーザーのtraQ ID（NeoShowcaseが自動で付与）
	XForwardedUser *XForwardedUser `json:"X-Forwarded-User,omitempty"`
}

// PutReviewsReviewIdParams defines parameters for PutReviewsReviewId.
type PutReviewsReviewIdParams struct {
	// XForwardedUser ログインしているユーザーのtraQ ID（NeoShowcaseが自動で付与）
	XForwardedUser *XForwardedUser `json:"X-Forwarded-User,omitempty"`
}

// PostEateriesJSONRequestBody defines body for PostEateries for application/json ContentType.
type PostEateriesJSONRequestBody = EateryCreate

// PutEateriesEateryIdJSONRequestBody defines body for PutEateriesEateryId for application/json ContentType.
type PutEateriesEateryIdJSONRequestBody = EateryUpdate

// PostEateriesEateryIdReviewsJSONRequestBody defines body for PostEateriesEateryIdReviews for application/json ContentType.
type PostEateriesEateryIdReviewsJSONRequestBody = ReviewCreate

// PostImagesMultipartRequestBody defines body for PostImages for multipart/form-data ContentType.
type PostImagesMultipartRequestBody PostImagesMultipartBody

// PutReviewsReviewIdJSONRequestBody defines body for PutReviewsReviewId for application/json ContentType.
type PutReviewsReviewIdJSONRequestBody = ReviewUpdate

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Get list of eateries
	// (GET /eateries)
	GetEateries(ctx echo.Context, params GetEateriesParams) error
	// Create a new eatery
	// (POST /eateries)
	PostEateries(ctx echo.Context, params PostEateriesParams) error
	// Get eatery by ID
	// (GET /eateries/{eateryId})
	GetEateriesEateryId(ctx echo.Context, eateryId openapi_types.UUID) error
	// Update eatery
	// (PUT /eateries/{eateryId})
	PutEateriesEateryId(ctx echo.Context, eateryId openapi_types.UUID, params PutEateriesEateryIdParams) error
	// Get reviews for a specific eatery
	// (GET /eateries/{eateryId}/reviews)
	GetEateriesEateryIdReviews(ctx echo.Context, eateryId openapi_types.UUID, params GetEateriesEateryIdReviewsParams) error
	// Create a review for an eatery
	// (POST /eateries/{eateryId}/reviews)
	PostEateriesEateryIdReviews(ctx echo.Context, eateryId openapi_types.UUID, params PostEateriesEateryIdReviewsParams) error
	// Upload image
	// (POST /images)
	PostImages(ctx echo.Context) error
	// Download image
	// (GET /images/{imageId})
	GetImagesImageId(ctx echo.Context, imageId openapi_types.UUID) error
	// Get all reviews (summary for top page)
	// (GET /reviews)
	GetReviews(ctx echo.Context, params GetReviewsParams) error
	// Delete review
	// (DELETE /reviews/{reviewId})
	DeleteReviewsReviewId(ctx echo.Context, reviewId openapi_types.UUID, params DeleteReviewsReviewIdParams) error
	// Get review details
	// (GET /reviews/{reviewId})
	GetReviewsReviewId(ctx echo.Context, reviewId openapi_types.UUID) error
	// Update review
	// (PUT /reviews/{reviewId})
	PutReviewsReviewId(ctx echo.Context, reviewId openapi_types.UUID, params PutReviewsReviewIdParams) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetEateries converts echo context to params.
func (w *ServerInterfaceWrapper) GetEateries(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetEateriesParams
	// ------------- Optional query parameter "query" -------------

	err = runtime.BindQueryParameter("form", true, false, "query", ctx.QueryParams(), &params.Query)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter query: %s", err))
	}

	// ------------- Optional query parameter "page" -------------

	err = runtime.BindQueryParameter("form", true, false, "page", ctx.QueryParams(), &params.Page)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter page: %s", err))
	}

	// ------------- Optional query parameter "limit" -------------

	err = runtime.BindQueryParameter("form", true, false, "limit", ctx.QueryParams(), &params.Limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter limit: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetEateries(ctx, params)
	return err
}

// PostEateries converts echo context to params.
func (w *ServerInterfaceWrapper) PostEateries(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params PostEateriesParams

	headers := ctx.Request().Header
	// ------------- Optional header parameter "X-Forwarded-User" -------------
	if valueList, found := headers[http.CanonicalHeaderKey("X-Forwarded-User")]; found {
		var XForwardedUser XForwardedUser
		n := len(valueList)
		if n != 1 {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Expected one value for X-Forwarded-User, got %d", n))
		}

		err = runtime.BindStyledParameterWithOptions("simple", "X-Forwarded-User", valueList[0], &XForwardedUser, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false})
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter X-Forwarded-User: %s", err))
		}

		params.XForwardedUser = &XForwardedUser
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PostEateries(ctx, params)
	return err
}

// GetEateriesEateryId converts echo context to params.
func (w *ServerInterfaceWrapper) GetEateriesEateryId(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "eateryId" -------------
	var eateryId openapi_types.UUID

	err = runtime.BindStyledParameterWithOptions("simple", "eateryId", ctx.Param("eateryId"), &eateryId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter eateryId: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetEateriesEateryId(ctx, eateryId)
	return err
}

// PutEateriesEateryId converts echo context to params.
func (w *ServerInterfaceWrapper) PutEateriesEateryId(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "eateryId" -------------
	var eateryId openapi_types.UUID

	err = runtime.BindStyledParameterWithOptions("simple", "eateryId", ctx.Param("eateryId"), &eateryId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter eateryId: %s", err))
	}

	// Parameter object where we will unmarshal all parameters from the context
	var params PutEateriesEateryIdParams

	headers := ctx.Request().Header
	// ------------- Optional header parameter "X-Forwarded-User" -------------
	if valueList, found := headers[http.CanonicalHeaderKey("X-Forwarded-User")]; found {
		var XForwardedUser XForwardedUser
		n := len(valueList)
		if n != 1 {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Expected one value for X-Forwarded-User, got %d", n))
		}

		err = runtime.BindStyledParameterWithOptions("simple", "X-Forwarded-User", valueList[0], &XForwardedUser, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false})
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter X-Forwarded-User: %s", err))
		}

		params.XForwardedUser = &XForwardedUser
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PutEateriesEateryId(ctx, eateryId, params)
	return err
}

// GetEateriesEateryIdReviews converts echo context to params.
func (w *ServerInterfaceWrapper) GetEateriesEateryIdReviews(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "eateryId" -------------
	var eateryId openapi_types.UUID

	err = runtime.BindStyledParameterWithOptions("simple", "eateryId", ctx.Param("eateryId"), &eateryId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter eateryId: %s", err))
	}

	// Parameter object where we will unmarshal all parameters from the context
	var params GetEateriesEateryIdReviewsParams
	// ------------- Optional query parameter "sortBy" -------------

	err = runtime.BindQueryParameter("form", true, false, "sortBy", ctx.QueryParams(), &params.SortBy)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter sortBy: %s", err))
	}

	// ------------- Optional query parameter "page" -------------

	err = runtime.BindQueryParameter("form", true, false, "page", ctx.QueryParams(), &params.Page)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter page: %s", err))
	}

	// ------------- Optional query parameter "limit" -------------

	err = runtime.BindQueryParameter("form", true, false, "limit", ctx.QueryParams(), &params.Limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter limit: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetEateriesEateryIdReviews(ctx, eateryId, params)
	return err
}

// PostEateriesEateryIdReviews converts echo context to params.
func (w *ServerInterfaceWrapper) PostEateriesEateryIdReviews(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "eateryId" -------------
	var eateryId openapi_types.UUID

	err = runtime.BindStyledParameterWithOptions("simple", "eateryId", ctx.Param("eateryId"), &eateryId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter eateryId: %s", err))
	}

	// Parameter object where we will unmarshal all parameters from the context
	var params PostEateriesEateryIdReviewsParams

	headers := ctx.Request().Header
	// ------------- Optional header parameter "X-Forwarded-User" -------------
	if valueList, found := headers[http.CanonicalHeaderKey("X-Forwarded-User")]; found {
		var XForwardedUser XForwardedUser
		n := len(valueList)
		if n != 1 {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Expected one value for X-Forwarded-User, got %d", n))
		}

		err = runtime.BindStyledParameterWithOptions("simple", "X-Forwarded-User", valueList[0], &XForwardedUser, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false})
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter X-Forwarded-User: %s", err))
		}

		params.XForwardedUser = &XForwardedUser
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PostEateriesEateryIdReviews(ctx, eateryId, params)
	return err
}

// PostImages converts echo context to params.
func (w *ServerInterfaceWrapper) PostImages(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PostImages(ctx)
	return err
}

// GetImagesImageId converts echo context to params.
func (w *ServerInterfaceWrapper) GetImagesImageId(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "imageId" -------------
	var imageId openapi_types.UUID

	err = runtime.BindStyledParameterWithOptions("simple", "imageId", ctx.Param("imageId"), &imageId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter imageId: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetImagesImageId(ctx, imageId)
	return err
}

// GetReviews converts echo context to params.
func (w *ServerInterfaceWrapper) GetReviews(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetReviewsParams
	// ------------- Optional query parameter "sortBy" -------------

	err = runtime.BindQueryParameter("form", true, false, "sortBy", ctx.QueryParams(), &params.SortBy)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter sortBy: %s", err))
	}

	// ------------- Optional query parameter "page" -------------

	err = runtime.BindQueryParameter("form", true, false, "page", ctx.QueryParams(), &params.Page)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter page: %s", err))
	}

	// ------------- Optional query parameter "limit" -------------

	err = runtime.BindQueryParameter("form", true, false, "limit", ctx.QueryParams(), &params.Limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter limit: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetReviews(ctx, params)
	return err
}

// DeleteReviewsReviewId converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteReviewsReviewId(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "reviewId" -------------
	var reviewId openapi_types.UUID

	err = runtime.BindStyledParameterWithOptions("simple", "reviewId", ctx.Param("reviewId"), &reviewId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter reviewId: %s", err))
	}

	// Parameter object where we will unmarshal all parameters from the context
	var params DeleteReviewsReviewIdParams

	headers := ctx.Request().Header
	// ------------- Optional header parameter "X-Forwarded-User" -------------
	if valueList, found := headers[http.CanonicalHeaderKey("X-Forwarded-User")]; found {
		var XForwardedUser XForwardedUser
		n := len(valueList)
		if n != 1 {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Expected one value for X-Forwarded-User, got %d", n))
		}

		err = runtime.BindStyledParameterWithOptions("simple", "X-Forwarded-User", valueList[0], &XForwardedUser, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false})
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter X-Forwarded-User: %s", err))
		}

		params.XForwardedUser = &XForwardedUser
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.DeleteReviewsReviewId(ctx, reviewId, params)
	return err
}

// GetReviewsReviewId converts echo context to params.
func (w *ServerInterfaceWrapper) GetReviewsReviewId(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "reviewId" -------------
	var reviewId openapi_types.UUID

	err = runtime.BindStyledParameterWithOptions("simple", "reviewId", ctx.Param("reviewId"), &reviewId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter reviewId: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetReviewsReviewId(ctx, reviewId)
	return err
}

// PutReviewsReviewId converts echo context to params.
func (w *ServerInterfaceWrapper) PutReviewsReviewId(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "reviewId" -------------
	var reviewId openapi_types.UUID

	err = runtime.BindStyledParameterWithOptions("simple", "reviewId", ctx.Param("reviewId"), &reviewId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter reviewId: %s", err))
	}

	// Parameter object where we will unmarshal all parameters from the context
	var params PutReviewsReviewIdParams

	headers := ctx.Request().Header
	// ------------- Optional header parameter "X-Forwarded-User" -------------
	if valueList, found := headers[http.CanonicalHeaderKey("X-Forwarded-User")]; found {
		var XForwardedUser XForwardedUser
		n := len(valueList)
		if n != 1 {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Expected one value for X-Forwarded-User, got %d", n))
		}

		err = runtime.BindStyledParameterWithOptions("simple", "X-Forwarded-User", valueList[0], &XForwardedUser, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false})
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter X-Forwarded-User: %s", err))
		}

		params.XForwardedUser = &XForwardedUser
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PutReviewsReviewId(ctx, reviewId, params)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/eateries", wrapper.GetEateries)
	router.POST(baseURL+"/eateries", wrapper.PostEateries)
	router.GET(baseURL+"/eateries/:eateryId", wrapper.GetEateriesEateryId)
	router.PUT(baseURL+"/eateries/:eateryId", wrapper.PutEateriesEateryId)
	router.GET(baseURL+"/eateries/:eateryId/reviews", wrapper.GetEateriesEateryIdReviews)
	router.POST(baseURL+"/eateries/:eateryId/reviews", wrapper.PostEateriesEateryIdReviews)
	router.POST(baseURL+"/images", wrapper.PostImages)
	router.GET(baseURL+"/images/:imageId", wrapper.GetImagesImageId)
	router.GET(baseURL+"/reviews", wrapper.GetReviews)
	router.DELETE(baseURL+"/reviews/:reviewId", wrapper.DeleteReviewsReviewId)
	router.GET(baseURL+"/reviews/:reviewId", wrapper.GetReviewsReviewId)
	router.PUT(baseURL+"/reviews/:reviewId", wrapper.PutReviewsReviewId)

}
