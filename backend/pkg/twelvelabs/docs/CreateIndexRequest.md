# CreateIndexRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IndexName** | **string** | The name of the index. Make sure you use a succinct and descriptive name.  | 
**Models** | [**[]CreateIndexRequestModelsInner**](CreateIndexRequestModelsInner.md) | An array that specifies the [video understanding models](/v1.3/docs/concepts/models) and the [model options](/v1.3/docs/concepts/modalities#model-options) to be enabled for this index. This determines how the platform processes your videos.  | 
**Addons** | Pointer to **[]string** | An array specifying which add-ons should be enabled. Each entry in the array is an addon, and the following values are supported: - &#x60;thumbnail&#x60;: Enables thumbnail generation.  If you don&#39;t provide this parameter, no add-ons will be enabled.  &lt;Note title&#x3D;\&quot;Notes\&quot;&gt; - You can only enable addons when using the Marengo video understanding model. - You cannot disable an add-on once the index has been created. &lt;/Note&gt;  | [optional] 

## Methods

### NewCreateIndexRequest

`func NewCreateIndexRequest(indexName string, models []CreateIndexRequestModelsInner, ) *CreateIndexRequest`

NewCreateIndexRequest instantiates a new CreateIndexRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateIndexRequestWithDefaults

`func NewCreateIndexRequestWithDefaults() *CreateIndexRequest`

NewCreateIndexRequestWithDefaults instantiates a new CreateIndexRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndexName

`func (o *CreateIndexRequest) GetIndexName() string`

GetIndexName returns the IndexName field if non-nil, zero value otherwise.

### GetIndexNameOk

`func (o *CreateIndexRequest) GetIndexNameOk() (*string, bool)`

GetIndexNameOk returns a tuple with the IndexName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexName

`func (o *CreateIndexRequest) SetIndexName(v string)`

SetIndexName sets IndexName field to given value.


### GetModels

`func (o *CreateIndexRequest) GetModels() []CreateIndexRequestModelsInner`

GetModels returns the Models field if non-nil, zero value otherwise.

### GetModelsOk

`func (o *CreateIndexRequest) GetModelsOk() (*[]CreateIndexRequestModelsInner, bool)`

GetModelsOk returns a tuple with the Models field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModels

`func (o *CreateIndexRequest) SetModels(v []CreateIndexRequestModelsInner)`

SetModels sets Models field to given value.


### GetAddons

`func (o *CreateIndexRequest) GetAddons() []string`

GetAddons returns the Addons field if non-nil, zero value otherwise.

### GetAddonsOk

`func (o *CreateIndexRequest) GetAddonsOk() (*[]string, bool)`

GetAddonsOk returns a tuple with the Addons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddons

`func (o *CreateIndexRequest) SetAddons(v []string)`

SetAddons sets Addons field to given value.

### HasAddons

`func (o *CreateIndexRequest) HasAddons() bool`

HasAddons returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


