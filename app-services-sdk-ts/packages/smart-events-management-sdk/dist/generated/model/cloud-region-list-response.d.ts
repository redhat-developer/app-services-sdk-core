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
import { CloudRegionResponse } from './cloud-region-response';
/**
 *
 * @export
 * @interface CloudRegionListResponse
 */
export interface CloudRegionListResponse {
    /**
     *
     * @type {string}
     * @memberof CloudRegionListResponse
     */
    'kind': string;
    /**
     *
     * @type {Array<CloudRegionResponse>}
     * @memberof CloudRegionListResponse
     */
    'items'?: Array<CloudRegionResponse>;
    /**
     *
     * @type {number}
     * @memberof CloudRegionListResponse
     */
    'page': number;
    /**
     *
     * @type {number}
     * @memberof CloudRegionListResponse
     */
    'size': number;
    /**
     *
     * @type {number}
     * @memberof CloudRegionListResponse
     */
    'total': number;
}
