# StreamResultOfEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | Pointer to [**Event**](Event.md) |  | [optional] 
**Error** | Pointer to [**GoogleRpcStatus**](GoogleRpcStatus.md) |  | [optional] 

## Methods

### NewStreamResultOfEvent

`func NewStreamResultOfEvent() *StreamResultOfEvent`

NewStreamResultOfEvent instantiates a new StreamResultOfEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamResultOfEventWithDefaults

`func NewStreamResultOfEventWithDefaults() *StreamResultOfEvent`

NewStreamResultOfEventWithDefaults instantiates a new StreamResultOfEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *StreamResultOfEvent) GetResult() Event`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *StreamResultOfEvent) GetResultOk() (*Event, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *StreamResultOfEvent) SetResult(v Event)`

SetResult sets Result field to given value.

### HasResult

`func (o *StreamResultOfEvent) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetError

`func (o *StreamResultOfEvent) GetError() GoogleRpcStatus`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *StreamResultOfEvent) GetErrorOk() (*GoogleRpcStatus, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *StreamResultOfEvent) SetError(v GoogleRpcStatus)`

SetError sets Error field to given value.

### HasError

`func (o *StreamResultOfEvent) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


