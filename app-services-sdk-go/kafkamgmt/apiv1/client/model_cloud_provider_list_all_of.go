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

// checks if the CloudProviderListAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CloudProviderListAllOf{}

// CloudProviderListAllOf struct for CloudProviderListAllOf
type CloudProviderListAllOf struct {
	Items []CloudProviderListAllOfItemsInner `json:"items"`
}

// NewCloudProviderListAllOf instantiates a new CloudProviderListAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudProviderListAllOf(items []CloudProviderListAllOfItemsInner) *CloudProviderListAllOf {
	this := CloudProviderListAllOf{}
	this.Items = items
	return &this
}

// NewCloudProviderListAllOfWithDefaults instantiates a new CloudProviderListAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudProviderListAllOfWithDefaults() *CloudProviderListAllOf {
	this := CloudProviderListAllOf{}
	return &this
}

// GetItems returns the Items field value
func (o *CloudProviderListAllOf) GetItems() []CloudProviderListAllOfItemsInner {
	if o == nil {
		var ret []CloudProviderListAllOfItemsInner
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *CloudProviderListAllOf) GetItemsOk() ([]CloudProviderListAllOfItemsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *CloudProviderListAllOf) SetItems(v []CloudProviderListAllOfItemsInner) {
	o.Items = v
}

func (o CloudProviderListAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CloudProviderListAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["items"] = o.Items
	return toSerialize, nil
}

type NullableCloudProviderListAllOf struct {
	value *CloudProviderListAllOf
	isSet bool
}

func (v NullableCloudProviderListAllOf) Get() *CloudProviderListAllOf {
	return v.value
}

func (v *NullableCloudProviderListAllOf) Set(val *CloudProviderListAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudProviderListAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudProviderListAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudProviderListAllOf(val *CloudProviderListAllOf) *NullableCloudProviderListAllOf {
	return &NullableCloudProviderListAllOf{value: val, isSet: true}
}

func (v NullableCloudProviderListAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudProviderListAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


