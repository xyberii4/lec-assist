# SearchResultsPageInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LimitPerPage** | Pointer to **int32** | The maximum number of items on each page. When grouping by video, this field represents the maximum number of videos per page. Otherwise, it represents the maximum number of video clips per page.  | [optional] 
**PageExpiresAt** | Pointer to **string** | A string representing the date and time, in the RFC 3339 format (\&quot;YYYY-MM-DDTHH:mm:ssZ\&quot;), that the page expires. | [optional] 
**TotalResults** | Pointer to **int32** | The total number of results. When grouping by video, this field represents the total number of video clips matching your query. Otherwise , this field represents the total number of videos.  | [optional] 
**TotalInnerMatches** | Pointer to **int32** | When grouping by video, the platform return this field that shows the total number of video clips matching your query.  | [optional] 
**NextPageToken** | Pointer to **string** | The unique identifier of the next page. | [optional] 

## Methods

### NewSearchResultsPageInfo

`func NewSearchResultsPageInfo() *SearchResultsPageInfo`

NewSearchResultsPageInfo instantiates a new SearchResultsPageInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchResultsPageInfoWithDefaults

`func NewSearchResultsPageInfoWithDefaults() *SearchResultsPageInfo`

NewSearchResultsPageInfoWithDefaults instantiates a new SearchResultsPageInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimitPerPage

`func (o *SearchResultsPageInfo) GetLimitPerPage() int32`

GetLimitPerPage returns the LimitPerPage field if non-nil, zero value otherwise.

### GetLimitPerPageOk

`func (o *SearchResultsPageInfo) GetLimitPerPageOk() (*int32, bool)`

GetLimitPerPageOk returns a tuple with the LimitPerPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitPerPage

`func (o *SearchResultsPageInfo) SetLimitPerPage(v int32)`

SetLimitPerPage sets LimitPerPage field to given value.

### HasLimitPerPage

`func (o *SearchResultsPageInfo) HasLimitPerPage() bool`

HasLimitPerPage returns a boolean if a field has been set.

### GetPageExpiresAt

`func (o *SearchResultsPageInfo) GetPageExpiresAt() string`

GetPageExpiresAt returns the PageExpiresAt field if non-nil, zero value otherwise.

### GetPageExpiresAtOk

`func (o *SearchResultsPageInfo) GetPageExpiresAtOk() (*string, bool)`

GetPageExpiresAtOk returns a tuple with the PageExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageExpiresAt

`func (o *SearchResultsPageInfo) SetPageExpiresAt(v string)`

SetPageExpiresAt sets PageExpiresAt field to given value.

### HasPageExpiresAt

`func (o *SearchResultsPageInfo) HasPageExpiresAt() bool`

HasPageExpiresAt returns a boolean if a field has been set.

### GetTotalResults

`func (o *SearchResultsPageInfo) GetTotalResults() int32`

GetTotalResults returns the TotalResults field if non-nil, zero value otherwise.

### GetTotalResultsOk

`func (o *SearchResultsPageInfo) GetTotalResultsOk() (*int32, bool)`

GetTotalResultsOk returns a tuple with the TotalResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalResults

`func (o *SearchResultsPageInfo) SetTotalResults(v int32)`

SetTotalResults sets TotalResults field to given value.

### HasTotalResults

`func (o *SearchResultsPageInfo) HasTotalResults() bool`

HasTotalResults returns a boolean if a field has been set.

### GetTotalInnerMatches

`func (o *SearchResultsPageInfo) GetTotalInnerMatches() int32`

GetTotalInnerMatches returns the TotalInnerMatches field if non-nil, zero value otherwise.

### GetTotalInnerMatchesOk

`func (o *SearchResultsPageInfo) GetTotalInnerMatchesOk() (*int32, bool)`

GetTotalInnerMatchesOk returns a tuple with the TotalInnerMatches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalInnerMatches

`func (o *SearchResultsPageInfo) SetTotalInnerMatches(v int32)`

SetTotalInnerMatches sets TotalInnerMatches field to given value.

### HasTotalInnerMatches

`func (o *SearchResultsPageInfo) HasTotalInnerMatches() bool`

HasTotalInnerMatches returns a boolean if a field has been set.

### GetNextPageToken

`func (o *SearchResultsPageInfo) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *SearchResultsPageInfo) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *SearchResultsPageInfo) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *SearchResultsPageInfo) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


