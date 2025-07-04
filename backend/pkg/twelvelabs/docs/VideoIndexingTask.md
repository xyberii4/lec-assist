# VideoIndexingTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | A string representing the unique identifier of the task. It is assigned by the platform when a new task is created.  | [optional] 
**VideoId** | Pointer to **string** | A string representing the unique identifier of the video associated with the specified video indexing task.  | [optional] 
**CreatedAt** | Pointer to **string** | A string indicating the date and time, in the RFC RFC 3339 format (\&quot;YYYY-MM-DDTHH:mm:ssZ\&quot;), that the task was created.  | [optional] 
**UpdatedAt** | Pointer to **string** | A string indicating the date and time, in the RFC RFC 3339 format (\&quot;YYYY-MM-DDTHH:mm:ssZ\&quot;), that the task object was last updated. The platform updates this field every time the video indexing task transitions to a different state.  | [optional] 
**Status** | Pointer to **string** | A string indicating the status of the video indexing task. See the [Task object](/v1.3/api-reference/tasks/the-task-object) page for a list of possible statuses.  | [optional] 
**IndexId** | Pointer to **string** | A string representing the unique identifier of the index to which the video must be uploaded.  | [optional] 
**SystemMetadata** | Pointer to [**VideoIndexingTaskSystemMetadata**](VideoIndexingTaskSystemMetadata.md) |  | [optional] 

## Methods

### NewVideoIndexingTask

`func NewVideoIndexingTask() *VideoIndexingTask`

NewVideoIndexingTask instantiates a new VideoIndexingTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVideoIndexingTaskWithDefaults

`func NewVideoIndexingTaskWithDefaults() *VideoIndexingTask`

NewVideoIndexingTaskWithDefaults instantiates a new VideoIndexingTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VideoIndexingTask) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VideoIndexingTask) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VideoIndexingTask) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VideoIndexingTask) HasId() bool`

HasId returns a boolean if a field has been set.

### GetVideoId

`func (o *VideoIndexingTask) GetVideoId() string`

GetVideoId returns the VideoId field if non-nil, zero value otherwise.

### GetVideoIdOk

`func (o *VideoIndexingTask) GetVideoIdOk() (*string, bool)`

GetVideoIdOk returns a tuple with the VideoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoId

`func (o *VideoIndexingTask) SetVideoId(v string)`

SetVideoId sets VideoId field to given value.

### HasVideoId

`func (o *VideoIndexingTask) HasVideoId() bool`

HasVideoId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *VideoIndexingTask) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *VideoIndexingTask) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *VideoIndexingTask) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *VideoIndexingTask) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *VideoIndexingTask) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *VideoIndexingTask) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *VideoIndexingTask) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *VideoIndexingTask) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetStatus

`func (o *VideoIndexingTask) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VideoIndexingTask) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VideoIndexingTask) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VideoIndexingTask) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetIndexId

`func (o *VideoIndexingTask) GetIndexId() string`

GetIndexId returns the IndexId field if non-nil, zero value otherwise.

### GetIndexIdOk

`func (o *VideoIndexingTask) GetIndexIdOk() (*string, bool)`

GetIndexIdOk returns a tuple with the IndexId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexId

`func (o *VideoIndexingTask) SetIndexId(v string)`

SetIndexId sets IndexId field to given value.

### HasIndexId

`func (o *VideoIndexingTask) HasIndexId() bool`

HasIndexId returns a boolean if a field has been set.

### GetSystemMetadata

`func (o *VideoIndexingTask) GetSystemMetadata() VideoIndexingTaskSystemMetadata`

GetSystemMetadata returns the SystemMetadata field if non-nil, zero value otherwise.

### GetSystemMetadataOk

`func (o *VideoIndexingTask) GetSystemMetadataOk() (*VideoIndexingTaskSystemMetadata, bool)`

GetSystemMetadataOk returns a tuple with the SystemMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemMetadata

`func (o *VideoIndexingTask) SetSystemMetadata(v VideoIndexingTaskSystemMetadata)`

SetSystemMetadata sets SystemMetadata field to given value.

### HasSystemMetadata

`func (o *VideoIndexingTask) HasSystemMetadata() bool`

HasSystemMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


