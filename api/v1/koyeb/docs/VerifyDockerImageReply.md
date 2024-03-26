# VerifyDockerImageReply

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**Reason** | Pointer to **string** |  | [optional] 
**Code** | Pointer to [**VerifyDockerImageReplyErrCode**](VerifyDockerImageReplyErrCode.md) |  | [optional] [default to VERIFYDOCKERIMAGEREPLYERRCODE_UNKNOWN]

## Methods

### NewVerifyDockerImageReply

`func NewVerifyDockerImageReply() *VerifyDockerImageReply`

NewVerifyDockerImageReply instantiates a new VerifyDockerImageReply object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerifyDockerImageReplyWithDefaults

`func NewVerifyDockerImageReplyWithDefaults() *VerifyDockerImageReply`

NewVerifyDockerImageReplyWithDefaults instantiates a new VerifyDockerImageReply object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *VerifyDockerImageReply) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *VerifyDockerImageReply) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *VerifyDockerImageReply) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *VerifyDockerImageReply) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetReason

`func (o *VerifyDockerImageReply) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *VerifyDockerImageReply) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *VerifyDockerImageReply) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *VerifyDockerImageReply) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetCode

`func (o *VerifyDockerImageReply) GetCode() VerifyDockerImageReplyErrCode`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *VerifyDockerImageReply) GetCodeOk() (*VerifyDockerImageReplyErrCode, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *VerifyDockerImageReply) SetCode(v VerifyDockerImageReplyErrCode)`

SetCode sets Code field to given value.

### HasCode

`func (o *VerifyDockerImageReply) HasCode() bool`

HasCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


