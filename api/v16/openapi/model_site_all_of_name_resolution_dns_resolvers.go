/*
Appgate SDP Controller REST API

# About   This specification documents the REST API calls for the Appgate SDP Controller.    Please refer to the REST API chapter in the manual or contact Appgate support with any questions about   this functionality. # Getting Started   Requirements for API scripting:   - Access to the Admin/API TLS Connection (default port 8443) of a Controller appliance.     (https://sdphelp.appgate.com/adminguide/appliance-function-configure.html?anchor=admin-api)   - An API user with relevant permissions.     (https://sdphelp.appgate.com/adminguide/administrative-roles-configure.html)   - In order to use the simple login API, Admin MFA must be disabled or the API user must be excluded.     (https://sdphelp.appgate.com/adminguide/mfa-for-admins.html) # Base path   HTTPS requests must be sent to the Admin Interface hostname and port, with **_/admin** path.    For example: **https://appgate.company.com:8443/admin**    All requests must have the **Accept** header as:    **application/vnd.appgate.peer-v16+json** # API Conventions   API conventions are  important to understand and follow strictly.    - While updating objects (via PUT), entire object must be sent with all fields.     - For example, in order to add a remedy method to the condition below:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": []       }       ```     - send the entire object with updated and non-updated fields:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": [{\"type\": \"DisplayMessage\", \"message\": \"test message\"}]       }       ```    - In case Controller returns an error (non-2xx HTTP status code), response body is JSON.     The \"message\" field contains information about the error.     HTTP 422 \"Unprocessable Entity\" has extra `errors` field to list all the issues with specific fields.    - Empty string (\"\") is considered a different value than \"null\" or field being omitted from JSON.     Omitting the field is recommend if no value is intended.     Empty string (\"\") will be almost always rejected as invalid value.    - There are common pattern between many objects:     - **Configuration Objects**: There are many objects with common fields, namely \"id\", \"name\", \"notes\", \"created\"       and \"updated\". These entities are listed, queried, created, updated and deleted in a similar fashion.     - **Distinguished Name**: Users and Devices are identified with what is called Distinguished Names, as used in        LDAP. The distinguished format that identifies a device and a user combination is        \"CN=\\<Device ID\\>,CN=\\<username\\>,OU=\\<Identity Provider Name\\>\". Some objects have the        \"userDistinguishedName\" field, which does not include the CN for Device ID.        This identifies a user on every device.

API version: API version 16.0
Contact: appgatesdp.support@appgate.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// SiteAllOfNameResolutionDnsResolvers struct for SiteAllOfNameResolutionDnsResolvers
type SiteAllOfNameResolutionDnsResolvers struct {
	// Identifier name. Has no functional effect.
	Name string `json:"name"`
	// How often will the resolver poll the server. In seconds.
	UpdateInterval *int32 `json:"updateInterval,omitempty"`
	// DNS Server addresses that will be used to resolve hostnames within the Site.
	Servers []string `json:"servers"`
	// DNS search domains that will be used to resolve hostnames within the Site.
	// Deprecated
	SearchDomains *[]string `json:"searchDomains,omitempty"`
	// The DNS resolver will only attempt to resolve names matching the match domains. If match domains are not specified the DNS resolver will attempt to resolve all hostnames.
	MatchDomains *[]string `json:"matchDomains,omitempty"`
}

// NewSiteAllOfNameResolutionDnsResolvers instantiates a new SiteAllOfNameResolutionDnsResolvers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSiteAllOfNameResolutionDnsResolvers(name string, servers []string) *SiteAllOfNameResolutionDnsResolvers {
	this := SiteAllOfNameResolutionDnsResolvers{}
	this.Name = name
	var updateInterval int32 = 60
	this.UpdateInterval = &updateInterval
	this.Servers = servers
	return &this
}

// NewSiteAllOfNameResolutionDnsResolversWithDefaults instantiates a new SiteAllOfNameResolutionDnsResolvers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSiteAllOfNameResolutionDnsResolversWithDefaults() *SiteAllOfNameResolutionDnsResolvers {
	this := SiteAllOfNameResolutionDnsResolvers{}
	var updateInterval int32 = 60
	this.UpdateInterval = &updateInterval
	return &this
}

// GetName returns the Name field value
func (o *SiteAllOfNameResolutionDnsResolvers) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SiteAllOfNameResolutionDnsResolvers) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SiteAllOfNameResolutionDnsResolvers) SetName(v string) {
	o.Name = v
}

// GetUpdateInterval returns the UpdateInterval field value if set, zero value otherwise.
func (o *SiteAllOfNameResolutionDnsResolvers) GetUpdateInterval() int32 {
	if o == nil || o.UpdateInterval == nil {
		var ret int32
		return ret
	}
	return *o.UpdateInterval
}

// GetUpdateIntervalOk returns a tuple with the UpdateInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteAllOfNameResolutionDnsResolvers) GetUpdateIntervalOk() (*int32, bool) {
	if o == nil || o.UpdateInterval == nil {
		return nil, false
	}
	return o.UpdateInterval, true
}

// HasUpdateInterval returns a boolean if a field has been set.
func (o *SiteAllOfNameResolutionDnsResolvers) HasUpdateInterval() bool {
	if o != nil && o.UpdateInterval != nil {
		return true
	}

	return false
}

// SetUpdateInterval gets a reference to the given int32 and assigns it to the UpdateInterval field.
func (o *SiteAllOfNameResolutionDnsResolvers) SetUpdateInterval(v int32) {
	o.UpdateInterval = &v
}

// GetServers returns the Servers field value
func (o *SiteAllOfNameResolutionDnsResolvers) GetServers() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Servers
}

// GetServersOk returns a tuple with the Servers field value
// and a boolean to check if the value has been set.
func (o *SiteAllOfNameResolutionDnsResolvers) GetServersOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Servers, true
}

// SetServers sets field value
func (o *SiteAllOfNameResolutionDnsResolvers) SetServers(v []string) {
	o.Servers = v
}

// GetSearchDomains returns the SearchDomains field value if set, zero value otherwise.
// Deprecated
func (o *SiteAllOfNameResolutionDnsResolvers) GetSearchDomains() []string {
	if o == nil || o.SearchDomains == nil {
		var ret []string
		return ret
	}
	return *o.SearchDomains
}

// GetSearchDomainsOk returns a tuple with the SearchDomains field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *SiteAllOfNameResolutionDnsResolvers) GetSearchDomainsOk() (*[]string, bool) {
	if o == nil || o.SearchDomains == nil {
		return nil, false
	}
	return o.SearchDomains, true
}

// HasSearchDomains returns a boolean if a field has been set.
func (o *SiteAllOfNameResolutionDnsResolvers) HasSearchDomains() bool {
	if o != nil && o.SearchDomains != nil {
		return true
	}

	return false
}

// SetSearchDomains gets a reference to the given []string and assigns it to the SearchDomains field.
// Deprecated
func (o *SiteAllOfNameResolutionDnsResolvers) SetSearchDomains(v []string) {
	o.SearchDomains = &v
}

// GetMatchDomains returns the MatchDomains field value if set, zero value otherwise.
func (o *SiteAllOfNameResolutionDnsResolvers) GetMatchDomains() []string {
	if o == nil || o.MatchDomains == nil {
		var ret []string
		return ret
	}
	return *o.MatchDomains
}

// GetMatchDomainsOk returns a tuple with the MatchDomains field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteAllOfNameResolutionDnsResolvers) GetMatchDomainsOk() (*[]string, bool) {
	if o == nil || o.MatchDomains == nil {
		return nil, false
	}
	return o.MatchDomains, true
}

// HasMatchDomains returns a boolean if a field has been set.
func (o *SiteAllOfNameResolutionDnsResolvers) HasMatchDomains() bool {
	if o != nil && o.MatchDomains != nil {
		return true
	}

	return false
}

// SetMatchDomains gets a reference to the given []string and assigns it to the MatchDomains field.
func (o *SiteAllOfNameResolutionDnsResolvers) SetMatchDomains(v []string) {
	o.MatchDomains = &v
}

func (o SiteAllOfNameResolutionDnsResolvers) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.UpdateInterval != nil {
		toSerialize["updateInterval"] = o.UpdateInterval
	}
	if true {
		toSerialize["servers"] = o.Servers
	}
	if o.SearchDomains != nil {
		toSerialize["searchDomains"] = o.SearchDomains
	}
	if o.MatchDomains != nil {
		toSerialize["matchDomains"] = o.MatchDomains
	}
	return json.Marshal(toSerialize)
}

type NullableSiteAllOfNameResolutionDnsResolvers struct {
	value *SiteAllOfNameResolutionDnsResolvers
	isSet bool
}

func (v NullableSiteAllOfNameResolutionDnsResolvers) Get() *SiteAllOfNameResolutionDnsResolvers {
	return v.value
}

func (v *NullableSiteAllOfNameResolutionDnsResolvers) Set(val *SiteAllOfNameResolutionDnsResolvers) {
	v.value = val
	v.isSet = true
}

func (v NullableSiteAllOfNameResolutionDnsResolvers) IsSet() bool {
	return v.isSet
}

func (v *NullableSiteAllOfNameResolutionDnsResolvers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSiteAllOfNameResolutionDnsResolvers(val *SiteAllOfNameResolutionDnsResolvers) *NullableSiteAllOfNameResolutionDnsResolvers {
	return &NullableSiteAllOfNameResolutionDnsResolvers{value: val, isSet: true}
}

func (v NullableSiteAllOfNameResolutionDnsResolvers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSiteAllOfNameResolutionDnsResolvers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
