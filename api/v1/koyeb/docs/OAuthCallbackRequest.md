# OAuthCallbackRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**SetupAction** | Pointer to **string** | setup_action is populated in the context of a GitHub app installation request. For logins and signups, it is not set. | [optional] 
**InstallationId** | Pointer to **string** | installation_id is populated in the context of a GitHub app installation request. For logins and signups, it is not set. | [optional] 

## Methods

### NewOAuthCallbackRequest

`func NewOAuthCallbackRequest() *OAuthCallbackRequest`

NewOAuthCallbackRequest instantiates a new OAuthCallbackRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAuthCallbackRequestWithDefaults

`func NewOAuthCallbackRequestWithDefaults() *OAuthCallbackRequest`

NewOAuthCallbackRequestWithDefaults instantiates a new OAuthCallbackRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *OAuthCallbackRequest) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *OAuthCallbackRequest) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *OAuthCallbackRequest) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *OAuthCallbackRequest) HasState() bool`

HasState returns a boolean if a field has been set.

### GetCode

`func (o *OAuthCallbackRequest) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *OAuthCallbackRequest) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *OAuthCallbackRequest) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *OAuthCallbackRequest) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetSetupAction

`func (o *OAuthCallbackRequest) GetSetupAction() string`

GetSetupAction returns the SetupAction field if non-nil, zero value otherwise.

### GetSetupActionOk

`func (o *OAuthCallbackRequest) GetSetupActionOk() (*string, bool)`

GetSetupActionOk returns a tuple with the SetupAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetupAction

`func (o *OAuthCallbackRequest) SetSetupAction(v string)`

SetSetupAction sets SetupAction field to given value.

### HasSetupAction

`func (o *OAuthCallbackRequest) HasSetupAction() bool`

HasSetupAction returns a boolean if a field has been set.

### GetInstallationId

`func (o *OAuthCallbackRequest) GetInstallationId() string`

GetInstallationId returns the InstallationId field if non-nil, zero value otherwise.

### GetInstallationIdOk

`func (o *OAuthCallbackRequest) GetInstallationIdOk() (*string, bool)`

GetInstallationIdOk returns a tuple with the InstallationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallationId

`func (o *OAuthCallbackRequest) SetInstallationId(v string)`

SetInstallationId sets InstallationId field to given value.

### HasInstallationId

`func (o *OAuthCallbackRequest) HasInstallationId() bool`

HasInstallationId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


