/*
Appgate SDP Controller REST API

# About   This specification documents the REST API calls for the Appgate SDP Controller.    Please refer to the REST API chapter in the manual or contact Appgate support with any questions about   this functionality. # Getting Started   Requirements for API scripting:   - Access to the Admin/API TLS Connection (default port 8443) of a Controller appliance.     (https://sdphelp.appgate.com/adminguide/appliance-function-configure.html?anchor=admin-api)   - An API user with relevant permissions.     (https://sdphelp.appgate.com/adminguide/administrative-roles-configure.html)   - In order to use the simple login API, Admin MFA must be disabled or the API user must be excluded.     (https://sdphelp.appgate.com/adminguide/mfa-for-admins.html) # Base path   HTTPS requests must be sent to the Admin Interface hostname and port, with **_/admin** path.    For example: **https://appgate.company.com:8443/admin**    All requests must have the **Accept** header as:    **application/vnd.appgate.peer-v17+json** # API Conventions   API conventions are  important to understand and follow strictly.    - While updating objects (via PUT), entire object must be sent with all fields.     - For example, in order to add a remedy method to the condition below:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": []       }       ```     - send the entire object with updated and non-updated fields:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": [{\"type\": \"DisplayMessage\", \"message\": \"test message\"}]       }       ```    - In case Controller returns an error (non-2xx HTTP status code), response body is JSON.     The \"message\" field contains information about the error.     HTTP 422 \"Unprocessable Entity\" has extra `errors` field to list all the issues with specific fields.    - Empty string (\"\") is considered a different value than \"null\" or field being omitted from JSON.     Omitting the field is recommended if no value is intended.     Empty string (\"\") will be almost always rejected as invalid value.    - There are common pattern between many objects:     - **Configuration Objects**: There are many objects with common fields, namely \"id\", \"name\", \"notes\", \"created\"       and \"updated\". These entities are listed, queried, created, updated and deleted in a similar fashion.     - **Distinguished Name**: Users and Devices are identified with what is called Distinguished Names, as used in        LDAP. The distinguished format that identifies a device and a user combination is        \"CN=\\<Device ID\\>,CN=\\<username\\>,OU=\\<Identity Provider Name\\>\". Some objects have the        \"userDistinguishedName\" field, which does not include the CN for Device ID.        This identifies a user on every device.

API version: API version 17.1
Contact: appgatesdp.support@appgate.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// AllowSourcesInner struct for AllowSourcesInner
type AllowSourcesInner struct {
	// IP address to allow connection.
	Address *string `json:"address,omitempty"`
	// Netmask to use with address for allowing connections.
	Netmask *int32 `json:"netmask,omitempty"`
	// NIC name to accept connections on.
	Nic *string `json:"nic,omitempty"`
}

// NewAllowSourcesInner instantiates a new AllowSourcesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAllowSourcesInner() *AllowSourcesInner {
	this := AllowSourcesInner{}
	return &this
}

// NewAllowSourcesInnerWithDefaults instantiates a new AllowSourcesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAllowSourcesInnerWithDefaults() *AllowSourcesInner {
	this := AllowSourcesInner{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *AllowSourcesInner) GetAddress() string {
	if o == nil || o.Address == nil {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllowSourcesInner) GetAddressOk() (*string, bool) {
	if o == nil || o.Address == nil {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *AllowSourcesInner) HasAddress() bool {
	if o != nil && o.Address != nil {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *AllowSourcesInner) SetAddress(v string) {
	o.Address = &v
}

// GetNetmask returns the Netmask field value if set, zero value otherwise.
func (o *AllowSourcesInner) GetNetmask() int32 {
	if o == nil || o.Netmask == nil {
		var ret int32
		return ret
	}
	return *o.Netmask
}

// GetNetmaskOk returns a tuple with the Netmask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllowSourcesInner) GetNetmaskOk() (*int32, bool) {
	if o == nil || o.Netmask == nil {
		return nil, false
	}
	return o.Netmask, true
}

// HasNetmask returns a boolean if a field has been set.
func (o *AllowSourcesInner) HasNetmask() bool {
	if o != nil && o.Netmask != nil {
		return true
	}

	return false
}

// SetNetmask gets a reference to the given int32 and assigns it to the Netmask field.
func (o *AllowSourcesInner) SetNetmask(v int32) {
	o.Netmask = &v
}

// GetNic returns the Nic field value if set, zero value otherwise.
func (o *AllowSourcesInner) GetNic() string {
	if o == nil || o.Nic == nil {
		var ret string
		return ret
	}
	return *o.Nic
}

// GetNicOk returns a tuple with the Nic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllowSourcesInner) GetNicOk() (*string, bool) {
	if o == nil || o.Nic == nil {
		return nil, false
	}
	return o.Nic, true
}

// HasNic returns a boolean if a field has been set.
func (o *AllowSourcesInner) HasNic() bool {
	if o != nil && o.Nic != nil {
		return true
	}

	return false
}

// SetNic gets a reference to the given string and assigns it to the Nic field.
func (o *AllowSourcesInner) SetNic(v string) {
	o.Nic = &v
}

func (o AllowSourcesInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Address != nil {
		toSerialize["address"] = o.Address
	}
	if o.Netmask != nil {
		toSerialize["netmask"] = o.Netmask
	}
	if o.Nic != nil {
		toSerialize["nic"] = o.Nic
	}
	return json.Marshal(toSerialize)
}

type NullableAllowSourcesInner struct {
	value *AllowSourcesInner
	isSet bool
}

func (v NullableAllowSourcesInner) Get() *AllowSourcesInner {
	return v.value
}

func (v *NullableAllowSourcesInner) Set(val *AllowSourcesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableAllowSourcesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableAllowSourcesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAllowSourcesInner(val *AllowSourcesInner) *NullableAllowSourcesInner {
	return &NullableAllowSourcesInner{value: val, isSet: true}
}

func (v NullableAllowSourcesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAllowSourcesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
