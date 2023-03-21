/*
Apicurio Registry API [v2]

Apicurio Registry is a datastore for standard event schemas and API designs. Apicurio Registry enables developers to manage and share the structure of their data using a REST interface. For example, client applications can dynamically push or pull the latest updates to or from the registry without needing to redeploy. Apicurio Registry also enables developers to create rules that govern how registry content can evolve over time. For example, this includes rules for content validation and version compatibility.  The Apicurio Registry REST API enables client applications to manage the artifacts in the registry. This API provides create, read, update, and delete operations for schema and API artifacts, rules, versions, and metadata.   The supported artifact types include: - Apache Avro schema - AsyncAPI specification - Google protocol buffers - GraphQL schema - JSON Schema - Kafka Connect schema - OpenAPI specification - Web Services Description Language - XML Schema Definition   **Important**: The Apicurio Registry REST API is available from `https://MY-REGISTRY-URL/apis/registry/v2` by default. Therefore you must prefix all API operation paths with `../apis/registry/v2` in this case. For example: `../apis/registry/v2/ids/globalIds/{globalId}`. 

API version: 2.4.x
Contact: apicurio@lists.jboss.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package registryinstanceclient

import (
	"encoding/json"
	"time"
)

// checks if the VersionMetaData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VersionMetaData{}

// VersionMetaData 
type VersionMetaData struct {
	Version string `json:"version"`
	Name *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	CreatedBy string `json:"createdBy"`
	CreatedOn time.Time `json:"createdOn"`
	// 
	Type string `json:"type"`
	// 
	GlobalId int64 `json:"globalId"`
	State *ArtifactState `json:"state,omitempty"`
	// The ID of a single artifact.
	Id string `json:"id"`
	// 
	Labels []string `json:"labels,omitempty"`
	// User-defined name-value pairs. Name and value must be strings.
	Properties *map[string]string `json:"properties,omitempty"`
	// An ID of a single artifact group.
	GroupId *string `json:"groupId,omitempty"`
	// 
	ContentId int64 `json:"contentId"`
}

// NewVersionMetaData instantiates a new VersionMetaData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVersionMetaData(version string, createdBy string, createdOn time.Time, type_ string, globalId int64, id string, contentId int64) *VersionMetaData {
	this := VersionMetaData{}
	this.Version = version
	this.CreatedBy = createdBy
	this.CreatedOn = createdOn
	this.Type = type_
	this.GlobalId = globalId
	this.Id = id
	this.ContentId = contentId
	return &this
}

// NewVersionMetaDataWithDefaults instantiates a new VersionMetaData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVersionMetaDataWithDefaults() *VersionMetaData {
	this := VersionMetaData{}
	return &this
}

// GetVersion returns the Version field value
func (o *VersionMetaData) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *VersionMetaData) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *VersionMetaData) SetVersion(v string) {
	o.Version = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *VersionMetaData) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionMetaData) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *VersionMetaData) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *VersionMetaData) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *VersionMetaData) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionMetaData) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *VersionMetaData) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *VersionMetaData) SetDescription(v string) {
	o.Description = &v
}

// GetCreatedBy returns the CreatedBy field value
func (o *VersionMetaData) GetCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *VersionMetaData) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value
func (o *VersionMetaData) SetCreatedBy(v string) {
	o.CreatedBy = v
}

// GetCreatedOn returns the CreatedOn field value
func (o *VersionMetaData) GetCreatedOn() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedOn
}

// GetCreatedOnOk returns a tuple with the CreatedOn field value
// and a boolean to check if the value has been set.
func (o *VersionMetaData) GetCreatedOnOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedOn, true
}

// SetCreatedOn sets field value
func (o *VersionMetaData) SetCreatedOn(v time.Time) {
	o.CreatedOn = v
}

// GetType returns the Type field value
func (o *VersionMetaData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *VersionMetaData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *VersionMetaData) SetType(v string) {
	o.Type = v
}

// GetGlobalId returns the GlobalId field value
func (o *VersionMetaData) GetGlobalId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.GlobalId
}

// GetGlobalIdOk returns a tuple with the GlobalId field value
// and a boolean to check if the value has been set.
func (o *VersionMetaData) GetGlobalIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GlobalId, true
}

// SetGlobalId sets field value
func (o *VersionMetaData) SetGlobalId(v int64) {
	o.GlobalId = v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *VersionMetaData) GetState() ArtifactState {
	if o == nil || IsNil(o.State) {
		var ret ArtifactState
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionMetaData) GetStateOk() (*ArtifactState, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *VersionMetaData) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given ArtifactState and assigns it to the State field.
func (o *VersionMetaData) SetState(v ArtifactState) {
	o.State = &v
}

// GetId returns the Id field value
func (o *VersionMetaData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *VersionMetaData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *VersionMetaData) SetId(v string) {
	o.Id = v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *VersionMetaData) GetLabels() []string {
	if o == nil || IsNil(o.Labels) {
		var ret []string
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionMetaData) GetLabelsOk() ([]string, bool) {
	if o == nil || IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *VersionMetaData) HasLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []string and assigns it to the Labels field.
func (o *VersionMetaData) SetLabels(v []string) {
	o.Labels = v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *VersionMetaData) GetProperties() map[string]string {
	if o == nil || IsNil(o.Properties) {
		var ret map[string]string
		return ret
	}
	return *o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionMetaData) GetPropertiesOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Properties) {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *VersionMetaData) HasProperties() bool {
	if o != nil && !IsNil(o.Properties) {
		return true
	}

	return false
}

// SetProperties gets a reference to the given map[string]string and assigns it to the Properties field.
func (o *VersionMetaData) SetProperties(v map[string]string) {
	o.Properties = &v
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *VersionMetaData) GetGroupId() string {
	if o == nil || IsNil(o.GroupId) {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionMetaData) GetGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.GroupId) {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *VersionMetaData) HasGroupId() bool {
	if o != nil && !IsNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *VersionMetaData) SetGroupId(v string) {
	o.GroupId = &v
}

// GetContentId returns the ContentId field value
func (o *VersionMetaData) GetContentId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ContentId
}

// GetContentIdOk returns a tuple with the ContentId field value
// and a boolean to check if the value has been set.
func (o *VersionMetaData) GetContentIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContentId, true
}

// SetContentId sets field value
func (o *VersionMetaData) SetContentId(v int64) {
	o.ContentId = v
}

func (o VersionMetaData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VersionMetaData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["version"] = o.Version
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["createdBy"] = o.CreatedBy
	toSerialize["createdOn"] = o.CreatedOn
	toSerialize["type"] = o.Type
	toSerialize["globalId"] = o.GlobalId
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	toSerialize["id"] = o.Id
	if !IsNil(o.Labels) {
		toSerialize["labels"] = o.Labels
	}
	if !IsNil(o.Properties) {
		toSerialize["properties"] = o.Properties
	}
	if !IsNil(o.GroupId) {
		toSerialize["groupId"] = o.GroupId
	}
	toSerialize["contentId"] = o.ContentId
	return toSerialize, nil
}

type NullableVersionMetaData struct {
	value *VersionMetaData
	isSet bool
}

func (v NullableVersionMetaData) Get() *VersionMetaData {
	return v.value
}

func (v *NullableVersionMetaData) Set(val *VersionMetaData) {
	v.value = val
	v.isSet = true
}

func (v NullableVersionMetaData) IsSet() bool {
	return v.isSet
}

func (v *NullableVersionMetaData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVersionMetaData(val *VersionMetaData) *NullableVersionMetaData {
	return &NullableVersionMetaData{value: val, isSet: true}
}

func (v NullableVersionMetaData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVersionMetaData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


