# NotificationList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Notifications** | Pointer to [**[]Notification**](Notification.md) |  | [optional] 
**Limit** | Pointer to **int64** |  | [optional] 
**Offset** | Pointer to **int64** |  | [optional] 
**Count** | Pointer to **int64** |  | [optional] 
**IsRead** | Pointer to **bool** |  | [optional] 
**IsSeen** | Pointer to **bool** |  | [optional] 
**Unread** | Pointer to **int64** |  | [optional] 
**Unseen** | Pointer to **int64** |  | [optional] 

## Methods

### NewNotificationList

`func NewNotificationList() *NotificationList`

NewNotificationList instantiates a new NotificationList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationListWithDefaults

`func NewNotificationListWithDefaults() *NotificationList`

NewNotificationListWithDefaults instantiates a new NotificationList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotifications

`func (o *NotificationList) GetNotifications() []Notification`

GetNotifications returns the Notifications field if non-nil, zero value otherwise.

### GetNotificationsOk

`func (o *NotificationList) GetNotificationsOk() (*[]Notification, bool)`

GetNotificationsOk returns a tuple with the Notifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifications

`func (o *NotificationList) SetNotifications(v []Notification)`

SetNotifications sets Notifications field to given value.

### HasNotifications

`func (o *NotificationList) HasNotifications() bool`

HasNotifications returns a boolean if a field has been set.

### GetLimit

`func (o *NotificationList) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *NotificationList) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *NotificationList) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *NotificationList) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *NotificationList) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *NotificationList) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *NotificationList) SetOffset(v int64)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *NotificationList) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetCount

`func (o *NotificationList) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *NotificationList) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *NotificationList) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *NotificationList) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetIsRead

`func (o *NotificationList) GetIsRead() bool`

GetIsRead returns the IsRead field if non-nil, zero value otherwise.

### GetIsReadOk

`func (o *NotificationList) GetIsReadOk() (*bool, bool)`

GetIsReadOk returns a tuple with the IsRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRead

`func (o *NotificationList) SetIsRead(v bool)`

SetIsRead sets IsRead field to given value.

### HasIsRead

`func (o *NotificationList) HasIsRead() bool`

HasIsRead returns a boolean if a field has been set.

### GetIsSeen

`func (o *NotificationList) GetIsSeen() bool`

GetIsSeen returns the IsSeen field if non-nil, zero value otherwise.

### GetIsSeenOk

`func (o *NotificationList) GetIsSeenOk() (*bool, bool)`

GetIsSeenOk returns a tuple with the IsSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSeen

`func (o *NotificationList) SetIsSeen(v bool)`

SetIsSeen sets IsSeen field to given value.

### HasIsSeen

`func (o *NotificationList) HasIsSeen() bool`

HasIsSeen returns a boolean if a field has been set.

### GetUnread

`func (o *NotificationList) GetUnread() int64`

GetUnread returns the Unread field if non-nil, zero value otherwise.

### GetUnreadOk

`func (o *NotificationList) GetUnreadOk() (*int64, bool)`

GetUnreadOk returns a tuple with the Unread field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnread

`func (o *NotificationList) SetUnread(v int64)`

SetUnread sets Unread field to given value.

### HasUnread

`func (o *NotificationList) HasUnread() bool`

HasUnread returns a boolean if a field has been set.

### GetUnseen

`func (o *NotificationList) GetUnseen() int64`

GetUnseen returns the Unseen field if non-nil, zero value otherwise.

### GetUnseenOk

`func (o *NotificationList) GetUnseenOk() (*int64, bool)`

GetUnseenOk returns a tuple with the Unseen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnseen

`func (o *NotificationList) SetUnseen(v int64)`

SetUnseen sets Unseen field to given value.

### HasUnseen

`func (o *NotificationList) HasUnseen() bool`

HasUnseen returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


