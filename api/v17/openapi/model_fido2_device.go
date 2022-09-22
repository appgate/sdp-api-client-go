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

// Fido2Device struct for Fido2Device
type Fido2Device struct {
	// Distinguished name of a user. Format: \"CN=,OU=\"
	UserDistinguishedName *string `json:"userDistinguishedName,omitempty"`
	// The username, same as the one in the user Distinguished Name.
	Username *string `json:"username,omitempty"`
	// The provider name of the user, same as the one in the user Distinguished Name.
	ProviderName *string `json:"providerName,omitempty"`
	// The device ID reported by the FIDO2 device during registration. May be empty.
	DeviceId *string `json:"deviceId,omitempty"`
	// The device name reported by the FIDO2 device during registration. May be empty.
	DeviceName *string `json:"deviceName,omitempty"`
}

// NewFido2Device instantiates a new Fido2Device object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFido2Device() *Fido2Device {
	this := Fido2Device{}
	return &this
}

// NewFido2DeviceWithDefaults instantiates a new Fido2Device object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFido2DeviceWithDefaults() *Fido2Device {
	this := Fido2Device{}
	return &this
}

// GetUserDistinguishedName returns the UserDistinguishedName field value if set, zero value otherwise.
func (o *Fido2Device) GetUserDistinguishedName() string {
	if o == nil || o.UserDistinguishedName == nil {
		var ret string
		return ret
	}
	return *o.UserDistinguishedName
}

// GetUserDistinguishedNameOk returns a tuple with the UserDistinguishedName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Fido2Device) GetUserDistinguishedNameOk() (*string, bool) {
	if o == nil || o.UserDistinguishedName == nil {
		return nil, false
	}
	return o.UserDistinguishedName, true
}

// HasUserDistinguishedName returns a boolean if a field has been set.
func (o *Fido2Device) HasUserDistinguishedName() bool {
	if o != nil && o.UserDistinguishedName != nil {
		return true
	}

	return false
}

// SetUserDistinguishedName gets a reference to the given string and assigns it to the UserDistinguishedName field.
func (o *Fido2Device) SetUserDistinguishedName(v string) {
	o.UserDistinguishedName = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *Fido2Device) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Fido2Device) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *Fido2Device) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *Fido2Device) SetUsername(v string) {
	o.Username = &v
}

// GetProviderName returns the ProviderName field value if set, zero value otherwise.
func (o *Fido2Device) GetProviderName() string {
	if o == nil || o.ProviderName == nil {
		var ret string
		return ret
	}
	return *o.ProviderName
}

// GetProviderNameOk returns a tuple with the ProviderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Fido2Device) GetProviderNameOk() (*string, bool) {
	if o == nil || o.ProviderName == nil {
		return nil, false
	}
	return o.ProviderName, true
}

// HasProviderName returns a boolean if a field has been set.
func (o *Fido2Device) HasProviderName() bool {
	if o != nil && o.ProviderName != nil {
		return true
	}

	return false
}

// SetProviderName gets a reference to the given string and assigns it to the ProviderName field.
func (o *Fido2Device) SetProviderName(v string) {
	o.ProviderName = &v
}

// GetDeviceId returns the DeviceId field value if set, zero value otherwise.
func (o *Fido2Device) GetDeviceId() string {
	if o == nil || o.DeviceId == nil {
		var ret string
		return ret
	}
	return *o.DeviceId
}

// GetDeviceIdOk returns a tuple with the DeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Fido2Device) GetDeviceIdOk() (*string, bool) {
	if o == nil || o.DeviceId == nil {
		return nil, false
	}
	return o.DeviceId, true
}

// HasDeviceId returns a boolean if a field has been set.
func (o *Fido2Device) HasDeviceId() bool {
	if o != nil && o.DeviceId != nil {
		return true
	}

	return false
}

// SetDeviceId gets a reference to the given string and assigns it to the DeviceId field.
func (o *Fido2Device) SetDeviceId(v string) {
	o.DeviceId = &v
}

// GetDeviceName returns the DeviceName field value if set, zero value otherwise.
func (o *Fido2Device) GetDeviceName() string {
	if o == nil || o.DeviceName == nil {
		var ret string
		return ret
	}
	return *o.DeviceName
}

// GetDeviceNameOk returns a tuple with the DeviceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Fido2Device) GetDeviceNameOk() (*string, bool) {
	if o == nil || o.DeviceName == nil {
		return nil, false
	}
	return o.DeviceName, true
}

// HasDeviceName returns a boolean if a field has been set.
func (o *Fido2Device) HasDeviceName() bool {
	if o != nil && o.DeviceName != nil {
		return true
	}

	return false
}

// SetDeviceName gets a reference to the given string and assigns it to the DeviceName field.
func (o *Fido2Device) SetDeviceName(v string) {
	o.DeviceName = &v
}

func (o Fido2Device) MarshalJSON() ([]byte, error) {
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
	if o.DeviceId != nil {
		toSerialize["deviceId"] = o.DeviceId
	}
	if o.DeviceName != nil {
		toSerialize["deviceName"] = o.DeviceName
	}
	return json.Marshal(toSerialize)
}

type NullableFido2Device struct {
	value *Fido2Device
	isSet bool
}

func (v NullableFido2Device) Get() *Fido2Device {
	return v.value
}

func (v *NullableFido2Device) Set(val *Fido2Device) {
	v.value = val
	v.isSet = true
}

func (v NullableFido2Device) IsSet() bool {
	return v.isSet
}

func (v *NullableFido2Device) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFido2Device(val *Fido2Device) *NullableFido2Device {
	return &NullableFido2Device{value: val, isSet: true}
}

func (v NullableFido2Device) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFido2Device) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
