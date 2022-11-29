/*
Appgate SDP Controller REST API

# About   This specification documents the REST API calls for the Appgate SDP Controller.    Please refer to the REST API chapter in the manual or contact Appgate support with any questions about   this functionality. # Getting Started   Requirements for API scripting:   - Access to the Admin/API TLS Connection (default port 8443) of a Controller appliance.     (https://sdphelp.appgate.com/adminguide/appliance-function-configure.html?anchor=admin-api)   - An API user with relevant permissions.     (https://sdphelp.appgate.com/adminguide/administrative-roles-configure.html)   - In order to use the simple login API, Admin MFA must be disabled or the API user must be excluded.     (https://sdphelp.appgate.com/adminguide/mfa-for-admins.html) # Base path   HTTPS requests must be sent to the Admin Interface hostname and port, with **_/admin** path.    For example: **https://appgate.company.com:8443/admin**    All requests must have the **Accept** header as:    **application/vnd.appgate.peer-v18+json** # API Conventions   API conventions are  important to understand and follow strictly.    - While updating objects (via PUT), entire object must be sent with all fields.     - For example, in order to add a remedy method to the condition below:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": []       }       ```     - send the entire object with updated and non-updated fields:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": [{\"type\": \"DisplayMessage\", \"message\": \"test message\"}]       }       ```    - In case Controller returns an error (non-2xx HTTP status code), response body is JSON.     The \"message\" field contains information about the error.     HTTP 422 \"Unprocessable Entity\" has extra `errors` field to list all the issues with specific fields.    - Empty string (\"\") is considered a different value than \"null\" or field being omitted from JSON.     Omitting the field is recommended if no value is intended.     Empty string (\"\") will be almost always rejected as invalid value.    - There are common pattern between many objects:     - **Configuration Objects**: There are many objects with common fields, namely \"id\", \"name\", \"notes\", \"created\"       and \"updated\". These entities are listed, queried, created, updated and deleted in a similar fashion.     - **Distinguished Name**: Users and Devices are identified with what is called Distinguished Names, as used in        LDAP. The distinguished format that identifies a device and a user combination is        \"CN=\\<Device ID\\>,CN=\\<username\\>,OU=\\<Identity Provider Name\\>\". Some objects have the        \"userDistinguishedName\" field, which does not include the CN for Device ID.        This identifies a user on every device.

API version: API version 18.0
Contact: appgatesdp.support@appgate.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// ApplianceAllOfNetworkingIpv6Dhcp IPv6 DHCP configuration for the NIC.
type ApplianceAllOfNetworkingIpv6Dhcp struct {
	// Whether DHCP for IPv6 is enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// Whether to use DHCP for setting IPv6 DNS settings on the Appliance.
	Dns *bool `json:"dns,omitempty"`
	// Whether to use DHCP for setting NTP on the appliance.
	Ntp *bool `json:"ntp,omitempty"`
	// Whether to use DHCP for setting MTU on the appliance.
	Mtu *bool `json:"mtu,omitempty"`
}

// NewApplianceAllOfNetworkingIpv6Dhcp instantiates a new ApplianceAllOfNetworkingIpv6Dhcp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplianceAllOfNetworkingIpv6Dhcp() *ApplianceAllOfNetworkingIpv6Dhcp {
	this := ApplianceAllOfNetworkingIpv6Dhcp{}
	return &this
}

// NewApplianceAllOfNetworkingIpv6DhcpWithDefaults instantiates a new ApplianceAllOfNetworkingIpv6Dhcp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplianceAllOfNetworkingIpv6DhcpWithDefaults() *ApplianceAllOfNetworkingIpv6Dhcp {
	this := ApplianceAllOfNetworkingIpv6Dhcp{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *ApplianceAllOfNetworkingIpv6Dhcp) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceAllOfNetworkingIpv6Dhcp) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *ApplianceAllOfNetworkingIpv6Dhcp) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *ApplianceAllOfNetworkingIpv6Dhcp) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetDns returns the Dns field value if set, zero value otherwise.
func (o *ApplianceAllOfNetworkingIpv6Dhcp) GetDns() bool {
	if o == nil || o.Dns == nil {
		var ret bool
		return ret
	}
	return *o.Dns
}

// GetDnsOk returns a tuple with the Dns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceAllOfNetworkingIpv6Dhcp) GetDnsOk() (*bool, bool) {
	if o == nil || o.Dns == nil {
		return nil, false
	}
	return o.Dns, true
}

// HasDns returns a boolean if a field has been set.
func (o *ApplianceAllOfNetworkingIpv6Dhcp) HasDns() bool {
	if o != nil && o.Dns != nil {
		return true
	}

	return false
}

// SetDns gets a reference to the given bool and assigns it to the Dns field.
func (o *ApplianceAllOfNetworkingIpv6Dhcp) SetDns(v bool) {
	o.Dns = &v
}

// GetNtp returns the Ntp field value if set, zero value otherwise.
func (o *ApplianceAllOfNetworkingIpv6Dhcp) GetNtp() bool {
	if o == nil || o.Ntp == nil {
		var ret bool
		return ret
	}
	return *o.Ntp
}

// GetNtpOk returns a tuple with the Ntp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceAllOfNetworkingIpv6Dhcp) GetNtpOk() (*bool, bool) {
	if o == nil || o.Ntp == nil {
		return nil, false
	}
	return o.Ntp, true
}

// HasNtp returns a boolean if a field has been set.
func (o *ApplianceAllOfNetworkingIpv6Dhcp) HasNtp() bool {
	if o != nil && o.Ntp != nil {
		return true
	}

	return false
}

// SetNtp gets a reference to the given bool and assigns it to the Ntp field.
func (o *ApplianceAllOfNetworkingIpv6Dhcp) SetNtp(v bool) {
	o.Ntp = &v
}

// GetMtu returns the Mtu field value if set, zero value otherwise.
func (o *ApplianceAllOfNetworkingIpv6Dhcp) GetMtu() bool {
	if o == nil || o.Mtu == nil {
		var ret bool
		return ret
	}
	return *o.Mtu
}

// GetMtuOk returns a tuple with the Mtu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceAllOfNetworkingIpv6Dhcp) GetMtuOk() (*bool, bool) {
	if o == nil || o.Mtu == nil {
		return nil, false
	}
	return o.Mtu, true
}

// HasMtu returns a boolean if a field has been set.
func (o *ApplianceAllOfNetworkingIpv6Dhcp) HasMtu() bool {
	if o != nil && o.Mtu != nil {
		return true
	}

	return false
}

// SetMtu gets a reference to the given bool and assigns it to the Mtu field.
func (o *ApplianceAllOfNetworkingIpv6Dhcp) SetMtu(v bool) {
	o.Mtu = &v
}

func (o ApplianceAllOfNetworkingIpv6Dhcp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.Dns != nil {
		toSerialize["dns"] = o.Dns
	}
	if o.Ntp != nil {
		toSerialize["ntp"] = o.Ntp
	}
	if o.Mtu != nil {
		toSerialize["mtu"] = o.Mtu
	}
	return json.Marshal(toSerialize)
}

type NullableApplianceAllOfNetworkingIpv6Dhcp struct {
	value *ApplianceAllOfNetworkingIpv6Dhcp
	isSet bool
}

func (v NullableApplianceAllOfNetworkingIpv6Dhcp) Get() *ApplianceAllOfNetworkingIpv6Dhcp {
	return v.value
}

func (v *NullableApplianceAllOfNetworkingIpv6Dhcp) Set(val *ApplianceAllOfNetworkingIpv6Dhcp) {
	v.value = val
	v.isSet = true
}

func (v NullableApplianceAllOfNetworkingIpv6Dhcp) IsSet() bool {
	return v.isSet
}

func (v *NullableApplianceAllOfNetworkingIpv6Dhcp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplianceAllOfNetworkingIpv6Dhcp(val *ApplianceAllOfNetworkingIpv6Dhcp) *NullableApplianceAllOfNetworkingIpv6Dhcp {
	return &NullableApplianceAllOfNetworkingIpv6Dhcp{value: val, isSet: true}
}

func (v NullableApplianceAllOfNetworkingIpv6Dhcp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplianceAllOfNetworkingIpv6Dhcp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
