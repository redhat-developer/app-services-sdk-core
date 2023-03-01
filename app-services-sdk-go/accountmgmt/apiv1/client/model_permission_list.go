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

// checks if the PermissionList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PermissionList{}

// PermissionList struct for PermissionList
type PermissionList struct {
	Kind string `json:"kind"`
	Page int32 `json:"page"`
	Size int32 `json:"size"`
	Total int32 `json:"total"`
	Items []Permission `json:"items"`
}

// NewPermissionList instantiates a new PermissionList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPermissionList(kind string, page int32, size int32, total int32, items []Permission) *PermissionList {
	this := PermissionList{}
	this.Kind = kind
	this.Page = page
	this.Size = size
	this.Total = total
	this.Items = items
	return &this
}

// NewPermissionListWithDefaults instantiates a new PermissionList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPermissionListWithDefaults() *PermissionList {
	this := PermissionList{}
	return &this
}

// GetKind returns the Kind field value
func (o *PermissionList) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *PermissionList) GetKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *PermissionList) SetKind(v string) {
	o.Kind = v
}

// GetPage returns the Page field value
func (o *PermissionList) GetPage() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Page
}

// GetPageOk returns a tuple with the Page field value
// and a boolean to check if the value has been set.
func (o *PermissionList) GetPageOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Page, true
}

// SetPage sets field value
func (o *PermissionList) SetPage(v int32) {
	o.Page = v
}

// GetSize returns the Size field value
func (o *PermissionList) GetSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *PermissionList) GetSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *PermissionList) SetSize(v int32) {
	o.Size = v
}

// GetTotal returns the Total field value
func (o *PermissionList) GetTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *PermissionList) GetTotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *PermissionList) SetTotal(v int32) {
	o.Total = v
}

// GetItems returns the Items field value
func (o *PermissionList) GetItems() []Permission {
	if o == nil {
		var ret []Permission
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *PermissionList) GetItemsOk() ([]Permission, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *PermissionList) SetItems(v []Permission) {
	o.Items = v
}

func (o PermissionList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PermissionList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["kind"] = o.Kind
	toSerialize["page"] = o.Page
	toSerialize["size"] = o.Size
	toSerialize["total"] = o.Total
	toSerialize["items"] = o.Items
	return toSerialize, nil
}

type NullablePermissionList struct {
	value *PermissionList
	isSet bool
}

func (v NullablePermissionList) Get() *PermissionList {
	return v.value
}

func (v *NullablePermissionList) Set(val *PermissionList) {
	v.value = val
	v.isSet = true
}

func (v NullablePermissionList) IsSet() bool {
	return v.isSet
}

func (v *NullablePermissionList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePermissionList(val *PermissionList) *NullablePermissionList {
	return &NullablePermissionList{value: val, isSet: true}
}

func (v NullablePermissionList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePermissionList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


