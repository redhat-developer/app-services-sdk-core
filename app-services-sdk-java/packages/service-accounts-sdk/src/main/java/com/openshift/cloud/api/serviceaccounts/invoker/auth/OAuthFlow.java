/*
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


package com.openshift.cloud.api.serviceaccounts.invoker.auth;

@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen")
public enum OAuthFlow {
    accessCode, //called authorizationCode  in OpenAPI 3.0
    implicit,
    password,
    application //called clientCredentials in OpenAPI 3.0
}
