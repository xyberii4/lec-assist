# InlineObject4

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The unique identifier of the video. | [optional] 
**CreatedAt** | Pointer to **string** | A string indicating the date and time, in the RFC 3339 format (\&quot;YYYY-MM-DDTHH:mm:ssZ\&quot;), that the video indexing task was created. | [optional] 
**UpdatedAt** | Pointer to **string** | A string indicating the date and time, in the RFC 3339 format (\&quot;YYYY-MM-DDTHH:mm:ssZ\&quot;), that the corresponding video indexing task was last updated. The platform updates this field every time the corresponding video indexing task transitions to a different state. | [optional] 
**IndexedAt** | Pointer to **string** | A string indicating the date and time, in the RFC 3339 format (\&quot;YYYY-MM-DDTHH:mm:ssZ\&quot;), that the video indexing task has been completed. | [optional] 
**SystemMetadata** | Pointer to [**InlineObject4SystemMetadata**](InlineObject4SystemMetadata.md) |  | [optional] 
**UserMetadata** | Pointer to **map[string]interface{}** | User-generated metadata about the video. | [optional] 
**Hls** | Pointer to [**HLSObject**](HLSObject.md) |  | [optional] 
**Embedding** | Pointer to [**InlineObject4Embedding**](InlineObject4Embedding.md) |  | [optional] 
**Transcription** | Pointer to [**[]TranscriptionDataInner**](TranscriptionDataInner.md) | An array of objects that contains the transcription. For each time range for which the platform finds spoken words, it returns an object that contains the fields below. If the platform doesn&#39;t find any spoken words, the &#x60;data&#x60; field is set to &#x60;null&#x60;. Note that the official SDKs will support this feature in a future release. | [optional] 

## Methods

### NewInlineObject4

`func NewInlineObject4() *InlineObject4`

NewInlineObject4 instantiates a new InlineObject4 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject4WithDefaults

`func NewInlineObject4WithDefaults() *InlineObject4`

NewInlineObject4WithDefaults instantiates a new InlineObject4 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InlineObject4) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineObject4) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineObject4) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InlineObject4) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *InlineObject4) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *InlineObject4) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *InlineObject4) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *InlineObject4) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *InlineObject4) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *InlineObject4) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *InlineObject4) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *InlineObject4) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetIndexedAt

`func (o *InlineObject4) GetIndexedAt() string`

GetIndexedAt returns the IndexedAt field if non-nil, zero value otherwise.

### GetIndexedAtOk

`func (o *InlineObject4) GetIndexedAtOk() (*string, bool)`

GetIndexedAtOk returns a tuple with the IndexedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexedAt

`func (o *InlineObject4) SetIndexedAt(v string)`

SetIndexedAt sets IndexedAt field to given value.

### HasIndexedAt

`func (o *InlineObject4) HasIndexedAt() bool`

HasIndexedAt returns a boolean if a field has been set.

### GetSystemMetadata

`func (o *InlineObject4) GetSystemMetadata() InlineObject4SystemMetadata`

GetSystemMetadata returns the SystemMetadata field if non-nil, zero value otherwise.

### GetSystemMetadataOk

`func (o *InlineObject4) GetSystemMetadataOk() (*InlineObject4SystemMetadata, bool)`

GetSystemMetadataOk returns a tuple with the SystemMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemMetadata

`func (o *InlineObject4) SetSystemMetadata(v InlineObject4SystemMetadata)`

SetSystemMetadata sets SystemMetadata field to given value.

### HasSystemMetadata

`func (o *InlineObject4) HasSystemMetadata() bool`

HasSystemMetadata returns a boolean if a field has been set.

### GetUserMetadata

`func (o *InlineObject4) GetUserMetadata() map[string]interface{}`

GetUserMetadata returns the UserMetadata field if non-nil, zero value otherwise.

### GetUserMetadataOk

`func (o *InlineObject4) GetUserMetadataOk() (*map[string]interface{}, bool)`

GetUserMetadataOk returns a tuple with the UserMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserMetadata

`func (o *InlineObject4) SetUserMetadata(v map[string]interface{})`

SetUserMetadata sets UserMetadata field to given value.

### HasUserMetadata

`func (o *InlineObject4) HasUserMetadata() bool`

HasUserMetadata returns a boolean if a field has been set.

### GetHls

`func (o *InlineObject4) GetHls() HLSObject`

GetHls returns the Hls field if non-nil, zero value otherwise.

### GetHlsOk

`func (o *InlineObject4) GetHlsOk() (*HLSObject, bool)`

GetHlsOk returns a tuple with the Hls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHls

`func (o *InlineObject4) SetHls(v HLSObject)`

SetHls sets Hls field to given value.

### HasHls

`func (o *InlineObject4) HasHls() bool`

HasHls returns a boolean if a field has been set.

### GetEmbedding

`func (o *InlineObject4) GetEmbedding() InlineObject4Embedding`

GetEmbedding returns the Embedding field if non-nil, zero value otherwise.

### GetEmbeddingOk

`func (o *InlineObject4) GetEmbeddingOk() (*InlineObject4Embedding, bool)`

GetEmbeddingOk returns a tuple with the Embedding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedding

`func (o *InlineObject4) SetEmbedding(v InlineObject4Embedding)`

SetEmbedding sets Embedding field to given value.

### HasEmbedding

`func (o *InlineObject4) HasEmbedding() bool`

HasEmbedding returns a boolean if a field has been set.

### GetTranscription

`func (o *InlineObject4) GetTranscription() []TranscriptionDataInner`

GetTranscription returns the Transcription field if non-nil, zero value otherwise.

### GetTranscriptionOk

`func (o *InlineObject4) GetTranscriptionOk() (*[]TranscriptionDataInner, bool)`

GetTranscriptionOk returns a tuple with the Transcription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscription

`func (o *InlineObject4) SetTranscription(v []TranscriptionDataInner)`

SetTranscription sets Transcription field to given value.

### HasTranscription

`func (o *InlineObject4) HasTranscription() bool`

HasTranscription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


