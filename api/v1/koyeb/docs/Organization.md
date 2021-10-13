# Organization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address1** | Pointer to **string** |  | [optional] 
**Address2** | Pointer to **string** |  | [optional] 
**City** | Pointer to **string** |  | [optional] 
**PostalCode** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**Country** | Pointer to **string** |  | [optional] 
**Company** | Pointer to **bool** |  | [optional] 
**VatNumber** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Plan** | Pointer to [**Plan**](Plan.md) |  | [optional] [default to PLAN_BASIC]

## Methods

### NewOrganization

`func NewOrganization() *Organization`

NewOrganization instantiates a new Organization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationWithDefaults

`func NewOrganizationWithDefaults() *Organization`

NewOrganizationWithDefaults instantiates a new Organization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress1

`func (o *Organization) GetAddress1() string`

GetAddress1 returns the Address1 field if non-nil, zero value otherwise.

### GetAddress1Ok

`func (o *Organization) GetAddress1Ok() (*string, bool)`

GetAddress1Ok returns a tuple with the Address1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress1

`func (o *Organization) SetAddress1(v string)`

SetAddress1 sets Address1 field to given value.

### HasAddress1

`func (o *Organization) HasAddress1() bool`

HasAddress1 returns a boolean if a field has been set.

### GetAddress2

`func (o *Organization) GetAddress2() string`

GetAddress2 returns the Address2 field if non-nil, zero value otherwise.

### GetAddress2Ok

`func (o *Organization) GetAddress2Ok() (*string, bool)`

GetAddress2Ok returns a tuple with the Address2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress2

`func (o *Organization) SetAddress2(v string)`

SetAddress2 sets Address2 field to given value.

### HasAddress2

`func (o *Organization) HasAddress2() bool`

HasAddress2 returns a boolean if a field has been set.

### GetCity

`func (o *Organization) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *Organization) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *Organization) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *Organization) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetPostalCode

`func (o *Organization) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *Organization) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *Organization) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *Organization) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetState

`func (o *Organization) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Organization) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Organization) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *Organization) HasState() bool`

HasState returns a boolean if a field has been set.

### GetCountry

`func (o *Organization) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *Organization) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *Organization) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *Organization) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCompany

`func (o *Organization) GetCompany() bool`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *Organization) GetCompanyOk() (*bool, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *Organization) SetCompany(v bool)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *Organization) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetVatNumber

`func (o *Organization) GetVatNumber() string`

GetVatNumber returns the VatNumber field if non-nil, zero value otherwise.

### GetVatNumberOk

`func (o *Organization) GetVatNumberOk() (*string, bool)`

GetVatNumberOk returns a tuple with the VatNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatNumber

`func (o *Organization) SetVatNumber(v string)`

SetVatNumber sets VatNumber field to given value.

### HasVatNumber

`func (o *Organization) HasVatNumber() bool`

HasVatNumber returns a boolean if a field has been set.

### GetName

`func (o *Organization) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Organization) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Organization) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Organization) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPlan

`func (o *Organization) GetPlan() Plan`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *Organization) GetPlanOk() (*Plan, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *Organization) SetPlan(v Plan)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *Organization) HasPlan() bool`

HasPlan returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


