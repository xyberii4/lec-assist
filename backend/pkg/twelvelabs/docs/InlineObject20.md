# InlineObject20

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FailedFiles** | Pointer to [**[]InlineObject20FailedFilesInner**](InlineObject20FailedFilesInner.md) | A list of the video files that failed to import. Typically, these files did not meet the upload requirements. To resolve this, review the requirements listed in the [Prerequisites](/v1.3/docs/guides/search#prerequisites) section of the **Search** page. | [optional] 
**Videos** | Pointer to [**[]InlineObject20VideosInner**](InlineObject20VideosInner.md) | A list of the videos that will be uploaded and indexed. | [optional] 

## Methods

### NewInlineObject20

`func NewInlineObject20() *InlineObject20`

NewInlineObject20 instantiates a new InlineObject20 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject20WithDefaults

`func NewInlineObject20WithDefaults() *InlineObject20`

NewInlineObject20WithDefaults instantiates a new InlineObject20 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFailedFiles

`func (o *InlineObject20) GetFailedFiles() []InlineObject20FailedFilesInner`

GetFailedFiles returns the FailedFiles field if non-nil, zero value otherwise.

### GetFailedFilesOk

`func (o *InlineObject20) GetFailedFilesOk() (*[]InlineObject20FailedFilesInner, bool)`

GetFailedFilesOk returns a tuple with the FailedFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedFiles

`func (o *InlineObject20) SetFailedFiles(v []InlineObject20FailedFilesInner)`

SetFailedFiles sets FailedFiles field to given value.

### HasFailedFiles

`func (o *InlineObject20) HasFailedFiles() bool`

HasFailedFiles returns a boolean if a field has been set.

### GetVideos

`func (o *InlineObject20) GetVideos() []InlineObject20VideosInner`

GetVideos returns the Videos field if non-nil, zero value otherwise.

### GetVideosOk

`func (o *InlineObject20) GetVideosOk() (*[]InlineObject20VideosInner, bool)`

GetVideosOk returns a tuple with the Videos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideos

`func (o *InlineObject20) SetVideos(v []InlineObject20VideosInner)`

SetVideos sets Videos field to given value.

### HasVideos

`func (o *InlineObject20) HasVideos() bool`

HasVideos returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


