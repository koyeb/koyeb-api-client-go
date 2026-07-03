# NetworkPolicyDestination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cidr** | Pointer to **string** | IPv4 or IPv6 CIDR (e.g. \&quot;10.0.0.0/8\&quot;, \&quot;2001:db8::/32\&quot;). Bare IPs are accepted at the API boundary and normalized to /32 (IPv4) or /128 (IPv6) before storage. | [optional] 

## Methods

### NewNetworkPolicyDestination

`func NewNetworkPolicyDestination() *NetworkPolicyDestination`

NewNetworkPolicyDestination instantiates a new NetworkPolicyDestination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkPolicyDestinationWithDefaults

`func NewNetworkPolicyDestinationWithDefaults() *NetworkPolicyDestination`

NewNetworkPolicyDestinationWithDefaults instantiates a new NetworkPolicyDestination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCidr

`func (o *NetworkPolicyDestination) GetCidr() string`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *NetworkPolicyDestination) GetCidrOk() (*string, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *NetworkPolicyDestination) SetCidr(v string)`

SetCidr sets Cidr field to given value.

### HasCidr

`func (o *NetworkPolicyDestination) HasCidr() bool`

HasCidr returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


