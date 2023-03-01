/*
sso.redhat.com API documentation

This is the API documentation for sso.redhat.com

API version: 5.0.19-SNAPSHOT
Contact: it-user-team-list@redhat.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package serviceaccountsclient

import (
	"encoding/json"
)

// checks if the RedHatErrorRepresentation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RedHatErrorRepresentation{}

// RedHatErrorRepresentation struct for RedHatErrorRepresentation
type RedHatErrorRepresentation struct {
	Error *string `json:"error,omitempty"`
	ErrorDescription *string `json:"error_description,omitempty"`
}

// NewRedHatErrorRepresentation instantiates a new RedHatErrorRepresentation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRedHatErrorRepresentation() *RedHatErrorRepresentation {
	this := RedHatErrorRepresentation{}
	return &this
}

// NewRedHatErrorRepresentationWithDefaults instantiates a new RedHatErrorRepresentation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRedHatErrorRepresentationWithDefaults() *RedHatErrorRepresentation {
	this := RedHatErrorRepresentation{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *RedHatErrorRepresentation) GetError() string {
	if o == nil || IsNil(o.Error) {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedHatErrorRepresentation) GetErrorOk() (*string, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *RedHatErrorRepresentation) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *RedHatErrorRepresentation) SetError(v string) {
	o.Error = &v
}

// GetErrorDescription returns the ErrorDescription field value if set, zero value otherwise.
func (o *RedHatErrorRepresentation) GetErrorDescription() string {
	if o == nil || IsNil(o.ErrorDescription) {
		var ret string
		return ret
	}
	return *o.ErrorDescription
}

// GetErrorDescriptionOk returns a tuple with the ErrorDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedHatErrorRepresentation) GetErrorDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorDescription) {
		return nil, false
	}
	return o.ErrorDescription, true
}

// HasErrorDescription returns a boolean if a field has been set.
func (o *RedHatErrorRepresentation) HasErrorDescription() bool {
	if o != nil && !IsNil(o.ErrorDescription) {
		return true
	}

	return false
}

// SetErrorDescription gets a reference to the given string and assigns it to the ErrorDescription field.
func (o *RedHatErrorRepresentation) SetErrorDescription(v string) {
	o.ErrorDescription = &v
}

func (o RedHatErrorRepresentation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RedHatErrorRepresentation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	if !IsNil(o.ErrorDescription) {
		toSerialize["error_description"] = o.ErrorDescription
	}
	return toSerialize, nil
}

type NullableRedHatErrorRepresentation struct {
	value *RedHatErrorRepresentation
	isSet bool
}

func (v NullableRedHatErrorRepresentation) Get() *RedHatErrorRepresentation {
	return v.value
}

func (v *NullableRedHatErrorRepresentation) Set(val *RedHatErrorRepresentation) {
	v.value = val
	v.isSet = true
}

func (v NullableRedHatErrorRepresentation) IsSet() bool {
	return v.isSet
}

func (v *NullableRedHatErrorRepresentation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRedHatErrorRepresentation(val *RedHatErrorRepresentation) *NullableRedHatErrorRepresentation {
	return &NullableRedHatErrorRepresentation{value: val, isSet: true}
}

func (v NullableRedHatErrorRepresentation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRedHatErrorRepresentation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


