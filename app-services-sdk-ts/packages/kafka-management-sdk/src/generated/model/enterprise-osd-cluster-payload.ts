/* tslint:disable */
/* eslint-disable */
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
 * Schema for the request body sent to /clusters POST
 * @export
 * @interface EnterpriseOsdClusterPayload
 */
export interface EnterpriseOsdClusterPayload {
    /**
     * Sets whether Kafkas created on this data plane cluster have to be accessed via private network
     * @type {boolean}
     * @memberof EnterpriseOsdClusterPayload
     */
    'access_kafkas_via_private_network': boolean;
    /**
     * The data plane cluster ID. This is the ID of the cluster obtained from OpenShift Cluster Manager (OCM) API
     * @type {string}
     * @memberof EnterpriseOsdClusterPayload
     */
    'cluster_id': string;
    /**
     * dns name of the cluster. Can be obtained from the response JSON of the /api/clusters_mgmt/v1/clusters/<cluster_id>/ingresses (dns_name)
     * @type {string}
     * @memberof EnterpriseOsdClusterPayload
     */
    'cluster_ingress_dns_name': string;
    /**
     * The node count given to the created kafka machine pool.  The machine pool must be created via /api/clusters_mgmt/v1/clusters/<cluster_id>/machine_pools prior to passing this value. The created machine pool must have a `bf2.org/kafkaInstanceProfileType=standard` label and a `bf2.org/kafkaInstanceProfileType=standard:NoExecute` taint. The name of the machine pool must be `kafka-standard`  The node count value has to be a multiple of 3 with a minimum of 3 nodes.
     * @type {number}
     * @memberof EnterpriseOsdClusterPayload
     */
    'kafka_machine_pool_node_count': number;
}
