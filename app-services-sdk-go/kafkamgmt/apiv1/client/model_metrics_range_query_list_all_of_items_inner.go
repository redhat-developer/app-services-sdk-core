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

// checks if the MetricsRangeQueryListAllOfItemsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MetricsRangeQueryListAllOfItemsInner{}

// MetricsRangeQueryListAllOfItemsInner struct for MetricsRangeQueryListAllOfItemsInner
type MetricsRangeQueryListAllOfItemsInner struct {
	Metric *map[string]string `json:"metric,omitempty"`
	Values []Values `json:"values,omitempty"`
}

// NewMetricsRangeQueryListAllOfItemsInner instantiates a new MetricsRangeQueryListAllOfItemsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetricsRangeQueryListAllOfItemsInner() *MetricsRangeQueryListAllOfItemsInner {
	this := MetricsRangeQueryListAllOfItemsInner{}
	return &this
}

// NewMetricsRangeQueryListAllOfItemsInnerWithDefaults instantiates a new MetricsRangeQueryListAllOfItemsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetricsRangeQueryListAllOfItemsInnerWithDefaults() *MetricsRangeQueryListAllOfItemsInner {
	this := MetricsRangeQueryListAllOfItemsInner{}
	return &this
}

// GetMetric returns the Metric field value if set, zero value otherwise.
func (o *MetricsRangeQueryListAllOfItemsInner) GetMetric() map[string]string {
	if o == nil || IsNil(o.Metric) {
		var ret map[string]string
		return ret
	}
	return *o.Metric
}

// GetMetricOk returns a tuple with the Metric field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsRangeQueryListAllOfItemsInner) GetMetricOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Metric) {
		return nil, false
	}
	return o.Metric, true
}

// HasMetric returns a boolean if a field has been set.
func (o *MetricsRangeQueryListAllOfItemsInner) HasMetric() bool {
	if o != nil && !IsNil(o.Metric) {
		return true
	}

	return false
}

// SetMetric gets a reference to the given map[string]string and assigns it to the Metric field.
func (o *MetricsRangeQueryListAllOfItemsInner) SetMetric(v map[string]string) {
	o.Metric = &v
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *MetricsRangeQueryListAllOfItemsInner) GetValues() []Values {
	if o == nil || IsNil(o.Values) {
		var ret []Values
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsRangeQueryListAllOfItemsInner) GetValuesOk() ([]Values, bool) {
	if o == nil || IsNil(o.Values) {
		return nil, false
	}
	return o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *MetricsRangeQueryListAllOfItemsInner) HasValues() bool {
	if o != nil && !IsNil(o.Values) {
		return true
	}

	return false
}

// SetValues gets a reference to the given []Values and assigns it to the Values field.
func (o *MetricsRangeQueryListAllOfItemsInner) SetValues(v []Values) {
	o.Values = v
}

func (o MetricsRangeQueryListAllOfItemsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MetricsRangeQueryListAllOfItemsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Metric) {
		toSerialize["metric"] = o.Metric
	}
	if !IsNil(o.Values) {
		toSerialize["values"] = o.Values
	}
	return toSerialize, nil
}

type NullableMetricsRangeQueryListAllOfItemsInner struct {
	value *MetricsRangeQueryListAllOfItemsInner
	isSet bool
}

func (v NullableMetricsRangeQueryListAllOfItemsInner) Get() *MetricsRangeQueryListAllOfItemsInner {
	return v.value
}

func (v *NullableMetricsRangeQueryListAllOfItemsInner) Set(val *MetricsRangeQueryListAllOfItemsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableMetricsRangeQueryListAllOfItemsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableMetricsRangeQueryListAllOfItemsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetricsRangeQueryListAllOfItemsInner(val *MetricsRangeQueryListAllOfItemsInner) *NullableMetricsRangeQueryListAllOfItemsInner {
	return &NullableMetricsRangeQueryListAllOfItemsInner{value: val, isSet: true}
}

func (v NullableMetricsRangeQueryListAllOfItemsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetricsRangeQueryListAllOfItemsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


