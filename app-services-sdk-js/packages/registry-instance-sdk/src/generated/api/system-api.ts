/* tslint:disable */
/* eslint-disable */
/**
 * Service Registry API
 * Service Registry Instance API  NOTE: This API cannot be called directly from the portal.
 *
 * The version of the OpenAPI document: 2.2.5.Final
 * Contact: apicurio@lists.jboss.org
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
import { Limits } from '../model';
// @ts-ignore
import { SystemInfo } from '../model';
/**
 * SystemApi - axios parameter creator
 * @export
 */
export const SystemApiAxiosParamCreator = function (configuration?: Configuration) {
    return {
        /**
         * This operation retrieves the list of limitations on used resources, that are applied on the current instance of Registry.
         * @summary Get resource limits information
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getResourceLimits: async (options: AxiosRequestConfig = {}): Promise<RequestArgs> => {
            const localVarPath = `/system/limits`;
            // use dummy base URL string because the URL constructor only accepts absolute URLs.
            const localVarUrlObj = new URL(localVarPath, DUMMY_BASE_URL);
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }

            const localVarRequestOptions = { method: 'GET', ...baseOptions, ...options};
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;


    
            setSearchParams(localVarUrlObj, localVarQueryParameter);
            let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
            localVarRequestOptions.headers = {...localVarHeaderParameter, ...headersFromBaseOptions, ...options.headers};

            return {
                url: toPathString(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
        /**
         * This operation retrieves information about the running registry system, such as the version of the software and when it was built.
         * @summary Get system information
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getSystemInfo: async (options: AxiosRequestConfig = {}): Promise<RequestArgs> => {
            const localVarPath = `/system/info`;
            // use dummy base URL string because the URL constructor only accepts absolute URLs.
            const localVarUrlObj = new URL(localVarPath, DUMMY_BASE_URL);
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }

            const localVarRequestOptions = { method: 'GET', ...baseOptions, ...options};
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;


    
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
 * SystemApi - functional programming interface
 * @export
 */
export const SystemApiFp = function(configuration?: Configuration) {
    const localVarAxiosParamCreator = SystemApiAxiosParamCreator(configuration)
    return {
        /**
         * This operation retrieves the list of limitations on used resources, that are applied on the current instance of Registry.
         * @summary Get resource limits information
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async getResourceLimits(options?: AxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Limits>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.getResourceLimits(options);
            return createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration);
        },
        /**
         * This operation retrieves information about the running registry system, such as the version of the software and when it was built.
         * @summary Get system information
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async getSystemInfo(options?: AxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<SystemInfo>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.getSystemInfo(options);
            return createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration);
        },
    }
};

/**
 * SystemApi - factory interface
 * @export
 */
export const SystemApiFactory = function (configuration?: Configuration, basePath?: string, axios?: AxiosInstance) {
    const localVarFp = SystemApiFp(configuration)
    return {
        /**
         * This operation retrieves the list of limitations on used resources, that are applied on the current instance of Registry.
         * @summary Get resource limits information
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getResourceLimits(options?: any): AxiosPromise<Limits> {
            return localVarFp.getResourceLimits(options).then((request) => request(axios, basePath));
        },
        /**
         * This operation retrieves information about the running registry system, such as the version of the software and when it was built.
         * @summary Get system information
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getSystemInfo(options?: any): AxiosPromise<SystemInfo> {
            return localVarFp.getSystemInfo(options).then((request) => request(axios, basePath));
        },
    };
};

/**
 * SystemApi - interface
 * @export
 * @interface SystemApi
 */
export interface SystemApiInterface {
    /**
     * This operation retrieves the list of limitations on used resources, that are applied on the current instance of Registry.
     * @summary Get resource limits information
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof SystemApiInterface
     */
    getResourceLimits(options?: AxiosRequestConfig): AxiosPromise<Limits>;

    /**
     * This operation retrieves information about the running registry system, such as the version of the software and when it was built.
     * @summary Get system information
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof SystemApiInterface
     */
    getSystemInfo(options?: AxiosRequestConfig): AxiosPromise<SystemInfo>;

}

/**
 * SystemApi - object-oriented interface
 * @export
 * @class SystemApi
 * @extends {BaseAPI}
 */
export class SystemApi extends BaseAPI implements SystemApiInterface {
    /**
     * This operation retrieves the list of limitations on used resources, that are applied on the current instance of Registry.
     * @summary Get resource limits information
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof SystemApi
     */
    public getResourceLimits(options?: AxiosRequestConfig) {
        return SystemApiFp(this.configuration).getResourceLimits(options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * This operation retrieves information about the running registry system, such as the version of the software and when it was built.
     * @summary Get system information
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof SystemApi
     */
    public getSystemInfo(options?: AxiosRequestConfig) {
        return SystemApiFp(this.configuration).getSystemInfo(options).then((request) => request(this.axios, this.basePath));
    }
}