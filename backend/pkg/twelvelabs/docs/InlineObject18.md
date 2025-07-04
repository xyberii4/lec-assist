# InlineObject18

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The unique identifier of the video embedding task. You can use the identifier to: - Retrieve the status of your task by invoking the [&#x60;GET&#x60;](/v1.3/api-reference/video-embeddings/retrieve-video-embedding-task-status) method of the &#x60;/embed/tasks/{task_id}/status&#x60; endpoint. - Retrieve the embedding by invoking the  [&#x60;GET&#x60;](/v1.3/api-reference/video-embeddings/retrieve-video-embeddings) method of the &#x60;/embed/tasks/{task_id}&#x60; endpoint.  | [optional] 

## Methods

### NewInlineObject18

`func NewInlineObject18() *InlineObject18`

NewInlineObject18 instantiates a new InlineObject18 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject18WithDefaults

`func NewInlineObject18WithDefaults() *InlineObject18`

NewInlineObject18WithDefaults instantiates a new InlineObject18 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InlineObject18) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineObject18) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineObject18) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InlineObject18) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


