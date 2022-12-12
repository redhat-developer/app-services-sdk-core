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
	"fmt"
)

// ArtifactType the model 'ArtifactType'
type ArtifactType string

// List of ArtifactType
const (
	ARTIFACTTYPE_AVRO ArtifactType = "AVRO"
	ARTIFACTTYPE_PROTOBUF ArtifactType = "PROTOBUF"
	ARTIFACTTYPE_JSON ArtifactType = "JSON"
	ARTIFACTTYPE_OPENAPI ArtifactType = "OPENAPI"
	ARTIFACTTYPE_ASYNCAPI ArtifactType = "ASYNCAPI"
	ARTIFACTTYPE_GRAPHQL ArtifactType = "GRAPHQL"
	ARTIFACTTYPE_KCONNECT ArtifactType = "KCONNECT"
	ARTIFACTTYPE_WSDL ArtifactType = "WSDL"
	ARTIFACTTYPE_XSD ArtifactType = "XSD"
	ARTIFACTTYPE_XML ArtifactType = "XML"
)

var allowedArtifactTypeEnumValues = []ArtifactType{
	"AVRO",
	"PROTOBUF",
	"JSON",
	"OPENAPI",
	"ASYNCAPI",
	"GRAPHQL",
	"KCONNECT",
	"WSDL",
	"XSD",
	"XML",
}

func (v *ArtifactType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ArtifactType(value)
	for _, existing := range allowedArtifactTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ArtifactType", value)
}

// NewArtifactTypeFromValue returns a pointer to a valid ArtifactType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewArtifactTypeFromValue(v string) (*ArtifactType, error) {
	ev := ArtifactType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ArtifactType: valid values are %v", v, allowedArtifactTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ArtifactType) IsValid() bool {
	for _, existing := range allowedArtifactTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ArtifactType value
func (v ArtifactType) Ptr() *ArtifactType {
	return &v
}

type NullableArtifactType struct {
	value *ArtifactType
	isSet bool
}

func (v NullableArtifactType) Get() *ArtifactType {
	return v.value
}

func (v *NullableArtifactType) Set(val *ArtifactType) {
	v.value = val
	v.isSet = true
}

func (v NullableArtifactType) IsSet() bool {
	return v.isSet
}

func (v *NullableArtifactType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArtifactType(val *ArtifactType) *NullableArtifactType {
	return &NullableArtifactType{value: val, isSet: true}
}

func (v NullableArtifactType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArtifactType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
