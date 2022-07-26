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


import { ModelError } from './model-error';
import { RuleViolationCause } from './rule-violation-cause';
import { RuleViolationErrorAllOf } from './rule-violation-error-all-of';

/**
 * @type RuleViolationError
 * All error responses, whether `4xx` or `5xx` will include one of these as the response body.
 * @export
 */
export type RuleViolationError = ModelError & RuleViolationErrorAllOf;


