# BaseSegment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Float** | Pointer to **[]float32** | An array of floating point numbers representing the embedding. You can use this array with cosine similarity for various downstream tasks. Note that the example response was truncated for brevity. | [optional] 

## Methods

### NewBaseSegment

`func NewBaseSegment() *BaseSegment`

NewBaseSegment instantiates a new BaseSegment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseSegmentWithDefaults

`func NewBaseSegmentWithDefaults() *BaseSegment`

NewBaseSegmentWithDefaults instantiates a new BaseSegment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFloat

`func (o *BaseSegment) GetFloat() []float32`

GetFloat returns the Float field if non-nil, zero value otherwise.

### GetFloatOk

`func (o *BaseSegment) GetFloatOk() (*[]float32, bool)`

GetFloatOk returns a tuple with the Float field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFloat

`func (o *BaseSegment) SetFloat(v []float32)`

SetFloat sets Float field to given value.

### HasFloat

`func (o *BaseSegment) HasFloat() bool`

HasFloat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


