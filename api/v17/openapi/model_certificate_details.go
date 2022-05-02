/*
Appgate SDP Controller REST API

# About   This specification documents the REST API calls for the Appgate SDP Controller.    Please refer to the REST API chapter in the manual or contact Appgate support with any questions about   this functionality. # Getting Started   Requirements for API scripting:   - Access to the Admin/API TLS Connection (default port 8443) of a Controller appliance.     (https://sdphelp.appgate.com/adminguide/appliance-function-configure.html?anchor=admin-api)   - An API user with relevant permissions.     (https://sdphelp.appgate.com/adminguide/administrative-roles-configure.html)   - In order to use the simple login API, Admin MFA must be disabled or the API user must be excluded.     (https://sdphelp.appgate.com/adminguide/mfa-for-admins.html) # Base path   HTTPS requests must be sent to the Admin Interface hostname and port, with **_/admin** path.    For example: **https://appgate.company.com:8443/admin**    All requests must have the **Accept** header as:    **application/vnd.appgate.peer-v16+json** # API Conventions   API conventions are  important to understand and follow strictly.    - While updating objects (via PUT), entire object must be sent with all fields.     - For example, in order to add a remedy method to the condition below:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": []       }       ```     - send the entire object with updated and non-updated fields:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": [{\"type\": \"DisplayMessage\", \"message\": \"test message\"}]       }       ```    - In case Controller returns an error (non-2xx HTTP status code), response body is JSON.     The \"message\" field contains information about the error.     HTTP 422 \"Unprocessable Entity\" has extra `errors` field to list all the issues with specific fields.    - Empty string (\"\") is considered a different value than \"null\" or field being omitted from JSON.     Omitting the field is recommend if no value is intended.     Empty string (\"\") will be almost always rejected as invalid value.    - There are common pattern between many objects:     - **Configuration Objects**: There are many objects with common fields, namely \"id\", \"name\", \"notes\", \"created\"       and \"updated\". These entities are listed, queried, created, updated and deleted in a similar fashion.     - **Distinguished Name**: Users and Devices are identified with what is called Distinguished Names, as used in        LDAP. The distinguished format that identifies a device and a user combination is        \"CN=\\<Device ID\\>,CN=\\<username\\>,OU=\\<Identity Provider Name\\>\". Some objects have the        \"userDistinguishedName\" field, which does not include the CN for Device ID.        This identifies a user on every device.

API version: API version 16.3
Contact: appgatesdp.support@appgate.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// CertificateDetails X509 certificate details.
type CertificateDetails struct {
	// X.509 certificate version.
	Version *float32 `json:"version,omitempty"`
	// X.509 certificate serial number.
	Serial *string `json:"serial,omitempty"`
	// The issuer name of the certificate.
	Issuer *string `json:"issuer,omitempty"`
	// The subject name of the certificate.
	Subject *string `json:"subject,omitempty"`
	// Since when the certificate is valid from.
	ValidFrom *time.Time `json:"validFrom,omitempty"`
	// Until when the certificate is valid.
	ValidTo *time.Time `json:"validTo,omitempty"`
	// SHA256 fingerprint of the certificate.
	Fingerprint *string `json:"fingerprint,omitempty"`
	// Base64 encoded binary of the certificate. Either DER or PEM formatted depending on the request.
	Certificate *string `json:"certificate,omitempty"`
	// Base64 encoded public key of the certificate.
	SubjectPublicKey *string `json:"subjectPublicKey,omitempty"`
}

// NewCertificateDetails instantiates a new CertificateDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCertificateDetails() *CertificateDetails {
	this := CertificateDetails{}
	return &this
}

// NewCertificateDetailsWithDefaults instantiates a new CertificateDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCertificateDetailsWithDefaults() *CertificateDetails {
	this := CertificateDetails{}
	return &this
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *CertificateDetails) GetVersion() float32 {
	if o == nil || o.Version == nil {
		var ret float32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateDetails) GetVersionOk() (*float32, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *CertificateDetails) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given float32 and assigns it to the Version field.
func (o *CertificateDetails) SetVersion(v float32) {
	o.Version = &v
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *CertificateDetails) GetSerial() string {
	if o == nil || o.Serial == nil {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateDetails) GetSerialOk() (*string, bool) {
	if o == nil || o.Serial == nil {
		return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *CertificateDetails) HasSerial() bool {
	if o != nil && o.Serial != nil {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *CertificateDetails) SetSerial(v string) {
	o.Serial = &v
}

// GetIssuer returns the Issuer field value if set, zero value otherwise.
func (o *CertificateDetails) GetIssuer() string {
	if o == nil || o.Issuer == nil {
		var ret string
		return ret
	}
	return *o.Issuer
}

// GetIssuerOk returns a tuple with the Issuer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateDetails) GetIssuerOk() (*string, bool) {
	if o == nil || o.Issuer == nil {
		return nil, false
	}
	return o.Issuer, true
}

// HasIssuer returns a boolean if a field has been set.
func (o *CertificateDetails) HasIssuer() bool {
	if o != nil && o.Issuer != nil {
		return true
	}

	return false
}

// SetIssuer gets a reference to the given string and assigns it to the Issuer field.
func (o *CertificateDetails) SetIssuer(v string) {
	o.Issuer = &v
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *CertificateDetails) GetSubject() string {
	if o == nil || o.Subject == nil {
		var ret string
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateDetails) GetSubjectOk() (*string, bool) {
	if o == nil || o.Subject == nil {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *CertificateDetails) HasSubject() bool {
	if o != nil && o.Subject != nil {
		return true
	}

	return false
}

// SetSubject gets a reference to the given string and assigns it to the Subject field.
func (o *CertificateDetails) SetSubject(v string) {
	o.Subject = &v
}

// GetValidFrom returns the ValidFrom field value if set, zero value otherwise.
func (o *CertificateDetails) GetValidFrom() time.Time {
	if o == nil || o.ValidFrom == nil {
		var ret time.Time
		return ret
	}
	return *o.ValidFrom
}

// GetValidFromOk returns a tuple with the ValidFrom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateDetails) GetValidFromOk() (*time.Time, bool) {
	if o == nil || o.ValidFrom == nil {
		return nil, false
	}
	return o.ValidFrom, true
}

// HasValidFrom returns a boolean if a field has been set.
func (o *CertificateDetails) HasValidFrom() bool {
	if o != nil && o.ValidFrom != nil {
		return true
	}

	return false
}

// SetValidFrom gets a reference to the given time.Time and assigns it to the ValidFrom field.
func (o *CertificateDetails) SetValidFrom(v time.Time) {
	o.ValidFrom = &v
}

// GetValidTo returns the ValidTo field value if set, zero value otherwise.
func (o *CertificateDetails) GetValidTo() time.Time {
	if o == nil || o.ValidTo == nil {
		var ret time.Time
		return ret
	}
	return *o.ValidTo
}

// GetValidToOk returns a tuple with the ValidTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateDetails) GetValidToOk() (*time.Time, bool) {
	if o == nil || o.ValidTo == nil {
		return nil, false
	}
	return o.ValidTo, true
}

// HasValidTo returns a boolean if a field has been set.
func (o *CertificateDetails) HasValidTo() bool {
	if o != nil && o.ValidTo != nil {
		return true
	}

	return false
}

// SetValidTo gets a reference to the given time.Time and assigns it to the ValidTo field.
func (o *CertificateDetails) SetValidTo(v time.Time) {
	o.ValidTo = &v
}

// GetFingerprint returns the Fingerprint field value if set, zero value otherwise.
func (o *CertificateDetails) GetFingerprint() string {
	if o == nil || o.Fingerprint == nil {
		var ret string
		return ret
	}
	return *o.Fingerprint
}

// GetFingerprintOk returns a tuple with the Fingerprint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateDetails) GetFingerprintOk() (*string, bool) {
	if o == nil || o.Fingerprint == nil {
		return nil, false
	}
	return o.Fingerprint, true
}

// HasFingerprint returns a boolean if a field has been set.
func (o *CertificateDetails) HasFingerprint() bool {
	if o != nil && o.Fingerprint != nil {
		return true
	}

	return false
}

// SetFingerprint gets a reference to the given string and assigns it to the Fingerprint field.
func (o *CertificateDetails) SetFingerprint(v string) {
	o.Fingerprint = &v
}

// GetCertificate returns the Certificate field value if set, zero value otherwise.
func (o *CertificateDetails) GetCertificate() string {
	if o == nil || o.Certificate == nil {
		var ret string
		return ret
	}
	return *o.Certificate
}

// GetCertificateOk returns a tuple with the Certificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateDetails) GetCertificateOk() (*string, bool) {
	if o == nil || o.Certificate == nil {
		return nil, false
	}
	return o.Certificate, true
}

// HasCertificate returns a boolean if a field has been set.
func (o *CertificateDetails) HasCertificate() bool {
	if o != nil && o.Certificate != nil {
		return true
	}

	return false
}

// SetCertificate gets a reference to the given string and assigns it to the Certificate field.
func (o *CertificateDetails) SetCertificate(v string) {
	o.Certificate = &v
}

// GetSubjectPublicKey returns the SubjectPublicKey field value if set, zero value otherwise.
func (o *CertificateDetails) GetSubjectPublicKey() string {
	if o == nil || o.SubjectPublicKey == nil {
		var ret string
		return ret
	}
	return *o.SubjectPublicKey
}

// GetSubjectPublicKeyOk returns a tuple with the SubjectPublicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateDetails) GetSubjectPublicKeyOk() (*string, bool) {
	if o == nil || o.SubjectPublicKey == nil {
		return nil, false
	}
	return o.SubjectPublicKey, true
}

// HasSubjectPublicKey returns a boolean if a field has been set.
func (o *CertificateDetails) HasSubjectPublicKey() bool {
	if o != nil && o.SubjectPublicKey != nil {
		return true
	}

	return false
}

// SetSubjectPublicKey gets a reference to the given string and assigns it to the SubjectPublicKey field.
func (o *CertificateDetails) SetSubjectPublicKey(v string) {
	o.SubjectPublicKey = &v
}

func (o CertificateDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	if o.Serial != nil {
		toSerialize["serial"] = o.Serial
	}
	if o.Issuer != nil {
		toSerialize["issuer"] = o.Issuer
	}
	if o.Subject != nil {
		toSerialize["subject"] = o.Subject
	}
	if o.ValidFrom != nil {
		toSerialize["validFrom"] = o.ValidFrom
	}
	if o.ValidTo != nil {
		toSerialize["validTo"] = o.ValidTo
	}
	if o.Fingerprint != nil {
		toSerialize["fingerprint"] = o.Fingerprint
	}
	if o.Certificate != nil {
		toSerialize["certificate"] = o.Certificate
	}
	if o.SubjectPublicKey != nil {
		toSerialize["subjectPublicKey"] = o.SubjectPublicKey
	}
	return json.Marshal(toSerialize)
}

type NullableCertificateDetails struct {
	value *CertificateDetails
	isSet bool
}

func (v NullableCertificateDetails) Get() *CertificateDetails {
	return v.value
}

func (v *NullableCertificateDetails) Set(val *CertificateDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableCertificateDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableCertificateDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCertificateDetails(val *CertificateDetails) *NullableCertificateDetails {
	return &NullableCertificateDetails{value: val, isSet: true}
}

func (v NullableCertificateDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCertificateDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
