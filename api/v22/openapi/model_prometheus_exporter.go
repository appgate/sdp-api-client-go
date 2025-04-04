/*
Appgate SDP Controller REST API

# About   This specification documents the REST API calls for the Appgate SDP Controller.    Please refer to the REST API chapter in the manual or contact Appgate support with any questions about   this functionality. # Getting Started   Requirements for API scripting:   - Access to the Admin/API TLS Connection (default port 8443) of a Controller appliance.     (https://sdphelp.appgate.com/adminguide/appliance-function-configure.html?anchor=admin-api)   - An API user with relevant permissions.     (https://sdphelp.appgate.com/adminguide/administrative-roles-configure.html)   - In order to use the simple login API, Admin MFA must be disabled or the API user must be excluded.     (https://sdphelp.appgate.com/adminguide/mfa-for-admins.html) # Base path   HTTPS requests must be sent to the Admin Interface hostname and port, with **_/admin** path.    For example: **https://appgate.company.com:8443/admin**    All requests must have the **Accept** header as:    **application/vnd.appgate.peer-v22+json**    An exception is made for the **_/admin/version** endpoint which instead expects an **application/json** Accept header. # API Conventions   API conventions are  important to understand and follow strictly.    - While updating objects (via PUT), entire object must be sent with all fields.     - For example, in order to add a remedy method to the condition below:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": []       }       ```     - send the entire object with updated and non-updated fields:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": [{\"type\": \"DisplayMessage\", \"message\": \"test message\"}]       }       ```    - In case Controller returns an error (non-2xx HTTP status code), response body is JSON.     The \"message\" field contains information about the error.     HTTP 422 \"Unprocessable Entity\" has extra `errors` field to list all the issues with specific fields.    - Empty string (\"\") is considered a different value than \"null\" or field being omitted from JSON.     Omitting the field is recommended if no value is intended.     Empty string (\"\") will be almost always rejected as invalid value.    - There are common pattern between many objects:     - **Configuration Objects**: There are many objects with common fields, namely \"id\", \"name\", \"notes\", \"created\"       and \"updated\". These entities are listed, queried, created, updated and deleted in a similar fashion.     - **Distinguished Name**: Users and Devices are identified with what is called Distinguished Names, as used in        LDAP. The distinguished format that identifies a device and a user combination is        \"CN=\\<Device ID\\>,CN=\\<username\\>,OU=\\<Identity Provider Name\\>\". Some objects have the        \"userDistinguishedName\" field, which does not include the CN for Device ID.        This identifies a user on every device.

API version: API version 22.0
Contact: appgatesdp.support@appgate.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// PrometheusExporter Prometheus Exporter configuration.
type PrometheusExporter struct {
	// Whether the Prometheus Exporter is enabled on this appliance or not.
	Enabled *bool `json:"enabled,omitempty"`
	// Port to connect for Prometheus Exporter.
	Port *int32 `json:"port,omitempty"`
	// Source configuration to allow via iptables.
	AllowSources []AllowSourcesInner `json:"allowSources,omitempty"`
	// Whether to use HTTP or HTTPS for the exporter.
	UseHTTPS *bool `json:"useHTTPS,omitempty"`
	HttpsP12 *P12  `json:"httpsP12,omitempty"`
	// Enable basic auth for Prometheus Exporter.
	BasicAuth *bool `json:"basicAuth,omitempty"`
	// Basic auth users.
	AllowedUsers []PrometheusExporterAllowedUsersInner `json:"allowedUsers,omitempty"`
	// List of labels to filter out.
	LabelsDisabled []string `json:"labelsDisabled,omitempty"`
}

// NewPrometheusExporter instantiates a new PrometheusExporter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrometheusExporter() *PrometheusExporter {
	this := PrometheusExporter{}
	var enabled bool = false
	this.Enabled = &enabled
	var port int32 = 5556
	this.Port = &port
	var useHTTPS bool = false
	this.UseHTTPS = &useHTTPS
	var basicAuth bool = false
	this.BasicAuth = &basicAuth
	return &this
}

// NewPrometheusExporterWithDefaults instantiates a new PrometheusExporter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrometheusExporterWithDefaults() *PrometheusExporter {
	this := PrometheusExporter{}
	var enabled bool = false
	this.Enabled = &enabled
	var port int32 = 5556
	this.Port = &port
	var useHTTPS bool = false
	this.UseHTTPS = &useHTTPS
	var basicAuth bool = false
	this.BasicAuth = &basicAuth
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *PrometheusExporter) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrometheusExporter) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *PrometheusExporter) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *PrometheusExporter) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *PrometheusExporter) GetPort() int32 {
	if o == nil || o.Port == nil {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrometheusExporter) GetPortOk() (*int32, bool) {
	if o == nil || o.Port == nil {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *PrometheusExporter) HasPort() bool {
	if o != nil && o.Port != nil {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *PrometheusExporter) SetPort(v int32) {
	o.Port = &v
}

// GetAllowSources returns the AllowSources field value if set, zero value otherwise.
func (o *PrometheusExporter) GetAllowSources() []AllowSourcesInner {
	if o == nil || o.AllowSources == nil {
		var ret []AllowSourcesInner
		return ret
	}
	return o.AllowSources
}

// GetAllowSourcesOk returns a tuple with the AllowSources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrometheusExporter) GetAllowSourcesOk() ([]AllowSourcesInner, bool) {
	if o == nil || o.AllowSources == nil {
		return nil, false
	}
	return o.AllowSources, true
}

// HasAllowSources returns a boolean if a field has been set.
func (o *PrometheusExporter) HasAllowSources() bool {
	if o != nil && o.AllowSources != nil {
		return true
	}

	return false
}

// SetAllowSources gets a reference to the given []AllowSourcesInner and assigns it to the AllowSources field.
func (o *PrometheusExporter) SetAllowSources(v []AllowSourcesInner) {
	o.AllowSources = v
}

// GetUseHTTPS returns the UseHTTPS field value if set, zero value otherwise.
func (o *PrometheusExporter) GetUseHTTPS() bool {
	if o == nil || o.UseHTTPS == nil {
		var ret bool
		return ret
	}
	return *o.UseHTTPS
}

// GetUseHTTPSOk returns a tuple with the UseHTTPS field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrometheusExporter) GetUseHTTPSOk() (*bool, bool) {
	if o == nil || o.UseHTTPS == nil {
		return nil, false
	}
	return o.UseHTTPS, true
}

// HasUseHTTPS returns a boolean if a field has been set.
func (o *PrometheusExporter) HasUseHTTPS() bool {
	if o != nil && o.UseHTTPS != nil {
		return true
	}

	return false
}

// SetUseHTTPS gets a reference to the given bool and assigns it to the UseHTTPS field.
func (o *PrometheusExporter) SetUseHTTPS(v bool) {
	o.UseHTTPS = &v
}

// GetHttpsP12 returns the HttpsP12 field value if set, zero value otherwise.
func (o *PrometheusExporter) GetHttpsP12() P12 {
	if o == nil || o.HttpsP12 == nil {
		var ret P12
		return ret
	}
	return *o.HttpsP12
}

// GetHttpsP12Ok returns a tuple with the HttpsP12 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrometheusExporter) GetHttpsP12Ok() (*P12, bool) {
	if o == nil || o.HttpsP12 == nil {
		return nil, false
	}
	return o.HttpsP12, true
}

// HasHttpsP12 returns a boolean if a field has been set.
func (o *PrometheusExporter) HasHttpsP12() bool {
	if o != nil && o.HttpsP12 != nil {
		return true
	}

	return false
}

// SetHttpsP12 gets a reference to the given P12 and assigns it to the HttpsP12 field.
func (o *PrometheusExporter) SetHttpsP12(v P12) {
	o.HttpsP12 = &v
}

// GetBasicAuth returns the BasicAuth field value if set, zero value otherwise.
func (o *PrometheusExporter) GetBasicAuth() bool {
	if o == nil || o.BasicAuth == nil {
		var ret bool
		return ret
	}
	return *o.BasicAuth
}

// GetBasicAuthOk returns a tuple with the BasicAuth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrometheusExporter) GetBasicAuthOk() (*bool, bool) {
	if o == nil || o.BasicAuth == nil {
		return nil, false
	}
	return o.BasicAuth, true
}

// HasBasicAuth returns a boolean if a field has been set.
func (o *PrometheusExporter) HasBasicAuth() bool {
	if o != nil && o.BasicAuth != nil {
		return true
	}

	return false
}

// SetBasicAuth gets a reference to the given bool and assigns it to the BasicAuth field.
func (o *PrometheusExporter) SetBasicAuth(v bool) {
	o.BasicAuth = &v
}

// GetAllowedUsers returns the AllowedUsers field value if set, zero value otherwise.
func (o *PrometheusExporter) GetAllowedUsers() []PrometheusExporterAllowedUsersInner {
	if o == nil || o.AllowedUsers == nil {
		var ret []PrometheusExporterAllowedUsersInner
		return ret
	}
	return o.AllowedUsers
}

// GetAllowedUsersOk returns a tuple with the AllowedUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrometheusExporter) GetAllowedUsersOk() ([]PrometheusExporterAllowedUsersInner, bool) {
	if o == nil || o.AllowedUsers == nil {
		return nil, false
	}
	return o.AllowedUsers, true
}

// HasAllowedUsers returns a boolean if a field has been set.
func (o *PrometheusExporter) HasAllowedUsers() bool {
	if o != nil && o.AllowedUsers != nil {
		return true
	}

	return false
}

// SetAllowedUsers gets a reference to the given []PrometheusExporterAllowedUsersInner and assigns it to the AllowedUsers field.
func (o *PrometheusExporter) SetAllowedUsers(v []PrometheusExporterAllowedUsersInner) {
	o.AllowedUsers = v
}

// GetLabelsDisabled returns the LabelsDisabled field value if set, zero value otherwise.
func (o *PrometheusExporter) GetLabelsDisabled() []string {
	if o == nil || o.LabelsDisabled == nil {
		var ret []string
		return ret
	}
	return o.LabelsDisabled
}

// GetLabelsDisabledOk returns a tuple with the LabelsDisabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrometheusExporter) GetLabelsDisabledOk() ([]string, bool) {
	if o == nil || o.LabelsDisabled == nil {
		return nil, false
	}
	return o.LabelsDisabled, true
}

// HasLabelsDisabled returns a boolean if a field has been set.
func (o *PrometheusExporter) HasLabelsDisabled() bool {
	if o != nil && o.LabelsDisabled != nil {
		return true
	}

	return false
}

// SetLabelsDisabled gets a reference to the given []string and assigns it to the LabelsDisabled field.
func (o *PrometheusExporter) SetLabelsDisabled(v []string) {
	o.LabelsDisabled = v
}

func (o PrometheusExporter) MarshalJSON() ([]byte, error) {
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
	if o.UseHTTPS != nil {
		toSerialize["useHTTPS"] = o.UseHTTPS
	}
	if o.HttpsP12 != nil {
		toSerialize["httpsP12"] = o.HttpsP12
	}
	if o.BasicAuth != nil {
		toSerialize["basicAuth"] = o.BasicAuth
	}
	if o.AllowedUsers != nil {
		toSerialize["allowedUsers"] = o.AllowedUsers
	}
	if o.LabelsDisabled != nil {
		toSerialize["labelsDisabled"] = o.LabelsDisabled
	}
	return json.Marshal(toSerialize)
}

type NullablePrometheusExporter struct {
	value *PrometheusExporter
	isSet bool
}

func (v NullablePrometheusExporter) Get() *PrometheusExporter {
	return v.value
}

func (v *NullablePrometheusExporter) Set(val *PrometheusExporter) {
	v.value = val
	v.isSet = true
}

func (v NullablePrometheusExporter) IsSet() bool {
	return v.isSet
}

func (v *NullablePrometheusExporter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrometheusExporter(val *PrometheusExporter) *NullablePrometheusExporter {
	return &NullablePrometheusExporter{value: val, isSet: true}
}

func (v NullablePrometheusExporter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrometheusExporter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
