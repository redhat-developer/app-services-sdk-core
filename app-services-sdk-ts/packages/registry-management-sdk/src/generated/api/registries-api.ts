/* tslint:disable */
/* eslint-disable */
/**
 * Service Registry Management API
 * Service Registry Management API is a REST API for managing Service Registry instances. Service Registry is a datastore for event schemas and API designs, which is based on the open source Apicurio Registry project.
 *
 * The version of the OpenAPI document: 1.0.0
 * Contact: rhosak-eval-support@redhat.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


import globalAxios, { AxiosPromise, AxiosInstance, AxiosRequestConfig } from 'axios';
import { Configuration } from '../configuration';
// Some imports not used depending on template conditions
// @ts-ignore
import { DUMMY_BASE_URL, assertParamExists, setApiKeyToObject, setBasicAuthToObject, setBearerAuthToObject, setOAuthToObject, setSearchParams, serializeDataIfNeeded, toPathString, createRequestFunction } from '../common';
// @ts-ignore
import { BASE_PATH, COLLECTION_FORMATS, RequestArgs, BaseAPI, RequiredError } from '../base';
// @ts-ignore
import { Empty } from '../model';
// @ts-ignore
import { Registry } from '../model';
// @ts-ignore
import { RegistryCreate } from '../model';
// @ts-ignore
import { RegistryList } from '../model';
/**
 * RegistriesApi - axios parameter creator
 * @export
 */
export const RegistriesApiAxiosParamCreator = function (configuration?: Configuration) {
    return {
        /**
         * Create a new Registry instance
         * @param {RegistryCreate} registryCreate A new &#x60;Registry&#x60; instance to be created.
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        createRegistry: async (registryCreate: RegistryCreate, options: AxiosRequestConfig = {}): Promise<RequestArgs> => {
            // verify required parameter 'registryCreate' is not null or undefined
            assertParamExists('createRegistry', 'registryCreate', registryCreate)
            const localVarPath = `/api/serviceregistry_mgmt/v1/registries`;
            // use dummy base URL string because the URL constructor only accepts absolute URLs.
            const localVarUrlObj = new URL(localVarPath, DUMMY_BASE_URL);
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }

            const localVarRequestOptions = { method: 'POST', ...baseOptions, ...options};
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;

            // authentication Bearer required
            // http bearer authentication required
            await setBearerAuthToObject(localVarHeaderParameter, configuration)


    
            localVarHeaderParameter['Content-Type'] = 'application/json';

            setSearchParams(localVarUrlObj, localVarQueryParameter);
            let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
            localVarRequestOptions.headers = {...localVarHeaderParameter, ...headersFromBaseOptions, ...options.headers};
            localVarRequestOptions.data = serializeDataIfNeeded(registryCreate, localVarRequestOptions, configuration)

            return {
                url: toPathString(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
        /**
         * Deletes an existing `Registry` instance and all of the data that it stores. Important: Users should export the registry data before deleting the instance, e.g., using the Service Registry web console, core REST API, or `rhoas` CLI.
         * @summary Delete a Registry instance
         * @param {string} id The id of the object you wish to interact with.
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        deleteRegistry: async (id: string, options: AxiosRequestConfig = {}): Promise<RequestArgs> => {
            // verify required parameter 'id' is not null or undefined
            assertParamExists('deleteRegistry', 'id', id)
            const localVarPath = `/api/serviceregistry_mgmt/v1/registries/{id}`
                .replace(`{${"id"}}`, encodeURIComponent(String(id)));
            // use dummy base URL string because the URL constructor only accepts absolute URLs.
            const localVarUrlObj = new URL(localVarPath, DUMMY_BASE_URL);
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }

            const localVarRequestOptions = { method: 'DELETE', ...baseOptions, ...options};
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;

            // authentication Bearer required
            // http bearer authentication required
            await setBearerAuthToObject(localVarHeaderParameter, configuration)


    
            setSearchParams(localVarUrlObj, localVarQueryParameter);
            let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
            localVarRequestOptions.headers = {...localVarHeaderParameter, ...headersFromBaseOptions, ...options.headers};

            return {
                url: toPathString(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
        /**
         * Get the list of all Registry instances
         * @param {number} [page] Page index.
         * @param {number} [size] Number of items in each page.
         * @param {string} [orderBy] Specifies the order by criteria. The syntax of this parameter is similar to the syntax of the _order by_ clause of an SQL statement. Each query can be ordered by any of the request fields. For example, to retrieve all Registry instances ordered by their name:  &#x60;&#x60;&#x60;sql name asc &#x60;&#x60;&#x60;  Or to retrieve all Registry instances ordered by their name _and_ created date:  &#x60;&#x60;&#x60;sql name asc, created_at asc &#x60;&#x60;&#x60;  If the parameter isn\&#39;t provided, or if the value is empty,  the results are ordered by name.
         * @param {string} [search] Search criteria.  The syntax of this parameter is similar to the syntax of the _where_ clause of an SQL statement. Allowed fields in the search are: &#x60;name&#x60;, &#x60;status&#x60;. Allowed comparators are &#x60;&#x3D;&#x60; or &#x60;LIKE&#x60;. Allowed joins are &#x60;AND&#x60; and &#x60;OR&#x60;, however there is a limit of max 10 joins in the search query.  Examples:  To retrieve a request with name equal &#x60;my-registry&#x60;, the value should be:  &#x60;&#x60;&#x60; name &#x3D; my-registry  &#x60;&#x60;&#x60;  To retrieve a request with its name starting with &#x60;my&#x60;, the value should be:  &#x60;&#x60;&#x60; name like my%25 &#x60;&#x60;&#x60;  If the parameter isn\&#39;t provided, or if the value is empty, all the Registry instances that the user has permission to see are returned.  Note: If the query is invalid, an error is returned. 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getRegistries: async (page?: number, size?: number, orderBy?: string, search?: string, options: AxiosRequestConfig = {}): Promise<RequestArgs> => {
            const localVarPath = `/api/serviceregistry_mgmt/v1/registries`;
            // use dummy base URL string because the URL constructor only accepts absolute URLs.
            const localVarUrlObj = new URL(localVarPath, DUMMY_BASE_URL);
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }

            const localVarRequestOptions = { method: 'GET', ...baseOptions, ...options};
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;

            // authentication Bearer required
            // http bearer authentication required
            await setBearerAuthToObject(localVarHeaderParameter, configuration)

            if (page !== undefined) {
                localVarQueryParameter['page'] = page;
            }

            if (size !== undefined) {
                localVarQueryParameter['size'] = size;
            }

            if (orderBy !== undefined) {
                localVarQueryParameter['orderBy'] = orderBy;
            }

            if (search !== undefined) {
                localVarQueryParameter['search'] = search;
            }


    
            setSearchParams(localVarUrlObj, localVarQueryParameter);
            let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
            localVarRequestOptions.headers = {...localVarHeaderParameter, ...headersFromBaseOptions, ...options.headers};

            return {
                url: toPathString(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
        /**
         * Gets the details of a single instance of a `Registry`.
         * @summary Get a Registry instance
         * @param {string} id The id of the object you wish to interact with.
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getRegistry: async (id: string, options: AxiosRequestConfig = {}): Promise<RequestArgs> => {
            // verify required parameter 'id' is not null or undefined
            assertParamExists('getRegistry', 'id', id)
            const localVarPath = `/api/serviceregistry_mgmt/v1/registries/{id}`
                .replace(`{${"id"}}`, encodeURIComponent(String(id)));
            // use dummy base URL string because the URL constructor only accepts absolute URLs.
            const localVarUrlObj = new URL(localVarPath, DUMMY_BASE_URL);
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }

            const localVarRequestOptions = { method: 'GET', ...baseOptions, ...options};
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;

            // authentication Bearer required
            // http bearer authentication required
            await setBearerAuthToObject(localVarHeaderParameter, configuration)


    
            setSearchParams(localVarUrlObj, localVarQueryParameter);
            let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
            localVarRequestOptions.headers = {...localVarHeaderParameter, ...headersFromBaseOptions, ...options.headers};

            return {
                url: toPathString(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
    }
};

/**
 * RegistriesApi - functional programming interface
 * @export
 */
export const RegistriesApiFp = function(configuration?: Configuration) {
    const localVarAxiosParamCreator = RegistriesApiAxiosParamCreator(configuration)
    return {
        /**
         * Create a new Registry instance
         * @param {RegistryCreate} registryCreate A new &#x60;Registry&#x60; instance to be created.
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async createRegistry(registryCreate: RegistryCreate, options?: AxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Registry>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.createRegistry(registryCreate, options);
            return createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration);
        },
        /**
         * Deletes an existing `Registry` instance and all of the data that it stores. Important: Users should export the registry data before deleting the instance, e.g., using the Service Registry web console, core REST API, or `rhoas` CLI.
         * @summary Delete a Registry instance
         * @param {string} id The id of the object you wish to interact with.
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async deleteRegistry(id: string, options?: AxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<void>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.deleteRegistry(id, options);
            return createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration);
        },
        /**
         * Get the list of all Registry instances
         * @param {number} [page] Page index.
         * @param {number} [size] Number of items in each page.
         * @param {string} [orderBy] Specifies the order by criteria. The syntax of this parameter is similar to the syntax of the _order by_ clause of an SQL statement. Each query can be ordered by any of the request fields. For example, to retrieve all Registry instances ordered by their name:  &#x60;&#x60;&#x60;sql name asc &#x60;&#x60;&#x60;  Or to retrieve all Registry instances ordered by their name _and_ created date:  &#x60;&#x60;&#x60;sql name asc, created_at asc &#x60;&#x60;&#x60;  If the parameter isn\&#39;t provided, or if the value is empty,  the results are ordered by name.
         * @param {string} [search] Search criteria.  The syntax of this parameter is similar to the syntax of the _where_ clause of an SQL statement. Allowed fields in the search are: &#x60;name&#x60;, &#x60;status&#x60;. Allowed comparators are &#x60;&#x3D;&#x60; or &#x60;LIKE&#x60;. Allowed joins are &#x60;AND&#x60; and &#x60;OR&#x60;, however there is a limit of max 10 joins in the search query.  Examples:  To retrieve a request with name equal &#x60;my-registry&#x60;, the value should be:  &#x60;&#x60;&#x60; name &#x3D; my-registry  &#x60;&#x60;&#x60;  To retrieve a request with its name starting with &#x60;my&#x60;, the value should be:  &#x60;&#x60;&#x60; name like my%25 &#x60;&#x60;&#x60;  If the parameter isn\&#39;t provided, or if the value is empty, all the Registry instances that the user has permission to see are returned.  Note: If the query is invalid, an error is returned. 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async getRegistries(page?: number, size?: number, orderBy?: string, search?: string, options?: AxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<RegistryList>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.getRegistries(page, size, orderBy, search, options);
            return createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration);
        },
        /**
         * Gets the details of a single instance of a `Registry`.
         * @summary Get a Registry instance
         * @param {string} id The id of the object you wish to interact with.
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async getRegistry(id: string, options?: AxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Registry>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.getRegistry(id, options);
            return createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration);
        },
    }
};

/**
 * RegistriesApi - factory interface
 * @export
 */
export const RegistriesApiFactory = function (configuration?: Configuration, basePath?: string, axios?: AxiosInstance) {
    const localVarFp = RegistriesApiFp(configuration)
    return {
        /**
         * Create a new Registry instance
         * @param {RegistryCreate} registryCreate A new &#x60;Registry&#x60; instance to be created.
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        createRegistry(registryCreate: RegistryCreate, options?: any): AxiosPromise<Registry> {
            return localVarFp.createRegistry(registryCreate, options).then((request) => request(axios, basePath));
        },
        /**
         * Deletes an existing `Registry` instance and all of the data that it stores. Important: Users should export the registry data before deleting the instance, e.g., using the Service Registry web console, core REST API, or `rhoas` CLI.
         * @summary Delete a Registry instance
         * @param {string} id The id of the object you wish to interact with.
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        deleteRegistry(id: string, options?: any): AxiosPromise<void> {
            return localVarFp.deleteRegistry(id, options).then((request) => request(axios, basePath));
        },
        /**
         * Get the list of all Registry instances
         * @param {number} [page] Page index.
         * @param {number} [size] Number of items in each page.
         * @param {string} [orderBy] Specifies the order by criteria. The syntax of this parameter is similar to the syntax of the _order by_ clause of an SQL statement. Each query can be ordered by any of the request fields. For example, to retrieve all Registry instances ordered by their name:  &#x60;&#x60;&#x60;sql name asc &#x60;&#x60;&#x60;  Or to retrieve all Registry instances ordered by their name _and_ created date:  &#x60;&#x60;&#x60;sql name asc, created_at asc &#x60;&#x60;&#x60;  If the parameter isn\&#39;t provided, or if the value is empty,  the results are ordered by name.
         * @param {string} [search] Search criteria.  The syntax of this parameter is similar to the syntax of the _where_ clause of an SQL statement. Allowed fields in the search are: &#x60;name&#x60;, &#x60;status&#x60;. Allowed comparators are &#x60;&#x3D;&#x60; or &#x60;LIKE&#x60;. Allowed joins are &#x60;AND&#x60; and &#x60;OR&#x60;, however there is a limit of max 10 joins in the search query.  Examples:  To retrieve a request with name equal &#x60;my-registry&#x60;, the value should be:  &#x60;&#x60;&#x60; name &#x3D; my-registry  &#x60;&#x60;&#x60;  To retrieve a request with its name starting with &#x60;my&#x60;, the value should be:  &#x60;&#x60;&#x60; name like my%25 &#x60;&#x60;&#x60;  If the parameter isn\&#39;t provided, or if the value is empty, all the Registry instances that the user has permission to see are returned.  Note: If the query is invalid, an error is returned. 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getRegistries(page?: number, size?: number, orderBy?: string, search?: string, options?: any): AxiosPromise<RegistryList> {
            return localVarFp.getRegistries(page, size, orderBy, search, options).then((request) => request(axios, basePath));
        },
        /**
         * Gets the details of a single instance of a `Registry`.
         * @summary Get a Registry instance
         * @param {string} id The id of the object you wish to interact with.
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getRegistry(id: string, options?: any): AxiosPromise<Registry> {
            return localVarFp.getRegistry(id, options).then((request) => request(axios, basePath));
        },
    };
};

/**
 * RegistriesApi - interface
 * @export
 * @interface RegistriesApi
 */
export interface RegistriesApiInterface {
    /**
     * Create a new Registry instance
     * @param {RegistryCreate} registryCreate A new &#x60;Registry&#x60; instance to be created.
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof RegistriesApiInterface
     */
    createRegistry(registryCreate: RegistryCreate, options?: AxiosRequestConfig): AxiosPromise<Registry>;

    /**
     * Deletes an existing `Registry` instance and all of the data that it stores. Important: Users should export the registry data before deleting the instance, e.g., using the Service Registry web console, core REST API, or `rhoas` CLI.
     * @summary Delete a Registry instance
     * @param {string} id The id of the object you wish to interact with.
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof RegistriesApiInterface
     */
    deleteRegistry(id: string, options?: AxiosRequestConfig): AxiosPromise<void>;

    /**
     * Get the list of all Registry instances
     * @param {number} [page] Page index.
     * @param {number} [size] Number of items in each page.
     * @param {string} [orderBy] Specifies the order by criteria. The syntax of this parameter is similar to the syntax of the _order by_ clause of an SQL statement. Each query can be ordered by any of the request fields. For example, to retrieve all Registry instances ordered by their name:  &#x60;&#x60;&#x60;sql name asc &#x60;&#x60;&#x60;  Or to retrieve all Registry instances ordered by their name _and_ created date:  &#x60;&#x60;&#x60;sql name asc, created_at asc &#x60;&#x60;&#x60;  If the parameter isn\&#39;t provided, or if the value is empty,  the results are ordered by name.
     * @param {string} [search] Search criteria.  The syntax of this parameter is similar to the syntax of the _where_ clause of an SQL statement. Allowed fields in the search are: &#x60;name&#x60;, &#x60;status&#x60;. Allowed comparators are &#x60;&#x3D;&#x60; or &#x60;LIKE&#x60;. Allowed joins are &#x60;AND&#x60; and &#x60;OR&#x60;, however there is a limit of max 10 joins in the search query.  Examples:  To retrieve a request with name equal &#x60;my-registry&#x60;, the value should be:  &#x60;&#x60;&#x60; name &#x3D; my-registry  &#x60;&#x60;&#x60;  To retrieve a request with its name starting with &#x60;my&#x60;, the value should be:  &#x60;&#x60;&#x60; name like my%25 &#x60;&#x60;&#x60;  If the parameter isn\&#39;t provided, or if the value is empty, all the Registry instances that the user has permission to see are returned.  Note: If the query is invalid, an error is returned. 
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof RegistriesApiInterface
     */
    getRegistries(page?: number, size?: number, orderBy?: string, search?: string, options?: AxiosRequestConfig): AxiosPromise<RegistryList>;

    /**
     * Gets the details of a single instance of a `Registry`.
     * @summary Get a Registry instance
     * @param {string} id The id of the object you wish to interact with.
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof RegistriesApiInterface
     */
    getRegistry(id: string, options?: AxiosRequestConfig): AxiosPromise<Registry>;

}

/**
 * RegistriesApi - object-oriented interface
 * @export
 * @class RegistriesApi
 * @extends {BaseAPI}
 */
export class RegistriesApi extends BaseAPI implements RegistriesApiInterface {
    /**
     * Create a new Registry instance
     * @param {RegistryCreate} registryCreate A new &#x60;Registry&#x60; instance to be created.
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof RegistriesApi
     */
    public createRegistry(registryCreate: RegistryCreate, options?: AxiosRequestConfig) {
        return RegistriesApiFp(this.configuration).createRegistry(registryCreate, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * Deletes an existing `Registry` instance and all of the data that it stores. Important: Users should export the registry data before deleting the instance, e.g., using the Service Registry web console, core REST API, or `rhoas` CLI.
     * @summary Delete a Registry instance
     * @param {string} id The id of the object you wish to interact with.
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof RegistriesApi
     */
    public deleteRegistry(id: string, options?: AxiosRequestConfig) {
        return RegistriesApiFp(this.configuration).deleteRegistry(id, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * Get the list of all Registry instances
     * @param {number} [page] Page index.
     * @param {number} [size] Number of items in each page.
     * @param {string} [orderBy] Specifies the order by criteria. The syntax of this parameter is similar to the syntax of the _order by_ clause of an SQL statement. Each query can be ordered by any of the request fields. For example, to retrieve all Registry instances ordered by their name:  &#x60;&#x60;&#x60;sql name asc &#x60;&#x60;&#x60;  Or to retrieve all Registry instances ordered by their name _and_ created date:  &#x60;&#x60;&#x60;sql name asc, created_at asc &#x60;&#x60;&#x60;  If the parameter isn\&#39;t provided, or if the value is empty,  the results are ordered by name.
     * @param {string} [search] Search criteria.  The syntax of this parameter is similar to the syntax of the _where_ clause of an SQL statement. Allowed fields in the search are: &#x60;name&#x60;, &#x60;status&#x60;. Allowed comparators are &#x60;&#x3D;&#x60; or &#x60;LIKE&#x60;. Allowed joins are &#x60;AND&#x60; and &#x60;OR&#x60;, however there is a limit of max 10 joins in the search query.  Examples:  To retrieve a request with name equal &#x60;my-registry&#x60;, the value should be:  &#x60;&#x60;&#x60; name &#x3D; my-registry  &#x60;&#x60;&#x60;  To retrieve a request with its name starting with &#x60;my&#x60;, the value should be:  &#x60;&#x60;&#x60; name like my%25 &#x60;&#x60;&#x60;  If the parameter isn\&#39;t provided, or if the value is empty, all the Registry instances that the user has permission to see are returned.  Note: If the query is invalid, an error is returned. 
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof RegistriesApi
     */
    public getRegistries(page?: number, size?: number, orderBy?: string, search?: string, options?: AxiosRequestConfig) {
        return RegistriesApiFp(this.configuration).getRegistries(page, size, orderBy, search, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * Gets the details of a single instance of a `Registry`.
     * @summary Get a Registry instance
     * @param {string} id The id of the object you wish to interact with.
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof RegistriesApi
     */
    public getRegistry(id: string, options?: AxiosRequestConfig) {
        return RegistriesApiFp(this.configuration).getRegistry(id, options).then((request) => request(this.axios, this.basePath));
    }
}