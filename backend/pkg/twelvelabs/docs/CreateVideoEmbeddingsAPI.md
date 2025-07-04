# \CreateVideoEmbeddingsAPI

All URIs are relative to *https://api.twelvelabs.io/v1.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateVideoEmbeddingTask**](CreateVideoEmbeddingsAPI.md#CreateVideoEmbeddingTask) | **Post** /embed/tasks | Create a video embedding task
[**ListVideoEmbeddingTasks**](CreateVideoEmbeddingsAPI.md#ListVideoEmbeddingTasks) | **Get** /embed/tasks | List video embedding tasks
[**RetrieveVideoEmbedding**](CreateVideoEmbeddingsAPI.md#RetrieveVideoEmbedding) | **Get** /embed/tasks/{task_id} | Retrieve video embeddings
[**RetrieveVideoEmbeddingTask**](CreateVideoEmbeddingsAPI.md#RetrieveVideoEmbeddingTask) | **Get** /embed/tasks/{task_id}/status | Retrieve the status of a video embedding task



## CreateVideoEmbeddingTask

> InlineObject18 CreateVideoEmbeddingTask(ctx).XApiKey(xApiKey).ContentType(contentType).ModelName(modelName).VideoFile(videoFile).VideoUrl(videoUrl).VideoStartOffsetSec(videoStartOffsetSec).VideoEndOffsetSec(videoEndOffsetSec).VideoClipLength(videoClipLength).VideoEmbeddingScope(videoEmbeddingScope).Execute()

Create a video embedding task



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
	modelName := "modelName_example" // string | The name of the model you want to use. The following models are available:   - `Marengo-retrieval-2.7` 
	videoFile := os.NewFile(1234, "some_file") // *os.File | Specify this parameter to upload a video from your local file system.  (optional)
	videoUrl := "videoUrl_example" // string | Specify this parameter to upload a video from a publicly accessible URL.  (optional)
	videoStartOffsetSec := float32(3.4) // float32 | The start offset in seconds from the beginning of the video where processing should begin. Specifying 0 means starting from the beginning of the video.  **Default**: 0 **Min**: 0 **Max**: Duration of the video minus video_clip_length  (optional)
	videoEndOffsetSec := float32(3.4) // float32 | The end offset in seconds from the beginning of the video where processing should stop.  Ensure the following when you specify this parameter: - The end offset does not exceed the total duration of the video file. - The end offset is greater than the start offset. - You must set both the start and end offsets. Setting only one of these offsets is not permitted, resulting in an error.  **Min**: video_start_offset + video_clip_length **Max**: Duration of the video file  (optional)
	videoClipLength := float32(8.14) // float32 | The desired duration in seconds for each clip for which the platform generates an embedding. Ensure that the clip length does not exceed the interval between the start and end offsets.  **Default**: 6 **Min**: 2 **Max**: 10  (optional)
	videoEmbeddingScope := "videoEmbeddingScope_example" // string | Defines the scope of video embedding generation. Valid values are the following: - `clip`: Creates embeddings for each video segment of `video_clip_length` seconds, from `video_start_offset_sec` to `video_end_offset_sec`. - `clip` and `video`: Creates embeddings for video segments and the entire video.  To create embeddings for segments and the entire video in the same request, include this parameter twice as shown below:  ```json --form video_embedding_scope=clip \\\\ --form video_embedding_scope=video ```  **Default**: `clip`  (optional) (default to "clip")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CreateVideoEmbeddingsAPI.CreateVideoEmbeddingTask(context.Background()).XApiKey(xApiKey).ContentType(contentType).ModelName(modelName).VideoFile(videoFile).VideoUrl(videoUrl).VideoStartOffsetSec(videoStartOffsetSec).VideoEndOffsetSec(videoEndOffsetSec).VideoClipLength(videoClipLength).VideoEmbeddingScope(videoEmbeddingScope).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreateVideoEmbeddingsAPI.CreateVideoEmbeddingTask``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateVideoEmbeddingTask`: InlineObject18
	fmt.Fprintf(os.Stdout, "Response from `CreateVideoEmbeddingsAPI.CreateVideoEmbeddingTask`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateVideoEmbeddingTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiKey** | **string** | Your API key.  &lt;Note title&#x3D;\&quot;Note\&quot;&gt; You can find your API key on the &lt;a href&#x3D;\&quot;https://playground.twelvelabs.io/dashboard/api-key\&quot; target&#x3D;\&quot;_blank\&quot;&gt;API Key&lt;/a&gt; page. &lt;/Note&gt;  | 
 **contentType** | **string** | Must be set to &#x60;multipart/form-data&#x60;. | [default to &quot;multipart/form-data&quot;]
 **modelName** | **string** | The name of the model you want to use. The following models are available:   - &#x60;Marengo-retrieval-2.7&#x60;  | 
 **videoFile** | ***os.File** | Specify this parameter to upload a video from your local file system.  | 
 **videoUrl** | **string** | Specify this parameter to upload a video from a publicly accessible URL.  | 
 **videoStartOffsetSec** | **float32** | The start offset in seconds from the beginning of the video where processing should begin. Specifying 0 means starting from the beginning of the video.  **Default**: 0 **Min**: 0 **Max**: Duration of the video minus video_clip_length  | 
 **videoEndOffsetSec** | **float32** | The end offset in seconds from the beginning of the video where processing should stop.  Ensure the following when you specify this parameter: - The end offset does not exceed the total duration of the video file. - The end offset is greater than the start offset. - You must set both the start and end offsets. Setting only one of these offsets is not permitted, resulting in an error.  **Min**: video_start_offset + video_clip_length **Max**: Duration of the video file  | 
 **videoClipLength** | **float32** | The desired duration in seconds for each clip for which the platform generates an embedding. Ensure that the clip length does not exceed the interval between the start and end offsets.  **Default**: 6 **Min**: 2 **Max**: 10  | 
 **videoEmbeddingScope** | **string** | Defines the scope of video embedding generation. Valid values are the following: - &#x60;clip&#x60;: Creates embeddings for each video segment of &#x60;video_clip_length&#x60; seconds, from &#x60;video_start_offset_sec&#x60; to &#x60;video_end_offset_sec&#x60;. - &#x60;clip&#x60; and &#x60;video&#x60;: Creates embeddings for video segments and the entire video.  To create embeddings for segments and the entire video in the same request, include this parameter twice as shown below:  &#x60;&#x60;&#x60;json --form video_embedding_scope&#x3D;clip \\\\ --form video_embedding_scope&#x3D;video &#x60;&#x60;&#x60;  **Default**: &#x60;clip&#x60;  | [default to &quot;clip&quot;]

### Return type

[**InlineObject18**](InlineObject18.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVideoEmbeddingTasks

> InlineObject16 ListVideoEmbeddingTasks(ctx).XApiKey(xApiKey).StartedAt(startedAt).EndedAt(endedAt).Status(status).Page(page).PageLimit(pageLimit).Execute()

List video embedding tasks



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
	startedAt := "2024-03-01T00:00:00Z" // string | Retrieve the video embedding tasks that were created after the given date and time, expressed in the RFC 3339 format (\"YYYY-MM-DDTHH:mm:ssZ\").  (optional)
	endedAt := "2024-03-01T00:00:00Z" // string | Retrieve the video embedding tasks that were created before the given date and time, expressed in the RFC 3339 format (\"YYYY-MM-DDTHH:mm:ssZ\").  (optional)
	status := "processing" // string | Filter video embedding tasks by their current status. Possible values are `processing`, `ready`, or `failed`. (optional)
	page := int32(1) // int32 | A number that identifies the page to retrieve.  **Default**: `1`.  (optional) (default to 1)
	pageLimit := int32(10) // int32 | The number of items to return on each page.  **Default**: `10`. **Max**: `50`.  (optional) (default to 10)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CreateVideoEmbeddingsAPI.ListVideoEmbeddingTasks(context.Background()).XApiKey(xApiKey).StartedAt(startedAt).EndedAt(endedAt).Status(status).Page(page).PageLimit(pageLimit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreateVideoEmbeddingsAPI.ListVideoEmbeddingTasks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListVideoEmbeddingTasks`: InlineObject16
	fmt.Fprintf(os.Stdout, "Response from `CreateVideoEmbeddingsAPI.ListVideoEmbeddingTasks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListVideoEmbeddingTasksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiKey** | **string** | Your API key.  &lt;Note title&#x3D;\&quot;Note\&quot;&gt; You can find your API key on the &lt;a href&#x3D;\&quot;https://playground.twelvelabs.io/dashboard/api-key\&quot; target&#x3D;\&quot;_blank\&quot;&gt;API Key&lt;/a&gt; page. &lt;/Note&gt;  | 
 **startedAt** | **string** | Retrieve the video embedding tasks that were created after the given date and time, expressed in the RFC 3339 format (\&quot;YYYY-MM-DDTHH:mm:ssZ\&quot;).  | 
 **endedAt** | **string** | Retrieve the video embedding tasks that were created before the given date and time, expressed in the RFC 3339 format (\&quot;YYYY-MM-DDTHH:mm:ssZ\&quot;).  | 
 **status** | **string** | Filter video embedding tasks by their current status. Possible values are &#x60;processing&#x60;, &#x60;ready&#x60;, or &#x60;failed&#x60;. | 
 **page** | **int32** | A number that identifies the page to retrieve.  **Default**: &#x60;1&#x60;.  | [default to 1]
 **pageLimit** | **int32** | The number of items to return on each page.  **Default**: &#x60;10&#x60;. **Max**: &#x60;50&#x60;.  | [default to 10]

### Return type

[**InlineObject16**](InlineObject16.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveVideoEmbedding

> InlineObject15 RetrieveVideoEmbedding(ctx, taskId).XApiKey(xApiKey).ContentType(contentType).EmbeddingOption(embeddingOption).Execute()

Retrieve video embeddings



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
	taskId := "663da73b31cdd0c1f638a8e6" // string | The unique identifier of your video embedding task. 
	embeddingOption := []string{"EmbeddingOption_example"} // []string | Specifies which types of embeddings to retrieve. You can include one or more of the following values:   - `visual-text`:  Returns visual embeddings optimized for text search.   - `audio`: Returns audio embeddings.  The platform returns all available embeddings if you don't provide this parameter.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CreateVideoEmbeddingsAPI.RetrieveVideoEmbedding(context.Background(), taskId).XApiKey(xApiKey).ContentType(contentType).EmbeddingOption(embeddingOption).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreateVideoEmbeddingsAPI.RetrieveVideoEmbedding``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveVideoEmbedding`: InlineObject15
	fmt.Fprintf(os.Stdout, "Response from `CreateVideoEmbeddingsAPI.RetrieveVideoEmbedding`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string** | The unique identifier of your video embedding task.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveVideoEmbeddingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiKey** | **string** | Your API key.  &lt;Note title&#x3D;\&quot;Note\&quot;&gt; You can find your API key on the &lt;a href&#x3D;\&quot;https://playground.twelvelabs.io/dashboard/api-key\&quot; target&#x3D;\&quot;_blank\&quot;&gt;API Key&lt;/a&gt; page. &lt;/Note&gt;  | 
 **contentType** | **string** | Must be set to &#x60;application/json&#x60;. | [default to &quot;application/json&quot;]

 **embeddingOption** | **[]string** | Specifies which types of embeddings to retrieve. You can include one or more of the following values:   - &#x60;visual-text&#x60;:  Returns visual embeddings optimized for text search.   - &#x60;audio&#x60;: Returns audio embeddings.  The platform returns all available embeddings if you don&#39;t provide this parameter.  | 

### Return type

[**InlineObject15**](InlineObject15.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveVideoEmbeddingTask

> InlineObject17 RetrieveVideoEmbeddingTask(ctx, taskId).XApiKey(xApiKey).ContentType(contentType).Execute()

Retrieve the status of a video embedding task



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
	taskId := "663da73b31cdd0c1f638a8e6" // string | The unique identifier of your video embedding task. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CreateVideoEmbeddingsAPI.RetrieveVideoEmbeddingTask(context.Background(), taskId).XApiKey(xApiKey).ContentType(contentType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreateVideoEmbeddingsAPI.RetrieveVideoEmbeddingTask``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveVideoEmbeddingTask`: InlineObject17
	fmt.Fprintf(os.Stdout, "Response from `CreateVideoEmbeddingsAPI.RetrieveVideoEmbeddingTask`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string** | The unique identifier of your video embedding task.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveVideoEmbeddingTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiKey** | **string** | Your API key.  &lt;Note title&#x3D;\&quot;Note\&quot;&gt; You can find your API key on the &lt;a href&#x3D;\&quot;https://playground.twelvelabs.io/dashboard/api-key\&quot; target&#x3D;\&quot;_blank\&quot;&gt;API Key&lt;/a&gt; page. &lt;/Note&gt;  | 
 **contentType** | **string** | Must be set to &#x60;application/json&#x60;. | [default to &quot;application/json&quot;]


### Return type

[**InlineObject17**](InlineObject17.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

