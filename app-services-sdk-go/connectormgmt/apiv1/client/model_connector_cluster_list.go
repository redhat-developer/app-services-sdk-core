/*
Connector Management API

Connector Management API is a REST API to manage connectors.

API version: 0.1.0
Contact: rhosak-support@redhat.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package connectormgmtclient

import (
	"encoding/json"
)

// checks if the ConnectorClusterList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConnectorClusterList{}

// ConnectorClusterList struct for ConnectorClusterList
type ConnectorClusterList struct {
	Kind string `json:"kind"`
	Page int32 `json:"page"`
	Size int32 `json:"size"`
	Total int32 `json:"total"`
	Items []ConnectorCluster `json:"items"`
}

// NewConnectorClusterList instantiates a new ConnectorClusterList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectorClusterList(kind string, page int32, size int32, total int32, items []ConnectorCluster) *ConnectorClusterList {
	this := ConnectorClusterList{}
	this.Kind = kind
	this.Page = page
	this.Size = size
	this.Total = total
	this.Items = items
	return &this
}

// NewConnectorClusterListWithDefaults instantiates a new ConnectorClusterList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectorClusterListWithDefaults() *ConnectorClusterList {
	this := ConnectorClusterList{}
	return &this
}

// GetKind returns the Kind field value
func (o *ConnectorClusterList) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *ConnectorClusterList) GetKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *ConnectorClusterList) SetKind(v string) {
	o.Kind = v
}

// GetPage returns the Page field value
func (o *ConnectorClusterList) GetPage() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Page
}

// GetPageOk returns a tuple with the Page field value
// and a boolean to check if the value has been set.
func (o *ConnectorClusterList) GetPageOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Page, true
}

// SetPage sets field value
func (o *ConnectorClusterList) SetPage(v int32) {
	o.Page = v
}

// GetSize returns the Size field value
func (o *ConnectorClusterList) GetSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *ConnectorClusterList) GetSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *ConnectorClusterList) SetSize(v int32) {
	o.Size = v
}

// GetTotal returns the Total field value
func (o *ConnectorClusterList) GetTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *ConnectorClusterList) GetTotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *ConnectorClusterList) SetTotal(v int32) {
	o.Total = v
}

// GetItems returns the Items field value
func (o *ConnectorClusterList) GetItems() []ConnectorCluster {
	if o == nil {
		var ret []ConnectorCluster
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *ConnectorClusterList) GetItemsOk() ([]ConnectorCluster, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *ConnectorClusterList) SetItems(v []ConnectorCluster) {
	o.Items = v
}

func (o ConnectorClusterList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConnectorClusterList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["kind"] = o.Kind
	toSerialize["page"] = o.Page
	toSerialize["size"] = o.Size
	toSerialize["total"] = o.Total
	toSerialize["items"] = o.Items
	return toSerialize, nil
}

type NullableConnectorClusterList struct {
	value *ConnectorClusterList
	isSet bool
}

func (v NullableConnectorClusterList) Get() *ConnectorClusterList {
	return v.value
}

func (v *NullableConnectorClusterList) Set(val *ConnectorClusterList) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectorClusterList) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectorClusterList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectorClusterList(val *ConnectorClusterList) *NullableConnectorClusterList {
	return &NullableConnectorClusterList{value: val, isSet: true}
}

func (v NullableConnectorClusterList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectorClusterList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


