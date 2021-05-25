# Notification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Activity** | Pointer to [**Activity**](Activity.md) |  | [optional] 
**IsRead** | Pointer to **bool** |  | [optional] 
**IsSeen** | Pointer to **bool** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewNotification

`func NewNotification() *Notification`

NewNotification instantiates a new Notification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationWithDefaults

`func NewNotificationWithDefaults() *Notification`

NewNotificationWithDefaults instantiates a new Notification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Notification) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Notification) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Notification) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Notification) HasId() bool`

HasId returns a boolean if a field has been set.

### GetActivity

`func (o *Notification) GetActivity() Activity`

GetActivity returns the Activity field if non-nil, zero value otherwise.

### GetActivityOk

`func (o *Notification) GetActivityOk() (*Activity, bool)`

GetActivityOk returns a tuple with the Activity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivity

`func (o *Notification) SetActivity(v Activity)`

SetActivity sets Activity field to given value.

### HasActivity

`func (o *Notification) HasActivity() bool`

HasActivity returns a boolean if a field has been set.

### GetIsRead

`func (o *Notification) GetIsRead() bool`

GetIsRead returns the IsRead field if non-nil, zero value otherwise.

### GetIsReadOk

`func (o *Notification) GetIsReadOk() (*bool, bool)`

GetIsReadOk returns a tuple with the IsRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRead

`func (o *Notification) SetIsRead(v bool)`

SetIsRead sets IsRead field to given value.

### HasIsRead

`func (o *Notification) HasIsRead() bool`

HasIsRead returns a boolean if a field has been set.

### GetIsSeen

`func (o *Notification) GetIsSeen() bool`

GetIsSeen returns the IsSeen field if non-nil, zero value otherwise.

### GetIsSeenOk

`func (o *Notification) GetIsSeenOk() (*bool, bool)`

GetIsSeenOk returns a tuple with the IsSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSeen

`func (o *Notification) SetIsSeen(v bool)`

SetIsSeen sets IsSeen field to given value.

### HasIsSeen

`func (o *Notification) HasIsSeen() bool`

HasIsSeen returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Notification) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Notification) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Notification) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Notification) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


