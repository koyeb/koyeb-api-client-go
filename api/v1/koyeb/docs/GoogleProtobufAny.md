# GoogleProtobufAny

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TypeUrl** | Pointer to **string** | A URL/resource name that uniquely identifies the type of the serialized protocol buffer message. This string must contain at least one \&quot;/\&quot; character. The last segment of the URL&#39;s path must represent the fully qualified name of the type (as in &#x60;path/google.protobuf.Duration&#x60;). The name should be in a canonical form (e.g., leading \&quot;.\&quot; is not accepted).  In practice, teams usually precompile into the binary all types that they expect it to use in the context of Any. However, for URLs which use the scheme &#x60;http&#x60;, &#x60;https&#x60;, or no scheme, one can optionally set up a type server that maps type URLs to message definitions as follows:  * If no scheme is provided, &#x60;https&#x60; is assumed. * An HTTP GET on the URL must yield a [google.protobuf.Type][]   value in binary format, or produce an error. * Applications are allowed to cache lookup results based on the   URL, or have them precompiled into a binary to avoid any   lookup. Therefore, binary compatibility needs to be preserved   on changes to types. (Use versioned type names to manage   breaking changes.)  Note: this functionality is not currently available in the official protobuf release, and it is not used for type URLs beginning with type.googleapis.com.  Schemes other than &#x60;http&#x60;, &#x60;https&#x60; (or the empty scheme) might be used with implementation specific semantics. | [optional] 
**Value** | Pointer to **string** | Must be a valid serialized protocol buffer of the above specified type. | [optional] 

## Methods

### NewGoogleProtobufAny

`func NewGoogleProtobufAny() *GoogleProtobufAny`

NewGoogleProtobufAny instantiates a new GoogleProtobufAny object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGoogleProtobufAnyWithDefaults

`func NewGoogleProtobufAnyWithDefaults() *GoogleProtobufAny`

NewGoogleProtobufAnyWithDefaults instantiates a new GoogleProtobufAny object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTypeUrl

`func (o *GoogleProtobufAny) GetTypeUrl() string`

GetTypeUrl returns the TypeUrl field if non-nil, zero value otherwise.

### GetTypeUrlOk

`func (o *GoogleProtobufAny) GetTypeUrlOk() (*string, bool)`

GetTypeUrlOk returns a tuple with the TypeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeUrl

`func (o *GoogleProtobufAny) SetTypeUrl(v string)`

SetTypeUrl sets TypeUrl field to given value.

### HasTypeUrl

`func (o *GoogleProtobufAny) HasTypeUrl() bool`

HasTypeUrl returns a boolean if a field has been set.

### GetValue

`func (o *GoogleProtobufAny) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *GoogleProtobufAny) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *GoogleProtobufAny) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *GoogleProtobufAny) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


