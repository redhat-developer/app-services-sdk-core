/**
 * Account Management Service API
 * Manage user subscriptions and clusters
 *
 * The version of the OpenAPI document: 0.0.1
 *
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import { LandingPageLinkShape } from './landing-page-link-shape';
/**
 *
 * @export
 * @interface LandingPageEstateItem
 */
export interface LandingPageEstateItem {
    /**
     *
     * @type {number}
     * @memberof LandingPageEstateItem
     */
    'count'?: number;
    /**
     *
     * @type {string}
     * @memberof LandingPageEstateItem
     */
    'id'?: string;
    /**
     *
     * @type {LandingPageLinkShape}
     * @memberof LandingPageEstateItem
     */
    'shape'?: LandingPageLinkShape;
}
