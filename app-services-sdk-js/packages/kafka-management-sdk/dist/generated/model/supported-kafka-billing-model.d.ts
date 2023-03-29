/**
 * Kafka Management API
 * Kafka Management API is a REST API to manage Kafka instances
 *
 * The version of the OpenAPI document: 1.15.0
 * Contact: rhosak-support@redhat.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
/**
 * Supported Kafka Billing Model
 * @export
 * @interface SupportedKafkaBillingModel
 */
export interface SupportedKafkaBillingModel {
    /**
     * Identifier for the Kafka billing model
     * @type {string}
     * @memberof SupportedKafkaBillingModel
     */
    'id': string;
    /**
     * AMS resource to be used. Accepted values: [\'rhosak\']
     * @type {string}
     * @memberof SupportedKafkaBillingModel
     */
    'ams_resource': string;
    /**
     * AMS product to be used. Accepted values: [\'RHOSAK\', \'RHOSAKTrial\', \'RHOSAKEval\', \'RHOSAKCC\']
     * @type {string}
     * @memberof SupportedKafkaBillingModel
     */
    'ams_product': string;
    /**
     * List of AMS available billing models: Accepted values: [\'marketplace\', \'marketplace-rhm\', \'marketplace-aws\']
     * @type {Array<string>}
     * @memberof SupportedKafkaBillingModel
     */
    'ams_billing_models': Array<string>;
}
