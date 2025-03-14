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

// TokenRevocationRequest Details for token revocation request.
type TokenRevocationRequest struct {
	// Optional reason text for the revocation. The value is stored and logged.
	RevocationReason *string `json:"revocationReason,omitempty"`
	// The delay time for client token revocation in minutes. Client will renew the token(s) at least 5 minutes before the revocation time, without losing connection.
	DelayMinutes *int32 `json:"delayMinutes,omitempty"`
	// Only used when revoking all Tokens. In order to spread the workload on the Controllers, tokens are revoked in batches according to this value.
	TokensPerSecond *float32 `json:"tokensPerSecond,omitempty"`
	// Only used when revoking all Tokens. Specific distinguished names can be defined to renew tokens in bulk for a specific list of devices.
	SpecificDistinguishedNames []string `json:"specificDistinguishedNames,omitempty"`
}

// NewTokenRevocationRequest instantiates a new TokenRevocationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenRevocationRequest() *TokenRevocationRequest {
	this := TokenRevocationRequest{}
	var delayMinutes int32 = 5
	this.DelayMinutes = &delayMinutes
	var tokensPerSecond float32 = 7
	this.TokensPerSecond = &tokensPerSecond
	return &this
}

// NewTokenRevocationRequestWithDefaults instantiates a new TokenRevocationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenRevocationRequestWithDefaults() *TokenRevocationRequest {
	this := TokenRevocationRequest{}
	var delayMinutes int32 = 5
	this.DelayMinutes = &delayMinutes
	var tokensPerSecond float32 = 7
	this.TokensPerSecond = &tokensPerSecond
	return &this
}

// GetRevocationReason returns the RevocationReason field value if set, zero value otherwise.
func (o *TokenRevocationRequest) GetRevocationReason() string {
	if o == nil || o.RevocationReason == nil {
		var ret string
		return ret
	}
	return *o.RevocationReason
}

// GetRevocationReasonOk returns a tuple with the RevocationReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenRevocationRequest) GetRevocationReasonOk() (*string, bool) {
	if o == nil || o.RevocationReason == nil {
		return nil, false
	}
	return o.RevocationReason, true
}

// HasRevocationReason returns a boolean if a field has been set.
func (o *TokenRevocationRequest) HasRevocationReason() bool {
	if o != nil && o.RevocationReason != nil {
		return true
	}

	return false
}

// SetRevocationReason gets a reference to the given string and assigns it to the RevocationReason field.
func (o *TokenRevocationRequest) SetRevocationReason(v string) {
	o.RevocationReason = &v
}

// GetDelayMinutes returns the DelayMinutes field value if set, zero value otherwise.
func (o *TokenRevocationRequest) GetDelayMinutes() int32 {
	if o == nil || o.DelayMinutes == nil {
		var ret int32
		return ret
	}
	return *o.DelayMinutes
}

// GetDelayMinutesOk returns a tuple with the DelayMinutes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenRevocationRequest) GetDelayMinutesOk() (*int32, bool) {
	if o == nil || o.DelayMinutes == nil {
		return nil, false
	}
	return o.DelayMinutes, true
}

// HasDelayMinutes returns a boolean if a field has been set.
func (o *TokenRevocationRequest) HasDelayMinutes() bool {
	if o != nil && o.DelayMinutes != nil {
		return true
	}

	return false
}

// SetDelayMinutes gets a reference to the given int32 and assigns it to the DelayMinutes field.
func (o *TokenRevocationRequest) SetDelayMinutes(v int32) {
	o.DelayMinutes = &v
}

// GetTokensPerSecond returns the TokensPerSecond field value if set, zero value otherwise.
func (o *TokenRevocationRequest) GetTokensPerSecond() float32 {
	if o == nil || o.TokensPerSecond == nil {
		var ret float32
		return ret
	}
	return *o.TokensPerSecond
}

// GetTokensPerSecondOk returns a tuple with the TokensPerSecond field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenRevocationRequest) GetTokensPerSecondOk() (*float32, bool) {
	if o == nil || o.TokensPerSecond == nil {
		return nil, false
	}
	return o.TokensPerSecond, true
}

// HasTokensPerSecond returns a boolean if a field has been set.
func (o *TokenRevocationRequest) HasTokensPerSecond() bool {
	if o != nil && o.TokensPerSecond != nil {
		return true
	}

	return false
}

// SetTokensPerSecond gets a reference to the given float32 and assigns it to the TokensPerSecond field.
func (o *TokenRevocationRequest) SetTokensPerSecond(v float32) {
	o.TokensPerSecond = &v
}

// GetSpecificDistinguishedNames returns the SpecificDistinguishedNames field value if set, zero value otherwise.
func (o *TokenRevocationRequest) GetSpecificDistinguishedNames() []string {
	if o == nil || o.SpecificDistinguishedNames == nil {
		var ret []string
		return ret
	}
	return o.SpecificDistinguishedNames
}

// GetSpecificDistinguishedNamesOk returns a tuple with the SpecificDistinguishedNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenRevocationRequest) GetSpecificDistinguishedNamesOk() ([]string, bool) {
	if o == nil || o.SpecificDistinguishedNames == nil {
		return nil, false
	}
	return o.SpecificDistinguishedNames, true
}

// HasSpecificDistinguishedNames returns a boolean if a field has been set.
func (o *TokenRevocationRequest) HasSpecificDistinguishedNames() bool {
	if o != nil && o.SpecificDistinguishedNames != nil {
		return true
	}

	return false
}

// SetSpecificDistinguishedNames gets a reference to the given []string and assigns it to the SpecificDistinguishedNames field.
func (o *TokenRevocationRequest) SetSpecificDistinguishedNames(v []string) {
	o.SpecificDistinguishedNames = v
}

func (o TokenRevocationRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RevocationReason != nil {
		toSerialize["revocationReason"] = o.RevocationReason
	}
	if o.DelayMinutes != nil {
		toSerialize["delayMinutes"] = o.DelayMinutes
	}
	if o.TokensPerSecond != nil {
		toSerialize["tokensPerSecond"] = o.TokensPerSecond
	}
	if o.SpecificDistinguishedNames != nil {
		toSerialize["specificDistinguishedNames"] = o.SpecificDistinguishedNames
	}
	return json.Marshal(toSerialize)
}

type NullableTokenRevocationRequest struct {
	value *TokenRevocationRequest
	isSet bool
}

func (v NullableTokenRevocationRequest) Get() *TokenRevocationRequest {
	return v.value
}

func (v *NullableTokenRevocationRequest) Set(val *TokenRevocationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenRevocationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenRevocationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenRevocationRequest(val *TokenRevocationRequest) *NullableTokenRevocationRequest {
	return &NullableTokenRevocationRequest{value: val, isSet: true}
}

func (v NullableTokenRevocationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenRevocationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
