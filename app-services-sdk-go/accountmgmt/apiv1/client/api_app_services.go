/*
Account Management Service API

Manage user subscriptions and clusters

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package accountmgmtclient

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


type AppServicesApi interface {

	/*
	ApiAccountsMgmtV1AccessTokenPost Return access token generated from registries in docker format

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiApiAccountsMgmtV1AccessTokenPostRequest
	*/
	ApiAccountsMgmtV1AccessTokenPost(ctx context.Context) ApiApiAccountsMgmtV1AccessTokenPostRequest

	// ApiAccountsMgmtV1AccessTokenPostExecute executes the request
	//  @return AccessTokenCfg
	ApiAccountsMgmtV1AccessTokenPostExecute(r ApiApiAccountsMgmtV1AccessTokenPostRequest) (*AccessTokenCfg, *http.Response, error)

	/*
	ApiAccountsMgmtV1CurrentAccountGet Get the authenticated account

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiApiAccountsMgmtV1CurrentAccountGetRequest
	*/
	ApiAccountsMgmtV1CurrentAccountGet(ctx context.Context) ApiApiAccountsMgmtV1CurrentAccountGetRequest

	// ApiAccountsMgmtV1CurrentAccountGetExecute executes the request
	//  @return Account
	ApiAccountsMgmtV1CurrentAccountGetExecute(r ApiApiAccountsMgmtV1CurrentAccountGetRequest) (*Account, *http.Response, error)

	/*
	ApiAccountsMgmtV1OrganizationsOrgIdQuotaCostGet Returns a summary of quota cost

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId The id of organization
	@return ApiApiAccountsMgmtV1OrganizationsOrgIdQuotaCostGetRequest
	*/
	ApiAccountsMgmtV1OrganizationsOrgIdQuotaCostGet(ctx context.Context, orgId string) ApiApiAccountsMgmtV1OrganizationsOrgIdQuotaCostGetRequest

	// ApiAccountsMgmtV1OrganizationsOrgIdQuotaCostGetExecute executes the request
	//  @return QuotaCostList
	ApiAccountsMgmtV1OrganizationsOrgIdQuotaCostGetExecute(r ApiApiAccountsMgmtV1OrganizationsOrgIdQuotaCostGetRequest) (*QuotaCostList, *http.Response, error)

	/*
	ApiAuthorizationsV1SelfTermsReviewPost Review your status of Terms

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiApiAuthorizationsV1SelfTermsReviewPostRequest
	*/
	ApiAuthorizationsV1SelfTermsReviewPost(ctx context.Context) ApiApiAuthorizationsV1SelfTermsReviewPostRequest

	// ApiAuthorizationsV1SelfTermsReviewPostExecute executes the request
	//  @return TermsReviewResponse
	ApiAuthorizationsV1SelfTermsReviewPostExecute(r ApiApiAuthorizationsV1SelfTermsReviewPostRequest) (*TermsReviewResponse, *http.Response, error)
}

// AppServicesApiService AppServicesApi service
type AppServicesApiService service

type ApiApiAccountsMgmtV1AccessTokenPostRequest struct {
	ctx context.Context
	ApiService AppServicesApi
}

func (r ApiApiAccountsMgmtV1AccessTokenPostRequest) Execute() (*AccessTokenCfg, *http.Response, error) {
	return r.ApiService.ApiAccountsMgmtV1AccessTokenPostExecute(r)
}

/*
ApiAccountsMgmtV1AccessTokenPost Return access token generated from registries in docker format

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiApiAccountsMgmtV1AccessTokenPostRequest
*/
func (a *AppServicesApiService) ApiAccountsMgmtV1AccessTokenPost(ctx context.Context) ApiApiAccountsMgmtV1AccessTokenPostRequest {
	return ApiApiAccountsMgmtV1AccessTokenPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AccessTokenCfg
func (a *AppServicesApiService) ApiAccountsMgmtV1AccessTokenPostExecute(r ApiApiAccountsMgmtV1AccessTokenPostRequest) (*AccessTokenCfg, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AccessTokenCfg
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AppServicesApiService.ApiAccountsMgmtV1AccessTokenPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/accounts_mgmt/v1/access_token"

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

type ApiApiAccountsMgmtV1CurrentAccountGetRequest struct {
	ctx context.Context
	ApiService AppServicesApi
	fetchLabels *bool
}

// If true, includes the labels on a subscription/organization/account in the output. Could slow request response time.
func (r ApiApiAccountsMgmtV1CurrentAccountGetRequest) FetchLabels(fetchLabels bool) ApiApiAccountsMgmtV1CurrentAccountGetRequest {
	r.fetchLabels = &fetchLabels
	return r
}

func (r ApiApiAccountsMgmtV1CurrentAccountGetRequest) Execute() (*Account, *http.Response, error) {
	return r.ApiService.ApiAccountsMgmtV1CurrentAccountGetExecute(r)
}

/*
ApiAccountsMgmtV1CurrentAccountGet Get the authenticated account

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiApiAccountsMgmtV1CurrentAccountGetRequest
*/
func (a *AppServicesApiService) ApiAccountsMgmtV1CurrentAccountGet(ctx context.Context) ApiApiAccountsMgmtV1CurrentAccountGetRequest {
	return ApiApiAccountsMgmtV1CurrentAccountGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return Account
func (a *AppServicesApiService) ApiAccountsMgmtV1CurrentAccountGetExecute(r ApiApiAccountsMgmtV1CurrentAccountGetRequest) (*Account, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Account
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AppServicesApiService.ApiAccountsMgmtV1CurrentAccountGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/accounts_mgmt/v1/current_account"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fetchLabels != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fetchLabels", r.fetchLabels, "")
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
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["AccessToken"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
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

type ApiApiAccountsMgmtV1OrganizationsOrgIdQuotaCostGetRequest struct {
	ctx context.Context
	ApiService AppServicesApi
	orgId string
	search *string
	fetchRelatedResources *bool
	forceRecalc *bool
	fetchCloudAccounts *bool
}

// Specifies the search criteria. The syntax of this parameter is similar to the syntax of the _where_ clause of an SQL statement, using the names of the json attributes / column names of the account. For example, in order to retrieve all the accounts with a username starting with &#x60;my&#x60;:  &#x60;&#x60;&#x60;sql username like &#39;my%&#39; &#x60;&#x60;&#x60;  The search criteria can also be applied on related resource. For example, in order to retrieve all the subscriptions labeled by &#x60;foo&#x3D;bar&#x60;,  &#x60;&#x60;&#x60;sql labels.key &#x3D; &#39;foo&#39; and labels.value &#x3D; &#39;bar&#39; &#x60;&#x60;&#x60;  If the parameter isn&#39;t provided, or if the value is empty, then all the accounts that the user has permission to see will be returned.
func (r ApiApiAccountsMgmtV1OrganizationsOrgIdQuotaCostGetRequest) Search(search string) ApiApiAccountsMgmtV1OrganizationsOrgIdQuotaCostGetRequest {
	r.search = &search
	return r
}

// If true, includes the related resources in the output. Could slow request response time.
func (r ApiApiAccountsMgmtV1OrganizationsOrgIdQuotaCostGetRequest) FetchRelatedResources(fetchRelatedResources bool) ApiApiAccountsMgmtV1OrganizationsOrgIdQuotaCostGetRequest {
	r.fetchRelatedResources = &fetchRelatedResources
	return r
}

// If true, includes that ConsumedQuota should be recalculated.
func (r ApiApiAccountsMgmtV1OrganizationsOrgIdQuotaCostGetRequest) ForceRecalc(forceRecalc bool) ApiApiAccountsMgmtV1OrganizationsOrgIdQuotaCostGetRequest {
	r.forceRecalc = &forceRecalc
	return r
}

// If true, includes the marketplace cloud accounts in the output. Could slow request response time.
func (r ApiApiAccountsMgmtV1OrganizationsOrgIdQuotaCostGetRequest) FetchCloudAccounts(fetchCloudAccounts bool) ApiApiAccountsMgmtV1OrganizationsOrgIdQuotaCostGetRequest {
	r.fetchCloudAccounts = &fetchCloudAccounts
	return r
}

func (r ApiApiAccountsMgmtV1OrganizationsOrgIdQuotaCostGetRequest) Execute() (*QuotaCostList, *http.Response, error) {
	return r.ApiService.ApiAccountsMgmtV1OrganizationsOrgIdQuotaCostGetExecute(r)
}

/*
ApiAccountsMgmtV1OrganizationsOrgIdQuotaCostGet Returns a summary of quota cost

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param orgId The id of organization
 @return ApiApiAccountsMgmtV1OrganizationsOrgIdQuotaCostGetRequest
*/
func (a *AppServicesApiService) ApiAccountsMgmtV1OrganizationsOrgIdQuotaCostGet(ctx context.Context, orgId string) ApiApiAccountsMgmtV1OrganizationsOrgIdQuotaCostGetRequest {
	return ApiApiAccountsMgmtV1OrganizationsOrgIdQuotaCostGetRequest{
		ApiService: a,
		ctx: ctx,
		orgId: orgId,
	}
}

// Execute executes the request
//  @return QuotaCostList
func (a *AppServicesApiService) ApiAccountsMgmtV1OrganizationsOrgIdQuotaCostGetExecute(r ApiApiAccountsMgmtV1OrganizationsOrgIdQuotaCostGetRequest) (*QuotaCostList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *QuotaCostList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AppServicesApiService.ApiAccountsMgmtV1OrganizationsOrgIdQuotaCostGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/accounts_mgmt/v1/organizations/{orgId}/quota_cost"
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(parameterValueToString(r.orgId, "orgId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.search != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search", r.search, "")
	}
	if r.fetchRelatedResources != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fetchRelatedResources", r.fetchRelatedResources, "")
	}
	if r.forceRecalc != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "forceRecalc", r.forceRecalc, "")
	}
	if r.fetchCloudAccounts != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fetchCloudAccounts", r.fetchCloudAccounts, "")
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

type ApiApiAuthorizationsV1SelfTermsReviewPostRequest struct {
	ctx context.Context
	ApiService AppServicesApi
	selfTermsReview *SelfTermsReview
}

// Data to check self terms for
func (r ApiApiAuthorizationsV1SelfTermsReviewPostRequest) SelfTermsReview(selfTermsReview SelfTermsReview) ApiApiAuthorizationsV1SelfTermsReviewPostRequest {
	r.selfTermsReview = &selfTermsReview
	return r
}

func (r ApiApiAuthorizationsV1SelfTermsReviewPostRequest) Execute() (*TermsReviewResponse, *http.Response, error) {
	return r.ApiService.ApiAuthorizationsV1SelfTermsReviewPostExecute(r)
}

/*
ApiAuthorizationsV1SelfTermsReviewPost Review your status of Terms

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiApiAuthorizationsV1SelfTermsReviewPostRequest
*/
func (a *AppServicesApiService) ApiAuthorizationsV1SelfTermsReviewPost(ctx context.Context) ApiApiAuthorizationsV1SelfTermsReviewPostRequest {
	return ApiApiAuthorizationsV1SelfTermsReviewPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return TermsReviewResponse
func (a *AppServicesApiService) ApiAuthorizationsV1SelfTermsReviewPostExecute(r ApiApiAuthorizationsV1SelfTermsReviewPostRequest) (*TermsReviewResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TermsReviewResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AppServicesApiService.ApiAuthorizationsV1SelfTermsReviewPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/authorizations/v1/self_terms_review"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.selfTermsReview == nil {
		return localVarReturnValue, nil, reportError("selfTermsReview is required and must be specified")
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
	localVarPostBody = r.selfTermsReview
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
		if localVarHTTPResponse.StatusCode == 422 {
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
