# PoolClaim

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**PoolId** | Pointer to **string** |  | [optional] 
**ServiceId** | Pointer to **string** |  | [optional] 
**RequestId** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**PoolClaimStatus**](PoolClaimStatus.md) |  | [optional] [default to POOLCLAIMSTATUS_UNSPECIFIED]
**OrganizationId** | Pointer to **string** |  | [optional] 
**WorkspaceId** | Pointer to **string** |  | [optional] 
**CustomerId** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**FulfilledAt** | Pointer to **time.Time** |  | [optional] 
**ReleasedAt** | Pointer to **time.Time** |  | [optional] 
**SandboxId** | Pointer to **string** |  | [optional] 
**PoolGeneration** | Pointer to **string** |  | [optional] 

## Methods

### NewPoolClaim

`func NewPoolClaim() *PoolClaim`

NewPoolClaim instantiates a new PoolClaim object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPoolClaimWithDefaults

`func NewPoolClaimWithDefaults() *PoolClaim`

NewPoolClaimWithDefaults instantiates a new PoolClaim object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PoolClaim) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PoolClaim) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PoolClaim) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PoolClaim) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPoolId

`func (o *PoolClaim) GetPoolId() string`

GetPoolId returns the PoolId field if non-nil, zero value otherwise.

### GetPoolIdOk

`func (o *PoolClaim) GetPoolIdOk() (*string, bool)`

GetPoolIdOk returns a tuple with the PoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolId

`func (o *PoolClaim) SetPoolId(v string)`

SetPoolId sets PoolId field to given value.

### HasPoolId

`func (o *PoolClaim) HasPoolId() bool`

HasPoolId returns a boolean if a field has been set.

### GetServiceId

`func (o *PoolClaim) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *PoolClaim) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *PoolClaim) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *PoolClaim) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetRequestId

`func (o *PoolClaim) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *PoolClaim) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *PoolClaim) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *PoolClaim) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetStatus

`func (o *PoolClaim) GetStatus() PoolClaimStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PoolClaim) GetStatusOk() (*PoolClaimStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PoolClaim) SetStatus(v PoolClaimStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PoolClaim) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetOrganizationId

`func (o *PoolClaim) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *PoolClaim) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *PoolClaim) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *PoolClaim) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetWorkspaceId

`func (o *PoolClaim) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *PoolClaim) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *PoolClaim) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.

### HasWorkspaceId

`func (o *PoolClaim) HasWorkspaceId() bool`

HasWorkspaceId returns a boolean if a field has been set.

### GetCustomerId

`func (o *PoolClaim) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *PoolClaim) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *PoolClaim) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *PoolClaim) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PoolClaim) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PoolClaim) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PoolClaim) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PoolClaim) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetFulfilledAt

`func (o *PoolClaim) GetFulfilledAt() time.Time`

GetFulfilledAt returns the FulfilledAt field if non-nil, zero value otherwise.

### GetFulfilledAtOk

`func (o *PoolClaim) GetFulfilledAtOk() (*time.Time, bool)`

GetFulfilledAtOk returns a tuple with the FulfilledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfilledAt

`func (o *PoolClaim) SetFulfilledAt(v time.Time)`

SetFulfilledAt sets FulfilledAt field to given value.

### HasFulfilledAt

`func (o *PoolClaim) HasFulfilledAt() bool`

HasFulfilledAt returns a boolean if a field has been set.

### GetReleasedAt

`func (o *PoolClaim) GetReleasedAt() time.Time`

GetReleasedAt returns the ReleasedAt field if non-nil, zero value otherwise.

### GetReleasedAtOk

`func (o *PoolClaim) GetReleasedAtOk() (*time.Time, bool)`

GetReleasedAtOk returns a tuple with the ReleasedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleasedAt

`func (o *PoolClaim) SetReleasedAt(v time.Time)`

SetReleasedAt sets ReleasedAt field to given value.

### HasReleasedAt

`func (o *PoolClaim) HasReleasedAt() bool`

HasReleasedAt returns a boolean if a field has been set.

### GetSandboxId

`func (o *PoolClaim) GetSandboxId() string`

GetSandboxId returns the SandboxId field if non-nil, zero value otherwise.

### GetSandboxIdOk

`func (o *PoolClaim) GetSandboxIdOk() (*string, bool)`

GetSandboxIdOk returns a tuple with the SandboxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSandboxId

`func (o *PoolClaim) SetSandboxId(v string)`

SetSandboxId sets SandboxId field to given value.

### HasSandboxId

`func (o *PoolClaim) HasSandboxId() bool`

HasSandboxId returns a boolean if a field has been set.

### GetPoolGeneration

`func (o *PoolClaim) GetPoolGeneration() string`

GetPoolGeneration returns the PoolGeneration field if non-nil, zero value otherwise.

### GetPoolGenerationOk

`func (o *PoolClaim) GetPoolGenerationOk() (*string, bool)`

GetPoolGenerationOk returns a tuple with the PoolGeneration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolGeneration

`func (o *PoolClaim) SetPoolGeneration(v string)`

SetPoolGeneration sets PoolGeneration field to given value.

### HasPoolGeneration

`func (o *PoolClaim) HasPoolGeneration() bool`

HasPoolGeneration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


