# HLSObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VideoUrl** | Pointer to **string** | A string representing the URL of the video. You can then use this URL to access the stream over the &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/HTTP_Live_Streaming\&quot; target&#x3D;\&quot;_blank\&quot;&gt;HLS&lt;/a&gt; protocol. | [optional] 
**ThumbnailUrls** | Pointer to **[]string** | An array containing the URL of the thumbnail. | [optional] 
**Status** | Pointer to **string** | A string representing the encoding status of the video file from its original format to a streamable format. Possible values: - &#x60;PROCESSING&#x60;: Video is currently being encoded and is not yet ready for streaming - &#x60;COMPLETE&#x60;: Encoding has successfully finished and the video is ready for streaming - &#x60;CANCELED&#x60;: Encoding was manually canceled before completion - &#x60;ERROR&#x60;: An error occurred during the encoding process  | [optional] 
**UpdatedAt** | Pointer to **string** | A string indicating the date and time, in the RFC 3339 format (\&quot;YYYY-MM-DDTHH:mm:ssZ\&quot;), that the encoding status was last updated. | [optional] 

## Methods

### NewHLSObject

`func NewHLSObject() *HLSObject`

NewHLSObject instantiates a new HLSObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHLSObjectWithDefaults

`func NewHLSObjectWithDefaults() *HLSObject`

NewHLSObjectWithDefaults instantiates a new HLSObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVideoUrl

`func (o *HLSObject) GetVideoUrl() string`

GetVideoUrl returns the VideoUrl field if non-nil, zero value otherwise.

### GetVideoUrlOk

`func (o *HLSObject) GetVideoUrlOk() (*string, bool)`

GetVideoUrlOk returns a tuple with the VideoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoUrl

`func (o *HLSObject) SetVideoUrl(v string)`

SetVideoUrl sets VideoUrl field to given value.

### HasVideoUrl

`func (o *HLSObject) HasVideoUrl() bool`

HasVideoUrl returns a boolean if a field has been set.

### GetThumbnailUrls

`func (o *HLSObject) GetThumbnailUrls() []string`

GetThumbnailUrls returns the ThumbnailUrls field if non-nil, zero value otherwise.

### GetThumbnailUrlsOk

`func (o *HLSObject) GetThumbnailUrlsOk() (*[]string, bool)`

GetThumbnailUrlsOk returns a tuple with the ThumbnailUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnailUrls

`func (o *HLSObject) SetThumbnailUrls(v []string)`

SetThumbnailUrls sets ThumbnailUrls field to given value.

### HasThumbnailUrls

`func (o *HLSObject) HasThumbnailUrls() bool`

HasThumbnailUrls returns a boolean if a field has been set.

### GetStatus

`func (o *HLSObject) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HLSObject) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HLSObject) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *HLSObject) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *HLSObject) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *HLSObject) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *HLSObject) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *HLSObject) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


