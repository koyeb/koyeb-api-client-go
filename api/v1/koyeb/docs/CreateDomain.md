# CreateDomain

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**DomainType**](DomainType.md) |  | [optional] [default to DOMAINTYPE_AUTOASSIGNED]
**AppId** | Pointer to **string** |  | [optional] 
**Cloudflare** | Pointer to **map[string]interface{}** |  | [optional] 
**Koyeb** | Pointer to [**DomainLoadBalancerKoyeb**](DomainLoadBalancerKoyeb.md) |  | [optional] 

## Methods

### NewCreateDomain

`func NewCreateDomain() *CreateDomain`

NewCreateDomain instantiates a new CreateDomain object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDomainWithDefaults

`func NewCreateDomainWithDefaults() *CreateDomain`

NewCreateDomainWithDefaults instantiates a new CreateDomain object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateDomain) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateDomain) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateDomain) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateDomain) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *CreateDomain) GetType() DomainType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateDomain) GetTypeOk() (*DomainType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateDomain) SetType(v DomainType)`

SetType sets Type field to given value.

### HasType

`func (o *CreateDomain) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAppId

`func (o *CreateDomain) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *CreateDomain) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *CreateDomain) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *CreateDomain) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### GetCloudflare

`func (o *CreateDomain) GetCloudflare() map[string]interface{}`

GetCloudflare returns the Cloudflare field if non-nil, zero value otherwise.

### GetCloudflareOk

`func (o *CreateDomain) GetCloudflareOk() (*map[string]interface{}, bool)`

GetCloudflareOk returns a tuple with the Cloudflare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudflare

`func (o *CreateDomain) SetCloudflare(v map[string]interface{})`

SetCloudflare sets Cloudflare field to given value.

### HasCloudflare

`func (o *CreateDomain) HasCloudflare() bool`

HasCloudflare returns a boolean if a field has been set.

### GetKoyeb

`func (o *CreateDomain) GetKoyeb() DomainLoadBalancerKoyeb`

GetKoyeb returns the Koyeb field if non-nil, zero value otherwise.

### GetKoyebOk

`func (o *CreateDomain) GetKoyebOk() (*DomainLoadBalancerKoyeb, bool)`

GetKoyebOk returns a tuple with the Koyeb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKoyeb

`func (o *CreateDomain) SetKoyeb(v DomainLoadBalancerKoyeb)`

SetKoyeb sets Koyeb field to given value.

### HasKoyeb

`func (o *CreateDomain) HasKoyeb() bool`

HasKoyeb returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


