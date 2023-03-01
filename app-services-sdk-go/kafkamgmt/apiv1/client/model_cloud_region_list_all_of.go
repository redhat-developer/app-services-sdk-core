/*
Kafka Management API

Kafka Management API is a REST API to manage Kafka instances

API version: 1.15.0
Contact: rhosak-support@redhat.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package kafkamgmtclient

import (
	"encoding/json"
)

// checks if the CloudRegionListAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CloudRegionListAllOf{}

// CloudRegionListAllOf struct for CloudRegionListAllOf
type CloudRegionListAllOf struct {
	Items []CloudRegion `json:"items"`
}

// NewCloudRegionListAllOf instantiates a new CloudRegionListAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudRegionListAllOf(items []CloudRegion) *CloudRegionListAllOf {
	this := CloudRegionListAllOf{}
	this.Items = items
	return &this
}

// NewCloudRegionListAllOfWithDefaults instantiates a new CloudRegionListAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudRegionListAllOfWithDefaults() *CloudRegionListAllOf {
	this := CloudRegionListAllOf{}
	return &this
}

// GetItems returns the Items field value
func (o *CloudRegionListAllOf) GetItems() []CloudRegion {
	if o == nil {
		var ret []CloudRegion
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *CloudRegionListAllOf) GetItemsOk() ([]CloudRegion, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *CloudRegionListAllOf) SetItems(v []CloudRegion) {
	o.Items = v
}

func (o CloudRegionListAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CloudRegionListAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["items"] = o.Items
	return toSerialize, nil
}

type NullableCloudRegionListAllOf struct {
	value *CloudRegionListAllOf
	isSet bool
}

func (v NullableCloudRegionListAllOf) Get() *CloudRegionListAllOf {
	return v.value
}

func (v *NullableCloudRegionListAllOf) Set(val *CloudRegionListAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudRegionListAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudRegionListAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudRegionListAllOf(val *CloudRegionListAllOf) *NullableCloudRegionListAllOf {
	return &NullableCloudRegionListAllOf{value: val, isSet: true}
}

func (v NullableCloudRegionListAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudRegionListAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


