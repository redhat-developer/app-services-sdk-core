/*
Kafka Instance API

API for interacting with Kafka Instance. Includes Produce, Consume and Admin APIs

API version: 0.13.1-SNAPSHOT
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package kafkainstanceclient

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


type GroupsApi interface {

	/*
	DeleteConsumerGroupById Delete a consumer group.

	Delete a consumer group, along with its consumers.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param consumerGroupId Consumer group identifier
	@return ApiDeleteConsumerGroupByIdRequest
	*/
	DeleteConsumerGroupById(ctx context.Context, consumerGroupId string) ApiDeleteConsumerGroupByIdRequest

	// DeleteConsumerGroupByIdExecute executes the request
	DeleteConsumerGroupByIdExecute(r ApiDeleteConsumerGroupByIdRequest) (*http.Response, error)

	/*
	GetConsumerGroupById Get a single consumer group by its unique ID.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param consumerGroupId Consumer group identifier
	@return ApiGetConsumerGroupByIdRequest
	*/
	GetConsumerGroupById(ctx context.Context, consumerGroupId string) ApiGetConsumerGroupByIdRequest

	// GetConsumerGroupByIdExecute executes the request
	//  @return ConsumerGroup
	GetConsumerGroupByIdExecute(r ApiGetConsumerGroupByIdRequest) (*ConsumerGroup, *http.Response, error)

	/*
	GetConsumerGroups List of consumer groups in the Kafka instance.

	Returns a list of all consumer groups for a particular Kafka instance. The consumer groups returned are limited to those records the requestor is authorized to view.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetConsumerGroupsRequest
	*/
	GetConsumerGroups(ctx context.Context) ApiGetConsumerGroupsRequest

	// GetConsumerGroupsExecute executes the request
	//  @return ConsumerGroupList
	GetConsumerGroupsExecute(r ApiGetConsumerGroupsRequest) (*ConsumerGroupList, *http.Response, error)

	/*
	ResetConsumerGroupOffset Reset the offset for a consumer group.

	Reset the offset for a particular consumer group.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param consumerGroupId Consumer group identifier
	@return ApiResetConsumerGroupOffsetRequest
	*/
	ResetConsumerGroupOffset(ctx context.Context, consumerGroupId string) ApiResetConsumerGroupOffsetRequest

	// ResetConsumerGroupOffsetExecute executes the request
	//  @return ConsumerGroupResetOffsetResult
	ResetConsumerGroupOffsetExecute(r ApiResetConsumerGroupOffsetRequest) (*ConsumerGroupResetOffsetResult, *http.Response, error)
}

// GroupsApiService GroupsApi service
type GroupsApiService service

type ApiDeleteConsumerGroupByIdRequest struct {
	ctx context.Context
	ApiService GroupsApi
	consumerGroupId string
}

func (r ApiDeleteConsumerGroupByIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteConsumerGroupByIdExecute(r)
}

/*
DeleteConsumerGroupById Delete a consumer group.

Delete a consumer group, along with its consumers.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param consumerGroupId Consumer group identifier
 @return ApiDeleteConsumerGroupByIdRequest
*/
func (a *GroupsApiService) DeleteConsumerGroupById(ctx context.Context, consumerGroupId string) ApiDeleteConsumerGroupByIdRequest {
	return ApiDeleteConsumerGroupByIdRequest{
		ApiService: a,
		ctx: ctx,
		consumerGroupId: consumerGroupId,
	}
}

// Execute executes the request
func (a *GroupsApiService) DeleteConsumerGroupByIdExecute(r ApiDeleteConsumerGroupByIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GroupsApiService.DeleteConsumerGroupById")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/consumer-groups/{consumerGroupId}"
	localVarPath = strings.Replace(localVarPath, "{"+"consumerGroupId"+"}", url.PathEscape(parameterValueToString(r.consumerGroupId, "consumerGroupId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 423 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 503 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiGetConsumerGroupByIdRequest struct {
	ctx context.Context
	ApiService GroupsApi
	consumerGroupId string
	order *SortDirection
	orderKey *ConsumerGroupDescriptionOrderKey
	partitionFilter *int32
	topic *string
}

// Order items are sorted
func (r ApiGetConsumerGroupByIdRequest) Order(order SortDirection) ApiGetConsumerGroupByIdRequest {
	r.order = &order
	return r
}

func (r ApiGetConsumerGroupByIdRequest) OrderKey(orderKey ConsumerGroupDescriptionOrderKey) ApiGetConsumerGroupByIdRequest {
	r.orderKey = &orderKey
	return r
}

// Value of partition to include. Value -1 means filter is not active.
func (r ApiGetConsumerGroupByIdRequest) PartitionFilter(partitionFilter int32) ApiGetConsumerGroupByIdRequest {
	r.partitionFilter = &partitionFilter
	return r
}

// Filter consumer groups for a specific topic
func (r ApiGetConsumerGroupByIdRequest) Topic(topic string) ApiGetConsumerGroupByIdRequest {
	r.topic = &topic
	return r
}

func (r ApiGetConsumerGroupByIdRequest) Execute() (*ConsumerGroup, *http.Response, error) {
	return r.ApiService.GetConsumerGroupByIdExecute(r)
}

/*
GetConsumerGroupById Get a single consumer group by its unique ID.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param consumerGroupId Consumer group identifier
 @return ApiGetConsumerGroupByIdRequest
*/
func (a *GroupsApiService) GetConsumerGroupById(ctx context.Context, consumerGroupId string) ApiGetConsumerGroupByIdRequest {
	return ApiGetConsumerGroupByIdRequest{
		ApiService: a,
		ctx: ctx,
		consumerGroupId: consumerGroupId,
	}
}

// Execute executes the request
//  @return ConsumerGroup
func (a *GroupsApiService) GetConsumerGroupByIdExecute(r ApiGetConsumerGroupByIdRequest) (*ConsumerGroup, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ConsumerGroup
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GroupsApiService.GetConsumerGroupById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/consumer-groups/{consumerGroupId}"
	localVarPath = strings.Replace(localVarPath, "{"+"consumerGroupId"+"}", url.PathEscape(parameterValueToString(r.consumerGroupId, "consumerGroupId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.order != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "order", r.order, "")
	}
	if r.orderKey != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "orderKey", r.orderKey, "")
	}
	if r.partitionFilter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "partitionFilter", r.partitionFilter, "")
	}
	if r.topic != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "topic", r.topic, "")
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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 503 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetConsumerGroupsRequest struct {
	ctx context.Context
	ApiService GroupsApi
	offset *int32
	limit *int32
	size *int32
	page *int32
	topic *string
	groupIdFilter *string
	order *SortDirection
	orderKey *ConsumerGroupOrderKey
}

// Offset of the first record to return, zero-based
// Deprecated
func (r ApiGetConsumerGroupsRequest) Offset(offset int32) ApiGetConsumerGroupsRequest {
	r.offset = &offset
	return r
}

// Maximum number of records to return
// Deprecated
func (r ApiGetConsumerGroupsRequest) Limit(limit int32) ApiGetConsumerGroupsRequest {
	r.limit = &limit
	return r
}

// Number of records per page
func (r ApiGetConsumerGroupsRequest) Size(size int32) ApiGetConsumerGroupsRequest {
	r.size = &size
	return r
}

// Page number
func (r ApiGetConsumerGroupsRequest) Page(page int32) ApiGetConsumerGroupsRequest {
	r.page = &page
	return r
}

// Return consumer groups where the topic name contains this value
func (r ApiGetConsumerGroupsRequest) Topic(topic string) ApiGetConsumerGroupsRequest {
	r.topic = &topic
	return r
}

// Return the consumer groups where the ID contains this value
func (r ApiGetConsumerGroupsRequest) GroupIdFilter(groupIdFilter string) ApiGetConsumerGroupsRequest {
	r.groupIdFilter = &groupIdFilter
	return r
}

// Order items are sorted
func (r ApiGetConsumerGroupsRequest) Order(order SortDirection) ApiGetConsumerGroupsRequest {
	r.order = &order
	return r
}

func (r ApiGetConsumerGroupsRequest) OrderKey(orderKey ConsumerGroupOrderKey) ApiGetConsumerGroupsRequest {
	r.orderKey = &orderKey
	return r
}

func (r ApiGetConsumerGroupsRequest) Execute() (*ConsumerGroupList, *http.Response, error) {
	return r.ApiService.GetConsumerGroupsExecute(r)
}

/*
GetConsumerGroups List of consumer groups in the Kafka instance.

Returns a list of all consumer groups for a particular Kafka instance. The consumer groups returned are limited to those records the requestor is authorized to view.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetConsumerGroupsRequest
*/
func (a *GroupsApiService) GetConsumerGroups(ctx context.Context) ApiGetConsumerGroupsRequest {
	return ApiGetConsumerGroupsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ConsumerGroupList
func (a *GroupsApiService) GetConsumerGroupsExecute(r ApiGetConsumerGroupsRequest) (*ConsumerGroupList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ConsumerGroupList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GroupsApiService.GetConsumerGroups")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/consumer-groups"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	}
	if r.size != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "size", r.size, "")
	}
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "")
	}
	if r.topic != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "topic", r.topic, "")
	}
	if r.groupIdFilter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "group-id-filter", r.groupIdFilter, "")
	}
	if r.order != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "order", r.order, "")
	}
	if r.orderKey != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "orderKey", r.orderKey, "")
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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 503 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiResetConsumerGroupOffsetRequest struct {
	ctx context.Context
	ApiService GroupsApi
	consumerGroupId string
	consumerGroupResetOffsetParameters *ConsumerGroupResetOffsetParameters
}

func (r ApiResetConsumerGroupOffsetRequest) ConsumerGroupResetOffsetParameters(consumerGroupResetOffsetParameters ConsumerGroupResetOffsetParameters) ApiResetConsumerGroupOffsetRequest {
	r.consumerGroupResetOffsetParameters = &consumerGroupResetOffsetParameters
	return r
}

func (r ApiResetConsumerGroupOffsetRequest) Execute() (*ConsumerGroupResetOffsetResult, *http.Response, error) {
	return r.ApiService.ResetConsumerGroupOffsetExecute(r)
}

/*
ResetConsumerGroupOffset Reset the offset for a consumer group.

Reset the offset for a particular consumer group.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param consumerGroupId Consumer group identifier
 @return ApiResetConsumerGroupOffsetRequest
*/
func (a *GroupsApiService) ResetConsumerGroupOffset(ctx context.Context, consumerGroupId string) ApiResetConsumerGroupOffsetRequest {
	return ApiResetConsumerGroupOffsetRequest{
		ApiService: a,
		ctx: ctx,
		consumerGroupId: consumerGroupId,
	}
}

// Execute executes the request
//  @return ConsumerGroupResetOffsetResult
func (a *GroupsApiService) ResetConsumerGroupOffsetExecute(r ApiResetConsumerGroupOffsetRequest) (*ConsumerGroupResetOffsetResult, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ConsumerGroupResetOffsetResult
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GroupsApiService.ResetConsumerGroupOffset")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/consumer-groups/{consumerGroupId}/reset-offset"
	localVarPath = strings.Replace(localVarPath, "{"+"consumerGroupId"+"}", url.PathEscape(parameterValueToString(r.consumerGroupId, "consumerGroupId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.consumerGroupResetOffsetParameters == nil {
		return localVarReturnValue, nil, reportError("consumerGroupResetOffsetParameters is required and must be specified")
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
	localVarPostBody = r.consumerGroupResetOffsetParameters
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 503 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
