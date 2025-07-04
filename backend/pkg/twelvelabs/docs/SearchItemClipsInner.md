# SearchItemClipsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Score** | Pointer to **float32** | A quantitative value determined by the AI model representing the level of confidence that the results match your search terms.  | [optional] 
**Start** | Pointer to **float32** | The start time of the matching video clip, expressed in seconds.  | [optional] 
**End** | Pointer to **float32** | The end time of the matching video clip, expressed in seconds.  | [optional] 
**Confidence** | Pointer to **string** | A qualitative indicator based on the value of the &#x60;score&#x60; field. This field can take one of the following values: - &#x60;high&#x60; - &#x60;medium&#x60; - &#x60;low&#x60; - &#x60;extremely low&#x60;  | [optional] 
**ThumbnailUrl** | Pointer to **string** | If thumbnail generation has been enabled for this index, the platform returns a string representing the URL of the thumbnail. Note that the URL expires in one hour.  | [optional] 
**Transcription** | Pointer to **string** | A transcription of the spoken words that are captured in the clip. Note that the official SDKs will support this feature in a future release. | [optional] 
**VideoId** | Pointer to **string** | A string representing the unique identifier of the video for the corresponding clip. | [optional] 
**UserMetadata** | Pointer to **map[string]interface{}** | Metadata that helps you categorize your videos. You can specify a list of keys and values. Keys must be of type &#x60;string&#x60;, and values can be of the following types: &#x60;string&#x60;, &#x60;integer&#x60;, &#x60;float&#x60; or &#x60;boolean&#x60;.  **Example**: &#x60;&#x60;&#x60;JSON \&quot;user_metadata\&quot;: {   \&quot;category\&quot;: \&quot;recentlyAdded\&quot;,   \&quot;batchNumber\&quot;: 5,   \&quot;rating\&quot;: 9.3,   \&quot;needsReview\&quot;: true } &#x60;&#x60;&#x60;  &lt;Note title&#x3D;\&quot;Notes\&quot;&gt; -  If you want to store other types of data such as objects or arrays,  you must convert your data into string values. - You cannot override the following system-generated metadata fields:   - &#x60;duration&#x60;   - &#x60;filename&#x60;   - &#x60;fps&#x60;   - &#x60;height&#x60;   - &#x60;model_names&#x60;   - &#x60;size&#x60;   - &#x60;video_title&#x60;   - &#x60;width&#x60; &lt;/Note&gt;  | [optional] 

## Methods

### NewSearchItemClipsInner

`func NewSearchItemClipsInner() *SearchItemClipsInner`

NewSearchItemClipsInner instantiates a new SearchItemClipsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchItemClipsInnerWithDefaults

`func NewSearchItemClipsInnerWithDefaults() *SearchItemClipsInner`

NewSearchItemClipsInnerWithDefaults instantiates a new SearchItemClipsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScore

`func (o *SearchItemClipsInner) GetScore() float32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *SearchItemClipsInner) GetScoreOk() (*float32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *SearchItemClipsInner) SetScore(v float32)`

SetScore sets Score field to given value.

### HasScore

`func (o *SearchItemClipsInner) HasScore() bool`

HasScore returns a boolean if a field has been set.

### GetStart

`func (o *SearchItemClipsInner) GetStart() float32`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *SearchItemClipsInner) GetStartOk() (*float32, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *SearchItemClipsInner) SetStart(v float32)`

SetStart sets Start field to given value.

### HasStart

`func (o *SearchItemClipsInner) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetEnd

`func (o *SearchItemClipsInner) GetEnd() float32`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *SearchItemClipsInner) GetEndOk() (*float32, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *SearchItemClipsInner) SetEnd(v float32)`

SetEnd sets End field to given value.

### HasEnd

`func (o *SearchItemClipsInner) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetConfidence

`func (o *SearchItemClipsInner) GetConfidence() string`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *SearchItemClipsInner) GetConfidenceOk() (*string, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *SearchItemClipsInner) SetConfidence(v string)`

SetConfidence sets Confidence field to given value.

### HasConfidence

`func (o *SearchItemClipsInner) HasConfidence() bool`

HasConfidence returns a boolean if a field has been set.

### GetThumbnailUrl

`func (o *SearchItemClipsInner) GetThumbnailUrl() string`

GetThumbnailUrl returns the ThumbnailUrl field if non-nil, zero value otherwise.

### GetThumbnailUrlOk

`func (o *SearchItemClipsInner) GetThumbnailUrlOk() (*string, bool)`

GetThumbnailUrlOk returns a tuple with the ThumbnailUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnailUrl

`func (o *SearchItemClipsInner) SetThumbnailUrl(v string)`

SetThumbnailUrl sets ThumbnailUrl field to given value.

### HasThumbnailUrl

`func (o *SearchItemClipsInner) HasThumbnailUrl() bool`

HasThumbnailUrl returns a boolean if a field has been set.

### GetTranscription

`func (o *SearchItemClipsInner) GetTranscription() string`

GetTranscription returns the Transcription field if non-nil, zero value otherwise.

### GetTranscriptionOk

`func (o *SearchItemClipsInner) GetTranscriptionOk() (*string, bool)`

GetTranscriptionOk returns a tuple with the Transcription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscription

`func (o *SearchItemClipsInner) SetTranscription(v string)`

SetTranscription sets Transcription field to given value.

### HasTranscription

`func (o *SearchItemClipsInner) HasTranscription() bool`

HasTranscription returns a boolean if a field has been set.

### GetVideoId

`func (o *SearchItemClipsInner) GetVideoId() string`

GetVideoId returns the VideoId field if non-nil, zero value otherwise.

### GetVideoIdOk

`func (o *SearchItemClipsInner) GetVideoIdOk() (*string, bool)`

GetVideoIdOk returns a tuple with the VideoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoId

`func (o *SearchItemClipsInner) SetVideoId(v string)`

SetVideoId sets VideoId field to given value.

### HasVideoId

`func (o *SearchItemClipsInner) HasVideoId() bool`

HasVideoId returns a boolean if a field has been set.

### GetUserMetadata

`func (o *SearchItemClipsInner) GetUserMetadata() map[string]interface{}`

GetUserMetadata returns the UserMetadata field if non-nil, zero value otherwise.

### GetUserMetadataOk

`func (o *SearchItemClipsInner) GetUserMetadataOk() (*map[string]interface{}, bool)`

GetUserMetadataOk returns a tuple with the UserMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserMetadata

`func (o *SearchItemClipsInner) SetUserMetadata(v map[string]interface{})`

SetUserMetadata sets UserMetadata field to given value.

### HasUserMetadata

`func (o *SearchItemClipsInner) HasUserMetadata() bool`

HasUserMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


