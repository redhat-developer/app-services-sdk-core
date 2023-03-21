/*
Apicurio Registry API [v2]

Apicurio Registry is a datastore for standard event schemas and API designs. Apicurio Registry enables developers to manage and share the structure of their data using a REST interface. For example, client applications can dynamically push or pull the latest updates to or from the registry without needing to redeploy. Apicurio Registry also enables developers to create rules that govern how registry content can evolve over time. For example, this includes rules for content validation and version compatibility.  The Apicurio Registry REST API enables client applications to manage the artifacts in the registry. This API provides create, read, update, and delete operations for schema and API artifacts, rules, versions, and metadata.   The supported artifact types include: - Apache Avro schema - AsyncAPI specification - Google protocol buffers - GraphQL schema - JSON Schema - Kafka Connect schema - OpenAPI specification - Web Services Description Language - XML Schema Definition   **Important**: The Apicurio Registry REST API is available from `https://MY-REGISTRY-URL/apis/registry/v2` by default. Therefore you must prefix all API operation paths with `../apis/registry/v2` in this case. For example: `../apis/registry/v2/ids/globalIds/{globalId}`. 

API version: 2.4.x
Contact: apicurio@lists.jboss.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package registryinstanceclient

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"os"
)


type SearchApi interface {

	/*
	SearchArtifacts Search for artifacts

	Returns a paginated list of all artifacts that match the provided filter criteria.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiSearchArtifactsRequest
	*/
	SearchArtifacts(ctx context.Context) ApiSearchArtifactsRequest

	// SearchArtifactsExecute executes the request
	//  @return ArtifactSearchResults
	SearchArtifactsExecute(r ApiSearchArtifactsRequest) (*ArtifactSearchResults, *http.Response, error)

	/*
	SearchArtifactsByContent Search for artifacts by content

	Returns a paginated list of all artifacts with at least one version that matches the
posted content.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiSearchArtifactsByContentRequest
	*/
	SearchArtifactsByContent(ctx context.Context) ApiSearchArtifactsByContentRequest

	// SearchArtifactsByContentExecute executes the request
	//  @return ArtifactSearchResults
	SearchArtifactsByContentExecute(r ApiSearchArtifactsByContentRequest) (*ArtifactSearchResults, *http.Response, error)
}

// SearchApiService SearchApi service
type SearchApiService service

type ApiSearchArtifactsRequest struct {
	ctx context.Context
	ApiService SearchApi
	name *string
	offset *int32
	limit *int32
	order *SortOrder
	orderby *SortBy
	labels *[]string
	properties *[]string
	description *string
	group *string
	globalId *int64
	contentId *int64
}

// Filter by artifact name.
func (r ApiSearchArtifactsRequest) Name(name string) ApiSearchArtifactsRequest {
	r.name = &name
	return r
}

// The number of artifacts to skip before starting to collect the result set.  Defaults to 0.
func (r ApiSearchArtifactsRequest) Offset(offset int32) ApiSearchArtifactsRequest {
	r.offset = &offset
	return r
}

// The number of artifacts to return.  Defaults to 20.
func (r ApiSearchArtifactsRequest) Limit(limit int32) ApiSearchArtifactsRequest {
	r.limit = &limit
	return r
}

// Sort order, ascending (&#x60;asc&#x60;) or descending (&#x60;desc&#x60;).
func (r ApiSearchArtifactsRequest) Order(order SortOrder) ApiSearchArtifactsRequest {
	r.order = &order
	return r
}

// The field to sort by.  Can be one of:  * &#x60;name&#x60; * &#x60;createdOn&#x60; 
func (r ApiSearchArtifactsRequest) Orderby(orderby SortBy) ApiSearchArtifactsRequest {
	r.orderby = &orderby
	return r
}

// Filter by label.  Include one or more label to only return artifacts containing all of the specified labels.
func (r ApiSearchArtifactsRequest) Labels(labels []string) ApiSearchArtifactsRequest {
	r.labels = &labels
	return r
}

// Filter by one or more name/value property.  Separate each name/value pair using a colon.  For example &#x60;properties&#x3D;foo:bar&#x60; will return only artifacts with a custom property named &#x60;foo&#x60; and value &#x60;bar&#x60;.
func (r ApiSearchArtifactsRequest) Properties(properties []string) ApiSearchArtifactsRequest {
	r.properties = &properties
	return r
}

// Filter by description.
func (r ApiSearchArtifactsRequest) Description(description string) ApiSearchArtifactsRequest {
	r.description = &description
	return r
}

// Filter by artifact group.
func (r ApiSearchArtifactsRequest) Group(group string) ApiSearchArtifactsRequest {
	r.group = &group
	return r
}

// Filter by globalId.
func (r ApiSearchArtifactsRequest) GlobalId(globalId int64) ApiSearchArtifactsRequest {
	r.globalId = &globalId
	return r
}

// Filter by contentId.
func (r ApiSearchArtifactsRequest) ContentId(contentId int64) ApiSearchArtifactsRequest {
	r.contentId = &contentId
	return r
}

func (r ApiSearchArtifactsRequest) Execute() (*ArtifactSearchResults, *http.Response, error) {
	return r.ApiService.SearchArtifactsExecute(r)
}

/*
SearchArtifacts Search for artifacts

Returns a paginated list of all artifacts that match the provided filter criteria.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSearchArtifactsRequest
*/
func (a *SearchApiService) SearchArtifacts(ctx context.Context) ApiSearchArtifactsRequest {
	return ApiSearchArtifactsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ArtifactSearchResults
func (a *SearchApiService) SearchArtifactsExecute(r ApiSearchArtifactsRequest) (*ArtifactSearchResults, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ArtifactSearchResults
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SearchApiService.SearchArtifacts")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/search/artifacts"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.name != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "name", r.name, "")
	}
	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	}
	if r.order != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "order", r.order, "")
	}
	if r.orderby != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "orderby", r.orderby, "")
	}
	if r.labels != nil {
		t := *r.labels
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "labels", s.Index(i), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "labels", t, "multi")
		}
	}
	if r.properties != nil {
		t := *r.properties
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "properties", s.Index(i), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "properties", t, "multi")
		}
	}
	if r.description != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "description", r.description, "")
	}
	if r.group != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "group", r.group, "")
	}
	if r.globalId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "globalId", r.globalId, "")
	}
	if r.contentId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "contentId", r.contentId, "")
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
		if localVarHTTPResponse.StatusCode == 500 {
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

type ApiSearchArtifactsByContentRequest struct {
	ctx context.Context
	ApiService SearchApi
	body *os.File
	canonical *bool
	artifactType *string
	offset *int32
	limit *int32
	order *string
	orderby *string
}

// The content to search for.
func (r ApiSearchArtifactsByContentRequest) Body(body *os.File) ApiSearchArtifactsByContentRequest {
	r.body = body
	return r
}

// Parameter that can be set to &#x60;true&#x60; to indicate that the server should \&quot;canonicalize\&quot; the content when searching for matching artifacts.  Canonicalization is unique to each artifact type, but typically involves removing any extra whitespace and formatting the content in a consistent manner.  Must be used along with the &#x60;artifactType&#x60; query parameter.
func (r ApiSearchArtifactsByContentRequest) Canonical(canonical bool) ApiSearchArtifactsByContentRequest {
	r.canonical = &canonical
	return r
}

// Indicates the type of artifact represented by the content being used for the search.  This is only needed when using the &#x60;canonical&#x60; query parameter, so that the server knows how to canonicalize the content prior to searching for matching artifacts.
func (r ApiSearchArtifactsByContentRequest) ArtifactType(artifactType string) ApiSearchArtifactsByContentRequest {
	r.artifactType = &artifactType
	return r
}

// The number of artifacts to skip before starting to collect the result set.  Defaults to 0.
func (r ApiSearchArtifactsByContentRequest) Offset(offset int32) ApiSearchArtifactsByContentRequest {
	r.offset = &offset
	return r
}

// The number of artifacts to return.  Defaults to 20.
func (r ApiSearchArtifactsByContentRequest) Limit(limit int32) ApiSearchArtifactsByContentRequest {
	r.limit = &limit
	return r
}

// Sort order, ascending (&#x60;asc&#x60;) or descending (&#x60;desc&#x60;).
func (r ApiSearchArtifactsByContentRequest) Order(order string) ApiSearchArtifactsByContentRequest {
	r.order = &order
	return r
}

// The field to sort by.  Can be one of:  * &#x60;name&#x60; * &#x60;createdOn&#x60; 
func (r ApiSearchArtifactsByContentRequest) Orderby(orderby string) ApiSearchArtifactsByContentRequest {
	r.orderby = &orderby
	return r
}

func (r ApiSearchArtifactsByContentRequest) Execute() (*ArtifactSearchResults, *http.Response, error) {
	return r.ApiService.SearchArtifactsByContentExecute(r)
}

/*
SearchArtifactsByContent Search for artifacts by content

Returns a paginated list of all artifacts with at least one version that matches the
posted content.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSearchArtifactsByContentRequest
*/
func (a *SearchApiService) SearchArtifactsByContent(ctx context.Context) ApiSearchArtifactsByContentRequest {
	return ApiSearchArtifactsByContentRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ArtifactSearchResults
func (a *SearchApiService) SearchArtifactsByContentExecute(r ApiSearchArtifactsByContentRequest) (*ArtifactSearchResults, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ArtifactSearchResults
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SearchApiService.SearchArtifactsByContent")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/search/artifacts"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, reportError("body is required and must be specified")
	}

	if r.canonical != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "canonical", r.canonical, "")
	}
	if r.artifactType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "artifactType", r.artifactType, "")
	}
	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	}
	if r.order != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "order", r.order, "")
	}
	if r.orderby != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "orderby", r.orderby, "")
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
	// body params
	localVarPostBody = r.body
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
		if localVarHTTPResponse.StatusCode == 500 {
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
