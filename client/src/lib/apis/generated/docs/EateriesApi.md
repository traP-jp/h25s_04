# EateriesApi

All URIs are relative to *http://localhost:8080/api/v1*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**eateriesEateryIdGet**](#eaterieseateryidget) | **GET** /eateries/{eateryId} | Get eatery by ID|
|[**eateriesEateryIdPut**](#eaterieseateryidput) | **PUT** /eateries/{eateryId} | Update eatery|
|[**eateriesGet**](#eateriesget) | **GET** /eateries | Get list of eateries|
|[**eateriesPost**](#eateriespost) | **POST** /eateries | Create a new eatery|

# **eateriesEateryIdGet**
> Eatery eateriesEateryIdGet()

Returns a specific eatery by its ID

### Example

```typescript
import {
    EateriesApi,
    Configuration
} from './api';

const configuration = new Configuration();
const apiInstance = new EateriesApi(configuration);

let eateryId: string; //UUID of the eatery (default to undefined)

const { status, data } = await apiInstance.eateriesEateryIdGet(
    eateryId
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **eateryId** | [**string**] | UUID of the eatery | defaults to undefined|


### Return type

**Eatery**

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

# **eateriesEateryIdPut**
> Eatery eateriesEateryIdPut(eateryUpdate)

Updates an existing eatery\'s information

### Example

```typescript
import {
    EateriesApi,
    Configuration,
    EateryUpdate
} from './api';

const configuration = new Configuration();
const apiInstance = new EateriesApi(configuration);

let eateryId: string; //UUID of the eatery (default to undefined)
let eateryUpdate: EateryUpdate; //
let xForwardedUser: string; //ログインしているユーザーのtraQ ID（NeoShowcaseが自動で付与） (optional) (default to undefined)

const { status, data } = await apiInstance.eateriesEateryIdPut(
    eateryId,
    eateryUpdate,
    xForwardedUser
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **eateryUpdate** | **EateryUpdate**|  | |
| **eateryId** | [**string**] | UUID of the eatery | defaults to undefined|
| **xForwardedUser** | [**string**] | ログインしているユーザーのtraQ ID（NeoShowcaseが自動で付与） | (optional) defaults to undefined|


### Return type

**Eatery**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Eatery updated successfully |  -  |
|**400** | Bad request |  -  |
|**404** | Eatery not found |  -  |
|**500** | Internal server error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **eateriesGet**
> EateryListResponse eateriesGet()

Returns a list of all eateries with optional query filtering

### Example

```typescript
import {
    EateriesApi,
    Configuration
} from './api';

const configuration = new Configuration();
const apiInstance = new EateriesApi(configuration);

let query: string; //Search query to filter eateries by name or description (optional) (default to undefined)
let page: number; //Page number for pagination (optional) (default to 1)
let limit: number; //Number of items per page (optional) (default to 20)

const { status, data } = await apiInstance.eateriesGet(
    query,
    page,
    limit
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **query** | [**string**] | Search query to filter eateries by name or description | (optional) defaults to undefined|
| **page** | [**number**] | Page number for pagination | (optional) defaults to 1|
| **limit** | [**number**] | Number of items per page | (optional) defaults to 20|


### Return type

**EateryListResponse**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Successful response |  -  |
|**400** | Bad request |  -  |
|**500** | Internal server error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **eateriesPost**
> Eatery eateriesPost(eateryCreate)

Creates a new eatery

### Example

```typescript
import {
    EateriesApi,
    Configuration,
    EateryCreate
} from './api';

const configuration = new Configuration();
const apiInstance = new EateriesApi(configuration);

let eateryCreate: EateryCreate; //
let xForwardedUser: string; //ログインしているユーザーのtraQ ID（NeoShowcaseが自動で付与） (optional) (default to undefined)

const { status, data } = await apiInstance.eateriesPost(
    eateryCreate,
    xForwardedUser
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **eateryCreate** | **EateryCreate**|  | |
| **xForwardedUser** | [**string**] | ログインしているユーザーのtraQ ID（NeoShowcaseが自動で付与） | (optional) defaults to undefined|


### Return type

**Eatery**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**201** | Eatery created successfully |  -  |
|**400** | Bad request |  -  |
|**500** | Internal server error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

