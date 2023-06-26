/*
Appgate SDP Controller REST API

# About   This specification documents the REST API calls for the Appgate SDP Controller.    Please refer to the REST API chapter in the manual or contact Appgate support with any questions about   this functionality. # Getting Started   Requirements for API scripting:   - Access to the Admin/API TLS Connection (default port 8443) of a Controller appliance.     (https://sdphelp.appgate.com/adminguide/appliance-function-configure.html?anchor=admin-api)   - An API user with relevant permissions.     (https://sdphelp.appgate.com/adminguide/administrative-roles-configure.html)   - In order to use the simple login API, Admin MFA must be disabled or the API user must be excluded.     (https://sdphelp.appgate.com/adminguide/mfa-for-admins.html) # Base path   HTTPS requests must be sent to the Admin Interface hostname and port, with **_/admin** path.    For example: **https://appgate.company.com:8443/admin**    All requests must have the **Accept** header as:    **application/vnd.appgate.peer-v19+json** # API Conventions   API conventions are  important to understand and follow strictly.    - While updating objects (via PUT), entire object must be sent with all fields.     - For example, in order to add a remedy method to the condition below:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": []       }       ```     - send the entire object with updated and non-updated fields:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": [{\"type\": \"DisplayMessage\", \"message\": \"test message\"}]       }       ```    - In case Controller returns an error (non-2xx HTTP status code), response body is JSON.     The \"message\" field contains information about the error.     HTTP 422 \"Unprocessable Entity\" has extra `errors` field to list all the issues with specific fields.    - Empty string (\"\") is considered a different value than \"null\" or field being omitted from JSON.     Omitting the field is recommended if no value is intended.     Empty string (\"\") will be almost always rejected as invalid value.    - There are common pattern between many objects:     - **Configuration Objects**: There are many objects with common fields, namely \"id\", \"name\", \"notes\", \"created\"       and \"updated\". These entities are listed, queried, created, updated and deleted in a similar fashion.     - **Distinguished Name**: Users and Devices are identified with what is called Distinguished Names, as used in        LDAP. The distinguished format that identifies a device and a user combination is        \"CN=\\<Device ID\\>,CN=\\<username\\>,OU=\\<Identity Provider Name\\>\". Some objects have the        \"userDistinguishedName\" field, which does not include the CN for Device ID.        This identifies a user on every device.

API version: API version 19.1
Contact: appgatesdp.support@appgate.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// ApplianceAllOfClientInterface The details of the Client connection interface.
type ApplianceAllOfClientInterface struct {
	// To enable/disable Proxy protocol on this Appliance.
	ProxyProtocol *bool `json:"proxyProtocol,omitempty"`
	// Hostname to connect by the Clients. It will be used to validate the Appliance Certificate.
	Hostname string `json:"hostname"`
	// Load hostname that would take effect with Local Site Detection feature.
	LocalHostname *string `json:"localHostname,omitempty"`
	// Port to connect for the Client specific services.
	HttpsPort *int32 `json:"httpsPort,omitempty"`
	// Port to connect for the Clients that connects to vpnd on DTLS if enabled.
	DtlsPort *int32 `json:"dtlsPort,omitempty"`
	// Source configuration to allow via iptables.
	AllowSources []AllowSourcesInner `json:"allowSources,omitempty"`
	// Override SPA mode for this appliance.
	OverrideSpaMode *string `json:"overrideSpaMode,omitempty"`
}

// NewApplianceAllOfClientInterface instantiates a new ApplianceAllOfClientInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplianceAllOfClientInterface(hostname string) *ApplianceAllOfClientInterface {
	this := ApplianceAllOfClientInterface{}
	var proxyProtocol bool = false
	this.ProxyProtocol = &proxyProtocol
	this.Hostname = hostname
	var httpsPort int32 = 443
	this.HttpsPort = &httpsPort
	var dtlsPort int32 = 443
	this.DtlsPort = &dtlsPort
	return &this
}

// NewApplianceAllOfClientInterfaceWithDefaults instantiates a new ApplianceAllOfClientInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplianceAllOfClientInterfaceWithDefaults() *ApplianceAllOfClientInterface {
	this := ApplianceAllOfClientInterface{}
	var proxyProtocol bool = false
	this.ProxyProtocol = &proxyProtocol
	var httpsPort int32 = 443
	this.HttpsPort = &httpsPort
	var dtlsPort int32 = 443
	this.DtlsPort = &dtlsPort
	return &this
}

// GetProxyProtocol returns the ProxyProtocol field value if set, zero value otherwise.
func (o *ApplianceAllOfClientInterface) GetProxyProtocol() bool {
	if o == nil || o.ProxyProtocol == nil {
		var ret bool
		return ret
	}
	return *o.ProxyProtocol
}

// GetProxyProtocolOk returns a tuple with the ProxyProtocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceAllOfClientInterface) GetProxyProtocolOk() (*bool, bool) {
	if o == nil || o.ProxyProtocol == nil {
		return nil, false
	}
	return o.ProxyProtocol, true
}

// HasProxyProtocol returns a boolean if a field has been set.
func (o *ApplianceAllOfClientInterface) HasProxyProtocol() bool {
	if o != nil && o.ProxyProtocol != nil {
		return true
	}

	return false
}

// SetProxyProtocol gets a reference to the given bool and assigns it to the ProxyProtocol field.
func (o *ApplianceAllOfClientInterface) SetProxyProtocol(v bool) {
	o.ProxyProtocol = &v
}

// GetHostname returns the Hostname field value
func (o *ApplianceAllOfClientInterface) GetHostname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value
// and a boolean to check if the value has been set.
func (o *ApplianceAllOfClientInterface) GetHostnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hostname, true
}

// SetHostname sets field value
func (o *ApplianceAllOfClientInterface) SetHostname(v string) {
	o.Hostname = v
}

// GetLocalHostname returns the LocalHostname field value if set, zero value otherwise.
func (o *ApplianceAllOfClientInterface) GetLocalHostname() string {
	if o == nil || o.LocalHostname == nil {
		var ret string
		return ret
	}
	return *o.LocalHostname
}

// GetLocalHostnameOk returns a tuple with the LocalHostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceAllOfClientInterface) GetLocalHostnameOk() (*string, bool) {
	if o == nil || o.LocalHostname == nil {
		return nil, false
	}
	return o.LocalHostname, true
}

// HasLocalHostname returns a boolean if a field has been set.
func (o *ApplianceAllOfClientInterface) HasLocalHostname() bool {
	if o != nil && o.LocalHostname != nil {
		return true
	}

	return false
}

// SetLocalHostname gets a reference to the given string and assigns it to the LocalHostname field.
func (o *ApplianceAllOfClientInterface) SetLocalHostname(v string) {
	o.LocalHostname = &v
}

// GetHttpsPort returns the HttpsPort field value if set, zero value otherwise.
func (o *ApplianceAllOfClientInterface) GetHttpsPort() int32 {
	if o == nil || o.HttpsPort == nil {
		var ret int32
		return ret
	}
	return *o.HttpsPort
}

// GetHttpsPortOk returns a tuple with the HttpsPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceAllOfClientInterface) GetHttpsPortOk() (*int32, bool) {
	if o == nil || o.HttpsPort == nil {
		return nil, false
	}
	return o.HttpsPort, true
}

// HasHttpsPort returns a boolean if a field has been set.
func (o *ApplianceAllOfClientInterface) HasHttpsPort() bool {
	if o != nil && o.HttpsPort != nil {
		return true
	}

	return false
}

// SetHttpsPort gets a reference to the given int32 and assigns it to the HttpsPort field.
func (o *ApplianceAllOfClientInterface) SetHttpsPort(v int32) {
	o.HttpsPort = &v
}

// GetDtlsPort returns the DtlsPort field value if set, zero value otherwise.
func (o *ApplianceAllOfClientInterface) GetDtlsPort() int32 {
	if o == nil || o.DtlsPort == nil {
		var ret int32
		return ret
	}
	return *o.DtlsPort
}

// GetDtlsPortOk returns a tuple with the DtlsPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceAllOfClientInterface) GetDtlsPortOk() (*int32, bool) {
	if o == nil || o.DtlsPort == nil {
		return nil, false
	}
	return o.DtlsPort, true
}

// HasDtlsPort returns a boolean if a field has been set.
func (o *ApplianceAllOfClientInterface) HasDtlsPort() bool {
	if o != nil && o.DtlsPort != nil {
		return true
	}

	return false
}

// SetDtlsPort gets a reference to the given int32 and assigns it to the DtlsPort field.
func (o *ApplianceAllOfClientInterface) SetDtlsPort(v int32) {
	o.DtlsPort = &v
}

// GetAllowSources returns the AllowSources field value if set, zero value otherwise.
func (o *ApplianceAllOfClientInterface) GetAllowSources() []AllowSourcesInner {
	if o == nil || o.AllowSources == nil {
		var ret []AllowSourcesInner
		return ret
	}
	return o.AllowSources
}

// GetAllowSourcesOk returns a tuple with the AllowSources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceAllOfClientInterface) GetAllowSourcesOk() ([]AllowSourcesInner, bool) {
	if o == nil || o.AllowSources == nil {
		return nil, false
	}
	return o.AllowSources, true
}

// HasAllowSources returns a boolean if a field has been set.
func (o *ApplianceAllOfClientInterface) HasAllowSources() bool {
	if o != nil && o.AllowSources != nil {
		return true
	}

	return false
}

// SetAllowSources gets a reference to the given []AllowSourcesInner and assigns it to the AllowSources field.
func (o *ApplianceAllOfClientInterface) SetAllowSources(v []AllowSourcesInner) {
	o.AllowSources = v
}

// GetOverrideSpaMode returns the OverrideSpaMode field value if set, zero value otherwise.
func (o *ApplianceAllOfClientInterface) GetOverrideSpaMode() string {
	if o == nil || o.OverrideSpaMode == nil {
		var ret string
		return ret
	}
	return *o.OverrideSpaMode
}

// GetOverrideSpaModeOk returns a tuple with the OverrideSpaMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceAllOfClientInterface) GetOverrideSpaModeOk() (*string, bool) {
	if o == nil || o.OverrideSpaMode == nil {
		return nil, false
	}
	return o.OverrideSpaMode, true
}

// HasOverrideSpaMode returns a boolean if a field has been set.
func (o *ApplianceAllOfClientInterface) HasOverrideSpaMode() bool {
	if o != nil && o.OverrideSpaMode != nil {
		return true
	}

	return false
}

// SetOverrideSpaMode gets a reference to the given string and assigns it to the OverrideSpaMode field.
func (o *ApplianceAllOfClientInterface) SetOverrideSpaMode(v string) {
	o.OverrideSpaMode = &v
}

func (o ApplianceAllOfClientInterface) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ProxyProtocol != nil {
		toSerialize["proxyProtocol"] = o.ProxyProtocol
	}
	if true {
		toSerialize["hostname"] = o.Hostname
	}
	if o.LocalHostname != nil {
		toSerialize["localHostname"] = o.LocalHostname
	}
	if o.HttpsPort != nil {
		toSerialize["httpsPort"] = o.HttpsPort
	}
	if o.DtlsPort != nil {
		toSerialize["dtlsPort"] = o.DtlsPort
	}
	if o.AllowSources != nil {
		toSerialize["allowSources"] = o.AllowSources
	}
	if o.OverrideSpaMode != nil {
		toSerialize["overrideSpaMode"] = o.OverrideSpaMode
	}
	return json.Marshal(toSerialize)
}

type NullableApplianceAllOfClientInterface struct {
	value *ApplianceAllOfClientInterface
	isSet bool
}

func (v NullableApplianceAllOfClientInterface) Get() *ApplianceAllOfClientInterface {
	return v.value
}

func (v *NullableApplianceAllOfClientInterface) Set(val *ApplianceAllOfClientInterface) {
	v.value = val
	v.isSet = true
}

func (v NullableApplianceAllOfClientInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableApplianceAllOfClientInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplianceAllOfClientInterface(val *ApplianceAllOfClientInterface) *NullableApplianceAllOfClientInterface {
	return &NullableApplianceAllOfClientInterface{value: val, isSet: true}
}

func (v NullableApplianceAllOfClientInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplianceAllOfClientInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
