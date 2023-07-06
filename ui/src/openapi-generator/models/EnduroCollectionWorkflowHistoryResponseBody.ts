/* tslint:disable */
/* eslint-disable */
/**
 * Enduro API
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * The version of the OpenAPI document: 
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
/**
 * WorkflowHistoryEvent describes a history event in Temporal. (default view)
 * @export
 * @interface EnduroCollectionWorkflowHistoryResponseBody
 */
export interface EnduroCollectionWorkflowHistoryResponseBody {
    /**
     * Contents of the event
     * @type {Blob}
     * @memberof EnduroCollectionWorkflowHistoryResponseBody
     */
    details?: Blob;
    /**
     * Identifier of collection
     * @type {number}
     * @memberof EnduroCollectionWorkflowHistoryResponseBody
     */
    id?: number;
    /**
     * Type of the event
     * @type {string}
     * @memberof EnduroCollectionWorkflowHistoryResponseBody
     */
    type?: string;
}

export function EnduroCollectionWorkflowHistoryResponseBodyFromJSON(json: any): EnduroCollectionWorkflowHistoryResponseBody {
    return EnduroCollectionWorkflowHistoryResponseBodyFromJSONTyped(json, false);
}

export function EnduroCollectionWorkflowHistoryResponseBodyFromJSONTyped(json: any, ignoreDiscriminator: boolean): EnduroCollectionWorkflowHistoryResponseBody {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'details': !exists(json, 'details') ? undefined : json['details'],
        'id': !exists(json, 'id') ? undefined : json['id'],
        'type': !exists(json, 'type') ? undefined : json['type'],
    };
}

export function EnduroCollectionWorkflowHistoryResponseBodyToJSON(value?: EnduroCollectionWorkflowHistoryResponseBody | null): any {
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


