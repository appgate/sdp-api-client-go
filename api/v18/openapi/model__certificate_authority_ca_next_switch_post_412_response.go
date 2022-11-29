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

// CertificateAuthorityCaNextSwitchPost412Response struct for CertificateAuthorityCaNextSwitchPost412Response
type CertificateAuthorityCaNextSwitchPost412Response struct {
	// Machine readable error code.
	Id *string `json:"id,omitempty"`
	// Human readable error details.
	Message *string `json:"message,omitempty"`
	// A dictionary of Appliance name and failure reason.
	FailedAppliances map[string]interface{} `json:"failedAppliances,omitempty"`
}

// NewCertificateAuthorityCaNextSwitchPost412Response instantiates a new CertificateAuthorityCaNextSwitchPost412Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCertificateAuthorityCaNextSwitchPost412Response() *CertificateAuthorityCaNextSwitchPost412Response {
	this := CertificateAuthorityCaNextSwitchPost412Response{}
	return &this
}

// NewCertificateAuthorityCaNextSwitchPost412ResponseWithDefaults instantiates a new CertificateAuthorityCaNextSwitchPost412Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCertificateAuthorityCaNextSwitchPost412ResponseWithDefaults() *CertificateAuthorityCaNextSwitchPost412Response {
	this := CertificateAuthorityCaNextSwitchPost412Response{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CertificateAuthorityCaNextSwitchPost412Response) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateAuthorityCaNextSwitchPost412Response) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CertificateAuthorityCaNextSwitchPost412Response) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CertificateAuthorityCaNextSwitchPost412Response) SetId(v string) {
	o.Id = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *CertificateAuthorityCaNextSwitchPost412Response) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateAuthorityCaNextSwitchPost412Response) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *CertificateAuthorityCaNextSwitchPost412Response) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *CertificateAuthorityCaNextSwitchPost412Response) SetMessage(v string) {
	o.Message = &v
}

// GetFailedAppliances returns the FailedAppliances field value if set, zero value otherwise.
func (o *CertificateAuthorityCaNextSwitchPost412Response) GetFailedAppliances() map[string]interface{} {
	if o == nil || o.FailedAppliances == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.FailedAppliances
}

// GetFailedAppliancesOk returns a tuple with the FailedAppliances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateAuthorityCaNextSwitchPost412Response) GetFailedAppliancesOk() (map[string]interface{}, bool) {
	if o == nil || o.FailedAppliances == nil {
		return nil, false
	}
	return o.FailedAppliances, true
}

// HasFailedAppliances returns a boolean if a field has been set.
func (o *CertificateAuthorityCaNextSwitchPost412Response) HasFailedAppliances() bool {
	if o != nil && o.FailedAppliances != nil {
		return true
	}

	return false
}

// SetFailedAppliances gets a reference to the given map[string]interface{} and assigns it to the FailedAppliances field.
func (o *CertificateAuthorityCaNextSwitchPost412Response) SetFailedAppliances(v map[string]interface{}) {
	o.FailedAppliances = v
}

func (o CertificateAuthorityCaNextSwitchPost412Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.FailedAppliances != nil {
		toSerialize["failedAppliances"] = o.FailedAppliances
	}
	return json.Marshal(toSerialize)
}

type NullableCertificateAuthorityCaNextSwitchPost412Response struct {
	value *CertificateAuthorityCaNextSwitchPost412Response
	isSet bool
}

func (v NullableCertificateAuthorityCaNextSwitchPost412Response) Get() *CertificateAuthorityCaNextSwitchPost412Response {
	return v.value
}

func (v *NullableCertificateAuthorityCaNextSwitchPost412Response) Set(val *CertificateAuthorityCaNextSwitchPost412Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCertificateAuthorityCaNextSwitchPost412Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCertificateAuthorityCaNextSwitchPost412Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCertificateAuthorityCaNextSwitchPost412Response(val *CertificateAuthorityCaNextSwitchPost412Response) *NullableCertificateAuthorityCaNextSwitchPost412Response {
	return &NullableCertificateAuthorityCaNextSwitchPost412Response{value: val, isSet: true}
}

func (v NullableCertificateAuthorityCaNextSwitchPost412Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCertificateAuthorityCaNextSwitchPost412Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
