# VideoSegment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Float** | Pointer to **[]float32** | An array of floating point numbers representing the embedding. You can use this array with cosine similarity for various downstream tasks. Note that the example response was truncated for brevity. | [optional] 
**StartOffsetSec** | Pointer to **float32** | The start time, in seconds, from which the platform generated the audio embedding. | [optional] 
**EndOffsetSec** | Pointer to **float32** | The end time, in seconds, of the video segment for this embedding. | [optional] 
**EmbeddingOption** | Pointer to **string** | The type of the embedding. | [optional] 
**EmbeddingScope** | Pointer to **string** | The scope of the video embedding. | [optional] 

## Methods

### NewVideoSegment

`func NewVideoSegment() *VideoSegment`

NewVideoSegment instantiates a new VideoSegment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVideoSegmentWithDefaults

`func NewVideoSegmentWithDefaults() *VideoSegment`

NewVideoSegmentWithDefaults instantiates a new VideoSegment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFloat

`func (o *VideoSegment) GetFloat() []float32`

GetFloat returns the Float field if non-nil, zero value otherwise.

### GetFloatOk

`func (o *VideoSegment) GetFloatOk() (*[]float32, bool)`

GetFloatOk returns a tuple with the Float field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFloat

`func (o *VideoSegment) SetFloat(v []float32)`

SetFloat sets Float field to given value.

### HasFloat

`func (o *VideoSegment) HasFloat() bool`

HasFloat returns a boolean if a field has been set.

### GetStartOffsetSec

`func (o *VideoSegment) GetStartOffsetSec() float32`

GetStartOffsetSec returns the StartOffsetSec field if non-nil, zero value otherwise.

### GetStartOffsetSecOk

`func (o *VideoSegment) GetStartOffsetSecOk() (*float32, bool)`

GetStartOffsetSecOk returns a tuple with the StartOffsetSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartOffsetSec

`func (o *VideoSegment) SetStartOffsetSec(v float32)`

SetStartOffsetSec sets StartOffsetSec field to given value.

### HasStartOffsetSec

`func (o *VideoSegment) HasStartOffsetSec() bool`

HasStartOffsetSec returns a boolean if a field has been set.

### GetEndOffsetSec

`func (o *VideoSegment) GetEndOffsetSec() float32`

GetEndOffsetSec returns the EndOffsetSec field if non-nil, zero value otherwise.

### GetEndOffsetSecOk

`func (o *VideoSegment) GetEndOffsetSecOk() (*float32, bool)`

GetEndOffsetSecOk returns a tuple with the EndOffsetSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndOffsetSec

`func (o *VideoSegment) SetEndOffsetSec(v float32)`

SetEndOffsetSec sets EndOffsetSec field to given value.

### HasEndOffsetSec

`func (o *VideoSegment) HasEndOffsetSec() bool`

HasEndOffsetSec returns a boolean if a field has been set.

### GetEmbeddingOption

`func (o *VideoSegment) GetEmbeddingOption() string`

GetEmbeddingOption returns the EmbeddingOption field if non-nil, zero value otherwise.

### GetEmbeddingOptionOk

`func (o *VideoSegment) GetEmbeddingOptionOk() (*string, bool)`

GetEmbeddingOptionOk returns a tuple with the EmbeddingOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbeddingOption

`func (o *VideoSegment) SetEmbeddingOption(v string)`

SetEmbeddingOption sets EmbeddingOption field to given value.

### HasEmbeddingOption

`func (o *VideoSegment) HasEmbeddingOption() bool`

HasEmbeddingOption returns a boolean if a field has been set.

### GetEmbeddingScope

`func (o *VideoSegment) GetEmbeddingScope() string`

GetEmbeddingScope returns the EmbeddingScope field if non-nil, zero value otherwise.

### GetEmbeddingScopeOk

`func (o *VideoSegment) GetEmbeddingScopeOk() (*string, bool)`

GetEmbeddingScopeOk returns a tuple with the EmbeddingScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbeddingScope

`func (o *VideoSegment) SetEmbeddingScope(v string)`

SetEmbeddingScope sets EmbeddingScope field to given value.

### HasEmbeddingScope

`func (o *VideoSegment) HasEmbeddingScope() bool`

HasEmbeddingScope returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


