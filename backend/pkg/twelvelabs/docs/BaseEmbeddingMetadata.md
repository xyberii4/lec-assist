# BaseEmbeddingMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InputUrl** | Pointer to **string** | The URL of the media file used to generate the embedding. Present if a URL was provided in the request.  | [optional] 
**InputFilename** | Pointer to **string** | The name of the media file used to generate the embedding. Present if a file was provided in the request.  | [optional] 

## Methods

### NewBaseEmbeddingMetadata

`func NewBaseEmbeddingMetadata() *BaseEmbeddingMetadata`

NewBaseEmbeddingMetadata instantiates a new BaseEmbeddingMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseEmbeddingMetadataWithDefaults

`func NewBaseEmbeddingMetadataWithDefaults() *BaseEmbeddingMetadata`

NewBaseEmbeddingMetadataWithDefaults instantiates a new BaseEmbeddingMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInputUrl

`func (o *BaseEmbeddingMetadata) GetInputUrl() string`

GetInputUrl returns the InputUrl field if non-nil, zero value otherwise.

### GetInputUrlOk

`func (o *BaseEmbeddingMetadata) GetInputUrlOk() (*string, bool)`

GetInputUrlOk returns a tuple with the InputUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputUrl

`func (o *BaseEmbeddingMetadata) SetInputUrl(v string)`

SetInputUrl sets InputUrl field to given value.

### HasInputUrl

`func (o *BaseEmbeddingMetadata) HasInputUrl() bool`

HasInputUrl returns a boolean if a field has been set.

### GetInputFilename

`func (o *BaseEmbeddingMetadata) GetInputFilename() string`

GetInputFilename returns the InputFilename field if non-nil, zero value otherwise.

### GetInputFilenameOk

`func (o *BaseEmbeddingMetadata) GetInputFilenameOk() (*string, bool)`

GetInputFilenameOk returns a tuple with the InputFilename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputFilename

`func (o *BaseEmbeddingMetadata) SetInputFilename(v string)`

SetInputFilename sets InputFilename field to given value.

### HasInputFilename

`func (o *BaseEmbeddingMetadata) HasInputFilename() bool`

HasInputFilename returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


