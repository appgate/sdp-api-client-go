/*
Appgate SDP Controller REST API

# About   This specification documents the REST API calls for the Appgate SDP Controller.    Please refer to the REST API chapter in the manual or contact Appgate support with any questions about   this functionality. # Getting Started   Requirements for API scripting:   - Access to the Admin/API TLS Connection (default port 8443) of a Controller appliance.     (https://sdphelp.appgate.com/adminguide/appliance-function-configure.html?anchor=admin-api)   - An API user with relevant permissions.     (https://sdphelp.appgate.com/adminguide/administrative-roles-configure.html)   - In order to use the simple login API, Admin MFA must be disabled or the API user must be excluded.     (https://sdphelp.appgate.com/adminguide/mfa-for-admins.html) # Base path   HTTPS requests must be sent to the Admin Interface hostname and port, with **_/admin** path.    For example: **https://appgate.company.com:8443/admin**    All requests must have the **Accept** header as:    **application/vnd.appgate.peer-v21+json**    An exception is made for the **_/admin/version** endpoint which instead expects an **application/json** Accept header. # API Conventions   API conventions are  important to understand and follow strictly.    - While updating objects (via PUT), entire object must be sent with all fields.     - For example, in order to add a remedy method to the condition below:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": []       }       ```     - send the entire object with updated and non-updated fields:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": [{\"type\": \"DisplayMessage\", \"message\": \"test message\"}]       }       ```    - In case Controller returns an error (non-2xx HTTP status code), response body is JSON.     The \"message\" field contains information about the error.     HTTP 422 \"Unprocessable Entity\" has extra `errors` field to list all the issues with specific fields.    - Empty string (\"\") is considered a different value than \"null\" or field being omitted from JSON.     Omitting the field is recommended if no value is intended.     Empty string (\"\") will be almost always rejected as invalid value.    - There are common pattern between many objects:     - **Configuration Objects**: There are many objects with common fields, namely \"id\", \"name\", \"notes\", \"created\"       and \"updated\". These entities are listed, queried, created, updated and deleted in a similar fashion.     - **Distinguished Name**: Users and Devices are identified with what is called Distinguished Names, as used in        LDAP. The distinguished format that identifies a device and a user combination is        \"CN=\\<Device ID\\>,CN=\\<username\\>,OU=\\<Identity Provider Name\\>\". Some objects have the        \"userDistinguishedName\" field, which does not include the CN for Device ID.        This identifies a user on every device.

API version: API version 21.1
Contact: appgatesdp.support@appgate.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// OnBoardedDeviceAllOf struct for OnBoardedDeviceAllOf
type OnBoardedDeviceAllOf struct {
	// Type of the registered device.
	DeviceType *string `json:"device_type,omitempty"`
	// Hostname of the Device at the time of registration, sent by the Device.
	Hostname *string `json:"hostname,omitempty"`
	// Registration time.
	OnBoardedAt *time.Time `json:"onBoardedAt,omitempty"`
	// The time when the device last signed in. 'null' if it has signed in last on a Controller that was older than 5.4 at the time.
	LastSeenAt *time.Time `json:"lastSeenAt,omitempty"`
}

// NewOnBoardedDeviceAllOf instantiates a new OnBoardedDeviceAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOnBoardedDeviceAllOf() *OnBoardedDeviceAllOf {
	this := OnBoardedDeviceAllOf{}
	return &this
}

// NewOnBoardedDeviceAllOfWithDefaults instantiates a new OnBoardedDeviceAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOnBoardedDeviceAllOfWithDefaults() *OnBoardedDeviceAllOf {
	this := OnBoardedDeviceAllOf{}
	return &this
}

// GetDeviceType returns the DeviceType field value if set, zero value otherwise.
func (o *OnBoardedDeviceAllOf) GetDeviceType() string {
	if o == nil || o.DeviceType == nil {
		var ret string
		return ret
	}
	return *o.DeviceType
}

// GetDeviceTypeOk returns a tuple with the DeviceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnBoardedDeviceAllOf) GetDeviceTypeOk() (*string, bool) {
	if o == nil || o.DeviceType == nil {
		return nil, false
	}
	return o.DeviceType, true
}

// HasDeviceType returns a boolean if a field has been set.
func (o *OnBoardedDeviceAllOf) HasDeviceType() bool {
	if o != nil && o.DeviceType != nil {
		return true
	}

	return false
}

// SetDeviceType gets a reference to the given string and assigns it to the DeviceType field.
func (o *OnBoardedDeviceAllOf) SetDeviceType(v string) {
	o.DeviceType = &v
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *OnBoardedDeviceAllOf) GetHostname() string {
	if o == nil || o.Hostname == nil {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnBoardedDeviceAllOf) GetHostnameOk() (*string, bool) {
	if o == nil || o.Hostname == nil {
		return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *OnBoardedDeviceAllOf) HasHostname() bool {
	if o != nil && o.Hostname != nil {
		return true
	}

	return false
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *OnBoardedDeviceAllOf) SetHostname(v string) {
	o.Hostname = &v
}

// GetOnBoardedAt returns the OnBoardedAt field value if set, zero value otherwise.
func (o *OnBoardedDeviceAllOf) GetOnBoardedAt() time.Time {
	if o == nil || o.OnBoardedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.OnBoardedAt
}

// GetOnBoardedAtOk returns a tuple with the OnBoardedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnBoardedDeviceAllOf) GetOnBoardedAtOk() (*time.Time, bool) {
	if o == nil || o.OnBoardedAt == nil {
		return nil, false
	}
	return o.OnBoardedAt, true
}

// HasOnBoardedAt returns a boolean if a field has been set.
func (o *OnBoardedDeviceAllOf) HasOnBoardedAt() bool {
	if o != nil && o.OnBoardedAt != nil {
		return true
	}

	return false
}

// SetOnBoardedAt gets a reference to the given time.Time and assigns it to the OnBoardedAt field.
func (o *OnBoardedDeviceAllOf) SetOnBoardedAt(v time.Time) {
	o.OnBoardedAt = &v
}

// GetLastSeenAt returns the LastSeenAt field value if set, zero value otherwise.
func (o *OnBoardedDeviceAllOf) GetLastSeenAt() time.Time {
	if o == nil || o.LastSeenAt == nil {
		var ret time.Time
		return ret
	}
	return *o.LastSeenAt
}

// GetLastSeenAtOk returns a tuple with the LastSeenAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnBoardedDeviceAllOf) GetLastSeenAtOk() (*time.Time, bool) {
	if o == nil || o.LastSeenAt == nil {
		return nil, false
	}
	return o.LastSeenAt, true
}

// HasLastSeenAt returns a boolean if a field has been set.
func (o *OnBoardedDeviceAllOf) HasLastSeenAt() bool {
	if o != nil && o.LastSeenAt != nil {
		return true
	}

	return false
}

// SetLastSeenAt gets a reference to the given time.Time and assigns it to the LastSeenAt field.
func (o *OnBoardedDeviceAllOf) SetLastSeenAt(v time.Time) {
	o.LastSeenAt = &v
}

func (o OnBoardedDeviceAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DeviceType != nil {
		toSerialize["device_type"] = o.DeviceType
	}
	if o.Hostname != nil {
		toSerialize["hostname"] = o.Hostname
	}
	if o.OnBoardedAt != nil {
		toSerialize["onBoardedAt"] = o.OnBoardedAt
	}
	if o.LastSeenAt != nil {
		toSerialize["lastSeenAt"] = o.LastSeenAt
	}
	return json.Marshal(toSerialize)
}

type NullableOnBoardedDeviceAllOf struct {
	value *OnBoardedDeviceAllOf
	isSet bool
}

func (v NullableOnBoardedDeviceAllOf) Get() *OnBoardedDeviceAllOf {
	return v.value
}

func (v *NullableOnBoardedDeviceAllOf) Set(val *OnBoardedDeviceAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableOnBoardedDeviceAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableOnBoardedDeviceAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOnBoardedDeviceAllOf(val *OnBoardedDeviceAllOf) *NullableOnBoardedDeviceAllOf {
	return &NullableOnBoardedDeviceAllOf{value: val, isSet: true}
}

func (v NullableOnBoardedDeviceAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOnBoardedDeviceAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
