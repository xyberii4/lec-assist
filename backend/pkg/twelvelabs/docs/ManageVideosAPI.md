# \ManageVideosAPI

All URIs are relative to *https://api.twelvelabs.io/v1.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteVideoInformation**](ManageVideosAPI.md#DeleteVideoInformation) | **Delete** /indexes/{index-id}/videos/{video-id} | Delete video information
[**ListVideos**](ManageVideosAPI.md#ListVideos) | **Get** /indexes/{index-id}/videos | List videos
[**RetrieveVideoInformation**](ManageVideosAPI.md#RetrieveVideoInformation) | **Get** /indexes/{index-id}/videos/{video-id} | Retrieve video information
[**UpdateVideoInformation**](ManageVideosAPI.md#UpdateVideoInformation) | **Put** /indexes/{index-id}/videos/{video-id} | Update video information



## DeleteVideoInformation

> DeleteVideoInformation(ctx, indexId, videoId).XApiKey(xApiKey).ContentType(contentType).Execute()

Delete video information



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
	indexId := "6298d673f1090f1100476d4c" // string | The unique identifier of the index to which the video has been uploaded. 
	videoId := "6298d673f1090f1100476d4c" // string | The unique identifier of the video to delete. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ManageVideosAPI.DeleteVideoInformation(context.Background(), indexId, videoId).XApiKey(xApiKey).ContentType(contentType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManageVideosAPI.DeleteVideoInformation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**indexId** | **string** | The unique identifier of the index to which the video has been uploaded.  | 
**videoId** | **string** | The unique identifier of the video to delete.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVideoInformationRequest struct via the builder pattern


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


## ListVideos

> InlineObject3 ListVideos(ctx, indexId).XApiKey(xApiKey).ContentType(contentType).Page(page).PageLimit(pageLimit).SortBy(sortBy).SortOption(sortOption).Filename(filename).Duration(duration).Fps(fps).Width(width).Height(height).Size(size).CreatedAt(createdAt).UpdatedAt(updatedAt).UserMetadata(userMetadata).Execute()

List videos



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
	indexId := "6298d673f1090f1100476d4c" // string | The unique identifier of the index for which the API will retrieve the videos.
	page := int32(1) // int32 | A number that identifies the page to retrieve.  **Default**: `1`.  (optional) (default to 1)
	pageLimit := int32(10) // int32 | The number of items to return on each page.  **Default**: `10`. **Max**: `50`.  (optional) (default to 10)
	sortBy := "created_at" // string | The field to sort on. The following options are available: - `updated_at`: Sorts by the time, in the RFC 3339 format (\"YYYY-MM-DDTHH:mm:ssZ\"), when the item was updated. - `created_at`: Sorts by the time, in the RFC 3339 format (\"YYYY-MM-DDTHH:mm:ssZ\"), when the item was created.  **Default**: `created_at`.  (optional) (default to "created_at")
	sortOption := "desc" // string | The sorting direction. The following options are available: - `asc` - `desc`  **Default**: `desc`.  (optional) (default to "desc")
	filename := "01.mp4" // string | Filter by filename.  (optional)
	duration := float32(10) // float32 | Filter by duration. Expressed in seconds.  (optional)
	fps := float32(25) // float32 | Filter by frames per second.  (optional)
	width := float32(1920) // float32 | Filter by width.  (optional)
	height := int32(1080) // int32 | Filter by height.  (optional)
	size := float32(1048576) // float32 | Filter by size. Expressed in bytes.  (optional)
	createdAt := "2024-08-16T16:53:59Z" // string | Filter videos by the creation date and time of their associated indexing tasks, in the RFC 3339 format (\"YYYY-MM-DDTHH:mm:ssZ\"). The platform returns the videos whose indexing tasks were created on the specified date at or after the given time.  (optional)
	updatedAt := "2024-08-16T16:53:59Z" // string | This filter applies only to videos updated using the [`PUT`](/v1.3/api-reference/videos/update) method of the `/indexes/{index-id}/videos/{video-id}` endpoint. It filters videos by the last update date and time, in the RFC 3339 format (\"YYYY-MM-DDTHH:mm:ssZ\"). The platform returns the video indexing tasks that were last updated on the specified date at or after the given time.  (optional)
	userMetadata := map[string]ListVideosUserMetadataParameterValue{"key": openapiclient.list_videos_user_metadata_parameter_value{Bool: new(bool)}} // map[string]ListVideosUserMetadataParameterValue | To enable filtering by custom fields, you must first add user-defined metadata to your video by calling the [`PUT`](/v1.3/api-reference/videos/update) method of the `/indexes/:index-id/videos/:video-id` endpoint.  Examples: - To filter on a string: `?category=recentlyAdded` - To filter on an integer: `?batchNumber=5` - To filter on a float: `?rating=9.3` - To filter on a boolean: `?needsReview=true`  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManageVideosAPI.ListVideos(context.Background(), indexId).XApiKey(xApiKey).ContentType(contentType).Page(page).PageLimit(pageLimit).SortBy(sortBy).SortOption(sortOption).Filename(filename).Duration(duration).Fps(fps).Width(width).Height(height).Size(size).CreatedAt(createdAt).UpdatedAt(updatedAt).UserMetadata(userMetadata).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManageVideosAPI.ListVideos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListVideos`: InlineObject3
	fmt.Fprintf(os.Stdout, "Response from `ManageVideosAPI.ListVideos`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**indexId** | **string** | The unique identifier of the index for which the API will retrieve the videos. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListVideosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiKey** | **string** | Your API key.  &lt;Note title&#x3D;\&quot;Note\&quot;&gt; You can find your API key on the &lt;a href&#x3D;\&quot;https://playground.twelvelabs.io/dashboard/api-key\&quot; target&#x3D;\&quot;_blank\&quot;&gt;API Key&lt;/a&gt; page. &lt;/Note&gt;  | 
 **contentType** | **string** | Must be set to &#x60;application/json&#x60;. | [default to &quot;application/json&quot;]

 **page** | **int32** | A number that identifies the page to retrieve.  **Default**: &#x60;1&#x60;.  | [default to 1]
 **pageLimit** | **int32** | The number of items to return on each page.  **Default**: &#x60;10&#x60;. **Max**: &#x60;50&#x60;.  | [default to 10]
 **sortBy** | **string** | The field to sort on. The following options are available: - &#x60;updated_at&#x60;: Sorts by the time, in the RFC 3339 format (\&quot;YYYY-MM-DDTHH:mm:ssZ\&quot;), when the item was updated. - &#x60;created_at&#x60;: Sorts by the time, in the RFC 3339 format (\&quot;YYYY-MM-DDTHH:mm:ssZ\&quot;), when the item was created.  **Default**: &#x60;created_at&#x60;.  | [default to &quot;created_at&quot;]
 **sortOption** | **string** | The sorting direction. The following options are available: - &#x60;asc&#x60; - &#x60;desc&#x60;  **Default**: &#x60;desc&#x60;.  | [default to &quot;desc&quot;]
 **filename** | **string** | Filter by filename.  | 
 **duration** | **float32** | Filter by duration. Expressed in seconds.  | 
 **fps** | **float32** | Filter by frames per second.  | 
 **width** | **float32** | Filter by width.  | 
 **height** | **int32** | Filter by height.  | 
 **size** | **float32** | Filter by size. Expressed in bytes.  | 
 **createdAt** | **string** | Filter videos by the creation date and time of their associated indexing tasks, in the RFC 3339 format (\&quot;YYYY-MM-DDTHH:mm:ssZ\&quot;). The platform returns the videos whose indexing tasks were created on the specified date at or after the given time.  | 
 **updatedAt** | **string** | This filter applies only to videos updated using the [&#x60;PUT&#x60;](/v1.3/api-reference/videos/update) method of the &#x60;/indexes/{index-id}/videos/{video-id}&#x60; endpoint. It filters videos by the last update date and time, in the RFC 3339 format (\&quot;YYYY-MM-DDTHH:mm:ssZ\&quot;). The platform returns the video indexing tasks that were last updated on the specified date at or after the given time.  | 
 **userMetadata** | [**map[string]ListVideosUserMetadataParameterValue**](ListVideosUserMetadataParameterValue.md) | To enable filtering by custom fields, you must first add user-defined metadata to your video by calling the [&#x60;PUT&#x60;](/v1.3/api-reference/videos/update) method of the &#x60;/indexes/:index-id/videos/:video-id&#x60; endpoint.  Examples: - To filter on a string: &#x60;?category&#x3D;recentlyAdded&#x60; - To filter on an integer: &#x60;?batchNumber&#x3D;5&#x60; - To filter on a float: &#x60;?rating&#x3D;9.3&#x60; - To filter on a boolean: &#x60;?needsReview&#x3D;true&#x60;  | 

### Return type

[**InlineObject3**](InlineObject3.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveVideoInformation

> InlineObject4 RetrieveVideoInformation(ctx, indexId, videoId).XApiKey(xApiKey).ContentType(contentType).EmbeddingOption(embeddingOption).Transcription(transcription).Execute()

Retrieve video information



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
	indexId := "6298d673f1090f1100476d4c" // string | The unique identifier of the index to which the video has been uploaded. 
	videoId := "6298d673f1090f1100476d4c" // string | The unique identifier of the video to retrieve. 
	embeddingOption := []string{"EmbeddingOption_example"} // []string | Specifies which types of embeddings to retrieve. You can include one or more of the following values: - `visual-text`:  Returns visual embeddings optimized for text search. - `audio`: Returns audio embeddings. <br/> To retrieve embeddings for a video, it must be indexed using the Marengo video understanding model version 2.7 or later. For details on enabling this model for an index, see the [Create an index](/reference/create-index) page.  The platform does not return embeddings if you don't provide this parameter.  The values you specify in `embedding_option` must be included in the `model_options` defined when the index was created. For example, if `model_options` is set to `visual,` you cannot set `embedding_option` to `audio` or  both `visual-text` and `audio`.  (optional)
	transcription := true // bool | The parameter indicates whether to retrieve a transcription of the spoken words for the indexed video. Note that the official SDKs will support this feature in a future release.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManageVideosAPI.RetrieveVideoInformation(context.Background(), indexId, videoId).XApiKey(xApiKey).ContentType(contentType).EmbeddingOption(embeddingOption).Transcription(transcription).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManageVideosAPI.RetrieveVideoInformation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveVideoInformation`: InlineObject4
	fmt.Fprintf(os.Stdout, "Response from `ManageVideosAPI.RetrieveVideoInformation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**indexId** | **string** | The unique identifier of the index to which the video has been uploaded.  | 
**videoId** | **string** | The unique identifier of the video to retrieve.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveVideoInformationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiKey** | **string** | Your API key.  &lt;Note title&#x3D;\&quot;Note\&quot;&gt; You can find your API key on the &lt;a href&#x3D;\&quot;https://playground.twelvelabs.io/dashboard/api-key\&quot; target&#x3D;\&quot;_blank\&quot;&gt;API Key&lt;/a&gt; page. &lt;/Note&gt;  | 
 **contentType** | **string** | Must be set to &#x60;application/json&#x60;. | [default to &quot;application/json&quot;]


 **embeddingOption** | **[]string** | Specifies which types of embeddings to retrieve. You can include one or more of the following values: - &#x60;visual-text&#x60;:  Returns visual embeddings optimized for text search. - &#x60;audio&#x60;: Returns audio embeddings. &lt;br/&gt; To retrieve embeddings for a video, it must be indexed using the Marengo video understanding model version 2.7 or later. For details on enabling this model for an index, see the [Create an index](/reference/create-index) page.  The platform does not return embeddings if you don&#39;t provide this parameter.  The values you specify in &#x60;embedding_option&#x60; must be included in the &#x60;model_options&#x60; defined when the index was created. For example, if &#x60;model_options&#x60; is set to &#x60;visual,&#x60; you cannot set &#x60;embedding_option&#x60; to &#x60;audio&#x60; or  both &#x60;visual-text&#x60; and &#x60;audio&#x60;.  | 
 **transcription** | **bool** | The parameter indicates whether to retrieve a transcription of the spoken words for the indexed video. Note that the official SDKs will support this feature in a future release.  | 

### Return type

[**InlineObject4**](InlineObject4.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateVideoInformation

> UpdateVideoInformation(ctx, indexId, videoId).XApiKey(xApiKey).ContentType(contentType).UpdateVideoInformationRequest(updateVideoInformationRequest).Execute()

Update video information



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
	indexId := "6298d673f1090f1100476d4c" // string | The unique identifier of the index to which the video has been uploaded. 
	videoId := "6298d673f1090f1100476d4c" // string | The unique identifier of the video to update. 
	updateVideoInformationRequest := *openapiclient.NewUpdateVideoInformationRequest() // UpdateVideoInformationRequest | Request to update the metadata of a video.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ManageVideosAPI.UpdateVideoInformation(context.Background(), indexId, videoId).XApiKey(xApiKey).ContentType(contentType).UpdateVideoInformationRequest(updateVideoInformationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManageVideosAPI.UpdateVideoInformation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**indexId** | **string** | The unique identifier of the index to which the video has been uploaded.  | 
**videoId** | **string** | The unique identifier of the video to update.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVideoInformationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiKey** | **string** | Your API key.  &lt;Note title&#x3D;\&quot;Note\&quot;&gt; You can find your API key on the &lt;a href&#x3D;\&quot;https://playground.twelvelabs.io/dashboard/api-key\&quot; target&#x3D;\&quot;_blank\&quot;&gt;API Key&lt;/a&gt; page. &lt;/Note&gt;  | 
 **contentType** | **string** | Must be set to &#x60;application/json&#x60;. | [default to &quot;application/json&quot;]


 **updateVideoInformationRequest** | [**UpdateVideoInformationRequest**](UpdateVideoInformationRequest.md) | Request to update the metadata of a video.  | 

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

