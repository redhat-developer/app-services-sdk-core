/* tslint:disable */
/* eslint-disable */
/**
 * sso.redhat.com API documentation
 * This is the API documentation for sso.redhat.com
 *
 * The version of the OpenAPI document: 5.0.19-SNAPSHOT
 * Contact: it-user-team-list@redhat.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */



/**
 * 
 * @export
 * @interface AcsClientRequestData
 */
export interface AcsClientRequestData {
    /**
     * 
     * @type {string}
     * @memberof AcsClientRequestData
     */
    'name'?: string;
    /**
     * 
     * @type {Set<string>}
     * @memberof AcsClientRequestData
     */
    'redirectUris': Set<string>;
    /**
     * 
     * @type {string}
     * @memberof AcsClientRequestData
     */
    'orgId': string;
}

