# Service

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**StartedAt** | Pointer to **time.Time** |  | [optional] 
**SucceededAt** | Pointer to **time.Time** |  | [optional] 
**PausedAt** | Pointer to **time.Time** |  | [optional] 
**ResumedAt** | Pointer to **time.Time** |  | [optional] 
**TerminatedAt** | Pointer to **time.Time** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**ServiceType**](ServiceType.md) |  | [optional] [default to SERVICETYPE_INVALID_TYPE]
**OrganizationId** | Pointer to **string** |  | [optional] 
**AppId** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**ServiceStatus**](ServiceStatus.md) |  | [optional] [default to SERVICESTATUS_STARTING]
**Messages** | Pointer to **[]string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**ActiveDeploymentId** | Pointer to **string** |  | [optional] 
**LatestDeploymentId** | Pointer to **string** |  | [optional] 
**LastProvisionedDeploymentId** | Pointer to **string** |  | [optional] 
**State** | Pointer to [**ServiceState**](ServiceState.md) |  | [optional] 

## Methods

### NewService

`func NewService() *Service`

NewService instantiates a new Service object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceWithDefaults

`func NewServiceWithDefaults() *Service`

NewServiceWithDefaults instantiates a new Service object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Service) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Service) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Service) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Service) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Service) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Service) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Service) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Service) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Service) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Service) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Service) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Service) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetStartedAt

`func (o *Service) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *Service) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *Service) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *Service) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetSucceededAt

`func (o *Service) GetSucceededAt() time.Time`

GetSucceededAt returns the SucceededAt field if non-nil, zero value otherwise.

### GetSucceededAtOk

`func (o *Service) GetSucceededAtOk() (*time.Time, bool)`

GetSucceededAtOk returns a tuple with the SucceededAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSucceededAt

`func (o *Service) SetSucceededAt(v time.Time)`

SetSucceededAt sets SucceededAt field to given value.

### HasSucceededAt

`func (o *Service) HasSucceededAt() bool`

HasSucceededAt returns a boolean if a field has been set.

### GetPausedAt

`func (o *Service) GetPausedAt() time.Time`

GetPausedAt returns the PausedAt field if non-nil, zero value otherwise.

### GetPausedAtOk

`func (o *Service) GetPausedAtOk() (*time.Time, bool)`

GetPausedAtOk returns a tuple with the PausedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPausedAt

`func (o *Service) SetPausedAt(v time.Time)`

SetPausedAt sets PausedAt field to given value.

### HasPausedAt

`func (o *Service) HasPausedAt() bool`

HasPausedAt returns a boolean if a field has been set.

### GetResumedAt

`func (o *Service) GetResumedAt() time.Time`

GetResumedAt returns the ResumedAt field if non-nil, zero value otherwise.

### GetResumedAtOk

`func (o *Service) GetResumedAtOk() (*time.Time, bool)`

GetResumedAtOk returns a tuple with the ResumedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResumedAt

`func (o *Service) SetResumedAt(v time.Time)`

SetResumedAt sets ResumedAt field to given value.

### HasResumedAt

`func (o *Service) HasResumedAt() bool`

HasResumedAt returns a boolean if a field has been set.

### GetTerminatedAt

`func (o *Service) GetTerminatedAt() time.Time`

GetTerminatedAt returns the TerminatedAt field if non-nil, zero value otherwise.

### GetTerminatedAtOk

`func (o *Service) GetTerminatedAtOk() (*time.Time, bool)`

GetTerminatedAtOk returns a tuple with the TerminatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminatedAt

`func (o *Service) SetTerminatedAt(v time.Time)`

SetTerminatedAt sets TerminatedAt field to given value.

### HasTerminatedAt

`func (o *Service) HasTerminatedAt() bool`

HasTerminatedAt returns a boolean if a field has been set.

### GetName

`func (o *Service) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Service) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Service) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Service) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *Service) GetType() ServiceType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Service) GetTypeOk() (*ServiceType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Service) SetType(v ServiceType)`

SetType sets Type field to given value.

### HasType

`func (o *Service) HasType() bool`

HasType returns a boolean if a field has been set.

### GetOrganizationId

`func (o *Service) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *Service) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *Service) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *Service) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetAppId

`func (o *Service) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *Service) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *Service) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *Service) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### GetStatus

`func (o *Service) GetStatus() ServiceStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Service) GetStatusOk() (*ServiceStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Service) SetStatus(v ServiceStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Service) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMessages

`func (o *Service) GetMessages() []string`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *Service) GetMessagesOk() (*[]string, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *Service) SetMessages(v []string)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *Service) HasMessages() bool`

HasMessages returns a boolean if a field has been set.

### GetVersion

`func (o *Service) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Service) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Service) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Service) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetActiveDeploymentId

`func (o *Service) GetActiveDeploymentId() string`

GetActiveDeploymentId returns the ActiveDeploymentId field if non-nil, zero value otherwise.

### GetActiveDeploymentIdOk

`func (o *Service) GetActiveDeploymentIdOk() (*string, bool)`

GetActiveDeploymentIdOk returns a tuple with the ActiveDeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveDeploymentId

`func (o *Service) SetActiveDeploymentId(v string)`

SetActiveDeploymentId sets ActiveDeploymentId field to given value.

### HasActiveDeploymentId

`func (o *Service) HasActiveDeploymentId() bool`

HasActiveDeploymentId returns a boolean if a field has been set.

### GetLatestDeploymentId

`func (o *Service) GetLatestDeploymentId() string`

GetLatestDeploymentId returns the LatestDeploymentId field if non-nil, zero value otherwise.

### GetLatestDeploymentIdOk

`func (o *Service) GetLatestDeploymentIdOk() (*string, bool)`

GetLatestDeploymentIdOk returns a tuple with the LatestDeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestDeploymentId

`func (o *Service) SetLatestDeploymentId(v string)`

SetLatestDeploymentId sets LatestDeploymentId field to given value.

### HasLatestDeploymentId

`func (o *Service) HasLatestDeploymentId() bool`

HasLatestDeploymentId returns a boolean if a field has been set.

### GetLastProvisionedDeploymentId

`func (o *Service) GetLastProvisionedDeploymentId() string`

GetLastProvisionedDeploymentId returns the LastProvisionedDeploymentId field if non-nil, zero value otherwise.

### GetLastProvisionedDeploymentIdOk

`func (o *Service) GetLastProvisionedDeploymentIdOk() (*string, bool)`

GetLastProvisionedDeploymentIdOk returns a tuple with the LastProvisionedDeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastProvisionedDeploymentId

`func (o *Service) SetLastProvisionedDeploymentId(v string)`

SetLastProvisionedDeploymentId sets LastProvisionedDeploymentId field to given value.

### HasLastProvisionedDeploymentId

`func (o *Service) HasLastProvisionedDeploymentId() bool`

HasLastProvisionedDeploymentId returns a boolean if a field has been set.

### GetState

`func (o *Service) GetState() ServiceState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Service) GetStateOk() (*ServiceState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Service) SetState(v ServiceState)`

SetState sets State field to given value.

### HasState

`func (o *Service) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


