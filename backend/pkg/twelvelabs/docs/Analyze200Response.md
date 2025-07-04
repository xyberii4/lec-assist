# Analyze200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventType** | Pointer to **string** | This field is always set to &#x60;stream_end&#x60; for this event.  | [optional] 
**Metadata** | Pointer to [**StreamEndResponseMetadata**](StreamEndResponseMetadata.md) |  | [optional] 
**Text** | Pointer to **string** | A fragment of the generated text.  | [optional] 
**Id** | Pointer to **string** | Unique identifier of the response.  | [optional] 
**Data** | Pointer to **string** | The generated text based on the prompt you provided.  | [optional] 
**Usage** | Pointer to [**TokenUsage**](TokenUsage.md) |  | [optional] 

## Methods

### NewAnalyze200Response

`func NewAnalyze200Response() *Analyze200Response`

NewAnalyze200Response instantiates a new Analyze200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalyze200ResponseWithDefaults

`func NewAnalyze200ResponseWithDefaults() *Analyze200Response`

NewAnalyze200ResponseWithDefaults instantiates a new Analyze200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventType

`func (o *Analyze200Response) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *Analyze200Response) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *Analyze200Response) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *Analyze200Response) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetMetadata

`func (o *Analyze200Response) GetMetadata() StreamEndResponseMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Analyze200Response) GetMetadataOk() (*StreamEndResponseMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Analyze200Response) SetMetadata(v StreamEndResponseMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Analyze200Response) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetText

`func (o *Analyze200Response) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *Analyze200Response) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *Analyze200Response) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *Analyze200Response) HasText() bool`

HasText returns a boolean if a field has been set.

### GetId

`func (o *Analyze200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Analyze200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Analyze200Response) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Analyze200Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetData

`func (o *Analyze200Response) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *Analyze200Response) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *Analyze200Response) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *Analyze200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetUsage

`func (o *Analyze200Response) GetUsage() TokenUsage`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *Analyze200Response) GetUsageOk() (*TokenUsage, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *Analyze200Response) SetUsage(v TokenUsage)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *Analyze200Response) HasUsage() bool`

HasUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


