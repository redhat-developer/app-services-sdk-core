/*
Red Hat Openshift SmartEvents Fleet Manager V2

The API exposed by the fleet manager of the SmartEvents service.

API version: 0.0.1
Contact: openbridge-dev@redhat.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package smarteventsmgmtclient

import (
	"encoding/json"
)

// checks if the ConnectorRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConnectorRequest{}

// ConnectorRequest struct for ConnectorRequest
type ConnectorRequest struct {
	// The name of the connector
	Name string `json:"name"`
	// The name of the connector
	ConnectorTypeId string `json:"connector_type_id"`
	// The Connector configuration payload
	Connector map[string]interface{} `json:"connector"`
}

// NewConnectorRequest instantiates a new ConnectorRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectorRequest(name string, connectorTypeId string, connector map[string]interface{}) *ConnectorRequest {
	this := ConnectorRequest{}
	this.Name = name
	this.ConnectorTypeId = connectorTypeId
	this.Connector = connector
	return &this
}

// NewConnectorRequestWithDefaults instantiates a new ConnectorRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectorRequestWithDefaults() *ConnectorRequest {
	this := ConnectorRequest{}
	return &this
}

// GetName returns the Name field value
func (o *ConnectorRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ConnectorRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ConnectorRequest) SetName(v string) {
	o.Name = v
}

// GetConnectorTypeId returns the ConnectorTypeId field value
func (o *ConnectorRequest) GetConnectorTypeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConnectorTypeId
}

// GetConnectorTypeIdOk returns a tuple with the ConnectorTypeId field value
// and a boolean to check if the value has been set.
func (o *ConnectorRequest) GetConnectorTypeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConnectorTypeId, true
}

// SetConnectorTypeId sets field value
func (o *ConnectorRequest) SetConnectorTypeId(v string) {
	o.ConnectorTypeId = v
}

// GetConnector returns the Connector field value
func (o *ConnectorRequest) GetConnector() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Connector
}

// GetConnectorOk returns a tuple with the Connector field value
// and a boolean to check if the value has been set.
func (o *ConnectorRequest) GetConnectorOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Connector, true
}

// SetConnector sets field value
func (o *ConnectorRequest) SetConnector(v map[string]interface{}) {
	o.Connector = v
}

func (o ConnectorRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConnectorRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["connector_type_id"] = o.ConnectorTypeId
	toSerialize["connector"] = o.Connector
	return toSerialize, nil
}

type NullableConnectorRequest struct {
	value *ConnectorRequest
	isSet bool
}

func (v NullableConnectorRequest) Get() *ConnectorRequest {
	return v.value
}

func (v *NullableConnectorRequest) Set(val *ConnectorRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectorRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectorRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectorRequest(val *ConnectorRequest) *NullableConnectorRequest {
	return &NullableConnectorRequest{value: val, isSet: true}
}

func (v NullableConnectorRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectorRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


