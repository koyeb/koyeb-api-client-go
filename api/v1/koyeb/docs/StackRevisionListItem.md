# StackRevisionListItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sha** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**ParentSha** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**StackRevisionStatus**](StackRevisionStatus.md) |  | [optional] [default to STACKREVISIONSTATUS_UNKNOWN]
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**CommitInfo** | Pointer to [**CommitInfo**](CommitInfo.md) |  | [optional] 
**ReleaseId** | Pointer to **string** |  | [optional] 

## Methods

### NewStackRevisionListItem

`func NewStackRevisionListItem() *StackRevisionListItem`

NewStackRevisionListItem instantiates a new StackRevisionListItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStackRevisionListItemWithDefaults

`func NewStackRevisionListItemWithDefaults() *StackRevisionListItem`

NewStackRevisionListItemWithDefaults instantiates a new StackRevisionListItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSha

`func (o *StackRevisionListItem) GetSha() string`

GetSha returns the Sha field if non-nil, zero value otherwise.

### GetShaOk

`func (o *StackRevisionListItem) GetShaOk() (*string, bool)`

GetShaOk returns a tuple with the Sha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha

`func (o *StackRevisionListItem) SetSha(v string)`

SetSha sets Sha field to given value.

### HasSha

`func (o *StackRevisionListItem) HasSha() bool`

HasSha returns a boolean if a field has been set.

### GetVersion

`func (o *StackRevisionListItem) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *StackRevisionListItem) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *StackRevisionListItem) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *StackRevisionListItem) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetParentSha

`func (o *StackRevisionListItem) GetParentSha() string`

GetParentSha returns the ParentSha field if non-nil, zero value otherwise.

### GetParentShaOk

`func (o *StackRevisionListItem) GetParentShaOk() (*string, bool)`

GetParentShaOk returns a tuple with the ParentSha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentSha

`func (o *StackRevisionListItem) SetParentSha(v string)`

SetParentSha sets ParentSha field to given value.

### HasParentSha

`func (o *StackRevisionListItem) HasParentSha() bool`

HasParentSha returns a boolean if a field has been set.

### GetStatus

`func (o *StackRevisionListItem) GetStatus() StackRevisionStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *StackRevisionListItem) GetStatusOk() (*StackRevisionStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *StackRevisionListItem) SetStatus(v StackRevisionStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *StackRevisionListItem) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCreatedAt

`func (o *StackRevisionListItem) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *StackRevisionListItem) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *StackRevisionListItem) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *StackRevisionListItem) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCommitInfo

`func (o *StackRevisionListItem) GetCommitInfo() CommitInfo`

GetCommitInfo returns the CommitInfo field if non-nil, zero value otherwise.

### GetCommitInfoOk

`func (o *StackRevisionListItem) GetCommitInfoOk() (*CommitInfo, bool)`

GetCommitInfoOk returns a tuple with the CommitInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitInfo

`func (o *StackRevisionListItem) SetCommitInfo(v CommitInfo)`

SetCommitInfo sets CommitInfo field to given value.

### HasCommitInfo

`func (o *StackRevisionListItem) HasCommitInfo() bool`

HasCommitInfo returns a boolean if a field has been set.

### GetReleaseId

`func (o *StackRevisionListItem) GetReleaseId() string`

GetReleaseId returns the ReleaseId field if non-nil, zero value otherwise.

### GetReleaseIdOk

`func (o *StackRevisionListItem) GetReleaseIdOk() (*string, bool)`

GetReleaseIdOk returns a tuple with the ReleaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseId

`func (o *StackRevisionListItem) SetReleaseId(v string)`

SetReleaseId sets ReleaseId field to given value.

### HasReleaseId

`func (o *StackRevisionListItem) HasReleaseId() bool`

HasReleaseId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


