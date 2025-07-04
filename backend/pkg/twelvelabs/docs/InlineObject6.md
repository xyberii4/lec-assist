# InlineObject6

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
**Hls** | Pointer to [**HLSObject**](HLSObject.md) |  | [optional] 

## Methods

### NewInlineObject6

`func NewInlineObject6() *InlineObject6`

NewInlineObject6 instantiates a new InlineObject6 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject6WithDefaults

`func NewInlineObject6WithDefaults() *InlineObject6`

NewInlineObject6WithDefaults instantiates a new InlineObject6 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InlineObject6) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineObject6) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineObject6) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InlineObject6) HasId() bool`

HasId returns a boolean if a field has been set.

### GetVideoId

`func (o *InlineObject6) GetVideoId() string`

GetVideoId returns the VideoId field if non-nil, zero value otherwise.

### GetVideoIdOk

`func (o *InlineObject6) GetVideoIdOk() (*string, bool)`

GetVideoIdOk returns a tuple with the VideoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoId

`func (o *InlineObject6) SetVideoId(v string)`

SetVideoId sets VideoId field to given value.

### HasVideoId

`func (o *InlineObject6) HasVideoId() bool`

HasVideoId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *InlineObject6) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *InlineObject6) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *InlineObject6) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *InlineObject6) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *InlineObject6) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *InlineObject6) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *InlineObject6) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *InlineObject6) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetStatus

`func (o *InlineObject6) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineObject6) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineObject6) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineObject6) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetIndexId

`func (o *InlineObject6) GetIndexId() string`

GetIndexId returns the IndexId field if non-nil, zero value otherwise.

### GetIndexIdOk

`func (o *InlineObject6) GetIndexIdOk() (*string, bool)`

GetIndexIdOk returns a tuple with the IndexId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexId

`func (o *InlineObject6) SetIndexId(v string)`

SetIndexId sets IndexId field to given value.

### HasIndexId

`func (o *InlineObject6) HasIndexId() bool`

HasIndexId returns a boolean if a field has been set.

### GetSystemMetadata

`func (o *InlineObject6) GetSystemMetadata() VideoIndexingTaskSystemMetadata`

GetSystemMetadata returns the SystemMetadata field if non-nil, zero value otherwise.

### GetSystemMetadataOk

`func (o *InlineObject6) GetSystemMetadataOk() (*VideoIndexingTaskSystemMetadata, bool)`

GetSystemMetadataOk returns a tuple with the SystemMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemMetadata

`func (o *InlineObject6) SetSystemMetadata(v VideoIndexingTaskSystemMetadata)`

SetSystemMetadata sets SystemMetadata field to given value.

### HasSystemMetadata

`func (o *InlineObject6) HasSystemMetadata() bool`

HasSystemMetadata returns a boolean if a field has been set.

### GetHls

`func (o *InlineObject6) GetHls() HLSObject`

GetHls returns the Hls field if non-nil, zero value otherwise.

### GetHlsOk

`func (o *InlineObject6) GetHlsOk() (*HLSObject, bool)`

GetHlsOk returns a tuple with the Hls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHls

`func (o *InlineObject6) SetHls(v HLSObject)`

SetHls sets Hls field to given value.

### HasHls

`func (o *InlineObject6) HasHls() bool`

HasHls returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


