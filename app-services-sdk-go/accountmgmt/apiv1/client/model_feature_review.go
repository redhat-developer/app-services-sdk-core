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

// checks if the FeatureReview type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FeatureReview{}

// FeatureReview struct for FeatureReview
type FeatureReview struct {
	AccountUsername *string `json:"account_username,omitempty"`
	Feature string `json:"feature"`
	OrganizationId *string `json:"organization_id,omitempty"`
}

// NewFeatureReview instantiates a new FeatureReview object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFeatureReview(feature string) *FeatureReview {
	this := FeatureReview{}
	this.Feature = feature
	return &this
}

// NewFeatureReviewWithDefaults instantiates a new FeatureReview object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFeatureReviewWithDefaults() *FeatureReview {
	this := FeatureReview{}
	return &this
}

// GetAccountUsername returns the AccountUsername field value if set, zero value otherwise.
func (o *FeatureReview) GetAccountUsername() string {
	if o == nil || IsNil(o.AccountUsername) {
		var ret string
		return ret
	}
	return *o.AccountUsername
}

// GetAccountUsernameOk returns a tuple with the AccountUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureReview) GetAccountUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.AccountUsername) {
		return nil, false
	}
	return o.AccountUsername, true
}

// HasAccountUsername returns a boolean if a field has been set.
func (o *FeatureReview) HasAccountUsername() bool {
	if o != nil && !IsNil(o.AccountUsername) {
		return true
	}

	return false
}

// SetAccountUsername gets a reference to the given string and assigns it to the AccountUsername field.
func (o *FeatureReview) SetAccountUsername(v string) {
	o.AccountUsername = &v
}

// GetFeature returns the Feature field value
func (o *FeatureReview) GetFeature() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Feature
}

// GetFeatureOk returns a tuple with the Feature field value
// and a boolean to check if the value has been set.
func (o *FeatureReview) GetFeatureOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Feature, true
}

// SetFeature sets field value
func (o *FeatureReview) SetFeature(v string) {
	o.Feature = v
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise.
func (o *FeatureReview) GetOrganizationId() string {
	if o == nil || IsNil(o.OrganizationId) {
		var ret string
		return ret
	}
	return *o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureReview) GetOrganizationIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrganizationId) {
		return nil, false
	}
	return o.OrganizationId, true
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *FeatureReview) HasOrganizationId() bool {
	if o != nil && !IsNil(o.OrganizationId) {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given string and assigns it to the OrganizationId field.
func (o *FeatureReview) SetOrganizationId(v string) {
	o.OrganizationId = &v
}

func (o FeatureReview) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FeatureReview) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccountUsername) {
		toSerialize["account_username"] = o.AccountUsername
	}
	toSerialize["feature"] = o.Feature
	if !IsNil(o.OrganizationId) {
		toSerialize["organization_id"] = o.OrganizationId
	}
	return toSerialize, nil
}

type NullableFeatureReview struct {
	value *FeatureReview
	isSet bool
}

func (v NullableFeatureReview) Get() *FeatureReview {
	return v.value
}

func (v *NullableFeatureReview) Set(val *FeatureReview) {
	v.value = val
	v.isSet = true
}

func (v NullableFeatureReview) IsSet() bool {
	return v.isSet
}

func (v *NullableFeatureReview) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeatureReview(val *FeatureReview) *NullableFeatureReview {
	return &NullableFeatureReview{value: val, isSet: true}
}

func (v NullableFeatureReview) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFeatureReview) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


