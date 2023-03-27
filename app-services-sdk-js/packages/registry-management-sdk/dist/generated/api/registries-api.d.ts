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
import { AxiosPromise, AxiosInstance, AxiosRequestConfig } from 'axios';
import { Configuration } from '../configuration';
import { RequestArgs, BaseAPI } from '../base';
import { Registry } from '../model';
import { RegistryCreate } from '../model';
import { RegistryList } from '../model';
/**
 * RegistriesApi - axios parameter creator
 * @export
 */
export declare const RegistriesApiAxiosParamCreator: (configuration?: Configuration) => {
    /**
     * Create a new Registry instance
     * @param {RegistryCreate} registryCreate A new &#x60;Registry&#x60; instance to be created.
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    createRegistry: (registryCreate: RegistryCreate, options?: AxiosRequestConfig) => Promise<RequestArgs>;
    /**
     * Deletes an existing `Registry` instance and all of the data that it stores. Important: Users should export the registry data before deleting the instance, e.g., using the Service Registry web console, core REST API, or `rhoas` CLI.
     * @summary Delete a Registry instance
     * @param {string} id The id of the object you wish to interact with.
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    deleteRegistry: (id: string, options?: AxiosRequestConfig) => Promise<RequestArgs>;
    /**
     * Get the list of all Registry instances
     * @param {number} [page] Page index.
     * @param {number} [size] Number of items in each page.
     * @param {string} [orderBy] Specifies the order by criteria. The syntax of this parameter is similar to the syntax of the _order by_ clause of an SQL statement. Each query can be ordered by any of the request fields. For example, to retrieve all Registry instances ordered by their name:  &#x60;&#x60;&#x60;sql name asc &#x60;&#x60;&#x60;  Or to retrieve all Registry instances ordered by their name _and_ created date:  &#x60;&#x60;&#x60;sql name asc, created_at asc &#x60;&#x60;&#x60;  If the parameter isn\&#39;t provided, or if the value is empty,  the results are ordered by name.
     * @param {string} [search] Search criteria.  The syntax of this parameter is similar to the syntax of the _where_ clause of an SQL statement. Allowed fields in the search are: &#x60;name&#x60;, &#x60;status&#x60;. Allowed comparators are &#x60;&#x3D;&#x60; or &#x60;LIKE&#x60;. Allowed joins are &#x60;AND&#x60; and &#x60;OR&#x60;, however there is a limit of max 10 joins in the search query.  Examples:  To retrieve a request with name equal &#x60;my-registry&#x60;, the value should be:  &#x60;&#x60;&#x60; name &#x3D; my-registry  &#x60;&#x60;&#x60;  To retrieve a request with its name starting with &#x60;my&#x60;, the value should be:  &#x60;&#x60;&#x60; name like my%25 &#x60;&#x60;&#x60;  If the parameter isn\&#39;t provided, or if the value is empty, all the Registry instances that the user has permission to see are returned.  Note: If the query is invalid, an error is returned.
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    getRegistries: (page?: number, size?: number, orderBy?: string, search?: string, options?: AxiosRequestConfig) => Promise<RequestArgs>;
    /**
     * Gets the details of a single instance of a `Registry`.
     * @summary Get a Registry instance
     * @param {string} id The id of the object you wish to interact with.
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    getRegistry: (id: string, options?: AxiosRequestConfig) => Promise<RequestArgs>;
};
/**
 * RegistriesApi - functional programming interface
 * @export
 */
export declare const RegistriesApiFp: (configuration?: Configuration) => {
    /**
     * Create a new Registry instance
     * @param {RegistryCreate} registryCreate A new &#x60;Registry&#x60; instance to be created.
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    createRegistry(registryCreate: RegistryCreate, options?: AxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Registry>>;
    /**
     * Deletes an existing `Registry` instance and all of the data that it stores. Important: Users should export the registry data before deleting the instance, e.g., using the Service Registry web console, core REST API, or `rhoas` CLI.
     * @summary Delete a Registry instance
     * @param {string} id The id of the object you wish to interact with.
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    deleteRegistry(id: string, options?: AxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<void>>;
    /**
     * Get the list of all Registry instances
     * @param {number} [page] Page index.
     * @param {number} [size] Number of items in each page.
     * @param {string} [orderBy] Specifies the order by criteria. The syntax of this parameter is similar to the syntax of the _order by_ clause of an SQL statement. Each query can be ordered by any of the request fields. For example, to retrieve all Registry instances ordered by their name:  &#x60;&#x60;&#x60;sql name asc &#x60;&#x60;&#x60;  Or to retrieve all Registry instances ordered by their name _and_ created date:  &#x60;&#x60;&#x60;sql name asc, created_at asc &#x60;&#x60;&#x60;  If the parameter isn\&#39;t provided, or if the value is empty,  the results are ordered by name.
     * @param {string} [search] Search criteria.  The syntax of this parameter is similar to the syntax of the _where_ clause of an SQL statement. Allowed fields in the search are: &#x60;name&#x60;, &#x60;status&#x60;. Allowed comparators are &#x60;&#x3D;&#x60; or &#x60;LIKE&#x60;. Allowed joins are &#x60;AND&#x60; and &#x60;OR&#x60;, however there is a limit of max 10 joins in the search query.  Examples:  To retrieve a request with name equal &#x60;my-registry&#x60;, the value should be:  &#x60;&#x60;&#x60; name &#x3D; my-registry  &#x60;&#x60;&#x60;  To retrieve a request with its name starting with &#x60;my&#x60;, the value should be:  &#x60;&#x60;&#x60; name like my%25 &#x60;&#x60;&#x60;  If the parameter isn\&#39;t provided, or if the value is empty, all the Registry instances that the user has permission to see are returned.  Note: If the query is invalid, an error is returned.
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    getRegistries(page?: number, size?: number, orderBy?: string, search?: string, options?: AxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<RegistryList>>;
    /**
     * Gets the details of a single instance of a `Registry`.
     * @summary Get a Registry instance
     * @param {string} id The id of the object you wish to interact with.
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    getRegistry(id: string, options?: AxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Registry>>;
};
/**
 * RegistriesApi - factory interface
 * @export
 */
export declare const RegistriesApiFactory: (configuration?: Configuration, basePath?: string, axios?: AxiosInstance) => {
    /**
     * Create a new Registry instance
     * @param {RegistryCreate} registryCreate A new &#x60;Registry&#x60; instance to be created.
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    createRegistry(registryCreate: RegistryCreate, options?: any): AxiosPromise<Registry>;
    /**
     * Deletes an existing `Registry` instance and all of the data that it stores. Important: Users should export the registry data before deleting the instance, e.g., using the Service Registry web console, core REST API, or `rhoas` CLI.
     * @summary Delete a Registry instance
     * @param {string} id The id of the object you wish to interact with.
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    deleteRegistry(id: string, options?: any): AxiosPromise<void>;
    /**
     * Get the list of all Registry instances
     * @param {number} [page] Page index.
     * @param {number} [size] Number of items in each page.
     * @param {string} [orderBy] Specifies the order by criteria. The syntax of this parameter is similar to the syntax of the _order by_ clause of an SQL statement. Each query can be ordered by any of the request fields. For example, to retrieve all Registry instances ordered by their name:  &#x60;&#x60;&#x60;sql name asc &#x60;&#x60;&#x60;  Or to retrieve all Registry instances ordered by their name _and_ created date:  &#x60;&#x60;&#x60;sql name asc, created_at asc &#x60;&#x60;&#x60;  If the parameter isn\&#39;t provided, or if the value is empty,  the results are ordered by name.
     * @param {string} [search] Search criteria.  The syntax of this parameter is similar to the syntax of the _where_ clause of an SQL statement. Allowed fields in the search are: &#x60;name&#x60;, &#x60;status&#x60;. Allowed comparators are &#x60;&#x3D;&#x60; or &#x60;LIKE&#x60;. Allowed joins are &#x60;AND&#x60; and &#x60;OR&#x60;, however there is a limit of max 10 joins in the search query.  Examples:  To retrieve a request with name equal &#x60;my-registry&#x60;, the value should be:  &#x60;&#x60;&#x60; name &#x3D; my-registry  &#x60;&#x60;&#x60;  To retrieve a request with its name starting with &#x60;my&#x60;, the value should be:  &#x60;&#x60;&#x60; name like my%25 &#x60;&#x60;&#x60;  If the parameter isn\&#39;t provided, or if the value is empty, all the Registry instances that the user has permission to see are returned.  Note: If the query is invalid, an error is returned.
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    getRegistries(page?: number, size?: number, orderBy?: string, search?: string, options?: any): AxiosPromise<RegistryList>;
    /**
     * Gets the details of a single instance of a `Registry`.
     * @summary Get a Registry instance
     * @param {string} id The id of the object you wish to interact with.
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    getRegistry(id: string, options?: any): AxiosPromise<Registry>;
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
export declare class RegistriesApi extends BaseAPI implements RegistriesApiInterface {
    /**
     * Create a new Registry instance
     * @param {RegistryCreate} registryCreate A new &#x60;Registry&#x60; instance to be created.
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof RegistriesApi
     */
    createRegistry(registryCreate: RegistryCreate, options?: AxiosRequestConfig): Promise<import("axios").AxiosResponse<Registry, any>>;
    /**
     * Deletes an existing `Registry` instance and all of the data that it stores. Important: Users should export the registry data before deleting the instance, e.g., using the Service Registry web console, core REST API, or `rhoas` CLI.
     * @summary Delete a Registry instance
     * @param {string} id The id of the object you wish to interact with.
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof RegistriesApi
     */
    deleteRegistry(id: string, options?: AxiosRequestConfig): Promise<import("axios").AxiosResponse<void, any>>;
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
    getRegistries(page?: number, size?: number, orderBy?: string, search?: string, options?: AxiosRequestConfig): Promise<import("axios").AxiosResponse<RegistryList, any>>;
    /**
     * Gets the details of a single instance of a `Registry`.
     * @summary Get a Registry instance
     * @param {string} id The id of the object you wish to interact with.
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof RegistriesApi
     */
    getRegistry(id: string, options?: AxiosRequestConfig): Promise<import("axios").AxiosResponse<Registry, any>>;
}
