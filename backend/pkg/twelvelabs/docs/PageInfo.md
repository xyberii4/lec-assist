# PageInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LimitPerPage** | Pointer to **int32** | The maximum number of items on each page. | [optional] 
**Page** | Pointer to **int32** | The page you retrieved. | [optional] 
**TotalPage** | Pointer to **int32** | The total number of pages. | [optional] 
**TotalResults** | Pointer to **int32** | The total number of results. | [optional] 

## Methods

### NewPageInfo

`func NewPageInfo() *PageInfo`

NewPageInfo instantiates a new PageInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPageInfoWithDefaults

`func NewPageInfoWithDefaults() *PageInfo`

NewPageInfoWithDefaults instantiates a new PageInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimitPerPage

`func (o *PageInfo) GetLimitPerPage() int32`

GetLimitPerPage returns the LimitPerPage field if non-nil, zero value otherwise.

### GetLimitPerPageOk

`func (o *PageInfo) GetLimitPerPageOk() (*int32, bool)`

GetLimitPerPageOk returns a tuple with the LimitPerPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitPerPage

`func (o *PageInfo) SetLimitPerPage(v int32)`

SetLimitPerPage sets LimitPerPage field to given value.

### HasLimitPerPage

`func (o *PageInfo) HasLimitPerPage() bool`

HasLimitPerPage returns a boolean if a field has been set.

### GetPage

`func (o *PageInfo) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *PageInfo) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *PageInfo) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *PageInfo) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetTotalPage

`func (o *PageInfo) GetTotalPage() int32`

GetTotalPage returns the TotalPage field if non-nil, zero value otherwise.

### GetTotalPageOk

`func (o *PageInfo) GetTotalPageOk() (*int32, bool)`

GetTotalPageOk returns a tuple with the TotalPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPage

`func (o *PageInfo) SetTotalPage(v int32)`

SetTotalPage sets TotalPage field to given value.

### HasTotalPage

`func (o *PageInfo) HasTotalPage() bool`

HasTotalPage returns a boolean if a field has been set.

### GetTotalResults

`func (o *PageInfo) GetTotalResults() int32`

GetTotalResults returns the TotalResults field if non-nil, zero value otherwise.

### GetTotalResultsOk

`func (o *PageInfo) GetTotalResultsOk() (*int32, bool)`

GetTotalResultsOk returns a tuple with the TotalResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalResults

`func (o *PageInfo) SetTotalResults(v int32)`

SetTotalResults sets TotalResults field to given value.

### HasTotalResults

`func (o *PageInfo) HasTotalResults() bool`

HasTotalResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


