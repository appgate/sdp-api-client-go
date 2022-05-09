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

// P121 struct for P121
type P121 struct {
	// P12 binary in Base64 format.
	P12 *string `json:"p12,omitempty"`
	// Password for the p12 file.
	Password *string `json:"password,omitempty"`
}

// NewP121 instantiates a new P121 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewP121() *P121 {
	this := P121{}
	var password string = ""
	this.Password = &password
	return &this
}

// NewP121WithDefaults instantiates a new P121 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewP121WithDefaults() *P121 {
	this := P121{}
	var password string = ""
	this.Password = &password
	return &this
}

// GetP12 returns the P12 field value if set, zero value otherwise.
func (o *P121) GetP12() string {
	if o == nil || o.P12 == nil {
		var ret string
		return ret
	}
	return *o.P12
}

// GetP12Ok returns a tuple with the P12 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *P121) GetP12Ok() (*string, bool) {
	if o == nil || o.P12 == nil {
		return nil, false
	}
	return o.P12, true
}

// HasP12 returns a boolean if a field has been set.
func (o *P121) HasP12() bool {
	if o != nil && o.P12 != nil {
		return true
	}

	return false
}

// SetP12 gets a reference to the given string and assigns it to the P12 field.
func (o *P121) SetP12(v string) {
	o.P12 = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *P121) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *P121) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *P121) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *P121) SetPassword(v string) {
	o.Password = &v
}

func (o P121) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.P12 != nil {
		toSerialize["p12"] = o.P12
	}
	if o.Password != nil {
		toSerialize["password"] = o.Password
	}
	return json.Marshal(toSerialize)
}

type NullableP121 struct {
	value *P121
	isSet bool
}

func (v NullableP121) Get() *P121 {
	return v.value
}

func (v *NullableP121) Set(val *P121) {
	v.value = val
	v.isSet = true
}

func (v NullableP121) IsSet() bool {
	return v.isSet
}

func (v *NullableP121) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableP121(val *P121) *NullableP121 {
	return &NullableP121{value: val, isSet: true}
}

func (v NullableP121) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableP121) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
