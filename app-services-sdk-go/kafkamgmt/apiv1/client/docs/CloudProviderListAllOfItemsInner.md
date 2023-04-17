# CloudProviderListAllOfItemsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | Pointer to **string** | Indicates the type of this object. Will be &#39;CloudProvider&#39; link. | [optional] 
**Id** | Pointer to **string** | Unique identifier of the object. | [optional] 
**DisplayName** | Pointer to **string** | Name of the cloud provider for display purposes. | [optional] 
**Name** | Pointer to **string** | Human friendly identifier of the cloud provider, for example &#x60;aws&#x60;. | [optional] 
**Enabled** | **bool** | Whether the cloud provider is enabled for deploying an OSD cluster. | 

## Methods

### NewCloudProviderListAllOfItemsInner

`func NewCloudProviderListAllOfItemsInner(enabled bool, ) *CloudProviderListAllOfItemsInner`

NewCloudProviderListAllOfItemsInner instantiates a new CloudProviderListAllOfItemsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudProviderListAllOfItemsInnerWithDefaults

`func NewCloudProviderListAllOfItemsInnerWithDefaults() *CloudProviderListAllOfItemsInner`

NewCloudProviderListAllOfItemsInnerWithDefaults instantiates a new CloudProviderListAllOfItemsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *CloudProviderListAllOfItemsInner) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CloudProviderListAllOfItemsInner) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CloudProviderListAllOfItemsInner) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *CloudProviderListAllOfItemsInner) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *CloudProviderListAllOfItemsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudProviderListAllOfItemsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudProviderListAllOfItemsInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CloudProviderListAllOfItemsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDisplayName

`func (o *CloudProviderListAllOfItemsInner) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CloudProviderListAllOfItemsInner) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CloudProviderListAllOfItemsInner) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *CloudProviderListAllOfItemsInner) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetName

`func (o *CloudProviderListAllOfItemsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudProviderListAllOfItemsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudProviderListAllOfItemsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudProviderListAllOfItemsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEnabled

`func (o *CloudProviderListAllOfItemsInner) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CloudProviderListAllOfItemsInner) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CloudProviderListAllOfItemsInner) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


