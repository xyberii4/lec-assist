# StreamEndResponseMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GenerationId** | Pointer to **string** | The same unique identifier provided in the &#x60;stream_start&#x60; event.  | [optional] 
**Usage** | Pointer to [**TokenUsage**](TokenUsage.md) |  | [optional] 

## Methods

### NewStreamEndResponseMetadata

`func NewStreamEndResponseMetadata() *StreamEndResponseMetadata`

NewStreamEndResponseMetadata instantiates a new StreamEndResponseMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamEndResponseMetadataWithDefaults

`func NewStreamEndResponseMetadataWithDefaults() *StreamEndResponseMetadata`

NewStreamEndResponseMetadataWithDefaults instantiates a new StreamEndResponseMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGenerationId

`func (o *StreamEndResponseMetadata) GetGenerationId() string`

GetGenerationId returns the GenerationId field if non-nil, zero value otherwise.

### GetGenerationIdOk

`func (o *StreamEndResponseMetadata) GetGenerationIdOk() (*string, bool)`

GetGenerationIdOk returns a tuple with the GenerationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerationId

`func (o *StreamEndResponseMetadata) SetGenerationId(v string)`

SetGenerationId sets GenerationId field to given value.

### HasGenerationId

`func (o *StreamEndResponseMetadata) HasGenerationId() bool`

HasGenerationId returns a boolean if a field has been set.

### GetUsage

`func (o *StreamEndResponseMetadata) GetUsage() TokenUsage`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *StreamEndResponseMetadata) GetUsageOk() (*TokenUsage, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *StreamEndResponseMetadata) SetUsage(v TokenUsage)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *StreamEndResponseMetadata) HasUsage() bool`

HasUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


