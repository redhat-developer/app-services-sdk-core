/* tslint:disable */
/* eslint-disable */
/**
 * Red Hat Openshift SmartEvents Fleet Manager V2
 * The API exposed by the fleet manager of the SmartEvents service.
 *
 * The version of the OpenAPI document: 0.0.1
 * Contact: openbridge-dev@redhat.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


import type { Configuration } from '../configuration';
import type { AxiosPromise, AxiosInstance, AxiosRequestConfig } from 'axios';
import globalAxios from 'axios';
// Some imports not used depending on template conditions
// @ts-ignore
import { DUMMY_BASE_URL, assertParamExists, setApiKeyToObject, setBasicAuthToObject, setBearerAuthToObject, setOAuthToObject, setSearchParams, serializeDataIfNeeded, toPathString, createRequestFunction } from '../common';
// @ts-ignore
import { BASE_PATH, COLLECTION_FORMATS, RequestArgs, BaseAPI, RequiredError } from '../base';
// @ts-ignore
import { BridgeError } from '../model';
// @ts-ignore
import { ErrorListResponse } from '../model';
// @ts-ignore
import { ErrorsList } from '../model';
/**
 * ErrorCatalogApi - axios parameter creator
 * @export
 */
export const ErrorCatalogApiAxiosParamCreator = function (configuration?: Configuration) {
    return {
        /**
         * Get an error from the error catalog.
         * @summary Get an error from the error catalog.
         * @param {number} id 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getError: async (id: number, options: AxiosRequestConfig = {}): Promise<RequestArgs> => {
            // verify required parameter 'id' is not null or undefined
            assertParamExists('getError', 'id', id)
            const localVarPath = `/api/smartevents_mgmt/v2/errors/{id}`
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


    
            setSearchParams(localVarUrlObj, localVarQueryParameter);
            let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
            localVarRequestOptions.headers = {...localVarHeaderParameter, ...headersFromBaseOptions, ...options.headers};

            return {
                url: toPathString(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
        /**
         * Get the list of errors from the error catalog.
         * @summary Get the list of errors.
         * @param {number} [page] 
         * @param {number} [size] 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getErrors: async (page?: number, size?: number, options: AxiosRequestConfig = {}): Promise<RequestArgs> => {
            const localVarPath = `/api/smartevents_mgmt/v2/errors`;
            // use dummy base URL string because the URL constructor only accepts absolute URLs.
            const localVarUrlObj = new URL(localVarPath, DUMMY_BASE_URL);
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }

            const localVarRequestOptions = { method: 'GET', ...baseOptions, ...options};
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;

            if (page !== undefined) {
                localVarQueryParameter['page'] = page;
            }

            if (size !== undefined) {
                localVarQueryParameter['size'] = size;
            }


    
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
 * ErrorCatalogApi - functional programming interface
 * @export
 */
export const ErrorCatalogApiFp = function(configuration?: Configuration) {
    const localVarAxiosParamCreator = ErrorCatalogApiAxiosParamCreator(configuration)
    return {
        /**
         * Get an error from the error catalog.
         * @summary Get an error from the error catalog.
         * @param {number} id 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async getError(id: number, options?: AxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<BridgeError>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.getError(id, options);
            return createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration);
        },
        /**
         * Get the list of errors from the error catalog.
         * @summary Get the list of errors.
         * @param {number} [page] 
         * @param {number} [size] 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async getErrors(page?: number, size?: number, options?: AxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<ErrorListResponse>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.getErrors(page, size, options);
            return createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration);
        },
    }
};

/**
 * ErrorCatalogApi - factory interface
 * @export
 */
export const ErrorCatalogApiFactory = function (configuration?: Configuration, basePath?: string, axios?: AxiosInstance) {
    const localVarFp = ErrorCatalogApiFp(configuration)
    return {
        /**
         * Get an error from the error catalog.
         * @summary Get an error from the error catalog.
         * @param {number} id 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getError(id: number, options?: any): AxiosPromise<BridgeError> {
            return localVarFp.getError(id, options).then((request) => request(axios, basePath));
        },
        /**
         * Get the list of errors from the error catalog.
         * @summary Get the list of errors.
         * @param {number} [page] 
         * @param {number} [size] 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getErrors(page?: number, size?: number, options?: any): AxiosPromise<ErrorListResponse> {
            return localVarFp.getErrors(page, size, options).then((request) => request(axios, basePath));
        },
    };
};

/**
 * ErrorCatalogApi - interface
 * @export
 * @interface ErrorCatalogApi
 */
export interface ErrorCatalogApiInterface {
    /**
     * Get an error from the error catalog.
     * @summary Get an error from the error catalog.
     * @param {number} id 
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof ErrorCatalogApiInterface
     */
    getError(id: number, options?: AxiosRequestConfig): AxiosPromise<BridgeError>;

    /**
     * Get the list of errors from the error catalog.
     * @summary Get the list of errors.
     * @param {number} [page] 
     * @param {number} [size] 
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof ErrorCatalogApiInterface
     */
    getErrors(page?: number, size?: number, options?: AxiosRequestConfig): AxiosPromise<ErrorListResponse>;

}

/**
 * ErrorCatalogApi - object-oriented interface
 * @export
 * @class ErrorCatalogApi
 * @extends {BaseAPI}
 */
export class ErrorCatalogApi extends BaseAPI implements ErrorCatalogApiInterface {
    /**
     * Get an error from the error catalog.
     * @summary Get an error from the error catalog.
     * @param {number} id 
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof ErrorCatalogApi
     */
    public getError(id: number, options?: AxiosRequestConfig) {
        return ErrorCatalogApiFp(this.configuration).getError(id, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * Get the list of errors from the error catalog.
     * @summary Get the list of errors.
     * @param {number} [page] 
     * @param {number} [size] 
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof ErrorCatalogApi
     */
    public getErrors(page?: number, size?: number, options?: AxiosRequestConfig) {
        return ErrorCatalogApiFp(this.configuration).getErrors(page, size, options).then((request) => request(this.axios, this.basePath));
    }
}
