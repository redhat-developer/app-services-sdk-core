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


import { ConnectorClusterMeta } from './connector-cluster-meta';
import { ConnectorClusterStatus } from './connector-cluster-status';
import { ConnectorClusterStatusStatus } from './connector-cluster-status-status';
import { ObjectReference } from './object-reference';

/**
 * @type ConnectorCluster
 * @export
 */
export type ConnectorCluster = ConnectorClusterMeta & ConnectorClusterStatus & ObjectReference;

