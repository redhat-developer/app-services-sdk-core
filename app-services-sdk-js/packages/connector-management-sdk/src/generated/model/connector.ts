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


// May contain unused imports in some cases
// @ts-ignore
import { Channel } from './channel';
// May contain unused imports in some cases
// @ts-ignore
import { ConnectorConfiguration } from './connector-configuration';
// May contain unused imports in some cases
// @ts-ignore
import { ConnectorDesiredState } from './connector-desired-state';
// May contain unused imports in some cases
// @ts-ignore
import { ConnectorMeta } from './connector-meta';
// May contain unused imports in some cases
// @ts-ignore
import { ConnectorStatus } from './connector-status';
// May contain unused imports in some cases
// @ts-ignore
import { ConnectorStatusStatus } from './connector-status-status';
// May contain unused imports in some cases
// @ts-ignore
import { KafkaConnectionSettings } from './kafka-connection-settings';
// May contain unused imports in some cases
// @ts-ignore
import { ObjectReference } from './object-reference';
// May contain unused imports in some cases
// @ts-ignore
import { SchemaRegistryConnectionSettings } from './schema-registry-connection-settings';
// May contain unused imports in some cases
// @ts-ignore
import { ServiceAccount } from './service-account';

/**
 * @type Connector
 * @export
 */
export type Connector = ConnectorConfiguration & ConnectorMeta & ConnectorStatus & ObjectReference;


