/*
Appgate SDP Controller REST API

# About   This specification documents the REST API calls for the Appgate SDP Controller.    Please refer to the REST API chapter in the manual or contact Appgate support with any questions about   this functionality. # Getting Started   Requirements for API scripting:   - Access to the Admin/API TLS Connection (default port 8443) of a Controller appliance.     (https://sdphelp.appgate.com/adminguide/appliance-function-configure.html?anchor=admin-api)   - An API user with relevant permissions.     (https://sdphelp.appgate.com/adminguide/administrative-roles-configure.html)   - In order to use the simple login API, Admin MFA must be disabled or the API user must be excluded.     (https://sdphelp.appgate.com/adminguide/mfa-for-admins.html) # Base path   HTTPS requests must be sent to the Admin Interface hostname and port, with **_/admin** path.    For example: **https://appgate.company.com:8443/admin**    All requests must have the **Accept** header as:    **application/vnd.appgate.peer-v22+json**    An exception is made for the **_/admin/version** endpoint which instead expects an **application/json** Accept header. # API Conventions   API conventions are  important to understand and follow strictly.    - While updating objects (via PUT), entire object must be sent with all fields.     - For example, in order to add a remedy method to the condition below:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": []       }       ```     - send the entire object with updated and non-updated fields:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": [{\"type\": \"DisplayMessage\", \"message\": \"test message\"}]       }       ```    - In case Controller returns an error (non-2xx HTTP status code), response body is JSON.     The \"message\" field contains information about the error.     HTTP 422 \"Unprocessable Entity\" has extra `errors` field to list all the issues with specific fields.    - Empty string (\"\") is considered a different value than \"null\" or field being omitted from JSON.     Omitting the field is recommended if no value is intended.     Empty string (\"\") will be almost always rejected as invalid value.    - There are common pattern between many objects:     - **Configuration Objects**: There are many objects with common fields, namely \"id\", \"name\", \"notes\", \"created\"       and \"updated\". These entities are listed, queried, created, updated and deleted in a similar fashion.     - **Distinguished Name**: Users and Devices are identified with what is called Distinguished Names, as used in        LDAP. The distinguished format that identifies a device and a user combination is        \"CN=\\<Device ID\\>,CN=\\<username\\>,OU=\\<Identity Provider Name\\>\". Some objects have the        \"userDistinguishedName\" field, which does not include the CN for Device ID.        This identifies a user on every device.

API version: API version 22.2
Contact: appgatesdp.support@appgate.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// AppDetailsAllOf struct for AppDetailsAllOf
type AppDetailsAllOf struct {
	// The rule applied.
	Rule *string `json:"rule,omitempty"`
	// The protocol used.
	Protocol *string `json:"protocol,omitempty"`
	// The direction of the access.
	Direction *string `json:"direction,omitempty"`
	// Users accessed this app.
	Users []User `json:"users,omitempty"`
	// Common groups of the users which have accessed this app.
	CommonGroups []CommonGroup `json:"commonGroups,omitempty"`
	// Data for hit count chart per day.
	HitChartData []map[string]float32 `json:"hitChartData,omitempty"`
	// Timestamp of the last time the app was reset.
	LastResetAt *time.Time `json:"lastResetAt,omitempty"`
	// Timestamp of the last time the analysis is done for all apps.
	LastAnalysisAt *time.Time `json:"lastAnalysisAt,omitempty"`
	// The ID of the Policy generated when access is configured to this app.
	PolicyId *string `json:"policyId,omitempty"`
	// The ID of the Entitlement generated when access is configured to this app.
	EntitlementId *string `json:"entitlementId,omitempty"`
	// History of the actions taken on this app.
	History []interface{} `json:"history,omitempty"`
}

// NewAppDetailsAllOf instantiates a new AppDetailsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppDetailsAllOf() *AppDetailsAllOf {
	this := AppDetailsAllOf{}
	return &this
}

// NewAppDetailsAllOfWithDefaults instantiates a new AppDetailsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppDetailsAllOfWithDefaults() *AppDetailsAllOf {
	this := AppDetailsAllOf{}
	return &this
}

// GetRule returns the Rule field value if set, zero value otherwise.
func (o *AppDetailsAllOf) GetRule() string {
	if o == nil || o.Rule == nil {
		var ret string
		return ret
	}
	return *o.Rule
}

// GetRuleOk returns a tuple with the Rule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppDetailsAllOf) GetRuleOk() (*string, bool) {
	if o == nil || o.Rule == nil {
		return nil, false
	}
	return o.Rule, true
}

// HasRule returns a boolean if a field has been set.
func (o *AppDetailsAllOf) HasRule() bool {
	if o != nil && o.Rule != nil {
		return true
	}

	return false
}

// SetRule gets a reference to the given string and assigns it to the Rule field.
func (o *AppDetailsAllOf) SetRule(v string) {
	o.Rule = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *AppDetailsAllOf) GetProtocol() string {
	if o == nil || o.Protocol == nil {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppDetailsAllOf) GetProtocolOk() (*string, bool) {
	if o == nil || o.Protocol == nil {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *AppDetailsAllOf) HasProtocol() bool {
	if o != nil && o.Protocol != nil {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *AppDetailsAllOf) SetProtocol(v string) {
	o.Protocol = &v
}

// GetDirection returns the Direction field value if set, zero value otherwise.
func (o *AppDetailsAllOf) GetDirection() string {
	if o == nil || o.Direction == nil {
		var ret string
		return ret
	}
	return *o.Direction
}

// GetDirectionOk returns a tuple with the Direction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppDetailsAllOf) GetDirectionOk() (*string, bool) {
	if o == nil || o.Direction == nil {
		return nil, false
	}
	return o.Direction, true
}

// HasDirection returns a boolean if a field has been set.
func (o *AppDetailsAllOf) HasDirection() bool {
	if o != nil && o.Direction != nil {
		return true
	}

	return false
}

// SetDirection gets a reference to the given string and assigns it to the Direction field.
func (o *AppDetailsAllOf) SetDirection(v string) {
	o.Direction = &v
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *AppDetailsAllOf) GetUsers() []User {
	if o == nil || o.Users == nil {
		var ret []User
		return ret
	}
	return o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppDetailsAllOf) GetUsersOk() ([]User, bool) {
	if o == nil || o.Users == nil {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *AppDetailsAllOf) HasUsers() bool {
	if o != nil && o.Users != nil {
		return true
	}

	return false
}

// SetUsers gets a reference to the given []User and assigns it to the Users field.
func (o *AppDetailsAllOf) SetUsers(v []User) {
	o.Users = v
}

// GetCommonGroups returns the CommonGroups field value if set, zero value otherwise.
func (o *AppDetailsAllOf) GetCommonGroups() []CommonGroup {
	if o == nil || o.CommonGroups == nil {
		var ret []CommonGroup
		return ret
	}
	return o.CommonGroups
}

// GetCommonGroupsOk returns a tuple with the CommonGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppDetailsAllOf) GetCommonGroupsOk() ([]CommonGroup, bool) {
	if o == nil || o.CommonGroups == nil {
		return nil, false
	}
	return o.CommonGroups, true
}

// HasCommonGroups returns a boolean if a field has been set.
func (o *AppDetailsAllOf) HasCommonGroups() bool {
	if o != nil && o.CommonGroups != nil {
		return true
	}

	return false
}

// SetCommonGroups gets a reference to the given []CommonGroup and assigns it to the CommonGroups field.
func (o *AppDetailsAllOf) SetCommonGroups(v []CommonGroup) {
	o.CommonGroups = v
}

// GetHitChartData returns the HitChartData field value if set, zero value otherwise.
func (o *AppDetailsAllOf) GetHitChartData() []map[string]float32 {
	if o == nil || o.HitChartData == nil {
		var ret []map[string]float32
		return ret
	}
	return o.HitChartData
}

// GetHitChartDataOk returns a tuple with the HitChartData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppDetailsAllOf) GetHitChartDataOk() ([]map[string]float32, bool) {
	if o == nil || o.HitChartData == nil {
		return nil, false
	}
	return o.HitChartData, true
}

// HasHitChartData returns a boolean if a field has been set.
func (o *AppDetailsAllOf) HasHitChartData() bool {
	if o != nil && o.HitChartData != nil {
		return true
	}

	return false
}

// SetHitChartData gets a reference to the given []map[string]float32 and assigns it to the HitChartData field.
func (o *AppDetailsAllOf) SetHitChartData(v []map[string]float32) {
	o.HitChartData = v
}

// GetLastResetAt returns the LastResetAt field value if set, zero value otherwise.
func (o *AppDetailsAllOf) GetLastResetAt() time.Time {
	if o == nil || o.LastResetAt == nil {
		var ret time.Time
		return ret
	}
	return *o.LastResetAt
}

// GetLastResetAtOk returns a tuple with the LastResetAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppDetailsAllOf) GetLastResetAtOk() (*time.Time, bool) {
	if o == nil || o.LastResetAt == nil {
		return nil, false
	}
	return o.LastResetAt, true
}

// HasLastResetAt returns a boolean if a field has been set.
func (o *AppDetailsAllOf) HasLastResetAt() bool {
	if o != nil && o.LastResetAt != nil {
		return true
	}

	return false
}

// SetLastResetAt gets a reference to the given time.Time and assigns it to the LastResetAt field.
func (o *AppDetailsAllOf) SetLastResetAt(v time.Time) {
	o.LastResetAt = &v
}

// GetLastAnalysisAt returns the LastAnalysisAt field value if set, zero value otherwise.
func (o *AppDetailsAllOf) GetLastAnalysisAt() time.Time {
	if o == nil || o.LastAnalysisAt == nil {
		var ret time.Time
		return ret
	}
	return *o.LastAnalysisAt
}

// GetLastAnalysisAtOk returns a tuple with the LastAnalysisAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppDetailsAllOf) GetLastAnalysisAtOk() (*time.Time, bool) {
	if o == nil || o.LastAnalysisAt == nil {
		return nil, false
	}
	return o.LastAnalysisAt, true
}

// HasLastAnalysisAt returns a boolean if a field has been set.
func (o *AppDetailsAllOf) HasLastAnalysisAt() bool {
	if o != nil && o.LastAnalysisAt != nil {
		return true
	}

	return false
}

// SetLastAnalysisAt gets a reference to the given time.Time and assigns it to the LastAnalysisAt field.
func (o *AppDetailsAllOf) SetLastAnalysisAt(v time.Time) {
	o.LastAnalysisAt = &v
}

// GetPolicyId returns the PolicyId field value if set, zero value otherwise.
func (o *AppDetailsAllOf) GetPolicyId() string {
	if o == nil || o.PolicyId == nil {
		var ret string
		return ret
	}
	return *o.PolicyId
}

// GetPolicyIdOk returns a tuple with the PolicyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppDetailsAllOf) GetPolicyIdOk() (*string, bool) {
	if o == nil || o.PolicyId == nil {
		return nil, false
	}
	return o.PolicyId, true
}

// HasPolicyId returns a boolean if a field has been set.
func (o *AppDetailsAllOf) HasPolicyId() bool {
	if o != nil && o.PolicyId != nil {
		return true
	}

	return false
}

// SetPolicyId gets a reference to the given string and assigns it to the PolicyId field.
func (o *AppDetailsAllOf) SetPolicyId(v string) {
	o.PolicyId = &v
}

// GetEntitlementId returns the EntitlementId field value if set, zero value otherwise.
func (o *AppDetailsAllOf) GetEntitlementId() string {
	if o == nil || o.EntitlementId == nil {
		var ret string
		return ret
	}
	return *o.EntitlementId
}

// GetEntitlementIdOk returns a tuple with the EntitlementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppDetailsAllOf) GetEntitlementIdOk() (*string, bool) {
	if o == nil || o.EntitlementId == nil {
		return nil, false
	}
	return o.EntitlementId, true
}

// HasEntitlementId returns a boolean if a field has been set.
func (o *AppDetailsAllOf) HasEntitlementId() bool {
	if o != nil && o.EntitlementId != nil {
		return true
	}

	return false
}

// SetEntitlementId gets a reference to the given string and assigns it to the EntitlementId field.
func (o *AppDetailsAllOf) SetEntitlementId(v string) {
	o.EntitlementId = &v
}

// GetHistory returns the History field value if set, zero value otherwise.
func (o *AppDetailsAllOf) GetHistory() []interface{} {
	if o == nil || o.History == nil {
		var ret []interface{}
		return ret
	}
	return o.History
}

// GetHistoryOk returns a tuple with the History field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppDetailsAllOf) GetHistoryOk() ([]interface{}, bool) {
	if o == nil || o.History == nil {
		return nil, false
	}
	return o.History, true
}

// HasHistory returns a boolean if a field has been set.
func (o *AppDetailsAllOf) HasHistory() bool {
	if o != nil && o.History != nil {
		return true
	}

	return false
}

// SetHistory gets a reference to the given []interface{} and assigns it to the History field.
func (o *AppDetailsAllOf) SetHistory(v []interface{}) {
	o.History = v
}

func (o AppDetailsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Rule != nil {
		toSerialize["rule"] = o.Rule
	}
	if o.Protocol != nil {
		toSerialize["protocol"] = o.Protocol
	}
	if o.Direction != nil {
		toSerialize["direction"] = o.Direction
	}
	if o.Users != nil {
		toSerialize["users"] = o.Users
	}
	if o.CommonGroups != nil {
		toSerialize["commonGroups"] = o.CommonGroups
	}
	if o.HitChartData != nil {
		toSerialize["hitChartData"] = o.HitChartData
	}
	if o.LastResetAt != nil {
		toSerialize["lastResetAt"] = o.LastResetAt
	}
	if o.LastAnalysisAt != nil {
		toSerialize["lastAnalysisAt"] = o.LastAnalysisAt
	}
	if o.PolicyId != nil {
		toSerialize["policyId"] = o.PolicyId
	}
	if o.EntitlementId != nil {
		toSerialize["entitlementId"] = o.EntitlementId
	}
	if o.History != nil {
		toSerialize["history"] = o.History
	}
	return json.Marshal(toSerialize)
}

type NullableAppDetailsAllOf struct {
	value *AppDetailsAllOf
	isSet bool
}

func (v NullableAppDetailsAllOf) Get() *AppDetailsAllOf {
	return v.value
}

func (v *NullableAppDetailsAllOf) Set(val *AppDetailsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAppDetailsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAppDetailsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppDetailsAllOf(val *AppDetailsAllOf) *NullableAppDetailsAllOf {
	return &NullableAppDetailsAllOf{value: val, isSet: true}
}

func (v NullableAppDetailsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppDetailsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
