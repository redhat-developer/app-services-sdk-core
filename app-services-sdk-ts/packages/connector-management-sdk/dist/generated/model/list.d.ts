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
 * @interface List
 */
export interface List {
    /**
     *
     * @type {string}
     * @memberof List
     */
    'kind': string;
    /**
     *
     * @type {number}
     * @memberof List
     */
    'page': number;
    /**
     *
     * @type {number}
     * @memberof List
     */
    'size': number;
    /**
     *
     * @type {number}
     * @memberof List
     */
    'total': number;
}
