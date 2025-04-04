/*
Appgate SDP Controller REST API

# About   This specification documents the REST API calls for the Appgate SDP Controller.    Please refer to the REST API chapter in the manual or contact Appgate support with any questions about   this functionality. # Getting Started   Requirements for API scripting:   - Access to the Admin/API TLS Connection (default port 8443) of a Controller appliance.     (https://sdphelp.appgate.com/adminguide/appliance-function-configure.html?anchor=admin-api)   - An API user with relevant permissions.     (https://sdphelp.appgate.com/adminguide/administrative-roles-configure.html)   - In order to use the simple login API, Admin MFA must be disabled or the API user must be excluded.     (https://sdphelp.appgate.com/adminguide/mfa-for-admins.html) # Base path   HTTPS requests must be sent to the Admin Interface hostname and port, with **_/admin** path.    For example: **https://appgate.company.com:8443/admin**    All requests must have the **Accept** header as:    **application/vnd.appgate.peer-v22+json**    An exception is made for the **_/admin/version** endpoint which instead expects an **application/json** Accept header. # API Conventions   API conventions are  important to understand and follow strictly.    - While updating objects (via PUT), entire object must be sent with all fields.     - For example, in order to add a remedy method to the condition below:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": []       }       ```     - send the entire object with updated and non-updated fields:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": [{\"type\": \"DisplayMessage\", \"message\": \"test message\"}]       }       ```    - In case Controller returns an error (non-2xx HTTP status code), response body is JSON.     The \"message\" field contains information about the error.     HTTP 422 \"Unprocessable Entity\" has extra `errors` field to list all the issues with specific fields.    - Empty string (\"\") is considered a different value than \"null\" or field being omitted from JSON.     Omitting the field is recommended if no value is intended.     Empty string (\"\") will be almost always rejected as invalid value.    - There are common pattern between many objects:     - **Configuration Objects**: There are many objects with common fields, namely \"id\", \"name\", \"notes\", \"created\"       and \"updated\". These entities are listed, queried, created, updated and deleted in a similar fashion.     - **Distinguished Name**: Users and Devices are identified with what is called Distinguished Names, as used in        LDAP. The distinguished format that identifies a device and a user combination is        \"CN=\\<Device ID\\>,CN=\\<username\\>,OU=\\<Identity Provider Name\\>\". Some objects have the        \"userDistinguishedName\" field, which does not include the CN for Device ID.        This identifies a user on every device.

API version: API version 22.0
Contact: appgatesdp.support@appgate.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// SessionInfoDistinguishedNameDataValueEntitlementInfosValue Entitlement details.
type SessionInfoDistinguishedNameDataValueEntitlementInfosValue struct {
	// Whether the Entitlement is accessible or not.
	Access *bool `json:"access,omitempty"`
	// Whether all the Conditions must succeed to have access to this Entitlement or just one.
	ConditionLogic *string `json:"conditionLogic,omitempty"`
	// Current results of the Condition evaluations. Entitlement access is decided according to these results. The key is the Condition name.
	ConditionResults *map[string]bool `json:"conditionResults,omitempty"`
	// Current Firewall Rules assigned after evaluating all the Entitlements, Conditions and Name Resolvers.
	FirewallRules []SessionInfoDistinguishedNameDataValueEntitlementInfosValueFirewallRulesInner `json:"firewallRules,omitempty"`
	// Names of the Policies that has granted this Entitlement.
	PolicyNames []string `json:"policyNames,omitempty"`
	// The name of the primary Site if this entitlements is currently active as part of backup Site feature.
	PrimarySite *string `json:"primarySite,omitempty"`
	// Whether the Entitlement has only domain:// type actions. When that is the case, the firewallRules being empty is a normal scenario.
	DomainEntitlement *bool `json:"domainEntitlement,omitempty"`
}

// NewSessionInfoDistinguishedNameDataValueEntitlementInfosValue instantiates a new SessionInfoDistinguishedNameDataValueEntitlementInfosValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSessionInfoDistinguishedNameDataValueEntitlementInfosValue() *SessionInfoDistinguishedNameDataValueEntitlementInfosValue {
	this := SessionInfoDistinguishedNameDataValueEntitlementInfosValue{}
	return &this
}

// NewSessionInfoDistinguishedNameDataValueEntitlementInfosValueWithDefaults instantiates a new SessionInfoDistinguishedNameDataValueEntitlementInfosValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSessionInfoDistinguishedNameDataValueEntitlementInfosValueWithDefaults() *SessionInfoDistinguishedNameDataValueEntitlementInfosValue {
	this := SessionInfoDistinguishedNameDataValueEntitlementInfosValue{}
	return &this
}

// GetAccess returns the Access field value if set, zero value otherwise.
func (o *SessionInfoDistinguishedNameDataValueEntitlementInfosValue) GetAccess() bool {
	if o == nil || o.Access == nil {
		var ret bool
		return ret
	}
	return *o.Access
}

// GetAccessOk returns a tuple with the Access field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionInfoDistinguishedNameDataValueEntitlementInfosValue) GetAccessOk() (*bool, bool) {
	if o == nil || o.Access == nil {
		return nil, false
	}
	return o.Access, true
}

// HasAccess returns a boolean if a field has been set.
func (o *SessionInfoDistinguishedNameDataValueEntitlementInfosValue) HasAccess() bool {
	if o != nil && o.Access != nil {
		return true
	}

	return false
}

// SetAccess gets a reference to the given bool and assigns it to the Access field.
func (o *SessionInfoDistinguishedNameDataValueEntitlementInfosValue) SetAccess(v bool) {
	o.Access = &v
}

// GetConditionLogic returns the ConditionLogic field value if set, zero value otherwise.
func (o *SessionInfoDistinguishedNameDataValueEntitlementInfosValue) GetConditionLogic() string {
	if o == nil || o.ConditionLogic == nil {
		var ret string
		return ret
	}
	return *o.ConditionLogic
}

// GetConditionLogicOk returns a tuple with the ConditionLogic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionInfoDistinguishedNameDataValueEntitlementInfosValue) GetConditionLogicOk() (*string, bool) {
	if o == nil || o.ConditionLogic == nil {
		return nil, false
	}
	return o.ConditionLogic, true
}

// HasConditionLogic returns a boolean if a field has been set.
func (o *SessionInfoDistinguishedNameDataValueEntitlementInfosValue) HasConditionLogic() bool {
	if o != nil && o.ConditionLogic != nil {
		return true
	}

	return false
}

// SetConditionLogic gets a reference to the given string and assigns it to the ConditionLogic field.
func (o *SessionInfoDistinguishedNameDataValueEntitlementInfosValue) SetConditionLogic(v string) {
	o.ConditionLogic = &v
}

// GetConditionResults returns the ConditionResults field value if set, zero value otherwise.
func (o *SessionInfoDistinguishedNameDataValueEntitlementInfosValue) GetConditionResults() map[string]bool {
	if o == nil || o.ConditionResults == nil {
		var ret map[string]bool
		return ret
	}
	return *o.ConditionResults
}

// GetConditionResultsOk returns a tuple with the ConditionResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionInfoDistinguishedNameDataValueEntitlementInfosValue) GetConditionResultsOk() (*map[string]bool, bool) {
	if o == nil || o.ConditionResults == nil {
		return nil, false
	}
	return o.ConditionResults, true
}

// HasConditionResults returns a boolean if a field has been set.
func (o *SessionInfoDistinguishedNameDataValueEntitlementInfosValue) HasConditionResults() bool {
	if o != nil && o.ConditionResults != nil {
		return true
	}

	return false
}

// SetConditionResults gets a reference to the given map[string]bool and assigns it to the ConditionResults field.
func (o *SessionInfoDistinguishedNameDataValueEntitlementInfosValue) SetConditionResults(v map[string]bool) {
	o.ConditionResults = &v
}

// GetFirewallRules returns the FirewallRules field value if set, zero value otherwise.
func (o *SessionInfoDistinguishedNameDataValueEntitlementInfosValue) GetFirewallRules() []SessionInfoDistinguishedNameDataValueEntitlementInfosValueFirewallRulesInner {
	if o == nil || o.FirewallRules == nil {
		var ret []SessionInfoDistinguishedNameDataValueEntitlementInfosValueFirewallRulesInner
		return ret
	}
	return o.FirewallRules
}

// GetFirewallRulesOk returns a tuple with the FirewallRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionInfoDistinguishedNameDataValueEntitlementInfosValue) GetFirewallRulesOk() ([]SessionInfoDistinguishedNameDataValueEntitlementInfosValueFirewallRulesInner, bool) {
	if o == nil || o.FirewallRules == nil {
		return nil, false
	}
	return o.FirewallRules, true
}

// HasFirewallRules returns a boolean if a field has been set.
func (o *SessionInfoDistinguishedNameDataValueEntitlementInfosValue) HasFirewallRules() bool {
	if o != nil && o.FirewallRules != nil {
		return true
	}

	return false
}

// SetFirewallRules gets a reference to the given []SessionInfoDistinguishedNameDataValueEntitlementInfosValueFirewallRulesInner and assigns it to the FirewallRules field.
func (o *SessionInfoDistinguishedNameDataValueEntitlementInfosValue) SetFirewallRules(v []SessionInfoDistinguishedNameDataValueEntitlementInfosValueFirewallRulesInner) {
	o.FirewallRules = v
}

// GetPolicyNames returns the PolicyNames field value if set, zero value otherwise.
func (o *SessionInfoDistinguishedNameDataValueEntitlementInfosValue) GetPolicyNames() []string {
	if o == nil || o.PolicyNames == nil {
		var ret []string
		return ret
	}
	return o.PolicyNames
}

// GetPolicyNamesOk returns a tuple with the PolicyNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionInfoDistinguishedNameDataValueEntitlementInfosValue) GetPolicyNamesOk() ([]string, bool) {
	if o == nil || o.PolicyNames == nil {
		return nil, false
	}
	return o.PolicyNames, true
}

// HasPolicyNames returns a boolean if a field has been set.
func (o *SessionInfoDistinguishedNameDataValueEntitlementInfosValue) HasPolicyNames() bool {
	if o != nil && o.PolicyNames != nil {
		return true
	}

	return false
}

// SetPolicyNames gets a reference to the given []string and assigns it to the PolicyNames field.
func (o *SessionInfoDistinguishedNameDataValueEntitlementInfosValue) SetPolicyNames(v []string) {
	o.PolicyNames = v
}

// GetPrimarySite returns the PrimarySite field value if set, zero value otherwise.
func (o *SessionInfoDistinguishedNameDataValueEntitlementInfosValue) GetPrimarySite() string {
	if o == nil || o.PrimarySite == nil {
		var ret string
		return ret
	}
	return *o.PrimarySite
}

// GetPrimarySiteOk returns a tuple with the PrimarySite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionInfoDistinguishedNameDataValueEntitlementInfosValue) GetPrimarySiteOk() (*string, bool) {
	if o == nil || o.PrimarySite == nil {
		return nil, false
	}
	return o.PrimarySite, true
}

// HasPrimarySite returns a boolean if a field has been set.
func (o *SessionInfoDistinguishedNameDataValueEntitlementInfosValue) HasPrimarySite() bool {
	if o != nil && o.PrimarySite != nil {
		return true
	}

	return false
}

// SetPrimarySite gets a reference to the given string and assigns it to the PrimarySite field.
func (o *SessionInfoDistinguishedNameDataValueEntitlementInfosValue) SetPrimarySite(v string) {
	o.PrimarySite = &v
}

// GetDomainEntitlement returns the DomainEntitlement field value if set, zero value otherwise.
func (o *SessionInfoDistinguishedNameDataValueEntitlementInfosValue) GetDomainEntitlement() bool {
	if o == nil || o.DomainEntitlement == nil {
		var ret bool
		return ret
	}
	return *o.DomainEntitlement
}

// GetDomainEntitlementOk returns a tuple with the DomainEntitlement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionInfoDistinguishedNameDataValueEntitlementInfosValue) GetDomainEntitlementOk() (*bool, bool) {
	if o == nil || o.DomainEntitlement == nil {
		return nil, false
	}
	return o.DomainEntitlement, true
}

// HasDomainEntitlement returns a boolean if a field has been set.
func (o *SessionInfoDistinguishedNameDataValueEntitlementInfosValue) HasDomainEntitlement() bool {
	if o != nil && o.DomainEntitlement != nil {
		return true
	}

	return false
}

// SetDomainEntitlement gets a reference to the given bool and assigns it to the DomainEntitlement field.
func (o *SessionInfoDistinguishedNameDataValueEntitlementInfosValue) SetDomainEntitlement(v bool) {
	o.DomainEntitlement = &v
}

func (o SessionInfoDistinguishedNameDataValueEntitlementInfosValue) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Access != nil {
		toSerialize["access"] = o.Access
	}
	if o.ConditionLogic != nil {
		toSerialize["conditionLogic"] = o.ConditionLogic
	}
	if o.ConditionResults != nil {
		toSerialize["conditionResults"] = o.ConditionResults
	}
	if o.FirewallRules != nil {
		toSerialize["firewallRules"] = o.FirewallRules
	}
	if o.PolicyNames != nil {
		toSerialize["policyNames"] = o.PolicyNames
	}
	if o.PrimarySite != nil {
		toSerialize["primarySite"] = o.PrimarySite
	}
	if o.DomainEntitlement != nil {
		toSerialize["domainEntitlement"] = o.DomainEntitlement
	}
	return json.Marshal(toSerialize)
}

type NullableSessionInfoDistinguishedNameDataValueEntitlementInfosValue struct {
	value *SessionInfoDistinguishedNameDataValueEntitlementInfosValue
	isSet bool
}

func (v NullableSessionInfoDistinguishedNameDataValueEntitlementInfosValue) Get() *SessionInfoDistinguishedNameDataValueEntitlementInfosValue {
	return v.value
}

func (v *NullableSessionInfoDistinguishedNameDataValueEntitlementInfosValue) Set(val *SessionInfoDistinguishedNameDataValueEntitlementInfosValue) {
	v.value = val
	v.isSet = true
}

func (v NullableSessionInfoDistinguishedNameDataValueEntitlementInfosValue) IsSet() bool {
	return v.isSet
}

func (v *NullableSessionInfoDistinguishedNameDataValueEntitlementInfosValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSessionInfoDistinguishedNameDataValueEntitlementInfosValue(val *SessionInfoDistinguishedNameDataValueEntitlementInfosValue) *NullableSessionInfoDistinguishedNameDataValueEntitlementInfosValue {
	return &NullableSessionInfoDistinguishedNameDataValueEntitlementInfosValue{value: val, isSet: true}
}

func (v NullableSessionInfoDistinguishedNameDataValueEntitlementInfosValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSessionInfoDistinguishedNameDataValueEntitlementInfosValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
