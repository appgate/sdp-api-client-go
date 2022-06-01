/*
Appgate SDP Controller REST API

# About   This specification documents the REST API calls for the Appgate SDP Controller.    Please refer to the REST API chapter in the manual or contact Appgate support with any questions about   this functionality. # Getting Started   Requirements for API scripting:   - Access to the Admin/API TLS Connection (default port 8443) of a Controller appliance.     (https://sdphelp.appgate.com/adminguide/appliance-function-configure.html?anchor=admin-api)   - An API user with relevant permissions.     (https://sdphelp.appgate.com/adminguide/administrative-roles-configure.html)   - In order to use the simple login API, Admin MFA must be disabled or the API user must be excluded.     (https://sdphelp.appgate.com/adminguide/mfa-for-admins.html) # Base path   HTTPS requests must be sent to the Admin Interface hostname and port, with **_/admin** path.    For example: **https://appgate.company.com:8443/admin**    All requests must have the **Accept** header as:    **application/vnd.appgate.peer-v17+json** # API Conventions   API conventions are  important to understand and follow strictly.    - While updating objects (via PUT), entire object must be sent with all fields.     - For example, in order to add a remedy method to the condition below:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": []       }       ```     - send the entire object with updated and non-updated fields:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": [{\"type\": \"DisplayMessage\", \"message\": \"test message\"}]       }       ```    - In case Controller returns an error (non-2xx HTTP status code), response body is JSON.     The \"message\" field contains information about the error.     HTTP 422 \"Unprocessable Entity\" has extra `errors` field to list all the issues with specific fields.    - Empty string (\"\") is considered a different value than \"null\" or field being omitted from JSON.     Omitting the field is recommended if no value is intended.     Empty string (\"\") will be almost always rejected as invalid value.    - There are common pattern between many objects:     - **Configuration Objects**: There are many objects with common fields, namely \"id\", \"name\", \"notes\", \"created\"       and \"updated\". These entities are listed, queried, created, updated and deleted in a similar fashion.     - **Distinguished Name**: Users and Devices are identified with what is called Distinguished Names, as used in        LDAP. The distinguished format that identifies a device and a user combination is        \"CN=\\<Device ID\\>,CN=\\<username\\>,OU=\\<Identity Provider Name\\>\". Some objects have the        \"userDistinguishedName\" field, which does not include the CN for Device ID.        This identifies a user on every device.

API version: API version 17.1
Contact: appgatesdp.support@appgate.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// SiteAllOfVpnRouteVia Override routing for tunnel traffic.
type SiteAllOfVpnRouteVia struct {
	// IPv4 address for routing tunnel traffic.
	Ipv4 *string `json:"ipv4,omitempty"`
	// IPv6 address for routing tunnel traffic.
	Ipv6 *string `json:"ipv6,omitempty"`
}

// NewSiteAllOfVpnRouteVia instantiates a new SiteAllOfVpnRouteVia object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSiteAllOfVpnRouteVia() *SiteAllOfVpnRouteVia {
	this := SiteAllOfVpnRouteVia{}
	return &this
}

// NewSiteAllOfVpnRouteViaWithDefaults instantiates a new SiteAllOfVpnRouteVia object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSiteAllOfVpnRouteViaWithDefaults() *SiteAllOfVpnRouteVia {
	this := SiteAllOfVpnRouteVia{}
	return &this
}

// GetIpv4 returns the Ipv4 field value if set, zero value otherwise.
func (o *SiteAllOfVpnRouteVia) GetIpv4() string {
	if o == nil || o.Ipv4 == nil {
		var ret string
		return ret
	}
	return *o.Ipv4
}

// GetIpv4Ok returns a tuple with the Ipv4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteAllOfVpnRouteVia) GetIpv4Ok() (*string, bool) {
	if o == nil || o.Ipv4 == nil {
		return nil, false
	}
	return o.Ipv4, true
}

// HasIpv4 returns a boolean if a field has been set.
func (o *SiteAllOfVpnRouteVia) HasIpv4() bool {
	if o != nil && o.Ipv4 != nil {
		return true
	}

	return false
}

// SetIpv4 gets a reference to the given string and assigns it to the Ipv4 field.
func (o *SiteAllOfVpnRouteVia) SetIpv4(v string) {
	o.Ipv4 = &v
}

// GetIpv6 returns the Ipv6 field value if set, zero value otherwise.
func (o *SiteAllOfVpnRouteVia) GetIpv6() string {
	if o == nil || o.Ipv6 == nil {
		var ret string
		return ret
	}
	return *o.Ipv6
}

// GetIpv6Ok returns a tuple with the Ipv6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteAllOfVpnRouteVia) GetIpv6Ok() (*string, bool) {
	if o == nil || o.Ipv6 == nil {
		return nil, false
	}
	return o.Ipv6, true
}

// HasIpv6 returns a boolean if a field has been set.
func (o *SiteAllOfVpnRouteVia) HasIpv6() bool {
	if o != nil && o.Ipv6 != nil {
		return true
	}

	return false
}

// SetIpv6 gets a reference to the given string and assigns it to the Ipv6 field.
func (o *SiteAllOfVpnRouteVia) SetIpv6(v string) {
	o.Ipv6 = &v
}

func (o SiteAllOfVpnRouteVia) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Ipv4 != nil {
		toSerialize["ipv4"] = o.Ipv4
	}
	if o.Ipv6 != nil {
		toSerialize["ipv6"] = o.Ipv6
	}
	return json.Marshal(toSerialize)
}

type NullableSiteAllOfVpnRouteVia struct {
	value *SiteAllOfVpnRouteVia
	isSet bool
}

func (v NullableSiteAllOfVpnRouteVia) Get() *SiteAllOfVpnRouteVia {
	return v.value
}

func (v *NullableSiteAllOfVpnRouteVia) Set(val *SiteAllOfVpnRouteVia) {
	v.value = val
	v.isSet = true
}

func (v NullableSiteAllOfVpnRouteVia) IsSet() bool {
	return v.isSet
}

func (v *NullableSiteAllOfVpnRouteVia) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSiteAllOfVpnRouteVia(val *SiteAllOfVpnRouteVia) *NullableSiteAllOfVpnRouteVia {
	return &NullableSiteAllOfVpnRouteVia{value: val, isSet: true}
}

func (v NullableSiteAllOfVpnRouteVia) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSiteAllOfVpnRouteVia) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
