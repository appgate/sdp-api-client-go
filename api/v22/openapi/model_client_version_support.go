/*
Appgate SDP Controller REST API

# About   This specification documents the REST API calls for the Appgate SDP Controller.    Please refer to the REST API chapter in the manual or contact Appgate support with any questions about   this functionality. # Getting Started   Requirements for API scripting:   - Access to the Admin/API TLS Connection (default port 8443) of a Controller appliance.     (https://sdphelp.appgate.com/adminguide/appliance-function-configure.html?anchor=admin-api)   - An API user with relevant permissions.     (https://sdphelp.appgate.com/adminguide/administrative-roles-configure.html)   - In order to use the simple login API, Admin MFA must be disabled or the API user must be excluded.     (https://sdphelp.appgate.com/adminguide/mfa-for-admins.html) # Base path   HTTPS requests must be sent to the Admin Interface hostname and port, with **_/admin** path.    For example: **https://appgate.company.com:8443/admin**    All requests must have the **Accept** header as:    **application/vnd.appgate.peer-v22+json**    An exception is made for the **_/admin/version** endpoint which instead expects an **application/json** Accept header. # API Conventions   API conventions are  important to understand and follow strictly.    - While updating objects (via PUT), entire object must be sent with all fields.     - For example, in order to add a remedy method to the condition below:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": []       }       ```     - send the entire object with updated and non-updated fields:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": [{\"type\": \"DisplayMessage\", \"message\": \"test message\"}]       }       ```    - In case Controller returns an error (non-2xx HTTP status code), response body is JSON.     The \"message\" field contains information about the error.     HTTP 422 \"Unprocessable Entity\" has extra `errors` field to list all the issues with specific fields.    - Empty string (\"\") is considered a different value than \"null\" or field being omitted from JSON.     Omitting the field is recommended if no value is intended.     Empty string (\"\") will be almost always rejected as invalid value.    - There are common pattern between many objects:     - **Configuration Objects**: There are many objects with common fields, namely \"id\", \"name\", \"notes\", \"created\"       and \"updated\". These entities are listed, queried, created, updated and deleted in a similar fashion.     - **Distinguished Name**: Users and Devices are identified with what is called Distinguished Names, as used in        LDAP. The distinguished format that identifies a device and a user combination is        \"CN=\\<Device ID\\>,CN=\\<username\\>,OU=\\<Identity Provider Name\\>\". Some objects have the        \"userDistinguishedName\" field, which does not include the CN for Device ID.        This identifies a user on every device.

API version: API version 22.0
Contact: appgatesdp.support@appgate.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// ClientVersionSupport Information about the client versions in use.
type ClientVersionSupport struct {
	// The support status of the client version.
	SupportStatus *string `json:"supportStatus,omitempty"`
	// The count of sessions using this client version.
	SessionCount *int32 `json:"sessionCount,omitempty"`
}

// NewClientVersionSupport instantiates a new ClientVersionSupport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClientVersionSupport() *ClientVersionSupport {
	this := ClientVersionSupport{}
	return &this
}

// NewClientVersionSupportWithDefaults instantiates a new ClientVersionSupport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClientVersionSupportWithDefaults() *ClientVersionSupport {
	this := ClientVersionSupport{}
	return &this
}

// GetSupportStatus returns the SupportStatus field value if set, zero value otherwise.
func (o *ClientVersionSupport) GetSupportStatus() string {
	if o == nil || o.SupportStatus == nil {
		var ret string
		return ret
	}
	return *o.SupportStatus
}

// GetSupportStatusOk returns a tuple with the SupportStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientVersionSupport) GetSupportStatusOk() (*string, bool) {
	if o == nil || o.SupportStatus == nil {
		return nil, false
	}
	return o.SupportStatus, true
}

// HasSupportStatus returns a boolean if a field has been set.
func (o *ClientVersionSupport) HasSupportStatus() bool {
	if o != nil && o.SupportStatus != nil {
		return true
	}

	return false
}

// SetSupportStatus gets a reference to the given string and assigns it to the SupportStatus field.
func (o *ClientVersionSupport) SetSupportStatus(v string) {
	o.SupportStatus = &v
}

// GetSessionCount returns the SessionCount field value if set, zero value otherwise.
func (o *ClientVersionSupport) GetSessionCount() int32 {
	if o == nil || o.SessionCount == nil {
		var ret int32
		return ret
	}
	return *o.SessionCount
}

// GetSessionCountOk returns a tuple with the SessionCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientVersionSupport) GetSessionCountOk() (*int32, bool) {
	if o == nil || o.SessionCount == nil {
		return nil, false
	}
	return o.SessionCount, true
}

// HasSessionCount returns a boolean if a field has been set.
func (o *ClientVersionSupport) HasSessionCount() bool {
	if o != nil && o.SessionCount != nil {
		return true
	}

	return false
}

// SetSessionCount gets a reference to the given int32 and assigns it to the SessionCount field.
func (o *ClientVersionSupport) SetSessionCount(v int32) {
	o.SessionCount = &v
}

func (o ClientVersionSupport) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SupportStatus != nil {
		toSerialize["supportStatus"] = o.SupportStatus
	}
	if o.SessionCount != nil {
		toSerialize["sessionCount"] = o.SessionCount
	}
	return json.Marshal(toSerialize)
}

type NullableClientVersionSupport struct {
	value *ClientVersionSupport
	isSet bool
}

func (v NullableClientVersionSupport) Get() *ClientVersionSupport {
	return v.value
}

func (v *NullableClientVersionSupport) Set(val *ClientVersionSupport) {
	v.value = val
	v.isSet = true
}

func (v NullableClientVersionSupport) IsSet() bool {
	return v.isSet
}

func (v *NullableClientVersionSupport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClientVersionSupport(val *ClientVersionSupport) *NullableClientVersionSupport {
	return &NullableClientVersionSupport{value: val, isSet: true}
}

func (v NullableClientVersionSupport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClientVersionSupport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
