/*
Appgate SDP Controller REST API

# About   This specification documents the REST API calls for the Appgate SDP Controller.    Please refer to the REST API chapter in the manual or contact Appgate support with any questions about   this functionality. # Getting Started   Requirements for API scripting:   - Access to the Admin/API TLS Connection (default port 8443) of a Controller appliance.     (https://sdphelp.appgate.com/adminguide/appliance-function-configure.html?anchor=admin-api)   - An API user with relevant permissions.     (https://sdphelp.appgate.com/adminguide/administrative-roles-configure.html)   - In order to use the simple login API, Admin MFA must be disabled or the API user must be excluded.     (https://sdphelp.appgate.com/adminguide/mfa-for-admins.html) # Base path   HTTPS requests must be sent to the Admin Interface hostname and port, with **_/admin** path.    For example: **https://appgate.company.com:8443/admin**    All requests must have the **Accept** header as:    **application/vnd.appgate.peer-v20+json**    An exception is made for the **_/admin/version** endpoint which instead expects an **application/json** Accept header. # API Conventions   API conventions are  important to understand and follow strictly.    - While updating objects (via PUT), entire object must be sent with all fields.     - For example, in order to add a remedy method to the condition below:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": []       }       ```     - send the entire object with updated and non-updated fields:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": [{\"type\": \"DisplayMessage\", \"message\": \"test message\"}]       }       ```    - In case Controller returns an error (non-2xx HTTP status code), response body is JSON.     The \"message\" field contains information about the error.     HTTP 422 \"Unprocessable Entity\" has extra `errors` field to list all the issues with specific fields.    - Empty string (\"\") is considered a different value than \"null\" or field being omitted from JSON.     Omitting the field is recommended if no value is intended.     Empty string (\"\") will be almost always rejected as invalid value.    - There are common pattern between many objects:     - **Configuration Objects**: There are many objects with common fields, namely \"id\", \"name\", \"notes\", \"created\"       and \"updated\". These entities are listed, queried, created, updated and deleted in a similar fashion.     - **Distinguished Name**: Users and Devices are identified with what is called Distinguished Names, as used in        LDAP. The distinguished format that identifies a device and a user combination is        \"CN=\\<Device ID\\>,CN=\\<username\\>,OU=\\<Identity Provider Name\\>\". Some objects have the        \"userDistinguishedName\" field, which does not include the CN for Device ID.        This identifies a user on every device.

API version: API version 20.3
Contact: appgatesdp.support@appgate.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// StatsAppliancesListAllOf struct for StatsAppliancesListAllOf
type StatsAppliancesListAllOf struct {
	// The queries applied to the list.
	Queries []string `json:"queries,omitempty"`
	// The filters applied to the list.
	FilterBy []FilterBy `json:"filterBy,omitempty"`
	// The number of active Appliances with the Controller role enabled.
	ControllerCount *float32 `json:"controllerCount,omitempty"`
	// The number of active Appliances with the Gateway role enabled.
	GatewayCount *float32 `json:"gatewayCount,omitempty"`
	// The number of active Appliances in total.
	ApplianceCount *float32 `json:"applianceCount,omitempty"`
	// The number of active Appliances with the LogServer role enabled.
	LogServerCount *float32 `json:"logServerCount,omitempty"`
	// The number of active Appliances with the LogForwarder role enabled.
	LogForwarderCount *float32 `json:"logForwarderCount,omitempty"`
	// The number of active Appliances with the Connector role enabled.
	ConnectorCount *float32 `json:"connectorCount,omitempty"`
	// The number of active Appliances with the Portal role enabled.
	PortalCount *float32                       `json:"portalCount,omitempty"`
	Data        []StatsAppliancesListAllOfData `json:"data,omitempty"`
}

// NewStatsAppliancesListAllOf instantiates a new StatsAppliancesListAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStatsAppliancesListAllOf() *StatsAppliancesListAllOf {
	this := StatsAppliancesListAllOf{}
	return &this
}

// NewStatsAppliancesListAllOfWithDefaults instantiates a new StatsAppliancesListAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStatsAppliancesListAllOfWithDefaults() *StatsAppliancesListAllOf {
	this := StatsAppliancesListAllOf{}
	return &this
}

// GetQueries returns the Queries field value if set, zero value otherwise.
func (o *StatsAppliancesListAllOf) GetQueries() []string {
	if o == nil || o.Queries == nil {
		var ret []string
		return ret
	}
	return o.Queries
}

// GetQueriesOk returns a tuple with the Queries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatsAppliancesListAllOf) GetQueriesOk() ([]string, bool) {
	if o == nil || o.Queries == nil {
		return nil, false
	}
	return o.Queries, true
}

// HasQueries returns a boolean if a field has been set.
func (o *StatsAppliancesListAllOf) HasQueries() bool {
	if o != nil && o.Queries != nil {
		return true
	}

	return false
}

// SetQueries gets a reference to the given []string and assigns it to the Queries field.
func (o *StatsAppliancesListAllOf) SetQueries(v []string) {
	o.Queries = v
}

// GetFilterBy returns the FilterBy field value if set, zero value otherwise.
func (o *StatsAppliancesListAllOf) GetFilterBy() []FilterBy {
	if o == nil || o.FilterBy == nil {
		var ret []FilterBy
		return ret
	}
	return o.FilterBy
}

// GetFilterByOk returns a tuple with the FilterBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatsAppliancesListAllOf) GetFilterByOk() ([]FilterBy, bool) {
	if o == nil || o.FilterBy == nil {
		return nil, false
	}
	return o.FilterBy, true
}

// HasFilterBy returns a boolean if a field has been set.
func (o *StatsAppliancesListAllOf) HasFilterBy() bool {
	if o != nil && o.FilterBy != nil {
		return true
	}

	return false
}

// SetFilterBy gets a reference to the given []FilterBy and assigns it to the FilterBy field.
func (o *StatsAppliancesListAllOf) SetFilterBy(v []FilterBy) {
	o.FilterBy = v
}

// GetControllerCount returns the ControllerCount field value if set, zero value otherwise.
func (o *StatsAppliancesListAllOf) GetControllerCount() float32 {
	if o == nil || o.ControllerCount == nil {
		var ret float32
		return ret
	}
	return *o.ControllerCount
}

// GetControllerCountOk returns a tuple with the ControllerCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatsAppliancesListAllOf) GetControllerCountOk() (*float32, bool) {
	if o == nil || o.ControllerCount == nil {
		return nil, false
	}
	return o.ControllerCount, true
}

// HasControllerCount returns a boolean if a field has been set.
func (o *StatsAppliancesListAllOf) HasControllerCount() bool {
	if o != nil && o.ControllerCount != nil {
		return true
	}

	return false
}

// SetControllerCount gets a reference to the given float32 and assigns it to the ControllerCount field.
func (o *StatsAppliancesListAllOf) SetControllerCount(v float32) {
	o.ControllerCount = &v
}

// GetGatewayCount returns the GatewayCount field value if set, zero value otherwise.
func (o *StatsAppliancesListAllOf) GetGatewayCount() float32 {
	if o == nil || o.GatewayCount == nil {
		var ret float32
		return ret
	}
	return *o.GatewayCount
}

// GetGatewayCountOk returns a tuple with the GatewayCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatsAppliancesListAllOf) GetGatewayCountOk() (*float32, bool) {
	if o == nil || o.GatewayCount == nil {
		return nil, false
	}
	return o.GatewayCount, true
}

// HasGatewayCount returns a boolean if a field has been set.
func (o *StatsAppliancesListAllOf) HasGatewayCount() bool {
	if o != nil && o.GatewayCount != nil {
		return true
	}

	return false
}

// SetGatewayCount gets a reference to the given float32 and assigns it to the GatewayCount field.
func (o *StatsAppliancesListAllOf) SetGatewayCount(v float32) {
	o.GatewayCount = &v
}

// GetApplianceCount returns the ApplianceCount field value if set, zero value otherwise.
func (o *StatsAppliancesListAllOf) GetApplianceCount() float32 {
	if o == nil || o.ApplianceCount == nil {
		var ret float32
		return ret
	}
	return *o.ApplianceCount
}

// GetApplianceCountOk returns a tuple with the ApplianceCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatsAppliancesListAllOf) GetApplianceCountOk() (*float32, bool) {
	if o == nil || o.ApplianceCount == nil {
		return nil, false
	}
	return o.ApplianceCount, true
}

// HasApplianceCount returns a boolean if a field has been set.
func (o *StatsAppliancesListAllOf) HasApplianceCount() bool {
	if o != nil && o.ApplianceCount != nil {
		return true
	}

	return false
}

// SetApplianceCount gets a reference to the given float32 and assigns it to the ApplianceCount field.
func (o *StatsAppliancesListAllOf) SetApplianceCount(v float32) {
	o.ApplianceCount = &v
}

// GetLogServerCount returns the LogServerCount field value if set, zero value otherwise.
func (o *StatsAppliancesListAllOf) GetLogServerCount() float32 {
	if o == nil || o.LogServerCount == nil {
		var ret float32
		return ret
	}
	return *o.LogServerCount
}

// GetLogServerCountOk returns a tuple with the LogServerCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatsAppliancesListAllOf) GetLogServerCountOk() (*float32, bool) {
	if o == nil || o.LogServerCount == nil {
		return nil, false
	}
	return o.LogServerCount, true
}

// HasLogServerCount returns a boolean if a field has been set.
func (o *StatsAppliancesListAllOf) HasLogServerCount() bool {
	if o != nil && o.LogServerCount != nil {
		return true
	}

	return false
}

// SetLogServerCount gets a reference to the given float32 and assigns it to the LogServerCount field.
func (o *StatsAppliancesListAllOf) SetLogServerCount(v float32) {
	o.LogServerCount = &v
}

// GetLogForwarderCount returns the LogForwarderCount field value if set, zero value otherwise.
func (o *StatsAppliancesListAllOf) GetLogForwarderCount() float32 {
	if o == nil || o.LogForwarderCount == nil {
		var ret float32
		return ret
	}
	return *o.LogForwarderCount
}

// GetLogForwarderCountOk returns a tuple with the LogForwarderCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatsAppliancesListAllOf) GetLogForwarderCountOk() (*float32, bool) {
	if o == nil || o.LogForwarderCount == nil {
		return nil, false
	}
	return o.LogForwarderCount, true
}

// HasLogForwarderCount returns a boolean if a field has been set.
func (o *StatsAppliancesListAllOf) HasLogForwarderCount() bool {
	if o != nil && o.LogForwarderCount != nil {
		return true
	}

	return false
}

// SetLogForwarderCount gets a reference to the given float32 and assigns it to the LogForwarderCount field.
func (o *StatsAppliancesListAllOf) SetLogForwarderCount(v float32) {
	o.LogForwarderCount = &v
}

// GetConnectorCount returns the ConnectorCount field value if set, zero value otherwise.
func (o *StatsAppliancesListAllOf) GetConnectorCount() float32 {
	if o == nil || o.ConnectorCount == nil {
		var ret float32
		return ret
	}
	return *o.ConnectorCount
}

// GetConnectorCountOk returns a tuple with the ConnectorCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatsAppliancesListAllOf) GetConnectorCountOk() (*float32, bool) {
	if o == nil || o.ConnectorCount == nil {
		return nil, false
	}
	return o.ConnectorCount, true
}

// HasConnectorCount returns a boolean if a field has been set.
func (o *StatsAppliancesListAllOf) HasConnectorCount() bool {
	if o != nil && o.ConnectorCount != nil {
		return true
	}

	return false
}

// SetConnectorCount gets a reference to the given float32 and assigns it to the ConnectorCount field.
func (o *StatsAppliancesListAllOf) SetConnectorCount(v float32) {
	o.ConnectorCount = &v
}

// GetPortalCount returns the PortalCount field value if set, zero value otherwise.
func (o *StatsAppliancesListAllOf) GetPortalCount() float32 {
	if o == nil || o.PortalCount == nil {
		var ret float32
		return ret
	}
	return *o.PortalCount
}

// GetPortalCountOk returns a tuple with the PortalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatsAppliancesListAllOf) GetPortalCountOk() (*float32, bool) {
	if o == nil || o.PortalCount == nil {
		return nil, false
	}
	return o.PortalCount, true
}

// HasPortalCount returns a boolean if a field has been set.
func (o *StatsAppliancesListAllOf) HasPortalCount() bool {
	if o != nil && o.PortalCount != nil {
		return true
	}

	return false
}

// SetPortalCount gets a reference to the given float32 and assigns it to the PortalCount field.
func (o *StatsAppliancesListAllOf) SetPortalCount(v float32) {
	o.PortalCount = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *StatsAppliancesListAllOf) GetData() []StatsAppliancesListAllOfData {
	if o == nil || o.Data == nil {
		var ret []StatsAppliancesListAllOfData
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatsAppliancesListAllOf) GetDataOk() ([]StatsAppliancesListAllOfData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *StatsAppliancesListAllOf) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []StatsAppliancesListAllOfData and assigns it to the Data field.
func (o *StatsAppliancesListAllOf) SetData(v []StatsAppliancesListAllOfData) {
	o.Data = v
}

func (o StatsAppliancesListAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Queries != nil {
		toSerialize["queries"] = o.Queries
	}
	if o.FilterBy != nil {
		toSerialize["filterBy"] = o.FilterBy
	}
	if o.ControllerCount != nil {
		toSerialize["controllerCount"] = o.ControllerCount
	}
	if o.GatewayCount != nil {
		toSerialize["gatewayCount"] = o.GatewayCount
	}
	if o.ApplianceCount != nil {
		toSerialize["applianceCount"] = o.ApplianceCount
	}
	if o.LogServerCount != nil {
		toSerialize["logServerCount"] = o.LogServerCount
	}
	if o.LogForwarderCount != nil {
		toSerialize["logForwarderCount"] = o.LogForwarderCount
	}
	if o.ConnectorCount != nil {
		toSerialize["connectorCount"] = o.ConnectorCount
	}
	if o.PortalCount != nil {
		toSerialize["portalCount"] = o.PortalCount
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableStatsAppliancesListAllOf struct {
	value *StatsAppliancesListAllOf
	isSet bool
}

func (v NullableStatsAppliancesListAllOf) Get() *StatsAppliancesListAllOf {
	return v.value
}

func (v *NullableStatsAppliancesListAllOf) Set(val *StatsAppliancesListAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableStatsAppliancesListAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableStatsAppliancesListAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatsAppliancesListAllOf(val *StatsAppliancesListAllOf) *NullableStatsAppliancesListAllOf {
	return &NullableStatsAppliancesListAllOf{value: val, isSet: true}
}

func (v NullableStatsAppliancesListAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatsAppliancesListAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
