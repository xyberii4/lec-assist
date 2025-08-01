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

// checks if the InlineObject21 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InlineObject21{}

// InlineObject21 struct for InlineObject21
type InlineObject21 struct {
	// An array of filenames that haven't yet been imported.
	NotImported []string `json:"not_imported,omitempty"`
	// An array of videos that are being validated.
	Validating []VideoItem `json:"validating,omitempty"`
	// An array of videos that are pending.
	Pending []VideoItem `json:"pending,omitempty"`
	// An array of videos that are queued for import.
	Queued []VideoItem `json:"queued,omitempty"`
	// An array of videos that are being indexed.
	Indexing []VideoItem `json:"indexing,omitempty"`
	// An array of videos that have successfully been imported.
	Ready []VideoItem `json:"ready,omitempty"`
	// An array of videos that failed to import.
	Failed []VideoItemFailed `json:"failed,omitempty"`
}

// NewInlineObject21 instantiates a new InlineObject21 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject21() *InlineObject21 {
	this := InlineObject21{}
	return &this
}

// NewInlineObject21WithDefaults instantiates a new InlineObject21 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject21WithDefaults() *InlineObject21 {
	this := InlineObject21{}
	return &this
}

// GetNotImported returns the NotImported field value if set, zero value otherwise.
func (o *InlineObject21) GetNotImported() []string {
	if o == nil || IsNil(o.NotImported) {
		var ret []string
		return ret
	}
	return o.NotImported
}

// GetNotImportedOk returns a tuple with the NotImported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject21) GetNotImportedOk() ([]string, bool) {
	if o == nil || IsNil(o.NotImported) {
		return nil, false
	}
	return o.NotImported, true
}

// HasNotImported returns a boolean if a field has been set.
func (o *InlineObject21) HasNotImported() bool {
	if o != nil && !IsNil(o.NotImported) {
		return true
	}

	return false
}

// SetNotImported gets a reference to the given []string and assigns it to the NotImported field.
func (o *InlineObject21) SetNotImported(v []string) {
	o.NotImported = v
}

// GetValidating returns the Validating field value if set, zero value otherwise.
func (o *InlineObject21) GetValidating() []VideoItem {
	if o == nil || IsNil(o.Validating) {
		var ret []VideoItem
		return ret
	}
	return o.Validating
}

// GetValidatingOk returns a tuple with the Validating field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject21) GetValidatingOk() ([]VideoItem, bool) {
	if o == nil || IsNil(o.Validating) {
		return nil, false
	}
	return o.Validating, true
}

// HasValidating returns a boolean if a field has been set.
func (o *InlineObject21) HasValidating() bool {
	if o != nil && !IsNil(o.Validating) {
		return true
	}

	return false
}

// SetValidating gets a reference to the given []VideoItem and assigns it to the Validating field.
func (o *InlineObject21) SetValidating(v []VideoItem) {
	o.Validating = v
}

// GetPending returns the Pending field value if set, zero value otherwise.
func (o *InlineObject21) GetPending() []VideoItem {
	if o == nil || IsNil(o.Pending) {
		var ret []VideoItem
		return ret
	}
	return o.Pending
}

// GetPendingOk returns a tuple with the Pending field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject21) GetPendingOk() ([]VideoItem, bool) {
	if o == nil || IsNil(o.Pending) {
		return nil, false
	}
	return o.Pending, true
}

// HasPending returns a boolean if a field has been set.
func (o *InlineObject21) HasPending() bool {
	if o != nil && !IsNil(o.Pending) {
		return true
	}

	return false
}

// SetPending gets a reference to the given []VideoItem and assigns it to the Pending field.
func (o *InlineObject21) SetPending(v []VideoItem) {
	o.Pending = v
}

// GetQueued returns the Queued field value if set, zero value otherwise.
func (o *InlineObject21) GetQueued() []VideoItem {
	if o == nil || IsNil(o.Queued) {
		var ret []VideoItem
		return ret
	}
	return o.Queued
}

// GetQueuedOk returns a tuple with the Queued field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject21) GetQueuedOk() ([]VideoItem, bool) {
	if o == nil || IsNil(o.Queued) {
		return nil, false
	}
	return o.Queued, true
}

// HasQueued returns a boolean if a field has been set.
func (o *InlineObject21) HasQueued() bool {
	if o != nil && !IsNil(o.Queued) {
		return true
	}

	return false
}

// SetQueued gets a reference to the given []VideoItem and assigns it to the Queued field.
func (o *InlineObject21) SetQueued(v []VideoItem) {
	o.Queued = v
}

// GetIndexing returns the Indexing field value if set, zero value otherwise.
func (o *InlineObject21) GetIndexing() []VideoItem {
	if o == nil || IsNil(o.Indexing) {
		var ret []VideoItem
		return ret
	}
	return o.Indexing
}

// GetIndexingOk returns a tuple with the Indexing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject21) GetIndexingOk() ([]VideoItem, bool) {
	if o == nil || IsNil(o.Indexing) {
		return nil, false
	}
	return o.Indexing, true
}

// HasIndexing returns a boolean if a field has been set.
func (o *InlineObject21) HasIndexing() bool {
	if o != nil && !IsNil(o.Indexing) {
		return true
	}

	return false
}

// SetIndexing gets a reference to the given []VideoItem and assigns it to the Indexing field.
func (o *InlineObject21) SetIndexing(v []VideoItem) {
	o.Indexing = v
}

// GetReady returns the Ready field value if set, zero value otherwise.
func (o *InlineObject21) GetReady() []VideoItem {
	if o == nil || IsNil(o.Ready) {
		var ret []VideoItem
		return ret
	}
	return o.Ready
}

// GetReadyOk returns a tuple with the Ready field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject21) GetReadyOk() ([]VideoItem, bool) {
	if o == nil || IsNil(o.Ready) {
		return nil, false
	}
	return o.Ready, true
}

// HasReady returns a boolean if a field has been set.
func (o *InlineObject21) HasReady() bool {
	if o != nil && !IsNil(o.Ready) {
		return true
	}

	return false
}

// SetReady gets a reference to the given []VideoItem and assigns it to the Ready field.
func (o *InlineObject21) SetReady(v []VideoItem) {
	o.Ready = v
}

// GetFailed returns the Failed field value if set, zero value otherwise.
func (o *InlineObject21) GetFailed() []VideoItemFailed {
	if o == nil || IsNil(o.Failed) {
		var ret []VideoItemFailed
		return ret
	}
	return o.Failed
}

// GetFailedOk returns a tuple with the Failed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject21) GetFailedOk() ([]VideoItemFailed, bool) {
	if o == nil || IsNil(o.Failed) {
		return nil, false
	}
	return o.Failed, true
}

// HasFailed returns a boolean if a field has been set.
func (o *InlineObject21) HasFailed() bool {
	if o != nil && !IsNil(o.Failed) {
		return true
	}

	return false
}

// SetFailed gets a reference to the given []VideoItemFailed and assigns it to the Failed field.
func (o *InlineObject21) SetFailed(v []VideoItemFailed) {
	o.Failed = v
}

func (o InlineObject21) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InlineObject21) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NotImported) {
		toSerialize["not_imported"] = o.NotImported
	}
	if !IsNil(o.Validating) {
		toSerialize["validating"] = o.Validating
	}
	if !IsNil(o.Pending) {
		toSerialize["pending"] = o.Pending
	}
	if !IsNil(o.Queued) {
		toSerialize["queued"] = o.Queued
	}
	if !IsNil(o.Indexing) {
		toSerialize["indexing"] = o.Indexing
	}
	if !IsNil(o.Ready) {
		toSerialize["ready"] = o.Ready
	}
	if !IsNil(o.Failed) {
		toSerialize["failed"] = o.Failed
	}
	return toSerialize, nil
}

type NullableInlineObject21 struct {
	value *InlineObject21
	isSet bool
}

func (v NullableInlineObject21) Get() *InlineObject21 {
	return v.value
}

func (v *NullableInlineObject21) Set(val *InlineObject21) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject21) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject21) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject21(val *InlineObject21) *NullableInlineObject21 {
	return &NullableInlineObject21{value: val, isSet: true}
}

func (v NullableInlineObject21) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject21) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


