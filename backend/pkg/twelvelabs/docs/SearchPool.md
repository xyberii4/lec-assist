# SearchPool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalCount** | Pointer to **int32** | The number of videos in the index you queried. | [optional] 
**TotalDuration** | Pointer to **float32** | The total duration of the videos. | [optional] 
**IndexId** | Pointer to **string** | The unique identifier of the index. | [optional] 

## Methods

### NewSearchPool

`func NewSearchPool() *SearchPool`

NewSearchPool instantiates a new SearchPool object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchPoolWithDefaults

`func NewSearchPoolWithDefaults() *SearchPool`

NewSearchPoolWithDefaults instantiates a new SearchPool object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalCount

`func (o *SearchPool) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *SearchPool) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *SearchPool) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *SearchPool) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetTotalDuration

`func (o *SearchPool) GetTotalDuration() float32`

GetTotalDuration returns the TotalDuration field if non-nil, zero value otherwise.

### GetTotalDurationOk

`func (o *SearchPool) GetTotalDurationOk() (*float32, bool)`

GetTotalDurationOk returns a tuple with the TotalDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDuration

`func (o *SearchPool) SetTotalDuration(v float32)`

SetTotalDuration sets TotalDuration field to given value.

### HasTotalDuration

`func (o *SearchPool) HasTotalDuration() bool`

HasTotalDuration returns a boolean if a field has been set.

### GetIndexId

`func (o *SearchPool) GetIndexId() string`

GetIndexId returns the IndexId field if non-nil, zero value otherwise.

### GetIndexIdOk

`func (o *SearchPool) GetIndexIdOk() (*string, bool)`

GetIndexIdOk returns a tuple with the IndexId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexId

`func (o *SearchPool) SetIndexId(v string)`

SetIndexId sets IndexId field to given value.

### HasIndexId

`func (o *SearchPool) HasIndexId() bool`

HasIndexId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


