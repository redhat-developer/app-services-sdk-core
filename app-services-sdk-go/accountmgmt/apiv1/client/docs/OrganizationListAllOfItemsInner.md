# OrganizationListAllOfItemsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Capabilities** | Pointer to [**[]Capability**](Capability.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**EbsAccountId** | Pointer to **string** |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 
**Labels** | Pointer to [**[]Label**](Label.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewOrganizationListAllOfItemsInner

`func NewOrganizationListAllOfItemsInner() *OrganizationListAllOfItemsInner`

NewOrganizationListAllOfItemsInner instantiates a new OrganizationListAllOfItemsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationListAllOfItemsInnerWithDefaults

`func NewOrganizationListAllOfItemsInnerWithDefaults() *OrganizationListAllOfItemsInner`

NewOrganizationListAllOfItemsInnerWithDefaults instantiates a new OrganizationListAllOfItemsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *OrganizationListAllOfItemsInner) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *OrganizationListAllOfItemsInner) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *OrganizationListAllOfItemsInner) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *OrganizationListAllOfItemsInner) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *OrganizationListAllOfItemsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrganizationListAllOfItemsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrganizationListAllOfItemsInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OrganizationListAllOfItemsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *OrganizationListAllOfItemsInner) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *OrganizationListAllOfItemsInner) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *OrganizationListAllOfItemsInner) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *OrganizationListAllOfItemsInner) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetCapabilities

`func (o *OrganizationListAllOfItemsInner) GetCapabilities() []Capability`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *OrganizationListAllOfItemsInner) GetCapabilitiesOk() (*[]Capability, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *OrganizationListAllOfItemsInner) SetCapabilities(v []Capability)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *OrganizationListAllOfItemsInner) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### GetCreatedAt

`func (o *OrganizationListAllOfItemsInner) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *OrganizationListAllOfItemsInner) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *OrganizationListAllOfItemsInner) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *OrganizationListAllOfItemsInner) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetEbsAccountId

`func (o *OrganizationListAllOfItemsInner) GetEbsAccountId() string`

GetEbsAccountId returns the EbsAccountId field if non-nil, zero value otherwise.

### GetEbsAccountIdOk

`func (o *OrganizationListAllOfItemsInner) GetEbsAccountIdOk() (*string, bool)`

GetEbsAccountIdOk returns a tuple with the EbsAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEbsAccountId

`func (o *OrganizationListAllOfItemsInner) SetEbsAccountId(v string)`

SetEbsAccountId sets EbsAccountId field to given value.

### HasEbsAccountId

`func (o *OrganizationListAllOfItemsInner) HasEbsAccountId() bool`

HasEbsAccountId returns a boolean if a field has been set.

### GetExternalId

`func (o *OrganizationListAllOfItemsInner) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *OrganizationListAllOfItemsInner) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *OrganizationListAllOfItemsInner) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *OrganizationListAllOfItemsInner) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetLabels

`func (o *OrganizationListAllOfItemsInner) GetLabels() []Label`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *OrganizationListAllOfItemsInner) GetLabelsOk() (*[]Label, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *OrganizationListAllOfItemsInner) SetLabels(v []Label)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *OrganizationListAllOfItemsInner) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetName

`func (o *OrganizationListAllOfItemsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrganizationListAllOfItemsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrganizationListAllOfItemsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OrganizationListAllOfItemsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *OrganizationListAllOfItemsInner) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *OrganizationListAllOfItemsInner) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *OrganizationListAllOfItemsInner) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *OrganizationListAllOfItemsInner) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


