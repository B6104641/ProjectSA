/* tslint:disable */
/* eslint-disable */
/**
 * SUT SA Example API Order
 * This is a sample server for SUT SE 2563
 *
 * The version of the OpenAPI document: 1.0
 * Contact: support@swagger.io
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import {
    EntOrderEdges,
    EntOrderEdgesFromJSON,
    EntOrderEdgesFromJSONTyped,
    EntOrderEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntOrder
 */
export interface EntOrder {
    /**
     * ADDEDTIME holds the value of the "ADDED_TIME" field.
     * @type {string}
     * @memberof EntOrder
     */
    aDDEDTIME?: string;
    /**
     * AMOUNT holds the value of the "AMOUNT" field.
     * @type {number}
     * @memberof EntOrder
     */
    aMOUNT?: number;
    /**
     * PRICE holds the value of the "PRICE" field.
     * @type {number}
     * @memberof EntOrder
     */
    pRICE?: number;
    /**
     * 
     * @type {EntOrderEdges}
     * @memberof EntOrder
     */
    edges?: EntOrderEdges;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntOrder
     */
    id?: number;
}

export function EntOrderFromJSON(json: any): EntOrder {
    return EntOrderFromJSONTyped(json, false);
}

export function EntOrderFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntOrder {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'aDDEDTIME': !exists(json, 'ADDED_TIME') ? undefined : json['ADDED_TIME'],
        'aMOUNT': !exists(json, 'AMOUNT') ? undefined : json['AMOUNT'],
        'pRICE': !exists(json, 'PRICE') ? undefined : json['PRICE'],
        'edges': !exists(json, 'edges') ? undefined : EntOrderEdgesFromJSON(json['edges']),
        'id': !exists(json, 'id') ? undefined : json['id'],
    };
}

export function EntOrderToJSON(value?: EntOrder | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'ADDED_TIME': value.aDDEDTIME,
        'AMOUNT': value.aMOUNT,
        'PRICE': value.pRICE,
        'edges': EntOrderEdgesToJSON(value.edges),
        'id': value.id,
    };
}


