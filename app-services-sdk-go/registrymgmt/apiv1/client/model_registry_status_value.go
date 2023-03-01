/*
Service Registry Management API

Service Registry Management API is a REST API for managing Service Registry instances. Service Registry is a datastore for event schemas and API designs, which is based on the open source Apicurio Registry project.

API version: 1.0.0
Contact: rhosak-eval-support@redhat.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package registrymgmtclient

import (
	"encoding/json"
	"fmt"
)

// RegistryStatusValue \"accepted\": Registry status when accepted for processing.  \"provisioning\": Registry status when provisioning a new instance.  \"ready\": Registry status when ready for use.  \"failed\": Registry status when the provisioning failed. When removing a Registry instance in this state, the status transitions directly to \"deleting\".   \"deprovision\": Registry status when accepted for deprovisioning.  \"deleting\": Registry status when deprovisioning. 
type RegistryStatusValue string

// List of RegistryStatusValue
const (
	REGISTRYSTATUSVALUE_ACCEPTED RegistryStatusValue = "accepted"
	REGISTRYSTATUSVALUE_PROVISIONING RegistryStatusValue = "provisioning"
	REGISTRYSTATUSVALUE_READY RegistryStatusValue = "ready"
	REGISTRYSTATUSVALUE_FAILED RegistryStatusValue = "failed"
	REGISTRYSTATUSVALUE_DEPROVISION RegistryStatusValue = "deprovision"
	REGISTRYSTATUSVALUE_DELETING RegistryStatusValue = "deleting"
)

// All allowed values of RegistryStatusValue enum
var AllowedRegistryStatusValueEnumValues = []RegistryStatusValue{
	"accepted",
	"provisioning",
	"ready",
	"failed",
	"deprovision",
	"deleting",
}

func (v *RegistryStatusValue) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RegistryStatusValue(value)
	for _, existing := range AllowedRegistryStatusValueEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RegistryStatusValue", value)
}

// NewRegistryStatusValueFromValue returns a pointer to a valid RegistryStatusValue
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRegistryStatusValueFromValue(v string) (*RegistryStatusValue, error) {
	ev := RegistryStatusValue(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RegistryStatusValue: valid values are %v", v, AllowedRegistryStatusValueEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RegistryStatusValue) IsValid() bool {
	for _, existing := range AllowedRegistryStatusValueEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RegistryStatusValue value
func (v RegistryStatusValue) Ptr() *RegistryStatusValue {
	return &v
}

type NullableRegistryStatusValue struct {
	value *RegistryStatusValue
	isSet bool
}

func (v NullableRegistryStatusValue) Get() *RegistryStatusValue {
	return v.value
}

func (v *NullableRegistryStatusValue) Set(val *RegistryStatusValue) {
	v.value = val
	v.isSet = true
}

func (v NullableRegistryStatusValue) IsSet() bool {
	return v.isSet
}

func (v *NullableRegistryStatusValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegistryStatusValue(val *RegistryStatusValue) *NullableRegistryStatusValue {
	return &NullableRegistryStatusValue{value: val, isSet: true}
}

func (v NullableRegistryStatusValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegistryStatusValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

