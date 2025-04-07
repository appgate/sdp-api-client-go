/*
Appgate SDP Controller REST API

# About   This specification documents the REST API calls for the Appgate SDP Controller.    Please refer to the REST API chapter in the manual or contact Appgate support with any questions about   this functionality. # Getting Started   Requirements for API scripting:   - Access to the Admin/API TLS Connection (default port 8443) of a Controller appliance.     (https://sdphelp.appgate.com/adminguide/appliance-function-configure.html?anchor=admin-api)   - An API user with relevant permissions.     (https://sdphelp.appgate.com/adminguide/administrative-roles-configure.html)   - In order to use the simple login API, Admin MFA must be disabled or the API user must be excluded.     (https://sdphelp.appgate.com/adminguide/mfa-for-admins.html) # Base path   HTTPS requests must be sent to the Admin Interface hostname and port, with **_/admin** path.    For example: **https://appgate.company.com:8443/admin**    All requests must have the **Accept** header as:    **application/vnd.appgate.peer-v22+json**    An exception is made for the **_/admin/version** endpoint which instead expects an **application/json** Accept header. # API Conventions   API conventions are  important to understand and follow strictly.    - While updating objects (via PUT), entire object must be sent with all fields.     - For example, in order to add a remedy method to the condition below:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": []       }       ```     - send the entire object with updated and non-updated fields:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": [{\"type\": \"DisplayMessage\", \"message\": \"test message\"}]       }       ```    - In case Controller returns an error (non-2xx HTTP status code), response body is JSON.     The \"message\" field contains information about the error.     HTTP 422 \"Unprocessable Entity\" has extra `errors` field to list all the issues with specific fields.    - Empty string (\"\") is considered a different value than \"null\" or field being omitted from JSON.     Omitting the field is recommended if no value is intended.     Empty string (\"\") will be almost always rejected as invalid value.    - There are common pattern between many objects:     - **Configuration Objects**: There are many objects with common fields, namely \"id\", \"name\", \"notes\", \"created\"       and \"updated\". These entities are listed, queried, created, updated and deleted in a similar fashion.     - **Distinguished Name**: Users and Devices are identified with what is called Distinguished Names, as used in        LDAP. The distinguished format that identifies a device and a user combination is        \"CN=\\<Device ID\\>,CN=\\<username\\>,OU=\\<Identity Provider Name\\>\". Some objects have the        \"userDistinguishedName\" field, which does not include the CN for Device ID.        This identifies a user on every device.

API version: API version 22.1
Contact: appgatesdp.support@appgate.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// IssuedCertificate Issued Certificate by Appgate CA.
type IssuedCertificate struct {
	// Random ID assigned to the certificate.
	Id *string `json:"id,omitempty"`
	// Type of the certificate.
	Type *string `json:"type,omitempty"`
	// The subject name of the certificate.
	Subject *string `json:"subject,omitempty"`
	// The issuer name of the certificate.
	Issuer *string `json:"issuer,omitempty"`
	// X.509 certificate serial number.
	Serial *string `json:"serial,omitempty"`
	// SHA256 fingerprint of the certificate.
	FingerprintSha256 *string `json:"fingerprintSha256,omitempty"`
	// Since when the certificate is valid from.
	ValidFrom *time.Time `json:"validFrom,omitempty"`
	// Until when the certificate is valid.
	ValidTo *time.Time `json:"validTo,omitempty"`
	// PEM formatted public certificate.
	Pem *string `json:"pem,omitempty"`
	// When the Controller issued the certificate.
	IssueTime *time.Time `json:"issueTime,omitempty"`
	// Whether the certificate is revoked or not.
	Revoked *bool `json:"revoked,omitempty"`
	// X509 certificate revocation reason. See RFC 5280.
	RevocationReason *string `json:"revocationReason,omitempty"`
	// The revocation time of the certificate.
	RevocationTime *time.Time `json:"revocationTime,omitempty"`
	// Free-text notes for revocation reason.
	RevocationNotes *string `json:"revocationNotes,omitempty"`
}

// NewIssuedCertificate instantiates a new IssuedCertificate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIssuedCertificate() *IssuedCertificate {
	this := IssuedCertificate{}
	var revoked bool = false
	this.Revoked = &revoked
	return &this
}

// NewIssuedCertificateWithDefaults instantiates a new IssuedCertificate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIssuedCertificateWithDefaults() *IssuedCertificate {
	this := IssuedCertificate{}
	var revoked bool = false
	this.Revoked = &revoked
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *IssuedCertificate) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssuedCertificate) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *IssuedCertificate) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *IssuedCertificate) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *IssuedCertificate) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssuedCertificate) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *IssuedCertificate) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *IssuedCertificate) SetType(v string) {
	o.Type = &v
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *IssuedCertificate) GetSubject() string {
	if o == nil || o.Subject == nil {
		var ret string
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssuedCertificate) GetSubjectOk() (*string, bool) {
	if o == nil || o.Subject == nil {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *IssuedCertificate) HasSubject() bool {
	if o != nil && o.Subject != nil {
		return true
	}

	return false
}

// SetSubject gets a reference to the given string and assigns it to the Subject field.
func (o *IssuedCertificate) SetSubject(v string) {
	o.Subject = &v
}

// GetIssuer returns the Issuer field value if set, zero value otherwise.
func (o *IssuedCertificate) GetIssuer() string {
	if o == nil || o.Issuer == nil {
		var ret string
		return ret
	}
	return *o.Issuer
}

// GetIssuerOk returns a tuple with the Issuer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssuedCertificate) GetIssuerOk() (*string, bool) {
	if o == nil || o.Issuer == nil {
		return nil, false
	}
	return o.Issuer, true
}

// HasIssuer returns a boolean if a field has been set.
func (o *IssuedCertificate) HasIssuer() bool {
	if o != nil && o.Issuer != nil {
		return true
	}

	return false
}

// SetIssuer gets a reference to the given string and assigns it to the Issuer field.
func (o *IssuedCertificate) SetIssuer(v string) {
	o.Issuer = &v
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *IssuedCertificate) GetSerial() string {
	if o == nil || o.Serial == nil {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssuedCertificate) GetSerialOk() (*string, bool) {
	if o == nil || o.Serial == nil {
		return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *IssuedCertificate) HasSerial() bool {
	if o != nil && o.Serial != nil {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *IssuedCertificate) SetSerial(v string) {
	o.Serial = &v
}

// GetFingerprintSha256 returns the FingerprintSha256 field value if set, zero value otherwise.
func (o *IssuedCertificate) GetFingerprintSha256() string {
	if o == nil || o.FingerprintSha256 == nil {
		var ret string
		return ret
	}
	return *o.FingerprintSha256
}

// GetFingerprintSha256Ok returns a tuple with the FingerprintSha256 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssuedCertificate) GetFingerprintSha256Ok() (*string, bool) {
	if o == nil || o.FingerprintSha256 == nil {
		return nil, false
	}
	return o.FingerprintSha256, true
}

// HasFingerprintSha256 returns a boolean if a field has been set.
func (o *IssuedCertificate) HasFingerprintSha256() bool {
	if o != nil && o.FingerprintSha256 != nil {
		return true
	}

	return false
}

// SetFingerprintSha256 gets a reference to the given string and assigns it to the FingerprintSha256 field.
func (o *IssuedCertificate) SetFingerprintSha256(v string) {
	o.FingerprintSha256 = &v
}

// GetValidFrom returns the ValidFrom field value if set, zero value otherwise.
func (o *IssuedCertificate) GetValidFrom() time.Time {
	if o == nil || o.ValidFrom == nil {
		var ret time.Time
		return ret
	}
	return *o.ValidFrom
}

// GetValidFromOk returns a tuple with the ValidFrom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssuedCertificate) GetValidFromOk() (*time.Time, bool) {
	if o == nil || o.ValidFrom == nil {
		return nil, false
	}
	return o.ValidFrom, true
}

// HasValidFrom returns a boolean if a field has been set.
func (o *IssuedCertificate) HasValidFrom() bool {
	if o != nil && o.ValidFrom != nil {
		return true
	}

	return false
}

// SetValidFrom gets a reference to the given time.Time and assigns it to the ValidFrom field.
func (o *IssuedCertificate) SetValidFrom(v time.Time) {
	o.ValidFrom = &v
}

// GetValidTo returns the ValidTo field value if set, zero value otherwise.
func (o *IssuedCertificate) GetValidTo() time.Time {
	if o == nil || o.ValidTo == nil {
		var ret time.Time
		return ret
	}
	return *o.ValidTo
}

// GetValidToOk returns a tuple with the ValidTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssuedCertificate) GetValidToOk() (*time.Time, bool) {
	if o == nil || o.ValidTo == nil {
		return nil, false
	}
	return o.ValidTo, true
}

// HasValidTo returns a boolean if a field has been set.
func (o *IssuedCertificate) HasValidTo() bool {
	if o != nil && o.ValidTo != nil {
		return true
	}

	return false
}

// SetValidTo gets a reference to the given time.Time and assigns it to the ValidTo field.
func (o *IssuedCertificate) SetValidTo(v time.Time) {
	o.ValidTo = &v
}

// GetPem returns the Pem field value if set, zero value otherwise.
func (o *IssuedCertificate) GetPem() string {
	if o == nil || o.Pem == nil {
		var ret string
		return ret
	}
	return *o.Pem
}

// GetPemOk returns a tuple with the Pem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssuedCertificate) GetPemOk() (*string, bool) {
	if o == nil || o.Pem == nil {
		return nil, false
	}
	return o.Pem, true
}

// HasPem returns a boolean if a field has been set.
func (o *IssuedCertificate) HasPem() bool {
	if o != nil && o.Pem != nil {
		return true
	}

	return false
}

// SetPem gets a reference to the given string and assigns it to the Pem field.
func (o *IssuedCertificate) SetPem(v string) {
	o.Pem = &v
}

// GetIssueTime returns the IssueTime field value if set, zero value otherwise.
func (o *IssuedCertificate) GetIssueTime() time.Time {
	if o == nil || o.IssueTime == nil {
		var ret time.Time
		return ret
	}
	return *o.IssueTime
}

// GetIssueTimeOk returns a tuple with the IssueTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssuedCertificate) GetIssueTimeOk() (*time.Time, bool) {
	if o == nil || o.IssueTime == nil {
		return nil, false
	}
	return o.IssueTime, true
}

// HasIssueTime returns a boolean if a field has been set.
func (o *IssuedCertificate) HasIssueTime() bool {
	if o != nil && o.IssueTime != nil {
		return true
	}

	return false
}

// SetIssueTime gets a reference to the given time.Time and assigns it to the IssueTime field.
func (o *IssuedCertificate) SetIssueTime(v time.Time) {
	o.IssueTime = &v
}

// GetRevoked returns the Revoked field value if set, zero value otherwise.
func (o *IssuedCertificate) GetRevoked() bool {
	if o == nil || o.Revoked == nil {
		var ret bool
		return ret
	}
	return *o.Revoked
}

// GetRevokedOk returns a tuple with the Revoked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssuedCertificate) GetRevokedOk() (*bool, bool) {
	if o == nil || o.Revoked == nil {
		return nil, false
	}
	return o.Revoked, true
}

// HasRevoked returns a boolean if a field has been set.
func (o *IssuedCertificate) HasRevoked() bool {
	if o != nil && o.Revoked != nil {
		return true
	}

	return false
}

// SetRevoked gets a reference to the given bool and assigns it to the Revoked field.
func (o *IssuedCertificate) SetRevoked(v bool) {
	o.Revoked = &v
}

// GetRevocationReason returns the RevocationReason field value if set, zero value otherwise.
func (o *IssuedCertificate) GetRevocationReason() string {
	if o == nil || o.RevocationReason == nil {
		var ret string
		return ret
	}
	return *o.RevocationReason
}

// GetRevocationReasonOk returns a tuple with the RevocationReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssuedCertificate) GetRevocationReasonOk() (*string, bool) {
	if o == nil || o.RevocationReason == nil {
		return nil, false
	}
	return o.RevocationReason, true
}

// HasRevocationReason returns a boolean if a field has been set.
func (o *IssuedCertificate) HasRevocationReason() bool {
	if o != nil && o.RevocationReason != nil {
		return true
	}

	return false
}

// SetRevocationReason gets a reference to the given string and assigns it to the RevocationReason field.
func (o *IssuedCertificate) SetRevocationReason(v string) {
	o.RevocationReason = &v
}

// GetRevocationTime returns the RevocationTime field value if set, zero value otherwise.
func (o *IssuedCertificate) GetRevocationTime() time.Time {
	if o == nil || o.RevocationTime == nil {
		var ret time.Time
		return ret
	}
	return *o.RevocationTime
}

// GetRevocationTimeOk returns a tuple with the RevocationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssuedCertificate) GetRevocationTimeOk() (*time.Time, bool) {
	if o == nil || o.RevocationTime == nil {
		return nil, false
	}
	return o.RevocationTime, true
}

// HasRevocationTime returns a boolean if a field has been set.
func (o *IssuedCertificate) HasRevocationTime() bool {
	if o != nil && o.RevocationTime != nil {
		return true
	}

	return false
}

// SetRevocationTime gets a reference to the given time.Time and assigns it to the RevocationTime field.
func (o *IssuedCertificate) SetRevocationTime(v time.Time) {
	o.RevocationTime = &v
}

// GetRevocationNotes returns the RevocationNotes field value if set, zero value otherwise.
func (o *IssuedCertificate) GetRevocationNotes() string {
	if o == nil || o.RevocationNotes == nil {
		var ret string
		return ret
	}
	return *o.RevocationNotes
}

// GetRevocationNotesOk returns a tuple with the RevocationNotes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssuedCertificate) GetRevocationNotesOk() (*string, bool) {
	if o == nil || o.RevocationNotes == nil {
		return nil, false
	}
	return o.RevocationNotes, true
}

// HasRevocationNotes returns a boolean if a field has been set.
func (o *IssuedCertificate) HasRevocationNotes() bool {
	if o != nil && o.RevocationNotes != nil {
		return true
	}

	return false
}

// SetRevocationNotes gets a reference to the given string and assigns it to the RevocationNotes field.
func (o *IssuedCertificate) SetRevocationNotes(v string) {
	o.RevocationNotes = &v
}

func (o IssuedCertificate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Subject != nil {
		toSerialize["subject"] = o.Subject
	}
	if o.Issuer != nil {
		toSerialize["issuer"] = o.Issuer
	}
	if o.Serial != nil {
		toSerialize["serial"] = o.Serial
	}
	if o.FingerprintSha256 != nil {
		toSerialize["fingerprintSha256"] = o.FingerprintSha256
	}
	if o.ValidFrom != nil {
		toSerialize["validFrom"] = o.ValidFrom
	}
	if o.ValidTo != nil {
		toSerialize["validTo"] = o.ValidTo
	}
	if o.Pem != nil {
		toSerialize["pem"] = o.Pem
	}
	if o.IssueTime != nil {
		toSerialize["issueTime"] = o.IssueTime
	}
	if o.Revoked != nil {
		toSerialize["revoked"] = o.Revoked
	}
	if o.RevocationReason != nil {
		toSerialize["revocationReason"] = o.RevocationReason
	}
	if o.RevocationTime != nil {
		toSerialize["revocationTime"] = o.RevocationTime
	}
	if o.RevocationNotes != nil {
		toSerialize["revocationNotes"] = o.RevocationNotes
	}
	return json.Marshal(toSerialize)
}

type NullableIssuedCertificate struct {
	value *IssuedCertificate
	isSet bool
}

func (v NullableIssuedCertificate) Get() *IssuedCertificate {
	return v.value
}

func (v *NullableIssuedCertificate) Set(val *IssuedCertificate) {
	v.value = val
	v.isSet = true
}

func (v NullableIssuedCertificate) IsSet() bool {
	return v.isSet
}

func (v *NullableIssuedCertificate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIssuedCertificate(val *IssuedCertificate) *NullableIssuedCertificate {
	return &NullableIssuedCertificate{value: val, isSet: true}
}

func (v NullableIssuedCertificate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIssuedCertificate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
