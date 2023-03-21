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


// May contain unused imports in some cases
// @ts-ignore
import { ManagedResourceStatus } from './managed-resource-status';

/**
 * 
 * @export
 * @interface SinkConnectorResponse
 */
export interface SinkConnectorResponse {
    /**
     * The kind (type) of this resource
     * @type {string}
     * @memberof SinkConnectorResponse
     */
    'kind': string;
    /**
     * The unique identifier of this resource
     * @type {string}
     * @memberof SinkConnectorResponse
     */
    'id': string;
    /**
     * The name of this resource
     * @type {string}
     * @memberof SinkConnectorResponse
     */
    'name': string;
    /**
     * The URL of this resource, without the protocol
     * @type {string}
     * @memberof SinkConnectorResponse
     */
    'href': string;
    /**
     * 
     * @type {string}
     * @memberof SinkConnectorResponse
     */
    'submitted_at': string;
    /**
     * 
     * @type {string}
     * @memberof SinkConnectorResponse
     */
    'published_at'?: string;
    /**
     * 
     * @type {string}
     * @memberof SinkConnectorResponse
     */
    'modified_at'?: string;
    /**
     * 
     * @type {ManagedResourceStatus}
     * @memberof SinkConnectorResponse
     */
    'status': ManagedResourceStatus;
    /**
     * The user that owns this resource
     * @type {string}
     * @memberof SinkConnectorResponse
     */
    'owner': string;
    /**
     * The connector type
     * @type {string}
     * @memberof SinkConnectorResponse
     */
    'connector_type_id': string;
    /**
     * The Connector configuration payload
     * @type {object}
     * @memberof SinkConnectorResponse
     */
    'connector': object;
    /**
     * A detailed status message in case there is a problem with the connector
     * @type {string}
     * @memberof SinkConnectorResponse
     */
    'status_message'?: string;
    /**
     * The URI to be used in Camel DSL to send data to this sink
     * @type {string}
     * @memberof SinkConnectorResponse
     */
    'uri_dsl': string;
}

