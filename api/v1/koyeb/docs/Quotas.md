# Quotas

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Apps** | Pointer to **string** |  | [optional] 
**Services** | Pointer to **string** |  | [optional] 
**Domains** | Pointer to **string** |  | [optional] 
**ServicesByApp** | Pointer to **string** |  | [optional] 
**ServiceProvisioningConcurrency** | Pointer to **string** |  | [optional] 
**MemoryMb** | Pointer to **string** |  | [optional] 
**InstanceTypes** | Pointer to **[]string** |  | [optional] 
**Regions** | Pointer to **[]string** |  | [optional] 
**MaxOrganizationMembers** | Pointer to **string** |  | [optional] 
**MaxInstancesByType** | Pointer to **map[string]string** |  | [optional] 
**PersistentVolumesByRegion** | Pointer to [**map[string]PersistentVolumeQuotas**](PersistentVolumeQuotas.md) |  | [optional] 
**CustomDomains** | Pointer to **string** |  | [optional] 
**DomainsLoadBalancer** | Pointer to [**DomainLoadBalancerQuotas**](DomainLoadBalancerQuotas.md) |  | [optional] 
**MetricsRetention** | Pointer to **int32** |  | [optional] 
**LogsRetention** | Pointer to **int32** |  | [optional] 
**AccessReservedSubdomains** | Pointer to **[]string** |  | [optional] 

## Methods

### NewQuotas

`func NewQuotas() *Quotas`

NewQuotas instantiates a new Quotas object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuotasWithDefaults

`func NewQuotasWithDefaults() *Quotas`

NewQuotasWithDefaults instantiates a new Quotas object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApps

`func (o *Quotas) GetApps() string`

GetApps returns the Apps field if non-nil, zero value otherwise.

### GetAppsOk

`func (o *Quotas) GetAppsOk() (*string, bool)`

GetAppsOk returns a tuple with the Apps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApps

`func (o *Quotas) SetApps(v string)`

SetApps sets Apps field to given value.

### HasApps

`func (o *Quotas) HasApps() bool`

HasApps returns a boolean if a field has been set.

### GetServices

`func (o *Quotas) GetServices() string`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *Quotas) GetServicesOk() (*string, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *Quotas) SetServices(v string)`

SetServices sets Services field to given value.

### HasServices

`func (o *Quotas) HasServices() bool`

HasServices returns a boolean if a field has been set.

### GetDomains

`func (o *Quotas) GetDomains() string`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *Quotas) GetDomainsOk() (*string, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *Quotas) SetDomains(v string)`

SetDomains sets Domains field to given value.

### HasDomains

`func (o *Quotas) HasDomains() bool`

HasDomains returns a boolean if a field has been set.

### GetServicesByApp

`func (o *Quotas) GetServicesByApp() string`

GetServicesByApp returns the ServicesByApp field if non-nil, zero value otherwise.

### GetServicesByAppOk

`func (o *Quotas) GetServicesByAppOk() (*string, bool)`

GetServicesByAppOk returns a tuple with the ServicesByApp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicesByApp

`func (o *Quotas) SetServicesByApp(v string)`

SetServicesByApp sets ServicesByApp field to given value.

### HasServicesByApp

`func (o *Quotas) HasServicesByApp() bool`

HasServicesByApp returns a boolean if a field has been set.

### GetServiceProvisioningConcurrency

`func (o *Quotas) GetServiceProvisioningConcurrency() string`

GetServiceProvisioningConcurrency returns the ServiceProvisioningConcurrency field if non-nil, zero value otherwise.

### GetServiceProvisioningConcurrencyOk

`func (o *Quotas) GetServiceProvisioningConcurrencyOk() (*string, bool)`

GetServiceProvisioningConcurrencyOk returns a tuple with the ServiceProvisioningConcurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceProvisioningConcurrency

`func (o *Quotas) SetServiceProvisioningConcurrency(v string)`

SetServiceProvisioningConcurrency sets ServiceProvisioningConcurrency field to given value.

### HasServiceProvisioningConcurrency

`func (o *Quotas) HasServiceProvisioningConcurrency() bool`

HasServiceProvisioningConcurrency returns a boolean if a field has been set.

### GetMemoryMb

`func (o *Quotas) GetMemoryMb() string`

GetMemoryMb returns the MemoryMb field if non-nil, zero value otherwise.

### GetMemoryMbOk

`func (o *Quotas) GetMemoryMbOk() (*string, bool)`

GetMemoryMbOk returns a tuple with the MemoryMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryMb

`func (o *Quotas) SetMemoryMb(v string)`

SetMemoryMb sets MemoryMb field to given value.

### HasMemoryMb

`func (o *Quotas) HasMemoryMb() bool`

HasMemoryMb returns a boolean if a field has been set.

### GetInstanceTypes

`func (o *Quotas) GetInstanceTypes() []string`

GetInstanceTypes returns the InstanceTypes field if non-nil, zero value otherwise.

### GetInstanceTypesOk

`func (o *Quotas) GetInstanceTypesOk() (*[]string, bool)`

GetInstanceTypesOk returns a tuple with the InstanceTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceTypes

`func (o *Quotas) SetInstanceTypes(v []string)`

SetInstanceTypes sets InstanceTypes field to given value.

### HasInstanceTypes

`func (o *Quotas) HasInstanceTypes() bool`

HasInstanceTypes returns a boolean if a field has been set.

### GetRegions

`func (o *Quotas) GetRegions() []string`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *Quotas) GetRegionsOk() (*[]string, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *Quotas) SetRegions(v []string)`

SetRegions sets Regions field to given value.

### HasRegions

`func (o *Quotas) HasRegions() bool`

HasRegions returns a boolean if a field has been set.

### GetMaxOrganizationMembers

`func (o *Quotas) GetMaxOrganizationMembers() string`

GetMaxOrganizationMembers returns the MaxOrganizationMembers field if non-nil, zero value otherwise.

### GetMaxOrganizationMembersOk

`func (o *Quotas) GetMaxOrganizationMembersOk() (*string, bool)`

GetMaxOrganizationMembersOk returns a tuple with the MaxOrganizationMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxOrganizationMembers

`func (o *Quotas) SetMaxOrganizationMembers(v string)`

SetMaxOrganizationMembers sets MaxOrganizationMembers field to given value.

### HasMaxOrganizationMembers

`func (o *Quotas) HasMaxOrganizationMembers() bool`

HasMaxOrganizationMembers returns a boolean if a field has been set.

### GetMaxInstancesByType

`func (o *Quotas) GetMaxInstancesByType() map[string]string`

GetMaxInstancesByType returns the MaxInstancesByType field if non-nil, zero value otherwise.

### GetMaxInstancesByTypeOk

`func (o *Quotas) GetMaxInstancesByTypeOk() (*map[string]string, bool)`

GetMaxInstancesByTypeOk returns a tuple with the MaxInstancesByType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxInstancesByType

`func (o *Quotas) SetMaxInstancesByType(v map[string]string)`

SetMaxInstancesByType sets MaxInstancesByType field to given value.

### HasMaxInstancesByType

`func (o *Quotas) HasMaxInstancesByType() bool`

HasMaxInstancesByType returns a boolean if a field has been set.

### GetPersistentVolumesByRegion

`func (o *Quotas) GetPersistentVolumesByRegion() map[string]PersistentVolumeQuotas`

GetPersistentVolumesByRegion returns the PersistentVolumesByRegion field if non-nil, zero value otherwise.

### GetPersistentVolumesByRegionOk

`func (o *Quotas) GetPersistentVolumesByRegionOk() (*map[string]PersistentVolumeQuotas, bool)`

GetPersistentVolumesByRegionOk returns a tuple with the PersistentVolumesByRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistentVolumesByRegion

`func (o *Quotas) SetPersistentVolumesByRegion(v map[string]PersistentVolumeQuotas)`

SetPersistentVolumesByRegion sets PersistentVolumesByRegion field to given value.

### HasPersistentVolumesByRegion

`func (o *Quotas) HasPersistentVolumesByRegion() bool`

HasPersistentVolumesByRegion returns a boolean if a field has been set.

### GetCustomDomains

`func (o *Quotas) GetCustomDomains() string`

GetCustomDomains returns the CustomDomains field if non-nil, zero value otherwise.

### GetCustomDomainsOk

`func (o *Quotas) GetCustomDomainsOk() (*string, bool)`

GetCustomDomainsOk returns a tuple with the CustomDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomDomains

`func (o *Quotas) SetCustomDomains(v string)`

SetCustomDomains sets CustomDomains field to given value.

### HasCustomDomains

`func (o *Quotas) HasCustomDomains() bool`

HasCustomDomains returns a boolean if a field has been set.

### GetDomainsLoadBalancer

`func (o *Quotas) GetDomainsLoadBalancer() DomainLoadBalancerQuotas`

GetDomainsLoadBalancer returns the DomainsLoadBalancer field if non-nil, zero value otherwise.

### GetDomainsLoadBalancerOk

`func (o *Quotas) GetDomainsLoadBalancerOk() (*DomainLoadBalancerQuotas, bool)`

GetDomainsLoadBalancerOk returns a tuple with the DomainsLoadBalancer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainsLoadBalancer

`func (o *Quotas) SetDomainsLoadBalancer(v DomainLoadBalancerQuotas)`

SetDomainsLoadBalancer sets DomainsLoadBalancer field to given value.

### HasDomainsLoadBalancer

`func (o *Quotas) HasDomainsLoadBalancer() bool`

HasDomainsLoadBalancer returns a boolean if a field has been set.

### GetMetricsRetention

`func (o *Quotas) GetMetricsRetention() int32`

GetMetricsRetention returns the MetricsRetention field if non-nil, zero value otherwise.

### GetMetricsRetentionOk

`func (o *Quotas) GetMetricsRetentionOk() (*int32, bool)`

GetMetricsRetentionOk returns a tuple with the MetricsRetention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricsRetention

`func (o *Quotas) SetMetricsRetention(v int32)`

SetMetricsRetention sets MetricsRetention field to given value.

### HasMetricsRetention

`func (o *Quotas) HasMetricsRetention() bool`

HasMetricsRetention returns a boolean if a field has been set.

### GetLogsRetention

`func (o *Quotas) GetLogsRetention() int32`

GetLogsRetention returns the LogsRetention field if non-nil, zero value otherwise.

### GetLogsRetentionOk

`func (o *Quotas) GetLogsRetentionOk() (*int32, bool)`

GetLogsRetentionOk returns a tuple with the LogsRetention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogsRetention

`func (o *Quotas) SetLogsRetention(v int32)`

SetLogsRetention sets LogsRetention field to given value.

### HasLogsRetention

`func (o *Quotas) HasLogsRetention() bool`

HasLogsRetention returns a boolean if a field has been set.

### GetAccessReservedSubdomains

`func (o *Quotas) GetAccessReservedSubdomains() []string`

GetAccessReservedSubdomains returns the AccessReservedSubdomains field if non-nil, zero value otherwise.

### GetAccessReservedSubdomainsOk

`func (o *Quotas) GetAccessReservedSubdomainsOk() (*[]string, bool)`

GetAccessReservedSubdomainsOk returns a tuple with the AccessReservedSubdomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessReservedSubdomains

`func (o *Quotas) SetAccessReservedSubdomains(v []string)`

SetAccessReservedSubdomains sets AccessReservedSubdomains field to given value.

### HasAccessReservedSubdomains

`func (o *Quotas) HasAccessReservedSubdomains() bool`

HasAccessReservedSubdomains returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


