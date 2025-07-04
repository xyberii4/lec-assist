# AudioEmbeddingResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Segments** | Pointer to [**[]AudioSegment**](AudioSegment.md) | An object that contains the embedding and its start time.  | [optional] 
**ErrorMessage** | Pointer to **string** | Error message if the embedding generation failed. | [optional] 
**Metadata** | Pointer to [**BaseEmbeddingMetadata**](BaseEmbeddingMetadata.md) |  | [optional] 

## Methods

### NewAudioEmbeddingResult

`func NewAudioEmbeddingResult() *AudioEmbeddingResult`

NewAudioEmbeddingResult instantiates a new AudioEmbeddingResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAudioEmbeddingResultWithDefaults

`func NewAudioEmbeddingResultWithDefaults() *AudioEmbeddingResult`

NewAudioEmbeddingResultWithDefaults instantiates a new AudioEmbeddingResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSegments

`func (o *AudioEmbeddingResult) GetSegments() []AudioSegment`

GetSegments returns the Segments field if non-nil, zero value otherwise.

### GetSegmentsOk

`func (o *AudioEmbeddingResult) GetSegmentsOk() (*[]AudioSegment, bool)`

GetSegmentsOk returns a tuple with the Segments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegments

`func (o *AudioEmbeddingResult) SetSegments(v []AudioSegment)`

SetSegments sets Segments field to given value.

### HasSegments

`func (o *AudioEmbeddingResult) HasSegments() bool`

HasSegments returns a boolean if a field has been set.

### GetErrorMessage

`func (o *AudioEmbeddingResult) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *AudioEmbeddingResult) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *AudioEmbeddingResult) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *AudioEmbeddingResult) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetMetadata

`func (o *AudioEmbeddingResult) GetMetadata() BaseEmbeddingMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AudioEmbeddingResult) GetMetadataOk() (*BaseEmbeddingMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AudioEmbeddingResult) SetMetadata(v BaseEmbeddingMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *AudioEmbeddingResult) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


