# AppUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppId** | Pointer to **string** |  | [optional] 
**AppName** | Pointer to **string** |  | [optional] 
**Services** | Pointer to [**[]ServiceUsage**](ServiceUsage.md) |  | [optional] 

## Methods

### NewAppUsage

`func NewAppUsage() *AppUsage`

NewAppUsage instantiates a new AppUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppUsageWithDefaults

`func NewAppUsageWithDefaults() *AppUsage`

NewAppUsageWithDefaults instantiates a new AppUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppId

`func (o *AppUsage) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *AppUsage) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *AppUsage) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *AppUsage) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### GetAppName

`func (o *AppUsage) GetAppName() string`

GetAppName returns the AppName field if non-nil, zero value otherwise.

### GetAppNameOk

`func (o *AppUsage) GetAppNameOk() (*string, bool)`

GetAppNameOk returns a tuple with the AppName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppName

`func (o *AppUsage) SetAppName(v string)`

SetAppName sets AppName field to given value.

### HasAppName

`func (o *AppUsage) HasAppName() bool`

HasAppName returns a boolean if a field has been set.

### GetServices

`func (o *AppUsage) GetServices() []ServiceUsage`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *AppUsage) GetServicesOk() (*[]ServiceUsage, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *AppUsage) SetServices(v []ServiceUsage)`

SetServices sets Services field to given value.

### HasServices

`func (o *AppUsage) HasServices() bool`

HasServices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


