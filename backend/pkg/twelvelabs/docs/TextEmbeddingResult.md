# TextEmbeddingResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorMessage** | Pointer to **string** | Error message if the embedding generation failed. | [optional] 
**Segments** | Pointer to [**[]BaseSegment**](BaseSegment.md) | An object that contains the embedding.  | [optional] 

## Methods

### NewTextEmbeddingResult

`func NewTextEmbeddingResult() *TextEmbeddingResult`

NewTextEmbeddingResult instantiates a new TextEmbeddingResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTextEmbeddingResultWithDefaults

`func NewTextEmbeddingResultWithDefaults() *TextEmbeddingResult`

NewTextEmbeddingResultWithDefaults instantiates a new TextEmbeddingResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorMessage

`func (o *TextEmbeddingResult) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *TextEmbeddingResult) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *TextEmbeddingResult) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *TextEmbeddingResult) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetSegments

`func (o *TextEmbeddingResult) GetSegments() []BaseSegment`

GetSegments returns the Segments field if non-nil, zero value otherwise.

### GetSegmentsOk

`func (o *TextEmbeddingResult) GetSegmentsOk() (*[]BaseSegment, bool)`

GetSegmentsOk returns a tuple with the Segments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegments

`func (o *TextEmbeddingResult) SetSegments(v []BaseSegment)`

SetSegments sets Segments field to given value.

### HasSegments

`func (o *TextEmbeddingResult) HasSegments() bool`

HasSegments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


