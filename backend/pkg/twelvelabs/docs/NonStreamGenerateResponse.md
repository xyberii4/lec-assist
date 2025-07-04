# NonStreamGenerateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique identifier of the response.  | [optional] 
**Data** | Pointer to **string** | The generated text based on the prompt you provided.  | [optional] 
**Usage** | Pointer to [**TokenUsage**](TokenUsage.md) |  | [optional] 

## Methods

### NewNonStreamGenerateResponse

`func NewNonStreamGenerateResponse() *NonStreamGenerateResponse`

NewNonStreamGenerateResponse instantiates a new NonStreamGenerateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNonStreamGenerateResponseWithDefaults

`func NewNonStreamGenerateResponseWithDefaults() *NonStreamGenerateResponse`

NewNonStreamGenerateResponseWithDefaults instantiates a new NonStreamGenerateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NonStreamGenerateResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NonStreamGenerateResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NonStreamGenerateResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *NonStreamGenerateResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetData

`func (o *NonStreamGenerateResponse) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *NonStreamGenerateResponse) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *NonStreamGenerateResponse) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *NonStreamGenerateResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetUsage

`func (o *NonStreamGenerateResponse) GetUsage() TokenUsage`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *NonStreamGenerateResponse) GetUsageOk() (*TokenUsage, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *NonStreamGenerateResponse) SetUsage(v TokenUsage)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *NonStreamGenerateResponse) HasUsage() bool`

HasUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


