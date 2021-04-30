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

// ConnectorProviderAllOf Represents a Connector Identity Provider.
type ConnectorProviderAllOf struct {
	// The type of the Identity Provider.
	Type string `json:"type"`
	// The IPv4 Pool ID the users in this Identity Provider are going to use to allocate IP addresses for the tunnels.
	IpPoolV4 *string `json:"ipPoolV4,omitempty"`
	// The IPv6 Pool ID the users in this Identity Provider are going to use to allocate IP addresses for the tunnels.
	IpPoolV6 *string `json:"ipPoolV6,omitempty"`
	// The mapping of Identity Provider attributes to claims.
	ClaimMappings *[]map[string]interface{} `json:"claimMappings,omitempty"`
	// The mapping of Identity Provider on demand attributes to claims.
	OnDemandClaimMappings *[]map[string]interface{} `json:"onDemandClaimMappings,omitempty"`
}

// NewConnectorProviderAllOf instantiates a new ConnectorProviderAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectorProviderAllOf(type_ string) *ConnectorProviderAllOf {
	this := ConnectorProviderAllOf{}
	this.Type = type_
	return &this
}

// NewConnectorProviderAllOfWithDefaults instantiates a new ConnectorProviderAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectorProviderAllOfWithDefaults() *ConnectorProviderAllOf {
	this := ConnectorProviderAllOf{}
	return &this
}

// GetType returns the Type field value
func (o *ConnectorProviderAllOf) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ConnectorProviderAllOf) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ConnectorProviderAllOf) SetType(v string) {
	o.Type = v
}

// GetIpPoolV4 returns the IpPoolV4 field value if set, zero value otherwise.
func (o *ConnectorProviderAllOf) GetIpPoolV4() string {
	if o == nil || o.IpPoolV4 == nil {
		var ret string
		return ret
	}
	return *o.IpPoolV4
}

// GetIpPoolV4Ok returns a tuple with the IpPoolV4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorProviderAllOf) GetIpPoolV4Ok() (*string, bool) {
	if o == nil || o.IpPoolV4 == nil {
		return nil, false
	}
	return o.IpPoolV4, true
}

// HasIpPoolV4 returns a boolean if a field has been set.
func (o *ConnectorProviderAllOf) HasIpPoolV4() bool {
	if o != nil && o.IpPoolV4 != nil {
		return true
	}

	return false
}

// SetIpPoolV4 gets a reference to the given string and assigns it to the IpPoolV4 field.
func (o *ConnectorProviderAllOf) SetIpPoolV4(v string) {
	o.IpPoolV4 = &v
}

// GetIpPoolV6 returns the IpPoolV6 field value if set, zero value otherwise.
func (o *ConnectorProviderAllOf) GetIpPoolV6() string {
	if o == nil || o.IpPoolV6 == nil {
		var ret string
		return ret
	}
	return *o.IpPoolV6
}

// GetIpPoolV6Ok returns a tuple with the IpPoolV6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorProviderAllOf) GetIpPoolV6Ok() (*string, bool) {
	if o == nil || o.IpPoolV6 == nil {
		return nil, false
	}
	return o.IpPoolV6, true
}

// HasIpPoolV6 returns a boolean if a field has been set.
func (o *ConnectorProviderAllOf) HasIpPoolV6() bool {
	if o != nil && o.IpPoolV6 != nil {
		return true
	}

	return false
}

// SetIpPoolV6 gets a reference to the given string and assigns it to the IpPoolV6 field.
func (o *ConnectorProviderAllOf) SetIpPoolV6(v string) {
	o.IpPoolV6 = &v
}

// GetClaimMappings returns the ClaimMappings field value if set, zero value otherwise.
func (o *ConnectorProviderAllOf) GetClaimMappings() []map[string]interface{} {
	if o == nil || o.ClaimMappings == nil {
		var ret []map[string]interface{}
		return ret
	}
	return *o.ClaimMappings
}

// GetClaimMappingsOk returns a tuple with the ClaimMappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorProviderAllOf) GetClaimMappingsOk() (*[]map[string]interface{}, bool) {
	if o == nil || o.ClaimMappings == nil {
		return nil, false
	}
	return o.ClaimMappings, true
}

// HasClaimMappings returns a boolean if a field has been set.
func (o *ConnectorProviderAllOf) HasClaimMappings() bool {
	if o != nil && o.ClaimMappings != nil {
		return true
	}

	return false
}

// SetClaimMappings gets a reference to the given []map[string]interface{} and assigns it to the ClaimMappings field.
func (o *ConnectorProviderAllOf) SetClaimMappings(v []map[string]interface{}) {
	o.ClaimMappings = &v
}

// GetOnDemandClaimMappings returns the OnDemandClaimMappings field value if set, zero value otherwise.
func (o *ConnectorProviderAllOf) GetOnDemandClaimMappings() []map[string]interface{} {
	if o == nil || o.OnDemandClaimMappings == nil {
		var ret []map[string]interface{}
		return ret
	}
	return *o.OnDemandClaimMappings
}

// GetOnDemandClaimMappingsOk returns a tuple with the OnDemandClaimMappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorProviderAllOf) GetOnDemandClaimMappingsOk() (*[]map[string]interface{}, bool) {
	if o == nil || o.OnDemandClaimMappings == nil {
		return nil, false
	}
	return o.OnDemandClaimMappings, true
}

// HasOnDemandClaimMappings returns a boolean if a field has been set.
func (o *ConnectorProviderAllOf) HasOnDemandClaimMappings() bool {
	if o != nil && o.OnDemandClaimMappings != nil {
		return true
	}

	return false
}

// SetOnDemandClaimMappings gets a reference to the given []map[string]interface{} and assigns it to the OnDemandClaimMappings field.
func (o *ConnectorProviderAllOf) SetOnDemandClaimMappings(v []map[string]interface{}) {
	o.OnDemandClaimMappings = &v
}

func (o ConnectorProviderAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if o.IpPoolV4 != nil {
		toSerialize["ipPoolV4"] = o.IpPoolV4
	}
	if o.IpPoolV6 != nil {
		toSerialize["ipPoolV6"] = o.IpPoolV6
	}
	if o.ClaimMappings != nil {
		toSerialize["claimMappings"] = o.ClaimMappings
	}
	if o.OnDemandClaimMappings != nil {
		toSerialize["onDemandClaimMappings"] = o.OnDemandClaimMappings
	}
	return json.Marshal(toSerialize)
}

type NullableConnectorProviderAllOf struct {
	value *ConnectorProviderAllOf
	isSet bool
}

func (v NullableConnectorProviderAllOf) Get() *ConnectorProviderAllOf {
	return v.value
}

func (v *NullableConnectorProviderAllOf) Set(val *ConnectorProviderAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectorProviderAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectorProviderAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectorProviderAllOf(val *ConnectorProviderAllOf) *NullableConnectorProviderAllOf {
	return &NullableConnectorProviderAllOf{value: val, isSet: true}
}

func (v NullableConnectorProviderAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectorProviderAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
