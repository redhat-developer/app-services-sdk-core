# KafkaRequestListAllOfItemsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Kind** | **string** |  | 
**Href** | **string** |  | 
**Status** | Pointer to **string** | Values: [accepted, preparing, provisioning, ready, failed, deprovision, deleting, suspending, suspended, resuming]  | [optional] 
**CloudProvider** | Pointer to **string** | Name of Cloud used to deploy. For example AWS | [optional] 
**MultiAz** | **bool** |  | 
**Region** | Pointer to **string** | Values will be regions of specific cloud provider. For example: us-east-1 for AWS | [optional] 
**Owner** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**BootstrapServerHost** | Pointer to **string** |  | [optional] 
**AdminApiServerUrl** | Pointer to **string** | The kafka admin server url to perform kafka admin operations e.g acl management etc. The value will be available when the Kafka has been fully provisioned i.e it reaches a &#39;ready&#39; state | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**ExpiresAt** | Pointer to **NullableTime** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**FailedReason** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**InstanceType** | Pointer to **string** |  | [optional] 
**InstanceTypeName** | Pointer to **string** | This field is now deprecated, please use the /api/kafkas_mgmt/v1/instance_types/{cloud_provider}/{cloud_region} endpoint to retrieve the field instead. | [optional] 
**ReauthenticationEnabled** | **bool** |  | 
**MaxDataRetentionSize** | Pointer to [**SupportedKafkaSizeBytesValueItem**](SupportedKafkaSizeBytesValueItem.md) |  | [optional] 
**BrowserUrl** | Pointer to **string** |  | [optional] 
**SizeId** | Pointer to **string** |  | [optional] 
**IngressThroughputPerSec** | Pointer to **string** | This field is now deprecated, please use the /api/kafkas_mgmt/v1/instance_types/{cloud_provider}/{cloud_region} endpoint to retrieve the field instead. | [optional] 
**EgressThroughputPerSec** | Pointer to **string** | This field is now deprecated, please use the /api/kafkas_mgmt/v1/instance_types/{cloud_provider}/{cloud_region} endpoint to retrieve the field instead. | [optional] 
**TotalMaxConnections** | Pointer to **int32** | This field is now deprecated, please use the /api/kafkas_mgmt/v1/instance_types/{cloud_provider}/{cloud_region} endpoint to retrieve the field instead. | [optional] 
**MaxPartitions** | Pointer to **int32** | This field is now deprecated, please use the /api/kafkas_mgmt/v1/instance_types/{cloud_provider}/{cloud_region} endpoint to retrieve the field instead. | [optional] 
**MaxDataRetentionPeriod** | Pointer to **string** | This field is now deprecated, please use the /api/kafkas_mgmt/v1/instance_types/{cloud_provider}/{cloud_region} endpoint to retrieve the field instead. | [optional] 
**MaxConnectionAttemptsPerSec** | Pointer to **int32** | This field is now deprecated, please use the /api/kafkas_mgmt/v1/instance_types/{cloud_provider}/{cloud_region} endpoint to retrieve the field instead. | [optional] 
**BillingCloudAccountId** | Pointer to **string** |  | [optional] 
**Marketplace** | Pointer to **string** |  | [optional] 
**BillingModel** | Pointer to **string** |  | [optional] 
**PromotionStatus** | Pointer to **string** | Status of the Kafka request promotion. Possible values: [&#39;promoting&#39;, &#39;failed&#39;]. If unset it means no promotion is in progress. | [optional] 
**ClusterId** | Pointer to **NullableString** | The ID of the data plane where Kafka is deployed on. This information is only returned for kafka whose billing model is enterprise | [optional] 
**PromotionDetails** | Pointer to **string** | Details of the Kafka request promotion. It can be set when a Kafka request promotion is in progress or has failed | [optional] 

## Methods

### NewKafkaRequestListAllOfItemsInner

`func NewKafkaRequestListAllOfItemsInner(id string, kind string, href string, multiAz bool, reauthenticationEnabled bool, ) *KafkaRequestListAllOfItemsInner`

NewKafkaRequestListAllOfItemsInner instantiates a new KafkaRequestListAllOfItemsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKafkaRequestListAllOfItemsInnerWithDefaults

`func NewKafkaRequestListAllOfItemsInnerWithDefaults() *KafkaRequestListAllOfItemsInner`

NewKafkaRequestListAllOfItemsInnerWithDefaults instantiates a new KafkaRequestListAllOfItemsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KafkaRequestListAllOfItemsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KafkaRequestListAllOfItemsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KafkaRequestListAllOfItemsInner) SetId(v string)`

SetId sets Id field to given value.


### GetKind

`func (o *KafkaRequestListAllOfItemsInner) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *KafkaRequestListAllOfItemsInner) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *KafkaRequestListAllOfItemsInner) SetKind(v string)`

SetKind sets Kind field to given value.


### GetHref

`func (o *KafkaRequestListAllOfItemsInner) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *KafkaRequestListAllOfItemsInner) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *KafkaRequestListAllOfItemsInner) SetHref(v string)`

SetHref sets Href field to given value.


### GetStatus

`func (o *KafkaRequestListAllOfItemsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *KafkaRequestListAllOfItemsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *KafkaRequestListAllOfItemsInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *KafkaRequestListAllOfItemsInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCloudProvider

`func (o *KafkaRequestListAllOfItemsInner) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *KafkaRequestListAllOfItemsInner) GetCloudProviderOk() (*string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *KafkaRequestListAllOfItemsInner) SetCloudProvider(v string)`

SetCloudProvider sets CloudProvider field to given value.

### HasCloudProvider

`func (o *KafkaRequestListAllOfItemsInner) HasCloudProvider() bool`

HasCloudProvider returns a boolean if a field has been set.

### GetMultiAz

`func (o *KafkaRequestListAllOfItemsInner) GetMultiAz() bool`

GetMultiAz returns the MultiAz field if non-nil, zero value otherwise.

### GetMultiAzOk

`func (o *KafkaRequestListAllOfItemsInner) GetMultiAzOk() (*bool, bool)`

GetMultiAzOk returns a tuple with the MultiAz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiAz

`func (o *KafkaRequestListAllOfItemsInner) SetMultiAz(v bool)`

SetMultiAz sets MultiAz field to given value.


### GetRegion

`func (o *KafkaRequestListAllOfItemsInner) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *KafkaRequestListAllOfItemsInner) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *KafkaRequestListAllOfItemsInner) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *KafkaRequestListAllOfItemsInner) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetOwner

`func (o *KafkaRequestListAllOfItemsInner) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *KafkaRequestListAllOfItemsInner) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *KafkaRequestListAllOfItemsInner) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *KafkaRequestListAllOfItemsInner) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetName

`func (o *KafkaRequestListAllOfItemsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KafkaRequestListAllOfItemsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KafkaRequestListAllOfItemsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KafkaRequestListAllOfItemsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetBootstrapServerHost

`func (o *KafkaRequestListAllOfItemsInner) GetBootstrapServerHost() string`

GetBootstrapServerHost returns the BootstrapServerHost field if non-nil, zero value otherwise.

### GetBootstrapServerHostOk

`func (o *KafkaRequestListAllOfItemsInner) GetBootstrapServerHostOk() (*string, bool)`

GetBootstrapServerHostOk returns a tuple with the BootstrapServerHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootstrapServerHost

`func (o *KafkaRequestListAllOfItemsInner) SetBootstrapServerHost(v string)`

SetBootstrapServerHost sets BootstrapServerHost field to given value.

### HasBootstrapServerHost

`func (o *KafkaRequestListAllOfItemsInner) HasBootstrapServerHost() bool`

HasBootstrapServerHost returns a boolean if a field has been set.

### GetAdminApiServerUrl

`func (o *KafkaRequestListAllOfItemsInner) GetAdminApiServerUrl() string`

GetAdminApiServerUrl returns the AdminApiServerUrl field if non-nil, zero value otherwise.

### GetAdminApiServerUrlOk

`func (o *KafkaRequestListAllOfItemsInner) GetAdminApiServerUrlOk() (*string, bool)`

GetAdminApiServerUrlOk returns a tuple with the AdminApiServerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminApiServerUrl

`func (o *KafkaRequestListAllOfItemsInner) SetAdminApiServerUrl(v string)`

SetAdminApiServerUrl sets AdminApiServerUrl field to given value.

### HasAdminApiServerUrl

`func (o *KafkaRequestListAllOfItemsInner) HasAdminApiServerUrl() bool`

HasAdminApiServerUrl returns a boolean if a field has been set.

### GetCreatedAt

`func (o *KafkaRequestListAllOfItemsInner) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *KafkaRequestListAllOfItemsInner) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *KafkaRequestListAllOfItemsInner) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *KafkaRequestListAllOfItemsInner) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetExpiresAt

`func (o *KafkaRequestListAllOfItemsInner) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *KafkaRequestListAllOfItemsInner) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *KafkaRequestListAllOfItemsInner) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *KafkaRequestListAllOfItemsInner) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### SetExpiresAtNil

`func (o *KafkaRequestListAllOfItemsInner) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *KafkaRequestListAllOfItemsInner) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil
### GetUpdatedAt

`func (o *KafkaRequestListAllOfItemsInner) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *KafkaRequestListAllOfItemsInner) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *KafkaRequestListAllOfItemsInner) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *KafkaRequestListAllOfItemsInner) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetFailedReason

`func (o *KafkaRequestListAllOfItemsInner) GetFailedReason() string`

GetFailedReason returns the FailedReason field if non-nil, zero value otherwise.

### GetFailedReasonOk

`func (o *KafkaRequestListAllOfItemsInner) GetFailedReasonOk() (*string, bool)`

GetFailedReasonOk returns a tuple with the FailedReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedReason

`func (o *KafkaRequestListAllOfItemsInner) SetFailedReason(v string)`

SetFailedReason sets FailedReason field to given value.

### HasFailedReason

`func (o *KafkaRequestListAllOfItemsInner) HasFailedReason() bool`

HasFailedReason returns a boolean if a field has been set.

### GetVersion

`func (o *KafkaRequestListAllOfItemsInner) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *KafkaRequestListAllOfItemsInner) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *KafkaRequestListAllOfItemsInner) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *KafkaRequestListAllOfItemsInner) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetInstanceType

`func (o *KafkaRequestListAllOfItemsInner) GetInstanceType() string`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *KafkaRequestListAllOfItemsInner) GetInstanceTypeOk() (*string, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *KafkaRequestListAllOfItemsInner) SetInstanceType(v string)`

SetInstanceType sets InstanceType field to given value.

### HasInstanceType

`func (o *KafkaRequestListAllOfItemsInner) HasInstanceType() bool`

HasInstanceType returns a boolean if a field has been set.

### GetInstanceTypeName

`func (o *KafkaRequestListAllOfItemsInner) GetInstanceTypeName() string`

GetInstanceTypeName returns the InstanceTypeName field if non-nil, zero value otherwise.

### GetInstanceTypeNameOk

`func (o *KafkaRequestListAllOfItemsInner) GetInstanceTypeNameOk() (*string, bool)`

GetInstanceTypeNameOk returns a tuple with the InstanceTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceTypeName

`func (o *KafkaRequestListAllOfItemsInner) SetInstanceTypeName(v string)`

SetInstanceTypeName sets InstanceTypeName field to given value.

### HasInstanceTypeName

`func (o *KafkaRequestListAllOfItemsInner) HasInstanceTypeName() bool`

HasInstanceTypeName returns a boolean if a field has been set.

### GetReauthenticationEnabled

`func (o *KafkaRequestListAllOfItemsInner) GetReauthenticationEnabled() bool`

GetReauthenticationEnabled returns the ReauthenticationEnabled field if non-nil, zero value otherwise.

### GetReauthenticationEnabledOk

`func (o *KafkaRequestListAllOfItemsInner) GetReauthenticationEnabledOk() (*bool, bool)`

GetReauthenticationEnabledOk returns a tuple with the ReauthenticationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReauthenticationEnabled

`func (o *KafkaRequestListAllOfItemsInner) SetReauthenticationEnabled(v bool)`

SetReauthenticationEnabled sets ReauthenticationEnabled field to given value.


### GetMaxDataRetentionSize

`func (o *KafkaRequestListAllOfItemsInner) GetMaxDataRetentionSize() SupportedKafkaSizeBytesValueItem`

GetMaxDataRetentionSize returns the MaxDataRetentionSize field if non-nil, zero value otherwise.

### GetMaxDataRetentionSizeOk

`func (o *KafkaRequestListAllOfItemsInner) GetMaxDataRetentionSizeOk() (*SupportedKafkaSizeBytesValueItem, bool)`

GetMaxDataRetentionSizeOk returns a tuple with the MaxDataRetentionSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDataRetentionSize

`func (o *KafkaRequestListAllOfItemsInner) SetMaxDataRetentionSize(v SupportedKafkaSizeBytesValueItem)`

SetMaxDataRetentionSize sets MaxDataRetentionSize field to given value.

### HasMaxDataRetentionSize

`func (o *KafkaRequestListAllOfItemsInner) HasMaxDataRetentionSize() bool`

HasMaxDataRetentionSize returns a boolean if a field has been set.

### GetBrowserUrl

`func (o *KafkaRequestListAllOfItemsInner) GetBrowserUrl() string`

GetBrowserUrl returns the BrowserUrl field if non-nil, zero value otherwise.

### GetBrowserUrlOk

`func (o *KafkaRequestListAllOfItemsInner) GetBrowserUrlOk() (*string, bool)`

GetBrowserUrlOk returns a tuple with the BrowserUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowserUrl

`func (o *KafkaRequestListAllOfItemsInner) SetBrowserUrl(v string)`

SetBrowserUrl sets BrowserUrl field to given value.

### HasBrowserUrl

`func (o *KafkaRequestListAllOfItemsInner) HasBrowserUrl() bool`

HasBrowserUrl returns a boolean if a field has been set.

### GetSizeId

`func (o *KafkaRequestListAllOfItemsInner) GetSizeId() string`

GetSizeId returns the SizeId field if non-nil, zero value otherwise.

### GetSizeIdOk

`func (o *KafkaRequestListAllOfItemsInner) GetSizeIdOk() (*string, bool)`

GetSizeIdOk returns a tuple with the SizeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeId

`func (o *KafkaRequestListAllOfItemsInner) SetSizeId(v string)`

SetSizeId sets SizeId field to given value.

### HasSizeId

`func (o *KafkaRequestListAllOfItemsInner) HasSizeId() bool`

HasSizeId returns a boolean if a field has been set.

### GetIngressThroughputPerSec

`func (o *KafkaRequestListAllOfItemsInner) GetIngressThroughputPerSec() string`

GetIngressThroughputPerSec returns the IngressThroughputPerSec field if non-nil, zero value otherwise.

### GetIngressThroughputPerSecOk

`func (o *KafkaRequestListAllOfItemsInner) GetIngressThroughputPerSecOk() (*string, bool)`

GetIngressThroughputPerSecOk returns a tuple with the IngressThroughputPerSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngressThroughputPerSec

`func (o *KafkaRequestListAllOfItemsInner) SetIngressThroughputPerSec(v string)`

SetIngressThroughputPerSec sets IngressThroughputPerSec field to given value.

### HasIngressThroughputPerSec

`func (o *KafkaRequestListAllOfItemsInner) HasIngressThroughputPerSec() bool`

HasIngressThroughputPerSec returns a boolean if a field has been set.

### GetEgressThroughputPerSec

`func (o *KafkaRequestListAllOfItemsInner) GetEgressThroughputPerSec() string`

GetEgressThroughputPerSec returns the EgressThroughputPerSec field if non-nil, zero value otherwise.

### GetEgressThroughputPerSecOk

`func (o *KafkaRequestListAllOfItemsInner) GetEgressThroughputPerSecOk() (*string, bool)`

GetEgressThroughputPerSecOk returns a tuple with the EgressThroughputPerSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEgressThroughputPerSec

`func (o *KafkaRequestListAllOfItemsInner) SetEgressThroughputPerSec(v string)`

SetEgressThroughputPerSec sets EgressThroughputPerSec field to given value.

### HasEgressThroughputPerSec

`func (o *KafkaRequestListAllOfItemsInner) HasEgressThroughputPerSec() bool`

HasEgressThroughputPerSec returns a boolean if a field has been set.

### GetTotalMaxConnections

`func (o *KafkaRequestListAllOfItemsInner) GetTotalMaxConnections() int32`

GetTotalMaxConnections returns the TotalMaxConnections field if non-nil, zero value otherwise.

### GetTotalMaxConnectionsOk

`func (o *KafkaRequestListAllOfItemsInner) GetTotalMaxConnectionsOk() (*int32, bool)`

GetTotalMaxConnectionsOk returns a tuple with the TotalMaxConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMaxConnections

`func (o *KafkaRequestListAllOfItemsInner) SetTotalMaxConnections(v int32)`

SetTotalMaxConnections sets TotalMaxConnections field to given value.

### HasTotalMaxConnections

`func (o *KafkaRequestListAllOfItemsInner) HasTotalMaxConnections() bool`

HasTotalMaxConnections returns a boolean if a field has been set.

### GetMaxPartitions

`func (o *KafkaRequestListAllOfItemsInner) GetMaxPartitions() int32`

GetMaxPartitions returns the MaxPartitions field if non-nil, zero value otherwise.

### GetMaxPartitionsOk

`func (o *KafkaRequestListAllOfItemsInner) GetMaxPartitionsOk() (*int32, bool)`

GetMaxPartitionsOk returns a tuple with the MaxPartitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPartitions

`func (o *KafkaRequestListAllOfItemsInner) SetMaxPartitions(v int32)`

SetMaxPartitions sets MaxPartitions field to given value.

### HasMaxPartitions

`func (o *KafkaRequestListAllOfItemsInner) HasMaxPartitions() bool`

HasMaxPartitions returns a boolean if a field has been set.

### GetMaxDataRetentionPeriod

`func (o *KafkaRequestListAllOfItemsInner) GetMaxDataRetentionPeriod() string`

GetMaxDataRetentionPeriod returns the MaxDataRetentionPeriod field if non-nil, zero value otherwise.

### GetMaxDataRetentionPeriodOk

`func (o *KafkaRequestListAllOfItemsInner) GetMaxDataRetentionPeriodOk() (*string, bool)`

GetMaxDataRetentionPeriodOk returns a tuple with the MaxDataRetentionPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDataRetentionPeriod

`func (o *KafkaRequestListAllOfItemsInner) SetMaxDataRetentionPeriod(v string)`

SetMaxDataRetentionPeriod sets MaxDataRetentionPeriod field to given value.

### HasMaxDataRetentionPeriod

`func (o *KafkaRequestListAllOfItemsInner) HasMaxDataRetentionPeriod() bool`

HasMaxDataRetentionPeriod returns a boolean if a field has been set.

### GetMaxConnectionAttemptsPerSec

`func (o *KafkaRequestListAllOfItemsInner) GetMaxConnectionAttemptsPerSec() int32`

GetMaxConnectionAttemptsPerSec returns the MaxConnectionAttemptsPerSec field if non-nil, zero value otherwise.

### GetMaxConnectionAttemptsPerSecOk

`func (o *KafkaRequestListAllOfItemsInner) GetMaxConnectionAttemptsPerSecOk() (*int32, bool)`

GetMaxConnectionAttemptsPerSecOk returns a tuple with the MaxConnectionAttemptsPerSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConnectionAttemptsPerSec

`func (o *KafkaRequestListAllOfItemsInner) SetMaxConnectionAttemptsPerSec(v int32)`

SetMaxConnectionAttemptsPerSec sets MaxConnectionAttemptsPerSec field to given value.

### HasMaxConnectionAttemptsPerSec

`func (o *KafkaRequestListAllOfItemsInner) HasMaxConnectionAttemptsPerSec() bool`

HasMaxConnectionAttemptsPerSec returns a boolean if a field has been set.

### GetBillingCloudAccountId

`func (o *KafkaRequestListAllOfItemsInner) GetBillingCloudAccountId() string`

GetBillingCloudAccountId returns the BillingCloudAccountId field if non-nil, zero value otherwise.

### GetBillingCloudAccountIdOk

`func (o *KafkaRequestListAllOfItemsInner) GetBillingCloudAccountIdOk() (*string, bool)`

GetBillingCloudAccountIdOk returns a tuple with the BillingCloudAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingCloudAccountId

`func (o *KafkaRequestListAllOfItemsInner) SetBillingCloudAccountId(v string)`

SetBillingCloudAccountId sets BillingCloudAccountId field to given value.

### HasBillingCloudAccountId

`func (o *KafkaRequestListAllOfItemsInner) HasBillingCloudAccountId() bool`

HasBillingCloudAccountId returns a boolean if a field has been set.

### GetMarketplace

`func (o *KafkaRequestListAllOfItemsInner) GetMarketplace() string`

GetMarketplace returns the Marketplace field if non-nil, zero value otherwise.

### GetMarketplaceOk

`func (o *KafkaRequestListAllOfItemsInner) GetMarketplaceOk() (*string, bool)`

GetMarketplaceOk returns a tuple with the Marketplace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketplace

`func (o *KafkaRequestListAllOfItemsInner) SetMarketplace(v string)`

SetMarketplace sets Marketplace field to given value.

### HasMarketplace

`func (o *KafkaRequestListAllOfItemsInner) HasMarketplace() bool`

HasMarketplace returns a boolean if a field has been set.

### GetBillingModel

`func (o *KafkaRequestListAllOfItemsInner) GetBillingModel() string`

GetBillingModel returns the BillingModel field if non-nil, zero value otherwise.

### GetBillingModelOk

`func (o *KafkaRequestListAllOfItemsInner) GetBillingModelOk() (*string, bool)`

GetBillingModelOk returns a tuple with the BillingModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingModel

`func (o *KafkaRequestListAllOfItemsInner) SetBillingModel(v string)`

SetBillingModel sets BillingModel field to given value.

### HasBillingModel

`func (o *KafkaRequestListAllOfItemsInner) HasBillingModel() bool`

HasBillingModel returns a boolean if a field has been set.

### GetPromotionStatus

`func (o *KafkaRequestListAllOfItemsInner) GetPromotionStatus() string`

GetPromotionStatus returns the PromotionStatus field if non-nil, zero value otherwise.

### GetPromotionStatusOk

`func (o *KafkaRequestListAllOfItemsInner) GetPromotionStatusOk() (*string, bool)`

GetPromotionStatusOk returns a tuple with the PromotionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromotionStatus

`func (o *KafkaRequestListAllOfItemsInner) SetPromotionStatus(v string)`

SetPromotionStatus sets PromotionStatus field to given value.

### HasPromotionStatus

`func (o *KafkaRequestListAllOfItemsInner) HasPromotionStatus() bool`

HasPromotionStatus returns a boolean if a field has been set.

### GetClusterId

`func (o *KafkaRequestListAllOfItemsInner) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *KafkaRequestListAllOfItemsInner) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *KafkaRequestListAllOfItemsInner) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *KafkaRequestListAllOfItemsInner) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### SetClusterIdNil

`func (o *KafkaRequestListAllOfItemsInner) SetClusterIdNil(b bool)`

 SetClusterIdNil sets the value for ClusterId to be an explicit nil

### UnsetClusterId
`func (o *KafkaRequestListAllOfItemsInner) UnsetClusterId()`

UnsetClusterId ensures that no value is present for ClusterId, not even an explicit nil
### GetPromotionDetails

`func (o *KafkaRequestListAllOfItemsInner) GetPromotionDetails() string`

GetPromotionDetails returns the PromotionDetails field if non-nil, zero value otherwise.

### GetPromotionDetailsOk

`func (o *KafkaRequestListAllOfItemsInner) GetPromotionDetailsOk() (*string, bool)`

GetPromotionDetailsOk returns a tuple with the PromotionDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromotionDetails

`func (o *KafkaRequestListAllOfItemsInner) SetPromotionDetails(v string)`

SetPromotionDetails sets PromotionDetails field to given value.

### HasPromotionDetails

`func (o *KafkaRequestListAllOfItemsInner) HasPromotionDetails() bool`

HasPromotionDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


