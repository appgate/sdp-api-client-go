/*
Appgate SDP Controller REST API

# About   This specification documents the REST API calls for the Appgate SDP Controller.    Please refer to the REST API chapter in the manual or contact Appgate support with any questions about   this functionality. # Getting Started   Requirements for API scripting:   - Access to the Admin/API TLS Connection (default port 8443) of a Controller appliance.     (https://sdphelp.appgate.com/adminguide/appliance-function-configure.html?anchor=admin-api)   - An API user with relevant permissions.     (https://sdphelp.appgate.com/adminguide/administrative-roles-configure.html)   - In order to use the simple login API, Admin MFA must be disabled or the API user must be excluded.     (https://sdphelp.appgate.com/adminguide/mfa-for-admins.html) # Base path   HTTPS requests must be sent to the Admin Interface hostname and port, with **_/admin** path.    For example: **https://appgate.company.com:8443/admin**    All requests must have the **Accept** header as:    **application/vnd.appgate.peer-v17+json** # API Conventions   API conventions are  important to understand and follow strictly.    - While updating objects (via PUT), entire object must be sent with all fields.     - For example, in order to add a remedy method to the condition below:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": []       }       ```     - send the entire object with updated and non-updated fields:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": [{\"type\": \"DisplayMessage\", \"message\": \"test message\"}]       }       ```    - In case Controller returns an error (non-2xx HTTP status code), response body is JSON.     The \"message\" field contains information about the error.     HTTP 422 \"Unprocessable Entity\" has extra `errors` field to list all the issues with specific fields.    - Empty string (\"\") is considered a different value than \"null\" or field being omitted from JSON.     Omitting the field is recommended if no value is intended.     Empty string (\"\") will be almost always rejected as invalid value.    - There are common pattern between many objects:     - **Configuration Objects**: There are many objects with common fields, namely \"id\", \"name\", \"notes\", \"created\"       and \"updated\". These entities are listed, queried, created, updated and deleted in a similar fashion.     - **Distinguished Name**: Users and Devices are identified with what is called Distinguished Names, as used in        LDAP. The distinguished format that identifies a device and a user combination is        \"CN=\\<Device ID\\>,CN=\\<username\\>,OU=\\<Identity Provider Name\\>\". Some objects have the        \"userDistinguishedName\" field, which does not include the CN for Device ID.        This identifies a user on every device.

API version: API version 17.0
Contact: appgatesdp.support@appgate.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// OtpSeed struct for OtpSeed
type OtpSeed struct {
	// Distinguished name of a user. Format: \"CN=,OU=\"
	UserDistinguishedName *string `json:"userDistinguishedName,omitempty"`
	// The username, same as the one in the user Distinguished Name.
	Username *string `json:"username,omitempty"`
	// The provider name of the user, same as the one in the user Distinguished Name.
	ProviderName *string `json:"providerName,omitempty"`
	// Whether the generated Seed successfully used by the User at least once or not.
	Verified *bool `json:"verified,omitempty"`
}

// NewOtpSeed instantiates a new OtpSeed object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOtpSeed() *OtpSeed {
	this := OtpSeed{}
	return &this
}

// NewOtpSeedWithDefaults instantiates a new OtpSeed object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOtpSeedWithDefaults() *OtpSeed {
	this := OtpSeed{}
	return &this
}

// GetUserDistinguishedName returns the UserDistinguishedName field value if set, zero value otherwise.
func (o *OtpSeed) GetUserDistinguishedName() string {
	if o == nil || o.UserDistinguishedName == nil {
		var ret string
		return ret
	}
	return *o.UserDistinguishedName
}

// GetUserDistinguishedNameOk returns a tuple with the UserDistinguishedName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OtpSeed) GetUserDistinguishedNameOk() (*string, bool) {
	if o == nil || o.UserDistinguishedName == nil {
		return nil, false
	}
	return o.UserDistinguishedName, true
}

// HasUserDistinguishedName returns a boolean if a field has been set.
func (o *OtpSeed) HasUserDistinguishedName() bool {
	if o != nil && o.UserDistinguishedName != nil {
		return true
	}

	return false
}

// SetUserDistinguishedName gets a reference to the given string and assigns it to the UserDistinguishedName field.
func (o *OtpSeed) SetUserDistinguishedName(v string) {
	o.UserDistinguishedName = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *OtpSeed) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OtpSeed) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *OtpSeed) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *OtpSeed) SetUsername(v string) {
	o.Username = &v
}

// GetProviderName returns the ProviderName field value if set, zero value otherwise.
func (o *OtpSeed) GetProviderName() string {
	if o == nil || o.ProviderName == nil {
		var ret string
		return ret
	}
	return *o.ProviderName
}

// GetProviderNameOk returns a tuple with the ProviderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OtpSeed) GetProviderNameOk() (*string, bool) {
	if o == nil || o.ProviderName == nil {
		return nil, false
	}
	return o.ProviderName, true
}

// HasProviderName returns a boolean if a field has been set.
func (o *OtpSeed) HasProviderName() bool {
	if o != nil && o.ProviderName != nil {
		return true
	}

	return false
}

// SetProviderName gets a reference to the given string and assigns it to the ProviderName field.
func (o *OtpSeed) SetProviderName(v string) {
	o.ProviderName = &v
}

// GetVerified returns the Verified field value if set, zero value otherwise.
func (o *OtpSeed) GetVerified() bool {
	if o == nil || o.Verified == nil {
		var ret bool
		return ret
	}
	return *o.Verified
}

// GetVerifiedOk returns a tuple with the Verified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OtpSeed) GetVerifiedOk() (*bool, bool) {
	if o == nil || o.Verified == nil {
		return nil, false
	}
	return o.Verified, true
}

// HasVerified returns a boolean if a field has been set.
func (o *OtpSeed) HasVerified() bool {
	if o != nil && o.Verified != nil {
		return true
	}

	return false
}

// SetVerified gets a reference to the given bool and assigns it to the Verified field.
func (o *OtpSeed) SetVerified(v bool) {
	o.Verified = &v
}

func (o OtpSeed) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UserDistinguishedName != nil {
		toSerialize["userDistinguishedName"] = o.UserDistinguishedName
	}
	if o.Username != nil {
		toSerialize["username"] = o.Username
	}
	if o.ProviderName != nil {
		toSerialize["providerName"] = o.ProviderName
	}
	if o.Verified != nil {
		toSerialize["verified"] = o.Verified
	}
	return json.Marshal(toSerialize)
}

type NullableOtpSeed struct {
	value *OtpSeed
	isSet bool
}

func (v NullableOtpSeed) Get() *OtpSeed {
	return v.value
}

func (v *NullableOtpSeed) Set(val *OtpSeed) {
	v.value = val
	v.isSet = true
}

func (v NullableOtpSeed) IsSet() bool {
	return v.isSet
}

func (v *NullableOtpSeed) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOtpSeed(val *OtpSeed) *NullableOtpSeed {
	return &NullableOtpSeed{value: val, isSet: true}
}

func (v NullableOtpSeed) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOtpSeed) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
