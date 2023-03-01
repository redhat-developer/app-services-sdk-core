/*
Apicurio Registry API [v2]

Apicurio Registry is a datastore for standard event schemas and API designs. Apicurio Registry enables developers to manage and share the structure of their data using a REST interface. For example, client applications can dynamically push or pull the latest updates to or from the registry without needing to redeploy. Apicurio Registry also enables developers to create rules that govern how registry content can evolve over time. For example, this includes rules for content validation and version compatibility.  The Apicurio Registry REST API enables client applications to manage the artifacts in the registry. This API provides create, read, update, and delete operations for schema and API artifacts, rules, versions, and metadata.   The supported artifact types include: - Apache Avro schema - AsyncAPI specification - Google protocol buffers - GraphQL schema - JSON Schema - Kafka Connect schema - OpenAPI specification - Web Services Description Language - XML Schema Definition   **Important**: The Apicurio Registry REST API is available from `https://MY-REGISTRY-URL/apis/registry/v2` by default. Therefore you must prefix all API operation paths with `../apis/registry/v2` in this case. For example: `../apis/registry/v2/ids/globalIds/{globalId}`. 

API version: 2.2.5.Final
Contact: apicurio@lists.jboss.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package registryinstanceclient

import (
	"encoding/json"
)

// checks if the ArtifactOwner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ArtifactOwner{}

// ArtifactOwner Describes the ownership of an artifact.
type ArtifactOwner struct {
	Owner *string `json:"owner,omitempty"`
}

// NewArtifactOwner instantiates a new ArtifactOwner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArtifactOwner() *ArtifactOwner {
	this := ArtifactOwner{}
	return &this
}

// NewArtifactOwnerWithDefaults instantiates a new ArtifactOwner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArtifactOwnerWithDefaults() *ArtifactOwner {
	this := ArtifactOwner{}
	return &this
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *ArtifactOwner) GetOwner() string {
	if o == nil || IsNil(o.Owner) {
		var ret string
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArtifactOwner) GetOwnerOk() (*string, bool) {
	if o == nil || IsNil(o.Owner) {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *ArtifactOwner) HasOwner() bool {
	if o != nil && !IsNil(o.Owner) {
		return true
	}

	return false
}

// SetOwner gets a reference to the given string and assigns it to the Owner field.
func (o *ArtifactOwner) SetOwner(v string) {
	o.Owner = &v
}

func (o ArtifactOwner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ArtifactOwner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Owner) {
		toSerialize["owner"] = o.Owner
	}
	return toSerialize, nil
}

type NullableArtifactOwner struct {
	value *ArtifactOwner
	isSet bool
}

func (v NullableArtifactOwner) Get() *ArtifactOwner {
	return v.value
}

func (v *NullableArtifactOwner) Set(val *ArtifactOwner) {
	v.value = val
	v.isSet = true
}

func (v NullableArtifactOwner) IsSet() bool {
	return v.isSet
}

func (v *NullableArtifactOwner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArtifactOwner(val *ArtifactOwner) *NullableArtifactOwner {
	return &NullableArtifactOwner{value: val, isSet: true}
}

func (v NullableArtifactOwner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArtifactOwner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


