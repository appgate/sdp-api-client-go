/*
Appgate SDP Controller REST API

# About   This specification documents the REST API calls for the Appgate SDP Controller.    Please refer to the REST API chapter in the manual or contact Appgate support with any questions about   this functionality. # Getting Started   Requirements for API scripting:   - Access to the Admin/API TLS Connection (default port 8443) of a Controller appliance.     (https://sdphelp.appgate.com/adminguide/appliance-function-configure.html?anchor=admin-api)   - An API user with relevant permissions.     (https://sdphelp.appgate.com/adminguide/administrative-roles-configure.html)   - In order to use the simple login API, Admin MFA must be disabled or the API user must be excluded.     (https://sdphelp.appgate.com/adminguide/mfa-for-admins.html) # Base path   HTTPS requests must be sent to the Admin Interface hostname and port, with **_/admin** path.    For example: **https://appgate.company.com:8443/admin**    All requests must have the **Accept** header as:    **application/vnd.appgate.peer-v18+json** # API Conventions   API conventions are  important to understand and follow strictly.    - While updating objects (via PUT), entire object must be sent with all fields.     - For example, in order to add a remedy method to the condition below:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": []       }       ```     - send the entire object with updated and non-updated fields:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": [{\"type\": \"DisplayMessage\", \"message\": \"test message\"}]       }       ```    - In case Controller returns an error (non-2xx HTTP status code), response body is JSON.     The \"message\" field contains information about the error.     HTTP 422 \"Unprocessable Entity\" has extra `errors` field to list all the issues with specific fields.    - Empty string (\"\") is considered a different value than \"null\" or field being omitted from JSON.     Omitting the field is recommended if no value is intended.     Empty string (\"\") will be almost always rejected as invalid value.    - There are common pattern between many objects:     - **Configuration Objects**: There are many objects with common fields, namely \"id\", \"name\", \"notes\", \"created\"       and \"updated\". These entities are listed, queried, created, updated and deleted in a similar fashion.     - **Distinguished Name**: Users and Devices are identified with what is called Distinguished Names, as used in        LDAP. The distinguished format that identifies a device and a user combination is        \"CN=\\<Device ID\\>,CN=\\<username\\>,OU=\\<Identity Provider Name\\>\". Some objects have the        \"userDistinguishedName\" field, which does not include the CN for Device ID.        This identifies a user on every device.

API version: API version 18.5
Contact: appgatesdp.support@appgate.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// CaConfig struct for CaConfig
type CaConfig struct {
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
	// Name constraints extension details for permitted hostnames or IPs.
	NameConstraintsPermitted []string `json:"nameConstraintsPermitted,omitempty"`
	// Name constraints extension details for excluded hostnames or IPs.
	NameConstraintsExcluded []string `json:"nameConstraintsExcluded,omitempty"`
	// The URL where the CRL is hosted.
	CrlUrl *string `json:"crlUrl,omitempty"`
}

// NewCaConfig instantiates a new CaConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCaConfig() *CaConfig {
	this := CaConfig{}
	return &this
}

// NewCaConfigWithDefaults instantiates a new CaConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCaConfigWithDefaults() *CaConfig {
	this := CaConfig{}
	return &this
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *CaConfig) GetVersion() float32 {
	if o == nil || o.Version == nil {
		var ret float32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaConfig) GetVersionOk() (*float32, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *CaConfig) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given float32 and assigns it to the Version field.
func (o *CaConfig) SetVersion(v float32) {
	o.Version = &v
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *CaConfig) GetSerial() string {
	if o == nil || o.Serial == nil {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaConfig) GetSerialOk() (*string, bool) {
	if o == nil || o.Serial == nil {
		return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *CaConfig) HasSerial() bool {
	if o != nil && o.Serial != nil {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *CaConfig) SetSerial(v string) {
	o.Serial = &v
}

// GetIssuer returns the Issuer field value if set, zero value otherwise.
func (o *CaConfig) GetIssuer() string {
	if o == nil || o.Issuer == nil {
		var ret string
		return ret
	}
	return *o.Issuer
}

// GetIssuerOk returns a tuple with the Issuer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaConfig) GetIssuerOk() (*string, bool) {
	if o == nil || o.Issuer == nil {
		return nil, false
	}
	return o.Issuer, true
}

// HasIssuer returns a boolean if a field has been set.
func (o *CaConfig) HasIssuer() bool {
	if o != nil && o.Issuer != nil {
		return true
	}

	return false
}

// SetIssuer gets a reference to the given string and assigns it to the Issuer field.
func (o *CaConfig) SetIssuer(v string) {
	o.Issuer = &v
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *CaConfig) GetSubject() string {
	if o == nil || o.Subject == nil {
		var ret string
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaConfig) GetSubjectOk() (*string, bool) {
	if o == nil || o.Subject == nil {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *CaConfig) HasSubject() bool {
	if o != nil && o.Subject != nil {
		return true
	}

	return false
}

// SetSubject gets a reference to the given string and assigns it to the Subject field.
func (o *CaConfig) SetSubject(v string) {
	o.Subject = &v
}

// GetValidFrom returns the ValidFrom field value if set, zero value otherwise.
func (o *CaConfig) GetValidFrom() time.Time {
	if o == nil || o.ValidFrom == nil {
		var ret time.Time
		return ret
	}
	return *o.ValidFrom
}

// GetValidFromOk returns a tuple with the ValidFrom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaConfig) GetValidFromOk() (*time.Time, bool) {
	if o == nil || o.ValidFrom == nil {
		return nil, false
	}
	return o.ValidFrom, true
}

// HasValidFrom returns a boolean if a field has been set.
func (o *CaConfig) HasValidFrom() bool {
	if o != nil && o.ValidFrom != nil {
		return true
	}

	return false
}

// SetValidFrom gets a reference to the given time.Time and assigns it to the ValidFrom field.
func (o *CaConfig) SetValidFrom(v time.Time) {
	o.ValidFrom = &v
}

// GetValidTo returns the ValidTo field value if set, zero value otherwise.
func (o *CaConfig) GetValidTo() time.Time {
	if o == nil || o.ValidTo == nil {
		var ret time.Time
		return ret
	}
	return *o.ValidTo
}

// GetValidToOk returns a tuple with the ValidTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaConfig) GetValidToOk() (*time.Time, bool) {
	if o == nil || o.ValidTo == nil {
		return nil, false
	}
	return o.ValidTo, true
}

// HasValidTo returns a boolean if a field has been set.
func (o *CaConfig) HasValidTo() bool {
	if o != nil && o.ValidTo != nil {
		return true
	}

	return false
}

// SetValidTo gets a reference to the given time.Time and assigns it to the ValidTo field.
func (o *CaConfig) SetValidTo(v time.Time) {
	o.ValidTo = &v
}

// GetFingerprint returns the Fingerprint field value if set, zero value otherwise.
func (o *CaConfig) GetFingerprint() string {
	if o == nil || o.Fingerprint == nil {
		var ret string
		return ret
	}
	return *o.Fingerprint
}

// GetFingerprintOk returns a tuple with the Fingerprint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaConfig) GetFingerprintOk() (*string, bool) {
	if o == nil || o.Fingerprint == nil {
		return nil, false
	}
	return o.Fingerprint, true
}

// HasFingerprint returns a boolean if a field has been set.
func (o *CaConfig) HasFingerprint() bool {
	if o != nil && o.Fingerprint != nil {
		return true
	}

	return false
}

// SetFingerprint gets a reference to the given string and assigns it to the Fingerprint field.
func (o *CaConfig) SetFingerprint(v string) {
	o.Fingerprint = &v
}

// GetCertificate returns the Certificate field value if set, zero value otherwise.
func (o *CaConfig) GetCertificate() string {
	if o == nil || o.Certificate == nil {
		var ret string
		return ret
	}
	return *o.Certificate
}

// GetCertificateOk returns a tuple with the Certificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaConfig) GetCertificateOk() (*string, bool) {
	if o == nil || o.Certificate == nil {
		return nil, false
	}
	return o.Certificate, true
}

// HasCertificate returns a boolean if a field has been set.
func (o *CaConfig) HasCertificate() bool {
	if o != nil && o.Certificate != nil {
		return true
	}

	return false
}

// SetCertificate gets a reference to the given string and assigns it to the Certificate field.
func (o *CaConfig) SetCertificate(v string) {
	o.Certificate = &v
}

// GetSubjectPublicKey returns the SubjectPublicKey field value if set, zero value otherwise.
func (o *CaConfig) GetSubjectPublicKey() string {
	if o == nil || o.SubjectPublicKey == nil {
		var ret string
		return ret
	}
	return *o.SubjectPublicKey
}

// GetSubjectPublicKeyOk returns a tuple with the SubjectPublicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaConfig) GetSubjectPublicKeyOk() (*string, bool) {
	if o == nil || o.SubjectPublicKey == nil {
		return nil, false
	}
	return o.SubjectPublicKey, true
}

// HasSubjectPublicKey returns a boolean if a field has been set.
func (o *CaConfig) HasSubjectPublicKey() bool {
	if o != nil && o.SubjectPublicKey != nil {
		return true
	}

	return false
}

// SetSubjectPublicKey gets a reference to the given string and assigns it to the SubjectPublicKey field.
func (o *CaConfig) SetSubjectPublicKey(v string) {
	o.SubjectPublicKey = &v
}

// GetNameConstraintsPermitted returns the NameConstraintsPermitted field value if set, zero value otherwise.
func (o *CaConfig) GetNameConstraintsPermitted() []string {
	if o == nil || o.NameConstraintsPermitted == nil {
		var ret []string
		return ret
	}
	return o.NameConstraintsPermitted
}

// GetNameConstraintsPermittedOk returns a tuple with the NameConstraintsPermitted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaConfig) GetNameConstraintsPermittedOk() ([]string, bool) {
	if o == nil || o.NameConstraintsPermitted == nil {
		return nil, false
	}
	return o.NameConstraintsPermitted, true
}

// HasNameConstraintsPermitted returns a boolean if a field has been set.
func (o *CaConfig) HasNameConstraintsPermitted() bool {
	if o != nil && o.NameConstraintsPermitted != nil {
		return true
	}

	return false
}

// SetNameConstraintsPermitted gets a reference to the given []string and assigns it to the NameConstraintsPermitted field.
func (o *CaConfig) SetNameConstraintsPermitted(v []string) {
	o.NameConstraintsPermitted = v
}

// GetNameConstraintsExcluded returns the NameConstraintsExcluded field value if set, zero value otherwise.
func (o *CaConfig) GetNameConstraintsExcluded() []string {
	if o == nil || o.NameConstraintsExcluded == nil {
		var ret []string
		return ret
	}
	return o.NameConstraintsExcluded
}

// GetNameConstraintsExcludedOk returns a tuple with the NameConstraintsExcluded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaConfig) GetNameConstraintsExcludedOk() ([]string, bool) {
	if o == nil || o.NameConstraintsExcluded == nil {
		return nil, false
	}
	return o.NameConstraintsExcluded, true
}

// HasNameConstraintsExcluded returns a boolean if a field has been set.
func (o *CaConfig) HasNameConstraintsExcluded() bool {
	if o != nil && o.NameConstraintsExcluded != nil {
		return true
	}

	return false
}

// SetNameConstraintsExcluded gets a reference to the given []string and assigns it to the NameConstraintsExcluded field.
func (o *CaConfig) SetNameConstraintsExcluded(v []string) {
	o.NameConstraintsExcluded = v
}

// GetCrlUrl returns the CrlUrl field value if set, zero value otherwise.
func (o *CaConfig) GetCrlUrl() string {
	if o == nil || o.CrlUrl == nil {
		var ret string
		return ret
	}
	return *o.CrlUrl
}

// GetCrlUrlOk returns a tuple with the CrlUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaConfig) GetCrlUrlOk() (*string, bool) {
	if o == nil || o.CrlUrl == nil {
		return nil, false
	}
	return o.CrlUrl, true
}

// HasCrlUrl returns a boolean if a field has been set.
func (o *CaConfig) HasCrlUrl() bool {
	if o != nil && o.CrlUrl != nil {
		return true
	}

	return false
}

// SetCrlUrl gets a reference to the given string and assigns it to the CrlUrl field.
func (o *CaConfig) SetCrlUrl(v string) {
	o.CrlUrl = &v
}

func (o CaConfig) MarshalJSON() ([]byte, error) {
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
	if o.NameConstraintsPermitted != nil {
		toSerialize["nameConstraintsPermitted"] = o.NameConstraintsPermitted
	}
	if o.NameConstraintsExcluded != nil {
		toSerialize["nameConstraintsExcluded"] = o.NameConstraintsExcluded
	}
	if o.CrlUrl != nil {
		toSerialize["crlUrl"] = o.CrlUrl
	}
	return json.Marshal(toSerialize)
}

type NullableCaConfig struct {
	value *CaConfig
	isSet bool
}

func (v NullableCaConfig) Get() *CaConfig {
	return v.value
}

func (v *NullableCaConfig) Set(val *CaConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableCaConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableCaConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCaConfig(val *CaConfig) *NullableCaConfig {
	return &NullableCaConfig{value: val, isSet: true}
}

func (v NullableCaConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCaConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
