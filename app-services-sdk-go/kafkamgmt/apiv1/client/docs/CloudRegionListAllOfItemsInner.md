# CloudRegionListAllOfItemsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | Pointer to **string** | Indicates the type of this object. Will be &#39;CloudRegion&#39;. | [optional] 
**Id** | Pointer to **string** | Unique identifier of the object. | [optional] 
**DisplayName** | Pointer to **string** | Name of the region for display purposes, for example &#x60;N. Virginia&#x60;. | [optional] 
**Enabled** | **bool** | Whether the region is enabled for deploying an OSD cluster. | [default to false]
**Capacity** | [**[]RegionCapacityListItem**](RegionCapacityListItem.md) | Indicates whether there is capacity left per instance type | 

## Methods

### NewCloudRegionListAllOfItemsInner

`func NewCloudRegionListAllOfItemsInner(enabled bool, capacity []RegionCapacityListItem, ) *CloudRegionListAllOfItemsInner`

NewCloudRegionListAllOfItemsInner instantiates a new CloudRegionListAllOfItemsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudRegionListAllOfItemsInnerWithDefaults

`func NewCloudRegionListAllOfItemsInnerWithDefaults() *CloudRegionListAllOfItemsInner`

NewCloudRegionListAllOfItemsInnerWithDefaults instantiates a new CloudRegionListAllOfItemsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *CloudRegionListAllOfItemsInner) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CloudRegionListAllOfItemsInner) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CloudRegionListAllOfItemsInner) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *CloudRegionListAllOfItemsInner) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *CloudRegionListAllOfItemsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudRegionListAllOfItemsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudRegionListAllOfItemsInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CloudRegionListAllOfItemsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDisplayName

`func (o *CloudRegionListAllOfItemsInner) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CloudRegionListAllOfItemsInner) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CloudRegionListAllOfItemsInner) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *CloudRegionListAllOfItemsInner) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetEnabled

`func (o *CloudRegionListAllOfItemsInner) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CloudRegionListAllOfItemsInner) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CloudRegionListAllOfItemsInner) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetCapacity

`func (o *CloudRegionListAllOfItemsInner) GetCapacity() []RegionCapacityListItem`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *CloudRegionListAllOfItemsInner) GetCapacityOk() (*[]RegionCapacityListItem, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *CloudRegionListAllOfItemsInner) SetCapacity(v []RegionCapacityListItem)`

SetCapacity sets Capacity field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

