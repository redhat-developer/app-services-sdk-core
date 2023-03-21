/*
Account Management Service API

Manage user subscriptions and clusters

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package accountmgmtclient

import (
	"encoding/json"
	"time"
)

// checks if the ClusterUpgrade type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClusterUpgrade{}

// ClusterUpgrade struct for ClusterUpgrade
type ClusterUpgrade struct {
	Available *bool `json:"available,omitempty"`
	State *string `json:"state,omitempty"`
	UpdatedTimestamp *time.Time `json:"updated_timestamp,omitempty"`
	Version *string `json:"version,omitempty"`
}

// NewClusterUpgrade instantiates a new ClusterUpgrade object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterUpgrade() *ClusterUpgrade {
	this := ClusterUpgrade{}
	return &this
}

// NewClusterUpgradeWithDefaults instantiates a new ClusterUpgrade object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterUpgradeWithDefaults() *ClusterUpgrade {
	this := ClusterUpgrade{}
	return &this
}

// GetAvailable returns the Available field value if set, zero value otherwise.
func (o *ClusterUpgrade) GetAvailable() bool {
	if o == nil || IsNil(o.Available) {
		var ret bool
		return ret
	}
	return *o.Available
}

// GetAvailableOk returns a tuple with the Available field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterUpgrade) GetAvailableOk() (*bool, bool) {
	if o == nil || IsNil(o.Available) {
		return nil, false
	}
	return o.Available, true
}

// HasAvailable returns a boolean if a field has been set.
func (o *ClusterUpgrade) HasAvailable() bool {
	if o != nil && !IsNil(o.Available) {
		return true
	}

	return false
}

// SetAvailable gets a reference to the given bool and assigns it to the Available field.
func (o *ClusterUpgrade) SetAvailable(v bool) {
	o.Available = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *ClusterUpgrade) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterUpgrade) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *ClusterUpgrade) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *ClusterUpgrade) SetState(v string) {
	o.State = &v
}

// GetUpdatedTimestamp returns the UpdatedTimestamp field value if set, zero value otherwise.
func (o *ClusterUpgrade) GetUpdatedTimestamp() time.Time {
	if o == nil || IsNil(o.UpdatedTimestamp) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedTimestamp
}

// GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterUpgrade) GetUpdatedTimestampOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedTimestamp) {
		return nil, false
	}
	return o.UpdatedTimestamp, true
}

// HasUpdatedTimestamp returns a boolean if a field has been set.
func (o *ClusterUpgrade) HasUpdatedTimestamp() bool {
	if o != nil && !IsNil(o.UpdatedTimestamp) {
		return true
	}

	return false
}

// SetUpdatedTimestamp gets a reference to the given time.Time and assigns it to the UpdatedTimestamp field.
func (o *ClusterUpgrade) SetUpdatedTimestamp(v time.Time) {
	o.UpdatedTimestamp = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *ClusterUpgrade) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterUpgrade) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *ClusterUpgrade) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *ClusterUpgrade) SetVersion(v string) {
	o.Version = &v
}

func (o ClusterUpgrade) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClusterUpgrade) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Available) {
		toSerialize["available"] = o.Available
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.UpdatedTimestamp) {
		toSerialize["updated_timestamp"] = o.UpdatedTimestamp
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return toSerialize, nil
}

type NullableClusterUpgrade struct {
	value *ClusterUpgrade
	isSet bool
}

func (v NullableClusterUpgrade) Get() *ClusterUpgrade {
	return v.value
}

func (v *NullableClusterUpgrade) Set(val *ClusterUpgrade) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterUpgrade) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterUpgrade) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterUpgrade(val *ClusterUpgrade) *NullableClusterUpgrade {
	return &NullableClusterUpgrade{value: val, isSet: true}
}

func (v NullableClusterUpgrade) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterUpgrade) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


