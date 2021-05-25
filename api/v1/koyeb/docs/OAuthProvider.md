# OAuthProvider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 

## Methods

### NewOAuthProvider

`func NewOAuthProvider() *OAuthProvider`

NewOAuthProvider instantiates a new OAuthProvider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAuthProviderWithDefaults

`func NewOAuthProviderWithDefaults() *OAuthProvider`

NewOAuthProviderWithDefaults instantiates a new OAuthProvider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OAuthProvider) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OAuthProvider) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OAuthProvider) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OAuthProvider) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUrl

`func (o *OAuthProvider) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *OAuthProvider) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *OAuthProvider) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *OAuthProvider) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetState

`func (o *OAuthProvider) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *OAuthProvider) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *OAuthProvider) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *OAuthProvider) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


