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

// SessionInfoDistinguishedName struct for SessionInfoDistinguishedName
type SessionInfoDistinguishedName struct {
	// The device ID, same as the one in the Distinguished Name.
	DeviceId *string `json:"deviceId,omitempty"`
	// The username, same as the one in the Distinguished Name.
	Username *string `json:"username,omitempty"`
	// The provider name of the user, same as the one in the Distinguished Name.
	ProviderName *string `json:"providerName,omitempty"`
	// The session details per Gateway. The key is the Appliance name and the value contains the Active Client Session details.
	Data *map[string]SessionInfoDistinguishedNameDataValue `json:"data,omitempty"`
}

// NewSessionInfoDistinguishedName instantiates a new SessionInfoDistinguishedName object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSessionInfoDistinguishedName() *SessionInfoDistinguishedName {
	this := SessionInfoDistinguishedName{}
	return &this
}

// NewSessionInfoDistinguishedNameWithDefaults instantiates a new SessionInfoDistinguishedName object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSessionInfoDistinguishedNameWithDefaults() *SessionInfoDistinguishedName {
	this := SessionInfoDistinguishedName{}
	return &this
}

// GetDeviceId returns the DeviceId field value if set, zero value otherwise.
func (o *SessionInfoDistinguishedName) GetDeviceId() string {
	if o == nil || o.DeviceId == nil {
		var ret string
		return ret
	}
	return *o.DeviceId
}

// GetDeviceIdOk returns a tuple with the DeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionInfoDistinguishedName) GetDeviceIdOk() (*string, bool) {
	if o == nil || o.DeviceId == nil {
		return nil, false
	}
	return o.DeviceId, true
}

// HasDeviceId returns a boolean if a field has been set.
func (o *SessionInfoDistinguishedName) HasDeviceId() bool {
	if o != nil && o.DeviceId != nil {
		return true
	}

	return false
}

// SetDeviceId gets a reference to the given string and assigns it to the DeviceId field.
func (o *SessionInfoDistinguishedName) SetDeviceId(v string) {
	o.DeviceId = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *SessionInfoDistinguishedName) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionInfoDistinguishedName) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *SessionInfoDistinguishedName) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *SessionInfoDistinguishedName) SetUsername(v string) {
	o.Username = &v
}

// GetProviderName returns the ProviderName field value if set, zero value otherwise.
func (o *SessionInfoDistinguishedName) GetProviderName() string {
	if o == nil || o.ProviderName == nil {
		var ret string
		return ret
	}
	return *o.ProviderName
}

// GetProviderNameOk returns a tuple with the ProviderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionInfoDistinguishedName) GetProviderNameOk() (*string, bool) {
	if o == nil || o.ProviderName == nil {
		return nil, false
	}
	return o.ProviderName, true
}

// HasProviderName returns a boolean if a field has been set.
func (o *SessionInfoDistinguishedName) HasProviderName() bool {
	if o != nil && o.ProviderName != nil {
		return true
	}

	return false
}

// SetProviderName gets a reference to the given string and assigns it to the ProviderName field.
func (o *SessionInfoDistinguishedName) SetProviderName(v string) {
	o.ProviderName = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *SessionInfoDistinguishedName) GetData() map[string]SessionInfoDistinguishedNameDataValue {
	if o == nil || o.Data == nil {
		var ret map[string]SessionInfoDistinguishedNameDataValue
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionInfoDistinguishedName) GetDataOk() (*map[string]SessionInfoDistinguishedNameDataValue, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *SessionInfoDistinguishedName) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given map[string]SessionInfoDistinguishedNameDataValue and assigns it to the Data field.
func (o *SessionInfoDistinguishedName) SetData(v map[string]SessionInfoDistinguishedNameDataValue) {
	o.Data = &v
}

func (o SessionInfoDistinguishedName) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DeviceId != nil {
		toSerialize["deviceId"] = o.DeviceId
	}
	if o.Username != nil {
		toSerialize["username"] = o.Username
	}
	if o.ProviderName != nil {
		toSerialize["providerName"] = o.ProviderName
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableSessionInfoDistinguishedName struct {
	value *SessionInfoDistinguishedName
	isSet bool
}

func (v NullableSessionInfoDistinguishedName) Get() *SessionInfoDistinguishedName {
	return v.value
}

func (v *NullableSessionInfoDistinguishedName) Set(val *SessionInfoDistinguishedName) {
	v.value = val
	v.isSet = true
}

func (v NullableSessionInfoDistinguishedName) IsSet() bool {
	return v.isSet
}

func (v *NullableSessionInfoDistinguishedName) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSessionInfoDistinguishedName(val *SessionInfoDistinguishedName) *NullableSessionInfoDistinguishedName {
	return &NullableSessionInfoDistinguishedName{value: val, isSet: true}
}

func (v NullableSessionInfoDistinguishedName) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSessionInfoDistinguishedName) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
