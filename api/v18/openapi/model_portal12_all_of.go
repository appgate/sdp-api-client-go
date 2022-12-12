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

// Portal12AllOf struct for Portal12AllOf
type Portal12AllOf struct {
	// Portal will verify upstream certificate of the endpoints.
	VerifyUpstream *bool `json:"verifyUpstream,omitempty"`
}

// NewPortal12AllOf instantiates a new Portal12AllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPortal12AllOf() *Portal12AllOf {
	this := Portal12AllOf{}
	var verifyUpstream bool = true
	this.VerifyUpstream = &verifyUpstream
	return &this
}

// NewPortal12AllOfWithDefaults instantiates a new Portal12AllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPortal12AllOfWithDefaults() *Portal12AllOf {
	this := Portal12AllOf{}
	var verifyUpstream bool = true
	this.VerifyUpstream = &verifyUpstream
	return &this
}

// GetVerifyUpstream returns the VerifyUpstream field value if set, zero value otherwise.
func (o *Portal12AllOf) GetVerifyUpstream() bool {
	if o == nil || o.VerifyUpstream == nil {
		var ret bool
		return ret
	}
	return *o.VerifyUpstream
}

// GetVerifyUpstreamOk returns a tuple with the VerifyUpstream field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Portal12AllOf) GetVerifyUpstreamOk() (*bool, bool) {
	if o == nil || o.VerifyUpstream == nil {
		return nil, false
	}
	return o.VerifyUpstream, true
}

// HasVerifyUpstream returns a boolean if a field has been set.
func (o *Portal12AllOf) HasVerifyUpstream() bool {
	if o != nil && o.VerifyUpstream != nil {
		return true
	}

	return false
}

// SetVerifyUpstream gets a reference to the given bool and assigns it to the VerifyUpstream field.
func (o *Portal12AllOf) SetVerifyUpstream(v bool) {
	o.VerifyUpstream = &v
}

func (o Portal12AllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.VerifyUpstream != nil {
		toSerialize["verifyUpstream"] = o.VerifyUpstream
	}
	return json.Marshal(toSerialize)
}

type NullablePortal12AllOf struct {
	value *Portal12AllOf
	isSet bool
}

func (v NullablePortal12AllOf) Get() *Portal12AllOf {
	return v.value
}

func (v *NullablePortal12AllOf) Set(val *Portal12AllOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePortal12AllOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePortal12AllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePortal12AllOf(val *Portal12AllOf) *NullablePortal12AllOf {
	return &NullablePortal12AllOf{value: val, isSet: true}
}

func (v NullablePortal12AllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePortal12AllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
