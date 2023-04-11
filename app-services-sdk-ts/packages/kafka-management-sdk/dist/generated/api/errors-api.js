"use strict";
/* tslint:disable */
/* eslint-disable */
/**
 * Kafka Management API
 * Kafka Management API is a REST API to manage Kafka instances
 *
 * The version of the OpenAPI document: 1.16.0
 * Contact: rhosak-support@redhat.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
var __awaiter = (this && this.__awaiter) || function (thisArg, _arguments, P, generator) {
    function adopt(value) { return value instanceof P ? value : new P(function (resolve) { resolve(value); }); }
    return new (P || (P = Promise))(function (resolve, reject) {
        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }
        function rejected(value) { try { step(generator["throw"](value)); } catch (e) { reject(e); } }
        function step(result) { result.done ? resolve(result.value) : adopt(result.value).then(fulfilled, rejected); }
        step((generator = generator.apply(thisArg, _arguments || [])).next());
    });
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.ErrorsApi = exports.ErrorsApiFactory = exports.ErrorsApiFp = exports.ErrorsApiAxiosParamCreator = void 0;
const axios_1 = require("axios");
// Some imports not used depending on template conditions
// @ts-ignore
const common_1 = require("../common");
// @ts-ignore
const base_1 = require("../base");
/**
 * ErrorsApi - axios parameter creator
 * @export
 */
const ErrorsApiAxiosParamCreator = function (configuration) {
    return {
        /**
         * Returns the error by Id
         * @param {string} id The ID of record
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getErrorById: (id, options = {}) => __awaiter(this, void 0, void 0, function* () {
            // verify required parameter 'id' is not null or undefined
            (0, common_1.assertParamExists)('getErrorById', 'id', id);
            const localVarPath = `/api/kafkas_mgmt/v1/errors/{id}`
                .replace(`{${"id"}}`, encodeURIComponent(String(id)));
            // use dummy base URL string because the URL constructor only accepts absolute URLs.
            const localVarUrlObj = new URL(localVarPath, common_1.DUMMY_BASE_URL);
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }
            const localVarRequestOptions = Object.assign(Object.assign({ method: 'GET' }, baseOptions), options);
            const localVarHeaderParameter = {};
            const localVarQueryParameter = {};
            (0, common_1.setSearchParams)(localVarUrlObj, localVarQueryParameter);
            let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
            localVarRequestOptions.headers = Object.assign(Object.assign(Object.assign({}, localVarHeaderParameter), headersFromBaseOptions), options.headers);
            return {
                url: (0, common_1.toPathString)(localVarUrlObj),
                options: localVarRequestOptions,
            };
        }),
        /**
         * Returns the list of possible API errors
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getErrors: (options = {}) => __awaiter(this, void 0, void 0, function* () {
            const localVarPath = `/api/kafkas_mgmt/v1/errors`;
            // use dummy base URL string because the URL constructor only accepts absolute URLs.
            const localVarUrlObj = new URL(localVarPath, common_1.DUMMY_BASE_URL);
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }
            const localVarRequestOptions = Object.assign(Object.assign({ method: 'GET' }, baseOptions), options);
            const localVarHeaderParameter = {};
            const localVarQueryParameter = {};
            (0, common_1.setSearchParams)(localVarUrlObj, localVarQueryParameter);
            let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
            localVarRequestOptions.headers = Object.assign(Object.assign(Object.assign({}, localVarHeaderParameter), headersFromBaseOptions), options.headers);
            return {
                url: (0, common_1.toPathString)(localVarUrlObj),
                options: localVarRequestOptions,
            };
        }),
    };
};
exports.ErrorsApiAxiosParamCreator = ErrorsApiAxiosParamCreator;
/**
 * ErrorsApi - functional programming interface
 * @export
 */
const ErrorsApiFp = function (configuration) {
    const localVarAxiosParamCreator = (0, exports.ErrorsApiAxiosParamCreator)(configuration);
    return {
        /**
         * Returns the error by Id
         * @param {string} id The ID of record
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getErrorById(id, options) {
            return __awaiter(this, void 0, void 0, function* () {
                const localVarAxiosArgs = yield localVarAxiosParamCreator.getErrorById(id, options);
                return (0, common_1.createRequestFunction)(localVarAxiosArgs, axios_1.default, base_1.BASE_PATH, configuration);
            });
        },
        /**
         * Returns the list of possible API errors
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getErrors(options) {
            return __awaiter(this, void 0, void 0, function* () {
                const localVarAxiosArgs = yield localVarAxiosParamCreator.getErrors(options);
                return (0, common_1.createRequestFunction)(localVarAxiosArgs, axios_1.default, base_1.BASE_PATH, configuration);
            });
        },
    };
};
exports.ErrorsApiFp = ErrorsApiFp;
/**
 * ErrorsApi - factory interface
 * @export
 */
const ErrorsApiFactory = function (configuration, basePath, axios) {
    const localVarFp = (0, exports.ErrorsApiFp)(configuration);
    return {
        /**
         * Returns the error by Id
         * @param {string} id The ID of record
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getErrorById(id, options) {
            return localVarFp.getErrorById(id, options).then((request) => request(axios, basePath));
        },
        /**
         * Returns the list of possible API errors
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getErrors(options) {
            return localVarFp.getErrors(options).then((request) => request(axios, basePath));
        },
    };
};
exports.ErrorsApiFactory = ErrorsApiFactory;
/**
 * ErrorsApi - object-oriented interface
 * @export
 * @class ErrorsApi
 * @extends {BaseAPI}
 */
class ErrorsApi extends base_1.BaseAPI {
    /**
     * Returns the error by Id
     * @param {string} id The ID of record
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof ErrorsApi
     */
    getErrorById(id, options) {
        return (0, exports.ErrorsApiFp)(this.configuration).getErrorById(id, options).then((request) => request(this.axios, this.basePath));
    }
    /**
     * Returns the list of possible API errors
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof ErrorsApi
     */
    getErrors(options) {
        return (0, exports.ErrorsApiFp)(this.configuration).getErrors(options).then((request) => request(this.axios, this.basePath));
    }
}
exports.ErrorsApi = ErrorsApi;
