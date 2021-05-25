# StackRevision

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sha** | Pointer to **string** |  | [optional] 
**Yaml** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**ParentSha** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**StackRevisionStatus**](StackRevisionStatus.md) |  | [optional] [default to STACKREVISIONSTATUS_UNKNOWN]
**StatusMessage** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**CommitInfo** | Pointer to [**CommitInfo**](CommitInfo.md) |  | [optional] 
**ReleaseId** | Pointer to **string** |  | [optional] 

## Methods

### NewStackRevision

`func NewStackRevision() *StackRevision`

NewStackRevision instantiates a new StackRevision object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStackRevisionWithDefaults

`func NewStackRevisionWithDefaults() *StackRevision`

NewStackRevisionWithDefaults instantiates a new StackRevision object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSha

`func (o *StackRevision) GetSha() string`

GetSha returns the Sha field if non-nil, zero value otherwise.

### GetShaOk

`func (o *StackRevision) GetShaOk() (*string, bool)`

GetShaOk returns a tuple with the Sha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha

`func (o *StackRevision) SetSha(v string)`

SetSha sets Sha field to given value.

### HasSha

`func (o *StackRevision) HasSha() bool`

HasSha returns a boolean if a field has been set.

### GetYaml

`func (o *StackRevision) GetYaml() string`

GetYaml returns the Yaml field if non-nil, zero value otherwise.

### GetYamlOk

`func (o *StackRevision) GetYamlOk() (*string, bool)`

GetYamlOk returns a tuple with the Yaml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYaml

`func (o *StackRevision) SetYaml(v string)`

SetYaml sets Yaml field to given value.

### HasYaml

`func (o *StackRevision) HasYaml() bool`

HasYaml returns a boolean if a field has been set.

### GetVersion

`func (o *StackRevision) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *StackRevision) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *StackRevision) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *StackRevision) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetParentSha

`func (o *StackRevision) GetParentSha() string`

GetParentSha returns the ParentSha field if non-nil, zero value otherwise.

### GetParentShaOk

`func (o *StackRevision) GetParentShaOk() (*string, bool)`

GetParentShaOk returns a tuple with the ParentSha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentSha

`func (o *StackRevision) SetParentSha(v string)`

SetParentSha sets ParentSha field to given value.

### HasParentSha

`func (o *StackRevision) HasParentSha() bool`

HasParentSha returns a boolean if a field has been set.

### GetStatus

`func (o *StackRevision) GetStatus() StackRevisionStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *StackRevision) GetStatusOk() (*StackRevisionStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *StackRevision) SetStatus(v StackRevisionStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *StackRevision) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusMessage

`func (o *StackRevision) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *StackRevision) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *StackRevision) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *StackRevision) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### GetCreatedAt

`func (o *StackRevision) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *StackRevision) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *StackRevision) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *StackRevision) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCommitInfo

`func (o *StackRevision) GetCommitInfo() CommitInfo`

GetCommitInfo returns the CommitInfo field if non-nil, zero value otherwise.

### GetCommitInfoOk

`func (o *StackRevision) GetCommitInfoOk() (*CommitInfo, bool)`

GetCommitInfoOk returns a tuple with the CommitInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitInfo

`func (o *StackRevision) SetCommitInfo(v CommitInfo)`

SetCommitInfo sets CommitInfo field to given value.

### HasCommitInfo

`func (o *StackRevision) HasCommitInfo() bool`

HasCommitInfo returns a boolean if a field has been set.

### GetReleaseId

`func (o *StackRevision) GetReleaseId() string`

GetReleaseId returns the ReleaseId field if non-nil, zero value otherwise.

### GetReleaseIdOk

`func (o *StackRevision) GetReleaseIdOk() (*string, bool)`

GetReleaseIdOk returns a tuple with the ReleaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseId

`func (o *StackRevision) SetReleaseId(v string)`

SetReleaseId sets ReleaseId field to given value.

### HasReleaseId

`func (o *StackRevision) HasReleaseId() bool`

HasReleaseId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


