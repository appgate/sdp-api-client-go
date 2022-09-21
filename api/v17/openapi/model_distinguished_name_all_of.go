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

// DistinguishedNameAllOf Represents a user and device combination.
type DistinguishedNameAllOf struct {
	// The last time a Token issued to this user&device.
	LastTokenIssuedAt *string `json:"lastTokenIssuedAt,omitempty"`
	// The hostname recorded for the given user&device during On-Boarding. It may be empty if the Client cannot resolve or Token belongs to an Admin UI session.
	Hostname *string `json:"hostname,omitempty"`
}

// NewDistinguishedNameAllOf instantiates a new DistinguishedNameAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDistinguishedNameAllOf() *DistinguishedNameAllOf {
	this := DistinguishedNameAllOf{}
	return &this
}

// NewDistinguishedNameAllOfWithDefaults instantiates a new DistinguishedNameAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDistinguishedNameAllOfWithDefaults() *DistinguishedNameAllOf {
	this := DistinguishedNameAllOf{}
	return &this
}

// GetLastTokenIssuedAt returns the LastTokenIssuedAt field value if set, zero value otherwise.
func (o *DistinguishedNameAllOf) GetLastTokenIssuedAt() string {
	if o == nil || o.LastTokenIssuedAt == nil {
		var ret string
		return ret
	}
	return *o.LastTokenIssuedAt
}

// GetLastTokenIssuedAtOk returns a tuple with the LastTokenIssuedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DistinguishedNameAllOf) GetLastTokenIssuedAtOk() (*string, bool) {
	if o == nil || o.LastTokenIssuedAt == nil {
		return nil, false
	}
	return o.LastTokenIssuedAt, true
}

// HasLastTokenIssuedAt returns a boolean if a field has been set.
func (o *DistinguishedNameAllOf) HasLastTokenIssuedAt() bool {
	if o != nil && o.LastTokenIssuedAt != nil {
		return true
	}

	return false
}

// SetLastTokenIssuedAt gets a reference to the given string and assigns it to the LastTokenIssuedAt field.
func (o *DistinguishedNameAllOf) SetLastTokenIssuedAt(v string) {
	o.LastTokenIssuedAt = &v
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *DistinguishedNameAllOf) GetHostname() string {
	if o == nil || o.Hostname == nil {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DistinguishedNameAllOf) GetHostnameOk() (*string, bool) {
	if o == nil || o.Hostname == nil {
		return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *DistinguishedNameAllOf) HasHostname() bool {
	if o != nil && o.Hostname != nil {
		return true
	}

	return false
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *DistinguishedNameAllOf) SetHostname(v string) {
	o.Hostname = &v
}

func (o DistinguishedNameAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.LastTokenIssuedAt != nil {
		toSerialize["lastTokenIssuedAt"] = o.LastTokenIssuedAt
	}
	if o.Hostname != nil {
		toSerialize["hostname"] = o.Hostname
	}
	return json.Marshal(toSerialize)
}

type NullableDistinguishedNameAllOf struct {
	value *DistinguishedNameAllOf
	isSet bool
}

func (v NullableDistinguishedNameAllOf) Get() *DistinguishedNameAllOf {
	return v.value
}

func (v *NullableDistinguishedNameAllOf) Set(val *DistinguishedNameAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableDistinguishedNameAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableDistinguishedNameAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDistinguishedNameAllOf(val *DistinguishedNameAllOf) *NullableDistinguishedNameAllOf {
	return &NullableDistinguishedNameAllOf{value: val, isSet: true}
}

func (v NullableDistinguishedNameAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDistinguishedNameAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
