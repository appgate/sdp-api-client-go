/*
Appgate SDP Controller REST API

# About   This specification documents the REST API calls for the Appgate SDP Controller.    Please refer to the REST API chapter in the manual or contact Appgate support with any questions about   this functionality. # Getting Started   Requirements for API scripting:   - Access to the Admin/API TLS Connection (default port 8443) of a Controller appliance.     (https://sdphelp.appgate.com/adminguide/appliance-function-configure.html?anchor=admin-api)   - An API user with relevant permissions.     (https://sdphelp.appgate.com/adminguide/administrative-roles-configure.html)   - In order to use the simple login API, Admin MFA must be disabled or the API user must be excluded.     (https://sdphelp.appgate.com/adminguide/mfa-for-admins.html) # Base path   HTTPS requests must be sent to the Admin Interface hostname and port, with **_/admin** path.    For example: **https://appgate.company.com:8443/admin**    All requests must have the **Accept** header as:    **application/vnd.appgate.peer-v16+json** # API Conventions   API conventions are  important to understand and follow strictly.    - While updating objects (via PUT), entire object must be sent with all fields.     - For example, in order to add a remedy method to the condition below:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": []       }       ```     - send the entire object with updated and non-updated fields:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": [{\"type\": \"DisplayMessage\", \"message\": \"test message\"}]       }       ```    - In case Controller returns an error (non-2xx HTTP status code), response body is JSON.     The \"message\" field contains information about the error.     HTTP 422 \"Unprocessable Entity\" has extra `errors` field to list all the issues with specific fields.    - Empty string (\"\") is considered a different value than \"null\" or field being omitted from JSON.     Omitting the field is recommend if no value is intended.     Empty string (\"\") will be almost always rejected as invalid value.    - There are common pattern between many objects:     - **Configuration Objects**: There are many objects with common fields, namely \"id\", \"name\", \"notes\", \"created\"       and \"updated\". These entities are listed, queried, created, updated and deleted in a similar fashion.     - **Distinguished Name**: Users and Devices are identified with what is called Distinguished Names, as used in        LDAP. The distinguished format that identifies a device and a user combination is        \"CN=\\<Device ID\\>,CN=\\<username\\>,OU=\\<Identity Provider Name\\>\". Some objects have the        \"userDistinguishedName\" field, which does not include the CN for Device ID.        This identifies a user on every device.

API version: API version 16.3
Contact: appgatesdp.support@appgate.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// ApplianceAllOfNetworkingRoutes struct for ApplianceAllOfNetworkingRoutes
type ApplianceAllOfNetworkingRoutes struct {
	// Address to route.
	Address string `json:"address"`
	// Netmask for the subnet to route.
	Netmask int32 `json:"netmask"`
	// Gateway to use for routing.
	Gateway *string `json:"gateway,omitempty"`
	// NIC name to use for routing.
	Nic *string `json:"nic,omitempty"`
}

// NewApplianceAllOfNetworkingRoutes instantiates a new ApplianceAllOfNetworkingRoutes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplianceAllOfNetworkingRoutes(address string, netmask int32) *ApplianceAllOfNetworkingRoutes {
	this := ApplianceAllOfNetworkingRoutes{}
	this.Address = address
	this.Netmask = netmask
	return &this
}

// NewApplianceAllOfNetworkingRoutesWithDefaults instantiates a new ApplianceAllOfNetworkingRoutes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplianceAllOfNetworkingRoutesWithDefaults() *ApplianceAllOfNetworkingRoutes {
	this := ApplianceAllOfNetworkingRoutes{}
	return &this
}

// GetAddress returns the Address field value
func (o *ApplianceAllOfNetworkingRoutes) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *ApplianceAllOfNetworkingRoutes) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *ApplianceAllOfNetworkingRoutes) SetAddress(v string) {
	o.Address = v
}

// GetNetmask returns the Netmask field value
func (o *ApplianceAllOfNetworkingRoutes) GetNetmask() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Netmask
}

// GetNetmaskOk returns a tuple with the Netmask field value
// and a boolean to check if the value has been set.
func (o *ApplianceAllOfNetworkingRoutes) GetNetmaskOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Netmask, true
}

// SetNetmask sets field value
func (o *ApplianceAllOfNetworkingRoutes) SetNetmask(v int32) {
	o.Netmask = v
}

// GetGateway returns the Gateway field value if set, zero value otherwise.
func (o *ApplianceAllOfNetworkingRoutes) GetGateway() string {
	if o == nil || o.Gateway == nil {
		var ret string
		return ret
	}
	return *o.Gateway
}

// GetGatewayOk returns a tuple with the Gateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceAllOfNetworkingRoutes) GetGatewayOk() (*string, bool) {
	if o == nil || o.Gateway == nil {
		return nil, false
	}
	return o.Gateway, true
}

// HasGateway returns a boolean if a field has been set.
func (o *ApplianceAllOfNetworkingRoutes) HasGateway() bool {
	if o != nil && o.Gateway != nil {
		return true
	}

	return false
}

// SetGateway gets a reference to the given string and assigns it to the Gateway field.
func (o *ApplianceAllOfNetworkingRoutes) SetGateway(v string) {
	o.Gateway = &v
}

// GetNic returns the Nic field value if set, zero value otherwise.
func (o *ApplianceAllOfNetworkingRoutes) GetNic() string {
	if o == nil || o.Nic == nil {
		var ret string
		return ret
	}
	return *o.Nic
}

// GetNicOk returns a tuple with the Nic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceAllOfNetworkingRoutes) GetNicOk() (*string, bool) {
	if o == nil || o.Nic == nil {
		return nil, false
	}
	return o.Nic, true
}

// HasNic returns a boolean if a field has been set.
func (o *ApplianceAllOfNetworkingRoutes) HasNic() bool {
	if o != nil && o.Nic != nil {
		return true
	}

	return false
}

// SetNic gets a reference to the given string and assigns it to the Nic field.
func (o *ApplianceAllOfNetworkingRoutes) SetNic(v string) {
	o.Nic = &v
}

func (o ApplianceAllOfNetworkingRoutes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["address"] = o.Address
	}
	if true {
		toSerialize["netmask"] = o.Netmask
	}
	if o.Gateway != nil {
		toSerialize["gateway"] = o.Gateway
	}
	if o.Nic != nil {
		toSerialize["nic"] = o.Nic
	}
	return json.Marshal(toSerialize)
}

type NullableApplianceAllOfNetworkingRoutes struct {
	value *ApplianceAllOfNetworkingRoutes
	isSet bool
}

func (v NullableApplianceAllOfNetworkingRoutes) Get() *ApplianceAllOfNetworkingRoutes {
	return v.value
}

func (v *NullableApplianceAllOfNetworkingRoutes) Set(val *ApplianceAllOfNetworkingRoutes) {
	v.value = val
	v.isSet = true
}

func (v NullableApplianceAllOfNetworkingRoutes) IsSet() bool {
	return v.isSet
}

func (v *NullableApplianceAllOfNetworkingRoutes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplianceAllOfNetworkingRoutes(val *ApplianceAllOfNetworkingRoutes) *NullableApplianceAllOfNetworkingRoutes {
	return &NullableApplianceAllOfNetworkingRoutes{value: val, isSet: true}
}

func (v NullableApplianceAllOfNetworkingRoutes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplianceAllOfNetworkingRoutes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
