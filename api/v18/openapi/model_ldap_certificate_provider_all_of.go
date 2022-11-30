/*
Appgate SDP Controller REST API

# About   This specification documents the REST API calls for the Appgate SDP Controller.    Please refer to the REST API chapter in the manual or contact Appgate support with any questions about   this functionality. # Getting Started   Requirements for API scripting:   - Access to the Admin/API TLS Connection (default port 8443) of a Controller appliance.     (https://sdphelp.appgate.com/adminguide/appliance-function-configure.html?anchor=admin-api)   - An API user with relevant permissions.     (https://sdphelp.appgate.com/adminguide/administrative-roles-configure.html)   - In order to use the simple login API, Admin MFA must be disabled or the API user must be excluded.     (https://sdphelp.appgate.com/adminguide/mfa-for-admins.html) # Base path   HTTPS requests must be sent to the Admin Interface hostname and port, with **_/admin** path.    For example: **https://appgate.company.com:8443/admin**    All requests must have the **Accept** header as:    **application/vnd.appgate.peer-v18+json** # API Conventions   API conventions are  important to understand and follow strictly.    - While updating objects (via PUT), entire object must be sent with all fields.     - For example, in order to add a remedy method to the condition below:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": []       }       ```     - send the entire object with updated and non-updated fields:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": [{\"type\": \"DisplayMessage\", \"message\": \"test message\"}]       }       ```    - In case Controller returns an error (non-2xx HTTP status code), response body is JSON.     The \"message\" field contains information about the error.     HTTP 422 \"Unprocessable Entity\" has extra `errors` field to list all the issues with specific fields.    - Empty string (\"\") is considered a different value than \"null\" or field being omitted from JSON.     Omitting the field is recommended if no value is intended.     Empty string (\"\") will be almost always rejected as invalid value.    - There are common pattern between many objects:     - **Configuration Objects**: There are many objects with common fields, namely \"id\", \"name\", \"notes\", \"created\"       and \"updated\". These entities are listed, queried, created, updated and deleted in a similar fashion.     - **Distinguished Name**: Users and Devices are identified with what is called Distinguished Names, as used in        LDAP. The distinguished format that identifies a device and a user combination is        \"CN=\\<Device ID\\>,CN=\\<username\\>,OU=\\<Identity Provider Name\\>\". Some objects have the        \"userDistinguishedName\" field, which does not include the CN for Device ID.        This identifies a user on every device.

API version: API version 18.0
Contact: appgatesdp.support@appgate.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// LdapCertificateProviderAllOf Represents an LDAP Certificate Identity Provider.
type LdapCertificateProviderAllOf struct {
	// CA certificates to verify the Client certificates. In PEM format.
	CaCertificates []string `json:"caCertificates"`
	// The LDAP attribute to compare the Client certificate's Subject Alternative Name.
	CertificateUserAttribute *string `json:"certificateUserAttribute,omitempty"`
	// The LDAP attribute to compare the Client certificate binary. Leave it null to skip this comparison.
	CertificateAttribute *string `json:"certificateAttribute,omitempty"`
	// By default, Controller contacts the endpoints on the certificate extensions in order to verify revocation status and pull the intermediate CA certificates. Set this flag in order to skip them.
	SkipX509ExternalChecks *bool `json:"skipX509ExternalChecks,omitempty"`
	// Client will order the available certificates according to the given priority list.
	CertificatePriorities []LdapCertificateProviderAllOfCertificatePriorities `json:"certificatePriorities,omitempty"`
}

// NewLdapCertificateProviderAllOf instantiates a new LdapCertificateProviderAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLdapCertificateProviderAllOf(caCertificates []string) *LdapCertificateProviderAllOf {
	this := LdapCertificateProviderAllOf{}
	this.CaCertificates = caCertificates
	var certificateUserAttribute string = "userPrincipalName"
	this.CertificateUserAttribute = &certificateUserAttribute
	return &this
}

// NewLdapCertificateProviderAllOfWithDefaults instantiates a new LdapCertificateProviderAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLdapCertificateProviderAllOfWithDefaults() *LdapCertificateProviderAllOf {
	this := LdapCertificateProviderAllOf{}
	var certificateUserAttribute string = "userPrincipalName"
	this.CertificateUserAttribute = &certificateUserAttribute
	return &this
}

// GetCaCertificates returns the CaCertificates field value
func (o *LdapCertificateProviderAllOf) GetCaCertificates() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.CaCertificates
}

// GetCaCertificatesOk returns a tuple with the CaCertificates field value
// and a boolean to check if the value has been set.
func (o *LdapCertificateProviderAllOf) GetCaCertificatesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CaCertificates, true
}

// SetCaCertificates sets field value
func (o *LdapCertificateProviderAllOf) SetCaCertificates(v []string) {
	o.CaCertificates = v
}

// GetCertificateUserAttribute returns the CertificateUserAttribute field value if set, zero value otherwise.
func (o *LdapCertificateProviderAllOf) GetCertificateUserAttribute() string {
	if o == nil || o.CertificateUserAttribute == nil {
		var ret string
		return ret
	}
	return *o.CertificateUserAttribute
}

// GetCertificateUserAttributeOk returns a tuple with the CertificateUserAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapCertificateProviderAllOf) GetCertificateUserAttributeOk() (*string, bool) {
	if o == nil || o.CertificateUserAttribute == nil {
		return nil, false
	}
	return o.CertificateUserAttribute, true
}

// HasCertificateUserAttribute returns a boolean if a field has been set.
func (o *LdapCertificateProviderAllOf) HasCertificateUserAttribute() bool {
	if o != nil && o.CertificateUserAttribute != nil {
		return true
	}

	return false
}

// SetCertificateUserAttribute gets a reference to the given string and assigns it to the CertificateUserAttribute field.
func (o *LdapCertificateProviderAllOf) SetCertificateUserAttribute(v string) {
	o.CertificateUserAttribute = &v
}

// GetCertificateAttribute returns the CertificateAttribute field value if set, zero value otherwise.
func (o *LdapCertificateProviderAllOf) GetCertificateAttribute() string {
	if o == nil || o.CertificateAttribute == nil {
		var ret string
		return ret
	}
	return *o.CertificateAttribute
}

// GetCertificateAttributeOk returns a tuple with the CertificateAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapCertificateProviderAllOf) GetCertificateAttributeOk() (*string, bool) {
	if o == nil || o.CertificateAttribute == nil {
		return nil, false
	}
	return o.CertificateAttribute, true
}

// HasCertificateAttribute returns a boolean if a field has been set.
func (o *LdapCertificateProviderAllOf) HasCertificateAttribute() bool {
	if o != nil && o.CertificateAttribute != nil {
		return true
	}

	return false
}

// SetCertificateAttribute gets a reference to the given string and assigns it to the CertificateAttribute field.
func (o *LdapCertificateProviderAllOf) SetCertificateAttribute(v string) {
	o.CertificateAttribute = &v
}

// GetSkipX509ExternalChecks returns the SkipX509ExternalChecks field value if set, zero value otherwise.
func (o *LdapCertificateProviderAllOf) GetSkipX509ExternalChecks() bool {
	if o == nil || o.SkipX509ExternalChecks == nil {
		var ret bool
		return ret
	}
	return *o.SkipX509ExternalChecks
}

// GetSkipX509ExternalChecksOk returns a tuple with the SkipX509ExternalChecks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapCertificateProviderAllOf) GetSkipX509ExternalChecksOk() (*bool, bool) {
	if o == nil || o.SkipX509ExternalChecks == nil {
		return nil, false
	}
	return o.SkipX509ExternalChecks, true
}

// HasSkipX509ExternalChecks returns a boolean if a field has been set.
func (o *LdapCertificateProviderAllOf) HasSkipX509ExternalChecks() bool {
	if o != nil && o.SkipX509ExternalChecks != nil {
		return true
	}

	return false
}

// SetSkipX509ExternalChecks gets a reference to the given bool and assigns it to the SkipX509ExternalChecks field.
func (o *LdapCertificateProviderAllOf) SetSkipX509ExternalChecks(v bool) {
	o.SkipX509ExternalChecks = &v
}

// GetCertificatePriorities returns the CertificatePriorities field value if set, zero value otherwise.
func (o *LdapCertificateProviderAllOf) GetCertificatePriorities() []LdapCertificateProviderAllOfCertificatePriorities {
	if o == nil || o.CertificatePriorities == nil {
		var ret []LdapCertificateProviderAllOfCertificatePriorities
		return ret
	}
	return o.CertificatePriorities
}

// GetCertificatePrioritiesOk returns a tuple with the CertificatePriorities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapCertificateProviderAllOf) GetCertificatePrioritiesOk() ([]LdapCertificateProviderAllOfCertificatePriorities, bool) {
	if o == nil || o.CertificatePriorities == nil {
		return nil, false
	}
	return o.CertificatePriorities, true
}

// HasCertificatePriorities returns a boolean if a field has been set.
func (o *LdapCertificateProviderAllOf) HasCertificatePriorities() bool {
	if o != nil && o.CertificatePriorities != nil {
		return true
	}

	return false
}

// SetCertificatePriorities gets a reference to the given []LdapCertificateProviderAllOfCertificatePriorities and assigns it to the CertificatePriorities field.
func (o *LdapCertificateProviderAllOf) SetCertificatePriorities(v []LdapCertificateProviderAllOfCertificatePriorities) {
	o.CertificatePriorities = v
}

func (o LdapCertificateProviderAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["caCertificates"] = o.CaCertificates
	}
	if o.CertificateUserAttribute != nil {
		toSerialize["certificateUserAttribute"] = o.CertificateUserAttribute
	}
	if o.CertificateAttribute != nil {
		toSerialize["certificateAttribute"] = o.CertificateAttribute
	}
	if o.SkipX509ExternalChecks != nil {
		toSerialize["skipX509ExternalChecks"] = o.SkipX509ExternalChecks
	}
	if o.CertificatePriorities != nil {
		toSerialize["certificatePriorities"] = o.CertificatePriorities
	}
	return json.Marshal(toSerialize)
}

type NullableLdapCertificateProviderAllOf struct {
	value *LdapCertificateProviderAllOf
	isSet bool
}

func (v NullableLdapCertificateProviderAllOf) Get() *LdapCertificateProviderAllOf {
	return v.value
}

func (v *NullableLdapCertificateProviderAllOf) Set(val *LdapCertificateProviderAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableLdapCertificateProviderAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableLdapCertificateProviderAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLdapCertificateProviderAllOf(val *LdapCertificateProviderAllOf) *NullableLdapCertificateProviderAllOf {
	return &NullableLdapCertificateProviderAllOf{value: val, isSet: true}
}

func (v NullableLdapCertificateProviderAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLdapCertificateProviderAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}