/*
Appgate SDP Controller REST API

# About   This specification documents the REST API calls for the Appgate SDP Controller.    Please refer to the REST API chapter in the manual or contact Appgate support with any questions about   this functionality. # Getting Started   Requirements for API scripting:   - Access to the Admin/API TLS Connection (default port 8443) of a Controller appliance.     (https://sdphelp.appgate.com/adminguide/appliance-function-configure.html?anchor=admin-api)   - An API user with relevant permissions.     (https://sdphelp.appgate.com/adminguide/administrative-roles-configure.html)   - In order to use the simple login API, Admin MFA must be disabled or the API user must be excluded.     (https://sdphelp.appgate.com/adminguide/mfa-for-admins.html) # Base path   HTTPS requests must be sent to the Admin Interface hostname and port, with **_/admin** path.    For example: **https://appgate.company.com:8443/admin**    All requests must have the **Accept** header as:    **application/vnd.appgate.peer-v22+json**    An exception is made for the **_/admin/version** endpoint which instead expects an **application/json** Accept header. # API Conventions   API conventions are  important to understand and follow strictly.    - While updating objects (via PUT), entire object must be sent with all fields.     - For example, in order to add a remedy method to the condition below:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": []       }       ```     - send the entire object with updated and non-updated fields:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": [{\"type\": \"DisplayMessage\", \"message\": \"test message\"}]       }       ```    - In case Controller returns an error (non-2xx HTTP status code), response body is JSON.     The \"message\" field contains information about the error.     HTTP 422 \"Unprocessable Entity\" has extra `errors` field to list all the issues with specific fields.    - Empty string (\"\") is considered a different value than \"null\" or field being omitted from JSON.     Omitting the field is recommended if no value is intended.     Empty string (\"\") will be almost always rejected as invalid value.    - There are common pattern between many objects:     - **Configuration Objects**: There are many objects with common fields, namely \"id\", \"name\", \"notes\", \"created\"       and \"updated\". These entities are listed, queried, created, updated and deleted in a similar fashion.     - **Distinguished Name**: Users and Devices are identified with what is called Distinguished Names, as used in        LDAP. The distinguished format that identifies a device and a user combination is        \"CN=\\<Device ID\\>,CN=\\<username\\>,OU=\\<Identity Provider Name\\>\". Some objects have the        \"userDistinguishedName\" field, which does not include the CN for Device ID.        This identifies a user on every device.

API version: API version 22.2
Contact: appgatesdp.support@appgate.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// AppStats Represents statistics for discovered apps.
type AppStats struct {
	Authentication *AppStatsEntry `json:"authentication,omitempty"`
	Access         *AppStatsEntry `json:"access,omitempty"`
	Dns            *AppStatsEntry `json:"dns,omitempty"`
}

// NewAppStats instantiates a new AppStats object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppStats() *AppStats {
	this := AppStats{}
	return &this
}

// NewAppStatsWithDefaults instantiates a new AppStats object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppStatsWithDefaults() *AppStats {
	this := AppStats{}
	return &this
}

// GetAuthentication returns the Authentication field value if set, zero value otherwise.
func (o *AppStats) GetAuthentication() AppStatsEntry {
	if o == nil || o.Authentication == nil {
		var ret AppStatsEntry
		return ret
	}
	return *o.Authentication
}

// GetAuthenticationOk returns a tuple with the Authentication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStats) GetAuthenticationOk() (*AppStatsEntry, bool) {
	if o == nil || o.Authentication == nil {
		return nil, false
	}
	return o.Authentication, true
}

// HasAuthentication returns a boolean if a field has been set.
func (o *AppStats) HasAuthentication() bool {
	if o != nil && o.Authentication != nil {
		return true
	}

	return false
}

// SetAuthentication gets a reference to the given AppStatsEntry and assigns it to the Authentication field.
func (o *AppStats) SetAuthentication(v AppStatsEntry) {
	o.Authentication = &v
}

// GetAccess returns the Access field value if set, zero value otherwise.
func (o *AppStats) GetAccess() AppStatsEntry {
	if o == nil || o.Access == nil {
		var ret AppStatsEntry
		return ret
	}
	return *o.Access
}

// GetAccessOk returns a tuple with the Access field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStats) GetAccessOk() (*AppStatsEntry, bool) {
	if o == nil || o.Access == nil {
		return nil, false
	}
	return o.Access, true
}

// HasAccess returns a boolean if a field has been set.
func (o *AppStats) HasAccess() bool {
	if o != nil && o.Access != nil {
		return true
	}

	return false
}

// SetAccess gets a reference to the given AppStatsEntry and assigns it to the Access field.
func (o *AppStats) SetAccess(v AppStatsEntry) {
	o.Access = &v
}

// GetDns returns the Dns field value if set, zero value otherwise.
func (o *AppStats) GetDns() AppStatsEntry {
	if o == nil || o.Dns == nil {
		var ret AppStatsEntry
		return ret
	}
	return *o.Dns
}

// GetDnsOk returns a tuple with the Dns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStats) GetDnsOk() (*AppStatsEntry, bool) {
	if o == nil || o.Dns == nil {
		return nil, false
	}
	return o.Dns, true
}

// HasDns returns a boolean if a field has been set.
func (o *AppStats) HasDns() bool {
	if o != nil && o.Dns != nil {
		return true
	}

	return false
}

// SetDns gets a reference to the given AppStatsEntry and assigns it to the Dns field.
func (o *AppStats) SetDns(v AppStatsEntry) {
	o.Dns = &v
}

func (o AppStats) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Authentication != nil {
		toSerialize["authentication"] = o.Authentication
	}
	if o.Access != nil {
		toSerialize["access"] = o.Access
	}
	if o.Dns != nil {
		toSerialize["dns"] = o.Dns
	}
	return json.Marshal(toSerialize)
}

type NullableAppStats struct {
	value *AppStats
	isSet bool
}

func (v NullableAppStats) Get() *AppStats {
	return v.value
}

func (v *NullableAppStats) Set(val *AppStats) {
	v.value = val
	v.isSet = true
}

func (v NullableAppStats) IsSet() bool {
	return v.isSet
}

func (v *NullableAppStats) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppStats(val *AppStats) *NullableAppStats {
	return &NullableAppStats{value: val, isSet: true}
}

func (v NullableAppStats) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppStats) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
