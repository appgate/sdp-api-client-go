/*
 * AppGate SDP Controller REST API
 *
 * # About   This specification documents the REST API calls for the AppGate SDP Controller.    Please refer to the Integration chapter in the manual or contact AppGate support with any questions about   this functionality. # Getting Started   Requirements for API scripting:   - Access to the peer interface (default port 444) of a Controller appliance.     (https://sdphelp.appgate.com/adminguide/appliances-configure.html?anchor=peer)   - An API user with relevant permissions.     (https://sdphelp.appgate.com/adminguide/administrative-roles-configure.html)   - In order to use the simple login API, Admin MFA must be disabled or the API user must be excluded.     (https://sdphelp.appgate.com/adminguide/mfa-for-admins.html) # Base path   HTTPS requests must be sent to the Peer Interface hostname and port, with **_/admin** path.    For example: **https://appgate.company.com:444/admin**    All requests must have the **Accept** header as:    **application/vnd.appgate.peer-v12+json** # API Conventions   API conventions are  important to understand and follow strictly.    - While updating objects (via PUT), entire object must be sent with all fields.     - For example, if in order to add a remedy method to the condition below:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": []       }       ```     - send the entire object with updated and non-updated fields:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": [{\"type\": \"DisplayMessage\", \"message\": \"test message\"}]       }       ```    - In case Controller returns an error (non-2xx HTTP status code), response body is JSON.     The \"message\" field contains information abot the error.     HTTP 422 \"Unprocessable Entity\" has extra `errors` field to list all the issues with specific fields.    - Empty string (\"\") is considered a different value than \"null\" or field being omitted from JSON.     Omitting the field is recommend if no value is intended.     Empty string (\"\") will be almost always rejected as invalid value.    - There are common pattern between many objects:     - **Configuration Objects**: There are many objects with common fields, namely \"id\", \"name\", \"notes\", \"created\"       and \"updated\". These entities are listed, queried, created, updated and deleted in a similar fashion.     - **Distinguished Name**: Users and Devices are identified with what is called Distinguished Names, as used in        LDAP. The distinguished format that identifies a device and a user combination is        \"CN=\\<Device ID\\>,CN=\\<username\\>,OU=\\<Identity Provider Name\\>\". Some objects have the        \"userDistinguishedName\" field, which does not include the CN for Device ID.        This identifies a user on every device.
 *
 * API version: API version 12
 * Contact: appgatesdp.support@appgate.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"encoding/json"
)

// ApplianceAllOfNetworking Networking configuration of the system.
type ApplianceAllOfNetworking struct {
	// /etc/hosts configuration
	Hosts *[]ApplianceAllOfNetworkingHosts `json:"hosts,omitempty"`
	// System NIC configuration
	Nics *[]ApplianceAllOfNetworkingNics `json:"nics,omitempty"`
	// DNS Server addresses.
	DnsServers *[]string `json:"dnsServers,omitempty"`
	// DNS Search domains.
	DnsDomains *[]string `json:"dnsDomains,omitempty"`
	// System route settings.
	Routes *[]ApplianceAllOfNetworkingRoutes `json:"routes,omitempty"`
}

// NewApplianceAllOfNetworking instantiates a new ApplianceAllOfNetworking object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplianceAllOfNetworking() *ApplianceAllOfNetworking {
	this := ApplianceAllOfNetworking{}
	return &this
}

// NewApplianceAllOfNetworkingWithDefaults instantiates a new ApplianceAllOfNetworking object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplianceAllOfNetworkingWithDefaults() *ApplianceAllOfNetworking {
	this := ApplianceAllOfNetworking{}
	return &this
}

// GetHosts returns the Hosts field value if set, zero value otherwise.
func (o *ApplianceAllOfNetworking) GetHosts() []ApplianceAllOfNetworkingHosts {
	if o == nil || o.Hosts == nil {
		var ret []ApplianceAllOfNetworkingHosts
		return ret
	}
	return *o.Hosts
}

// GetHostsOk returns a tuple with the Hosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceAllOfNetworking) GetHostsOk() (*[]ApplianceAllOfNetworkingHosts, bool) {
	if o == nil || o.Hosts == nil {
		return nil, false
	}
	return o.Hosts, true
}

// HasHosts returns a boolean if a field has been set.
func (o *ApplianceAllOfNetworking) HasHosts() bool {
	if o != nil && o.Hosts != nil {
		return true
	}

	return false
}

// SetHosts gets a reference to the given []ApplianceAllOfNetworkingHosts and assigns it to the Hosts field.
func (o *ApplianceAllOfNetworking) SetHosts(v []ApplianceAllOfNetworkingHosts) {
	o.Hosts = &v
}

// GetNics returns the Nics field value if set, zero value otherwise.
func (o *ApplianceAllOfNetworking) GetNics() []ApplianceAllOfNetworkingNics {
	if o == nil || o.Nics == nil {
		var ret []ApplianceAllOfNetworkingNics
		return ret
	}
	return *o.Nics
}

// GetNicsOk returns a tuple with the Nics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceAllOfNetworking) GetNicsOk() (*[]ApplianceAllOfNetworkingNics, bool) {
	if o == nil || o.Nics == nil {
		return nil, false
	}
	return o.Nics, true
}

// HasNics returns a boolean if a field has been set.
func (o *ApplianceAllOfNetworking) HasNics() bool {
	if o != nil && o.Nics != nil {
		return true
	}

	return false
}

// SetNics gets a reference to the given []ApplianceAllOfNetworkingNics and assigns it to the Nics field.
func (o *ApplianceAllOfNetworking) SetNics(v []ApplianceAllOfNetworkingNics) {
	o.Nics = &v
}

// GetDnsServers returns the DnsServers field value if set, zero value otherwise.
func (o *ApplianceAllOfNetworking) GetDnsServers() []string {
	if o == nil || o.DnsServers == nil {
		var ret []string
		return ret
	}
	return *o.DnsServers
}

// GetDnsServersOk returns a tuple with the DnsServers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceAllOfNetworking) GetDnsServersOk() (*[]string, bool) {
	if o == nil || o.DnsServers == nil {
		return nil, false
	}
	return o.DnsServers, true
}

// HasDnsServers returns a boolean if a field has been set.
func (o *ApplianceAllOfNetworking) HasDnsServers() bool {
	if o != nil && o.DnsServers != nil {
		return true
	}

	return false
}

// SetDnsServers gets a reference to the given []string and assigns it to the DnsServers field.
func (o *ApplianceAllOfNetworking) SetDnsServers(v []string) {
	o.DnsServers = &v
}

// GetDnsDomains returns the DnsDomains field value if set, zero value otherwise.
func (o *ApplianceAllOfNetworking) GetDnsDomains() []string {
	if o == nil || o.DnsDomains == nil {
		var ret []string
		return ret
	}
	return *o.DnsDomains
}

// GetDnsDomainsOk returns a tuple with the DnsDomains field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceAllOfNetworking) GetDnsDomainsOk() (*[]string, bool) {
	if o == nil || o.DnsDomains == nil {
		return nil, false
	}
	return o.DnsDomains, true
}

// HasDnsDomains returns a boolean if a field has been set.
func (o *ApplianceAllOfNetworking) HasDnsDomains() bool {
	if o != nil && o.DnsDomains != nil {
		return true
	}

	return false
}

// SetDnsDomains gets a reference to the given []string and assigns it to the DnsDomains field.
func (o *ApplianceAllOfNetworking) SetDnsDomains(v []string) {
	o.DnsDomains = &v
}

// GetRoutes returns the Routes field value if set, zero value otherwise.
func (o *ApplianceAllOfNetworking) GetRoutes() []ApplianceAllOfNetworkingRoutes {
	if o == nil || o.Routes == nil {
		var ret []ApplianceAllOfNetworkingRoutes
		return ret
	}
	return *o.Routes
}

// GetRoutesOk returns a tuple with the Routes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceAllOfNetworking) GetRoutesOk() (*[]ApplianceAllOfNetworkingRoutes, bool) {
	if o == nil || o.Routes == nil {
		return nil, false
	}
	return o.Routes, true
}

// HasRoutes returns a boolean if a field has been set.
func (o *ApplianceAllOfNetworking) HasRoutes() bool {
	if o != nil && o.Routes != nil {
		return true
	}

	return false
}

// SetRoutes gets a reference to the given []ApplianceAllOfNetworkingRoutes and assigns it to the Routes field.
func (o *ApplianceAllOfNetworking) SetRoutes(v []ApplianceAllOfNetworkingRoutes) {
	o.Routes = &v
}

func (o ApplianceAllOfNetworking) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Hosts != nil {
		toSerialize["hosts"] = o.Hosts
	}
	if o.Nics != nil {
		toSerialize["nics"] = o.Nics
	}
	if o.DnsServers != nil {
		toSerialize["dnsServers"] = o.DnsServers
	}
	if o.DnsDomains != nil {
		toSerialize["dnsDomains"] = o.DnsDomains
	}
	if o.Routes != nil {
		toSerialize["routes"] = o.Routes
	}
	return json.Marshal(toSerialize)
}

type NullableApplianceAllOfNetworking struct {
	value *ApplianceAllOfNetworking
	isSet bool
}

func (v NullableApplianceAllOfNetworking) Get() *ApplianceAllOfNetworking {
	return v.value
}

func (v *NullableApplianceAllOfNetworking) Set(val *ApplianceAllOfNetworking) {
	v.value = val
	v.isSet = true
}

func (v NullableApplianceAllOfNetworking) IsSet() bool {
	return v.isSet
}

func (v *NullableApplianceAllOfNetworking) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplianceAllOfNetworking(val *ApplianceAllOfNetworking) *NullableApplianceAllOfNetworking {
	return &NullableApplianceAllOfNetworking{value: val, isSet: true}
}

func (v NullableApplianceAllOfNetworking) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplianceAllOfNetworking) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
