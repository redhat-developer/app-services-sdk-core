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

// checks if the MetricsRangeQueryListAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MetricsRangeQueryListAllOf{}

// MetricsRangeQueryListAllOf struct for MetricsRangeQueryListAllOf
type MetricsRangeQueryListAllOf struct {
	Kind *string `json:"kind,omitempty"`
	Id *string `json:"id,omitempty"`
	Items []MetricsRangeQueryListAllOfItemsInner `json:"items,omitempty"`
}

// NewMetricsRangeQueryListAllOf instantiates a new MetricsRangeQueryListAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetricsRangeQueryListAllOf() *MetricsRangeQueryListAllOf {
	this := MetricsRangeQueryListAllOf{}
	return &this
}

// NewMetricsRangeQueryListAllOfWithDefaults instantiates a new MetricsRangeQueryListAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetricsRangeQueryListAllOfWithDefaults() *MetricsRangeQueryListAllOf {
	this := MetricsRangeQueryListAllOf{}
	return &this
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *MetricsRangeQueryListAllOf) GetKind() string {
	if o == nil || IsNil(o.Kind) {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsRangeQueryListAllOf) GetKindOk() (*string, bool) {
	if o == nil || IsNil(o.Kind) {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *MetricsRangeQueryListAllOf) HasKind() bool {
	if o != nil && !IsNil(o.Kind) {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *MetricsRangeQueryListAllOf) SetKind(v string) {
	o.Kind = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MetricsRangeQueryListAllOf) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsRangeQueryListAllOf) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MetricsRangeQueryListAllOf) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MetricsRangeQueryListAllOf) SetId(v string) {
	o.Id = &v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *MetricsRangeQueryListAllOf) GetItems() []MetricsRangeQueryListAllOfItemsInner {
	if o == nil || IsNil(o.Items) {
		var ret []MetricsRangeQueryListAllOfItemsInner
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsRangeQueryListAllOf) GetItemsOk() ([]MetricsRangeQueryListAllOfItemsInner, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *MetricsRangeQueryListAllOf) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []MetricsRangeQueryListAllOfItemsInner and assigns it to the Items field.
func (o *MetricsRangeQueryListAllOf) SetItems(v []MetricsRangeQueryListAllOfItemsInner) {
	o.Items = v
}

func (o MetricsRangeQueryListAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MetricsRangeQueryListAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Kind) {
		toSerialize["kind"] = o.Kind
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	return toSerialize, nil
}

type NullableMetricsRangeQueryListAllOf struct {
	value *MetricsRangeQueryListAllOf
	isSet bool
}

func (v NullableMetricsRangeQueryListAllOf) Get() *MetricsRangeQueryListAllOf {
	return v.value
}

func (v *NullableMetricsRangeQueryListAllOf) Set(val *MetricsRangeQueryListAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableMetricsRangeQueryListAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableMetricsRangeQueryListAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetricsRangeQueryListAllOf(val *MetricsRangeQueryListAllOf) *NullableMetricsRangeQueryListAllOf {
	return &NullableMetricsRangeQueryListAllOf{value: val, isSet: true}
}

func (v NullableMetricsRangeQueryListAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetricsRangeQueryListAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


