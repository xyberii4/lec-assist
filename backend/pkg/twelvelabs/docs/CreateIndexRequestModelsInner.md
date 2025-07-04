# CreateIndexRequestModelsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ModelName** | **string** | The name of the model. The following models are available:    - **Embedding**: These models are proficient at performing tasks such as search and classification, enabling enhanced video understanding.      - &#x60;marengo2.7&#x60;    - **Generative**: These models generate text based on your videos.      - &#x60;pegasus1.2&#x60;  &lt;Note title&#x3D;\&quot;Note\&quot;&gt; You cannot change the models once the index has been created. &lt;/Note&gt;  For more details, see the [Video understanding models](/v1.3/docs/concepts/models) page.  | 
**ModelOptions** | **[]string** | An array that specifies how the platform will process the videos uploaded to this index. For the Marengo and Pegasus models, you can specify one or both of the following model options: &#x60;visual&#x60; and &#x60;audio&#x60;. For more details, see the [model options](/v1.3/docs/concepts/model-options) page.  | 

## Methods

### NewCreateIndexRequestModelsInner

`func NewCreateIndexRequestModelsInner(modelName string, modelOptions []string, ) *CreateIndexRequestModelsInner`

NewCreateIndexRequestModelsInner instantiates a new CreateIndexRequestModelsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateIndexRequestModelsInnerWithDefaults

`func NewCreateIndexRequestModelsInnerWithDefaults() *CreateIndexRequestModelsInner`

NewCreateIndexRequestModelsInnerWithDefaults instantiates a new CreateIndexRequestModelsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModelName

`func (o *CreateIndexRequestModelsInner) GetModelName() string`

GetModelName returns the ModelName field if non-nil, zero value otherwise.

### GetModelNameOk

`func (o *CreateIndexRequestModelsInner) GetModelNameOk() (*string, bool)`

GetModelNameOk returns a tuple with the ModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelName

`func (o *CreateIndexRequestModelsInner) SetModelName(v string)`

SetModelName sets ModelName field to given value.


### GetModelOptions

`func (o *CreateIndexRequestModelsInner) GetModelOptions() []string`

GetModelOptions returns the ModelOptions field if non-nil, zero value otherwise.

### GetModelOptionsOk

`func (o *CreateIndexRequestModelsInner) GetModelOptionsOk() (*[]string, bool)`

GetModelOptionsOk returns a tuple with the ModelOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelOptions

`func (o *CreateIndexRequestModelsInner) SetModelOptions(v []string)`

SetModelOptions sets ModelOptions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


