# StreamStartResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventType** | Pointer to **string** | This field is always set to &#x60;stream_start&#x60; for this event.  | [optional] 
**Metadata** | Pointer to [**StreamStartResponseMetadata**](StreamStartResponseMetadata.md) |  | [optional] 

## Methods

### NewStreamStartResponse

`func NewStreamStartResponse() *StreamStartResponse`

NewStreamStartResponse instantiates a new StreamStartResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamStartResponseWithDefaults

`func NewStreamStartResponseWithDefaults() *StreamStartResponse`

NewStreamStartResponseWithDefaults instantiates a new StreamStartResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventType

`func (o *StreamStartResponse) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *StreamStartResponse) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *StreamStartResponse) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *StreamStartResponse) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetMetadata

`func (o *StreamStartResponse) GetMetadata() StreamStartResponseMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *StreamStartResponse) GetMetadataOk() (*StreamStartResponseMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *StreamStartResponse) SetMetadata(v StreamStartResponseMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *StreamStartResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


