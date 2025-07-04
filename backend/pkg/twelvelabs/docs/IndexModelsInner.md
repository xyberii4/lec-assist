# IndexModelsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ModelName** | Pointer to **string** | A string representing the name of the model. | [optional] 
**ModelOptions** | Pointer to **[]string** | An array of strings that contains the [model options](/v1.3/docs/concepts/modalities#model-options) enabled for this index.  | [optional] 

## Methods

### NewIndexModelsInner

`func NewIndexModelsInner() *IndexModelsInner`

NewIndexModelsInner instantiates a new IndexModelsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIndexModelsInnerWithDefaults

`func NewIndexModelsInnerWithDefaults() *IndexModelsInner`

NewIndexModelsInnerWithDefaults instantiates a new IndexModelsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModelName

`func (o *IndexModelsInner) GetModelName() string`

GetModelName returns the ModelName field if non-nil, zero value otherwise.

### GetModelNameOk

`func (o *IndexModelsInner) GetModelNameOk() (*string, bool)`

GetModelNameOk returns a tuple with the ModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelName

`func (o *IndexModelsInner) SetModelName(v string)`

SetModelName sets ModelName field to given value.

### HasModelName

`func (o *IndexModelsInner) HasModelName() bool`

HasModelName returns a boolean if a field has been set.

### GetModelOptions

`func (o *IndexModelsInner) GetModelOptions() []string`

GetModelOptions returns the ModelOptions field if non-nil, zero value otherwise.

### GetModelOptionsOk

`func (o *IndexModelsInner) GetModelOptionsOk() (*[]string, bool)`

GetModelOptionsOk returns a tuple with the ModelOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelOptions

`func (o *IndexModelsInner) SetModelOptions(v []string)`

SetModelOptions sets ModelOptions field to given value.

### HasModelOptions

`func (o *IndexModelsInner) HasModelOptions() bool`

HasModelOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


