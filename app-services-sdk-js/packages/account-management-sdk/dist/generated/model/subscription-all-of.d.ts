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
import { Account } from './account';
import { AccountReference } from './account-reference';
import { Capability } from './capability';
import { Label } from './label';
import { OneMetric } from './one-metric';
import { Plan } from './plan';
/**
 *
 * @export
 * @interface SubscriptionAllOf
 */
export interface SubscriptionAllOf {
    /**
     *
     * @type {Array<Capability>}
     * @memberof SubscriptionAllOf
     */
    'capabilities'?: Array<Capability>;
    /**
     *
     * @type {string}
     * @memberof SubscriptionAllOf
     */
    'created_at'?: string;
    /**
     *
     * @type {AccountReference}
     * @memberof SubscriptionAllOf
     */
    'creator'?: AccountReference;
    /**
     * Calulated as the subscription created date + 60 days
     * @type {string}
     * @memberof SubscriptionAllOf
     */
    'eval_expiration_date'?: string;
    /**
     *
     * @type {Array<Label>}
     * @memberof SubscriptionAllOf
     */
    'labels'?: Array<Label>;
    /**
     *
     * @type {Array<OneMetric>}
     * @memberof SubscriptionAllOf
     */
    'metrics'?: Array<OneMetric>;
    /**
     *
     * @type {Array<Account>}
     * @memberof SubscriptionAllOf
     */
    'notification_contacts'?: Array<Account>;
    /**
     *
     * @type {Plan}
     * @memberof SubscriptionAllOf
     */
    'plan'?: Plan;
    /**
     *
     * @type {string}
     * @memberof SubscriptionAllOf
     */
    'updated_at'?: string;
}
