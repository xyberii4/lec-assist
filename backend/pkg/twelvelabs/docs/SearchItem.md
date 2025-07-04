# SearchItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Score** | Pointer to **float32** | A quantitative value determined by the AI model representing the level of confidence that the results match your search terms.  | [optional] 
**Start** | Pointer to **float32** | The start time of the matching video clip, expressed in seconds.  | [optional] 
**End** | Pointer to **float32** | The end time of the matching video clip, expressed in seconds.  | [optional] 
**VideoId** | Pointer to **string** | A string representing the unique identifier of the video. Once the platform indexes a video, it assigns a unique identifier. Note that this is different from the identifier of the video indexing task. | [optional] 
**Confidence** | Pointer to **string** | A qualitative indicator based on the value of the &#x60;score&#x60; field. This field can take one of the following values: - &#x60;high&#x60; - &#x60;medium&#x60; - &#x60;low&#x60; - &#x60;extremely low&#x60;  | [optional] 
**ThumbnailUrl** | Pointer to **string** | If thumbnail generation has been enabled for this index, the platform returns a string representing the URL of the thumbnail. Note that the URL expires in one hour.  | [optional] 
**Transcription** | Pointer to **string** | A transcription of the spoken words that are captured in the video. Note that the official SDKs will support this feature in a future release. | [optional] 
**Id** | Pointer to **string** | A string representing the unique identifier of the video. It only appears when the &#x60;group_by&#x3D;video&#x60; parameter is used in the request. | [optional] 
**Clips** | Pointer to [**[]SearchItemClipsInner**](SearchItemClipsInner.md) | An array that contains detailed information about the clips that match your query. The platform returns this array only when the &#x60;group_by&#x60; parameter is set to &#x60;video&#x60; in the request. | [optional] 

## Methods

### NewSearchItem

`func NewSearchItem() *SearchItem`

NewSearchItem instantiates a new SearchItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchItemWithDefaults

`func NewSearchItemWithDefaults() *SearchItem`

NewSearchItemWithDefaults instantiates a new SearchItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScore

`func (o *SearchItem) GetScore() float32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *SearchItem) GetScoreOk() (*float32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *SearchItem) SetScore(v float32)`

SetScore sets Score field to given value.

### HasScore

`func (o *SearchItem) HasScore() bool`

HasScore returns a boolean if a field has been set.

### GetStart

`func (o *SearchItem) GetStart() float32`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *SearchItem) GetStartOk() (*float32, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *SearchItem) SetStart(v float32)`

SetStart sets Start field to given value.

### HasStart

`func (o *SearchItem) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetEnd

`func (o *SearchItem) GetEnd() float32`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *SearchItem) GetEndOk() (*float32, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *SearchItem) SetEnd(v float32)`

SetEnd sets End field to given value.

### HasEnd

`func (o *SearchItem) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetVideoId

`func (o *SearchItem) GetVideoId() string`

GetVideoId returns the VideoId field if non-nil, zero value otherwise.

### GetVideoIdOk

`func (o *SearchItem) GetVideoIdOk() (*string, bool)`

GetVideoIdOk returns a tuple with the VideoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoId

`func (o *SearchItem) SetVideoId(v string)`

SetVideoId sets VideoId field to given value.

### HasVideoId

`func (o *SearchItem) HasVideoId() bool`

HasVideoId returns a boolean if a field has been set.

### GetConfidence

`func (o *SearchItem) GetConfidence() string`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *SearchItem) GetConfidenceOk() (*string, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *SearchItem) SetConfidence(v string)`

SetConfidence sets Confidence field to given value.

### HasConfidence

`func (o *SearchItem) HasConfidence() bool`

HasConfidence returns a boolean if a field has been set.

### GetThumbnailUrl

`func (o *SearchItem) GetThumbnailUrl() string`

GetThumbnailUrl returns the ThumbnailUrl field if non-nil, zero value otherwise.

### GetThumbnailUrlOk

`func (o *SearchItem) GetThumbnailUrlOk() (*string, bool)`

GetThumbnailUrlOk returns a tuple with the ThumbnailUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnailUrl

`func (o *SearchItem) SetThumbnailUrl(v string)`

SetThumbnailUrl sets ThumbnailUrl field to given value.

### HasThumbnailUrl

`func (o *SearchItem) HasThumbnailUrl() bool`

HasThumbnailUrl returns a boolean if a field has been set.

### GetTranscription

`func (o *SearchItem) GetTranscription() string`

GetTranscription returns the Transcription field if non-nil, zero value otherwise.

### GetTranscriptionOk

`func (o *SearchItem) GetTranscriptionOk() (*string, bool)`

GetTranscriptionOk returns a tuple with the Transcription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscription

`func (o *SearchItem) SetTranscription(v string)`

SetTranscription sets Transcription field to given value.

### HasTranscription

`func (o *SearchItem) HasTranscription() bool`

HasTranscription returns a boolean if a field has been set.

### GetId

`func (o *SearchItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SearchItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SearchItem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SearchItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetClips

`func (o *SearchItem) GetClips() []SearchItemClipsInner`

GetClips returns the Clips field if non-nil, zero value otherwise.

### GetClipsOk

`func (o *SearchItem) GetClipsOk() (*[]SearchItemClipsInner, bool)`

GetClipsOk returns a tuple with the Clips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClips

`func (o *SearchItem) SetClips(v []SearchItemClipsInner)`

SetClips sets Clips field to given value.

### HasClips

`func (o *SearchItem) HasClips() bool`

HasClips returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


