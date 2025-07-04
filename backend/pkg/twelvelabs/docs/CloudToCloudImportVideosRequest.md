# CloudToCloudImportVideosRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IndexId** | **string** | The unique identifier of the index to which the videos are being uploaded.  | 
**IncrementalImport** | Pointer to **bool** | Specifies whether or not incremental sync is enabled. If set to &#x60;false&#x60;, the platform will synchronize all the files in the bucket.  **Default**: &#x60;true&#x60;.  | [optional] [default to true]
**RetryFailed** | Pointer to **bool** | Determines whether the platform retries failed uploads. When set to &#x60;true&#x60;, the platform attempts to re-upload files that failed during the initial upload process.  **Default**: &#x60;false&#x60;.  | [optional] [default to false]
**UserMetadata** | Pointer to **map[string]interface{}** | Metadata that helps you categorize your videos. You can specify a list of keys and values. Keys must be of type &#x60;string&#x60;, and values can be of the following types: &#x60;string&#x60;, &#x60;integer&#x60;, &#x60;float&#x60; or &#x60;boolean&#x60;.  &lt;Note title&#x3D;\&quot;Notes\&quot;&gt; - The metadata you specify when calling this method applies to all videos imported in this request. -  If you want to store other types of data such as objects or arrays, you must convert your data into string values. - You cannot override any of the predefined metadata (example: duration, width, length, etc) associated with a video. &lt;/Note&gt;  | [optional] 

## Methods

### NewCloudToCloudImportVideosRequest

`func NewCloudToCloudImportVideosRequest(indexId string, ) *CloudToCloudImportVideosRequest`

NewCloudToCloudImportVideosRequest instantiates a new CloudToCloudImportVideosRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudToCloudImportVideosRequestWithDefaults

`func NewCloudToCloudImportVideosRequestWithDefaults() *CloudToCloudImportVideosRequest`

NewCloudToCloudImportVideosRequestWithDefaults instantiates a new CloudToCloudImportVideosRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndexId

`func (o *CloudToCloudImportVideosRequest) GetIndexId() string`

GetIndexId returns the IndexId field if non-nil, zero value otherwise.

### GetIndexIdOk

`func (o *CloudToCloudImportVideosRequest) GetIndexIdOk() (*string, bool)`

GetIndexIdOk returns a tuple with the IndexId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexId

`func (o *CloudToCloudImportVideosRequest) SetIndexId(v string)`

SetIndexId sets IndexId field to given value.


### GetIncrementalImport

`func (o *CloudToCloudImportVideosRequest) GetIncrementalImport() bool`

GetIncrementalImport returns the IncrementalImport field if non-nil, zero value otherwise.

### GetIncrementalImportOk

`func (o *CloudToCloudImportVideosRequest) GetIncrementalImportOk() (*bool, bool)`

GetIncrementalImportOk returns a tuple with the IncrementalImport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncrementalImport

`func (o *CloudToCloudImportVideosRequest) SetIncrementalImport(v bool)`

SetIncrementalImport sets IncrementalImport field to given value.

### HasIncrementalImport

`func (o *CloudToCloudImportVideosRequest) HasIncrementalImport() bool`

HasIncrementalImport returns a boolean if a field has been set.

### GetRetryFailed

`func (o *CloudToCloudImportVideosRequest) GetRetryFailed() bool`

GetRetryFailed returns the RetryFailed field if non-nil, zero value otherwise.

### GetRetryFailedOk

`func (o *CloudToCloudImportVideosRequest) GetRetryFailedOk() (*bool, bool)`

GetRetryFailedOk returns a tuple with the RetryFailed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryFailed

`func (o *CloudToCloudImportVideosRequest) SetRetryFailed(v bool)`

SetRetryFailed sets RetryFailed field to given value.

### HasRetryFailed

`func (o *CloudToCloudImportVideosRequest) HasRetryFailed() bool`

HasRetryFailed returns a boolean if a field has been set.

### GetUserMetadata

`func (o *CloudToCloudImportVideosRequest) GetUserMetadata() map[string]interface{}`

GetUserMetadata returns the UserMetadata field if non-nil, zero value otherwise.

### GetUserMetadataOk

`func (o *CloudToCloudImportVideosRequest) GetUserMetadataOk() (*map[string]interface{}, bool)`

GetUserMetadataOk returns a tuple with the UserMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserMetadata

`func (o *CloudToCloudImportVideosRequest) SetUserMetadata(v map[string]interface{})`

SetUserMetadata sets UserMetadata field to given value.

### HasUserMetadata

`func (o *CloudToCloudImportVideosRequest) HasUserMetadata() bool`

HasUserMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


