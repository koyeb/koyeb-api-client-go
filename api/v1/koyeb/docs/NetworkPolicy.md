# NetworkPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Egress** | Pointer to [**EgressPolicy**](EgressPolicy.md) |  | [optional] 

## Methods

### NewNetworkPolicy

`func NewNetworkPolicy() *NetworkPolicy`

NewNetworkPolicy instantiates a new NetworkPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkPolicyWithDefaults

`func NewNetworkPolicyWithDefaults() *NetworkPolicy`

NewNetworkPolicyWithDefaults instantiates a new NetworkPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEgress

`func (o *NetworkPolicy) GetEgress() EgressPolicy`

GetEgress returns the Egress field if non-nil, zero value otherwise.

### GetEgressOk

`func (o *NetworkPolicy) GetEgressOk() (*EgressPolicy, bool)`

GetEgressOk returns a tuple with the Egress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEgress

`func (o *NetworkPolicy) SetEgress(v EgressPolicy)`

SetEgress sets Egress field to given value.

### HasEgress

`func (o *NetworkPolicy) HasEgress() bool`

HasEgress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


