# Highlight

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique identifier of the response.  | [optional] 
**Highlights** | Pointer to [**[]HighlightHighlightsInner**](HighlightHighlightsInner.md) | An array that contains the highlights.  | [optional] 
**Usage** | Pointer to [**TokenUsage**](TokenUsage.md) |  | [optional] 

## Methods

### NewHighlight

`func NewHighlight() *Highlight`

NewHighlight instantiates a new Highlight object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHighlightWithDefaults

`func NewHighlightWithDefaults() *Highlight`

NewHighlightWithDefaults instantiates a new Highlight object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Highlight) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Highlight) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Highlight) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Highlight) HasId() bool`

HasId returns a boolean if a field has been set.

### GetHighlights

`func (o *Highlight) GetHighlights() []HighlightHighlightsInner`

GetHighlights returns the Highlights field if non-nil, zero value otherwise.

### GetHighlightsOk

`func (o *Highlight) GetHighlightsOk() (*[]HighlightHighlightsInner, bool)`

GetHighlightsOk returns a tuple with the Highlights field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighlights

`func (o *Highlight) SetHighlights(v []HighlightHighlightsInner)`

SetHighlights sets Highlights field to given value.

### HasHighlights

`func (o *Highlight) HasHighlights() bool`

HasHighlights returns a boolean if a field has been set.

### GetUsage

`func (o *Highlight) GetUsage() TokenUsage`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *Highlight) GetUsageOk() (*TokenUsage, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *Highlight) SetUsage(v TokenUsage)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *Highlight) HasUsage() bool`

HasUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


