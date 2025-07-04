# Index

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | A string representing the unique identifier of the index. It is assigned by the API when an index is created. | [optional] 
**CreatedAt** | Pointer to **string** | A string representing the date and time, in the RFC 3339 format (\&quot;YYYY-MM-DDTHH:mm:ssZ\&quot;), that the index was created. | [optional] 
**UpdatedAt** | Pointer to **string** | A string representing the date and time, in the RFC 3339 format (\&quot;YYYY-MM-DDTHH:mm:ssZ\&quot;), that the index has been updated. | [optional] 
**ExpiresAt** | Pointer to **string** | A string representing the date and time, in the RFC 3339 format (\&quot;YYYY-MM-DDTHH:mm:ssZ\&quot;), when your index will expire.  If you&#39;re on the Free plan, the platform retains your index data for 90 days from creation. After this period, the platform deletes your index data, and this action cannot be undone. To continue using your index beyond this period, consider upgrading to the Developer plan, which offers unlimited index retention. For details, see the [Upgrade your plan](/v1.3/docs/get-started/manage-your-plan#upgrade-your-plan) section.  If you&#39;re on the Developer plan, this field is set to &#x60;null&#x60;, indicating no expiration.  | [optional] 
**IndexName** | Pointer to **string** | A string representing the name of the index. | [optional] 
**TotalDuration** | Pointer to **float32** | A number representing the total duration, in seconds, of the videos in the index. | [optional] 
**VideoCount** | Pointer to **float32** | The number of videos uploaded to this index. | [optional] 
**Models** | Pointer to [**[]IndexModelsInner**](IndexModelsInner.md) | An array containing the list of the [video understanding models](/v1.3/docs/concepts/models) enabled for this index. | [optional] 
**Addons** | Pointer to **[]string** | The list of the add-ons that are enabled for this index. | [optional] 

## Methods

### NewIndex

`func NewIndex() *Index`

NewIndex instantiates a new Index object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIndexWithDefaults

`func NewIndexWithDefaults() *Index`

NewIndexWithDefaults instantiates a new Index object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Index) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Index) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Index) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Index) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Index) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Index) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Index) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Index) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Index) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Index) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Index) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Index) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetExpiresAt

`func (o *Index) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *Index) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *Index) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *Index) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetIndexName

`func (o *Index) GetIndexName() string`

GetIndexName returns the IndexName field if non-nil, zero value otherwise.

### GetIndexNameOk

`func (o *Index) GetIndexNameOk() (*string, bool)`

GetIndexNameOk returns a tuple with the IndexName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexName

`func (o *Index) SetIndexName(v string)`

SetIndexName sets IndexName field to given value.

### HasIndexName

`func (o *Index) HasIndexName() bool`

HasIndexName returns a boolean if a field has been set.

### GetTotalDuration

`func (o *Index) GetTotalDuration() float32`

GetTotalDuration returns the TotalDuration field if non-nil, zero value otherwise.

### GetTotalDurationOk

`func (o *Index) GetTotalDurationOk() (*float32, bool)`

GetTotalDurationOk returns a tuple with the TotalDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDuration

`func (o *Index) SetTotalDuration(v float32)`

SetTotalDuration sets TotalDuration field to given value.

### HasTotalDuration

`func (o *Index) HasTotalDuration() bool`

HasTotalDuration returns a boolean if a field has been set.

### GetVideoCount

`func (o *Index) GetVideoCount() float32`

GetVideoCount returns the VideoCount field if non-nil, zero value otherwise.

### GetVideoCountOk

`func (o *Index) GetVideoCountOk() (*float32, bool)`

GetVideoCountOk returns a tuple with the VideoCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoCount

`func (o *Index) SetVideoCount(v float32)`

SetVideoCount sets VideoCount field to given value.

### HasVideoCount

`func (o *Index) HasVideoCount() bool`

HasVideoCount returns a boolean if a field has been set.

### GetModels

`func (o *Index) GetModels() []IndexModelsInner`

GetModels returns the Models field if non-nil, zero value otherwise.

### GetModelsOk

`func (o *Index) GetModelsOk() (*[]IndexModelsInner, bool)`

GetModelsOk returns a tuple with the Models field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModels

`func (o *Index) SetModels(v []IndexModelsInner)`

SetModels sets Models field to given value.

### HasModels

`func (o *Index) HasModels() bool`

HasModels returns a boolean if a field has been set.

### GetAddons

`func (o *Index) GetAddons() []string`

GetAddons returns the Addons field if non-nil, zero value otherwise.

### GetAddonsOk

`func (o *Index) GetAddonsOk() (*[]string, bool)`

GetAddonsOk returns a tuple with the Addons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddons

`func (o *Index) SetAddons(v []string)`

SetAddons sets Addons field to given value.

### HasAddons

`func (o *Index) HasAddons() bool`

HasAddons returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


