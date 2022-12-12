/* tslint:disable */
/* eslint-disable */
/**
 * Service Registry Management API
 * Service Registry Management API is a REST API for managing Service Registry instances. Service Registry is a datastore for event schemas and API designs, which is based on the open source Apicurio Registry project.
 *
 * The version of the OpenAPI document: 0.0.6
 * Contact: rhosak-eval-support@redhat.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


import { RegistryInstanceTypeValue } from './registry-instance-type-value';
import { RegistryStatusValue } from './registry-status-value';

/**
 * Service Registry instance in a multi-tenant deployment.
 * @export
 * @interface RootTypeForRegistry
 */
export interface RootTypeForRegistry {
    /**
     * 
     * @type {string}
     * @memberof RootTypeForRegistry
     */
    'id': string;
    /**
     * 
     * @type {RegistryStatusValue}
     * @memberof RootTypeForRegistry
     */
    'status': RegistryStatusValue;
    /**
     * 
     * @type {string}
     * @memberof RootTypeForRegistry
     */
    'registryUrl'?: string;
    /**
     * 
     * @type {string}
     * @memberof RootTypeForRegistry
     */
    'browserUrl'?: string;
    /**
     * User-defined Registry instance name. Does not have to be unique.
     * @type {string}
     * @memberof RootTypeForRegistry
     */
    'name': string;
    /**
     * Identifier of a multi-tenant deployment, where this Service Registry instance resides.
     * @type {number}
     * @memberof RootTypeForRegistry
     */
    'registryDeploymentId'?: number;
    /**
     * Registry instance owner.
     * @type {string}
     * @memberof RootTypeForRegistry
     */
    'owner': string;
    /**
     * Description of the Registry instance.
     * @type {string}
     * @memberof RootTypeForRegistry
     */
    'description'?: string;
    /**
     * ISO 8601 UTC timestamp.
     * @type {string}
     * @memberof RootTypeForRegistry
     */
    'created_at': string;
    /**
     * ISO 8601 UTC timestamp.
     * @type {string}
     * @memberof RootTypeForRegistry
     */
    'updated_at': string;
    /**
     * 
     * @type {RegistryInstanceTypeValue}
     * @memberof RootTypeForRegistry
     */
    'instance_type': RegistryInstanceTypeValue;
}
