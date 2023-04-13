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
 * @interface RedHatErrorRepresentation
 */
export interface RedHatErrorRepresentation {
    /**
     * 
     * @type {string}
     * @memberof RedHatErrorRepresentation
     */
    'error'?: RedHatErrorRepresentationErrorEnum;
    /**
     * 
     * @type {string}
     * @memberof RedHatErrorRepresentation
     */
    'error_description'?: string;
}

export const RedHatErrorRepresentationErrorEnum = {
    ServiceAccountLimitExceeded: 'service_account_limit_exceeded',
    ServiceAccountNotFound: 'service_account_not_found',
    ServiceAccountUserNotFound: 'service_account_user_not_found',
    ServiceAccountAccessInvalid: 'service_account_access_invalid',
    AcsTenantLimitExceeded: 'acs_tenant_limit_exceeded',
    AcsTenantNotFound: 'acs_tenant_not_found',
    AcsAccessInvalid: 'acs_access_invalid',
    AcsInvalidRedirectUri: 'acs_invalid_redirect_uri',
    AcsInvalidClient: 'acs_invalid_client',
    AcsDisabled: 'acs_disabled',
    SmoketestAccessInvalid: 'smoketest_access_invalid',
    SmoketestNotFound: 'smoketest_not_found',
    GeneralFailure: 'general_failure',
    OrganizationApiAccessInvalid: 'organization_api_access_invalid'
} as const;

export type RedHatErrorRepresentationErrorEnum = typeof RedHatErrorRepresentationErrorEnum[keyof typeof RedHatErrorRepresentationErrorEnum];

