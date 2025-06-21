# ReviewDetail

Detailed review information

## Properties

| Name           | Type                    | Description                               | Notes                  |
| -------------- | ----------------------- | ----------------------------------------- | ---------------------- |
| **id**         | **string**              | Unique identifier for the review          | [default to undefined] |
| **content**    | **string**              | Full content of the review                | [default to undefined] |
| **authorId**   | **string**              | ID of the review author                   | [default to undefined] |
| **eateryId**   | **string**              | ID of the eatery being reviewed           | [default to undefined] |
| **eateryName** | **string**              | Name of the eatery being reviewed         | [default to undefined] |
| **imageIds**   | **Array&lt;string&gt;** | Array of image IDs attached to the review | [default to undefined] |
| **createdAt**  | **string**              | Creation timestamp                        | [default to undefined] |
| **updatedAt**  | **string**              | Last update timestamp                     | [default to undefined] |

## Example

```typescript
import { ReviewDetail } from './api'

const instance: ReviewDetail = {
  id,
  content,
  authorId,
  eateryId,
  eateryName,
  imageIds,
  createdAt,
  updatedAt,
}
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
