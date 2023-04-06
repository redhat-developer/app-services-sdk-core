/* tslint:disable */
/* eslint-disable */
/**
 * Red Hat Openshift SmartEvents Fleet Manager V2
 * The API exposed by the fleet manager of the SmartEvents service.
 *
 * The version of the OpenAPI document: 0.0.1
 * Contact: openbridge-dev@redhat.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


import { SinkConnectorResponse } from './sink-connector-response';

/**
 * 
 * @export
 * @interface SinkConnectorListResponse
 */
export interface SinkConnectorListResponse {
    /**
     * 
     * @type {string}
     * @memberof SinkConnectorListResponse
     */
    'kind': string;
    /**
     * 
     * @type {Array<SinkConnectorResponse>}
     * @memberof SinkConnectorListResponse
     */
    'items'?: Array<SinkConnectorResponse>;
    /**
     * 
     * @type {number}
     * @memberof SinkConnectorListResponse
     */
    'page': number;
    /**
     * 
     * @type {number}
     * @memberof SinkConnectorListResponse
     */
    'size': number;
    /**
     * 
     * @type {number}
     * @memberof SinkConnectorListResponse
     */
    'total': number;
}
