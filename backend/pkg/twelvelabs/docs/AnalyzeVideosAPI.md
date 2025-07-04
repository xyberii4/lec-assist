# \AnalyzeVideosAPI

All URIs are relative to *https://api.twelvelabs.io/v1.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Analyze**](AnalyzeVideosAPI.md#Analyze) | **Post** /generate | Open-ended texts
[**GenerateTextRepresentation**](AnalyzeVideosAPI.md#GenerateTextRepresentation) | **Post** /analyze | Open-ended analysis
[**Gist**](AnalyzeVideosAPI.md#Gist) | **Post** /gist | Titles, topics, and hashtags
[**Summarize**](AnalyzeVideosAPI.md#Summarize) | **Post** /summarize | Summaries, chapters, or highlights



## Analyze

> Analyze200Response Analyze(ctx).XApiKey(xApiKey).ContentType(contentType).AnalyzeRequest(analyzeRequest).Execute()

Open-ended texts



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
	analyzeRequest := *openapiclient.NewAnalyzeRequest("6298d673f1090f1100476d4c", "I want to generate a description for my video with the following format- Title of the video, followed by a summary in 2-3 sentences, highlighting the main topic, key events, and concluding remarks.") // AnalyzeRequest | Request to generate a text representation of a video. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalyzeVideosAPI.Analyze(context.Background()).XApiKey(xApiKey).ContentType(contentType).AnalyzeRequest(analyzeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalyzeVideosAPI.Analyze``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Analyze`: Analyze200Response
	fmt.Fprintf(os.Stdout, "Response from `AnalyzeVideosAPI.Analyze`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAnalyzeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiKey** | **string** | Your API key.  &lt;Note title&#x3D;\&quot;Note\&quot;&gt; You can find your API key on the &lt;a href&#x3D;\&quot;https://playground.twelvelabs.io/dashboard/api-key\&quot; target&#x3D;\&quot;_blank\&quot;&gt;API Key&lt;/a&gt; page. &lt;/Note&gt;  | 
 **contentType** | **string** | Must be set to &#x60;application/json&#x60;. | [default to &quot;application/json&quot;]
 **analyzeRequest** | [**AnalyzeRequest**](AnalyzeRequest.md) | Request to generate a text representation of a video.  | 

### Return type

[**Analyze200Response**](Analyze200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenerateTextRepresentation

> Analyze200Response GenerateTextRepresentation(ctx).XApiKey(xApiKey).ContentType(contentType).AnalyzeRequest(analyzeRequest).Execute()

Open-ended analysis



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
	analyzeRequest := *openapiclient.NewAnalyzeRequest("6298d673f1090f1100476d4c", "I want to generate a description for my video with the following format- Title of the video, followed by a summary in 2-3 sentences, highlighting the main topic, key events, and concluding remarks.") // AnalyzeRequest | Request to generate a text representation of a video. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalyzeVideosAPI.GenerateTextRepresentation(context.Background()).XApiKey(xApiKey).ContentType(contentType).AnalyzeRequest(analyzeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalyzeVideosAPI.GenerateTextRepresentation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GenerateTextRepresentation`: Analyze200Response
	fmt.Fprintf(os.Stdout, "Response from `AnalyzeVideosAPI.GenerateTextRepresentation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenerateTextRepresentationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiKey** | **string** | Your API key.  &lt;Note title&#x3D;\&quot;Note\&quot;&gt; You can find your API key on the &lt;a href&#x3D;\&quot;https://playground.twelvelabs.io/dashboard/api-key\&quot; target&#x3D;\&quot;_blank\&quot;&gt;API Key&lt;/a&gt; page. &lt;/Note&gt;  | 
 **contentType** | **string** | Must be set to &#x60;application/json&#x60;. | [default to &quot;application/json&quot;]
 **analyzeRequest** | [**AnalyzeRequest**](AnalyzeRequest.md) | Request to generate a text representation of a video.  | 

### Return type

[**Analyze200Response**](Analyze200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Gist

> Gist Gist(ctx).XApiKey(xApiKey).ContentType(contentType).GistRequest(gistRequest).Execute()

Titles, topics, and hashtags



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
	gistRequest := *openapiclient.NewGistRequest("6298d673f1090f1100476d4c", []string{"title"}) // GistRequest | Request to generate a gist for a video.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalyzeVideosAPI.Gist(context.Background()).XApiKey(xApiKey).ContentType(contentType).GistRequest(gistRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalyzeVideosAPI.Gist``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Gist`: Gist
	fmt.Fprintf(os.Stdout, "Response from `AnalyzeVideosAPI.Gist`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiKey** | **string** | Your API key.  &lt;Note title&#x3D;\&quot;Note\&quot;&gt; You can find your API key on the &lt;a href&#x3D;\&quot;https://playground.twelvelabs.io/dashboard/api-key\&quot; target&#x3D;\&quot;_blank\&quot;&gt;API Key&lt;/a&gt; page. &lt;/Note&gt;  | 
 **contentType** | **string** | Must be set to &#x60;application/json&#x60;. | [default to &quot;application/json&quot;]
 **gistRequest** | [**GistRequest**](GistRequest.md) | Request to generate a gist for a video.  | 

### Return type

[**Gist**](Gist.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Summarize

> InlineObject19 Summarize(ctx).XApiKey(xApiKey).ContentType(contentType).SummarizeRequest(summarizeRequest).Execute()

Summaries, chapters, or highlights



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
	summarizeRequest := *openapiclient.NewSummarizeRequest("6298d673f1090f1100476d4c", "summary") // SummarizeRequest | Request to generate a summary of a video.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalyzeVideosAPI.Summarize(context.Background()).XApiKey(xApiKey).ContentType(contentType).SummarizeRequest(summarizeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalyzeVideosAPI.Summarize``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Summarize`: InlineObject19
	fmt.Fprintf(os.Stdout, "Response from `AnalyzeVideosAPI.Summarize`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSummarizeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiKey** | **string** | Your API key.  &lt;Note title&#x3D;\&quot;Note\&quot;&gt; You can find your API key on the &lt;a href&#x3D;\&quot;https://playground.twelvelabs.io/dashboard/api-key\&quot; target&#x3D;\&quot;_blank\&quot;&gt;API Key&lt;/a&gt; page. &lt;/Note&gt;  | 
 **contentType** | **string** | Must be set to &#x60;application/json&#x60;. | [default to &quot;application/json&quot;]
 **summarizeRequest** | [**SummarizeRequest**](SummarizeRequest.md) | Request to generate a summary of a video.  | 

### Return type

[**InlineObject19**](InlineObject19.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

