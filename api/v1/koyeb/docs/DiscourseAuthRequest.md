# DiscourseAuthRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Payload** | Pointer to **string** |  | [optional] 
**Sig** | Pointer to **string** |  | [optional] 

## Methods

### NewDiscourseAuthRequest

`func NewDiscourseAuthRequest() *DiscourseAuthRequest`

NewDiscourseAuthRequest instantiates a new DiscourseAuthRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiscourseAuthRequestWithDefaults

`func NewDiscourseAuthRequestWithDefaults() *DiscourseAuthRequest`

NewDiscourseAuthRequestWithDefaults instantiates a new DiscourseAuthRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPayload

`func (o *DiscourseAuthRequest) GetPayload() string`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *DiscourseAuthRequest) GetPayloadOk() (*string, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *DiscourseAuthRequest) SetPayload(v string)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *DiscourseAuthRequest) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### GetSig

`func (o *DiscourseAuthRequest) GetSig() string`

GetSig returns the Sig field if non-nil, zero value otherwise.

### GetSigOk

`func (o *DiscourseAuthRequest) GetSigOk() (*string, bool)`

GetSigOk returns a tuple with the Sig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSig

`func (o *DiscourseAuthRequest) SetSig(v string)`

SetSig sets Sig field to given value.

### HasSig

`func (o *DiscourseAuthRequest) HasSig() bool`

HasSig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


