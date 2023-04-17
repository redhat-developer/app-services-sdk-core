/*
Account Management Service API

Manage user subscriptions and clusters

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package accountmgmtclient

import (
	"encoding/json"
)

// checks if the ReservedResourceList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReservedResourceList{}

// ReservedResourceList struct for ReservedResourceList
type ReservedResourceList struct {
	Kind string `json:"kind"`
	Page int32 `json:"page"`
	Size int32 `json:"size"`
	Total int32 `json:"total"`
	Items []ReservedResource `json:"items"`
}

// NewReservedResourceList instantiates a new ReservedResourceList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReservedResourceList(kind string, page int32, size int32, total int32, items []ReservedResource) *ReservedResourceList {
	this := ReservedResourceList{}
	this.Kind = kind
	this.Page = page
	this.Size = size
	this.Total = total
	this.Items = items
	return &this
}

// NewReservedResourceListWithDefaults instantiates a new ReservedResourceList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReservedResourceListWithDefaults() *ReservedResourceList {
	this := ReservedResourceList{}
	return &this
}

// GetKind returns the Kind field value
func (o *ReservedResourceList) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *ReservedResourceList) GetKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *ReservedResourceList) SetKind(v string) {
	o.Kind = v
}

// GetPage returns the Page field value
func (o *ReservedResourceList) GetPage() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Page
}

// GetPageOk returns a tuple with the Page field value
// and a boolean to check if the value has been set.
func (o *ReservedResourceList) GetPageOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Page, true
}

// SetPage sets field value
func (o *ReservedResourceList) SetPage(v int32) {
	o.Page = v
}

// GetSize returns the Size field value
func (o *ReservedResourceList) GetSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *ReservedResourceList) GetSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *ReservedResourceList) SetSize(v int32) {
	o.Size = v
}

// GetTotal returns the Total field value
func (o *ReservedResourceList) GetTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *ReservedResourceList) GetTotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *ReservedResourceList) SetTotal(v int32) {
	o.Total = v
}

// GetItems returns the Items field value
func (o *ReservedResourceList) GetItems() []ReservedResource {
	if o == nil {
		var ret []ReservedResource
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *ReservedResourceList) GetItemsOk() ([]ReservedResource, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *ReservedResourceList) SetItems(v []ReservedResource) {
	o.Items = v
}

func (o ReservedResourceList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReservedResourceList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["kind"] = o.Kind
	toSerialize["page"] = o.Page
	toSerialize["size"] = o.Size
	toSerialize["total"] = o.Total
	toSerialize["items"] = o.Items
	return toSerialize, nil
}

type NullableReservedResourceList struct {
	value *ReservedResourceList
	isSet bool
}

func (v NullableReservedResourceList) Get() *ReservedResourceList {
	return v.value
}

func (v *NullableReservedResourceList) Set(val *ReservedResourceList) {
	v.value = val
	v.isSet = true
}

func (v NullableReservedResourceList) IsSet() bool {
	return v.isSet
}

func (v *NullableReservedResourceList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReservedResourceList(val *ReservedResourceList) *NullableReservedResourceList {
	return &NullableReservedResourceList{value: val, isSet: true}
}

func (v NullableReservedResourceList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReservedResourceList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


