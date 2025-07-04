# \ManageIndexesAPI

All URIs are relative to *https://api.twelvelabs.io/v1.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateIndex**](ManageIndexesAPI.md#CreateIndex) | **Post** /indexes | Create an index
[**DeleteIndex**](ManageIndexesAPI.md#DeleteIndex) | **Delete** /indexes/{index-id} | Delete an index
[**ListIndexes**](ManageIndexesAPI.md#ListIndexes) | **Get** /indexes | List indexes
[**RetrieveIndex**](ManageIndexesAPI.md#RetrieveIndex) | **Get** /indexes/{index-id} | Retrieve an index
[**UpdateIndex**](ManageIndexesAPI.md#UpdateIndex) | **Put** /indexes/{index-id} | Update an index



## CreateIndex

> InlineObject9 CreateIndex(ctx).XApiKey(xApiKey).ContentType(contentType).CreateIndexRequest(createIndexRequest).Execute()

Create an index



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/xyberii4/lec-assist/backend/pkg/twelvelabs"
)

func main() {
	xApiKey := "tlk_1234567890" // string | Your API key.  <Note title=\"Note\"> You can find your API key on the <a href=\"https://playground.twelvelabs.io/dashboard/api-key\" target=\"_blank\">API Key</a> page. </Note> 
	contentType := "application/json" // string | Must be set to `application/json`. (default to "application/json")
	createIndexRequest := *openapiclient.NewCreateIndexRequest("myIndex", []openapiclient.CreateIndexRequestModelsInner{*openapiclient.NewCreateIndexRequestModelsInner("marengo2.7", []string{"visual"})}) // CreateIndexRequest | Request to create an index. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManageIndexesAPI.CreateIndex(context.Background()).XApiKey(xApiKey).ContentType(contentType).CreateIndexRequest(createIndexRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManageIndexesAPI.CreateIndex``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateIndex`: InlineObject9
	fmt.Fprintf(os.Stdout, "Response from `ManageIndexesAPI.CreateIndex`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiKey** | **string** | Your API key.  &lt;Note title&#x3D;\&quot;Note\&quot;&gt; You can find your API key on the &lt;a href&#x3D;\&quot;https://playground.twelvelabs.io/dashboard/api-key\&quot; target&#x3D;\&quot;_blank\&quot;&gt;API Key&lt;/a&gt; page. &lt;/Note&gt;  | 
 **contentType** | **string** | Must be set to &#x60;application/json&#x60;. | [default to &quot;application/json&quot;]
 **createIndexRequest** | [**CreateIndexRequest**](CreateIndexRequest.md) | Request to create an index.  | 

### Return type

[**InlineObject9**](InlineObject9.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIndex

> DeleteIndex(ctx, indexId).XApiKey(xApiKey).ContentType(contentType).Execute()

Delete an index



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/xyberii4/lec-assist/backend/pkg/twelvelabs"
)

func main() {
	xApiKey := "tlk_1234567890" // string | Your API key.  <Note title=\"Note\"> You can find your API key on the <a href=\"https://playground.twelvelabs.io/dashboard/api-key\" target=\"_blank\">API Key</a> page. </Note> 
	contentType := "application/json" // string | Must be set to `application/json`. (default to "application/json")
	indexId := "6298d673f1090f1100476d4c" // string | Unique identifier of the index to delete. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ManageIndexesAPI.DeleteIndex(context.Background(), indexId).XApiKey(xApiKey).ContentType(contentType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManageIndexesAPI.DeleteIndex``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**indexId** | **string** | Unique identifier of the index to delete.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiKey** | **string** | Your API key.  &lt;Note title&#x3D;\&quot;Note\&quot;&gt; You can find your API key on the &lt;a href&#x3D;\&quot;https://playground.twelvelabs.io/dashboard/api-key\&quot; target&#x3D;\&quot;_blank\&quot;&gt;API Key&lt;/a&gt; page. &lt;/Note&gt;  | 
 **contentType** | **string** | Must be set to &#x60;application/json&#x60;. | [default to &quot;application/json&quot;]


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIndexes

> InlineObject7 ListIndexes(ctx).XApiKey(xApiKey).ContentType(contentType).Page(page).PageLimit(pageLimit).SortBy(sortBy).SortOption(sortOption).IndexName(indexName).ModelOptions(modelOptions).ModelFamily(modelFamily).CreatedAt(createdAt).UpdatedAt(updatedAt).Execute()

List indexes



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/xyberii4/lec-assist/backend/pkg/twelvelabs"
)

func main() {
	xApiKey := "tlk_1234567890" // string | Your API key.  <Note title=\"Note\"> You can find your API key on the <a href=\"https://playground.twelvelabs.io/dashboard/api-key\" target=\"_blank\">API Key</a> page. </Note> 
	contentType := "application/json" // string | Must be set to `application/json`. (default to "application/json")
	page := int32(1) // int32 | A number that identifies the page to retrieve.  **Default**: `1`.  (optional) (default to 1)
	pageLimit := int32(10) // int32 | The number of items to return on each page.  **Default**: `10`. **Max**: `50`.  (optional) (default to 10)
	sortBy := "created_at" // string | The field to sort on. The following options are available: - `updated_at`: Sorts by the time, in the RFC 3339 format (\"YYYY-MM-DDTHH:mm:ssZ\"), when the item was updated. - `created_at`: Sorts by the time, in the RFC 3339 format (\"YYYY-MM-DDTHH:mm:ssZ\"), when the item was created.  **Default**: `created_at`.  (optional) (default to "created_at")
	sortOption := "desc" // string | The sorting direction. The following options are available: - `asc` - `desc`  **Default**: `desc`.  (optional) (default to "desc")
	indexName := "myIndex" // string | Filter by the name of an index. (optional)
	modelOptions := "visual,audio" // string | Filter by the model options. When filtering by multiple model options, the values must be comma-separated.  (optional)
	modelFamily := "marengo" // string | Filter by the model family. This parameter can take one of the following values: `marengo` or `pegasus`. You can specify a single value.  (optional)
	createdAt := "2024-08-16T16:53:59Z" // string | Filter indexes by the creation date and time, in the RFC 3339 format (\"YYYY-MM-DDTHH:mm:ssZ\"). The platform returns the indexes that were created on the specified date at or after the given time.  (optional)
	updatedAt := "2024-08-16T16:55:59Z" // string | Filter indexes by the last update date and time, in the RFC 3339 format (\"YYYY-MM-DDTHH:mm:ssZ\"). The platform returns the indexes that were last updated on the specified date at or after the given time.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManageIndexesAPI.ListIndexes(context.Background()).XApiKey(xApiKey).ContentType(contentType).Page(page).PageLimit(pageLimit).SortBy(sortBy).SortOption(sortOption).IndexName(indexName).ModelOptions(modelOptions).ModelFamily(modelFamily).CreatedAt(createdAt).UpdatedAt(updatedAt).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManageIndexesAPI.ListIndexes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListIndexes`: InlineObject7
	fmt.Fprintf(os.Stdout, "Response from `ManageIndexesAPI.ListIndexes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListIndexesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiKey** | **string** | Your API key.  &lt;Note title&#x3D;\&quot;Note\&quot;&gt; You can find your API key on the &lt;a href&#x3D;\&quot;https://playground.twelvelabs.io/dashboard/api-key\&quot; target&#x3D;\&quot;_blank\&quot;&gt;API Key&lt;/a&gt; page. &lt;/Note&gt;  | 
 **contentType** | **string** | Must be set to &#x60;application/json&#x60;. | [default to &quot;application/json&quot;]
 **page** | **int32** | A number that identifies the page to retrieve.  **Default**: &#x60;1&#x60;.  | [default to 1]
 **pageLimit** | **int32** | The number of items to return on each page.  **Default**: &#x60;10&#x60;. **Max**: &#x60;50&#x60;.  | [default to 10]
 **sortBy** | **string** | The field to sort on. The following options are available: - &#x60;updated_at&#x60;: Sorts by the time, in the RFC 3339 format (\&quot;YYYY-MM-DDTHH:mm:ssZ\&quot;), when the item was updated. - &#x60;created_at&#x60;: Sorts by the time, in the RFC 3339 format (\&quot;YYYY-MM-DDTHH:mm:ssZ\&quot;), when the item was created.  **Default**: &#x60;created_at&#x60;.  | [default to &quot;created_at&quot;]
 **sortOption** | **string** | The sorting direction. The following options are available: - &#x60;asc&#x60; - &#x60;desc&#x60;  **Default**: &#x60;desc&#x60;.  | [default to &quot;desc&quot;]
 **indexName** | **string** | Filter by the name of an index. | 
 **modelOptions** | **string** | Filter by the model options. When filtering by multiple model options, the values must be comma-separated.  | 
 **modelFamily** | **string** | Filter by the model family. This parameter can take one of the following values: &#x60;marengo&#x60; or &#x60;pegasus&#x60;. You can specify a single value.  | 
 **createdAt** | **string** | Filter indexes by the creation date and time, in the RFC 3339 format (\&quot;YYYY-MM-DDTHH:mm:ssZ\&quot;). The platform returns the indexes that were created on the specified date at or after the given time.  | 
 **updatedAt** | **string** | Filter indexes by the last update date and time, in the RFC 3339 format (\&quot;YYYY-MM-DDTHH:mm:ssZ\&quot;). The platform returns the indexes that were last updated on the specified date at or after the given time.  | 

### Return type

[**InlineObject7**](InlineObject7.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveIndex

> Index RetrieveIndex(ctx, indexId).XApiKey(xApiKey).ContentType(contentType).Execute()

Retrieve an index



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/xyberii4/lec-assist/backend/pkg/twelvelabs"
)

func main() {
	xApiKey := "tlk_1234567890" // string | Your API key.  <Note title=\"Note\"> You can find your API key on the <a href=\"https://playground.twelvelabs.io/dashboard/api-key\" target=\"_blank\">API Key</a> page. </Note> 
	contentType := "application/json" // string | Must be set to `application/json`. (default to "application/json")
	indexId := "6298d673f1090f1100476d4c" // string | Unique identifier of the index to retrieve. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManageIndexesAPI.RetrieveIndex(context.Background(), indexId).XApiKey(xApiKey).ContentType(contentType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManageIndexesAPI.RetrieveIndex``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveIndex`: Index
	fmt.Fprintf(os.Stdout, "Response from `ManageIndexesAPI.RetrieveIndex`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**indexId** | **string** | Unique identifier of the index to retrieve.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiKey** | **string** | Your API key.  &lt;Note title&#x3D;\&quot;Note\&quot;&gt; You can find your API key on the &lt;a href&#x3D;\&quot;https://playground.twelvelabs.io/dashboard/api-key\&quot; target&#x3D;\&quot;_blank\&quot;&gt;API Key&lt;/a&gt; page. &lt;/Note&gt;  | 
 **contentType** | **string** | Must be set to &#x60;application/json&#x60;. | [default to &quot;application/json&quot;]


### Return type

[**Index**](Index.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIndex

> UpdateIndex(ctx, indexId).XApiKey(xApiKey).ContentType(contentType).UpdateIndexRequest(updateIndexRequest).Execute()

Update an index



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/xyberii4/lec-assist/backend/pkg/twelvelabs"
)

func main() {
	xApiKey := "tlk_1234567890" // string | Your API key.  <Note title=\"Note\"> You can find your API key on the <a href=\"https://playground.twelvelabs.io/dashboard/api-key\" target=\"_blank\">API Key</a> page. </Note> 
	contentType := "application/json" // string | Must be set to `application/json`. (default to "application/json")
	indexId := "6298d673f1090f1100476d4c" // string | Unique identifier of the index to update. 
	updateIndexRequest := *openapiclient.NewUpdateIndexRequest("myIndex") // UpdateIndexRequest | Request to update the name of an index.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ManageIndexesAPI.UpdateIndex(context.Background(), indexId).XApiKey(xApiKey).ContentType(contentType).UpdateIndexRequest(updateIndexRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManageIndexesAPI.UpdateIndex``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**indexId** | **string** | Unique identifier of the index to update.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiKey** | **string** | Your API key.  &lt;Note title&#x3D;\&quot;Note\&quot;&gt; You can find your API key on the &lt;a href&#x3D;\&quot;https://playground.twelvelabs.io/dashboard/api-key\&quot; target&#x3D;\&quot;_blank\&quot;&gt;API Key&lt;/a&gt; page. &lt;/Note&gt;  | 
 **contentType** | **string** | Must be set to &#x60;application/json&#x60;. | [default to &quot;application/json&quot;]

 **updateIndexRequest** | [**UpdateIndexRequest**](UpdateIndexRequest.md) | Request to update the name of an index.  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

