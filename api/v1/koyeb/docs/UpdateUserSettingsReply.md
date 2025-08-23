# UpdateUserSettingsReply

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Settings** | Pointer to [**UserSettings**](UserSettings.md) |  | [optional] 

## Methods

### NewUpdateUserSettingsReply

`func NewUpdateUserSettingsReply() *UpdateUserSettingsReply`

NewUpdateUserSettingsReply instantiates a new UpdateUserSettingsReply object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateUserSettingsReplyWithDefaults

`func NewUpdateUserSettingsReplyWithDefaults() *UpdateUserSettingsReply`

NewUpdateUserSettingsReplyWithDefaults instantiates a new UpdateUserSettingsReply object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSettings

`func (o *UpdateUserSettingsReply) GetSettings() UserSettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *UpdateUserSettingsReply) GetSettingsOk() (*UserSettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *UpdateUserSettingsReply) SetSettings(v UserSettings)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *UpdateUserSettingsReply) HasSettings() bool`

HasSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


