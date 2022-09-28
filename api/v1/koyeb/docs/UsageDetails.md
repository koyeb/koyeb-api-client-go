# UsageDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrganizationId** | Pointer to **string** |  | [optional] 
**InstanceId** | Pointer to **string** |  | [optional] 
**AppId** | Pointer to **string** |  | [optional] 
**AppName** | Pointer to **string** |  | [optional] 
**ServiceId** | Pointer to **string** |  | [optional] 
**ServiceName** | Pointer to **string** |  | [optional] 
**RegionalDeploymentId** | Pointer to **string** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**DeploymentId** | Pointer to **string** |  | [optional] 
**InstanceType** | Pointer to **string** |  | [optional] 
**DurationSeconds** | Pointer to **int64** |  | [optional] 
**StartedAt** | Pointer to **time.Time** |  | [optional] 
**TerminatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewUsageDetails

`func NewUsageDetails() *UsageDetails`

NewUsageDetails instantiates a new UsageDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageDetailsWithDefaults

`func NewUsageDetailsWithDefaults() *UsageDetails`

NewUsageDetailsWithDefaults instantiates a new UsageDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrganizationId

`func (o *UsageDetails) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *UsageDetails) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *UsageDetails) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *UsageDetails) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetInstanceId

`func (o *UsageDetails) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *UsageDetails) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *UsageDetails) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.

### HasInstanceId

`func (o *UsageDetails) HasInstanceId() bool`

HasInstanceId returns a boolean if a field has been set.

### GetAppId

`func (o *UsageDetails) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *UsageDetails) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *UsageDetails) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *UsageDetails) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### GetAppName

`func (o *UsageDetails) GetAppName() string`

GetAppName returns the AppName field if non-nil, zero value otherwise.

### GetAppNameOk

`func (o *UsageDetails) GetAppNameOk() (*string, bool)`

GetAppNameOk returns a tuple with the AppName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppName

`func (o *UsageDetails) SetAppName(v string)`

SetAppName sets AppName field to given value.

### HasAppName

`func (o *UsageDetails) HasAppName() bool`

HasAppName returns a boolean if a field has been set.

### GetServiceId

`func (o *UsageDetails) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *UsageDetails) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *UsageDetails) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *UsageDetails) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetServiceName

`func (o *UsageDetails) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *UsageDetails) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *UsageDetails) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *UsageDetails) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.

### GetRegionalDeploymentId

`func (o *UsageDetails) GetRegionalDeploymentId() string`

GetRegionalDeploymentId returns the RegionalDeploymentId field if non-nil, zero value otherwise.

### GetRegionalDeploymentIdOk

`func (o *UsageDetails) GetRegionalDeploymentIdOk() (*string, bool)`

GetRegionalDeploymentIdOk returns a tuple with the RegionalDeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionalDeploymentId

`func (o *UsageDetails) SetRegionalDeploymentId(v string)`

SetRegionalDeploymentId sets RegionalDeploymentId field to given value.

### HasRegionalDeploymentId

`func (o *UsageDetails) HasRegionalDeploymentId() bool`

HasRegionalDeploymentId returns a boolean if a field has been set.

### GetRegion

`func (o *UsageDetails) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *UsageDetails) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *UsageDetails) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *UsageDetails) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetDeploymentId

`func (o *UsageDetails) GetDeploymentId() string`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *UsageDetails) GetDeploymentIdOk() (*string, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *UsageDetails) SetDeploymentId(v string)`

SetDeploymentId sets DeploymentId field to given value.

### HasDeploymentId

`func (o *UsageDetails) HasDeploymentId() bool`

HasDeploymentId returns a boolean if a field has been set.

### GetInstanceType

`func (o *UsageDetails) GetInstanceType() string`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *UsageDetails) GetInstanceTypeOk() (*string, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *UsageDetails) SetInstanceType(v string)`

SetInstanceType sets InstanceType field to given value.

### HasInstanceType

`func (o *UsageDetails) HasInstanceType() bool`

HasInstanceType returns a boolean if a field has been set.

### GetDurationSeconds

`func (o *UsageDetails) GetDurationSeconds() int64`

GetDurationSeconds returns the DurationSeconds field if non-nil, zero value otherwise.

### GetDurationSecondsOk

`func (o *UsageDetails) GetDurationSecondsOk() (*int64, bool)`

GetDurationSecondsOk returns a tuple with the DurationSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationSeconds

`func (o *UsageDetails) SetDurationSeconds(v int64)`

SetDurationSeconds sets DurationSeconds field to given value.

### HasDurationSeconds

`func (o *UsageDetails) HasDurationSeconds() bool`

HasDurationSeconds returns a boolean if a field has been set.

### GetStartedAt

`func (o *UsageDetails) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *UsageDetails) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *UsageDetails) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *UsageDetails) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetTerminatedAt

`func (o *UsageDetails) GetTerminatedAt() time.Time`

GetTerminatedAt returns the TerminatedAt field if non-nil, zero value otherwise.

### GetTerminatedAtOk

`func (o *UsageDetails) GetTerminatedAtOk() (*time.Time, bool)`

GetTerminatedAtOk returns a tuple with the TerminatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminatedAt

`func (o *UsageDetails) SetTerminatedAt(v time.Time)`

SetTerminatedAt sets TerminatedAt field to given value.

### HasTerminatedAt

`func (o *UsageDetails) HasTerminatedAt() bool`

HasTerminatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


