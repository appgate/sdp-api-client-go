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

// AutoUpdateSettings struct for AutoUpdateSettings
type AutoUpdateSettings struct {
	// Whether the Client Auto-Update is enabled or not.
	Enabled *bool `json:"enabled,omitempty"`
	// The Criteria Script to evaluate the Client claims during authorization in order to decide whether the Client Auto-Update will be applied or not.
	CriteriaScript *string `json:"criteriaScript,omitempty"`
	Windows        *Client `json:"windows,omitempty"`
	MacOS          *Client `json:"macOS,omitempty"`
	Ubuntu         *Client `json:"ubuntu,omitempty"`
	Fedora         *Client `json:"fedora,omitempty"`
	RedHat8        *Client `json:"redHat8,omitempty"`
}

// NewAutoUpdateSettings instantiates a new AutoUpdateSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAutoUpdateSettings() *AutoUpdateSettings {
	this := AutoUpdateSettings{}
	return &this
}

// NewAutoUpdateSettingsWithDefaults instantiates a new AutoUpdateSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAutoUpdateSettingsWithDefaults() *AutoUpdateSettings {
	this := AutoUpdateSettings{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *AutoUpdateSettings) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoUpdateSettings) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *AutoUpdateSettings) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *AutoUpdateSettings) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetCriteriaScript returns the CriteriaScript field value if set, zero value otherwise.
func (o *AutoUpdateSettings) GetCriteriaScript() string {
	if o == nil || o.CriteriaScript == nil {
		var ret string
		return ret
	}
	return *o.CriteriaScript
}

// GetCriteriaScriptOk returns a tuple with the CriteriaScript field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoUpdateSettings) GetCriteriaScriptOk() (*string, bool) {
	if o == nil || o.CriteriaScript == nil {
		return nil, false
	}
	return o.CriteriaScript, true
}

// HasCriteriaScript returns a boolean if a field has been set.
func (o *AutoUpdateSettings) HasCriteriaScript() bool {
	if o != nil && o.CriteriaScript != nil {
		return true
	}

	return false
}

// SetCriteriaScript gets a reference to the given string and assigns it to the CriteriaScript field.
func (o *AutoUpdateSettings) SetCriteriaScript(v string) {
	o.CriteriaScript = &v
}

// GetWindows returns the Windows field value if set, zero value otherwise.
func (o *AutoUpdateSettings) GetWindows() Client {
	if o == nil || o.Windows == nil {
		var ret Client
		return ret
	}
	return *o.Windows
}

// GetWindowsOk returns a tuple with the Windows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoUpdateSettings) GetWindowsOk() (*Client, bool) {
	if o == nil || o.Windows == nil {
		return nil, false
	}
	return o.Windows, true
}

// HasWindows returns a boolean if a field has been set.
func (o *AutoUpdateSettings) HasWindows() bool {
	if o != nil && o.Windows != nil {
		return true
	}

	return false
}

// SetWindows gets a reference to the given Client and assigns it to the Windows field.
func (o *AutoUpdateSettings) SetWindows(v Client) {
	o.Windows = &v
}

// GetMacOS returns the MacOS field value if set, zero value otherwise.
func (o *AutoUpdateSettings) GetMacOS() Client {
	if o == nil || o.MacOS == nil {
		var ret Client
		return ret
	}
	return *o.MacOS
}

// GetMacOSOk returns a tuple with the MacOS field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoUpdateSettings) GetMacOSOk() (*Client, bool) {
	if o == nil || o.MacOS == nil {
		return nil, false
	}
	return o.MacOS, true
}

// HasMacOS returns a boolean if a field has been set.
func (o *AutoUpdateSettings) HasMacOS() bool {
	if o != nil && o.MacOS != nil {
		return true
	}

	return false
}

// SetMacOS gets a reference to the given Client and assigns it to the MacOS field.
func (o *AutoUpdateSettings) SetMacOS(v Client) {
	o.MacOS = &v
}

// GetUbuntu returns the Ubuntu field value if set, zero value otherwise.
func (o *AutoUpdateSettings) GetUbuntu() Client {
	if o == nil || o.Ubuntu == nil {
		var ret Client
		return ret
	}
	return *o.Ubuntu
}

// GetUbuntuOk returns a tuple with the Ubuntu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoUpdateSettings) GetUbuntuOk() (*Client, bool) {
	if o == nil || o.Ubuntu == nil {
		return nil, false
	}
	return o.Ubuntu, true
}

// HasUbuntu returns a boolean if a field has been set.
func (o *AutoUpdateSettings) HasUbuntu() bool {
	if o != nil && o.Ubuntu != nil {
		return true
	}

	return false
}

// SetUbuntu gets a reference to the given Client and assigns it to the Ubuntu field.
func (o *AutoUpdateSettings) SetUbuntu(v Client) {
	o.Ubuntu = &v
}

// GetFedora returns the Fedora field value if set, zero value otherwise.
func (o *AutoUpdateSettings) GetFedora() Client {
	if o == nil || o.Fedora == nil {
		var ret Client
		return ret
	}
	return *o.Fedora
}

// GetFedoraOk returns a tuple with the Fedora field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoUpdateSettings) GetFedoraOk() (*Client, bool) {
	if o == nil || o.Fedora == nil {
		return nil, false
	}
	return o.Fedora, true
}

// HasFedora returns a boolean if a field has been set.
func (o *AutoUpdateSettings) HasFedora() bool {
	if o != nil && o.Fedora != nil {
		return true
	}

	return false
}

// SetFedora gets a reference to the given Client and assigns it to the Fedora field.
func (o *AutoUpdateSettings) SetFedora(v Client) {
	o.Fedora = &v
}

// GetRedHat8 returns the RedHat8 field value if set, zero value otherwise.
func (o *AutoUpdateSettings) GetRedHat8() Client {
	if o == nil || o.RedHat8 == nil {
		var ret Client
		return ret
	}
	return *o.RedHat8
}

// GetRedHat8Ok returns a tuple with the RedHat8 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoUpdateSettings) GetRedHat8Ok() (*Client, bool) {
	if o == nil || o.RedHat8 == nil {
		return nil, false
	}
	return o.RedHat8, true
}

// HasRedHat8 returns a boolean if a field has been set.
func (o *AutoUpdateSettings) HasRedHat8() bool {
	if o != nil && o.RedHat8 != nil {
		return true
	}

	return false
}

// SetRedHat8 gets a reference to the given Client and assigns it to the RedHat8 field.
func (o *AutoUpdateSettings) SetRedHat8(v Client) {
	o.RedHat8 = &v
}

func (o AutoUpdateSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.CriteriaScript != nil {
		toSerialize["criteriaScript"] = o.CriteriaScript
	}
	if o.Windows != nil {
		toSerialize["windows"] = o.Windows
	}
	if o.MacOS != nil {
		toSerialize["macOS"] = o.MacOS
	}
	if o.Ubuntu != nil {
		toSerialize["ubuntu"] = o.Ubuntu
	}
	if o.Fedora != nil {
		toSerialize["fedora"] = o.Fedora
	}
	if o.RedHat8 != nil {
		toSerialize["redHat8"] = o.RedHat8
	}
	return json.Marshal(toSerialize)
}

type NullableAutoUpdateSettings struct {
	value *AutoUpdateSettings
	isSet bool
}

func (v NullableAutoUpdateSettings) Get() *AutoUpdateSettings {
	return v.value
}

func (v *NullableAutoUpdateSettings) Set(val *AutoUpdateSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableAutoUpdateSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableAutoUpdateSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAutoUpdateSettings(val *AutoUpdateSettings) *NullableAutoUpdateSettings {
	return &NullableAutoUpdateSettings{value: val, isSet: true}
}

func (v NullableAutoUpdateSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAutoUpdateSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
