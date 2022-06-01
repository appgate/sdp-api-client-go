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

// ApplianceAllOfNetworkingIpv4 IPv4 settings for this NIC.
type ApplianceAllOfNetworkingIpv4 struct {
	Dhcp *ApplianceAllOfNetworkingIpv4Dhcp `json:"dhcp,omitempty"`
	// IPv4 static NIC configuration for the NIC.
	Static *[]ApplianceAllOfNetworkingIpv4Static `json:"static,omitempty"`
	// Virtual IP to use for IPv4.
	VirtualIp *string `json:"virtualIp,omitempty"`
}

// NewApplianceAllOfNetworkingIpv4 instantiates a new ApplianceAllOfNetworkingIpv4 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplianceAllOfNetworkingIpv4() *ApplianceAllOfNetworkingIpv4 {
	this := ApplianceAllOfNetworkingIpv4{}
	return &this
}

// NewApplianceAllOfNetworkingIpv4WithDefaults instantiates a new ApplianceAllOfNetworkingIpv4 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplianceAllOfNetworkingIpv4WithDefaults() *ApplianceAllOfNetworkingIpv4 {
	this := ApplianceAllOfNetworkingIpv4{}
	return &this
}

// GetDhcp returns the Dhcp field value if set, zero value otherwise.
func (o *ApplianceAllOfNetworkingIpv4) GetDhcp() ApplianceAllOfNetworkingIpv4Dhcp {
	if o == nil || o.Dhcp == nil {
		var ret ApplianceAllOfNetworkingIpv4Dhcp
		return ret
	}
	return *o.Dhcp
}

// GetDhcpOk returns a tuple with the Dhcp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceAllOfNetworkingIpv4) GetDhcpOk() (*ApplianceAllOfNetworkingIpv4Dhcp, bool) {
	if o == nil || o.Dhcp == nil {
		return nil, false
	}
	return o.Dhcp, true
}

// HasDhcp returns a boolean if a field has been set.
func (o *ApplianceAllOfNetworkingIpv4) HasDhcp() bool {
	if o != nil && o.Dhcp != nil {
		return true
	}

	return false
}

// SetDhcp gets a reference to the given ApplianceAllOfNetworkingIpv4Dhcp and assigns it to the Dhcp field.
func (o *ApplianceAllOfNetworkingIpv4) SetDhcp(v ApplianceAllOfNetworkingIpv4Dhcp) {
	o.Dhcp = &v
}

// GetStatic returns the Static field value if set, zero value otherwise.
func (o *ApplianceAllOfNetworkingIpv4) GetStatic() []ApplianceAllOfNetworkingIpv4Static {
	if o == nil || o.Static == nil {
		var ret []ApplianceAllOfNetworkingIpv4Static
		return ret
	}
	return *o.Static
}

// GetStaticOk returns a tuple with the Static field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceAllOfNetworkingIpv4) GetStaticOk() (*[]ApplianceAllOfNetworkingIpv4Static, bool) {
	if o == nil || o.Static == nil {
		return nil, false
	}
	return o.Static, true
}

// HasStatic returns a boolean if a field has been set.
func (o *ApplianceAllOfNetworkingIpv4) HasStatic() bool {
	if o != nil && o.Static != nil {
		return true
	}

	return false
}

// SetStatic gets a reference to the given []ApplianceAllOfNetworkingIpv4Static and assigns it to the Static field.
func (o *ApplianceAllOfNetworkingIpv4) SetStatic(v []ApplianceAllOfNetworkingIpv4Static) {
	o.Static = &v
}

// GetVirtualIp returns the VirtualIp field value if set, zero value otherwise.
func (o *ApplianceAllOfNetworkingIpv4) GetVirtualIp() string {
	if o == nil || o.VirtualIp == nil {
		var ret string
		return ret
	}
	return *o.VirtualIp
}

// GetVirtualIpOk returns a tuple with the VirtualIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceAllOfNetworkingIpv4) GetVirtualIpOk() (*string, bool) {
	if o == nil || o.VirtualIp == nil {
		return nil, false
	}
	return o.VirtualIp, true
}

// HasVirtualIp returns a boolean if a field has been set.
func (o *ApplianceAllOfNetworkingIpv4) HasVirtualIp() bool {
	if o != nil && o.VirtualIp != nil {
		return true
	}

	return false
}

// SetVirtualIp gets a reference to the given string and assigns it to the VirtualIp field.
func (o *ApplianceAllOfNetworkingIpv4) SetVirtualIp(v string) {
	o.VirtualIp = &v
}

func (o ApplianceAllOfNetworkingIpv4) MarshalJSON() ([]byte, error) {
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

type NullableApplianceAllOfNetworkingIpv4 struct {
	value *ApplianceAllOfNetworkingIpv4
	isSet bool
}

func (v NullableApplianceAllOfNetworkingIpv4) Get() *ApplianceAllOfNetworkingIpv4 {
	return v.value
}

func (v *NullableApplianceAllOfNetworkingIpv4) Set(val *ApplianceAllOfNetworkingIpv4) {
	v.value = val
	v.isSet = true
}

func (v NullableApplianceAllOfNetworkingIpv4) IsSet() bool {
	return v.isSet
}

func (v *NullableApplianceAllOfNetworkingIpv4) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplianceAllOfNetworkingIpv4(val *ApplianceAllOfNetworkingIpv4) *NullableApplianceAllOfNetworkingIpv4 {
	return &NullableApplianceAllOfNetworkingIpv4{value: val, isSet: true}
}

func (v NullableApplianceAllOfNetworkingIpv4) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplianceAllOfNetworkingIpv4) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
