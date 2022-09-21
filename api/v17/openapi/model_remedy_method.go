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

// RemedyMethod struct for RemedyMethod
type RemedyMethod struct {
	// User Action type.
	Type string `json:"type"`
	// Message to be shown to the user. Required for all remedy method.
	Message string `json:"message"`
	// Suffix to be added to the claim. Required for OtpAuthentication, PasswordAuthentication and Reason remedy methods.
	ClaimSuffix *string `json:"claimSuffix,omitempty"`
	// MFA Provider Id or Identity Provider Id. Required for some remedy method.
	ProviderId *string `json:"providerId,omitempty"`
}

// NewRemedyMethod instantiates a new RemedyMethod object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRemedyMethod(type_ string, message string) *RemedyMethod {
	this := RemedyMethod{}
	this.Type = type_
	this.Message = message
	return &this
}

// NewRemedyMethodWithDefaults instantiates a new RemedyMethod object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRemedyMethodWithDefaults() *RemedyMethod {
	this := RemedyMethod{}
	return &this
}

// GetType returns the Type field value
func (o *RemedyMethod) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RemedyMethod) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *RemedyMethod) SetType(v string) {
	o.Type = v
}

// GetMessage returns the Message field value
func (o *RemedyMethod) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *RemedyMethod) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *RemedyMethod) SetMessage(v string) {
	o.Message = v
}

// GetClaimSuffix returns the ClaimSuffix field value if set, zero value otherwise.
func (o *RemedyMethod) GetClaimSuffix() string {
	if o == nil || o.ClaimSuffix == nil {
		var ret string
		return ret
	}
	return *o.ClaimSuffix
}

// GetClaimSuffixOk returns a tuple with the ClaimSuffix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemedyMethod) GetClaimSuffixOk() (*string, bool) {
	if o == nil || o.ClaimSuffix == nil {
		return nil, false
	}
	return o.ClaimSuffix, true
}

// HasClaimSuffix returns a boolean if a field has been set.
func (o *RemedyMethod) HasClaimSuffix() bool {
	if o != nil && o.ClaimSuffix != nil {
		return true
	}

	return false
}

// SetClaimSuffix gets a reference to the given string and assigns it to the ClaimSuffix field.
func (o *RemedyMethod) SetClaimSuffix(v string) {
	o.ClaimSuffix = &v
}

// GetProviderId returns the ProviderId field value if set, zero value otherwise.
func (o *RemedyMethod) GetProviderId() string {
	if o == nil || o.ProviderId == nil {
		var ret string
		return ret
	}
	return *o.ProviderId
}

// GetProviderIdOk returns a tuple with the ProviderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemedyMethod) GetProviderIdOk() (*string, bool) {
	if o == nil || o.ProviderId == nil {
		return nil, false
	}
	return o.ProviderId, true
}

// HasProviderId returns a boolean if a field has been set.
func (o *RemedyMethod) HasProviderId() bool {
	if o != nil && o.ProviderId != nil {
		return true
	}

	return false
}

// SetProviderId gets a reference to the given string and assigns it to the ProviderId field.
func (o *RemedyMethod) SetProviderId(v string) {
	o.ProviderId = &v
}

func (o RemedyMethod) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["message"] = o.Message
	}
	if o.ClaimSuffix != nil {
		toSerialize["claimSuffix"] = o.ClaimSuffix
	}
	if o.ProviderId != nil {
		toSerialize["providerId"] = o.ProviderId
	}
	return json.Marshal(toSerialize)
}

type NullableRemedyMethod struct {
	value *RemedyMethod
	isSet bool
}

func (v NullableRemedyMethod) Get() *RemedyMethod {
	return v.value
}

func (v *NullableRemedyMethod) Set(val *RemedyMethod) {
	v.value = val
	v.isSet = true
}

func (v NullableRemedyMethod) IsSet() bool {
	return v.isSet
}

func (v *NullableRemedyMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRemedyMethod(val *RemedyMethod) *NullableRemedyMethod {
	return &NullableRemedyMethod{value: val, isSet: true}
}

func (v NullableRemedyMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRemedyMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
