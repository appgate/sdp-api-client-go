/*
Appgate SDP Controller REST API

# About   This specification documents the REST API calls for the Appgate SDP Controller.    Please refer to the REST API chapter in the manual or contact Appgate support with any questions about   this functionality. # Getting Started   Requirements for API scripting:   - Access to the Admin/API TLS Connection (default port 8443) of a Controller appliance.     (https://sdphelp.appgate.com/adminguide/appliance-function-configure.html?anchor=admin-api)   - An API user with relevant permissions.     (https://sdphelp.appgate.com/adminguide/administrative-roles-configure.html)   - In order to use the simple login API, Admin MFA must be disabled or the API user must be excluded.     (https://sdphelp.appgate.com/adminguide/mfa-for-admins.html) # Base path   HTTPS requests must be sent to the Admin Interface hostname and port, with **_/admin** path.    For example: **https://appgate.company.com:8443/admin**    All requests must have the **Accept** header as:    **application/vnd.appgate.peer-v20+json**    An exception is made for the **_/admin/version** endpoint which instead expects an **application/json** Accept header. # API Conventions   API conventions are  important to understand and follow strictly.    - While updating objects (via PUT), entire object must be sent with all fields.     - For example, in order to add a remedy method to the condition below:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": []       }       ```     - send the entire object with updated and non-updated fields:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": [{\"type\": \"DisplayMessage\", \"message\": \"test message\"}]       }       ```    - In case Controller returns an error (non-2xx HTTP status code), response body is JSON.     The \"message\" field contains information about the error.     HTTP 422 \"Unprocessable Entity\" has extra `errors` field to list all the issues with specific fields.    - Empty string (\"\") is considered a different value than \"null\" or field being omitted from JSON.     Omitting the field is recommended if no value is intended.     Empty string (\"\") will be almost always rejected as invalid value.    - There are common pattern between many objects:     - **Configuration Objects**: There are many objects with common fields, namely \"id\", \"name\", \"notes\", \"created\"       and \"updated\". These entities are listed, queried, created, updated and deleted in a similar fashion.     - **Distinguished Name**: Users and Devices are identified with what is called Distinguished Names, as used in        LDAP. The distinguished format that identifies a device and a user combination is        \"CN=\\<Device ID\\>,CN=\\<username\\>,OU=\\<Identity Provider Name\\>\". Some objects have the        \"userDistinguishedName\" field, which does not include the CN for Device ID.        This identifies a user on every device.

API version: API version 20.0
Contact: appgatesdp.support@appgate.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// ApplianceAllOfNetworkingNics struct for ApplianceAllOfNetworkingNics
type ApplianceAllOfNetworkingNics struct {
	// Whether the NIC is active or not.
	Enabled *bool `json:"enabled,omitempty"`
	// NIC name
	Name string                        `json:"name"`
	Ipv4 *ApplianceAllOfNetworkingIpv4 `json:"ipv4,omitempty"`
	Ipv6 *ApplianceAllOfNetworkingIpv6 `json:"ipv6,omitempty"`
	// MTU setting for the NIC. If left empty, appliance default will be used.
	Mtu *int32 `json:"mtu,omitempty"`
}

// NewApplianceAllOfNetworkingNics instantiates a new ApplianceAllOfNetworkingNics object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplianceAllOfNetworkingNics(name string) *ApplianceAllOfNetworkingNics {
	this := ApplianceAllOfNetworkingNics{}
	this.Name = name
	return &this
}

// NewApplianceAllOfNetworkingNicsWithDefaults instantiates a new ApplianceAllOfNetworkingNics object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplianceAllOfNetworkingNicsWithDefaults() *ApplianceAllOfNetworkingNics {
	this := ApplianceAllOfNetworkingNics{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *ApplianceAllOfNetworkingNics) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceAllOfNetworkingNics) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *ApplianceAllOfNetworkingNics) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *ApplianceAllOfNetworkingNics) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetName returns the Name field value
func (o *ApplianceAllOfNetworkingNics) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ApplianceAllOfNetworkingNics) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ApplianceAllOfNetworkingNics) SetName(v string) {
	o.Name = v
}

// GetIpv4 returns the Ipv4 field value if set, zero value otherwise.
func (o *ApplianceAllOfNetworkingNics) GetIpv4() ApplianceAllOfNetworkingIpv4 {
	if o == nil || o.Ipv4 == nil {
		var ret ApplianceAllOfNetworkingIpv4
		return ret
	}
	return *o.Ipv4
}

// GetIpv4Ok returns a tuple with the Ipv4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceAllOfNetworkingNics) GetIpv4Ok() (*ApplianceAllOfNetworkingIpv4, bool) {
	if o == nil || o.Ipv4 == nil {
		return nil, false
	}
	return o.Ipv4, true
}

// HasIpv4 returns a boolean if a field has been set.
func (o *ApplianceAllOfNetworkingNics) HasIpv4() bool {
	if o != nil && o.Ipv4 != nil {
		return true
	}

	return false
}

// SetIpv4 gets a reference to the given ApplianceAllOfNetworkingIpv4 and assigns it to the Ipv4 field.
func (o *ApplianceAllOfNetworkingNics) SetIpv4(v ApplianceAllOfNetworkingIpv4) {
	o.Ipv4 = &v
}

// GetIpv6 returns the Ipv6 field value if set, zero value otherwise.
func (o *ApplianceAllOfNetworkingNics) GetIpv6() ApplianceAllOfNetworkingIpv6 {
	if o == nil || o.Ipv6 == nil {
		var ret ApplianceAllOfNetworkingIpv6
		return ret
	}
	return *o.Ipv6
}

// GetIpv6Ok returns a tuple with the Ipv6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceAllOfNetworkingNics) GetIpv6Ok() (*ApplianceAllOfNetworkingIpv6, bool) {
	if o == nil || o.Ipv6 == nil {
		return nil, false
	}
	return o.Ipv6, true
}

// HasIpv6 returns a boolean if a field has been set.
func (o *ApplianceAllOfNetworkingNics) HasIpv6() bool {
	if o != nil && o.Ipv6 != nil {
		return true
	}

	return false
}

// SetIpv6 gets a reference to the given ApplianceAllOfNetworkingIpv6 and assigns it to the Ipv6 field.
func (o *ApplianceAllOfNetworkingNics) SetIpv6(v ApplianceAllOfNetworkingIpv6) {
	o.Ipv6 = &v
}

// GetMtu returns the Mtu field value if set, zero value otherwise.
func (o *ApplianceAllOfNetworkingNics) GetMtu() int32 {
	if o == nil || o.Mtu == nil {
		var ret int32
		return ret
	}
	return *o.Mtu
}

// GetMtuOk returns a tuple with the Mtu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceAllOfNetworkingNics) GetMtuOk() (*int32, bool) {
	if o == nil || o.Mtu == nil {
		return nil, false
	}
	return o.Mtu, true
}

// HasMtu returns a boolean if a field has been set.
func (o *ApplianceAllOfNetworkingNics) HasMtu() bool {
	if o != nil && o.Mtu != nil {
		return true
	}

	return false
}

// SetMtu gets a reference to the given int32 and assigns it to the Mtu field.
func (o *ApplianceAllOfNetworkingNics) SetMtu(v int32) {
	o.Mtu = &v
}

func (o ApplianceAllOfNetworkingNics) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Ipv4 != nil {
		toSerialize["ipv4"] = o.Ipv4
	}
	if o.Ipv6 != nil {
		toSerialize["ipv6"] = o.Ipv6
	}
	if o.Mtu != nil {
		toSerialize["mtu"] = o.Mtu
	}
	return json.Marshal(toSerialize)
}

type NullableApplianceAllOfNetworkingNics struct {
	value *ApplianceAllOfNetworkingNics
	isSet bool
}

func (v NullableApplianceAllOfNetworkingNics) Get() *ApplianceAllOfNetworkingNics {
	return v.value
}

func (v *NullableApplianceAllOfNetworkingNics) Set(val *ApplianceAllOfNetworkingNics) {
	v.value = val
	v.isSet = true
}

func (v NullableApplianceAllOfNetworkingNics) IsSet() bool {
	return v.isSet
}

func (v *NullableApplianceAllOfNetworkingNics) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplianceAllOfNetworkingNics(val *ApplianceAllOfNetworkingNics) *NullableApplianceAllOfNetworkingNics {
	return &NullableApplianceAllOfNetworkingNics{value: val, isSet: true}
}

func (v NullableApplianceAllOfNetworkingNics) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplianceAllOfNetworkingNics) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
