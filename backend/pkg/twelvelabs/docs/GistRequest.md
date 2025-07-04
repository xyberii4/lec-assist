# GistRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VideoId** | **string** | The unique identifier of the video that you want to generate a gist for.  | 
**Types** | **[]string** | Specifies the type of gist. Use one of the following values:   - &#x60;title&#x60;: A title succinctly captures a video&#39;s main theme, such as \&quot;From Consumerism to Minimalism: A Journey Toward Sustainable Living,\&quot; guiding viewers to its content and themes.   - &#x60;topic&#x60;: A topic is the central theme of a video, such as \&quot;Shopping Vlog Lifestyle\&quot;, summarizing its content for efficient categorization and reference.   - &#x60;hashtag&#x60;: A hashtag, like \&quot;#BlackFriday\&quot;, represents key themes in a video, enhancing its discoverability and categorization on social media platforms.  | 

## Methods

### NewGistRequest

`func NewGistRequest(videoId string, types []string, ) *GistRequest`

NewGistRequest instantiates a new GistRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGistRequestWithDefaults

`func NewGistRequestWithDefaults() *GistRequest`

NewGistRequestWithDefaults instantiates a new GistRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVideoId

`func (o *GistRequest) GetVideoId() string`

GetVideoId returns the VideoId field if non-nil, zero value otherwise.

### GetVideoIdOk

`func (o *GistRequest) GetVideoIdOk() (*string, bool)`

GetVideoIdOk returns a tuple with the VideoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoId

`func (o *GistRequest) SetVideoId(v string)`

SetVideoId sets VideoId field to given value.


### GetTypes

`func (o *GistRequest) GetTypes() []string`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *GistRequest) GetTypesOk() (*[]string, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *GistRequest) SetTypes(v []string)`

SetTypes sets Types field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


