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

// ApplianceAllOfMetricsAggregator Metrics Aggregator settings. It collects metrics from the appliances in the given sites and provides Prometheus APIs for consumption.
type ApplianceAllOfMetricsAggregator struct {
	// Whether the Metrics Aggregator is enabled on this appliance or not.
	Enabled            *bool               `json:"enabled,omitempty"`
	PrometheusExporter *PrometheusExporter `json:"prometheusExporter,omitempty"`
	// The sites to collect metrics from.
	Sites []string `json:"sites,omitempty"`
}

// NewApplianceAllOfMetricsAggregator instantiates a new ApplianceAllOfMetricsAggregator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplianceAllOfMetricsAggregator() *ApplianceAllOfMetricsAggregator {
	this := ApplianceAllOfMetricsAggregator{}
	var enabled bool = false
	this.Enabled = &enabled
	return &this
}

// NewApplianceAllOfMetricsAggregatorWithDefaults instantiates a new ApplianceAllOfMetricsAggregator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplianceAllOfMetricsAggregatorWithDefaults() *ApplianceAllOfMetricsAggregator {
	this := ApplianceAllOfMetricsAggregator{}
	var enabled bool = false
	this.Enabled = &enabled
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *ApplianceAllOfMetricsAggregator) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceAllOfMetricsAggregator) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *ApplianceAllOfMetricsAggregator) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *ApplianceAllOfMetricsAggregator) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetPrometheusExporter returns the PrometheusExporter field value if set, zero value otherwise.
func (o *ApplianceAllOfMetricsAggregator) GetPrometheusExporter() PrometheusExporter {
	if o == nil || o.PrometheusExporter == nil {
		var ret PrometheusExporter
		return ret
	}
	return *o.PrometheusExporter
}

// GetPrometheusExporterOk returns a tuple with the PrometheusExporter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceAllOfMetricsAggregator) GetPrometheusExporterOk() (*PrometheusExporter, bool) {
	if o == nil || o.PrometheusExporter == nil {
		return nil, false
	}
	return o.PrometheusExporter, true
}

// HasPrometheusExporter returns a boolean if a field has been set.
func (o *ApplianceAllOfMetricsAggregator) HasPrometheusExporter() bool {
	if o != nil && o.PrometheusExporter != nil {
		return true
	}

	return false
}

// SetPrometheusExporter gets a reference to the given PrometheusExporter and assigns it to the PrometheusExporter field.
func (o *ApplianceAllOfMetricsAggregator) SetPrometheusExporter(v PrometheusExporter) {
	o.PrometheusExporter = &v
}

// GetSites returns the Sites field value if set, zero value otherwise.
func (o *ApplianceAllOfMetricsAggregator) GetSites() []string {
	if o == nil || o.Sites == nil {
		var ret []string
		return ret
	}
	return o.Sites
}

// GetSitesOk returns a tuple with the Sites field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceAllOfMetricsAggregator) GetSitesOk() ([]string, bool) {
	if o == nil || o.Sites == nil {
		return nil, false
	}
	return o.Sites, true
}

// HasSites returns a boolean if a field has been set.
func (o *ApplianceAllOfMetricsAggregator) HasSites() bool {
	if o != nil && o.Sites != nil {
		return true
	}

	return false
}

// SetSites gets a reference to the given []string and assigns it to the Sites field.
func (o *ApplianceAllOfMetricsAggregator) SetSites(v []string) {
	o.Sites = v
}

func (o ApplianceAllOfMetricsAggregator) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.PrometheusExporter != nil {
		toSerialize["prometheusExporter"] = o.PrometheusExporter
	}
	if o.Sites != nil {
		toSerialize["sites"] = o.Sites
	}
	return json.Marshal(toSerialize)
}

type NullableApplianceAllOfMetricsAggregator struct {
	value *ApplianceAllOfMetricsAggregator
	isSet bool
}

func (v NullableApplianceAllOfMetricsAggregator) Get() *ApplianceAllOfMetricsAggregator {
	return v.value
}

func (v *NullableApplianceAllOfMetricsAggregator) Set(val *ApplianceAllOfMetricsAggregator) {
	v.value = val
	v.isSet = true
}

func (v NullableApplianceAllOfMetricsAggregator) IsSet() bool {
	return v.isSet
}

func (v *NullableApplianceAllOfMetricsAggregator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplianceAllOfMetricsAggregator(val *ApplianceAllOfMetricsAggregator) *NullableApplianceAllOfMetricsAggregator {
	return &NullableApplianceAllOfMetricsAggregator{value: val, isSet: true}
}

func (v NullableApplianceAllOfMetricsAggregator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplianceAllOfMetricsAggregator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
