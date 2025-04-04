/*
Appgate SDP Controller REST API

# About   This specification documents the REST API calls for the Appgate SDP Controller.    Please refer to the REST API chapter in the manual or contact Appgate support with any questions about   this functionality. # Getting Started   Requirements for API scripting:   - Access to the Admin/API TLS Connection (default port 8443) of a Controller appliance.     (https://sdphelp.appgate.com/adminguide/appliance-function-configure.html?anchor=admin-api)   - An API user with relevant permissions.     (https://sdphelp.appgate.com/adminguide/administrative-roles-configure.html)   - In order to use the simple login API, Admin MFA must be disabled or the API user must be excluded.     (https://sdphelp.appgate.com/adminguide/mfa-for-admins.html) # Base path   HTTPS requests must be sent to the Admin Interface hostname and port, with **_/admin** path.    For example: **https://appgate.company.com:8443/admin**    All requests must have the **Accept** header as:    **application/vnd.appgate.peer-v22+json**    An exception is made for the **_/admin/version** endpoint which instead expects an **application/json** Accept header. # API Conventions   API conventions are  important to understand and follow strictly.    - While updating objects (via PUT), entire object must be sent with all fields.     - For example, in order to add a remedy method to the condition below:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": []       }       ```     - send the entire object with updated and non-updated fields:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": [{\"type\": \"DisplayMessage\", \"message\": \"test message\"}]       }       ```    - In case Controller returns an error (non-2xx HTTP status code), response body is JSON.     The \"message\" field contains information about the error.     HTTP 422 \"Unprocessable Entity\" has extra `errors` field to list all the issues with specific fields.    - Empty string (\"\") is considered a different value than \"null\" or field being omitted from JSON.     Omitting the field is recommended if no value is intended.     Empty string (\"\") will be almost always rejected as invalid value.    - There are common pattern between many objects:     - **Configuration Objects**: There are many objects with common fields, namely \"id\", \"name\", \"notes\", \"created\"       and \"updated\". These entities are listed, queried, created, updated and deleted in a similar fashion.     - **Distinguished Name**: Users and Devices are identified with what is called Distinguished Names, as used in        LDAP. The distinguished format that identifies a device and a user combination is        \"CN=\\<Device ID\\>,CN=\\<username\\>,OU=\\<Identity Provider Name\\>\". Some objects have the        \"userDistinguishedName\" field, which does not include the CN for Device ID.        This identifies a user on every device.

API version: API version 22.0
Contact: appgatesdp.support@appgate.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// SiteAllOfNameResolutionDnsForwarding DNS Forwarding feature. Always enabled and will be filled if there is no object is passed.
type SiteAllOfNameResolutionDnsForwarding struct {
	// DNS Forwarder Site IPv4 address.
	SiteIpv4 *string `json:"siteIpv4,omitempty"`
	// DNS Forwarder Site IPv6 address.
	SiteIpv6 *string `json:"siteIpv6,omitempty"`
	// DNS Servers to use for resolving endpoints. Leave it empty to use the Gateways' own DNS configuration.
	DnsServers []string `json:"dnsServers,omitempty"`
	// A list of subnets to allow access.
	AllowDestinations []AllowResourcesInner `json:"allowDestinations"`
	// Deprecated as of 6.4. This will apply whenever Gateway gets a DNS response which has no TTL set.
	// Deprecated
	DefaultTtlSeconds *int32 `json:"defaultTtlSeconds,omitempty"`
	// The match domains to use for automatic Client DNS configuration.
	MatchDomains []string `json:"matchDomains,omitempty"`
	// This will configure Client machines' DNS according to this forwarder if the Client connects to this Site.
	AutoClientDns *bool `json:"autoClientDns,omitempty"`
}

// NewSiteAllOfNameResolutionDnsForwarding instantiates a new SiteAllOfNameResolutionDnsForwarding object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSiteAllOfNameResolutionDnsForwarding(allowDestinations []AllowResourcesInner) *SiteAllOfNameResolutionDnsForwarding {
	this := SiteAllOfNameResolutionDnsForwarding{}
	this.AllowDestinations = allowDestinations
	var defaultTtlSeconds int32 = 300
	this.DefaultTtlSeconds = &defaultTtlSeconds
	return &this
}

// NewSiteAllOfNameResolutionDnsForwardingWithDefaults instantiates a new SiteAllOfNameResolutionDnsForwarding object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSiteAllOfNameResolutionDnsForwardingWithDefaults() *SiteAllOfNameResolutionDnsForwarding {
	this := SiteAllOfNameResolutionDnsForwarding{}
	var defaultTtlSeconds int32 = 300
	this.DefaultTtlSeconds = &defaultTtlSeconds
	return &this
}

// GetSiteIpv4 returns the SiteIpv4 field value if set, zero value otherwise.
func (o *SiteAllOfNameResolutionDnsForwarding) GetSiteIpv4() string {
	if o == nil || o.SiteIpv4 == nil {
		var ret string
		return ret
	}
	return *o.SiteIpv4
}

// GetSiteIpv4Ok returns a tuple with the SiteIpv4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteAllOfNameResolutionDnsForwarding) GetSiteIpv4Ok() (*string, bool) {
	if o == nil || o.SiteIpv4 == nil {
		return nil, false
	}
	return o.SiteIpv4, true
}

// HasSiteIpv4 returns a boolean if a field has been set.
func (o *SiteAllOfNameResolutionDnsForwarding) HasSiteIpv4() bool {
	if o != nil && o.SiteIpv4 != nil {
		return true
	}

	return false
}

// SetSiteIpv4 gets a reference to the given string and assigns it to the SiteIpv4 field.
func (o *SiteAllOfNameResolutionDnsForwarding) SetSiteIpv4(v string) {
	o.SiteIpv4 = &v
}

// GetSiteIpv6 returns the SiteIpv6 field value if set, zero value otherwise.
func (o *SiteAllOfNameResolutionDnsForwarding) GetSiteIpv6() string {
	if o == nil || o.SiteIpv6 == nil {
		var ret string
		return ret
	}
	return *o.SiteIpv6
}

// GetSiteIpv6Ok returns a tuple with the SiteIpv6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteAllOfNameResolutionDnsForwarding) GetSiteIpv6Ok() (*string, bool) {
	if o == nil || o.SiteIpv6 == nil {
		return nil, false
	}
	return o.SiteIpv6, true
}

// HasSiteIpv6 returns a boolean if a field has been set.
func (o *SiteAllOfNameResolutionDnsForwarding) HasSiteIpv6() bool {
	if o != nil && o.SiteIpv6 != nil {
		return true
	}

	return false
}

// SetSiteIpv6 gets a reference to the given string and assigns it to the SiteIpv6 field.
func (o *SiteAllOfNameResolutionDnsForwarding) SetSiteIpv6(v string) {
	o.SiteIpv6 = &v
}

// GetDnsServers returns the DnsServers field value if set, zero value otherwise.
func (o *SiteAllOfNameResolutionDnsForwarding) GetDnsServers() []string {
	if o == nil || o.DnsServers == nil {
		var ret []string
		return ret
	}
	return o.DnsServers
}

// GetDnsServersOk returns a tuple with the DnsServers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteAllOfNameResolutionDnsForwarding) GetDnsServersOk() ([]string, bool) {
	if o == nil || o.DnsServers == nil {
		return nil, false
	}
	return o.DnsServers, true
}

// HasDnsServers returns a boolean if a field has been set.
func (o *SiteAllOfNameResolutionDnsForwarding) HasDnsServers() bool {
	if o != nil && o.DnsServers != nil {
		return true
	}

	return false
}

// SetDnsServers gets a reference to the given []string and assigns it to the DnsServers field.
func (o *SiteAllOfNameResolutionDnsForwarding) SetDnsServers(v []string) {
	o.DnsServers = v
}

// GetAllowDestinations returns the AllowDestinations field value
func (o *SiteAllOfNameResolutionDnsForwarding) GetAllowDestinations() []AllowResourcesInner {
	if o == nil {
		var ret []AllowResourcesInner
		return ret
	}

	return o.AllowDestinations
}

// GetAllowDestinationsOk returns a tuple with the AllowDestinations field value
// and a boolean to check if the value has been set.
func (o *SiteAllOfNameResolutionDnsForwarding) GetAllowDestinationsOk() ([]AllowResourcesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.AllowDestinations, true
}

// SetAllowDestinations sets field value
func (o *SiteAllOfNameResolutionDnsForwarding) SetAllowDestinations(v []AllowResourcesInner) {
	o.AllowDestinations = v
}

// GetDefaultTtlSeconds returns the DefaultTtlSeconds field value if set, zero value otherwise.
// Deprecated
func (o *SiteAllOfNameResolutionDnsForwarding) GetDefaultTtlSeconds() int32 {
	if o == nil || o.DefaultTtlSeconds == nil {
		var ret int32
		return ret
	}
	return *o.DefaultTtlSeconds
}

// GetDefaultTtlSecondsOk returns a tuple with the DefaultTtlSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *SiteAllOfNameResolutionDnsForwarding) GetDefaultTtlSecondsOk() (*int32, bool) {
	if o == nil || o.DefaultTtlSeconds == nil {
		return nil, false
	}
	return o.DefaultTtlSeconds, true
}

// HasDefaultTtlSeconds returns a boolean if a field has been set.
func (o *SiteAllOfNameResolutionDnsForwarding) HasDefaultTtlSeconds() bool {
	if o != nil && o.DefaultTtlSeconds != nil {
		return true
	}

	return false
}

// SetDefaultTtlSeconds gets a reference to the given int32 and assigns it to the DefaultTtlSeconds field.
// Deprecated
func (o *SiteAllOfNameResolutionDnsForwarding) SetDefaultTtlSeconds(v int32) {
	o.DefaultTtlSeconds = &v
}

// GetMatchDomains returns the MatchDomains field value if set, zero value otherwise.
func (o *SiteAllOfNameResolutionDnsForwarding) GetMatchDomains() []string {
	if o == nil || o.MatchDomains == nil {
		var ret []string
		return ret
	}
	return o.MatchDomains
}

// GetMatchDomainsOk returns a tuple with the MatchDomains field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteAllOfNameResolutionDnsForwarding) GetMatchDomainsOk() ([]string, bool) {
	if o == nil || o.MatchDomains == nil {
		return nil, false
	}
	return o.MatchDomains, true
}

// HasMatchDomains returns a boolean if a field has been set.
func (o *SiteAllOfNameResolutionDnsForwarding) HasMatchDomains() bool {
	if o != nil && o.MatchDomains != nil {
		return true
	}

	return false
}

// SetMatchDomains gets a reference to the given []string and assigns it to the MatchDomains field.
func (o *SiteAllOfNameResolutionDnsForwarding) SetMatchDomains(v []string) {
	o.MatchDomains = v
}

// GetAutoClientDns returns the AutoClientDns field value if set, zero value otherwise.
func (o *SiteAllOfNameResolutionDnsForwarding) GetAutoClientDns() bool {
	if o == nil || o.AutoClientDns == nil {
		var ret bool
		return ret
	}
	return *o.AutoClientDns
}

// GetAutoClientDnsOk returns a tuple with the AutoClientDns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteAllOfNameResolutionDnsForwarding) GetAutoClientDnsOk() (*bool, bool) {
	if o == nil || o.AutoClientDns == nil {
		return nil, false
	}
	return o.AutoClientDns, true
}

// HasAutoClientDns returns a boolean if a field has been set.
func (o *SiteAllOfNameResolutionDnsForwarding) HasAutoClientDns() bool {
	if o != nil && o.AutoClientDns != nil {
		return true
	}

	return false
}

// SetAutoClientDns gets a reference to the given bool and assigns it to the AutoClientDns field.
func (o *SiteAllOfNameResolutionDnsForwarding) SetAutoClientDns(v bool) {
	o.AutoClientDns = &v
}

func (o SiteAllOfNameResolutionDnsForwarding) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SiteIpv4 != nil {
		toSerialize["siteIpv4"] = o.SiteIpv4
	}
	if o.SiteIpv6 != nil {
		toSerialize["siteIpv6"] = o.SiteIpv6
	}
	if o.DnsServers != nil {
		toSerialize["dnsServers"] = o.DnsServers
	}
	if true {
		toSerialize["allowDestinations"] = o.AllowDestinations
	}
	if o.DefaultTtlSeconds != nil {
		toSerialize["defaultTtlSeconds"] = o.DefaultTtlSeconds
	}
	if o.MatchDomains != nil {
		toSerialize["matchDomains"] = o.MatchDomains
	}
	if o.AutoClientDns != nil {
		toSerialize["autoClientDns"] = o.AutoClientDns
	}
	return json.Marshal(toSerialize)
}

type NullableSiteAllOfNameResolutionDnsForwarding struct {
	value *SiteAllOfNameResolutionDnsForwarding
	isSet bool
}

func (v NullableSiteAllOfNameResolutionDnsForwarding) Get() *SiteAllOfNameResolutionDnsForwarding {
	return v.value
}

func (v *NullableSiteAllOfNameResolutionDnsForwarding) Set(val *SiteAllOfNameResolutionDnsForwarding) {
	v.value = val
	v.isSet = true
}

func (v NullableSiteAllOfNameResolutionDnsForwarding) IsSet() bool {
	return v.isSet
}

func (v *NullableSiteAllOfNameResolutionDnsForwarding) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSiteAllOfNameResolutionDnsForwarding(val *SiteAllOfNameResolutionDnsForwarding) *NullableSiteAllOfNameResolutionDnsForwarding {
	return &NullableSiteAllOfNameResolutionDnsForwarding{value: val, isSet: true}
}

func (v NullableSiteAllOfNameResolutionDnsForwarding) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSiteAllOfNameResolutionDnsForwarding) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
