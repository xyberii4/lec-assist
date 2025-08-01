/*
TwelveLabs Video Understanding API

Use the TwelveLabs Video Understanding API to extract information from your videos and make it available to your applications. The API is organized around REST and returns responses in the JSON format. It is compatible with most programming languages, and you can also use Postman or other REST clients to send requests and view responses. 

API version: 1.3.0
Contact: support@twelvelabs.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package twelvelabs

import (
	"encoding/json"
)

// checks if the Index type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Index{}

// Index An index groups one or more videos stored as vectors and is the most granular level at which you can perform a search. 
type Index struct {
	// A string representing the unique identifier of the index. It is assigned by the API when an index is created.
	Id *string `json:"_id,omitempty"`
	// A string representing the date and time, in the RFC 3339 format (\"YYYY-MM-DDTHH:mm:ssZ\"), that the index was created.
	CreatedAt *string `json:"created_at,omitempty"`
	// A string representing the date and time, in the RFC 3339 format (\"YYYY-MM-DDTHH:mm:ssZ\"), that the index has been updated.
	UpdatedAt *string `json:"updated_at,omitempty"`
	// A string representing the date and time, in the RFC 3339 format (\"YYYY-MM-DDTHH:mm:ssZ\"), when your index will expire.  If you're on the Free plan, the platform retains your index data for 90 days from creation. After this period, the platform deletes your index data, and this action cannot be undone. To continue using your index beyond this period, consider upgrading to the Developer plan, which offers unlimited index retention. For details, see the [Upgrade your plan](/v1.3/docs/get-started/manage-your-plan#upgrade-your-plan) section.  If you're on the Developer plan, this field is set to `null`, indicating no expiration. 
	ExpiresAt *string `json:"expires_at,omitempty"`
	// A string representing the name of the index.
	IndexName *string `json:"index_name,omitempty"`
	// A number representing the total duration, in seconds, of the videos in the index.
	TotalDuration *float32 `json:"total_duration,omitempty"`
	// The number of videos uploaded to this index.
	VideoCount *float32 `json:"video_count,omitempty"`
	// An array containing the list of the [video understanding models](/v1.3/docs/concepts/models) enabled for this index.
	Models []IndexModelsInner `json:"models,omitempty"`
	// The list of the add-ons that are enabled for this index.
	Addons []string `json:"addons,omitempty"`
}

// NewIndex instantiates a new Index object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIndex() *Index {
	this := Index{}
	return &this
}

// NewIndexWithDefaults instantiates a new Index object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIndexWithDefaults() *Index {
	this := Index{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Index) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Index) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Index) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Index) SetId(v string) {
	o.Id = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Index) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Index) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Index) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *Index) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Index) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Index) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Index) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *Index) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *Index) GetExpiresAt() string {
	if o == nil || IsNil(o.ExpiresAt) {
		var ret string
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Index) GetExpiresAtOk() (*string, bool) {
	if o == nil || IsNil(o.ExpiresAt) {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *Index) HasExpiresAt() bool {
	if o != nil && !IsNil(o.ExpiresAt) {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given string and assigns it to the ExpiresAt field.
func (o *Index) SetExpiresAt(v string) {
	o.ExpiresAt = &v
}

// GetIndexName returns the IndexName field value if set, zero value otherwise.
func (o *Index) GetIndexName() string {
	if o == nil || IsNil(o.IndexName) {
		var ret string
		return ret
	}
	return *o.IndexName
}

// GetIndexNameOk returns a tuple with the IndexName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Index) GetIndexNameOk() (*string, bool) {
	if o == nil || IsNil(o.IndexName) {
		return nil, false
	}
	return o.IndexName, true
}

// HasIndexName returns a boolean if a field has been set.
func (o *Index) HasIndexName() bool {
	if o != nil && !IsNil(o.IndexName) {
		return true
	}

	return false
}

// SetIndexName gets a reference to the given string and assigns it to the IndexName field.
func (o *Index) SetIndexName(v string) {
	o.IndexName = &v
}

// GetTotalDuration returns the TotalDuration field value if set, zero value otherwise.
func (o *Index) GetTotalDuration() float32 {
	if o == nil || IsNil(o.TotalDuration) {
		var ret float32
		return ret
	}
	return *o.TotalDuration
}

// GetTotalDurationOk returns a tuple with the TotalDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Index) GetTotalDurationOk() (*float32, bool) {
	if o == nil || IsNil(o.TotalDuration) {
		return nil, false
	}
	return o.TotalDuration, true
}

// HasTotalDuration returns a boolean if a field has been set.
func (o *Index) HasTotalDuration() bool {
	if o != nil && !IsNil(o.TotalDuration) {
		return true
	}

	return false
}

// SetTotalDuration gets a reference to the given float32 and assigns it to the TotalDuration field.
func (o *Index) SetTotalDuration(v float32) {
	o.TotalDuration = &v
}

// GetVideoCount returns the VideoCount field value if set, zero value otherwise.
func (o *Index) GetVideoCount() float32 {
	if o == nil || IsNil(o.VideoCount) {
		var ret float32
		return ret
	}
	return *o.VideoCount
}

// GetVideoCountOk returns a tuple with the VideoCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Index) GetVideoCountOk() (*float32, bool) {
	if o == nil || IsNil(o.VideoCount) {
		return nil, false
	}
	return o.VideoCount, true
}

// HasVideoCount returns a boolean if a field has been set.
func (o *Index) HasVideoCount() bool {
	if o != nil && !IsNil(o.VideoCount) {
		return true
	}

	return false
}

// SetVideoCount gets a reference to the given float32 and assigns it to the VideoCount field.
func (o *Index) SetVideoCount(v float32) {
	o.VideoCount = &v
}

// GetModels returns the Models field value if set, zero value otherwise.
func (o *Index) GetModels() []IndexModelsInner {
	if o == nil || IsNil(o.Models) {
		var ret []IndexModelsInner
		return ret
	}
	return o.Models
}

// GetModelsOk returns a tuple with the Models field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Index) GetModelsOk() ([]IndexModelsInner, bool) {
	if o == nil || IsNil(o.Models) {
		return nil, false
	}
	return o.Models, true
}

// HasModels returns a boolean if a field has been set.
func (o *Index) HasModels() bool {
	if o != nil && !IsNil(o.Models) {
		return true
	}

	return false
}

// SetModels gets a reference to the given []IndexModelsInner and assigns it to the Models field.
func (o *Index) SetModels(v []IndexModelsInner) {
	o.Models = v
}

// GetAddons returns the Addons field value if set, zero value otherwise.
func (o *Index) GetAddons() []string {
	if o == nil || IsNil(o.Addons) {
		var ret []string
		return ret
	}
	return o.Addons
}

// GetAddonsOk returns a tuple with the Addons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Index) GetAddonsOk() ([]string, bool) {
	if o == nil || IsNil(o.Addons) {
		return nil, false
	}
	return o.Addons, true
}

// HasAddons returns a boolean if a field has been set.
func (o *Index) HasAddons() bool {
	if o != nil && !IsNil(o.Addons) {
		return true
	}

	return false
}

// SetAddons gets a reference to the given []string and assigns it to the Addons field.
func (o *Index) SetAddons(v []string) {
	o.Addons = v
}

func (o Index) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Index) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["_id"] = o.Id
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !IsNil(o.ExpiresAt) {
		toSerialize["expires_at"] = o.ExpiresAt
	}
	if !IsNil(o.IndexName) {
		toSerialize["index_name"] = o.IndexName
	}
	if !IsNil(o.TotalDuration) {
		toSerialize["total_duration"] = o.TotalDuration
	}
	if !IsNil(o.VideoCount) {
		toSerialize["video_count"] = o.VideoCount
	}
	if !IsNil(o.Models) {
		toSerialize["models"] = o.Models
	}
	if !IsNil(o.Addons) {
		toSerialize["addons"] = o.Addons
	}
	return toSerialize, nil
}

type NullableIndex struct {
	value *Index
	isSet bool
}

func (v NullableIndex) Get() *Index {
	return v.value
}

func (v *NullableIndex) Set(val *Index) {
	v.value = val
	v.isSet = true
}

func (v NullableIndex) IsSet() bool {
	return v.isSet
}

func (v *NullableIndex) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIndex(val *Index) *NullableIndex {
	return &NullableIndex{value: val, isSet: true}
}

func (v NullableIndex) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIndex) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


