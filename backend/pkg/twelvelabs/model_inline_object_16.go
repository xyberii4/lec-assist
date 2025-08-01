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

// checks if the InlineObject16 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InlineObject16{}

// InlineObject16 struct for InlineObject16
type InlineObject16 struct {
	// An array that contains up to `page_limit` video embedding tasks. 
	Data []VideoEmbeddingTask `json:"data,omitempty"`
	PageInfo *InlineObject16PageInfo `json:"page_info,omitempty"`
}

// NewInlineObject16 instantiates a new InlineObject16 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject16() *InlineObject16 {
	this := InlineObject16{}
	return &this
}

// NewInlineObject16WithDefaults instantiates a new InlineObject16 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject16WithDefaults() *InlineObject16 {
	this := InlineObject16{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *InlineObject16) GetData() []VideoEmbeddingTask {
	if o == nil || IsNil(o.Data) {
		var ret []VideoEmbeddingTask
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject16) GetDataOk() ([]VideoEmbeddingTask, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *InlineObject16) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []VideoEmbeddingTask and assigns it to the Data field.
func (o *InlineObject16) SetData(v []VideoEmbeddingTask) {
	o.Data = v
}

// GetPageInfo returns the PageInfo field value if set, zero value otherwise.
func (o *InlineObject16) GetPageInfo() InlineObject16PageInfo {
	if o == nil || IsNil(o.PageInfo) {
		var ret InlineObject16PageInfo
		return ret
	}
	return *o.PageInfo
}

// GetPageInfoOk returns a tuple with the PageInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject16) GetPageInfoOk() (*InlineObject16PageInfo, bool) {
	if o == nil || IsNil(o.PageInfo) {
		return nil, false
	}
	return o.PageInfo, true
}

// HasPageInfo returns a boolean if a field has been set.
func (o *InlineObject16) HasPageInfo() bool {
	if o != nil && !IsNil(o.PageInfo) {
		return true
	}

	return false
}

// SetPageInfo gets a reference to the given InlineObject16PageInfo and assigns it to the PageInfo field.
func (o *InlineObject16) SetPageInfo(v InlineObject16PageInfo) {
	o.PageInfo = &v
}

func (o InlineObject16) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InlineObject16) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.PageInfo) {
		toSerialize["page_info"] = o.PageInfo
	}
	return toSerialize, nil
}

type NullableInlineObject16 struct {
	value *InlineObject16
	isSet bool
}

func (v NullableInlineObject16) Get() *InlineObject16 {
	return v.value
}

func (v *NullableInlineObject16) Set(val *InlineObject16) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject16) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject16) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject16(val *InlineObject16) *NullableInlineObject16 {
	return &NullableInlineObject16{value: val, isSet: true}
}

func (v NullableInlineObject16) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject16) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


