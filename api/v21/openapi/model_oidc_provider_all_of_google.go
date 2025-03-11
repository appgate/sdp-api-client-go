/*
Appgate SDP Controller REST API

# About   This specification documents the REST API calls for the Appgate SDP Controller.    Please refer to the REST API chapter in the manual or contact Appgate support with any questions about   this functionality. # Getting Started   Requirements for API scripting:   - Access to the Admin/API TLS Connection (default port 8443) of a Controller appliance.     (https://sdphelp.appgate.com/adminguide/appliance-function-configure.html?anchor=admin-api)   - An API user with relevant permissions.     (https://sdphelp.appgate.com/adminguide/administrative-roles-configure.html)   - In order to use the simple login API, Admin MFA must be disabled or the API user must be excluded.     (https://sdphelp.appgate.com/adminguide/mfa-for-admins.html) # Base path   HTTPS requests must be sent to the Admin Interface hostname and port, with **_/admin** path.    For example: **https://appgate.company.com:8443/admin**    All requests must have the **Accept** header as:    **application/vnd.appgate.peer-v21+json**    An exception is made for the **_/admin/version** endpoint which instead expects an **application/json** Accept header. # API Conventions   API conventions are  important to understand and follow strictly.    - While updating objects (via PUT), entire object must be sent with all fields.     - For example, in order to add a remedy method to the condition below:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": []       }       ```     - send the entire object with updated and non-updated fields:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": [{\"type\": \"DisplayMessage\", \"message\": \"test message\"}]       }       ```    - In case Controller returns an error (non-2xx HTTP status code), response body is JSON.     The \"message\" field contains information about the error.     HTTP 422 \"Unprocessable Entity\" has extra `errors` field to list all the issues with specific fields.    - Empty string (\"\") is considered a different value than \"null\" or field being omitted from JSON.     Omitting the field is recommended if no value is intended.     Empty string (\"\") will be almost always rejected as invalid value.    - There are common pattern between many objects:     - **Configuration Objects**: There are many objects with common fields, namely \"id\", \"name\", \"notes\", \"created\"       and \"updated\". These entities are listed, queried, created, updated and deleted in a similar fashion.     - **Distinguished Name**: Users and Devices are identified with what is called Distinguished Names, as used in        LDAP. The distinguished format that identifies a device and a user combination is        \"CN=\\<Device ID\\>,CN=\\<username\\>,OU=\\<Identity Provider Name\\>\". Some objects have the        \"userDistinguishedName\" field, which does not include the CN for Device ID.        This identifies a user on every device.

API version: API version 21.1
Contact: appgatesdp.support@appgate.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// OidcProviderAllOfGoogle Google specific OIDC settings.
type OidcProviderAllOfGoogle struct {
	// Whether to enable Google OIDC settings or not.
	Enabled *bool `json:"enabled,omitempty"`
	// Client secret.
	ClientSecret *string `json:"clientSecret,omitempty"`
	// Whether to enable refresh token with Google OIDC or not.
	RefreshToken *bool `json:"refreshToken,omitempty"`
}

// NewOidcProviderAllOfGoogle instantiates a new OidcProviderAllOfGoogle object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOidcProviderAllOfGoogle() *OidcProviderAllOfGoogle {
	this := OidcProviderAllOfGoogle{}
	var enabled bool = false
	this.Enabled = &enabled
	return &this
}

// NewOidcProviderAllOfGoogleWithDefaults instantiates a new OidcProviderAllOfGoogle object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOidcProviderAllOfGoogleWithDefaults() *OidcProviderAllOfGoogle {
	this := OidcProviderAllOfGoogle{}
	var enabled bool = false
	this.Enabled = &enabled
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *OidcProviderAllOfGoogle) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OidcProviderAllOfGoogle) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *OidcProviderAllOfGoogle) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *OidcProviderAllOfGoogle) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetClientSecret returns the ClientSecret field value if set, zero value otherwise.
func (o *OidcProviderAllOfGoogle) GetClientSecret() string {
	if o == nil || o.ClientSecret == nil {
		var ret string
		return ret
	}
	return *o.ClientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OidcProviderAllOfGoogle) GetClientSecretOk() (*string, bool) {
	if o == nil || o.ClientSecret == nil {
		return nil, false
	}
	return o.ClientSecret, true
}

// HasClientSecret returns a boolean if a field has been set.
func (o *OidcProviderAllOfGoogle) HasClientSecret() bool {
	if o != nil && o.ClientSecret != nil {
		return true
	}

	return false
}

// SetClientSecret gets a reference to the given string and assigns it to the ClientSecret field.
func (o *OidcProviderAllOfGoogle) SetClientSecret(v string) {
	o.ClientSecret = &v
}

// GetRefreshToken returns the RefreshToken field value if set, zero value otherwise.
func (o *OidcProviderAllOfGoogle) GetRefreshToken() bool {
	if o == nil || o.RefreshToken == nil {
		var ret bool
		return ret
	}
	return *o.RefreshToken
}

// GetRefreshTokenOk returns a tuple with the RefreshToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OidcProviderAllOfGoogle) GetRefreshTokenOk() (*bool, bool) {
	if o == nil || o.RefreshToken == nil {
		return nil, false
	}
	return o.RefreshToken, true
}

// HasRefreshToken returns a boolean if a field has been set.
func (o *OidcProviderAllOfGoogle) HasRefreshToken() bool {
	if o != nil && o.RefreshToken != nil {
		return true
	}

	return false
}

// SetRefreshToken gets a reference to the given bool and assigns it to the RefreshToken field.
func (o *OidcProviderAllOfGoogle) SetRefreshToken(v bool) {
	o.RefreshToken = &v
}

func (o OidcProviderAllOfGoogle) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.ClientSecret != nil {
		toSerialize["clientSecret"] = o.ClientSecret
	}
	if o.RefreshToken != nil {
		toSerialize["refreshToken"] = o.RefreshToken
	}
	return json.Marshal(toSerialize)
}

type NullableOidcProviderAllOfGoogle struct {
	value *OidcProviderAllOfGoogle
	isSet bool
}

func (v NullableOidcProviderAllOfGoogle) Get() *OidcProviderAllOfGoogle {
	return v.value
}

func (v *NullableOidcProviderAllOfGoogle) Set(val *OidcProviderAllOfGoogle) {
	v.value = val
	v.isSet = true
}

func (v NullableOidcProviderAllOfGoogle) IsSet() bool {
	return v.isSet
}

func (v *NullableOidcProviderAllOfGoogle) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOidcProviderAllOfGoogle(val *OidcProviderAllOfGoogle) *NullableOidcProviderAllOfGoogle {
	return &NullableOidcProviderAllOfGoogle{value: val, isSet: true}
}

func (v NullableOidcProviderAllOfGoogle) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOidcProviderAllOfGoogle) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
