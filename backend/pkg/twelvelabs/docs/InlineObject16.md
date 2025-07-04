# InlineObject16

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]VideoEmbeddingTask**](VideoEmbeddingTask.md) | An array that contains up to &#x60;page_limit&#x60; video embedding tasks.  | [optional] 
**PageInfo** | Pointer to [**InlineObject16PageInfo**](InlineObject16PageInfo.md) |  | [optional] 

## Methods

### NewInlineObject16

`func NewInlineObject16() *InlineObject16`

NewInlineObject16 instantiates a new InlineObject16 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject16WithDefaults

`func NewInlineObject16WithDefaults() *InlineObject16`

NewInlineObject16WithDefaults instantiates a new InlineObject16 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *InlineObject16) GetData() []VideoEmbeddingTask`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *InlineObject16) GetDataOk() (*[]VideoEmbeddingTask, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *InlineObject16) SetData(v []VideoEmbeddingTask)`

SetData sets Data field to given value.

### HasData

`func (o *InlineObject16) HasData() bool`

HasData returns a boolean if a field has been set.

### GetPageInfo

`func (o *InlineObject16) GetPageInfo() InlineObject16PageInfo`

GetPageInfo returns the PageInfo field if non-nil, zero value otherwise.

### GetPageInfoOk

`func (o *InlineObject16) GetPageInfoOk() (*InlineObject16PageInfo, bool)`

GetPageInfoOk returns a tuple with the PageInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageInfo

`func (o *InlineObject16) SetPageInfo(v InlineObject16PageInfo)`

SetPageInfo sets PageInfo field to given value.

### HasPageInfo

`func (o *InlineObject16) HasPageInfo() bool`

HasPageInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


