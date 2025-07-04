# StreamTextResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventType** | Pointer to **string** | This field is always set to &#x60;text_generation&#x60; for this event.  | [optional] 
**Text** | Pointer to **string** | A fragment of the generated text.  | [optional] 

## Methods

### NewStreamTextResponse

`func NewStreamTextResponse() *StreamTextResponse`

NewStreamTextResponse instantiates a new StreamTextResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamTextResponseWithDefaults

`func NewStreamTextResponseWithDefaults() *StreamTextResponse`

NewStreamTextResponseWithDefaults instantiates a new StreamTextResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventType

`func (o *StreamTextResponse) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *StreamTextResponse) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *StreamTextResponse) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *StreamTextResponse) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetText

`func (o *StreamTextResponse) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *StreamTextResponse) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *StreamTextResponse) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *StreamTextResponse) HasText() bool`

HasText returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


