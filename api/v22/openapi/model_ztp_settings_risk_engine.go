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

// ZtpSettingsRiskEngine Settings for Risk Engine integration.
type ZtpSettingsRiskEngine struct {
	// Whether the Risk Engine integration is enabled or not.
	Enabled        *bool                                `json:"enabled,omitempty"`
	FallbackValues *ZtpSettingsRiskEngineFallbackValues `json:"fallbackValues,omitempty"`
}

// NewZtpSettingsRiskEngine instantiates a new ZtpSettingsRiskEngine object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewZtpSettingsRiskEngine() *ZtpSettingsRiskEngine {
	this := ZtpSettingsRiskEngine{}
	var enabled bool = true
	this.Enabled = &enabled
	return &this
}

// NewZtpSettingsRiskEngineWithDefaults instantiates a new ZtpSettingsRiskEngine object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewZtpSettingsRiskEngineWithDefaults() *ZtpSettingsRiskEngine {
	this := ZtpSettingsRiskEngine{}
	var enabled bool = true
	this.Enabled = &enabled
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *ZtpSettingsRiskEngine) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZtpSettingsRiskEngine) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *ZtpSettingsRiskEngine) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *ZtpSettingsRiskEngine) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetFallbackValues returns the FallbackValues field value if set, zero value otherwise.
func (o *ZtpSettingsRiskEngine) GetFallbackValues() ZtpSettingsRiskEngineFallbackValues {
	if o == nil || o.FallbackValues == nil {
		var ret ZtpSettingsRiskEngineFallbackValues
		return ret
	}
	return *o.FallbackValues
}

// GetFallbackValuesOk returns a tuple with the FallbackValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZtpSettingsRiskEngine) GetFallbackValuesOk() (*ZtpSettingsRiskEngineFallbackValues, bool) {
	if o == nil || o.FallbackValues == nil {
		return nil, false
	}
	return o.FallbackValues, true
}

// HasFallbackValues returns a boolean if a field has been set.
func (o *ZtpSettingsRiskEngine) HasFallbackValues() bool {
	if o != nil && o.FallbackValues != nil {
		return true
	}

	return false
}

// SetFallbackValues gets a reference to the given ZtpSettingsRiskEngineFallbackValues and assigns it to the FallbackValues field.
func (o *ZtpSettingsRiskEngine) SetFallbackValues(v ZtpSettingsRiskEngineFallbackValues) {
	o.FallbackValues = &v
}

func (o ZtpSettingsRiskEngine) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.FallbackValues != nil {
		toSerialize["fallbackValues"] = o.FallbackValues
	}
	return json.Marshal(toSerialize)
}

type NullableZtpSettingsRiskEngine struct {
	value *ZtpSettingsRiskEngine
	isSet bool
}

func (v NullableZtpSettingsRiskEngine) Get() *ZtpSettingsRiskEngine {
	return v.value
}

func (v *NullableZtpSettingsRiskEngine) Set(val *ZtpSettingsRiskEngine) {
	v.value = val
	v.isSet = true
}

func (v NullableZtpSettingsRiskEngine) IsSet() bool {
	return v.isSet
}

func (v *NullableZtpSettingsRiskEngine) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableZtpSettingsRiskEngine(val *ZtpSettingsRiskEngine) *NullableZtpSettingsRiskEngine {
	return &NullableZtpSettingsRiskEngine{value: val, isSet: true}
}

func (v NullableZtpSettingsRiskEngine) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableZtpSettingsRiskEngine) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
