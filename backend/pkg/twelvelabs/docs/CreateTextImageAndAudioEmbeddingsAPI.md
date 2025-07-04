# \CreateTextImageAndAudioEmbeddingsAPI

All URIs are relative to *https://api.twelvelabs.io/v1.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTextImageAudioEmbedding**](CreateTextImageAndAudioEmbeddingsAPI.md#CreateTextImageAudioEmbedding) | **Post** /embed | Create embeddings for text, image, and audio



## CreateTextImageAudioEmbedding

> EmbeddingResponse CreateTextImageAudioEmbedding(ctx).XApiKey(xApiKey).ContentType(contentType).ModelName(modelName).Text(text).TextTruncate(textTruncate).ImageUrl(imageUrl).ImageFile(imageFile).AudioUrl(audioUrl).AudioFile(audioFile).AudioStartOffsetSec(audioStartOffsetSec).Execute()

Create embeddings for text, image, and audio



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
	text := "text_example" // string | The text for which you wish to create an embedding.  <Note title=\\\"Note\\\"> Text embeddings are limited to 77 tokens. If the text exceeds this limit, the platform truncates it according to the value of the `text_truncate` parameter described below. </Note>  **Example**: \\\"Man with a dog crossing the street\\\"  (optional)
	textTruncate := "textTruncate_example" // string | Specifies how the platform truncates text that exceeds 77 tokens to fit the maximum length allowed for an embedding. This parameter can take one of the following values: - `start`: The platform will truncate the start of the provided text. - `end`: The platform will truncate the end of the provided text. - `none`: The platform will return an error if the text is longer than the maximum token limit.  **Default**: `end`  (optional) (default to "end")
	imageUrl := "imageUrl_example" // string | The publicly accessible URL of the image for which you wish to create an embedding. This parameter is required for image embeddings if `image_file` is not provided.  (optional)
	imageFile := os.NewFile(1234, "some_file") // *os.File | The image file for which you wish to create an embedding as a local file. This parameter is required for image embeddings if `image_url` is not provided.  (optional)
	audioUrl := "audioUrl_example" // string | The publicly accessible URL of the audio file for which you wish to creae an emebdding. This parameter is required for audio embeddings if `audio_file` is not provided.  (optional)
	audioFile := os.NewFile(1234, "some_file") // *os.File | The audio file for which you wish to create an embedding as a local file. This parameter is required for audio embeddings if `audio_url` is not provided.  (optional)
	audioStartOffsetSec := float32(8.14) // float32 | Specifies the start time, in seconds, from which the platform generates the audio embeddings. This parameter allows you to skip the initial portion of the audio during processing. **Default**: `0`.  (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CreateTextImageAndAudioEmbeddingsAPI.CreateTextImageAudioEmbedding(context.Background()).XApiKey(xApiKey).ContentType(contentType).ModelName(modelName).Text(text).TextTruncate(textTruncate).ImageUrl(imageUrl).ImageFile(imageFile).AudioUrl(audioUrl).AudioFile(audioFile).AudioStartOffsetSec(audioStartOffsetSec).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreateTextImageAndAudioEmbeddingsAPI.CreateTextImageAudioEmbedding``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTextImageAudioEmbedding`: EmbeddingResponse
	fmt.Fprintf(os.Stdout, "Response from `CreateTextImageAndAudioEmbeddingsAPI.CreateTextImageAudioEmbedding`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTextImageAudioEmbeddingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiKey** | **string** | Your API key.  &lt;Note title&#x3D;\&quot;Note\&quot;&gt; You can find your API key on the &lt;a href&#x3D;\&quot;https://playground.twelvelabs.io/dashboard/api-key\&quot; target&#x3D;\&quot;_blank\&quot;&gt;API Key&lt;/a&gt; page. &lt;/Note&gt;  | 
 **contentType** | **string** | Must be set to &#x60;multipart/form-data&#x60;. | [default to &quot;multipart/form-data&quot;]
 **modelName** | **string** | The name of the model you want to use. The following models are available:   - &#x60;Marengo-retrieval-2.7&#x60;  | 
 **text** | **string** | The text for which you wish to create an embedding.  &lt;Note title&#x3D;\\\&quot;Note\\\&quot;&gt; Text embeddings are limited to 77 tokens. If the text exceeds this limit, the platform truncates it according to the value of the &#x60;text_truncate&#x60; parameter described below. &lt;/Note&gt;  **Example**: \\\&quot;Man with a dog crossing the street\\\&quot;  | 
 **textTruncate** | **string** | Specifies how the platform truncates text that exceeds 77 tokens to fit the maximum length allowed for an embedding. This parameter can take one of the following values: - &#x60;start&#x60;: The platform will truncate the start of the provided text. - &#x60;end&#x60;: The platform will truncate the end of the provided text. - &#x60;none&#x60;: The platform will return an error if the text is longer than the maximum token limit.  **Default**: &#x60;end&#x60;  | [default to &quot;end&quot;]
 **imageUrl** | **string** | The publicly accessible URL of the image for which you wish to create an embedding. This parameter is required for image embeddings if &#x60;image_file&#x60; is not provided.  | 
 **imageFile** | ***os.File** | The image file for which you wish to create an embedding as a local file. This parameter is required for image embeddings if &#x60;image_url&#x60; is not provided.  | 
 **audioUrl** | **string** | The publicly accessible URL of the audio file for which you wish to creae an emebdding. This parameter is required for audio embeddings if &#x60;audio_file&#x60; is not provided.  | 
 **audioFile** | ***os.File** | The audio file for which you wish to create an embedding as a local file. This parameter is required for audio embeddings if &#x60;audio_url&#x60; is not provided.  | 
 **audioStartOffsetSec** | **float32** | Specifies the start time, in seconds, from which the platform generates the audio embeddings. This parameter allows you to skip the initial portion of the audio during processing. **Default**: &#x60;0&#x60;.  | [default to 0]

### Return type

[**EmbeddingResponse**](EmbeddingResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

