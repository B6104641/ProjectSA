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
    EntCompany,
    EntCompanyFromJSON,
    EntCompanyFromJSONTyped,
    EntCompanyToJSON,
    EntEquipment,
    EntEquipmentFromJSON,
    EntEquipmentFromJSONTyped,
    EntEquipmentToJSON,
    EntUser,
    EntUserFromJSON,
    EntUserFromJSONTyped,
    EntUserToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntOrderEdges
 */
export interface EntOrderEdges {
    /**
     * 
     * @type {EntCompany}
     * @memberof EntOrderEdges
     */
    company?: EntCompany;
    /**
     * 
     * @type {EntEquipment}
     * @memberof EntOrderEdges
     */
    equipment?: EntEquipment;
    /**
     * 
     * @type {EntUser}
     * @memberof EntOrderEdges
     */
    user?: EntUser;
}

export function EntOrderEdgesFromJSON(json: any): EntOrderEdges {
    return EntOrderEdgesFromJSONTyped(json, false);
}

export function EntOrderEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntOrderEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'company': !exists(json, 'Company') ? undefined : EntCompanyFromJSON(json['Company']),
        'equipment': !exists(json, 'Equipment') ? undefined : EntEquipmentFromJSON(json['Equipment']),
        'user': !exists(json, 'User') ? undefined : EntUserFromJSON(json['User']),
    };
}

export function EntOrderEdgesToJSON(value?: EntOrderEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'company': EntCompanyToJSON(value.company),
        'equipment': EntEquipmentToJSON(value.equipment),
        'user': EntUserToJSON(value.user),
    };
}

