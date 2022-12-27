/*
Appgate SDP Controller REST API

# About   This specification documents the REST API calls for the Appgate SDP Controller.    Please refer to the REST API chapter in the manual or contact Appgate support with any questions about   this functionality. # Getting Started   Requirements for API scripting:   - Access to the Admin/API TLS Connection (default port 8443) of a Controller appliance.     (https://sdphelp.appgate.com/adminguide/appliance-function-configure.html?anchor=admin-api)   - An API user with relevant permissions.     (https://sdphelp.appgate.com/adminguide/administrative-roles-configure.html)   - In order to use the simple login API, Admin MFA must be disabled or the API user must be excluded.     (https://sdphelp.appgate.com/adminguide/mfa-for-admins.html) # Base path   HTTPS requests must be sent to the Admin Interface hostname and port, with **_/admin** path.    For example: **https://appgate.company.com:8443/admin**    All requests must have the **Accept** header as:    **application/vnd.appgate.peer-v18+json** # API Conventions   API conventions are  important to understand and follow strictly.    - While updating objects (via PUT), entire object must be sent with all fields.     - For example, in order to add a remedy method to the condition below:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": []       }       ```     - send the entire object with updated and non-updated fields:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": [{\"type\": \"DisplayMessage\", \"message\": \"test message\"}]       }       ```    - In case Controller returns an error (non-2xx HTTP status code), response body is JSON.     The \"message\" field contains information about the error.     HTTP 422 \"Unprocessable Entity\" has extra `errors` field to list all the issues with specific fields.    - Empty string (\"\") is considered a different value than \"null\" or field being omitted from JSON.     Omitting the field is recommended if no value is intended.     Empty string (\"\") will be almost always rejected as invalid value.    - There are common pattern between many objects:     - **Configuration Objects**: There are many objects with common fields, namely \"id\", \"name\", \"notes\", \"created\"       and \"updated\". These entities are listed, queried, created, updated and deleted in a similar fashion.     - **Distinguished Name**: Users and Devices are identified with what is called Distinguished Names, as used in        LDAP. The distinguished format that identifies a device and a user combination is        \"CN=\\<Device ID\\>,CN=\\<username\\>,OU=\\<Identity Provider Name\\>\". Some objects have the        \"userDistinguishedName\" field, which does not include the CN for Device ID.        This identifies a user on every device.

API version: API version 18.2
Contact: appgatesdp.support@appgate.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// OnBoardedDevice struct for OnBoardedDevice
type OnBoardedDevice struct {
	// Distinguished name of a user&device combination. Format: \"CN=,CN=,OU=\"
	DistinguishedName *string `json:"distinguishedName,omitempty"`
	// The device ID, same as the one in the Distinguished Name.
	DeviceId *string `json:"deviceId,omitempty"`
	// The username, same as the one in the Distinguished Name.
	Username *string `json:"username,omitempty"`
	// The provider name of the user, same as the one in the Distinguished Name.
	ProviderName *string `json:"providerName,omitempty"`
	// Hostname of the Device at the time of registration, sent by the Device.
	Hostname *string `json:"hostname,omitempty"`
	// Registration time.
	OnBoardedAt *time.Time `json:"onBoardedAt,omitempty"`
	// The time when the device last signed in. 'null' if it has signed in last on a Controller that was older than 5.4 at the time.
	LastSeenAt *time.Time `json:"lastSeenAt,omitempty"`
}

// NewOnBoardedDevice instantiates a new OnBoardedDevice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOnBoardedDevice() *OnBoardedDevice {
	this := OnBoardedDevice{}
	return &this
}

// NewOnBoardedDeviceWithDefaults instantiates a new OnBoardedDevice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOnBoardedDeviceWithDefaults() *OnBoardedDevice {
	this := OnBoardedDevice{}
	return &this
}

// GetDistinguishedName returns the DistinguishedName field value if set, zero value otherwise.
func (o *OnBoardedDevice) GetDistinguishedName() string {
	if o == nil || o.DistinguishedName == nil {
		var ret string
		return ret
	}
	return *o.DistinguishedName
}

// GetDistinguishedNameOk returns a tuple with the DistinguishedName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnBoardedDevice) GetDistinguishedNameOk() (*string, bool) {
	if o == nil || o.DistinguishedName == nil {
		return nil, false
	}
	return o.DistinguishedName, true
}

// HasDistinguishedName returns a boolean if a field has been set.
func (o *OnBoardedDevice) HasDistinguishedName() bool {
	if o != nil && o.DistinguishedName != nil {
		return true
	}

	return false
}

// SetDistinguishedName gets a reference to the given string and assigns it to the DistinguishedName field.
func (o *OnBoardedDevice) SetDistinguishedName(v string) {
	o.DistinguishedName = &v
}

// GetDeviceId returns the DeviceId field value if set, zero value otherwise.
func (o *OnBoardedDevice) GetDeviceId() string {
	if o == nil || o.DeviceId == nil {
		var ret string
		return ret
	}
	return *o.DeviceId
}

// GetDeviceIdOk returns a tuple with the DeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnBoardedDevice) GetDeviceIdOk() (*string, bool) {
	if o == nil || o.DeviceId == nil {
		return nil, false
	}
	return o.DeviceId, true
}

// HasDeviceId returns a boolean if a field has been set.
func (o *OnBoardedDevice) HasDeviceId() bool {
	if o != nil && o.DeviceId != nil {
		return true
	}

	return false
}

// SetDeviceId gets a reference to the given string and assigns it to the DeviceId field.
func (o *OnBoardedDevice) SetDeviceId(v string) {
	o.DeviceId = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *OnBoardedDevice) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnBoardedDevice) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *OnBoardedDevice) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *OnBoardedDevice) SetUsername(v string) {
	o.Username = &v
}

// GetProviderName returns the ProviderName field value if set, zero value otherwise.
func (o *OnBoardedDevice) GetProviderName() string {
	if o == nil || o.ProviderName == nil {
		var ret string
		return ret
	}
	return *o.ProviderName
}

// GetProviderNameOk returns a tuple with the ProviderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnBoardedDevice) GetProviderNameOk() (*string, bool) {
	if o == nil || o.ProviderName == nil {
		return nil, false
	}
	return o.ProviderName, true
}

// HasProviderName returns a boolean if a field has been set.
func (o *OnBoardedDevice) HasProviderName() bool {
	if o != nil && o.ProviderName != nil {
		return true
	}

	return false
}

// SetProviderName gets a reference to the given string and assigns it to the ProviderName field.
func (o *OnBoardedDevice) SetProviderName(v string) {
	o.ProviderName = &v
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *OnBoardedDevice) GetHostname() string {
	if o == nil || o.Hostname == nil {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnBoardedDevice) GetHostnameOk() (*string, bool) {
	if o == nil || o.Hostname == nil {
		return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *OnBoardedDevice) HasHostname() bool {
	if o != nil && o.Hostname != nil {
		return true
	}

	return false
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *OnBoardedDevice) SetHostname(v string) {
	o.Hostname = &v
}

// GetOnBoardedAt returns the OnBoardedAt field value if set, zero value otherwise.
func (o *OnBoardedDevice) GetOnBoardedAt() time.Time {
	if o == nil || o.OnBoardedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.OnBoardedAt
}

// GetOnBoardedAtOk returns a tuple with the OnBoardedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnBoardedDevice) GetOnBoardedAtOk() (*time.Time, bool) {
	if o == nil || o.OnBoardedAt == nil {
		return nil, false
	}
	return o.OnBoardedAt, true
}

// HasOnBoardedAt returns a boolean if a field has been set.
func (o *OnBoardedDevice) HasOnBoardedAt() bool {
	if o != nil && o.OnBoardedAt != nil {
		return true
	}

	return false
}

// SetOnBoardedAt gets a reference to the given time.Time and assigns it to the OnBoardedAt field.
func (o *OnBoardedDevice) SetOnBoardedAt(v time.Time) {
	o.OnBoardedAt = &v
}

// GetLastSeenAt returns the LastSeenAt field value if set, zero value otherwise.
func (o *OnBoardedDevice) GetLastSeenAt() time.Time {
	if o == nil || o.LastSeenAt == nil {
		var ret time.Time
		return ret
	}
	return *o.LastSeenAt
}

// GetLastSeenAtOk returns a tuple with the LastSeenAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnBoardedDevice) GetLastSeenAtOk() (*time.Time, bool) {
	if o == nil || o.LastSeenAt == nil {
		return nil, false
	}
	return o.LastSeenAt, true
}

// HasLastSeenAt returns a boolean if a field has been set.
func (o *OnBoardedDevice) HasLastSeenAt() bool {
	if o != nil && o.LastSeenAt != nil {
		return true
	}

	return false
}

// SetLastSeenAt gets a reference to the given time.Time and assigns it to the LastSeenAt field.
func (o *OnBoardedDevice) SetLastSeenAt(v time.Time) {
	o.LastSeenAt = &v
}

func (o OnBoardedDevice) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DistinguishedName != nil {
		toSerialize["distinguishedName"] = o.DistinguishedName
	}
	if o.DeviceId != nil {
		toSerialize["deviceId"] = o.DeviceId
	}
	if o.Username != nil {
		toSerialize["username"] = o.Username
	}
	if o.ProviderName != nil {
		toSerialize["providerName"] = o.ProviderName
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

type NullableOnBoardedDevice struct {
	value *OnBoardedDevice
	isSet bool
}

func (v NullableOnBoardedDevice) Get() *OnBoardedDevice {
	return v.value
}

func (v *NullableOnBoardedDevice) Set(val *OnBoardedDevice) {
	v.value = val
	v.isSet = true
}

func (v NullableOnBoardedDevice) IsSet() bool {
	return v.isSet
}

func (v *NullableOnBoardedDevice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOnBoardedDevice(val *OnBoardedDevice) *NullableOnBoardedDevice {
	return &NullableOnBoardedDevice{value: val, isSet: true}
}

func (v NullableOnBoardedDevice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOnBoardedDevice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
