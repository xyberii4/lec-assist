# TokenUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OutputTokens** | Pointer to **int32** | The number of tokens in the generated text.  | [optional] 

## Methods

### NewTokenUsage

`func NewTokenUsage() *TokenUsage`

NewTokenUsage instantiates a new TokenUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenUsageWithDefaults

`func NewTokenUsageWithDefaults() *TokenUsage`

NewTokenUsageWithDefaults instantiates a new TokenUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOutputTokens

`func (o *TokenUsage) GetOutputTokens() int32`

GetOutputTokens returns the OutputTokens field if non-nil, zero value otherwise.

### GetOutputTokensOk

`func (o *TokenUsage) GetOutputTokensOk() (*int32, bool)`

GetOutputTokensOk returns a tuple with the OutputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputTokens

`func (o *TokenUsage) SetOutputTokens(v int32)`

SetOutputTokens sets OutputTokens field to given value.

### HasOutputTokens

`func (o *TokenUsage) HasOutputTokens() bool`

HasOutputTokens returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


