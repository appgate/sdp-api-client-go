/*
Appgate SDP Controller REST API

# About   This specification documents the REST API calls for the Appgate SDP Controller.    Please refer to the REST API chapter in the manual or contact Appgate support with any questions about   this functionality. # Getting Started   Requirements for API scripting:   - Access to the Admin/API TLS Connection (default port 8443) of a Controller appliance.     (https://sdphelp.appgate.com/adminguide/appliance-function-configure.html?anchor=admin-api)   - An API user with relevant permissions.     (https://sdphelp.appgate.com/adminguide/administrative-roles-configure.html)   - In order to use the simple login API, Admin MFA must be disabled or the API user must be excluded.     (https://sdphelp.appgate.com/adminguide/mfa-for-admins.html) # Base path   HTTPS requests must be sent to the Admin Interface hostname and port, with **_/admin** path.    For example: **https://appgate.company.com:8443/admin**    All requests must have the **Accept** header as:    **application/vnd.appgate.peer-v16+json** # API Conventions   API conventions are  important to understand and follow strictly.    - While updating objects (via PUT), entire object must be sent with all fields.     - For example, in order to add a remedy method to the condition below:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": []       }       ```     - send the entire object with updated and non-updated fields:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": [{\"type\": \"DisplayMessage\", \"message\": \"test message\"}]       }       ```    - In case Controller returns an error (non-2xx HTTP status code), response body is JSON.     The \"message\" field contains information about the error.     HTTP 422 \"Unprocessable Entity\" has extra `errors` field to list all the issues with specific fields.    - Empty string (\"\") is considered a different value than \"null\" or field being omitted from JSON.     Omitting the field is recommend if no value is intended.     Empty string (\"\") will be almost always rejected as invalid value.    - There are common pattern between many objects:     - **Configuration Objects**: There are many objects with common fields, namely \"id\", \"name\", \"notes\", \"created\"       and \"updated\". These entities are listed, queried, created, updated and deleted in a similar fashion.     - **Distinguished Name**: Users and Devices are identified with what is called Distinguished Names, as used in        LDAP. The distinguished format that identifies a device and a user combination is        \"CN=\\<Device ID\\>,CN=\\<username\\>,OU=\\<Identity Provider Name\\>\". Some objects have the        \"userDistinguishedName\" field, which does not include the CN for Device ID.        This identifies a user on every device.

API version: API version 16.0
Contact: appgatesdp.support@appgate.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// ApplianceAllOfNtpServers NTP server.
type ApplianceAllOfNtpServers struct {
	// Hostname or IP of the NTP server.
	Hostname string `json:"hostname"`
	// Type of key to use for secure NTP communication.
	KeyType *string `json:"keyType,omitempty"`
	// Identifier number for the key.
	KeyNo *int32 `json:"keyNo,omitempty"`
	// Key to use for secure NTP communication.
	Key *string `json:"key,omitempty"`
}

// NewApplianceAllOfNtpServers instantiates a new ApplianceAllOfNtpServers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplianceAllOfNtpServers(hostname string) *ApplianceAllOfNtpServers {
	this := ApplianceAllOfNtpServers{}
	this.Hostname = hostname
	return &this
}

// NewApplianceAllOfNtpServersWithDefaults instantiates a new ApplianceAllOfNtpServers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplianceAllOfNtpServersWithDefaults() *ApplianceAllOfNtpServers {
	this := ApplianceAllOfNtpServers{}
	return &this
}

// GetHostname returns the Hostname field value
func (o *ApplianceAllOfNtpServers) GetHostname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value
// and a boolean to check if the value has been set.
func (o *ApplianceAllOfNtpServers) GetHostnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hostname, true
}

// SetHostname sets field value
func (o *ApplianceAllOfNtpServers) SetHostname(v string) {
	o.Hostname = v
}

// GetKeyType returns the KeyType field value if set, zero value otherwise.
func (o *ApplianceAllOfNtpServers) GetKeyType() string {
	if o == nil || o.KeyType == nil {
		var ret string
		return ret
	}
	return *o.KeyType
}

// GetKeyTypeOk returns a tuple with the KeyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceAllOfNtpServers) GetKeyTypeOk() (*string, bool) {
	if o == nil || o.KeyType == nil {
		return nil, false
	}
	return o.KeyType, true
}

// HasKeyType returns a boolean if a field has been set.
func (o *ApplianceAllOfNtpServers) HasKeyType() bool {
	if o != nil && o.KeyType != nil {
		return true
	}

	return false
}

// SetKeyType gets a reference to the given string and assigns it to the KeyType field.
func (o *ApplianceAllOfNtpServers) SetKeyType(v string) {
	o.KeyType = &v
}

// GetKeyNo returns the KeyNo field value if set, zero value otherwise.
func (o *ApplianceAllOfNtpServers) GetKeyNo() int32 {
	if o == nil || o.KeyNo == nil {
		var ret int32
		return ret
	}
	return *o.KeyNo
}

// GetKeyNoOk returns a tuple with the KeyNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceAllOfNtpServers) GetKeyNoOk() (*int32, bool) {
	if o == nil || o.KeyNo == nil {
		return nil, false
	}
	return o.KeyNo, true
}

// HasKeyNo returns a boolean if a field has been set.
func (o *ApplianceAllOfNtpServers) HasKeyNo() bool {
	if o != nil && o.KeyNo != nil {
		return true
	}

	return false
}

// SetKeyNo gets a reference to the given int32 and assigns it to the KeyNo field.
func (o *ApplianceAllOfNtpServers) SetKeyNo(v int32) {
	o.KeyNo = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *ApplianceAllOfNtpServers) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceAllOfNtpServers) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *ApplianceAllOfNtpServers) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *ApplianceAllOfNtpServers) SetKey(v string) {
	o.Key = &v
}

func (o ApplianceAllOfNtpServers) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["hostname"] = o.Hostname
	}
	if o.KeyType != nil {
		toSerialize["keyType"] = o.KeyType
	}
	if o.KeyNo != nil {
		toSerialize["keyNo"] = o.KeyNo
	}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	return json.Marshal(toSerialize)
}

type NullableApplianceAllOfNtpServers struct {
	value *ApplianceAllOfNtpServers
	isSet bool
}

func (v NullableApplianceAllOfNtpServers) Get() *ApplianceAllOfNtpServers {
	return v.value
}

func (v *NullableApplianceAllOfNtpServers) Set(val *ApplianceAllOfNtpServers) {
	v.value = val
	v.isSet = true
}

func (v NullableApplianceAllOfNtpServers) IsSet() bool {
	return v.isSet
}

func (v *NullableApplianceAllOfNtpServers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplianceAllOfNtpServers(val *ApplianceAllOfNtpServers) *NullableApplianceAllOfNtpServers {
	return &NullableApplianceAllOfNtpServers{value: val, isSet: true}
}

func (v NullableApplianceAllOfNtpServers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplianceAllOfNtpServers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
