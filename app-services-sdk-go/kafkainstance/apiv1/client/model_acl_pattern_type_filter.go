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

// AclPatternTypeFilter Use value 'MATCH' to perform pattern matching.
type AclPatternTypeFilter string

// List of AclPatternTypeFilter
const (
	ACLPATTERNTYPEFILTER_LITERAL AclPatternTypeFilter = "LITERAL"
	ACLPATTERNTYPEFILTER_PREFIXED AclPatternTypeFilter = "PREFIXED"
	ACLPATTERNTYPEFILTER_ANY AclPatternTypeFilter = "ANY"
	ACLPATTERNTYPEFILTER_MATCH AclPatternTypeFilter = "MATCH"
)

// All allowed values of AclPatternTypeFilter enum
var AllowedAclPatternTypeFilterEnumValues = []AclPatternTypeFilter{
	"LITERAL",
	"PREFIXED",
	"ANY",
	"MATCH",
}

func (v *AclPatternTypeFilter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AclPatternTypeFilter(value)
	for _, existing := range AllowedAclPatternTypeFilterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AclPatternTypeFilter", value)
}

// NewAclPatternTypeFilterFromValue returns a pointer to a valid AclPatternTypeFilter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAclPatternTypeFilterFromValue(v string) (*AclPatternTypeFilter, error) {
	ev := AclPatternTypeFilter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AclPatternTypeFilter: valid values are %v", v, AllowedAclPatternTypeFilterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AclPatternTypeFilter) IsValid() bool {
	for _, existing := range AllowedAclPatternTypeFilterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AclPatternTypeFilter value
func (v AclPatternTypeFilter) Ptr() *AclPatternTypeFilter {
	return &v
}

type NullableAclPatternTypeFilter struct {
	value *AclPatternTypeFilter
	isSet bool
}

func (v NullableAclPatternTypeFilter) Get() *AclPatternTypeFilter {
	return v.value
}

func (v *NullableAclPatternTypeFilter) Set(val *AclPatternTypeFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableAclPatternTypeFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableAclPatternTypeFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAclPatternTypeFilter(val *AclPatternTypeFilter) *NullableAclPatternTypeFilter {
	return &NullableAclPatternTypeFilter{value: val, isSet: true}
}

func (v NullableAclPatternTypeFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAclPatternTypeFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

