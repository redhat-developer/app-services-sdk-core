/**
 * Kafka Management API
 * Kafka Management API is a REST API to manage Kafka instances
 *
 * The version of the OpenAPI document: 1.16.0
 * Contact: rhosak-support@redhat.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
/**
 * Returns the capacity related information
 * @export
 * @interface EnterpriseClusterAllOfCapacityInformation
 */
export interface EnterpriseClusterAllOfCapacityInformation {
    /**
     * The kafka machine pool node count provided during cluster registration
     * @type {number}
     * @memberof EnterpriseClusterAllOfCapacityInformation
     */
    'kafka_machine_pool_node_count': number;
    /**
     * The maximum number of Kafka streaming units that can be created on this cluster
     * @type {number}
     * @memberof EnterpriseClusterAllOfCapacityInformation
     */
    'maximum_kafka_streaming_units': number;
    /**
     * The remaining number of Kafka streaming units that can be still be created on this cluster
     * @type {number}
     * @memberof EnterpriseClusterAllOfCapacityInformation
     */
    'remaining_kafka_streaming_units': number;
    /**
     * The number of Kafka streaming units that have been consumed on this cluster
     * @type {number}
     * @memberof EnterpriseClusterAllOfCapacityInformation
     */
    'consumed_kafka_streaming_units': number;
}
