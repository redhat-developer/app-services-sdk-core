/*
 * Red Hat Openshift SmartEvents Fleet Manager V1
 *
 * The API exposed by the fleet manager of the SmartEvents service.
 *
 * API version: 0.0.1
 * Contact: openbridge-dev@redhat.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package smarteventsmgmtclient

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
)

// Linger please
var (
	_ _context.Context
)

type ErrorCatalogApi interface {

	/*
	 * ErrorsAPIGetError Get an error from the error catalog.
	 * Get an error from the error catalog.
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @param id
	 * @return ApiErrorsAPIGetErrorRequest
	 */
	ErrorsAPIGetError(ctx _context.Context, id int32) ApiErrorsAPIGetErrorRequest

	/*
	 * ErrorsAPIGetErrorExecute executes the request
	 * @return BridgeError
	 */
	ErrorsAPIGetErrorExecute(r ApiErrorsAPIGetErrorRequest) (BridgeError, *_nethttp.Response, error)

	/*
	 * ErrorsAPIGetErrors Get the list of errors.
	 * Get the list of errors from the error catalog.
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @return ApiErrorsAPIGetErrorsRequest
	 */
	ErrorsAPIGetErrors(ctx _context.Context) ApiErrorsAPIGetErrorsRequest

	/*
	 * ErrorsAPIGetErrorsExecute executes the request
	 * @return ErrorListResponse
	 */
	ErrorsAPIGetErrorsExecute(r ApiErrorsAPIGetErrorsRequest) (ErrorListResponse, *_nethttp.Response, error)
}

// ErrorCatalogApiService ErrorCatalogApi service
type ErrorCatalogApiService service

type ApiErrorsAPIGetErrorRequest struct {
	ctx _context.Context
	ApiService ErrorCatalogApi
	id int32
}


func (r ApiErrorsAPIGetErrorRequest) Execute() (BridgeError, *_nethttp.Response, error) {
	return r.ApiService.ErrorsAPIGetErrorExecute(r)
}

/*
 * ErrorsAPIGetError Get an error from the error catalog.
 * Get an error from the error catalog.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id
 * @return ApiErrorsAPIGetErrorRequest
 */
func (a *ErrorCatalogApiService) ErrorsAPIGetError(ctx _context.Context, id int32) ApiErrorsAPIGetErrorRequest {
	return ApiErrorsAPIGetErrorRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

/*
 * Execute executes the request
 * @return BridgeError
 */
func (a *ErrorCatalogApiService) ErrorsAPIGetErrorExecute(r ApiErrorsAPIGetErrorRequest) (BridgeError, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  BridgeError
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ErrorCatalogApiService.ErrorsAPIGetError")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/smartevents_mgmt/v1/errors/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorsList
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ErrorsList
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiErrorsAPIGetErrorsRequest struct {
	ctx _context.Context
	ApiService ErrorCatalogApi
	page *int32
	size *int32
}

func (r ApiErrorsAPIGetErrorsRequest) Page(page int32) ApiErrorsAPIGetErrorsRequest {
	r.page = &page
	return r
}
func (r ApiErrorsAPIGetErrorsRequest) Size(size int32) ApiErrorsAPIGetErrorsRequest {
	r.size = &size
	return r
}

func (r ApiErrorsAPIGetErrorsRequest) Execute() (ErrorListResponse, *_nethttp.Response, error) {
	return r.ApiService.ErrorsAPIGetErrorsExecute(r)
}

/*
 * ErrorsAPIGetErrors Get the list of errors.
 * Get the list of errors from the error catalog.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiErrorsAPIGetErrorsRequest
 */
func (a *ErrorCatalogApiService) ErrorsAPIGetErrors(ctx _context.Context) ApiErrorsAPIGetErrorsRequest {
	return ApiErrorsAPIGetErrorsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return ErrorListResponse
 */
func (a *ErrorCatalogApiService) ErrorsAPIGetErrorsExecute(r ApiErrorsAPIGetErrorsRequest) (ErrorListResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ErrorListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ErrorCatalogApiService.ErrorsAPIGetErrors")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/smartevents_mgmt/v1/errors"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.page != nil {
		localVarQueryParams.Add("page", parameterToString(*r.page, ""))
	}
	if r.size != nil {
		localVarQueryParams.Add("size", parameterToString(*r.size, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorsList
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ErrorsList
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
