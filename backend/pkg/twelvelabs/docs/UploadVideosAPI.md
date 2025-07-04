# \UploadVideosAPI

All URIs are relative to *https://api.twelvelabs.io/v1.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudToCloudImportVideos**](UploadVideosAPI.md#CloudToCloudImportVideos) | **Post** /tasks/transfers/import/{integration-id} | Import videos
[**CloudToCloudRetrieveImportLogs**](UploadVideosAPI.md#CloudToCloudRetrieveImportLogs) | **Get** /tasks/transfers/import/{integration-id}/logs | Retrieve import logs
[**CloudToCloudRetrieveStatus**](UploadVideosAPI.md#CloudToCloudRetrieveStatus) | **Get** /tasks/transfers/import/{integration-id}/status | Retrieve import status
[**CreateVideoIndexingTask**](UploadVideosAPI.md#CreateVideoIndexingTask) | **Post** /tasks | Create a video indexing task
[**DeleteVideoIndexingTask**](UploadVideosAPI.md#DeleteVideoIndexingTask) | **Delete** /tasks/{task_id} | Delete a video indexing task
[**ListVideoIndexingTasks**](UploadVideosAPI.md#ListVideoIndexingTasks) | **Get** /tasks | List video indexing tasks
[**RetrieveVideoIndexingTask**](UploadVideosAPI.md#RetrieveVideoIndexingTask) | **Get** /tasks/{task_id} | Retrieve a video indexing task



## CloudToCloudImportVideos

> InlineObject20 CloudToCloudImportVideos(ctx, integrationId).XApiKey(xApiKey).ContentType(contentType).CloudToCloudImportVideosRequest(cloudToCloudImportVideosRequest).Execute()

Import videos



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
	integrationId := "6298d673f1090f1100476d4c" // string | The unique identifier of the integration for which you want to import videos. You can retrieve it from the [Integrations](https://playground.twelvelabs.io/dashboard/integrations) page.
	cloudToCloudImportVideosRequest := *openapiclient.NewCloudToCloudImportVideosRequest("6298d673f1090f1100476d4c") // CloudToCloudImportVideosRequest | Request to import videos from a cloud storage bucket to an index. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UploadVideosAPI.CloudToCloudImportVideos(context.Background(), integrationId).XApiKey(xApiKey).ContentType(contentType).CloudToCloudImportVideosRequest(cloudToCloudImportVideosRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UploadVideosAPI.CloudToCloudImportVideos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudToCloudImportVideos`: InlineObject20
	fmt.Fprintf(os.Stdout, "Response from `UploadVideosAPI.CloudToCloudImportVideos`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**integrationId** | **string** | The unique identifier of the integration for which you want to import videos. You can retrieve it from the [Integrations](https://playground.twelvelabs.io/dashboard/integrations) page. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudToCloudImportVideosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiKey** | **string** | Your API key.  &lt;Note title&#x3D;\&quot;Note\&quot;&gt; You can find your API key on the &lt;a href&#x3D;\&quot;https://playground.twelvelabs.io/dashboard/api-key\&quot; target&#x3D;\&quot;_blank\&quot;&gt;API Key&lt;/a&gt; page. &lt;/Note&gt;  | 
 **contentType** | **string** | Must be set to &#x60;application/json&#x60;. | [default to &quot;application/json&quot;]

 **cloudToCloudImportVideosRequest** | [**CloudToCloudImportVideosRequest**](CloudToCloudImportVideosRequest.md) | Request to import videos from a cloud storage bucket to an index.  | 

### Return type

[**InlineObject20**](InlineObject20.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudToCloudRetrieveImportLogs

> InlineObject CloudToCloudRetrieveImportLogs(ctx, integrationId).XApiKey(xApiKey).ContentType(contentType).Execute()

Retrieve import logs



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
	integrationId := "6298d673f1090f1100476d4c" // string | The unique identifier of the integration for which you want to retrieve the import logs. You can retrieve it from the [Integrations](https://playground.twelvelabs.io/dashboard/integrations) page.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UploadVideosAPI.CloudToCloudRetrieveImportLogs(context.Background(), integrationId).XApiKey(xApiKey).ContentType(contentType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UploadVideosAPI.CloudToCloudRetrieveImportLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudToCloudRetrieveImportLogs`: InlineObject
	fmt.Fprintf(os.Stdout, "Response from `UploadVideosAPI.CloudToCloudRetrieveImportLogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**integrationId** | **string** | The unique identifier of the integration for which you want to retrieve the import logs. You can retrieve it from the [Integrations](https://playground.twelvelabs.io/dashboard/integrations) page. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudToCloudRetrieveImportLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiKey** | **string** | Your API key.  &lt;Note title&#x3D;\&quot;Note\&quot;&gt; You can find your API key on the &lt;a href&#x3D;\&quot;https://playground.twelvelabs.io/dashboard/api-key\&quot; target&#x3D;\&quot;_blank\&quot;&gt;API Key&lt;/a&gt; page. &lt;/Note&gt;  | 
 **contentType** | **string** | Must be set to &#x60;application/json&#x60;. | [default to &quot;application/json&quot;]


### Return type

[**InlineObject**](InlineObject.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudToCloudRetrieveStatus

> InlineObject21 CloudToCloudRetrieveStatus(ctx, integrationId).XApiKey(xApiKey).ContentType(contentType).IndexId(indexId).Execute()

Retrieve import status



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
	integrationId := "6298d673f1090f1100476d4c" // string | The unique identifier of the integration for which you want to retrieve the status of your imported videos. You can retrieve it from the [Integrations](https://playground.twelvelabs.io/dashboard/integrations) page.
	indexId := "6298d673f1090f1100476d4c" // string | The unique identifier of the index for which you want to retrieve the status of your imported videos. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UploadVideosAPI.CloudToCloudRetrieveStatus(context.Background(), integrationId).XApiKey(xApiKey).ContentType(contentType).IndexId(indexId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UploadVideosAPI.CloudToCloudRetrieveStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudToCloudRetrieveStatus`: InlineObject21
	fmt.Fprintf(os.Stdout, "Response from `UploadVideosAPI.CloudToCloudRetrieveStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**integrationId** | **string** | The unique identifier of the integration for which you want to retrieve the status of your imported videos. You can retrieve it from the [Integrations](https://playground.twelvelabs.io/dashboard/integrations) page. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudToCloudRetrieveStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiKey** | **string** | Your API key.  &lt;Note title&#x3D;\&quot;Note\&quot;&gt; You can find your API key on the &lt;a href&#x3D;\&quot;https://playground.twelvelabs.io/dashboard/api-key\&quot; target&#x3D;\&quot;_blank\&quot;&gt;API Key&lt;/a&gt; page. &lt;/Note&gt;  | 
 **contentType** | **string** | Must be set to &#x60;application/json&#x60;. | [default to &quot;application/json&quot;]

 **indexId** | **string** | The unique identifier of the index for which you want to retrieve the status of your imported videos.  | 

### Return type

[**InlineObject21**](InlineObject21.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateVideoIndexingTask

> InlineObject8 CreateVideoIndexingTask(ctx).XApiKey(xApiKey).ContentType(contentType).IndexId(indexId).VideoFile(videoFile).VideoUrl(videoUrl).EnableVideoStream(enableVideoStream).Execute()

Create a video indexing task



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
	contentType := "multipart/form-data" // string | Must be set to `multipart/form-data`. (default to "multipart/form-data")
	indexId := "indexId_example" // string | The unique identifier of the index to which the video is being uploaded. 
	videoFile := os.NewFile(1234, "some_file") // *os.File | Specify this parameter to upload a video from your local file system.  (optional)
	videoUrl := "videoUrl_example" // string | Specify this parameter to upload a video from a publicly accessible URL.  (optional)
	enableVideoStream := true // bool | This parameter indicates if the platform stores the video for streaming. When set to `true`, the platform stores the video, and you can retrieve its URL by calling the [`GET`](/v1.3/api-reference/videos/retrieve) method of the `/indexes/{index-id}/videos/{video-id}` endpoint. You can then use this URL to access the stream over the <a href=\\\"https://en.wikipedia.org/wiki/HTTP_Live_Streaming\\\" target=\\\"_blank\\\">HLS</a> protocol.  (optional) (default to true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UploadVideosAPI.CreateVideoIndexingTask(context.Background()).XApiKey(xApiKey).ContentType(contentType).IndexId(indexId).VideoFile(videoFile).VideoUrl(videoUrl).EnableVideoStream(enableVideoStream).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UploadVideosAPI.CreateVideoIndexingTask``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateVideoIndexingTask`: InlineObject8
	fmt.Fprintf(os.Stdout, "Response from `UploadVideosAPI.CreateVideoIndexingTask`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateVideoIndexingTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiKey** | **string** | Your API key.  &lt;Note title&#x3D;\&quot;Note\&quot;&gt; You can find your API key on the &lt;a href&#x3D;\&quot;https://playground.twelvelabs.io/dashboard/api-key\&quot; target&#x3D;\&quot;_blank\&quot;&gt;API Key&lt;/a&gt; page. &lt;/Note&gt;  | 
 **contentType** | **string** | Must be set to &#x60;multipart/form-data&#x60;. | [default to &quot;multipart/form-data&quot;]
 **indexId** | **string** | The unique identifier of the index to which the video is being uploaded.  | 
 **videoFile** | ***os.File** | Specify this parameter to upload a video from your local file system.  | 
 **videoUrl** | **string** | Specify this parameter to upload a video from a publicly accessible URL.  | 
 **enableVideoStream** | **bool** | This parameter indicates if the platform stores the video for streaming. When set to &#x60;true&#x60;, the platform stores the video, and you can retrieve its URL by calling the [&#x60;GET&#x60;](/v1.3/api-reference/videos/retrieve) method of the &#x60;/indexes/{index-id}/videos/{video-id}&#x60; endpoint. You can then use this URL to access the stream over the &lt;a href&#x3D;\\\&quot;https://en.wikipedia.org/wiki/HTTP_Live_Streaming\\\&quot; target&#x3D;\\\&quot;_blank\\\&quot;&gt;HLS&lt;/a&gt; protocol.  | [default to true]

### Return type

[**InlineObject8**](InlineObject8.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteVideoIndexingTask

> DeleteVideoIndexingTask(ctx, taskId).XApiKey(xApiKey).ContentType(contentType).Execute()

Delete a video indexing task



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
	taskId := "6298d673f1090f1100476d4c" // string | The unique identifier of the video indexing task you want to delete. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UploadVideosAPI.DeleteVideoIndexingTask(context.Background(), taskId).XApiKey(xApiKey).ContentType(contentType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UploadVideosAPI.DeleteVideoIndexingTask``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string** | The unique identifier of the video indexing task you want to delete.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVideoIndexingTaskRequest struct via the builder pattern


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


## ListVideoIndexingTasks

> InlineObject5 ListVideoIndexingTasks(ctx).XApiKey(xApiKey).ContentType(contentType).Page(page).PageLimit(pageLimit).SortBy(sortBy).SortOption(sortOption).IndexId(indexId).Status(status).Filename(filename).Duration(duration).Width(width).Height(height).CreatedAt(createdAt).UpdatedAt(updatedAt).Execute()

List video indexing tasks



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
	indexId := "630aff993fcee0532cb809d0" // string | Filter by the unique identifier of an index.  (optional)
	status := []string{"ready"} // []string | Filter by one or more video indexing task statuses. The following options are available: - `ready`: The video has been successfully uploaded and indexed. - `uploading`: The video is being uploaded. - `validating`: The video is being validated against the prerequisites. - `pending`: The video is pending. - `queued`: The video is queued. - `indexing`: The video is being indexed. - `failed`: The video indexing task failed.  To filter by multiple statuses, specify the `status` parameter for each value: ``` status=ready&status=validating ```  (optional)
	filename := "01.mp4" // string | Filter by filename.  (optional)
	duration := float32(531.998133) // float32 | Filter by duration. Expressed in seconds.  (optional)
	width := int32(640) // int32 | Filter by width.  (optional)
	height := int32(360) // int32 | Filter by height.  (optional)
	createdAt := "2024-03-01T00:00:00Z" // string | Filter video indexing tasks by the creation date and time, in the RFC 3339 format (\"YYYY-MM-DDTHH:mm:ssZ\"). The platform returns the video indexing tasks that were created on the specified date at or after the given time.  (optional)
	updatedAt := "2024-03-01T00:00:00Z" // string | Filter video indexing tasks by the last update date and time, in the RFC 3339 format (\"YYYY-MM-DDTHH:mm:ssZ\"). The platform returns the video indexing tasks that were updated on the specified date at or after the given time.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UploadVideosAPI.ListVideoIndexingTasks(context.Background()).XApiKey(xApiKey).ContentType(contentType).Page(page).PageLimit(pageLimit).SortBy(sortBy).SortOption(sortOption).IndexId(indexId).Status(status).Filename(filename).Duration(duration).Width(width).Height(height).CreatedAt(createdAt).UpdatedAt(updatedAt).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UploadVideosAPI.ListVideoIndexingTasks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListVideoIndexingTasks`: InlineObject5
	fmt.Fprintf(os.Stdout, "Response from `UploadVideosAPI.ListVideoIndexingTasks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListVideoIndexingTasksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiKey** | **string** | Your API key.  &lt;Note title&#x3D;\&quot;Note\&quot;&gt; You can find your API key on the &lt;a href&#x3D;\&quot;https://playground.twelvelabs.io/dashboard/api-key\&quot; target&#x3D;\&quot;_blank\&quot;&gt;API Key&lt;/a&gt; page. &lt;/Note&gt;  | 
 **contentType** | **string** | Must be set to &#x60;application/json&#x60;. | [default to &quot;application/json&quot;]
 **page** | **int32** | A number that identifies the page to retrieve.  **Default**: &#x60;1&#x60;.  | [default to 1]
 **pageLimit** | **int32** | The number of items to return on each page.  **Default**: &#x60;10&#x60;. **Max**: &#x60;50&#x60;.  | [default to 10]
 **sortBy** | **string** | The field to sort on. The following options are available: - &#x60;updated_at&#x60;: Sorts by the time, in the RFC 3339 format (\&quot;YYYY-MM-DDTHH:mm:ssZ\&quot;), when the item was updated. - &#x60;created_at&#x60;: Sorts by the time, in the RFC 3339 format (\&quot;YYYY-MM-DDTHH:mm:ssZ\&quot;), when the item was created.  **Default**: &#x60;created_at&#x60;.  | [default to &quot;created_at&quot;]
 **sortOption** | **string** | The sorting direction. The following options are available: - &#x60;asc&#x60; - &#x60;desc&#x60;  **Default**: &#x60;desc&#x60;.  | [default to &quot;desc&quot;]
 **indexId** | **string** | Filter by the unique identifier of an index.  | 
 **status** | **[]string** | Filter by one or more video indexing task statuses. The following options are available: - &#x60;ready&#x60;: The video has been successfully uploaded and indexed. - &#x60;uploading&#x60;: The video is being uploaded. - &#x60;validating&#x60;: The video is being validated against the prerequisites. - &#x60;pending&#x60;: The video is pending. - &#x60;queued&#x60;: The video is queued. - &#x60;indexing&#x60;: The video is being indexed. - &#x60;failed&#x60;: The video indexing task failed.  To filter by multiple statuses, specify the &#x60;status&#x60; parameter for each value: &#x60;&#x60;&#x60; status&#x3D;ready&amp;status&#x3D;validating &#x60;&#x60;&#x60;  | 
 **filename** | **string** | Filter by filename.  | 
 **duration** | **float32** | Filter by duration. Expressed in seconds.  | 
 **width** | **int32** | Filter by width.  | 
 **height** | **int32** | Filter by height.  | 
 **createdAt** | **string** | Filter video indexing tasks by the creation date and time, in the RFC 3339 format (\&quot;YYYY-MM-DDTHH:mm:ssZ\&quot;). The platform returns the video indexing tasks that were created on the specified date at or after the given time.  | 
 **updatedAt** | **string** | Filter video indexing tasks by the last update date and time, in the RFC 3339 format (\&quot;YYYY-MM-DDTHH:mm:ssZ\&quot;). The platform returns the video indexing tasks that were updated on the specified date at or after the given time.  | 

### Return type

[**InlineObject5**](InlineObject5.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveVideoIndexingTask

> InlineObject6 RetrieveVideoIndexingTask(ctx, taskId).XApiKey(xApiKey).ContentType(contentType).Execute()

Retrieve a video indexing task



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
	taskId := "6298d673f1090f1100476d4c" // string | The unique identifier of the video indexing task to retrieve. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UploadVideosAPI.RetrieveVideoIndexingTask(context.Background(), taskId).XApiKey(xApiKey).ContentType(contentType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UploadVideosAPI.RetrieveVideoIndexingTask``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveVideoIndexingTask`: InlineObject6
	fmt.Fprintf(os.Stdout, "Response from `UploadVideosAPI.RetrieveVideoIndexingTask`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string** | The unique identifier of the video indexing task to retrieve.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveVideoIndexingTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiKey** | **string** | Your API key.  &lt;Note title&#x3D;\&quot;Note\&quot;&gt; You can find your API key on the &lt;a href&#x3D;\&quot;https://playground.twelvelabs.io/dashboard/api-key\&quot; target&#x3D;\&quot;_blank\&quot;&gt;API Key&lt;/a&gt; page. &lt;/Note&gt;  | 
 **contentType** | **string** | Must be set to &#x60;application/json&#x60;. | [default to &quot;application/json&quot;]


### Return type

[**InlineObject6**](InlineObject6.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

