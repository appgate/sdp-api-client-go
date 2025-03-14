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

// UnhealthyAppActionsInner struct for UnhealthyAppActionsInner
type UnhealthyAppActionsInner struct {
	// IP of the resource.
	Ip *string `json:"ip,omitempty"`
	// Action index of the Entitlement.
	Index *int32 `json:"index,omitempty"`
	// Report count of the unhealthy action.
	ReportCount *int32 `json:"reportCount,omitempty"`
	// User count of the unhealthy action.
	UserCount *int32 `json:"userCount,omitempty"`
	// Port range affected on the unhealthy action.
	PortRange *string `json:"portRange,omitempty"`
	// The reason for the unhealth.
	Reason *string `json:"reason,omitempty"`
	// Timestamp of the last time the action was reported unhealthy.
	LastReported *time.Time `json:"lastReported,omitempty"`
}

// NewUnhealthyAppActionsInner instantiates a new UnhealthyAppActionsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnhealthyAppActionsInner() *UnhealthyAppActionsInner {
	this := UnhealthyAppActionsInner{}
	return &this
}

// NewUnhealthyAppActionsInnerWithDefaults instantiates a new UnhealthyAppActionsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnhealthyAppActionsInnerWithDefaults() *UnhealthyAppActionsInner {
	this := UnhealthyAppActionsInner{}
	return &this
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *UnhealthyAppActionsInner) GetIp() string {
	if o == nil || o.Ip == nil {
		var ret string
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnhealthyAppActionsInner) GetIpOk() (*string, bool) {
	if o == nil || o.Ip == nil {
		return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *UnhealthyAppActionsInner) HasIp() bool {
	if o != nil && o.Ip != nil {
		return true
	}

	return false
}

// SetIp gets a reference to the given string and assigns it to the Ip field.
func (o *UnhealthyAppActionsInner) SetIp(v string) {
	o.Ip = &v
}

// GetIndex returns the Index field value if set, zero value otherwise.
func (o *UnhealthyAppActionsInner) GetIndex() int32 {
	if o == nil || o.Index == nil {
		var ret int32
		return ret
	}
	return *o.Index
}

// GetIndexOk returns a tuple with the Index field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnhealthyAppActionsInner) GetIndexOk() (*int32, bool) {
	if o == nil || o.Index == nil {
		return nil, false
	}
	return o.Index, true
}

// HasIndex returns a boolean if a field has been set.
func (o *UnhealthyAppActionsInner) HasIndex() bool {
	if o != nil && o.Index != nil {
		return true
	}

	return false
}

// SetIndex gets a reference to the given int32 and assigns it to the Index field.
func (o *UnhealthyAppActionsInner) SetIndex(v int32) {
	o.Index = &v
}

// GetReportCount returns the ReportCount field value if set, zero value otherwise.
func (o *UnhealthyAppActionsInner) GetReportCount() int32 {
	if o == nil || o.ReportCount == nil {
		var ret int32
		return ret
	}
	return *o.ReportCount
}

// GetReportCountOk returns a tuple with the ReportCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnhealthyAppActionsInner) GetReportCountOk() (*int32, bool) {
	if o == nil || o.ReportCount == nil {
		return nil, false
	}
	return o.ReportCount, true
}

// HasReportCount returns a boolean if a field has been set.
func (o *UnhealthyAppActionsInner) HasReportCount() bool {
	if o != nil && o.ReportCount != nil {
		return true
	}

	return false
}

// SetReportCount gets a reference to the given int32 and assigns it to the ReportCount field.
func (o *UnhealthyAppActionsInner) SetReportCount(v int32) {
	o.ReportCount = &v
}

// GetUserCount returns the UserCount field value if set, zero value otherwise.
func (o *UnhealthyAppActionsInner) GetUserCount() int32 {
	if o == nil || o.UserCount == nil {
		var ret int32
		return ret
	}
	return *o.UserCount
}

// GetUserCountOk returns a tuple with the UserCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnhealthyAppActionsInner) GetUserCountOk() (*int32, bool) {
	if o == nil || o.UserCount == nil {
		return nil, false
	}
	return o.UserCount, true
}

// HasUserCount returns a boolean if a field has been set.
func (o *UnhealthyAppActionsInner) HasUserCount() bool {
	if o != nil && o.UserCount != nil {
		return true
	}

	return false
}

// SetUserCount gets a reference to the given int32 and assigns it to the UserCount field.
func (o *UnhealthyAppActionsInner) SetUserCount(v int32) {
	o.UserCount = &v
}

// GetPortRange returns the PortRange field value if set, zero value otherwise.
func (o *UnhealthyAppActionsInner) GetPortRange() string {
	if o == nil || o.PortRange == nil {
		var ret string
		return ret
	}
	return *o.PortRange
}

// GetPortRangeOk returns a tuple with the PortRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnhealthyAppActionsInner) GetPortRangeOk() (*string, bool) {
	if o == nil || o.PortRange == nil {
		return nil, false
	}
	return o.PortRange, true
}

// HasPortRange returns a boolean if a field has been set.
func (o *UnhealthyAppActionsInner) HasPortRange() bool {
	if o != nil && o.PortRange != nil {
		return true
	}

	return false
}

// SetPortRange gets a reference to the given string and assigns it to the PortRange field.
func (o *UnhealthyAppActionsInner) SetPortRange(v string) {
	o.PortRange = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *UnhealthyAppActionsInner) GetReason() string {
	if o == nil || o.Reason == nil {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnhealthyAppActionsInner) GetReasonOk() (*string, bool) {
	if o == nil || o.Reason == nil {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *UnhealthyAppActionsInner) HasReason() bool {
	if o != nil && o.Reason != nil {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *UnhealthyAppActionsInner) SetReason(v string) {
	o.Reason = &v
}

// GetLastReported returns the LastReported field value if set, zero value otherwise.
func (o *UnhealthyAppActionsInner) GetLastReported() time.Time {
	if o == nil || o.LastReported == nil {
		var ret time.Time
		return ret
	}
	return *o.LastReported
}

// GetLastReportedOk returns a tuple with the LastReported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnhealthyAppActionsInner) GetLastReportedOk() (*time.Time, bool) {
	if o == nil || o.LastReported == nil {
		return nil, false
	}
	return o.LastReported, true
}

// HasLastReported returns a boolean if a field has been set.
func (o *UnhealthyAppActionsInner) HasLastReported() bool {
	if o != nil && o.LastReported != nil {
		return true
	}

	return false
}

// SetLastReported gets a reference to the given time.Time and assigns it to the LastReported field.
func (o *UnhealthyAppActionsInner) SetLastReported(v time.Time) {
	o.LastReported = &v
}

func (o UnhealthyAppActionsInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Ip != nil {
		toSerialize["ip"] = o.Ip
	}
	if o.Index != nil {
		toSerialize["index"] = o.Index
	}
	if o.ReportCount != nil {
		toSerialize["reportCount"] = o.ReportCount
	}
	if o.UserCount != nil {
		toSerialize["userCount"] = o.UserCount
	}
	if o.PortRange != nil {
		toSerialize["portRange"] = o.PortRange
	}
	if o.Reason != nil {
		toSerialize["reason"] = o.Reason
	}
	if o.LastReported != nil {
		toSerialize["lastReported"] = o.LastReported
	}
	return json.Marshal(toSerialize)
}

type NullableUnhealthyAppActionsInner struct {
	value *UnhealthyAppActionsInner
	isSet bool
}

func (v NullableUnhealthyAppActionsInner) Get() *UnhealthyAppActionsInner {
	return v.value
}

func (v *NullableUnhealthyAppActionsInner) Set(val *UnhealthyAppActionsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableUnhealthyAppActionsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableUnhealthyAppActionsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnhealthyAppActionsInner(val *UnhealthyAppActionsInner) *NullableUnhealthyAppActionsInner {
	return &NullableUnhealthyAppActionsInner{value: val, isSet: true}
}

func (v NullableUnhealthyAppActionsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnhealthyAppActionsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
