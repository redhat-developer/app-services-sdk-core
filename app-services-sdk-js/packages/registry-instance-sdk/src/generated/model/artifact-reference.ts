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



/**
 * A reference to a different artifact. Typically used with artifact types that can have dependencies like Protobuf.
 * @export
 * @interface ArtifactReference
 */
export interface ArtifactReference {
    /**
     * 
     * @type {string}
     * @memberof ArtifactReference
     */
    'groupId': string;
    /**
     * 
     * @type {string}
     * @memberof ArtifactReference
     */
    'artifactId': string;
    /**
     * 
     * @type {string}
     * @memberof ArtifactReference
     */
    'version'?: string;
    /**
     * 
     * @type {string}
     * @memberof ArtifactReference
     */
    'name': string;
}

