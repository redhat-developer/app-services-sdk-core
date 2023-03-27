/**
 * Kafka Instance API
 * API for interacting with Kafka Instance. Includes Produce, Consume and Admin APIs
 *
 * The version of the OpenAPI document: 0.14.1-SNAPSHOT
 *
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import { Consumer } from './consumer';
import { ConsumerGroupMetrics } from './consumer-group-metrics';
import { ConsumerGroupState } from './consumer-group-state';
/**
 * A group of Kafka consumers
 * @export
 * @interface ConsumerGroupAllOf
 */
export interface ConsumerGroupAllOf {
    /**
     * Unique identifier for the consumer group
     * @type {string}
     * @memberof ConsumerGroupAllOf
     */
    'groupId': string;
    /**
     *
     * @type {ConsumerGroupState}
     * @memberof ConsumerGroupAllOf
     */
    'state'?: ConsumerGroupState;
    /**
     * The list of consumers associated with this consumer group
     * @type {Array<Consumer>}
     * @memberof ConsumerGroupAllOf
     */
    'consumers': Array<Consumer>;
    /**
     *
     * @type {ConsumerGroupMetrics}
     * @memberof ConsumerGroupAllOf
     */
    'metrics'?: ConsumerGroupMetrics;
}
