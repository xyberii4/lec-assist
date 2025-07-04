# InlineObject20VideosInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VideoId** | Pointer to **string** | The unique identifier of a video. This identifier serves a dual purpose: - It identifies the video itself. - It identifies the associated video indexing task.  For example, due to this dual functionality, you can use this identifier to: - Retrieve details about the video itself by invoking the [&#x60;GET&#x60;](/v1.3/api-reference/videos/retrieve) method of the &#x60;/indexes/{index-id}/videos/{video-id}&#x60; endpoint. - Retrieve the status of the video indexing task by invoking the [&#x60;GET&#x60;](/v1.3/api-reference/tasks/retrieve) method of the &#x60;/tasks/{task_id}&#x60; endpoint.  | [optional] 
**Filename** | Pointer to **string** | The filename of the video. | [optional] 

## Methods

### NewInlineObject20VideosInner

`func NewInlineObject20VideosInner() *InlineObject20VideosInner`

NewInlineObject20VideosInner instantiates a new InlineObject20VideosInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject20VideosInnerWithDefaults

`func NewInlineObject20VideosInnerWithDefaults() *InlineObject20VideosInner`

NewInlineObject20VideosInnerWithDefaults instantiates a new InlineObject20VideosInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVideoId

`func (o *InlineObject20VideosInner) GetVideoId() string`

GetVideoId returns the VideoId field if non-nil, zero value otherwise.

### GetVideoIdOk

`func (o *InlineObject20VideosInner) GetVideoIdOk() (*string, bool)`

GetVideoIdOk returns a tuple with the VideoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoId

`func (o *InlineObject20VideosInner) SetVideoId(v string)`

SetVideoId sets VideoId field to given value.

### HasVideoId

`func (o *InlineObject20VideosInner) HasVideoId() bool`

HasVideoId returns a boolean if a field has been set.

### GetFilename

`func (o *InlineObject20VideosInner) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *InlineObject20VideosInner) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *InlineObject20VideosInner) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *InlineObject20VideosInner) HasFilename() bool`

HasFilename returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


