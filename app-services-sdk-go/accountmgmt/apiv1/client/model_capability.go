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

// checks if the Capability type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Capability{}

// Capability struct for Capability
type Capability struct {
	Href *string `json:"href,omitempty"`
	Id *string `json:"id,omitempty"`
	Kind *string `json:"kind,omitempty"`
	Inherited bool `json:"inherited"`
	Name string `json:"name"`
	Value string `json:"value"`
}

// NewCapability instantiates a new Capability object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCapability(inherited bool, name string, value string) *Capability {
	this := Capability{}
	this.Inherited = inherited
	this.Name = name
	this.Value = value
	return &this
}

// NewCapabilityWithDefaults instantiates a new Capability object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCapabilityWithDefaults() *Capability {
	this := Capability{}
	return &this
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *Capability) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Capability) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *Capability) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *Capability) SetHref(v string) {
	o.Href = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Capability) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Capability) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Capability) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Capability) SetId(v string) {
	o.Id = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *Capability) GetKind() string {
	if o == nil || IsNil(o.Kind) {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Capability) GetKindOk() (*string, bool) {
	if o == nil || IsNil(o.Kind) {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *Capability) HasKind() bool {
	if o != nil && !IsNil(o.Kind) {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *Capability) SetKind(v string) {
	o.Kind = &v
}

// GetInherited returns the Inherited field value
func (o *Capability) GetInherited() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Inherited
}

// GetInheritedOk returns a tuple with the Inherited field value
// and a boolean to check if the value has been set.
func (o *Capability) GetInheritedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Inherited, true
}

// SetInherited sets field value
func (o *Capability) SetInherited(v bool) {
	o.Inherited = v
}

// GetName returns the Name field value
func (o *Capability) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Capability) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Capability) SetName(v string) {
	o.Name = v
}

// GetValue returns the Value field value
func (o *Capability) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *Capability) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *Capability) SetValue(v string) {
	o.Value = v
}

func (o Capability) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Capability) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Href) {
		toSerialize["href"] = o.Href
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Kind) {
		toSerialize["kind"] = o.Kind
	}
	toSerialize["inherited"] = o.Inherited
	toSerialize["name"] = o.Name
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

type NullableCapability struct {
	value *Capability
	isSet bool
}

func (v NullableCapability) Get() *Capability {
	return v.value
}

func (v *NullableCapability) Set(val *Capability) {
	v.value = val
	v.isSet = true
}

func (v NullableCapability) IsSet() bool {
	return v.isSet
}

func (v *NullableCapability) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCapability(val *Capability) *NullableCapability {
	return &NullableCapability{value: val, isSet: true}
}

func (v NullableCapability) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCapability) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


