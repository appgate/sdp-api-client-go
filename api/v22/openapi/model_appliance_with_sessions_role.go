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

// ApplianceWithSessionsRole struct for ApplianceWithSessionsRole
type ApplianceWithSessionsRole struct {
	// Status of the role.
	Status *string `json:"status,omitempty"`
	// Optional details for extra information like error messages.
	Details *string `json:"details,omitempty"`
	// Number of active sessions on the Gateway or Portal.
	NumberOfSessions *int32 `json:"numberOfSessions,omitempty"`
}

// NewApplianceWithSessionsRole instantiates a new ApplianceWithSessionsRole object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplianceWithSessionsRole() *ApplianceWithSessionsRole {
	this := ApplianceWithSessionsRole{}
	return &this
}

// NewApplianceWithSessionsRoleWithDefaults instantiates a new ApplianceWithSessionsRole object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplianceWithSessionsRoleWithDefaults() *ApplianceWithSessionsRole {
	this := ApplianceWithSessionsRole{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ApplianceWithSessionsRole) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceWithSessionsRole) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ApplianceWithSessionsRole) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ApplianceWithSessionsRole) SetStatus(v string) {
	o.Status = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *ApplianceWithSessionsRole) GetDetails() string {
	if o == nil || o.Details == nil {
		var ret string
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceWithSessionsRole) GetDetailsOk() (*string, bool) {
	if o == nil || o.Details == nil {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *ApplianceWithSessionsRole) HasDetails() bool {
	if o != nil && o.Details != nil {
		return true
	}

	return false
}

// SetDetails gets a reference to the given string and assigns it to the Details field.
func (o *ApplianceWithSessionsRole) SetDetails(v string) {
	o.Details = &v
}

// GetNumberOfSessions returns the NumberOfSessions field value if set, zero value otherwise.
func (o *ApplianceWithSessionsRole) GetNumberOfSessions() int32 {
	if o == nil || o.NumberOfSessions == nil {
		var ret int32
		return ret
	}
	return *o.NumberOfSessions
}

// GetNumberOfSessionsOk returns a tuple with the NumberOfSessions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceWithSessionsRole) GetNumberOfSessionsOk() (*int32, bool) {
	if o == nil || o.NumberOfSessions == nil {
		return nil, false
	}
	return o.NumberOfSessions, true
}

// HasNumberOfSessions returns a boolean if a field has been set.
func (o *ApplianceWithSessionsRole) HasNumberOfSessions() bool {
	if o != nil && o.NumberOfSessions != nil {
		return true
	}

	return false
}

// SetNumberOfSessions gets a reference to the given int32 and assigns it to the NumberOfSessions field.
func (o *ApplianceWithSessionsRole) SetNumberOfSessions(v int32) {
	o.NumberOfSessions = &v
}

func (o ApplianceWithSessionsRole) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Details != nil {
		toSerialize["details"] = o.Details
	}
	if o.NumberOfSessions != nil {
		toSerialize["numberOfSessions"] = o.NumberOfSessions
	}
	return json.Marshal(toSerialize)
}

type NullableApplianceWithSessionsRole struct {
	value *ApplianceWithSessionsRole
	isSet bool
}

func (v NullableApplianceWithSessionsRole) Get() *ApplianceWithSessionsRole {
	return v.value
}

func (v *NullableApplianceWithSessionsRole) Set(val *ApplianceWithSessionsRole) {
	v.value = val
	v.isSet = true
}

func (v NullableApplianceWithSessionsRole) IsSet() bool {
	return v.isSet
}

func (v *NullableApplianceWithSessionsRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplianceWithSessionsRole(val *ApplianceWithSessionsRole) *NullableApplianceWithSessionsRole {
	return &NullableApplianceWithSessionsRole{value: val, isSet: true}
}

func (v NullableApplianceWithSessionsRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplianceWithSessionsRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
