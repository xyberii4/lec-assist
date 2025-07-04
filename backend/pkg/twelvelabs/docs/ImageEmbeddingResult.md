# ImageEmbeddingResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorMessage** | Pointer to **string** | Error message if the embedding generation failed. | [optional] 
**Segments** | Pointer to [**[]BaseSegment**](BaseSegment.md) | An object that contains the embedding.  | [optional] 
**Metadata** | Pointer to [**BaseEmbeddingMetadata**](BaseEmbeddingMetadata.md) |  | [optional] 

## Methods

### NewImageEmbeddingResult

`func NewImageEmbeddingResult() *ImageEmbeddingResult`

NewImageEmbeddingResult instantiates a new ImageEmbeddingResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageEmbeddingResultWithDefaults

`func NewImageEmbeddingResultWithDefaults() *ImageEmbeddingResult`

NewImageEmbeddingResultWithDefaults instantiates a new ImageEmbeddingResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorMessage

`func (o *ImageEmbeddingResult) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *ImageEmbeddingResult) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *ImageEmbeddingResult) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *ImageEmbeddingResult) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetSegments

`func (o *ImageEmbeddingResult) GetSegments() []BaseSegment`

GetSegments returns the Segments field if non-nil, zero value otherwise.

### GetSegmentsOk

`func (o *ImageEmbeddingResult) GetSegmentsOk() (*[]BaseSegment, bool)`

GetSegmentsOk returns a tuple with the Segments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegments

`func (o *ImageEmbeddingResult) SetSegments(v []BaseSegment)`

SetSegments sets Segments field to given value.

### HasSegments

`func (o *ImageEmbeddingResult) HasSegments() bool`

HasSegments returns a boolean if a field has been set.

### GetMetadata

`func (o *ImageEmbeddingResult) GetMetadata() BaseEmbeddingMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ImageEmbeddingResult) GetMetadataOk() (*BaseEmbeddingMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ImageEmbeddingResult) SetMetadata(v BaseEmbeddingMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ImageEmbeddingResult) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


