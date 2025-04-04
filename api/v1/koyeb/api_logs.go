/*
Koyeb Rest API

The Koyeb API allows you to interact with the Koyeb platform in a simple, programmatic way using conventional HTTP requests. 

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
	instanceId *string
	stream *string
	regionalDeploymentId *string
	start *time.Time
	end *time.Time
	order *string
	limit *string
	regex *string
	text *string
}

func (r ApiQueryLogsRequest) Type_(type_ string) ApiQueryLogsRequest {
	r.type_ = &type_
	return r
}

func (r ApiQueryLogsRequest) AppId(appId string) ApiQueryLogsRequest {
	r.appId = &appId
	return r
}

func (r ApiQueryLogsRequest) ServiceId(serviceId string) ApiQueryLogsRequest {
	r.serviceId = &serviceId
	return r
}

func (r ApiQueryLogsRequest) DeploymentId(deploymentId string) ApiQueryLogsRequest {
	r.deploymentId = &deploymentId
	return r
}

func (r ApiQueryLogsRequest) InstanceId(instanceId string) ApiQueryLogsRequest {
	r.instanceId = &instanceId
	return r
}

func (r ApiQueryLogsRequest) Stream(stream string) ApiQueryLogsRequest {
	r.stream = &stream
	return r
}

func (r ApiQueryLogsRequest) RegionalDeploymentId(regionalDeploymentId string) ApiQueryLogsRequest {
	r.regionalDeploymentId = &regionalDeploymentId
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
	if r.instanceId != nil {
		localVarQueryParams.Add("instance_id", parameterToString(*r.instanceId, ""))
	}
	if r.stream != nil {
		localVarQueryParams.Add("stream", parameterToString(*r.stream, ""))
	}
	if r.regionalDeploymentId != nil {
		localVarQueryParams.Add("regional_deployment_id", parameterToString(*r.regionalDeploymentId, ""))
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
	stream *string
	start *time.Time
	limit *string
	regex *string
	text *string
}

func (r ApiTailLogsRequest) Type_(type_ string) ApiTailLogsRequest {
	r.type_ = &type_
	return r
}

func (r ApiTailLogsRequest) AppId(appId string) ApiTailLogsRequest {
	r.appId = &appId
	return r
}

func (r ApiTailLogsRequest) ServiceId(serviceId string) ApiTailLogsRequest {
	r.serviceId = &serviceId
	return r
}

func (r ApiTailLogsRequest) DeploymentId(deploymentId string) ApiTailLogsRequest {
	r.deploymentId = &deploymentId
	return r
}

func (r ApiTailLogsRequest) RegionalDeploymentId(regionalDeploymentId string) ApiTailLogsRequest {
	r.regionalDeploymentId = &regionalDeploymentId
	return r
}

func (r ApiTailLogsRequest) InstanceId(instanceId string) ApiTailLogsRequest {
	r.instanceId = &instanceId
	return r
}

func (r ApiTailLogsRequest) Stream(stream string) ApiTailLogsRequest {
	r.stream = &stream
	return r
}

func (r ApiTailLogsRequest) Start(start time.Time) ApiTailLogsRequest {
	r.start = &start
	return r
}

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
	if r.stream != nil {
		localVarQueryParams.Add("stream", parameterToString(*r.stream, ""))
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
