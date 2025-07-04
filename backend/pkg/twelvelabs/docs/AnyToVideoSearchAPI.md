# \AnyToVideoSearchAPI

All URIs are relative to *https://api.twelvelabs.io/v1.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AnyToVideoRetrieveSpecificPage**](AnyToVideoSearchAPI.md#AnyToVideoRetrieveSpecificPage) | **Get** /search/{page-token} | Retrieve a specific page of search results
[**AnyToVideoSearch**](AnyToVideoSearchAPI.md#AnyToVideoSearch) | **Post** /search | Make any-to-video search requests



## AnyToVideoRetrieveSpecificPage

> InlineObject1 AnyToVideoRetrieveSpecificPage(ctx, pageToken).XApiKey(xApiKey).ContentType(contentType).Execute()

Retrieve a specific page of search results



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
	pageToken := "1234567890" // string | A token that identifies the page to retrieve. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnyToVideoSearchAPI.AnyToVideoRetrieveSpecificPage(context.Background(), pageToken).XApiKey(xApiKey).ContentType(contentType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnyToVideoSearchAPI.AnyToVideoRetrieveSpecificPage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnyToVideoRetrieveSpecificPage`: InlineObject1
	fmt.Fprintf(os.Stdout, "Response from `AnyToVideoSearchAPI.AnyToVideoRetrieveSpecificPage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageToken** | **string** | A token that identifies the page to retrieve.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAnyToVideoRetrieveSpecificPageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiKey** | **string** | Your API key.  &lt;Note title&#x3D;\&quot;Note\&quot;&gt; You can find your API key on the &lt;a href&#x3D;\&quot;https://playground.twelvelabs.io/dashboard/api-key\&quot; target&#x3D;\&quot;_blank\&quot;&gt;API Key&lt;/a&gt; page. &lt;/Note&gt;  | 
 **contentType** | **string** | Must be set to &#x60;application/json&#x60;. | [default to &quot;application/json&quot;]


### Return type

[**InlineObject1**](InlineObject1.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AnyToVideoSearch

> SearchResults AnyToVideoSearch(ctx).XApiKey(xApiKey).ContentType(contentType).IndexId(indexId).SearchOptions(searchOptions).QueryMediaType(queryMediaType).QueryMediaUrl(queryMediaUrl).QueryMediaFile(queryMediaFile).QueryText(queryText).AdjustConfidenceLevel(adjustConfidenceLevel).GroupBy(groupBy).Threshold(threshold).SortOption(sortOption).Operator(operator).PageLimit(pageLimit).Filter(filter).Execute()

Make any-to-video search requests



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
	indexId := "indexId_example" // string | The unique identifier of the index to search. 
	searchOptions := []string{"visual"} // []string | Specifies the [sources of information](/v1.3/docs/concepts/modalities#search-options) the platform uses when performing a search. You must include the `search_options` parameter separately for each desired source of information.  <Note title=\\\"Notes\\\"> - The search options you specify must be a subset of the [model options](/v1.3/docs/concepts/modalities#model-options) used when you created the index. - You can specify multiple search options in conjunction with the `operator` parameter described below to broaden or narrow your search.  Example: To search using both visual and audio cues, include this parameter twice in the request as shown below: ```JSON --form search_options=visual \\\\ --form search_options=audio \\\\ ``` </Note> 
	queryMediaType := "queryMediaType_example" // string | The type of media you wish to use. This parameter is required for media queries. For example, to perform an image-based search, set this parameter to `image`.  (optional)
	queryMediaUrl := "queryMediaUrl_example" // string | The publicly accessible URL of the media file you wish to use. This parameter is required for media queries if `query_media_file` is not provided.  (optional)
	queryMediaFile := os.NewFile(1234, "some_file") // *os.File | The media you wish to use as a local file. This parameter is required for media queries if `query_media_url` is not provided.  (optional)
	queryText := "queryText_example" // string | The text query to search for. This parameter is required for text queries. Note that the platform supports full natural language-based search.  (optional)
	adjustConfidenceLevel := float32(3.4) // float32 | This parameter specifies the strictness of the thresholds for assigning the high, medium, or low confidence levels to search results. If you use a lower value, the thresholds become more relaxed, and more search results will be classified as having high, medium, or low confidence levels. You can use this parameter to include a broader range of potentially relevant video clips, even if some results might be less precise.  **Min**: 0 **Max**: 1 **Default:** 0.5  (optional) (default to 0.5)
	groupBy := "groupBy_example" // string | Use this parameter to group or ungroup items in a response. It can take one of the following values: - `video`:  The platform will group the matching video clips in the response by video. - `clip`: The matching video clips in the response will not be grouped.  **Default:** `clip`  (optional) (default to "clip")
	threshold := openapiclient.threshold_search("high") // ThresholdSearch |  (optional) (default to "low")
	sortOption := "sortOption_example" // string | Use this parameter to specify the sort order for the response.  When performing a search, the platform determines the level of confidence that each video clip matches your search terms. By default, the search results are sorted on the level of confidence in descending order.  If you set this parameter to `score` and `group_by` is set to `video`, the platform will determine the maximum value of the `score` field for each video and sort the videos in the response by the maximum value of this field. For each video, the matching video clips will be sorted by the level of confidence.  If you set this parameter to `clip_count` and `group_by` is set to `video`, the platform will sort the videos in the response by the number of clips. For each video, the matching video clips will be sorted by the level of confidence. You can use `clip_count` only when the matching video clips are grouped by video.   **Default:** `score`  (optional) (default to "score")
	operator := "operator_example" // string | When you perform a search specifying multiple [sources of information](/v1.3/docs/concepts/modalities#search-options), you can use the this parameter to broaden or narrow your search.    The following logical operators are supported:    - `or`    - `and`    For details and examples, see the [Using multiple sources of information](/v1.3/docs/guides/search/queries/text-queries#visual-and-audio) section.     **Default**: `or`.  (optional) (default to "or")
	pageLimit := int32(56) // int32 | The number of items to return on each page. When grouping by video, this parameter represents the number of videos per page. Otherwise, it represents the maximum number of video clips per page.  **Max**: `50`.  (optional) (default to 10)
	filter := "filter_example" // string | Specifies a stringified JSON object to filter your search results. Supports both system-generated metadata (example: video ID, duration) and user-defined metadata.  **Syntax for filtering**  The following table describes the supported data types, operators, and filter syntax:   | Data type | Operator | Description | Syntax | |:----------|:---------|:------------|:-------| | String | `=` | Matches results equal to the specified value. | `{\\\"field\\\": \\\"value\\\"}`  | Array of strings | `=` | Matches results with any value in the specified array. Supported only for `id`. | `{\\\"id\\\": [\\\"value1\\\", \\\"value2\\\"]}` | | Numeric (integer, float) | `=`, `lte`, `gte` | Matches results equal to or within a range of the specified value. | `{\\\"field\\\": number}` or `{\\\"field\\\": { \\\"gte\\\": number, \\\"lte\\\": number }}` | | Boolean | `=` | Matches results equal to the specified boolean value. | `{\\\"field\\\": true}` or `{\\\"field\\\": false}`. |  <br/> **System-generated metadata**  The table below describes the system-generated metadata available for filtering your  search results:  | Field name | Description | Type | Example | |:-----------|:------------|:-----|:--------| | `id` | Filters by specific video IDs. | Array of strings | `{\\\"id\\\": [\\\"67cec9caf45d9b64a58340fc\\\", \\\"67cec9baf45d9b64a58340fa\\\"]}`. | | `duration` | Filters based on the duration of the video containing the segment that matches your query. | Number or object with `gte` and `lte` | `{\\\"duration\\\": 600}` or `{\\\"duration\\\": { \\\"gte\\\": 600, \\\"lte\\\": 800 }}` | | `width` | Filters by video width (in pixels). | Number or object with `gte` and `lte` | `{\\\"width\\\": 1920}` or `{\\\"width\\\": { \\\"gte\\\": 1280, \\\"lte\\\": 1920}}` | | `height` | Filters by video height (in pixels). | Number or object with `gte` and `lte`. | `{\\\"height\\\": 1080}` or `{\\\"height\\\": { \\\"gte\\\": 720, \\\"lte\\\": 1080 }}`. | | `size` | Filters by video size (in bytes) | Number or object with `gte` and `lte`. | `{\\\"size\\\": 1048576}` or `{\\\"size\\\": { \\\"gte\\\": 1048576, \\\"lte\\\": 5242880}}` |  | `filename` | Filters by the exact file name. | String | `{\\\"filename\\\": \\\"Animal Encounters part 1\\\"}` |   <br/> **User-defined metadata**  To filter by user-defined metadata: 1. Add metadata to your video by calling the [`PUT`](/v1.3/api-reference/videos/update) method of the `/indexes/:index-id/videos/:video-id` endpoint 2. Reference the custom field in your filter object. For example, to filter videos where a custom field named `needsReview` of type boolean is `true`, use `{\\\"needs_review\\\": true}`.  For more details and examples, see the [Filter search results](/v1.3/docs/guides/search/filtering) page.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnyToVideoSearchAPI.AnyToVideoSearch(context.Background()).XApiKey(xApiKey).ContentType(contentType).IndexId(indexId).SearchOptions(searchOptions).QueryMediaType(queryMediaType).QueryMediaUrl(queryMediaUrl).QueryMediaFile(queryMediaFile).QueryText(queryText).AdjustConfidenceLevel(adjustConfidenceLevel).GroupBy(groupBy).Threshold(threshold).SortOption(sortOption).Operator(operator).PageLimit(pageLimit).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnyToVideoSearchAPI.AnyToVideoSearch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnyToVideoSearch`: SearchResults
	fmt.Fprintf(os.Stdout, "Response from `AnyToVideoSearchAPI.AnyToVideoSearch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAnyToVideoSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiKey** | **string** | Your API key.  &lt;Note title&#x3D;\&quot;Note\&quot;&gt; You can find your API key on the &lt;a href&#x3D;\&quot;https://playground.twelvelabs.io/dashboard/api-key\&quot; target&#x3D;\&quot;_blank\&quot;&gt;API Key&lt;/a&gt; page. &lt;/Note&gt;  | 
 **contentType** | **string** | Must be set to &#x60;multipart/form-data&#x60;. | [default to &quot;multipart/form-data&quot;]
 **indexId** | **string** | The unique identifier of the index to search.  | 
 **searchOptions** | **[]string** | Specifies the [sources of information](/v1.3/docs/concepts/modalities#search-options) the platform uses when performing a search. You must include the &#x60;search_options&#x60; parameter separately for each desired source of information.  &lt;Note title&#x3D;\\\&quot;Notes\\\&quot;&gt; - The search options you specify must be a subset of the [model options](/v1.3/docs/concepts/modalities#model-options) used when you created the index. - You can specify multiple search options in conjunction with the &#x60;operator&#x60; parameter described below to broaden or narrow your search.  Example: To search using both visual and audio cues, include this parameter twice in the request as shown below: &#x60;&#x60;&#x60;JSON --form search_options&#x3D;visual \\\\ --form search_options&#x3D;audio \\\\ &#x60;&#x60;&#x60; &lt;/Note&gt;  | 
 **queryMediaType** | **string** | The type of media you wish to use. This parameter is required for media queries. For example, to perform an image-based search, set this parameter to &#x60;image&#x60;.  | 
 **queryMediaUrl** | **string** | The publicly accessible URL of the media file you wish to use. This parameter is required for media queries if &#x60;query_media_file&#x60; is not provided.  | 
 **queryMediaFile** | ***os.File** | The media you wish to use as a local file. This parameter is required for media queries if &#x60;query_media_url&#x60; is not provided.  | 
 **queryText** | **string** | The text query to search for. This parameter is required for text queries. Note that the platform supports full natural language-based search.  | 
 **adjustConfidenceLevel** | **float32** | This parameter specifies the strictness of the thresholds for assigning the high, medium, or low confidence levels to search results. If you use a lower value, the thresholds become more relaxed, and more search results will be classified as having high, medium, or low confidence levels. You can use this parameter to include a broader range of potentially relevant video clips, even if some results might be less precise.  **Min**: 0 **Max**: 1 **Default:** 0.5  | [default to 0.5]
 **groupBy** | **string** | Use this parameter to group or ungroup items in a response. It can take one of the following values: - &#x60;video&#x60;:  The platform will group the matching video clips in the response by video. - &#x60;clip&#x60;: The matching video clips in the response will not be grouped.  **Default:** &#x60;clip&#x60;  | [default to &quot;clip&quot;]
 **threshold** | [**ThresholdSearch**](ThresholdSearch.md) |  | [default to &quot;low&quot;]
 **sortOption** | **string** | Use this parameter to specify the sort order for the response.  When performing a search, the platform determines the level of confidence that each video clip matches your search terms. By default, the search results are sorted on the level of confidence in descending order.  If you set this parameter to &#x60;score&#x60; and &#x60;group_by&#x60; is set to &#x60;video&#x60;, the platform will determine the maximum value of the &#x60;score&#x60; field for each video and sort the videos in the response by the maximum value of this field. For each video, the matching video clips will be sorted by the level of confidence.  If you set this parameter to &#x60;clip_count&#x60; and &#x60;group_by&#x60; is set to &#x60;video&#x60;, the platform will sort the videos in the response by the number of clips. For each video, the matching video clips will be sorted by the level of confidence. You can use &#x60;clip_count&#x60; only when the matching video clips are grouped by video.   **Default:** &#x60;score&#x60;  | [default to &quot;score&quot;]
 **operator** | **string** | When you perform a search specifying multiple [sources of information](/v1.3/docs/concepts/modalities#search-options), you can use the this parameter to broaden or narrow your search.    The following logical operators are supported:    - &#x60;or&#x60;    - &#x60;and&#x60;    For details and examples, see the [Using multiple sources of information](/v1.3/docs/guides/search/queries/text-queries#visual-and-audio) section.     **Default**: &#x60;or&#x60;.  | [default to &quot;or&quot;]
 **pageLimit** | **int32** | The number of items to return on each page. When grouping by video, this parameter represents the number of videos per page. Otherwise, it represents the maximum number of video clips per page.  **Max**: &#x60;50&#x60;.  | [default to 10]
 **filter** | **string** | Specifies a stringified JSON object to filter your search results. Supports both system-generated metadata (example: video ID, duration) and user-defined metadata.  **Syntax for filtering**  The following table describes the supported data types, operators, and filter syntax:   | Data type | Operator | Description | Syntax | |:----------|:---------|:------------|:-------| | String | &#x60;&#x3D;&#x60; | Matches results equal to the specified value. | &#x60;{\\\&quot;field\\\&quot;: \\\&quot;value\\\&quot;}&#x60;  | Array of strings | &#x60;&#x3D;&#x60; | Matches results with any value in the specified array. Supported only for &#x60;id&#x60;. | &#x60;{\\\&quot;id\\\&quot;: [\\\&quot;value1\\\&quot;, \\\&quot;value2\\\&quot;]}&#x60; | | Numeric (integer, float) | &#x60;&#x3D;&#x60;, &#x60;lte&#x60;, &#x60;gte&#x60; | Matches results equal to or within a range of the specified value. | &#x60;{\\\&quot;field\\\&quot;: number}&#x60; or &#x60;{\\\&quot;field\\\&quot;: { \\\&quot;gte\\\&quot;: number, \\\&quot;lte\\\&quot;: number }}&#x60; | | Boolean | &#x60;&#x3D;&#x60; | Matches results equal to the specified boolean value. | &#x60;{\\\&quot;field\\\&quot;: true}&#x60; or &#x60;{\\\&quot;field\\\&quot;: false}&#x60;. |  &lt;br/&gt; **System-generated metadata**  The table below describes the system-generated metadata available for filtering your  search results:  | Field name | Description | Type | Example | |:-----------|:------------|:-----|:--------| | &#x60;id&#x60; | Filters by specific video IDs. | Array of strings | &#x60;{\\\&quot;id\\\&quot;: [\\\&quot;67cec9caf45d9b64a58340fc\\\&quot;, \\\&quot;67cec9baf45d9b64a58340fa\\\&quot;]}&#x60;. | | &#x60;duration&#x60; | Filters based on the duration of the video containing the segment that matches your query. | Number or object with &#x60;gte&#x60; and &#x60;lte&#x60; | &#x60;{\\\&quot;duration\\\&quot;: 600}&#x60; or &#x60;{\\\&quot;duration\\\&quot;: { \\\&quot;gte\\\&quot;: 600, \\\&quot;lte\\\&quot;: 800 }}&#x60; | | &#x60;width&#x60; | Filters by video width (in pixels). | Number or object with &#x60;gte&#x60; and &#x60;lte&#x60; | &#x60;{\\\&quot;width\\\&quot;: 1920}&#x60; or &#x60;{\\\&quot;width\\\&quot;: { \\\&quot;gte\\\&quot;: 1280, \\\&quot;lte\\\&quot;: 1920}}&#x60; | | &#x60;height&#x60; | Filters by video height (in pixels). | Number or object with &#x60;gte&#x60; and &#x60;lte&#x60;. | &#x60;{\\\&quot;height\\\&quot;: 1080}&#x60; or &#x60;{\\\&quot;height\\\&quot;: { \\\&quot;gte\\\&quot;: 720, \\\&quot;lte\\\&quot;: 1080 }}&#x60;. | | &#x60;size&#x60; | Filters by video size (in bytes) | Number or object with &#x60;gte&#x60; and &#x60;lte&#x60;. | &#x60;{\\\&quot;size\\\&quot;: 1048576}&#x60; or &#x60;{\\\&quot;size\\\&quot;: { \\\&quot;gte\\\&quot;: 1048576, \\\&quot;lte\\\&quot;: 5242880}}&#x60; |  | &#x60;filename&#x60; | Filters by the exact file name. | String | &#x60;{\\\&quot;filename\\\&quot;: \\\&quot;Animal Encounters part 1\\\&quot;}&#x60; |   &lt;br/&gt; **User-defined metadata**  To filter by user-defined metadata: 1. Add metadata to your video by calling the [&#x60;PUT&#x60;](/v1.3/api-reference/videos/update) method of the &#x60;/indexes/:index-id/videos/:video-id&#x60; endpoint 2. Reference the custom field in your filter object. For example, to filter videos where a custom field named &#x60;needsReview&#x60; of type boolean is &#x60;true&#x60;, use &#x60;{\\\&quot;needs_review\\\&quot;: true}&#x60;.  For more details and examples, see the [Filter search results](/v1.3/docs/guides/search/filtering) page.  | 

### Return type

[**SearchResults**](SearchResults.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

