# AccountGroupListAllOfItemsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Description** | **string** |  | 
**ManagedBy** | **string** |  | 
**Name** | **string** |  | 
**OrganizationId** | **string** |  | 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewAccountGroupListAllOfItemsInner

`func NewAccountGroupListAllOfItemsInner(description string, managedBy string, name string, organizationId string, ) *AccountGroupListAllOfItemsInner`

NewAccountGroupListAllOfItemsInner instantiates a new AccountGroupListAllOfItemsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountGroupListAllOfItemsInnerWithDefaults

`func NewAccountGroupListAllOfItemsInnerWithDefaults() *AccountGroupListAllOfItemsInner`

NewAccountGroupListAllOfItemsInnerWithDefaults instantiates a new AccountGroupListAllOfItemsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *AccountGroupListAllOfItemsInner) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *AccountGroupListAllOfItemsInner) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *AccountGroupListAllOfItemsInner) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *AccountGroupListAllOfItemsInner) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *AccountGroupListAllOfItemsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccountGroupListAllOfItemsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccountGroupListAllOfItemsInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AccountGroupListAllOfItemsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *AccountGroupListAllOfItemsInner) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *AccountGroupListAllOfItemsInner) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *AccountGroupListAllOfItemsInner) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *AccountGroupListAllOfItemsInner) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AccountGroupListAllOfItemsInner) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AccountGroupListAllOfItemsInner) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AccountGroupListAllOfItemsInner) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AccountGroupListAllOfItemsInner) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *AccountGroupListAllOfItemsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AccountGroupListAllOfItemsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AccountGroupListAllOfItemsInner) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetManagedBy

`func (o *AccountGroupListAllOfItemsInner) GetManagedBy() string`

GetManagedBy returns the ManagedBy field if non-nil, zero value otherwise.

### GetManagedByOk

`func (o *AccountGroupListAllOfItemsInner) GetManagedByOk() (*string, bool)`

GetManagedByOk returns a tuple with the ManagedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedBy

`func (o *AccountGroupListAllOfItemsInner) SetManagedBy(v string)`

SetManagedBy sets ManagedBy field to given value.


### GetName

`func (o *AccountGroupListAllOfItemsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccountGroupListAllOfItemsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccountGroupListAllOfItemsInner) SetName(v string)`

SetName sets Name field to given value.


### GetOrganizationId

`func (o *AccountGroupListAllOfItemsInner) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *AccountGroupListAllOfItemsInner) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *AccountGroupListAllOfItemsInner) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.


### GetUpdatedAt

`func (o *AccountGroupListAllOfItemsInner) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AccountGroupListAllOfItemsInner) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AccountGroupListAllOfItemsInner) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AccountGroupListAllOfItemsInner) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


