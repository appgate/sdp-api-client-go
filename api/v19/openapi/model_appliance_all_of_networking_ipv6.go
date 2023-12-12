/*
Appgate SDP Controller REST API

# About   This specification documents the REST API calls for the Appgate SDP Controller.    Please refer to the REST API chapter in the manual or contact Appgate support with any questions about   this functionality. # Getting Started   Requirements for API scripting:   - Access to the Admin/API TLS Connection (default port 8443) of a Controller appliance.     (https://sdphelp.appgate.com/adminguide/appliance-function-configure.html?anchor=admin-api)   - An API user with relevant permissions.     (https://sdphelp.appgate.com/adminguide/administrative-roles-configure.html)   - In order to use the simple login API, Admin MFA must be disabled or the API user must be excluded.     (https://sdphelp.appgate.com/adminguide/mfa-for-admins.html) # Base path   HTTPS requests must be sent to the Admin Interface hostname and port, with **_/admin** path.    For example: **https://appgate.company.com:8443/admin**    All requests must have the **Accept** header as:    **application/vnd.appgate.peer-v19+json** # API Conventions   API conventions are  important to understand and follow strictly.    - While updating objects (via PUT), entire object must be sent with all fields.     - For example, in order to add a remedy method to the condition below:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": []       }       ```     - send the entire object with updated and non-updated fields:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": [{\"type\": \"DisplayMessage\", \"message\": \"test message\"}]       }       ```    - In case Controller returns an error (non-2xx HTTP status code), response body is JSON.     The \"message\" field contains information about the error.     HTTP 422 \"Unprocessable Entity\" has extra `errors` field to list all the issues with specific fields.    - Empty string (\"\") is considered a different value than \"null\" or field being omitted from JSON.     Omitting the field is recommended if no value is intended.     Empty string (\"\") will be almost always rejected as invalid value.    - There are common pattern between many objects:     - **Configuration Objects**: There are many objects with common fields, namely \"id\", \"name\", \"notes\", \"created\"       and \"updated\". These entities are listed, queried, created, updated and deleted in a similar fashion.     - **Distinguished Name**: Users and Devices are identified with what is called Distinguished Names, as used in        LDAP. The distinguished format that identifies a device and a user combination is        \"CN=\\<Device ID\\>,CN=\\<username\\>,OU=\\<Identity Provider Name\\>\". Some objects have the        \"userDistinguishedName\" field, which does not include the CN for Device ID.        This identifies a user on every device.

API version: API version 19.2
Contact: appgatesdp.support@appgate.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// ApplianceAllOfNetworkingIpv6 IPv6 settings for this NIC.
type ApplianceAllOfNetworkingIpv6 struct {
	Dhcp *ApplianceAllOfNetworkingIpv6Dhcp `json:"dhcp,omitempty"`
	// IPv6 static NIC configuration for the NIC.
	Static []ApplianceAllOfNetworkingIpv6Static `json:"static,omitempty"`
	// Virtual IP to use for IPv6.
	VirtualIp *string `json:"virtualIp,omitempty"`
}

// NewApplianceAllOfNetworkingIpv6 instantiates a new ApplianceAllOfNetworkingIpv6 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplianceAllOfNetworkingIpv6() *ApplianceAllOfNetworkingIpv6 {
	this := ApplianceAllOfNetworkingIpv6{}
	return &this
}

// NewApplianceAllOfNetworkingIpv6WithDefaults instantiates a new ApplianceAllOfNetworkingIpv6 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplianceAllOfNetworkingIpv6WithDefaults() *ApplianceAllOfNetworkingIpv6 {
	this := ApplianceAllOfNetworkingIpv6{}
	return &this
}

// GetDhcp returns the Dhcp field value if set, zero value otherwise.
func (o *ApplianceAllOfNetworkingIpv6) GetDhcp() ApplianceAllOfNetworkingIpv6Dhcp {
	if o == nil || o.Dhcp == nil {
		var ret ApplianceAllOfNetworkingIpv6Dhcp
		return ret
	}
	return *o.Dhcp
}

// GetDhcpOk returns a tuple with the Dhcp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceAllOfNetworkingIpv6) GetDhcpOk() (*ApplianceAllOfNetworkingIpv6Dhcp, bool) {
	if o == nil || o.Dhcp == nil {
		return nil, false
	}
	return o.Dhcp, true
}

// HasDhcp returns a boolean if a field has been set.
func (o *ApplianceAllOfNetworkingIpv6) HasDhcp() bool {
	if o != nil && o.Dhcp != nil {
		return true
	}

	return false
}

// SetDhcp gets a reference to the given ApplianceAllOfNetworkingIpv6Dhcp and assigns it to the Dhcp field.
func (o *ApplianceAllOfNetworkingIpv6) SetDhcp(v ApplianceAllOfNetworkingIpv6Dhcp) {
	o.Dhcp = &v
}

// GetStatic returns the Static field value if set, zero value otherwise.
func (o *ApplianceAllOfNetworkingIpv6) GetStatic() []ApplianceAllOfNetworkingIpv6Static {
	if o == nil || o.Static == nil {
		var ret []ApplianceAllOfNetworkingIpv6Static
		return ret
	}
	return o.Static
}

// GetStaticOk returns a tuple with the Static field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceAllOfNetworkingIpv6) GetStaticOk() ([]ApplianceAllOfNetworkingIpv6Static, bool) {
	if o == nil || o.Static == nil {
		return nil, false
	}
	return o.Static, true
}

// HasStatic returns a boolean if a field has been set.
func (o *ApplianceAllOfNetworkingIpv6) HasStatic() bool {
	if o != nil && o.Static != nil {
		return true
	}

	return false
}

// SetStatic gets a reference to the given []ApplianceAllOfNetworkingIpv6Static and assigns it to the Static field.
func (o *ApplianceAllOfNetworkingIpv6) SetStatic(v []ApplianceAllOfNetworkingIpv6Static) {
	o.Static = v
}

// GetVirtualIp returns the VirtualIp field value if set, zero value otherwise.
func (o *ApplianceAllOfNetworkingIpv6) GetVirtualIp() string {
	if o == nil || o.VirtualIp == nil {
		var ret string
		return ret
	}
	return *o.VirtualIp
}

// GetVirtualIpOk returns a tuple with the VirtualIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceAllOfNetworkingIpv6) GetVirtualIpOk() (*string, bool) {
	if o == nil || o.VirtualIp == nil {
		return nil, false
	}
	return o.VirtualIp, true
}

// HasVirtualIp returns a boolean if a field has been set.
func (o *ApplianceAllOfNetworkingIpv6) HasVirtualIp() bool {
	if o != nil && o.VirtualIp != nil {
		return true
	}

	return false
}

// SetVirtualIp gets a reference to the given string and assigns it to the VirtualIp field.
func (o *ApplianceAllOfNetworkingIpv6) SetVirtualIp(v string) {
	o.VirtualIp = &v
}

func (o ApplianceAllOfNetworkingIpv6) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Dhcp != nil {
		toSerialize["dhcp"] = o.Dhcp
	}
	if o.Static != nil {
		toSerialize["static"] = o.Static
	}
	if o.VirtualIp != nil {
		toSerialize["virtualIp"] = o.VirtualIp
	}
	return json.Marshal(toSerialize)
}

type NullableApplianceAllOfNetworkingIpv6 struct {
	value *ApplianceAllOfNetworkingIpv6
	isSet bool
}

func (v NullableApplianceAllOfNetworkingIpv6) Get() *ApplianceAllOfNetworkingIpv6 {
	return v.value
}

func (v *NullableApplianceAllOfNetworkingIpv6) Set(val *ApplianceAllOfNetworkingIpv6) {
	v.value = val
	v.isSet = true
}

func (v NullableApplianceAllOfNetworkingIpv6) IsSet() bool {
	return v.isSet
}

func (v *NullableApplianceAllOfNetworkingIpv6) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplianceAllOfNetworkingIpv6(val *ApplianceAllOfNetworkingIpv6) *NullableApplianceAllOfNetworkingIpv6 {
	return &NullableApplianceAllOfNetworkingIpv6{value: val, isSet: true}
}

func (v NullableApplianceAllOfNetworkingIpv6) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplianceAllOfNetworkingIpv6) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
