# GetGithubInstallationReply

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstallationId** | Pointer to **string** |  | [optional] 
**InstallationUrl** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**AvatarUrl** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**KgitproxyGithubInstallationStatus**](KgitproxyGithubInstallationStatus.md) |  | [optional] [default to KGITPROXYGITHUBINSTALLATIONSTATUS_INVALID]
**InstalledAt** | Pointer to **time.Time** |  | [optional] 
**SuspendedAt** | Pointer to **time.Time** |  | [optional] 
**IndexingStatus** | Pointer to [**KgitproxyIndexingStatus**](KgitproxyIndexingStatus.md) |  | [optional] [default to KGITPROXYINDEXINGSTATUS_INVALID_INDEXING_STATUS]
**IndexedRepositories** | Pointer to **int64** |  | [optional] 
**TotalRepositories** | Pointer to **int64** |  | [optional] 

## Methods

### NewGetGithubInstallationReply

`func NewGetGithubInstallationReply() *GetGithubInstallationReply`

NewGetGithubInstallationReply instantiates a new GetGithubInstallationReply object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetGithubInstallationReplyWithDefaults

`func NewGetGithubInstallationReplyWithDefaults() *GetGithubInstallationReply`

NewGetGithubInstallationReplyWithDefaults instantiates a new GetGithubInstallationReply object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstallationId

`func (o *GetGithubInstallationReply) GetInstallationId() string`

GetInstallationId returns the InstallationId field if non-nil, zero value otherwise.

### GetInstallationIdOk

`func (o *GetGithubInstallationReply) GetInstallationIdOk() (*string, bool)`

GetInstallationIdOk returns a tuple with the InstallationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallationId

`func (o *GetGithubInstallationReply) SetInstallationId(v string)`

SetInstallationId sets InstallationId field to given value.

### HasInstallationId

`func (o *GetGithubInstallationReply) HasInstallationId() bool`

HasInstallationId returns a boolean if a field has been set.

### GetInstallationUrl

`func (o *GetGithubInstallationReply) GetInstallationUrl() string`

GetInstallationUrl returns the InstallationUrl field if non-nil, zero value otherwise.

### GetInstallationUrlOk

`func (o *GetGithubInstallationReply) GetInstallationUrlOk() (*string, bool)`

GetInstallationUrlOk returns a tuple with the InstallationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallationUrl

`func (o *GetGithubInstallationReply) SetInstallationUrl(v string)`

SetInstallationUrl sets InstallationUrl field to given value.

### HasInstallationUrl

`func (o *GetGithubInstallationReply) HasInstallationUrl() bool`

HasInstallationUrl returns a boolean if a field has been set.

### GetName

`func (o *GetGithubInstallationReply) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetGithubInstallationReply) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetGithubInstallationReply) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetGithubInstallationReply) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAvatarUrl

`func (o *GetGithubInstallationReply) GetAvatarUrl() string`

GetAvatarUrl returns the AvatarUrl field if non-nil, zero value otherwise.

### GetAvatarUrlOk

`func (o *GetGithubInstallationReply) GetAvatarUrlOk() (*string, bool)`

GetAvatarUrlOk returns a tuple with the AvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarUrl

`func (o *GetGithubInstallationReply) SetAvatarUrl(v string)`

SetAvatarUrl sets AvatarUrl field to given value.

### HasAvatarUrl

`func (o *GetGithubInstallationReply) HasAvatarUrl() bool`

HasAvatarUrl returns a boolean if a field has been set.

### GetStatus

`func (o *GetGithubInstallationReply) GetStatus() KgitproxyGithubInstallationStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetGithubInstallationReply) GetStatusOk() (*KgitproxyGithubInstallationStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetGithubInstallationReply) SetStatus(v KgitproxyGithubInstallationStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetGithubInstallationReply) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetInstalledAt

`func (o *GetGithubInstallationReply) GetInstalledAt() time.Time`

GetInstalledAt returns the InstalledAt field if non-nil, zero value otherwise.

### GetInstalledAtOk

`func (o *GetGithubInstallationReply) GetInstalledAtOk() (*time.Time, bool)`

GetInstalledAtOk returns a tuple with the InstalledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstalledAt

`func (o *GetGithubInstallationReply) SetInstalledAt(v time.Time)`

SetInstalledAt sets InstalledAt field to given value.

### HasInstalledAt

`func (o *GetGithubInstallationReply) HasInstalledAt() bool`

HasInstalledAt returns a boolean if a field has been set.

### GetSuspendedAt

`func (o *GetGithubInstallationReply) GetSuspendedAt() time.Time`

GetSuspendedAt returns the SuspendedAt field if non-nil, zero value otherwise.

### GetSuspendedAtOk

`func (o *GetGithubInstallationReply) GetSuspendedAtOk() (*time.Time, bool)`

GetSuspendedAtOk returns a tuple with the SuspendedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspendedAt

`func (o *GetGithubInstallationReply) SetSuspendedAt(v time.Time)`

SetSuspendedAt sets SuspendedAt field to given value.

### HasSuspendedAt

`func (o *GetGithubInstallationReply) HasSuspendedAt() bool`

HasSuspendedAt returns a boolean if a field has been set.

### GetIndexingStatus

`func (o *GetGithubInstallationReply) GetIndexingStatus() KgitproxyIndexingStatus`

GetIndexingStatus returns the IndexingStatus field if non-nil, zero value otherwise.

### GetIndexingStatusOk

`func (o *GetGithubInstallationReply) GetIndexingStatusOk() (*KgitproxyIndexingStatus, bool)`

GetIndexingStatusOk returns a tuple with the IndexingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexingStatus

`func (o *GetGithubInstallationReply) SetIndexingStatus(v KgitproxyIndexingStatus)`

SetIndexingStatus sets IndexingStatus field to given value.

### HasIndexingStatus

`func (o *GetGithubInstallationReply) HasIndexingStatus() bool`

HasIndexingStatus returns a boolean if a field has been set.

### GetIndexedRepositories

`func (o *GetGithubInstallationReply) GetIndexedRepositories() int64`

GetIndexedRepositories returns the IndexedRepositories field if non-nil, zero value otherwise.

### GetIndexedRepositoriesOk

`func (o *GetGithubInstallationReply) GetIndexedRepositoriesOk() (*int64, bool)`

GetIndexedRepositoriesOk returns a tuple with the IndexedRepositories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexedRepositories

`func (o *GetGithubInstallationReply) SetIndexedRepositories(v int64)`

SetIndexedRepositories sets IndexedRepositories field to given value.

### HasIndexedRepositories

`func (o *GetGithubInstallationReply) HasIndexedRepositories() bool`

HasIndexedRepositories returns a boolean if a field has been set.

### GetTotalRepositories

`func (o *GetGithubInstallationReply) GetTotalRepositories() int64`

GetTotalRepositories returns the TotalRepositories field if non-nil, zero value otherwise.

### GetTotalRepositoriesOk

`func (o *GetGithubInstallationReply) GetTotalRepositoriesOk() (*int64, bool)`

GetTotalRepositoriesOk returns a tuple with the TotalRepositories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRepositories

`func (o *GetGithubInstallationReply) SetTotalRepositories(v int64)`

SetTotalRepositories sets TotalRepositories field to given value.

### HasTotalRepositories

`func (o *GetGithubInstallationReply) HasTotalRepositories() bool`

HasTotalRepositories returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


