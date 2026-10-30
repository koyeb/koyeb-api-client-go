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
	"time"
)


type LogsApi interface {

	/*
	QueryLogs Query logs

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiQueryLogsRequest
	*/
	QueryLogs(ctx context.Context) ApiQueryLogsRequest

	// QueryLogsExecute executes the request
	//  @return QueryLogsReply
	QueryLogsExecute(r ApiQueryLogsRequest) (*QueryLogsReply, *http.Response, error)

	/*
	TailLogs Tails logs

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiTailLogsRequest
	*/
	TailLogs(ctx context.Context) ApiTailLogsRequest

	// TailLogsExecute executes the request
	//  @return StreamResultOfLogEntry
	TailLogsExecute(r ApiTailLogsRequest) (*StreamResultOfLogEntry, *http.Response, error)
}

// LogsApiService LogsApi service
type LogsApiService service

type ApiQueryLogsRequest struct {
	ctx context.Context
	ApiService LogsApi
	type_ *string
	appId *string
	serviceId *string
	deploymentId *string
	regionalDeploymentId *string
	instanceId *string
	instanceIds *[]string
	stream *string
	streams *[]string
	start *time.Time
	end *time.Time
	order *string
	limit *string
	regex *string
	text *string
	regions *[]string
}

// Type of logs to retrieve, either \&quot;build\&quot; or \&quot;runtime\&quot;. Defaults to \&quot;runtime\&quot;.
func (r ApiQueryLogsRequest) Type_(type_ string) ApiQueryLogsRequest {
	r.type_ = &type_
	return r
}

// (Optional) Filter on the provided app_id. At least one of app_id, service_id, deployment_id, regional_deployment_id or instance_ids must be set.
func (r ApiQueryLogsRequest) AppId(appId string) ApiQueryLogsRequest {
	r.appId = &appId
	return r
}

// (Optional) Filter on the provided service_id. At least one of app_id, service_id, deployment_id, regional_deployment_id or instance_ids must be set.
func (r ApiQueryLogsRequest) ServiceId(serviceId string) ApiQueryLogsRequest {
	r.serviceId = &serviceId
	return r
}

// (Optional) Filter on the provided deployment_id. At least one of app_id, service_id, deployment_id, regional_deployment_id or instance_ids must be set.
func (r ApiQueryLogsRequest) DeploymentId(deploymentId string) ApiQueryLogsRequest {
	r.deploymentId = &deploymentId
	return r
}

// (Optional) Filter on the provided regional_deployment_id. At least one of app_id, service_id, deployment_id, regional_deployment_id or instance_ids must be set.
func (r ApiQueryLogsRequest) RegionalDeploymentId(regionalDeploymentId string) ApiQueryLogsRequest {
	r.regionalDeploymentId = &regionalDeploymentId
	return r
}

// Deprecated, prefer using instance_ids instead.
func (r ApiQueryLogsRequest) InstanceId(instanceId string) ApiQueryLogsRequest {
	r.instanceId = &instanceId
	return r
}

// (Optional) Filter on the provided instance_ids. At least one of app_id, service_id, deployment_id, regional_deployment_id or instance_ids must be set.
func (r ApiQueryLogsRequest) InstanceIds(instanceIds []string) ApiQueryLogsRequest {
	r.instanceIds = &instanceIds
	return r
}

// Deprecated, prefer using streams instead.
func (r ApiQueryLogsRequest) Stream(stream string) ApiQueryLogsRequest {
	r.stream = &stream
	return r
}

// (Optional) Filter on stream: either \&quot;stdout\&quot;, \&quot;stderr\&quot; or \&quot;koyeb\&quot; (for system logs).
func (r ApiQueryLogsRequest) Streams(streams []string) ApiQueryLogsRequest {
	r.streams = &streams
	return r
}

// (Optional) Must always be before &#x60;end&#x60;. Defaults to 15 minutes ago.
func (r ApiQueryLogsRequest) Start(start time.Time) ApiQueryLogsRequest {
	r.start = &start
	return r
}

// (Optional) Must always be after &#x60;start&#x60;. Defaults to now.
func (r ApiQueryLogsRequest) End(end time.Time) ApiQueryLogsRequest {
	r.end = &end
	return r
}

// (Optional) &#x60;asc&#x60; or &#x60;desc&#x60;. Defaults to &#x60;desc&#x60;.
func (r ApiQueryLogsRequest) Order(order string) ApiQueryLogsRequest {
	r.order = &order
	return r
}

// (Optional) Defaults to 100. Maximum of 1000.
func (r ApiQueryLogsRequest) Limit(limit string) ApiQueryLogsRequest {
	r.limit = &limit
	return r
}

// (Optional) Apply a regex to filter logs. Can&#39;t be used with &#x60;text&#x60;.
func (r ApiQueryLogsRequest) Regex(regex string) ApiQueryLogsRequest {
	r.regex = &regex
	return r
}

// (Optional) Looks for this string in logs. Can&#39;t be used with &#x60;regex&#x60;.
func (r ApiQueryLogsRequest) Text(text string) ApiQueryLogsRequest {
	r.text = &text
	return r
}

// (Optional) Filter on the provided regions (e.g. [\&quot;fra\&quot;, \&quot;was\&quot;]).
func (r ApiQueryLogsRequest) Regions(regions []string) ApiQueryLogsRequest {
	r.regions = &regions
	return r
}

func (r ApiQueryLogsRequest) Execute() (*QueryLogsReply, *http.Response, error) {
	return r.ApiService.QueryLogsExecute(r)
}

/*
QueryLogs Query logs

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiQueryLogsRequest
*/
func (a *LogsApiService) QueryLogs(ctx context.Context) ApiQueryLogsRequest {
	return ApiQueryLogsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return QueryLogsReply
func (a *LogsApiService) QueryLogsExecute(r ApiQueryLogsRequest) (*QueryLogsReply, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *QueryLogsReply
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LogsApiService.QueryLogs")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/streams/logs/query"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.type_ != nil {
		localVarQueryParams.Add("type", parameterToString(*r.type_, ""))
	}
	if r.appId != nil {
		localVarQueryParams.Add("app_id", parameterToString(*r.appId, ""))
	}
	if r.serviceId != nil {
		localVarQueryParams.Add("service_id", parameterToString(*r.serviceId, ""))
	}
	if r.deploymentId != nil {
		localVarQueryParams.Add("deployment_id", parameterToString(*r.deploymentId, ""))
	}
	if r.regionalDeploymentId != nil {
		localVarQueryParams.Add("regional_deployment_id", parameterToString(*r.regionalDeploymentId, ""))
	}
	if r.instanceId != nil {
		localVarQueryParams.Add("instance_id", parameterToString(*r.instanceId, ""))
	}
	if r.instanceIds != nil {
		t := *r.instanceIds
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("instance_ids", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("instance_ids", parameterToString(t, "multi"))
		}
	}
	if r.stream != nil {
		localVarQueryParams.Add("stream", parameterToString(*r.stream, ""))
	}
	if r.streams != nil {
		t := *r.streams
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("streams", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("streams", parameterToString(t, "multi"))
		}
	}
	if r.start != nil {
		localVarQueryParams.Add("start", parameterToString(*r.start, ""))
	}
	if r.end != nil {
		localVarQueryParams.Add("end", parameterToString(*r.end, ""))
	}
	if r.order != nil {
		localVarQueryParams.Add("order", parameterToString(*r.order, ""))
	}
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.regex != nil {
		localVarQueryParams.Add("regex", parameterToString(*r.regex, ""))
	}
	if r.text != nil {
		localVarQueryParams.Add("text", parameterToString(*r.text, ""))
	}
	if r.regions != nil {
		t := *r.regions
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("regions", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("regions", parameterToString(t, "multi"))
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

type ApiTailLogsRequest struct {
	ctx context.Context
	ApiService LogsApi
	type_ *string
	appId *string
	serviceId *string
	deploymentId *string
	regionalDeploymentId *string
	instanceId *string
	instanceIds *[]string
	stream *string
	streams *[]string
	start *time.Time
	limit *string
	regex *string
	text *string
	regions *[]string
}

// Type of logs to retrieve, either \&quot;build\&quot; or \&quot;runtime\&quot;. Defaults to \&quot;runtime\&quot;.
func (r ApiTailLogsRequest) Type_(type_ string) ApiTailLogsRequest {
	r.type_ = &type_
	return r
}

// (Optional) Filter on the provided app_id. At least one of app_id, service_id, deployment_id, regional_deployment_id or instance_ids must be set.
func (r ApiTailLogsRequest) AppId(appId string) ApiTailLogsRequest {
	r.appId = &appId
	return r
}

// (Optional) Filter on the provided service_id. At least one of app_id, service_id, deployment_id, regional_deployment_id or instance_ids must be set.
func (r ApiTailLogsRequest) ServiceId(serviceId string) ApiTailLogsRequest {
	r.serviceId = &serviceId
	return r
}

// (Optional) Filter on the provided deployment_id. At least one of app_id, service_id, deployment_id, regional_deployment_id or instance_ids must be set.
func (r ApiTailLogsRequest) DeploymentId(deploymentId string) ApiTailLogsRequest {
	r.deploymentId = &deploymentId
	return r
}

// (Optional) Filter on the provided regional_deployment_id. At least one of app_id, service_id, deployment_id, regional_deployment_id or instance_ids must be set.
func (r ApiTailLogsRequest) RegionalDeploymentId(regionalDeploymentId string) ApiTailLogsRequest {
	r.regionalDeploymentId = &regionalDeploymentId
	return r
}

// Deprecated, prefer using instance_ids instead.
func (r ApiTailLogsRequest) InstanceId(instanceId string) ApiTailLogsRequest {
	r.instanceId = &instanceId
	return r
}

// (Optional) Filter on the provided instance_ids. At least one of app_id, service_id, deployment_id, regional_deployment_id or instance_ids must be set.
func (r ApiTailLogsRequest) InstanceIds(instanceIds []string) ApiTailLogsRequest {
	r.instanceIds = &instanceIds
	return r
}

// Deprecated, prefer using streams instead.
func (r ApiTailLogsRequest) Stream(stream string) ApiTailLogsRequest {
	r.stream = &stream
	return r
}

// (Optional) Filter on stream: either \&quot;stdout\&quot;, \&quot;stderr\&quot; or \&quot;koyeb\&quot; (for system logs).
func (r ApiTailLogsRequest) Streams(streams []string) ApiTailLogsRequest {
	r.streams = &streams
	return r
}

// (Optional) Defaults to 24 hours ago.
func (r ApiTailLogsRequest) Start(start time.Time) ApiTailLogsRequest {
	r.start = &start
	return r
}

// (Optional) Defaults to 1000. Maximum of 1000.
func (r ApiTailLogsRequest) Limit(limit string) ApiTailLogsRequest {
	r.limit = &limit
	return r
}

// (Optional) Apply a regex to filter logs. Can&#39;t be used with &#x60;text&#x60;.
func (r ApiTailLogsRequest) Regex(regex string) ApiTailLogsRequest {
	r.regex = &regex
	return r
}

// (Optional) Looks for this string in logs. Can&#39;t be used with &#x60;regex&#x60;.
func (r ApiTailLogsRequest) Text(text string) ApiTailLogsRequest {
	r.text = &text
	return r
}

// (Optional) Filter on the provided regions (e.g. [\&quot;fra\&quot;, \&quot;was\&quot;]).
func (r ApiTailLogsRequest) Regions(regions []string) ApiTailLogsRequest {
	r.regions = &regions
	return r
}

func (r ApiTailLogsRequest) Execute() (*StreamResultOfLogEntry, *http.Response, error) {
	return r.ApiService.TailLogsExecute(r)
}

/*
TailLogs Tails logs

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiTailLogsRequest
*/
func (a *LogsApiService) TailLogs(ctx context.Context) ApiTailLogsRequest {
	return ApiTailLogsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return StreamResultOfLogEntry
func (a *LogsApiService) TailLogsExecute(r ApiTailLogsRequest) (*StreamResultOfLogEntry, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *StreamResultOfLogEntry
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LogsApiService.TailLogs")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/streams/logs/tail"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.type_ != nil {
		localVarQueryParams.Add("type", parameterToString(*r.type_, ""))
	}
	if r.appId != nil {
		localVarQueryParams.Add("app_id", parameterToString(*r.appId, ""))
	}
	if r.serviceId != nil {
		localVarQueryParams.Add("service_id", parameterToString(*r.serviceId, ""))
	}
	if r.deploymentId != nil {
		localVarQueryParams.Add("deployment_id", parameterToString(*r.deploymentId, ""))
	}
	if r.regionalDeploymentId != nil {
		localVarQueryParams.Add("regional_deployment_id", parameterToString(*r.regionalDeploymentId, ""))
	}
	if r.instanceId != nil {
		localVarQueryParams.Add("instance_id", parameterToString(*r.instanceId, ""))
	}
	if r.instanceIds != nil {
		t := *r.instanceIds
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("instance_ids", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("instance_ids", parameterToString(t, "multi"))
		}
	}
	if r.stream != nil {
		localVarQueryParams.Add("stream", parameterToString(*r.stream, ""))
	}
	if r.streams != nil {
		t := *r.streams
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("streams", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("streams", parameterToString(t, "multi"))
		}
	}
	if r.start != nil {
		localVarQueryParams.Add("start", parameterToString(*r.start, ""))
	}
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.regex != nil {
		localVarQueryParams.Add("regex", parameterToString(*r.regex, ""))
	}
	if r.text != nil {
		localVarQueryParams.Add("text", parameterToString(*r.text, ""))
	}
	if r.regions != nil {
		t := *r.regions
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("regions", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("regions", parameterToString(t, "multi"))
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
