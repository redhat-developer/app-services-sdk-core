/* tslint:disable */
/* eslint-disable */
/**
 * Red Hat Openshift SmartEvents Fleet Manager V1
 * The API exposed by the fleet manager of the SmartEvents service.
 *
 * The version of the OpenAPI document: 0.0.1
 * Contact: openbridge-dev@redhat.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


import { Action } from './action';
import { BaseFilter } from './base-filter';
import { ManagedResourceStatus } from './managed-resource-status';
import { ProcessorType } from './processor-type';
import { Source } from './source';

/**
 * 
 * @export
 * @interface ProcessorResponse
 */
export interface ProcessorResponse {
    /**
     * 
     * @type {string}
     * @memberof ProcessorResponse
     */
    'kind': string;
    /**
     * 
     * @type {string}
     * @memberof ProcessorResponse
     */
    'id': string;
    /**
     * 
     * @type {string}
     * @memberof ProcessorResponse
     */
    'name'?: string;
    /**
     * 
     * @type {string}
     * @memberof ProcessorResponse
     */
    'href': string;
    /**
     * 
     * @type {string}
     * @memberof ProcessorResponse
     */
    'submitted_at': string;
    /**
     * 
     * @type {string}
     * @memberof ProcessorResponse
     */
    'published_at'?: string;
    /**
     * 
     * @type {string}
     * @memberof ProcessorResponse
     */
    'modified_at'?: string;
    /**
     * 
     * @type {ManagedResourceStatus}
     * @memberof ProcessorResponse
     */
    'status': ManagedResourceStatus;
    /**
     * 
     * @type {string}
     * @memberof ProcessorResponse
     */
    'owner': string;
    /**
     * 
     * @type {ProcessorType}
     * @memberof ProcessorResponse
     */
    'type': ProcessorType;
    /**
     * 
     * @type {Set<BaseFilter>}
     * @memberof ProcessorResponse
     */
    'filters'?: Set<BaseFilter>;
    /**
     * 
     * @type {string}
     * @memberof ProcessorResponse
     */
    'transformationTemplate'?: string;
    /**
     * 
     * @type {Action}
     * @memberof ProcessorResponse
     */
    'action'?: Action;
    /**
     * 
     * @type {Source}
     * @memberof ProcessorResponse
     */
    'source'?: Source;
    /**
     * 
     * @type {string}
     * @memberof ProcessorResponse
     */
    'status_message'?: string;
}

