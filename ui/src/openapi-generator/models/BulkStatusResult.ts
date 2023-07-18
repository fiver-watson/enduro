/* tslint:disable */
/* eslint-disable */
/**
 * Enduro API
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * The version of the OpenAPI document: 1.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
/**
 * 
 * @export
 * @interface BulkStatusResult
 */
export interface BulkStatusResult {
    /**
     * 
     * @type {Date}
     * @memberof BulkStatusResult
     */
    closedAt?: Date;
    /**
     * 
     * @type {string}
     * @memberof BulkStatusResult
     */
    runId?: string;
    /**
     * 
     * @type {boolean}
     * @memberof BulkStatusResult
     */
    running: boolean;
    /**
     * 
     * @type {Date}
     * @memberof BulkStatusResult
     */
    startedAt?: Date;
    /**
     * 
     * @type {string}
     * @memberof BulkStatusResult
     */
    status?: string;
    /**
     * 
     * @type {string}
     * @memberof BulkStatusResult
     */
    workflowId?: string;
}

/**
 * Check if a given object implements the BulkStatusResult interface.
 */
export function instanceOfBulkStatusResult(value: object): boolean {
    let isInstance = true;
    isInstance = isInstance && "running" in value;

    return isInstance;
}

export function BulkStatusResultFromJSON(json: any): BulkStatusResult {
    return BulkStatusResultFromJSONTyped(json, false);
}

export function BulkStatusResultFromJSONTyped(json: any, ignoreDiscriminator: boolean): BulkStatusResult {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'closedAt': !exists(json, 'closed_at') ? undefined : (new Date(json['closed_at'])),
        'runId': !exists(json, 'run_id') ? undefined : json['run_id'],
        'running': json['running'],
        'startedAt': !exists(json, 'started_at') ? undefined : (new Date(json['started_at'])),
        'status': !exists(json, 'status') ? undefined : json['status'],
        'workflowId': !exists(json, 'workflow_id') ? undefined : json['workflow_id'],
    };
}

export function BulkStatusResultToJSON(value?: BulkStatusResult | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'closed_at': value.closedAt === undefined ? undefined : (value.closedAt.toISOString()),
        'run_id': value.runId,
        'running': value.running,
        'started_at': value.startedAt === undefined ? undefined : (value.startedAt.toISOString()),
        'status': value.status,
        'workflow_id': value.workflowId,
    };
}

