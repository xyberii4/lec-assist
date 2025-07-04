# VideoVector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | A string representing the unique identifier of a video. The platform creates a new &#x60;video_vector&#x60; object and assigns it a unique identifier when the video has successfully been indexed. Note that video IDs are different from task IDs.  | [optional] 
**CreatedAt** | Pointer to **string** | A string indicating the date and time, in the RFC RFC 3339 format (\&quot;YYYY-MM-DDTHH:mm:ssZ\&quot;), that the video indexing task was created.  | [optional] 
**UpdatedAt** | Pointer to **string** | A string indicating the date and time, in the RFC RFC 3339 format (\&quot;YYYY-MM-DDTHH:mm:ssZ\&quot;), that the video indexing task object was last updated. The platform updates this field every time the video indexing task transitions to a different state.  | [optional] 
**IndexedAt** | Pointer to **string** | A string indicating the date and time, in the RFC RFC 3339 format (\&quot;YYYY-MM-DDTHH:mm:ssZ\&quot;), that the video indexing task has been completed.  | [optional] 
**SystemMetadata** | Pointer to [**VideoVectorSystemMetadata**](VideoVectorSystemMetadata.md) |  | [optional] 

## Methods

### NewVideoVector

`func NewVideoVector() *VideoVector`

NewVideoVector instantiates a new VideoVector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVideoVectorWithDefaults

`func NewVideoVectorWithDefaults() *VideoVector`

NewVideoVectorWithDefaults instantiates a new VideoVector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VideoVector) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VideoVector) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VideoVector) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VideoVector) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *VideoVector) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *VideoVector) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *VideoVector) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *VideoVector) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *VideoVector) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *VideoVector) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *VideoVector) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *VideoVector) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetIndexedAt

`func (o *VideoVector) GetIndexedAt() string`

GetIndexedAt returns the IndexedAt field if non-nil, zero value otherwise.

### GetIndexedAtOk

`func (o *VideoVector) GetIndexedAtOk() (*string, bool)`

GetIndexedAtOk returns a tuple with the IndexedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexedAt

`func (o *VideoVector) SetIndexedAt(v string)`

SetIndexedAt sets IndexedAt field to given value.

### HasIndexedAt

`func (o *VideoVector) HasIndexedAt() bool`

HasIndexedAt returns a boolean if a field has been set.

### GetSystemMetadata

`func (o *VideoVector) GetSystemMetadata() VideoVectorSystemMetadata`

GetSystemMetadata returns the SystemMetadata field if non-nil, zero value otherwise.

### GetSystemMetadataOk

`func (o *VideoVector) GetSystemMetadataOk() (*VideoVectorSystemMetadata, bool)`

GetSystemMetadataOk returns a tuple with the SystemMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemMetadata

`func (o *VideoVector) SetSystemMetadata(v VideoVectorSystemMetadata)`

SetSystemMetadata sets SystemMetadata field to given value.

### HasSystemMetadata

`func (o *VideoVector) HasSystemMetadata() bool`

HasSystemMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


