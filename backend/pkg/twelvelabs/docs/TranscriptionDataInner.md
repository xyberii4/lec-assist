# TranscriptionDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Start** | Pointer to **float32** | The start of the time range, expressed in seconds. | [optional] 
**End** | Pointer to **float32** | The end of the time range, expressed in seconds. | [optional] 
**Value** | Pointer to **string** | Text representing the spoken words within this time range. | [optional] 

## Methods

### NewTranscriptionDataInner

`func NewTranscriptionDataInner() *TranscriptionDataInner`

NewTranscriptionDataInner instantiates a new TranscriptionDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTranscriptionDataInnerWithDefaults

`func NewTranscriptionDataInnerWithDefaults() *TranscriptionDataInner`

NewTranscriptionDataInnerWithDefaults instantiates a new TranscriptionDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStart

`func (o *TranscriptionDataInner) GetStart() float32`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *TranscriptionDataInner) GetStartOk() (*float32, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *TranscriptionDataInner) SetStart(v float32)`

SetStart sets Start field to given value.

### HasStart

`func (o *TranscriptionDataInner) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetEnd

`func (o *TranscriptionDataInner) GetEnd() float32`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *TranscriptionDataInner) GetEndOk() (*float32, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *TranscriptionDataInner) SetEnd(v float32)`

SetEnd sets End field to given value.

### HasEnd

`func (o *TranscriptionDataInner) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetValue

`func (o *TranscriptionDataInner) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *TranscriptionDataInner) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *TranscriptionDataInner) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *TranscriptionDataInner) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


