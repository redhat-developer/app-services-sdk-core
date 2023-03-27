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
/**
 * An individual record consumed from a topic or produced to a topic
 * @export
 * @interface RecordAllOf
 */
export interface RecordAllOf {
    /**
     * The record\'s partition within the topic
     * @type {number}
     * @memberof RecordAllOf
     */
    'partition'?: number;
    /**
     * The record\'s offset within the topic partition
     * @type {number}
     * @memberof RecordAllOf
     */
    'offset'?: number;
    /**
     * Timestamp associated with the record. The type is indicated by `timestampType`. When producing a record, this value will be used as the record\'s `CREATE_TIME`.
     * @type {string}
     * @memberof RecordAllOf
     */
    'timestamp'?: string;
    /**
     * Type of timestamp associated with the record
     * @type {string}
     * @memberof RecordAllOf
     */
    'timestampType'?: string;
    /**
     * Record headers, key/value pairs
     * @type {{ [key: string]: string; }}
     * @memberof RecordAllOf
     */
    'headers'?: {
        [key: string]: string;
    };
    /**
     * Record key
     * @type {string}
     * @memberof RecordAllOf
     */
    'key'?: string;
    /**
     * Record value
     * @type {string}
     * @memberof RecordAllOf
     */
    'value': string;
}
