# InlineObject19

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique identifier of the response.  | [optional] 
**Summary** | Pointer to **string** | A brief report of the main points of the video.  | [optional] 
**Usage** | Pointer to [**TokenUsage**](TokenUsage.md) |  | [optional] 
**Chapters** | Pointer to [**[]ChapterChaptersInner**](ChapterChaptersInner.md) | An array that contains details about the detected chapters and their content.  | [optional] 
**Highlights** | Pointer to [**[]HighlightHighlightsInner**](HighlightHighlightsInner.md) | An array that contains the highlights.  | [optional] 

## Methods

### NewInlineObject19

`func NewInlineObject19() *InlineObject19`

NewInlineObject19 instantiates a new InlineObject19 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject19WithDefaults

`func NewInlineObject19WithDefaults() *InlineObject19`

NewInlineObject19WithDefaults instantiates a new InlineObject19 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InlineObject19) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineObject19) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineObject19) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InlineObject19) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSummary

`func (o *InlineObject19) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *InlineObject19) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *InlineObject19) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *InlineObject19) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetUsage

`func (o *InlineObject19) GetUsage() TokenUsage`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *InlineObject19) GetUsageOk() (*TokenUsage, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *InlineObject19) SetUsage(v TokenUsage)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *InlineObject19) HasUsage() bool`

HasUsage returns a boolean if a field has been set.

### GetChapters

`func (o *InlineObject19) GetChapters() []ChapterChaptersInner`

GetChapters returns the Chapters field if non-nil, zero value otherwise.

### GetChaptersOk

`func (o *InlineObject19) GetChaptersOk() (*[]ChapterChaptersInner, bool)`

GetChaptersOk returns a tuple with the Chapters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChapters

`func (o *InlineObject19) SetChapters(v []ChapterChaptersInner)`

SetChapters sets Chapters field to given value.

### HasChapters

`func (o *InlineObject19) HasChapters() bool`

HasChapters returns a boolean if a field has been set.

### GetHighlights

`func (o *InlineObject19) GetHighlights() []HighlightHighlightsInner`

GetHighlights returns the Highlights field if non-nil, zero value otherwise.

### GetHighlightsOk

`func (o *InlineObject19) GetHighlightsOk() (*[]HighlightHighlightsInner, bool)`

GetHighlightsOk returns a tuple with the Highlights field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighlights

`func (o *InlineObject19) SetHighlights(v []HighlightHighlightsInner)`

SetHighlights sets Highlights field to given value.

### HasHighlights

`func (o *InlineObject19) HasHighlights() bool`

HasHighlights returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


