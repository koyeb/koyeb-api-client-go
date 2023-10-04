# BuildpackBuilder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BuildCommand** | Pointer to **string** |  | [optional] 
**RunCommand** | Pointer to **string** |  | [optional] 
**Privileged** | Pointer to **bool** |  | [optional] 

## Methods

### NewBuildpackBuilder

`func NewBuildpackBuilder() *BuildpackBuilder`

NewBuildpackBuilder instantiates a new BuildpackBuilder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBuildpackBuilderWithDefaults

`func NewBuildpackBuilderWithDefaults() *BuildpackBuilder`

NewBuildpackBuilderWithDefaults instantiates a new BuildpackBuilder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuildCommand

`func (o *BuildpackBuilder) GetBuildCommand() string`

GetBuildCommand returns the BuildCommand field if non-nil, zero value otherwise.

### GetBuildCommandOk

`func (o *BuildpackBuilder) GetBuildCommandOk() (*string, bool)`

GetBuildCommandOk returns a tuple with the BuildCommand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildCommand

`func (o *BuildpackBuilder) SetBuildCommand(v string)`

SetBuildCommand sets BuildCommand field to given value.

### HasBuildCommand

`func (o *BuildpackBuilder) HasBuildCommand() bool`

HasBuildCommand returns a boolean if a field has been set.

### GetRunCommand

`func (o *BuildpackBuilder) GetRunCommand() string`

GetRunCommand returns the RunCommand field if non-nil, zero value otherwise.

### GetRunCommandOk

`func (o *BuildpackBuilder) GetRunCommandOk() (*string, bool)`

GetRunCommandOk returns a tuple with the RunCommand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunCommand

`func (o *BuildpackBuilder) SetRunCommand(v string)`

SetRunCommand sets RunCommand field to given value.

### HasRunCommand

`func (o *BuildpackBuilder) HasRunCommand() bool`

HasRunCommand returns a boolean if a field has been set.

### GetPrivileged

`func (o *BuildpackBuilder) GetPrivileged() bool`

GetPrivileged returns the Privileged field if non-nil, zero value otherwise.

### GetPrivilegedOk

`func (o *BuildpackBuilder) GetPrivilegedOk() (*bool, bool)`

GetPrivilegedOk returns a tuple with the Privileged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivileged

`func (o *BuildpackBuilder) SetPrivileged(v bool)`

SetPrivileged sets Privileged field to given value.

### HasPrivileged

`func (o *BuildpackBuilder) HasPrivileged() bool`

HasPrivileged returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


