# InlineObject15AllOfVideoEmbedding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**VideoEmbeddingMetadata**](VideoEmbeddingMetadata.md) |  | [optional] 
**Segments** | Pointer to [**[]VideoSegment**](VideoSegment.md) | An array of objects containing the embeddings for each video segment and the associated information.  | [optional] 

## Methods

### NewInlineObject15AllOfVideoEmbedding

`func NewInlineObject15AllOfVideoEmbedding() *InlineObject15AllOfVideoEmbedding`

NewInlineObject15AllOfVideoEmbedding instantiates a new InlineObject15AllOfVideoEmbedding object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject15AllOfVideoEmbeddingWithDefaults

`func NewInlineObject15AllOfVideoEmbeddingWithDefaults() *InlineObject15AllOfVideoEmbedding`

NewInlineObject15AllOfVideoEmbeddingWithDefaults instantiates a new InlineObject15AllOfVideoEmbedding object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *InlineObject15AllOfVideoEmbedding) GetMetadata() VideoEmbeddingMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *InlineObject15AllOfVideoEmbedding) GetMetadataOk() (*VideoEmbeddingMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *InlineObject15AllOfVideoEmbedding) SetMetadata(v VideoEmbeddingMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *InlineObject15AllOfVideoEmbedding) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSegments

`func (o *InlineObject15AllOfVideoEmbedding) GetSegments() []VideoSegment`

GetSegments returns the Segments field if non-nil, zero value otherwise.

### GetSegmentsOk

`func (o *InlineObject15AllOfVideoEmbedding) GetSegmentsOk() (*[]VideoSegment, bool)`

GetSegmentsOk returns a tuple with the Segments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegments

`func (o *InlineObject15AllOfVideoEmbedding) SetSegments(v []VideoSegment)`

SetSegments sets Segments field to given value.

### HasSegments

`func (o *InlineObject15AllOfVideoEmbedding) HasSegments() bool`

HasSegments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


