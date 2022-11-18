/*
 * Service Registry API
 *
 * Service Registry Instance API  NOTE: This API cannot be called directly from the portal.
 *
 * API version: 2.2.5.Final
 * Contact: apicurio@lists.jboss.org
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package registryinstanceclient

import (
	"encoding/json"
)

// ArtifactOwner Describes the ownership of an artifact.
type ArtifactOwner struct {
	Owner *string `json:"owner,omitempty"`
}

// NewArtifactOwner instantiates a new ArtifactOwner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArtifactOwner() *ArtifactOwner {
	this := ArtifactOwner{}
	return &this
}

// NewArtifactOwnerWithDefaults instantiates a new ArtifactOwner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArtifactOwnerWithDefaults() *ArtifactOwner {
	this := ArtifactOwner{}
	return &this
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *ArtifactOwner) GetOwner() string {
	if o == nil || o.Owner == nil {
		var ret string
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArtifactOwner) GetOwnerOk() (*string, bool) {
	if o == nil || o.Owner == nil {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *ArtifactOwner) HasOwner() bool {
	if o != nil && o.Owner != nil {
		return true
	}

	return false
}

// SetOwner gets a reference to the given string and assigns it to the Owner field.
func (o *ArtifactOwner) SetOwner(v string) {
	o.Owner = &v
}

func (o ArtifactOwner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Owner != nil {
		toSerialize["owner"] = o.Owner
	}
	return json.Marshal(toSerialize)
}

type NullableArtifactOwner struct {
	value *ArtifactOwner
	isSet bool
}

func (v NullableArtifactOwner) Get() *ArtifactOwner {
	return v.value
}

func (v *NullableArtifactOwner) Set(val *ArtifactOwner) {
	v.value = val
	v.isSet = true
}

func (v NullableArtifactOwner) IsSet() bool {
	return v.isSet
}

func (v *NullableArtifactOwner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArtifactOwner(val *ArtifactOwner) *NullableArtifactOwner {
	return &NullableArtifactOwner{value: val, isSet: true}
}

func (v NullableArtifactOwner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArtifactOwner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

