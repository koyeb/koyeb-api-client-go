# DatabaseUsageDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrganizationId** | Pointer to **string** |  | [optional] 
**AppId** | Pointer to **string** |  | [optional] 
**AppName** | Pointer to **string** |  | [optional] 
**ServiceId** | Pointer to **string** |  | [optional] 
**ServiceName** | Pointer to **string** |  | [optional] 
**ComputeTimeSeconds** | Pointer to **int64** |  | [optional] 
**DataStorageMegabytesHour** | Pointer to **int64** |  | [optional] 
**StartedAt** | Pointer to **time.Time** |  | [optional] 
**TerminatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewDatabaseUsageDetails

`func NewDatabaseUsageDetails() *DatabaseUsageDetails`

NewDatabaseUsageDetails instantiates a new DatabaseUsageDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatabaseUsageDetailsWithDefaults

`func NewDatabaseUsageDetailsWithDefaults() *DatabaseUsageDetails`

NewDatabaseUsageDetailsWithDefaults instantiates a new DatabaseUsageDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrganizationId

`func (o *DatabaseUsageDetails) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *DatabaseUsageDetails) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *DatabaseUsageDetails) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *DatabaseUsageDetails) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetAppId

`func (o *DatabaseUsageDetails) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *DatabaseUsageDetails) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *DatabaseUsageDetails) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *DatabaseUsageDetails) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### GetAppName

`func (o *DatabaseUsageDetails) GetAppName() string`

GetAppName returns the AppName field if non-nil, zero value otherwise.

### GetAppNameOk

`func (o *DatabaseUsageDetails) GetAppNameOk() (*string, bool)`

GetAppNameOk returns a tuple with the AppName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppName

`func (o *DatabaseUsageDetails) SetAppName(v string)`

SetAppName sets AppName field to given value.

### HasAppName

`func (o *DatabaseUsageDetails) HasAppName() bool`

HasAppName returns a boolean if a field has been set.

### GetServiceId

`func (o *DatabaseUsageDetails) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *DatabaseUsageDetails) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *DatabaseUsageDetails) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *DatabaseUsageDetails) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetServiceName

`func (o *DatabaseUsageDetails) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *DatabaseUsageDetails) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *DatabaseUsageDetails) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *DatabaseUsageDetails) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.

### GetComputeTimeSeconds

`func (o *DatabaseUsageDetails) GetComputeTimeSeconds() int64`

GetComputeTimeSeconds returns the ComputeTimeSeconds field if non-nil, zero value otherwise.

### GetComputeTimeSecondsOk

`func (o *DatabaseUsageDetails) GetComputeTimeSecondsOk() (*int64, bool)`

GetComputeTimeSecondsOk returns a tuple with the ComputeTimeSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeTimeSeconds

`func (o *DatabaseUsageDetails) SetComputeTimeSeconds(v int64)`

SetComputeTimeSeconds sets ComputeTimeSeconds field to given value.

### HasComputeTimeSeconds

`func (o *DatabaseUsageDetails) HasComputeTimeSeconds() bool`

HasComputeTimeSeconds returns a boolean if a field has been set.

### GetDataStorageMegabytesHour

`func (o *DatabaseUsageDetails) GetDataStorageMegabytesHour() int64`

GetDataStorageMegabytesHour returns the DataStorageMegabytesHour field if non-nil, zero value otherwise.

### GetDataStorageMegabytesHourOk

`func (o *DatabaseUsageDetails) GetDataStorageMegabytesHourOk() (*int64, bool)`

GetDataStorageMegabytesHourOk returns a tuple with the DataStorageMegabytesHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataStorageMegabytesHour

`func (o *DatabaseUsageDetails) SetDataStorageMegabytesHour(v int64)`

SetDataStorageMegabytesHour sets DataStorageMegabytesHour field to given value.

### HasDataStorageMegabytesHour

`func (o *DatabaseUsageDetails) HasDataStorageMegabytesHour() bool`

HasDataStorageMegabytesHour returns a boolean if a field has been set.

### GetStartedAt

`func (o *DatabaseUsageDetails) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *DatabaseUsageDetails) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *DatabaseUsageDetails) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *DatabaseUsageDetails) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetTerminatedAt

`func (o *DatabaseUsageDetails) GetTerminatedAt() time.Time`

GetTerminatedAt returns the TerminatedAt field if non-nil, zero value otherwise.

### GetTerminatedAtOk

`func (o *DatabaseUsageDetails) GetTerminatedAtOk() (*time.Time, bool)`

GetTerminatedAtOk returns a tuple with the TerminatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminatedAt

`func (o *DatabaseUsageDetails) SetTerminatedAt(v time.Time)`

SetTerminatedAt sets TerminatedAt field to given value.

### HasTerminatedAt

`func (o *DatabaseUsageDetails) HasTerminatedAt() bool`

HasTerminatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


