# VideoItemFailed

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filename** | **string** | The name of the video file. | 
**CreatedAt** | **time.Time** | The date and time, in the RFC 3339 format (\&quot;YYYY-MM-DDTHH:mm:ssZ\&quot;), when the video was added to the import process. | 
**ErrorMessage** | **string** | The error message explaining why the video failed to import. | 

## Methods

### NewVideoItemFailed

`func NewVideoItemFailed(filename string, createdAt time.Time, errorMessage string, ) *VideoItemFailed`

NewVideoItemFailed instantiates a new VideoItemFailed object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVideoItemFailedWithDefaults

`func NewVideoItemFailedWithDefaults() *VideoItemFailed`

NewVideoItemFailedWithDefaults instantiates a new VideoItemFailed object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilename

`func (o *VideoItemFailed) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *VideoItemFailed) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *VideoItemFailed) SetFilename(v string)`

SetFilename sets Filename field to given value.


### GetCreatedAt

`func (o *VideoItemFailed) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *VideoItemFailed) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *VideoItemFailed) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetErrorMessage

`func (o *VideoItemFailed) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *VideoItemFailed) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *VideoItemFailed) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


