# StreamGenerateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventType** | Pointer to **string** | This field is always set to &#x60;stream_end&#x60; for this event.  | [optional] 
**Metadata** | Pointer to [**StreamEndResponseMetadata**](StreamEndResponseMetadata.md) |  | [optional] 
**Text** | Pointer to **string** | A fragment of the generated text.  | [optional] 

## Methods

### NewStreamGenerateResponse

`func NewStreamGenerateResponse() *StreamGenerateResponse`

NewStreamGenerateResponse instantiates a new StreamGenerateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamGenerateResponseWithDefaults

`func NewStreamGenerateResponseWithDefaults() *StreamGenerateResponse`

NewStreamGenerateResponseWithDefaults instantiates a new StreamGenerateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventType

`func (o *StreamGenerateResponse) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *StreamGenerateResponse) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *StreamGenerateResponse) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *StreamGenerateResponse) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetMetadata

`func (o *StreamGenerateResponse) GetMetadata() StreamEndResponseMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *StreamGenerateResponse) GetMetadataOk() (*StreamEndResponseMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *StreamGenerateResponse) SetMetadata(v StreamEndResponseMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *StreamGenerateResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetText

`func (o *StreamGenerateResponse) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *StreamGenerateResponse) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *StreamGenerateResponse) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *StreamGenerateResponse) HasText() bool`

HasText returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


