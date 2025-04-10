/*
Appgate SDP Controller REST API

# About   This specification documents the REST API calls for the Appgate SDP Controller.    Please refer to the REST API chapter in the manual or contact Appgate support with any questions about   this functionality. # Getting Started   Requirements for API scripting:   - Access to the Admin/API TLS Connection (default port 8443) of a Controller appliance.     (https://sdphelp.appgate.com/adminguide/appliance-function-configure.html?anchor=admin-api)   - An API user with relevant permissions.     (https://sdphelp.appgate.com/adminguide/administrative-roles-configure.html)   - In order to use the simple login API, Admin MFA must be disabled or the API user must be excluded.     (https://sdphelp.appgate.com/adminguide/mfa-for-admins.html) # Base path   HTTPS requests must be sent to the Admin Interface hostname and port, with **_/admin** path.    For example: **https://appgate.company.com:8443/admin**    All requests must have the **Accept** header as:    **application/vnd.appgate.peer-v22+json**    An exception is made for the **_/admin/version** endpoint which instead expects an **application/json** Accept header. # API Conventions   API conventions are  important to understand and follow strictly.    - While updating objects (via PUT), entire object must be sent with all fields.     - For example, in order to add a remedy method to the condition below:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": []       }       ```     - send the entire object with updated and non-updated fields:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": [{\"type\": \"DisplayMessage\", \"message\": \"test message\"}]       }       ```    - In case Controller returns an error (non-2xx HTTP status code), response body is JSON.     The \"message\" field contains information about the error.     HTTP 422 \"Unprocessable Entity\" has extra `errors` field to list all the issues with specific fields.    - Empty string (\"\") is considered a different value than \"null\" or field being omitted from JSON.     Omitting the field is recommended if no value is intended.     Empty string (\"\") will be almost always rejected as invalid value.    - There are common pattern between many objects:     - **Configuration Objects**: There are many objects with common fields, namely \"id\", \"name\", \"notes\", \"created\"       and \"updated\". These entities are listed, queried, created, updated and deleted in a similar fashion.     - **Distinguished Name**: Users and Devices are identified with what is called Distinguished Names, as used in        LDAP. The distinguished format that identifies a device and a user combination is        \"CN=\\<Device ID\\>,CN=\\<username\\>,OU=\\<Identity Provider Name\\>\". Some objects have the        \"userDistinguishedName\" field, which does not include the CN for Device ID.        This identifies a user on every device.

API version: API version 22.2
Contact: appgatesdp.support@appgate.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// LicenseDetailsUsageIdentityProviders The amount of non-built-in Identity Providers in the system at present broken out by type.
type LicenseDetailsUsageIdentityProviders struct {
	Ldap            *float32 `json:"ldap,omitempty"`
	LdapCertificate *float32 `json:"ldapCertificate,omitempty"`
	Radius          *float32 `json:"radius,omitempty"`
	Saml            *float32 `json:"saml,omitempty"`
	Oidc            *float32 `json:"oidc,omitempty"`
}

// NewLicenseDetailsUsageIdentityProviders instantiates a new LicenseDetailsUsageIdentityProviders object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLicenseDetailsUsageIdentityProviders() *LicenseDetailsUsageIdentityProviders {
	this := LicenseDetailsUsageIdentityProviders{}
	return &this
}

// NewLicenseDetailsUsageIdentityProvidersWithDefaults instantiates a new LicenseDetailsUsageIdentityProviders object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLicenseDetailsUsageIdentityProvidersWithDefaults() *LicenseDetailsUsageIdentityProviders {
	this := LicenseDetailsUsageIdentityProviders{}
	return &this
}

// GetLdap returns the Ldap field value if set, zero value otherwise.
func (o *LicenseDetailsUsageIdentityProviders) GetLdap() float32 {
	if o == nil || o.Ldap == nil {
		var ret float32
		return ret
	}
	return *o.Ldap
}

// GetLdapOk returns a tuple with the Ldap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseDetailsUsageIdentityProviders) GetLdapOk() (*float32, bool) {
	if o == nil || o.Ldap == nil {
		return nil, false
	}
	return o.Ldap, true
}

// HasLdap returns a boolean if a field has been set.
func (o *LicenseDetailsUsageIdentityProviders) HasLdap() bool {
	if o != nil && o.Ldap != nil {
		return true
	}

	return false
}

// SetLdap gets a reference to the given float32 and assigns it to the Ldap field.
func (o *LicenseDetailsUsageIdentityProviders) SetLdap(v float32) {
	o.Ldap = &v
}

// GetLdapCertificate returns the LdapCertificate field value if set, zero value otherwise.
func (o *LicenseDetailsUsageIdentityProviders) GetLdapCertificate() float32 {
	if o == nil || o.LdapCertificate == nil {
		var ret float32
		return ret
	}
	return *o.LdapCertificate
}

// GetLdapCertificateOk returns a tuple with the LdapCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseDetailsUsageIdentityProviders) GetLdapCertificateOk() (*float32, bool) {
	if o == nil || o.LdapCertificate == nil {
		return nil, false
	}
	return o.LdapCertificate, true
}

// HasLdapCertificate returns a boolean if a field has been set.
func (o *LicenseDetailsUsageIdentityProviders) HasLdapCertificate() bool {
	if o != nil && o.LdapCertificate != nil {
		return true
	}

	return false
}

// SetLdapCertificate gets a reference to the given float32 and assigns it to the LdapCertificate field.
func (o *LicenseDetailsUsageIdentityProviders) SetLdapCertificate(v float32) {
	o.LdapCertificate = &v
}

// GetRadius returns the Radius field value if set, zero value otherwise.
func (o *LicenseDetailsUsageIdentityProviders) GetRadius() float32 {
	if o == nil || o.Radius == nil {
		var ret float32
		return ret
	}
	return *o.Radius
}

// GetRadiusOk returns a tuple with the Radius field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseDetailsUsageIdentityProviders) GetRadiusOk() (*float32, bool) {
	if o == nil || o.Radius == nil {
		return nil, false
	}
	return o.Radius, true
}

// HasRadius returns a boolean if a field has been set.
func (o *LicenseDetailsUsageIdentityProviders) HasRadius() bool {
	if o != nil && o.Radius != nil {
		return true
	}

	return false
}

// SetRadius gets a reference to the given float32 and assigns it to the Radius field.
func (o *LicenseDetailsUsageIdentityProviders) SetRadius(v float32) {
	o.Radius = &v
}

// GetSaml returns the Saml field value if set, zero value otherwise.
func (o *LicenseDetailsUsageIdentityProviders) GetSaml() float32 {
	if o == nil || o.Saml == nil {
		var ret float32
		return ret
	}
	return *o.Saml
}

// GetSamlOk returns a tuple with the Saml field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseDetailsUsageIdentityProviders) GetSamlOk() (*float32, bool) {
	if o == nil || o.Saml == nil {
		return nil, false
	}
	return o.Saml, true
}

// HasSaml returns a boolean if a field has been set.
func (o *LicenseDetailsUsageIdentityProviders) HasSaml() bool {
	if o != nil && o.Saml != nil {
		return true
	}

	return false
}

// SetSaml gets a reference to the given float32 and assigns it to the Saml field.
func (o *LicenseDetailsUsageIdentityProviders) SetSaml(v float32) {
	o.Saml = &v
}

// GetOidc returns the Oidc field value if set, zero value otherwise.
func (o *LicenseDetailsUsageIdentityProviders) GetOidc() float32 {
	if o == nil || o.Oidc == nil {
		var ret float32
		return ret
	}
	return *o.Oidc
}

// GetOidcOk returns a tuple with the Oidc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseDetailsUsageIdentityProviders) GetOidcOk() (*float32, bool) {
	if o == nil || o.Oidc == nil {
		return nil, false
	}
	return o.Oidc, true
}

// HasOidc returns a boolean if a field has been set.
func (o *LicenseDetailsUsageIdentityProviders) HasOidc() bool {
	if o != nil && o.Oidc != nil {
		return true
	}

	return false
}

// SetOidc gets a reference to the given float32 and assigns it to the Oidc field.
func (o *LicenseDetailsUsageIdentityProviders) SetOidc(v float32) {
	o.Oidc = &v
}

func (o LicenseDetailsUsageIdentityProviders) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Ldap != nil {
		toSerialize["ldap"] = o.Ldap
	}
	if o.LdapCertificate != nil {
		toSerialize["ldapCertificate"] = o.LdapCertificate
	}
	if o.Radius != nil {
		toSerialize["radius"] = o.Radius
	}
	if o.Saml != nil {
		toSerialize["saml"] = o.Saml
	}
	if o.Oidc != nil {
		toSerialize["oidc"] = o.Oidc
	}
	return json.Marshal(toSerialize)
}

type NullableLicenseDetailsUsageIdentityProviders struct {
	value *LicenseDetailsUsageIdentityProviders
	isSet bool
}

func (v NullableLicenseDetailsUsageIdentityProviders) Get() *LicenseDetailsUsageIdentityProviders {
	return v.value
}

func (v *NullableLicenseDetailsUsageIdentityProviders) Set(val *LicenseDetailsUsageIdentityProviders) {
	v.value = val
	v.isSet = true
}

func (v NullableLicenseDetailsUsageIdentityProviders) IsSet() bool {
	return v.isSet
}

func (v *NullableLicenseDetailsUsageIdentityProviders) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLicenseDetailsUsageIdentityProviders(val *LicenseDetailsUsageIdentityProviders) *NullableLicenseDetailsUsageIdentityProviders {
	return &NullableLicenseDetailsUsageIdentityProviders{value: val, isSet: true}
}

func (v NullableLicenseDetailsUsageIdentityProviders) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLicenseDetailsUsageIdentityProviders) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
