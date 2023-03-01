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

// checks if the ConnectorNamespaceMetaAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConnectorNamespaceMetaAllOf{}

// ConnectorNamespaceMetaAllOf struct for ConnectorNamespaceMetaAllOf
type ConnectorNamespaceMetaAllOf struct {
	ResourceVersion *int64 `json:"resource_version,omitempty"`
	Quota *ConnectorNamespaceQuota `json:"quota,omitempty"`
}

// NewConnectorNamespaceMetaAllOf instantiates a new ConnectorNamespaceMetaAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectorNamespaceMetaAllOf() *ConnectorNamespaceMetaAllOf {
	this := ConnectorNamespaceMetaAllOf{}
	return &this
}

// NewConnectorNamespaceMetaAllOfWithDefaults instantiates a new ConnectorNamespaceMetaAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectorNamespaceMetaAllOfWithDefaults() *ConnectorNamespaceMetaAllOf {
	this := ConnectorNamespaceMetaAllOf{}
	return &this
}

// GetResourceVersion returns the ResourceVersion field value if set, zero value otherwise.
func (o *ConnectorNamespaceMetaAllOf) GetResourceVersion() int64 {
	if o == nil || IsNil(o.ResourceVersion) {
		var ret int64
		return ret
	}
	return *o.ResourceVersion
}

// GetResourceVersionOk returns a tuple with the ResourceVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorNamespaceMetaAllOf) GetResourceVersionOk() (*int64, bool) {
	if o == nil || IsNil(o.ResourceVersion) {
		return nil, false
	}
	return o.ResourceVersion, true
}

// HasResourceVersion returns a boolean if a field has been set.
func (o *ConnectorNamespaceMetaAllOf) HasResourceVersion() bool {
	if o != nil && !IsNil(o.ResourceVersion) {
		return true
	}

	return false
}

// SetResourceVersion gets a reference to the given int64 and assigns it to the ResourceVersion field.
func (o *ConnectorNamespaceMetaAllOf) SetResourceVersion(v int64) {
	o.ResourceVersion = &v
}

// GetQuota returns the Quota field value if set, zero value otherwise.
func (o *ConnectorNamespaceMetaAllOf) GetQuota() ConnectorNamespaceQuota {
	if o == nil || IsNil(o.Quota) {
		var ret ConnectorNamespaceQuota
		return ret
	}
	return *o.Quota
}

// GetQuotaOk returns a tuple with the Quota field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorNamespaceMetaAllOf) GetQuotaOk() (*ConnectorNamespaceQuota, bool) {
	if o == nil || IsNil(o.Quota) {
		return nil, false
	}
	return o.Quota, true
}

// HasQuota returns a boolean if a field has been set.
func (o *ConnectorNamespaceMetaAllOf) HasQuota() bool {
	if o != nil && !IsNil(o.Quota) {
		return true
	}

	return false
}

// SetQuota gets a reference to the given ConnectorNamespaceQuota and assigns it to the Quota field.
func (o *ConnectorNamespaceMetaAllOf) SetQuota(v ConnectorNamespaceQuota) {
	o.Quota = &v
}

func (o ConnectorNamespaceMetaAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConnectorNamespaceMetaAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ResourceVersion) {
		toSerialize["resource_version"] = o.ResourceVersion
	}
	if !IsNil(o.Quota) {
		toSerialize["quota"] = o.Quota
	}
	return toSerialize, nil
}

type NullableConnectorNamespaceMetaAllOf struct {
	value *ConnectorNamespaceMetaAllOf
	isSet bool
}

func (v NullableConnectorNamespaceMetaAllOf) Get() *ConnectorNamespaceMetaAllOf {
	return v.value
}

func (v *NullableConnectorNamespaceMetaAllOf) Set(val *ConnectorNamespaceMetaAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectorNamespaceMetaAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectorNamespaceMetaAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectorNamespaceMetaAllOf(val *ConnectorNamespaceMetaAllOf) *NullableConnectorNamespaceMetaAllOf {
	return &NullableConnectorNamespaceMetaAllOf{value: val, isSet: true}
}

func (v NullableConnectorNamespaceMetaAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectorNamespaceMetaAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


