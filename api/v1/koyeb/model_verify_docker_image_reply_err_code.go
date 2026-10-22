/*
Koyeb Rest API

Welcome to the Koyeb API documentation.  The Koyeb API allows you to interact with the Koyeb platform in a simple, programmatic way using conventional HTTP requests. The API has predictable resource-oriented URLs, allowing you to request, retrieve and execute actions on resources efficiently.  The API returns JSON-encoded responses and uses standard HTTP response codes, and verbs.  All of the functionality available in the Koyeb Control Panel is also available via the API, allowing you to automate repetitive and complex actions as required.  All production API requests are made to:  - https://app.koyeb.com/<version>  All data is sent and received as JSON.  ## Constructing Requests  Requests are made of three components:  - Base URL: `https://app.koyeb.com` - API Version: `v1` - Resource path: `apps`  To construct a proper request, you need to format the URL as follows:  `https://app.koyeb.com/{version}/{resource}`  Example: `https://app.koyeb.com/v1/apps`  The following code is an example request to retrieve the list of Apps.  ``` $ curl -H 'Authorization: Bearer 1111-1111-1111-1111-1111' https://app.koyeb.com/v1/apps {   \"apps\": [     {       \"id\": \"e267e09d-a2ef-4672-95c7-bf88d787a274\",       \"name\": \"demo\",       \"organization_id\": \"c402bcad-d06d-4fd6-8af2-08b790c0e249\",       \"updated_at\": \"2021-04-26T08:28:59.221274Z\",       \"created_at\": \"2021-04-26T08:28:59.221274Z\",       \"domains\": [         {           \"name\": \"demo-koyeb.koyeb.app\",           \"deployment_group\": \"prod\"         }       ]     },     {       \"id\": \"e38caea5-5acb-4fec-aa65-81749546b470\",       \"name\": \"www\",       \"organization_id\": \"c402bcad-d06d-4fd6-8af2-08b790c0e249\",       \"updated_at\": \"2021-05-04T08:56:40.557792Z\",       \"created_at\": \"2021-05-04T08:56:40.557792Z\",       \"domains\": [         {           \"name\": \"www-koyeb.koyeb.app\",           \"deployment_group\": \"prod\"         }       ]     }   ],   \"limit\": 20,   \"offset\": 0,   \"count\": 2 } ```  # Authentication  The Koyeb API uses API keys to authenticate requests. You can view and manage your API keys in the Koyeb Control Panel.  Your API keys carry many privileges. Take care to keep them secure! Do not share your secret API keys in publicly accessible areas such as GitHub, Slack, client-side code, and so forth.  To make an authenticated request, include a bearer-type Authorization header containing your API key, prefixed with `Bearer `:  ``` Authorization: Bearer <your-api-key> ```  All API requests must be made over HTTPS. Calls made over plain HTTP will fail. API requests without authentication will also fail.  # Errors  We use conventional HTTP response codes to indicate the success or failure of an API request.  In general, codes in the 2xx range indicate success, codes in the 4xx range indicate an error that resulted from the provided information (e.g., a required parameter is missing, resource not found), and codes in the 5xx range indicate an error with our servers.  ## HTTP Status Code Summary  - `200 OK` - Everything worked as expected. - `400 Bad Request` - Often missing a required parameter. - `401 Unauthorized` - No valid API key provided. - `402 Request Failed` - Parameters were valid, but the request failed. - `403 Forbidden` - Insufficient privileges to access the requested resource. - `404 Not Found` - The requested item doesn't exist. - `50x` Server errors - Something went wrong on the API domain's end.  Input and output data must be valid JSON with the appropriate Content-Type header set.  ## Attributes  - code:    - `invalid_argument`: Occur when your request has invalid parameters.   - `unknown_resource`: Occur when the resource doesn't exist.   - `authentication_error`: Arise when there is a problem with authentication.   - `rate_limit_error`: Too many requests hit the API too quickly.   - `authorization_required`: You don't have sufficient right to access the resource.   - `api_error`: API errors use in case of a problem with api domain's servers (extremely uncommon).  - message:    - A human-readable error giving more details about the error  - status:    - The HTTP code returned by the server.  - fields (Optional):    - An array of parameters with a human-readable message giving more details about the error.  ## Errors responses example  Response 400 (application/json)  ``` {   \"status\": 400,   \"code\": \"invalid_argument\",   \"message\": \"Invalid argument\",   \"fields\": [     {       \"field\": \"definition.docker.image\",       \"description\": \"cannot be blank\"     }   ] } ```  Response 401 (application/json)  ``` {   \"status\": 401,   \"code\": \"authentication_error\",   \"message\": \"Request unauthenticated with bearer\" } ```  Response 403 (application/json)  ``` {   \"status\": 403,   \"code\": \"authentication_required\",   \"message\": \"You don't have the requested permission to access to resource.\" } ```   Response 404 (application/json)  ``` {   \"status\": 404,   \"code\": \"unknown_resource\",   \"message\": \"App not found\" } ```  # Pagination  All top-level API resources receive a paginated response. For instance, list Apps, list Services, and others share a standard structure, taking two parameters: limit and offset.  **Paginated request**  Requests against paginated endpoints accept two query string arguments:  * *limit*, a positive integer lower or equal to 100 to select the number of items to display. * *offset*, an integer to indicate the position to start at.  Paginated endpoints usually also accept filters to search and sort results. These filters are documented along with each endpoint documentation. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package koyeb

import (
	"encoding/json"
	"fmt"
)

// VerifyDockerImageReplyErrCode - UNKNOWN: Default value  - AUTH_ACCESS_DENIED: The registry denied access to an authenticated request  - ANON_ACCESS_DENIED: The registry denied access to an anonymous request  - AUTH_NOT_FOUND: The image has not been found after an authenticated request  - ANON_NOT_FOUND: The image has not been found after an anonymous request  - REGISTRY_ERROR: The registry returned an error  - TIMEOUT: The request to the registry timed out  - DNS: There was an error trying to resolve the name of the registry  - MALFORMED: The provided image name is malformed  - INVALID_OS: The operating system is not supported  - INVALID_ARCH: The architecture is not supported  - INVALID_SCHEME: The scheme is not https  - GENERIC: Generic catch-all error code
type VerifyDockerImageReplyErrCode string

// List of VerifyDockerImageReply.ErrCode
const (
	VERIFYDOCKERIMAGEREPLYERRCODE_UNKNOWN VerifyDockerImageReplyErrCode = "UNKNOWN"
	VERIFYDOCKERIMAGEREPLYERRCODE_AUTH_ACCESS_DENIED VerifyDockerImageReplyErrCode = "AUTH_ACCESS_DENIED"
	VERIFYDOCKERIMAGEREPLYERRCODE_ANON_ACCESS_DENIED VerifyDockerImageReplyErrCode = "ANON_ACCESS_DENIED"
	VERIFYDOCKERIMAGEREPLYERRCODE_AUTH_NOT_FOUND VerifyDockerImageReplyErrCode = "AUTH_NOT_FOUND"
	VERIFYDOCKERIMAGEREPLYERRCODE_ANON_NOT_FOUND VerifyDockerImageReplyErrCode = "ANON_NOT_FOUND"
	VERIFYDOCKERIMAGEREPLYERRCODE_REGISTRY_ERROR VerifyDockerImageReplyErrCode = "REGISTRY_ERROR"
	VERIFYDOCKERIMAGEREPLYERRCODE_TIMEOUT VerifyDockerImageReplyErrCode = "TIMEOUT"
	VERIFYDOCKERIMAGEREPLYERRCODE_DNS VerifyDockerImageReplyErrCode = "DNS"
	VERIFYDOCKERIMAGEREPLYERRCODE_MALFORMED VerifyDockerImageReplyErrCode = "MALFORMED"
	VERIFYDOCKERIMAGEREPLYERRCODE_INVALID_OS VerifyDockerImageReplyErrCode = "INVALID_OS"
	VERIFYDOCKERIMAGEREPLYERRCODE_INVALID_ARCH VerifyDockerImageReplyErrCode = "INVALID_ARCH"
	VERIFYDOCKERIMAGEREPLYERRCODE_INVALID_SCHEME VerifyDockerImageReplyErrCode = "INVALID_SCHEME"
	VERIFYDOCKERIMAGEREPLYERRCODE_GENERIC VerifyDockerImageReplyErrCode = "GENERIC"
)

// All allowed values of VerifyDockerImageReplyErrCode enum
var AllowedVerifyDockerImageReplyErrCodeEnumValues = []VerifyDockerImageReplyErrCode{
	"UNKNOWN",
	"AUTH_ACCESS_DENIED",
	"ANON_ACCESS_DENIED",
	"AUTH_NOT_FOUND",
	"ANON_NOT_FOUND",
	"REGISTRY_ERROR",
	"TIMEOUT",
	"DNS",
	"MALFORMED",
	"INVALID_OS",
	"INVALID_ARCH",
	"INVALID_SCHEME",
	"GENERIC",
}

func (v *VerifyDockerImageReplyErrCode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := VerifyDockerImageReplyErrCode(value)
	for _, existing := range AllowedVerifyDockerImageReplyErrCodeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid VerifyDockerImageReplyErrCode", value)
}

// NewVerifyDockerImageReplyErrCodeFromValue returns a pointer to a valid VerifyDockerImageReplyErrCode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewVerifyDockerImageReplyErrCodeFromValue(v string) (*VerifyDockerImageReplyErrCode, error) {
	ev := VerifyDockerImageReplyErrCode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for VerifyDockerImageReplyErrCode: valid values are %v", v, AllowedVerifyDockerImageReplyErrCodeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v VerifyDockerImageReplyErrCode) IsValid() bool {
	for _, existing := range AllowedVerifyDockerImageReplyErrCodeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to VerifyDockerImageReply.ErrCode value
func (v VerifyDockerImageReplyErrCode) Ptr() *VerifyDockerImageReplyErrCode {
	return &v
}

type NullableVerifyDockerImageReplyErrCode struct {
	value *VerifyDockerImageReplyErrCode
	isSet bool
}

func (v NullableVerifyDockerImageReplyErrCode) Get() *VerifyDockerImageReplyErrCode {
	return v.value
}

func (v *NullableVerifyDockerImageReplyErrCode) Set(val *VerifyDockerImageReplyErrCode) {
	v.value = val
	v.isSet = true
}

func (v NullableVerifyDockerImageReplyErrCode) IsSet() bool {
	return v.isSet
}

func (v *NullableVerifyDockerImageReplyErrCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVerifyDockerImageReplyErrCode(val *VerifyDockerImageReplyErrCode) *NullableVerifyDockerImageReplyErrCode {
	return &NullableVerifyDockerImageReplyErrCode{value: val, isSet: true}
}

func (v NullableVerifyDockerImageReplyErrCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVerifyDockerImageReplyErrCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

