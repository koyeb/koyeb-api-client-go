# EgressPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | Pointer to [**EgressPolicyMode**](EgressPolicyMode.md) |  | [optional] [default to EGRESSPOLICYMODE_DEFAULT]
**AllowList** | Pointer to [**[]NetworkPolicyDestination**](NetworkPolicyDestination.md) | Allowed destinations (deny-by-default semantics under DENY_ALL). Ignored when mode is DEFAULT. | [optional] 

## Methods

### NewEgressPolicy

`func NewEgressPolicy() *EgressPolicy`

NewEgressPolicy instantiates a new EgressPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEgressPolicyWithDefaults

`func NewEgressPolicyWithDefaults() *EgressPolicy`

NewEgressPolicyWithDefaults instantiates a new EgressPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *EgressPolicy) GetMode() EgressPolicyMode`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *EgressPolicy) GetModeOk() (*EgressPolicyMode, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *EgressPolicy) SetMode(v EgressPolicyMode)`

SetMode sets Mode field to given value.

### HasMode

`func (o *EgressPolicy) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetAllowList

`func (o *EgressPolicy) GetAllowList() []NetworkPolicyDestination`

GetAllowList returns the AllowList field if non-nil, zero value otherwise.

### GetAllowListOk

`func (o *EgressPolicy) GetAllowListOk() (*[]NetworkPolicyDestination, bool)`

GetAllowListOk returns a tuple with the AllowList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowList

`func (o *EgressPolicy) SetAllowList(v []NetworkPolicyDestination)`

SetAllowList sets AllowList field to given value.

### HasAllowList

`func (o *EgressPolicy) HasAllowList() bool`

HasAllowList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


