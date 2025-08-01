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

// checks if the NonStreamGenerateResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NonStreamGenerateResponse{}

// NonStreamGenerateResponse When the value of the `stream` parameter is set to `false`, the response is as follows: 
type NonStreamGenerateResponse struct {
	// Unique identifier of the response. 
	Id *string `json:"id,omitempty"`
	// The generated text based on the prompt you provided. 
	Data *string `json:"data,omitempty"`
	Usage *TokenUsage `json:"usage,omitempty"`
}

// NewNonStreamGenerateResponse instantiates a new NonStreamGenerateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNonStreamGenerateResponse() *NonStreamGenerateResponse {
	this := NonStreamGenerateResponse{}
	return &this
}

// NewNonStreamGenerateResponseWithDefaults instantiates a new NonStreamGenerateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNonStreamGenerateResponseWithDefaults() *NonStreamGenerateResponse {
	this := NonStreamGenerateResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *NonStreamGenerateResponse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonStreamGenerateResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *NonStreamGenerateResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *NonStreamGenerateResponse) SetId(v string) {
	o.Id = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *NonStreamGenerateResponse) GetData() string {
	if o == nil || IsNil(o.Data) {
		var ret string
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonStreamGenerateResponse) GetDataOk() (*string, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *NonStreamGenerateResponse) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given string and assigns it to the Data field.
func (o *NonStreamGenerateResponse) SetData(v string) {
	o.Data = &v
}

// GetUsage returns the Usage field value if set, zero value otherwise.
func (o *NonStreamGenerateResponse) GetUsage() TokenUsage {
	if o == nil || IsNil(o.Usage) {
		var ret TokenUsage
		return ret
	}
	return *o.Usage
}

// GetUsageOk returns a tuple with the Usage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonStreamGenerateResponse) GetUsageOk() (*TokenUsage, bool) {
	if o == nil || IsNil(o.Usage) {
		return nil, false
	}
	return o.Usage, true
}

// HasUsage returns a boolean if a field has been set.
func (o *NonStreamGenerateResponse) HasUsage() bool {
	if o != nil && !IsNil(o.Usage) {
		return true
	}

	return false
}

// SetUsage gets a reference to the given TokenUsage and assigns it to the Usage field.
func (o *NonStreamGenerateResponse) SetUsage(v TokenUsage) {
	o.Usage = &v
}

func (o NonStreamGenerateResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NonStreamGenerateResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Usage) {
		toSerialize["usage"] = o.Usage
	}
	return toSerialize, nil
}

type NullableNonStreamGenerateResponse struct {
	value *NonStreamGenerateResponse
	isSet bool
}

func (v NullableNonStreamGenerateResponse) Get() *NonStreamGenerateResponse {
	return v.value
}

func (v *NullableNonStreamGenerateResponse) Set(val *NonStreamGenerateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableNonStreamGenerateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableNonStreamGenerateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNonStreamGenerateResponse(val *NonStreamGenerateResponse) *NullableNonStreamGenerateResponse {
	return &NullableNonStreamGenerateResponse{value: val, isSet: true}
}

func (v NullableNonStreamGenerateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNonStreamGenerateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


