/*
Appgate SDP Controller REST API

# About   This specification documents the REST API calls for the Appgate SDP Controller.    Please refer to the REST API chapter in the manual or contact Appgate support with any questions about   this functionality. # Getting Started   Requirements for API scripting:   - Access to the Admin/API TLS Connection (default port 8443) of a Controller appliance.     (https://sdphelp.appgate.com/adminguide/appliance-function-configure.html?anchor=admin-api)   - An API user with relevant permissions.     (https://sdphelp.appgate.com/adminguide/administrative-roles-configure.html)   - In order to use the simple login API, Admin MFA must be disabled or the API user must be excluded.     (https://sdphelp.appgate.com/adminguide/mfa-for-admins.html) # Base path   HTTPS requests must be sent to the Admin Interface hostname and port, with **_/admin** path.    For example: **https://appgate.company.com:8443/admin**    All requests must have the **Accept** header as:    **application/vnd.appgate.peer-v22+json**    An exception is made for the **_/admin/version** endpoint which instead expects an **application/json** Accept header. # API Conventions   API conventions are  important to understand and follow strictly.    - While updating objects (via PUT), entire object must be sent with all fields.     - For example, in order to add a remedy method to the condition below:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": []       }       ```     - send the entire object with updated and non-updated fields:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": [{\"type\": \"DisplayMessage\", \"message\": \"test message\"}]       }       ```    - In case Controller returns an error (non-2xx HTTP status code), response body is JSON.     The \"message\" field contains information about the error.     HTTP 422 \"Unprocessable Entity\" has extra `errors` field to list all the issues with specific fields.    - Empty string (\"\") is considered a different value than \"null\" or field being omitted from JSON.     Omitting the field is recommended if no value is intended.     Empty string (\"\") will be almost always rejected as invalid value.    - There are common pattern between many objects:     - **Configuration Objects**: There are many objects with common fields, namely \"id\", \"name\", \"notes\", \"created\"       and \"updated\". These entities are listed, queried, created, updated and deleted in a similar fashion.     - **Distinguished Name**: Users and Devices are identified with what is called Distinguished Names, as used in        LDAP. The distinguished format that identifies a device and a user combination is        \"CN=\\<Device ID\\>,CN=\\<username\\>,OU=\\<Identity Provider Name\\>\". Some objects have the        \"userDistinguishedName\" field, which does not include the CN for Device ID.        This identifies a user on every device.

API version: API version 22.4
Contact: appgatesdp.support@appgate.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// LocalProviderSettings struct for LocalProviderSettings
type LocalProviderSettings struct {
	// After how many failed authentication attempts will a local user be locked out from authenticating again.
	UserLockoutThreshold *int32 `json:"userLockoutThreshold,omitempty"`
	// For how long lockout will last for local users.
	UserLockoutDurationMinutes *int32 `json:"userLockoutDurationMinutes,omitempty"`
	// Minimum password length requirement for local users.
	MinPasswordLength *int32 `json:"minPasswordLength,omitempty"`
}

// NewLocalProviderSettings instantiates a new LocalProviderSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLocalProviderSettings() *LocalProviderSettings {
	this := LocalProviderSettings{}
	var userLockoutThreshold int32 = 5
	this.UserLockoutThreshold = &userLockoutThreshold
	var userLockoutDurationMinutes int32 = 1
	this.UserLockoutDurationMinutes = &userLockoutDurationMinutes
	var minPasswordLength int32 = 0
	this.MinPasswordLength = &minPasswordLength
	return &this
}

// NewLocalProviderSettingsWithDefaults instantiates a new LocalProviderSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLocalProviderSettingsWithDefaults() *LocalProviderSettings {
	this := LocalProviderSettings{}
	var userLockoutThreshold int32 = 5
	this.UserLockoutThreshold = &userLockoutThreshold
	var userLockoutDurationMinutes int32 = 1
	this.UserLockoutDurationMinutes = &userLockoutDurationMinutes
	var minPasswordLength int32 = 0
	this.MinPasswordLength = &minPasswordLength
	return &this
}

// GetUserLockoutThreshold returns the UserLockoutThreshold field value if set, zero value otherwise.
func (o *LocalProviderSettings) GetUserLockoutThreshold() int32 {
	if o == nil || o.UserLockoutThreshold == nil {
		var ret int32
		return ret
	}
	return *o.UserLockoutThreshold
}

// GetUserLockoutThresholdOk returns a tuple with the UserLockoutThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocalProviderSettings) GetUserLockoutThresholdOk() (*int32, bool) {
	if o == nil || o.UserLockoutThreshold == nil {
		return nil, false
	}
	return o.UserLockoutThreshold, true
}

// HasUserLockoutThreshold returns a boolean if a field has been set.
func (o *LocalProviderSettings) HasUserLockoutThreshold() bool {
	if o != nil && o.UserLockoutThreshold != nil {
		return true
	}

	return false
}

// SetUserLockoutThreshold gets a reference to the given int32 and assigns it to the UserLockoutThreshold field.
func (o *LocalProviderSettings) SetUserLockoutThreshold(v int32) {
	o.UserLockoutThreshold = &v
}

// GetUserLockoutDurationMinutes returns the UserLockoutDurationMinutes field value if set, zero value otherwise.
func (o *LocalProviderSettings) GetUserLockoutDurationMinutes() int32 {
	if o == nil || o.UserLockoutDurationMinutes == nil {
		var ret int32
		return ret
	}
	return *o.UserLockoutDurationMinutes
}

// GetUserLockoutDurationMinutesOk returns a tuple with the UserLockoutDurationMinutes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocalProviderSettings) GetUserLockoutDurationMinutesOk() (*int32, bool) {
	if o == nil || o.UserLockoutDurationMinutes == nil {
		return nil, false
	}
	return o.UserLockoutDurationMinutes, true
}

// HasUserLockoutDurationMinutes returns a boolean if a field has been set.
func (o *LocalProviderSettings) HasUserLockoutDurationMinutes() bool {
	if o != nil && o.UserLockoutDurationMinutes != nil {
		return true
	}

	return false
}

// SetUserLockoutDurationMinutes gets a reference to the given int32 and assigns it to the UserLockoutDurationMinutes field.
func (o *LocalProviderSettings) SetUserLockoutDurationMinutes(v int32) {
	o.UserLockoutDurationMinutes = &v
}

// GetMinPasswordLength returns the MinPasswordLength field value if set, zero value otherwise.
func (o *LocalProviderSettings) GetMinPasswordLength() int32 {
	if o == nil || o.MinPasswordLength == nil {
		var ret int32
		return ret
	}
	return *o.MinPasswordLength
}

// GetMinPasswordLengthOk returns a tuple with the MinPasswordLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocalProviderSettings) GetMinPasswordLengthOk() (*int32, bool) {
	if o == nil || o.MinPasswordLength == nil {
		return nil, false
	}
	return o.MinPasswordLength, true
}

// HasMinPasswordLength returns a boolean if a field has been set.
func (o *LocalProviderSettings) HasMinPasswordLength() bool {
	if o != nil && o.MinPasswordLength != nil {
		return true
	}

	return false
}

// SetMinPasswordLength gets a reference to the given int32 and assigns it to the MinPasswordLength field.
func (o *LocalProviderSettings) SetMinPasswordLength(v int32) {
	o.MinPasswordLength = &v
}

func (o LocalProviderSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UserLockoutThreshold != nil {
		toSerialize["userLockoutThreshold"] = o.UserLockoutThreshold
	}
	if o.UserLockoutDurationMinutes != nil {
		toSerialize["userLockoutDurationMinutes"] = o.UserLockoutDurationMinutes
	}
	if o.MinPasswordLength != nil {
		toSerialize["minPasswordLength"] = o.MinPasswordLength
	}
	return json.Marshal(toSerialize)
}

type NullableLocalProviderSettings struct {
	value *LocalProviderSettings
	isSet bool
}

func (v NullableLocalProviderSettings) Get() *LocalProviderSettings {
	return v.value
}

func (v *NullableLocalProviderSettings) Set(val *LocalProviderSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableLocalProviderSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableLocalProviderSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocalProviderSettings(val *LocalProviderSettings) *NullableLocalProviderSettings {
	return &NullableLocalProviderSettings{value: val, isSet: true}
}

func (v NullableLocalProviderSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocalProviderSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
