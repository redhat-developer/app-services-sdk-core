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

// checks if the ListResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListResponse{}

// ListResponse struct for ListResponse
type ListResponse struct {
	Kind string `json:"kind"`
	Items []map[string]interface{} `json:"items,omitempty"`
}

// NewListResponse instantiates a new ListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListResponse(kind string) *ListResponse {
	this := ListResponse{}
	this.Kind = kind
	return &this
}

// NewListResponseWithDefaults instantiates a new ListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListResponseWithDefaults() *ListResponse {
	this := ListResponse{}
	return &this
}

// GetKind returns the Kind field value
func (o *ListResponse) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *ListResponse) GetKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *ListResponse) SetKind(v string) {
	o.Kind = v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *ListResponse) GetItems() []map[string]interface{} {
	if o == nil || IsNil(o.Items) {
		var ret []map[string]interface{}
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListResponse) GetItemsOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *ListResponse) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []map[string]interface{} and assigns it to the Items field.
func (o *ListResponse) SetItems(v []map[string]interface{}) {
	o.Items = v
}

func (o ListResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["kind"] = o.Kind
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	return toSerialize, nil
}

type NullableListResponse struct {
	value *ListResponse
	isSet bool
}

func (v NullableListResponse) Get() *ListResponse {
	return v.value
}

func (v *NullableListResponse) Set(val *ListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListResponse(val *ListResponse) *NullableListResponse {
	return &NullableListResponse{value: val, isSet: true}
}

func (v NullableListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


