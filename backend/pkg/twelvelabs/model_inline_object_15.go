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
	"time"
)

// checks if the InlineObject15 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InlineObject15{}

// InlineObject15 struct for InlineObject15
type InlineObject15 struct {
	// The unique identifier of the video embedding task. 
	Id *string `json:"_id,omitempty"`
	// The name of the video understanding model the platform used to create the embedding. 
	ModelName *string `json:"model_name,omitempty"`
	// A string indicating the status of the video indexing task. It can take one of the following values: `processing`, `ready` or `failed`. 
	Status *string `json:"status,omitempty"`
	// The date and time, in the RFC 3339 format (\"YYYY-MM-DDTHH:mm:ssZ\"), that the video embedding task was created. 
	CreatedAt *time.Time `json:"created_at,omitempty"`
	VideoEmbedding *InlineObject15AllOfVideoEmbedding `json:"video_embedding,omitempty"`
}

// NewInlineObject15 instantiates a new InlineObject15 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject15() *InlineObject15 {
	this := InlineObject15{}
	return &this
}

// NewInlineObject15WithDefaults instantiates a new InlineObject15 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject15WithDefaults() *InlineObject15 {
	this := InlineObject15{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InlineObject15) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject15) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InlineObject15) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InlineObject15) SetId(v string) {
	o.Id = &v
}

// GetModelName returns the ModelName field value if set, zero value otherwise.
func (o *InlineObject15) GetModelName() string {
	if o == nil || IsNil(o.ModelName) {
		var ret string
		return ret
	}
	return *o.ModelName
}

// GetModelNameOk returns a tuple with the ModelName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject15) GetModelNameOk() (*string, bool) {
	if o == nil || IsNil(o.ModelName) {
		return nil, false
	}
	return o.ModelName, true
}

// HasModelName returns a boolean if a field has been set.
func (o *InlineObject15) HasModelName() bool {
	if o != nil && !IsNil(o.ModelName) {
		return true
	}

	return false
}

// SetModelName gets a reference to the given string and assigns it to the ModelName field.
func (o *InlineObject15) SetModelName(v string) {
	o.ModelName = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *InlineObject15) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject15) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *InlineObject15) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *InlineObject15) SetStatus(v string) {
	o.Status = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *InlineObject15) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject15) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *InlineObject15) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *InlineObject15) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetVideoEmbedding returns the VideoEmbedding field value if set, zero value otherwise.
func (o *InlineObject15) GetVideoEmbedding() InlineObject15AllOfVideoEmbedding {
	if o == nil || IsNil(o.VideoEmbedding) {
		var ret InlineObject15AllOfVideoEmbedding
		return ret
	}
	return *o.VideoEmbedding
}

// GetVideoEmbeddingOk returns a tuple with the VideoEmbedding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject15) GetVideoEmbeddingOk() (*InlineObject15AllOfVideoEmbedding, bool) {
	if o == nil || IsNil(o.VideoEmbedding) {
		return nil, false
	}
	return o.VideoEmbedding, true
}

// HasVideoEmbedding returns a boolean if a field has been set.
func (o *InlineObject15) HasVideoEmbedding() bool {
	if o != nil && !IsNil(o.VideoEmbedding) {
		return true
	}

	return false
}

// SetVideoEmbedding gets a reference to the given InlineObject15AllOfVideoEmbedding and assigns it to the VideoEmbedding field.
func (o *InlineObject15) SetVideoEmbedding(v InlineObject15AllOfVideoEmbedding) {
	o.VideoEmbedding = &v
}

func (o InlineObject15) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InlineObject15) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["_id"] = o.Id
	}
	if !IsNil(o.ModelName) {
		toSerialize["model_name"] = o.ModelName
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.VideoEmbedding) {
		toSerialize["video_embedding"] = o.VideoEmbedding
	}
	return toSerialize, nil
}

type NullableInlineObject15 struct {
	value *InlineObject15
	isSet bool
}

func (v NullableInlineObject15) Get() *InlineObject15 {
	return v.value
}

func (v *NullableInlineObject15) Set(val *InlineObject15) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject15) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject15) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject15(val *InlineObject15) *NullableInlineObject15 {
	return &NullableInlineObject15{value: val, isSet: true}
}

func (v NullableInlineObject15) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject15) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


