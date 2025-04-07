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

// Sensitivity Sensitivity settings.
type Sensitivity struct {
	LowRisk    *Risk `json:"lowRisk,omitempty"`
	MediumRisk *Risk `json:"mediumRisk,omitempty"`
	HighRisk   *Risk `json:"highRisk,omitempty"`
}

// NewSensitivity instantiates a new Sensitivity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSensitivity() *Sensitivity {
	this := Sensitivity{}
	return &this
}

// NewSensitivityWithDefaults instantiates a new Sensitivity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSensitivityWithDefaults() *Sensitivity {
	this := Sensitivity{}
	return &this
}

// GetLowRisk returns the LowRisk field value if set, zero value otherwise.
func (o *Sensitivity) GetLowRisk() Risk {
	if o == nil || o.LowRisk == nil {
		var ret Risk
		return ret
	}
	return *o.LowRisk
}

// GetLowRiskOk returns a tuple with the LowRisk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Sensitivity) GetLowRiskOk() (*Risk, bool) {
	if o == nil || o.LowRisk == nil {
		return nil, false
	}
	return o.LowRisk, true
}

// HasLowRisk returns a boolean if a field has been set.
func (o *Sensitivity) HasLowRisk() bool {
	if o != nil && o.LowRisk != nil {
		return true
	}

	return false
}

// SetLowRisk gets a reference to the given Risk and assigns it to the LowRisk field.
func (o *Sensitivity) SetLowRisk(v Risk) {
	o.LowRisk = &v
}

// GetMediumRisk returns the MediumRisk field value if set, zero value otherwise.
func (o *Sensitivity) GetMediumRisk() Risk {
	if o == nil || o.MediumRisk == nil {
		var ret Risk
		return ret
	}
	return *o.MediumRisk
}

// GetMediumRiskOk returns a tuple with the MediumRisk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Sensitivity) GetMediumRiskOk() (*Risk, bool) {
	if o == nil || o.MediumRisk == nil {
		return nil, false
	}
	return o.MediumRisk, true
}

// HasMediumRisk returns a boolean if a field has been set.
func (o *Sensitivity) HasMediumRisk() bool {
	if o != nil && o.MediumRisk != nil {
		return true
	}

	return false
}

// SetMediumRisk gets a reference to the given Risk and assigns it to the MediumRisk field.
func (o *Sensitivity) SetMediumRisk(v Risk) {
	o.MediumRisk = &v
}

// GetHighRisk returns the HighRisk field value if set, zero value otherwise.
func (o *Sensitivity) GetHighRisk() Risk {
	if o == nil || o.HighRisk == nil {
		var ret Risk
		return ret
	}
	return *o.HighRisk
}

// GetHighRiskOk returns a tuple with the HighRisk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Sensitivity) GetHighRiskOk() (*Risk, bool) {
	if o == nil || o.HighRisk == nil {
		return nil, false
	}
	return o.HighRisk, true
}

// HasHighRisk returns a boolean if a field has been set.
func (o *Sensitivity) HasHighRisk() bool {
	if o != nil && o.HighRisk != nil {
		return true
	}

	return false
}

// SetHighRisk gets a reference to the given Risk and assigns it to the HighRisk field.
func (o *Sensitivity) SetHighRisk(v Risk) {
	o.HighRisk = &v
}

func (o Sensitivity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.LowRisk != nil {
		toSerialize["lowRisk"] = o.LowRisk
	}
	if o.MediumRisk != nil {
		toSerialize["mediumRisk"] = o.MediumRisk
	}
	if o.HighRisk != nil {
		toSerialize["highRisk"] = o.HighRisk
	}
	return json.Marshal(toSerialize)
}

type NullableSensitivity struct {
	value *Sensitivity
	isSet bool
}

func (v NullableSensitivity) Get() *Sensitivity {
	return v.value
}

func (v *NullableSensitivity) Set(val *Sensitivity) {
	v.value = val
	v.isSet = true
}

func (v NullableSensitivity) IsSet() bool {
	return v.isSet
}

func (v *NullableSensitivity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSensitivity(val *Sensitivity) *NullableSensitivity {
	return &NullableSensitivity{value: val, isSet: true}
}

func (v NullableSensitivity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSensitivity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
