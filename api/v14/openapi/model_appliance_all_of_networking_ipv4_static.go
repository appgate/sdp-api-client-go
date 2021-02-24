/*
 * Appgate SDP Controller REST API
 *
 * # About   This specification documents the REST API calls for the Appgate SDP Controller.    Please refer to the Integration chapter in the manual or contact Appgate support with any questions about   this functionality. # Getting Started   Requirements for API scripting:   - Access to the Admin Interface (default port 8443) of a Controller appliance.     (https://sdphelp.appgate.com/adminguide/appliances-admin-interface-configure.html)   - An API user with relevant permissions.     (https://sdphelp.appgate.com/adminguide/administrative-roles-configure.html)   - In order to use the simple login API, Admin MFA must be disabled or the API user must be excluded.     (https://sdphelp.appgate.com/adminguide/mfa-for-admins.html) # Base path   HTTPS requests must be sent to the Admin Interface hostname and port, with **_/admin** path.    For example: **https://appgate.company.com:8443/admin**    All requests must have the **Accept** header as:    **application/vnd.appgate.peer-v14+json** # API Conventions   API conventions are  important to understand and follow strictly.    - While updating objects (via PUT), entire object must be sent with all fields.     - For example, if in order to add a remedy method to the condition below:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": []       }       ```     - send the entire object with updated and non-updated fields:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": [{\"type\": \"DisplayMessage\", \"message\": \"test message\"}]       }       ```    - In case Controller returns an error (non-2xx HTTP status code), response body is JSON.     The \"message\" field contains information about the error.     HTTP 422 \"Unprocessable Entity\" has extra `errors` field to list all the issues with specific fields.    - Empty string (\"\") is considered a different value than \"null\" or field being omitted from JSON.     Omitting the field is recommend if no value is intended.     Empty string (\"\") will be almost always rejected as invalid value.    - There are common pattern between many objects:     - **Configuration Objects**: There are many objects with common fields, namely \"id\", \"name\", \"notes\", \"created\"       and \"updated\". These entities are listed, queried, created, updated and deleted in a similar fashion.     - **Distinguished Name**: Users and Devices are identified with what is called Distinguished Names, as used in        LDAP. The distinguished format that identifies a device and a user combination is        \"CN=\\<Device ID\\>,CN=\\<username\\>,OU=\\<Identity Provider Name\\>\". Some objects have the        \"userDistinguishedName\" field, which does not include the CN for Device ID.        This identifies a user on every device.
 *
 * API version: API version 14
 * Contact: appgatesdp.support@appgate.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"encoding/json"
)

// ApplianceAllOfNetworkingIpv4Static struct for ApplianceAllOfNetworkingIpv4Static
type ApplianceAllOfNetworkingIpv4Static struct {
	// IPv4 Address of the network interface.
	Address *string `json:"address,omitempty"`
	// Netmask of the network interface.
	Netmask *int32 `json:"netmask,omitempty"`
	// NIC hostname.
	Hostname *string `json:"hostname,omitempty"`
	// Enable SNAT on this IP.
	Snat *bool `json:"snat,omitempty"`
}

// NewApplianceAllOfNetworkingIpv4Static instantiates a new ApplianceAllOfNetworkingIpv4Static object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplianceAllOfNetworkingIpv4Static() *ApplianceAllOfNetworkingIpv4Static {
	this := ApplianceAllOfNetworkingIpv4Static{}
	return &this
}

// NewApplianceAllOfNetworkingIpv4StaticWithDefaults instantiates a new ApplianceAllOfNetworkingIpv4Static object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplianceAllOfNetworkingIpv4StaticWithDefaults() *ApplianceAllOfNetworkingIpv4Static {
	this := ApplianceAllOfNetworkingIpv4Static{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *ApplianceAllOfNetworkingIpv4Static) GetAddress() string {
	if o == nil || o.Address == nil {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceAllOfNetworkingIpv4Static) GetAddressOk() (*string, bool) {
	if o == nil || o.Address == nil {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *ApplianceAllOfNetworkingIpv4Static) HasAddress() bool {
	if o != nil && o.Address != nil {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *ApplianceAllOfNetworkingIpv4Static) SetAddress(v string) {
	o.Address = &v
}

// GetNetmask returns the Netmask field value if set, zero value otherwise.
func (o *ApplianceAllOfNetworkingIpv4Static) GetNetmask() int32 {
	if o == nil || o.Netmask == nil {
		var ret int32
		return ret
	}
	return *o.Netmask
}

// GetNetmaskOk returns a tuple with the Netmask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceAllOfNetworkingIpv4Static) GetNetmaskOk() (*int32, bool) {
	if o == nil || o.Netmask == nil {
		return nil, false
	}
	return o.Netmask, true
}

// HasNetmask returns a boolean if a field has been set.
func (o *ApplianceAllOfNetworkingIpv4Static) HasNetmask() bool {
	if o != nil && o.Netmask != nil {
		return true
	}

	return false
}

// SetNetmask gets a reference to the given int32 and assigns it to the Netmask field.
func (o *ApplianceAllOfNetworkingIpv4Static) SetNetmask(v int32) {
	o.Netmask = &v
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *ApplianceAllOfNetworkingIpv4Static) GetHostname() string {
	if o == nil || o.Hostname == nil {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceAllOfNetworkingIpv4Static) GetHostnameOk() (*string, bool) {
	if o == nil || o.Hostname == nil {
		return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *ApplianceAllOfNetworkingIpv4Static) HasHostname() bool {
	if o != nil && o.Hostname != nil {
		return true
	}

	return false
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *ApplianceAllOfNetworkingIpv4Static) SetHostname(v string) {
	o.Hostname = &v
}

// GetSnat returns the Snat field value if set, zero value otherwise.
func (o *ApplianceAllOfNetworkingIpv4Static) GetSnat() bool {
	if o == nil || o.Snat == nil {
		var ret bool
		return ret
	}
	return *o.Snat
}

// GetSnatOk returns a tuple with the Snat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceAllOfNetworkingIpv4Static) GetSnatOk() (*bool, bool) {
	if o == nil || o.Snat == nil {
		return nil, false
	}
	return o.Snat, true
}

// HasSnat returns a boolean if a field has been set.
func (o *ApplianceAllOfNetworkingIpv4Static) HasSnat() bool {
	if o != nil && o.Snat != nil {
		return true
	}

	return false
}

// SetSnat gets a reference to the given bool and assigns it to the Snat field.
func (o *ApplianceAllOfNetworkingIpv4Static) SetSnat(v bool) {
	o.Snat = &v
}

func (o ApplianceAllOfNetworkingIpv4Static) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Address != nil {
		toSerialize["address"] = o.Address
	}
	if o.Netmask != nil {
		toSerialize["netmask"] = o.Netmask
	}
	if o.Hostname != nil {
		toSerialize["hostname"] = o.Hostname
	}
	if o.Snat != nil {
		toSerialize["snat"] = o.Snat
	}
	return json.Marshal(toSerialize)
}

type NullableApplianceAllOfNetworkingIpv4Static struct {
	value *ApplianceAllOfNetworkingIpv4Static
	isSet bool
}

func (v NullableApplianceAllOfNetworkingIpv4Static) Get() *ApplianceAllOfNetworkingIpv4Static {
	return v.value
}

func (v *NullableApplianceAllOfNetworkingIpv4Static) Set(val *ApplianceAllOfNetworkingIpv4Static) {
	v.value = val
	v.isSet = true
}

func (v NullableApplianceAllOfNetworkingIpv4Static) IsSet() bool {
	return v.isSet
}

func (v *NullableApplianceAllOfNetworkingIpv4Static) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplianceAllOfNetworkingIpv4Static(val *ApplianceAllOfNetworkingIpv4Static) *NullableApplianceAllOfNetworkingIpv4Static {
	return &NullableApplianceAllOfNetworkingIpv4Static{value: val, isSet: true}
}

func (v NullableApplianceAllOfNetworkingIpv4Static) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplianceAllOfNetworkingIpv4Static) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
