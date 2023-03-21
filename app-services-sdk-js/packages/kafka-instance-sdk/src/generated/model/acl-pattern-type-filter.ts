/* tslint:disable */
/* eslint-disable */
/**
 * Kafka Instance API
 * API for interacting with Kafka Instance. Includes Produce, Consume and Admin APIs
 *
 * The version of the OpenAPI document: 0.14.1-SNAPSHOT
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */



/**
 * Use value \'MATCH\' to perform pattern matching.
 * @export
 * @enum {string}
 */

export const AclPatternTypeFilter = {
    Literal: 'LITERAL',
    Prefixed: 'PREFIXED',
    Any: 'ANY',
    Match: 'MATCH'
} as const;

export type AclPatternTypeFilter = typeof AclPatternTypeFilter[keyof typeof AclPatternTypeFilter];



