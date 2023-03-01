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

// checks if the FeatureToggleQueryRequestAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FeatureToggleQueryRequestAllOf{}

// FeatureToggleQueryRequestAllOf struct for FeatureToggleQueryRequestAllOf
type FeatureToggleQueryRequestAllOf struct {
	OrganizationId string `json:"organization_id"`
}

// NewFeatureToggleQueryRequestAllOf instantiates a new FeatureToggleQueryRequestAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFeatureToggleQueryRequestAllOf(organizationId string) *FeatureToggleQueryRequestAllOf {
	this := FeatureToggleQueryRequestAllOf{}
	this.OrganizationId = organizationId
	return &this
}

// NewFeatureToggleQueryRequestAllOfWithDefaults instantiates a new FeatureToggleQueryRequestAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFeatureToggleQueryRequestAllOfWithDefaults() *FeatureToggleQueryRequestAllOf {
	this := FeatureToggleQueryRequestAllOf{}
	return &this
}

// GetOrganizationId returns the OrganizationId field value
func (o *FeatureToggleQueryRequestAllOf) GetOrganizationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value
// and a boolean to check if the value has been set.
func (o *FeatureToggleQueryRequestAllOf) GetOrganizationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrganizationId, true
}

// SetOrganizationId sets field value
func (o *FeatureToggleQueryRequestAllOf) SetOrganizationId(v string) {
	o.OrganizationId = v
}

func (o FeatureToggleQueryRequestAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FeatureToggleQueryRequestAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["organization_id"] = o.OrganizationId
	return toSerialize, nil
}

type NullableFeatureToggleQueryRequestAllOf struct {
	value *FeatureToggleQueryRequestAllOf
	isSet bool
}

func (v NullableFeatureToggleQueryRequestAllOf) Get() *FeatureToggleQueryRequestAllOf {
	return v.value
}

func (v *NullableFeatureToggleQueryRequestAllOf) Set(val *FeatureToggleQueryRequestAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableFeatureToggleQueryRequestAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableFeatureToggleQueryRequestAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeatureToggleQueryRequestAllOf(val *FeatureToggleQueryRequestAllOf) *NullableFeatureToggleQueryRequestAllOf {
	return &NullableFeatureToggleQueryRequestAllOf{value: val, isSet: true}
}

func (v NullableFeatureToggleQueryRequestAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFeatureToggleQueryRequestAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


