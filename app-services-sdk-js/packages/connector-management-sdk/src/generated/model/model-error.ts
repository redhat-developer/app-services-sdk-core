/* tslint:disable */
/* eslint-disable */
/**
 * Connector Management API
 * Connector Management API is a REST API to manage connectors.
 *
 * The version of the OpenAPI document: 0.1.0
 * Contact: rhosak-support@redhat.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */



/**
 * 
 * @export
 * @interface ModelError
 */
export interface ModelError {
    /**
     * 
     * @type {string}
     * @memberof ModelError
     */
    'reason': string;
    /**
     * 
     * @type {string}
     * @memberof ModelError
     */
    'operation_id'?: string;
    /**
     * 
     * @type {string}
     * @memberof ModelError
     */
    'id': string;
    /**
     * 
     * @type {string}
     * @memberof ModelError
     */
    'kind': string;
    /**
     * 
     * @type {string}
     * @memberof ModelError
     */
    'href': string;
    /**
     * 
     * @type {string}
     * @memberof ModelError
     */
    'code': string;
}
