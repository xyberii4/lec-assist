# InlineObject4Embedding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ModelName** | Pointer to **string** | The name of the video understanding model used to create the embedding. | [optional] 
**VideoEmbedding** | Pointer to [**InlineObject4EmbeddingVideoEmbedding**](InlineObject4EmbeddingVideoEmbedding.md) |  | [optional] 

## Methods

### NewInlineObject4Embedding

`func NewInlineObject4Embedding() *InlineObject4Embedding`

NewInlineObject4Embedding instantiates a new InlineObject4Embedding object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject4EmbeddingWithDefaults

`func NewInlineObject4EmbeddingWithDefaults() *InlineObject4Embedding`

NewInlineObject4EmbeddingWithDefaults instantiates a new InlineObject4Embedding object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModelName

`func (o *InlineObject4Embedding) GetModelName() string`

GetModelName returns the ModelName field if non-nil, zero value otherwise.

### GetModelNameOk

`func (o *InlineObject4Embedding) GetModelNameOk() (*string, bool)`

GetModelNameOk returns a tuple with the ModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelName

`func (o *InlineObject4Embedding) SetModelName(v string)`

SetModelName sets ModelName field to given value.

### HasModelName

`func (o *InlineObject4Embedding) HasModelName() bool`

HasModelName returns a boolean if a field has been set.

### GetVideoEmbedding

`func (o *InlineObject4Embedding) GetVideoEmbedding() InlineObject4EmbeddingVideoEmbedding`

GetVideoEmbedding returns the VideoEmbedding field if non-nil, zero value otherwise.

### GetVideoEmbeddingOk

`func (o *InlineObject4Embedding) GetVideoEmbeddingOk() (*InlineObject4EmbeddingVideoEmbedding, bool)`

GetVideoEmbeddingOk returns a tuple with the VideoEmbedding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoEmbedding

`func (o *InlineObject4Embedding) SetVideoEmbedding(v InlineObject4EmbeddingVideoEmbedding)`

SetVideoEmbedding sets VideoEmbedding field to given value.

### HasVideoEmbedding

`func (o *InlineObject4Embedding) HasVideoEmbedding() bool`

HasVideoEmbedding returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


