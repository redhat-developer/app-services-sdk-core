package com.openshift.cloud.api.accountmanagement.api.accounts_mgmt.v1.subscriptions.item.reserved_resources;

import com.microsoft.kiota.HttpMethod;
import com.microsoft.kiota.RequestAdapter;
import com.microsoft.kiota.RequestHeaders;
import com.microsoft.kiota.RequestInformation;
import com.microsoft.kiota.RequestOption;
import com.microsoft.kiota.serialization.Parsable;
import com.microsoft.kiota.serialization.ParsableFactory;
import com.openshift.cloud.api.accountmanagement.models.Error;
import com.openshift.cloud.api.accountmanagement.models.ReservedResourceList;
import java.net.URISyntaxException;
import java.util.Collection;
import java.util.Collections;
import java.util.HashMap;
import java.util.Map;
import java.util.Objects;
/**
 * Builds and executes requests for operations under /api/accounts_mgmt/v1/subscriptions/{id}/reserved_resources
 */
public class Reserved_resourcesRequestBuilder {
    /** Path parameters for the request */
    private HashMap<String, Object> pathParameters;
    /** The request adapter to use to execute the requests. */
    private RequestAdapter requestAdapter;
    /** Url template to use to build the URL for the current request builder */
    private String urlTemplate;
    /**
     * Instantiates a new Reserved_resourcesRequestBuilder and sets the default values.
     * @param pathParameters Path parameters for the request
     * @param requestAdapter The request adapter to use to execute the requests.
     * @return a void
     */
    @javax.annotation.Nullable
    public Reserved_resourcesRequestBuilder(@javax.annotation.Nonnull final HashMap<String, Object> pathParameters, @javax.annotation.Nonnull final RequestAdapter requestAdapter) {
        Objects.requireNonNull(pathParameters);
        Objects.requireNonNull(requestAdapter);
        this.urlTemplate = "{+baseurl}/api/accounts_mgmt/v1/subscriptions/{id}/reserved_resources{?page*,size*,search*,orderBy*}";
        final HashMap<String, Object> urlTplParams = new HashMap<String, Object>(pathParameters);
        this.pathParameters = urlTplParams;
        this.requestAdapter = requestAdapter;
    }
    /**
     * Instantiates a new Reserved_resourcesRequestBuilder and sets the default values.
     * @param rawUrl The raw URL to use for the request builder.
     * @param requestAdapter The request adapter to use to execute the requests.
     * @return a void
     */
    @javax.annotation.Nullable
    public Reserved_resourcesRequestBuilder(@javax.annotation.Nonnull final String rawUrl, @javax.annotation.Nonnull final RequestAdapter requestAdapter) {
        this.urlTemplate = "{+baseurl}/api/accounts_mgmt/v1/subscriptions/{id}/reserved_resources{?page*,size*,search*,orderBy*}";
        final HashMap<String, Object> urlTplParams = new HashMap<String, Object>();
        urlTplParams.put("request-raw-url", rawUrl);
        this.pathParameters = urlTplParams;
        this.requestAdapter = requestAdapter;
    }
    /**
     * Returns a list of reserved resources
     * @return a CompletableFuture of ReservedResourceList
     */
    @javax.annotation.Nonnull
    public java.util.concurrent.CompletableFuture<ReservedResourceList> get() {
        try {
            final RequestInformation requestInfo = toGetRequestInformation(null);
            final HashMap<String, ParsableFactory<? extends Parsable>> errorMapping = new HashMap<String, ParsableFactory<? extends Parsable>>();
            errorMapping.put("401", Error::createFromDiscriminatorValue);
            errorMapping.put("403", Error::createFromDiscriminatorValue);
            errorMapping.put("500", Error::createFromDiscriminatorValue);
            return this.requestAdapter.sendAsync(requestInfo, ReservedResourceList::createFromDiscriminatorValue, errorMapping);
        } catch (URISyntaxException ex) {
            final java.util.concurrent.CompletableFuture<ReservedResourceList> executionException = new java.util.concurrent.CompletableFuture<ReservedResourceList>();
            executionException.completeExceptionally(ex);
            return executionException;
        }
    }
    /**
     * Returns a list of reserved resources
     * @param requestConfiguration Configuration for the request such as headers, query parameters, and middleware options.
     * @return a CompletableFuture of ReservedResourceList
     */
    @javax.annotation.Nonnull
    public java.util.concurrent.CompletableFuture<ReservedResourceList> get(@javax.annotation.Nullable final java.util.function.Consumer<GetRequestConfiguration> requestConfiguration) {
        try {
            final RequestInformation requestInfo = toGetRequestInformation(requestConfiguration);
            final HashMap<String, ParsableFactory<? extends Parsable>> errorMapping = new HashMap<String, ParsableFactory<? extends Parsable>>();
            errorMapping.put("401", Error::createFromDiscriminatorValue);
            errorMapping.put("403", Error::createFromDiscriminatorValue);
            errorMapping.put("500", Error::createFromDiscriminatorValue);
            return this.requestAdapter.sendAsync(requestInfo, ReservedResourceList::createFromDiscriminatorValue, errorMapping);
        } catch (URISyntaxException ex) {
            final java.util.concurrent.CompletableFuture<ReservedResourceList> executionException = new java.util.concurrent.CompletableFuture<ReservedResourceList>();
            executionException.completeExceptionally(ex);
            return executionException;
        }
    }
    /**
     * Returns a list of reserved resources
     * @return a RequestInformation
     */
    @javax.annotation.Nonnull
    public RequestInformation toGetRequestInformation() throws URISyntaxException {
        return toGetRequestInformation(null);
    }
    /**
     * Returns a list of reserved resources
     * @param requestConfiguration Configuration for the request such as headers, query parameters, and middleware options.
     * @return a RequestInformation
     */
    @javax.annotation.Nonnull
    public RequestInformation toGetRequestInformation(@javax.annotation.Nullable final java.util.function.Consumer<GetRequestConfiguration> requestConfiguration) throws URISyntaxException {
        final RequestInformation requestInfo = new RequestInformation();
        requestInfo.httpMethod = HttpMethod.GET;
        requestInfo.urlTemplate = urlTemplate;
        requestInfo.pathParameters = pathParameters;
        requestInfo.headers.add("Accept", "application/json");
        if (requestConfiguration != null) {
            final GetRequestConfiguration requestConfig = new GetRequestConfiguration();
            requestConfiguration.accept(requestConfig);
            requestInfo.addQueryParameters(requestConfig.queryParameters);
            requestInfo.headers.putAll(requestConfig.headers);
            requestInfo.addRequestOptions(requestConfig.options);
        }
        return requestInfo;
    }
    /**
     * Returns a list of reserved resources
     */
    public class GetQueryParameters {
        /** Specifies the order by criteria. The syntax of this parameter issimilar to the syntax of the _order by_ clause of an SQL statement,but using the names of the json attributes / column of the account.For example, in order to retrieve all accounts ordered by username:```sqlusername asc```Or in order to retrieve all accounts ordered by username _and_ first name:```sqlusername asc, firstName asc```If the parameter isn't provided, or if the value is empty, thenno explicit ordering will be applied. */
        @javax.annotation.Nullable
        public String orderBy;
        /** Page number of record list when record list exceeds specified page size */
        @javax.annotation.Nullable
        public Integer page;
        /** Specifies the search criteria. The syntax of this parameter issimilar to the syntax of the _where_ clause of an SQL statement,using the names of the json attributes / column names of the account.For example, in order to retrieve all the accounts with a usernamestarting with `my`:```sqlusername like 'my%'```The search criteria can also be applied on related resource.For example, in order to retrieve all the subscriptions labeled by `foo=bar`,```sqllabels.key = 'foo' and labels.value = 'bar'```If the parameter isn't provided, or if the value is empty, thenall the accounts that the user has permission to see will bereturned. */
        @javax.annotation.Nullable
        public String search;
        /** Maximum number of records to return */
        @javax.annotation.Nullable
        public Integer size;
    }
    /**
     * Configuration for the request such as headers, query parameters, and middleware options.
     */
    public class GetRequestConfiguration {
        /** Request headers */
        @javax.annotation.Nullable
        public RequestHeaders headers = new RequestHeaders();
        /** Request options */
        @javax.annotation.Nullable
        public java.util.List<RequestOption> options = Collections.emptyList();
        /** Request query parameters */
        @javax.annotation.Nullable
        public GetQueryParameters queryParameters = new GetQueryParameters();
        /**
         * Instantiates a new GetRequestConfiguration and sets the default values.
         * @return a void
         */
        @javax.annotation.Nullable
        public GetRequestConfiguration() {
        }
    }
}
