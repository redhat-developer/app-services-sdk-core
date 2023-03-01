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
	"time"
)

// checks if the Registry type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Registry{}

// Registry struct for Registry
type Registry struct {
	Id string `json:"id"`
	Kind string `json:"kind"`
	Href string `json:"href"`
	Status RegistryStatusValue `json:"status"`
	RegistryUrl *string `json:"registryUrl,omitempty"`
	BrowserUrl *string `json:"browserUrl,omitempty"`
	// User-defined Registry instance name. Does not have to be unique.
	Name string `json:"name"`
	// Identifier of a multi-tenant deployment, where this Service Registry instance resides.
	RegistryDeploymentId *int32 `json:"registryDeploymentId,omitempty"`
	// Registry instance owner.
	Owner string `json:"owner"`
	// Description of the Registry instance.
	Description *string `json:"description,omitempty"`
	// ISO 8601 UTC timestamp.
	CreatedAt time.Time `json:"created_at"`
	// ISO 8601 UTC timestamp.
	UpdatedAt time.Time `json:"updated_at"`
	InstanceType RegistryInstanceTypeValue `json:"instance_type"`
}

// NewRegistry instantiates a new Registry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegistry(id string, kind string, href string, status RegistryStatusValue, name string, owner string, createdAt time.Time, updatedAt time.Time, instanceType RegistryInstanceTypeValue) *Registry {
	this := Registry{}
	this.Id = id
	this.Kind = kind
	this.Href = href
	this.Status = status
	this.Name = name
	this.Owner = owner
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	this.InstanceType = instanceType
	return &this
}

// NewRegistryWithDefaults instantiates a new Registry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegistryWithDefaults() *Registry {
	this := Registry{}
	return &this
}

// GetId returns the Id field value
func (o *Registry) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Registry) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Registry) SetId(v string) {
	o.Id = v
}

// GetKind returns the Kind field value
func (o *Registry) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *Registry) GetKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *Registry) SetKind(v string) {
	o.Kind = v
}

// GetHref returns the Href field value
func (o *Registry) GetHref() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Href
}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
func (o *Registry) GetHrefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Href, true
}

// SetHref sets field value
func (o *Registry) SetHref(v string) {
	o.Href = v
}

// GetStatus returns the Status field value
func (o *Registry) GetStatus() RegistryStatusValue {
	if o == nil {
		var ret RegistryStatusValue
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Registry) GetStatusOk() (*RegistryStatusValue, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *Registry) SetStatus(v RegistryStatusValue) {
	o.Status = v
}

// GetRegistryUrl returns the RegistryUrl field value if set, zero value otherwise.
func (o *Registry) GetRegistryUrl() string {
	if o == nil || IsNil(o.RegistryUrl) {
		var ret string
		return ret
	}
	return *o.RegistryUrl
}

// GetRegistryUrlOk returns a tuple with the RegistryUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Registry) GetRegistryUrlOk() (*string, bool) {
	if o == nil || IsNil(o.RegistryUrl) {
		return nil, false
	}
	return o.RegistryUrl, true
}

// HasRegistryUrl returns a boolean if a field has been set.
func (o *Registry) HasRegistryUrl() bool {
	if o != nil && !IsNil(o.RegistryUrl) {
		return true
	}

	return false
}

// SetRegistryUrl gets a reference to the given string and assigns it to the RegistryUrl field.
func (o *Registry) SetRegistryUrl(v string) {
	o.RegistryUrl = &v
}

// GetBrowserUrl returns the BrowserUrl field value if set, zero value otherwise.
func (o *Registry) GetBrowserUrl() string {
	if o == nil || IsNil(o.BrowserUrl) {
		var ret string
		return ret
	}
	return *o.BrowserUrl
}

// GetBrowserUrlOk returns a tuple with the BrowserUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Registry) GetBrowserUrlOk() (*string, bool) {
	if o == nil || IsNil(o.BrowserUrl) {
		return nil, false
	}
	return o.BrowserUrl, true
}

// HasBrowserUrl returns a boolean if a field has been set.
func (o *Registry) HasBrowserUrl() bool {
	if o != nil && !IsNil(o.BrowserUrl) {
		return true
	}

	return false
}

// SetBrowserUrl gets a reference to the given string and assigns it to the BrowserUrl field.
func (o *Registry) SetBrowserUrl(v string) {
	o.BrowserUrl = &v
}

// GetName returns the Name field value
func (o *Registry) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Registry) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Registry) SetName(v string) {
	o.Name = v
}

// GetRegistryDeploymentId returns the RegistryDeploymentId field value if set, zero value otherwise.
func (o *Registry) GetRegistryDeploymentId() int32 {
	if o == nil || IsNil(o.RegistryDeploymentId) {
		var ret int32
		return ret
	}
	return *o.RegistryDeploymentId
}

// GetRegistryDeploymentIdOk returns a tuple with the RegistryDeploymentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Registry) GetRegistryDeploymentIdOk() (*int32, bool) {
	if o == nil || IsNil(o.RegistryDeploymentId) {
		return nil, false
	}
	return o.RegistryDeploymentId, true
}

// HasRegistryDeploymentId returns a boolean if a field has been set.
func (o *Registry) HasRegistryDeploymentId() bool {
	if o != nil && !IsNil(o.RegistryDeploymentId) {
		return true
	}

	return false
}

// SetRegistryDeploymentId gets a reference to the given int32 and assigns it to the RegistryDeploymentId field.
func (o *Registry) SetRegistryDeploymentId(v int32) {
	o.RegistryDeploymentId = &v
}

// GetOwner returns the Owner field value
func (o *Registry) GetOwner() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value
// and a boolean to check if the value has been set.
func (o *Registry) GetOwnerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Owner, true
}

// SetOwner sets field value
func (o *Registry) SetOwner(v string) {
	o.Owner = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Registry) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Registry) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Registry) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Registry) SetDescription(v string) {
	o.Description = &v
}

// GetCreatedAt returns the CreatedAt field value
func (o *Registry) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Registry) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *Registry) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *Registry) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *Registry) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *Registry) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetInstanceType returns the InstanceType field value
func (o *Registry) GetInstanceType() RegistryInstanceTypeValue {
	if o == nil {
		var ret RegistryInstanceTypeValue
		return ret
	}

	return o.InstanceType
}

// GetInstanceTypeOk returns a tuple with the InstanceType field value
// and a boolean to check if the value has been set.
func (o *Registry) GetInstanceTypeOk() (*RegistryInstanceTypeValue, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InstanceType, true
}

// SetInstanceType sets field value
func (o *Registry) SetInstanceType(v RegistryInstanceTypeValue) {
	o.InstanceType = v
}

func (o Registry) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Registry) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["kind"] = o.Kind
	toSerialize["href"] = o.Href
	toSerialize["status"] = o.Status
	if !IsNil(o.RegistryUrl) {
		toSerialize["registryUrl"] = o.RegistryUrl
	}
	if !IsNil(o.BrowserUrl) {
		toSerialize["browserUrl"] = o.BrowserUrl
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.RegistryDeploymentId) {
		toSerialize["registryDeploymentId"] = o.RegistryDeploymentId
	}
	toSerialize["owner"] = o.Owner
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["updated_at"] = o.UpdatedAt
	toSerialize["instance_type"] = o.InstanceType
	return toSerialize, nil
}

type NullableRegistry struct {
	value *Registry
	isSet bool
}

func (v NullableRegistry) Get() *Registry {
	return v.value
}

func (v *NullableRegistry) Set(val *Registry) {
	v.value = val
	v.isSet = true
}

func (v NullableRegistry) IsSet() bool {
	return v.isSet
}

func (v *NullableRegistry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegistry(val *Registry) *NullableRegistry {
	return &NullableRegistry{value: val, isSet: true}
}

func (v NullableRegistry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegistry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


