# ComposeReply

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**App** | Pointer to [**App**](App.md) |  | [optional] 
**Services** | Pointer to [**[]Service**](Service.md) |  | [optional] 

## Methods

### NewComposeReply

`func NewComposeReply() *ComposeReply`

NewComposeReply instantiates a new ComposeReply object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComposeReplyWithDefaults

`func NewComposeReplyWithDefaults() *ComposeReply`

NewComposeReplyWithDefaults instantiates a new ComposeReply object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApp

`func (o *ComposeReply) GetApp() App`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *ComposeReply) GetAppOk() (*App, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *ComposeReply) SetApp(v App)`

SetApp sets App field to given value.

### HasApp

`func (o *ComposeReply) HasApp() bool`

HasApp returns a boolean if a field has been set.

### GetServices

`func (o *ComposeReply) GetServices() []Service`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *ComposeReply) GetServicesOk() (*[]Service, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *ComposeReply) SetServices(v []Service)`

SetServices sets Services field to given value.

### HasServices

`func (o *ComposeReply) HasServices() bool`

HasServices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


