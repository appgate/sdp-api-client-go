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

// ApplianceAllOfHealthcheckServer Healthcheck Server configuration.
type ApplianceAllOfHealthcheckServer struct {
	// Whether the Healthcheck Server is enabled on this appliance or not.
	Enabled *bool `json:"enabled,omitempty"`
	// Port to connect for Healthcheck Server.
	Port *int32 `json:"port,omitempty"`
	// Source configuration to allow via iptables.
	AllowSources *[]interface{} `json:"allowSources,omitempty"`
}

// NewApplianceAllOfHealthcheckServer instantiates a new ApplianceAllOfHealthcheckServer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplianceAllOfHealthcheckServer() *ApplianceAllOfHealthcheckServer {
	this := ApplianceAllOfHealthcheckServer{}
	var enabled bool = false
	this.Enabled = &enabled
	var port int32 = 5555
	this.Port = &port
	return &this
}

// NewApplianceAllOfHealthcheckServerWithDefaults instantiates a new ApplianceAllOfHealthcheckServer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplianceAllOfHealthcheckServerWithDefaults() *ApplianceAllOfHealthcheckServer {
	this := ApplianceAllOfHealthcheckServer{}
	var enabled bool = false
	this.Enabled = &enabled
	var port int32 = 5555
	this.Port = &port
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *ApplianceAllOfHealthcheckServer) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceAllOfHealthcheckServer) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *ApplianceAllOfHealthcheckServer) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *ApplianceAllOfHealthcheckServer) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *ApplianceAllOfHealthcheckServer) GetPort() int32 {
	if o == nil || o.Port == nil {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceAllOfHealthcheckServer) GetPortOk() (*int32, bool) {
	if o == nil || o.Port == nil {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *ApplianceAllOfHealthcheckServer) HasPort() bool {
	if o != nil && o.Port != nil {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *ApplianceAllOfHealthcheckServer) SetPort(v int32) {
	o.Port = &v
}

// GetAllowSources returns the AllowSources field value if set, zero value otherwise.
func (o *ApplianceAllOfHealthcheckServer) GetAllowSources() []interface{} {
	if o == nil || o.AllowSources == nil {
		var ret []interface{}
		return ret
	}
	return *o.AllowSources
}

// GetAllowSourcesOk returns a tuple with the AllowSources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceAllOfHealthcheckServer) GetAllowSourcesOk() (*[]interface{}, bool) {
	if o == nil || o.AllowSources == nil {
		return nil, false
	}
	return o.AllowSources, true
}

// HasAllowSources returns a boolean if a field has been set.
func (o *ApplianceAllOfHealthcheckServer) HasAllowSources() bool {
	if o != nil && o.AllowSources != nil {
		return true
	}

	return false
}

// SetAllowSources gets a reference to the given []interface{} and assigns it to the AllowSources field.
func (o *ApplianceAllOfHealthcheckServer) SetAllowSources(v []interface{}) {
	o.AllowSources = &v
}

func (o ApplianceAllOfHealthcheckServer) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.Port != nil {
		toSerialize["port"] = o.Port
	}
	if o.AllowSources != nil {
		toSerialize["allowSources"] = o.AllowSources
	}
	return json.Marshal(toSerialize)
}

type NullableApplianceAllOfHealthcheckServer struct {
	value *ApplianceAllOfHealthcheckServer
	isSet bool
}

func (v NullableApplianceAllOfHealthcheckServer) Get() *ApplianceAllOfHealthcheckServer {
	return v.value
}

func (v *NullableApplianceAllOfHealthcheckServer) Set(val *ApplianceAllOfHealthcheckServer) {
	v.value = val
	v.isSet = true
}

func (v NullableApplianceAllOfHealthcheckServer) IsSet() bool {
	return v.isSet
}

func (v *NullableApplianceAllOfHealthcheckServer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplianceAllOfHealthcheckServer(val *ApplianceAllOfHealthcheckServer) *NullableApplianceAllOfHealthcheckServer {
	return &NullableApplianceAllOfHealthcheckServer{value: val, isSet: true}
}

func (v NullableApplianceAllOfHealthcheckServer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplianceAllOfHealthcheckServer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
