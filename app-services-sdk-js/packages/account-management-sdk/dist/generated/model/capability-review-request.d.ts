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
/**
 *
 * @export
 * @interface CapabilityReviewRequest
 */
export interface CapabilityReviewRequest {
    /**
     *
     * @type {string}
     * @memberof CapabilityReviewRequest
     */
    'account_username': string;
    /**
     *
     * @type {string}
     * @memberof CapabilityReviewRequest
     */
    'capability': CapabilityReviewRequestCapabilityEnum;
    /**
     *
     * @type {string}
     * @memberof CapabilityReviewRequest
     */
    'cluster_id'?: string;
    /**
     *
     * @type {string}
     * @memberof CapabilityReviewRequest
     */
    'organization_id'?: string;
    /**
     *
     * @type {string}
     * @memberof CapabilityReviewRequest
     */
    'subscription_id'?: string;
    /**
     *
     * @type {string}
     * @memberof CapabilityReviewRequest
     */
    'type': CapabilityReviewRequestTypeEnum;
}
export declare const CapabilityReviewRequestCapabilityEnum: {
    readonly ManageClusterAdmin: "manage_cluster_admin";
};
export type CapabilityReviewRequestCapabilityEnum = typeof CapabilityReviewRequestCapabilityEnum[keyof typeof CapabilityReviewRequestCapabilityEnum];
export declare const CapabilityReviewRequestTypeEnum: {
    readonly Cluster: "Cluster";
};
export type CapabilityReviewRequestTypeEnum = typeof CapabilityReviewRequestTypeEnum[keyof typeof CapabilityReviewRequestTypeEnum];
