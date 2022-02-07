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
 * 
 * @export
 * @interface BatchSubmitRequestBody
 */
export interface BatchSubmitRequestBody {
    /**
     * 
     * @type {string}
     * @memberof BatchSubmitRequestBody
     */
    completedDir?: string;
    /**
     * 
     * @type {string}
     * @memberof BatchSubmitRequestBody
     */
    path: string;
    /**
     * 
     * @type {string}
     * @memberof BatchSubmitRequestBody
     */
    pipeline: string;
    /**
     * 
     * @type {string}
     * @memberof BatchSubmitRequestBody
     */
    processingConfig?: string;
    /**
     * 
     * @type {string}
     * @memberof BatchSubmitRequestBody
     */
    retentionPeriod?: string;
}

export function BatchSubmitRequestBodyFromJSON(json: any): BatchSubmitRequestBody {
    return BatchSubmitRequestBodyFromJSONTyped(json, false);
}

export function BatchSubmitRequestBodyFromJSONTyped(json: any, ignoreDiscriminator: boolean): BatchSubmitRequestBody {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'completedDir': !exists(json, 'completed_dir') ? undefined : json['completed_dir'],
        'path': json['path'],
        'pipeline': json['pipeline'],
        'processingConfig': !exists(json, 'processing_config') ? undefined : json['processing_config'],
        'retentionPeriod': !exists(json, 'retention_period') ? undefined : json['retention_period'],
    };
}

export function BatchSubmitRequestBodyToJSON(value?: BatchSubmitRequestBody | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'completed_dir': value.completedDir,
        'path': value.path,
        'pipeline': value.pipeline,
        'processing_config': value.processingConfig,
        'retention_period': value.retentionPeriod,
    };
}


