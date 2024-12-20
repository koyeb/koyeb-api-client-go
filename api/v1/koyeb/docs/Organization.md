# Organization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Address1** | Pointer to **string** |  | [optional] 
**Address2** | Pointer to **string** |  | [optional] 
**City** | Pointer to **string** |  | [optional] 
**PostalCode** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**Country** | Pointer to **string** |  | [optional] 
**Company** | Pointer to **bool** |  | [optional] 
**VatNumber** | Pointer to **string** |  | [optional] 
**BillingName** | Pointer to **string** |  | [optional] 
**BillingEmail** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Plan** | Pointer to [**Plan**](Plan.md) |  | [optional] [default to PLAN_HOBBY]
**PlanUpdatedAt** | Pointer to **time.Time** |  | [optional] 
**HasPaymentMethod** | Pointer to **bool** |  | [optional] 
**SubscriptionId** | Pointer to **string** |  | [optional] 
**CurrentSubscriptionId** | Pointer to **string** |  | [optional] 
**LatestSubscriptionId** | Pointer to **string** |  | [optional] 
**SignupQualification** | Pointer to **map[string]interface{}** |  | [optional] 
**Status** | Pointer to [**OrganizationStatus**](OrganizationStatus.md) |  | [optional] [default to ORGANIZATIONSTATUS_WARNING]
**StatusMessage** | Pointer to [**OrganizationDetailedStatus**](OrganizationDetailedStatus.md) |  | [optional] [default to ORGANIZATIONDETAILEDSTATUS_NEW]
**DeactivationReason** | Pointer to [**OrganizationDeactivationReason**](OrganizationDeactivationReason.md) |  | [optional] [default to ORGANIZATIONDEACTIVATIONREASON_INVALID]
**Verified** | Pointer to **bool** |  | [optional] 
**QualifiesForHobby23** | Pointer to **bool** |  | [optional] 
**ReprocessAfter** | Pointer to **time.Time** |  | [optional] 
**Trialing** | Pointer to **bool** |  | [optional] 
**TrialStartsAt** | Pointer to **time.Time** |  | [optional] 
**TrialEndsAt** | Pointer to **time.Time** |  | [optional] 

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

### GetId

`func (o *Organization) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Organization) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Organization) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Organization) HasId() bool`

HasId returns a boolean if a field has been set.

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

### GetBillingName

`func (o *Organization) GetBillingName() string`

GetBillingName returns the BillingName field if non-nil, zero value otherwise.

### GetBillingNameOk

`func (o *Organization) GetBillingNameOk() (*string, bool)`

GetBillingNameOk returns a tuple with the BillingName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingName

`func (o *Organization) SetBillingName(v string)`

SetBillingName sets BillingName field to given value.

### HasBillingName

`func (o *Organization) HasBillingName() bool`

HasBillingName returns a boolean if a field has been set.

### GetBillingEmail

`func (o *Organization) GetBillingEmail() string`

GetBillingEmail returns the BillingEmail field if non-nil, zero value otherwise.

### GetBillingEmailOk

`func (o *Organization) GetBillingEmailOk() (*string, bool)`

GetBillingEmailOk returns a tuple with the BillingEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingEmail

`func (o *Organization) SetBillingEmail(v string)`

SetBillingEmail sets BillingEmail field to given value.

### HasBillingEmail

`func (o *Organization) HasBillingEmail() bool`

HasBillingEmail returns a boolean if a field has been set.

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

### GetPlanUpdatedAt

`func (o *Organization) GetPlanUpdatedAt() time.Time`

GetPlanUpdatedAt returns the PlanUpdatedAt field if non-nil, zero value otherwise.

### GetPlanUpdatedAtOk

`func (o *Organization) GetPlanUpdatedAtOk() (*time.Time, bool)`

GetPlanUpdatedAtOk returns a tuple with the PlanUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanUpdatedAt

`func (o *Organization) SetPlanUpdatedAt(v time.Time)`

SetPlanUpdatedAt sets PlanUpdatedAt field to given value.

### HasPlanUpdatedAt

`func (o *Organization) HasPlanUpdatedAt() bool`

HasPlanUpdatedAt returns a boolean if a field has been set.

### GetHasPaymentMethod

`func (o *Organization) GetHasPaymentMethod() bool`

GetHasPaymentMethod returns the HasPaymentMethod field if non-nil, zero value otherwise.

### GetHasPaymentMethodOk

`func (o *Organization) GetHasPaymentMethodOk() (*bool, bool)`

GetHasPaymentMethodOk returns a tuple with the HasPaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPaymentMethod

`func (o *Organization) SetHasPaymentMethod(v bool)`

SetHasPaymentMethod sets HasPaymentMethod field to given value.

### HasHasPaymentMethod

`func (o *Organization) HasHasPaymentMethod() bool`

HasHasPaymentMethod returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *Organization) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *Organization) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *Organization) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *Organization) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetCurrentSubscriptionId

`func (o *Organization) GetCurrentSubscriptionId() string`

GetCurrentSubscriptionId returns the CurrentSubscriptionId field if non-nil, zero value otherwise.

### GetCurrentSubscriptionIdOk

`func (o *Organization) GetCurrentSubscriptionIdOk() (*string, bool)`

GetCurrentSubscriptionIdOk returns a tuple with the CurrentSubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentSubscriptionId

`func (o *Organization) SetCurrentSubscriptionId(v string)`

SetCurrentSubscriptionId sets CurrentSubscriptionId field to given value.

### HasCurrentSubscriptionId

`func (o *Organization) HasCurrentSubscriptionId() bool`

HasCurrentSubscriptionId returns a boolean if a field has been set.

### GetLatestSubscriptionId

`func (o *Organization) GetLatestSubscriptionId() string`

GetLatestSubscriptionId returns the LatestSubscriptionId field if non-nil, zero value otherwise.

### GetLatestSubscriptionIdOk

`func (o *Organization) GetLatestSubscriptionIdOk() (*string, bool)`

GetLatestSubscriptionIdOk returns a tuple with the LatestSubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestSubscriptionId

`func (o *Organization) SetLatestSubscriptionId(v string)`

SetLatestSubscriptionId sets LatestSubscriptionId field to given value.

### HasLatestSubscriptionId

`func (o *Organization) HasLatestSubscriptionId() bool`

HasLatestSubscriptionId returns a boolean if a field has been set.

### GetSignupQualification

`func (o *Organization) GetSignupQualification() map[string]interface{}`

GetSignupQualification returns the SignupQualification field if non-nil, zero value otherwise.

### GetSignupQualificationOk

`func (o *Organization) GetSignupQualificationOk() (*map[string]interface{}, bool)`

GetSignupQualificationOk returns a tuple with the SignupQualification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignupQualification

`func (o *Organization) SetSignupQualification(v map[string]interface{})`

SetSignupQualification sets SignupQualification field to given value.

### HasSignupQualification

`func (o *Organization) HasSignupQualification() bool`

HasSignupQualification returns a boolean if a field has been set.

### GetStatus

`func (o *Organization) GetStatus() OrganizationStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Organization) GetStatusOk() (*OrganizationStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Organization) SetStatus(v OrganizationStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Organization) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusMessage

`func (o *Organization) GetStatusMessage() OrganizationDetailedStatus`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *Organization) GetStatusMessageOk() (*OrganizationDetailedStatus, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *Organization) SetStatusMessage(v OrganizationDetailedStatus)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *Organization) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### GetDeactivationReason

`func (o *Organization) GetDeactivationReason() OrganizationDeactivationReason`

GetDeactivationReason returns the DeactivationReason field if non-nil, zero value otherwise.

### GetDeactivationReasonOk

`func (o *Organization) GetDeactivationReasonOk() (*OrganizationDeactivationReason, bool)`

GetDeactivationReasonOk returns a tuple with the DeactivationReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeactivationReason

`func (o *Organization) SetDeactivationReason(v OrganizationDeactivationReason)`

SetDeactivationReason sets DeactivationReason field to given value.

### HasDeactivationReason

`func (o *Organization) HasDeactivationReason() bool`

HasDeactivationReason returns a boolean if a field has been set.

### GetVerified

`func (o *Organization) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *Organization) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *Organization) SetVerified(v bool)`

SetVerified sets Verified field to given value.

### HasVerified

`func (o *Organization) HasVerified() bool`

HasVerified returns a boolean if a field has been set.

### GetQualifiesForHobby23

`func (o *Organization) GetQualifiesForHobby23() bool`

GetQualifiesForHobby23 returns the QualifiesForHobby23 field if non-nil, zero value otherwise.

### GetQualifiesForHobby23Ok

`func (o *Organization) GetQualifiesForHobby23Ok() (*bool, bool)`

GetQualifiesForHobby23Ok returns a tuple with the QualifiesForHobby23 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualifiesForHobby23

`func (o *Organization) SetQualifiesForHobby23(v bool)`

SetQualifiesForHobby23 sets QualifiesForHobby23 field to given value.

### HasQualifiesForHobby23

`func (o *Organization) HasQualifiesForHobby23() bool`

HasQualifiesForHobby23 returns a boolean if a field has been set.

### GetReprocessAfter

`func (o *Organization) GetReprocessAfter() time.Time`

GetReprocessAfter returns the ReprocessAfter field if non-nil, zero value otherwise.

### GetReprocessAfterOk

`func (o *Organization) GetReprocessAfterOk() (*time.Time, bool)`

GetReprocessAfterOk returns a tuple with the ReprocessAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReprocessAfter

`func (o *Organization) SetReprocessAfter(v time.Time)`

SetReprocessAfter sets ReprocessAfter field to given value.

### HasReprocessAfter

`func (o *Organization) HasReprocessAfter() bool`

HasReprocessAfter returns a boolean if a field has been set.

### GetTrialing

`func (o *Organization) GetTrialing() bool`

GetTrialing returns the Trialing field if non-nil, zero value otherwise.

### GetTrialingOk

`func (o *Organization) GetTrialingOk() (*bool, bool)`

GetTrialingOk returns a tuple with the Trialing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialing

`func (o *Organization) SetTrialing(v bool)`

SetTrialing sets Trialing field to given value.

### HasTrialing

`func (o *Organization) HasTrialing() bool`

HasTrialing returns a boolean if a field has been set.

### GetTrialStartsAt

`func (o *Organization) GetTrialStartsAt() time.Time`

GetTrialStartsAt returns the TrialStartsAt field if non-nil, zero value otherwise.

### GetTrialStartsAtOk

`func (o *Organization) GetTrialStartsAtOk() (*time.Time, bool)`

GetTrialStartsAtOk returns a tuple with the TrialStartsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialStartsAt

`func (o *Organization) SetTrialStartsAt(v time.Time)`

SetTrialStartsAt sets TrialStartsAt field to given value.

### HasTrialStartsAt

`func (o *Organization) HasTrialStartsAt() bool`

HasTrialStartsAt returns a boolean if a field has been set.

### GetTrialEndsAt

`func (o *Organization) GetTrialEndsAt() time.Time`

GetTrialEndsAt returns the TrialEndsAt field if non-nil, zero value otherwise.

### GetTrialEndsAtOk

`func (o *Organization) GetTrialEndsAtOk() (*time.Time, bool)`

GetTrialEndsAtOk returns a tuple with the TrialEndsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialEndsAt

`func (o *Organization) SetTrialEndsAt(v time.Time)`

SetTrialEndsAt sets TrialEndsAt field to given value.

### HasTrialEndsAt

`func (o *Organization) HasTrialEndsAt() bool`

HasTrialEndsAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


