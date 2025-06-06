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

// ConfigurableIdentityProviderAllOfOnBoarding2FA On-boarding two-factor authentication settings. Leave it empty keep it disabled.
type ConfigurableIdentityProviderAllOfOnBoarding2FA struct {
	// MFA provider ID to use for the authentication.
	MfaProviderId string `json:"mfaProviderId"`
	// On-boarding MFA message to be displayed on the Client UI during the second-factor authentication.
	Message *string `json:"message,omitempty"`
	// Upon successful on-boarding, the claim will be added as if MFA remedy action is fulfilled.
	ClaimSuffix *string `json:"claimSuffix,omitempty"`
	// If enabled, MFA will be required on every authentication.
	AlwaysRequired *bool `json:"alwaysRequired,omitempty"`
}

// NewConfigurableIdentityProviderAllOfOnBoarding2FA instantiates a new ConfigurableIdentityProviderAllOfOnBoarding2FA object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigurableIdentityProviderAllOfOnBoarding2FA(mfaProviderId string) *ConfigurableIdentityProviderAllOfOnBoarding2FA {
	this := ConfigurableIdentityProviderAllOfOnBoarding2FA{}
	this.MfaProviderId = mfaProviderId
	var claimSuffix string = "onBoarding"
	this.ClaimSuffix = &claimSuffix
	return &this
}

// NewConfigurableIdentityProviderAllOfOnBoarding2FAWithDefaults instantiates a new ConfigurableIdentityProviderAllOfOnBoarding2FA object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigurableIdentityProviderAllOfOnBoarding2FAWithDefaults() *ConfigurableIdentityProviderAllOfOnBoarding2FA {
	this := ConfigurableIdentityProviderAllOfOnBoarding2FA{}
	var claimSuffix string = "onBoarding"
	this.ClaimSuffix = &claimSuffix
	return &this
}

// GetMfaProviderId returns the MfaProviderId field value
func (o *ConfigurableIdentityProviderAllOfOnBoarding2FA) GetMfaProviderId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MfaProviderId
}

// GetMfaProviderIdOk returns a tuple with the MfaProviderId field value
// and a boolean to check if the value has been set.
func (o *ConfigurableIdentityProviderAllOfOnBoarding2FA) GetMfaProviderIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MfaProviderId, true
}

// SetMfaProviderId sets field value
func (o *ConfigurableIdentityProviderAllOfOnBoarding2FA) SetMfaProviderId(v string) {
	o.MfaProviderId = v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ConfigurableIdentityProviderAllOfOnBoarding2FA) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurableIdentityProviderAllOfOnBoarding2FA) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ConfigurableIdentityProviderAllOfOnBoarding2FA) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ConfigurableIdentityProviderAllOfOnBoarding2FA) SetMessage(v string) {
	o.Message = &v
}

// GetClaimSuffix returns the ClaimSuffix field value if set, zero value otherwise.
func (o *ConfigurableIdentityProviderAllOfOnBoarding2FA) GetClaimSuffix() string {
	if o == nil || o.ClaimSuffix == nil {
		var ret string
		return ret
	}
	return *o.ClaimSuffix
}

// GetClaimSuffixOk returns a tuple with the ClaimSuffix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurableIdentityProviderAllOfOnBoarding2FA) GetClaimSuffixOk() (*string, bool) {
	if o == nil || o.ClaimSuffix == nil {
		return nil, false
	}
	return o.ClaimSuffix, true
}

// HasClaimSuffix returns a boolean if a field has been set.
func (o *ConfigurableIdentityProviderAllOfOnBoarding2FA) HasClaimSuffix() bool {
	if o != nil && o.ClaimSuffix != nil {
		return true
	}

	return false
}

// SetClaimSuffix gets a reference to the given string and assigns it to the ClaimSuffix field.
func (o *ConfigurableIdentityProviderAllOfOnBoarding2FA) SetClaimSuffix(v string) {
	o.ClaimSuffix = &v
}

// GetAlwaysRequired returns the AlwaysRequired field value if set, zero value otherwise.
func (o *ConfigurableIdentityProviderAllOfOnBoarding2FA) GetAlwaysRequired() bool {
	if o == nil || o.AlwaysRequired == nil {
		var ret bool
		return ret
	}
	return *o.AlwaysRequired
}

// GetAlwaysRequiredOk returns a tuple with the AlwaysRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurableIdentityProviderAllOfOnBoarding2FA) GetAlwaysRequiredOk() (*bool, bool) {
	if o == nil || o.AlwaysRequired == nil {
		return nil, false
	}
	return o.AlwaysRequired, true
}

// HasAlwaysRequired returns a boolean if a field has been set.
func (o *ConfigurableIdentityProviderAllOfOnBoarding2FA) HasAlwaysRequired() bool {
	if o != nil && o.AlwaysRequired != nil {
		return true
	}

	return false
}

// SetAlwaysRequired gets a reference to the given bool and assigns it to the AlwaysRequired field.
func (o *ConfigurableIdentityProviderAllOfOnBoarding2FA) SetAlwaysRequired(v bool) {
	o.AlwaysRequired = &v
}

func (o ConfigurableIdentityProviderAllOfOnBoarding2FA) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["mfaProviderId"] = o.MfaProviderId
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.ClaimSuffix != nil {
		toSerialize["claimSuffix"] = o.ClaimSuffix
	}
	if o.AlwaysRequired != nil {
		toSerialize["alwaysRequired"] = o.AlwaysRequired
	}
	return json.Marshal(toSerialize)
}

type NullableConfigurableIdentityProviderAllOfOnBoarding2FA struct {
	value *ConfigurableIdentityProviderAllOfOnBoarding2FA
	isSet bool
}

func (v NullableConfigurableIdentityProviderAllOfOnBoarding2FA) Get() *ConfigurableIdentityProviderAllOfOnBoarding2FA {
	return v.value
}

func (v *NullableConfigurableIdentityProviderAllOfOnBoarding2FA) Set(val *ConfigurableIdentityProviderAllOfOnBoarding2FA) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigurableIdentityProviderAllOfOnBoarding2FA) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigurableIdentityProviderAllOfOnBoarding2FA) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigurableIdentityProviderAllOfOnBoarding2FA(val *ConfigurableIdentityProviderAllOfOnBoarding2FA) *NullableConfigurableIdentityProviderAllOfOnBoarding2FA {
	return &NullableConfigurableIdentityProviderAllOfOnBoarding2FA{value: val, isSet: true}
}

func (v NullableConfigurableIdentityProviderAllOfOnBoarding2FA) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigurableIdentityProviderAllOfOnBoarding2FA) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
