# ReviewsApi

All URIs are relative to *http://localhost:8080/api/v1*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**eateriesEateryIdReviewsGet**](#eaterieseateryidreviewsget) | **GET** /eateries/{eateryId}/reviews | Get reviews for a specific eatery|
|[**eateriesEateryIdReviewsPost**](#eaterieseateryidreviewspost) | **POST** /eateries/{eateryId}/reviews | Create a review for an eatery|
|[**reviewsGet**](#reviewsget) | **GET** /reviews | Get all reviews (summary for top page)|
|[**reviewsReviewIdDelete**](#reviewsreviewiddelete) | **DELETE** /reviews/{reviewId} | Delete review|
|[**reviewsReviewIdGet**](#reviewsreviewidget) | **GET** /reviews/{reviewId} | Get review details|
|[**reviewsReviewIdPut**](#reviewsreviewidput) | **PUT** /reviews/{reviewId} | Update review|

# **eateriesEateryIdReviewsGet**
> EateriesEateryIdReviewsGet200Response eateriesEateryIdReviewsGet()

Returns all reviews for a specific eatery

### Example

```typescript
import {
    ReviewsApi,
    Configuration
} from './api';

const configuration = new Configuration();
const apiInstance = new ReviewsApi(configuration);

let eateryId: string; //UUID of the eatery (default to undefined)
let sortBy: 'newest' | 'oldest'; //Sort order for reviews (optional) (default to 'newest')
let page: number; //Page number for pagination (optional) (default to 1)
let limit: number; //Number of items per page (optional) (default to 20)

const { status, data } = await apiInstance.eateriesEateryIdReviewsGet(
    eateryId,
    sortBy,
    page,
    limit
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **eateryId** | [**string**] | UUID of the eatery | defaults to undefined|
| **sortBy** | [**&#39;newest&#39; | &#39;oldest&#39;**]**Array<&#39;newest&#39; &#124; &#39;oldest&#39;>** | Sort order for reviews | (optional) defaults to 'newest'|
| **page** | [**number**] | Page number for pagination | (optional) defaults to 1|
| **limit** | [**number**] | Number of items per page | (optional) defaults to 20|


### Return type

**EateriesEateryIdReviewsGet200Response**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Successful response |  -  |
|**404** | Eatery not found |  -  |
|**500** | Internal server error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **eateriesEateryIdReviewsPost**
> ReviewDetail eateriesEateryIdReviewsPost(reviewCreate)

Creates a new review for a specific eatery

### Example

```typescript
import {
    ReviewsApi,
    Configuration,
    ReviewCreate
} from './api';

const configuration = new Configuration();
const apiInstance = new ReviewsApi(configuration);

let eateryId: string; //UUID of the eatery (default to undefined)
let reviewCreate: ReviewCreate; //

const { status, data } = await apiInstance.eateriesEateryIdReviewsPost(
    eateryId,
    reviewCreate
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **reviewCreate** | **ReviewCreate**|  | |
| **eateryId** | [**string**] | UUID of the eatery | defaults to undefined|


### Return type

**ReviewDetail**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**201** | Review created successfully |  -  |
|**400** | Bad request |  -  |
|**404** | Eatery not found |  -  |
|**500** | Internal server error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **reviewsGet**
> EateriesEateryIdReviewsGet200Response reviewsGet()

Returns a list of all reviews across all eateries with images only, optimized for the top page

### Example

```typescript
import {
    ReviewsApi,
    Configuration
} from './api';

const configuration = new Configuration();
const apiInstance = new ReviewsApi(configuration);

let sortBy: 'newest' | 'oldest'; //Sort order for reviews (optional) (default to 'newest')
let page: number; //Page number for pagination (optional) (default to 1)
let limit: number; //Number of items per page (optional) (default to 20)

const { status, data } = await apiInstance.reviewsGet(
    sortBy,
    page,
    limit
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **sortBy** | [**&#39;newest&#39; | &#39;oldest&#39;**]**Array<&#39;newest&#39; &#124; &#39;oldest&#39;>** | Sort order for reviews | (optional) defaults to 'newest'|
| **page** | [**number**] | Page number for pagination | (optional) defaults to 1|
| **limit** | [**number**] | Number of items per page | (optional) defaults to 20|


### Return type

**EateriesEateryIdReviewsGet200Response**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Successful response |  -  |
|**500** | Internal server error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **reviewsReviewIdDelete**
> reviewsReviewIdDelete()

Deletes a specific review

### Example

```typescript
import {
    ReviewsApi,
    Configuration
} from './api';

const configuration = new Configuration();
const apiInstance = new ReviewsApi(configuration);

let reviewId: string; //UUID of the review (default to undefined)

const { status, data } = await apiInstance.reviewsReviewIdDelete(
    reviewId
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **reviewId** | [**string**] | UUID of the review | defaults to undefined|


### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**204** | Review deleted successfully |  -  |
|**404** | Review not found |  -  |
|**500** | Internal server error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **reviewsReviewIdGet**
> ReviewDetail reviewsReviewIdGet()

Returns detailed information about a specific review

### Example

```typescript
import {
    ReviewsApi,
    Configuration
} from './api';

const configuration = new Configuration();
const apiInstance = new ReviewsApi(configuration);

let reviewId: string; //UUID of the review (default to undefined)

const { status, data } = await apiInstance.reviewsReviewIdGet(
    reviewId
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **reviewId** | [**string**] | UUID of the review | defaults to undefined|


### Return type

**ReviewDetail**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Successful response |  -  |
|**404** | Review not found |  -  |
|**500** | Internal server error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **reviewsReviewIdPut**
> ReviewDetail reviewsReviewIdPut(reviewUpdate)

Updates an existing review

### Example

```typescript
import {
    ReviewsApi,
    Configuration,
    ReviewUpdate
} from './api';

const configuration = new Configuration();
const apiInstance = new ReviewsApi(configuration);

let reviewId: string; //UUID of the review (default to undefined)
let reviewUpdate: ReviewUpdate; //

const { status, data } = await apiInstance.reviewsReviewIdPut(
    reviewId,
    reviewUpdate
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **reviewUpdate** | **ReviewUpdate**|  | |
| **reviewId** | [**string**] | UUID of the review | defaults to undefined|


### Return type

**ReviewDetail**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Review updated successfully |  -  |
|**400** | Bad request |  -  |
|**404** | Review not found |  -  |
|**500** | Internal server error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

