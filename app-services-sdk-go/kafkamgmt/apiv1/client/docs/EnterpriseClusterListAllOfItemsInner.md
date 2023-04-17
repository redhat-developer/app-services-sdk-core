# EnterpriseClusterListAllOfItemsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Kind** | **string** |  | 
**Href** | **string** |  | 
**AccessKafkasViaPrivateNetwork** | **bool** | Indicates whether Kafkas created on this data plane cluster have to be accessed via private network | 
**ClusterId** | Pointer to **string** | The OCM&#39;s cluster id of the registered Enterprise cluster. | [optional] 
**Status** | Pointer to **string** | The status of Enterprise cluster registration | [optional] 
**CloudProvider** | Pointer to **string** | The cloud provider for this cluster. This valus will be used as the Kafka&#39;s cloud provider value when a Kafka is created on this cluster | [optional] 
**Region** | Pointer to **string** | The region of this cluster. This valus will be used as the Kafka&#39;s region value when a Kafka is created on this cluster | [optional] 
**MultiAz** | **bool** | A flag indicating whether this cluster is available on multiple availability zones or not | 
**SupportedInstanceTypes** | Pointer to [**SupportedKafkaInstanceTypesList**](SupportedKafkaInstanceTypesList.md) |  | [optional] 
**CapacityInformation** | Pointer to [**EnterpriseClusterAllOfCapacityInformation**](EnterpriseClusterAllOfCapacityInformation.md) |  | [optional] 

## Methods

### NewEnterpriseClusterListAllOfItemsInner

`func NewEnterpriseClusterListAllOfItemsInner(id string, kind string, href string, accessKafkasViaPrivateNetwork bool, multiAz bool, ) *EnterpriseClusterListAllOfItemsInner`

NewEnterpriseClusterListAllOfItemsInner instantiates a new EnterpriseClusterListAllOfItemsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnterpriseClusterListAllOfItemsInnerWithDefaults

`func NewEnterpriseClusterListAllOfItemsInnerWithDefaults() *EnterpriseClusterListAllOfItemsInner`

NewEnterpriseClusterListAllOfItemsInnerWithDefaults instantiates a new EnterpriseClusterListAllOfItemsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EnterpriseClusterListAllOfItemsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EnterpriseClusterListAllOfItemsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EnterpriseClusterListAllOfItemsInner) SetId(v string)`

SetId sets Id field to given value.


### GetKind

`func (o *EnterpriseClusterListAllOfItemsInner) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *EnterpriseClusterListAllOfItemsInner) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *EnterpriseClusterListAllOfItemsInner) SetKind(v string)`

SetKind sets Kind field to given value.


### GetHref

`func (o *EnterpriseClusterListAllOfItemsInner) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *EnterpriseClusterListAllOfItemsInner) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *EnterpriseClusterListAllOfItemsInner) SetHref(v string)`

SetHref sets Href field to given value.


### GetAccessKafkasViaPrivateNetwork

`func (o *EnterpriseClusterListAllOfItemsInner) GetAccessKafkasViaPrivateNetwork() bool`

GetAccessKafkasViaPrivateNetwork returns the AccessKafkasViaPrivateNetwork field if non-nil, zero value otherwise.

### GetAccessKafkasViaPrivateNetworkOk

`func (o *EnterpriseClusterListAllOfItemsInner) GetAccessKafkasViaPrivateNetworkOk() (*bool, bool)`

GetAccessKafkasViaPrivateNetworkOk returns a tuple with the AccessKafkasViaPrivateNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKafkasViaPrivateNetwork

`func (o *EnterpriseClusterListAllOfItemsInner) SetAccessKafkasViaPrivateNetwork(v bool)`

SetAccessKafkasViaPrivateNetwork sets AccessKafkasViaPrivateNetwork field to given value.


### GetClusterId

`func (o *EnterpriseClusterListAllOfItemsInner) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *EnterpriseClusterListAllOfItemsInner) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *EnterpriseClusterListAllOfItemsInner) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *EnterpriseClusterListAllOfItemsInner) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### GetStatus

`func (o *EnterpriseClusterListAllOfItemsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EnterpriseClusterListAllOfItemsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EnterpriseClusterListAllOfItemsInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EnterpriseClusterListAllOfItemsInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCloudProvider

`func (o *EnterpriseClusterListAllOfItemsInner) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *EnterpriseClusterListAllOfItemsInner) GetCloudProviderOk() (*string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *EnterpriseClusterListAllOfItemsInner) SetCloudProvider(v string)`

SetCloudProvider sets CloudProvider field to given value.

### HasCloudProvider

`func (o *EnterpriseClusterListAllOfItemsInner) HasCloudProvider() bool`

HasCloudProvider returns a boolean if a field has been set.

### GetRegion

`func (o *EnterpriseClusterListAllOfItemsInner) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *EnterpriseClusterListAllOfItemsInner) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *EnterpriseClusterListAllOfItemsInner) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *EnterpriseClusterListAllOfItemsInner) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetMultiAz

`func (o *EnterpriseClusterListAllOfItemsInner) GetMultiAz() bool`

GetMultiAz returns the MultiAz field if non-nil, zero value otherwise.

### GetMultiAzOk

`func (o *EnterpriseClusterListAllOfItemsInner) GetMultiAzOk() (*bool, bool)`

GetMultiAzOk returns a tuple with the MultiAz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiAz

`func (o *EnterpriseClusterListAllOfItemsInner) SetMultiAz(v bool)`

SetMultiAz sets MultiAz field to given value.


### GetSupportedInstanceTypes

`func (o *EnterpriseClusterListAllOfItemsInner) GetSupportedInstanceTypes() SupportedKafkaInstanceTypesList`

GetSupportedInstanceTypes returns the SupportedInstanceTypes field if non-nil, zero value otherwise.

### GetSupportedInstanceTypesOk

`func (o *EnterpriseClusterListAllOfItemsInner) GetSupportedInstanceTypesOk() (*SupportedKafkaInstanceTypesList, bool)`

GetSupportedInstanceTypesOk returns a tuple with the SupportedInstanceTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedInstanceTypes

`func (o *EnterpriseClusterListAllOfItemsInner) SetSupportedInstanceTypes(v SupportedKafkaInstanceTypesList)`

SetSupportedInstanceTypes sets SupportedInstanceTypes field to given value.

### HasSupportedInstanceTypes

`func (o *EnterpriseClusterListAllOfItemsInner) HasSupportedInstanceTypes() bool`

HasSupportedInstanceTypes returns a boolean if a field has been set.

### GetCapacityInformation

`func (o *EnterpriseClusterListAllOfItemsInner) GetCapacityInformation() EnterpriseClusterAllOfCapacityInformation`

GetCapacityInformation returns the CapacityInformation field if non-nil, zero value otherwise.

### GetCapacityInformationOk

`func (o *EnterpriseClusterListAllOfItemsInner) GetCapacityInformationOk() (*EnterpriseClusterAllOfCapacityInformation, bool)`

GetCapacityInformationOk returns a tuple with the CapacityInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacityInformation

`func (o *EnterpriseClusterListAllOfItemsInner) SetCapacityInformation(v EnterpriseClusterAllOfCapacityInformation)`

SetCapacityInformation sets CapacityInformation field to given value.

### HasCapacityInformation

`func (o *EnterpriseClusterListAllOfItemsInner) HasCapacityInformation() bool`

HasCapacityInformation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


