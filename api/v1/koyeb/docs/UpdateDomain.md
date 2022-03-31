# UpdateDomain

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**DomainStatus**](DomainStatus.md) |  | [optional] [default to DOMAINSTATUS_PENDING]
**AppId** | Pointer to **string** |  | [optional] 

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

### GetStatus

`func (o *UpdateDomain) GetStatus() DomainStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UpdateDomain) GetStatusOk() (*DomainStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UpdateDomain) SetStatus(v DomainStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UpdateDomain) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


