/*
Appgate SDP Controller REST API

# About   This specification documents the REST API calls for the Appgate SDP Controller.    Please refer to the REST API chapter in the manual or contact Appgate support with any questions about   this functionality. # Getting Started   Requirements for API scripting:   - Access to the Admin/API TLS Connection (default port 8443) of a Controller appliance.     (https://sdphelp.appgate.com/adminguide/appliance-function-configure.html?anchor=admin-api)   - An API user with relevant permissions.     (https://sdphelp.appgate.com/adminguide/administrative-roles-configure.html)   - In order to use the simple login API, Admin MFA must be disabled or the API user must be excluded.     (https://sdphelp.appgate.com/adminguide/mfa-for-admins.html) # Base path   HTTPS requests must be sent to the Admin Interface hostname and port, with **_/admin** path.    For example: **https://appgate.company.com:8443/admin**    All requests must have the **Accept** header as:    **application/vnd.appgate.peer-v22+json**    An exception is made for the **_/admin/version** endpoint which instead expects an **application/json** Accept header. # API Conventions   API conventions are  important to understand and follow strictly.    - While updating objects (via PUT), entire object must be sent with all fields.     - For example, in order to add a remedy method to the condition below:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": []       }       ```     - send the entire object with updated and non-updated fields:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": [{\"type\": \"DisplayMessage\", \"message\": \"test message\"}]       }       ```    - In case Controller returns an error (non-2xx HTTP status code), response body is JSON.     The \"message\" field contains information about the error.     HTTP 422 \"Unprocessable Entity\" has extra `errors` field to list all the issues with specific fields.    - Empty string (\"\") is considered a different value than \"null\" or field being omitted from JSON.     Omitting the field is recommended if no value is intended.     Empty string (\"\") will be almost always rejected as invalid value.    - There are common pattern between many objects:     - **Configuration Objects**: There are many objects with common fields, namely \"id\", \"name\", \"notes\", \"created\"       and \"updated\". These entities are listed, queried, created, updated and deleted in a similar fashion.     - **Distinguished Name**: Users and Devices are identified with what is called Distinguished Names, as used in        LDAP. The distinguished format that identifies a device and a user combination is        \"CN=\\<Device ID\\>,CN=\\<username\\>,OU=\\<Identity Provider Name\\>\". Some objects have the        \"userDistinguishedName\" field, which does not include the CN for Device ID.        This identifies a user on every device.

API version: API version 22.1
Contact: appgatesdp.support@appgate.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// AppliedBy struct for AppliedBy
type AppliedBy struct {
	// The distinguished name of who updated the message metadata.
	AppliedBy *string `json:"appliedBy,omitempty"`
	// The username of the user who updated the message metadata.
	AppliedByUsername *string `json:"appliedByUsername,omitempty"`
	// The provider name of the user who updated the message metadata.
	AppliedByProviderName *string `json:"appliedByProviderName,omitempty"`
	// The device ID of the user who updated the message metadata.
	AppliedByDeviceId *string `json:"appliedByDeviceId,omitempty"`
}

// NewAppliedBy instantiates a new AppliedBy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppliedBy() *AppliedBy {
	this := AppliedBy{}
	return &this
}

// NewAppliedByWithDefaults instantiates a new AppliedBy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppliedByWithDefaults() *AppliedBy {
	this := AppliedBy{}
	return &this
}

// GetAppliedBy returns the AppliedBy field value if set, zero value otherwise.
func (o *AppliedBy) GetAppliedBy() string {
	if o == nil || o.AppliedBy == nil {
		var ret string
		return ret
	}
	return *o.AppliedBy
}

// GetAppliedByOk returns a tuple with the AppliedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppliedBy) GetAppliedByOk() (*string, bool) {
	if o == nil || o.AppliedBy == nil {
		return nil, false
	}
	return o.AppliedBy, true
}

// HasAppliedBy returns a boolean if a field has been set.
func (o *AppliedBy) HasAppliedBy() bool {
	if o != nil && o.AppliedBy != nil {
		return true
	}

	return false
}

// SetAppliedBy gets a reference to the given string and assigns it to the AppliedBy field.
func (o *AppliedBy) SetAppliedBy(v string) {
	o.AppliedBy = &v
}

// GetAppliedByUsername returns the AppliedByUsername field value if set, zero value otherwise.
func (o *AppliedBy) GetAppliedByUsername() string {
	if o == nil || o.AppliedByUsername == nil {
		var ret string
		return ret
	}
	return *o.AppliedByUsername
}

// GetAppliedByUsernameOk returns a tuple with the AppliedByUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppliedBy) GetAppliedByUsernameOk() (*string, bool) {
	if o == nil || o.AppliedByUsername == nil {
		return nil, false
	}
	return o.AppliedByUsername, true
}

// HasAppliedByUsername returns a boolean if a field has been set.
func (o *AppliedBy) HasAppliedByUsername() bool {
	if o != nil && o.AppliedByUsername != nil {
		return true
	}

	return false
}

// SetAppliedByUsername gets a reference to the given string and assigns it to the AppliedByUsername field.
func (o *AppliedBy) SetAppliedByUsername(v string) {
	o.AppliedByUsername = &v
}

// GetAppliedByProviderName returns the AppliedByProviderName field value if set, zero value otherwise.
func (o *AppliedBy) GetAppliedByProviderName() string {
	if o == nil || o.AppliedByProviderName == nil {
		var ret string
		return ret
	}
	return *o.AppliedByProviderName
}

// GetAppliedByProviderNameOk returns a tuple with the AppliedByProviderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppliedBy) GetAppliedByProviderNameOk() (*string, bool) {
	if o == nil || o.AppliedByProviderName == nil {
		return nil, false
	}
	return o.AppliedByProviderName, true
}

// HasAppliedByProviderName returns a boolean if a field has been set.
func (o *AppliedBy) HasAppliedByProviderName() bool {
	if o != nil && o.AppliedByProviderName != nil {
		return true
	}

	return false
}

// SetAppliedByProviderName gets a reference to the given string and assigns it to the AppliedByProviderName field.
func (o *AppliedBy) SetAppliedByProviderName(v string) {
	o.AppliedByProviderName = &v
}

// GetAppliedByDeviceId returns the AppliedByDeviceId field value if set, zero value otherwise.
func (o *AppliedBy) GetAppliedByDeviceId() string {
	if o == nil || o.AppliedByDeviceId == nil {
		var ret string
		return ret
	}
	return *o.AppliedByDeviceId
}

// GetAppliedByDeviceIdOk returns a tuple with the AppliedByDeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppliedBy) GetAppliedByDeviceIdOk() (*string, bool) {
	if o == nil || o.AppliedByDeviceId == nil {
		return nil, false
	}
	return o.AppliedByDeviceId, true
}

// HasAppliedByDeviceId returns a boolean if a field has been set.
func (o *AppliedBy) HasAppliedByDeviceId() bool {
	if o != nil && o.AppliedByDeviceId != nil {
		return true
	}

	return false
}

// SetAppliedByDeviceId gets a reference to the given string and assigns it to the AppliedByDeviceId field.
func (o *AppliedBy) SetAppliedByDeviceId(v string) {
	o.AppliedByDeviceId = &v
}

func (o AppliedBy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AppliedBy != nil {
		toSerialize["appliedBy"] = o.AppliedBy
	}
	if o.AppliedByUsername != nil {
		toSerialize["appliedByUsername"] = o.AppliedByUsername
	}
	if o.AppliedByProviderName != nil {
		toSerialize["appliedByProviderName"] = o.AppliedByProviderName
	}
	if o.AppliedByDeviceId != nil {
		toSerialize["appliedByDeviceId"] = o.AppliedByDeviceId
	}
	return json.Marshal(toSerialize)
}

type NullableAppliedBy struct {
	value *AppliedBy
	isSet bool
}

func (v NullableAppliedBy) Get() *AppliedBy {
	return v.value
}

func (v *NullableAppliedBy) Set(val *AppliedBy) {
	v.value = val
	v.isSet = true
}

func (v NullableAppliedBy) IsSet() bool {
	return v.isSet
}

func (v *NullableAppliedBy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppliedBy(val *AppliedBy) *NullableAppliedBy {
	return &NullableAppliedBy{value: val, isSet: true}
}

func (v NullableAppliedBy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppliedBy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
