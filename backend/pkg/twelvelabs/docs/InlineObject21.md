# InlineObject21

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NotImported** | Pointer to **[]string** | An array of filenames that haven&#39;t yet been imported. | [optional] 
**Validating** | Pointer to [**[]VideoItem**](VideoItem.md) | An array of videos that are being validated. | [optional] 
**Pending** | Pointer to [**[]VideoItem**](VideoItem.md) | An array of videos that are pending. | [optional] 
**Queued** | Pointer to [**[]VideoItem**](VideoItem.md) | An array of videos that are queued for import. | [optional] 
**Indexing** | Pointer to [**[]VideoItem**](VideoItem.md) | An array of videos that are being indexed. | [optional] 
**Ready** | Pointer to [**[]VideoItem**](VideoItem.md) | An array of videos that have successfully been imported. | [optional] 
**Failed** | Pointer to [**[]VideoItemFailed**](VideoItemFailed.md) | An array of videos that failed to import. | [optional] 

## Methods

### NewInlineObject21

`func NewInlineObject21() *InlineObject21`

NewInlineObject21 instantiates a new InlineObject21 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject21WithDefaults

`func NewInlineObject21WithDefaults() *InlineObject21`

NewInlineObject21WithDefaults instantiates a new InlineObject21 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotImported

`func (o *InlineObject21) GetNotImported() []string`

GetNotImported returns the NotImported field if non-nil, zero value otherwise.

### GetNotImportedOk

`func (o *InlineObject21) GetNotImportedOk() (*[]string, bool)`

GetNotImportedOk returns a tuple with the NotImported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotImported

`func (o *InlineObject21) SetNotImported(v []string)`

SetNotImported sets NotImported field to given value.

### HasNotImported

`func (o *InlineObject21) HasNotImported() bool`

HasNotImported returns a boolean if a field has been set.

### GetValidating

`func (o *InlineObject21) GetValidating() []VideoItem`

GetValidating returns the Validating field if non-nil, zero value otherwise.

### GetValidatingOk

`func (o *InlineObject21) GetValidatingOk() (*[]VideoItem, bool)`

GetValidatingOk returns a tuple with the Validating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidating

`func (o *InlineObject21) SetValidating(v []VideoItem)`

SetValidating sets Validating field to given value.

### HasValidating

`func (o *InlineObject21) HasValidating() bool`

HasValidating returns a boolean if a field has been set.

### GetPending

`func (o *InlineObject21) GetPending() []VideoItem`

GetPending returns the Pending field if non-nil, zero value otherwise.

### GetPendingOk

`func (o *InlineObject21) GetPendingOk() (*[]VideoItem, bool)`

GetPendingOk returns a tuple with the Pending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPending

`func (o *InlineObject21) SetPending(v []VideoItem)`

SetPending sets Pending field to given value.

### HasPending

`func (o *InlineObject21) HasPending() bool`

HasPending returns a boolean if a field has been set.

### GetQueued

`func (o *InlineObject21) GetQueued() []VideoItem`

GetQueued returns the Queued field if non-nil, zero value otherwise.

### GetQueuedOk

`func (o *InlineObject21) GetQueuedOk() (*[]VideoItem, bool)`

GetQueuedOk returns a tuple with the Queued field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueued

`func (o *InlineObject21) SetQueued(v []VideoItem)`

SetQueued sets Queued field to given value.

### HasQueued

`func (o *InlineObject21) HasQueued() bool`

HasQueued returns a boolean if a field has been set.

### GetIndexing

`func (o *InlineObject21) GetIndexing() []VideoItem`

GetIndexing returns the Indexing field if non-nil, zero value otherwise.

### GetIndexingOk

`func (o *InlineObject21) GetIndexingOk() (*[]VideoItem, bool)`

GetIndexingOk returns a tuple with the Indexing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexing

`func (o *InlineObject21) SetIndexing(v []VideoItem)`

SetIndexing sets Indexing field to given value.

### HasIndexing

`func (o *InlineObject21) HasIndexing() bool`

HasIndexing returns a boolean if a field has been set.

### GetReady

`func (o *InlineObject21) GetReady() []VideoItem`

GetReady returns the Ready field if non-nil, zero value otherwise.

### GetReadyOk

`func (o *InlineObject21) GetReadyOk() (*[]VideoItem, bool)`

GetReadyOk returns a tuple with the Ready field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReady

`func (o *InlineObject21) SetReady(v []VideoItem)`

SetReady sets Ready field to given value.

### HasReady

`func (o *InlineObject21) HasReady() bool`

HasReady returns a boolean if a field has been set.

### GetFailed

`func (o *InlineObject21) GetFailed() []VideoItemFailed`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *InlineObject21) GetFailedOk() (*[]VideoItemFailed, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *InlineObject21) SetFailed(v []VideoItemFailed)`

SetFailed sets Failed field to given value.

### HasFailed

`func (o *InlineObject21) HasFailed() bool`

HasFailed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


