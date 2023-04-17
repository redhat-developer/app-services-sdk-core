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

// checks if the CloudProviderListAllOfItemsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CloudProviderListAllOfItemsInner{}

// CloudProviderListAllOfItemsInner struct for CloudProviderListAllOfItemsInner
type CloudProviderListAllOfItemsInner struct {
	// Indicates the type of this object. Will be 'CloudProvider' link.
	Kind *string `json:"kind,omitempty"`
	// Unique identifier of the object.
	Id *string `json:"id,omitempty"`
	// Name of the cloud provider for display purposes.
	DisplayName *string `json:"display_name,omitempty"`
	// Human friendly identifier of the cloud provider, for example `aws`.
	Name *string `json:"name,omitempty"`
	// Whether the cloud provider is enabled for deploying an OSD cluster.
	Enabled bool `json:"enabled"`
}

// NewCloudProviderListAllOfItemsInner instantiates a new CloudProviderListAllOfItemsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudProviderListAllOfItemsInner(enabled bool) *CloudProviderListAllOfItemsInner {
	this := CloudProviderListAllOfItemsInner{}
	this.Enabled = enabled
	return &this
}

// NewCloudProviderListAllOfItemsInnerWithDefaults instantiates a new CloudProviderListAllOfItemsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudProviderListAllOfItemsInnerWithDefaults() *CloudProviderListAllOfItemsInner {
	this := CloudProviderListAllOfItemsInner{}
	return &this
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *CloudProviderListAllOfItemsInner) GetKind() string {
	if o == nil || IsNil(o.Kind) {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudProviderListAllOfItemsInner) GetKindOk() (*string, bool) {
	if o == nil || IsNil(o.Kind) {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *CloudProviderListAllOfItemsInner) HasKind() bool {
	if o != nil && !IsNil(o.Kind) {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *CloudProviderListAllOfItemsInner) SetKind(v string) {
	o.Kind = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CloudProviderListAllOfItemsInner) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudProviderListAllOfItemsInner) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CloudProviderListAllOfItemsInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CloudProviderListAllOfItemsInner) SetId(v string) {
	o.Id = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *CloudProviderListAllOfItemsInner) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudProviderListAllOfItemsInner) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *CloudProviderListAllOfItemsInner) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *CloudProviderListAllOfItemsInner) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CloudProviderListAllOfItemsInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudProviderListAllOfItemsInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CloudProviderListAllOfItemsInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CloudProviderListAllOfItemsInner) SetName(v string) {
	o.Name = &v
}

// GetEnabled returns the Enabled field value
func (o *CloudProviderListAllOfItemsInner) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *CloudProviderListAllOfItemsInner) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *CloudProviderListAllOfItemsInner) SetEnabled(v bool) {
	o.Enabled = v
}

func (o CloudProviderListAllOfItemsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CloudProviderListAllOfItemsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Kind) {
		toSerialize["kind"] = o.Kind
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.DisplayName) {
		toSerialize["display_name"] = o.DisplayName
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	toSerialize["enabled"] = o.Enabled
	return toSerialize, nil
}

type NullableCloudProviderListAllOfItemsInner struct {
	value *CloudProviderListAllOfItemsInner
	isSet bool
}

func (v NullableCloudProviderListAllOfItemsInner) Get() *CloudProviderListAllOfItemsInner {
	return v.value
}

func (v *NullableCloudProviderListAllOfItemsInner) Set(val *CloudProviderListAllOfItemsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudProviderListAllOfItemsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudProviderListAllOfItemsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudProviderListAllOfItemsInner(val *CloudProviderListAllOfItemsInner) *NullableCloudProviderListAllOfItemsInner {
	return &NullableCloudProviderListAllOfItemsInner{value: val, isSet: true}
}

func (v NullableCloudProviderListAllOfItemsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudProviderListAllOfItemsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

