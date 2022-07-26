/* tslint:disable */
/* eslint-disable */
/**
 * Service Registry API
 * Service Registry Instance API  NOTE: This API cannot be called directly from the portal.
 *
 * The version of the OpenAPI document: 2.2.5.Final
 * Contact: apicurio@lists.jboss.org
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


import { ArtifactReference } from './artifact-reference';
import { ArtifactState } from './artifact-state';
import { ArtifactType } from './artifact-type';

/**
 * Models a single artifact from the result set returned when searching for artifacts.
 * @export
 * @interface SearchedVersion
 */
export interface SearchedVersion {
    /**
     * 
     * @type {string}
     * @memberof SearchedVersion
     */
    'name'?: string;
    /**
     * 
     * @type {string}
     * @memberof SearchedVersion
     */
    'description'?: string;
    /**
     * 
     * @type {string}
     * @memberof SearchedVersion
     */
    'createdOn': string;
    /**
     * 
     * @type {string}
     * @memberof SearchedVersion
     */
    'createdBy': string;
    /**
     * 
     * @type {ArtifactType}
     * @memberof SearchedVersion
     */
    'type': ArtifactType;
    /**
     * 
     * @type {Array<string>}
     * @memberof SearchedVersion
     */
    'labels'?: Array<string>;
    /**
     * 
     * @type {ArtifactState}
     * @memberof SearchedVersion
     */
    'state': ArtifactState;
    /**
     * 
     * @type {number}
     * @memberof SearchedVersion
     */
    'globalId': number;
    /**
     * 
     * @type {string}
     * @memberof SearchedVersion
     */
    'version': string;
    /**
     * User-defined name-value pairs. Name and value must be strings.
     * @type {{ [key: string]: string; }}
     * @memberof SearchedVersion
     */
    'properties'?: { [key: string]: string; };
    /**
     * 
     * @type {number}
     * @memberof SearchedVersion
     */
    'contentId': number;
    /**
     * 
     * @type {Array<ArtifactReference>}
     * @memberof SearchedVersion
     */
    'references': Array<ArtifactReference>;
}

