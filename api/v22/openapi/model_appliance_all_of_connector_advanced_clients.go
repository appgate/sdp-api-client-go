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

// ApplianceAllOfConnectorAdvancedClients struct for ApplianceAllOfConnectorAdvancedClients
type ApplianceAllOfConnectorAdvancedClients struct {
	// Name for the Client. It will be mapped to the user claim 'clientName'.
	Name string `json:"name"`
	// The device ID to assign to this Client. It will be used to generate device distinguished name.
	DeviceId *string `json:"deviceId,omitempty"`
	// Source configuration to allow via iptables.
	AllowResources []AllowSourcesInner `json:"allowResources,omitempty"`
	// Use Source NAT for the Client tunnel.
	SnatToTunnel *bool `json:"snatToTunnel,omitempty"`
	// Use Source NAT for the resources.
	SnatToResources *bool `json:"snatToResources,omitempty"`
	// Apply destination NAT to traffic from tunnel into a resource
	DnatToResource *bool `json:"dnatToResource,omitempty"`
	// Use this connector client as a default gw for local resources
	DefaultGateway *bool                             `json:"defaultGateway,omitempty"`
	DhcpRelay      *ApplianceAllOfConnectorDhcpRelay `json:"dhcpRelay,omitempty"`
}

// NewApplianceAllOfConnectorAdvancedClients instantiates a new ApplianceAllOfConnectorAdvancedClients object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplianceAllOfConnectorAdvancedClients(name string) *ApplianceAllOfConnectorAdvancedClients {
	this := ApplianceAllOfConnectorAdvancedClients{}
	this.Name = name
	var snatToTunnel bool = true
	this.SnatToTunnel = &snatToTunnel
	var snatToResources bool = true
	this.SnatToResources = &snatToResources
	var dnatToResource bool = false
	this.DnatToResource = &dnatToResource
	var defaultGateway bool = false
	this.DefaultGateway = &defaultGateway
	return &this
}

// NewApplianceAllOfConnectorAdvancedClientsWithDefaults instantiates a new ApplianceAllOfConnectorAdvancedClients object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplianceAllOfConnectorAdvancedClientsWithDefaults() *ApplianceAllOfConnectorAdvancedClients {
	this := ApplianceAllOfConnectorAdvancedClients{}
	var snatToTunnel bool = true
	this.SnatToTunnel = &snatToTunnel
	var snatToResources bool = true
	this.SnatToResources = &snatToResources
	var dnatToResource bool = false
	this.DnatToResource = &dnatToResource
	var defaultGateway bool = false
	this.DefaultGateway = &defaultGateway
	return &this
}

// GetName returns the Name field value
func (o *ApplianceAllOfConnectorAdvancedClients) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ApplianceAllOfConnectorAdvancedClients) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ApplianceAllOfConnectorAdvancedClients) SetName(v string) {
	o.Name = v
}

// GetDeviceId returns the DeviceId field value if set, zero value otherwise.
func (o *ApplianceAllOfConnectorAdvancedClients) GetDeviceId() string {
	if o == nil || o.DeviceId == nil {
		var ret string
		return ret
	}
	return *o.DeviceId
}

// GetDeviceIdOk returns a tuple with the DeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceAllOfConnectorAdvancedClients) GetDeviceIdOk() (*string, bool) {
	if o == nil || o.DeviceId == nil {
		return nil, false
	}
	return o.DeviceId, true
}

// HasDeviceId returns a boolean if a field has been set.
func (o *ApplianceAllOfConnectorAdvancedClients) HasDeviceId() bool {
	if o != nil && o.DeviceId != nil {
		return true
	}

	return false
}

// SetDeviceId gets a reference to the given string and assigns it to the DeviceId field.
func (o *ApplianceAllOfConnectorAdvancedClients) SetDeviceId(v string) {
	o.DeviceId = &v
}

// GetAllowResources returns the AllowResources field value if set, zero value otherwise.
func (o *ApplianceAllOfConnectorAdvancedClients) GetAllowResources() []AllowSourcesInner {
	if o == nil || o.AllowResources == nil {
		var ret []AllowSourcesInner
		return ret
	}
	return o.AllowResources
}

// GetAllowResourcesOk returns a tuple with the AllowResources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceAllOfConnectorAdvancedClients) GetAllowResourcesOk() ([]AllowSourcesInner, bool) {
	if o == nil || o.AllowResources == nil {
		return nil, false
	}
	return o.AllowResources, true
}

// HasAllowResources returns a boolean if a field has been set.
func (o *ApplianceAllOfConnectorAdvancedClients) HasAllowResources() bool {
	if o != nil && o.AllowResources != nil {
		return true
	}

	return false
}

// SetAllowResources gets a reference to the given []AllowSourcesInner and assigns it to the AllowResources field.
func (o *ApplianceAllOfConnectorAdvancedClients) SetAllowResources(v []AllowSourcesInner) {
	o.AllowResources = v
}

// GetSnatToTunnel returns the SnatToTunnel field value if set, zero value otherwise.
func (o *ApplianceAllOfConnectorAdvancedClients) GetSnatToTunnel() bool {
	if o == nil || o.SnatToTunnel == nil {
		var ret bool
		return ret
	}
	return *o.SnatToTunnel
}

// GetSnatToTunnelOk returns a tuple with the SnatToTunnel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceAllOfConnectorAdvancedClients) GetSnatToTunnelOk() (*bool, bool) {
	if o == nil || o.SnatToTunnel == nil {
		return nil, false
	}
	return o.SnatToTunnel, true
}

// HasSnatToTunnel returns a boolean if a field has been set.
func (o *ApplianceAllOfConnectorAdvancedClients) HasSnatToTunnel() bool {
	if o != nil && o.SnatToTunnel != nil {
		return true
	}

	return false
}

// SetSnatToTunnel gets a reference to the given bool and assigns it to the SnatToTunnel field.
func (o *ApplianceAllOfConnectorAdvancedClients) SetSnatToTunnel(v bool) {
	o.SnatToTunnel = &v
}

// GetSnatToResources returns the SnatToResources field value if set, zero value otherwise.
func (o *ApplianceAllOfConnectorAdvancedClients) GetSnatToResources() bool {
	if o == nil || o.SnatToResources == nil {
		var ret bool
		return ret
	}
	return *o.SnatToResources
}

// GetSnatToResourcesOk returns a tuple with the SnatToResources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceAllOfConnectorAdvancedClients) GetSnatToResourcesOk() (*bool, bool) {
	if o == nil || o.SnatToResources == nil {
		return nil, false
	}
	return o.SnatToResources, true
}

// HasSnatToResources returns a boolean if a field has been set.
func (o *ApplianceAllOfConnectorAdvancedClients) HasSnatToResources() bool {
	if o != nil && o.SnatToResources != nil {
		return true
	}

	return false
}

// SetSnatToResources gets a reference to the given bool and assigns it to the SnatToResources field.
func (o *ApplianceAllOfConnectorAdvancedClients) SetSnatToResources(v bool) {
	o.SnatToResources = &v
}

// GetDnatToResource returns the DnatToResource field value if set, zero value otherwise.
func (o *ApplianceAllOfConnectorAdvancedClients) GetDnatToResource() bool {
	if o == nil || o.DnatToResource == nil {
		var ret bool
		return ret
	}
	return *o.DnatToResource
}

// GetDnatToResourceOk returns a tuple with the DnatToResource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceAllOfConnectorAdvancedClients) GetDnatToResourceOk() (*bool, bool) {
	if o == nil || o.DnatToResource == nil {
		return nil, false
	}
	return o.DnatToResource, true
}

// HasDnatToResource returns a boolean if a field has been set.
func (o *ApplianceAllOfConnectorAdvancedClients) HasDnatToResource() bool {
	if o != nil && o.DnatToResource != nil {
		return true
	}

	return false
}

// SetDnatToResource gets a reference to the given bool and assigns it to the DnatToResource field.
func (o *ApplianceAllOfConnectorAdvancedClients) SetDnatToResource(v bool) {
	o.DnatToResource = &v
}

// GetDefaultGateway returns the DefaultGateway field value if set, zero value otherwise.
func (o *ApplianceAllOfConnectorAdvancedClients) GetDefaultGateway() bool {
	if o == nil || o.DefaultGateway == nil {
		var ret bool
		return ret
	}
	return *o.DefaultGateway
}

// GetDefaultGatewayOk returns a tuple with the DefaultGateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceAllOfConnectorAdvancedClients) GetDefaultGatewayOk() (*bool, bool) {
	if o == nil || o.DefaultGateway == nil {
		return nil, false
	}
	return o.DefaultGateway, true
}

// HasDefaultGateway returns a boolean if a field has been set.
func (o *ApplianceAllOfConnectorAdvancedClients) HasDefaultGateway() bool {
	if o != nil && o.DefaultGateway != nil {
		return true
	}

	return false
}

// SetDefaultGateway gets a reference to the given bool and assigns it to the DefaultGateway field.
func (o *ApplianceAllOfConnectorAdvancedClients) SetDefaultGateway(v bool) {
	o.DefaultGateway = &v
}

// GetDhcpRelay returns the DhcpRelay field value if set, zero value otherwise.
func (o *ApplianceAllOfConnectorAdvancedClients) GetDhcpRelay() ApplianceAllOfConnectorDhcpRelay {
	if o == nil || o.DhcpRelay == nil {
		var ret ApplianceAllOfConnectorDhcpRelay
		return ret
	}
	return *o.DhcpRelay
}

// GetDhcpRelayOk returns a tuple with the DhcpRelay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceAllOfConnectorAdvancedClients) GetDhcpRelayOk() (*ApplianceAllOfConnectorDhcpRelay, bool) {
	if o == nil || o.DhcpRelay == nil {
		return nil, false
	}
	return o.DhcpRelay, true
}

// HasDhcpRelay returns a boolean if a field has been set.
func (o *ApplianceAllOfConnectorAdvancedClients) HasDhcpRelay() bool {
	if o != nil && o.DhcpRelay != nil {
		return true
	}

	return false
}

// SetDhcpRelay gets a reference to the given ApplianceAllOfConnectorDhcpRelay and assigns it to the DhcpRelay field.
func (o *ApplianceAllOfConnectorAdvancedClients) SetDhcpRelay(v ApplianceAllOfConnectorDhcpRelay) {
	o.DhcpRelay = &v
}

func (o ApplianceAllOfConnectorAdvancedClients) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.DeviceId != nil {
		toSerialize["deviceId"] = o.DeviceId
	}
	if o.AllowResources != nil {
		toSerialize["allowResources"] = o.AllowResources
	}
	if o.SnatToTunnel != nil {
		toSerialize["snatToTunnel"] = o.SnatToTunnel
	}
	if o.SnatToResources != nil {
		toSerialize["snatToResources"] = o.SnatToResources
	}
	if o.DnatToResource != nil {
		toSerialize["dnatToResource"] = o.DnatToResource
	}
	if o.DefaultGateway != nil {
		toSerialize["defaultGateway"] = o.DefaultGateway
	}
	if o.DhcpRelay != nil {
		toSerialize["dhcpRelay"] = o.DhcpRelay
	}
	return json.Marshal(toSerialize)
}

type NullableApplianceAllOfConnectorAdvancedClients struct {
	value *ApplianceAllOfConnectorAdvancedClients
	isSet bool
}

func (v NullableApplianceAllOfConnectorAdvancedClients) Get() *ApplianceAllOfConnectorAdvancedClients {
	return v.value
}

func (v *NullableApplianceAllOfConnectorAdvancedClients) Set(val *ApplianceAllOfConnectorAdvancedClients) {
	v.value = val
	v.isSet = true
}

func (v NullableApplianceAllOfConnectorAdvancedClients) IsSet() bool {
	return v.isSet
}

func (v *NullableApplianceAllOfConnectorAdvancedClients) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplianceAllOfConnectorAdvancedClients(val *ApplianceAllOfConnectorAdvancedClients) *NullableApplianceAllOfConnectorAdvancedClients {
	return &NullableApplianceAllOfConnectorAdvancedClients{value: val, isSet: true}
}

func (v NullableApplianceAllOfConnectorAdvancedClients) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplianceAllOfConnectorAdvancedClients) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
