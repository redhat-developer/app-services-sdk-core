/*
Kafka Instance API

API for interacting with Kafka Instance. Includes Produce, Consume and Admin APIs

API version: 0.13.1-SNAPSHOT
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package kafkainstanceclient

import (
	"encoding/json"
)

// checks if the ErrorAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ErrorAllOf{}

// ErrorAllOf General error response
type ErrorAllOf struct {
	// General reason for the error. Does not change between specific occurrences.
	Reason *string `json:"reason,omitempty"`
	// Detail specific to an error occurrence. May be different depending on the condition(s) that trigger the error.
	Detail *string `json:"detail,omitempty"`
	Code *int32 `json:"code,omitempty"`
	// Deprecated
	ErrorMessage *string `json:"error_message,omitempty"`
	// Deprecated
	Class *string `json:"class,omitempty"`
}

// NewErrorAllOf instantiates a new ErrorAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorAllOf() *ErrorAllOf {
	this := ErrorAllOf{}
	return &this
}

// NewErrorAllOfWithDefaults instantiates a new ErrorAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorAllOfWithDefaults() *ErrorAllOf {
	this := ErrorAllOf{}
	return &this
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *ErrorAllOf) GetReason() string {
	if o == nil || IsNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorAllOf) GetReasonOk() (*string, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *ErrorAllOf) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *ErrorAllOf) SetReason(v string) {
	o.Reason = &v
}

// GetDetail returns the Detail field value if set, zero value otherwise.
func (o *ErrorAllOf) GetDetail() string {
	if o == nil || IsNil(o.Detail) {
		var ret string
		return ret
	}
	return *o.Detail
}

// GetDetailOk returns a tuple with the Detail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorAllOf) GetDetailOk() (*string, bool) {
	if o == nil || IsNil(o.Detail) {
		return nil, false
	}
	return o.Detail, true
}

// HasDetail returns a boolean if a field has been set.
func (o *ErrorAllOf) HasDetail() bool {
	if o != nil && !IsNil(o.Detail) {
		return true
	}

	return false
}

// SetDetail gets a reference to the given string and assigns it to the Detail field.
func (o *ErrorAllOf) SetDetail(v string) {
	o.Detail = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ErrorAllOf) GetCode() int32 {
	if o == nil || IsNil(o.Code) {
		var ret int32
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorAllOf) GetCodeOk() (*int32, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ErrorAllOf) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given int32 and assigns it to the Code field.
func (o *ErrorAllOf) SetCode(v int32) {
	o.Code = &v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
// Deprecated
func (o *ErrorAllOf) GetErrorMessage() string {
	if o == nil || IsNil(o.ErrorMessage) {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *ErrorAllOf) GetErrorMessageOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorMessage) {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *ErrorAllOf) HasErrorMessage() bool {
	if o != nil && !IsNil(o.ErrorMessage) {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
// Deprecated
func (o *ErrorAllOf) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

// GetClass returns the Class field value if set, zero value otherwise.
// Deprecated
func (o *ErrorAllOf) GetClass() string {
	if o == nil || IsNil(o.Class) {
		var ret string
		return ret
	}
	return *o.Class
}

// GetClassOk returns a tuple with the Class field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *ErrorAllOf) GetClassOk() (*string, bool) {
	if o == nil || IsNil(o.Class) {
		return nil, false
	}
	return o.Class, true
}

// HasClass returns a boolean if a field has been set.
func (o *ErrorAllOf) HasClass() bool {
	if o != nil && !IsNil(o.Class) {
		return true
	}

	return false
}

// SetClass gets a reference to the given string and assigns it to the Class field.
// Deprecated
func (o *ErrorAllOf) SetClass(v string) {
	o.Class = &v
}

func (o ErrorAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ErrorAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	if !IsNil(o.Detail) {
		toSerialize["detail"] = o.Detail
	}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.ErrorMessage) {
		toSerialize["error_message"] = o.ErrorMessage
	}
	if !IsNil(o.Class) {
		toSerialize["class"] = o.Class
	}
	return toSerialize, nil
}

type NullableErrorAllOf struct {
	value *ErrorAllOf
	isSet bool
}

func (v NullableErrorAllOf) Get() *ErrorAllOf {
	return v.value
}

func (v *NullableErrorAllOf) Set(val *ErrorAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorAllOf(val *ErrorAllOf) *NullableErrorAllOf {
	return &NullableErrorAllOf{value: val, isSet: true}
}

func (v NullableErrorAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


