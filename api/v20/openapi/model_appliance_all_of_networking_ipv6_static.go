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

// ApplianceAllOfNetworkingIpv6Static struct for ApplianceAllOfNetworkingIpv6Static
type ApplianceAllOfNetworkingIpv6Static struct {
	// IPv6 Address of the network interface.
	Address string `json:"address"`
	// Netmask of the network interface.
	Netmask int32 `json:"netmask"`
	// Enable SNAT on this IP.
	Snat *bool `json:"snat,omitempty"`
}

// NewApplianceAllOfNetworkingIpv6Static instantiates a new ApplianceAllOfNetworkingIpv6Static object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplianceAllOfNetworkingIpv6Static(address string, netmask int32) *ApplianceAllOfNetworkingIpv6Static {
	this := ApplianceAllOfNetworkingIpv6Static{}
	this.Address = address
	this.Netmask = netmask
	return &this
}

// NewApplianceAllOfNetworkingIpv6StaticWithDefaults instantiates a new ApplianceAllOfNetworkingIpv6Static object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplianceAllOfNetworkingIpv6StaticWithDefaults() *ApplianceAllOfNetworkingIpv6Static {
	this := ApplianceAllOfNetworkingIpv6Static{}
	return &this
}

// GetAddress returns the Address field value
func (o *ApplianceAllOfNetworkingIpv6Static) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *ApplianceAllOfNetworkingIpv6Static) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *ApplianceAllOfNetworkingIpv6Static) SetAddress(v string) {
	o.Address = v
}

// GetNetmask returns the Netmask field value
func (o *ApplianceAllOfNetworkingIpv6Static) GetNetmask() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Netmask
}

// GetNetmaskOk returns a tuple with the Netmask field value
// and a boolean to check if the value has been set.
func (o *ApplianceAllOfNetworkingIpv6Static) GetNetmaskOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Netmask, true
}

// SetNetmask sets field value
func (o *ApplianceAllOfNetworkingIpv6Static) SetNetmask(v int32) {
	o.Netmask = v
}

// GetSnat returns the Snat field value if set, zero value otherwise.
func (o *ApplianceAllOfNetworkingIpv6Static) GetSnat() bool {
	if o == nil || o.Snat == nil {
		var ret bool
		return ret
	}
	return *o.Snat
}

// GetSnatOk returns a tuple with the Snat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceAllOfNetworkingIpv6Static) GetSnatOk() (*bool, bool) {
	if o == nil || o.Snat == nil {
		return nil, false
	}
	return o.Snat, true
}

// HasSnat returns a boolean if a field has been set.
func (o *ApplianceAllOfNetworkingIpv6Static) HasSnat() bool {
	if o != nil && o.Snat != nil {
		return true
	}

	return false
}

// SetSnat gets a reference to the given bool and assigns it to the Snat field.
func (o *ApplianceAllOfNetworkingIpv6Static) SetSnat(v bool) {
	o.Snat = &v
}

func (o ApplianceAllOfNetworkingIpv6Static) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["address"] = o.Address
	}
	if true {
		toSerialize["netmask"] = o.Netmask
	}
	if o.Snat != nil {
		toSerialize["snat"] = o.Snat
	}
	return json.Marshal(toSerialize)
}

type NullableApplianceAllOfNetworkingIpv6Static struct {
	value *ApplianceAllOfNetworkingIpv6Static
	isSet bool
}

func (v NullableApplianceAllOfNetworkingIpv6Static) Get() *ApplianceAllOfNetworkingIpv6Static {
	return v.value
}

func (v *NullableApplianceAllOfNetworkingIpv6Static) Set(val *ApplianceAllOfNetworkingIpv6Static) {
	v.value = val
	v.isSet = true
}

func (v NullableApplianceAllOfNetworkingIpv6Static) IsSet() bool {
	return v.isSet
}

func (v *NullableApplianceAllOfNetworkingIpv6Static) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplianceAllOfNetworkingIpv6Static(val *ApplianceAllOfNetworkingIpv6Static) *NullableApplianceAllOfNetworkingIpv6Static {
	return &NullableApplianceAllOfNetworkingIpv6Static{value: val, isSet: true}
}

func (v NullableApplianceAllOfNetworkingIpv6Static) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplianceAllOfNetworkingIpv6Static) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
