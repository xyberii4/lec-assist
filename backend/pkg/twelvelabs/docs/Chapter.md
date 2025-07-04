# Chapter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique identifier of the response.  | [optional] 
**Chapters** | Pointer to [**[]ChapterChaptersInner**](ChapterChaptersInner.md) | An array that contains details about the detected chapters and their content.  | [optional] 
**Usage** | Pointer to [**TokenUsage**](TokenUsage.md) |  | [optional] 

## Methods

### NewChapter

`func NewChapter() *Chapter`

NewChapter instantiates a new Chapter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChapterWithDefaults

`func NewChapterWithDefaults() *Chapter`

NewChapterWithDefaults instantiates a new Chapter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Chapter) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Chapter) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Chapter) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Chapter) HasId() bool`

HasId returns a boolean if a field has been set.

### GetChapters

`func (o *Chapter) GetChapters() []ChapterChaptersInner`

GetChapters returns the Chapters field if non-nil, zero value otherwise.

### GetChaptersOk

`func (o *Chapter) GetChaptersOk() (*[]ChapterChaptersInner, bool)`

GetChaptersOk returns a tuple with the Chapters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChapters

`func (o *Chapter) SetChapters(v []ChapterChaptersInner)`

SetChapters sets Chapters field to given value.

### HasChapters

`func (o *Chapter) HasChapters() bool`

HasChapters returns a boolean if a field has been set.

### GetUsage

`func (o *Chapter) GetUsage() TokenUsage`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *Chapter) GetUsageOk() (*TokenUsage, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *Chapter) SetUsage(v TokenUsage)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *Chapter) HasUsage() bool`

HasUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


