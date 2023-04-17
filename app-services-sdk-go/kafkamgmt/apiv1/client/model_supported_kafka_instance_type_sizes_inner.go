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

// checks if the SupportedKafkaInstanceTypeSizesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SupportedKafkaInstanceTypeSizesInner{}

// SupportedKafkaInstanceTypeSizesInner struct for SupportedKafkaInstanceTypeSizesInner
type SupportedKafkaInstanceTypeSizesInner struct {
	// Unique identifier of this Kafka instance size.
	Id *string `json:"id,omitempty"`
	// Display name of this Kafka instance size.
	DisplayName *string `json:"display_name,omitempty"`
	IngressThroughputPerSec SupportedKafkaSizeBytesValueItem `json:"ingress_throughput_per_sec,omitempty"`
	EgressThroughputPerSec SupportedKafkaSizeBytesValueItem `json:"egress_throughput_per_sec,omitempty"`
	// Maximum amount of total connections available to this Kafka instance size.
	TotalMaxConnections *int32 `json:"total_max_connections,omitempty"`
	MaxDataRetentionSize SupportedKafkaSizeBytesValueItem `json:"max_data_retention_size,omitempty"`
	// Maximum amount of total partitions available to this Kafka instance size.
	MaxPartitions *int32 `json:"max_partitions,omitempty"`
	// Maximum data retention period available to this Kafka instance size.
	MaxDataRetentionPeriod *string `json:"max_data_retention_period,omitempty"`
	// Maximium connection attempts per second available to this Kafka instance size.
	MaxConnectionAttemptsPerSec *int32 `json:"max_connection_attempts_per_sec,omitempty"`
	MaxMessageSize SupportedKafkaSizeBytesValueItem `json:"max_message_size,omitempty"`
	// Minimum number of in-sync replicas.
	MinInSyncReplicas *int32 `json:"min_in_sync_replicas,omitempty"`
	// Replication factor available to this Kafka instance size.
	ReplicationFactor *int32 `json:"replication_factor,omitempty"`
	// List of Availability Zone modes that this Kafka instance size supports. The possible values are \"single\", \"multi\".
	SupportedAzModes []string `json:"supported_az_modes,omitempty"`
	// The limit lifespan of the kafka instance in seconds. If not specified then the instance never expires.
	LifespanSeconds NullableInt32 `json:"lifespan_seconds,omitempty"`
	// Quota consumed by this Kafka instance size.
	QuotaConsumed *int32 `json:"quota_consumed,omitempty"`
	// Quota type used by this Kafka instance size. This is now deprecated, please refer to supported_billing_models at instance-type level instead.
	// Deprecated
	QuotaType *string `json:"quota_type,omitempty"`
	// Data plane cluster capacity consumed by this Kafka instance size.
	CapacityConsumed *int32 `json:"capacity_consumed,omitempty"`
	// Maturity level of the size. Can be \"stable\" or \"preview\".
	MaturityStatus *string `json:"maturity_status,omitempty"`
}

// NewSupportedKafkaInstanceTypeSizesInner instantiates a new SupportedKafkaInstanceTypeSizesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSupportedKafkaInstanceTypeSizesInner() *SupportedKafkaInstanceTypeSizesInner {
	this := SupportedKafkaInstanceTypeSizesInner{}
	return &this
}

// NewSupportedKafkaInstanceTypeSizesInnerWithDefaults instantiates a new SupportedKafkaInstanceTypeSizesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSupportedKafkaInstanceTypeSizesInnerWithDefaults() *SupportedKafkaInstanceTypeSizesInner {
	this := SupportedKafkaInstanceTypeSizesInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SupportedKafkaInstanceTypeSizesInner) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportedKafkaInstanceTypeSizesInner) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SupportedKafkaInstanceTypeSizesInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SupportedKafkaInstanceTypeSizesInner) SetId(v string) {
	o.Id = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *SupportedKafkaInstanceTypeSizesInner) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportedKafkaInstanceTypeSizesInner) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *SupportedKafkaInstanceTypeSizesInner) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *SupportedKafkaInstanceTypeSizesInner) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetIngressThroughputPerSec returns the IngressThroughputPerSec field value if set, zero value otherwise.
func (o *SupportedKafkaInstanceTypeSizesInner) GetIngressThroughputPerSec() SupportedKafkaSizeBytesValueItem {
	if o == nil || IsNil(o.IngressThroughputPerSec) {
		var ret SupportedKafkaSizeBytesValueItem
		return ret
	}
	return o.IngressThroughputPerSec
}

// GetIngressThroughputPerSecOk returns a tuple with the IngressThroughputPerSec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportedKafkaInstanceTypeSizesInner) GetIngressThroughputPerSecOk() (SupportedKafkaSizeBytesValueItem, bool) {
	if o == nil || IsNil(o.IngressThroughputPerSec) {
		return SupportedKafkaSizeBytesValueItem{}, false
	}
	return o.IngressThroughputPerSec, true
}

// HasIngressThroughputPerSec returns a boolean if a field has been set.
func (o *SupportedKafkaInstanceTypeSizesInner) HasIngressThroughputPerSec() bool {
	if o != nil && !IsNil(o.IngressThroughputPerSec) {
		return true
	}

	return false
}

// SetIngressThroughputPerSec gets a reference to the given SupportedKafkaSizeBytesValueItem and assigns it to the IngressThroughputPerSec field.
func (o *SupportedKafkaInstanceTypeSizesInner) SetIngressThroughputPerSec(v SupportedKafkaSizeBytesValueItem) {
	o.IngressThroughputPerSec = v
}

// GetEgressThroughputPerSec returns the EgressThroughputPerSec field value if set, zero value otherwise.
func (o *SupportedKafkaInstanceTypeSizesInner) GetEgressThroughputPerSec() SupportedKafkaSizeBytesValueItem {
	if o == nil || IsNil(o.EgressThroughputPerSec) {
		var ret SupportedKafkaSizeBytesValueItem
		return ret
	}
	return o.EgressThroughputPerSec
}

// GetEgressThroughputPerSecOk returns a tuple with the EgressThroughputPerSec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportedKafkaInstanceTypeSizesInner) GetEgressThroughputPerSecOk() (SupportedKafkaSizeBytesValueItem, bool) {
	if o == nil || IsNil(o.EgressThroughputPerSec) {
		return SupportedKafkaSizeBytesValueItem{}, false
	}
	return o.EgressThroughputPerSec, true
}

// HasEgressThroughputPerSec returns a boolean if a field has been set.
func (o *SupportedKafkaInstanceTypeSizesInner) HasEgressThroughputPerSec() bool {
	if o != nil && !IsNil(o.EgressThroughputPerSec) {
		return true
	}

	return false
}

// SetEgressThroughputPerSec gets a reference to the given SupportedKafkaSizeBytesValueItem and assigns it to the EgressThroughputPerSec field.
func (o *SupportedKafkaInstanceTypeSizesInner) SetEgressThroughputPerSec(v SupportedKafkaSizeBytesValueItem) {
	o.EgressThroughputPerSec = v
}

// GetTotalMaxConnections returns the TotalMaxConnections field value if set, zero value otherwise.
func (o *SupportedKafkaInstanceTypeSizesInner) GetTotalMaxConnections() int32 {
	if o == nil || IsNil(o.TotalMaxConnections) {
		var ret int32
		return ret
	}
	return *o.TotalMaxConnections
}

// GetTotalMaxConnectionsOk returns a tuple with the TotalMaxConnections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportedKafkaInstanceTypeSizesInner) GetTotalMaxConnectionsOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalMaxConnections) {
		return nil, false
	}
	return o.TotalMaxConnections, true
}

// HasTotalMaxConnections returns a boolean if a field has been set.
func (o *SupportedKafkaInstanceTypeSizesInner) HasTotalMaxConnections() bool {
	if o != nil && !IsNil(o.TotalMaxConnections) {
		return true
	}

	return false
}

// SetTotalMaxConnections gets a reference to the given int32 and assigns it to the TotalMaxConnections field.
func (o *SupportedKafkaInstanceTypeSizesInner) SetTotalMaxConnections(v int32) {
	o.TotalMaxConnections = &v
}

// GetMaxDataRetentionSize returns the MaxDataRetentionSize field value if set, zero value otherwise.
func (o *SupportedKafkaInstanceTypeSizesInner) GetMaxDataRetentionSize() SupportedKafkaSizeBytesValueItem {
	if o == nil || IsNil(o.MaxDataRetentionSize) {
		var ret SupportedKafkaSizeBytesValueItem
		return ret
	}
	return o.MaxDataRetentionSize
}

// GetMaxDataRetentionSizeOk returns a tuple with the MaxDataRetentionSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportedKafkaInstanceTypeSizesInner) GetMaxDataRetentionSizeOk() (SupportedKafkaSizeBytesValueItem, bool) {
	if o == nil || IsNil(o.MaxDataRetentionSize) {
		return SupportedKafkaSizeBytesValueItem{}, false
	}
	return o.MaxDataRetentionSize, true
}

// HasMaxDataRetentionSize returns a boolean if a field has been set.
func (o *SupportedKafkaInstanceTypeSizesInner) HasMaxDataRetentionSize() bool {
	if o != nil && !IsNil(o.MaxDataRetentionSize) {
		return true
	}

	return false
}

// SetMaxDataRetentionSize gets a reference to the given SupportedKafkaSizeBytesValueItem and assigns it to the MaxDataRetentionSize field.
func (o *SupportedKafkaInstanceTypeSizesInner) SetMaxDataRetentionSize(v SupportedKafkaSizeBytesValueItem) {
	o.MaxDataRetentionSize = v
}

// GetMaxPartitions returns the MaxPartitions field value if set, zero value otherwise.
func (o *SupportedKafkaInstanceTypeSizesInner) GetMaxPartitions() int32 {
	if o == nil || IsNil(o.MaxPartitions) {
		var ret int32
		return ret
	}
	return *o.MaxPartitions
}

// GetMaxPartitionsOk returns a tuple with the MaxPartitions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportedKafkaInstanceTypeSizesInner) GetMaxPartitionsOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxPartitions) {
		return nil, false
	}
	return o.MaxPartitions, true
}

// HasMaxPartitions returns a boolean if a field has been set.
func (o *SupportedKafkaInstanceTypeSizesInner) HasMaxPartitions() bool {
	if o != nil && !IsNil(o.MaxPartitions) {
		return true
	}

	return false
}

// SetMaxPartitions gets a reference to the given int32 and assigns it to the MaxPartitions field.
func (o *SupportedKafkaInstanceTypeSizesInner) SetMaxPartitions(v int32) {
	o.MaxPartitions = &v
}

// GetMaxDataRetentionPeriod returns the MaxDataRetentionPeriod field value if set, zero value otherwise.
func (o *SupportedKafkaInstanceTypeSizesInner) GetMaxDataRetentionPeriod() string {
	if o == nil || IsNil(o.MaxDataRetentionPeriod) {
		var ret string
		return ret
	}
	return *o.MaxDataRetentionPeriod
}

// GetMaxDataRetentionPeriodOk returns a tuple with the MaxDataRetentionPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportedKafkaInstanceTypeSizesInner) GetMaxDataRetentionPeriodOk() (*string, bool) {
	if o == nil || IsNil(o.MaxDataRetentionPeriod) {
		return nil, false
	}
	return o.MaxDataRetentionPeriod, true
}

// HasMaxDataRetentionPeriod returns a boolean if a field has been set.
func (o *SupportedKafkaInstanceTypeSizesInner) HasMaxDataRetentionPeriod() bool {
	if o != nil && !IsNil(o.MaxDataRetentionPeriod) {
		return true
	}

	return false
}

// SetMaxDataRetentionPeriod gets a reference to the given string and assigns it to the MaxDataRetentionPeriod field.
func (o *SupportedKafkaInstanceTypeSizesInner) SetMaxDataRetentionPeriod(v string) {
	o.MaxDataRetentionPeriod = &v
}

// GetMaxConnectionAttemptsPerSec returns the MaxConnectionAttemptsPerSec field value if set, zero value otherwise.
func (o *SupportedKafkaInstanceTypeSizesInner) GetMaxConnectionAttemptsPerSec() int32 {
	if o == nil || IsNil(o.MaxConnectionAttemptsPerSec) {
		var ret int32
		return ret
	}
	return *o.MaxConnectionAttemptsPerSec
}

// GetMaxConnectionAttemptsPerSecOk returns a tuple with the MaxConnectionAttemptsPerSec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportedKafkaInstanceTypeSizesInner) GetMaxConnectionAttemptsPerSecOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxConnectionAttemptsPerSec) {
		return nil, false
	}
	return o.MaxConnectionAttemptsPerSec, true
}

// HasMaxConnectionAttemptsPerSec returns a boolean if a field has been set.
func (o *SupportedKafkaInstanceTypeSizesInner) HasMaxConnectionAttemptsPerSec() bool {
	if o != nil && !IsNil(o.MaxConnectionAttemptsPerSec) {
		return true
	}

	return false
}

// SetMaxConnectionAttemptsPerSec gets a reference to the given int32 and assigns it to the MaxConnectionAttemptsPerSec field.
func (o *SupportedKafkaInstanceTypeSizesInner) SetMaxConnectionAttemptsPerSec(v int32) {
	o.MaxConnectionAttemptsPerSec = &v
}

// GetMaxMessageSize returns the MaxMessageSize field value if set, zero value otherwise.
func (o *SupportedKafkaInstanceTypeSizesInner) GetMaxMessageSize() SupportedKafkaSizeBytesValueItem {
	if o == nil || IsNil(o.MaxMessageSize) {
		var ret SupportedKafkaSizeBytesValueItem
		return ret
	}
	return o.MaxMessageSize
}

// GetMaxMessageSizeOk returns a tuple with the MaxMessageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportedKafkaInstanceTypeSizesInner) GetMaxMessageSizeOk() (SupportedKafkaSizeBytesValueItem, bool) {
	if o == nil || IsNil(o.MaxMessageSize) {
		return SupportedKafkaSizeBytesValueItem{}, false
	}
	return o.MaxMessageSize, true
}

// HasMaxMessageSize returns a boolean if a field has been set.
func (o *SupportedKafkaInstanceTypeSizesInner) HasMaxMessageSize() bool {
	if o != nil && !IsNil(o.MaxMessageSize) {
		return true
	}

	return false
}

// SetMaxMessageSize gets a reference to the given SupportedKafkaSizeBytesValueItem and assigns it to the MaxMessageSize field.
func (o *SupportedKafkaInstanceTypeSizesInner) SetMaxMessageSize(v SupportedKafkaSizeBytesValueItem) {
	o.MaxMessageSize = v
}

// GetMinInSyncReplicas returns the MinInSyncReplicas field value if set, zero value otherwise.
func (o *SupportedKafkaInstanceTypeSizesInner) GetMinInSyncReplicas() int32 {
	if o == nil || IsNil(o.MinInSyncReplicas) {
		var ret int32
		return ret
	}
	return *o.MinInSyncReplicas
}

// GetMinInSyncReplicasOk returns a tuple with the MinInSyncReplicas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportedKafkaInstanceTypeSizesInner) GetMinInSyncReplicasOk() (*int32, bool) {
	if o == nil || IsNil(o.MinInSyncReplicas) {
		return nil, false
	}
	return o.MinInSyncReplicas, true
}

// HasMinInSyncReplicas returns a boolean if a field has been set.
func (o *SupportedKafkaInstanceTypeSizesInner) HasMinInSyncReplicas() bool {
	if o != nil && !IsNil(o.MinInSyncReplicas) {
		return true
	}

	return false
}

// SetMinInSyncReplicas gets a reference to the given int32 and assigns it to the MinInSyncReplicas field.
func (o *SupportedKafkaInstanceTypeSizesInner) SetMinInSyncReplicas(v int32) {
	o.MinInSyncReplicas = &v
}

// GetReplicationFactor returns the ReplicationFactor field value if set, zero value otherwise.
func (o *SupportedKafkaInstanceTypeSizesInner) GetReplicationFactor() int32 {
	if o == nil || IsNil(o.ReplicationFactor) {
		var ret int32
		return ret
	}
	return *o.ReplicationFactor
}

// GetReplicationFactorOk returns a tuple with the ReplicationFactor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportedKafkaInstanceTypeSizesInner) GetReplicationFactorOk() (*int32, bool) {
	if o == nil || IsNil(o.ReplicationFactor) {
		return nil, false
	}
	return o.ReplicationFactor, true
}

// HasReplicationFactor returns a boolean if a field has been set.
func (o *SupportedKafkaInstanceTypeSizesInner) HasReplicationFactor() bool {
	if o != nil && !IsNil(o.ReplicationFactor) {
		return true
	}

	return false
}

// SetReplicationFactor gets a reference to the given int32 and assigns it to the ReplicationFactor field.
func (o *SupportedKafkaInstanceTypeSizesInner) SetReplicationFactor(v int32) {
	o.ReplicationFactor = &v
}

// GetSupportedAzModes returns the SupportedAzModes field value if set, zero value otherwise.
func (o *SupportedKafkaInstanceTypeSizesInner) GetSupportedAzModes() []string {
	if o == nil || IsNil(o.SupportedAzModes) {
		var ret []string
		return ret
	}
	return o.SupportedAzModes
}

// GetSupportedAzModesOk returns a tuple with the SupportedAzModes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportedKafkaInstanceTypeSizesInner) GetSupportedAzModesOk() ([]string, bool) {
	if o == nil || IsNil(o.SupportedAzModes) {
		return nil, false
	}
	return o.SupportedAzModes, true
}

// HasSupportedAzModes returns a boolean if a field has been set.
func (o *SupportedKafkaInstanceTypeSizesInner) HasSupportedAzModes() bool {
	if o != nil && !IsNil(o.SupportedAzModes) {
		return true
	}

	return false
}

// SetSupportedAzModes gets a reference to the given []string and assigns it to the SupportedAzModes field.
func (o *SupportedKafkaInstanceTypeSizesInner) SetSupportedAzModes(v []string) {
	o.SupportedAzModes = v
}

// GetLifespanSeconds returns the LifespanSeconds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SupportedKafkaInstanceTypeSizesInner) GetLifespanSeconds() int32 {
	if o == nil || IsNil(o.LifespanSeconds.Get()) {
		var ret int32
		return ret
	}
	return *o.LifespanSeconds.Get()
}

// GetLifespanSecondsOk returns a tuple with the LifespanSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SupportedKafkaInstanceTypeSizesInner) GetLifespanSecondsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.LifespanSeconds.Get(), o.LifespanSeconds.IsSet()
}

// HasLifespanSeconds returns a boolean if a field has been set.
func (o *SupportedKafkaInstanceTypeSizesInner) HasLifespanSeconds() bool {
	if o != nil && o.LifespanSeconds.IsSet() {
		return true
	}

	return false
}

// SetLifespanSeconds gets a reference to the given NullableInt32 and assigns it to the LifespanSeconds field.
func (o *SupportedKafkaInstanceTypeSizesInner) SetLifespanSeconds(v int32) {
	o.LifespanSeconds.Set(&v)
}
// SetLifespanSecondsNil sets the value for LifespanSeconds to be an explicit nil
func (o *SupportedKafkaInstanceTypeSizesInner) SetLifespanSecondsNil() {
	o.LifespanSeconds.Set(nil)
}

// UnsetLifespanSeconds ensures that no value is present for LifespanSeconds, not even an explicit nil
func (o *SupportedKafkaInstanceTypeSizesInner) UnsetLifespanSeconds() {
	o.LifespanSeconds.Unset()
}

// GetQuotaConsumed returns the QuotaConsumed field value if set, zero value otherwise.
func (o *SupportedKafkaInstanceTypeSizesInner) GetQuotaConsumed() int32 {
	if o == nil || IsNil(o.QuotaConsumed) {
		var ret int32
		return ret
	}
	return *o.QuotaConsumed
}

// GetQuotaConsumedOk returns a tuple with the QuotaConsumed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportedKafkaInstanceTypeSizesInner) GetQuotaConsumedOk() (*int32, bool) {
	if o == nil || IsNil(o.QuotaConsumed) {
		return nil, false
	}
	return o.QuotaConsumed, true
}

// HasQuotaConsumed returns a boolean if a field has been set.
func (o *SupportedKafkaInstanceTypeSizesInner) HasQuotaConsumed() bool {
	if o != nil && !IsNil(o.QuotaConsumed) {
		return true
	}

	return false
}

// SetQuotaConsumed gets a reference to the given int32 and assigns it to the QuotaConsumed field.
func (o *SupportedKafkaInstanceTypeSizesInner) SetQuotaConsumed(v int32) {
	o.QuotaConsumed = &v
}

// GetQuotaType returns the QuotaType field value if set, zero value otherwise.
// Deprecated
func (o *SupportedKafkaInstanceTypeSizesInner) GetQuotaType() string {
	if o == nil || IsNil(o.QuotaType) {
		var ret string
		return ret
	}
	return *o.QuotaType
}

// GetQuotaTypeOk returns a tuple with the QuotaType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *SupportedKafkaInstanceTypeSizesInner) GetQuotaTypeOk() (*string, bool) {
	if o == nil || IsNil(o.QuotaType) {
		return nil, false
	}
	return o.QuotaType, true
}

// HasQuotaType returns a boolean if a field has been set.
func (o *SupportedKafkaInstanceTypeSizesInner) HasQuotaType() bool {
	if o != nil && !IsNil(o.QuotaType) {
		return true
	}

	return false
}

// SetQuotaType gets a reference to the given string and assigns it to the QuotaType field.
// Deprecated
func (o *SupportedKafkaInstanceTypeSizesInner) SetQuotaType(v string) {
	o.QuotaType = &v
}

// GetCapacityConsumed returns the CapacityConsumed field value if set, zero value otherwise.
func (o *SupportedKafkaInstanceTypeSizesInner) GetCapacityConsumed() int32 {
	if o == nil || IsNil(o.CapacityConsumed) {
		var ret int32
		return ret
	}
	return *o.CapacityConsumed
}

// GetCapacityConsumedOk returns a tuple with the CapacityConsumed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportedKafkaInstanceTypeSizesInner) GetCapacityConsumedOk() (*int32, bool) {
	if o == nil || IsNil(o.CapacityConsumed) {
		return nil, false
	}
	return o.CapacityConsumed, true
}

// HasCapacityConsumed returns a boolean if a field has been set.
func (o *SupportedKafkaInstanceTypeSizesInner) HasCapacityConsumed() bool {
	if o != nil && !IsNil(o.CapacityConsumed) {
		return true
	}

	return false
}

// SetCapacityConsumed gets a reference to the given int32 and assigns it to the CapacityConsumed field.
func (o *SupportedKafkaInstanceTypeSizesInner) SetCapacityConsumed(v int32) {
	o.CapacityConsumed = &v
}

// GetMaturityStatus returns the MaturityStatus field value if set, zero value otherwise.
func (o *SupportedKafkaInstanceTypeSizesInner) GetMaturityStatus() string {
	if o == nil || IsNil(o.MaturityStatus) {
		var ret string
		return ret
	}
	return *o.MaturityStatus
}

// GetMaturityStatusOk returns a tuple with the MaturityStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportedKafkaInstanceTypeSizesInner) GetMaturityStatusOk() (*string, bool) {
	if o == nil || IsNil(o.MaturityStatus) {
		return nil, false
	}
	return o.MaturityStatus, true
}

// HasMaturityStatus returns a boolean if a field has been set.
func (o *SupportedKafkaInstanceTypeSizesInner) HasMaturityStatus() bool {
	if o != nil && !IsNil(o.MaturityStatus) {
		return true
	}

	return false
}

// SetMaturityStatus gets a reference to the given string and assigns it to the MaturityStatus field.
func (o *SupportedKafkaInstanceTypeSizesInner) SetMaturityStatus(v string) {
	o.MaturityStatus = &v
}

func (o SupportedKafkaInstanceTypeSizesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SupportedKafkaInstanceTypeSizesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.DisplayName) {
		toSerialize["display_name"] = o.DisplayName
	}
	if !IsNil(o.IngressThroughputPerSec) {
		toSerialize["ingress_throughput_per_sec"] = o.IngressThroughputPerSec
	}
	if !IsNil(o.EgressThroughputPerSec) {
		toSerialize["egress_throughput_per_sec"] = o.EgressThroughputPerSec
	}
	if !IsNil(o.TotalMaxConnections) {
		toSerialize["total_max_connections"] = o.TotalMaxConnections
	}
	if !IsNil(o.MaxDataRetentionSize) {
		toSerialize["max_data_retention_size"] = o.MaxDataRetentionSize
	}
	if !IsNil(o.MaxPartitions) {
		toSerialize["max_partitions"] = o.MaxPartitions
	}
	if !IsNil(o.MaxDataRetentionPeriod) {
		toSerialize["max_data_retention_period"] = o.MaxDataRetentionPeriod
	}
	if !IsNil(o.MaxConnectionAttemptsPerSec) {
		toSerialize["max_connection_attempts_per_sec"] = o.MaxConnectionAttemptsPerSec
	}
	if !IsNil(o.MaxMessageSize) {
		toSerialize["max_message_size"] = o.MaxMessageSize
	}
	if !IsNil(o.MinInSyncReplicas) {
		toSerialize["min_in_sync_replicas"] = o.MinInSyncReplicas
	}
	if !IsNil(o.ReplicationFactor) {
		toSerialize["replication_factor"] = o.ReplicationFactor
	}
	if !IsNil(o.SupportedAzModes) {
		toSerialize["supported_az_modes"] = o.SupportedAzModes
	}
	if o.LifespanSeconds.IsSet() {
		toSerialize["lifespan_seconds"] = o.LifespanSeconds.Get()
	}
	if !IsNil(o.QuotaConsumed) {
		toSerialize["quota_consumed"] = o.QuotaConsumed
	}
	if !IsNil(o.QuotaType) {
		toSerialize["quota_type"] = o.QuotaType
	}
	if !IsNil(o.CapacityConsumed) {
		toSerialize["capacity_consumed"] = o.CapacityConsumed
	}
	if !IsNil(o.MaturityStatus) {
		toSerialize["maturity_status"] = o.MaturityStatus
	}
	return toSerialize, nil
}

type NullableSupportedKafkaInstanceTypeSizesInner struct {
	value *SupportedKafkaInstanceTypeSizesInner
	isSet bool
}

func (v NullableSupportedKafkaInstanceTypeSizesInner) Get() *SupportedKafkaInstanceTypeSizesInner {
	return v.value
}

func (v *NullableSupportedKafkaInstanceTypeSizesInner) Set(val *SupportedKafkaInstanceTypeSizesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableSupportedKafkaInstanceTypeSizesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableSupportedKafkaInstanceTypeSizesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSupportedKafkaInstanceTypeSizesInner(val *SupportedKafkaInstanceTypeSizesInner) *NullableSupportedKafkaInstanceTypeSizesInner {
	return &NullableSupportedKafkaInstanceTypeSizesInner{value: val, isSet: true}
}

func (v NullableSupportedKafkaInstanceTypeSizesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSupportedKafkaInstanceTypeSizesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


