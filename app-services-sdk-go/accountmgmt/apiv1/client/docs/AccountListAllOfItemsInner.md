# AccountListAllOfItemsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**BanCode** | Pointer to **string** |  | [optional] 
**BanDescription** | Pointer to **string** |  | [optional] 
**Banned** | Pointer to **bool** |  | [optional] [default to false]
**Capabilities** | Pointer to [**[]Capability**](Capability.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**FirstName** | Pointer to **string** |  | [optional] 
**Labels** | Pointer to [**[]Label**](Label.md) |  | [optional] 
**LastName** | Pointer to **string** |  | [optional] 
**Organization** | Pointer to [**Organization**](Organization.md) |  | [optional] 
**OrganizationId** | Pointer to **string** |  | [optional] 
**RhitAccountId** | Pointer to **string** |  | [optional] 
**RhitWebUserId** | Pointer to **string** |  | [optional] 
**ServiceAccount** | Pointer to **bool** |  | [optional] [default to false]
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**Username** | **string** |  | 

## Methods

### NewAccountListAllOfItemsInner

`func NewAccountListAllOfItemsInner(username string, ) *AccountListAllOfItemsInner`

NewAccountListAllOfItemsInner instantiates a new AccountListAllOfItemsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountListAllOfItemsInnerWithDefaults

`func NewAccountListAllOfItemsInnerWithDefaults() *AccountListAllOfItemsInner`

NewAccountListAllOfItemsInnerWithDefaults instantiates a new AccountListAllOfItemsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *AccountListAllOfItemsInner) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *AccountListAllOfItemsInner) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *AccountListAllOfItemsInner) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *AccountListAllOfItemsInner) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *AccountListAllOfItemsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccountListAllOfItemsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccountListAllOfItemsInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AccountListAllOfItemsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *AccountListAllOfItemsInner) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *AccountListAllOfItemsInner) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *AccountListAllOfItemsInner) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *AccountListAllOfItemsInner) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetBanCode

`func (o *AccountListAllOfItemsInner) GetBanCode() string`

GetBanCode returns the BanCode field if non-nil, zero value otherwise.

### GetBanCodeOk

`func (o *AccountListAllOfItemsInner) GetBanCodeOk() (*string, bool)`

GetBanCodeOk returns a tuple with the BanCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBanCode

`func (o *AccountListAllOfItemsInner) SetBanCode(v string)`

SetBanCode sets BanCode field to given value.

### HasBanCode

`func (o *AccountListAllOfItemsInner) HasBanCode() bool`

HasBanCode returns a boolean if a field has been set.

### GetBanDescription

`func (o *AccountListAllOfItemsInner) GetBanDescription() string`

GetBanDescription returns the BanDescription field if non-nil, zero value otherwise.

### GetBanDescriptionOk

`func (o *AccountListAllOfItemsInner) GetBanDescriptionOk() (*string, bool)`

GetBanDescriptionOk returns a tuple with the BanDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBanDescription

`func (o *AccountListAllOfItemsInner) SetBanDescription(v string)`

SetBanDescription sets BanDescription field to given value.

### HasBanDescription

`func (o *AccountListAllOfItemsInner) HasBanDescription() bool`

HasBanDescription returns a boolean if a field has been set.

### GetBanned

`func (o *AccountListAllOfItemsInner) GetBanned() bool`

GetBanned returns the Banned field if non-nil, zero value otherwise.

### GetBannedOk

`func (o *AccountListAllOfItemsInner) GetBannedOk() (*bool, bool)`

GetBannedOk returns a tuple with the Banned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBanned

`func (o *AccountListAllOfItemsInner) SetBanned(v bool)`

SetBanned sets Banned field to given value.

### HasBanned

`func (o *AccountListAllOfItemsInner) HasBanned() bool`

HasBanned returns a boolean if a field has been set.

### GetCapabilities

`func (o *AccountListAllOfItemsInner) GetCapabilities() []Capability`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *AccountListAllOfItemsInner) GetCapabilitiesOk() (*[]Capability, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *AccountListAllOfItemsInner) SetCapabilities(v []Capability)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *AccountListAllOfItemsInner) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AccountListAllOfItemsInner) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AccountListAllOfItemsInner) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AccountListAllOfItemsInner) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AccountListAllOfItemsInner) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetEmail

`func (o *AccountListAllOfItemsInner) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *AccountListAllOfItemsInner) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *AccountListAllOfItemsInner) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *AccountListAllOfItemsInner) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetFirstName

`func (o *AccountListAllOfItemsInner) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *AccountListAllOfItemsInner) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *AccountListAllOfItemsInner) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *AccountListAllOfItemsInner) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLabels

`func (o *AccountListAllOfItemsInner) GetLabels() []Label`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *AccountListAllOfItemsInner) GetLabelsOk() (*[]Label, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *AccountListAllOfItemsInner) SetLabels(v []Label)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *AccountListAllOfItemsInner) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetLastName

`func (o *AccountListAllOfItemsInner) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *AccountListAllOfItemsInner) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *AccountListAllOfItemsInner) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *AccountListAllOfItemsInner) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetOrganization

`func (o *AccountListAllOfItemsInner) GetOrganization() Organization`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *AccountListAllOfItemsInner) GetOrganizationOk() (*Organization, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *AccountListAllOfItemsInner) SetOrganization(v Organization)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *AccountListAllOfItemsInner) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetOrganizationId

`func (o *AccountListAllOfItemsInner) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *AccountListAllOfItemsInner) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *AccountListAllOfItemsInner) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *AccountListAllOfItemsInner) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetRhitAccountId

`func (o *AccountListAllOfItemsInner) GetRhitAccountId() string`

GetRhitAccountId returns the RhitAccountId field if non-nil, zero value otherwise.

### GetRhitAccountIdOk

`func (o *AccountListAllOfItemsInner) GetRhitAccountIdOk() (*string, bool)`

GetRhitAccountIdOk returns a tuple with the RhitAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRhitAccountId

`func (o *AccountListAllOfItemsInner) SetRhitAccountId(v string)`

SetRhitAccountId sets RhitAccountId field to given value.

### HasRhitAccountId

`func (o *AccountListAllOfItemsInner) HasRhitAccountId() bool`

HasRhitAccountId returns a boolean if a field has been set.

### GetRhitWebUserId

`func (o *AccountListAllOfItemsInner) GetRhitWebUserId() string`

GetRhitWebUserId returns the RhitWebUserId field if non-nil, zero value otherwise.

### GetRhitWebUserIdOk

`func (o *AccountListAllOfItemsInner) GetRhitWebUserIdOk() (*string, bool)`

GetRhitWebUserIdOk returns a tuple with the RhitWebUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRhitWebUserId

`func (o *AccountListAllOfItemsInner) SetRhitWebUserId(v string)`

SetRhitWebUserId sets RhitWebUserId field to given value.

### HasRhitWebUserId

`func (o *AccountListAllOfItemsInner) HasRhitWebUserId() bool`

HasRhitWebUserId returns a boolean if a field has been set.

### GetServiceAccount

`func (o *AccountListAllOfItemsInner) GetServiceAccount() bool`

GetServiceAccount returns the ServiceAccount field if non-nil, zero value otherwise.

### GetServiceAccountOk

`func (o *AccountListAllOfItemsInner) GetServiceAccountOk() (*bool, bool)`

GetServiceAccountOk returns a tuple with the ServiceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccount

`func (o *AccountListAllOfItemsInner) SetServiceAccount(v bool)`

SetServiceAccount sets ServiceAccount field to given value.

### HasServiceAccount

`func (o *AccountListAllOfItemsInner) HasServiceAccount() bool`

HasServiceAccount returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AccountListAllOfItemsInner) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AccountListAllOfItemsInner) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AccountListAllOfItemsInner) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AccountListAllOfItemsInner) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUsername

`func (o *AccountListAllOfItemsInner) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *AccountListAllOfItemsInner) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *AccountListAllOfItemsInner) SetUsername(v string)`

SetUsername sets Username field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


