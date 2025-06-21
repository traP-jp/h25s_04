# ImagesApi

All URIs are relative to _http://localhost:8080/api/v1_

| Method                                    | HTTP request              | Description    |
| ----------------------------------------- | ------------------------- | -------------- |
| [**imagesImageIdGet**](#imagesimageidget) | **GET** /images/{imageId} | Download image |
| [**imagesPost**](#imagespost)             | **POST** /images          | Upload image   |

# **imagesImageIdGet**

> File imagesImageIdGet()

Downloads an image by its ID

### Example

```typescript
import { ImagesApi, Configuration } from './api'

const configuration = new Configuration()
const apiInstance = new ImagesApi(configuration)

let imageId: string //UUID of the image (default to undefined)

const { status, data } = await apiInstance.imagesImageIdGet(imageId)
```

### Parameters

| Name        | Type         | Description       | Notes                 |
| ----------- | ------------ | ----------------- | --------------------- |
| **imageId** | [**string**] | UUID of the image | defaults to undefined |

### Return type

**File**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: image/jpeg, image/png, image/gif, image/webp, application/json

### HTTP response details

| Status code | Description           | Response headers |
| ----------- | --------------------- | ---------------- |
| **200**     | Image file            | -                |
| **404**     | Image not found       | -                |
| **500**     | Internal server error | -                |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **imagesPost**

> ImageUploadResponse imagesPost()

Uploads an image file and returns its ID

### Example

```typescript
import { ImagesApi, Configuration } from './api'

const configuration = new Configuration()
const apiInstance = new ImagesApi(configuration)

let image: File //Image file to upload (default to undefined)

const { status, data } = await apiInstance.imagesPost(image)
```

### Parameters

| Name      | Type       | Description          | Notes                 |
| --------- | ---------- | -------------------- | --------------------- |
| **image** | [**File**] | Image file to upload | defaults to undefined |

### Return type

**ImageUploadResponse**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

### HTTP response details

| Status code | Description                 | Response headers |
| ----------- | --------------------------- | ---------------- |
| **201**     | Image uploaded successfully | -                |
| **400**     | Bad request                 | -                |
| **500**     | Internal server error       | -                |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)
