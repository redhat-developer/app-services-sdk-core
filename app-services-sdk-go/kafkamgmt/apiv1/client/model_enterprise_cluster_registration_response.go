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
)

// checks if the EnterpriseClusterRegistrationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EnterpriseClusterRegistrationResponse{}

// EnterpriseClusterRegistrationResponse Enterprise Cluster registration response. It returns additional privileged information compared to The information returned by EnterpriseCluster
type EnterpriseClusterRegistrationResponse struct {
	Id string `json:"id"`
	Kind string `json:"kind"`
	Href string `json:"href"`
	// Indicates whether Kafkas created on this data plane cluster have to be accessed via private network
	AccessKafkasViaPrivateNetwork bool `json:"access_kafkas_via_private_network"`
	// The OCM's cluster id of the registered Enterprise cluster.
	ClusterId *string `json:"cluster_id,omitempty"`
	// The status of Enterprise cluster registration
	Status *string `json:"status,omitempty"`
	// The cloud provider for this cluster. This valus will be used as the Kafka's cloud provider value when a Kafka is created on this cluster
	CloudProvider *string `json:"cloud_provider,omitempty"`
	// The region of this cluster. This valus will be used as the Kafka's region value when a Kafka is created on this cluster
	Region *string `json:"region,omitempty"`
	// A flag indicating whether this cluster is available on multiple availability zones or not
	MultiAz bool `json:"multi_az"`
	SupportedInstanceTypes SupportedKafkaInstanceTypesList `json:"supported_instance_types,omitempty"`
	CapacityInformation *EnterpriseClusterAllOfCapacityInformation `json:"capacity_information,omitempty"`
	// Enterprise Cluster fleetshard parameters array
	FleetshardParameters []FleetshardParameter `json:"fleetshard_parameters,omitempty"`
}

// NewEnterpriseClusterRegistrationResponse instantiates a new EnterpriseClusterRegistrationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnterpriseClusterRegistrationResponse(id string, kind string, href string, accessKafkasViaPrivateNetwork bool, multiAz bool) *EnterpriseClusterRegistrationResponse {
	this := EnterpriseClusterRegistrationResponse{}
	this.Id = id
	this.Kind = kind
	this.Href = href
	this.AccessKafkasViaPrivateNetwork = accessKafkasViaPrivateNetwork
	this.MultiAz = multiAz
	return &this
}

// NewEnterpriseClusterRegistrationResponseWithDefaults instantiates a new EnterpriseClusterRegistrationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnterpriseClusterRegistrationResponseWithDefaults() *EnterpriseClusterRegistrationResponse {
	this := EnterpriseClusterRegistrationResponse{}
	return &this
}

// GetId returns the Id field value
func (o *EnterpriseClusterRegistrationResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *EnterpriseClusterRegistrationResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *EnterpriseClusterRegistrationResponse) SetId(v string) {
	o.Id = v
}

// GetKind returns the Kind field value
func (o *EnterpriseClusterRegistrationResponse) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *EnterpriseClusterRegistrationResponse) GetKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *EnterpriseClusterRegistrationResponse) SetKind(v string) {
	o.Kind = v
}

// GetHref returns the Href field value
func (o *EnterpriseClusterRegistrationResponse) GetHref() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Href
}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
func (o *EnterpriseClusterRegistrationResponse) GetHrefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Href, true
}

// SetHref sets field value
func (o *EnterpriseClusterRegistrationResponse) SetHref(v string) {
	o.Href = v
}

// GetAccessKafkasViaPrivateNetwork returns the AccessKafkasViaPrivateNetwork field value
func (o *EnterpriseClusterRegistrationResponse) GetAccessKafkasViaPrivateNetwork() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AccessKafkasViaPrivateNetwork
}

// GetAccessKafkasViaPrivateNetworkOk returns a tuple with the AccessKafkasViaPrivateNetwork field value
// and a boolean to check if the value has been set.
func (o *EnterpriseClusterRegistrationResponse) GetAccessKafkasViaPrivateNetworkOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccessKafkasViaPrivateNetwork, true
}

// SetAccessKafkasViaPrivateNetwork sets field value
func (o *EnterpriseClusterRegistrationResponse) SetAccessKafkasViaPrivateNetwork(v bool) {
	o.AccessKafkasViaPrivateNetwork = v
}

// GetClusterId returns the ClusterId field value if set, zero value otherwise.
func (o *EnterpriseClusterRegistrationResponse) GetClusterId() string {
	if o == nil || IsNil(o.ClusterId) {
		var ret string
		return ret
	}
	return *o.ClusterId
}

// GetClusterIdOk returns a tuple with the ClusterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnterpriseClusterRegistrationResponse) GetClusterIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClusterId) {
		return nil, false
	}
	return o.ClusterId, true
}

// HasClusterId returns a boolean if a field has been set.
func (o *EnterpriseClusterRegistrationResponse) HasClusterId() bool {
	if o != nil && !IsNil(o.ClusterId) {
		return true
	}

	return false
}

// SetClusterId gets a reference to the given string and assigns it to the ClusterId field.
func (o *EnterpriseClusterRegistrationResponse) SetClusterId(v string) {
	o.ClusterId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *EnterpriseClusterRegistrationResponse) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnterpriseClusterRegistrationResponse) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *EnterpriseClusterRegistrationResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *EnterpriseClusterRegistrationResponse) SetStatus(v string) {
	o.Status = &v
}

// GetCloudProvider returns the CloudProvider field value if set, zero value otherwise.
func (o *EnterpriseClusterRegistrationResponse) GetCloudProvider() string {
	if o == nil || IsNil(o.CloudProvider) {
		var ret string
		return ret
	}
	return *o.CloudProvider
}

// GetCloudProviderOk returns a tuple with the CloudProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnterpriseClusterRegistrationResponse) GetCloudProviderOk() (*string, bool) {
	if o == nil || IsNil(o.CloudProvider) {
		return nil, false
	}
	return o.CloudProvider, true
}

// HasCloudProvider returns a boolean if a field has been set.
func (o *EnterpriseClusterRegistrationResponse) HasCloudProvider() bool {
	if o != nil && !IsNil(o.CloudProvider) {
		return true
	}

	return false
}

// SetCloudProvider gets a reference to the given string and assigns it to the CloudProvider field.
func (o *EnterpriseClusterRegistrationResponse) SetCloudProvider(v string) {
	o.CloudProvider = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *EnterpriseClusterRegistrationResponse) GetRegion() string {
	if o == nil || IsNil(o.Region) {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnterpriseClusterRegistrationResponse) GetRegionOk() (*string, bool) {
	if o == nil || IsNil(o.Region) {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *EnterpriseClusterRegistrationResponse) HasRegion() bool {
	if o != nil && !IsNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *EnterpriseClusterRegistrationResponse) SetRegion(v string) {
	o.Region = &v
}

// GetMultiAz returns the MultiAz field value
func (o *EnterpriseClusterRegistrationResponse) GetMultiAz() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.MultiAz
}

// GetMultiAzOk returns a tuple with the MultiAz field value
// and a boolean to check if the value has been set.
func (o *EnterpriseClusterRegistrationResponse) GetMultiAzOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MultiAz, true
}

// SetMultiAz sets field value
func (o *EnterpriseClusterRegistrationResponse) SetMultiAz(v bool) {
	o.MultiAz = v
}

// GetSupportedInstanceTypes returns the SupportedInstanceTypes field value if set, zero value otherwise.
func (o *EnterpriseClusterRegistrationResponse) GetSupportedInstanceTypes() SupportedKafkaInstanceTypesList {
	if o == nil || IsNil(o.SupportedInstanceTypes) {
		var ret SupportedKafkaInstanceTypesList
		return ret
	}
	return o.SupportedInstanceTypes
}

// GetSupportedInstanceTypesOk returns a tuple with the SupportedInstanceTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnterpriseClusterRegistrationResponse) GetSupportedInstanceTypesOk() (SupportedKafkaInstanceTypesList, bool) {
	if o == nil || IsNil(o.SupportedInstanceTypes) {
		return SupportedKafkaInstanceTypesList{}, false
	}
	return o.SupportedInstanceTypes, true
}

// HasSupportedInstanceTypes returns a boolean if a field has been set.
func (o *EnterpriseClusterRegistrationResponse) HasSupportedInstanceTypes() bool {
	if o != nil && !IsNil(o.SupportedInstanceTypes) {
		return true
	}

	return false
}

// SetSupportedInstanceTypes gets a reference to the given SupportedKafkaInstanceTypesList and assigns it to the SupportedInstanceTypes field.
func (o *EnterpriseClusterRegistrationResponse) SetSupportedInstanceTypes(v SupportedKafkaInstanceTypesList) {
	o.SupportedInstanceTypes = v
}

// GetCapacityInformation returns the CapacityInformation field value if set, zero value otherwise.
func (o *EnterpriseClusterRegistrationResponse) GetCapacityInformation() EnterpriseClusterAllOfCapacityInformation {
	if o == nil || IsNil(o.CapacityInformation) {
		var ret EnterpriseClusterAllOfCapacityInformation
		return ret
	}
	return *o.CapacityInformation
}

// GetCapacityInformationOk returns a tuple with the CapacityInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnterpriseClusterRegistrationResponse) GetCapacityInformationOk() (*EnterpriseClusterAllOfCapacityInformation, bool) {
	if o == nil || IsNil(o.CapacityInformation) {
		return nil, false
	}
	return o.CapacityInformation, true
}

// HasCapacityInformation returns a boolean if a field has been set.
func (o *EnterpriseClusterRegistrationResponse) HasCapacityInformation() bool {
	if o != nil && !IsNil(o.CapacityInformation) {
		return true
	}

	return false
}

// SetCapacityInformation gets a reference to the given EnterpriseClusterAllOfCapacityInformation and assigns it to the CapacityInformation field.
func (o *EnterpriseClusterRegistrationResponse) SetCapacityInformation(v EnterpriseClusterAllOfCapacityInformation) {
	o.CapacityInformation = &v
}

// GetFleetshardParameters returns the FleetshardParameters field value if set, zero value otherwise.
func (o *EnterpriseClusterRegistrationResponse) GetFleetshardParameters() []FleetshardParameter {
	if o == nil || IsNil(o.FleetshardParameters) {
		var ret []FleetshardParameter
		return ret
	}
	return o.FleetshardParameters
}

// GetFleetshardParametersOk returns a tuple with the FleetshardParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnterpriseClusterRegistrationResponse) GetFleetshardParametersOk() ([]FleetshardParameter, bool) {
	if o == nil || IsNil(o.FleetshardParameters) {
		return nil, false
	}
	return o.FleetshardParameters, true
}

// HasFleetshardParameters returns a boolean if a field has been set.
func (o *EnterpriseClusterRegistrationResponse) HasFleetshardParameters() bool {
	if o != nil && !IsNil(o.FleetshardParameters) {
		return true
	}

	return false
}

// SetFleetshardParameters gets a reference to the given []FleetshardParameter and assigns it to the FleetshardParameters field.
func (o *EnterpriseClusterRegistrationResponse) SetFleetshardParameters(v []FleetshardParameter) {
	o.FleetshardParameters = v
}

func (o EnterpriseClusterRegistrationResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EnterpriseClusterRegistrationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["kind"] = o.Kind
	toSerialize["href"] = o.Href
	toSerialize["access_kafkas_via_private_network"] = o.AccessKafkasViaPrivateNetwork
	if !IsNil(o.ClusterId) {
		toSerialize["cluster_id"] = o.ClusterId
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.CloudProvider) {
		toSerialize["cloud_provider"] = o.CloudProvider
	}
	if !IsNil(o.Region) {
		toSerialize["region"] = o.Region
	}
	toSerialize["multi_az"] = o.MultiAz
	if !IsNil(o.SupportedInstanceTypes) {
		toSerialize["supported_instance_types"] = o.SupportedInstanceTypes
	}
	if !IsNil(o.CapacityInformation) {
		toSerialize["capacity_information"] = o.CapacityInformation
	}
	if !IsNil(o.FleetshardParameters) {
		toSerialize["fleetshard_parameters"] = o.FleetshardParameters
	}
	return toSerialize, nil
}

type NullableEnterpriseClusterRegistrationResponse struct {
	value *EnterpriseClusterRegistrationResponse
	isSet bool
}

func (v NullableEnterpriseClusterRegistrationResponse) Get() *EnterpriseClusterRegistrationResponse {
	return v.value
}

func (v *NullableEnterpriseClusterRegistrationResponse) Set(val *EnterpriseClusterRegistrationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEnterpriseClusterRegistrationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEnterpriseClusterRegistrationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnterpriseClusterRegistrationResponse(val *EnterpriseClusterRegistrationResponse) *NullableEnterpriseClusterRegistrationResponse {
	return &NullableEnterpriseClusterRegistrationResponse{value: val, isSet: true}
}

func (v NullableEnterpriseClusterRegistrationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnterpriseClusterRegistrationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


