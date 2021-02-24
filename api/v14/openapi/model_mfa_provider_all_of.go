/*
 * Appgate SDP Controller REST API
 *
 * # About   This specification documents the REST API calls for the Appgate SDP Controller.    Please refer to the Integration chapter in the manual or contact Appgate support with any questions about   this functionality. # Getting Started   Requirements for API scripting:   - Access to the Admin Interface (default port 8443) of a Controller appliance.     (https://sdphelp.appgate.com/adminguide/appliances-admin-interface-configure.html)   - An API user with relevant permissions.     (https://sdphelp.appgate.com/adminguide/administrative-roles-configure.html)   - In order to use the simple login API, Admin MFA must be disabled or the API user must be excluded.     (https://sdphelp.appgate.com/adminguide/mfa-for-admins.html) # Base path   HTTPS requests must be sent to the Admin Interface hostname and port, with **_/admin** path.    For example: **https://appgate.company.com:8443/admin**    All requests must have the **Accept** header as:    **application/vnd.appgate.peer-v14+json** # API Conventions   API conventions are  important to understand and follow strictly.    - While updating objects (via PUT), entire object must be sent with all fields.     - For example, if in order to add a remedy method to the condition below:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": []       }       ```     - send the entire object with updated and non-updated fields:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": [{\"type\": \"DisplayMessage\", \"message\": \"test message\"}]       }       ```    - In case Controller returns an error (non-2xx HTTP status code), response body is JSON.     The \"message\" field contains information about the error.     HTTP 422 \"Unprocessable Entity\" has extra `errors` field to list all the issues with specific fields.    - Empty string (\"\") is considered a different value than \"null\" or field being omitted from JSON.     Omitting the field is recommend if no value is intended.     Empty string (\"\") will be almost always rejected as invalid value.    - There are common pattern between many objects:     - **Configuration Objects**: There are many objects with common fields, namely \"id\", \"name\", \"notes\", \"created\"       and \"updated\". These entities are listed, queried, created, updated and deleted in a similar fashion.     - **Distinguished Name**: Users and Devices are identified with what is called Distinguished Names, as used in        LDAP. The distinguished format that identifies a device and a user combination is        \"CN=\\<Device ID\\>,CN=\\<username\\>,OU=\\<Identity Provider Name\\>\". Some objects have the        \"userDistinguishedName\" field, which does not include the CN for Device ID.        This identifies a user on every device.
 *
 * API version: API version 14
 * Contact: appgatesdp.support@appgate.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"encoding/json"
)

// MfaProviderAllOf Represents a MFA Provider.
type MfaProviderAllOf struct {
	// The type of the MFA Provider. \"DefaultTimeBased\" is built-in, a new one cannot be created.
	Type string `json:"type"`
	// Hostnames/IP addresses to connect.
	Hostnames []string `json:"hostnames"`
	// Port to connect.
	Port float32 `json:"port"`
	// The input type used in the client to enter the MFA code.   * \"Masked\" - The input is masked the same way as a password field.  * \"Numeric\" - The input is marked as a numeric input.  * \"Text\" - The input is handled as a regular plain text field.
	InputType *string `json:"inputType,omitempty"`
	// Radius shared secret to authenticate to the server.
	SharedSecret *string `json:"sharedSecret,omitempty"`
	// Radius protocol to use while authenticating users.
	AuthenticationProtocol *string `json:"authenticationProtocol,omitempty"`
	// Timeout in seconds before giving up on response.
	Timeout *float32 `json:"timeout,omitempty"`
	// Defines the multi-factor authentication flow for RADIUS.  * \"OneFactor\" - The input from the user is sent as password and the response is used for result.  * \"Challenge\" - Before prompting the user, Controller sends a challenge request to the RADIUS server  using \"challengeSharedSecret\" or the user password. Data from the response is used with user input to  send the second RADIUS authentication request.  * \"Push\" - \"challengeSharedSecret\" or the user password is sent to RADIUS which triggers an external  authentication flow. When the external authentication flow returns success, the MFA attempt is  authenticated.
	Mode *string `json:"mode,omitempty"`
	// -> If enabled, the Client will send the cached password instead of using challengeSharedSecret\" to initiate the multi-factor authentication.
	UseUserPassword *bool `json:"useUserPassword,omitempty"`
	// -> Password sent to RADIUS to initiate multi-factor authentication. Required if \"useUserPassword\" is not enabled.
	ChallengeSharedSecret *string `json:"challengeSharedSecret,omitempty"`
}

// NewMfaProviderAllOf instantiates a new MfaProviderAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMfaProviderAllOf(type_ string, hostnames []string, port float32) *MfaProviderAllOf {
	this := MfaProviderAllOf{}
	this.Type = type_
	this.Hostnames = hostnames
	this.Port = port
	var inputType string = "Masked"
	this.InputType = &inputType
	var authenticationProtocol string = "CHAP"
	this.AuthenticationProtocol = &authenticationProtocol
	var timeout float32 = 6
	this.Timeout = &timeout
	var mode string = "Challenge"
	this.Mode = &mode
	return &this
}

// NewMfaProviderAllOfWithDefaults instantiates a new MfaProviderAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMfaProviderAllOfWithDefaults() *MfaProviderAllOf {
	this := MfaProviderAllOf{}
	var inputType string = "Masked"
	this.InputType = &inputType
	var authenticationProtocol string = "CHAP"
	this.AuthenticationProtocol = &authenticationProtocol
	var timeout float32 = 6
	this.Timeout = &timeout
	var mode string = "Challenge"
	this.Mode = &mode
	return &this
}

// GetType returns the Type field value
func (o *MfaProviderAllOf) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *MfaProviderAllOf) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *MfaProviderAllOf) SetType(v string) {
	o.Type = v
}

// GetHostnames returns the Hostnames field value
func (o *MfaProviderAllOf) GetHostnames() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Hostnames
}

// GetHostnamesOk returns a tuple with the Hostnames field value
// and a boolean to check if the value has been set.
func (o *MfaProviderAllOf) GetHostnamesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hostnames, true
}

// SetHostnames sets field value
func (o *MfaProviderAllOf) SetHostnames(v []string) {
	o.Hostnames = v
}

// GetPort returns the Port field value
func (o *MfaProviderAllOf) GetPort() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Port
}

// GetPortOk returns a tuple with the Port field value
// and a boolean to check if the value has been set.
func (o *MfaProviderAllOf) GetPortOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Port, true
}

// SetPort sets field value
func (o *MfaProviderAllOf) SetPort(v float32) {
	o.Port = v
}

// GetInputType returns the InputType field value if set, zero value otherwise.
func (o *MfaProviderAllOf) GetInputType() string {
	if o == nil || o.InputType == nil {
		var ret string
		return ret
	}
	return *o.InputType
}

// GetInputTypeOk returns a tuple with the InputType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MfaProviderAllOf) GetInputTypeOk() (*string, bool) {
	if o == nil || o.InputType == nil {
		return nil, false
	}
	return o.InputType, true
}

// HasInputType returns a boolean if a field has been set.
func (o *MfaProviderAllOf) HasInputType() bool {
	if o != nil && o.InputType != nil {
		return true
	}

	return false
}

// SetInputType gets a reference to the given string and assigns it to the InputType field.
func (o *MfaProviderAllOf) SetInputType(v string) {
	o.InputType = &v
}

// GetSharedSecret returns the SharedSecret field value if set, zero value otherwise.
func (o *MfaProviderAllOf) GetSharedSecret() string {
	if o == nil || o.SharedSecret == nil {
		var ret string
		return ret
	}
	return *o.SharedSecret
}

// GetSharedSecretOk returns a tuple with the SharedSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MfaProviderAllOf) GetSharedSecretOk() (*string, bool) {
	if o == nil || o.SharedSecret == nil {
		return nil, false
	}
	return o.SharedSecret, true
}

// HasSharedSecret returns a boolean if a field has been set.
func (o *MfaProviderAllOf) HasSharedSecret() bool {
	if o != nil && o.SharedSecret != nil {
		return true
	}

	return false
}

// SetSharedSecret gets a reference to the given string and assigns it to the SharedSecret field.
func (o *MfaProviderAllOf) SetSharedSecret(v string) {
	o.SharedSecret = &v
}

// GetAuthenticationProtocol returns the AuthenticationProtocol field value if set, zero value otherwise.
func (o *MfaProviderAllOf) GetAuthenticationProtocol() string {
	if o == nil || o.AuthenticationProtocol == nil {
		var ret string
		return ret
	}
	return *o.AuthenticationProtocol
}

// GetAuthenticationProtocolOk returns a tuple with the AuthenticationProtocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MfaProviderAllOf) GetAuthenticationProtocolOk() (*string, bool) {
	if o == nil || o.AuthenticationProtocol == nil {
		return nil, false
	}
	return o.AuthenticationProtocol, true
}

// HasAuthenticationProtocol returns a boolean if a field has been set.
func (o *MfaProviderAllOf) HasAuthenticationProtocol() bool {
	if o != nil && o.AuthenticationProtocol != nil {
		return true
	}

	return false
}

// SetAuthenticationProtocol gets a reference to the given string and assigns it to the AuthenticationProtocol field.
func (o *MfaProviderAllOf) SetAuthenticationProtocol(v string) {
	o.AuthenticationProtocol = &v
}

// GetTimeout returns the Timeout field value if set, zero value otherwise.
func (o *MfaProviderAllOf) GetTimeout() float32 {
	if o == nil || o.Timeout == nil {
		var ret float32
		return ret
	}
	return *o.Timeout
}

// GetTimeoutOk returns a tuple with the Timeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MfaProviderAllOf) GetTimeoutOk() (*float32, bool) {
	if o == nil || o.Timeout == nil {
		return nil, false
	}
	return o.Timeout, true
}

// HasTimeout returns a boolean if a field has been set.
func (o *MfaProviderAllOf) HasTimeout() bool {
	if o != nil && o.Timeout != nil {
		return true
	}

	return false
}

// SetTimeout gets a reference to the given float32 and assigns it to the Timeout field.
func (o *MfaProviderAllOf) SetTimeout(v float32) {
	o.Timeout = &v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *MfaProviderAllOf) GetMode() string {
	if o == nil || o.Mode == nil {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MfaProviderAllOf) GetModeOk() (*string, bool) {
	if o == nil || o.Mode == nil {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *MfaProviderAllOf) HasMode() bool {
	if o != nil && o.Mode != nil {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *MfaProviderAllOf) SetMode(v string) {
	o.Mode = &v
}

// GetUseUserPassword returns the UseUserPassword field value if set, zero value otherwise.
func (o *MfaProviderAllOf) GetUseUserPassword() bool {
	if o == nil || o.UseUserPassword == nil {
		var ret bool
		return ret
	}
	return *o.UseUserPassword
}

// GetUseUserPasswordOk returns a tuple with the UseUserPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MfaProviderAllOf) GetUseUserPasswordOk() (*bool, bool) {
	if o == nil || o.UseUserPassword == nil {
		return nil, false
	}
	return o.UseUserPassword, true
}

// HasUseUserPassword returns a boolean if a field has been set.
func (o *MfaProviderAllOf) HasUseUserPassword() bool {
	if o != nil && o.UseUserPassword != nil {
		return true
	}

	return false
}

// SetUseUserPassword gets a reference to the given bool and assigns it to the UseUserPassword field.
func (o *MfaProviderAllOf) SetUseUserPassword(v bool) {
	o.UseUserPassword = &v
}

// GetChallengeSharedSecret returns the ChallengeSharedSecret field value if set, zero value otherwise.
func (o *MfaProviderAllOf) GetChallengeSharedSecret() string {
	if o == nil || o.ChallengeSharedSecret == nil {
		var ret string
		return ret
	}
	return *o.ChallengeSharedSecret
}

// GetChallengeSharedSecretOk returns a tuple with the ChallengeSharedSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MfaProviderAllOf) GetChallengeSharedSecretOk() (*string, bool) {
	if o == nil || o.ChallengeSharedSecret == nil {
		return nil, false
	}
	return o.ChallengeSharedSecret, true
}

// HasChallengeSharedSecret returns a boolean if a field has been set.
func (o *MfaProviderAllOf) HasChallengeSharedSecret() bool {
	if o != nil && o.ChallengeSharedSecret != nil {
		return true
	}

	return false
}

// SetChallengeSharedSecret gets a reference to the given string and assigns it to the ChallengeSharedSecret field.
func (o *MfaProviderAllOf) SetChallengeSharedSecret(v string) {
	o.ChallengeSharedSecret = &v
}

func (o MfaProviderAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["hostnames"] = o.Hostnames
	}
	if true {
		toSerialize["port"] = o.Port
	}
	if o.InputType != nil {
		toSerialize["inputType"] = o.InputType
	}
	if o.SharedSecret != nil {
		toSerialize["sharedSecret"] = o.SharedSecret
	}
	if o.AuthenticationProtocol != nil {
		toSerialize["authenticationProtocol"] = o.AuthenticationProtocol
	}
	if o.Timeout != nil {
		toSerialize["timeout"] = o.Timeout
	}
	if o.Mode != nil {
		toSerialize["mode"] = o.Mode
	}
	if o.UseUserPassword != nil {
		toSerialize["useUserPassword"] = o.UseUserPassword
	}
	if o.ChallengeSharedSecret != nil {
		toSerialize["challengeSharedSecret"] = o.ChallengeSharedSecret
	}
	return json.Marshal(toSerialize)
}

type NullableMfaProviderAllOf struct {
	value *MfaProviderAllOf
	isSet bool
}

func (v NullableMfaProviderAllOf) Get() *MfaProviderAllOf {
	return v.value
}

func (v *NullableMfaProviderAllOf) Set(val *MfaProviderAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableMfaProviderAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableMfaProviderAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMfaProviderAllOf(val *MfaProviderAllOf) *NullableMfaProviderAllOf {
	return &NullableMfaProviderAllOf{value: val, isSet: true}
}

func (v NullableMfaProviderAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMfaProviderAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
