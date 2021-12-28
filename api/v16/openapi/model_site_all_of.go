/*
Appgate SDP Controller REST API

# About   This specification documents the REST API calls for the Appgate SDP Controller.    Please refer to the REST API chapter in the manual or contact Appgate support with any questions about   this functionality. # Getting Started   Requirements for API scripting:   - Access to the Admin/API TLS Connection (default port 8443) of a Controller appliance.     (https://sdphelp.appgate.com/adminguide/appliance-function-configure.html?anchor=admin-api)   - An API user with relevant permissions.     (https://sdphelp.appgate.com/adminguide/administrative-roles-configure.html)   - In order to use the simple login API, Admin MFA must be disabled or the API user must be excluded.     (https://sdphelp.appgate.com/adminguide/mfa-for-admins.html) # Base path   HTTPS requests must be sent to the Admin Interface hostname and port, with **_/admin** path.    For example: **https://appgate.company.com:8443/admin**    All requests must have the **Accept** header as:    **application/vnd.appgate.peer-v16+json** # API Conventions   API conventions are  important to understand and follow strictly.    - While updating objects (via PUT), entire object must be sent with all fields.     - For example, in order to add a remedy method to the condition below:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": []       }       ```     - send the entire object with updated and non-updated fields:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": [{\"type\": \"DisplayMessage\", \"message\": \"test message\"}]       }       ```    - In case Controller returns an error (non-2xx HTTP status code), response body is JSON.     The \"message\" field contains information about the error.     HTTP 422 \"Unprocessable Entity\" has extra `errors` field to list all the issues with specific fields.    - Empty string (\"\") is considered a different value than \"null\" or field being omitted from JSON.     Omitting the field is recommend if no value is intended.     Empty string (\"\") will be almost always rejected as invalid value.    - There are common pattern between many objects:     - **Configuration Objects**: There are many objects with common fields, namely \"id\", \"name\", \"notes\", \"created\"       and \"updated\". These entities are listed, queried, created, updated and deleted in a similar fashion.     - **Distinguished Name**: Users and Devices are identified with what is called Distinguished Names, as used in        LDAP. The distinguished format that identifies a device and a user combination is        \"CN=\\<Device ID\\>,CN=\\<username\\>,OU=\\<Identity Provider Name\\>\". Some objects have the        \"userDistinguishedName\" field, which does not include the CN for Device ID.        This identifies a user on every device.

API version: API version 16.2
Contact: appgatesdp.support@appgate.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// SiteAllOf Represents a Site.
type SiteAllOf struct {
	// A short 4 letter name for the Site to be displayed on the Client.
	ShortName *string `json:"shortName,omitempty"`
	// Description of the Site to be displayed on the Client.
	Description *string `json:"description,omitempty"`
	// Network subnets in CIDR format to define the Site's boundaries. They are added as routes by the Client.
	NetworkSubnets *[]string `json:"networkSubnets,omitempty"`
	// List of IP Pool mappings for this specific Site. When IPs are allocated this Site, they will be mapped to a new one using this setting.
	IpPoolMappings *[]SiteAllOfIpPoolMappings `json:"ipPoolMappings,omitempty"`
	DefaultGateway *SiteAllOfDefaultGateway   `json:"defaultGateway,omitempty"`
	// When enabled, the routes are sent to the Client by the Gateways according to the user's Entitlements \"networkSubnets\" should be left be empty if it's enabled.
	EntitlementBasedRouting *bool                    `json:"entitlementBasedRouting,omitempty"`
	Vpn                     *SiteAllOfVpn            `json:"vpn,omitempty"`
	NameResolution          *SiteAllOfNameResolution `json:"nameResolution,omitempty"`
}

// NewSiteAllOf instantiates a new SiteAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSiteAllOf() *SiteAllOf {
	this := SiteAllOf{}
	var entitlementBasedRouting bool = false
	this.EntitlementBasedRouting = &entitlementBasedRouting
	return &this
}

// NewSiteAllOfWithDefaults instantiates a new SiteAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSiteAllOfWithDefaults() *SiteAllOf {
	this := SiteAllOf{}
	var entitlementBasedRouting bool = false
	this.EntitlementBasedRouting = &entitlementBasedRouting
	return &this
}

// GetShortName returns the ShortName field value if set, zero value otherwise.
func (o *SiteAllOf) GetShortName() string {
	if o == nil || o.ShortName == nil {
		var ret string
		return ret
	}
	return *o.ShortName
}

// GetShortNameOk returns a tuple with the ShortName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteAllOf) GetShortNameOk() (*string, bool) {
	if o == nil || o.ShortName == nil {
		return nil, false
	}
	return o.ShortName, true
}

// HasShortName returns a boolean if a field has been set.
func (o *SiteAllOf) HasShortName() bool {
	if o != nil && o.ShortName != nil {
		return true
	}

	return false
}

// SetShortName gets a reference to the given string and assigns it to the ShortName field.
func (o *SiteAllOf) SetShortName(v string) {
	o.ShortName = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SiteAllOf) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteAllOf) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SiteAllOf) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SiteAllOf) SetDescription(v string) {
	o.Description = &v
}

// GetNetworkSubnets returns the NetworkSubnets field value if set, zero value otherwise.
func (o *SiteAllOf) GetNetworkSubnets() []string {
	if o == nil || o.NetworkSubnets == nil {
		var ret []string
		return ret
	}
	return *o.NetworkSubnets
}

// GetNetworkSubnetsOk returns a tuple with the NetworkSubnets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteAllOf) GetNetworkSubnetsOk() (*[]string, bool) {
	if o == nil || o.NetworkSubnets == nil {
		return nil, false
	}
	return o.NetworkSubnets, true
}

// HasNetworkSubnets returns a boolean if a field has been set.
func (o *SiteAllOf) HasNetworkSubnets() bool {
	if o != nil && o.NetworkSubnets != nil {
		return true
	}

	return false
}

// SetNetworkSubnets gets a reference to the given []string and assigns it to the NetworkSubnets field.
func (o *SiteAllOf) SetNetworkSubnets(v []string) {
	o.NetworkSubnets = &v
}

// GetIpPoolMappings returns the IpPoolMappings field value if set, zero value otherwise.
func (o *SiteAllOf) GetIpPoolMappings() []SiteAllOfIpPoolMappings {
	if o == nil || o.IpPoolMappings == nil {
		var ret []SiteAllOfIpPoolMappings
		return ret
	}
	return *o.IpPoolMappings
}

// GetIpPoolMappingsOk returns a tuple with the IpPoolMappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteAllOf) GetIpPoolMappingsOk() (*[]SiteAllOfIpPoolMappings, bool) {
	if o == nil || o.IpPoolMappings == nil {
		return nil, false
	}
	return o.IpPoolMappings, true
}

// HasIpPoolMappings returns a boolean if a field has been set.
func (o *SiteAllOf) HasIpPoolMappings() bool {
	if o != nil && o.IpPoolMappings != nil {
		return true
	}

	return false
}

// SetIpPoolMappings gets a reference to the given []SiteAllOfIpPoolMappings and assigns it to the IpPoolMappings field.
func (o *SiteAllOf) SetIpPoolMappings(v []SiteAllOfIpPoolMappings) {
	o.IpPoolMappings = &v
}

// GetDefaultGateway returns the DefaultGateway field value if set, zero value otherwise.
func (o *SiteAllOf) GetDefaultGateway() SiteAllOfDefaultGateway {
	if o == nil || o.DefaultGateway == nil {
		var ret SiteAllOfDefaultGateway
		return ret
	}
	return *o.DefaultGateway
}

// GetDefaultGatewayOk returns a tuple with the DefaultGateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteAllOf) GetDefaultGatewayOk() (*SiteAllOfDefaultGateway, bool) {
	if o == nil || o.DefaultGateway == nil {
		return nil, false
	}
	return o.DefaultGateway, true
}

// HasDefaultGateway returns a boolean if a field has been set.
func (o *SiteAllOf) HasDefaultGateway() bool {
	if o != nil && o.DefaultGateway != nil {
		return true
	}

	return false
}

// SetDefaultGateway gets a reference to the given SiteAllOfDefaultGateway and assigns it to the DefaultGateway field.
func (o *SiteAllOf) SetDefaultGateway(v SiteAllOfDefaultGateway) {
	o.DefaultGateway = &v
}

// GetEntitlementBasedRouting returns the EntitlementBasedRouting field value if set, zero value otherwise.
func (o *SiteAllOf) GetEntitlementBasedRouting() bool {
	if o == nil || o.EntitlementBasedRouting == nil {
		var ret bool
		return ret
	}
	return *o.EntitlementBasedRouting
}

// GetEntitlementBasedRoutingOk returns a tuple with the EntitlementBasedRouting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteAllOf) GetEntitlementBasedRoutingOk() (*bool, bool) {
	if o == nil || o.EntitlementBasedRouting == nil {
		return nil, false
	}
	return o.EntitlementBasedRouting, true
}

// HasEntitlementBasedRouting returns a boolean if a field has been set.
func (o *SiteAllOf) HasEntitlementBasedRouting() bool {
	if o != nil && o.EntitlementBasedRouting != nil {
		return true
	}

	return false
}

// SetEntitlementBasedRouting gets a reference to the given bool and assigns it to the EntitlementBasedRouting field.
func (o *SiteAllOf) SetEntitlementBasedRouting(v bool) {
	o.EntitlementBasedRouting = &v
}

// GetVpn returns the Vpn field value if set, zero value otherwise.
func (o *SiteAllOf) GetVpn() SiteAllOfVpn {
	if o == nil || o.Vpn == nil {
		var ret SiteAllOfVpn
		return ret
	}
	return *o.Vpn
}

// GetVpnOk returns a tuple with the Vpn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteAllOf) GetVpnOk() (*SiteAllOfVpn, bool) {
	if o == nil || o.Vpn == nil {
		return nil, false
	}
	return o.Vpn, true
}

// HasVpn returns a boolean if a field has been set.
func (o *SiteAllOf) HasVpn() bool {
	if o != nil && o.Vpn != nil {
		return true
	}

	return false
}

// SetVpn gets a reference to the given SiteAllOfVpn and assigns it to the Vpn field.
func (o *SiteAllOf) SetVpn(v SiteAllOfVpn) {
	o.Vpn = &v
}

// GetNameResolution returns the NameResolution field value if set, zero value otherwise.
func (o *SiteAllOf) GetNameResolution() SiteAllOfNameResolution {
	if o == nil || o.NameResolution == nil {
		var ret SiteAllOfNameResolution
		return ret
	}
	return *o.NameResolution
}

// GetNameResolutionOk returns a tuple with the NameResolution field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteAllOf) GetNameResolutionOk() (*SiteAllOfNameResolution, bool) {
	if o == nil || o.NameResolution == nil {
		return nil, false
	}
	return o.NameResolution, true
}

// HasNameResolution returns a boolean if a field has been set.
func (o *SiteAllOf) HasNameResolution() bool {
	if o != nil && o.NameResolution != nil {
		return true
	}

	return false
}

// SetNameResolution gets a reference to the given SiteAllOfNameResolution and assigns it to the NameResolution field.
func (o *SiteAllOf) SetNameResolution(v SiteAllOfNameResolution) {
	o.NameResolution = &v
}

func (o SiteAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ShortName != nil {
		toSerialize["shortName"] = o.ShortName
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.NetworkSubnets != nil {
		toSerialize["networkSubnets"] = o.NetworkSubnets
	}
	if o.IpPoolMappings != nil {
		toSerialize["ipPoolMappings"] = o.IpPoolMappings
	}
	if o.DefaultGateway != nil {
		toSerialize["defaultGateway"] = o.DefaultGateway
	}
	if o.EntitlementBasedRouting != nil {
		toSerialize["entitlementBasedRouting"] = o.EntitlementBasedRouting
	}
	if o.Vpn != nil {
		toSerialize["vpn"] = o.Vpn
	}
	if o.NameResolution != nil {
		toSerialize["nameResolution"] = o.NameResolution
	}
	return json.Marshal(toSerialize)
}

type NullableSiteAllOf struct {
	value *SiteAllOf
	isSet bool
}

func (v NullableSiteAllOf) Get() *SiteAllOf {
	return v.value
}

func (v *NullableSiteAllOf) Set(val *SiteAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSiteAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSiteAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSiteAllOf(val *SiteAllOf) *NullableSiteAllOf {
	return &NullableSiteAllOf{value: val, isSet: true}
}

func (v NullableSiteAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSiteAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
