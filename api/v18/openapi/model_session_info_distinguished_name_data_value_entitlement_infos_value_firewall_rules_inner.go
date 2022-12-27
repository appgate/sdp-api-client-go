/*
Appgate SDP Controller REST API

# About   This specification documents the REST API calls for the Appgate SDP Controller.    Please refer to the REST API chapter in the manual or contact Appgate support with any questions about   this functionality. # Getting Started   Requirements for API scripting:   - Access to the Admin/API TLS Connection (default port 8443) of a Controller appliance.     (https://sdphelp.appgate.com/adminguide/appliance-function-configure.html?anchor=admin-api)   - An API user with relevant permissions.     (https://sdphelp.appgate.com/adminguide/administrative-roles-configure.html)   - In order to use the simple login API, Admin MFA must be disabled or the API user must be excluded.     (https://sdphelp.appgate.com/adminguide/mfa-for-admins.html) # Base path   HTTPS requests must be sent to the Admin Interface hostname and port, with **_/admin** path.    For example: **https://appgate.company.com:8443/admin**    All requests must have the **Accept** header as:    **application/vnd.appgate.peer-v18+json** # API Conventions   API conventions are  important to understand and follow strictly.    - While updating objects (via PUT), entire object must be sent with all fields.     - For example, in order to add a remedy method to the condition below:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": []       }       ```     - send the entire object with updated and non-updated fields:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": [{\"type\": \"DisplayMessage\", \"message\": \"test message\"}]       }       ```    - In case Controller returns an error (non-2xx HTTP status code), response body is JSON.     The \"message\" field contains information about the error.     HTTP 422 \"Unprocessable Entity\" has extra `errors` field to list all the issues with specific fields.    - Empty string (\"\") is considered a different value than \"null\" or field being omitted from JSON.     Omitting the field is recommended if no value is intended.     Empty string (\"\") will be almost always rejected as invalid value.    - There are common pattern between many objects:     - **Configuration Objects**: There are many objects with common fields, namely \"id\", \"name\", \"notes\", \"created\"       and \"updated\". These entities are listed, queried, created, updated and deleted in a similar fashion.     - **Distinguished Name**: Users and Devices are identified with what is called Distinguished Names, as used in        LDAP. The distinguished format that identifies a device and a user combination is        \"CN=\\<Device ID\\>,CN=\\<username\\>,OU=\\<Identity Provider Name\\>\". Some objects have the        \"userDistinguishedName\" field, which does not include the CN for Device ID.        This identifies a user on every device.

API version: API version 18.2
Contact: appgatesdp.support@appgate.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// SessionInfoDistinguishedNameDataValueEntitlementInfosValueFirewallRulesInner Firewall Rule.
type SessionInfoDistinguishedNameDataValueEntitlementInfosValueFirewallRulesInner struct {
	// Name of the rule.
	Name *string `json:"name,omitempty"`
	// The protocol for the Firewall Rule.
	Protocol *string `json:"protocol,omitempty"`
	// The direction of the Firewall Rrule.
	Direction *string `json:"direction,omitempty"`
	// The action for the Firewall Rule.
	Action *string `json:"action,omitempty"`
	// The subnets the Firewall Rule applies to.
	Subnets []string `json:"subnets,omitempty"`
	// The URLs the Firewall Rule applies to in case of http_up subtype.
	Urls []string `json:"urls,omitempty"`
	// The ports the Firewall Rule applies to.
	Ports []string `json:"ports,omitempty"`
	// The ICMP types the Firewall Rule applies to. Valid for ICMP.
	Types []string `json:"types,omitempty"`
}

// NewSessionInfoDistinguishedNameDataValueEntitlementInfosValueFirewallRulesInner instantiates a new SessionInfoDistinguishedNameDataValueEntitlementInfosValueFirewallRulesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSessionInfoDistinguishedNameDataValueEntitlementInfosValueFirewallRulesInner() *SessionInfoDistinguishedNameDataValueEntitlementInfosValueFirewallRulesInner {
	this := SessionInfoDistinguishedNameDataValueEntitlementInfosValueFirewallRulesInner{}
	return &this
}

// NewSessionInfoDistinguishedNameDataValueEntitlementInfosValueFirewallRulesInnerWithDefaults instantiates a new SessionInfoDistinguishedNameDataValueEntitlementInfosValueFirewallRulesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSessionInfoDistinguishedNameDataValueEntitlementInfosValueFirewallRulesInnerWithDefaults() *SessionInfoDistinguishedNameDataValueEntitlementInfosValueFirewallRulesInner {
	this := SessionInfoDistinguishedNameDataValueEntitlementInfosValueFirewallRulesInner{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SessionInfoDistinguishedNameDataValueEntitlementInfosValueFirewallRulesInner) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionInfoDistinguishedNameDataValueEntitlementInfosValueFirewallRulesInner) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SessionInfoDistinguishedNameDataValueEntitlementInfosValueFirewallRulesInner) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SessionInfoDistinguishedNameDataValueEntitlementInfosValueFirewallRulesInner) SetName(v string) {
	o.Name = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *SessionInfoDistinguishedNameDataValueEntitlementInfosValueFirewallRulesInner) GetProtocol() string {
	if o == nil || o.Protocol == nil {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionInfoDistinguishedNameDataValueEntitlementInfosValueFirewallRulesInner) GetProtocolOk() (*string, bool) {
	if o == nil || o.Protocol == nil {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *SessionInfoDistinguishedNameDataValueEntitlementInfosValueFirewallRulesInner) HasProtocol() bool {
	if o != nil && o.Protocol != nil {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *SessionInfoDistinguishedNameDataValueEntitlementInfosValueFirewallRulesInner) SetProtocol(v string) {
	o.Protocol = &v
}

// GetDirection returns the Direction field value if set, zero value otherwise.
func (o *SessionInfoDistinguishedNameDataValueEntitlementInfosValueFirewallRulesInner) GetDirection() string {
	if o == nil || o.Direction == nil {
		var ret string
		return ret
	}
	return *o.Direction
}

// GetDirectionOk returns a tuple with the Direction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionInfoDistinguishedNameDataValueEntitlementInfosValueFirewallRulesInner) GetDirectionOk() (*string, bool) {
	if o == nil || o.Direction == nil {
		return nil, false
	}
	return o.Direction, true
}

// HasDirection returns a boolean if a field has been set.
func (o *SessionInfoDistinguishedNameDataValueEntitlementInfosValueFirewallRulesInner) HasDirection() bool {
	if o != nil && o.Direction != nil {
		return true
	}

	return false
}

// SetDirection gets a reference to the given string and assigns it to the Direction field.
func (o *SessionInfoDistinguishedNameDataValueEntitlementInfosValueFirewallRulesInner) SetDirection(v string) {
	o.Direction = &v
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *SessionInfoDistinguishedNameDataValueEntitlementInfosValueFirewallRulesInner) GetAction() string {
	if o == nil || o.Action == nil {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionInfoDistinguishedNameDataValueEntitlementInfosValueFirewallRulesInner) GetActionOk() (*string, bool) {
	if o == nil || o.Action == nil {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *SessionInfoDistinguishedNameDataValueEntitlementInfosValueFirewallRulesInner) HasAction() bool {
	if o != nil && o.Action != nil {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *SessionInfoDistinguishedNameDataValueEntitlementInfosValueFirewallRulesInner) SetAction(v string) {
	o.Action = &v
}

// GetSubnets returns the Subnets field value if set, zero value otherwise.
func (o *SessionInfoDistinguishedNameDataValueEntitlementInfosValueFirewallRulesInner) GetSubnets() []string {
	if o == nil || o.Subnets == nil {
		var ret []string
		return ret
	}
	return o.Subnets
}

// GetSubnetsOk returns a tuple with the Subnets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionInfoDistinguishedNameDataValueEntitlementInfosValueFirewallRulesInner) GetSubnetsOk() ([]string, bool) {
	if o == nil || o.Subnets == nil {
		return nil, false
	}
	return o.Subnets, true
}

// HasSubnets returns a boolean if a field has been set.
func (o *SessionInfoDistinguishedNameDataValueEntitlementInfosValueFirewallRulesInner) HasSubnets() bool {
	if o != nil && o.Subnets != nil {
		return true
	}

	return false
}

// SetSubnets gets a reference to the given []string and assigns it to the Subnets field.
func (o *SessionInfoDistinguishedNameDataValueEntitlementInfosValueFirewallRulesInner) SetSubnets(v []string) {
	o.Subnets = v
}

// GetUrls returns the Urls field value if set, zero value otherwise.
func (o *SessionInfoDistinguishedNameDataValueEntitlementInfosValueFirewallRulesInner) GetUrls() []string {
	if o == nil || o.Urls == nil {
		var ret []string
		return ret
	}
	return o.Urls
}

// GetUrlsOk returns a tuple with the Urls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionInfoDistinguishedNameDataValueEntitlementInfosValueFirewallRulesInner) GetUrlsOk() ([]string, bool) {
	if o == nil || o.Urls == nil {
		return nil, false
	}
	return o.Urls, true
}

// HasUrls returns a boolean if a field has been set.
func (o *SessionInfoDistinguishedNameDataValueEntitlementInfosValueFirewallRulesInner) HasUrls() bool {
	if o != nil && o.Urls != nil {
		return true
	}

	return false
}

// SetUrls gets a reference to the given []string and assigns it to the Urls field.
func (o *SessionInfoDistinguishedNameDataValueEntitlementInfosValueFirewallRulesInner) SetUrls(v []string) {
	o.Urls = v
}

// GetPorts returns the Ports field value if set, zero value otherwise.
func (o *SessionInfoDistinguishedNameDataValueEntitlementInfosValueFirewallRulesInner) GetPorts() []string {
	if o == nil || o.Ports == nil {
		var ret []string
		return ret
	}
	return o.Ports
}

// GetPortsOk returns a tuple with the Ports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionInfoDistinguishedNameDataValueEntitlementInfosValueFirewallRulesInner) GetPortsOk() ([]string, bool) {
	if o == nil || o.Ports == nil {
		return nil, false
	}
	return o.Ports, true
}

// HasPorts returns a boolean if a field has been set.
func (o *SessionInfoDistinguishedNameDataValueEntitlementInfosValueFirewallRulesInner) HasPorts() bool {
	if o != nil && o.Ports != nil {
		return true
	}

	return false
}

// SetPorts gets a reference to the given []string and assigns it to the Ports field.
func (o *SessionInfoDistinguishedNameDataValueEntitlementInfosValueFirewallRulesInner) SetPorts(v []string) {
	o.Ports = v
}

// GetTypes returns the Types field value if set, zero value otherwise.
func (o *SessionInfoDistinguishedNameDataValueEntitlementInfosValueFirewallRulesInner) GetTypes() []string {
	if o == nil || o.Types == nil {
		var ret []string
		return ret
	}
	return o.Types
}

// GetTypesOk returns a tuple with the Types field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionInfoDistinguishedNameDataValueEntitlementInfosValueFirewallRulesInner) GetTypesOk() ([]string, bool) {
	if o == nil || o.Types == nil {
		return nil, false
	}
	return o.Types, true
}

// HasTypes returns a boolean if a field has been set.
func (o *SessionInfoDistinguishedNameDataValueEntitlementInfosValueFirewallRulesInner) HasTypes() bool {
	if o != nil && o.Types != nil {
		return true
	}

	return false
}

// SetTypes gets a reference to the given []string and assigns it to the Types field.
func (o *SessionInfoDistinguishedNameDataValueEntitlementInfosValueFirewallRulesInner) SetTypes(v []string) {
	o.Types = v
}

func (o SessionInfoDistinguishedNameDataValueEntitlementInfosValueFirewallRulesInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Protocol != nil {
		toSerialize["protocol"] = o.Protocol
	}
	if o.Direction != nil {
		toSerialize["direction"] = o.Direction
	}
	if o.Action != nil {
		toSerialize["action"] = o.Action
	}
	if o.Subnets != nil {
		toSerialize["subnets"] = o.Subnets
	}
	if o.Urls != nil {
		toSerialize["urls"] = o.Urls
	}
	if o.Ports != nil {
		toSerialize["ports"] = o.Ports
	}
	if o.Types != nil {
		toSerialize["types"] = o.Types
	}
	return json.Marshal(toSerialize)
}

type NullableSessionInfoDistinguishedNameDataValueEntitlementInfosValueFirewallRulesInner struct {
	value *SessionInfoDistinguishedNameDataValueEntitlementInfosValueFirewallRulesInner
	isSet bool
}

func (v NullableSessionInfoDistinguishedNameDataValueEntitlementInfosValueFirewallRulesInner) Get() *SessionInfoDistinguishedNameDataValueEntitlementInfosValueFirewallRulesInner {
	return v.value
}

func (v *NullableSessionInfoDistinguishedNameDataValueEntitlementInfosValueFirewallRulesInner) Set(val *SessionInfoDistinguishedNameDataValueEntitlementInfosValueFirewallRulesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableSessionInfoDistinguishedNameDataValueEntitlementInfosValueFirewallRulesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableSessionInfoDistinguishedNameDataValueEntitlementInfosValueFirewallRulesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSessionInfoDistinguishedNameDataValueEntitlementInfosValueFirewallRulesInner(val *SessionInfoDistinguishedNameDataValueEntitlementInfosValueFirewallRulesInner) *NullableSessionInfoDistinguishedNameDataValueEntitlementInfosValueFirewallRulesInner {
	return &NullableSessionInfoDistinguishedNameDataValueEntitlementInfosValueFirewallRulesInner{value: val, isSet: true}
}

func (v NullableSessionInfoDistinguishedNameDataValueEntitlementInfosValueFirewallRulesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSessionInfoDistinguishedNameDataValueEntitlementInfosValueFirewallRulesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
