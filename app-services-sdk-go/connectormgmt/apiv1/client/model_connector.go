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

// checks if the Connector type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Connector{}

// Connector struct for Connector
type Connector struct {
	Id *string `json:"id,omitempty"`
	Kind *string `json:"kind,omitempty"`
	Href *string `json:"href,omitempty"`
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
	Kafka KafkaConnectionSettings `json:"kafka"`
	ServiceAccount ServiceAccount `json:"service_account"`
	SchemaRegistry *SchemaRegistryConnectionSettings `json:"schema_registry,omitempty"`
	Connector map[string]interface{} `json:"connector"`
	Status *ConnectorStatusStatus `json:"status,omitempty"`
}

// NewConnector instantiates a new Connector object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnector(name string, connectorTypeId string, namespaceId string, desiredState ConnectorDesiredState, kafka KafkaConnectionSettings, serviceAccount ServiceAccount, connector map[string]interface{}) *Connector {
	this := Connector{}
	this.Name = name
	this.ConnectorTypeId = connectorTypeId
	this.NamespaceId = namespaceId
	var channel Channel = CHANNEL_STABLE
	this.Channel = &channel
	this.DesiredState = desiredState
	this.Kafka = kafka
	this.ServiceAccount = serviceAccount
	this.Connector = connector
	return &this
}

// NewConnectorWithDefaults instantiates a new Connector object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectorWithDefaults() *Connector {
	this := Connector{}
	var channel Channel = CHANNEL_STABLE
	this.Channel = &channel
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Connector) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Connector) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Connector) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Connector) SetId(v string) {
	o.Id = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *Connector) GetKind() string {
	if o == nil || IsNil(o.Kind) {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Connector) GetKindOk() (*string, bool) {
	if o == nil || IsNil(o.Kind) {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *Connector) HasKind() bool {
	if o != nil && !IsNil(o.Kind) {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *Connector) SetKind(v string) {
	o.Kind = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *Connector) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Connector) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *Connector) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *Connector) SetHref(v string) {
	o.Href = &v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *Connector) GetOwner() string {
	if o == nil || IsNil(o.Owner) {
		var ret string
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Connector) GetOwnerOk() (*string, bool) {
	if o == nil || IsNil(o.Owner) {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *Connector) HasOwner() bool {
	if o != nil && !IsNil(o.Owner) {
		return true
	}

	return false
}

// SetOwner gets a reference to the given string and assigns it to the Owner field.
func (o *Connector) SetOwner(v string) {
	o.Owner = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Connector) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Connector) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Connector) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *Connector) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetModifiedAt returns the ModifiedAt field value if set, zero value otherwise.
func (o *Connector) GetModifiedAt() time.Time {
	if o == nil || IsNil(o.ModifiedAt) {
		var ret time.Time
		return ret
	}
	return *o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Connector) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ModifiedAt) {
		return nil, false
	}
	return o.ModifiedAt, true
}

// HasModifiedAt returns a boolean if a field has been set.
func (o *Connector) HasModifiedAt() bool {
	if o != nil && !IsNil(o.ModifiedAt) {
		return true
	}

	return false
}

// SetModifiedAt gets a reference to the given time.Time and assigns it to the ModifiedAt field.
func (o *Connector) SetModifiedAt(v time.Time) {
	o.ModifiedAt = &v
}

// GetName returns the Name field value
func (o *Connector) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Connector) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Connector) SetName(v string) {
	o.Name = v
}

// GetConnectorTypeId returns the ConnectorTypeId field value
func (o *Connector) GetConnectorTypeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConnectorTypeId
}

// GetConnectorTypeIdOk returns a tuple with the ConnectorTypeId field value
// and a boolean to check if the value has been set.
func (o *Connector) GetConnectorTypeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConnectorTypeId, true
}

// SetConnectorTypeId sets field value
func (o *Connector) SetConnectorTypeId(v string) {
	o.ConnectorTypeId = v
}

// GetNamespaceId returns the NamespaceId field value
func (o *Connector) GetNamespaceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NamespaceId
}

// GetNamespaceIdOk returns a tuple with the NamespaceId field value
// and a boolean to check if the value has been set.
func (o *Connector) GetNamespaceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NamespaceId, true
}

// SetNamespaceId sets field value
func (o *Connector) SetNamespaceId(v string) {
	o.NamespaceId = v
}

// GetChannel returns the Channel field value if set, zero value otherwise.
func (o *Connector) GetChannel() Channel {
	if o == nil || IsNil(o.Channel) {
		var ret Channel
		return ret
	}
	return *o.Channel
}

// GetChannelOk returns a tuple with the Channel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Connector) GetChannelOk() (*Channel, bool) {
	if o == nil || IsNil(o.Channel) {
		return nil, false
	}
	return o.Channel, true
}

// HasChannel returns a boolean if a field has been set.
func (o *Connector) HasChannel() bool {
	if o != nil && !IsNil(o.Channel) {
		return true
	}

	return false
}

// SetChannel gets a reference to the given Channel and assigns it to the Channel field.
func (o *Connector) SetChannel(v Channel) {
	o.Channel = &v
}

// GetDesiredState returns the DesiredState field value
func (o *Connector) GetDesiredState() ConnectorDesiredState {
	if o == nil {
		var ret ConnectorDesiredState
		return ret
	}

	return o.DesiredState
}

// GetDesiredStateOk returns a tuple with the DesiredState field value
// and a boolean to check if the value has been set.
func (o *Connector) GetDesiredStateOk() (*ConnectorDesiredState, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DesiredState, true
}

// SetDesiredState sets field value
func (o *Connector) SetDesiredState(v ConnectorDesiredState) {
	o.DesiredState = v
}

// GetAnnotations returns the Annotations field value if set, zero value otherwise.
func (o *Connector) GetAnnotations() map[string]string {
	if o == nil || IsNil(o.Annotations) {
		var ret map[string]string
		return ret
	}
	return *o.Annotations
}

// GetAnnotationsOk returns a tuple with the Annotations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Connector) GetAnnotationsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Annotations) {
		return nil, false
	}
	return o.Annotations, true
}

// HasAnnotations returns a boolean if a field has been set.
func (o *Connector) HasAnnotations() bool {
	if o != nil && !IsNil(o.Annotations) {
		return true
	}

	return false
}

// SetAnnotations gets a reference to the given map[string]string and assigns it to the Annotations field.
func (o *Connector) SetAnnotations(v map[string]string) {
	o.Annotations = &v
}

// GetResourceVersion returns the ResourceVersion field value if set, zero value otherwise.
func (o *Connector) GetResourceVersion() int64 {
	if o == nil || IsNil(o.ResourceVersion) {
		var ret int64
		return ret
	}
	return *o.ResourceVersion
}

// GetResourceVersionOk returns a tuple with the ResourceVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Connector) GetResourceVersionOk() (*int64, bool) {
	if o == nil || IsNil(o.ResourceVersion) {
		return nil, false
	}
	return o.ResourceVersion, true
}

// HasResourceVersion returns a boolean if a field has been set.
func (o *Connector) HasResourceVersion() bool {
	if o != nil && !IsNil(o.ResourceVersion) {
		return true
	}

	return false
}

// SetResourceVersion gets a reference to the given int64 and assigns it to the ResourceVersion field.
func (o *Connector) SetResourceVersion(v int64) {
	o.ResourceVersion = &v
}

// GetKafka returns the Kafka field value
func (o *Connector) GetKafka() KafkaConnectionSettings {
	if o == nil {
		var ret KafkaConnectionSettings
		return ret
	}

	return o.Kafka
}

// GetKafkaOk returns a tuple with the Kafka field value
// and a boolean to check if the value has been set.
func (o *Connector) GetKafkaOk() (*KafkaConnectionSettings, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kafka, true
}

// SetKafka sets field value
func (o *Connector) SetKafka(v KafkaConnectionSettings) {
	o.Kafka = v
}

// GetServiceAccount returns the ServiceAccount field value
func (o *Connector) GetServiceAccount() ServiceAccount {
	if o == nil {
		var ret ServiceAccount
		return ret
	}

	return o.ServiceAccount
}

// GetServiceAccountOk returns a tuple with the ServiceAccount field value
// and a boolean to check if the value has been set.
func (o *Connector) GetServiceAccountOk() (*ServiceAccount, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceAccount, true
}

// SetServiceAccount sets field value
func (o *Connector) SetServiceAccount(v ServiceAccount) {
	o.ServiceAccount = v
}

// GetSchemaRegistry returns the SchemaRegistry field value if set, zero value otherwise.
func (o *Connector) GetSchemaRegistry() SchemaRegistryConnectionSettings {
	if o == nil || IsNil(o.SchemaRegistry) {
		var ret SchemaRegistryConnectionSettings
		return ret
	}
	return *o.SchemaRegistry
}

// GetSchemaRegistryOk returns a tuple with the SchemaRegistry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Connector) GetSchemaRegistryOk() (*SchemaRegistryConnectionSettings, bool) {
	if o == nil || IsNil(o.SchemaRegistry) {
		return nil, false
	}
	return o.SchemaRegistry, true
}

// HasSchemaRegistry returns a boolean if a field has been set.
func (o *Connector) HasSchemaRegistry() bool {
	if o != nil && !IsNil(o.SchemaRegistry) {
		return true
	}

	return false
}

// SetSchemaRegistry gets a reference to the given SchemaRegistryConnectionSettings and assigns it to the SchemaRegistry field.
func (o *Connector) SetSchemaRegistry(v SchemaRegistryConnectionSettings) {
	o.SchemaRegistry = &v
}

// GetConnector returns the Connector field value
func (o *Connector) GetConnector() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Connector
}

// GetConnectorOk returns a tuple with the Connector field value
// and a boolean to check if the value has been set.
func (o *Connector) GetConnectorOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Connector, true
}

// SetConnector sets field value
func (o *Connector) SetConnector(v map[string]interface{}) {
	o.Connector = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Connector) GetStatus() ConnectorStatusStatus {
	if o == nil || IsNil(o.Status) {
		var ret ConnectorStatusStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Connector) GetStatusOk() (*ConnectorStatusStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Connector) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given ConnectorStatusStatus and assigns it to the Status field.
func (o *Connector) SetStatus(v ConnectorStatusStatus) {
	o.Status = &v
}

func (o Connector) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Connector) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Kind) {
		toSerialize["kind"] = o.Kind
	}
	if !IsNil(o.Href) {
		toSerialize["href"] = o.Href
	}
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
	toSerialize["kafka"] = o.Kafka
	toSerialize["service_account"] = o.ServiceAccount
	if !IsNil(o.SchemaRegistry) {
		toSerialize["schema_registry"] = o.SchemaRegistry
	}
	toSerialize["connector"] = o.Connector
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableConnector struct {
	value *Connector
	isSet bool
}

func (v NullableConnector) Get() *Connector {
	return v.value
}

func (v *NullableConnector) Set(val *Connector) {
	v.value = val
	v.isSet = true
}

func (v NullableConnector) IsSet() bool {
	return v.isSet
}

func (v *NullableConnector) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnector(val *Connector) *NullableConnector {
	return &NullableConnector{value: val, isSet: true}
}

func (v NullableConnector) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnector) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


