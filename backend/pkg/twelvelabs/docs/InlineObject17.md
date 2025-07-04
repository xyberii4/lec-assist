# InlineObject17

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The unique identifier of the video embedding task.  | [optional] 
**Status** | Pointer to **string** | A string indicating the status of the video indexing task. It can take one of the following values: &#x60;processing&#x60;, &#x60;ready&#x60; or &#x60;failed&#x60;.  | [optional] 
**ModelName** | Pointer to **string** | The name of the video understanding model the platform used to create the embedding.  | [optional] 
**VideoEmbedding** | Pointer to [**VideoEmbeddingTaskVideoEmbedding**](VideoEmbeddingTaskVideoEmbedding.md) |  | [optional] 

## Methods

### NewInlineObject17

`func NewInlineObject17() *InlineObject17`

NewInlineObject17 instantiates a new InlineObject17 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject17WithDefaults

`func NewInlineObject17WithDefaults() *InlineObject17`

NewInlineObject17WithDefaults instantiates a new InlineObject17 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InlineObject17) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineObject17) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineObject17) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InlineObject17) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *InlineObject17) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineObject17) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineObject17) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineObject17) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetModelName

`func (o *InlineObject17) GetModelName() string`

GetModelName returns the ModelName field if non-nil, zero value otherwise.

### GetModelNameOk

`func (o *InlineObject17) GetModelNameOk() (*string, bool)`

GetModelNameOk returns a tuple with the ModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelName

`func (o *InlineObject17) SetModelName(v string)`

SetModelName sets ModelName field to given value.

### HasModelName

`func (o *InlineObject17) HasModelName() bool`

HasModelName returns a boolean if a field has been set.

### GetVideoEmbedding

`func (o *InlineObject17) GetVideoEmbedding() VideoEmbeddingTaskVideoEmbedding`

GetVideoEmbedding returns the VideoEmbedding field if non-nil, zero value otherwise.

### GetVideoEmbeddingOk

`func (o *InlineObject17) GetVideoEmbeddingOk() (*VideoEmbeddingTaskVideoEmbedding, bool)`

GetVideoEmbeddingOk returns a tuple with the VideoEmbedding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoEmbedding

`func (o *InlineObject17) SetVideoEmbedding(v VideoEmbeddingTaskVideoEmbedding)`

SetVideoEmbedding sets VideoEmbedding field to given value.

### HasVideoEmbedding

`func (o *InlineObject17) HasVideoEmbedding() bool`

HasVideoEmbedding returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


