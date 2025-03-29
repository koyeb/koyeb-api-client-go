# CreateCompose

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**App** | Pointer to [**CreateApp**](CreateApp.md) |  | [optional] 
**Services** | Pointer to [**[]CreateService**](CreateService.md) |  | [optional] 

## Methods

### NewCreateCompose

`func NewCreateCompose() *CreateCompose`

NewCreateCompose instantiates a new CreateCompose object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateComposeWithDefaults

`func NewCreateComposeWithDefaults() *CreateCompose`

NewCreateComposeWithDefaults instantiates a new CreateCompose object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApp

`func (o *CreateCompose) GetApp() CreateApp`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *CreateCompose) GetAppOk() (*CreateApp, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *CreateCompose) SetApp(v CreateApp)`

SetApp sets App field to given value.

### HasApp

`func (o *CreateCompose) HasApp() bool`

HasApp returns a boolean if a field has been set.

### GetServices

`func (o *CreateCompose) GetServices() []CreateService`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *CreateCompose) GetServicesOk() (*[]CreateService, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *CreateCompose) SetServices(v []CreateService)`

SetServices sets Services field to given value.

### HasServices

`func (o *CreateCompose) HasServices() bool`

HasServices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


