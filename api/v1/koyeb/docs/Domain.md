# Domain

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**OrganizationId** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**Status** | Pointer to [**DomainStatus**](DomainStatus.md) |  | [optional] [default to DOMAINSTATUS_PENDING]
**Type** | Pointer to [**DomainType**](DomainType.md) |  | [optional] [default to DOMAINTYPE_AUTOASSIGNED]
**AppId** | Pointer to **string** |  | [optional] 
**DeploymentGroup** | Pointer to **string** |  | [optional] 
**VerifiedAt** | Pointer to **time.Time** |  | [optional] 
**IntendedCname** | Pointer to **string** |  | [optional] 
**Messages** | Pointer to **[]string** |  | [optional] 

## Methods

### NewDomain

`func NewDomain() *Domain`

NewDomain instantiates a new Domain object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainWithDefaults

`func NewDomainWithDefaults() *Domain`

NewDomainWithDefaults instantiates a new Domain object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Domain) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Domain) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Domain) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Domain) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOrganizationId

`func (o *Domain) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *Domain) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *Domain) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *Domain) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetName

`func (o *Domain) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Domain) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Domain) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Domain) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Domain) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Domain) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Domain) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Domain) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Domain) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Domain) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Domain) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Domain) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetStatus

`func (o *Domain) GetStatus() DomainStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Domain) GetStatusOk() (*DomainStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Domain) SetStatus(v DomainStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Domain) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *Domain) GetType() DomainType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Domain) GetTypeOk() (*DomainType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Domain) SetType(v DomainType)`

SetType sets Type field to given value.

### HasType

`func (o *Domain) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAppId

`func (o *Domain) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *Domain) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *Domain) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *Domain) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### GetDeploymentGroup

`func (o *Domain) GetDeploymentGroup() string`

GetDeploymentGroup returns the DeploymentGroup field if non-nil, zero value otherwise.

### GetDeploymentGroupOk

`func (o *Domain) GetDeploymentGroupOk() (*string, bool)`

GetDeploymentGroupOk returns a tuple with the DeploymentGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentGroup

`func (o *Domain) SetDeploymentGroup(v string)`

SetDeploymentGroup sets DeploymentGroup field to given value.

### HasDeploymentGroup

`func (o *Domain) HasDeploymentGroup() bool`

HasDeploymentGroup returns a boolean if a field has been set.

### GetVerifiedAt

`func (o *Domain) GetVerifiedAt() time.Time`

GetVerifiedAt returns the VerifiedAt field if non-nil, zero value otherwise.

### GetVerifiedAtOk

`func (o *Domain) GetVerifiedAtOk() (*time.Time, bool)`

GetVerifiedAtOk returns a tuple with the VerifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifiedAt

`func (o *Domain) SetVerifiedAt(v time.Time)`

SetVerifiedAt sets VerifiedAt field to given value.

### HasVerifiedAt

`func (o *Domain) HasVerifiedAt() bool`

HasVerifiedAt returns a boolean if a field has been set.

### GetIntendedCname

`func (o *Domain) GetIntendedCname() string`

GetIntendedCname returns the IntendedCname field if non-nil, zero value otherwise.

### GetIntendedCnameOk

`func (o *Domain) GetIntendedCnameOk() (*string, bool)`

GetIntendedCnameOk returns a tuple with the IntendedCname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntendedCname

`func (o *Domain) SetIntendedCname(v string)`

SetIntendedCname sets IntendedCname field to given value.

### HasIntendedCname

`func (o *Domain) HasIntendedCname() bool`

HasIntendedCname returns a boolean if a field has been set.

### GetMessages

`func (o *Domain) GetMessages() []string`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *Domain) GetMessagesOk() (*[]string, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *Domain) SetMessages(v []string)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *Domain) HasMessages() bool`

HasMessages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


