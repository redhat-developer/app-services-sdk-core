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

// checks if the KafkaRequestList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KafkaRequestList{}

// KafkaRequestList struct for KafkaRequestList
type KafkaRequestList struct {
	Kind string `json:"kind"`
	Page int32 `json:"page"`
	Size int32 `json:"size"`
	Total int32 `json:"total"`
	Items []KafkaRequestListAllOfItemsInner `json:"items"`
}

// NewKafkaRequestList instantiates a new KafkaRequestList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKafkaRequestList(kind string, page int32, size int32, total int32, items []KafkaRequestListAllOfItemsInner) *KafkaRequestList {
	this := KafkaRequestList{}
	this.Kind = kind
	this.Page = page
	this.Size = size
	this.Total = total
	this.Items = items
	return &this
}

// NewKafkaRequestListWithDefaults instantiates a new KafkaRequestList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKafkaRequestListWithDefaults() *KafkaRequestList {
	this := KafkaRequestList{}
	return &this
}

// GetKind returns the Kind field value
func (o *KafkaRequestList) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *KafkaRequestList) GetKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *KafkaRequestList) SetKind(v string) {
	o.Kind = v
}

// GetPage returns the Page field value
func (o *KafkaRequestList) GetPage() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Page
}

// GetPageOk returns a tuple with the Page field value
// and a boolean to check if the value has been set.
func (o *KafkaRequestList) GetPageOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Page, true
}

// SetPage sets field value
func (o *KafkaRequestList) SetPage(v int32) {
	o.Page = v
}

// GetSize returns the Size field value
func (o *KafkaRequestList) GetSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *KafkaRequestList) GetSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *KafkaRequestList) SetSize(v int32) {
	o.Size = v
}

// GetTotal returns the Total field value
func (o *KafkaRequestList) GetTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *KafkaRequestList) GetTotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *KafkaRequestList) SetTotal(v int32) {
	o.Total = v
}

// GetItems returns the Items field value
func (o *KafkaRequestList) GetItems() []KafkaRequestListAllOfItemsInner {
	if o == nil {
		var ret []KafkaRequestListAllOfItemsInner
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *KafkaRequestList) GetItemsOk() ([]KafkaRequestListAllOfItemsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *KafkaRequestList) SetItems(v []KafkaRequestListAllOfItemsInner) {
	o.Items = v
}

func (o KafkaRequestList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KafkaRequestList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["kind"] = o.Kind
	toSerialize["page"] = o.Page
	toSerialize["size"] = o.Size
	toSerialize["total"] = o.Total
	toSerialize["items"] = o.Items
	return toSerialize, nil
}

type NullableKafkaRequestList struct {
	value *KafkaRequestList
	isSet bool
}

func (v NullableKafkaRequestList) Get() *KafkaRequestList {
	return v.value
}

func (v *NullableKafkaRequestList) Set(val *KafkaRequestList) {
	v.value = val
	v.isSet = true
}

func (v NullableKafkaRequestList) IsSet() bool {
	return v.isSet
}

func (v *NullableKafkaRequestList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKafkaRequestList(val *KafkaRequestList) *NullableKafkaRequestList {
	return &NullableKafkaRequestList{value: val, isSet: true}
}

func (v NullableKafkaRequestList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKafkaRequestList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


