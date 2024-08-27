/*
Appgate SDP Controller REST API

# About   This specification documents the REST API calls for the Appgate SDP Controller.    Please refer to the REST API chapter in the manual or contact Appgate support with any questions about   this functionality. # Getting Started   Requirements for API scripting:   - Access to the Admin/API TLS Connection (default port 8443) of a Controller appliance.     (https://sdphelp.appgate.com/adminguide/appliance-function-configure.html?anchor=admin-api)   - An API user with relevant permissions.     (https://sdphelp.appgate.com/adminguide/administrative-roles-configure.html)   - In order to use the simple login API, Admin MFA must be disabled or the API user must be excluded.     (https://sdphelp.appgate.com/adminguide/mfa-for-admins.html) # Base path   HTTPS requests must be sent to the Admin Interface hostname and port, with **_/admin** path.    For example: **https://appgate.company.com:8443/admin**    All requests must have the **Accept** header as:    **application/vnd.appgate.peer-v20+json**    An exception is made for the **_/admin/version** endpoint which instead expects an **application/json** Accept header. # API Conventions   API conventions are  important to understand and follow strictly.    - While updating objects (via PUT), entire object must be sent with all fields.     - For example, in order to add a remedy method to the condition below:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": []       }       ```     - send the entire object with updated and non-updated fields:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": [{\"type\": \"DisplayMessage\", \"message\": \"test message\"}]       }       ```    - In case Controller returns an error (non-2xx HTTP status code), response body is JSON.     The \"message\" field contains information about the error.     HTTP 422 \"Unprocessable Entity\" has extra `errors` field to list all the issues with specific fields.    - Empty string (\"\") is considered a different value than \"null\" or field being omitted from JSON.     Omitting the field is recommended if no value is intended.     Empty string (\"\") will be almost always rejected as invalid value.    - There are common pattern between many objects:     - **Configuration Objects**: There are many objects with common fields, namely \"id\", \"name\", \"notes\", \"created\"       and \"updated\". These entities are listed, queried, created, updated and deleted in a similar fashion.     - **Distinguished Name**: Users and Devices are identified with what is called Distinguished Names, as used in        LDAP. The distinguished format that identifies a device and a user combination is        \"CN=\\<Device ID\\>,CN=\\<username\\>,OU=\\<Identity Provider Name\\>\". Some objects have the        \"userDistinguishedName\" field, which does not include the CN for Device ID.        This identifies a user on every device.

API version: API version 20.3
Contact: appgatesdp.support@appgate.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// DnsSettings struct for DnsSettings
type DnsSettings struct {
	// The DNS servers to be assigned to the Clients of the users in this Identity Provider.
	DnsServers []string `json:"dnsServers,omitempty"`
	// The DNS search domains to be assigned to Clients of the users in this Identity Provider.
	DnsSearchDomains []string `json:"dnsSearchDomains,omitempty"`
}

// NewDnsSettings instantiates a new DnsSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDnsSettings() *DnsSettings {
	this := DnsSettings{}
	return &this
}

// NewDnsSettingsWithDefaults instantiates a new DnsSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDnsSettingsWithDefaults() *DnsSettings {
	this := DnsSettings{}
	return &this
}

// GetDnsServers returns the DnsServers field value if set, zero value otherwise.
func (o *DnsSettings) GetDnsServers() []string {
	if o == nil || o.DnsServers == nil {
		var ret []string
		return ret
	}
	return o.DnsServers
}

// GetDnsServersOk returns a tuple with the DnsServers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsSettings) GetDnsServersOk() ([]string, bool) {
	if o == nil || o.DnsServers == nil {
		return nil, false
	}
	return o.DnsServers, true
}

// HasDnsServers returns a boolean if a field has been set.
func (o *DnsSettings) HasDnsServers() bool {
	if o != nil && o.DnsServers != nil {
		return true
	}

	return false
}

// SetDnsServers gets a reference to the given []string and assigns it to the DnsServers field.
func (o *DnsSettings) SetDnsServers(v []string) {
	o.DnsServers = v
}

// GetDnsSearchDomains returns the DnsSearchDomains field value if set, zero value otherwise.
func (o *DnsSettings) GetDnsSearchDomains() []string {
	if o == nil || o.DnsSearchDomains == nil {
		var ret []string
		return ret
	}
	return o.DnsSearchDomains
}

// GetDnsSearchDomainsOk returns a tuple with the DnsSearchDomains field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsSettings) GetDnsSearchDomainsOk() ([]string, bool) {
	if o == nil || o.DnsSearchDomains == nil {
		return nil, false
	}
	return o.DnsSearchDomains, true
}

// HasDnsSearchDomains returns a boolean if a field has been set.
func (o *DnsSettings) HasDnsSearchDomains() bool {
	if o != nil && o.DnsSearchDomains != nil {
		return true
	}

	return false
}

// SetDnsSearchDomains gets a reference to the given []string and assigns it to the DnsSearchDomains field.
func (o *DnsSettings) SetDnsSearchDomains(v []string) {
	o.DnsSearchDomains = v
}

func (o DnsSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DnsServers != nil {
		toSerialize["dnsServers"] = o.DnsServers
	}
	if o.DnsSearchDomains != nil {
		toSerialize["dnsSearchDomains"] = o.DnsSearchDomains
	}
	return json.Marshal(toSerialize)
}

type NullableDnsSettings struct {
	value *DnsSettings
	isSet bool
}

func (v NullableDnsSettings) Get() *DnsSettings {
	return v.value
}

func (v *NullableDnsSettings) Set(val *DnsSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableDnsSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableDnsSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDnsSettings(val *DnsSettings) *NullableDnsSettings {
	return &NullableDnsSettings{value: val, isSet: true}
}

func (v NullableDnsSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDnsSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
