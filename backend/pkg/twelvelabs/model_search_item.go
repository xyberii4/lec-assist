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

// checks if the SearchItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchItem{}

// SearchItem An object that contains the search results. 
type SearchItem struct {
	// A quantitative value determined by the AI model representing the level of confidence that the results match your search terms. 
	Score *float32 `json:"score,omitempty"`
	// The start time of the matching video clip, expressed in seconds. 
	Start *float32 `json:"start,omitempty"`
	// The end time of the matching video clip, expressed in seconds. 
	End *float32 `json:"end,omitempty"`
	// A string representing the unique identifier of the video. Once the platform indexes a video, it assigns a unique identifier. Note that this is different from the identifier of the video indexing task.
	VideoId *string `json:"video_id,omitempty"`
	// A qualitative indicator based on the value of the `score` field. This field can take one of the following values: - `high` - `medium` - `low` - `extremely low` 
	Confidence *string `json:"confidence,omitempty"`
	// If thumbnail generation has been enabled for this index, the platform returns a string representing the URL of the thumbnail. Note that the URL expires in one hour. 
	ThumbnailUrl *string `json:"thumbnail_url,omitempty"`
	// A transcription of the spoken words that are captured in the video. Note that the official SDKs will support this feature in a future release.
	Transcription *string `json:"transcription,omitempty"`
	// A string representing the unique identifier of the video. It only appears when the `group_by=video` parameter is used in the request.
	Id *string `json:"id,omitempty"`
	// An array that contains detailed information about the clips that match your query. The platform returns this array only when the `group_by` parameter is set to `video` in the request.
	Clips []SearchItemClipsInner `json:"clips,omitempty"`
}

// NewSearchItem instantiates a new SearchItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchItem() *SearchItem {
	this := SearchItem{}
	return &this
}

// NewSearchItemWithDefaults instantiates a new SearchItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchItemWithDefaults() *SearchItem {
	this := SearchItem{}
	return &this
}

// GetScore returns the Score field value if set, zero value otherwise.
func (o *SearchItem) GetScore() float32 {
	if o == nil || IsNil(o.Score) {
		var ret float32
		return ret
	}
	return *o.Score
}

// GetScoreOk returns a tuple with the Score field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchItem) GetScoreOk() (*float32, bool) {
	if o == nil || IsNil(o.Score) {
		return nil, false
	}
	return o.Score, true
}

// HasScore returns a boolean if a field has been set.
func (o *SearchItem) HasScore() bool {
	if o != nil && !IsNil(o.Score) {
		return true
	}

	return false
}

// SetScore gets a reference to the given float32 and assigns it to the Score field.
func (o *SearchItem) SetScore(v float32) {
	o.Score = &v
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *SearchItem) GetStart() float32 {
	if o == nil || IsNil(o.Start) {
		var ret float32
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchItem) GetStartOk() (*float32, bool) {
	if o == nil || IsNil(o.Start) {
		return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *SearchItem) HasStart() bool {
	if o != nil && !IsNil(o.Start) {
		return true
	}

	return false
}

// SetStart gets a reference to the given float32 and assigns it to the Start field.
func (o *SearchItem) SetStart(v float32) {
	o.Start = &v
}

// GetEnd returns the End field value if set, zero value otherwise.
func (o *SearchItem) GetEnd() float32 {
	if o == nil || IsNil(o.End) {
		var ret float32
		return ret
	}
	return *o.End
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchItem) GetEndOk() (*float32, bool) {
	if o == nil || IsNil(o.End) {
		return nil, false
	}
	return o.End, true
}

// HasEnd returns a boolean if a field has been set.
func (o *SearchItem) HasEnd() bool {
	if o != nil && !IsNil(o.End) {
		return true
	}

	return false
}

// SetEnd gets a reference to the given float32 and assigns it to the End field.
func (o *SearchItem) SetEnd(v float32) {
	o.End = &v
}

// GetVideoId returns the VideoId field value if set, zero value otherwise.
func (o *SearchItem) GetVideoId() string {
	if o == nil || IsNil(o.VideoId) {
		var ret string
		return ret
	}
	return *o.VideoId
}

// GetVideoIdOk returns a tuple with the VideoId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchItem) GetVideoIdOk() (*string, bool) {
	if o == nil || IsNil(o.VideoId) {
		return nil, false
	}
	return o.VideoId, true
}

// HasVideoId returns a boolean if a field has been set.
func (o *SearchItem) HasVideoId() bool {
	if o != nil && !IsNil(o.VideoId) {
		return true
	}

	return false
}

// SetVideoId gets a reference to the given string and assigns it to the VideoId field.
func (o *SearchItem) SetVideoId(v string) {
	o.VideoId = &v
}

// GetConfidence returns the Confidence field value if set, zero value otherwise.
func (o *SearchItem) GetConfidence() string {
	if o == nil || IsNil(o.Confidence) {
		var ret string
		return ret
	}
	return *o.Confidence
}

// GetConfidenceOk returns a tuple with the Confidence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchItem) GetConfidenceOk() (*string, bool) {
	if o == nil || IsNil(o.Confidence) {
		return nil, false
	}
	return o.Confidence, true
}

// HasConfidence returns a boolean if a field has been set.
func (o *SearchItem) HasConfidence() bool {
	if o != nil && !IsNil(o.Confidence) {
		return true
	}

	return false
}

// SetConfidence gets a reference to the given string and assigns it to the Confidence field.
func (o *SearchItem) SetConfidence(v string) {
	o.Confidence = &v
}

// GetThumbnailUrl returns the ThumbnailUrl field value if set, zero value otherwise.
func (o *SearchItem) GetThumbnailUrl() string {
	if o == nil || IsNil(o.ThumbnailUrl) {
		var ret string
		return ret
	}
	return *o.ThumbnailUrl
}

// GetThumbnailUrlOk returns a tuple with the ThumbnailUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchItem) GetThumbnailUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ThumbnailUrl) {
		return nil, false
	}
	return o.ThumbnailUrl, true
}

// HasThumbnailUrl returns a boolean if a field has been set.
func (o *SearchItem) HasThumbnailUrl() bool {
	if o != nil && !IsNil(o.ThumbnailUrl) {
		return true
	}

	return false
}

// SetThumbnailUrl gets a reference to the given string and assigns it to the ThumbnailUrl field.
func (o *SearchItem) SetThumbnailUrl(v string) {
	o.ThumbnailUrl = &v
}

// GetTranscription returns the Transcription field value if set, zero value otherwise.
func (o *SearchItem) GetTranscription() string {
	if o == nil || IsNil(o.Transcription) {
		var ret string
		return ret
	}
	return *o.Transcription
}

// GetTranscriptionOk returns a tuple with the Transcription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchItem) GetTranscriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Transcription) {
		return nil, false
	}
	return o.Transcription, true
}

// HasTranscription returns a boolean if a field has been set.
func (o *SearchItem) HasTranscription() bool {
	if o != nil && !IsNil(o.Transcription) {
		return true
	}

	return false
}

// SetTranscription gets a reference to the given string and assigns it to the Transcription field.
func (o *SearchItem) SetTranscription(v string) {
	o.Transcription = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SearchItem) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchItem) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SearchItem) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SearchItem) SetId(v string) {
	o.Id = &v
}

// GetClips returns the Clips field value if set, zero value otherwise.
func (o *SearchItem) GetClips() []SearchItemClipsInner {
	if o == nil || IsNil(o.Clips) {
		var ret []SearchItemClipsInner
		return ret
	}
	return o.Clips
}

// GetClipsOk returns a tuple with the Clips field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchItem) GetClipsOk() ([]SearchItemClipsInner, bool) {
	if o == nil || IsNil(o.Clips) {
		return nil, false
	}
	return o.Clips, true
}

// HasClips returns a boolean if a field has been set.
func (o *SearchItem) HasClips() bool {
	if o != nil && !IsNil(o.Clips) {
		return true
	}

	return false
}

// SetClips gets a reference to the given []SearchItemClipsInner and assigns it to the Clips field.
func (o *SearchItem) SetClips(v []SearchItemClipsInner) {
	o.Clips = v
}

func (o SearchItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Score) {
		toSerialize["score"] = o.Score
	}
	if !IsNil(o.Start) {
		toSerialize["start"] = o.Start
	}
	if !IsNil(o.End) {
		toSerialize["end"] = o.End
	}
	if !IsNil(o.VideoId) {
		toSerialize["video_id"] = o.VideoId
	}
	if !IsNil(o.Confidence) {
		toSerialize["confidence"] = o.Confidence
	}
	if !IsNil(o.ThumbnailUrl) {
		toSerialize["thumbnail_url"] = o.ThumbnailUrl
	}
	if !IsNil(o.Transcription) {
		toSerialize["transcription"] = o.Transcription
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Clips) {
		toSerialize["clips"] = o.Clips
	}
	return toSerialize, nil
}

type NullableSearchItem struct {
	value *SearchItem
	isSet bool
}

func (v NullableSearchItem) Get() *SearchItem {
	return v.value
}

func (v *NullableSearchItem) Set(val *SearchItem) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchItem) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchItem(val *SearchItem) *NullableSearchItem {
	return &NullableSearchItem{value: val, isSet: true}
}

func (v NullableSearchItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


