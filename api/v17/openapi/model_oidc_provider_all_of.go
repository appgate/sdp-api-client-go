/*
Appgate SDP Controller REST API

# About   This specification documents the REST API calls for the Appgate SDP Controller.    Please refer to the REST API chapter in the manual or contact Appgate support with any questions about   this functionality. # Getting Started   Requirements for API scripting:   - Access to the Admin/API TLS Connection (default port 8443) of a Controller appliance.     (https://sdphelp.appgate.com/adminguide/appliance-function-configure.html?anchor=admin-api)   - An API user with relevant permissions.     (https://sdphelp.appgate.com/adminguide/administrative-roles-configure.html)   - In order to use the simple login API, Admin MFA must be disabled or the API user must be excluded.     (https://sdphelp.appgate.com/adminguide/mfa-for-admins.html) # Base path   HTTPS requests must be sent to the Admin Interface hostname and port, with **_/admin** path.    For example: **https://appgate.company.com:8443/admin**    All requests must have the **Accept** header as:    **application/vnd.appgate.peer-v17+json** # API Conventions   API conventions are  important to understand and follow strictly.    - While updating objects (via PUT), entire object must be sent with all fields.     - For example, in order to add a remedy method to the condition below:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": []       }       ```     - send the entire object with updated and non-updated fields:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": [{\"type\": \"DisplayMessage\", \"message\": \"test message\"}]       }       ```    - In case Controller returns an error (non-2xx HTTP status code), response body is JSON.     The \"message\" field contains information about the error.     HTTP 422 \"Unprocessable Entity\" has extra `errors` field to list all the issues with specific fields.    - Empty string (\"\") is considered a different value than \"null\" or field being omitted from JSON.     Omitting the field is recommended if no value is intended.     Empty string (\"\") will be almost always rejected as invalid value.    - There are common pattern between many objects:     - **Configuration Objects**: There are many objects with common fields, namely \"id\", \"name\", \"notes\", \"created\"       and \"updated\". These entities are listed, queried, created, updated and deleted in a similar fashion.     - **Distinguished Name**: Users and Devices are identified with what is called Distinguished Names, as used in        LDAP. The distinguished format that identifies a device and a user combination is        \"CN=\\<Device ID\\>,CN=\\<username\\>,OU=\\<Identity Provider Name\\>\". Some objects have the        \"userDistinguishedName\" field, which does not include the CN for Device ID.        This identifies a user on every device.

API version: API version 17.4
Contact: appgatesdp.support@appgate.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// OidcProviderAllOf Represents an OIDC Identity Provider.
type OidcProviderAllOf struct {
	// OIDC issuer URL.
	Issuer string `json:"issuer"`
	// Audience/Client ID to make sure the recipient of the Token is this Controller.
	Audience string `json:"audience"`
	// Scope to use for tokens.
	Scope *string `json:"scope,omitempty"`
}

// NewOidcProviderAllOf instantiates a new OidcProviderAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOidcProviderAllOf(issuer string, audience string) *OidcProviderAllOf {
	this := OidcProviderAllOf{}
	this.Issuer = issuer
	this.Audience = audience
	var scope string = "openid profile email offline_access"
	this.Scope = &scope
	return &this
}

// NewOidcProviderAllOfWithDefaults instantiates a new OidcProviderAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOidcProviderAllOfWithDefaults() *OidcProviderAllOf {
	this := OidcProviderAllOf{}
	var scope string = "openid profile email offline_access"
	this.Scope = &scope
	return &this
}

// GetIssuer returns the Issuer field value
func (o *OidcProviderAllOf) GetIssuer() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Issuer
}

// GetIssuerOk returns a tuple with the Issuer field value
// and a boolean to check if the value has been set.
func (o *OidcProviderAllOf) GetIssuerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Issuer, true
}

// SetIssuer sets field value
func (o *OidcProviderAllOf) SetIssuer(v string) {
	o.Issuer = v
}

// GetAudience returns the Audience field value
func (o *OidcProviderAllOf) GetAudience() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Audience
}

// GetAudienceOk returns a tuple with the Audience field value
// and a boolean to check if the value has been set.
func (o *OidcProviderAllOf) GetAudienceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Audience, true
}

// SetAudience sets field value
func (o *OidcProviderAllOf) SetAudience(v string) {
	o.Audience = v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *OidcProviderAllOf) GetScope() string {
	if o == nil || o.Scope == nil {
		var ret string
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OidcProviderAllOf) GetScopeOk() (*string, bool) {
	if o == nil || o.Scope == nil {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *OidcProviderAllOf) HasScope() bool {
	if o != nil && o.Scope != nil {
		return true
	}

	return false
}

// SetScope gets a reference to the given string and assigns it to the Scope field.
func (o *OidcProviderAllOf) SetScope(v string) {
	o.Scope = &v
}

func (o OidcProviderAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["issuer"] = o.Issuer
	}
	if true {
		toSerialize["audience"] = o.Audience
	}
	if o.Scope != nil {
		toSerialize["scope"] = o.Scope
	}
	return json.Marshal(toSerialize)
}

type NullableOidcProviderAllOf struct {
	value *OidcProviderAllOf
	isSet bool
}

func (v NullableOidcProviderAllOf) Get() *OidcProviderAllOf {
	return v.value
}

func (v *NullableOidcProviderAllOf) Set(val *OidcProviderAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableOidcProviderAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableOidcProviderAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOidcProviderAllOf(val *OidcProviderAllOf) *NullableOidcProviderAllOf {
	return &NullableOidcProviderAllOf{value: val, isSet: true}
}

func (v NullableOidcProviderAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOidcProviderAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
