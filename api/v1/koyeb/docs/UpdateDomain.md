# UpdateDomain

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppId** | Pointer to **string** | To attach or detach from an app for custom domain. | [optional] 
**Subdomain** | Pointer to **string** | To change subdomain for auto-assigned domain. | [optional] 

## Methods

### NewUpdateDomain

`func NewUpdateDomain() *UpdateDomain`

NewUpdateDomain instantiates a new UpdateDomain object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDomainWithDefaults

`func NewUpdateDomainWithDefaults() *UpdateDomain`

NewUpdateDomainWithDefaults instantiates a new UpdateDomain object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppId

`func (o *UpdateDomain) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *UpdateDomain) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *UpdateDomain) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *UpdateDomain) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### GetSubdomain

`func (o *UpdateDomain) GetSubdomain() string`

GetSubdomain returns the Subdomain field if non-nil, zero value otherwise.

### GetSubdomainOk

`func (o *UpdateDomain) GetSubdomainOk() (*string, bool)`

GetSubdomainOk returns a tuple with the Subdomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdomain

`func (o *UpdateDomain) SetSubdomain(v string)`

SetSubdomain sets Subdomain field to given value.

### HasSubdomain

`func (o *UpdateDomain) HasSubdomain() bool`

HasSubdomain returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


