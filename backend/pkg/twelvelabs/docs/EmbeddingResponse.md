# EmbeddingResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ModelName** | **string** | The name of the video understanding model the platform has used to create this embedding. | 
**TextEmbedding** | Pointer to [**TextEmbeddingResult**](TextEmbeddingResult.md) | An object that contains the generated text embedding vector and associated information. Present when a text was processed. | [optional] 
**ImageEmbedding** | Pointer to [**ImageEmbeddingResult**](ImageEmbeddingResult.md) | An object that contains the generated image embedding vector and associated information. Present when an image was processed. | [optional] 
**AudioEmbedding** | Pointer to [**AudioEmbeddingResult**](AudioEmbeddingResult.md) | An object that contains the generated audio embedding vector and associated information. Present when an audio file was processed. | [optional] 

## Methods

### NewEmbeddingResponse

`func NewEmbeddingResponse(modelName string, ) *EmbeddingResponse`

NewEmbeddingResponse instantiates a new EmbeddingResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmbeddingResponseWithDefaults

`func NewEmbeddingResponseWithDefaults() *EmbeddingResponse`

NewEmbeddingResponseWithDefaults instantiates a new EmbeddingResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModelName

`func (o *EmbeddingResponse) GetModelName() string`

GetModelName returns the ModelName field if non-nil, zero value otherwise.

### GetModelNameOk

`func (o *EmbeddingResponse) GetModelNameOk() (*string, bool)`

GetModelNameOk returns a tuple with the ModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelName

`func (o *EmbeddingResponse) SetModelName(v string)`

SetModelName sets ModelName field to given value.


### GetTextEmbedding

`func (o *EmbeddingResponse) GetTextEmbedding() TextEmbeddingResult`

GetTextEmbedding returns the TextEmbedding field if non-nil, zero value otherwise.

### GetTextEmbeddingOk

`func (o *EmbeddingResponse) GetTextEmbeddingOk() (*TextEmbeddingResult, bool)`

GetTextEmbeddingOk returns a tuple with the TextEmbedding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextEmbedding

`func (o *EmbeddingResponse) SetTextEmbedding(v TextEmbeddingResult)`

SetTextEmbedding sets TextEmbedding field to given value.

### HasTextEmbedding

`func (o *EmbeddingResponse) HasTextEmbedding() bool`

HasTextEmbedding returns a boolean if a field has been set.

### GetImageEmbedding

`func (o *EmbeddingResponse) GetImageEmbedding() ImageEmbeddingResult`

GetImageEmbedding returns the ImageEmbedding field if non-nil, zero value otherwise.

### GetImageEmbeddingOk

`func (o *EmbeddingResponse) GetImageEmbeddingOk() (*ImageEmbeddingResult, bool)`

GetImageEmbeddingOk returns a tuple with the ImageEmbedding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageEmbedding

`func (o *EmbeddingResponse) SetImageEmbedding(v ImageEmbeddingResult)`

SetImageEmbedding sets ImageEmbedding field to given value.

### HasImageEmbedding

`func (o *EmbeddingResponse) HasImageEmbedding() bool`

HasImageEmbedding returns a boolean if a field has been set.

### GetAudioEmbedding

`func (o *EmbeddingResponse) GetAudioEmbedding() AudioEmbeddingResult`

GetAudioEmbedding returns the AudioEmbedding field if non-nil, zero value otherwise.

### GetAudioEmbeddingOk

`func (o *EmbeddingResponse) GetAudioEmbeddingOk() (*AudioEmbeddingResult, bool)`

GetAudioEmbeddingOk returns a tuple with the AudioEmbedding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioEmbedding

`func (o *EmbeddingResponse) SetAudioEmbedding(v AudioEmbeddingResult)`

SetAudioEmbedding sets AudioEmbedding field to given value.

### HasAudioEmbedding

`func (o *EmbeddingResponse) HasAudioEmbedding() bool`

HasAudioEmbedding returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


