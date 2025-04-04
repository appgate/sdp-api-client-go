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

// ZtpSettingsRiskEngineFallbackValues Configure fallback risk values when certain scenarios occur.
type ZtpSettingsRiskEngineFallbackValues struct {
	// Fallback risk value when the relevant rule or adapter is not working. 1 - LOW, 2 - MEDIUM, 3 - HIGH
	RuleNotWorking *float32 `json:"ruleNotWorking,omitempty"`
	// Fallback risk value when the Risk Engine cannot find the information for the device. 1 - LOW, 2 - MEDIUM, 3 - HIGH
	DeviceInfoNotAvailable *float32 `json:"deviceInfoNotAvailable,omitempty"`
	// Fallback risk value when Controller cannot reach ZTP services. 1 - LOW, 2 - MEDIUM, 3 - HIGH
	ZtpUnreachable *float32 `json:"ztpUnreachable,omitempty"`
}

// NewZtpSettingsRiskEngineFallbackValues instantiates a new ZtpSettingsRiskEngineFallbackValues object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewZtpSettingsRiskEngineFallbackValues() *ZtpSettingsRiskEngineFallbackValues {
	this := ZtpSettingsRiskEngineFallbackValues{}
	var ruleNotWorking float32 = 3
	this.RuleNotWorking = &ruleNotWorking
	var deviceInfoNotAvailable float32 = 3
	this.DeviceInfoNotAvailable = &deviceInfoNotAvailable
	var ztpUnreachable float32 = 3
	this.ZtpUnreachable = &ztpUnreachable
	return &this
}

// NewZtpSettingsRiskEngineFallbackValuesWithDefaults instantiates a new ZtpSettingsRiskEngineFallbackValues object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewZtpSettingsRiskEngineFallbackValuesWithDefaults() *ZtpSettingsRiskEngineFallbackValues {
	this := ZtpSettingsRiskEngineFallbackValues{}
	var ruleNotWorking float32 = 3
	this.RuleNotWorking = &ruleNotWorking
	var deviceInfoNotAvailable float32 = 3
	this.DeviceInfoNotAvailable = &deviceInfoNotAvailable
	var ztpUnreachable float32 = 3
	this.ZtpUnreachable = &ztpUnreachable
	return &this
}

// GetRuleNotWorking returns the RuleNotWorking field value if set, zero value otherwise.
func (o *ZtpSettingsRiskEngineFallbackValues) GetRuleNotWorking() float32 {
	if o == nil || o.RuleNotWorking == nil {
		var ret float32
		return ret
	}
	return *o.RuleNotWorking
}

// GetRuleNotWorkingOk returns a tuple with the RuleNotWorking field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZtpSettingsRiskEngineFallbackValues) GetRuleNotWorkingOk() (*float32, bool) {
	if o == nil || o.RuleNotWorking == nil {
		return nil, false
	}
	return o.RuleNotWorking, true
}

// HasRuleNotWorking returns a boolean if a field has been set.
func (o *ZtpSettingsRiskEngineFallbackValues) HasRuleNotWorking() bool {
	if o != nil && o.RuleNotWorking != nil {
		return true
	}

	return false
}

// SetRuleNotWorking gets a reference to the given float32 and assigns it to the RuleNotWorking field.
func (o *ZtpSettingsRiskEngineFallbackValues) SetRuleNotWorking(v float32) {
	o.RuleNotWorking = &v
}

// GetDeviceInfoNotAvailable returns the DeviceInfoNotAvailable field value if set, zero value otherwise.
func (o *ZtpSettingsRiskEngineFallbackValues) GetDeviceInfoNotAvailable() float32 {
	if o == nil || o.DeviceInfoNotAvailable == nil {
		var ret float32
		return ret
	}
	return *o.DeviceInfoNotAvailable
}

// GetDeviceInfoNotAvailableOk returns a tuple with the DeviceInfoNotAvailable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZtpSettingsRiskEngineFallbackValues) GetDeviceInfoNotAvailableOk() (*float32, bool) {
	if o == nil || o.DeviceInfoNotAvailable == nil {
		return nil, false
	}
	return o.DeviceInfoNotAvailable, true
}

// HasDeviceInfoNotAvailable returns a boolean if a field has been set.
func (o *ZtpSettingsRiskEngineFallbackValues) HasDeviceInfoNotAvailable() bool {
	if o != nil && o.DeviceInfoNotAvailable != nil {
		return true
	}

	return false
}

// SetDeviceInfoNotAvailable gets a reference to the given float32 and assigns it to the DeviceInfoNotAvailable field.
func (o *ZtpSettingsRiskEngineFallbackValues) SetDeviceInfoNotAvailable(v float32) {
	o.DeviceInfoNotAvailable = &v
}

// GetZtpUnreachable returns the ZtpUnreachable field value if set, zero value otherwise.
func (o *ZtpSettingsRiskEngineFallbackValues) GetZtpUnreachable() float32 {
	if o == nil || o.ZtpUnreachable == nil {
		var ret float32
		return ret
	}
	return *o.ZtpUnreachable
}

// GetZtpUnreachableOk returns a tuple with the ZtpUnreachable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZtpSettingsRiskEngineFallbackValues) GetZtpUnreachableOk() (*float32, bool) {
	if o == nil || o.ZtpUnreachable == nil {
		return nil, false
	}
	return o.ZtpUnreachable, true
}

// HasZtpUnreachable returns a boolean if a field has been set.
func (o *ZtpSettingsRiskEngineFallbackValues) HasZtpUnreachable() bool {
	if o != nil && o.ZtpUnreachable != nil {
		return true
	}

	return false
}

// SetZtpUnreachable gets a reference to the given float32 and assigns it to the ZtpUnreachable field.
func (o *ZtpSettingsRiskEngineFallbackValues) SetZtpUnreachable(v float32) {
	o.ZtpUnreachable = &v
}

func (o ZtpSettingsRiskEngineFallbackValues) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RuleNotWorking != nil {
		toSerialize["ruleNotWorking"] = o.RuleNotWorking
	}
	if o.DeviceInfoNotAvailable != nil {
		toSerialize["deviceInfoNotAvailable"] = o.DeviceInfoNotAvailable
	}
	if o.ZtpUnreachable != nil {
		toSerialize["ztpUnreachable"] = o.ZtpUnreachable
	}
	return json.Marshal(toSerialize)
}

type NullableZtpSettingsRiskEngineFallbackValues struct {
	value *ZtpSettingsRiskEngineFallbackValues
	isSet bool
}

func (v NullableZtpSettingsRiskEngineFallbackValues) Get() *ZtpSettingsRiskEngineFallbackValues {
	return v.value
}

func (v *NullableZtpSettingsRiskEngineFallbackValues) Set(val *ZtpSettingsRiskEngineFallbackValues) {
	v.value = val
	v.isSet = true
}

func (v NullableZtpSettingsRiskEngineFallbackValues) IsSet() bool {
	return v.isSet
}

func (v *NullableZtpSettingsRiskEngineFallbackValues) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableZtpSettingsRiskEngineFallbackValues(val *ZtpSettingsRiskEngineFallbackValues) *NullableZtpSettingsRiskEngineFallbackValues {
	return &NullableZtpSettingsRiskEngineFallbackValues{value: val, isSet: true}
}

func (v NullableZtpSettingsRiskEngineFallbackValues) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableZtpSettingsRiskEngineFallbackValues) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
