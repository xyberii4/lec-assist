# ImportLog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IndexId** | Pointer to **string** | The unique identifier of the index associated with this import. | [optional] 
**IndexName** | Pointer to **string** | The name of the index associated with this import. | [optional] 
**CreatedAt** | Pointer to **time.Time** | The date and time, in the RFC 3339 format when the import process was initiated. | [optional] 
**EndedAt** | Pointer to **time.Time** | The date and time, in the RFC 3339 format, when the platform completed importing your videos. A &#x60;null&#x60; value indicates that the import process is still ongoing. | [optional] 
**VideoStatus** | Pointer to [**ImportLogVideoStatus**](ImportLogVideoStatus.md) |  | [optional] 
**FailedFiles** | Pointer to [**[]ImportLogFailedFilesInner**](ImportLogFailedFilesInner.md) | An array containing the video files that failed to import, along with details about the error. | [optional] 

## Methods

### NewImportLog

`func NewImportLog() *ImportLog`

NewImportLog instantiates a new ImportLog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportLogWithDefaults

`func NewImportLogWithDefaults() *ImportLog`

NewImportLogWithDefaults instantiates a new ImportLog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndexId

`func (o *ImportLog) GetIndexId() string`

GetIndexId returns the IndexId field if non-nil, zero value otherwise.

### GetIndexIdOk

`func (o *ImportLog) GetIndexIdOk() (*string, bool)`

GetIndexIdOk returns a tuple with the IndexId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexId

`func (o *ImportLog) SetIndexId(v string)`

SetIndexId sets IndexId field to given value.

### HasIndexId

`func (o *ImportLog) HasIndexId() bool`

HasIndexId returns a boolean if a field has been set.

### GetIndexName

`func (o *ImportLog) GetIndexName() string`

GetIndexName returns the IndexName field if non-nil, zero value otherwise.

### GetIndexNameOk

`func (o *ImportLog) GetIndexNameOk() (*string, bool)`

GetIndexNameOk returns a tuple with the IndexName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexName

`func (o *ImportLog) SetIndexName(v string)`

SetIndexName sets IndexName field to given value.

### HasIndexName

`func (o *ImportLog) HasIndexName() bool`

HasIndexName returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ImportLog) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ImportLog) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ImportLog) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ImportLog) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetEndedAt

`func (o *ImportLog) GetEndedAt() time.Time`

GetEndedAt returns the EndedAt field if non-nil, zero value otherwise.

### GetEndedAtOk

`func (o *ImportLog) GetEndedAtOk() (*time.Time, bool)`

GetEndedAtOk returns a tuple with the EndedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndedAt

`func (o *ImportLog) SetEndedAt(v time.Time)`

SetEndedAt sets EndedAt field to given value.

### HasEndedAt

`func (o *ImportLog) HasEndedAt() bool`

HasEndedAt returns a boolean if a field has been set.

### GetVideoStatus

`func (o *ImportLog) GetVideoStatus() ImportLogVideoStatus`

GetVideoStatus returns the VideoStatus field if non-nil, zero value otherwise.

### GetVideoStatusOk

`func (o *ImportLog) GetVideoStatusOk() (*ImportLogVideoStatus, bool)`

GetVideoStatusOk returns a tuple with the VideoStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoStatus

`func (o *ImportLog) SetVideoStatus(v ImportLogVideoStatus)`

SetVideoStatus sets VideoStatus field to given value.

### HasVideoStatus

`func (o *ImportLog) HasVideoStatus() bool`

HasVideoStatus returns a boolean if a field has been set.

### GetFailedFiles

`func (o *ImportLog) GetFailedFiles() []ImportLogFailedFilesInner`

GetFailedFiles returns the FailedFiles field if non-nil, zero value otherwise.

### GetFailedFilesOk

`func (o *ImportLog) GetFailedFilesOk() (*[]ImportLogFailedFilesInner, bool)`

GetFailedFilesOk returns a tuple with the FailedFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedFiles

`func (o *ImportLog) SetFailedFiles(v []ImportLogFailedFilesInner)`

SetFailedFiles sets FailedFiles field to given value.

### HasFailedFiles

`func (o *ImportLog) HasFailedFiles() bool`

HasFailedFiles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


