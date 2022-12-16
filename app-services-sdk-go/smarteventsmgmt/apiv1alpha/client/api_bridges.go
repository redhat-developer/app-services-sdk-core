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

type BridgesApi interface {

	/*
	 * BridgesAPICreateBridge Create a Bridge instance
	 * Create a Bridge instance for the authenticated user.
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @return ApiBridgesAPICreateBridgeRequest
	 */
	BridgesAPICreateBridge(ctx _context.Context) ApiBridgesAPICreateBridgeRequest

	/*
	 * BridgesAPICreateBridgeExecute executes the request
	 * @return BridgeResponse
	 */
	BridgesAPICreateBridgeExecute(r ApiBridgesAPICreateBridgeRequest) (BridgeResponse, *_nethttp.Response, error)

	/*
	 * BridgesAPIDeleteBridge Delete a Bridge instance
	 * Delete a Bridge instance of the authenticated user by ID.
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @param bridgeId
	 * @return ApiBridgesAPIDeleteBridgeRequest
	 */
	BridgesAPIDeleteBridge(ctx _context.Context, bridgeId string) ApiBridgesAPIDeleteBridgeRequest

	/*
	 * BridgesAPIDeleteBridgeExecute executes the request
	 */
	BridgesAPIDeleteBridgeExecute(r ApiBridgesAPIDeleteBridgeRequest) (*_nethttp.Response, error)

	/*
	 * BridgesAPIGetBridge Get a Bridge instance
	 * Get a Bridge instance of the authenticated user by ID.
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @param bridgeId
	 * @return ApiBridgesAPIGetBridgeRequest
	 */
	BridgesAPIGetBridge(ctx _context.Context, bridgeId string) ApiBridgesAPIGetBridgeRequest

	/*
	 * BridgesAPIGetBridgeExecute executes the request
	 * @return BridgeResponse
	 */
	BridgesAPIGetBridgeExecute(r ApiBridgesAPIGetBridgeRequest) (BridgeResponse, *_nethttp.Response, error)

	/*
	 * BridgesAPIGetBridges Get the list of Bridge instances
	 * Get the list of Bridge instances for the authenticated user.
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @return ApiBridgesAPIGetBridgesRequest
	 */
	BridgesAPIGetBridges(ctx _context.Context) ApiBridgesAPIGetBridgesRequest

	/*
	 * BridgesAPIGetBridgesExecute executes the request
	 * @return BridgeListResponse
	 */
	BridgesAPIGetBridgesExecute(r ApiBridgesAPIGetBridgesRequest) (BridgeListResponse, *_nethttp.Response, error)

	/*
	 * BridgesAPIUpdateBridge Update a Bridge instance
	 * Update a Bridge instance for the authenticated user.
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @param bridgeId
	 * @return ApiBridgesAPIUpdateBridgeRequest
	 */
	BridgesAPIUpdateBridge(ctx _context.Context, bridgeId string) ApiBridgesAPIUpdateBridgeRequest

	/*
	 * BridgesAPIUpdateBridgeExecute executes the request
	 * @return BridgeResponse
	 */
	BridgesAPIUpdateBridgeExecute(r ApiBridgesAPIUpdateBridgeRequest) (BridgeResponse, *_nethttp.Response, error)
}

// BridgesApiService BridgesApi service
type BridgesApiService service

type ApiBridgesAPICreateBridgeRequest struct {
	ctx _context.Context
	ApiService BridgesApi
	bridgeRequest *BridgeRequest
}

func (r ApiBridgesAPICreateBridgeRequest) BridgeRequest(bridgeRequest BridgeRequest) ApiBridgesAPICreateBridgeRequest {
	r.bridgeRequest = &bridgeRequest
	return r
}

func (r ApiBridgesAPICreateBridgeRequest) Execute() (BridgeResponse, *_nethttp.Response, error) {
	return r.ApiService.BridgesAPICreateBridgeExecute(r)
}

/*
 * BridgesAPICreateBridge Create a Bridge instance
 * Create a Bridge instance for the authenticated user.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiBridgesAPICreateBridgeRequest
 */
func (a *BridgesApiService) BridgesAPICreateBridge(ctx _context.Context) ApiBridgesAPICreateBridgeRequest {
	return ApiBridgesAPICreateBridgeRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return BridgeResponse
 */
func (a *BridgesApiService) BridgesAPICreateBridgeExecute(r ApiBridgesAPICreateBridgeRequest) (BridgeResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  BridgeResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BridgesApiService.BridgesAPICreateBridge")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/smartevents_mgmt/v2/bridges"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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
	localVarPostBody = r.bridgeRequest
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

type ApiBridgesAPIDeleteBridgeRequest struct {
	ctx _context.Context
	ApiService BridgesApi
	bridgeId string
}


func (r ApiBridgesAPIDeleteBridgeRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.BridgesAPIDeleteBridgeExecute(r)
}

/*
 * BridgesAPIDeleteBridge Delete a Bridge instance
 * Delete a Bridge instance of the authenticated user by ID.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param bridgeId
 * @return ApiBridgesAPIDeleteBridgeRequest
 */
func (a *BridgesApiService) BridgesAPIDeleteBridge(ctx _context.Context, bridgeId string) ApiBridgesAPIDeleteBridgeRequest {
	return ApiBridgesAPIDeleteBridgeRequest{
		ApiService: a,
		ctx: ctx,
		bridgeId: bridgeId,
	}
}

/*
 * Execute executes the request
 */
func (a *BridgesApiService) BridgesAPIDeleteBridgeExecute(r ApiBridgesAPIDeleteBridgeRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BridgesApiService.BridgesAPIDeleteBridge")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/smartevents_mgmt/v2/bridges/{bridgeId}"
	localVarPath = strings.Replace(localVarPath, "{"+"bridgeId"+"}", _neturl.PathEscape(parameterToString(r.bridgeId, "")), -1)

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

type ApiBridgesAPIGetBridgeRequest struct {
	ctx _context.Context
	ApiService BridgesApi
	bridgeId string
}


func (r ApiBridgesAPIGetBridgeRequest) Execute() (BridgeResponse, *_nethttp.Response, error) {
	return r.ApiService.BridgesAPIGetBridgeExecute(r)
}

/*
 * BridgesAPIGetBridge Get a Bridge instance
 * Get a Bridge instance of the authenticated user by ID.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param bridgeId
 * @return ApiBridgesAPIGetBridgeRequest
 */
func (a *BridgesApiService) BridgesAPIGetBridge(ctx _context.Context, bridgeId string) ApiBridgesAPIGetBridgeRequest {
	return ApiBridgesAPIGetBridgeRequest{
		ApiService: a,
		ctx: ctx,
		bridgeId: bridgeId,
	}
}

/*
 * Execute executes the request
 * @return BridgeResponse
 */
func (a *BridgesApiService) BridgesAPIGetBridgeExecute(r ApiBridgesAPIGetBridgeRequest) (BridgeResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  BridgeResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BridgesApiService.BridgesAPIGetBridge")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/smartevents_mgmt/v2/bridges/{bridgeId}"
	localVarPath = strings.Replace(localVarPath, "{"+"bridgeId"+"}", _neturl.PathEscape(parameterToString(r.bridgeId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if strlen(r.bridgeId) < 1 {
		return localVarReturnValue, nil, reportError("bridgeId must have at least 1 elements")
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

type ApiBridgesAPIGetBridgesRequest struct {
	ctx _context.Context
	ApiService BridgesApi
	name *string
	page *int32
	size *int32
	status *[]ManagedResourceStatus
}

func (r ApiBridgesAPIGetBridgesRequest) Name(name string) ApiBridgesAPIGetBridgesRequest {
	r.name = &name
	return r
}
func (r ApiBridgesAPIGetBridgesRequest) Page(page int32) ApiBridgesAPIGetBridgesRequest {
	r.page = &page
	return r
}
func (r ApiBridgesAPIGetBridgesRequest) Size(size int32) ApiBridgesAPIGetBridgesRequest {
	r.size = &size
	return r
}
func (r ApiBridgesAPIGetBridgesRequest) Status(status []ManagedResourceStatus) ApiBridgesAPIGetBridgesRequest {
	r.status = &status
	return r
}

func (r ApiBridgesAPIGetBridgesRequest) Execute() (BridgeListResponse, *_nethttp.Response, error) {
	return r.ApiService.BridgesAPIGetBridgesExecute(r)
}

/*
 * BridgesAPIGetBridges Get the list of Bridge instances
 * Get the list of Bridge instances for the authenticated user.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiBridgesAPIGetBridgesRequest
 */
func (a *BridgesApiService) BridgesAPIGetBridges(ctx _context.Context) ApiBridgesAPIGetBridgesRequest {
	return ApiBridgesAPIGetBridgesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return BridgeListResponse
 */
func (a *BridgesApiService) BridgesAPIGetBridgesExecute(r ApiBridgesAPIGetBridgesRequest) (BridgeListResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  BridgeListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BridgesApiService.BridgesAPIGetBridges")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/smartevents_mgmt/v2/bridges"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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
		if localVarHTTPResponse.StatusCode == 401 {
			var v ErrorsList
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
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

type ApiBridgesAPIUpdateBridgeRequest struct {
	ctx _context.Context
	ApiService BridgesApi
	bridgeId string
	bridgeRequest *BridgeRequest
}

func (r ApiBridgesAPIUpdateBridgeRequest) BridgeRequest(bridgeRequest BridgeRequest) ApiBridgesAPIUpdateBridgeRequest {
	r.bridgeRequest = &bridgeRequest
	return r
}

func (r ApiBridgesAPIUpdateBridgeRequest) Execute() (BridgeResponse, *_nethttp.Response, error) {
	return r.ApiService.BridgesAPIUpdateBridgeExecute(r)
}

/*
 * BridgesAPIUpdateBridge Update a Bridge instance
 * Update a Bridge instance for the authenticated user.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param bridgeId
 * @return ApiBridgesAPIUpdateBridgeRequest
 */
func (a *BridgesApiService) BridgesAPIUpdateBridge(ctx _context.Context, bridgeId string) ApiBridgesAPIUpdateBridgeRequest {
	return ApiBridgesAPIUpdateBridgeRequest{
		ApiService: a,
		ctx: ctx,
		bridgeId: bridgeId,
	}
}

/*
 * Execute executes the request
 * @return BridgeResponse
 */
func (a *BridgesApiService) BridgesAPIUpdateBridgeExecute(r ApiBridgesAPIUpdateBridgeRequest) (BridgeResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  BridgeResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BridgesApiService.BridgesAPIUpdateBridge")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/smartevents_mgmt/v2/bridges/{bridgeId}"
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
	localVarPostBody = r.bridgeRequest
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
		if localVarHTTPResponse.StatusCode == 401 {
			var v ErrorsList
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
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
