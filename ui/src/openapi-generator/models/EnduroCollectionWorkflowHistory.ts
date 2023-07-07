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
 * WorkflowHistoryEvent describes a history event in Temporal.
 * @export
 * @interface EnduroCollectionWorkflowHistory
 */
export interface EnduroCollectionWorkflowHistory {
    /**
     * Contents of the event
     * @type {Blob}
     * @memberof EnduroCollectionWorkflowHistory
     */
    details?: Blob;
    /**
     * Identifier of collection
     * @type {number}
     * @memberof EnduroCollectionWorkflowHistory
     */
    id?: number;
    /**
     * Type of the event
     * @type {string}
     * @memberof EnduroCollectionWorkflowHistory
     */
    type?: string;
}

/**
 * Check if a given object implements the EnduroCollectionWorkflowHistory interface.
 */
export function instanceOfEnduroCollectionWorkflowHistory(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function EnduroCollectionWorkflowHistoryFromJSON(json: any): EnduroCollectionWorkflowHistory {
    return EnduroCollectionWorkflowHistoryFromJSONTyped(json, false);
}

export function EnduroCollectionWorkflowHistoryFromJSONTyped(json: any, ignoreDiscriminator: boolean): EnduroCollectionWorkflowHistory {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'details': !exists(json, 'details') ? undefined : json['details'],
        'id': !exists(json, 'id') ? undefined : json['id'],
        'type': !exists(json, 'type') ? undefined : json['type'],
    };
}

export function EnduroCollectionWorkflowHistoryToJSON(value?: EnduroCollectionWorkflowHistory | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'details': value.details,
        'id': value.id,
        'type': value.type,
    };
}
