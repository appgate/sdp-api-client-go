/*
Appgate SDP Controller REST API

# About   This specification documents the REST API calls for the Appgate SDP Controller.    Please refer to the REST API chapter in the manual or contact Appgate support with any questions about   this functionality. # Getting Started   Requirements for API scripting:   - Access to the Admin/API TLS Connection (default port 8443) of a Controller appliance.     (https://sdphelp.appgate.com/adminguide/appliance-function-configure.html?anchor=admin-api)   - An API user with relevant permissions.     (https://sdphelp.appgate.com/adminguide/administrative-roles-configure.html)   - In order to use the simple login API, Admin MFA must be disabled or the API user must be excluded.     (https://sdphelp.appgate.com/adminguide/mfa-for-admins.html) # Base path   HTTPS requests must be sent to the Admin Interface hostname and port, with **_/admin** path.    For example: **https://appgate.company.com:8443/admin**    All requests must have the **Accept** header as:    **application/vnd.appgate.peer-v22+json**    An exception is made for the **_/admin/version** endpoint which instead expects an **application/json** Accept header. # API Conventions   API conventions are  important to understand and follow strictly.    - While updating objects (via PUT), entire object must be sent with all fields.     - For example, in order to add a remedy method to the condition below:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": []       }       ```     - send the entire object with updated and non-updated fields:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": [{\"type\": \"DisplayMessage\", \"message\": \"test message\"}]       }       ```    - In case Controller returns an error (non-2xx HTTP status code), response body is JSON.     The \"message\" field contains information about the error.     HTTP 422 \"Unprocessable Entity\" has extra `errors` field to list all the issues with specific fields.    - Empty string (\"\") is considered a different value than \"null\" or field being omitted from JSON.     Omitting the field is recommended if no value is intended.     Empty string (\"\") will be almost always rejected as invalid value.    - There are common pattern between many objects:     - **Configuration Objects**: There are many objects with common fields, namely \"id\", \"name\", \"notes\", \"created\"       and \"updated\". These entities are listed, queried, created, updated and deleted in a similar fashion.     - **Distinguished Name**: Users and Devices are identified with what is called Distinguished Names, as used in        LDAP. The distinguished format that identifies a device and a user combination is        \"CN=\\<Device ID\\>,CN=\\<username\\>,OU=\\<Identity Provider Name\\>\". Some objects have the        \"userDistinguishedName\" field, which does not include the CN for Device ID.        This identifies a user on every device.

API version: API version 22.4
Contact: appgatesdp.support@appgate.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// ApplianceAllOfGatewayVpn VPN configuration.
type ApplianceAllOfGatewayVpn struct {
	// Load balancing weight.
	Weight *int32 `json:"weight,omitempty"`
	// Load balancing weight that would take effect with Local Site Detection feature.
	LocalWeight *int32 `json:"localWeight,omitempty"`
	// Destinations to allow tunnels to.
	AllowDestinations []ApplianceAllOfGatewayVpnAllowDestinations `json:"allowDestinations,omitempty"`
}

// NewApplianceAllOfGatewayVpn instantiates a new ApplianceAllOfGatewayVpn object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplianceAllOfGatewayVpn() *ApplianceAllOfGatewayVpn {
	this := ApplianceAllOfGatewayVpn{}
	var weight int32 = 100
	this.Weight = &weight
	return &this
}

// NewApplianceAllOfGatewayVpnWithDefaults instantiates a new ApplianceAllOfGatewayVpn object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplianceAllOfGatewayVpnWithDefaults() *ApplianceAllOfGatewayVpn {
	this := ApplianceAllOfGatewayVpn{}
	var weight int32 = 100
	this.Weight = &weight
	return &this
}

// GetWeight returns the Weight field value if set, zero value otherwise.
func (o *ApplianceAllOfGatewayVpn) GetWeight() int32 {
	if o == nil || o.Weight == nil {
		var ret int32
		return ret
	}
	return *o.Weight
}

// GetWeightOk returns a tuple with the Weight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceAllOfGatewayVpn) GetWeightOk() (*int32, bool) {
	if o == nil || o.Weight == nil {
		return nil, false
	}
	return o.Weight, true
}

// HasWeight returns a boolean if a field has been set.
func (o *ApplianceAllOfGatewayVpn) HasWeight() bool {
	if o != nil && o.Weight != nil {
		return true
	}

	return false
}

// SetWeight gets a reference to the given int32 and assigns it to the Weight field.
func (o *ApplianceAllOfGatewayVpn) SetWeight(v int32) {
	o.Weight = &v
}

// GetLocalWeight returns the LocalWeight field value if set, zero value otherwise.
func (o *ApplianceAllOfGatewayVpn) GetLocalWeight() int32 {
	if o == nil || o.LocalWeight == nil {
		var ret int32
		return ret
	}
	return *o.LocalWeight
}

// GetLocalWeightOk returns a tuple with the LocalWeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceAllOfGatewayVpn) GetLocalWeightOk() (*int32, bool) {
	if o == nil || o.LocalWeight == nil {
		return nil, false
	}
	return o.LocalWeight, true
}

// HasLocalWeight returns a boolean if a field has been set.
func (o *ApplianceAllOfGatewayVpn) HasLocalWeight() bool {
	if o != nil && o.LocalWeight != nil {
		return true
	}

	return false
}

// SetLocalWeight gets a reference to the given int32 and assigns it to the LocalWeight field.
func (o *ApplianceAllOfGatewayVpn) SetLocalWeight(v int32) {
	o.LocalWeight = &v
}

// GetAllowDestinations returns the AllowDestinations field value if set, zero value otherwise.
func (o *ApplianceAllOfGatewayVpn) GetAllowDestinations() []ApplianceAllOfGatewayVpnAllowDestinations {
	if o == nil || o.AllowDestinations == nil {
		var ret []ApplianceAllOfGatewayVpnAllowDestinations
		return ret
	}
	return o.AllowDestinations
}

// GetAllowDestinationsOk returns a tuple with the AllowDestinations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceAllOfGatewayVpn) GetAllowDestinationsOk() ([]ApplianceAllOfGatewayVpnAllowDestinations, bool) {
	if o == nil || o.AllowDestinations == nil {
		return nil, false
	}
	return o.AllowDestinations, true
}

// HasAllowDestinations returns a boolean if a field has been set.
func (o *ApplianceAllOfGatewayVpn) HasAllowDestinations() bool {
	if o != nil && o.AllowDestinations != nil {
		return true
	}

	return false
}

// SetAllowDestinations gets a reference to the given []ApplianceAllOfGatewayVpnAllowDestinations and assigns it to the AllowDestinations field.
func (o *ApplianceAllOfGatewayVpn) SetAllowDestinations(v []ApplianceAllOfGatewayVpnAllowDestinations) {
	o.AllowDestinations = v
}

func (o ApplianceAllOfGatewayVpn) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Weight != nil {
		toSerialize["weight"] = o.Weight
	}
	if o.LocalWeight != nil {
		toSerialize["localWeight"] = o.LocalWeight
	}
	if o.AllowDestinations != nil {
		toSerialize["allowDestinations"] = o.AllowDestinations
	}
	return json.Marshal(toSerialize)
}

type NullableApplianceAllOfGatewayVpn struct {
	value *ApplianceAllOfGatewayVpn
	isSet bool
}

func (v NullableApplianceAllOfGatewayVpn) Get() *ApplianceAllOfGatewayVpn {
	return v.value
}

func (v *NullableApplianceAllOfGatewayVpn) Set(val *ApplianceAllOfGatewayVpn) {
	v.value = val
	v.isSet = true
}

func (v NullableApplianceAllOfGatewayVpn) IsSet() bool {
	return v.isSet
}

func (v *NullableApplianceAllOfGatewayVpn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplianceAllOfGatewayVpn(val *ApplianceAllOfGatewayVpn) *NullableApplianceAllOfGatewayVpn {
	return &NullableApplianceAllOfGatewayVpn{value: val, isSet: true}
}

func (v NullableApplianceAllOfGatewayVpn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplianceAllOfGatewayVpn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
