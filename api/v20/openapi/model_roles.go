/*
Appgate SDP Controller REST API

# About   This specification documents the REST API calls for the Appgate SDP Controller.    Please refer to the REST API chapter in the manual or contact Appgate support with any questions about   this functionality. # Getting Started   Requirements for API scripting:   - Access to the Admin/API TLS Connection (default port 8443) of a Controller appliance.     (https://sdphelp.appgate.com/adminguide/appliance-function-configure.html?anchor=admin-api)   - An API user with relevant permissions.     (https://sdphelp.appgate.com/adminguide/administrative-roles-configure.html)   - In order to use the simple login API, Admin MFA must be disabled or the API user must be excluded.     (https://sdphelp.appgate.com/adminguide/mfa-for-admins.html) # Base path   HTTPS requests must be sent to the Admin Interface hostname and port, with **_/admin** path.    For example: **https://appgate.company.com:8443/admin**    All requests must have the **Accept** header as:    **application/vnd.appgate.peer-v20+json**    An exception is made for the **_/admin/version** endpoint which instead expects an **application/json** Accept header. # API Conventions   API conventions are  important to understand and follow strictly.    - While updating objects (via PUT), entire object must be sent with all fields.     - For example, in order to add a remedy method to the condition below:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": []       }       ```     - send the entire object with updated and non-updated fields:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": [{\"type\": \"DisplayMessage\", \"message\": \"test message\"}]       }       ```    - In case Controller returns an error (non-2xx HTTP status code), response body is JSON.     The \"message\" field contains information about the error.     HTTP 422 \"Unprocessable Entity\" has extra `errors` field to list all the issues with specific fields.    - Empty string (\"\") is considered a different value than \"null\" or field being omitted from JSON.     Omitting the field is recommended if no value is intended.     Empty string (\"\") will be almost always rejected as invalid value.    - There are common pattern between many objects:     - **Configuration Objects**: There are many objects with common fields, namely \"id\", \"name\", \"notes\", \"created\"       and \"updated\". These entities are listed, queried, created, updated and deleted in a similar fashion.     - **Distinguished Name**: Users and Devices are identified with what is called Distinguished Names, as used in        LDAP. The distinguished format that identifies a device and a user combination is        \"CN=\\<Device ID\\>,CN=\\<username\\>,OU=\\<Identity Provider Name\\>\". Some objects have the        \"userDistinguishedName\" field, which does not include the CN for Device ID.        This identifies a user on every device.

API version: API version 20.0
Contact: appgatesdp.support@appgate.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// Roles Appliance roles.
type Roles struct {
	Controller        *ControllerRole            `json:"controller,omitempty"`
	LogServer         *ApplianceRole             `json:"logServer,omitempty"`
	LogForwarder      *ApplianceRole             `json:"logForwarder,omitempty"`
	MetricsAggregator *ApplianceRole             `json:"metricsAggregator,omitempty"`
	Gateway           *ApplianceWithSessionsRole `json:"gateway,omitempty"`
	Connector         *ApplianceWithSessionsRole `json:"connector,omitempty"`
	Portal            *ApplianceWithSessionsRole `json:"portal,omitempty"`
	Appliance         *ApplianceRole             `json:"appliance,omitempty"`
}

// NewRoles instantiates a new Roles object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoles() *Roles {
	this := Roles{}
	return &this
}

// NewRolesWithDefaults instantiates a new Roles object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRolesWithDefaults() *Roles {
	this := Roles{}
	return &this
}

// GetController returns the Controller field value if set, zero value otherwise.
func (o *Roles) GetController() ControllerRole {
	if o == nil || o.Controller == nil {
		var ret ControllerRole
		return ret
	}
	return *o.Controller
}

// GetControllerOk returns a tuple with the Controller field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Roles) GetControllerOk() (*ControllerRole, bool) {
	if o == nil || o.Controller == nil {
		return nil, false
	}
	return o.Controller, true
}

// HasController returns a boolean if a field has been set.
func (o *Roles) HasController() bool {
	if o != nil && o.Controller != nil {
		return true
	}

	return false
}

// SetController gets a reference to the given ControllerRole and assigns it to the Controller field.
func (o *Roles) SetController(v ControllerRole) {
	o.Controller = &v
}

// GetLogServer returns the LogServer field value if set, zero value otherwise.
func (o *Roles) GetLogServer() ApplianceRole {
	if o == nil || o.LogServer == nil {
		var ret ApplianceRole
		return ret
	}
	return *o.LogServer
}

// GetLogServerOk returns a tuple with the LogServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Roles) GetLogServerOk() (*ApplianceRole, bool) {
	if o == nil || o.LogServer == nil {
		return nil, false
	}
	return o.LogServer, true
}

// HasLogServer returns a boolean if a field has been set.
func (o *Roles) HasLogServer() bool {
	if o != nil && o.LogServer != nil {
		return true
	}

	return false
}

// SetLogServer gets a reference to the given ApplianceRole and assigns it to the LogServer field.
func (o *Roles) SetLogServer(v ApplianceRole) {
	o.LogServer = &v
}

// GetLogForwarder returns the LogForwarder field value if set, zero value otherwise.
func (o *Roles) GetLogForwarder() ApplianceRole {
	if o == nil || o.LogForwarder == nil {
		var ret ApplianceRole
		return ret
	}
	return *o.LogForwarder
}

// GetLogForwarderOk returns a tuple with the LogForwarder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Roles) GetLogForwarderOk() (*ApplianceRole, bool) {
	if o == nil || o.LogForwarder == nil {
		return nil, false
	}
	return o.LogForwarder, true
}

// HasLogForwarder returns a boolean if a field has been set.
func (o *Roles) HasLogForwarder() bool {
	if o != nil && o.LogForwarder != nil {
		return true
	}

	return false
}

// SetLogForwarder gets a reference to the given ApplianceRole and assigns it to the LogForwarder field.
func (o *Roles) SetLogForwarder(v ApplianceRole) {
	o.LogForwarder = &v
}

// GetMetricsAggregator returns the MetricsAggregator field value if set, zero value otherwise.
func (o *Roles) GetMetricsAggregator() ApplianceRole {
	if o == nil || o.MetricsAggregator == nil {
		var ret ApplianceRole
		return ret
	}
	return *o.MetricsAggregator
}

// GetMetricsAggregatorOk returns a tuple with the MetricsAggregator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Roles) GetMetricsAggregatorOk() (*ApplianceRole, bool) {
	if o == nil || o.MetricsAggregator == nil {
		return nil, false
	}
	return o.MetricsAggregator, true
}

// HasMetricsAggregator returns a boolean if a field has been set.
func (o *Roles) HasMetricsAggregator() bool {
	if o != nil && o.MetricsAggregator != nil {
		return true
	}

	return false
}

// SetMetricsAggregator gets a reference to the given ApplianceRole and assigns it to the MetricsAggregator field.
func (o *Roles) SetMetricsAggregator(v ApplianceRole) {
	o.MetricsAggregator = &v
}

// GetGateway returns the Gateway field value if set, zero value otherwise.
func (o *Roles) GetGateway() ApplianceWithSessionsRole {
	if o == nil || o.Gateway == nil {
		var ret ApplianceWithSessionsRole
		return ret
	}
	return *o.Gateway
}

// GetGatewayOk returns a tuple with the Gateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Roles) GetGatewayOk() (*ApplianceWithSessionsRole, bool) {
	if o == nil || o.Gateway == nil {
		return nil, false
	}
	return o.Gateway, true
}

// HasGateway returns a boolean if a field has been set.
func (o *Roles) HasGateway() bool {
	if o != nil && o.Gateway != nil {
		return true
	}

	return false
}

// SetGateway gets a reference to the given ApplianceWithSessionsRole and assigns it to the Gateway field.
func (o *Roles) SetGateway(v ApplianceWithSessionsRole) {
	o.Gateway = &v
}

// GetConnector returns the Connector field value if set, zero value otherwise.
func (o *Roles) GetConnector() ApplianceWithSessionsRole {
	if o == nil || o.Connector == nil {
		var ret ApplianceWithSessionsRole
		return ret
	}
	return *o.Connector
}

// GetConnectorOk returns a tuple with the Connector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Roles) GetConnectorOk() (*ApplianceWithSessionsRole, bool) {
	if o == nil || o.Connector == nil {
		return nil, false
	}
	return o.Connector, true
}

// HasConnector returns a boolean if a field has been set.
func (o *Roles) HasConnector() bool {
	if o != nil && o.Connector != nil {
		return true
	}

	return false
}

// SetConnector gets a reference to the given ApplianceWithSessionsRole and assigns it to the Connector field.
func (o *Roles) SetConnector(v ApplianceWithSessionsRole) {
	o.Connector = &v
}

// GetPortal returns the Portal field value if set, zero value otherwise.
func (o *Roles) GetPortal() ApplianceWithSessionsRole {
	if o == nil || o.Portal == nil {
		var ret ApplianceWithSessionsRole
		return ret
	}
	return *o.Portal
}

// GetPortalOk returns a tuple with the Portal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Roles) GetPortalOk() (*ApplianceWithSessionsRole, bool) {
	if o == nil || o.Portal == nil {
		return nil, false
	}
	return o.Portal, true
}

// HasPortal returns a boolean if a field has been set.
func (o *Roles) HasPortal() bool {
	if o != nil && o.Portal != nil {
		return true
	}

	return false
}

// SetPortal gets a reference to the given ApplianceWithSessionsRole and assigns it to the Portal field.
func (o *Roles) SetPortal(v ApplianceWithSessionsRole) {
	o.Portal = &v
}

// GetAppliance returns the Appliance field value if set, zero value otherwise.
func (o *Roles) GetAppliance() ApplianceRole {
	if o == nil || o.Appliance == nil {
		var ret ApplianceRole
		return ret
	}
	return *o.Appliance
}

// GetApplianceOk returns a tuple with the Appliance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Roles) GetApplianceOk() (*ApplianceRole, bool) {
	if o == nil || o.Appliance == nil {
		return nil, false
	}
	return o.Appliance, true
}

// HasAppliance returns a boolean if a field has been set.
func (o *Roles) HasAppliance() bool {
	if o != nil && o.Appliance != nil {
		return true
	}

	return false
}

// SetAppliance gets a reference to the given ApplianceRole and assigns it to the Appliance field.
func (o *Roles) SetAppliance(v ApplianceRole) {
	o.Appliance = &v
}

func (o Roles) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Controller != nil {
		toSerialize["controller"] = o.Controller
	}
	if o.LogServer != nil {
		toSerialize["logServer"] = o.LogServer
	}
	if o.LogForwarder != nil {
		toSerialize["logForwarder"] = o.LogForwarder
	}
	if o.MetricsAggregator != nil {
		toSerialize["metricsAggregator"] = o.MetricsAggregator
	}
	if o.Gateway != nil {
		toSerialize["gateway"] = o.Gateway
	}
	if o.Connector != nil {
		toSerialize["connector"] = o.Connector
	}
	if o.Portal != nil {
		toSerialize["portal"] = o.Portal
	}
	if o.Appliance != nil {
		toSerialize["appliance"] = o.Appliance
	}
	return json.Marshal(toSerialize)
}

type NullableRoles struct {
	value *Roles
	isSet bool
}

func (v NullableRoles) Get() *Roles {
	return v.value
}

func (v *NullableRoles) Set(val *Roles) {
	v.value = val
	v.isSet = true
}

func (v NullableRoles) IsSet() bool {
	return v.isSet
}

func (v *NullableRoles) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoles(val *Roles) *NullableRoles {
	return &NullableRoles{value: val, isSet: true}
}

func (v NullableRoles) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoles) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
