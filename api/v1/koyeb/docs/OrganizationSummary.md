# OrganizationSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrganizationId** | Pointer to **string** |  | [optional] 
**Instances** | Pointer to [**InstancesSummary**](InstancesSummary.md) |  | [optional] 
**Apps** | Pointer to [**AppsSummary**](AppsSummary.md) |  | [optional] 
**Services** | Pointer to [**map[string]ServiceSummary**](ServiceSummary.md) |  | [optional] 
**Domains** | Pointer to [**DomainsSummary**](DomainsSummary.md) |  | [optional] 
**Secrets** | Pointer to [**SecretsSummary**](SecretsSummary.md) |  | [optional] 
**NeonPostgres** | Pointer to [**NeonPostgresSummary**](NeonPostgresSummary.md) |  | [optional] 
**Members** | Pointer to [**MembersSummary**](MembersSummary.md) |  | [optional] 

## Methods

### NewOrganizationSummary

`func NewOrganizationSummary() *OrganizationSummary`

NewOrganizationSummary instantiates a new OrganizationSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationSummaryWithDefaults

`func NewOrganizationSummaryWithDefaults() *OrganizationSummary`

NewOrganizationSummaryWithDefaults instantiates a new OrganizationSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrganizationId

`func (o *OrganizationSummary) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *OrganizationSummary) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *OrganizationSummary) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *OrganizationSummary) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetInstances

`func (o *OrganizationSummary) GetInstances() InstancesSummary`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *OrganizationSummary) GetInstancesOk() (*InstancesSummary, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstances

`func (o *OrganizationSummary) SetInstances(v InstancesSummary)`

SetInstances sets Instances field to given value.

### HasInstances

`func (o *OrganizationSummary) HasInstances() bool`

HasInstances returns a boolean if a field has been set.

### GetApps

`func (o *OrganizationSummary) GetApps() AppsSummary`

GetApps returns the Apps field if non-nil, zero value otherwise.

### GetAppsOk

`func (o *OrganizationSummary) GetAppsOk() (*AppsSummary, bool)`

GetAppsOk returns a tuple with the Apps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApps

`func (o *OrganizationSummary) SetApps(v AppsSummary)`

SetApps sets Apps field to given value.

### HasApps

`func (o *OrganizationSummary) HasApps() bool`

HasApps returns a boolean if a field has been set.

### GetServices

`func (o *OrganizationSummary) GetServices() map[string]ServiceSummary`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *OrganizationSummary) GetServicesOk() (*map[string]ServiceSummary, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *OrganizationSummary) SetServices(v map[string]ServiceSummary)`

SetServices sets Services field to given value.

### HasServices

`func (o *OrganizationSummary) HasServices() bool`

HasServices returns a boolean if a field has been set.

### GetDomains

`func (o *OrganizationSummary) GetDomains() DomainsSummary`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *OrganizationSummary) GetDomainsOk() (*DomainsSummary, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *OrganizationSummary) SetDomains(v DomainsSummary)`

SetDomains sets Domains field to given value.

### HasDomains

`func (o *OrganizationSummary) HasDomains() bool`

HasDomains returns a boolean if a field has been set.

### GetSecrets

`func (o *OrganizationSummary) GetSecrets() SecretsSummary`

GetSecrets returns the Secrets field if non-nil, zero value otherwise.

### GetSecretsOk

`func (o *OrganizationSummary) GetSecretsOk() (*SecretsSummary, bool)`

GetSecretsOk returns a tuple with the Secrets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecrets

`func (o *OrganizationSummary) SetSecrets(v SecretsSummary)`

SetSecrets sets Secrets field to given value.

### HasSecrets

`func (o *OrganizationSummary) HasSecrets() bool`

HasSecrets returns a boolean if a field has been set.

### GetNeonPostgres

`func (o *OrganizationSummary) GetNeonPostgres() NeonPostgresSummary`

GetNeonPostgres returns the NeonPostgres field if non-nil, zero value otherwise.

### GetNeonPostgresOk

`func (o *OrganizationSummary) GetNeonPostgresOk() (*NeonPostgresSummary, bool)`

GetNeonPostgresOk returns a tuple with the NeonPostgres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeonPostgres

`func (o *OrganizationSummary) SetNeonPostgres(v NeonPostgresSummary)`

SetNeonPostgres sets NeonPostgres field to given value.

### HasNeonPostgres

`func (o *OrganizationSummary) HasNeonPostgres() bool`

HasNeonPostgres returns a boolean if a field has been set.

### GetMembers

`func (o *OrganizationSummary) GetMembers() MembersSummary`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *OrganizationSummary) GetMembersOk() (*MembersSummary, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *OrganizationSummary) SetMembers(v MembersSummary)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *OrganizationSummary) HasMembers() bool`

HasMembers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


