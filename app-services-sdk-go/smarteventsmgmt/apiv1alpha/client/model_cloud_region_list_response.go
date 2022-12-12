/*
 * Red Hat Openshift SmartEvents Fleet Manager V2
 *
 * The API exposed by the fleet manager of the SmartEvents service.
 *
 * API version: 0.0.1
 * Contact: openbridge-dev@redhat.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package smarteventsmgmtclient

import (
	"encoding/json"
)

// CloudRegionListResponse struct for CloudRegionListResponse
type CloudRegionListResponse struct {
	Kind string `json:"kind"`
	Items *[]CloudRegionResponse `json:"items,omitempty"`
	Page int64 `json:"page"`
	Size int64 `json:"size"`
	Total int64 `json:"total"`
}

// NewCloudRegionListResponse instantiates a new CloudRegionListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudRegionListResponse(kind string, page int64, size int64, total int64) *CloudRegionListResponse {
	this := CloudRegionListResponse{}
	this.Kind = kind
	this.Page = page
	this.Size = size
	this.Total = total
	return &this
}

// NewCloudRegionListResponseWithDefaults instantiates a new CloudRegionListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudRegionListResponseWithDefaults() *CloudRegionListResponse {
	this := CloudRegionListResponse{}
	return &this
}

// GetKind returns the Kind field value
func (o *CloudRegionListResponse) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *CloudRegionListResponse) GetKindOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *CloudRegionListResponse) SetKind(v string) {
	o.Kind = v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *CloudRegionListResponse) GetItems() []CloudRegionResponse {
	if o == nil || o.Items == nil {
		var ret []CloudRegionResponse
		return ret
	}
	return *o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudRegionListResponse) GetItemsOk() (*[]CloudRegionResponse, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *CloudRegionListResponse) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []CloudRegionResponse and assigns it to the Items field.
func (o *CloudRegionListResponse) SetItems(v []CloudRegionResponse) {
	o.Items = &v
}

// GetPage returns the Page field value
func (o *CloudRegionListResponse) GetPage() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Page
}

// GetPageOk returns a tuple with the Page field value
// and a boolean to check if the value has been set.
func (o *CloudRegionListResponse) GetPageOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Page, true
}

// SetPage sets field value
func (o *CloudRegionListResponse) SetPage(v int64) {
	o.Page = v
}

// GetSize returns the Size field value
func (o *CloudRegionListResponse) GetSize() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *CloudRegionListResponse) GetSizeOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *CloudRegionListResponse) SetSize(v int64) {
	o.Size = v
}

// GetTotal returns the Total field value
func (o *CloudRegionListResponse) GetTotal() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *CloudRegionListResponse) GetTotalOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *CloudRegionListResponse) SetTotal(v int64) {
	o.Total = v
}

func (o CloudRegionListResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["kind"] = o.Kind
	}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	if true {
		toSerialize["page"] = o.Page
	}
	if true {
		toSerialize["size"] = o.Size
	}
	if true {
		toSerialize["total"] = o.Total
	}
	return json.Marshal(toSerialize)
}

type NullableCloudRegionListResponse struct {
	value *CloudRegionListResponse
	isSet bool
}

func (v NullableCloudRegionListResponse) Get() *CloudRegionListResponse {
	return v.value
}

func (v *NullableCloudRegionListResponse) Set(val *CloudRegionListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudRegionListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudRegionListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudRegionListResponse(val *CloudRegionListResponse) *NullableCloudRegionListResponse {
	return &NullableCloudRegionListResponse{value: val, isSet: true}
}

func (v NullableCloudRegionListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudRegionListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

