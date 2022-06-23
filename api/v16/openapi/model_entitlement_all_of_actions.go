/*
Appgate SDP Controller REST API

# About   This specification documents the REST API calls for the Appgate SDP Controller.    Please refer to the REST API chapter in the manual or contact Appgate support with any questions about   this functionality. # Getting Started   Requirements for API scripting:   - Access to the Admin/API TLS Connection (default port 8443) of a Controller appliance.     (https://sdphelp.appgate.com/adminguide/appliance-function-configure.html?anchor=admin-api)   - An API user with relevant permissions.     (https://sdphelp.appgate.com/adminguide/administrative-roles-configure.html)   - In order to use the simple login API, Admin MFA must be disabled or the API user must be excluded.     (https://sdphelp.appgate.com/adminguide/mfa-for-admins.html) # Base path   HTTPS requests must be sent to the Admin Interface hostname and port, with **_/admin** path.    For example: **https://appgate.company.com:8443/admin**    All requests must have the **Accept** header as:    **application/vnd.appgate.peer-v16+json** # API Conventions   API conventions are  important to understand and follow strictly.    - While updating objects (via PUT), entire object must be sent with all fields.     - For example, in order to add a remedy method to the condition below:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": []       }       ```     - send the entire object with updated and non-updated fields:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": [{\"type\": \"DisplayMessage\", \"message\": \"test message\"}]       }       ```    - In case Controller returns an error (non-2xx HTTP status code), response body is JSON.     The \"message\" field contains information about the error.     HTTP 422 \"Unprocessable Entity\" has extra `errors` field to list all the issues with specific fields.    - Empty string (\"\") is considered a different value than \"null\" or field being omitted from JSON.     Omitting the field is recommend if no value is intended.     Empty string (\"\") will be almost always rejected as invalid value.    - There are common pattern between many objects:     - **Configuration Objects**: There are many objects with common fields, namely \"id\", \"name\", \"notes\", \"created\"       and \"updated\". These entities are listed, queried, created, updated and deleted in a similar fashion.     - **Distinguished Name**: Users and Devices are identified with what is called Distinguished Names, as used in        LDAP. The distinguished format that identifies a device and a user combination is        \"CN=\\<Device ID\\>,CN=\\<username\\>,OU=\\<Identity Provider Name\\>\". Some objects have the        \"userDistinguishedName\" field, which does not include the CN for Device ID.        This identifies a user on every device.

API version: API version 16.5
Contact: appgatesdp.support@appgate.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// EntitlementAllOfActions IP Access action.
type EntitlementAllOfActions struct {
	// Type of the IP Access action.
	Subtype string `json:"subtype"`
	// Applied action to the traffic.
	Action string `json:"action"`
	// Hosts to apply the action to. See admin manual for possible values.
	Hosts []string `json:"hosts"`
	// Destination port. Multiple ports can be entered comma separated. Port ranges can be entered dash separated. Only valid for tcp and udp subtypes
	Ports *[]string `json:"ports,omitempty"`
	// ICMP type. Only valid for icmp subtypes.
	Types   *[]string                `json:"types,omitempty"`
	Monitor *EntitlementAllOfMonitor `json:"monitor,omitempty"`
}

// NewEntitlementAllOfActions instantiates a new EntitlementAllOfActions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitlementAllOfActions(subtype string, action string, hosts []string) *EntitlementAllOfActions {
	this := EntitlementAllOfActions{}
	this.Subtype = subtype
	this.Action = action
	this.Hosts = hosts
	return &this
}

// NewEntitlementAllOfActionsWithDefaults instantiates a new EntitlementAllOfActions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitlementAllOfActionsWithDefaults() *EntitlementAllOfActions {
	this := EntitlementAllOfActions{}
	return &this
}

// GetSubtype returns the Subtype field value
func (o *EntitlementAllOfActions) GetSubtype() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Subtype
}

// GetSubtypeOk returns a tuple with the Subtype field value
// and a boolean to check if the value has been set.
func (o *EntitlementAllOfActions) GetSubtypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subtype, true
}

// SetSubtype sets field value
func (o *EntitlementAllOfActions) SetSubtype(v string) {
	o.Subtype = v
}

// GetAction returns the Action field value
func (o *EntitlementAllOfActions) GetAction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *EntitlementAllOfActions) GetActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *EntitlementAllOfActions) SetAction(v string) {
	o.Action = v
}

// GetHosts returns the Hosts field value
func (o *EntitlementAllOfActions) GetHosts() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Hosts
}

// GetHostsOk returns a tuple with the Hosts field value
// and a boolean to check if the value has been set.
func (o *EntitlementAllOfActions) GetHostsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hosts, true
}

// SetHosts sets field value
func (o *EntitlementAllOfActions) SetHosts(v []string) {
	o.Hosts = v
}

// GetPorts returns the Ports field value if set, zero value otherwise.
func (o *EntitlementAllOfActions) GetPorts() []string {
	if o == nil || o.Ports == nil {
		var ret []string
		return ret
	}
	return *o.Ports
}

// GetPortsOk returns a tuple with the Ports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementAllOfActions) GetPortsOk() (*[]string, bool) {
	if o == nil || o.Ports == nil {
		return nil, false
	}
	return o.Ports, true
}

// HasPorts returns a boolean if a field has been set.
func (o *EntitlementAllOfActions) HasPorts() bool {
	if o != nil && o.Ports != nil {
		return true
	}

	return false
}

// SetPorts gets a reference to the given []string and assigns it to the Ports field.
func (o *EntitlementAllOfActions) SetPorts(v []string) {
	o.Ports = &v
}

// GetTypes returns the Types field value if set, zero value otherwise.
func (o *EntitlementAllOfActions) GetTypes() []string {
	if o == nil || o.Types == nil {
		var ret []string
		return ret
	}
	return *o.Types
}

// GetTypesOk returns a tuple with the Types field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementAllOfActions) GetTypesOk() (*[]string, bool) {
	if o == nil || o.Types == nil {
		return nil, false
	}
	return o.Types, true
}

// HasTypes returns a boolean if a field has been set.
func (o *EntitlementAllOfActions) HasTypes() bool {
	if o != nil && o.Types != nil {
		return true
	}

	return false
}

// SetTypes gets a reference to the given []string and assigns it to the Types field.
func (o *EntitlementAllOfActions) SetTypes(v []string) {
	o.Types = &v
}

// GetMonitor returns the Monitor field value if set, zero value otherwise.
func (o *EntitlementAllOfActions) GetMonitor() EntitlementAllOfMonitor {
	if o == nil || o.Monitor == nil {
		var ret EntitlementAllOfMonitor
		return ret
	}
	return *o.Monitor
}

// GetMonitorOk returns a tuple with the Monitor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementAllOfActions) GetMonitorOk() (*EntitlementAllOfMonitor, bool) {
	if o == nil || o.Monitor == nil {
		return nil, false
	}
	return o.Monitor, true
}

// HasMonitor returns a boolean if a field has been set.
func (o *EntitlementAllOfActions) HasMonitor() bool {
	if o != nil && o.Monitor != nil {
		return true
	}

	return false
}

// SetMonitor gets a reference to the given EntitlementAllOfMonitor and assigns it to the Monitor field.
func (o *EntitlementAllOfActions) SetMonitor(v EntitlementAllOfMonitor) {
	o.Monitor = &v
}

func (o EntitlementAllOfActions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["subtype"] = o.Subtype
	}
	if true {
		toSerialize["action"] = o.Action
	}
	if true {
		toSerialize["hosts"] = o.Hosts
	}
	if o.Ports != nil {
		toSerialize["ports"] = o.Ports
	}
	if o.Types != nil {
		toSerialize["types"] = o.Types
	}
	if o.Monitor != nil {
		toSerialize["monitor"] = o.Monitor
	}
	return json.Marshal(toSerialize)
}

type NullableEntitlementAllOfActions struct {
	value *EntitlementAllOfActions
	isSet bool
}

func (v NullableEntitlementAllOfActions) Get() *EntitlementAllOfActions {
	return v.value
}

func (v *NullableEntitlementAllOfActions) Set(val *EntitlementAllOfActions) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitlementAllOfActions) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitlementAllOfActions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitlementAllOfActions(val *EntitlementAllOfActions) *NullableEntitlementAllOfActions {
	return &NullableEntitlementAllOfActions{value: val, isSet: true}
}

func (v NullableEntitlementAllOfActions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitlementAllOfActions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
