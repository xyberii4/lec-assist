/*
TwelveLabs Video Understanding API

Use the TwelveLabs Video Understanding API to extract information from your videos and make it available to your applications. The API is organized around REST and returns responses in the JSON format. It is compatible with most programming languages, and you can also use Postman or other REST clients to send requests and view responses. 

API version: 1.3.0
Contact: support@twelvelabs.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package twelvelabs

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
	"os"
)


// AnyToVideoSearchAPIService AnyToVideoSearchAPI service
type AnyToVideoSearchAPIService service

type ApiAnyToVideoRetrieveSpecificPageRequest struct {
	ctx context.Context
	ApiService *AnyToVideoSearchAPIService
	xApiKey *string
	contentType *string
	pageToken string
}

// Your API key.  &lt;Note title&#x3D;\&quot;Note\&quot;&gt; You can find your API key on the &lt;a href&#x3D;\&quot;https://playground.twelvelabs.io/dashboard/api-key\&quot; target&#x3D;\&quot;_blank\&quot;&gt;API Key&lt;/a&gt; page. &lt;/Note&gt; 
func (r ApiAnyToVideoRetrieveSpecificPageRequest) XApiKey(xApiKey string) ApiAnyToVideoRetrieveSpecificPageRequest {
	r.xApiKey = &xApiKey
	return r
}

// Must be set to &#x60;application/json&#x60;.
func (r ApiAnyToVideoRetrieveSpecificPageRequest) ContentType(contentType string) ApiAnyToVideoRetrieveSpecificPageRequest {
	r.contentType = &contentType
	return r
}

func (r ApiAnyToVideoRetrieveSpecificPageRequest) Execute() (*InlineObject1, *http.Response, error) {
	return r.ApiService.AnyToVideoRetrieveSpecificPageExecute(r)
}

/*
AnyToVideoRetrieveSpecificPage Retrieve a specific page of search results

Use this endpoint to retrieve a specific page of search results.

<Note title="Note">
When you use pagination, you will not be charged for retrieving subsequent pages of results.
</Note>


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pageToken A token that identifies the page to retrieve. 
 @return ApiAnyToVideoRetrieveSpecificPageRequest
*/
func (a *AnyToVideoSearchAPIService) AnyToVideoRetrieveSpecificPage(ctx context.Context, pageToken string) ApiAnyToVideoRetrieveSpecificPageRequest {
	return ApiAnyToVideoRetrieveSpecificPageRequest{
		ApiService: a,
		ctx: ctx,
		pageToken: pageToken,
	}
}

// Execute executes the request
//  @return InlineObject1
func (a *AnyToVideoSearchAPIService) AnyToVideoRetrieveSpecificPageExecute(r ApiAnyToVideoRetrieveSpecificPageRequest) (*InlineObject1, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InlineObject1
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AnyToVideoSearchAPIService.AnyToVideoRetrieveSpecificPage")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/search/{page-token}"
	localVarPath = strings.Replace(localVarPath, "{"+"page-token"+"}", url.PathEscape(parameterValueToString(r.pageToken, "pageToken")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.xApiKey == nil {
		return localVarReturnValue, nil, reportError("xApiKey is required and must be specified")
	}
	if r.contentType == nil {
		return localVarReturnValue, nil, reportError("contentType is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "x-api-key", r.xApiKey, "simple", "")
	parameterAddToHeaderOrQuery(localVarHeaderParams, "Content-Type", r.contentType, "simple", "")
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v InlineObject10
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiAnyToVideoSearchRequest struct {
	ctx context.Context
	ApiService *AnyToVideoSearchAPIService
	xApiKey *string
	contentType *string
	indexId *string
	searchOptions *[]string
	queryMediaType *string
	queryMediaUrl *string
	queryMediaFile *os.File
	queryText *string
	adjustConfidenceLevel *float32
	groupBy *string
	threshold *ThresholdSearch
	sortOption *string
	operator *string
	pageLimit *int32
	filter *string
}

// Your API key.  &lt;Note title&#x3D;\&quot;Note\&quot;&gt; You can find your API key on the &lt;a href&#x3D;\&quot;https://playground.twelvelabs.io/dashboard/api-key\&quot; target&#x3D;\&quot;_blank\&quot;&gt;API Key&lt;/a&gt; page. &lt;/Note&gt; 
func (r ApiAnyToVideoSearchRequest) XApiKey(xApiKey string) ApiAnyToVideoSearchRequest {
	r.xApiKey = &xApiKey
	return r
}

// Must be set to &#x60;multipart/form-data&#x60;.
func (r ApiAnyToVideoSearchRequest) ContentType(contentType string) ApiAnyToVideoSearchRequest {
	r.contentType = &contentType
	return r
}

// The unique identifier of the index to search. 
func (r ApiAnyToVideoSearchRequest) IndexId(indexId string) ApiAnyToVideoSearchRequest {
	r.indexId = &indexId
	return r
}

// Specifies the [sources of information](/v1.3/docs/concepts/modalities#search-options) the platform uses when performing a search. You must include the &#x60;search_options&#x60; parameter separately for each desired source of information.  &lt;Note title&#x3D;\\\&quot;Notes\\\&quot;&gt; - The search options you specify must be a subset of the [model options](/v1.3/docs/concepts/modalities#model-options) used when you created the index. - You can specify multiple search options in conjunction with the &#x60;operator&#x60; parameter described below to broaden or narrow your search.  Example: To search using both visual and audio cues, include this parameter twice in the request as shown below: &#x60;&#x60;&#x60;JSON --form search_options&#x3D;visual \\\\ --form search_options&#x3D;audio \\\\ &#x60;&#x60;&#x60; &lt;/Note&gt; 
func (r ApiAnyToVideoSearchRequest) SearchOptions(searchOptions []string) ApiAnyToVideoSearchRequest {
	r.searchOptions = &searchOptions
	return r
}

// The type of media you wish to use. This parameter is required for media queries. For example, to perform an image-based search, set this parameter to &#x60;image&#x60;. 
func (r ApiAnyToVideoSearchRequest) QueryMediaType(queryMediaType string) ApiAnyToVideoSearchRequest {
	r.queryMediaType = &queryMediaType
	return r
}

// The publicly accessible URL of the media file you wish to use. This parameter is required for media queries if &#x60;query_media_file&#x60; is not provided. 
func (r ApiAnyToVideoSearchRequest) QueryMediaUrl(queryMediaUrl string) ApiAnyToVideoSearchRequest {
	r.queryMediaUrl = &queryMediaUrl
	return r
}

// The media you wish to use as a local file. This parameter is required for media queries if &#x60;query_media_url&#x60; is not provided. 
func (r ApiAnyToVideoSearchRequest) QueryMediaFile(queryMediaFile *os.File) ApiAnyToVideoSearchRequest {
	r.queryMediaFile = queryMediaFile
	return r
}

// The text query to search for. This parameter is required for text queries. Note that the platform supports full natural language-based search. 
func (r ApiAnyToVideoSearchRequest) QueryText(queryText string) ApiAnyToVideoSearchRequest {
	r.queryText = &queryText
	return r
}

// This parameter specifies the strictness of the thresholds for assigning the high, medium, or low confidence levels to search results. If you use a lower value, the thresholds become more relaxed, and more search results will be classified as having high, medium, or low confidence levels. You can use this parameter to include a broader range of potentially relevant video clips, even if some results might be less precise.  **Min**: 0 **Max**: 1 **Default:** 0.5 
func (r ApiAnyToVideoSearchRequest) AdjustConfidenceLevel(adjustConfidenceLevel float32) ApiAnyToVideoSearchRequest {
	r.adjustConfidenceLevel = &adjustConfidenceLevel
	return r
}

// Use this parameter to group or ungroup items in a response. It can take one of the following values: - &#x60;video&#x60;:  The platform will group the matching video clips in the response by video. - &#x60;clip&#x60;: The matching video clips in the response will not be grouped.  **Default:** &#x60;clip&#x60; 
func (r ApiAnyToVideoSearchRequest) GroupBy(groupBy string) ApiAnyToVideoSearchRequest {
	r.groupBy = &groupBy
	return r
}

func (r ApiAnyToVideoSearchRequest) Threshold(threshold ThresholdSearch) ApiAnyToVideoSearchRequest {
	r.threshold = &threshold
	return r
}

// Use this parameter to specify the sort order for the response.  When performing a search, the platform determines the level of confidence that each video clip matches your search terms. By default, the search results are sorted on the level of confidence in descending order.  If you set this parameter to &#x60;score&#x60; and &#x60;group_by&#x60; is set to &#x60;video&#x60;, the platform will determine the maximum value of the &#x60;score&#x60; field for each video and sort the videos in the response by the maximum value of this field. For each video, the matching video clips will be sorted by the level of confidence.  If you set this parameter to &#x60;clip_count&#x60; and &#x60;group_by&#x60; is set to &#x60;video&#x60;, the platform will sort the videos in the response by the number of clips. For each video, the matching video clips will be sorted by the level of confidence. You can use &#x60;clip_count&#x60; only when the matching video clips are grouped by video.   **Default:** &#x60;score&#x60; 
func (r ApiAnyToVideoSearchRequest) SortOption(sortOption string) ApiAnyToVideoSearchRequest {
	r.sortOption = &sortOption
	return r
}

// When you perform a search specifying multiple [sources of information](/v1.3/docs/concepts/modalities#search-options), you can use the this parameter to broaden or narrow your search.    The following logical operators are supported:    - &#x60;or&#x60;    - &#x60;and&#x60;    For details and examples, see the [Using multiple sources of information](/v1.3/docs/guides/search/queries/text-queries#visual-and-audio) section.     **Default**: &#x60;or&#x60;. 
func (r ApiAnyToVideoSearchRequest) Operator(operator string) ApiAnyToVideoSearchRequest {
	r.operator = &operator
	return r
}

// The number of items to return on each page. When grouping by video, this parameter represents the number of videos per page. Otherwise, it represents the maximum number of video clips per page.  **Max**: &#x60;50&#x60;. 
func (r ApiAnyToVideoSearchRequest) PageLimit(pageLimit int32) ApiAnyToVideoSearchRequest {
	r.pageLimit = &pageLimit
	return r
}

// Specifies a stringified JSON object to filter your search results. Supports both system-generated metadata (example: video ID, duration) and user-defined metadata.  **Syntax for filtering**  The following table describes the supported data types, operators, and filter syntax:   | Data type | Operator | Description | Syntax | |:----------|:---------|:------------|:-------| | String | &#x60;&#x3D;&#x60; | Matches results equal to the specified value. | &#x60;{\\\&quot;field\\\&quot;: \\\&quot;value\\\&quot;}&#x60;  | Array of strings | &#x60;&#x3D;&#x60; | Matches results with any value in the specified array. Supported only for &#x60;id&#x60;. | &#x60;{\\\&quot;id\\\&quot;: [\\\&quot;value1\\\&quot;, \\\&quot;value2\\\&quot;]}&#x60; | | Numeric (integer, float) | &#x60;&#x3D;&#x60;, &#x60;lte&#x60;, &#x60;gte&#x60; | Matches results equal to or within a range of the specified value. | &#x60;{\\\&quot;field\\\&quot;: number}&#x60; or &#x60;{\\\&quot;field\\\&quot;: { \\\&quot;gte\\\&quot;: number, \\\&quot;lte\\\&quot;: number }}&#x60; | | Boolean | &#x60;&#x3D;&#x60; | Matches results equal to the specified boolean value. | &#x60;{\\\&quot;field\\\&quot;: true}&#x60; or &#x60;{\\\&quot;field\\\&quot;: false}&#x60;. |  &lt;br/&gt; **System-generated metadata**  The table below describes the system-generated metadata available for filtering your  search results:  | Field name | Description | Type | Example | |:-----------|:------------|:-----|:--------| | &#x60;id&#x60; | Filters by specific video IDs. | Array of strings | &#x60;{\\\&quot;id\\\&quot;: [\\\&quot;67cec9caf45d9b64a58340fc\\\&quot;, \\\&quot;67cec9baf45d9b64a58340fa\\\&quot;]}&#x60;. | | &#x60;duration&#x60; | Filters based on the duration of the video containing the segment that matches your query. | Number or object with &#x60;gte&#x60; and &#x60;lte&#x60; | &#x60;{\\\&quot;duration\\\&quot;: 600}&#x60; or &#x60;{\\\&quot;duration\\\&quot;: { \\\&quot;gte\\\&quot;: 600, \\\&quot;lte\\\&quot;: 800 }}&#x60; | | &#x60;width&#x60; | Filters by video width (in pixels). | Number or object with &#x60;gte&#x60; and &#x60;lte&#x60; | &#x60;{\\\&quot;width\\\&quot;: 1920}&#x60; or &#x60;{\\\&quot;width\\\&quot;: { \\\&quot;gte\\\&quot;: 1280, \\\&quot;lte\\\&quot;: 1920}}&#x60; | | &#x60;height&#x60; | Filters by video height (in pixels). | Number or object with &#x60;gte&#x60; and &#x60;lte&#x60;. | &#x60;{\\\&quot;height\\\&quot;: 1080}&#x60; or &#x60;{\\\&quot;height\\\&quot;: { \\\&quot;gte\\\&quot;: 720, \\\&quot;lte\\\&quot;: 1080 }}&#x60;. | | &#x60;size&#x60; | Filters by video size (in bytes) | Number or object with &#x60;gte&#x60; and &#x60;lte&#x60;. | &#x60;{\\\&quot;size\\\&quot;: 1048576}&#x60; or &#x60;{\\\&quot;size\\\&quot;: { \\\&quot;gte\\\&quot;: 1048576, \\\&quot;lte\\\&quot;: 5242880}}&#x60; |  | &#x60;filename&#x60; | Filters by the exact file name. | String | &#x60;{\\\&quot;filename\\\&quot;: \\\&quot;Animal Encounters part 1\\\&quot;}&#x60; |   &lt;br/&gt; **User-defined metadata**  To filter by user-defined metadata: 1. Add metadata to your video by calling the [&#x60;PUT&#x60;](/v1.3/api-reference/videos/update) method of the &#x60;/indexes/:index-id/videos/:video-id&#x60; endpoint 2. Reference the custom field in your filter object. For example, to filter videos where a custom field named &#x60;needsReview&#x60; of type boolean is &#x60;true&#x60;, use &#x60;{\\\&quot;needs_review\\\&quot;: true}&#x60;.  For more details and examples, see the [Filter search results](/v1.3/docs/guides/search/filtering) page. 
func (r ApiAnyToVideoSearchRequest) Filter(filter string) ApiAnyToVideoSearchRequest {
	r.filter = &filter
	return r
}

func (r ApiAnyToVideoSearchRequest) Execute() (*SearchResults, *http.Response, error) {
	return r.ApiService.AnyToVideoSearchExecute(r)
}

/*
AnyToVideoSearch Make any-to-video search requests

Use this endpoint to search for relevant matches in an index using text or various media queries.

**Text queries**:
- Use the `query_text` parameter to specify your query.

**Media queries**:
- Set the `query_media_type` parameter to the corresponding media type (example: `image`).
- Specify either one of the following parameters:
  - `query_media_url`: Publicly accessible URL of your media file.
  - `query_media_file`: Local media file.
  If both `query_media_url` and `query_media_file` are specified in the same request, `query_media_url` takes precedence.
<Accordion title="Image requirements">
Your images must meet the following requirements:
  - **Format**: JPEG and PNG.
  - **Dimension**: Must be at least 64 x 64 pixels.
  - **Size**: Must not exceed 5MB.
  - **Object visibility**: Ensure that the objects of interest are visible and occupy at least 50% of the video frame. This helps the platform accurately identify and match the objects.
</Accordion>

<Note title="Note">
This endpoint is rate-limited. For details, see the [Rate limits](/v1.3/docs/get-started/rate-limits) page.
</Note>


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAnyToVideoSearchRequest
*/
func (a *AnyToVideoSearchAPIService) AnyToVideoSearch(ctx context.Context) ApiAnyToVideoSearchRequest {
	return ApiAnyToVideoSearchRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return SearchResults
func (a *AnyToVideoSearchAPIService) AnyToVideoSearchExecute(r ApiAnyToVideoSearchRequest) (*SearchResults, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SearchResults
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AnyToVideoSearchAPIService.AnyToVideoSearch")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/search"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.xApiKey == nil {
		return localVarReturnValue, nil, reportError("xApiKey is required and must be specified")
	}
	if r.contentType == nil {
		return localVarReturnValue, nil, reportError("contentType is required and must be specified")
	}
	if r.indexId == nil {
		return localVarReturnValue, nil, reportError("indexId is required and must be specified")
	}
	if r.searchOptions == nil {
		return localVarReturnValue, nil, reportError("searchOptions is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "x-api-key", r.xApiKey, "simple", "")
	parameterAddToHeaderOrQuery(localVarHeaderParams, "Content-Type", r.contentType, "simple", "")
	if r.queryMediaType != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "query_media_type", r.queryMediaType, "", "")
	}
	if r.queryMediaUrl != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "query_media_url", r.queryMediaUrl, "", "")
	}
	var queryMediaFileLocalVarFormFileName string
	var queryMediaFileLocalVarFileName     string
	var queryMediaFileLocalVarFileBytes    []byte

	queryMediaFileLocalVarFormFileName = "query_media_file"
	queryMediaFileLocalVarFile := r.queryMediaFile

	if queryMediaFileLocalVarFile != nil {
		fbs, _ := io.ReadAll(queryMediaFileLocalVarFile)

		queryMediaFileLocalVarFileBytes = fbs
		queryMediaFileLocalVarFileName = queryMediaFileLocalVarFile.Name()
		queryMediaFileLocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: queryMediaFileLocalVarFileBytes, fileName: queryMediaFileLocalVarFileName, formFileName: queryMediaFileLocalVarFormFileName})
	}
	if r.queryText != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "query_text", r.queryText, "", "")
	}
	parameterAddToHeaderOrQuery(localVarFormParams, "index_id", r.indexId, "", "")
	parameterAddToHeaderOrQuery(localVarFormParams, "search_options", r.searchOptions, "form", "multi")
	if r.adjustConfidenceLevel != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "adjust_confidence_level", r.adjustConfidenceLevel, "", "")
	}
	if r.groupBy != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "group_by", r.groupBy, "", "")
	}
	if r.threshold != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "threshold", r.threshold, "", "")
	}
	if r.sortOption != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "sort_option", r.sortOption, "", "")
	}
	if r.operator != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "operator", r.operator, "", "")
	}
	if r.pageLimit != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "page_limit", r.pageLimit, "", "")
	}
	if r.filter != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "filter", r.filter, "", "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v InlineObject13
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
