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
	"time"
)

// checks if the ConnectorMeta type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConnectorMeta{}

// ConnectorMeta struct for ConnectorMeta
type ConnectorMeta struct {
	Owner *string `json:"owner,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	ModifiedAt *time.Time `json:"modified_at,omitempty"`
	Name string `json:"name"`
	ConnectorTypeId string `json:"connector_type_id"`
	NamespaceId string `json:"namespace_id"`
	Channel *Channel `json:"channel,omitempty"`
	DesiredState ConnectorDesiredState `json:"desired_state"`
	// Name-value string annotations for resource
	Annotations *map[string]string `json:"annotations,omitempty"`
	ResourceVersion *int64 `json:"resource_version,omitempty"`
}

// NewConnectorMeta instantiates a new ConnectorMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectorMeta(name string, connectorTypeId string, namespaceId string, desiredState ConnectorDesiredState) *ConnectorMeta {
	this := ConnectorMeta{}
	this.Name = name
	this.ConnectorTypeId = connectorTypeId
	this.NamespaceId = namespaceId
	var channel Channel = CHANNEL_STABLE
	this.Channel = &channel
	this.DesiredState = desiredState
	return &this
}

// NewConnectorMetaWithDefaults instantiates a new ConnectorMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectorMetaWithDefaults() *ConnectorMeta {
	this := ConnectorMeta{}
	var channel Channel = CHANNEL_STABLE
	this.Channel = &channel
	return &this
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *ConnectorMeta) GetOwner() string {
	if o == nil || IsNil(o.Owner) {
		var ret string
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorMeta) GetOwnerOk() (*string, bool) {
	if o == nil || IsNil(o.Owner) {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *ConnectorMeta) HasOwner() bool {
	if o != nil && !IsNil(o.Owner) {
		return true
	}

	return false
}

// SetOwner gets a reference to the given string and assigns it to the Owner field.
func (o *ConnectorMeta) SetOwner(v string) {
	o.Owner = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ConnectorMeta) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorMeta) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ConnectorMeta) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *ConnectorMeta) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetModifiedAt returns the ModifiedAt field value if set, zero value otherwise.
func (o *ConnectorMeta) GetModifiedAt() time.Time {
	if o == nil || IsNil(o.ModifiedAt) {
		var ret time.Time
		return ret
	}
	return *o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorMeta) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ModifiedAt) {
		return nil, false
	}
	return o.ModifiedAt, true
}

// HasModifiedAt returns a boolean if a field has been set.
func (o *ConnectorMeta) HasModifiedAt() bool {
	if o != nil && !IsNil(o.ModifiedAt) {
		return true
	}

	return false
}

// SetModifiedAt gets a reference to the given time.Time and assigns it to the ModifiedAt field.
func (o *ConnectorMeta) SetModifiedAt(v time.Time) {
	o.ModifiedAt = &v
}

// GetName returns the Name field value
func (o *ConnectorMeta) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ConnectorMeta) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ConnectorMeta) SetName(v string) {
	o.Name = v
}

// GetConnectorTypeId returns the ConnectorTypeId field value
func (o *ConnectorMeta) GetConnectorTypeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConnectorTypeId
}

// GetConnectorTypeIdOk returns a tuple with the ConnectorTypeId field value
// and a boolean to check if the value has been set.
func (o *ConnectorMeta) GetConnectorTypeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConnectorTypeId, true
}

// SetConnectorTypeId sets field value
func (o *ConnectorMeta) SetConnectorTypeId(v string) {
	o.ConnectorTypeId = v
}

// GetNamespaceId returns the NamespaceId field value
func (o *ConnectorMeta) GetNamespaceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NamespaceId
}

// GetNamespaceIdOk returns a tuple with the NamespaceId field value
// and a boolean to check if the value has been set.
func (o *ConnectorMeta) GetNamespaceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NamespaceId, true
}

// SetNamespaceId sets field value
func (o *ConnectorMeta) SetNamespaceId(v string) {
	o.NamespaceId = v
}

// GetChannel returns the Channel field value if set, zero value otherwise.
func (o *ConnectorMeta) GetChannel() Channel {
	if o == nil || IsNil(o.Channel) {
		var ret Channel
		return ret
	}
	return *o.Channel
}

// GetChannelOk returns a tuple with the Channel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorMeta) GetChannelOk() (*Channel, bool) {
	if o == nil || IsNil(o.Channel) {
		return nil, false
	}
	return o.Channel, true
}

// HasChannel returns a boolean if a field has been set.
func (o *ConnectorMeta) HasChannel() bool {
	if o != nil && !IsNil(o.Channel) {
		return true
	}

	return false
}

// SetChannel gets a reference to the given Channel and assigns it to the Channel field.
func (o *ConnectorMeta) SetChannel(v Channel) {
	o.Channel = &v
}

// GetDesiredState returns the DesiredState field value
func (o *ConnectorMeta) GetDesiredState() ConnectorDesiredState {
	if o == nil {
		var ret ConnectorDesiredState
		return ret
	}

	return o.DesiredState
}

// GetDesiredStateOk returns a tuple with the DesiredState field value
// and a boolean to check if the value has been set.
func (o *ConnectorMeta) GetDesiredStateOk() (*ConnectorDesiredState, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DesiredState, true
}

// SetDesiredState sets field value
func (o *ConnectorMeta) SetDesiredState(v ConnectorDesiredState) {
	o.DesiredState = v
}

// GetAnnotations returns the Annotations field value if set, zero value otherwise.
func (o *ConnectorMeta) GetAnnotations() map[string]string {
	if o == nil || IsNil(o.Annotations) {
		var ret map[string]string
		return ret
	}
	return *o.Annotations
}

// GetAnnotationsOk returns a tuple with the Annotations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorMeta) GetAnnotationsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Annotations) {
		return nil, false
	}
	return o.Annotations, true
}

// HasAnnotations returns a boolean if a field has been set.
func (o *ConnectorMeta) HasAnnotations() bool {
	if o != nil && !IsNil(o.Annotations) {
		return true
	}

	return false
}

// SetAnnotations gets a reference to the given map[string]string and assigns it to the Annotations field.
func (o *ConnectorMeta) SetAnnotations(v map[string]string) {
	o.Annotations = &v
}

// GetResourceVersion returns the ResourceVersion field value if set, zero value otherwise.
func (o *ConnectorMeta) GetResourceVersion() int64 {
	if o == nil || IsNil(o.ResourceVersion) {
		var ret int64
		return ret
	}
	return *o.ResourceVersion
}

// GetResourceVersionOk returns a tuple with the ResourceVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorMeta) GetResourceVersionOk() (*int64, bool) {
	if o == nil || IsNil(o.ResourceVersion) {
		return nil, false
	}
	return o.ResourceVersion, true
}

// HasResourceVersion returns a boolean if a field has been set.
func (o *ConnectorMeta) HasResourceVersion() bool {
	if o != nil && !IsNil(o.ResourceVersion) {
		return true
	}

	return false
}

// SetResourceVersion gets a reference to the given int64 and assigns it to the ResourceVersion field.
func (o *ConnectorMeta) SetResourceVersion(v int64) {
	o.ResourceVersion = &v
}

func (o ConnectorMeta) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConnectorMeta) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Owner) {
		toSerialize["owner"] = o.Owner
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.ModifiedAt) {
		toSerialize["modified_at"] = o.ModifiedAt
	}
	toSerialize["name"] = o.Name
	toSerialize["connector_type_id"] = o.ConnectorTypeId
	toSerialize["namespace_id"] = o.NamespaceId
	if !IsNil(o.Channel) {
		toSerialize["channel"] = o.Channel
	}
	toSerialize["desired_state"] = o.DesiredState
	if !IsNil(o.Annotations) {
		toSerialize["annotations"] = o.Annotations
	}
	if !IsNil(o.ResourceVersion) {
		toSerialize["resource_version"] = o.ResourceVersion
	}
	return toSerialize, nil
}

type NullableConnectorMeta struct {
	value *ConnectorMeta
	isSet bool
}

func (v NullableConnectorMeta) Get() *ConnectorMeta {
	return v.value
}

func (v *NullableConnectorMeta) Set(val *ConnectorMeta) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectorMeta) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectorMeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectorMeta(val *ConnectorMeta) *NullableConnectorMeta {
	return &NullableConnectorMeta{value: val, isSet: true}
}

func (v NullableConnectorMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectorMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


