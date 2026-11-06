/*
Koyeb Rest API

Welcome to the Koyeb API documentation.  The Koyeb API allows you to interact with the Koyeb platform in a simple, programmatic way using conventional HTTP requests. The API has predictable resource-oriented URLs, allowing you to request, retrieve and execute actions on resources efficiently.  The API returns JSON-encoded responses and uses standard HTTP response codes, and verbs.  All of the functionality available in the Koyeb Control Panel is also available via the API, allowing you to automate repetitive and complex actions as required.  All production API requests are made to:  - https://app.koyeb.com/<version>  All data is sent and received as JSON.  ## Constructing Requests  Requests are made of three components:  - Base URL: `https://app.koyeb.com` - API Version: `v1` - Resource path: `apps`  To construct a proper request, you need to format the URL as follows:  `https://app.koyeb.com/{version}/{resource}`  Example: `https://app.koyeb.com/v1/apps`  The following code is an example request to retrieve the list of Apps.  ``` $ curl -H 'Authorization: Bearer 1111-1111-1111-1111-1111' https://app.koyeb.com/v1/apps {   \"apps\": [     {       \"id\": \"e267e09d-a2ef-4672-95c7-bf88d787a274\",       \"name\": \"demo\",       \"organization_id\": \"c402bcad-d06d-4fd6-8af2-08b790c0e249\",       \"updated_at\": \"2021-04-26T08:28:59.221274Z\",       \"created_at\": \"2021-04-26T08:28:59.221274Z\",       \"domains\": [         {           \"name\": \"demo-koyeb.koyeb.app\",           \"deployment_group\": \"prod\"         }       ]     },     {       \"id\": \"e38caea5-5acb-4fec-aa65-81749546b470\",       \"name\": \"www\",       \"organization_id\": \"c402bcad-d06d-4fd6-8af2-08b790c0e249\",       \"updated_at\": \"2021-05-04T08:56:40.557792Z\",       \"created_at\": \"2021-05-04T08:56:40.557792Z\",       \"domains\": [         {           \"name\": \"www-koyeb.koyeb.app\",           \"deployment_group\": \"prod\"         }       ]     }   ],   \"limit\": 20,   \"offset\": 0,   \"count\": 2 } ```  # Authentication  The Koyeb API uses API keys to authenticate requests. You can view and manage your API keys in the Koyeb Control Panel.  Your API keys carry many privileges. Take care to keep them secure! Do not share your secret API keys in publicly accessible areas such as GitHub, Slack, client-side code, and so forth.  To make an authenticated request, include a bearer-type Authorization header containing your API key, prefixed with `Bearer `:  ``` Authorization: Bearer <your-api-key> ```  All API requests must be made over HTTPS. Calls made over plain HTTP will fail. API requests without authentication will also fail.  # Errors  We use conventional HTTP response codes to indicate the success or failure of an API request.  In general, codes in the 2xx range indicate success, codes in the 4xx range indicate an error that resulted from the provided information (e.g., a required parameter is missing, resource not found), and codes in the 5xx range indicate an error with our servers.  ## HTTP Status Code Summary  - `200 OK` - Everything worked as expected. - `400 Bad Request` - Often missing a required parameter. - `401 Unauthorized` - No valid API key provided. - `402 Request Failed` - Parameters were valid, but the request failed. - `403 Forbidden` - Insufficient privileges to access the requested resource. - `404 Not Found` - The requested item doesn't exist. - `50x` Server errors - Something went wrong on the API domain's end.  Input and output data must be valid JSON with the appropriate Content-Type header set.  ## Attributes  - code:    - `invalid_argument`: Occur when your request has invalid parameters.   - `unknown_resource`: Occur when the resource doesn't exist.   - `authentication_error`: Arise when there is a problem with authentication.   - `rate_limit_error`: Too many requests hit the API too quickly.   - `authorization_required`: You don't have sufficient right to access the resource.   - `api_error`: API errors use in case of a problem with api domain's servers (extremely uncommon).  - message:    - A human-readable error giving more details about the error  - status:    - The HTTP code returned by the server.  - fields (Optional):    - An array of parameters with a human-readable message giving more details about the error.  ## Errors responses example  Response 400 (application/json)  ``` {   \"status\": 400,   \"code\": \"invalid_argument\",   \"message\": \"Invalid argument\",   \"fields\": [     {       \"field\": \"definition.docker.image\",       \"description\": \"cannot be blank\"     }   ] } ```  Response 401 (application/json)  ``` {   \"status\": 401,   \"code\": \"authentication_error\",   \"message\": \"Request unauthenticated with bearer\" } ```  Response 403 (application/json)  ``` {   \"status\": 403,   \"code\": \"authentication_required\",   \"message\": \"You don't have the requested permission to access to resource.\" } ```   Response 404 (application/json)  ``` {   \"status\": 404,   \"code\": \"unknown_resource\",   \"message\": \"App not found\" } ```  # Pagination  All top-level API resources receive a paginated response. For instance, list Apps, list Services, and others share a standard structure, taking two parameters: limit and offset.  **Paginated request**  Requests against paginated endpoints accept two query string arguments:  * *limit*, a positive integer lower or equal to 100 to select the number of items to display. * *offset*, an integer to indicate the position to start at.  Paginated endpoints usually also accept filters to search and sort results. These filters are documented along with each endpoint documentation. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package koyeb

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
)


type ServiceViewsApi interface {

	/*
	QueryServicesView Query the Services view

	Unstable: this endpoint may change without notice.

Returns Services with their App and Deployment state. Filters, search,
sorting, grouping, counts and facets are all computed on the full
authorized result before pagination.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiQueryServicesViewRequest
	*/
	QueryServicesView(ctx context.Context) ApiQueryServicesViewRequest

	// QueryServicesViewExecute executes the request
	//  @return QueryServicesViewReply
	QueryServicesViewExecute(r ApiQueryServicesViewRequest) (*QueryServicesViewReply, *http.Response, error)
}

// ServiceViewsApiService ServiceViewsApi service
type ServiceViewsApiService service

type ApiQueryServicesViewRequest struct {
	ctx context.Context
	ApiService ServiceViewsApi
	search *string
	filtersTypes *[]string
	filtersStatuses *[]string
	filtersRegions *[]string
	filtersAppIds *[]string
	filtersActiveDeploymentStates *[]string
	filtersLatestDeploymentStates *[]string
	filtersNeedsAttention *bool
	sortField *string
	sortDirection *string
	groupBy *string
	limit *string
	offset *string
	facets *[]string
}

// (Optional) Case-insensitive search on the Service name and the App name
func (r ApiQueryServicesViewRequest) Search(search string) ApiQueryServicesViewRequest {
	r.search = &search
	return r
}

// (Optional) Filter on Service types
func (r ApiQueryServicesViewRequest) FiltersTypes(filtersTypes []string) ApiQueryServicesViewRequest {
	r.filtersTypes = &filtersTypes
	return r
}

// (Optional) Filter on Service statuses
func (r ApiQueryServicesViewRequest) FiltersStatuses(filtersStatuses []string) ApiQueryServicesViewRequest {
	r.filtersStatuses = &filtersStatuses
	return r
}

// (Optional) Filter on regions of the latest Deployment
func (r ApiQueryServicesViewRequest) FiltersRegions(filtersRegions []string) ApiQueryServicesViewRequest {
	r.filtersRegions = &filtersRegions
	return r
}

// (Optional) Filter on App ids
func (r ApiQueryServicesViewRequest) FiltersAppIds(filtersAppIds []string) ApiQueryServicesViewRequest {
	r.filtersAppIds = &filtersAppIds
	return r
}

// (Optional) Filter on active Deployment states   - SLEEPING: The active Deployment is scaled to zero and wakes up on traffic
func (r ApiQueryServicesViewRequest) FiltersActiveDeploymentStates(filtersActiveDeploymentStates []string) ApiQueryServicesViewRequest {
	r.filtersActiveDeploymentStates = &filtersActiveDeploymentStates
	return r
}

// (Optional) Filter on latest Deployment states
func (r ApiQueryServicesViewRequest) FiltersLatestDeploymentStates(filtersLatestDeploymentStates []string) ApiQueryServicesViewRequest {
	r.filtersLatestDeploymentStates = &filtersLatestDeploymentStates
	return r
}

// (Optional) Filter on Services that need attention
func (r ApiQueryServicesViewRequest) FiltersNeedsAttention(filtersNeedsAttention bool) ApiQueryServicesViewRequest {
	r.filtersNeedsAttention = &filtersNeedsAttention
	return r
}

func (r ApiQueryServicesViewRequest) SortField(sortField string) ApiQueryServicesViewRequest {
	r.sortField = &sortField
	return r
}

func (r ApiQueryServicesViewRequest) SortDirection(sortDirection string) ApiQueryServicesViewRequest {
	r.sortDirection = &sortDirection
	return r
}

// (Optional) Group rows by App. Groups are ordered by App name and keep a stable order across pages. Pagination always counts Services
func (r ApiQueryServicesViewRequest) GroupBy(groupBy string) ApiQueryServicesViewRequest {
	r.groupBy = &groupBy
	return r
}

// (Optional) The number of items to return
func (r ApiQueryServicesViewRequest) Limit(limit string) ApiQueryServicesViewRequest {
	r.limit = &limit
	return r
}

// (Optional) The offset in the list of item to return
func (r ApiQueryServicesViewRequest) Offset(offset string) ApiQueryServicesViewRequest {
	r.offset = &offset
	return r
}

// (Optional) Facets to compute on the full result
func (r ApiQueryServicesViewRequest) Facets(facets []string) ApiQueryServicesViewRequest {
	r.facets = &facets
	return r
}

func (r ApiQueryServicesViewRequest) Execute() (*QueryServicesViewReply, *http.Response, error) {
	return r.ApiService.QueryServicesViewExecute(r)
}

/*
QueryServicesView Query the Services view

Unstable: this endpoint may change without notice.

Returns Services with their App and Deployment state. Filters, search,
sorting, grouping, counts and facets are all computed on the full
authorized result before pagination.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiQueryServicesViewRequest
*/
func (a *ServiceViewsApiService) QueryServicesView(ctx context.Context) ApiQueryServicesViewRequest {
	return ApiQueryServicesViewRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return QueryServicesViewReply
func (a *ServiceViewsApiService) QueryServicesViewExecute(r ApiQueryServicesViewRequest) (*QueryServicesViewReply, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *QueryServicesViewReply
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServiceViewsApiService.QueryServicesView")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/views/services"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.search != nil {
		localVarQueryParams.Add("search", parameterToString(*r.search, ""))
	}
	if r.filtersTypes != nil {
		t := *r.filtersTypes
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("filters.types", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("filters.types", parameterToString(t, "multi"))
		}
	}
	if r.filtersStatuses != nil {
		t := *r.filtersStatuses
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("filters.statuses", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("filters.statuses", parameterToString(t, "multi"))
		}
	}
	if r.filtersRegions != nil {
		t := *r.filtersRegions
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("filters.regions", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("filters.regions", parameterToString(t, "multi"))
		}
	}
	if r.filtersAppIds != nil {
		t := *r.filtersAppIds
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("filters.app_ids", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("filters.app_ids", parameterToString(t, "multi"))
		}
	}
	if r.filtersActiveDeploymentStates != nil {
		t := *r.filtersActiveDeploymentStates
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("filters.active_deployment_states", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("filters.active_deployment_states", parameterToString(t, "multi"))
		}
	}
	if r.filtersLatestDeploymentStates != nil {
		t := *r.filtersLatestDeploymentStates
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("filters.latest_deployment_states", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("filters.latest_deployment_states", parameterToString(t, "multi"))
		}
	}
	if r.filtersNeedsAttention != nil {
		localVarQueryParams.Add("filters.needs_attention", parameterToString(*r.filtersNeedsAttention, ""))
	}
	if r.sortField != nil {
		localVarQueryParams.Add("sort.field", parameterToString(*r.sortField, ""))
	}
	if r.sortDirection != nil {
		localVarQueryParams.Add("sort.direction", parameterToString(*r.sortDirection, ""))
	}
	if r.groupBy != nil {
		localVarQueryParams.Add("group_by", parameterToString(*r.groupBy, ""))
	}
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.offset != nil {
		localVarQueryParams.Add("offset", parameterToString(*r.offset, ""))
	}
	if r.facets != nil {
		t := *r.facets
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("facets", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("facets", parameterToString(t, "multi"))
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
	localVarHTTPHeaderAccepts := []string{"*/*"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["Bearer"]; ok {
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

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorWithFields
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
			var v GoogleRpcStatus
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
            		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
            		newErr.model = v
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
