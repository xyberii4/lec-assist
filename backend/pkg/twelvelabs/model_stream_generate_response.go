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
	"fmt"
	"gopkg.in/validator.v2"
)

// StreamGenerateResponse - When the value of the `stream` parameter is set to `true`, the platform provides a streaming response in the NDJSON format.  The stream contains three types of events: 1. Stream start 2. Text generation 3. Stream end  To integrate the response into your application, follow the guidelines below: - Parse each line of the response as a separate JSON object. - Check the `event_type` field to determine how to handle the event. - For `text_generation` events, process the `text` field as it arrives. Depending on your application's requirements, this may involve displaying the text incrementally, storing it for later use, or performing any tasks. - Use the `stream_start` and `stream_end` events to manage the lifecycle of your streaming session. 
type StreamGenerateResponse struct {
	StreamEndResponse *StreamEndResponse
	StreamStartResponse *StreamStartResponse
	StreamTextResponse *StreamTextResponse
}

// StreamEndResponseAsStreamGenerateResponse is a convenience function that returns StreamEndResponse wrapped in StreamGenerateResponse
func StreamEndResponseAsStreamGenerateResponse(v *StreamEndResponse) StreamGenerateResponse {
	return StreamGenerateResponse{
		StreamEndResponse: v,
	}
}

// StreamStartResponseAsStreamGenerateResponse is a convenience function that returns StreamStartResponse wrapped in StreamGenerateResponse
func StreamStartResponseAsStreamGenerateResponse(v *StreamStartResponse) StreamGenerateResponse {
	return StreamGenerateResponse{
		StreamStartResponse: v,
	}
}

// StreamTextResponseAsStreamGenerateResponse is a convenience function that returns StreamTextResponse wrapped in StreamGenerateResponse
func StreamTextResponseAsStreamGenerateResponse(v *StreamTextResponse) StreamGenerateResponse {
	return StreamGenerateResponse{
		StreamTextResponse: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *StreamGenerateResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into StreamEndResponse
	err = newStrictDecoder(data).Decode(&dst.StreamEndResponse)
	if err == nil {
		jsonStreamEndResponse, _ := json.Marshal(dst.StreamEndResponse)
		if string(jsonStreamEndResponse) == "{}" { // empty struct
			dst.StreamEndResponse = nil
		} else {
			if err = validator.Validate(dst.StreamEndResponse); err != nil {
				dst.StreamEndResponse = nil
			} else {
				match++
			}
		}
	} else {
		dst.StreamEndResponse = nil
	}

	// try to unmarshal data into StreamStartResponse
	err = newStrictDecoder(data).Decode(&dst.StreamStartResponse)
	if err == nil {
		jsonStreamStartResponse, _ := json.Marshal(dst.StreamStartResponse)
		if string(jsonStreamStartResponse) == "{}" { // empty struct
			dst.StreamStartResponse = nil
		} else {
			if err = validator.Validate(dst.StreamStartResponse); err != nil {
				dst.StreamStartResponse = nil
			} else {
				match++
			}
		}
	} else {
		dst.StreamStartResponse = nil
	}

	// try to unmarshal data into StreamTextResponse
	err = newStrictDecoder(data).Decode(&dst.StreamTextResponse)
	if err == nil {
		jsonStreamTextResponse, _ := json.Marshal(dst.StreamTextResponse)
		if string(jsonStreamTextResponse) == "{}" { // empty struct
			dst.StreamTextResponse = nil
		} else {
			if err = validator.Validate(dst.StreamTextResponse); err != nil {
				dst.StreamTextResponse = nil
			} else {
				match++
			}
		}
	} else {
		dst.StreamTextResponse = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.StreamEndResponse = nil
		dst.StreamStartResponse = nil
		dst.StreamTextResponse = nil

		return fmt.Errorf("data matches more than one schema in oneOf(StreamGenerateResponse)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(StreamGenerateResponse)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src StreamGenerateResponse) MarshalJSON() ([]byte, error) {
	if src.StreamEndResponse != nil {
		return json.Marshal(&src.StreamEndResponse)
	}

	if src.StreamStartResponse != nil {
		return json.Marshal(&src.StreamStartResponse)
	}

	if src.StreamTextResponse != nil {
		return json.Marshal(&src.StreamTextResponse)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *StreamGenerateResponse) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.StreamEndResponse != nil {
		return obj.StreamEndResponse
	}

	if obj.StreamStartResponse != nil {
		return obj.StreamStartResponse
	}

	if obj.StreamTextResponse != nil {
		return obj.StreamTextResponse
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj StreamGenerateResponse) GetActualInstanceValue() (interface{}) {
	if obj.StreamEndResponse != nil {
		return *obj.StreamEndResponse
	}

	if obj.StreamStartResponse != nil {
		return *obj.StreamStartResponse
	}

	if obj.StreamTextResponse != nil {
		return *obj.StreamTextResponse
	}

	// all schemas are nil
	return nil
}

type NullableStreamGenerateResponse struct {
	value *StreamGenerateResponse
	isSet bool
}

func (v NullableStreamGenerateResponse) Get() *StreamGenerateResponse {
	return v.value
}

func (v *NullableStreamGenerateResponse) Set(val *StreamGenerateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableStreamGenerateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableStreamGenerateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStreamGenerateResponse(val *StreamGenerateResponse) *NullableStreamGenerateResponse {
	return &NullableStreamGenerateResponse{value: val, isSet: true}
}

func (v NullableStreamGenerateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStreamGenerateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


