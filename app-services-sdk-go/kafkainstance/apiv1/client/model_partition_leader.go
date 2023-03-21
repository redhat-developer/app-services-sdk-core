/*
Kafka Instance API

API for interacting with Kafka Instance. Includes Produce, Consume and Admin APIs

API version: 0.14.1-SNAPSHOT
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package kafkainstanceclient

import (
	"encoding/json"
)

// checks if the PartitionLeader type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PartitionLeader{}

// PartitionLeader struct for PartitionLeader
type PartitionLeader struct {
	Id *int32 `json:"id,omitempty"`
}

// NewPartitionLeader instantiates a new PartitionLeader object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPartitionLeader() *PartitionLeader {
	this := PartitionLeader{}
	return &this
}

// NewPartitionLeaderWithDefaults instantiates a new PartitionLeader object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPartitionLeaderWithDefaults() *PartitionLeader {
	this := PartitionLeader{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PartitionLeader) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartitionLeader) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PartitionLeader) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *PartitionLeader) SetId(v int32) {
	o.Id = &v
}

func (o PartitionLeader) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PartitionLeader) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullablePartitionLeader struct {
	value *PartitionLeader
	isSet bool
}

func (v NullablePartitionLeader) Get() *PartitionLeader {
	return v.value
}

func (v *NullablePartitionLeader) Set(val *PartitionLeader) {
	v.value = val
	v.isSet = true
}

func (v NullablePartitionLeader) IsSet() bool {
	return v.isSet
}

func (v *NullablePartitionLeader) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePartitionLeader(val *PartitionLeader) *NullablePartitionLeader {
	return &NullablePartitionLeader{value: val, isSet: true}
}

func (v NullablePartitionLeader) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePartitionLeader) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


