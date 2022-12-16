/* tslint:disable */
/* eslint-disable */
/**
 * Kafka Instance API
 * API for interacting with Kafka Instance. Includes Produce, Consume and Admin APIs
 *
 * The version of the OpenAPI document: 0.13.0-SNAPSHOT
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */



/**
 * 
 * @export
 * @enum {string}
 */

export const AclResourceType = {
    Group: 'GROUP',
    Topic: 'TOPIC',
    Cluster: 'CLUSTER',
    TransactionalId: 'TRANSACTIONAL_ID'
} as const;

export type AclResourceType = typeof AclResourceType[keyof typeof AclResourceType];



