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

// ColorCode Represents a Color Code.
type ColorCode struct {
	// ID of the color code.
	Id *float32 `json:"id,omitempty"`
	// Display name of the color.
	ColorName *string `json:"colorName,omitempty"`
}

// NewColorCode instantiates a new ColorCode object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewColorCode() *ColorCode {
	this := ColorCode{}
	return &this
}

// NewColorCodeWithDefaults instantiates a new ColorCode object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewColorCodeWithDefaults() *ColorCode {
	this := ColorCode{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ColorCode) GetId() float32 {
	if o == nil || o.Id == nil {
		var ret float32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ColorCode) GetIdOk() (*float32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ColorCode) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given float32 and assigns it to the Id field.
func (o *ColorCode) SetId(v float32) {
	o.Id = &v
}

// GetColorName returns the ColorName field value if set, zero value otherwise.
func (o *ColorCode) GetColorName() string {
	if o == nil || o.ColorName == nil {
		var ret string
		return ret
	}
	return *o.ColorName
}

// GetColorNameOk returns a tuple with the ColorName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ColorCode) GetColorNameOk() (*string, bool) {
	if o == nil || o.ColorName == nil {
		return nil, false
	}
	return o.ColorName, true
}

// HasColorName returns a boolean if a field has been set.
func (o *ColorCode) HasColorName() bool {
	if o != nil && o.ColorName != nil {
		return true
	}

	return false
}

// SetColorName gets a reference to the given string and assigns it to the ColorName field.
func (o *ColorCode) SetColorName(v string) {
	o.ColorName = &v
}

func (o ColorCode) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.ColorName != nil {
		toSerialize["colorName"] = o.ColorName
	}
	return json.Marshal(toSerialize)
}

type NullableColorCode struct {
	value *ColorCode
	isSet bool
}

func (v NullableColorCode) Get() *ColorCode {
	return v.value
}

func (v *NullableColorCode) Set(val *ColorCode) {
	v.value = val
	v.isSet = true
}

func (v NullableColorCode) IsSet() bool {
	return v.isSet
}

func (v *NullableColorCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableColorCode(val *ColorCode) *NullableColorCode {
	return &NullableColorCode{value: val, isSet: true}
}

func (v NullableColorCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableColorCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
