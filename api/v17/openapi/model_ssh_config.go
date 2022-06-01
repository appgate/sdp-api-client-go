/*
Appgate SDP Controller REST API

# About   This specification documents the REST API calls for the Appgate SDP Controller.    Please refer to the REST API chapter in the manual or contact Appgate support with any questions about   this functionality. # Getting Started   Requirements for API scripting:   - Access to the Admin/API TLS Connection (default port 8443) of a Controller appliance.     (https://sdphelp.appgate.com/adminguide/appliance-function-configure.html?anchor=admin-api)   - An API user with relevant permissions.     (https://sdphelp.appgate.com/adminguide/administrative-roles-configure.html)   - In order to use the simple login API, Admin MFA must be disabled or the API user must be excluded.     (https://sdphelp.appgate.com/adminguide/mfa-for-admins.html) # Base path   HTTPS requests must be sent to the Admin Interface hostname and port, with **_/admin** path.    For example: **https://appgate.company.com:8443/admin**    All requests must have the **Accept** header as:    **application/vnd.appgate.peer-v17+json** # API Conventions   API conventions are  important to understand and follow strictly.    - While updating objects (via PUT), entire object must be sent with all fields.     - For example, in order to add a remedy method to the condition below:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": []       }       ```     - send the entire object with updated and non-updated fields:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": [{\"type\": \"DisplayMessage\", \"message\": \"test message\"}]       }       ```    - In case Controller returns an error (non-2xx HTTP status code), response body is JSON.     The \"message\" field contains information about the error.     HTTP 422 \"Unprocessable Entity\" has extra `errors` field to list all the issues with specific fields.    - Empty string (\"\") is considered a different value than \"null\" or field being omitted from JSON.     Omitting the field is recommended if no value is intended.     Empty string (\"\") will be almost always rejected as invalid value.    - There are common pattern between many objects:     - **Configuration Objects**: There are many objects with common fields, namely \"id\", \"name\", \"notes\", \"created\"       and \"updated\". These entities are listed, queried, created, updated and deleted in a similar fashion.     - **Distinguished Name**: Users and Devices are identified with what is called Distinguished Names, as used in        LDAP. The distinguished format that identifies a device and a user combination is        \"CN=\\<Device ID\\>,CN=\\<username\\>,OU=\\<Identity Provider Name\\>\". Some objects have the        \"userDistinguishedName\" field, which does not include the CN for Device ID.        This identifies a user on every device.

API version: API version 17.1
Contact: appgatesdp.support@appgate.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// SSHConfig SSH configuration during seeding.
type SSHConfig struct {
	// Tells appliance to use the key generated by AWS or Azure.
	ProvideCloudSSHKey *bool `json:"provideCloudSSHKey,omitempty"`
	// SSH public key to allow.
	SshKey *string `json:"sshKey,omitempty"`
	// Appliance's CZ user password.
	Password *string `json:"password,omitempty"`
	// Whether the Appliance should allow customizations or not.
	AllowCustomization *bool `json:"allowCustomization,omitempty"`
	// How many days the seed should be valid for.
	ValidityDays *float32 `json:"validityDays,omitempty"`
}

// NewSSHConfig instantiates a new SSHConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSSHConfig() *SSHConfig {
	this := SSHConfig{}
	var allowCustomization bool = true
	this.AllowCustomization = &allowCustomization
	var validityDays float32 = 1
	this.ValidityDays = &validityDays
	return &this
}

// NewSSHConfigWithDefaults instantiates a new SSHConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSSHConfigWithDefaults() *SSHConfig {
	this := SSHConfig{}
	var allowCustomization bool = true
	this.AllowCustomization = &allowCustomization
	var validityDays float32 = 1
	this.ValidityDays = &validityDays
	return &this
}

// GetProvideCloudSSHKey returns the ProvideCloudSSHKey field value if set, zero value otherwise.
func (o *SSHConfig) GetProvideCloudSSHKey() bool {
	if o == nil || o.ProvideCloudSSHKey == nil {
		var ret bool
		return ret
	}
	return *o.ProvideCloudSSHKey
}

// GetProvideCloudSSHKeyOk returns a tuple with the ProvideCloudSSHKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SSHConfig) GetProvideCloudSSHKeyOk() (*bool, bool) {
	if o == nil || o.ProvideCloudSSHKey == nil {
		return nil, false
	}
	return o.ProvideCloudSSHKey, true
}

// HasProvideCloudSSHKey returns a boolean if a field has been set.
func (o *SSHConfig) HasProvideCloudSSHKey() bool {
	if o != nil && o.ProvideCloudSSHKey != nil {
		return true
	}

	return false
}

// SetProvideCloudSSHKey gets a reference to the given bool and assigns it to the ProvideCloudSSHKey field.
func (o *SSHConfig) SetProvideCloudSSHKey(v bool) {
	o.ProvideCloudSSHKey = &v
}

// GetSshKey returns the SshKey field value if set, zero value otherwise.
func (o *SSHConfig) GetSshKey() string {
	if o == nil || o.SshKey == nil {
		var ret string
		return ret
	}
	return *o.SshKey
}

// GetSshKeyOk returns a tuple with the SshKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SSHConfig) GetSshKeyOk() (*string, bool) {
	if o == nil || o.SshKey == nil {
		return nil, false
	}
	return o.SshKey, true
}

// HasSshKey returns a boolean if a field has been set.
func (o *SSHConfig) HasSshKey() bool {
	if o != nil && o.SshKey != nil {
		return true
	}

	return false
}

// SetSshKey gets a reference to the given string and assigns it to the SshKey field.
func (o *SSHConfig) SetSshKey(v string) {
	o.SshKey = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *SSHConfig) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SSHConfig) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *SSHConfig) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *SSHConfig) SetPassword(v string) {
	o.Password = &v
}

// GetAllowCustomization returns the AllowCustomization field value if set, zero value otherwise.
func (o *SSHConfig) GetAllowCustomization() bool {
	if o == nil || o.AllowCustomization == nil {
		var ret bool
		return ret
	}
	return *o.AllowCustomization
}

// GetAllowCustomizationOk returns a tuple with the AllowCustomization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SSHConfig) GetAllowCustomizationOk() (*bool, bool) {
	if o == nil || o.AllowCustomization == nil {
		return nil, false
	}
	return o.AllowCustomization, true
}

// HasAllowCustomization returns a boolean if a field has been set.
func (o *SSHConfig) HasAllowCustomization() bool {
	if o != nil && o.AllowCustomization != nil {
		return true
	}

	return false
}

// SetAllowCustomization gets a reference to the given bool and assigns it to the AllowCustomization field.
func (o *SSHConfig) SetAllowCustomization(v bool) {
	o.AllowCustomization = &v
}

// GetValidityDays returns the ValidityDays field value if set, zero value otherwise.
func (o *SSHConfig) GetValidityDays() float32 {
	if o == nil || o.ValidityDays == nil {
		var ret float32
		return ret
	}
	return *o.ValidityDays
}

// GetValidityDaysOk returns a tuple with the ValidityDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SSHConfig) GetValidityDaysOk() (*float32, bool) {
	if o == nil || o.ValidityDays == nil {
		return nil, false
	}
	return o.ValidityDays, true
}

// HasValidityDays returns a boolean if a field has been set.
func (o *SSHConfig) HasValidityDays() bool {
	if o != nil && o.ValidityDays != nil {
		return true
	}

	return false
}

// SetValidityDays gets a reference to the given float32 and assigns it to the ValidityDays field.
func (o *SSHConfig) SetValidityDays(v float32) {
	o.ValidityDays = &v
}

func (o SSHConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ProvideCloudSSHKey != nil {
		toSerialize["provideCloudSSHKey"] = o.ProvideCloudSSHKey
	}
	if o.SshKey != nil {
		toSerialize["sshKey"] = o.SshKey
	}
	if o.Password != nil {
		toSerialize["password"] = o.Password
	}
	if o.AllowCustomization != nil {
		toSerialize["allowCustomization"] = o.AllowCustomization
	}
	if o.ValidityDays != nil {
		toSerialize["validityDays"] = o.ValidityDays
	}
	return json.Marshal(toSerialize)
}

type NullableSSHConfig struct {
	value *SSHConfig
	isSet bool
}

func (v NullableSSHConfig) Get() *SSHConfig {
	return v.value
}

func (v *NullableSSHConfig) Set(val *SSHConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableSSHConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableSSHConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSSHConfig(val *SSHConfig) *NullableSSHConfig {
	return &NullableSSHConfig{value: val, isSet: true}
}

func (v NullableSSHConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSSHConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
