# SearchResults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]SearchItem**](SearchItem.md) | An array that contains your search results. For each match found, the model returns the following fields:  | [optional] 
**PageInfo** | Pointer to [**SearchResultsPageInfo**](SearchResultsPageInfo.md) |  | [optional] 
**SearchPool** | Pointer to [**SearchPool**](SearchPool.md) |  | [optional] 

## Methods

### NewSearchResults

`func NewSearchResults() *SearchResults`

NewSearchResults instantiates a new SearchResults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchResultsWithDefaults

`func NewSearchResultsWithDefaults() *SearchResults`

NewSearchResultsWithDefaults instantiates a new SearchResults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *SearchResults) GetData() []SearchItem`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SearchResults) GetDataOk() (*[]SearchItem, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SearchResults) SetData(v []SearchItem)`

SetData sets Data field to given value.

### HasData

`func (o *SearchResults) HasData() bool`

HasData returns a boolean if a field has been set.

### GetPageInfo

`func (o *SearchResults) GetPageInfo() SearchResultsPageInfo`

GetPageInfo returns the PageInfo field if non-nil, zero value otherwise.

### GetPageInfoOk

`func (o *SearchResults) GetPageInfoOk() (*SearchResultsPageInfo, bool)`

GetPageInfoOk returns a tuple with the PageInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageInfo

`func (o *SearchResults) SetPageInfo(v SearchResultsPageInfo)`

SetPageInfo sets PageInfo field to given value.

### HasPageInfo

`func (o *SearchResults) HasPageInfo() bool`

HasPageInfo returns a boolean if a field has been set.

### GetSearchPool

`func (o *SearchResults) GetSearchPool() SearchPool`

GetSearchPool returns the SearchPool field if non-nil, zero value otherwise.

### GetSearchPoolOk

`func (o *SearchResults) GetSearchPoolOk() (*SearchPool, bool)`

GetSearchPoolOk returns a tuple with the SearchPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchPool

`func (o *SearchResults) SetSearchPool(v SearchPool)`

SetSearchPool sets SearchPool field to given value.

### HasSearchPool

`func (o *SearchResults) HasSearchPool() bool`

HasSearchPool returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


