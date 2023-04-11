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
/**
 *
 * @export
 * @interface BridgeRequest
 */
export interface BridgeRequest {
    /**
     * The name of the bridge
     * @type {string}
     * @memberof BridgeRequest
     */
    'name': string;
    /**
     * The cloud provider where the bridge resides
     * @type {string}
     * @memberof BridgeRequest
     */
    'cloud_provider': string;
    /**
     * The cloud provider region where the bridge resides
     * @type {string}
     * @memberof BridgeRequest
     */
    'region': string;
}
