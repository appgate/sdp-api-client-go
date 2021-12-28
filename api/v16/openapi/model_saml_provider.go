/*
Appgate SDP Controller REST API

# About   This specification documents the REST API calls for the Appgate SDP Controller.    Please refer to the REST API chapter in the manual or contact Appgate support with any questions about   this functionality. # Getting Started   Requirements for API scripting:   - Access to the Admin/API TLS Connection (default port 8443) of a Controller appliance.     (https://sdphelp.appgate.com/adminguide/appliance-function-configure.html?anchor=admin-api)   - An API user with relevant permissions.     (https://sdphelp.appgate.com/adminguide/administrative-roles-configure.html)   - In order to use the simple login API, Admin MFA must be disabled or the API user must be excluded.     (https://sdphelp.appgate.com/adminguide/mfa-for-admins.html) # Base path   HTTPS requests must be sent to the Admin Interface hostname and port, with **_/admin** path.    For example: **https://appgate.company.com:8443/admin**    All requests must have the **Accept** header as:    **application/vnd.appgate.peer-v16+json** # API Conventions   API conventions are  important to understand and follow strictly.    - While updating objects (via PUT), entire object must be sent with all fields.     - For example, in order to add a remedy method to the condition below:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": []       }       ```     - send the entire object with updated and non-updated fields:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": [{\"type\": \"DisplayMessage\", \"message\": \"test message\"}]       }       ```    - In case Controller returns an error (non-2xx HTTP status code), response body is JSON.     The \"message\" field contains information about the error.     HTTP 422 \"Unprocessable Entity\" has extra `errors` field to list all the issues with specific fields.    - Empty string (\"\") is considered a different value than \"null\" or field being omitted from JSON.     Omitting the field is recommend if no value is intended.     Empty string (\"\") will be almost always rejected as invalid value.    - There are common pattern between many objects:     - **Configuration Objects**: There are many objects with common fields, namely \"id\", \"name\", \"notes\", \"created\"       and \"updated\". These entities are listed, queried, created, updated and deleted in a similar fashion.     - **Distinguished Name**: Users and Devices are identified with what is called Distinguished Names, as used in        LDAP. The distinguished format that identifies a device and a user combination is        \"CN=\\<Device ID\\>,CN=\\<username\\>,OU=\\<Identity Provider Name\\>\". Some objects have the        \"userDistinguishedName\" field, which does not include the CN for Device ID.        This identifies a user on every device.

API version: API version 16.2
Contact: appgatesdp.support@appgate.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// SamlProvider struct for SamlProvider
type SamlProvider struct {
	IdentityProvider
	// SAML audience to make sure the recipient of the Token is this Controller.
	Audience string `json:"audience"`
	// The private key to decrypt encrypted assertions if there is any. In PEM format.
	DecryptionKey *string `json:"decryptionKey,omitempty"`
	// Enables ForceAuthn flag in the SAML Request. If the SAML Provider supports this flag, it will require user to enter their credentials every time Client requires SAML authentication.
	ForceAuthn *bool `json:"forceAuthn,omitempty"`
	// SAML issuer ID to make sure the sender of the Token is the expected SAML provider.
	Issuer string `json:"issuer"`
	// The certificate of the SAML provider to verify the SAML tokens. In PEM format.
	ProviderCertificate string `json:"providerCertificate"`
	// The URL to redirect the user browsers to authenticate against the SAML Server. Also known as Single Sign-on URL. AuthNRequest will be added automatically.
	RedirectUrl string `json:"redirectUrl"`
}

// NewSamlProvider instantiates a new SamlProvider object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSamlProvider(audience string, issuer string, providerCertificate string, redirectUrl string, id string, name string, type_ string) *SamlProvider {
	this := SamlProvider{}
	this.Id = id
	this.Name = name
	this.Type = type_
	var adminProvider bool = false
	this.AdminProvider = &adminProvider
	var deviceLimitPerUser int32 = 100
	this.DeviceLimitPerUser = &deviceLimitPerUser
	var inactivityTimeoutMinutes int32 = 0
	this.InactivityTimeoutMinutes = &inactivityTimeoutMinutes
	var blockLocalDnsRequests bool = false
	this.BlockLocalDnsRequests = &blockLocalDnsRequests
	this.Audience = audience
	this.Issuer = issuer
	this.ProviderCertificate = providerCertificate
	this.RedirectUrl = redirectUrl
	return &this
}

// NewSamlProviderWithDefaults instantiates a new SamlProvider object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSamlProviderWithDefaults() *SamlProvider {
	this := SamlProvider{}
	return &this
}

// GetAudience returns the Audience field value
func (o *SamlProvider) GetAudience() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Audience
}

// GetAudienceOk returns a tuple with the Audience field value
// and a boolean to check if the value has been set.
func (o *SamlProvider) GetAudienceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Audience, true
}

// SetAudience sets field value
func (o *SamlProvider) SetAudience(v string) {
	o.Audience = v
}

// GetDecryptionKey returns the DecryptionKey field value if set, zero value otherwise.
func (o *SamlProvider) GetDecryptionKey() string {
	if o == nil || o.DecryptionKey == nil {
		var ret string
		return ret
	}
	return *o.DecryptionKey
}

// GetDecryptionKeyOk returns a tuple with the DecryptionKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlProvider) GetDecryptionKeyOk() (*string, bool) {
	if o == nil || o.DecryptionKey == nil {
		return nil, false
	}
	return o.DecryptionKey, true
}

// HasDecryptionKey returns a boolean if a field has been set.
func (o *SamlProvider) HasDecryptionKey() bool {
	if o != nil && o.DecryptionKey != nil {
		return true
	}

	return false
}

// SetDecryptionKey gets a reference to the given string and assigns it to the DecryptionKey field.
func (o *SamlProvider) SetDecryptionKey(v string) {
	o.DecryptionKey = &v
}

// GetForceAuthn returns the ForceAuthn field value if set, zero value otherwise.
func (o *SamlProvider) GetForceAuthn() bool {
	if o == nil || o.ForceAuthn == nil {
		var ret bool
		return ret
	}
	return *o.ForceAuthn
}

// GetForceAuthnOk returns a tuple with the ForceAuthn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlProvider) GetForceAuthnOk() (*bool, bool) {
	if o == nil || o.ForceAuthn == nil {
		return nil, false
	}
	return o.ForceAuthn, true
}

// HasForceAuthn returns a boolean if a field has been set.
func (o *SamlProvider) HasForceAuthn() bool {
	if o != nil && o.ForceAuthn != nil {
		return true
	}

	return false
}

// SetForceAuthn gets a reference to the given bool and assigns it to the ForceAuthn field.
func (o *SamlProvider) SetForceAuthn(v bool) {
	o.ForceAuthn = &v
}

// GetIssuer returns the Issuer field value
func (o *SamlProvider) GetIssuer() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Issuer
}

// GetIssuerOk returns a tuple with the Issuer field value
// and a boolean to check if the value has been set.
func (o *SamlProvider) GetIssuerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Issuer, true
}

// SetIssuer sets field value
func (o *SamlProvider) SetIssuer(v string) {
	o.Issuer = v
}

// GetProviderCertificate returns the ProviderCertificate field value
func (o *SamlProvider) GetProviderCertificate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProviderCertificate
}

// GetProviderCertificateOk returns a tuple with the ProviderCertificate field value
// and a boolean to check if the value has been set.
func (o *SamlProvider) GetProviderCertificateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProviderCertificate, true
}

// SetProviderCertificate sets field value
func (o *SamlProvider) SetProviderCertificate(v string) {
	o.ProviderCertificate = v
}

// GetRedirectUrl returns the RedirectUrl field value
func (o *SamlProvider) GetRedirectUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RedirectUrl
}

// GetRedirectUrlOk returns a tuple with the RedirectUrl field value
// and a boolean to check if the value has been set.
func (o *SamlProvider) GetRedirectUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RedirectUrl, true
}

// SetRedirectUrl sets field value
func (o *SamlProvider) SetRedirectUrl(v string) {
	o.RedirectUrl = v
}

func (o SamlProvider) MarshalJSON() ([]byte, error) {
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
		toSerialize["audience"] = o.Audience
	}
	if o.DecryptionKey != nil {
		toSerialize["decryptionKey"] = o.DecryptionKey
	}
	if o.ForceAuthn != nil {
		toSerialize["forceAuthn"] = o.ForceAuthn
	}
	if true {
		toSerialize["issuer"] = o.Issuer
	}
	if true {
		toSerialize["providerCertificate"] = o.ProviderCertificate
	}
	if true {
		toSerialize["redirectUrl"] = o.RedirectUrl
	}
	return json.Marshal(toSerialize)
}

type NullableSamlProvider struct {
	value *SamlProvider
	isSet bool
}

func (v NullableSamlProvider) Get() *SamlProvider {
	return v.value
}

func (v *NullableSamlProvider) Set(val *SamlProvider) {
	v.value = val
	v.isSet = true
}

func (v NullableSamlProvider) IsSet() bool {
	return v.isSet
}

func (v *NullableSamlProvider) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSamlProvider(val *SamlProvider) *NullableSamlProvider {
	return &NullableSamlProvider{value: val, isSet: true}
}

func (v NullableSamlProvider) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSamlProvider) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
