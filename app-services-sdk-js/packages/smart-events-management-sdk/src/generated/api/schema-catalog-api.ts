/* tslint:disable */
/* eslint-disable */
/**
 * Red Hat Openshift SmartEvents Fleet Manager V1
 * The API exposed by the fleet manager of the SmartEvents service.
 *
 * The version of the OpenAPI document: 0.0.1
 * Contact: openbridge-dev@redhat.com
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
import { ErrorsList } from '../model';
// @ts-ignore
import { ProcessorCatalogResponse } from '../model';
/**
 * SchemaCatalogApi - axios parameter creator
 * @export
 */
export const SchemaCatalogApiAxiosParamCreator = function (configuration?: Configuration) {
    return {
        /**
         * Get the action processor JSON schema.
         * @summary Get action processor schema
         * @param {string} id 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getActionProcessorSchema: async (id: string, options: AxiosRequestConfig = {}): Promise<RequestArgs> => {
            // verify required parameter 'id' is not null or undefined
            assertParamExists('getActionProcessorSchema', 'id', id)
            const localVarPath = `/api/smartevents_mgmt/v1/schemas/actions/{id}`
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

            // authentication bearer required
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
         * Get the processor catalog with all the available sources and actions.
         * @summary Get processor catalog
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getCatalog: async (options: AxiosRequestConfig = {}): Promise<RequestArgs> => {
            const localVarPath = `/api/smartevents_mgmt/v1/schemas`;
            // use dummy base URL string because the URL constructor only accepts absolute URLs.
            const localVarUrlObj = new URL(localVarPath, DUMMY_BASE_URL);
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }

            const localVarRequestOptions = { method: 'GET', ...baseOptions, ...options};
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;

            // authentication bearer required
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
         * Get the source processor JSON schema.
         * @summary Get source processor schema
         * @param {string} id 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getSourceProcessorSchema: async (id: string, options: AxiosRequestConfig = {}): Promise<RequestArgs> => {
            // verify required parameter 'id' is not null or undefined
            assertParamExists('getSourceProcessorSchema', 'id', id)
            const localVarPath = `/api/smartevents_mgmt/v1/schemas/sources/{id}`
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

            // authentication bearer required
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
 * SchemaCatalogApi - functional programming interface
 * @export
 */
export const SchemaCatalogApiFp = function(configuration?: Configuration) {
    const localVarAxiosParamCreator = SchemaCatalogApiAxiosParamCreator(configuration)
    return {
        /**
         * Get the action processor JSON schema.
         * @summary Get action processor schema
         * @param {string} id 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async getActionProcessorSchema(id: string, options?: AxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<object>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.getActionProcessorSchema(id, options);
            return createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration);
        },
        /**
         * Get the processor catalog with all the available sources and actions.
         * @summary Get processor catalog
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async getCatalog(options?: AxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<ProcessorCatalogResponse>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.getCatalog(options);
            return createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration);
        },
        /**
         * Get the source processor JSON schema.
         * @summary Get source processor schema
         * @param {string} id 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async getSourceProcessorSchema(id: string, options?: AxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<object>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.getSourceProcessorSchema(id, options);
            return createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration);
        },
    }
};

/**
 * SchemaCatalogApi - factory interface
 * @export
 */
export const SchemaCatalogApiFactory = function (configuration?: Configuration, basePath?: string, axios?: AxiosInstance) {
    const localVarFp = SchemaCatalogApiFp(configuration)
    return {
        /**
         * Get the action processor JSON schema.
         * @summary Get action processor schema
         * @param {string} id 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getActionProcessorSchema(id: string, options?: any): AxiosPromise<object> {
            return localVarFp.getActionProcessorSchema(id, options).then((request) => request(axios, basePath));
        },
        /**
         * Get the processor catalog with all the available sources and actions.
         * @summary Get processor catalog
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getCatalog(options?: any): AxiosPromise<ProcessorCatalogResponse> {
            return localVarFp.getCatalog(options).then((request) => request(axios, basePath));
        },
        /**
         * Get the source processor JSON schema.
         * @summary Get source processor schema
         * @param {string} id 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getSourceProcessorSchema(id: string, options?: any): AxiosPromise<object> {
            return localVarFp.getSourceProcessorSchema(id, options).then((request) => request(axios, basePath));
        },
    };
};

/**
 * SchemaCatalogApi - interface
 * @export
 * @interface SchemaCatalogApi
 */
export interface SchemaCatalogApiInterface {
    /**
     * Get the action processor JSON schema.
     * @summary Get action processor schema
     * @param {string} id 
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof SchemaCatalogApiInterface
     */
    getActionProcessorSchema(id: string, options?: AxiosRequestConfig): AxiosPromise<object>;

    /**
     * Get the processor catalog with all the available sources and actions.
     * @summary Get processor catalog
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof SchemaCatalogApiInterface
     */
    getCatalog(options?: AxiosRequestConfig): AxiosPromise<ProcessorCatalogResponse>;

    /**
     * Get the source processor JSON schema.
     * @summary Get source processor schema
     * @param {string} id 
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof SchemaCatalogApiInterface
     */
    getSourceProcessorSchema(id: string, options?: AxiosRequestConfig): AxiosPromise<object>;

}

/**
 * SchemaCatalogApi - object-oriented interface
 * @export
 * @class SchemaCatalogApi
 * @extends {BaseAPI}
 */
export class SchemaCatalogApi extends BaseAPI implements SchemaCatalogApiInterface {
    /**
     * Get the action processor JSON schema.
     * @summary Get action processor schema
     * @param {string} id 
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof SchemaCatalogApi
     */
    public getActionProcessorSchema(id: string, options?: AxiosRequestConfig) {
        return SchemaCatalogApiFp(this.configuration).getActionProcessorSchema(id, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * Get the processor catalog with all the available sources and actions.
     * @summary Get processor catalog
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof SchemaCatalogApi
     */
    public getCatalog(options?: AxiosRequestConfig) {
        return SchemaCatalogApiFp(this.configuration).getCatalog(options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * Get the source processor JSON schema.
     * @summary Get source processor schema
     * @param {string} id 
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof SchemaCatalogApi
     */
    public getSourceProcessorSchema(id: string, options?: AxiosRequestConfig) {
        return SchemaCatalogApiFp(this.configuration).getSourceProcessorSchema(id, options).then((request) => request(this.axios, this.basePath));
    }
}
