/*
 * Koyeb Rest API
 *
 * The Koyeb API allows you to interact with the Koyeb platform in a simple, programmatic way using conventional HTTP requests. 
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package koyeb

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

type InstancesApi interface {

	/*
	 * ExecCommand Exec Command
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @return ApiExecCommandRequest
	 */
	ExecCommand(ctx _context.Context) ApiExecCommandRequest

	/*
	 * ExecCommandExecute executes the request
	 * @return StreamResultOfExecCommandReply
	 */
	ExecCommandExecute(r ApiExecCommandRequest) (StreamResultOfExecCommandReply, *_nethttp.Response, error)

	/*
	 * ListServiceInstances List Instances for a service
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @param appIdOrName Name or id of the application
	 * @param serviceIdOrName Name or id of the service
	 * @return ApiListServiceInstancesRequest
	 */
	ListServiceInstances(ctx _context.Context, appIdOrName string, serviceIdOrName string) ApiListServiceInstancesRequest

	/*
	 * ListServiceInstancesExecute executes the request
	 * @return ListServiceInstancesReply
	 */
	ListServiceInstancesExecute(r ApiListServiceInstancesRequest) (ListServiceInstancesReply, *_nethttp.Response, error)
}

// InstancesApiService InstancesApi service
type InstancesApiService service

type ApiExecCommandRequest struct {
	ctx _context.Context
	ApiService InstancesApi
	id *string
	bodyCommand *[]string
	bodyTtySizeHeight *int32
	bodyTtySizeWidth *int32
	bodyStdinData *string
}

func (r ApiExecCommandRequest) Id(id string) ApiExecCommandRequest {
	r.id = &id
	return r
}
func (r ApiExecCommandRequest) BodyCommand(bodyCommand []string) ApiExecCommandRequest {
	r.bodyCommand = &bodyCommand
	return r
}
func (r ApiExecCommandRequest) BodyTtySizeHeight(bodyTtySizeHeight int32) ApiExecCommandRequest {
	r.bodyTtySizeHeight = &bodyTtySizeHeight
	return r
}
func (r ApiExecCommandRequest) BodyTtySizeWidth(bodyTtySizeWidth int32) ApiExecCommandRequest {
	r.bodyTtySizeWidth = &bodyTtySizeWidth
	return r
}
func (r ApiExecCommandRequest) BodyStdinData(bodyStdinData string) ApiExecCommandRequest {
	r.bodyStdinData = &bodyStdinData
	return r
}

func (r ApiExecCommandRequest) Execute() (StreamResultOfExecCommandReply, *_nethttp.Response, error) {
	return r.ApiService.ExecCommandExecute(r)
}

/*
 * ExecCommand Exec Command
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiExecCommandRequest
 */
func (a *InstancesApiService) ExecCommand(ctx _context.Context) ApiExecCommandRequest {
	return ApiExecCommandRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return StreamResultOfExecCommandReply
 */
func (a *InstancesApiService) ExecCommandExecute(r ApiExecCommandRequest) (StreamResultOfExecCommandReply, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  StreamResultOfExecCommandReply
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InstancesApiService.ExecCommand")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/instances/exec"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.id != nil {
		localVarQueryParams.Add("id", parameterToString(*r.id, ""))
	}
	if r.bodyCommand != nil {
		t := *r.bodyCommand
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("body.command", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("body.command", parameterToString(t, "multi"))
		}
	}
	if r.bodyTtySizeHeight != nil {
		localVarQueryParams.Add("body.tty_size.height", parameterToString(*r.bodyTtySizeHeight, ""))
	}
	if r.bodyTtySizeWidth != nil {
		localVarQueryParams.Add("body.tty_size.width", parameterToString(*r.bodyTtySizeWidth, ""))
	}
	if r.bodyStdinData != nil {
		localVarQueryParams.Add("body.stdin.data", parameterToString(*r.bodyStdinData, ""))
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
			var v ErrorWithFields
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
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
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
			var v GoogleRpcStatus
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
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

type ApiListServiceInstancesRequest struct {
	ctx _context.Context
	ApiService InstancesApi
	appIdOrName string
	serviceIdOrName string
	limit *string
	offset *string
}

func (r ApiListServiceInstancesRequest) Limit(limit string) ApiListServiceInstancesRequest {
	r.limit = &limit
	return r
}
func (r ApiListServiceInstancesRequest) Offset(offset string) ApiListServiceInstancesRequest {
	r.offset = &offset
	return r
}

func (r ApiListServiceInstancesRequest) Execute() (ListServiceInstancesReply, *_nethttp.Response, error) {
	return r.ApiService.ListServiceInstancesExecute(r)
}

/*
 * ListServiceInstances List Instances for a service
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param appIdOrName Name or id of the application
 * @param serviceIdOrName Name or id of the service
 * @return ApiListServiceInstancesRequest
 */
func (a *InstancesApiService) ListServiceInstances(ctx _context.Context, appIdOrName string, serviceIdOrName string) ApiListServiceInstancesRequest {
	return ApiListServiceInstancesRequest{
		ApiService: a,
		ctx: ctx,
		appIdOrName: appIdOrName,
		serviceIdOrName: serviceIdOrName,
	}
}

/*
 * Execute executes the request
 * @return ListServiceInstancesReply
 */
func (a *InstancesApiService) ListServiceInstancesExecute(r ApiListServiceInstancesRequest) (ListServiceInstancesReply, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ListServiceInstancesReply
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InstancesApiService.ListServiceInstances")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/apps/{app_id_or_name}/services/{service_id_or_name}/instances"
	localVarPath = strings.Replace(localVarPath, "{"+"app_id_or_name"+"}", _neturl.PathEscape(parameterToString(r.appIdOrName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"service_id_or_name"+"}", _neturl.PathEscape(parameterToString(r.serviceIdOrName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.offset != nil {
		localVarQueryParams.Add("offset", parameterToString(*r.offset, ""))
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
			var v ErrorWithFields
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
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
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
			var v GoogleRpcStatus
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
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
