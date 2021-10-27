/*
Appgate SDP Controller REST API

# About   This specification documents the REST API calls for the Appgate SDP Controller.    Please refer to the REST API chapter in the manual or contact Appgate support with any questions about   this functionality. # Getting Started   Requirements for API scripting:   - Access to the Admin/API TLS Connection (default port 8443) of a Controller appliance.     (https://sdphelp.appgate.com/adminguide/appliance-function-configure.html?anchor=admin-api)   - An API user with relevant permissions.     (https://sdphelp.appgate.com/adminguide/administrative-roles-configure.html)   - In order to use the simple login API, Admin MFA must be disabled or the API user must be excluded.     (https://sdphelp.appgate.com/adminguide/mfa-for-admins.html) # Base path   HTTPS requests must be sent to the Admin Interface hostname and port, with **_/admin** path.    For example: **https://appgate.company.com:8443/admin**    All requests must have the **Accept** header as:    **application/vnd.appgate.peer-v16+json** # API Conventions   API conventions are  important to understand and follow strictly.    - While updating objects (via PUT), entire object must be sent with all fields.     - For example, in order to add a remedy method to the condition below:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": []       }       ```     - send the entire object with updated and non-updated fields:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": [{\"type\": \"DisplayMessage\", \"message\": \"test message\"}]       }       ```    - In case Controller returns an error (non-2xx HTTP status code), response body is JSON.     The \"message\" field contains information about the error.     HTTP 422 \"Unprocessable Entity\" has extra `errors` field to list all the issues with specific fields.    - Empty string (\"\") is considered a different value than \"null\" or field being omitted from JSON.     Omitting the field is recommend if no value is intended.     Empty string (\"\") will be almost always rejected as invalid value.    - There are common pattern between many objects:     - **Configuration Objects**: There are many objects with common fields, namely \"id\", \"name\", \"notes\", \"created\"       and \"updated\". These entities are listed, queried, created, updated and deleted in a similar fashion.     - **Distinguished Name**: Users and Devices are identified with what is called Distinguished Names, as used in        LDAP. The distinguished format that identifies a device and a user combination is        \"CN=\\<Device ID\\>,CN=\\<username\\>,OU=\\<Identity Provider Name\\>\". Some objects have the        \"userDistinguishedName\" field, which does not include the CN for Device ID.        This identifies a user on every device.

API version: API version 16.0
Contact: appgatesdp.support@appgate.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// SessionInfoDistinguishedNameEntitlementInfos Entitlement details.
type SessionInfoDistinguishedNameEntitlementInfos struct {
	// Current results of the Condition evaluations. Entitlement access is decided according to these results. The key is the Condition name.
	ConditionResults *map[string]bool `json:"conditionResults,omitempty"`
	// Current Firewall Rules assigned after evaluating all the Entitlements, Conditions and Name Resolvers.
	FirewallRules *[]SessionInfoDistinguishedNameFirewallRules `json:"firewallRules,omitempty"`
	// Names of the Policies that has granted this Entitlement.
	PolicyNames *[]string `json:"policyNames,omitempty"`
}

// NewSessionInfoDistinguishedNameEntitlementInfos instantiates a new SessionInfoDistinguishedNameEntitlementInfos object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSessionInfoDistinguishedNameEntitlementInfos() *SessionInfoDistinguishedNameEntitlementInfos {
	this := SessionInfoDistinguishedNameEntitlementInfos{}
	return &this
}

// NewSessionInfoDistinguishedNameEntitlementInfosWithDefaults instantiates a new SessionInfoDistinguishedNameEntitlementInfos object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSessionInfoDistinguishedNameEntitlementInfosWithDefaults() *SessionInfoDistinguishedNameEntitlementInfos {
	this := SessionInfoDistinguishedNameEntitlementInfos{}
	return &this
}

// GetConditionResults returns the ConditionResults field value if set, zero value otherwise.
func (o *SessionInfoDistinguishedNameEntitlementInfos) GetConditionResults() map[string]bool {
	if o == nil || o.ConditionResults == nil {
		var ret map[string]bool
		return ret
	}
	return *o.ConditionResults
}

// GetConditionResultsOk returns a tuple with the ConditionResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionInfoDistinguishedNameEntitlementInfos) GetConditionResultsOk() (*map[string]bool, bool) {
	if o == nil || o.ConditionResults == nil {
		return nil, false
	}
	return o.ConditionResults, true
}

// HasConditionResults returns a boolean if a field has been set.
func (o *SessionInfoDistinguishedNameEntitlementInfos) HasConditionResults() bool {
	if o != nil && o.ConditionResults != nil {
		return true
	}

	return false
}

// SetConditionResults gets a reference to the given map[string]bool and assigns it to the ConditionResults field.
func (o *SessionInfoDistinguishedNameEntitlementInfos) SetConditionResults(v map[string]bool) {
	o.ConditionResults = &v
}

// GetFirewallRules returns the FirewallRules field value if set, zero value otherwise.
func (o *SessionInfoDistinguishedNameEntitlementInfos) GetFirewallRules() []SessionInfoDistinguishedNameFirewallRules {
	if o == nil || o.FirewallRules == nil {
		var ret []SessionInfoDistinguishedNameFirewallRules
		return ret
	}
	return *o.FirewallRules
}

// GetFirewallRulesOk returns a tuple with the FirewallRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionInfoDistinguishedNameEntitlementInfos) GetFirewallRulesOk() (*[]SessionInfoDistinguishedNameFirewallRules, bool) {
	if o == nil || o.FirewallRules == nil {
		return nil, false
	}
	return o.FirewallRules, true
}

// HasFirewallRules returns a boolean if a field has been set.
func (o *SessionInfoDistinguishedNameEntitlementInfos) HasFirewallRules() bool {
	if o != nil && o.FirewallRules != nil {
		return true
	}

	return false
}

// SetFirewallRules gets a reference to the given []SessionInfoDistinguishedNameFirewallRules and assigns it to the FirewallRules field.
func (o *SessionInfoDistinguishedNameEntitlementInfos) SetFirewallRules(v []SessionInfoDistinguishedNameFirewallRules) {
	o.FirewallRules = &v
}

// GetPolicyNames returns the PolicyNames field value if set, zero value otherwise.
func (o *SessionInfoDistinguishedNameEntitlementInfos) GetPolicyNames() []string {
	if o == nil || o.PolicyNames == nil {
		var ret []string
		return ret
	}
	return *o.PolicyNames
}

// GetPolicyNamesOk returns a tuple with the PolicyNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionInfoDistinguishedNameEntitlementInfos) GetPolicyNamesOk() (*[]string, bool) {
	if o == nil || o.PolicyNames == nil {
		return nil, false
	}
	return o.PolicyNames, true
}

// HasPolicyNames returns a boolean if a field has been set.
func (o *SessionInfoDistinguishedNameEntitlementInfos) HasPolicyNames() bool {
	if o != nil && o.PolicyNames != nil {
		return true
	}

	return false
}

// SetPolicyNames gets a reference to the given []string and assigns it to the PolicyNames field.
func (o *SessionInfoDistinguishedNameEntitlementInfos) SetPolicyNames(v []string) {
	o.PolicyNames = &v
}

func (o SessionInfoDistinguishedNameEntitlementInfos) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ConditionResults != nil {
		toSerialize["conditionResults"] = o.ConditionResults
	}
	if o.FirewallRules != nil {
		toSerialize["firewallRules"] = o.FirewallRules
	}
	if o.PolicyNames != nil {
		toSerialize["policyNames"] = o.PolicyNames
	}
	return json.Marshal(toSerialize)
}

type NullableSessionInfoDistinguishedNameEntitlementInfos struct {
	value *SessionInfoDistinguishedNameEntitlementInfos
	isSet bool
}

func (v NullableSessionInfoDistinguishedNameEntitlementInfos) Get() *SessionInfoDistinguishedNameEntitlementInfos {
	return v.value
}

func (v *NullableSessionInfoDistinguishedNameEntitlementInfos) Set(val *SessionInfoDistinguishedNameEntitlementInfos) {
	v.value = val
	v.isSet = true
}

func (v NullableSessionInfoDistinguishedNameEntitlementInfos) IsSet() bool {
	return v.isSet
}

func (v *NullableSessionInfoDistinguishedNameEntitlementInfos) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSessionInfoDistinguishedNameEntitlementInfos(val *SessionInfoDistinguishedNameEntitlementInfos) *NullableSessionInfoDistinguishedNameEntitlementInfos {
	return &NullableSessionInfoDistinguishedNameEntitlementInfos{value: val, isSet: true}
}

func (v NullableSessionInfoDistinguishedNameEntitlementInfos) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSessionInfoDistinguishedNameEntitlementInfos) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
