/*
Appgate SDP Controller REST API

# About   This specification documents the REST API calls for the Appgate SDP Controller.    Please refer to the REST API chapter in the manual or contact Appgate support with any questions about   this functionality. # Getting Started   Requirements for API scripting:   - Access to the Admin/API TLS Connection (default port 8443) of a Controller appliance.     (https://sdphelp.appgate.com/adminguide/appliance-function-configure.html?anchor=admin-api)   - An API user with relevant permissions.     (https://sdphelp.appgate.com/adminguide/administrative-roles-configure.html)   - In order to use the simple login API, Admin MFA must be disabled or the API user must be excluded.     (https://sdphelp.appgate.com/adminguide/mfa-for-admins.html) # Base path   HTTPS requests must be sent to the Admin Interface hostname and port, with **_/admin** path.    For example: **https://appgate.company.com:8443/admin**    All requests must have the **Accept** header as:    **application/vnd.appgate.peer-v22+json**    An exception is made for the **_/admin/version** endpoint which instead expects an **application/json** Accept header. # API Conventions   API conventions are  important to understand and follow strictly.    - While updating objects (via PUT), entire object must be sent with all fields.     - For example, in order to add a remedy method to the condition below:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": []       }       ```     - send the entire object with updated and non-updated fields:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": [{\"type\": \"DisplayMessage\", \"message\": \"test message\"}]       }       ```    - In case Controller returns an error (non-2xx HTTP status code), response body is JSON.     The \"message\" field contains information about the error.     HTTP 422 \"Unprocessable Entity\" has extra `errors` field to list all the issues with specific fields.    - Empty string (\"\") is considered a different value than \"null\" or field being omitted from JSON.     Omitting the field is recommended if no value is intended.     Empty string (\"\") will be almost always rejected as invalid value.    - There are common pattern between many objects:     - **Configuration Objects**: There are many objects with common fields, namely \"id\", \"name\", \"notes\", \"created\"       and \"updated\". These entities are listed, queried, created, updated and deleted in a similar fashion.     - **Distinguished Name**: Users and Devices are identified with what is called Distinguished Names, as used in        LDAP. The distinguished format that identifies a device and a user combination is        \"CN=\\<Device ID\\>,CN=\\<username\\>,OU=\\<Identity Provider Name\\>\". Some objects have the        \"userDistinguishedName\" field, which does not include the CN for Device ID.        This identifies a user on every device.

API version: API version 22.1
Contact: appgatesdp.support@appgate.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// RiskModel struct for RiskModel
type RiskModel struct {
	// ID of the rule in ZTP to use as the source of the risk score.
	ZtpRuleId         *string      `json:"ztpRuleId,omitempty"`
	LowSensitivity    *Sensitivity `json:"lowSensitivity,omitempty"`
	MediumSensitivity *Sensitivity `json:"mediumSensitivity,omitempty"`
	HighSensitivity   *Sensitivity `json:"highSensitivity,omitempty"`
}

// NewRiskModel instantiates a new RiskModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRiskModel() *RiskModel {
	this := RiskModel{}
	return &this
}

// NewRiskModelWithDefaults instantiates a new RiskModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRiskModelWithDefaults() *RiskModel {
	this := RiskModel{}
	return &this
}

// GetZtpRuleId returns the ZtpRuleId field value if set, zero value otherwise.
func (o *RiskModel) GetZtpRuleId() string {
	if o == nil || o.ZtpRuleId == nil {
		var ret string
		return ret
	}
	return *o.ZtpRuleId
}

// GetZtpRuleIdOk returns a tuple with the ZtpRuleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskModel) GetZtpRuleIdOk() (*string, bool) {
	if o == nil || o.ZtpRuleId == nil {
		return nil, false
	}
	return o.ZtpRuleId, true
}

// HasZtpRuleId returns a boolean if a field has been set.
func (o *RiskModel) HasZtpRuleId() bool {
	if o != nil && o.ZtpRuleId != nil {
		return true
	}

	return false
}

// SetZtpRuleId gets a reference to the given string and assigns it to the ZtpRuleId field.
func (o *RiskModel) SetZtpRuleId(v string) {
	o.ZtpRuleId = &v
}

// GetLowSensitivity returns the LowSensitivity field value if set, zero value otherwise.
func (o *RiskModel) GetLowSensitivity() Sensitivity {
	if o == nil || o.LowSensitivity == nil {
		var ret Sensitivity
		return ret
	}
	return *o.LowSensitivity
}

// GetLowSensitivityOk returns a tuple with the LowSensitivity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskModel) GetLowSensitivityOk() (*Sensitivity, bool) {
	if o == nil || o.LowSensitivity == nil {
		return nil, false
	}
	return o.LowSensitivity, true
}

// HasLowSensitivity returns a boolean if a field has been set.
func (o *RiskModel) HasLowSensitivity() bool {
	if o != nil && o.LowSensitivity != nil {
		return true
	}

	return false
}

// SetLowSensitivity gets a reference to the given Sensitivity and assigns it to the LowSensitivity field.
func (o *RiskModel) SetLowSensitivity(v Sensitivity) {
	o.LowSensitivity = &v
}

// GetMediumSensitivity returns the MediumSensitivity field value if set, zero value otherwise.
func (o *RiskModel) GetMediumSensitivity() Sensitivity {
	if o == nil || o.MediumSensitivity == nil {
		var ret Sensitivity
		return ret
	}
	return *o.MediumSensitivity
}

// GetMediumSensitivityOk returns a tuple with the MediumSensitivity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskModel) GetMediumSensitivityOk() (*Sensitivity, bool) {
	if o == nil || o.MediumSensitivity == nil {
		return nil, false
	}
	return o.MediumSensitivity, true
}

// HasMediumSensitivity returns a boolean if a field has been set.
func (o *RiskModel) HasMediumSensitivity() bool {
	if o != nil && o.MediumSensitivity != nil {
		return true
	}

	return false
}

// SetMediumSensitivity gets a reference to the given Sensitivity and assigns it to the MediumSensitivity field.
func (o *RiskModel) SetMediumSensitivity(v Sensitivity) {
	o.MediumSensitivity = &v
}

// GetHighSensitivity returns the HighSensitivity field value if set, zero value otherwise.
func (o *RiskModel) GetHighSensitivity() Sensitivity {
	if o == nil || o.HighSensitivity == nil {
		var ret Sensitivity
		return ret
	}
	return *o.HighSensitivity
}

// GetHighSensitivityOk returns a tuple with the HighSensitivity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskModel) GetHighSensitivityOk() (*Sensitivity, bool) {
	if o == nil || o.HighSensitivity == nil {
		return nil, false
	}
	return o.HighSensitivity, true
}

// HasHighSensitivity returns a boolean if a field has been set.
func (o *RiskModel) HasHighSensitivity() bool {
	if o != nil && o.HighSensitivity != nil {
		return true
	}

	return false
}

// SetHighSensitivity gets a reference to the given Sensitivity and assigns it to the HighSensitivity field.
func (o *RiskModel) SetHighSensitivity(v Sensitivity) {
	o.HighSensitivity = &v
}

func (o RiskModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ZtpRuleId != nil {
		toSerialize["ztpRuleId"] = o.ZtpRuleId
	}
	if o.LowSensitivity != nil {
		toSerialize["lowSensitivity"] = o.LowSensitivity
	}
	if o.MediumSensitivity != nil {
		toSerialize["mediumSensitivity"] = o.MediumSensitivity
	}
	if o.HighSensitivity != nil {
		toSerialize["highSensitivity"] = o.HighSensitivity
	}
	return json.Marshal(toSerialize)
}

type NullableRiskModel struct {
	value *RiskModel
	isSet bool
}

func (v NullableRiskModel) Get() *RiskModel {
	return v.value
}

func (v *NullableRiskModel) Set(val *RiskModel) {
	v.value = val
	v.isSet = true
}

func (v NullableRiskModel) IsSet() bool {
	return v.isSet
}

func (v *NullableRiskModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRiskModel(val *RiskModel) *NullableRiskModel {
	return &NullableRiskModel{value: val, isSet: true}
}

func (v NullableRiskModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRiskModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
