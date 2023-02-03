/*
 * Kafka Management API
 *
 * Kafka Management API is a REST API to manage Kafka instances
 *
 * API version: 1.15.0
 * Contact: rhosak-support@redhat.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package kafkamgmtclient

import (
	"encoding/json"
)

// EnterpriseClusterListAllOf struct for EnterpriseClusterListAllOf
type EnterpriseClusterListAllOf struct {
	Items []EnterpriseCluster `json:"items"`
}

// NewEnterpriseClusterListAllOf instantiates a new EnterpriseClusterListAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnterpriseClusterListAllOf(items []EnterpriseCluster) *EnterpriseClusterListAllOf {
	this := EnterpriseClusterListAllOf{}
	this.Items = items
	return &this
}

// NewEnterpriseClusterListAllOfWithDefaults instantiates a new EnterpriseClusterListAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnterpriseClusterListAllOfWithDefaults() *EnterpriseClusterListAllOf {
	this := EnterpriseClusterListAllOf{}
	return &this
}

// GetItems returns the Items field value
func (o *EnterpriseClusterListAllOf) GetItems() []EnterpriseCluster {
	if o == nil {
		var ret []EnterpriseCluster
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *EnterpriseClusterListAllOf) GetItemsOk() (*[]EnterpriseCluster, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Items, true
}

// SetItems sets field value
func (o *EnterpriseClusterListAllOf) SetItems(v []EnterpriseCluster) {
	o.Items = v
}

func (o EnterpriseClusterListAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["items"] = o.Items
	}
	return json.Marshal(toSerialize)
}

type NullableEnterpriseClusterListAllOf struct {
	value *EnterpriseClusterListAllOf
	isSet bool
}

func (v NullableEnterpriseClusterListAllOf) Get() *EnterpriseClusterListAllOf {
	return v.value
}

func (v *NullableEnterpriseClusterListAllOf) Set(val *EnterpriseClusterListAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableEnterpriseClusterListAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableEnterpriseClusterListAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnterpriseClusterListAllOf(val *EnterpriseClusterListAllOf) *NullableEnterpriseClusterListAllOf {
	return &NullableEnterpriseClusterListAllOf{value: val, isSet: true}
}

func (v NullableEnterpriseClusterListAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnterpriseClusterListAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


