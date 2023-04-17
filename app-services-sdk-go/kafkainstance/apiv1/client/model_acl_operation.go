/*
Kafka Instance API

API for interacting with Kafka Instance. Includes Produce, Consume and Admin APIs

API version: 0.14.1-SNAPSHOT
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package kafkainstanceclient

import (
	"encoding/json"
	"fmt"
)

// AclOperation the model 'AclOperation'
type AclOperation string

// List of AclOperation
const (
	ACLOPERATION_ALL AclOperation = "ALL"
	ACLOPERATION_READ AclOperation = "READ"
	ACLOPERATION_WRITE AclOperation = "WRITE"
	ACLOPERATION_CREATE AclOperation = "CREATE"
	ACLOPERATION_DELETE AclOperation = "DELETE"
	ACLOPERATION_ALTER AclOperation = "ALTER"
	ACLOPERATION_DESCRIBE AclOperation = "DESCRIBE"
	ACLOPERATION_DESCRIBE_CONFIGS AclOperation = "DESCRIBE_CONFIGS"
	ACLOPERATION_ALTER_CONFIGS AclOperation = "ALTER_CONFIGS"
)

// All allowed values of AclOperation enum
var AllowedAclOperationEnumValues = []AclOperation{
	"ALL",
	"READ",
	"WRITE",
	"CREATE",
	"DELETE",
	"ALTER",
	"DESCRIBE",
	"DESCRIBE_CONFIGS",
	"ALTER_CONFIGS",
}

func (v *AclOperation) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AclOperation(value)
	for _, existing := range AllowedAclOperationEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AclOperation", value)
}

// NewAclOperationFromValue returns a pointer to a valid AclOperation
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAclOperationFromValue(v string) (*AclOperation, error) {
	ev := AclOperation(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AclOperation: valid values are %v", v, AllowedAclOperationEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AclOperation) IsValid() bool {
	for _, existing := range AllowedAclOperationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AclOperation value
func (v AclOperation) Ptr() *AclOperation {
	return &v
}

type NullableAclOperation struct {
	value *AclOperation
	isSet bool
}

func (v NullableAclOperation) Get() *AclOperation {
	return v.value
}

func (v *NullableAclOperation) Set(val *AclOperation) {
	v.value = val
	v.isSet = true
}

func (v NullableAclOperation) IsSet() bool {
	return v.isSet
}

func (v *NullableAclOperation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAclOperation(val *AclOperation) *NullableAclOperation {
	return &NullableAclOperation{value: val, isSet: true}
}

func (v NullableAclOperation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAclOperation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

