/*
Appgate SDP Controller REST API

# About   This specification documents the REST API calls for the Appgate SDP Controller.    Please refer to the REST API chapter in the manual or contact Appgate support with any questions about   this functionality. # Getting Started   Requirements for API scripting:   - Access to the Admin/API TLS Connection (default port 8443) of a Controller appliance.     (https://sdphelp.appgate.com/adminguide/appliance-function-configure.html?anchor=admin-api)   - An API user with relevant permissions.     (https://sdphelp.appgate.com/adminguide/administrative-roles-configure.html)   - In order to use the simple login API, Admin MFA must be disabled or the API user must be excluded.     (https://sdphelp.appgate.com/adminguide/mfa-for-admins.html) # Base path   HTTPS requests must be sent to the Admin Interface hostname and port, with **_/admin** path.    For example: **https://appgate.company.com:8443/admin**    All requests must have the **Accept** header as:    **application/vnd.appgate.peer-v21+json**    An exception is made for the **_/admin/version** endpoint which instead expects an **application/json** Accept header. # API Conventions   API conventions are  important to understand and follow strictly.    - While updating objects (via PUT), entire object must be sent with all fields.     - For example, in order to add a remedy method to the condition below:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": []       }       ```     - send the entire object with updated and non-updated fields:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": [{\"type\": \"DisplayMessage\", \"message\": \"test message\"}]       }       ```    - In case Controller returns an error (non-2xx HTTP status code), response body is JSON.     The \"message\" field contains information about the error.     HTTP 422 \"Unprocessable Entity\" has extra `errors` field to list all the issues with specific fields.    - Empty string (\"\") is considered a different value than \"null\" or field being omitted from JSON.     Omitting the field is recommended if no value is intended.     Empty string (\"\") will be almost always rejected as invalid value.    - There are common pattern between many objects:     - **Configuration Objects**: There are many objects with common fields, namely \"id\", \"name\", \"notes\", \"created\"       and \"updated\". These entities are listed, queried, created, updated and deleted in a similar fashion.     - **Distinguished Name**: Users and Devices are identified with what is called Distinguished Names, as used in        LDAP. The distinguished format that identifies a device and a user combination is        \"CN=\\<Device ID\\>,CN=\\<username\\>,OU=\\<Identity Provider Name\\>\". Some objects have the        \"userDistinguishedName\" field, which does not include the CN for Device ID.        This identifies a user on every device.

API version: API version 21.1
Contact: appgatesdp.support@appgate.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// AppDetails struct for AppDetails
type AppDetails struct {
	// ID of the app.
	Id *string `json:"id,omitempty"`
	// Host of the app.
	Host *string `json:"host,omitempty"`
	// Port of the app.
	Port *int32 `json:"port,omitempty"`
	// Total user count for the app.
	UsersCount *int32 `json:"usersCount,omitempty"`
	// Timestamp of the last time the app was accessed by a user.
	LastAccessedAt *time.Time `json:"lastAccessedAt,omitempty"`
	// Status of the discovered app.
	Status *string `json:"status,omitempty"`
	// The alias of the app if access is granted.
	Alias *string `json:"alias,omitempty"`
	// Timestamp of the last time an action is taken on this discovered app.
	LastModified *time.Time `json:"lastModified,omitempty"`
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
	// The ID of the Policy generated when access is granted to this app.
	PolicyId *string `json:"policyId,omitempty"`
	// The ID of the Entitlement generated when access is granted to this app.
	EntitlementId *string `json:"entitlementId,omitempty"`
	// History of the actions taken on this app.
	History []interface{} `json:"history,omitempty"`
}

// NewAppDetails instantiates a new AppDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppDetails() *AppDetails {
	this := AppDetails{}
	return &this
}

// NewAppDetailsWithDefaults instantiates a new AppDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppDetailsWithDefaults() *AppDetails {
	this := AppDetails{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AppDetails) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppDetails) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AppDetails) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AppDetails) SetId(v string) {
	o.Id = &v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *AppDetails) GetHost() string {
	if o == nil || o.Host == nil {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppDetails) GetHostOk() (*string, bool) {
	if o == nil || o.Host == nil {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *AppDetails) HasHost() bool {
	if o != nil && o.Host != nil {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *AppDetails) SetHost(v string) {
	o.Host = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *AppDetails) GetPort() int32 {
	if o == nil || o.Port == nil {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppDetails) GetPortOk() (*int32, bool) {
	if o == nil || o.Port == nil {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *AppDetails) HasPort() bool {
	if o != nil && o.Port != nil {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *AppDetails) SetPort(v int32) {
	o.Port = &v
}

// GetUsersCount returns the UsersCount field value if set, zero value otherwise.
func (o *AppDetails) GetUsersCount() int32 {
	if o == nil || o.UsersCount == nil {
		var ret int32
		return ret
	}
	return *o.UsersCount
}

// GetUsersCountOk returns a tuple with the UsersCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppDetails) GetUsersCountOk() (*int32, bool) {
	if o == nil || o.UsersCount == nil {
		return nil, false
	}
	return o.UsersCount, true
}

// HasUsersCount returns a boolean if a field has been set.
func (o *AppDetails) HasUsersCount() bool {
	if o != nil && o.UsersCount != nil {
		return true
	}

	return false
}

// SetUsersCount gets a reference to the given int32 and assigns it to the UsersCount field.
func (o *AppDetails) SetUsersCount(v int32) {
	o.UsersCount = &v
}

// GetLastAccessedAt returns the LastAccessedAt field value if set, zero value otherwise.
func (o *AppDetails) GetLastAccessedAt() time.Time {
	if o == nil || o.LastAccessedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.LastAccessedAt
}

// GetLastAccessedAtOk returns a tuple with the LastAccessedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppDetails) GetLastAccessedAtOk() (*time.Time, bool) {
	if o == nil || o.LastAccessedAt == nil {
		return nil, false
	}
	return o.LastAccessedAt, true
}

// HasLastAccessedAt returns a boolean if a field has been set.
func (o *AppDetails) HasLastAccessedAt() bool {
	if o != nil && o.LastAccessedAt != nil {
		return true
	}

	return false
}

// SetLastAccessedAt gets a reference to the given time.Time and assigns it to the LastAccessedAt field.
func (o *AppDetails) SetLastAccessedAt(v time.Time) {
	o.LastAccessedAt = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AppDetails) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppDetails) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AppDetails) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *AppDetails) SetStatus(v string) {
	o.Status = &v
}

// GetAlias returns the Alias field value if set, zero value otherwise.
func (o *AppDetails) GetAlias() string {
	if o == nil || o.Alias == nil {
		var ret string
		return ret
	}
	return *o.Alias
}

// GetAliasOk returns a tuple with the Alias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppDetails) GetAliasOk() (*string, bool) {
	if o == nil || o.Alias == nil {
		return nil, false
	}
	return o.Alias, true
}

// HasAlias returns a boolean if a field has been set.
func (o *AppDetails) HasAlias() bool {
	if o != nil && o.Alias != nil {
		return true
	}

	return false
}

// SetAlias gets a reference to the given string and assigns it to the Alias field.
func (o *AppDetails) SetAlias(v string) {
	o.Alias = &v
}

// GetLastModified returns the LastModified field value if set, zero value otherwise.
func (o *AppDetails) GetLastModified() time.Time {
	if o == nil || o.LastModified == nil {
		var ret time.Time
		return ret
	}
	return *o.LastModified
}

// GetLastModifiedOk returns a tuple with the LastModified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppDetails) GetLastModifiedOk() (*time.Time, bool) {
	if o == nil || o.LastModified == nil {
		return nil, false
	}
	return o.LastModified, true
}

// HasLastModified returns a boolean if a field has been set.
func (o *AppDetails) HasLastModified() bool {
	if o != nil && o.LastModified != nil {
		return true
	}

	return false
}

// SetLastModified gets a reference to the given time.Time and assigns it to the LastModified field.
func (o *AppDetails) SetLastModified(v time.Time) {
	o.LastModified = &v
}

// GetRule returns the Rule field value if set, zero value otherwise.
func (o *AppDetails) GetRule() string {
	if o == nil || o.Rule == nil {
		var ret string
		return ret
	}
	return *o.Rule
}

// GetRuleOk returns a tuple with the Rule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppDetails) GetRuleOk() (*string, bool) {
	if o == nil || o.Rule == nil {
		return nil, false
	}
	return o.Rule, true
}

// HasRule returns a boolean if a field has been set.
func (o *AppDetails) HasRule() bool {
	if o != nil && o.Rule != nil {
		return true
	}

	return false
}

// SetRule gets a reference to the given string and assigns it to the Rule field.
func (o *AppDetails) SetRule(v string) {
	o.Rule = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *AppDetails) GetProtocol() string {
	if o == nil || o.Protocol == nil {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppDetails) GetProtocolOk() (*string, bool) {
	if o == nil || o.Protocol == nil {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *AppDetails) HasProtocol() bool {
	if o != nil && o.Protocol != nil {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *AppDetails) SetProtocol(v string) {
	o.Protocol = &v
}

// GetDirection returns the Direction field value if set, zero value otherwise.
func (o *AppDetails) GetDirection() string {
	if o == nil || o.Direction == nil {
		var ret string
		return ret
	}
	return *o.Direction
}

// GetDirectionOk returns a tuple with the Direction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppDetails) GetDirectionOk() (*string, bool) {
	if o == nil || o.Direction == nil {
		return nil, false
	}
	return o.Direction, true
}

// HasDirection returns a boolean if a field has been set.
func (o *AppDetails) HasDirection() bool {
	if o != nil && o.Direction != nil {
		return true
	}

	return false
}

// SetDirection gets a reference to the given string and assigns it to the Direction field.
func (o *AppDetails) SetDirection(v string) {
	o.Direction = &v
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *AppDetails) GetUsers() []User {
	if o == nil || o.Users == nil {
		var ret []User
		return ret
	}
	return o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppDetails) GetUsersOk() ([]User, bool) {
	if o == nil || o.Users == nil {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *AppDetails) HasUsers() bool {
	if o != nil && o.Users != nil {
		return true
	}

	return false
}

// SetUsers gets a reference to the given []User and assigns it to the Users field.
func (o *AppDetails) SetUsers(v []User) {
	o.Users = v
}

// GetCommonGroups returns the CommonGroups field value if set, zero value otherwise.
func (o *AppDetails) GetCommonGroups() []CommonGroup {
	if o == nil || o.CommonGroups == nil {
		var ret []CommonGroup
		return ret
	}
	return o.CommonGroups
}

// GetCommonGroupsOk returns a tuple with the CommonGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppDetails) GetCommonGroupsOk() ([]CommonGroup, bool) {
	if o == nil || o.CommonGroups == nil {
		return nil, false
	}
	return o.CommonGroups, true
}

// HasCommonGroups returns a boolean if a field has been set.
func (o *AppDetails) HasCommonGroups() bool {
	if o != nil && o.CommonGroups != nil {
		return true
	}

	return false
}

// SetCommonGroups gets a reference to the given []CommonGroup and assigns it to the CommonGroups field.
func (o *AppDetails) SetCommonGroups(v []CommonGroup) {
	o.CommonGroups = v
}

// GetHitChartData returns the HitChartData field value if set, zero value otherwise.
func (o *AppDetails) GetHitChartData() []map[string]float32 {
	if o == nil || o.HitChartData == nil {
		var ret []map[string]float32
		return ret
	}
	return o.HitChartData
}

// GetHitChartDataOk returns a tuple with the HitChartData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppDetails) GetHitChartDataOk() ([]map[string]float32, bool) {
	if o == nil || o.HitChartData == nil {
		return nil, false
	}
	return o.HitChartData, true
}

// HasHitChartData returns a boolean if a field has been set.
func (o *AppDetails) HasHitChartData() bool {
	if o != nil && o.HitChartData != nil {
		return true
	}

	return false
}

// SetHitChartData gets a reference to the given []map[string]float32 and assigns it to the HitChartData field.
func (o *AppDetails) SetHitChartData(v []map[string]float32) {
	o.HitChartData = v
}

// GetLastResetAt returns the LastResetAt field value if set, zero value otherwise.
func (o *AppDetails) GetLastResetAt() time.Time {
	if o == nil || o.LastResetAt == nil {
		var ret time.Time
		return ret
	}
	return *o.LastResetAt
}

// GetLastResetAtOk returns a tuple with the LastResetAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppDetails) GetLastResetAtOk() (*time.Time, bool) {
	if o == nil || o.LastResetAt == nil {
		return nil, false
	}
	return o.LastResetAt, true
}

// HasLastResetAt returns a boolean if a field has been set.
func (o *AppDetails) HasLastResetAt() bool {
	if o != nil && o.LastResetAt != nil {
		return true
	}

	return false
}

// SetLastResetAt gets a reference to the given time.Time and assigns it to the LastResetAt field.
func (o *AppDetails) SetLastResetAt(v time.Time) {
	o.LastResetAt = &v
}

// GetLastAnalysisAt returns the LastAnalysisAt field value if set, zero value otherwise.
func (o *AppDetails) GetLastAnalysisAt() time.Time {
	if o == nil || o.LastAnalysisAt == nil {
		var ret time.Time
		return ret
	}
	return *o.LastAnalysisAt
}

// GetLastAnalysisAtOk returns a tuple with the LastAnalysisAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppDetails) GetLastAnalysisAtOk() (*time.Time, bool) {
	if o == nil || o.LastAnalysisAt == nil {
		return nil, false
	}
	return o.LastAnalysisAt, true
}

// HasLastAnalysisAt returns a boolean if a field has been set.
func (o *AppDetails) HasLastAnalysisAt() bool {
	if o != nil && o.LastAnalysisAt != nil {
		return true
	}

	return false
}

// SetLastAnalysisAt gets a reference to the given time.Time and assigns it to the LastAnalysisAt field.
func (o *AppDetails) SetLastAnalysisAt(v time.Time) {
	o.LastAnalysisAt = &v
}

// GetPolicyId returns the PolicyId field value if set, zero value otherwise.
func (o *AppDetails) GetPolicyId() string {
	if o == nil || o.PolicyId == nil {
		var ret string
		return ret
	}
	return *o.PolicyId
}

// GetPolicyIdOk returns a tuple with the PolicyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppDetails) GetPolicyIdOk() (*string, bool) {
	if o == nil || o.PolicyId == nil {
		return nil, false
	}
	return o.PolicyId, true
}

// HasPolicyId returns a boolean if a field has been set.
func (o *AppDetails) HasPolicyId() bool {
	if o != nil && o.PolicyId != nil {
		return true
	}

	return false
}

// SetPolicyId gets a reference to the given string and assigns it to the PolicyId field.
func (o *AppDetails) SetPolicyId(v string) {
	o.PolicyId = &v
}

// GetEntitlementId returns the EntitlementId field value if set, zero value otherwise.
func (o *AppDetails) GetEntitlementId() string {
	if o == nil || o.EntitlementId == nil {
		var ret string
		return ret
	}
	return *o.EntitlementId
}

// GetEntitlementIdOk returns a tuple with the EntitlementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppDetails) GetEntitlementIdOk() (*string, bool) {
	if o == nil || o.EntitlementId == nil {
		return nil, false
	}
	return o.EntitlementId, true
}

// HasEntitlementId returns a boolean if a field has been set.
func (o *AppDetails) HasEntitlementId() bool {
	if o != nil && o.EntitlementId != nil {
		return true
	}

	return false
}

// SetEntitlementId gets a reference to the given string and assigns it to the EntitlementId field.
func (o *AppDetails) SetEntitlementId(v string) {
	o.EntitlementId = &v
}

// GetHistory returns the History field value if set, zero value otherwise.
func (o *AppDetails) GetHistory() []interface{} {
	if o == nil || o.History == nil {
		var ret []interface{}
		return ret
	}
	return o.History
}

// GetHistoryOk returns a tuple with the History field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppDetails) GetHistoryOk() ([]interface{}, bool) {
	if o == nil || o.History == nil {
		return nil, false
	}
	return o.History, true
}

// HasHistory returns a boolean if a field has been set.
func (o *AppDetails) HasHistory() bool {
	if o != nil && o.History != nil {
		return true
	}

	return false
}

// SetHistory gets a reference to the given []interface{} and assigns it to the History field.
func (o *AppDetails) SetHistory(v []interface{}) {
	o.History = v
}

func (o AppDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Host != nil {
		toSerialize["host"] = o.Host
	}
	if o.Port != nil {
		toSerialize["port"] = o.Port
	}
	if o.UsersCount != nil {
		toSerialize["usersCount"] = o.UsersCount
	}
	if o.LastAccessedAt != nil {
		toSerialize["lastAccessedAt"] = o.LastAccessedAt
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Alias != nil {
		toSerialize["alias"] = o.Alias
	}
	if o.LastModified != nil {
		toSerialize["lastModified"] = o.LastModified
	}
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

type NullableAppDetails struct {
	value *AppDetails
	isSet bool
}

func (v NullableAppDetails) Get() *AppDetails {
	return v.value
}

func (v *NullableAppDetails) Set(val *AppDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableAppDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableAppDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppDetails(val *AppDetails) *NullableAppDetails {
	return &NullableAppDetails{value: val, isSet: true}
}

func (v NullableAppDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
