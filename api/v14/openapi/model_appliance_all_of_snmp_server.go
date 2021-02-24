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

// ApplianceAllOfSnmpServer SNMP Server configuration.
type ApplianceAllOfSnmpServer struct {
	// Whether the SNMP Server os enabled on this appliance or not.
	Enabled *bool `json:"enabled,omitempty"`
	// TCP port for SNMP Server.
	TcpPort *int32 `json:"tcpPort,omitempty"`
	// UDP port for SNMP Server.
	UdpPort *int32 `json:"udpPort,omitempty"`
	// Raw SNMP configuration.
	SnmpdConf *string `json:"snmpd.conf,omitempty"`
	// Source configuration to allow via iptables.
	AllowSources *[]map[string]interface{} `json:"allowSources,omitempty"`
}

// NewApplianceAllOfSnmpServer instantiates a new ApplianceAllOfSnmpServer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplianceAllOfSnmpServer() *ApplianceAllOfSnmpServer {
	this := ApplianceAllOfSnmpServer{}
	var enabled bool = false
	this.Enabled = &enabled
	return &this
}

// NewApplianceAllOfSnmpServerWithDefaults instantiates a new ApplianceAllOfSnmpServer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplianceAllOfSnmpServerWithDefaults() *ApplianceAllOfSnmpServer {
	this := ApplianceAllOfSnmpServer{}
	var enabled bool = false
	this.Enabled = &enabled
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *ApplianceAllOfSnmpServer) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceAllOfSnmpServer) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *ApplianceAllOfSnmpServer) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *ApplianceAllOfSnmpServer) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetTcpPort returns the TcpPort field value if set, zero value otherwise.
func (o *ApplianceAllOfSnmpServer) GetTcpPort() int32 {
	if o == nil || o.TcpPort == nil {
		var ret int32
		return ret
	}
	return *o.TcpPort
}

// GetTcpPortOk returns a tuple with the TcpPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceAllOfSnmpServer) GetTcpPortOk() (*int32, bool) {
	if o == nil || o.TcpPort == nil {
		return nil, false
	}
	return o.TcpPort, true
}

// HasTcpPort returns a boolean if a field has been set.
func (o *ApplianceAllOfSnmpServer) HasTcpPort() bool {
	if o != nil && o.TcpPort != nil {
		return true
	}

	return false
}

// SetTcpPort gets a reference to the given int32 and assigns it to the TcpPort field.
func (o *ApplianceAllOfSnmpServer) SetTcpPort(v int32) {
	o.TcpPort = &v
}

// GetUdpPort returns the UdpPort field value if set, zero value otherwise.
func (o *ApplianceAllOfSnmpServer) GetUdpPort() int32 {
	if o == nil || o.UdpPort == nil {
		var ret int32
		return ret
	}
	return *o.UdpPort
}

// GetUdpPortOk returns a tuple with the UdpPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceAllOfSnmpServer) GetUdpPortOk() (*int32, bool) {
	if o == nil || o.UdpPort == nil {
		return nil, false
	}
	return o.UdpPort, true
}

// HasUdpPort returns a boolean if a field has been set.
func (o *ApplianceAllOfSnmpServer) HasUdpPort() bool {
	if o != nil && o.UdpPort != nil {
		return true
	}

	return false
}

// SetUdpPort gets a reference to the given int32 and assigns it to the UdpPort field.
func (o *ApplianceAllOfSnmpServer) SetUdpPort(v int32) {
	o.UdpPort = &v
}

// GetSnmpdConf returns the SnmpdConf field value if set, zero value otherwise.
func (o *ApplianceAllOfSnmpServer) GetSnmpdConf() string {
	if o == nil || o.SnmpdConf == nil {
		var ret string
		return ret
	}
	return *o.SnmpdConf
}

// GetSnmpdConfOk returns a tuple with the SnmpdConf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceAllOfSnmpServer) GetSnmpdConfOk() (*string, bool) {
	if o == nil || o.SnmpdConf == nil {
		return nil, false
	}
	return o.SnmpdConf, true
}

// HasSnmpdConf returns a boolean if a field has been set.
func (o *ApplianceAllOfSnmpServer) HasSnmpdConf() bool {
	if o != nil && o.SnmpdConf != nil {
		return true
	}

	return false
}

// SetSnmpdConf gets a reference to the given string and assigns it to the SnmpdConf field.
func (o *ApplianceAllOfSnmpServer) SetSnmpdConf(v string) {
	o.SnmpdConf = &v
}

// GetAllowSources returns the AllowSources field value if set, zero value otherwise.
func (o *ApplianceAllOfSnmpServer) GetAllowSources() []map[string]interface{} {
	if o == nil || o.AllowSources == nil {
		var ret []map[string]interface{}
		return ret
	}
	return *o.AllowSources
}

// GetAllowSourcesOk returns a tuple with the AllowSources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceAllOfSnmpServer) GetAllowSourcesOk() (*[]map[string]interface{}, bool) {
	if o == nil || o.AllowSources == nil {
		return nil, false
	}
	return o.AllowSources, true
}

// HasAllowSources returns a boolean if a field has been set.
func (o *ApplianceAllOfSnmpServer) HasAllowSources() bool {
	if o != nil && o.AllowSources != nil {
		return true
	}

	return false
}

// SetAllowSources gets a reference to the given []map[string]interface{} and assigns it to the AllowSources field.
func (o *ApplianceAllOfSnmpServer) SetAllowSources(v []map[string]interface{}) {
	o.AllowSources = &v
}

func (o ApplianceAllOfSnmpServer) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.TcpPort != nil {
		toSerialize["tcpPort"] = o.TcpPort
	}
	if o.UdpPort != nil {
		toSerialize["udpPort"] = o.UdpPort
	}
	if o.SnmpdConf != nil {
		toSerialize["snmpd.conf"] = o.SnmpdConf
	}
	if o.AllowSources != nil {
		toSerialize["allowSources"] = o.AllowSources
	}
	return json.Marshal(toSerialize)
}

type NullableApplianceAllOfSnmpServer struct {
	value *ApplianceAllOfSnmpServer
	isSet bool
}

func (v NullableApplianceAllOfSnmpServer) Get() *ApplianceAllOfSnmpServer {
	return v.value
}

func (v *NullableApplianceAllOfSnmpServer) Set(val *ApplianceAllOfSnmpServer) {
	v.value = val
	v.isSet = true
}

func (v NullableApplianceAllOfSnmpServer) IsSet() bool {
	return v.isSet
}

func (v *NullableApplianceAllOfSnmpServer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplianceAllOfSnmpServer(val *ApplianceAllOfSnmpServer) *NullableApplianceAllOfSnmpServer {
	return &NullableApplianceAllOfSnmpServer{value: val, isSet: true}
}

func (v NullableApplianceAllOfSnmpServer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplianceAllOfSnmpServer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
