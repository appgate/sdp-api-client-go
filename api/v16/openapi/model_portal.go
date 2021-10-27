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

// Portal Portal settings.
type Portal struct {
	// Whether the Portal is enabled on this appliance or not.
	Enabled  *bool `json:"enabled,omitempty"`
	HttpsP12 *P12  `json:"httpsP12,omitempty"`
	// Automatic 80->443 redirection for Portal.
	HttpRedirect *bool `json:"httpRedirect,omitempty"`
	// Ports that can be proxied via Portal.
	ProxyPorts *[]int32 `json:"proxyPorts,omitempty"`
	// P12 files for proxying traffic to HTTPS endpoints.
	ProxyP12s *[]Portal12 `json:"proxyP12s,omitempty"`
	// Names of the profiles in this Collective to use in the Portal.
	Profiles *[]string `json:"profiles,omitempty"`
	// Profiles from other Collectives to use in the Portal.
	ExternalProfiles    *[]PortalExternalProfiles  `json:"externalProfiles,omitempty"`
	SignInCustomization *PortalSignInCustomization `json:"signInCustomization,omitempty"`
}

// NewPortal instantiates a new Portal object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPortal() *Portal {
	this := Portal{}
	var enabled bool = false
	this.Enabled = &enabled
	var httpRedirect bool = true
	this.HttpRedirect = &httpRedirect
	return &this
}

// NewPortalWithDefaults instantiates a new Portal object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPortalWithDefaults() *Portal {
	this := Portal{}
	var enabled bool = false
	this.Enabled = &enabled
	var httpRedirect bool = true
	this.HttpRedirect = &httpRedirect
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *Portal) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Portal) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *Portal) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *Portal) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetHttpsP12 returns the HttpsP12 field value if set, zero value otherwise.
func (o *Portal) GetHttpsP12() P12 {
	if o == nil || o.HttpsP12 == nil {
		var ret P12
		return ret
	}
	return *o.HttpsP12
}

// GetHttpsP12Ok returns a tuple with the HttpsP12 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Portal) GetHttpsP12Ok() (*P12, bool) {
	if o == nil || o.HttpsP12 == nil {
		return nil, false
	}
	return o.HttpsP12, true
}

// HasHttpsP12 returns a boolean if a field has been set.
func (o *Portal) HasHttpsP12() bool {
	if o != nil && o.HttpsP12 != nil {
		return true
	}

	return false
}

// SetHttpsP12 gets a reference to the given P12 and assigns it to the HttpsP12 field.
func (o *Portal) SetHttpsP12(v P12) {
	o.HttpsP12 = &v
}

// GetHttpRedirect returns the HttpRedirect field value if set, zero value otherwise.
func (o *Portal) GetHttpRedirect() bool {
	if o == nil || o.HttpRedirect == nil {
		var ret bool
		return ret
	}
	return *o.HttpRedirect
}

// GetHttpRedirectOk returns a tuple with the HttpRedirect field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Portal) GetHttpRedirectOk() (*bool, bool) {
	if o == nil || o.HttpRedirect == nil {
		return nil, false
	}
	return o.HttpRedirect, true
}

// HasHttpRedirect returns a boolean if a field has been set.
func (o *Portal) HasHttpRedirect() bool {
	if o != nil && o.HttpRedirect != nil {
		return true
	}

	return false
}

// SetHttpRedirect gets a reference to the given bool and assigns it to the HttpRedirect field.
func (o *Portal) SetHttpRedirect(v bool) {
	o.HttpRedirect = &v
}

// GetProxyPorts returns the ProxyPorts field value if set, zero value otherwise.
func (o *Portal) GetProxyPorts() []int32 {
	if o == nil || o.ProxyPorts == nil {
		var ret []int32
		return ret
	}
	return *o.ProxyPorts
}

// GetProxyPortsOk returns a tuple with the ProxyPorts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Portal) GetProxyPortsOk() (*[]int32, bool) {
	if o == nil || o.ProxyPorts == nil {
		return nil, false
	}
	return o.ProxyPorts, true
}

// HasProxyPorts returns a boolean if a field has been set.
func (o *Portal) HasProxyPorts() bool {
	if o != nil && o.ProxyPorts != nil {
		return true
	}

	return false
}

// SetProxyPorts gets a reference to the given []int32 and assigns it to the ProxyPorts field.
func (o *Portal) SetProxyPorts(v []int32) {
	o.ProxyPorts = &v
}

// GetProxyP12s returns the ProxyP12s field value if set, zero value otherwise.
func (o *Portal) GetProxyP12s() []Portal12 {
	if o == nil || o.ProxyP12s == nil {
		var ret []Portal12
		return ret
	}
	return *o.ProxyP12s
}

// GetProxyP12sOk returns a tuple with the ProxyP12s field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Portal) GetProxyP12sOk() (*[]Portal12, bool) {
	if o == nil || o.ProxyP12s == nil {
		return nil, false
	}
	return o.ProxyP12s, true
}

// HasProxyP12s returns a boolean if a field has been set.
func (o *Portal) HasProxyP12s() bool {
	if o != nil && o.ProxyP12s != nil {
		return true
	}

	return false
}

// SetProxyP12s gets a reference to the given []Portal12 and assigns it to the ProxyP12s field.
func (o *Portal) SetProxyP12s(v []Portal12) {
	o.ProxyP12s = &v
}

// GetProfiles returns the Profiles field value if set, zero value otherwise.
func (o *Portal) GetProfiles() []string {
	if o == nil || o.Profiles == nil {
		var ret []string
		return ret
	}
	return *o.Profiles
}

// GetProfilesOk returns a tuple with the Profiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Portal) GetProfilesOk() (*[]string, bool) {
	if o == nil || o.Profiles == nil {
		return nil, false
	}
	return o.Profiles, true
}

// HasProfiles returns a boolean if a field has been set.
func (o *Portal) HasProfiles() bool {
	if o != nil && o.Profiles != nil {
		return true
	}

	return false
}

// SetProfiles gets a reference to the given []string and assigns it to the Profiles field.
func (o *Portal) SetProfiles(v []string) {
	o.Profiles = &v
}

// GetExternalProfiles returns the ExternalProfiles field value if set, zero value otherwise.
func (o *Portal) GetExternalProfiles() []PortalExternalProfiles {
	if o == nil || o.ExternalProfiles == nil {
		var ret []PortalExternalProfiles
		return ret
	}
	return *o.ExternalProfiles
}

// GetExternalProfilesOk returns a tuple with the ExternalProfiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Portal) GetExternalProfilesOk() (*[]PortalExternalProfiles, bool) {
	if o == nil || o.ExternalProfiles == nil {
		return nil, false
	}
	return o.ExternalProfiles, true
}

// HasExternalProfiles returns a boolean if a field has been set.
func (o *Portal) HasExternalProfiles() bool {
	if o != nil && o.ExternalProfiles != nil {
		return true
	}

	return false
}

// SetExternalProfiles gets a reference to the given []PortalExternalProfiles and assigns it to the ExternalProfiles field.
func (o *Portal) SetExternalProfiles(v []PortalExternalProfiles) {
	o.ExternalProfiles = &v
}

// GetSignInCustomization returns the SignInCustomization field value if set, zero value otherwise.
func (o *Portal) GetSignInCustomization() PortalSignInCustomization {
	if o == nil || o.SignInCustomization == nil {
		var ret PortalSignInCustomization
		return ret
	}
	return *o.SignInCustomization
}

// GetSignInCustomizationOk returns a tuple with the SignInCustomization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Portal) GetSignInCustomizationOk() (*PortalSignInCustomization, bool) {
	if o == nil || o.SignInCustomization == nil {
		return nil, false
	}
	return o.SignInCustomization, true
}

// HasSignInCustomization returns a boolean if a field has been set.
func (o *Portal) HasSignInCustomization() bool {
	if o != nil && o.SignInCustomization != nil {
		return true
	}

	return false
}

// SetSignInCustomization gets a reference to the given PortalSignInCustomization and assigns it to the SignInCustomization field.
func (o *Portal) SetSignInCustomization(v PortalSignInCustomization) {
	o.SignInCustomization = &v
}

func (o Portal) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.HttpsP12 != nil {
		toSerialize["httpsP12"] = o.HttpsP12
	}
	if o.HttpRedirect != nil {
		toSerialize["httpRedirect"] = o.HttpRedirect
	}
	if o.ProxyPorts != nil {
		toSerialize["proxyPorts"] = o.ProxyPorts
	}
	if o.ProxyP12s != nil {
		toSerialize["proxyP12s"] = o.ProxyP12s
	}
	if o.Profiles != nil {
		toSerialize["profiles"] = o.Profiles
	}
	if o.ExternalProfiles != nil {
		toSerialize["externalProfiles"] = o.ExternalProfiles
	}
	if o.SignInCustomization != nil {
		toSerialize["signInCustomization"] = o.SignInCustomization
	}
	return json.Marshal(toSerialize)
}

type NullablePortal struct {
	value *Portal
	isSet bool
}

func (v NullablePortal) Get() *Portal {
	return v.value
}

func (v *NullablePortal) Set(val *Portal) {
	v.value = val
	v.isSet = true
}

func (v NullablePortal) IsSet() bool {
	return v.isSet
}

func (v *NullablePortal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePortal(val *Portal) *NullablePortal {
	return &NullablePortal{value: val, isSet: true}
}

func (v NullablePortal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePortal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
