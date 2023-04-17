/*
Kafka Management API

Kafka Management API is a REST API to manage Kafka instances

API version: 1.16.0
Contact: rhosak-support@redhat.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package kafkamgmtclient

import (
	"encoding/json"
)

// checks if the RangeQuery type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RangeQuery{}

// RangeQuery struct for RangeQuery
type RangeQuery struct {
	Metric *map[string]string `json:"metric,omitempty"`
	Values []Values `json:"values,omitempty"`
}

// NewRangeQuery instantiates a new RangeQuery object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRangeQuery() *RangeQuery {
	this := RangeQuery{}
	return &this
}

// NewRangeQueryWithDefaults instantiates a new RangeQuery object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRangeQueryWithDefaults() *RangeQuery {
	this := RangeQuery{}
	return &this
}

// GetMetric returns the Metric field value if set, zero value otherwise.
func (o *RangeQuery) GetMetric() map[string]string {
	if o == nil || IsNil(o.Metric) {
		var ret map[string]string
		return ret
	}
	return *o.Metric
}

// GetMetricOk returns a tuple with the Metric field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RangeQuery) GetMetricOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Metric) {
		return nil, false
	}
	return o.Metric, true
}

// HasMetric returns a boolean if a field has been set.
func (o *RangeQuery) HasMetric() bool {
	if o != nil && !IsNil(o.Metric) {
		return true
	}

	return false
}

// SetMetric gets a reference to the given map[string]string and assigns it to the Metric field.
func (o *RangeQuery) SetMetric(v map[string]string) {
	o.Metric = &v
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *RangeQuery) GetValues() []Values {
	if o == nil || IsNil(o.Values) {
		var ret []Values
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RangeQuery) GetValuesOk() ([]Values, bool) {
	if o == nil || IsNil(o.Values) {
		return nil, false
	}
	return o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *RangeQuery) HasValues() bool {
	if o != nil && !IsNil(o.Values) {
		return true
	}

	return false
}

// SetValues gets a reference to the given []Values and assigns it to the Values field.
func (o *RangeQuery) SetValues(v []Values) {
	o.Values = v
}

func (o RangeQuery) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RangeQuery) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Metric) {
		toSerialize["metric"] = o.Metric
	}
	if !IsNil(o.Values) {
		toSerialize["values"] = o.Values
	}
	return toSerialize, nil
}

type NullableRangeQuery struct {
	value *RangeQuery
	isSet bool
}

func (v NullableRangeQuery) Get() *RangeQuery {
	return v.value
}

func (v *NullableRangeQuery) Set(val *RangeQuery) {
	v.value = val
	v.isSet = true
}

func (v NullableRangeQuery) IsSet() bool {
	return v.isSet
}

func (v *NullableRangeQuery) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRangeQuery(val *RangeQuery) *NullableRangeQuery {
	return &NullableRangeQuery{value: val, isSet: true}
}

func (v NullableRangeQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRangeQuery) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


