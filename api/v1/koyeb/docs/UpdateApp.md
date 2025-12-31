# UpdateApp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**LifeCycle** | Pointer to [**AppLifeCycle**](AppLifeCycle.md) |  | [optional] 

## Methods

### NewUpdateApp

`func NewUpdateApp() *UpdateApp`

NewUpdateApp instantiates a new UpdateApp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAppWithDefaults

`func NewUpdateAppWithDefaults() *UpdateApp`

NewUpdateAppWithDefaults instantiates a new UpdateApp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateApp) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateApp) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateApp) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateApp) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLifeCycle

`func (o *UpdateApp) GetLifeCycle() AppLifeCycle`

GetLifeCycle returns the LifeCycle field if non-nil, zero value otherwise.

### GetLifeCycleOk

`func (o *UpdateApp) GetLifeCycleOk() (*AppLifeCycle, bool)`

GetLifeCycleOk returns a tuple with the LifeCycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifeCycle

`func (o *UpdateApp) SetLifeCycle(v AppLifeCycle)`

SetLifeCycle sets LifeCycle field to given value.

### HasLifeCycle

`func (o *UpdateApp) HasLifeCycle() bool`

HasLifeCycle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


