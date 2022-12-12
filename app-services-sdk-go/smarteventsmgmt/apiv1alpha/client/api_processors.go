/*
 * Red Hat Openshift SmartEvents Fleet Manager V2
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
	"reflect"
)

// Linger please
var (
	_ _context.Context
)

type ProcessorsApi interface {

	/*
	 * ProcessorsAPICreateProcessor Create a Processor of a Bridge instance
	 * Create a Processor of a Bridge instance for the authenticated user.
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @param bridgeId
	 * @return ApiProcessorsAPICreateProcessorRequest
	 */
	ProcessorsAPICreateProcessor(ctx _context.Context, bridgeId string) ApiProcessorsAPICreateProcessorRequest

	/*
	 * ProcessorsAPICreateProcessorExecute executes the request
	 * @return ProcessorResponse
	 */
	ProcessorsAPICreateProcessorExecute(r ApiProcessorsAPICreateProcessorRequest) (ProcessorResponse, *_nethttp.Response, error)

	/*
	 * ProcessorsAPIDeleteProcessor Delete a Processor of a Bridge instance
	 * Delete a Processor of a Bridge instance for the authenticated user.
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @param bridgeId
	 * @param processorId
	 * @return ApiProcessorsAPIDeleteProcessorRequest
	 */
	ProcessorsAPIDeleteProcessor(ctx _context.Context, bridgeId string, processorId string) ApiProcessorsAPIDeleteProcessorRequest

	/*
	 * ProcessorsAPIDeleteProcessorExecute executes the request
	 */
	ProcessorsAPIDeleteProcessorExecute(r ApiProcessorsAPIDeleteProcessorRequest) (*_nethttp.Response, error)

	/*
	 * ProcessorsAPIGetProcessor Get a Processor of a Bridge instance
	 * Get a Processor of a Bridge instance for the authenticated user.
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @param bridgeId
	 * @param processorId
	 * @return ApiProcessorsAPIGetProcessorRequest
	 */
	ProcessorsAPIGetProcessor(ctx _context.Context, bridgeId string, processorId string) ApiProcessorsAPIGetProcessorRequest

	/*
	 * ProcessorsAPIGetProcessorExecute executes the request
	 * @return ProcessorResponse
	 */
	ProcessorsAPIGetProcessorExecute(r ApiProcessorsAPIGetProcessorRequest) (ProcessorResponse, *_nethttp.Response, error)

	/*
	 * ProcessorsAPIGetProcessors Get the list of Processors of a Bridge instance
	 * Get the list of Processors of a Bridge instance for the authenticated user.
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @param bridgeId
	 * @return ApiProcessorsAPIGetProcessorsRequest
	 */
	ProcessorsAPIGetProcessors(ctx _context.Context, bridgeId string) ApiProcessorsAPIGetProcessorsRequest

	/*
	 * ProcessorsAPIGetProcessorsExecute executes the request
	 * @return ProcessorListResponse
	 */
	ProcessorsAPIGetProcessorsExecute(r ApiProcessorsAPIGetProcessorsRequest) (ProcessorListResponse, *_nethttp.Response, error)

	/*
	 * ProcessorsAPIUpdateProcessor Update a Processor instance.
	 * Update a Processor instance for the authenticated user.
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @param bridgeId
	 * @param processorId
	 * @return ApiProcessorsAPIUpdateProcessorRequest
	 */
	ProcessorsAPIUpdateProcessor(ctx _context.Context, bridgeId string, processorId string) ApiProcessorsAPIUpdateProcessorRequest

	/*
	 * ProcessorsAPIUpdateProcessorExecute executes the request
	 * @return ProcessorResponse
	 */
	ProcessorsAPIUpdateProcessorExecute(r ApiProcessorsAPIUpdateProcessorRequest) (ProcessorResponse, *_nethttp.Response, error)
}

// ProcessorsApiService ProcessorsApi service
type ProcessorsApiService service

type ApiProcessorsAPICreateProcessorRequest struct {
	ctx _context.Context
	ApiService ProcessorsApi
	bridgeId string
	processorRequest *ProcessorRequest
}

func (r ApiProcessorsAPICreateProcessorRequest) ProcessorRequest(processorRequest ProcessorRequest) ApiProcessorsAPICreateProcessorRequest {
	r.processorRequest = &processorRequest
	return r
}

func (r ApiProcessorsAPICreateProcessorRequest) Execute() (ProcessorResponse, *_nethttp.Response, error) {
	return r.ApiService.ProcessorsAPICreateProcessorExecute(r)
}

/*
 * ProcessorsAPICreateProcessor Create a Processor of a Bridge instance
 * Create a Processor of a Bridge instance for the authenticated user.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param bridgeId
 * @return ApiProcessorsAPICreateProcessorRequest
 */
func (a *ProcessorsApiService) ProcessorsAPICreateProcessor(ctx _context.Context, bridgeId string) ApiProcessorsAPICreateProcessorRequest {
	return ApiProcessorsAPICreateProcessorRequest{
		ApiService: a,
		ctx: ctx,
		bridgeId: bridgeId,
	}
}

/*
 * Execute executes the request
 * @return ProcessorResponse
 */
func (a *ProcessorsApiService) ProcessorsAPICreateProcessorExecute(r ApiProcessorsAPICreateProcessorRequest) (ProcessorResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ProcessorResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProcessorsApiService.ProcessorsAPICreateProcessor")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/smartevents_mgmt/v2/bridges/{bridgeId}/processors"
	localVarPath = strings.Replace(localVarPath, "{"+"bridgeId"+"}", _neturl.PathEscape(parameterToString(r.bridgeId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if strlen(r.bridgeId) < 1 {
		return localVarReturnValue, nil, reportError("bridgeId must have at least 1 elements")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.processorRequest
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
		if localVarHTTPResponse.StatusCode == 402 {
			var v ErrorsList
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
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

type ApiProcessorsAPIDeleteProcessorRequest struct {
	ctx _context.Context
	ApiService ProcessorsApi
	bridgeId string
	processorId string
}


func (r ApiProcessorsAPIDeleteProcessorRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.ProcessorsAPIDeleteProcessorExecute(r)
}

/*
 * ProcessorsAPIDeleteProcessor Delete a Processor of a Bridge instance
 * Delete a Processor of a Bridge instance for the authenticated user.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param bridgeId
 * @param processorId
 * @return ApiProcessorsAPIDeleteProcessorRequest
 */
func (a *ProcessorsApiService) ProcessorsAPIDeleteProcessor(ctx _context.Context, bridgeId string, processorId string) ApiProcessorsAPIDeleteProcessorRequest {
	return ApiProcessorsAPIDeleteProcessorRequest{
		ApiService: a,
		ctx: ctx,
		bridgeId: bridgeId,
		processorId: processorId,
	}
}

/*
 * Execute executes the request
 */
func (a *ProcessorsApiService) ProcessorsAPIDeleteProcessorExecute(r ApiProcessorsAPIDeleteProcessorRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProcessorsApiService.ProcessorsAPIDeleteProcessor")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/smartevents_mgmt/v2/bridges/{bridgeId}/processors/{processorId}"
	localVarPath = strings.Replace(localVarPath, "{"+"bridgeId"+"}", _neturl.PathEscape(parameterToString(r.bridgeId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"processorId"+"}", _neturl.PathEscape(parameterToString(r.processorId, "")), -1)

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
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
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
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ErrorsList
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ErrorsList
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiProcessorsAPIGetProcessorRequest struct {
	ctx _context.Context
	ApiService ProcessorsApi
	bridgeId string
	processorId string
}


func (r ApiProcessorsAPIGetProcessorRequest) Execute() (ProcessorResponse, *_nethttp.Response, error) {
	return r.ApiService.ProcessorsAPIGetProcessorExecute(r)
}

/*
 * ProcessorsAPIGetProcessor Get a Processor of a Bridge instance
 * Get a Processor of a Bridge instance for the authenticated user.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param bridgeId
 * @param processorId
 * @return ApiProcessorsAPIGetProcessorRequest
 */
func (a *ProcessorsApiService) ProcessorsAPIGetProcessor(ctx _context.Context, bridgeId string, processorId string) ApiProcessorsAPIGetProcessorRequest {
	return ApiProcessorsAPIGetProcessorRequest{
		ApiService: a,
		ctx: ctx,
		bridgeId: bridgeId,
		processorId: processorId,
	}
}

/*
 * Execute executes the request
 * @return ProcessorResponse
 */
func (a *ProcessorsApiService) ProcessorsAPIGetProcessorExecute(r ApiProcessorsAPIGetProcessorRequest) (ProcessorResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ProcessorResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProcessorsApiService.ProcessorsAPIGetProcessor")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/smartevents_mgmt/v2/bridges/{bridgeId}/processors/{processorId}"
	localVarPath = strings.Replace(localVarPath, "{"+"bridgeId"+"}", _neturl.PathEscape(parameterToString(r.bridgeId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"processorId"+"}", _neturl.PathEscape(parameterToString(r.processorId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if strlen(r.bridgeId) < 1 {
		return localVarReturnValue, nil, reportError("bridgeId must have at least 1 elements")
	}
	if strlen(r.processorId) < 1 {
		return localVarReturnValue, nil, reportError("processorId must have at least 1 elements")
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
		if localVarHTTPResponse.StatusCode == 404 {
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

type ApiProcessorsAPIGetProcessorsRequest struct {
	ctx _context.Context
	ApiService ProcessorsApi
	bridgeId string
	name *string
	page *int32
	size *int32
	status *[]ManagedResourceStatus
}

func (r ApiProcessorsAPIGetProcessorsRequest) Name(name string) ApiProcessorsAPIGetProcessorsRequest {
	r.name = &name
	return r
}
func (r ApiProcessorsAPIGetProcessorsRequest) Page(page int32) ApiProcessorsAPIGetProcessorsRequest {
	r.page = &page
	return r
}
func (r ApiProcessorsAPIGetProcessorsRequest) Size(size int32) ApiProcessorsAPIGetProcessorsRequest {
	r.size = &size
	return r
}
func (r ApiProcessorsAPIGetProcessorsRequest) Status(status []ManagedResourceStatus) ApiProcessorsAPIGetProcessorsRequest {
	r.status = &status
	return r
}

func (r ApiProcessorsAPIGetProcessorsRequest) Execute() (ProcessorListResponse, *_nethttp.Response, error) {
	return r.ApiService.ProcessorsAPIGetProcessorsExecute(r)
}

/*
 * ProcessorsAPIGetProcessors Get the list of Processors of a Bridge instance
 * Get the list of Processors of a Bridge instance for the authenticated user.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param bridgeId
 * @return ApiProcessorsAPIGetProcessorsRequest
 */
func (a *ProcessorsApiService) ProcessorsAPIGetProcessors(ctx _context.Context, bridgeId string) ApiProcessorsAPIGetProcessorsRequest {
	return ApiProcessorsAPIGetProcessorsRequest{
		ApiService: a,
		ctx: ctx,
		bridgeId: bridgeId,
	}
}

/*
 * Execute executes the request
 * @return ProcessorListResponse
 */
func (a *ProcessorsApiService) ProcessorsAPIGetProcessorsExecute(r ApiProcessorsAPIGetProcessorsRequest) (ProcessorListResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ProcessorListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProcessorsApiService.ProcessorsAPIGetProcessors")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/smartevents_mgmt/v2/bridges/{bridgeId}/processors"
	localVarPath = strings.Replace(localVarPath, "{"+"bridgeId"+"}", _neturl.PathEscape(parameterToString(r.bridgeId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if strlen(r.bridgeId) < 1 {
		return localVarReturnValue, nil, reportError("bridgeId must have at least 1 elements")
	}

	if r.name != nil {
		localVarQueryParams.Add("name", parameterToString(*r.name, ""))
	}
	if r.page != nil {
		localVarQueryParams.Add("page", parameterToString(*r.page, ""))
	}
	if r.size != nil {
		localVarQueryParams.Add("size", parameterToString(*r.size, ""))
	}
	if r.status != nil {
		t := *r.status
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("status", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("status", parameterToString(t, "multi"))
		}
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
		if localVarHTTPResponse.StatusCode == 404 {
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

type ApiProcessorsAPIUpdateProcessorRequest struct {
	ctx _context.Context
	ApiService ProcessorsApi
	bridgeId string
	processorId string
	processorRequest *ProcessorRequest
}

func (r ApiProcessorsAPIUpdateProcessorRequest) ProcessorRequest(processorRequest ProcessorRequest) ApiProcessorsAPIUpdateProcessorRequest {
	r.processorRequest = &processorRequest
	return r
}

func (r ApiProcessorsAPIUpdateProcessorRequest) Execute() (ProcessorResponse, *_nethttp.Response, error) {
	return r.ApiService.ProcessorsAPIUpdateProcessorExecute(r)
}

/*
 * ProcessorsAPIUpdateProcessor Update a Processor instance.
 * Update a Processor instance for the authenticated user.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param bridgeId
 * @param processorId
 * @return ApiProcessorsAPIUpdateProcessorRequest
 */
func (a *ProcessorsApiService) ProcessorsAPIUpdateProcessor(ctx _context.Context, bridgeId string, processorId string) ApiProcessorsAPIUpdateProcessorRequest {
	return ApiProcessorsAPIUpdateProcessorRequest{
		ApiService: a,
		ctx: ctx,
		bridgeId: bridgeId,
		processorId: processorId,
	}
}

/*
 * Execute executes the request
 * @return ProcessorResponse
 */
func (a *ProcessorsApiService) ProcessorsAPIUpdateProcessorExecute(r ApiProcessorsAPIUpdateProcessorRequest) (ProcessorResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ProcessorResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProcessorsApiService.ProcessorsAPIUpdateProcessor")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/smartevents_mgmt/v2/bridges/{bridgeId}/processors/{processorId}"
	localVarPath = strings.Replace(localVarPath, "{"+"bridgeId"+"}", _neturl.PathEscape(parameterToString(r.bridgeId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"processorId"+"}", _neturl.PathEscape(parameterToString(r.processorId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if strlen(r.bridgeId) < 1 {
		return localVarReturnValue, nil, reportError("bridgeId must have at least 1 elements")
	}
	if strlen(r.processorId) < 1 {
		return localVarReturnValue, nil, reportError("processorId must have at least 1 elements")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.processorRequest
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
		if localVarHTTPResponse.StatusCode == 404 {
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