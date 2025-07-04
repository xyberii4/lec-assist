# VideoEmbeddingMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InputUrl** | Pointer to **string** | The URL of the media file used to generate the embedding. Present if a URL was provided in the request.  | [optional] 
**InputFilename** | Pointer to **string** | The name of the media file used to generate the embedding. Present if a file was provided in the request.  | [optional] 
**VideoClipLength** | Pointer to **float32** | The duration for each clip in seconds, as specified by the &#x60;video_clip_length&#x60; parameter of the [&#x60;POST&#x60;](/v1.3/api-reference/video-embeddings/create) method of the &#x60;/embed/task&#x60; endpoint. Note that the platform automatically truncates video segments shorter than 2 seconds. For a 31-second video divided into 6-second segments, the final 1-second segment will be truncated. This truncation only applies to the last segment if it does not meet the minimum length requirement of 2 seconds. | [optional] 
**VideoEmbeddingScope** | Pointer to **[]string** | The scope you&#39;ve specified in the request. It can take one of the following values: [&#39;clip&#39;] or [&#39;clip&#39;, &#39;video&#39;].  | [optional] 
**Duration** | Pointer to **float32** | The total duration of the video in seconds. | [optional] 

## Methods

### NewVideoEmbeddingMetadata

`func NewVideoEmbeddingMetadata() *VideoEmbeddingMetadata`

NewVideoEmbeddingMetadata instantiates a new VideoEmbeddingMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVideoEmbeddingMetadataWithDefaults

`func NewVideoEmbeddingMetadataWithDefaults() *VideoEmbeddingMetadata`

NewVideoEmbeddingMetadataWithDefaults instantiates a new VideoEmbeddingMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInputUrl

`func (o *VideoEmbeddingMetadata) GetInputUrl() string`

GetInputUrl returns the InputUrl field if non-nil, zero value otherwise.

### GetInputUrlOk

`func (o *VideoEmbeddingMetadata) GetInputUrlOk() (*string, bool)`

GetInputUrlOk returns a tuple with the InputUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputUrl

`func (o *VideoEmbeddingMetadata) SetInputUrl(v string)`

SetInputUrl sets InputUrl field to given value.

### HasInputUrl

`func (o *VideoEmbeddingMetadata) HasInputUrl() bool`

HasInputUrl returns a boolean if a field has been set.

### GetInputFilename

`func (o *VideoEmbeddingMetadata) GetInputFilename() string`

GetInputFilename returns the InputFilename field if non-nil, zero value otherwise.

### GetInputFilenameOk

`func (o *VideoEmbeddingMetadata) GetInputFilenameOk() (*string, bool)`

GetInputFilenameOk returns a tuple with the InputFilename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputFilename

`func (o *VideoEmbeddingMetadata) SetInputFilename(v string)`

SetInputFilename sets InputFilename field to given value.

### HasInputFilename

`func (o *VideoEmbeddingMetadata) HasInputFilename() bool`

HasInputFilename returns a boolean if a field has been set.

### GetVideoClipLength

`func (o *VideoEmbeddingMetadata) GetVideoClipLength() float32`

GetVideoClipLength returns the VideoClipLength field if non-nil, zero value otherwise.

### GetVideoClipLengthOk

`func (o *VideoEmbeddingMetadata) GetVideoClipLengthOk() (*float32, bool)`

GetVideoClipLengthOk returns a tuple with the VideoClipLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoClipLength

`func (o *VideoEmbeddingMetadata) SetVideoClipLength(v float32)`

SetVideoClipLength sets VideoClipLength field to given value.

### HasVideoClipLength

`func (o *VideoEmbeddingMetadata) HasVideoClipLength() bool`

HasVideoClipLength returns a boolean if a field has been set.

### GetVideoEmbeddingScope

`func (o *VideoEmbeddingMetadata) GetVideoEmbeddingScope() []string`

GetVideoEmbeddingScope returns the VideoEmbeddingScope field if non-nil, zero value otherwise.

### GetVideoEmbeddingScopeOk

`func (o *VideoEmbeddingMetadata) GetVideoEmbeddingScopeOk() (*[]string, bool)`

GetVideoEmbeddingScopeOk returns a tuple with the VideoEmbeddingScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoEmbeddingScope

`func (o *VideoEmbeddingMetadata) SetVideoEmbeddingScope(v []string)`

SetVideoEmbeddingScope sets VideoEmbeddingScope field to given value.

### HasVideoEmbeddingScope

`func (o *VideoEmbeddingMetadata) HasVideoEmbeddingScope() bool`

HasVideoEmbeddingScope returns a boolean if a field has been set.

### GetDuration

`func (o *VideoEmbeddingMetadata) GetDuration() float32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *VideoEmbeddingMetadata) GetDurationOk() (*float32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *VideoEmbeddingMetadata) SetDuration(v float32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *VideoEmbeddingMetadata) HasDuration() bool`

HasDuration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


