/* tslint:disable */
/* eslint-disable */
/**
 * Connector Management API
 * Connector Management API is a REST API to manage connectors.
 *
 * The version of the OpenAPI document: 0.1.0
 * Contact: rhosak-support@redhat.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


import { ConnectorNamespaceMetaAllOf } from './connector-namespace-meta-all-of';
import { ConnectorNamespaceQuota } from './connector-namespace-quota';
import { ConnectorNamespaceRequestMeta } from './connector-namespace-request-meta';
import { ObjectMeta } from './object-meta';

/**
 * @type ConnectorNamespaceMeta
 * @export
 */
export type ConnectorNamespaceMeta = ConnectorNamespaceMetaAllOf & ConnectorNamespaceRequestMeta & ObjectMeta;

