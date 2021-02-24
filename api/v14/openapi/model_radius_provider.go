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

// RadiusProvider struct for RadiusProvider
type RadiusProvider struct {
	IdentityProvider
	// Hostnames/IP addresses to connect.
	Hostnames []string `json:"hostnames"`
	// Port to connect.
	Port *int32 `json:"port,omitempty"`
	// Radius shared secret to authenticate to the server.
	SharedSecret string `json:"sharedSecret"`
	// Radius protocol to use while authenticating users.
	AuthenticationProtocol *string `json:"authenticationProtocol,omitempty"`
}

// NewRadiusProvider instantiates a new RadiusProvider object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRadiusProvider(hostnames []string, sharedSecret string) *RadiusProvider {
	this := RadiusProvider{}
	this.Hostnames = hostnames
	var port int32 = 1812
	this.Port = &port
	this.SharedSecret = sharedSecret
	var authenticationProtocol string = "CHAP"
	this.AuthenticationProtocol = &authenticationProtocol
	return &this
}

// NewRadiusProviderWithDefaults instantiates a new RadiusProvider object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRadiusProviderWithDefaults() *RadiusProvider {
	this := RadiusProvider{}
	var port int32 = 1812
	this.Port = &port
	var authenticationProtocol string = "CHAP"
	this.AuthenticationProtocol = &authenticationProtocol
	return &this
}

// GetHostnames returns the Hostnames field value
func (o *RadiusProvider) GetHostnames() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Hostnames
}

// GetHostnamesOk returns a tuple with the Hostnames field value
// and a boolean to check if the value has been set.
func (o *RadiusProvider) GetHostnamesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hostnames, true
}

// SetHostnames sets field value
func (o *RadiusProvider) SetHostnames(v []string) {
	o.Hostnames = v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *RadiusProvider) GetPort() int32 {
	if o == nil || o.Port == nil {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RadiusProvider) GetPortOk() (*int32, bool) {
	if o == nil || o.Port == nil {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *RadiusProvider) HasPort() bool {
	if o != nil && o.Port != nil {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *RadiusProvider) SetPort(v int32) {
	o.Port = &v
}

// GetSharedSecret returns the SharedSecret field value
func (o *RadiusProvider) GetSharedSecret() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SharedSecret
}

// GetSharedSecretOk returns a tuple with the SharedSecret field value
// and a boolean to check if the value has been set.
func (o *RadiusProvider) GetSharedSecretOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SharedSecret, true
}

// SetSharedSecret sets field value
func (o *RadiusProvider) SetSharedSecret(v string) {
	o.SharedSecret = v
}

// GetAuthenticationProtocol returns the AuthenticationProtocol field value if set, zero value otherwise.
func (o *RadiusProvider) GetAuthenticationProtocol() string {
	if o == nil || o.AuthenticationProtocol == nil {
		var ret string
		return ret
	}
	return *o.AuthenticationProtocol
}

// GetAuthenticationProtocolOk returns a tuple with the AuthenticationProtocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RadiusProvider) GetAuthenticationProtocolOk() (*string, bool) {
	if o == nil || o.AuthenticationProtocol == nil {
		return nil, false
	}
	return o.AuthenticationProtocol, true
}

// HasAuthenticationProtocol returns a boolean if a field has been set.
func (o *RadiusProvider) HasAuthenticationProtocol() bool {
	if o != nil && o.AuthenticationProtocol != nil {
		return true
	}

	return false
}

// SetAuthenticationProtocol gets a reference to the given string and assigns it to the AuthenticationProtocol field.
func (o *RadiusProvider) SetAuthenticationProtocol(v string) {
	o.AuthenticationProtocol = &v
}

func (o RadiusProvider) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedIdentityProvider, errIdentityProvider := json.Marshal(o.IdentityProvider)
	if errIdentityProvider != nil {
		return []byte{}, errIdentityProvider
	}
	errIdentityProvider = json.Unmarshal([]byte(serializedIdentityProvider), &toSerialize)
	if errIdentityProvider != nil {
		return []byte{}, errIdentityProvider
	}
	if true {
		toSerialize["hostnames"] = o.Hostnames
	}
	if o.Port != nil {
		toSerialize["port"] = o.Port
	}
	if true {
		toSerialize["sharedSecret"] = o.SharedSecret
	}
	if o.AuthenticationProtocol != nil {
		toSerialize["authenticationProtocol"] = o.AuthenticationProtocol
	}
	return json.Marshal(toSerialize)
}

type NullableRadiusProvider struct {
	value *RadiusProvider
	isSet bool
}

func (v NullableRadiusProvider) Get() *RadiusProvider {
	return v.value
}

func (v *NullableRadiusProvider) Set(val *RadiusProvider) {
	v.value = val
	v.isSet = true
}

func (v NullableRadiusProvider) IsSet() bool {
	return v.isSet
}

func (v *NullableRadiusProvider) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRadiusProvider(val *RadiusProvider) *NullableRadiusProvider {
	return &NullableRadiusProvider{value: val, isSet: true}
}

func (v NullableRadiusProvider) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRadiusProvider) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
