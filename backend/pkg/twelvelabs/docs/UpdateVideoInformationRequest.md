# UpdateVideoInformationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserMetadata** | Pointer to **map[string]interface{}** | Metadata that helps you categorize your videos. You can specify a list of keys and values. Keys must be of type &#x60;string&#x60;, and values can be of the following types: &#x60;string&#x60;, &#x60;integer&#x60;, &#x60;float&#x60; or &#x60;boolean&#x60;.  **Example**: &#x60;&#x60;&#x60;JSON \&quot;user_metadata\&quot;: {   \&quot;category\&quot;: \&quot;recentlyAdded\&quot;,   \&quot;batchNumber\&quot;: 5,   \&quot;rating\&quot;: 9.3,   \&quot;needsReview\&quot;: true } &#x60;&#x60;&#x60;  &lt;Note title&#x3D;\&quot;Notes\&quot;&gt; -  If you want to store other types of data such as objects or arrays,  you must convert your data into string values. - You cannot override the following system-generated metadata fields:   - &#x60;duration&#x60;   - &#x60;filename&#x60;   - &#x60;fps&#x60;   - &#x60;height&#x60;   - &#x60;model_names&#x60;   - &#x60;size&#x60;   - &#x60;video_title&#x60;   - &#x60;width&#x60; &lt;/Note&gt;  | [optional] 

## Methods

### NewUpdateVideoInformationRequest

`func NewUpdateVideoInformationRequest() *UpdateVideoInformationRequest`

NewUpdateVideoInformationRequest instantiates a new UpdateVideoInformationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateVideoInformationRequestWithDefaults

`func NewUpdateVideoInformationRequestWithDefaults() *UpdateVideoInformationRequest`

NewUpdateVideoInformationRequestWithDefaults instantiates a new UpdateVideoInformationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserMetadata

`func (o *UpdateVideoInformationRequest) GetUserMetadata() map[string]interface{}`

GetUserMetadata returns the UserMetadata field if non-nil, zero value otherwise.

### GetUserMetadataOk

`func (o *UpdateVideoInformationRequest) GetUserMetadataOk() (*map[string]interface{}, bool)`

GetUserMetadataOk returns a tuple with the UserMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserMetadata

`func (o *UpdateVideoInformationRequest) SetUserMetadata(v map[string]interface{})`

SetUserMetadata sets UserMetadata field to given value.

### HasUserMetadata

`func (o *UpdateVideoInformationRequest) HasUserMetadata() bool`

HasUserMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


