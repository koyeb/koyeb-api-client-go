# ServiceViewSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**ServiceViewSourceType**](ServiceViewSourceType.md) |  | [optional] [default to SERVICEVIEWSOURCETYPE_UNKNOWN]
**Reference** | Pointer to **string** |  | [optional] 

## Methods

### NewServiceViewSource

`func NewServiceViewSource() *ServiceViewSource`

NewServiceViewSource instantiates a new ServiceViewSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceViewSourceWithDefaults

`func NewServiceViewSourceWithDefaults() *ServiceViewSource`

NewServiceViewSourceWithDefaults instantiates a new ServiceViewSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ServiceViewSource) GetType() ServiceViewSourceType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ServiceViewSource) GetTypeOk() (*ServiceViewSourceType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ServiceViewSource) SetType(v ServiceViewSourceType)`

SetType sets Type field to given value.

### HasType

`func (o *ServiceViewSource) HasType() bool`

HasType returns a boolean if a field has been set.

### GetReference

`func (o *ServiceViewSource) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *ServiceViewSource) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *ServiceViewSource) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *ServiceViewSource) HasReference() bool`

HasReference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


