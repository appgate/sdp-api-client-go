/*
Appgate SDP Controller REST API

# About   This specification documents the REST API calls for the Appgate SDP Controller.    Please refer to the REST API chapter in the manual or contact Appgate support with any questions about   this functionality. # Getting Started   Requirements for API scripting:   - Access to the Admin/API TLS Connection (default port 8443) of a Controller appliance.     (https://sdphelp.appgate.com/adminguide/appliance-function-configure.html?anchor=admin-api)   - An API user with relevant permissions.     (https://sdphelp.appgate.com/adminguide/administrative-roles-configure.html)   - In order to use the simple login API, Admin MFA must be disabled or the API user must be excluded.     (https://sdphelp.appgate.com/adminguide/mfa-for-admins.html) # Base path   HTTPS requests must be sent to the Admin Interface hostname and port, with **_/admin** path.    For example: **https://appgate.company.com:8443/admin**    All requests must have the **Accept** header as:    **application/vnd.appgate.peer-v21+json**    An exception is made for the **_/admin/version** endpoint which instead expects an **application/json** Accept header. # API Conventions   API conventions are  important to understand and follow strictly.    - While updating objects (via PUT), entire object must be sent with all fields.     - For example, in order to add a remedy method to the condition below:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": []       }       ```     - send the entire object with updated and non-updated fields:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": [{\"type\": \"DisplayMessage\", \"message\": \"test message\"}]       }       ```    - In case Controller returns an error (non-2xx HTTP status code), response body is JSON.     The \"message\" field contains information about the error.     HTTP 422 \"Unprocessable Entity\" has extra `errors` field to list all the issues with specific fields.    - Empty string (\"\") is considered a different value than \"null\" or field being omitted from JSON.     Omitting the field is recommended if no value is intended.     Empty string (\"\") will be almost always rejected as invalid value.    - There are common pattern between many objects:     - **Configuration Objects**: There are many objects with common fields, namely \"id\", \"name\", \"notes\", \"created\"       and \"updated\". These entities are listed, queried, created, updated and deleted in a similar fashion.     - **Distinguished Name**: Users and Devices are identified with what is called Distinguished Names, as used in        LDAP. The distinguished format that identifies a device and a user combination is        \"CN=\\<Device ID\\>,CN=\\<username\\>,OU=\\<Identity Provider Name\\>\". Some objects have the        \"userDistinguishedName\" field, which does not include the CN for Device ID.        This identifies a user on every device.

API version: API version 21.0
Contact: appgatesdp.support@appgate.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// ApplianceAllOfLogForwarder LogForwarder settings. LogForwarder collects audit logs from the appliances in the given sites and sends them to the given endpoints.
type ApplianceAllOfLogForwarder struct {
	// Whether the LogForwarder is enabled on this appliance or not.
	Enabled       *bool                 `json:"enabled,omitempty"`
	Elasticsearch NullableElasticsearch `json:"elasticsearch,omitempty"`
	// TCP endpoints to connect and send the audit logs with the given format.
	TcpClients []TcpClient `json:"tcpClients,omitempty"`
	// AWS Kinesis endpoints to connect and send the audit logs with the given format.
	AwsKineses []AwsKinesis `json:"awsKineses,omitempty"`
	// SumoLogic endpoints to connect and send the audit logs to.
	SumoLogicClients []SumoLogic `json:"sumoLogicClients,omitempty"`
	// Splunk endpoints to connect and send the audit logs to.
	SplunkClients []Splunk `json:"splunkClients,omitempty"`
	// Azure Monitor endpoints to connect and send the audit logs to.
	AzureMonitors []AzureMonitor `json:"azureMonitors,omitempty"`
	// Falcon LogScale endpoints to connect and send the audit logs to.
	FalconLogScales []FalconLogScale `json:"falconLogScales,omitempty"`
	// Datadog endpoints to connect and send the audit logs to.
	Datadogs []Datadog `json:"datadogs,omitempty"`
	// Coralogix endpoints to connect and send the audit logs to.
	Coralogixs []Coralogix `json:"coralogixs,omitempty"`
	// The sites to collect logs from and forward.
	Sites []string `json:"sites,omitempty"`
}

// NewApplianceAllOfLogForwarder instantiates a new ApplianceAllOfLogForwarder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplianceAllOfLogForwarder() *ApplianceAllOfLogForwarder {
	this := ApplianceAllOfLogForwarder{}
	var enabled bool = false
	this.Enabled = &enabled
	return &this
}

// NewApplianceAllOfLogForwarderWithDefaults instantiates a new ApplianceAllOfLogForwarder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplianceAllOfLogForwarderWithDefaults() *ApplianceAllOfLogForwarder {
	this := ApplianceAllOfLogForwarder{}
	var enabled bool = false
	this.Enabled = &enabled
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *ApplianceAllOfLogForwarder) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceAllOfLogForwarder) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *ApplianceAllOfLogForwarder) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *ApplianceAllOfLogForwarder) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetElasticsearch returns the Elasticsearch field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplianceAllOfLogForwarder) GetElasticsearch() Elasticsearch {
	if o == nil || o.Elasticsearch.Get() == nil {
		var ret Elasticsearch
		return ret
	}
	return *o.Elasticsearch.Get()
}

// GetElasticsearchOk returns a tuple with the Elasticsearch field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplianceAllOfLogForwarder) GetElasticsearchOk() (*Elasticsearch, bool) {
	if o == nil {
		return nil, false
	}
	return o.Elasticsearch.Get(), o.Elasticsearch.IsSet()
}

// HasElasticsearch returns a boolean if a field has been set.
func (o *ApplianceAllOfLogForwarder) HasElasticsearch() bool {
	if o != nil && o.Elasticsearch.IsSet() {
		return true
	}

	return false
}

// SetElasticsearch gets a reference to the given NullableElasticsearch and assigns it to the Elasticsearch field.
func (o *ApplianceAllOfLogForwarder) SetElasticsearch(v Elasticsearch) {
	o.Elasticsearch.Set(&v)
}

// SetElasticsearchNil sets the value for Elasticsearch to be an explicit nil
func (o *ApplianceAllOfLogForwarder) SetElasticsearchNil() {
	o.Elasticsearch.Set(nil)
}

// UnsetElasticsearch ensures that no value is present for Elasticsearch, not even an explicit nil
func (o *ApplianceAllOfLogForwarder) UnsetElasticsearch() {
	o.Elasticsearch.Unset()
}

// GetTcpClients returns the TcpClients field value if set, zero value otherwise.
func (o *ApplianceAllOfLogForwarder) GetTcpClients() []TcpClient {
	if o == nil || o.TcpClients == nil {
		var ret []TcpClient
		return ret
	}
	return o.TcpClients
}

// GetTcpClientsOk returns a tuple with the TcpClients field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceAllOfLogForwarder) GetTcpClientsOk() ([]TcpClient, bool) {
	if o == nil || o.TcpClients == nil {
		return nil, false
	}
	return o.TcpClients, true
}

// HasTcpClients returns a boolean if a field has been set.
func (o *ApplianceAllOfLogForwarder) HasTcpClients() bool {
	if o != nil && o.TcpClients != nil {
		return true
	}

	return false
}

// SetTcpClients gets a reference to the given []TcpClient and assigns it to the TcpClients field.
func (o *ApplianceAllOfLogForwarder) SetTcpClients(v []TcpClient) {
	o.TcpClients = v
}

// GetAwsKineses returns the AwsKineses field value if set, zero value otherwise.
func (o *ApplianceAllOfLogForwarder) GetAwsKineses() []AwsKinesis {
	if o == nil || o.AwsKineses == nil {
		var ret []AwsKinesis
		return ret
	}
	return o.AwsKineses
}

// GetAwsKinesesOk returns a tuple with the AwsKineses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceAllOfLogForwarder) GetAwsKinesesOk() ([]AwsKinesis, bool) {
	if o == nil || o.AwsKineses == nil {
		return nil, false
	}
	return o.AwsKineses, true
}

// HasAwsKineses returns a boolean if a field has been set.
func (o *ApplianceAllOfLogForwarder) HasAwsKineses() bool {
	if o != nil && o.AwsKineses != nil {
		return true
	}

	return false
}

// SetAwsKineses gets a reference to the given []AwsKinesis and assigns it to the AwsKineses field.
func (o *ApplianceAllOfLogForwarder) SetAwsKineses(v []AwsKinesis) {
	o.AwsKineses = v
}

// GetSumoLogicClients returns the SumoLogicClients field value if set, zero value otherwise.
func (o *ApplianceAllOfLogForwarder) GetSumoLogicClients() []SumoLogic {
	if o == nil || o.SumoLogicClients == nil {
		var ret []SumoLogic
		return ret
	}
	return o.SumoLogicClients
}

// GetSumoLogicClientsOk returns a tuple with the SumoLogicClients field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceAllOfLogForwarder) GetSumoLogicClientsOk() ([]SumoLogic, bool) {
	if o == nil || o.SumoLogicClients == nil {
		return nil, false
	}
	return o.SumoLogicClients, true
}

// HasSumoLogicClients returns a boolean if a field has been set.
func (o *ApplianceAllOfLogForwarder) HasSumoLogicClients() bool {
	if o != nil && o.SumoLogicClients != nil {
		return true
	}

	return false
}

// SetSumoLogicClients gets a reference to the given []SumoLogic and assigns it to the SumoLogicClients field.
func (o *ApplianceAllOfLogForwarder) SetSumoLogicClients(v []SumoLogic) {
	o.SumoLogicClients = v
}

// GetSplunkClients returns the SplunkClients field value if set, zero value otherwise.
func (o *ApplianceAllOfLogForwarder) GetSplunkClients() []Splunk {
	if o == nil || o.SplunkClients == nil {
		var ret []Splunk
		return ret
	}
	return o.SplunkClients
}

// GetSplunkClientsOk returns a tuple with the SplunkClients field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceAllOfLogForwarder) GetSplunkClientsOk() ([]Splunk, bool) {
	if o == nil || o.SplunkClients == nil {
		return nil, false
	}
	return o.SplunkClients, true
}

// HasSplunkClients returns a boolean if a field has been set.
func (o *ApplianceAllOfLogForwarder) HasSplunkClients() bool {
	if o != nil && o.SplunkClients != nil {
		return true
	}

	return false
}

// SetSplunkClients gets a reference to the given []Splunk and assigns it to the SplunkClients field.
func (o *ApplianceAllOfLogForwarder) SetSplunkClients(v []Splunk) {
	o.SplunkClients = v
}

// GetAzureMonitors returns the AzureMonitors field value if set, zero value otherwise.
func (o *ApplianceAllOfLogForwarder) GetAzureMonitors() []AzureMonitor {
	if o == nil || o.AzureMonitors == nil {
		var ret []AzureMonitor
		return ret
	}
	return o.AzureMonitors
}

// GetAzureMonitorsOk returns a tuple with the AzureMonitors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceAllOfLogForwarder) GetAzureMonitorsOk() ([]AzureMonitor, bool) {
	if o == nil || o.AzureMonitors == nil {
		return nil, false
	}
	return o.AzureMonitors, true
}

// HasAzureMonitors returns a boolean if a field has been set.
func (o *ApplianceAllOfLogForwarder) HasAzureMonitors() bool {
	if o != nil && o.AzureMonitors != nil {
		return true
	}

	return false
}

// SetAzureMonitors gets a reference to the given []AzureMonitor and assigns it to the AzureMonitors field.
func (o *ApplianceAllOfLogForwarder) SetAzureMonitors(v []AzureMonitor) {
	o.AzureMonitors = v
}

// GetFalconLogScales returns the FalconLogScales field value if set, zero value otherwise.
func (o *ApplianceAllOfLogForwarder) GetFalconLogScales() []FalconLogScale {
	if o == nil || o.FalconLogScales == nil {
		var ret []FalconLogScale
		return ret
	}
	return o.FalconLogScales
}

// GetFalconLogScalesOk returns a tuple with the FalconLogScales field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceAllOfLogForwarder) GetFalconLogScalesOk() ([]FalconLogScale, bool) {
	if o == nil || o.FalconLogScales == nil {
		return nil, false
	}
	return o.FalconLogScales, true
}

// HasFalconLogScales returns a boolean if a field has been set.
func (o *ApplianceAllOfLogForwarder) HasFalconLogScales() bool {
	if o != nil && o.FalconLogScales != nil {
		return true
	}

	return false
}

// SetFalconLogScales gets a reference to the given []FalconLogScale and assigns it to the FalconLogScales field.
func (o *ApplianceAllOfLogForwarder) SetFalconLogScales(v []FalconLogScale) {
	o.FalconLogScales = v
}

// GetDatadogs returns the Datadogs field value if set, zero value otherwise.
func (o *ApplianceAllOfLogForwarder) GetDatadogs() []Datadog {
	if o == nil || o.Datadogs == nil {
		var ret []Datadog
		return ret
	}
	return o.Datadogs
}

// GetDatadogsOk returns a tuple with the Datadogs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceAllOfLogForwarder) GetDatadogsOk() ([]Datadog, bool) {
	if o == nil || o.Datadogs == nil {
		return nil, false
	}
	return o.Datadogs, true
}

// HasDatadogs returns a boolean if a field has been set.
func (o *ApplianceAllOfLogForwarder) HasDatadogs() bool {
	if o != nil && o.Datadogs != nil {
		return true
	}

	return false
}

// SetDatadogs gets a reference to the given []Datadog and assigns it to the Datadogs field.
func (o *ApplianceAllOfLogForwarder) SetDatadogs(v []Datadog) {
	o.Datadogs = v
}

// GetCoralogixs returns the Coralogixs field value if set, zero value otherwise.
func (o *ApplianceAllOfLogForwarder) GetCoralogixs() []Coralogix {
	if o == nil || o.Coralogixs == nil {
		var ret []Coralogix
		return ret
	}
	return o.Coralogixs
}

// GetCoralogixsOk returns a tuple with the Coralogixs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceAllOfLogForwarder) GetCoralogixsOk() ([]Coralogix, bool) {
	if o == nil || o.Coralogixs == nil {
		return nil, false
	}
	return o.Coralogixs, true
}

// HasCoralogixs returns a boolean if a field has been set.
func (o *ApplianceAllOfLogForwarder) HasCoralogixs() bool {
	if o != nil && o.Coralogixs != nil {
		return true
	}

	return false
}

// SetCoralogixs gets a reference to the given []Coralogix and assigns it to the Coralogixs field.
func (o *ApplianceAllOfLogForwarder) SetCoralogixs(v []Coralogix) {
	o.Coralogixs = v
}

// GetSites returns the Sites field value if set, zero value otherwise.
func (o *ApplianceAllOfLogForwarder) GetSites() []string {
	if o == nil || o.Sites == nil {
		var ret []string
		return ret
	}
	return o.Sites
}

// GetSitesOk returns a tuple with the Sites field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceAllOfLogForwarder) GetSitesOk() ([]string, bool) {
	if o == nil || o.Sites == nil {
		return nil, false
	}
	return o.Sites, true
}

// HasSites returns a boolean if a field has been set.
func (o *ApplianceAllOfLogForwarder) HasSites() bool {
	if o != nil && o.Sites != nil {
		return true
	}

	return false
}

// SetSites gets a reference to the given []string and assigns it to the Sites field.
func (o *ApplianceAllOfLogForwarder) SetSites(v []string) {
	o.Sites = v
}

func (o ApplianceAllOfLogForwarder) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.Elasticsearch.IsSet() {
		toSerialize["elasticsearch"] = o.Elasticsearch.Get()
	}
	if o.TcpClients != nil {
		toSerialize["tcpClients"] = o.TcpClients
	}
	if o.AwsKineses != nil {
		toSerialize["awsKineses"] = o.AwsKineses
	}
	if o.SumoLogicClients != nil {
		toSerialize["sumoLogicClients"] = o.SumoLogicClients
	}
	if o.SplunkClients != nil {
		toSerialize["splunkClients"] = o.SplunkClients
	}
	if o.AzureMonitors != nil {
		toSerialize["azureMonitors"] = o.AzureMonitors
	}
	if o.FalconLogScales != nil {
		toSerialize["falconLogScales"] = o.FalconLogScales
	}
	if o.Datadogs != nil {
		toSerialize["datadogs"] = o.Datadogs
	}
	if o.Coralogixs != nil {
		toSerialize["coralogixs"] = o.Coralogixs
	}
	if o.Sites != nil {
		toSerialize["sites"] = o.Sites
	}
	return json.Marshal(toSerialize)
}

type NullableApplianceAllOfLogForwarder struct {
	value *ApplianceAllOfLogForwarder
	isSet bool
}

func (v NullableApplianceAllOfLogForwarder) Get() *ApplianceAllOfLogForwarder {
	return v.value
}

func (v *NullableApplianceAllOfLogForwarder) Set(val *ApplianceAllOfLogForwarder) {
	v.value = val
	v.isSet = true
}

func (v NullableApplianceAllOfLogForwarder) IsSet() bool {
	return v.isSet
}

func (v *NullableApplianceAllOfLogForwarder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplianceAllOfLogForwarder(val *ApplianceAllOfLogForwarder) *NullableApplianceAllOfLogForwarder {
	return &NullableApplianceAllOfLogForwarder{value: val, isSet: true}
}

func (v NullableApplianceAllOfLogForwarder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplianceAllOfLogForwarder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}