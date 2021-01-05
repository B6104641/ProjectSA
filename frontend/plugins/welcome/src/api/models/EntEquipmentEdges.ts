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
    EntOrder,
    EntOrderFromJSON,
    EntOrderFromJSONTyped,
    EntOrderToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntEquipmentEdges
 */
export interface EntEquipmentEdges {
    /**
     * OrderEquipmentt holds the value of the order_equipmentt edge.
     * @type {Array<EntOrder>}
     * @memberof EntEquipmentEdges
     */
    orderEquipmentt?: Array<EntOrder>;
}

export function EntEquipmentEdgesFromJSON(json: any): EntEquipmentEdges {
    return EntEquipmentEdgesFromJSONTyped(json, false);
}

export function EntEquipmentEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntEquipmentEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'orderEquipmentt': !exists(json, 'orderEquipmentt') ? undefined : ((json['orderEquipmentt'] as Array<any>).map(EntOrderFromJSON)),
    };
}

export function EntEquipmentEdgesToJSON(value?: EntEquipmentEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'orderEquipmentt': value.orderEquipmentt === undefined ? undefined : ((value.orderEquipmentt as Array<any>).map(EntOrderToJSON)),
    };
}

