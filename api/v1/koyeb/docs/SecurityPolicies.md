# SecurityPolicies

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BasicAuths** | Pointer to [**[]BasicAuthPolicy**](BasicAuthPolicy.md) |  | [optional] 
**ApiKeys** | Pointer to **[]string** |  | [optional] 

## Methods

### NewSecurityPolicies

`func NewSecurityPolicies() *SecurityPolicies`

NewSecurityPolicies instantiates a new SecurityPolicies object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityPoliciesWithDefaults

`func NewSecurityPoliciesWithDefaults() *SecurityPolicies`

NewSecurityPoliciesWithDefaults instantiates a new SecurityPolicies object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBasicAuths

`func (o *SecurityPolicies) GetBasicAuths() []BasicAuthPolicy`

GetBasicAuths returns the BasicAuths field if non-nil, zero value otherwise.

### GetBasicAuthsOk

`func (o *SecurityPolicies) GetBasicAuthsOk() (*[]BasicAuthPolicy, bool)`

GetBasicAuthsOk returns a tuple with the BasicAuths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicAuths

`func (o *SecurityPolicies) SetBasicAuths(v []BasicAuthPolicy)`

SetBasicAuths sets BasicAuths field to given value.

### HasBasicAuths

`func (o *SecurityPolicies) HasBasicAuths() bool`

HasBasicAuths returns a boolean if a field has been set.

### GetApiKeys

`func (o *SecurityPolicies) GetApiKeys() []string`

GetApiKeys returns the ApiKeys field if non-nil, zero value otherwise.

### GetApiKeysOk

`func (o *SecurityPolicies) GetApiKeysOk() (*[]string, bool)`

GetApiKeysOk returns a tuple with the ApiKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKeys

`func (o *SecurityPolicies) SetApiKeys(v []string)`

SetApiKeys sets ApiKeys field to given value.

### HasApiKeys

`func (o *SecurityPolicies) HasApiKeys() bool`

HasApiKeys returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


