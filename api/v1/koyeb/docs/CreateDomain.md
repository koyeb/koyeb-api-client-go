# CreateDomain

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**DomainType**](DomainType.md) |  | [optional] [default to DOMAINTYPE_AUTOASSIGNED]
**AppId** | Pointer to **string** |  | [optional] 

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


