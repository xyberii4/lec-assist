# VideoEmbeddingTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The unique identifier of the video embedding task.  | [optional] 
**ModelName** | Pointer to **string** | The name of the video understanding model the platform used to create the embedding.  | [optional] 
**Status** | Pointer to **string** | A string indicating the status of the video indexing task. It can take one of the following values: &#x60;processing&#x60;, &#x60;ready&#x60; or &#x60;failed&#x60;.  | [optional] 
**CreatedAt** | Pointer to **time.Time** | The date and time, in the RFC 3339 format (\&quot;YYYY-MM-DDTHH:mm:ssZ\&quot;), that the video embedding task was created.  | [optional] 
**VideoEmbedding** | Pointer to [**VideoEmbeddingTaskVideoEmbedding**](VideoEmbeddingTaskVideoEmbedding.md) |  | [optional] 

## Methods

### NewVideoEmbeddingTask

`func NewVideoEmbeddingTask() *VideoEmbeddingTask`

NewVideoEmbeddingTask instantiates a new VideoEmbeddingTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVideoEmbeddingTaskWithDefaults

`func NewVideoEmbeddingTaskWithDefaults() *VideoEmbeddingTask`

NewVideoEmbeddingTaskWithDefaults instantiates a new VideoEmbeddingTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VideoEmbeddingTask) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VideoEmbeddingTask) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VideoEmbeddingTask) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VideoEmbeddingTask) HasId() bool`

HasId returns a boolean if a field has been set.

### GetModelName

`func (o *VideoEmbeddingTask) GetModelName() string`

GetModelName returns the ModelName field if non-nil, zero value otherwise.

### GetModelNameOk

`func (o *VideoEmbeddingTask) GetModelNameOk() (*string, bool)`

GetModelNameOk returns a tuple with the ModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelName

`func (o *VideoEmbeddingTask) SetModelName(v string)`

SetModelName sets ModelName field to given value.

### HasModelName

`func (o *VideoEmbeddingTask) HasModelName() bool`

HasModelName returns a boolean if a field has been set.

### GetStatus

`func (o *VideoEmbeddingTask) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VideoEmbeddingTask) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VideoEmbeddingTask) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VideoEmbeddingTask) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCreatedAt

`func (o *VideoEmbeddingTask) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *VideoEmbeddingTask) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *VideoEmbeddingTask) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *VideoEmbeddingTask) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetVideoEmbedding

`func (o *VideoEmbeddingTask) GetVideoEmbedding() VideoEmbeddingTaskVideoEmbedding`

GetVideoEmbedding returns the VideoEmbedding field if non-nil, zero value otherwise.

### GetVideoEmbeddingOk

`func (o *VideoEmbeddingTask) GetVideoEmbeddingOk() (*VideoEmbeddingTaskVideoEmbedding, bool)`

GetVideoEmbeddingOk returns a tuple with the VideoEmbedding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoEmbedding

`func (o *VideoEmbeddingTask) SetVideoEmbedding(v VideoEmbeddingTaskVideoEmbedding)`

SetVideoEmbedding sets VideoEmbedding field to given value.

### HasVideoEmbedding

`func (o *VideoEmbeddingTask) HasVideoEmbedding() bool`

HasVideoEmbedding returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


