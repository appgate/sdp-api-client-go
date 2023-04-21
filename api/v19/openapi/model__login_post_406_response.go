/*
Appgate SDP Controller REST API

# About   This specification documents the REST API calls for the Appgate SDP Controller.    Please refer to the REST API chapter in the manual or contact Appgate support with any questions about   this functionality. # Getting Started   Requirements for API scripting:   - Access to the Admin/API TLS Connection (default port 8443) of a Controller appliance.     (https://sdphelp.appgate.com/adminguide/appliance-function-configure.html?anchor=admin-api)   - An API user with relevant permissions.     (https://sdphelp.appgate.com/adminguide/administrative-roles-configure.html)   - In order to use the simple login API, Admin MFA must be disabled or the API user must be excluded.     (https://sdphelp.appgate.com/adminguide/mfa-for-admins.html) # Base path   HTTPS requests must be sent to the Admin Interface hostname and port, with **_/admin** path.    For example: **https://appgate.company.com:8443/admin**    All requests must have the **Accept** header as:    **application/vnd.appgate.peer-v19+json** # API Conventions   API conventions are  important to understand and follow strictly.    - While updating objects (via PUT), entire object must be sent with all fields.     - For example, in order to add a remedy method to the condition below:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": []       }       ```     - send the entire object with updated and non-updated fields:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": [{\"type\": \"DisplayMessage\", \"message\": \"test message\"}]       }       ```    - In case Controller returns an error (non-2xx HTTP status code), response body is JSON.     The \"message\" field contains information about the error.     HTTP 422 \"Unprocessable Entity\" has extra `errors` field to list all the issues with specific fields.    - Empty string (\"\") is considered a different value than \"null\" or field being omitted from JSON.     Omitting the field is recommended if no value is intended.     Empty string (\"\") will be almost always rejected as invalid value.    - There are common pattern between many objects:     - **Configuration Objects**: There are many objects with common fields, namely \"id\", \"name\", \"notes\", \"created\"       and \"updated\". These entities are listed, queried, created, updated and deleted in a similar fashion.     - **Distinguished Name**: Users and Devices are identified with what is called Distinguished Names, as used in        LDAP. The distinguished format that identifies a device and a user combination is        \"CN=\\<Device ID\\>,CN=\\<username\\>,OU=\\<Identity Provider Name\\>\". Some objects have the        \"userDistinguishedName\" field, which does not include the CN for Device ID.        This identifies a user on every device.

API version: API version 19.0
Contact: appgatesdp.support@appgate.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// LoginPost406Response struct for LoginPost406Response
type LoginPost406Response struct {
	// N/A
	Id *string `json:"id,omitempty"`
	// Human readable error details.
	Message *string `json:"message,omitempty"`
	// Minimum supported version.
	MinSupportedVersion *int32 `json:"minSupportedVersion,omitempty"`
	// Maximum supported version.
	MaxSupportedVersion *int32 `json:"maxSupportedVersion,omitempty"`
}

// NewLoginPost406Response instantiates a new LoginPost406Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoginPost406Response() *LoginPost406Response {
	this := LoginPost406Response{}
	return &this
}

// NewLoginPost406ResponseWithDefaults instantiates a new LoginPost406Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoginPost406ResponseWithDefaults() *LoginPost406Response {
	this := LoginPost406Response{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *LoginPost406Response) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoginPost406Response) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *LoginPost406Response) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *LoginPost406Response) SetId(v string) {
	o.Id = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *LoginPost406Response) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoginPost406Response) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *LoginPost406Response) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *LoginPost406Response) SetMessage(v string) {
	o.Message = &v
}

// GetMinSupportedVersion returns the MinSupportedVersion field value if set, zero value otherwise.
func (o *LoginPost406Response) GetMinSupportedVersion() int32 {
	if o == nil || o.MinSupportedVersion == nil {
		var ret int32
		return ret
	}
	return *o.MinSupportedVersion
}

// GetMinSupportedVersionOk returns a tuple with the MinSupportedVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoginPost406Response) GetMinSupportedVersionOk() (*int32, bool) {
	if o == nil || o.MinSupportedVersion == nil {
		return nil, false
	}
	return o.MinSupportedVersion, true
}

// HasMinSupportedVersion returns a boolean if a field has been set.
func (o *LoginPost406Response) HasMinSupportedVersion() bool {
	if o != nil && o.MinSupportedVersion != nil {
		return true
	}

	return false
}

// SetMinSupportedVersion gets a reference to the given int32 and assigns it to the MinSupportedVersion field.
func (o *LoginPost406Response) SetMinSupportedVersion(v int32) {
	o.MinSupportedVersion = &v
}

// GetMaxSupportedVersion returns the MaxSupportedVersion field value if set, zero value otherwise.
func (o *LoginPost406Response) GetMaxSupportedVersion() int32 {
	if o == nil || o.MaxSupportedVersion == nil {
		var ret int32
		return ret
	}
	return *o.MaxSupportedVersion
}

// GetMaxSupportedVersionOk returns a tuple with the MaxSupportedVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoginPost406Response) GetMaxSupportedVersionOk() (*int32, bool) {
	if o == nil || o.MaxSupportedVersion == nil {
		return nil, false
	}
	return o.MaxSupportedVersion, true
}

// HasMaxSupportedVersion returns a boolean if a field has been set.
func (o *LoginPost406Response) HasMaxSupportedVersion() bool {
	if o != nil && o.MaxSupportedVersion != nil {
		return true
	}

	return false
}

// SetMaxSupportedVersion gets a reference to the given int32 and assigns it to the MaxSupportedVersion field.
func (o *LoginPost406Response) SetMaxSupportedVersion(v int32) {
	o.MaxSupportedVersion = &v
}

func (o LoginPost406Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.MinSupportedVersion != nil {
		toSerialize["minSupportedVersion"] = o.MinSupportedVersion
	}
	if o.MaxSupportedVersion != nil {
		toSerialize["maxSupportedVersion"] = o.MaxSupportedVersion
	}
	return json.Marshal(toSerialize)
}

type NullableLoginPost406Response struct {
	value *LoginPost406Response
	isSet bool
}

func (v NullableLoginPost406Response) Get() *LoginPost406Response {
	return v.value
}

func (v *NullableLoginPost406Response) Set(val *LoginPost406Response) {
	v.value = val
	v.isSet = true
}

func (v NullableLoginPost406Response) IsSet() bool {
	return v.isSet
}

func (v *NullableLoginPost406Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoginPost406Response(val *LoginPost406Response) *NullableLoginPost406Response {
	return &NullableLoginPost406Response{value: val, isSet: true}
}

func (v NullableLoginPost406Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoginPost406Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
