# MetricsInstantQueryListAllOfItemsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metric** | Pointer to **map[string]string** |  | [optional] 
**Timestamp** | Pointer to **int64** |  | [optional] 
**Value** | **float64** |  | 

## Methods

### NewMetricsInstantQueryListAllOfItemsInner

`func NewMetricsInstantQueryListAllOfItemsInner(value float64, ) *MetricsInstantQueryListAllOfItemsInner`

NewMetricsInstantQueryListAllOfItemsInner instantiates a new MetricsInstantQueryListAllOfItemsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricsInstantQueryListAllOfItemsInnerWithDefaults

`func NewMetricsInstantQueryListAllOfItemsInnerWithDefaults() *MetricsInstantQueryListAllOfItemsInner`

NewMetricsInstantQueryListAllOfItemsInnerWithDefaults instantiates a new MetricsInstantQueryListAllOfItemsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetric

`func (o *MetricsInstantQueryListAllOfItemsInner) GetMetric() map[string]string`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *MetricsInstantQueryListAllOfItemsInner) GetMetricOk() (*map[string]string, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *MetricsInstantQueryListAllOfItemsInner) SetMetric(v map[string]string)`

SetMetric sets Metric field to given value.

### HasMetric

`func (o *MetricsInstantQueryListAllOfItemsInner) HasMetric() bool`

HasMetric returns a boolean if a field has been set.

### GetTimestamp

`func (o *MetricsInstantQueryListAllOfItemsInner) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *MetricsInstantQueryListAllOfItemsInner) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *MetricsInstantQueryListAllOfItemsInner) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *MetricsInstantQueryListAllOfItemsInner) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetValue

`func (o *MetricsInstantQueryListAllOfItemsInner) GetValue() float64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *MetricsInstantQueryListAllOfItemsInner) GetValueOk() (*float64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *MetricsInstantQueryListAllOfItemsInner) SetValue(v float64)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


