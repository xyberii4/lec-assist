# Gist

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique identifier of the response.  | [optional] 
**Title** | Pointer to **string** | Suggested title for the video.  | [optional] 
**Topics** | Pointer to **[]string** | An array of topics that are relevant to the video.  | [optional] 
**Hashtags** | Pointer to **[]string** | An array of hashtags that are relevant to the video.  | [optional] 
**Usage** | Pointer to [**TokenUsage**](TokenUsage.md) |  | [optional] 

## Methods

### NewGist

`func NewGist() *Gist`

NewGist instantiates a new Gist object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGistWithDefaults

`func NewGistWithDefaults() *Gist`

NewGistWithDefaults instantiates a new Gist object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Gist) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Gist) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Gist) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Gist) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTitle

`func (o *Gist) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Gist) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Gist) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Gist) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetTopics

`func (o *Gist) GetTopics() []string`

GetTopics returns the Topics field if non-nil, zero value otherwise.

### GetTopicsOk

`func (o *Gist) GetTopicsOk() (*[]string, bool)`

GetTopicsOk returns a tuple with the Topics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopics

`func (o *Gist) SetTopics(v []string)`

SetTopics sets Topics field to given value.

### HasTopics

`func (o *Gist) HasTopics() bool`

HasTopics returns a boolean if a field has been set.

### GetHashtags

`func (o *Gist) GetHashtags() []string`

GetHashtags returns the Hashtags field if non-nil, zero value otherwise.

### GetHashtagsOk

`func (o *Gist) GetHashtagsOk() (*[]string, bool)`

GetHashtagsOk returns a tuple with the Hashtags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashtags

`func (o *Gist) SetHashtags(v []string)`

SetHashtags sets Hashtags field to given value.

### HasHashtags

`func (o *Gist) HasHashtags() bool`

HasHashtags returns a boolean if a field has been set.

### GetUsage

`func (o *Gist) GetUsage() TokenUsage`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *Gist) GetUsageOk() (*TokenUsage, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *Gist) SetUsage(v TokenUsage)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *Gist) HasUsage() bool`

HasUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


