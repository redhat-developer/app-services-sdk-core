/*
Service Registry Management API

Service Registry Management API is a REST API for managing Service Registry instances. Service Registry is a datastore for event schemas and API designs, which is based on the open source Apicurio Registry project.

API version: 1.0.0
Contact: rhosak-eval-support@redhat.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package registrymgmtclient

import (
	"encoding/json"
)

// checks if the AbstractList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AbstractList{}

// AbstractList struct for AbstractList
type AbstractList struct {
	Kind string `json:"kind"`
	Page int32 `json:"page"`
	Size int32 `json:"size"`
	Total int32 `json:"total"`
}

// NewAbstractList instantiates a new AbstractList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAbstractList(kind string, page int32, size int32, total int32) *AbstractList {
	this := AbstractList{}
	this.Kind = kind
	this.Page = page
	this.Size = size
	this.Total = total
	return &this
}

// NewAbstractListWithDefaults instantiates a new AbstractList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAbstractListWithDefaults() *AbstractList {
	this := AbstractList{}
	return &this
}

// GetKind returns the Kind field value
func (o *AbstractList) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *AbstractList) GetKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *AbstractList) SetKind(v string) {
	o.Kind = v
}

// GetPage returns the Page field value
func (o *AbstractList) GetPage() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Page
}

// GetPageOk returns a tuple with the Page field value
// and a boolean to check if the value has been set.
func (o *AbstractList) GetPageOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Page, true
}

// SetPage sets field value
func (o *AbstractList) SetPage(v int32) {
	o.Page = v
}

// GetSize returns the Size field value
func (o *AbstractList) GetSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *AbstractList) GetSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *AbstractList) SetSize(v int32) {
	o.Size = v
}

// GetTotal returns the Total field value
func (o *AbstractList) GetTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *AbstractList) GetTotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *AbstractList) SetTotal(v int32) {
	o.Total = v
}

func (o AbstractList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AbstractList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["kind"] = o.Kind
	toSerialize["page"] = o.Page
	toSerialize["size"] = o.Size
	toSerialize["total"] = o.Total
	return toSerialize, nil
}

type NullableAbstractList struct {
	value *AbstractList
	isSet bool
}

func (v NullableAbstractList) Get() *AbstractList {
	return v.value
}

func (v *NullableAbstractList) Set(val *AbstractList) {
	v.value = val
	v.isSet = true
}

func (v NullableAbstractList) IsSet() bool {
	return v.isSet
}

func (v *NullableAbstractList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAbstractList(val *AbstractList) *NullableAbstractList {
	return &NullableAbstractList{value: val, isSet: true}
}

func (v NullableAbstractList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAbstractList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


