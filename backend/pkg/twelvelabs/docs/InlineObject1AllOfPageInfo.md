# InlineObject1AllOfPageInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LimitPerPage** | Pointer to **int32** | The maximum number of items on each page. When grouping by video, this field represents the maximum number of videos per page. Otherwise, it represents the maximum number of video clips per page.  | [optional] 
**PageExpiresAt** | Pointer to **string** | A string representing the date and time, in the RFC 3339 format (\&quot;YYYY-MM-DDTHH:mm:ssZ\&quot;), that the page expires.  | [optional] 
**TotalResults** | Pointer to **int32** | The total number of results. When grouping by video, this field represents the total number of video clips matching your query. Otherwise , this field represents the total number of videos.  | [optional] 
**TotalInnerMatches** | Pointer to **int32** | When grouping by video, the platform return this field that shows the total number of video clips matching your query.  | [optional] 
**NextPageToken** | Pointer to **string** | The unique identifier of the next page. | [optional] 
**PrevPageToken** | Pointer to **string** | The unique identifier of the previous page. | [optional] 

## Methods

### NewInlineObject1AllOfPageInfo

`func NewInlineObject1AllOfPageInfo() *InlineObject1AllOfPageInfo`

NewInlineObject1AllOfPageInfo instantiates a new InlineObject1AllOfPageInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject1AllOfPageInfoWithDefaults

`func NewInlineObject1AllOfPageInfoWithDefaults() *InlineObject1AllOfPageInfo`

NewInlineObject1AllOfPageInfoWithDefaults instantiates a new InlineObject1AllOfPageInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimitPerPage

`func (o *InlineObject1AllOfPageInfo) GetLimitPerPage() int32`

GetLimitPerPage returns the LimitPerPage field if non-nil, zero value otherwise.

### GetLimitPerPageOk

`func (o *InlineObject1AllOfPageInfo) GetLimitPerPageOk() (*int32, bool)`

GetLimitPerPageOk returns a tuple with the LimitPerPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitPerPage

`func (o *InlineObject1AllOfPageInfo) SetLimitPerPage(v int32)`

SetLimitPerPage sets LimitPerPage field to given value.

### HasLimitPerPage

`func (o *InlineObject1AllOfPageInfo) HasLimitPerPage() bool`

HasLimitPerPage returns a boolean if a field has been set.

### GetPageExpiresAt

`func (o *InlineObject1AllOfPageInfo) GetPageExpiresAt() string`

GetPageExpiresAt returns the PageExpiresAt field if non-nil, zero value otherwise.

### GetPageExpiresAtOk

`func (o *InlineObject1AllOfPageInfo) GetPageExpiresAtOk() (*string, bool)`

GetPageExpiresAtOk returns a tuple with the PageExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageExpiresAt

`func (o *InlineObject1AllOfPageInfo) SetPageExpiresAt(v string)`

SetPageExpiresAt sets PageExpiresAt field to given value.

### HasPageExpiresAt

`func (o *InlineObject1AllOfPageInfo) HasPageExpiresAt() bool`

HasPageExpiresAt returns a boolean if a field has been set.

### GetTotalResults

`func (o *InlineObject1AllOfPageInfo) GetTotalResults() int32`

GetTotalResults returns the TotalResults field if non-nil, zero value otherwise.

### GetTotalResultsOk

`func (o *InlineObject1AllOfPageInfo) GetTotalResultsOk() (*int32, bool)`

GetTotalResultsOk returns a tuple with the TotalResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalResults

`func (o *InlineObject1AllOfPageInfo) SetTotalResults(v int32)`

SetTotalResults sets TotalResults field to given value.

### HasTotalResults

`func (o *InlineObject1AllOfPageInfo) HasTotalResults() bool`

HasTotalResults returns a boolean if a field has been set.

### GetTotalInnerMatches

`func (o *InlineObject1AllOfPageInfo) GetTotalInnerMatches() int32`

GetTotalInnerMatches returns the TotalInnerMatches field if non-nil, zero value otherwise.

### GetTotalInnerMatchesOk

`func (o *InlineObject1AllOfPageInfo) GetTotalInnerMatchesOk() (*int32, bool)`

GetTotalInnerMatchesOk returns a tuple with the TotalInnerMatches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalInnerMatches

`func (o *InlineObject1AllOfPageInfo) SetTotalInnerMatches(v int32)`

SetTotalInnerMatches sets TotalInnerMatches field to given value.

### HasTotalInnerMatches

`func (o *InlineObject1AllOfPageInfo) HasTotalInnerMatches() bool`

HasTotalInnerMatches returns a boolean if a field has been set.

### GetNextPageToken

`func (o *InlineObject1AllOfPageInfo) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *InlineObject1AllOfPageInfo) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *InlineObject1AllOfPageInfo) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *InlineObject1AllOfPageInfo) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### GetPrevPageToken

`func (o *InlineObject1AllOfPageInfo) GetPrevPageToken() string`

GetPrevPageToken returns the PrevPageToken field if non-nil, zero value otherwise.

### GetPrevPageTokenOk

`func (o *InlineObject1AllOfPageInfo) GetPrevPageTokenOk() (*string, bool)`

GetPrevPageTokenOk returns a tuple with the PrevPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevPageToken

`func (o *InlineObject1AllOfPageInfo) SetPrevPageToken(v string)`

SetPrevPageToken sets PrevPageToken field to given value.

### HasPrevPageToken

`func (o *InlineObject1AllOfPageInfo) HasPrevPageToken() bool`

HasPrevPageToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


