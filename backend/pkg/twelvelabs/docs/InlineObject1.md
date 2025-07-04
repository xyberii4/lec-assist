# InlineObject1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]SearchItem**](SearchItem.md) | An array that contains your search results. For each match found, the model returns the following fields:  | [optional] 
**PageInfo** | Pointer to [**InlineObject1AllOfPageInfo**](InlineObject1AllOfPageInfo.md) |  | [optional] 
**SearchPool** | Pointer to [**SearchPool**](SearchPool.md) |  | [optional] 

## Methods

### NewInlineObject1

`func NewInlineObject1() *InlineObject1`

NewInlineObject1 instantiates a new InlineObject1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject1WithDefaults

`func NewInlineObject1WithDefaults() *InlineObject1`

NewInlineObject1WithDefaults instantiates a new InlineObject1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *InlineObject1) GetData() []SearchItem`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *InlineObject1) GetDataOk() (*[]SearchItem, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *InlineObject1) SetData(v []SearchItem)`

SetData sets Data field to given value.

### HasData

`func (o *InlineObject1) HasData() bool`

HasData returns a boolean if a field has been set.

### GetPageInfo

`func (o *InlineObject1) GetPageInfo() InlineObject1AllOfPageInfo`

GetPageInfo returns the PageInfo field if non-nil, zero value otherwise.

### GetPageInfoOk

`func (o *InlineObject1) GetPageInfoOk() (*InlineObject1AllOfPageInfo, bool)`

GetPageInfoOk returns a tuple with the PageInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageInfo

`func (o *InlineObject1) SetPageInfo(v InlineObject1AllOfPageInfo)`

SetPageInfo sets PageInfo field to given value.

### HasPageInfo

`func (o *InlineObject1) HasPageInfo() bool`

HasPageInfo returns a boolean if a field has been set.

### GetSearchPool

`func (o *InlineObject1) GetSearchPool() SearchPool`

GetSearchPool returns the SearchPool field if non-nil, zero value otherwise.

### GetSearchPoolOk

`func (o *InlineObject1) GetSearchPoolOk() (*SearchPool, bool)`

GetSearchPoolOk returns a tuple with the SearchPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchPool

`func (o *InlineObject1) SetSearchPool(v SearchPool)`

SetSearchPool sets SearchPool field to given value.

### HasSearchPool

`func (o *InlineObject1) HasSearchPool() bool`

HasSearchPool returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


