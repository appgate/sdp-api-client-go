/*
Appgate SDP Controller REST API

# About   This specification documents the REST API calls for the Appgate SDP Controller.    Please refer to the REST API chapter in the manual or contact Appgate support with any questions about   this functionality. # Getting Started   Requirements for API scripting:   - Access to the Admin/API TLS Connection (default port 8443) of a Controller appliance.     (https://sdphelp.appgate.com/adminguide/appliance-function-configure.html?anchor=admin-api)   - An API user with relevant permissions.     (https://sdphelp.appgate.com/adminguide/administrative-roles-configure.html)   - In order to use the simple login API, Admin MFA must be disabled or the API user must be excluded.     (https://sdphelp.appgate.com/adminguide/mfa-for-admins.html) # Base path   HTTPS requests must be sent to the Admin Interface hostname and port, with **_/admin** path.    For example: **https://appgate.company.com:8443/admin**    All requests must have the **Accept** header as:    **application/vnd.appgate.peer-v16+json** # API Conventions   API conventions are  important to understand and follow strictly.    - While updating objects (via PUT), entire object must be sent with all fields.     - For example, in order to add a remedy method to the condition below:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": []       }       ```     - send the entire object with updated and non-updated fields:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": [{\"type\": \"DisplayMessage\", \"message\": \"test message\"}]       }       ```    - In case Controller returns an error (non-2xx HTTP status code), response body is JSON.     The \"message\" field contains information about the error.     HTTP 422 \"Unprocessable Entity\" has extra `errors` field to list all the issues with specific fields.    - Empty string (\"\") is considered a different value than \"null\" or field being omitted from JSON.     Omitting the field is recommend if no value is intended.     Empty string (\"\") will be almost always rejected as invalid value.    - There are common pattern between many objects:     - **Configuration Objects**: There are many objects with common fields, namely \"id\", \"name\", \"notes\", \"created\"       and \"updated\". These entities are listed, queried, created, updated and deleted in a similar fashion.     - **Distinguished Name**: Users and Devices are identified with what is called Distinguished Names, as used in        LDAP. The distinguished format that identifies a device and a user combination is        \"CN=\\<Device ID\\>,CN=\\<username\\>,OU=\\<Identity Provider Name\\>\". Some objects have the        \"userDistinguishedName\" field, which does not include the CN for Device ID.        This identifies a user on every device.

API version: API version 16.1
Contact: appgatesdp.support@appgate.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// ActiveSessionsAllOf struct for ActiveSessionsAllOf
type ActiveSessionsAllOf struct {
	// The number of total Distinct Users currently active.
	DistinctUserCount *float32         `json:"distinctUserCount,omitempty"`
	Data              *[]DeviceAndUser `json:"data,omitempty"`
}

// NewActiveSessionsAllOf instantiates a new ActiveSessionsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActiveSessionsAllOf() *ActiveSessionsAllOf {
	this := ActiveSessionsAllOf{}
	return &this
}

// NewActiveSessionsAllOfWithDefaults instantiates a new ActiveSessionsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActiveSessionsAllOfWithDefaults() *ActiveSessionsAllOf {
	this := ActiveSessionsAllOf{}
	return &this
}

// GetDistinctUserCount returns the DistinctUserCount field value if set, zero value otherwise.
func (o *ActiveSessionsAllOf) GetDistinctUserCount() float32 {
	if o == nil || o.DistinctUserCount == nil {
		var ret float32
		return ret
	}
	return *o.DistinctUserCount
}

// GetDistinctUserCountOk returns a tuple with the DistinctUserCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActiveSessionsAllOf) GetDistinctUserCountOk() (*float32, bool) {
	if o == nil || o.DistinctUserCount == nil {
		return nil, false
	}
	return o.DistinctUserCount, true
}

// HasDistinctUserCount returns a boolean if a field has been set.
func (o *ActiveSessionsAllOf) HasDistinctUserCount() bool {
	if o != nil && o.DistinctUserCount != nil {
		return true
	}

	return false
}

// SetDistinctUserCount gets a reference to the given float32 and assigns it to the DistinctUserCount field.
func (o *ActiveSessionsAllOf) SetDistinctUserCount(v float32) {
	o.DistinctUserCount = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ActiveSessionsAllOf) GetData() []DeviceAndUser {
	if o == nil || o.Data == nil {
		var ret []DeviceAndUser
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActiveSessionsAllOf) GetDataOk() (*[]DeviceAndUser, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ActiveSessionsAllOf) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []DeviceAndUser and assigns it to the Data field.
func (o *ActiveSessionsAllOf) SetData(v []DeviceAndUser) {
	o.Data = &v
}

func (o ActiveSessionsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DistinctUserCount != nil {
		toSerialize["distinctUserCount"] = o.DistinctUserCount
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableActiveSessionsAllOf struct {
	value *ActiveSessionsAllOf
	isSet bool
}

func (v NullableActiveSessionsAllOf) Get() *ActiveSessionsAllOf {
	return v.value
}

func (v *NullableActiveSessionsAllOf) Set(val *ActiveSessionsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableActiveSessionsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableActiveSessionsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActiveSessionsAllOf(val *ActiveSessionsAllOf) *NullableActiveSessionsAllOf {
	return &NullableActiveSessionsAllOf{value: val, isSet: true}
}

func (v NullableActiveSessionsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActiveSessionsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
