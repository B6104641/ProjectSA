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


import * as runtime from '../runtime';
import {
    ControllersOrder,
    ControllersOrderFromJSON,
    ControllersOrderToJSON,
    EntCompany,
    EntCompanyFromJSON,
    EntCompanyToJSON,
    EntEquipment,
    EntEquipmentFromJSON,
    EntEquipmentToJSON,
    EntOrder,
    EntOrderFromJSON,
    EntOrderToJSON,
    EntUser,
    EntUserFromJSON,
    EntUserToJSON,
} from '../models';

export interface CreateCompanyRequest {
    company: EntCompany;
}

export interface CreateEquipmentRequest {
    equipment: EntEquipment;
}

export interface CreateOrderRequest {
    order: ControllersOrder;
}

export interface CreateUserRequest {
    user: EntUser;
}

export interface DeleteCompanyRequest {
    id: number;
}

export interface DeleteEquipmentRequest {
    id: number;
}

export interface DeleteOrderRequest {
    id: number;
}

export interface DeleteUserRequest {
    id: number;
}

export interface GetCompanyRequest {
    id: number;
}

export interface GetEquipmentRequest {
    id: number;
}

export interface GetOrderRequest {
    id: number;
}

export interface GetUserRequest {
    id: number;
}

export interface ListCompanyRequest {
    limit?: number;
    offset?: number;
}

export interface ListEquipmentRequest {
    limit?: number;
    offset?: number;
}

export interface ListOrderRequest {
    limit?: number;
    offset?: number;
}

export interface ListUserRequest {
    limit?: number;
    offset?: number;
}

export interface UpdateCompanyRequest {
    id: number;
    company: EntCompany;
}

export interface UpdateEquipmentRequest {
    id: number;
    equipment: EntEquipment;
}

export interface UpdateOrderRequest {
    id: number;
    order: EntOrder;
}

export interface UpdateUserRequest {
    id: number;
    user: EntUser;
}

/**
 * 
 */
export class DefaultApi extends runtime.BaseAPI {

    /**
     * Create company
     * Create company
     */
    async createCompanyRaw(requestParameters: CreateCompanyRequest): Promise<runtime.ApiResponse<EntCompany>> {
        if (requestParameters.company === null || requestParameters.company === undefined) {
            throw new runtime.RequiredError('company','Required parameter requestParameters.company was null or undefined when calling createCompany.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/companys`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: EntCompanyToJSON(requestParameters.company),
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => EntCompanyFromJSON(jsonValue));
    }

    /**
     * Create company
     * Create company
     */
    async createCompany(requestParameters: CreateCompanyRequest): Promise<EntCompany> {
        const response = await this.createCompanyRaw(requestParameters);
        return await response.value();
    }

    /**
     * Create equipment
     * Create equipment
     */
    async createEquipmentRaw(requestParameters: CreateEquipmentRequest): Promise<runtime.ApiResponse<EntEquipment>> {
        if (requestParameters.equipment === null || requestParameters.equipment === undefined) {
            throw new runtime.RequiredError('equipment','Required parameter requestParameters.equipment was null or undefined when calling createEquipment.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/equipments`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: EntEquipmentToJSON(requestParameters.equipment),
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => EntEquipmentFromJSON(jsonValue));
    }

    /**
     * Create equipment
     * Create equipment
     */
    async createEquipment(requestParameters: CreateEquipmentRequest): Promise<EntEquipment> {
        const response = await this.createEquipmentRaw(requestParameters);
        return await response.value();
    }

    /**
     * Create order
     * Create order
     */
    async createOrderRaw(requestParameters: CreateOrderRequest): Promise<runtime.ApiResponse<EntOrder>> {
        if (requestParameters.order === null || requestParameters.order === undefined) {
            throw new runtime.RequiredError('order','Required parameter requestParameters.order was null or undefined when calling createOrder.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/orders`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: ControllersOrderToJSON(requestParameters.order),
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => EntOrderFromJSON(jsonValue));
    }

    /**
     * Create order
     * Create order
     */
    async createOrder(requestParameters: CreateOrderRequest): Promise<EntOrder> {
        const response = await this.createOrderRaw(requestParameters);
        return await response.value();
    }

    /**
     * Create user
     * Create user
     */
    async createUserRaw(requestParameters: CreateUserRequest): Promise<runtime.ApiResponse<EntUser>> {
        if (requestParameters.user === null || requestParameters.user === undefined) {
            throw new runtime.RequiredError('user','Required parameter requestParameters.user was null or undefined when calling createUser.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/users`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: EntUserToJSON(requestParameters.user),
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => EntUserFromJSON(jsonValue));
    }

    /**
     * Create user
     * Create user
     */
    async createUser(requestParameters: CreateUserRequest): Promise<EntUser> {
        const response = await this.createUserRaw(requestParameters);
        return await response.value();
    }

    /**
     * get company by ID
     * Delete a company entity by ID
     */
    async deleteCompanyRaw(requestParameters: DeleteCompanyRequest): Promise<runtime.ApiResponse<object>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling deleteCompany.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/companys/{id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'DELETE',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse<any>(response);
    }

    /**
     * get company by ID
     * Delete a company entity by ID
     */
    async deleteCompany(requestParameters: DeleteCompanyRequest): Promise<object> {
        const response = await this.deleteCompanyRaw(requestParameters);
        return await response.value();
    }

    /**
     * get equipment by ID
     * Delete a equipment entity by ID
     */
    async deleteEquipmentRaw(requestParameters: DeleteEquipmentRequest): Promise<runtime.ApiResponse<object>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling deleteEquipment.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/equipments/{id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'DELETE',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse<any>(response);
    }

    /**
     * get equipment by ID
     * Delete a equipment entity by ID
     */
    async deleteEquipment(requestParameters: DeleteEquipmentRequest): Promise<object> {
        const response = await this.deleteEquipmentRaw(requestParameters);
        return await response.value();
    }

    /**
     * get order by ID
     * Delete a order entity by ID
     */
    async deleteOrderRaw(requestParameters: DeleteOrderRequest): Promise<runtime.ApiResponse<object>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling deleteOrder.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/orders/{id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'DELETE',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse<any>(response);
    }

    /**
     * get order by ID
     * Delete a order entity by ID
     */
    async deleteOrder(requestParameters: DeleteOrderRequest): Promise<object> {
        const response = await this.deleteOrderRaw(requestParameters);
        return await response.value();
    }

    /**
     * get user by ID
     * Delete a user entity by ID
     */
    async deleteUserRaw(requestParameters: DeleteUserRequest): Promise<runtime.ApiResponse<object>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling deleteUser.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/users/{id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'DELETE',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse<any>(response);
    }

    /**
     * get user by ID
     * Delete a user entity by ID
     */
    async deleteUser(requestParameters: DeleteUserRequest): Promise<object> {
        const response = await this.deleteUserRaw(requestParameters);
        return await response.value();
    }

    /**
     * get company by ID
     * Get a company entity by ID
     */
    async getCompanyRaw(requestParameters: GetCompanyRequest): Promise<runtime.ApiResponse<EntCompany>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling getCompany.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/companys/{id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => EntCompanyFromJSON(jsonValue));
    }

    /**
     * get company by ID
     * Get a company entity by ID
     */
    async getCompany(requestParameters: GetCompanyRequest): Promise<EntCompany> {
        const response = await this.getCompanyRaw(requestParameters);
        return await response.value();
    }

    /**
     * get equipment by ID
     * Get a equipment entity by ID
     */
    async getEquipmentRaw(requestParameters: GetEquipmentRequest): Promise<runtime.ApiResponse<EntEquipment>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling getEquipment.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/equipments/{id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => EntEquipmentFromJSON(jsonValue));
    }

    /**
     * get equipment by ID
     * Get a equipment entity by ID
     */
    async getEquipment(requestParameters: GetEquipmentRequest): Promise<EntEquipment> {
        const response = await this.getEquipmentRaw(requestParameters);
        return await response.value();
    }

    /**
     * get order by ID
     * Get a order entity by ID
     */
    async getOrderRaw(requestParameters: GetOrderRequest): Promise<runtime.ApiResponse<EntOrder>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling getOrder.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/orders/{id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => EntOrderFromJSON(jsonValue));
    }

    /**
     * get order by ID
     * Get a order entity by ID
     */
    async getOrder(requestParameters: GetOrderRequest): Promise<EntOrder> {
        const response = await this.getOrderRaw(requestParameters);
        return await response.value();
    }

    /**
     * get user by ID
     * Get a user entity by ID
     */
    async getUserRaw(requestParameters: GetUserRequest): Promise<runtime.ApiResponse<EntUser>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling getUser.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/users/{id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => EntUserFromJSON(jsonValue));
    }

    /**
     * get user by ID
     * Get a user entity by ID
     */
    async getUser(requestParameters: GetUserRequest): Promise<EntUser> {
        const response = await this.getUserRaw(requestParameters);
        return await response.value();
    }

    /**
     * list company entities
     * List company entities
     */
    async listCompanyRaw(requestParameters: ListCompanyRequest): Promise<runtime.ApiResponse<Array<EntCompany>>> {
        const queryParameters: runtime.HTTPQuery = {};

        if (requestParameters.limit !== undefined) {
            queryParameters['limit'] = requestParameters.limit;
        }

        if (requestParameters.offset !== undefined) {
            queryParameters['offset'] = requestParameters.offset;
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/companys`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => jsonValue.map(EntCompanyFromJSON));
    }

    /**
     * list company entities
     * List company entities
     */
    async listCompany(requestParameters: ListCompanyRequest): Promise<Array<EntCompany>> {
        const response = await this.listCompanyRaw(requestParameters);
        return await response.value();
    }

    /**
     * list equipment entities
     * List equipment entities
     */
    async listEquipmentRaw(requestParameters: ListEquipmentRequest): Promise<runtime.ApiResponse<Array<EntEquipment>>> {
        const queryParameters: runtime.HTTPQuery = {};

        if (requestParameters.limit !== undefined) {
            queryParameters['limit'] = requestParameters.limit;
        }

        if (requestParameters.offset !== undefined) {
            queryParameters['offset'] = requestParameters.offset;
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/equipments`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => jsonValue.map(EntEquipmentFromJSON));
    }

    /**
     * list equipment entities
     * List equipment entities
     */
    async listEquipment(requestParameters: ListEquipmentRequest): Promise<Array<EntEquipment>> {
        const response = await this.listEquipmentRaw(requestParameters);
        return await response.value();
    }

    /**
     * list order entities
     * List order entities
     */
    async listOrderRaw(requestParameters: ListOrderRequest): Promise<runtime.ApiResponse<Array<EntOrder>>> {
        const queryParameters: runtime.HTTPQuery = {};

        if (requestParameters.limit !== undefined) {
            queryParameters['limit'] = requestParameters.limit;
        }

        if (requestParameters.offset !== undefined) {
            queryParameters['offset'] = requestParameters.offset;
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/orders`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => jsonValue.map(EntOrderFromJSON));
    }

    /**
     * list order entities
     * List order entities
     */
    async listOrder(requestParameters: ListOrderRequest): Promise<Array<EntOrder>> {
        const response = await this.listOrderRaw(requestParameters);
        return await response.value();
    }

    /**
     * list user entities
     * List user entities
     */
    async listUserRaw(requestParameters: ListUserRequest): Promise<runtime.ApiResponse<Array<EntUser>>> {
        const queryParameters: runtime.HTTPQuery = {};

        if (requestParameters.limit !== undefined) {
            queryParameters['limit'] = requestParameters.limit;
        }

        if (requestParameters.offset !== undefined) {
            queryParameters['offset'] = requestParameters.offset;
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/users`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => jsonValue.map(EntUserFromJSON));
    }

    /**
     * list user entities
     * List user entities
     */
    async listUser(requestParameters: ListUserRequest): Promise<Array<EntUser>> {
        const response = await this.listUserRaw(requestParameters);
        return await response.value();
    }

    /**
     * update company by ID
     * Update a company entity by ID
     */
    async updateCompanyRaw(requestParameters: UpdateCompanyRequest): Promise<runtime.ApiResponse<EntCompany>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling updateCompany.');
        }

        if (requestParameters.company === null || requestParameters.company === undefined) {
            throw new runtime.RequiredError('company','Required parameter requestParameters.company was null or undefined when calling updateCompany.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/companys/{id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'PUT',
            headers: headerParameters,
            query: queryParameters,
            body: EntCompanyToJSON(requestParameters.company),
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => EntCompanyFromJSON(jsonValue));
    }

    /**
     * update company by ID
     * Update a company entity by ID
     */
    async updateCompany(requestParameters: UpdateCompanyRequest): Promise<EntCompany> {
        const response = await this.updateCompanyRaw(requestParameters);
        return await response.value();
    }

    /**
     * update equipment by ID
     * Update a equipment entity by ID
     */
    async updateEquipmentRaw(requestParameters: UpdateEquipmentRequest): Promise<runtime.ApiResponse<EntEquipment>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling updateEquipment.');
        }

        if (requestParameters.equipment === null || requestParameters.equipment === undefined) {
            throw new runtime.RequiredError('equipment','Required parameter requestParameters.equipment was null or undefined when calling updateEquipment.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/equipments/{id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'PUT',
            headers: headerParameters,
            query: queryParameters,
            body: EntEquipmentToJSON(requestParameters.equipment),
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => EntEquipmentFromJSON(jsonValue));
    }

    /**
     * update equipment by ID
     * Update a equipment entity by ID
     */
    async updateEquipment(requestParameters: UpdateEquipmentRequest): Promise<EntEquipment> {
        const response = await this.updateEquipmentRaw(requestParameters);
        return await response.value();
    }

    /**
     * update order by ID
     * Update a order entity by ID
     */
    async updateOrderRaw(requestParameters: UpdateOrderRequest): Promise<runtime.ApiResponse<EntOrder>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling updateOrder.');
        }

        if (requestParameters.order === null || requestParameters.order === undefined) {
            throw new runtime.RequiredError('order','Required parameter requestParameters.order was null or undefined when calling updateOrder.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/orders/{id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'PUT',
            headers: headerParameters,
            query: queryParameters,
            body: EntOrderToJSON(requestParameters.order),
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => EntOrderFromJSON(jsonValue));
    }

    /**
     * update order by ID
     * Update a order entity by ID
     */
    async updateOrder(requestParameters: UpdateOrderRequest): Promise<EntOrder> {
        const response = await this.updateOrderRaw(requestParameters);
        return await response.value();
    }

    /**
     * update user by ID
     * Update a user entity by ID
     */
    async updateUserRaw(requestParameters: UpdateUserRequest): Promise<runtime.ApiResponse<EntUser>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling updateUser.');
        }

        if (requestParameters.user === null || requestParameters.user === undefined) {
            throw new runtime.RequiredError('user','Required parameter requestParameters.user was null or undefined when calling updateUser.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/users/{id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'PUT',
            headers: headerParameters,
            query: queryParameters,
            body: EntUserToJSON(requestParameters.user),
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => EntUserFromJSON(jsonValue));
    }

    /**
     * update user by ID
     * Update a user entity by ID
     */
    async updateUser(requestParameters: UpdateUserRequest): Promise<EntUser> {
        const response = await this.updateUserRaw(requestParameters);
        return await response.value();
    }

}