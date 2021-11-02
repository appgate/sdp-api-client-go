/*
Appgate SDP Controller REST API

# About   This specification documents the REST API calls for the Appgate SDP Controller.    Please refer to the REST API chapter in the manual or contact Appgate support with any questions about   this functionality. # Getting Started   Requirements for API scripting:   - Access to the Admin/API TLS Connection (default port 8443) of a Controller appliance.     (https://sdphelp.appgate.com/adminguide/appliance-function-configure.html?anchor=admin-api)   - An API user with relevant permissions.     (https://sdphelp.appgate.com/adminguide/administrative-roles-configure.html)   - In order to use the simple login API, Admin MFA must be disabled or the API user must be excluded.     (https://sdphelp.appgate.com/adminguide/mfa-for-admins.html) # Base path   HTTPS requests must be sent to the Admin Interface hostname and port, with **_/admin** path.    For example: **https://appgate.company.com:8443/admin**    All requests must have the **Accept** header as:    **application/vnd.appgate.peer-v16+json** # API Conventions   API conventions are  important to understand and follow strictly.    - While updating objects (via PUT), entire object must be sent with all fields.     - For example, in order to add a remedy method to the condition below:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": []       }       ```     - send the entire object with updated and non-updated fields:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": [{\"type\": \"DisplayMessage\", \"message\": \"test message\"}]       }       ```    - In case Controller returns an error (non-2xx HTTP status code), response body is JSON.     The \"message\" field contains information about the error.     HTTP 422 \"Unprocessable Entity\" has extra `errors` field to list all the issues with specific fields.    - Empty string (\"\") is considered a different value than \"null\" or field being omitted from JSON.     Omitting the field is recommend if no value is intended.     Empty string (\"\") will be almost always rejected as invalid value.    - There are common pattern between many objects:     - **Configuration Objects**: There are many objects with common fields, namely \"id\", \"name\", \"notes\", \"created\"       and \"updated\". These entities are listed, queried, created, updated and deleted in a similar fashion.     - **Distinguished Name**: Users and Devices are identified with what is called Distinguished Names, as used in        LDAP. The distinguished format that identifies a device and a user combination is        \"CN=\\<Device ID\\>,CN=\\<username\\>,OU=\\<Identity Provider Name\\>\". Some objects have the        \"userDistinguishedName\" field, which does not include the CN for Device ID.        This identifies a user on every device.

API version: API version 16.1
Contact: appgatesdp.support@appgate.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// ApplianceAllOfAdminInterface The details of the admin connection interface. Required on Controllers and LogServers.
type ApplianceAllOfAdminInterface struct {
	// Hostname to connect to the admin interface. This hostname will be used to validate the appliance certificate.
	Hostname string `json:"hostname"`
	// Port to connect for admin services.
	HttpsPort *int32 `json:"httpsPort,omitempty"`
	// The type of TLS ciphers to allow. See: https://www.openssl.org/docs/man1.0.2/apps/ciphers.html for all supported ciphers.
	HttpsCiphers []string `json:"httpsCiphers"`
	// Source configuration to allow via iptables.
	AllowSources *[]map[string]interface{} `json:"allowSources,omitempty"`
	HttpsP12     *P12                      `json:"httpsP12,omitempty"`
}

// NewApplianceAllOfAdminInterface instantiates a new ApplianceAllOfAdminInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplianceAllOfAdminInterface(hostname string, httpsCiphers []string) *ApplianceAllOfAdminInterface {
	this := ApplianceAllOfAdminInterface{}
	this.Hostname = hostname
	var httpsPort int32 = 8443
	this.HttpsPort = &httpsPort
	this.HttpsCiphers = httpsCiphers
	return &this
}

// NewApplianceAllOfAdminInterfaceWithDefaults instantiates a new ApplianceAllOfAdminInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplianceAllOfAdminInterfaceWithDefaults() *ApplianceAllOfAdminInterface {
	this := ApplianceAllOfAdminInterface{}
	var httpsPort int32 = 8443
	this.HttpsPort = &httpsPort
	return &this
}

// GetHostname returns the Hostname field value
func (o *ApplianceAllOfAdminInterface) GetHostname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value
// and a boolean to check if the value has been set.
func (o *ApplianceAllOfAdminInterface) GetHostnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hostname, true
}

// SetHostname sets field value
func (o *ApplianceAllOfAdminInterface) SetHostname(v string) {
	o.Hostname = v
}

// GetHttpsPort returns the HttpsPort field value if set, zero value otherwise.
func (o *ApplianceAllOfAdminInterface) GetHttpsPort() int32 {
	if o == nil || o.HttpsPort == nil {
		var ret int32
		return ret
	}
	return *o.HttpsPort
}

// GetHttpsPortOk returns a tuple with the HttpsPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceAllOfAdminInterface) GetHttpsPortOk() (*int32, bool) {
	if o == nil || o.HttpsPort == nil {
		return nil, false
	}
	return o.HttpsPort, true
}

// HasHttpsPort returns a boolean if a field has been set.
func (o *ApplianceAllOfAdminInterface) HasHttpsPort() bool {
	if o != nil && o.HttpsPort != nil {
		return true
	}

	return false
}

// SetHttpsPort gets a reference to the given int32 and assigns it to the HttpsPort field.
func (o *ApplianceAllOfAdminInterface) SetHttpsPort(v int32) {
	o.HttpsPort = &v
}

// GetHttpsCiphers returns the HttpsCiphers field value
func (o *ApplianceAllOfAdminInterface) GetHttpsCiphers() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.HttpsCiphers
}

// GetHttpsCiphersOk returns a tuple with the HttpsCiphers field value
// and a boolean to check if the value has been set.
func (o *ApplianceAllOfAdminInterface) GetHttpsCiphersOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HttpsCiphers, true
}

// SetHttpsCiphers sets field value
func (o *ApplianceAllOfAdminInterface) SetHttpsCiphers(v []string) {
	o.HttpsCiphers = v
}

// GetAllowSources returns the AllowSources field value if set, zero value otherwise.
func (o *ApplianceAllOfAdminInterface) GetAllowSources() []map[string]interface{} {
	if o == nil || o.AllowSources == nil {
		var ret []map[string]interface{}
		return ret
	}
	return *o.AllowSources
}

// GetAllowSourcesOk returns a tuple with the AllowSources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceAllOfAdminInterface) GetAllowSourcesOk() (*[]map[string]interface{}, bool) {
	if o == nil || o.AllowSources == nil {
		return nil, false
	}
	return o.AllowSources, true
}

// HasAllowSources returns a boolean if a field has been set.
func (o *ApplianceAllOfAdminInterface) HasAllowSources() bool {
	if o != nil && o.AllowSources != nil {
		return true
	}

	return false
}

// SetAllowSources gets a reference to the given []map[string]interface{} and assigns it to the AllowSources field.
func (o *ApplianceAllOfAdminInterface) SetAllowSources(v []map[string]interface{}) {
	o.AllowSources = &v
}

// GetHttpsP12 returns the HttpsP12 field value if set, zero value otherwise.
func (o *ApplianceAllOfAdminInterface) GetHttpsP12() P12 {
	if o == nil || o.HttpsP12 == nil {
		var ret P12
		return ret
	}
	return *o.HttpsP12
}

// GetHttpsP12Ok returns a tuple with the HttpsP12 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceAllOfAdminInterface) GetHttpsP12Ok() (*P12, bool) {
	if o == nil || o.HttpsP12 == nil {
		return nil, false
	}
	return o.HttpsP12, true
}

// HasHttpsP12 returns a boolean if a field has been set.
func (o *ApplianceAllOfAdminInterface) HasHttpsP12() bool {
	if o != nil && o.HttpsP12 != nil {
		return true
	}

	return false
}

// SetHttpsP12 gets a reference to the given P12 and assigns it to the HttpsP12 field.
func (o *ApplianceAllOfAdminInterface) SetHttpsP12(v P12) {
	o.HttpsP12 = &v
}

func (o ApplianceAllOfAdminInterface) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["hostname"] = o.Hostname
	}
	if o.HttpsPort != nil {
		toSerialize["httpsPort"] = o.HttpsPort
	}
	if true {
		toSerialize["httpsCiphers"] = o.HttpsCiphers
	}
	if o.AllowSources != nil {
		toSerialize["allowSources"] = o.AllowSources
	}
	if o.HttpsP12 != nil {
		toSerialize["httpsP12"] = o.HttpsP12
	}
	return json.Marshal(toSerialize)
}

type NullableApplianceAllOfAdminInterface struct {
	value *ApplianceAllOfAdminInterface
	isSet bool
}

func (v NullableApplianceAllOfAdminInterface) Get() *ApplianceAllOfAdminInterface {
	return v.value
}

func (v *NullableApplianceAllOfAdminInterface) Set(val *ApplianceAllOfAdminInterface) {
	v.value = val
	v.isSet = true
}

func (v NullableApplianceAllOfAdminInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableApplianceAllOfAdminInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplianceAllOfAdminInterface(val *ApplianceAllOfAdminInterface) *NullableApplianceAllOfAdminInterface {
	return &NullableApplianceAllOfAdminInterface{value: val, isSet: true}
}

func (v NullableApplianceAllOfAdminInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplianceAllOfAdminInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
