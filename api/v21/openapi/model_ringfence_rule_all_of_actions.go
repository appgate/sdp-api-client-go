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

// RingfenceRuleAllOfActions struct for RingfenceRuleAllOfActions
type RingfenceRuleAllOfActions struct {
	// Protocol of the ringfence action.
	Protocol string `json:"protocol"`
	// The direction of the action
	Direction string `json:"direction"`
	// Applied action to the traffic.
	Action string `json:"action"`
	// Destination address. IP address or hostname.
	Hosts []string `json:"hosts"`
	// Destination port. Multiple ports can be entered comma separated. Port ranges can be entered dash separated. Only valid for tcp and udp subtypes.
	Ports []string `json:"ports,omitempty"`
	// ICMP type. Only valid for icmp protocol.
	Types []string `json:"types,omitempty"`
}

// NewRingfenceRuleAllOfActions instantiates a new RingfenceRuleAllOfActions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRingfenceRuleAllOfActions(protocol string, direction string, action string, hosts []string) *RingfenceRuleAllOfActions {
	this := RingfenceRuleAllOfActions{}
	this.Protocol = protocol
	this.Direction = direction
	this.Action = action
	this.Hosts = hosts
	return &this
}

// NewRingfenceRuleAllOfActionsWithDefaults instantiates a new RingfenceRuleAllOfActions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRingfenceRuleAllOfActionsWithDefaults() *RingfenceRuleAllOfActions {
	this := RingfenceRuleAllOfActions{}
	return &this
}

// GetProtocol returns the Protocol field value
func (o *RingfenceRuleAllOfActions) GetProtocol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value
// and a boolean to check if the value has been set.
func (o *RingfenceRuleAllOfActions) GetProtocolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Protocol, true
}

// SetProtocol sets field value
func (o *RingfenceRuleAllOfActions) SetProtocol(v string) {
	o.Protocol = v
}

// GetDirection returns the Direction field value
func (o *RingfenceRuleAllOfActions) GetDirection() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Direction
}

// GetDirectionOk returns a tuple with the Direction field value
// and a boolean to check if the value has been set.
func (o *RingfenceRuleAllOfActions) GetDirectionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Direction, true
}

// SetDirection sets field value
func (o *RingfenceRuleAllOfActions) SetDirection(v string) {
	o.Direction = v
}

// GetAction returns the Action field value
func (o *RingfenceRuleAllOfActions) GetAction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *RingfenceRuleAllOfActions) GetActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *RingfenceRuleAllOfActions) SetAction(v string) {
	o.Action = v
}

// GetHosts returns the Hosts field value
func (o *RingfenceRuleAllOfActions) GetHosts() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Hosts
}

// GetHostsOk returns a tuple with the Hosts field value
// and a boolean to check if the value has been set.
func (o *RingfenceRuleAllOfActions) GetHostsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Hosts, true
}

// SetHosts sets field value
func (o *RingfenceRuleAllOfActions) SetHosts(v []string) {
	o.Hosts = v
}

// GetPorts returns the Ports field value if set, zero value otherwise.
func (o *RingfenceRuleAllOfActions) GetPorts() []string {
	if o == nil || o.Ports == nil {
		var ret []string
		return ret
	}
	return o.Ports
}

// GetPortsOk returns a tuple with the Ports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RingfenceRuleAllOfActions) GetPortsOk() ([]string, bool) {
	if o == nil || o.Ports == nil {
		return nil, false
	}
	return o.Ports, true
}

// HasPorts returns a boolean if a field has been set.
func (o *RingfenceRuleAllOfActions) HasPorts() bool {
	if o != nil && o.Ports != nil {
		return true
	}

	return false
}

// SetPorts gets a reference to the given []string and assigns it to the Ports field.
func (o *RingfenceRuleAllOfActions) SetPorts(v []string) {
	o.Ports = v
}

// GetTypes returns the Types field value if set, zero value otherwise.
func (o *RingfenceRuleAllOfActions) GetTypes() []string {
	if o == nil || o.Types == nil {
		var ret []string
		return ret
	}
	return o.Types
}

// GetTypesOk returns a tuple with the Types field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RingfenceRuleAllOfActions) GetTypesOk() ([]string, bool) {
	if o == nil || o.Types == nil {
		return nil, false
	}
	return o.Types, true
}

// HasTypes returns a boolean if a field has been set.
func (o *RingfenceRuleAllOfActions) HasTypes() bool {
	if o != nil && o.Types != nil {
		return true
	}

	return false
}

// SetTypes gets a reference to the given []string and assigns it to the Types field.
func (o *RingfenceRuleAllOfActions) SetTypes(v []string) {
	o.Types = v
}

func (o RingfenceRuleAllOfActions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["protocol"] = o.Protocol
	}
	if true {
		toSerialize["direction"] = o.Direction
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
	return json.Marshal(toSerialize)
}

type NullableRingfenceRuleAllOfActions struct {
	value *RingfenceRuleAllOfActions
	isSet bool
}

func (v NullableRingfenceRuleAllOfActions) Get() *RingfenceRuleAllOfActions {
	return v.value
}

func (v *NullableRingfenceRuleAllOfActions) Set(val *RingfenceRuleAllOfActions) {
	v.value = val
	v.isSet = true
}

func (v NullableRingfenceRuleAllOfActions) IsSet() bool {
	return v.isSet
}

func (v *NullableRingfenceRuleAllOfActions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRingfenceRuleAllOfActions(val *RingfenceRuleAllOfActions) *NullableRingfenceRuleAllOfActions {
	return &NullableRingfenceRuleAllOfActions{value: val, isSet: true}
}

func (v NullableRingfenceRuleAllOfActions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRingfenceRuleAllOfActions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
