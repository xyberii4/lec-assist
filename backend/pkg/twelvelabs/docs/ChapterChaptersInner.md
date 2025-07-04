# ChapterChaptersInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChapterNumber** | Pointer to **int32** | Represents the sequence number of the chapter. Note that this field starts at 0. Ensure to interpret it accordingly in your application.  | [optional] 
**Start** | Pointer to **int32** | The starting time of the chapter, measured in seconds from the beginning of the video.  | [optional] 
**End** | Pointer to **int32** | The ending time of the chapter, measured in seconds from the beginning of the video.  | [optional] 
**ChapterTitle** | Pointer to **string** | The title of the chapter.  | [optional] 
**ChapterSummary** | Pointer to **string** | A brief summary describing the content of the chapter.  | [optional] 

## Methods

### NewChapterChaptersInner

`func NewChapterChaptersInner() *ChapterChaptersInner`

NewChapterChaptersInner instantiates a new ChapterChaptersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChapterChaptersInnerWithDefaults

`func NewChapterChaptersInnerWithDefaults() *ChapterChaptersInner`

NewChapterChaptersInnerWithDefaults instantiates a new ChapterChaptersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChapterNumber

`func (o *ChapterChaptersInner) GetChapterNumber() int32`

GetChapterNumber returns the ChapterNumber field if non-nil, zero value otherwise.

### GetChapterNumberOk

`func (o *ChapterChaptersInner) GetChapterNumberOk() (*int32, bool)`

GetChapterNumberOk returns a tuple with the ChapterNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChapterNumber

`func (o *ChapterChaptersInner) SetChapterNumber(v int32)`

SetChapterNumber sets ChapterNumber field to given value.

### HasChapterNumber

`func (o *ChapterChaptersInner) HasChapterNumber() bool`

HasChapterNumber returns a boolean if a field has been set.

### GetStart

`func (o *ChapterChaptersInner) GetStart() int32`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *ChapterChaptersInner) GetStartOk() (*int32, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *ChapterChaptersInner) SetStart(v int32)`

SetStart sets Start field to given value.

### HasStart

`func (o *ChapterChaptersInner) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetEnd

`func (o *ChapterChaptersInner) GetEnd() int32`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *ChapterChaptersInner) GetEndOk() (*int32, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *ChapterChaptersInner) SetEnd(v int32)`

SetEnd sets End field to given value.

### HasEnd

`func (o *ChapterChaptersInner) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetChapterTitle

`func (o *ChapterChaptersInner) GetChapterTitle() string`

GetChapterTitle returns the ChapterTitle field if non-nil, zero value otherwise.

### GetChapterTitleOk

`func (o *ChapterChaptersInner) GetChapterTitleOk() (*string, bool)`

GetChapterTitleOk returns a tuple with the ChapterTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChapterTitle

`func (o *ChapterChaptersInner) SetChapterTitle(v string)`

SetChapterTitle sets ChapterTitle field to given value.

### HasChapterTitle

`func (o *ChapterChaptersInner) HasChapterTitle() bool`

HasChapterTitle returns a boolean if a field has been set.

### GetChapterSummary

`func (o *ChapterChaptersInner) GetChapterSummary() string`

GetChapterSummary returns the ChapterSummary field if non-nil, zero value otherwise.

### GetChapterSummaryOk

`func (o *ChapterChaptersInner) GetChapterSummaryOk() (*string, bool)`

GetChapterSummaryOk returns a tuple with the ChapterSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChapterSummary

`func (o *ChapterChaptersInner) SetChapterSummary(v string)`

SetChapterSummary sets ChapterSummary field to given value.

### HasChapterSummary

`func (o *ChapterChaptersInner) HasChapterSummary() bool`

HasChapterSummary returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


