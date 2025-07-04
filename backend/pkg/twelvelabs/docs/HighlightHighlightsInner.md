# HighlightHighlightsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Start** | Pointer to **int32** | The starting time of the highlight, measured in seconds from the beginning of the video.  | [optional] 
**End** | Pointer to **int32** | The ending time of the highlight, measured in seconds from the beginning of the video.  | [optional] 
**Highlight** | Pointer to **string** | The tile of the highlight.  | [optional] 
**HighlightSummary** | Pointer to **string** | A brief description that captures the essence of this part of the video.  | [optional] 

## Methods

### NewHighlightHighlightsInner

`func NewHighlightHighlightsInner() *HighlightHighlightsInner`

NewHighlightHighlightsInner instantiates a new HighlightHighlightsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHighlightHighlightsInnerWithDefaults

`func NewHighlightHighlightsInnerWithDefaults() *HighlightHighlightsInner`

NewHighlightHighlightsInnerWithDefaults instantiates a new HighlightHighlightsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStart

`func (o *HighlightHighlightsInner) GetStart() int32`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *HighlightHighlightsInner) GetStartOk() (*int32, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *HighlightHighlightsInner) SetStart(v int32)`

SetStart sets Start field to given value.

### HasStart

`func (o *HighlightHighlightsInner) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetEnd

`func (o *HighlightHighlightsInner) GetEnd() int32`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *HighlightHighlightsInner) GetEndOk() (*int32, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *HighlightHighlightsInner) SetEnd(v int32)`

SetEnd sets End field to given value.

### HasEnd

`func (o *HighlightHighlightsInner) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetHighlight

`func (o *HighlightHighlightsInner) GetHighlight() string`

GetHighlight returns the Highlight field if non-nil, zero value otherwise.

### GetHighlightOk

`func (o *HighlightHighlightsInner) GetHighlightOk() (*string, bool)`

GetHighlightOk returns a tuple with the Highlight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighlight

`func (o *HighlightHighlightsInner) SetHighlight(v string)`

SetHighlight sets Highlight field to given value.

### HasHighlight

`func (o *HighlightHighlightsInner) HasHighlight() bool`

HasHighlight returns a boolean if a field has been set.

### GetHighlightSummary

`func (o *HighlightHighlightsInner) GetHighlightSummary() string`

GetHighlightSummary returns the HighlightSummary field if non-nil, zero value otherwise.

### GetHighlightSummaryOk

`func (o *HighlightHighlightsInner) GetHighlightSummaryOk() (*string, bool)`

GetHighlightSummaryOk returns a tuple with the HighlightSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighlightSummary

`func (o *HighlightHighlightsInner) SetHighlightSummary(v string)`

SetHighlightSummary sets HighlightSummary field to given value.

### HasHighlightSummary

`func (o *HighlightHighlightsInner) HasHighlightSummary() bool`

HasHighlightSummary returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


