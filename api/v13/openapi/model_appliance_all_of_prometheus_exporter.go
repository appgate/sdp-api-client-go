/*
 * AppGate SDP Controller REST API
 *
 * # About   This specification documents the REST API calls for the AppGate SDP Controller.    Please refer to the Integration chapter in the manual or contact AppGate support with any questions about   this functionality. # Getting Started   Requirements for API scripting:   - Access to the peer interface (default port 444) of a Controller appliance.     (https://sdphelp.appgate.com/adminguide/appliances-configure.html?anchor=peer)   - An API user with relevant permissions.     (https://sdphelp.appgate.com/adminguide/administrative-roles-configure.html)   - In order to use the simple login API, Admin MFA must be disabled or the API user must be excluded.     (https://sdphelp.appgate.com/adminguide/mfa-for-admins.html) # Base path   HTTPS requests must be sent to the Peer Interface hostname and port, with **_/admin** path.    For example: **https://appgate.company.com:444/admin**    All requests must have the **Accept** header as:    **application/vnd.appgate.peer-v13+json** # API Conventions   API conventions are  important to understand and follow strictly.    - While updating objects (via PUT), entire object must be sent with all fields.     - For example, if in order to add a remedy method to the condition below:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": []       }       ```     - send the entire object with updated and non-updated fields:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": [{\"type\": \"DisplayMessage\", \"message\": \"test message\"}]       }       ```    - In case Controller returns an error (non-2xx HTTP status code), response body is JSON.     The \"message\" field contains information about the error.     HTTP 422 \"Unprocessable Entity\" has extra `errors` field to list all the issues with specific fields.    - Empty string (\"\") is considered a different value than \"null\" or field being omitted from JSON.     Omitting the field is recommend if no value is intended.     Empty string (\"\") will be almost always rejected as invalid value.    - There are common pattern between many objects:     - **Configuration Objects**: There are many objects with common fields, namely \"id\", \"name\", \"notes\", \"created\"       and \"updated\". These entities are listed, queried, created, updated and deleted in a similar fashion.     - **Distinguished Name**: Users and Devices are identified with what is called Distinguished Names, as used in        LDAP. The distinguished format that identifies a device and a user combination is        \"CN=\\<Device ID\\>,CN=\\<username\\>,OU=\\<Identity Provider Name\\>\". Some objects have the        \"userDistinguishedName\" field, which does not include the CN for Device ID.        This identifies a user on every device.
 *
 * API version: API version 13
 * Contact: appgatesdp.support@appgate.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// ApplianceAllOfPrometheusExporter Prometheus Exporter configuration.
type ApplianceAllOfPrometheusExporter struct {
	// Whether the Prometheus Exporter is enabled on this appliance or not.
	Enabled *bool `json:"enabled,omitempty"`
	// Port to connect for Prometheus Exporter.
	Port *int32 `json:"port,omitempty"`
	// Source configuration to allow via iptables.
	AllowSources *[]map[string]interface{} `json:"allowSources,omitempty"`
}

// NewApplianceAllOfPrometheusExporter instantiates a new ApplianceAllOfPrometheusExporter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplianceAllOfPrometheusExporter() *ApplianceAllOfPrometheusExporter {
	this := ApplianceAllOfPrometheusExporter{}
	var enabled bool = false
	this.Enabled = &enabled
	var port int32 = 5556
	this.Port = &port
	return &this
}

// NewApplianceAllOfPrometheusExporterWithDefaults instantiates a new ApplianceAllOfPrometheusExporter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplianceAllOfPrometheusExporterWithDefaults() *ApplianceAllOfPrometheusExporter {
	this := ApplianceAllOfPrometheusExporter{}
	var enabled bool = false
	this.Enabled = &enabled
	var port int32 = 5556
	this.Port = &port
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *ApplianceAllOfPrometheusExporter) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceAllOfPrometheusExporter) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *ApplianceAllOfPrometheusExporter) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *ApplianceAllOfPrometheusExporter) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *ApplianceAllOfPrometheusExporter) GetPort() int32 {
	if o == nil || o.Port == nil {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceAllOfPrometheusExporter) GetPortOk() (*int32, bool) {
	if o == nil || o.Port == nil {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *ApplianceAllOfPrometheusExporter) HasPort() bool {
	if o != nil && o.Port != nil {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *ApplianceAllOfPrometheusExporter) SetPort(v int32) {
	o.Port = &v
}

// GetAllowSources returns the AllowSources field value if set, zero value otherwise.
func (o *ApplianceAllOfPrometheusExporter) GetAllowSources() []map[string]interface{} {
	if o == nil || o.AllowSources == nil {
		var ret []map[string]interface{}
		return ret
	}
	return *o.AllowSources
}

// GetAllowSourcesOk returns a tuple with the AllowSources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceAllOfPrometheusExporter) GetAllowSourcesOk() (*[]map[string]interface{}, bool) {
	if o == nil || o.AllowSources == nil {
		return nil, false
	}
	return o.AllowSources, true
}

// HasAllowSources returns a boolean if a field has been set.
func (o *ApplianceAllOfPrometheusExporter) HasAllowSources() bool {
	if o != nil && o.AllowSources != nil {
		return true
	}

	return false
}

// SetAllowSources gets a reference to the given []map[string]interface{} and assigns it to the AllowSources field.
func (o *ApplianceAllOfPrometheusExporter) SetAllowSources(v []map[string]interface{}) {
	o.AllowSources = &v
}

func (o ApplianceAllOfPrometheusExporter) MarshalJSON() ([]byte, error) {
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

type NullableApplianceAllOfPrometheusExporter struct {
	value *ApplianceAllOfPrometheusExporter
	isSet bool
}

func (v NullableApplianceAllOfPrometheusExporter) Get() *ApplianceAllOfPrometheusExporter {
	return v.value
}

func (v *NullableApplianceAllOfPrometheusExporter) Set(val *ApplianceAllOfPrometheusExporter) {
	v.value = val
	v.isSet = true
}

func (v NullableApplianceAllOfPrometheusExporter) IsSet() bool {
	return v.isSet
}

func (v *NullableApplianceAllOfPrometheusExporter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplianceAllOfPrometheusExporter(val *ApplianceAllOfPrometheusExporter) *NullableApplianceAllOfPrometheusExporter {
	return &NullableApplianceAllOfPrometheusExporter{value: val, isSet: true}
}

func (v NullableApplianceAllOfPrometheusExporter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplianceAllOfPrometheusExporter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
