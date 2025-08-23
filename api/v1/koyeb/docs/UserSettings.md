# UserSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 
**FailedDeploymentEmailNotification** | Pointer to **bool** |  | [optional] 

## Methods

### NewUserSettings

`func NewUserSettings() *UserSettings`

NewUserSettings instantiates a new UserSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserSettingsWithDefaults

`func NewUserSettingsWithDefaults() *UserSettings`

NewUserSettingsWithDefaults instantiates a new UserSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserSettings) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserSettings) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserSettings) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UserSettings) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUserId

`func (o *UserSettings) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UserSettings) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UserSettings) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *UserSettings) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetFailedDeploymentEmailNotification

`func (o *UserSettings) GetFailedDeploymentEmailNotification() bool`

GetFailedDeploymentEmailNotification returns the FailedDeploymentEmailNotification field if non-nil, zero value otherwise.

### GetFailedDeploymentEmailNotificationOk

`func (o *UserSettings) GetFailedDeploymentEmailNotificationOk() (*bool, bool)`

GetFailedDeploymentEmailNotificationOk returns a tuple with the FailedDeploymentEmailNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedDeploymentEmailNotification

`func (o *UserSettings) SetFailedDeploymentEmailNotification(v bool)`

SetFailedDeploymentEmailNotification sets FailedDeploymentEmailNotification field to given value.

### HasFailedDeploymentEmailNotification

`func (o *UserSettings) HasFailedDeploymentEmailNotification() bool`

HasFailedDeploymentEmailNotification returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


