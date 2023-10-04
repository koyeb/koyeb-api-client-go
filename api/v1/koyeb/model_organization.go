/*
Koyeb Rest API

The Koyeb API allows you to interact with the Koyeb platform in a simple, programmatic way using conventional HTTP requests. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package koyeb

import (
	"encoding/json"
	"time"
)

// Organization struct for Organization
type Organization struct {
	Id *string `json:"id,omitempty"`
	Address1 *string `json:"address1,omitempty"`
	Address2 *string `json:"address2,omitempty"`
	City *string `json:"city,omitempty"`
	PostalCode *string `json:"postal_code,omitempty"`
	State *string `json:"state,omitempty"`
	Country *string `json:"country,omitempty"`
	Company *bool `json:"company,omitempty"`
	VatNumber *string `json:"vat_number,omitempty"`
	BillingName *string `json:"billing_name,omitempty"`
	BillingEmail *string `json:"billing_email,omitempty"`
	Name *string `json:"name,omitempty"`
	Plan *Plan `json:"plan,omitempty"`
	PlanUpdatedAt *time.Time `json:"plan_updated_at,omitempty"`
	HasPaymentMethod *bool `json:"has_payment_method,omitempty"`
	SubscriptionId *string `json:"subscription_id,omitempty"`
	CurrentSubscriptionId *string `json:"current_subscription_id,omitempty"`
	LatestSubscriptionId *string `json:"latest_subscription_id,omitempty"`
	SignupQualification map[string]interface{} `json:"signup_qualification,omitempty"`
	Status *OrganizationStatus `json:"status,omitempty"`
	StatusMessage *OrganizationDetailedStatus `json:"status_message,omitempty"`
	DeactivationReason *OrganizationDeactivationReason `json:"deactivation_reason,omitempty"`
	Verified *bool `json:"verified,omitempty"`
	QualifiesForHobby23 *bool `json:"qualifies_for_hobby23,omitempty"`
}

// NewOrganization instantiates a new Organization object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganization() *Organization {
	this := Organization{}
	var plan Plan = PLAN_HOBBY
	this.Plan = &plan
	var status OrganizationStatus = ORGANIZATIONSTATUS_WARNING
	this.Status = &status
	var statusMessage OrganizationDetailedStatus = ORGANIZATIONDETAILEDSTATUS_NEW
	this.StatusMessage = &statusMessage
	var deactivationReason OrganizationDeactivationReason = ORGANIZATIONDEACTIVATIONREASON_INVALID
	this.DeactivationReason = &deactivationReason
	return &this
}

// NewOrganizationWithDefaults instantiates a new Organization object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationWithDefaults() *Organization {
	this := Organization{}
	var plan Plan = PLAN_HOBBY
	this.Plan = &plan
	var status OrganizationStatus = ORGANIZATIONSTATUS_WARNING
	this.Status = &status
	var statusMessage OrganizationDetailedStatus = ORGANIZATIONDETAILEDSTATUS_NEW
	this.StatusMessage = &statusMessage
	var deactivationReason OrganizationDeactivationReason = ORGANIZATIONDEACTIVATIONREASON_INVALID
	this.DeactivationReason = &deactivationReason
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Organization) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organization) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Organization) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Organization) SetId(v string) {
	o.Id = &v
}

// GetAddress1 returns the Address1 field value if set, zero value otherwise.
func (o *Organization) GetAddress1() string {
	if o == nil || isNil(o.Address1) {
		var ret string
		return ret
	}
	return *o.Address1
}

// GetAddress1Ok returns a tuple with the Address1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organization) GetAddress1Ok() (*string, bool) {
	if o == nil || isNil(o.Address1) {
    return nil, false
	}
	return o.Address1, true
}

// HasAddress1 returns a boolean if a field has been set.
func (o *Organization) HasAddress1() bool {
	if o != nil && !isNil(o.Address1) {
		return true
	}

	return false
}

// SetAddress1 gets a reference to the given string and assigns it to the Address1 field.
func (o *Organization) SetAddress1(v string) {
	o.Address1 = &v
}

// GetAddress2 returns the Address2 field value if set, zero value otherwise.
func (o *Organization) GetAddress2() string {
	if o == nil || isNil(o.Address2) {
		var ret string
		return ret
	}
	return *o.Address2
}

// GetAddress2Ok returns a tuple with the Address2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organization) GetAddress2Ok() (*string, bool) {
	if o == nil || isNil(o.Address2) {
    return nil, false
	}
	return o.Address2, true
}

// HasAddress2 returns a boolean if a field has been set.
func (o *Organization) HasAddress2() bool {
	if o != nil && !isNil(o.Address2) {
		return true
	}

	return false
}

// SetAddress2 gets a reference to the given string and assigns it to the Address2 field.
func (o *Organization) SetAddress2(v string) {
	o.Address2 = &v
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *Organization) GetCity() string {
	if o == nil || isNil(o.City) {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organization) GetCityOk() (*string, bool) {
	if o == nil || isNil(o.City) {
    return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *Organization) HasCity() bool {
	if o != nil && !isNil(o.City) {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *Organization) SetCity(v string) {
	o.City = &v
}

// GetPostalCode returns the PostalCode field value if set, zero value otherwise.
func (o *Organization) GetPostalCode() string {
	if o == nil || isNil(o.PostalCode) {
		var ret string
		return ret
	}
	return *o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organization) GetPostalCodeOk() (*string, bool) {
	if o == nil || isNil(o.PostalCode) {
    return nil, false
	}
	return o.PostalCode, true
}

// HasPostalCode returns a boolean if a field has been set.
func (o *Organization) HasPostalCode() bool {
	if o != nil && !isNil(o.PostalCode) {
		return true
	}

	return false
}

// SetPostalCode gets a reference to the given string and assigns it to the PostalCode field.
func (o *Organization) SetPostalCode(v string) {
	o.PostalCode = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *Organization) GetState() string {
	if o == nil || isNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organization) GetStateOk() (*string, bool) {
	if o == nil || isNil(o.State) {
    return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *Organization) HasState() bool {
	if o != nil && !isNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *Organization) SetState(v string) {
	o.State = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *Organization) GetCountry() string {
	if o == nil || isNil(o.Country) {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organization) GetCountryOk() (*string, bool) {
	if o == nil || isNil(o.Country) {
    return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *Organization) HasCountry() bool {
	if o != nil && !isNil(o.Country) {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *Organization) SetCountry(v string) {
	o.Country = &v
}

// GetCompany returns the Company field value if set, zero value otherwise.
func (o *Organization) GetCompany() bool {
	if o == nil || isNil(o.Company) {
		var ret bool
		return ret
	}
	return *o.Company
}

// GetCompanyOk returns a tuple with the Company field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organization) GetCompanyOk() (*bool, bool) {
	if o == nil || isNil(o.Company) {
    return nil, false
	}
	return o.Company, true
}

// HasCompany returns a boolean if a field has been set.
func (o *Organization) HasCompany() bool {
	if o != nil && !isNil(o.Company) {
		return true
	}

	return false
}

// SetCompany gets a reference to the given bool and assigns it to the Company field.
func (o *Organization) SetCompany(v bool) {
	o.Company = &v
}

// GetVatNumber returns the VatNumber field value if set, zero value otherwise.
func (o *Organization) GetVatNumber() string {
	if o == nil || isNil(o.VatNumber) {
		var ret string
		return ret
	}
	return *o.VatNumber
}

// GetVatNumberOk returns a tuple with the VatNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organization) GetVatNumberOk() (*string, bool) {
	if o == nil || isNil(o.VatNumber) {
    return nil, false
	}
	return o.VatNumber, true
}

// HasVatNumber returns a boolean if a field has been set.
func (o *Organization) HasVatNumber() bool {
	if o != nil && !isNil(o.VatNumber) {
		return true
	}

	return false
}

// SetVatNumber gets a reference to the given string and assigns it to the VatNumber field.
func (o *Organization) SetVatNumber(v string) {
	o.VatNumber = &v
}

// GetBillingName returns the BillingName field value if set, zero value otherwise.
func (o *Organization) GetBillingName() string {
	if o == nil || isNil(o.BillingName) {
		var ret string
		return ret
	}
	return *o.BillingName
}

// GetBillingNameOk returns a tuple with the BillingName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organization) GetBillingNameOk() (*string, bool) {
	if o == nil || isNil(o.BillingName) {
    return nil, false
	}
	return o.BillingName, true
}

// HasBillingName returns a boolean if a field has been set.
func (o *Organization) HasBillingName() bool {
	if o != nil && !isNil(o.BillingName) {
		return true
	}

	return false
}

// SetBillingName gets a reference to the given string and assigns it to the BillingName field.
func (o *Organization) SetBillingName(v string) {
	o.BillingName = &v
}

// GetBillingEmail returns the BillingEmail field value if set, zero value otherwise.
func (o *Organization) GetBillingEmail() string {
	if o == nil || isNil(o.BillingEmail) {
		var ret string
		return ret
	}
	return *o.BillingEmail
}

// GetBillingEmailOk returns a tuple with the BillingEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organization) GetBillingEmailOk() (*string, bool) {
	if o == nil || isNil(o.BillingEmail) {
    return nil, false
	}
	return o.BillingEmail, true
}

// HasBillingEmail returns a boolean if a field has been set.
func (o *Organization) HasBillingEmail() bool {
	if o != nil && !isNil(o.BillingEmail) {
		return true
	}

	return false
}

// SetBillingEmail gets a reference to the given string and assigns it to the BillingEmail field.
func (o *Organization) SetBillingEmail(v string) {
	o.BillingEmail = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Organization) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organization) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Organization) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Organization) SetName(v string) {
	o.Name = &v
}

// GetPlan returns the Plan field value if set, zero value otherwise.
func (o *Organization) GetPlan() Plan {
	if o == nil || isNil(o.Plan) {
		var ret Plan
		return ret
	}
	return *o.Plan
}

// GetPlanOk returns a tuple with the Plan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organization) GetPlanOk() (*Plan, bool) {
	if o == nil || isNil(o.Plan) {
    return nil, false
	}
	return o.Plan, true
}

// HasPlan returns a boolean if a field has been set.
func (o *Organization) HasPlan() bool {
	if o != nil && !isNil(o.Plan) {
		return true
	}

	return false
}

// SetPlan gets a reference to the given Plan and assigns it to the Plan field.
func (o *Organization) SetPlan(v Plan) {
	o.Plan = &v
}

// GetPlanUpdatedAt returns the PlanUpdatedAt field value if set, zero value otherwise.
func (o *Organization) GetPlanUpdatedAt() time.Time {
	if o == nil || isNil(o.PlanUpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.PlanUpdatedAt
}

// GetPlanUpdatedAtOk returns a tuple with the PlanUpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organization) GetPlanUpdatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.PlanUpdatedAt) {
    return nil, false
	}
	return o.PlanUpdatedAt, true
}

// HasPlanUpdatedAt returns a boolean if a field has been set.
func (o *Organization) HasPlanUpdatedAt() bool {
	if o != nil && !isNil(o.PlanUpdatedAt) {
		return true
	}

	return false
}

// SetPlanUpdatedAt gets a reference to the given time.Time and assigns it to the PlanUpdatedAt field.
func (o *Organization) SetPlanUpdatedAt(v time.Time) {
	o.PlanUpdatedAt = &v
}

// GetHasPaymentMethod returns the HasPaymentMethod field value if set, zero value otherwise.
func (o *Organization) GetHasPaymentMethod() bool {
	if o == nil || isNil(o.HasPaymentMethod) {
		var ret bool
		return ret
	}
	return *o.HasPaymentMethod
}

// GetHasPaymentMethodOk returns a tuple with the HasPaymentMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organization) GetHasPaymentMethodOk() (*bool, bool) {
	if o == nil || isNil(o.HasPaymentMethod) {
    return nil, false
	}
	return o.HasPaymentMethod, true
}

// HasHasPaymentMethod returns a boolean if a field has been set.
func (o *Organization) HasHasPaymentMethod() bool {
	if o != nil && !isNil(o.HasPaymentMethod) {
		return true
	}

	return false
}

// SetHasPaymentMethod gets a reference to the given bool and assigns it to the HasPaymentMethod field.
func (o *Organization) SetHasPaymentMethod(v bool) {
	o.HasPaymentMethod = &v
}

// GetSubscriptionId returns the SubscriptionId field value if set, zero value otherwise.
func (o *Organization) GetSubscriptionId() string {
	if o == nil || isNil(o.SubscriptionId) {
		var ret string
		return ret
	}
	return *o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organization) GetSubscriptionIdOk() (*string, bool) {
	if o == nil || isNil(o.SubscriptionId) {
    return nil, false
	}
	return o.SubscriptionId, true
}

// HasSubscriptionId returns a boolean if a field has been set.
func (o *Organization) HasSubscriptionId() bool {
	if o != nil && !isNil(o.SubscriptionId) {
		return true
	}

	return false
}

// SetSubscriptionId gets a reference to the given string and assigns it to the SubscriptionId field.
func (o *Organization) SetSubscriptionId(v string) {
	o.SubscriptionId = &v
}

// GetCurrentSubscriptionId returns the CurrentSubscriptionId field value if set, zero value otherwise.
func (o *Organization) GetCurrentSubscriptionId() string {
	if o == nil || isNil(o.CurrentSubscriptionId) {
		var ret string
		return ret
	}
	return *o.CurrentSubscriptionId
}

// GetCurrentSubscriptionIdOk returns a tuple with the CurrentSubscriptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organization) GetCurrentSubscriptionIdOk() (*string, bool) {
	if o == nil || isNil(o.CurrentSubscriptionId) {
    return nil, false
	}
	return o.CurrentSubscriptionId, true
}

// HasCurrentSubscriptionId returns a boolean if a field has been set.
func (o *Organization) HasCurrentSubscriptionId() bool {
	if o != nil && !isNil(o.CurrentSubscriptionId) {
		return true
	}

	return false
}

// SetCurrentSubscriptionId gets a reference to the given string and assigns it to the CurrentSubscriptionId field.
func (o *Organization) SetCurrentSubscriptionId(v string) {
	o.CurrentSubscriptionId = &v
}

// GetLatestSubscriptionId returns the LatestSubscriptionId field value if set, zero value otherwise.
func (o *Organization) GetLatestSubscriptionId() string {
	if o == nil || isNil(o.LatestSubscriptionId) {
		var ret string
		return ret
	}
	return *o.LatestSubscriptionId
}

// GetLatestSubscriptionIdOk returns a tuple with the LatestSubscriptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organization) GetLatestSubscriptionIdOk() (*string, bool) {
	if o == nil || isNil(o.LatestSubscriptionId) {
    return nil, false
	}
	return o.LatestSubscriptionId, true
}

// HasLatestSubscriptionId returns a boolean if a field has been set.
func (o *Organization) HasLatestSubscriptionId() bool {
	if o != nil && !isNil(o.LatestSubscriptionId) {
		return true
	}

	return false
}

// SetLatestSubscriptionId gets a reference to the given string and assigns it to the LatestSubscriptionId field.
func (o *Organization) SetLatestSubscriptionId(v string) {
	o.LatestSubscriptionId = &v
}

// GetSignupQualification returns the SignupQualification field value if set, zero value otherwise.
func (o *Organization) GetSignupQualification() map[string]interface{} {
	if o == nil || isNil(o.SignupQualification) {
		var ret map[string]interface{}
		return ret
	}
	return o.SignupQualification
}

// GetSignupQualificationOk returns a tuple with the SignupQualification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organization) GetSignupQualificationOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.SignupQualification) {
    return map[string]interface{}{}, false
	}
	return o.SignupQualification, true
}

// HasSignupQualification returns a boolean if a field has been set.
func (o *Organization) HasSignupQualification() bool {
	if o != nil && !isNil(o.SignupQualification) {
		return true
	}

	return false
}

// SetSignupQualification gets a reference to the given map[string]interface{} and assigns it to the SignupQualification field.
func (o *Organization) SetSignupQualification(v map[string]interface{}) {
	o.SignupQualification = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Organization) GetStatus() OrganizationStatus {
	if o == nil || isNil(o.Status) {
		var ret OrganizationStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organization) GetStatusOk() (*OrganizationStatus, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Organization) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given OrganizationStatus and assigns it to the Status field.
func (o *Organization) SetStatus(v OrganizationStatus) {
	o.Status = &v
}

// GetStatusMessage returns the StatusMessage field value if set, zero value otherwise.
func (o *Organization) GetStatusMessage() OrganizationDetailedStatus {
	if o == nil || isNil(o.StatusMessage) {
		var ret OrganizationDetailedStatus
		return ret
	}
	return *o.StatusMessage
}

// GetStatusMessageOk returns a tuple with the StatusMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organization) GetStatusMessageOk() (*OrganizationDetailedStatus, bool) {
	if o == nil || isNil(o.StatusMessage) {
    return nil, false
	}
	return o.StatusMessage, true
}

// HasStatusMessage returns a boolean if a field has been set.
func (o *Organization) HasStatusMessage() bool {
	if o != nil && !isNil(o.StatusMessage) {
		return true
	}

	return false
}

// SetStatusMessage gets a reference to the given OrganizationDetailedStatus and assigns it to the StatusMessage field.
func (o *Organization) SetStatusMessage(v OrganizationDetailedStatus) {
	o.StatusMessage = &v
}

// GetDeactivationReason returns the DeactivationReason field value if set, zero value otherwise.
func (o *Organization) GetDeactivationReason() OrganizationDeactivationReason {
	if o == nil || isNil(o.DeactivationReason) {
		var ret OrganizationDeactivationReason
		return ret
	}
	return *o.DeactivationReason
}

// GetDeactivationReasonOk returns a tuple with the DeactivationReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organization) GetDeactivationReasonOk() (*OrganizationDeactivationReason, bool) {
	if o == nil || isNil(o.DeactivationReason) {
    return nil, false
	}
	return o.DeactivationReason, true
}

// HasDeactivationReason returns a boolean if a field has been set.
func (o *Organization) HasDeactivationReason() bool {
	if o != nil && !isNil(o.DeactivationReason) {
		return true
	}

	return false
}

// SetDeactivationReason gets a reference to the given OrganizationDeactivationReason and assigns it to the DeactivationReason field.
func (o *Organization) SetDeactivationReason(v OrganizationDeactivationReason) {
	o.DeactivationReason = &v
}

// GetVerified returns the Verified field value if set, zero value otherwise.
func (o *Organization) GetVerified() bool {
	if o == nil || isNil(o.Verified) {
		var ret bool
		return ret
	}
	return *o.Verified
}

// GetVerifiedOk returns a tuple with the Verified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organization) GetVerifiedOk() (*bool, bool) {
	if o == nil || isNil(o.Verified) {
    return nil, false
	}
	return o.Verified, true
}

// HasVerified returns a boolean if a field has been set.
func (o *Organization) HasVerified() bool {
	if o != nil && !isNil(o.Verified) {
		return true
	}

	return false
}

// SetVerified gets a reference to the given bool and assigns it to the Verified field.
func (o *Organization) SetVerified(v bool) {
	o.Verified = &v
}

// GetQualifiesForHobby23 returns the QualifiesForHobby23 field value if set, zero value otherwise.
func (o *Organization) GetQualifiesForHobby23() bool {
	if o == nil || isNil(o.QualifiesForHobby23) {
		var ret bool
		return ret
	}
	return *o.QualifiesForHobby23
}

// GetQualifiesForHobby23Ok returns a tuple with the QualifiesForHobby23 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organization) GetQualifiesForHobby23Ok() (*bool, bool) {
	if o == nil || isNil(o.QualifiesForHobby23) {
    return nil, false
	}
	return o.QualifiesForHobby23, true
}

// HasQualifiesForHobby23 returns a boolean if a field has been set.
func (o *Organization) HasQualifiesForHobby23() bool {
	if o != nil && !isNil(o.QualifiesForHobby23) {
		return true
	}

	return false
}

// SetQualifiesForHobby23 gets a reference to the given bool and assigns it to the QualifiesForHobby23 field.
func (o *Organization) SetQualifiesForHobby23(v bool) {
	o.QualifiesForHobby23 = &v
}

func (o Organization) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Address1) {
		toSerialize["address1"] = o.Address1
	}
	if !isNil(o.Address2) {
		toSerialize["address2"] = o.Address2
	}
	if !isNil(o.City) {
		toSerialize["city"] = o.City
	}
	if !isNil(o.PostalCode) {
		toSerialize["postal_code"] = o.PostalCode
	}
	if !isNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !isNil(o.Country) {
		toSerialize["country"] = o.Country
	}
	if !isNil(o.Company) {
		toSerialize["company"] = o.Company
	}
	if !isNil(o.VatNumber) {
		toSerialize["vat_number"] = o.VatNumber
	}
	if !isNil(o.BillingName) {
		toSerialize["billing_name"] = o.BillingName
	}
	if !isNil(o.BillingEmail) {
		toSerialize["billing_email"] = o.BillingEmail
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Plan) {
		toSerialize["plan"] = o.Plan
	}
	if !isNil(o.PlanUpdatedAt) {
		toSerialize["plan_updated_at"] = o.PlanUpdatedAt
	}
	if !isNil(o.HasPaymentMethod) {
		toSerialize["has_payment_method"] = o.HasPaymentMethod
	}
	if !isNil(o.SubscriptionId) {
		toSerialize["subscription_id"] = o.SubscriptionId
	}
	if !isNil(o.CurrentSubscriptionId) {
		toSerialize["current_subscription_id"] = o.CurrentSubscriptionId
	}
	if !isNil(o.LatestSubscriptionId) {
		toSerialize["latest_subscription_id"] = o.LatestSubscriptionId
	}
	if !isNil(o.SignupQualification) {
		toSerialize["signup_qualification"] = o.SignupQualification
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !isNil(o.StatusMessage) {
		toSerialize["status_message"] = o.StatusMessage
	}
	if !isNil(o.DeactivationReason) {
		toSerialize["deactivation_reason"] = o.DeactivationReason
	}
	if !isNil(o.Verified) {
		toSerialize["verified"] = o.Verified
	}
	if !isNil(o.QualifiesForHobby23) {
		toSerialize["qualifies_for_hobby23"] = o.QualifiesForHobby23
	}
	return json.Marshal(toSerialize)
}

type NullableOrganization struct {
	value *Organization
	isSet bool
}

func (v NullableOrganization) Get() *Organization {
	return v.value
}

func (v *NullableOrganization) Set(val *Organization) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganization) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganization) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganization(val *Organization) *NullableOrganization {
	return &NullableOrganization{value: val, isSet: true}
}

func (v NullableOrganization) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganization) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


