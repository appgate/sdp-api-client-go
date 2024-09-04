/*
Appgate SDP Controller REST API

# About   This specification documents the REST API calls for the Appgate SDP Controller.    Please refer to the REST API chapter in the manual or contact Appgate support with any questions about   this functionality. # Getting Started   Requirements for API scripting:   - Access to the Admin/API TLS Connection (default port 8443) of a Controller appliance.     (https://sdphelp.appgate.com/adminguide/appliance-function-configure.html?anchor=admin-api)   - An API user with relevant permissions.     (https://sdphelp.appgate.com/adminguide/administrative-roles-configure.html)   - In order to use the simple login API, Admin MFA must be disabled or the API user must be excluded.     (https://sdphelp.appgate.com/adminguide/mfa-for-admins.html) # Base path   HTTPS requests must be sent to the Admin Interface hostname and port, with **_/admin** path.    For example: **https://appgate.company.com:8443/admin**    All requests must have the **Accept** header as:    **application/vnd.appgate.peer-v20+json**    An exception is made for the **_/admin/version** endpoint which instead expects an **application/json** Accept header. # API Conventions   API conventions are  important to understand and follow strictly.    - While updating objects (via PUT), entire object must be sent with all fields.     - For example, in order to add a remedy method to the condition below:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": []       }       ```     - send the entire object with updated and non-updated fields:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": [{\"type\": \"DisplayMessage\", \"message\": \"test message\"}]       }       ```    - In case Controller returns an error (non-2xx HTTP status code), response body is JSON.     The \"message\" field contains information about the error.     HTTP 422 \"Unprocessable Entity\" has extra `errors` field to list all the issues with specific fields.    - Empty string (\"\") is considered a different value than \"null\" or field being omitted from JSON.     Omitting the field is recommended if no value is intended.     Empty string (\"\") will be almost always rejected as invalid value.    - There are common pattern between many objects:     - **Configuration Objects**: There are many objects with common fields, namely \"id\", \"name\", \"notes\", \"created\"       and \"updated\". These entities are listed, queried, created, updated and deleted in a similar fashion.     - **Distinguished Name**: Users and Devices are identified with what is called Distinguished Names, as used in        LDAP. The distinguished format that identifies a device and a user combination is        \"CN=\\<Device ID\\>,CN=\\<username\\>,OU=\\<Identity Provider Name\\>\". Some objects have the        \"userDistinguishedName\" field, which does not include the CN for Device ID.        This identifies a user on every device.

API version: API version 20.3
Contact: appgatesdp.support@appgate.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// RadiusProviderAllOf Represents a Radius Identity Provider.
type RadiusProviderAllOf struct {
	// Hostnames/IP addresses to connect.
	Hostnames []string `json:"hostnames"`
	// Port to connect.
	Port *int32 `json:"port,omitempty"`
	// Radius shared secret to authenticate to the server.
	SharedSecret string `json:"sharedSecret"`
	// Radius protocol to use while authenticating users.
	AuthenticationProtocol *string `json:"authenticationProtocol,omitempty"`
}

// NewRadiusProviderAllOf instantiates a new RadiusProviderAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRadiusProviderAllOf(hostnames []string, sharedSecret string) *RadiusProviderAllOf {
	this := RadiusProviderAllOf{}
	this.Hostnames = hostnames
	var port int32 = 1812
	this.Port = &port
	this.SharedSecret = sharedSecret
	var authenticationProtocol string = "CHAP"
	this.AuthenticationProtocol = &authenticationProtocol
	return &this
}

// NewRadiusProviderAllOfWithDefaults instantiates a new RadiusProviderAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRadiusProviderAllOfWithDefaults() *RadiusProviderAllOf {
	this := RadiusProviderAllOf{}
	var port int32 = 1812
	this.Port = &port
	var authenticationProtocol string = "CHAP"
	this.AuthenticationProtocol = &authenticationProtocol
	return &this
}

// GetHostnames returns the Hostnames field value
func (o *RadiusProviderAllOf) GetHostnames() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Hostnames
}

// GetHostnamesOk returns a tuple with the Hostnames field value
// and a boolean to check if the value has been set.
func (o *RadiusProviderAllOf) GetHostnamesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Hostnames, true
}

// SetHostnames sets field value
func (o *RadiusProviderAllOf) SetHostnames(v []string) {
	o.Hostnames = v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *RadiusProviderAllOf) GetPort() int32 {
	if o == nil || o.Port == nil {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RadiusProviderAllOf) GetPortOk() (*int32, bool) {
	if o == nil || o.Port == nil {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *RadiusProviderAllOf) HasPort() bool {
	if o != nil && o.Port != nil {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *RadiusProviderAllOf) SetPort(v int32) {
	o.Port = &v
}

// GetSharedSecret returns the SharedSecret field value
func (o *RadiusProviderAllOf) GetSharedSecret() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SharedSecret
}

// GetSharedSecretOk returns a tuple with the SharedSecret field value
// and a boolean to check if the value has been set.
func (o *RadiusProviderAllOf) GetSharedSecretOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SharedSecret, true
}

// SetSharedSecret sets field value
func (o *RadiusProviderAllOf) SetSharedSecret(v string) {
	o.SharedSecret = v
}

// GetAuthenticationProtocol returns the AuthenticationProtocol field value if set, zero value otherwise.
func (o *RadiusProviderAllOf) GetAuthenticationProtocol() string {
	if o == nil || o.AuthenticationProtocol == nil {
		var ret string
		return ret
	}
	return *o.AuthenticationProtocol
}

// GetAuthenticationProtocolOk returns a tuple with the AuthenticationProtocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RadiusProviderAllOf) GetAuthenticationProtocolOk() (*string, bool) {
	if o == nil || o.AuthenticationProtocol == nil {
		return nil, false
	}
	return o.AuthenticationProtocol, true
}

// HasAuthenticationProtocol returns a boolean if a field has been set.
func (o *RadiusProviderAllOf) HasAuthenticationProtocol() bool {
	if o != nil && o.AuthenticationProtocol != nil {
		return true
	}

	return false
}

// SetAuthenticationProtocol gets a reference to the given string and assigns it to the AuthenticationProtocol field.
func (o *RadiusProviderAllOf) SetAuthenticationProtocol(v string) {
	o.AuthenticationProtocol = &v
}

func (o RadiusProviderAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
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

type NullableRadiusProviderAllOf struct {
	value *RadiusProviderAllOf
	isSet bool
}

func (v NullableRadiusProviderAllOf) Get() *RadiusProviderAllOf {
	return v.value
}

func (v *NullableRadiusProviderAllOf) Set(val *RadiusProviderAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableRadiusProviderAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableRadiusProviderAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRadiusProviderAllOf(val *RadiusProviderAllOf) *NullableRadiusProviderAllOf {
	return &NullableRadiusProviderAllOf{value: val, isSet: true}
}

func (v NullableRadiusProviderAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRadiusProviderAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
