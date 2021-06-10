# AutoRelease

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Groups** | Pointer to [**[]AutoReleaseGroup**](AutoReleaseGroup.md) |  | [optional] 

## Methods

### NewAutoRelease

`func NewAutoRelease() *AutoRelease`

NewAutoRelease instantiates a new AutoRelease object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoReleaseWithDefaults

`func NewAutoReleaseWithDefaults() *AutoRelease`

NewAutoReleaseWithDefaults instantiates a new AutoRelease object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroups

`func (o *AutoRelease) GetGroups() []AutoReleaseGroup`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *AutoRelease) GetGroupsOk() (*[]AutoReleaseGroup, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *AutoRelease) SetGroups(v []AutoReleaseGroup)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *AutoRelease) HasGroups() bool`

HasGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


