/*
Kafka Management API

Kafka Management API is a REST API to manage Kafka instances

API version: 1.16.0
Contact: rhosak-support@redhat.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package kafkamgmtclient

import (
	"encoding/json"
	"time"
)

// checks if the ServiceAccountListAllOfItemsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceAccountListAllOfItemsInner{}

// ServiceAccountListAllOfItemsInner struct for ServiceAccountListAllOfItemsInner
type ServiceAccountListAllOfItemsInner struct {
	// server generated unique id of the service account
	Id string `json:"id"`
	Kind string `json:"kind"`
	Href string `json:"href"`
	// client id of the service account
	ClientId *string `json:"client_id,omitempty"`
	// name of the service account
	Name *string `json:"name,omitempty"`
	// owner of the service account
	// Deprecated
	Owner *string `json:"owner,omitempty"`
	// service account created by the user
	CreatedBy *string `json:"created_by,omitempty"`
	// service account creation timestamp
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// description of the service account
	Description *string `json:"description,omitempty"`
}

// NewServiceAccountListAllOfItemsInner instantiates a new ServiceAccountListAllOfItemsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceAccountListAllOfItemsInner(id string, kind string, href string) *ServiceAccountListAllOfItemsInner {
	this := ServiceAccountListAllOfItemsInner{}
	this.Id = id
	this.Kind = kind
	this.Href = href
	return &this
}

// NewServiceAccountListAllOfItemsInnerWithDefaults instantiates a new ServiceAccountListAllOfItemsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceAccountListAllOfItemsInnerWithDefaults() *ServiceAccountListAllOfItemsInner {
	this := ServiceAccountListAllOfItemsInner{}
	return &this
}

// GetId returns the Id field value
func (o *ServiceAccountListAllOfItemsInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ServiceAccountListAllOfItemsInner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ServiceAccountListAllOfItemsInner) SetId(v string) {
	o.Id = v
}

// GetKind returns the Kind field value
func (o *ServiceAccountListAllOfItemsInner) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *ServiceAccountListAllOfItemsInner) GetKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *ServiceAccountListAllOfItemsInner) SetKind(v string) {
	o.Kind = v
}

// GetHref returns the Href field value
func (o *ServiceAccountListAllOfItemsInner) GetHref() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Href
}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
func (o *ServiceAccountListAllOfItemsInner) GetHrefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Href, true
}

// SetHref sets field value
func (o *ServiceAccountListAllOfItemsInner) SetHref(v string) {
	o.Href = v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *ServiceAccountListAllOfItemsInner) GetClientId() string {
	if o == nil || IsNil(o.ClientId) {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccountListAllOfItemsInner) GetClientIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClientId) {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *ServiceAccountListAllOfItemsInner) HasClientId() bool {
	if o != nil && !IsNil(o.ClientId) {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *ServiceAccountListAllOfItemsInner) SetClientId(v string) {
	o.ClientId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ServiceAccountListAllOfItemsInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccountListAllOfItemsInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ServiceAccountListAllOfItemsInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ServiceAccountListAllOfItemsInner) SetName(v string) {
	o.Name = &v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
// Deprecated
func (o *ServiceAccountListAllOfItemsInner) GetOwner() string {
	if o == nil || IsNil(o.Owner) {
		var ret string
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *ServiceAccountListAllOfItemsInner) GetOwnerOk() (*string, bool) {
	if o == nil || IsNil(o.Owner) {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *ServiceAccountListAllOfItemsInner) HasOwner() bool {
	if o != nil && !IsNil(o.Owner) {
		return true
	}

	return false
}

// SetOwner gets a reference to the given string and assigns it to the Owner field.
// Deprecated
func (o *ServiceAccountListAllOfItemsInner) SetOwner(v string) {
	o.Owner = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *ServiceAccountListAllOfItemsInner) GetCreatedBy() string {
	if o == nil || IsNil(o.CreatedBy) {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccountListAllOfItemsInner) GetCreatedByOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedBy) {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *ServiceAccountListAllOfItemsInner) HasCreatedBy() bool {
	if o != nil && !IsNil(o.CreatedBy) {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *ServiceAccountListAllOfItemsInner) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ServiceAccountListAllOfItemsInner) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccountListAllOfItemsInner) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ServiceAccountListAllOfItemsInner) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *ServiceAccountListAllOfItemsInner) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ServiceAccountListAllOfItemsInner) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccountListAllOfItemsInner) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ServiceAccountListAllOfItemsInner) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ServiceAccountListAllOfItemsInner) SetDescription(v string) {
	o.Description = &v
}

func (o ServiceAccountListAllOfItemsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceAccountListAllOfItemsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["kind"] = o.Kind
	toSerialize["href"] = o.Href
	if !IsNil(o.ClientId) {
		toSerialize["client_id"] = o.ClientId
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Owner) {
		toSerialize["owner"] = o.Owner
	}
	if !IsNil(o.CreatedBy) {
		toSerialize["created_by"] = o.CreatedBy
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return toSerialize, nil
}

type NullableServiceAccountListAllOfItemsInner struct {
	value *ServiceAccountListAllOfItemsInner
	isSet bool
}

func (v NullableServiceAccountListAllOfItemsInner) Get() *ServiceAccountListAllOfItemsInner {
	return v.value
}

func (v *NullableServiceAccountListAllOfItemsInner) Set(val *ServiceAccountListAllOfItemsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceAccountListAllOfItemsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceAccountListAllOfItemsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceAccountListAllOfItemsInner(val *ServiceAccountListAllOfItemsInner) *NullableServiceAccountListAllOfItemsInner {
	return &NullableServiceAccountListAllOfItemsInner{value: val, isSet: true}
}

func (v NullableServiceAccountListAllOfItemsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceAccountListAllOfItemsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


