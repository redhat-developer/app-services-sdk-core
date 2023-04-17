/*
Account Management Service API

Manage user subscriptions and clusters

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package accountmgmtclient

import (
	"encoding/json"
	"time"
)

// checks if the SubscriptionPatchRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionPatchRequest{}

// SubscriptionPatchRequest struct for SubscriptionPatchRequest
type SubscriptionPatchRequest struct {
	BillingExpirationDate *time.Time `json:"billing_expiration_date,omitempty"`
	CloudAccountId *string `json:"cloud_account_id,omitempty"`
	CloudProviderId *string `json:"cloud_provider_id,omitempty"`
	ClusterBillingModel *string `json:"cluster_billing_model,omitempty"`
	ClusterId *string `json:"cluster_id,omitempty"`
	ConsoleUrl *string `json:"console_url,omitempty"`
	ConsumerUuid *string `json:"consumer_uuid,omitempty"`
	CpuTotal *int32 `json:"cpu_total,omitempty"`
	CreatorId *string `json:"creator_id,omitempty"`
	DisplayName *string `json:"display_name,omitempty"`
	ExternalClusterId *string `json:"external_cluster_id,omitempty"`
	Managed *bool `json:"managed,omitempty"`
	OrganizationId *string `json:"organization_id,omitempty"`
	PlanId *string `json:"plan_id,omitempty"`
	ProductBundle *string `json:"product_bundle,omitempty"`
	Provenance *string `json:"provenance,omitempty"`
	RegionId *string `json:"region_id,omitempty"`
	Released *bool `json:"released,omitempty"`
	ServiceLevel *string `json:"service_level,omitempty"`
	SocketTotal *int32 `json:"socket_total,omitempty"`
	Status *string `json:"status,omitempty"`
	SupportLevel *string `json:"support_level,omitempty"`
	SystemUnits *string `json:"system_units,omitempty"`
	TrialEndDate *time.Time `json:"trial_end_date,omitempty"`
	Usage *string `json:"usage,omitempty"`
}

// NewSubscriptionPatchRequest instantiates a new SubscriptionPatchRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionPatchRequest() *SubscriptionPatchRequest {
	this := SubscriptionPatchRequest{}
	return &this
}

// NewSubscriptionPatchRequestWithDefaults instantiates a new SubscriptionPatchRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionPatchRequestWithDefaults() *SubscriptionPatchRequest {
	this := SubscriptionPatchRequest{}
	return &this
}

// GetBillingExpirationDate returns the BillingExpirationDate field value if set, zero value otherwise.
func (o *SubscriptionPatchRequest) GetBillingExpirationDate() time.Time {
	if o == nil || IsNil(o.BillingExpirationDate) {
		var ret time.Time
		return ret
	}
	return *o.BillingExpirationDate
}

// GetBillingExpirationDateOk returns a tuple with the BillingExpirationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPatchRequest) GetBillingExpirationDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.BillingExpirationDate) {
		return nil, false
	}
	return o.BillingExpirationDate, true
}

// HasBillingExpirationDate returns a boolean if a field has been set.
func (o *SubscriptionPatchRequest) HasBillingExpirationDate() bool {
	if o != nil && !IsNil(o.BillingExpirationDate) {
		return true
	}

	return false
}

// SetBillingExpirationDate gets a reference to the given time.Time and assigns it to the BillingExpirationDate field.
func (o *SubscriptionPatchRequest) SetBillingExpirationDate(v time.Time) {
	o.BillingExpirationDate = &v
}

// GetCloudAccountId returns the CloudAccountId field value if set, zero value otherwise.
func (o *SubscriptionPatchRequest) GetCloudAccountId() string {
	if o == nil || IsNil(o.CloudAccountId) {
		var ret string
		return ret
	}
	return *o.CloudAccountId
}

// GetCloudAccountIdOk returns a tuple with the CloudAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPatchRequest) GetCloudAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.CloudAccountId) {
		return nil, false
	}
	return o.CloudAccountId, true
}

// HasCloudAccountId returns a boolean if a field has been set.
func (o *SubscriptionPatchRequest) HasCloudAccountId() bool {
	if o != nil && !IsNil(o.CloudAccountId) {
		return true
	}

	return false
}

// SetCloudAccountId gets a reference to the given string and assigns it to the CloudAccountId field.
func (o *SubscriptionPatchRequest) SetCloudAccountId(v string) {
	o.CloudAccountId = &v
}

// GetCloudProviderId returns the CloudProviderId field value if set, zero value otherwise.
func (o *SubscriptionPatchRequest) GetCloudProviderId() string {
	if o == nil || IsNil(o.CloudProviderId) {
		var ret string
		return ret
	}
	return *o.CloudProviderId
}

// GetCloudProviderIdOk returns a tuple with the CloudProviderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPatchRequest) GetCloudProviderIdOk() (*string, bool) {
	if o == nil || IsNil(o.CloudProviderId) {
		return nil, false
	}
	return o.CloudProviderId, true
}

// HasCloudProviderId returns a boolean if a field has been set.
func (o *SubscriptionPatchRequest) HasCloudProviderId() bool {
	if o != nil && !IsNil(o.CloudProviderId) {
		return true
	}

	return false
}

// SetCloudProviderId gets a reference to the given string and assigns it to the CloudProviderId field.
func (o *SubscriptionPatchRequest) SetCloudProviderId(v string) {
	o.CloudProviderId = &v
}

// GetClusterBillingModel returns the ClusterBillingModel field value if set, zero value otherwise.
func (o *SubscriptionPatchRequest) GetClusterBillingModel() string {
	if o == nil || IsNil(o.ClusterBillingModel) {
		var ret string
		return ret
	}
	return *o.ClusterBillingModel
}

// GetClusterBillingModelOk returns a tuple with the ClusterBillingModel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPatchRequest) GetClusterBillingModelOk() (*string, bool) {
	if o == nil || IsNil(o.ClusterBillingModel) {
		return nil, false
	}
	return o.ClusterBillingModel, true
}

// HasClusterBillingModel returns a boolean if a field has been set.
func (o *SubscriptionPatchRequest) HasClusterBillingModel() bool {
	if o != nil && !IsNil(o.ClusterBillingModel) {
		return true
	}

	return false
}

// SetClusterBillingModel gets a reference to the given string and assigns it to the ClusterBillingModel field.
func (o *SubscriptionPatchRequest) SetClusterBillingModel(v string) {
	o.ClusterBillingModel = &v
}

// GetClusterId returns the ClusterId field value if set, zero value otherwise.
func (o *SubscriptionPatchRequest) GetClusterId() string {
	if o == nil || IsNil(o.ClusterId) {
		var ret string
		return ret
	}
	return *o.ClusterId
}

// GetClusterIdOk returns a tuple with the ClusterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPatchRequest) GetClusterIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClusterId) {
		return nil, false
	}
	return o.ClusterId, true
}

// HasClusterId returns a boolean if a field has been set.
func (o *SubscriptionPatchRequest) HasClusterId() bool {
	if o != nil && !IsNil(o.ClusterId) {
		return true
	}

	return false
}

// SetClusterId gets a reference to the given string and assigns it to the ClusterId field.
func (o *SubscriptionPatchRequest) SetClusterId(v string) {
	o.ClusterId = &v
}

// GetConsoleUrl returns the ConsoleUrl field value if set, zero value otherwise.
func (o *SubscriptionPatchRequest) GetConsoleUrl() string {
	if o == nil || IsNil(o.ConsoleUrl) {
		var ret string
		return ret
	}
	return *o.ConsoleUrl
}

// GetConsoleUrlOk returns a tuple with the ConsoleUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPatchRequest) GetConsoleUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ConsoleUrl) {
		return nil, false
	}
	return o.ConsoleUrl, true
}

// HasConsoleUrl returns a boolean if a field has been set.
func (o *SubscriptionPatchRequest) HasConsoleUrl() bool {
	if o != nil && !IsNil(o.ConsoleUrl) {
		return true
	}

	return false
}

// SetConsoleUrl gets a reference to the given string and assigns it to the ConsoleUrl field.
func (o *SubscriptionPatchRequest) SetConsoleUrl(v string) {
	o.ConsoleUrl = &v
}

// GetConsumerUuid returns the ConsumerUuid field value if set, zero value otherwise.
func (o *SubscriptionPatchRequest) GetConsumerUuid() string {
	if o == nil || IsNil(o.ConsumerUuid) {
		var ret string
		return ret
	}
	return *o.ConsumerUuid
}

// GetConsumerUuidOk returns a tuple with the ConsumerUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPatchRequest) GetConsumerUuidOk() (*string, bool) {
	if o == nil || IsNil(o.ConsumerUuid) {
		return nil, false
	}
	return o.ConsumerUuid, true
}

// HasConsumerUuid returns a boolean if a field has been set.
func (o *SubscriptionPatchRequest) HasConsumerUuid() bool {
	if o != nil && !IsNil(o.ConsumerUuid) {
		return true
	}

	return false
}

// SetConsumerUuid gets a reference to the given string and assigns it to the ConsumerUuid field.
func (o *SubscriptionPatchRequest) SetConsumerUuid(v string) {
	o.ConsumerUuid = &v
}

// GetCpuTotal returns the CpuTotal field value if set, zero value otherwise.
func (o *SubscriptionPatchRequest) GetCpuTotal() int32 {
	if o == nil || IsNil(o.CpuTotal) {
		var ret int32
		return ret
	}
	return *o.CpuTotal
}

// GetCpuTotalOk returns a tuple with the CpuTotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPatchRequest) GetCpuTotalOk() (*int32, bool) {
	if o == nil || IsNil(o.CpuTotal) {
		return nil, false
	}
	return o.CpuTotal, true
}

// HasCpuTotal returns a boolean if a field has been set.
func (o *SubscriptionPatchRequest) HasCpuTotal() bool {
	if o != nil && !IsNil(o.CpuTotal) {
		return true
	}

	return false
}

// SetCpuTotal gets a reference to the given int32 and assigns it to the CpuTotal field.
func (o *SubscriptionPatchRequest) SetCpuTotal(v int32) {
	o.CpuTotal = &v
}

// GetCreatorId returns the CreatorId field value if set, zero value otherwise.
func (o *SubscriptionPatchRequest) GetCreatorId() string {
	if o == nil || IsNil(o.CreatorId) {
		var ret string
		return ret
	}
	return *o.CreatorId
}

// GetCreatorIdOk returns a tuple with the CreatorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPatchRequest) GetCreatorIdOk() (*string, bool) {
	if o == nil || IsNil(o.CreatorId) {
		return nil, false
	}
	return o.CreatorId, true
}

// HasCreatorId returns a boolean if a field has been set.
func (o *SubscriptionPatchRequest) HasCreatorId() bool {
	if o != nil && !IsNil(o.CreatorId) {
		return true
	}

	return false
}

// SetCreatorId gets a reference to the given string and assigns it to the CreatorId field.
func (o *SubscriptionPatchRequest) SetCreatorId(v string) {
	o.CreatorId = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *SubscriptionPatchRequest) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPatchRequest) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *SubscriptionPatchRequest) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *SubscriptionPatchRequest) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetExternalClusterId returns the ExternalClusterId field value if set, zero value otherwise.
func (o *SubscriptionPatchRequest) GetExternalClusterId() string {
	if o == nil || IsNil(o.ExternalClusterId) {
		var ret string
		return ret
	}
	return *o.ExternalClusterId
}

// GetExternalClusterIdOk returns a tuple with the ExternalClusterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPatchRequest) GetExternalClusterIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalClusterId) {
		return nil, false
	}
	return o.ExternalClusterId, true
}

// HasExternalClusterId returns a boolean if a field has been set.
func (o *SubscriptionPatchRequest) HasExternalClusterId() bool {
	if o != nil && !IsNil(o.ExternalClusterId) {
		return true
	}

	return false
}

// SetExternalClusterId gets a reference to the given string and assigns it to the ExternalClusterId field.
func (o *SubscriptionPatchRequest) SetExternalClusterId(v string) {
	o.ExternalClusterId = &v
}

// GetManaged returns the Managed field value if set, zero value otherwise.
func (o *SubscriptionPatchRequest) GetManaged() bool {
	if o == nil || IsNil(o.Managed) {
		var ret bool
		return ret
	}
	return *o.Managed
}

// GetManagedOk returns a tuple with the Managed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPatchRequest) GetManagedOk() (*bool, bool) {
	if o == nil || IsNil(o.Managed) {
		return nil, false
	}
	return o.Managed, true
}

// HasManaged returns a boolean if a field has been set.
func (o *SubscriptionPatchRequest) HasManaged() bool {
	if o != nil && !IsNil(o.Managed) {
		return true
	}

	return false
}

// SetManaged gets a reference to the given bool and assigns it to the Managed field.
func (o *SubscriptionPatchRequest) SetManaged(v bool) {
	o.Managed = &v
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise.
func (o *SubscriptionPatchRequest) GetOrganizationId() string {
	if o == nil || IsNil(o.OrganizationId) {
		var ret string
		return ret
	}
	return *o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPatchRequest) GetOrganizationIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrganizationId) {
		return nil, false
	}
	return o.OrganizationId, true
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *SubscriptionPatchRequest) HasOrganizationId() bool {
	if o != nil && !IsNil(o.OrganizationId) {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given string and assigns it to the OrganizationId field.
func (o *SubscriptionPatchRequest) SetOrganizationId(v string) {
	o.OrganizationId = &v
}

// GetPlanId returns the PlanId field value if set, zero value otherwise.
func (o *SubscriptionPatchRequest) GetPlanId() string {
	if o == nil || IsNil(o.PlanId) {
		var ret string
		return ret
	}
	return *o.PlanId
}

// GetPlanIdOk returns a tuple with the PlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPatchRequest) GetPlanIdOk() (*string, bool) {
	if o == nil || IsNil(o.PlanId) {
		return nil, false
	}
	return o.PlanId, true
}

// HasPlanId returns a boolean if a field has been set.
func (o *SubscriptionPatchRequest) HasPlanId() bool {
	if o != nil && !IsNil(o.PlanId) {
		return true
	}

	return false
}

// SetPlanId gets a reference to the given string and assigns it to the PlanId field.
func (o *SubscriptionPatchRequest) SetPlanId(v string) {
	o.PlanId = &v
}

// GetProductBundle returns the ProductBundle field value if set, zero value otherwise.
func (o *SubscriptionPatchRequest) GetProductBundle() string {
	if o == nil || IsNil(o.ProductBundle) {
		var ret string
		return ret
	}
	return *o.ProductBundle
}

// GetProductBundleOk returns a tuple with the ProductBundle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPatchRequest) GetProductBundleOk() (*string, bool) {
	if o == nil || IsNil(o.ProductBundle) {
		return nil, false
	}
	return o.ProductBundle, true
}

// HasProductBundle returns a boolean if a field has been set.
func (o *SubscriptionPatchRequest) HasProductBundle() bool {
	if o != nil && !IsNil(o.ProductBundle) {
		return true
	}

	return false
}

// SetProductBundle gets a reference to the given string and assigns it to the ProductBundle field.
func (o *SubscriptionPatchRequest) SetProductBundle(v string) {
	o.ProductBundle = &v
}

// GetProvenance returns the Provenance field value if set, zero value otherwise.
func (o *SubscriptionPatchRequest) GetProvenance() string {
	if o == nil || IsNil(o.Provenance) {
		var ret string
		return ret
	}
	return *o.Provenance
}

// GetProvenanceOk returns a tuple with the Provenance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPatchRequest) GetProvenanceOk() (*string, bool) {
	if o == nil || IsNil(o.Provenance) {
		return nil, false
	}
	return o.Provenance, true
}

// HasProvenance returns a boolean if a field has been set.
func (o *SubscriptionPatchRequest) HasProvenance() bool {
	if o != nil && !IsNil(o.Provenance) {
		return true
	}

	return false
}

// SetProvenance gets a reference to the given string and assigns it to the Provenance field.
func (o *SubscriptionPatchRequest) SetProvenance(v string) {
	o.Provenance = &v
}

// GetRegionId returns the RegionId field value if set, zero value otherwise.
func (o *SubscriptionPatchRequest) GetRegionId() string {
	if o == nil || IsNil(o.RegionId) {
		var ret string
		return ret
	}
	return *o.RegionId
}

// GetRegionIdOk returns a tuple with the RegionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPatchRequest) GetRegionIdOk() (*string, bool) {
	if o == nil || IsNil(o.RegionId) {
		return nil, false
	}
	return o.RegionId, true
}

// HasRegionId returns a boolean if a field has been set.
func (o *SubscriptionPatchRequest) HasRegionId() bool {
	if o != nil && !IsNil(o.RegionId) {
		return true
	}

	return false
}

// SetRegionId gets a reference to the given string and assigns it to the RegionId field.
func (o *SubscriptionPatchRequest) SetRegionId(v string) {
	o.RegionId = &v
}

// GetReleased returns the Released field value if set, zero value otherwise.
func (o *SubscriptionPatchRequest) GetReleased() bool {
	if o == nil || IsNil(o.Released) {
		var ret bool
		return ret
	}
	return *o.Released
}

// GetReleasedOk returns a tuple with the Released field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPatchRequest) GetReleasedOk() (*bool, bool) {
	if o == nil || IsNil(o.Released) {
		return nil, false
	}
	return o.Released, true
}

// HasReleased returns a boolean if a field has been set.
func (o *SubscriptionPatchRequest) HasReleased() bool {
	if o != nil && !IsNil(o.Released) {
		return true
	}

	return false
}

// SetReleased gets a reference to the given bool and assigns it to the Released field.
func (o *SubscriptionPatchRequest) SetReleased(v bool) {
	o.Released = &v
}

// GetServiceLevel returns the ServiceLevel field value if set, zero value otherwise.
func (o *SubscriptionPatchRequest) GetServiceLevel() string {
	if o == nil || IsNil(o.ServiceLevel) {
		var ret string
		return ret
	}
	return *o.ServiceLevel
}

// GetServiceLevelOk returns a tuple with the ServiceLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPatchRequest) GetServiceLevelOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceLevel) {
		return nil, false
	}
	return o.ServiceLevel, true
}

// HasServiceLevel returns a boolean if a field has been set.
func (o *SubscriptionPatchRequest) HasServiceLevel() bool {
	if o != nil && !IsNil(o.ServiceLevel) {
		return true
	}

	return false
}

// SetServiceLevel gets a reference to the given string and assigns it to the ServiceLevel field.
func (o *SubscriptionPatchRequest) SetServiceLevel(v string) {
	o.ServiceLevel = &v
}

// GetSocketTotal returns the SocketTotal field value if set, zero value otherwise.
func (o *SubscriptionPatchRequest) GetSocketTotal() int32 {
	if o == nil || IsNil(o.SocketTotal) {
		var ret int32
		return ret
	}
	return *o.SocketTotal
}

// GetSocketTotalOk returns a tuple with the SocketTotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPatchRequest) GetSocketTotalOk() (*int32, bool) {
	if o == nil || IsNil(o.SocketTotal) {
		return nil, false
	}
	return o.SocketTotal, true
}

// HasSocketTotal returns a boolean if a field has been set.
func (o *SubscriptionPatchRequest) HasSocketTotal() bool {
	if o != nil && !IsNil(o.SocketTotal) {
		return true
	}

	return false
}

// SetSocketTotal gets a reference to the given int32 and assigns it to the SocketTotal field.
func (o *SubscriptionPatchRequest) SetSocketTotal(v int32) {
	o.SocketTotal = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SubscriptionPatchRequest) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPatchRequest) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SubscriptionPatchRequest) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *SubscriptionPatchRequest) SetStatus(v string) {
	o.Status = &v
}

// GetSupportLevel returns the SupportLevel field value if set, zero value otherwise.
func (o *SubscriptionPatchRequest) GetSupportLevel() string {
	if o == nil || IsNil(o.SupportLevel) {
		var ret string
		return ret
	}
	return *o.SupportLevel
}

// GetSupportLevelOk returns a tuple with the SupportLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPatchRequest) GetSupportLevelOk() (*string, bool) {
	if o == nil || IsNil(o.SupportLevel) {
		return nil, false
	}
	return o.SupportLevel, true
}

// HasSupportLevel returns a boolean if a field has been set.
func (o *SubscriptionPatchRequest) HasSupportLevel() bool {
	if o != nil && !IsNil(o.SupportLevel) {
		return true
	}

	return false
}

// SetSupportLevel gets a reference to the given string and assigns it to the SupportLevel field.
func (o *SubscriptionPatchRequest) SetSupportLevel(v string) {
	o.SupportLevel = &v
}

// GetSystemUnits returns the SystemUnits field value if set, zero value otherwise.
func (o *SubscriptionPatchRequest) GetSystemUnits() string {
	if o == nil || IsNil(o.SystemUnits) {
		var ret string
		return ret
	}
	return *o.SystemUnits
}

// GetSystemUnitsOk returns a tuple with the SystemUnits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPatchRequest) GetSystemUnitsOk() (*string, bool) {
	if o == nil || IsNil(o.SystemUnits) {
		return nil, false
	}
	return o.SystemUnits, true
}

// HasSystemUnits returns a boolean if a field has been set.
func (o *SubscriptionPatchRequest) HasSystemUnits() bool {
	if o != nil && !IsNil(o.SystemUnits) {
		return true
	}

	return false
}

// SetSystemUnits gets a reference to the given string and assigns it to the SystemUnits field.
func (o *SubscriptionPatchRequest) SetSystemUnits(v string) {
	o.SystemUnits = &v
}

// GetTrialEndDate returns the TrialEndDate field value if set, zero value otherwise.
func (o *SubscriptionPatchRequest) GetTrialEndDate() time.Time {
	if o == nil || IsNil(o.TrialEndDate) {
		var ret time.Time
		return ret
	}
	return *o.TrialEndDate
}

// GetTrialEndDateOk returns a tuple with the TrialEndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPatchRequest) GetTrialEndDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.TrialEndDate) {
		return nil, false
	}
	return o.TrialEndDate, true
}

// HasTrialEndDate returns a boolean if a field has been set.
func (o *SubscriptionPatchRequest) HasTrialEndDate() bool {
	if o != nil && !IsNil(o.TrialEndDate) {
		return true
	}

	return false
}

// SetTrialEndDate gets a reference to the given time.Time and assigns it to the TrialEndDate field.
func (o *SubscriptionPatchRequest) SetTrialEndDate(v time.Time) {
	o.TrialEndDate = &v
}

// GetUsage returns the Usage field value if set, zero value otherwise.
func (o *SubscriptionPatchRequest) GetUsage() string {
	if o == nil || IsNil(o.Usage) {
		var ret string
		return ret
	}
	return *o.Usage
}

// GetUsageOk returns a tuple with the Usage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPatchRequest) GetUsageOk() (*string, bool) {
	if o == nil || IsNil(o.Usage) {
		return nil, false
	}
	return o.Usage, true
}

// HasUsage returns a boolean if a field has been set.
func (o *SubscriptionPatchRequest) HasUsage() bool {
	if o != nil && !IsNil(o.Usage) {
		return true
	}

	return false
}

// SetUsage gets a reference to the given string and assigns it to the Usage field.
func (o *SubscriptionPatchRequest) SetUsage(v string) {
	o.Usage = &v
}

func (o SubscriptionPatchRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionPatchRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BillingExpirationDate) {
		toSerialize["billing_expiration_date"] = o.BillingExpirationDate
	}
	if !IsNil(o.CloudAccountId) {
		toSerialize["cloud_account_id"] = o.CloudAccountId
	}
	if !IsNil(o.CloudProviderId) {
		toSerialize["cloud_provider_id"] = o.CloudProviderId
	}
	if !IsNil(o.ClusterBillingModel) {
		toSerialize["cluster_billing_model"] = o.ClusterBillingModel
	}
	if !IsNil(o.ClusterId) {
		toSerialize["cluster_id"] = o.ClusterId
	}
	if !IsNil(o.ConsoleUrl) {
		toSerialize["console_url"] = o.ConsoleUrl
	}
	if !IsNil(o.ConsumerUuid) {
		toSerialize["consumer_uuid"] = o.ConsumerUuid
	}
	if !IsNil(o.CpuTotal) {
		toSerialize["cpu_total"] = o.CpuTotal
	}
	if !IsNil(o.CreatorId) {
		toSerialize["creator_id"] = o.CreatorId
	}
	if !IsNil(o.DisplayName) {
		toSerialize["display_name"] = o.DisplayName
	}
	if !IsNil(o.ExternalClusterId) {
		toSerialize["external_cluster_id"] = o.ExternalClusterId
	}
	if !IsNil(o.Managed) {
		toSerialize["managed"] = o.Managed
	}
	if !IsNil(o.OrganizationId) {
		toSerialize["organization_id"] = o.OrganizationId
	}
	if !IsNil(o.PlanId) {
		toSerialize["plan_id"] = o.PlanId
	}
	if !IsNil(o.ProductBundle) {
		toSerialize["product_bundle"] = o.ProductBundle
	}
	if !IsNil(o.Provenance) {
		toSerialize["provenance"] = o.Provenance
	}
	if !IsNil(o.RegionId) {
		toSerialize["region_id"] = o.RegionId
	}
	if !IsNil(o.Released) {
		toSerialize["released"] = o.Released
	}
	if !IsNil(o.ServiceLevel) {
		toSerialize["service_level"] = o.ServiceLevel
	}
	if !IsNil(o.SocketTotal) {
		toSerialize["socket_total"] = o.SocketTotal
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.SupportLevel) {
		toSerialize["support_level"] = o.SupportLevel
	}
	if !IsNil(o.SystemUnits) {
		toSerialize["system_units"] = o.SystemUnits
	}
	if !IsNil(o.TrialEndDate) {
		toSerialize["trial_end_date"] = o.TrialEndDate
	}
	if !IsNil(o.Usage) {
		toSerialize["usage"] = o.Usage
	}
	return toSerialize, nil
}

type NullableSubscriptionPatchRequest struct {
	value *SubscriptionPatchRequest
	isSet bool
}

func (v NullableSubscriptionPatchRequest) Get() *SubscriptionPatchRequest {
	return v.value
}

func (v *NullableSubscriptionPatchRequest) Set(val *SubscriptionPatchRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionPatchRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionPatchRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionPatchRequest(val *SubscriptionPatchRequest) *NullableSubscriptionPatchRequest {
	return &NullableSubscriptionPatchRequest{value: val, isSet: true}
}

func (v NullableSubscriptionPatchRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionPatchRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


