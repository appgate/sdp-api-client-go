/*
Appgate SDP Controller REST API

# About   This specification documents the REST API calls for the Appgate SDP Controller.    Please refer to the REST API chapter in the manual or contact Appgate support with any questions about   this functionality. # Getting Started   Requirements for API scripting:   - Access to the Admin/API TLS Connection (default port 8443) of a Controller appliance.     (https://sdphelp.appgate.com/adminguide/appliance-function-configure.html?anchor=admin-api)   - An API user with relevant permissions.     (https://sdphelp.appgate.com/adminguide/administrative-roles-configure.html)   - In order to use the simple login API, Admin MFA must be disabled or the API user must be excluded.     (https://sdphelp.appgate.com/adminguide/mfa-for-admins.html) # Base path   HTTPS requests must be sent to the Admin Interface hostname and port, with **_/admin** path.    For example: **https://appgate.company.com:8443/admin**    All requests must have the **Accept** header as:    **application/vnd.appgate.peer-v18+json** # API Conventions   API conventions are  important to understand and follow strictly.    - While updating objects (via PUT), entire object must be sent with all fields.     - For example, in order to add a remedy method to the condition below:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": []       }       ```     - send the entire object with updated and non-updated fields:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": [{\"type\": \"DisplayMessage\", \"message\": \"test message\"}]       }       ```    - In case Controller returns an error (non-2xx HTTP status code), response body is JSON.     The \"message\" field contains information about the error.     HTTP 422 \"Unprocessable Entity\" has extra `errors` field to list all the issues with specific fields.    - Empty string (\"\") is considered a different value than \"null\" or field being omitted from JSON.     Omitting the field is recommended if no value is intended.     Empty string (\"\") will be almost always rejected as invalid value.    - There are common pattern between many objects:     - **Configuration Objects**: There are many objects with common fields, namely \"id\", \"name\", \"notes\", \"created\"       and \"updated\". These entities are listed, queried, created, updated and deleted in a similar fashion.     - **Distinguished Name**: Users and Devices are identified with what is called Distinguished Names, as used in        LDAP. The distinguished format that identifies a device and a user combination is        \"CN=\\<Device ID\\>,CN=\\<username\\>,OU=\\<Identity Provider Name\\>\". Some objects have the        \"userDistinguishedName\" field, which does not include the CN for Device ID.        This identifies a user on every device.

API version: API version 18.4
Contact: appgatesdp.support@appgate.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// LdapProviderAllOfPasswordWarning Password warning configuration for Active Directory. If enabled, the client will display the configured message before the password expiration.
type LdapProviderAllOfPasswordWarning struct {
	// Whether to check and warn the users for password expiration.
	Enabled *bool `json:"enabled,omitempty"`
	// How many days before the password warning to be displayed to the user.
	ThresholdDays *int32 `json:"thresholdDays,omitempty"`
	// The given message will be displayed to the user. Use this field to guide the users on how to change their passwords. The expiration time will displayed on the client on a separate section.
	Message *string `json:"message,omitempty"`
}

// NewLdapProviderAllOfPasswordWarning instantiates a new LdapProviderAllOfPasswordWarning object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLdapProviderAllOfPasswordWarning() *LdapProviderAllOfPasswordWarning {
	this := LdapProviderAllOfPasswordWarning{}
	var thresholdDays int32 = 5
	this.ThresholdDays = &thresholdDays
	return &this
}

// NewLdapProviderAllOfPasswordWarningWithDefaults instantiates a new LdapProviderAllOfPasswordWarning object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLdapProviderAllOfPasswordWarningWithDefaults() *LdapProviderAllOfPasswordWarning {
	this := LdapProviderAllOfPasswordWarning{}
	var thresholdDays int32 = 5
	this.ThresholdDays = &thresholdDays
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *LdapProviderAllOfPasswordWarning) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapProviderAllOfPasswordWarning) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *LdapProviderAllOfPasswordWarning) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *LdapProviderAllOfPasswordWarning) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetThresholdDays returns the ThresholdDays field value if set, zero value otherwise.
func (o *LdapProviderAllOfPasswordWarning) GetThresholdDays() int32 {
	if o == nil || o.ThresholdDays == nil {
		var ret int32
		return ret
	}
	return *o.ThresholdDays
}

// GetThresholdDaysOk returns a tuple with the ThresholdDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapProviderAllOfPasswordWarning) GetThresholdDaysOk() (*int32, bool) {
	if o == nil || o.ThresholdDays == nil {
		return nil, false
	}
	return o.ThresholdDays, true
}

// HasThresholdDays returns a boolean if a field has been set.
func (o *LdapProviderAllOfPasswordWarning) HasThresholdDays() bool {
	if o != nil && o.ThresholdDays != nil {
		return true
	}

	return false
}

// SetThresholdDays gets a reference to the given int32 and assigns it to the ThresholdDays field.
func (o *LdapProviderAllOfPasswordWarning) SetThresholdDays(v int32) {
	o.ThresholdDays = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *LdapProviderAllOfPasswordWarning) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapProviderAllOfPasswordWarning) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *LdapProviderAllOfPasswordWarning) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *LdapProviderAllOfPasswordWarning) SetMessage(v string) {
	o.Message = &v
}

func (o LdapProviderAllOfPasswordWarning) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.ThresholdDays != nil {
		toSerialize["thresholdDays"] = o.ThresholdDays
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	return json.Marshal(toSerialize)
}

type NullableLdapProviderAllOfPasswordWarning struct {
	value *LdapProviderAllOfPasswordWarning
	isSet bool
}

func (v NullableLdapProviderAllOfPasswordWarning) Get() *LdapProviderAllOfPasswordWarning {
	return v.value
}

func (v *NullableLdapProviderAllOfPasswordWarning) Set(val *LdapProviderAllOfPasswordWarning) {
	v.value = val
	v.isSet = true
}

func (v NullableLdapProviderAllOfPasswordWarning) IsSet() bool {
	return v.isSet
}

func (v *NullableLdapProviderAllOfPasswordWarning) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLdapProviderAllOfPasswordWarning(val *LdapProviderAllOfPasswordWarning) *NullableLdapProviderAllOfPasswordWarning {
	return &NullableLdapProviderAllOfPasswordWarning{value: val, isSet: true}
}

func (v NullableLdapProviderAllOfPasswordWarning) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLdapProviderAllOfPasswordWarning) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
