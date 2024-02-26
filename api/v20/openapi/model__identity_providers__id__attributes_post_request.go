/*
Appgate SDP Controller REST API

# About   This specification documents the REST API calls for the Appgate SDP Controller.    Please refer to the REST API chapter in the manual or contact Appgate support with any questions about   this functionality. # Getting Started   Requirements for API scripting:   - Access to the Admin/API TLS Connection (default port 8443) of a Controller appliance.     (https://sdphelp.appgate.com/adminguide/appliance-function-configure.html?anchor=admin-api)   - An API user with relevant permissions.     (https://sdphelp.appgate.com/adminguide/administrative-roles-configure.html)   - In order to use the simple login API, Admin MFA must be disabled or the API user must be excluded.     (https://sdphelp.appgate.com/adminguide/mfa-for-admins.html) # Base path   HTTPS requests must be sent to the Admin Interface hostname and port, with **_/admin** path.    For example: **https://appgate.company.com:8443/admin**    All requests must have the **Accept** header as:    **application/vnd.appgate.peer-v20+json**    An exception is made for the **_/admin/version** endpoint which instead expects an **application/json** Accept header. # API Conventions   API conventions are  important to understand and follow strictly.    - While updating objects (via PUT), entire object must be sent with all fields.     - For example, in order to add a remedy method to the condition below:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": []       }       ```     - send the entire object with updated and non-updated fields:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": [{\"type\": \"DisplayMessage\", \"message\": \"test message\"}]       }       ```    - In case Controller returns an error (non-2xx HTTP status code), response body is JSON.     The \"message\" field contains information about the error.     HTTP 422 \"Unprocessable Entity\" has extra `errors` field to list all the issues with specific fields.    - Empty string (\"\") is considered a different value than \"null\" or field being omitted from JSON.     Omitting the field is recommended if no value is intended.     Empty string (\"\") will be almost always rejected as invalid value.    - There are common pattern between many objects:     - **Configuration Objects**: There are many objects with common fields, namely \"id\", \"name\", \"notes\", \"created\"       and \"updated\". These entities are listed, queried, created, updated and deleted in a similar fashion.     - **Distinguished Name**: Users and Devices are identified with what is called Distinguished Names, as used in        LDAP. The distinguished format that identifies a device and a user combination is        \"CN=\\<Device ID\\>,CN=\\<username\\>,OU=\\<Identity Provider Name\\>\". Some objects have the        \"userDistinguishedName\" field, which does not include the CN for Device ID.        This identifies a user on every device.

API version: API version 20.0
Contact: appgatesdp.support@appgate.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// IdentityProvidersIdAttributesPostRequest struct for IdentityProvidersIdAttributesPostRequest
type IdentityProvidersIdAttributesPostRequest struct {
	// Required for Ldap, Radius and LocalDatabase providers.
	Username *string `json:"username,omitempty"`
	// Required for Radius provider.
	Password *string `json:"password,omitempty"`
	// A sample SAML token to extract attributes from. Required for SAML provider.
	SamlResponse *string `json:"samlResponse,omitempty"`
}

// NewIdentityProvidersIdAttributesPostRequest instantiates a new IdentityProvidersIdAttributesPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityProvidersIdAttributesPostRequest() *IdentityProvidersIdAttributesPostRequest {
	this := IdentityProvidersIdAttributesPostRequest{}
	return &this
}

// NewIdentityProvidersIdAttributesPostRequestWithDefaults instantiates a new IdentityProvidersIdAttributesPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityProvidersIdAttributesPostRequestWithDefaults() *IdentityProvidersIdAttributesPostRequest {
	this := IdentityProvidersIdAttributesPostRequest{}
	return &this
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *IdentityProvidersIdAttributesPostRequest) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProvidersIdAttributesPostRequest) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *IdentityProvidersIdAttributesPostRequest) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *IdentityProvidersIdAttributesPostRequest) SetUsername(v string) {
	o.Username = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *IdentityProvidersIdAttributesPostRequest) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProvidersIdAttributesPostRequest) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *IdentityProvidersIdAttributesPostRequest) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *IdentityProvidersIdAttributesPostRequest) SetPassword(v string) {
	o.Password = &v
}

// GetSamlResponse returns the SamlResponse field value if set, zero value otherwise.
func (o *IdentityProvidersIdAttributesPostRequest) GetSamlResponse() string {
	if o == nil || o.SamlResponse == nil {
		var ret string
		return ret
	}
	return *o.SamlResponse
}

// GetSamlResponseOk returns a tuple with the SamlResponse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProvidersIdAttributesPostRequest) GetSamlResponseOk() (*string, bool) {
	if o == nil || o.SamlResponse == nil {
		return nil, false
	}
	return o.SamlResponse, true
}

// HasSamlResponse returns a boolean if a field has been set.
func (o *IdentityProvidersIdAttributesPostRequest) HasSamlResponse() bool {
	if o != nil && o.SamlResponse != nil {
		return true
	}

	return false
}

// SetSamlResponse gets a reference to the given string and assigns it to the SamlResponse field.
func (o *IdentityProvidersIdAttributesPostRequest) SetSamlResponse(v string) {
	o.SamlResponse = &v
}

func (o IdentityProvidersIdAttributesPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Username != nil {
		toSerialize["username"] = o.Username
	}
	if o.Password != nil {
		toSerialize["password"] = o.Password
	}
	if o.SamlResponse != nil {
		toSerialize["samlResponse"] = o.SamlResponse
	}
	return json.Marshal(toSerialize)
}

type NullableIdentityProvidersIdAttributesPostRequest struct {
	value *IdentityProvidersIdAttributesPostRequest
	isSet bool
}

func (v NullableIdentityProvidersIdAttributesPostRequest) Get() *IdentityProvidersIdAttributesPostRequest {
	return v.value
}

func (v *NullableIdentityProvidersIdAttributesPostRequest) Set(val *IdentityProvidersIdAttributesPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityProvidersIdAttributesPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityProvidersIdAttributesPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityProvidersIdAttributesPostRequest(val *IdentityProvidersIdAttributesPostRequest) *NullableIdentityProvidersIdAttributesPostRequest {
	return &NullableIdentityProvidersIdAttributesPostRequest{value: val, isSet: true}
}

func (v NullableIdentityProvidersIdAttributesPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityProvidersIdAttributesPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
