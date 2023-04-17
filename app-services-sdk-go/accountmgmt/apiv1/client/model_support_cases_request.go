/*
Account Management Service API

Manage user subscriptions and clusters

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package accountmgmtclient

import (
	"encoding/json"
)

// checks if the SupportCasesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SupportCasesRequest{}

// SupportCasesRequest struct for SupportCasesRequest
type SupportCasesRequest struct {
	AccountNumber *string `json:"account_number,omitempty"`
	CaseLanguage *string `json:"case_language,omitempty"`
	ClusterId *string `json:"cluster_id,omitempty"`
	ClusterUuid *string `json:"cluster_uuid,omitempty"`
	ContactSsoName *string `json:"contact_sso_name,omitempty"`
	Description string `json:"description"`
	EventStreamId *string `json:"event_stream_id,omitempty"`
	OpenshiftClusterId *string `json:"openshift_cluster_id,omitempty"`
	Product *string `json:"product,omitempty"`
	Severity string `json:"severity"`
	SubscriptionId *string `json:"subscription_id,omitempty"`
	Summary string `json:"summary"`
	Version *string `json:"version,omitempty"`
}

// NewSupportCasesRequest instantiates a new SupportCasesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSupportCasesRequest(description string, severity string, summary string) *SupportCasesRequest {
	this := SupportCasesRequest{}
	this.Description = description
	var product string = "OpenShift Container Platform"
	this.Product = &product
	this.Severity = severity
	this.Summary = summary
	var version string = "4.10"
	this.Version = &version
	return &this
}

// NewSupportCasesRequestWithDefaults instantiates a new SupportCasesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSupportCasesRequestWithDefaults() *SupportCasesRequest {
	this := SupportCasesRequest{}
	var product string = "OpenShift Container Platform"
	this.Product = &product
	var version string = "4.10"
	this.Version = &version
	return &this
}

// GetAccountNumber returns the AccountNumber field value if set, zero value otherwise.
func (o *SupportCasesRequest) GetAccountNumber() string {
	if o == nil || IsNil(o.AccountNumber) {
		var ret string
		return ret
	}
	return *o.AccountNumber
}

// GetAccountNumberOk returns a tuple with the AccountNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportCasesRequest) GetAccountNumberOk() (*string, bool) {
	if o == nil || IsNil(o.AccountNumber) {
		return nil, false
	}
	return o.AccountNumber, true
}

// HasAccountNumber returns a boolean if a field has been set.
func (o *SupportCasesRequest) HasAccountNumber() bool {
	if o != nil && !IsNil(o.AccountNumber) {
		return true
	}

	return false
}

// SetAccountNumber gets a reference to the given string and assigns it to the AccountNumber field.
func (o *SupportCasesRequest) SetAccountNumber(v string) {
	o.AccountNumber = &v
}

// GetCaseLanguage returns the CaseLanguage field value if set, zero value otherwise.
func (o *SupportCasesRequest) GetCaseLanguage() string {
	if o == nil || IsNil(o.CaseLanguage) {
		var ret string
		return ret
	}
	return *o.CaseLanguage
}

// GetCaseLanguageOk returns a tuple with the CaseLanguage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportCasesRequest) GetCaseLanguageOk() (*string, bool) {
	if o == nil || IsNil(o.CaseLanguage) {
		return nil, false
	}
	return o.CaseLanguage, true
}

// HasCaseLanguage returns a boolean if a field has been set.
func (o *SupportCasesRequest) HasCaseLanguage() bool {
	if o != nil && !IsNil(o.CaseLanguage) {
		return true
	}

	return false
}

// SetCaseLanguage gets a reference to the given string and assigns it to the CaseLanguage field.
func (o *SupportCasesRequest) SetCaseLanguage(v string) {
	o.CaseLanguage = &v
}

// GetClusterId returns the ClusterId field value if set, zero value otherwise.
func (o *SupportCasesRequest) GetClusterId() string {
	if o == nil || IsNil(o.ClusterId) {
		var ret string
		return ret
	}
	return *o.ClusterId
}

// GetClusterIdOk returns a tuple with the ClusterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportCasesRequest) GetClusterIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClusterId) {
		return nil, false
	}
	return o.ClusterId, true
}

// HasClusterId returns a boolean if a field has been set.
func (o *SupportCasesRequest) HasClusterId() bool {
	if o != nil && !IsNil(o.ClusterId) {
		return true
	}

	return false
}

// SetClusterId gets a reference to the given string and assigns it to the ClusterId field.
func (o *SupportCasesRequest) SetClusterId(v string) {
	o.ClusterId = &v
}

// GetClusterUuid returns the ClusterUuid field value if set, zero value otherwise.
func (o *SupportCasesRequest) GetClusterUuid() string {
	if o == nil || IsNil(o.ClusterUuid) {
		var ret string
		return ret
	}
	return *o.ClusterUuid
}

// GetClusterUuidOk returns a tuple with the ClusterUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportCasesRequest) GetClusterUuidOk() (*string, bool) {
	if o == nil || IsNil(o.ClusterUuid) {
		return nil, false
	}
	return o.ClusterUuid, true
}

// HasClusterUuid returns a boolean if a field has been set.
func (o *SupportCasesRequest) HasClusterUuid() bool {
	if o != nil && !IsNil(o.ClusterUuid) {
		return true
	}

	return false
}

// SetClusterUuid gets a reference to the given string and assigns it to the ClusterUuid field.
func (o *SupportCasesRequest) SetClusterUuid(v string) {
	o.ClusterUuid = &v
}

// GetContactSsoName returns the ContactSsoName field value if set, zero value otherwise.
func (o *SupportCasesRequest) GetContactSsoName() string {
	if o == nil || IsNil(o.ContactSsoName) {
		var ret string
		return ret
	}
	return *o.ContactSsoName
}

// GetContactSsoNameOk returns a tuple with the ContactSsoName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportCasesRequest) GetContactSsoNameOk() (*string, bool) {
	if o == nil || IsNil(o.ContactSsoName) {
		return nil, false
	}
	return o.ContactSsoName, true
}

// HasContactSsoName returns a boolean if a field has been set.
func (o *SupportCasesRequest) HasContactSsoName() bool {
	if o != nil && !IsNil(o.ContactSsoName) {
		return true
	}

	return false
}

// SetContactSsoName gets a reference to the given string and assigns it to the ContactSsoName field.
func (o *SupportCasesRequest) SetContactSsoName(v string) {
	o.ContactSsoName = &v
}

// GetDescription returns the Description field value
func (o *SupportCasesRequest) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *SupportCasesRequest) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *SupportCasesRequest) SetDescription(v string) {
	o.Description = v
}

// GetEventStreamId returns the EventStreamId field value if set, zero value otherwise.
func (o *SupportCasesRequest) GetEventStreamId() string {
	if o == nil || IsNil(o.EventStreamId) {
		var ret string
		return ret
	}
	return *o.EventStreamId
}

// GetEventStreamIdOk returns a tuple with the EventStreamId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportCasesRequest) GetEventStreamIdOk() (*string, bool) {
	if o == nil || IsNil(o.EventStreamId) {
		return nil, false
	}
	return o.EventStreamId, true
}

// HasEventStreamId returns a boolean if a field has been set.
func (o *SupportCasesRequest) HasEventStreamId() bool {
	if o != nil && !IsNil(o.EventStreamId) {
		return true
	}

	return false
}

// SetEventStreamId gets a reference to the given string and assigns it to the EventStreamId field.
func (o *SupportCasesRequest) SetEventStreamId(v string) {
	o.EventStreamId = &v
}

// GetOpenshiftClusterId returns the OpenshiftClusterId field value if set, zero value otherwise.
func (o *SupportCasesRequest) GetOpenshiftClusterId() string {
	if o == nil || IsNil(o.OpenshiftClusterId) {
		var ret string
		return ret
	}
	return *o.OpenshiftClusterId
}

// GetOpenshiftClusterIdOk returns a tuple with the OpenshiftClusterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportCasesRequest) GetOpenshiftClusterIdOk() (*string, bool) {
	if o == nil || IsNil(o.OpenshiftClusterId) {
		return nil, false
	}
	return o.OpenshiftClusterId, true
}

// HasOpenshiftClusterId returns a boolean if a field has been set.
func (o *SupportCasesRequest) HasOpenshiftClusterId() bool {
	if o != nil && !IsNil(o.OpenshiftClusterId) {
		return true
	}

	return false
}

// SetOpenshiftClusterId gets a reference to the given string and assigns it to the OpenshiftClusterId field.
func (o *SupportCasesRequest) SetOpenshiftClusterId(v string) {
	o.OpenshiftClusterId = &v
}

// GetProduct returns the Product field value if set, zero value otherwise.
func (o *SupportCasesRequest) GetProduct() string {
	if o == nil || IsNil(o.Product) {
		var ret string
		return ret
	}
	return *o.Product
}

// GetProductOk returns a tuple with the Product field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportCasesRequest) GetProductOk() (*string, bool) {
	if o == nil || IsNil(o.Product) {
		return nil, false
	}
	return o.Product, true
}

// HasProduct returns a boolean if a field has been set.
func (o *SupportCasesRequest) HasProduct() bool {
	if o != nil && !IsNil(o.Product) {
		return true
	}

	return false
}

// SetProduct gets a reference to the given string and assigns it to the Product field.
func (o *SupportCasesRequest) SetProduct(v string) {
	o.Product = &v
}

// GetSeverity returns the Severity field value
func (o *SupportCasesRequest) GetSeverity() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value
// and a boolean to check if the value has been set.
func (o *SupportCasesRequest) GetSeverityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Severity, true
}

// SetSeverity sets field value
func (o *SupportCasesRequest) SetSeverity(v string) {
	o.Severity = v
}

// GetSubscriptionId returns the SubscriptionId field value if set, zero value otherwise.
func (o *SupportCasesRequest) GetSubscriptionId() string {
	if o == nil || IsNil(o.SubscriptionId) {
		var ret string
		return ret
	}
	return *o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportCasesRequest) GetSubscriptionIdOk() (*string, bool) {
	if o == nil || IsNil(o.SubscriptionId) {
		return nil, false
	}
	return o.SubscriptionId, true
}

// HasSubscriptionId returns a boolean if a field has been set.
func (o *SupportCasesRequest) HasSubscriptionId() bool {
	if o != nil && !IsNil(o.SubscriptionId) {
		return true
	}

	return false
}

// SetSubscriptionId gets a reference to the given string and assigns it to the SubscriptionId field.
func (o *SupportCasesRequest) SetSubscriptionId(v string) {
	o.SubscriptionId = &v
}

// GetSummary returns the Summary field value
func (o *SupportCasesRequest) GetSummary() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value
// and a boolean to check if the value has been set.
func (o *SupportCasesRequest) GetSummaryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Summary, true
}

// SetSummary sets field value
func (o *SupportCasesRequest) SetSummary(v string) {
	o.Summary = v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *SupportCasesRequest) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportCasesRequest) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *SupportCasesRequest) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *SupportCasesRequest) SetVersion(v string) {
	o.Version = &v
}

func (o SupportCasesRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SupportCasesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccountNumber) {
		toSerialize["account_number"] = o.AccountNumber
	}
	if !IsNil(o.CaseLanguage) {
		toSerialize["case_language"] = o.CaseLanguage
	}
	if !IsNil(o.ClusterId) {
		toSerialize["cluster_id"] = o.ClusterId
	}
	if !IsNil(o.ClusterUuid) {
		toSerialize["cluster_uuid"] = o.ClusterUuid
	}
	if !IsNil(o.ContactSsoName) {
		toSerialize["contact_sso_name"] = o.ContactSsoName
	}
	toSerialize["description"] = o.Description
	if !IsNil(o.EventStreamId) {
		toSerialize["event_stream_id"] = o.EventStreamId
	}
	if !IsNil(o.OpenshiftClusterId) {
		toSerialize["openshift_cluster_id"] = o.OpenshiftClusterId
	}
	if !IsNil(o.Product) {
		toSerialize["product"] = o.Product
	}
	toSerialize["severity"] = o.Severity
	if !IsNil(o.SubscriptionId) {
		toSerialize["subscription_id"] = o.SubscriptionId
	}
	toSerialize["summary"] = o.Summary
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return toSerialize, nil
}

type NullableSupportCasesRequest struct {
	value *SupportCasesRequest
	isSet bool
}

func (v NullableSupportCasesRequest) Get() *SupportCasesRequest {
	return v.value
}

func (v *NullableSupportCasesRequest) Set(val *SupportCasesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSupportCasesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSupportCasesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSupportCasesRequest(val *SupportCasesRequest) *NullableSupportCasesRequest {
	return &NullableSupportCasesRequest{value: val, isSet: true}
}

func (v NullableSupportCasesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSupportCasesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


