# SupportedKafkaInstanceTypesListAllOfInstanceTypesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique identifier of the Kafka instance type. | [optional] 
**DisplayName** | Pointer to **string** | Human readable name of the supported Kafka instance type | [optional] 
**SupportedBillingModels** | [**[]SupportedKafkaBillingModel**](SupportedKafkaBillingModel.md) | A list of available kafka billing models for the instance type. Each kafka billing model item has a unique &#39;id&#39; | 
**Sizes** | [**[]SupportedKafkaInstanceTypeSizesInner**](SupportedKafkaInstanceTypeSizesInner.md) | A list of Kafka instance sizes available for this instance type | 

## Methods

### NewSupportedKafkaInstanceTypesListAllOfInstanceTypesInner

`func NewSupportedKafkaInstanceTypesListAllOfInstanceTypesInner(supportedBillingModels []SupportedKafkaBillingModel, sizes []SupportedKafkaInstanceTypeSizesInner, ) *SupportedKafkaInstanceTypesListAllOfInstanceTypesInner`

NewSupportedKafkaInstanceTypesListAllOfInstanceTypesInner instantiates a new SupportedKafkaInstanceTypesListAllOfInstanceTypesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupportedKafkaInstanceTypesListAllOfInstanceTypesInnerWithDefaults

`func NewSupportedKafkaInstanceTypesListAllOfInstanceTypesInnerWithDefaults() *SupportedKafkaInstanceTypesListAllOfInstanceTypesInner`

NewSupportedKafkaInstanceTypesListAllOfInstanceTypesInnerWithDefaults instantiates a new SupportedKafkaInstanceTypesListAllOfInstanceTypesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SupportedKafkaInstanceTypesListAllOfInstanceTypesInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SupportedKafkaInstanceTypesListAllOfInstanceTypesInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SupportedKafkaInstanceTypesListAllOfInstanceTypesInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SupportedKafkaInstanceTypesListAllOfInstanceTypesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDisplayName

`func (o *SupportedKafkaInstanceTypesListAllOfInstanceTypesInner) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *SupportedKafkaInstanceTypesListAllOfInstanceTypesInner) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *SupportedKafkaInstanceTypesListAllOfInstanceTypesInner) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *SupportedKafkaInstanceTypesListAllOfInstanceTypesInner) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetSupportedBillingModels

`func (o *SupportedKafkaInstanceTypesListAllOfInstanceTypesInner) GetSupportedBillingModels() []SupportedKafkaBillingModel`

GetSupportedBillingModels returns the SupportedBillingModels field if non-nil, zero value otherwise.

### GetSupportedBillingModelsOk

`func (o *SupportedKafkaInstanceTypesListAllOfInstanceTypesInner) GetSupportedBillingModelsOk() (*[]SupportedKafkaBillingModel, bool)`

GetSupportedBillingModelsOk returns a tuple with the SupportedBillingModels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedBillingModels

`func (o *SupportedKafkaInstanceTypesListAllOfInstanceTypesInner) SetSupportedBillingModels(v []SupportedKafkaBillingModel)`

SetSupportedBillingModels sets SupportedBillingModels field to given value.


### GetSizes

`func (o *SupportedKafkaInstanceTypesListAllOfInstanceTypesInner) GetSizes() []SupportedKafkaInstanceTypeSizesInner`

GetSizes returns the Sizes field if non-nil, zero value otherwise.

### GetSizesOk

`func (o *SupportedKafkaInstanceTypesListAllOfInstanceTypesInner) GetSizesOk() (*[]SupportedKafkaInstanceTypeSizesInner, bool)`

GetSizesOk returns a tuple with the Sizes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizes

`func (o *SupportedKafkaInstanceTypesListAllOfInstanceTypesInner) SetSizes(v []SupportedKafkaInstanceTypeSizesInner)`

SetSizes sets Sizes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

