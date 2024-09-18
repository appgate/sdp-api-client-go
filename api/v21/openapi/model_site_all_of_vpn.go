/*
Appgate SDP Controller REST API

# About   This specification documents the REST API calls for the Appgate SDP Controller.    Please refer to the REST API chapter in the manual or contact Appgate support with any questions about   this functionality. # Getting Started   Requirements for API scripting:   - Access to the Admin/API TLS Connection (default port 8443) of a Controller appliance.     (https://sdphelp.appgate.com/adminguide/appliance-function-configure.html?anchor=admin-api)   - An API user with relevant permissions.     (https://sdphelp.appgate.com/adminguide/administrative-roles-configure.html)   - In order to use the simple login API, Admin MFA must be disabled or the API user must be excluded.     (https://sdphelp.appgate.com/adminguide/mfa-for-admins.html) # Base path   HTTPS requests must be sent to the Admin Interface hostname and port, with **_/admin** path.    For example: **https://appgate.company.com:8443/admin**    All requests must have the **Accept** header as:    **application/vnd.appgate.peer-v21+json**    An exception is made for the **_/admin/version** endpoint which instead expects an **application/json** Accept header. # API Conventions   API conventions are  important to understand and follow strictly.    - While updating objects (via PUT), entire object must be sent with all fields.     - For example, in order to add a remedy method to the condition below:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": []       }       ```     - send the entire object with updated and non-updated fields:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": [{\"type\": \"DisplayMessage\", \"message\": \"test message\"}]       }       ```    - In case Controller returns an error (non-2xx HTTP status code), response body is JSON.     The \"message\" field contains information about the error.     HTTP 422 \"Unprocessable Entity\" has extra `errors` field to list all the issues with specific fields.    - Empty string (\"\") is considered a different value than \"null\" or field being omitted from JSON.     Omitting the field is recommended if no value is intended.     Empty string (\"\") will be almost always rejected as invalid value.    - There are common pattern between many objects:     - **Configuration Objects**: There are many objects with common fields, namely \"id\", \"name\", \"notes\", \"created\"       and \"updated\". These entities are listed, queried, created, updated and deleted in a similar fashion.     - **Distinguished Name**: Users and Devices are identified with what is called Distinguished Names, as used in        LDAP. The distinguished format that identifies a device and a user combination is        \"CN=\\<Device ID\\>,CN=\\<username\\>,OU=\\<Identity Provider Name\\>\". Some objects have the        \"userDistinguishedName\" field, which does not include the CN for Device ID.        This identifies a user on every device.

API version: API version 21.0
Contact: appgatesdp.support@appgate.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// SiteAllOfVpn VPN configuration for this Site.
type SiteAllOfVpn struct {
	// Source NAT.
	Snat     *bool                 `json:"snat,omitempty"`
	Tls      *SiteAllOfVpnTls      `json:"tls,omitempty"`
	Dtls     *SiteAllOfVpnDtls     `json:"dtls,omitempty"`
	RouteVia *SiteAllOfVpnRouteVia `json:"routeVia,omitempty"`
	// Whether to enable URL Access feature or not.
	UrlAccessEnabled *bool `json:"urlAccessEnabled,omitempty"`
	// P12 files for proxying traffic for URL Access feature.
	UrlAccessP12s []P12 `json:"urlAccessP12s,omitempty"`
	// Frequency configuration for generating IP Access audit logs for a connection.
	IpAccessLogIntervalSeconds *float32 `json:"ipAccessLogIntervalSeconds,omitempty"`
	// Whether to log NAT traffic or not.
	LogNatIpAndNatPort *bool `json:"logNatIpAndNatPort,omitempty"`
}

// NewSiteAllOfVpn instantiates a new SiteAllOfVpn object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSiteAllOfVpn() *SiteAllOfVpn {
	this := SiteAllOfVpn{}
	var snat bool = false
	this.Snat = &snat
	var ipAccessLogIntervalSeconds float32 = 120
	this.IpAccessLogIntervalSeconds = &ipAccessLogIntervalSeconds
	var logNatIpAndNatPort bool = false
	this.LogNatIpAndNatPort = &logNatIpAndNatPort
	return &this
}

// NewSiteAllOfVpnWithDefaults instantiates a new SiteAllOfVpn object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSiteAllOfVpnWithDefaults() *SiteAllOfVpn {
	this := SiteAllOfVpn{}
	var snat bool = false
	this.Snat = &snat
	var ipAccessLogIntervalSeconds float32 = 120
	this.IpAccessLogIntervalSeconds = &ipAccessLogIntervalSeconds
	var logNatIpAndNatPort bool = false
	this.LogNatIpAndNatPort = &logNatIpAndNatPort
	return &this
}

// GetSnat returns the Snat field value if set, zero value otherwise.
func (o *SiteAllOfVpn) GetSnat() bool {
	if o == nil || o.Snat == nil {
		var ret bool
		return ret
	}
	return *o.Snat
}

// GetSnatOk returns a tuple with the Snat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteAllOfVpn) GetSnatOk() (*bool, bool) {
	if o == nil || o.Snat == nil {
		return nil, false
	}
	return o.Snat, true
}

// HasSnat returns a boolean if a field has been set.
func (o *SiteAllOfVpn) HasSnat() bool {
	if o != nil && o.Snat != nil {
		return true
	}

	return false
}

// SetSnat gets a reference to the given bool and assigns it to the Snat field.
func (o *SiteAllOfVpn) SetSnat(v bool) {
	o.Snat = &v
}

// GetTls returns the Tls field value if set, zero value otherwise.
func (o *SiteAllOfVpn) GetTls() SiteAllOfVpnTls {
	if o == nil || o.Tls == nil {
		var ret SiteAllOfVpnTls
		return ret
	}
	return *o.Tls
}

// GetTlsOk returns a tuple with the Tls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteAllOfVpn) GetTlsOk() (*SiteAllOfVpnTls, bool) {
	if o == nil || o.Tls == nil {
		return nil, false
	}
	return o.Tls, true
}

// HasTls returns a boolean if a field has been set.
func (o *SiteAllOfVpn) HasTls() bool {
	if o != nil && o.Tls != nil {
		return true
	}

	return false
}

// SetTls gets a reference to the given SiteAllOfVpnTls and assigns it to the Tls field.
func (o *SiteAllOfVpn) SetTls(v SiteAllOfVpnTls) {
	o.Tls = &v
}

// GetDtls returns the Dtls field value if set, zero value otherwise.
func (o *SiteAllOfVpn) GetDtls() SiteAllOfVpnDtls {
	if o == nil || o.Dtls == nil {
		var ret SiteAllOfVpnDtls
		return ret
	}
	return *o.Dtls
}

// GetDtlsOk returns a tuple with the Dtls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteAllOfVpn) GetDtlsOk() (*SiteAllOfVpnDtls, bool) {
	if o == nil || o.Dtls == nil {
		return nil, false
	}
	return o.Dtls, true
}

// HasDtls returns a boolean if a field has been set.
func (o *SiteAllOfVpn) HasDtls() bool {
	if o != nil && o.Dtls != nil {
		return true
	}

	return false
}

// SetDtls gets a reference to the given SiteAllOfVpnDtls and assigns it to the Dtls field.
func (o *SiteAllOfVpn) SetDtls(v SiteAllOfVpnDtls) {
	o.Dtls = &v
}

// GetRouteVia returns the RouteVia field value if set, zero value otherwise.
func (o *SiteAllOfVpn) GetRouteVia() SiteAllOfVpnRouteVia {
	if o == nil || o.RouteVia == nil {
		var ret SiteAllOfVpnRouteVia
		return ret
	}
	return *o.RouteVia
}

// GetRouteViaOk returns a tuple with the RouteVia field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteAllOfVpn) GetRouteViaOk() (*SiteAllOfVpnRouteVia, bool) {
	if o == nil || o.RouteVia == nil {
		return nil, false
	}
	return o.RouteVia, true
}

// HasRouteVia returns a boolean if a field has been set.
func (o *SiteAllOfVpn) HasRouteVia() bool {
	if o != nil && o.RouteVia != nil {
		return true
	}

	return false
}

// SetRouteVia gets a reference to the given SiteAllOfVpnRouteVia and assigns it to the RouteVia field.
func (o *SiteAllOfVpn) SetRouteVia(v SiteAllOfVpnRouteVia) {
	o.RouteVia = &v
}

// GetUrlAccessEnabled returns the UrlAccessEnabled field value if set, zero value otherwise.
func (o *SiteAllOfVpn) GetUrlAccessEnabled() bool {
	if o == nil || o.UrlAccessEnabled == nil {
		var ret bool
		return ret
	}
	return *o.UrlAccessEnabled
}

// GetUrlAccessEnabledOk returns a tuple with the UrlAccessEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteAllOfVpn) GetUrlAccessEnabledOk() (*bool, bool) {
	if o == nil || o.UrlAccessEnabled == nil {
		return nil, false
	}
	return o.UrlAccessEnabled, true
}

// HasUrlAccessEnabled returns a boolean if a field has been set.
func (o *SiteAllOfVpn) HasUrlAccessEnabled() bool {
	if o != nil && o.UrlAccessEnabled != nil {
		return true
	}

	return false
}

// SetUrlAccessEnabled gets a reference to the given bool and assigns it to the UrlAccessEnabled field.
func (o *SiteAllOfVpn) SetUrlAccessEnabled(v bool) {
	o.UrlAccessEnabled = &v
}

// GetUrlAccessP12s returns the UrlAccessP12s field value if set, zero value otherwise.
func (o *SiteAllOfVpn) GetUrlAccessP12s() []P12 {
	if o == nil || o.UrlAccessP12s == nil {
		var ret []P12
		return ret
	}
	return o.UrlAccessP12s
}

// GetUrlAccessP12sOk returns a tuple with the UrlAccessP12s field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteAllOfVpn) GetUrlAccessP12sOk() ([]P12, bool) {
	if o == nil || o.UrlAccessP12s == nil {
		return nil, false
	}
	return o.UrlAccessP12s, true
}

// HasUrlAccessP12s returns a boolean if a field has been set.
func (o *SiteAllOfVpn) HasUrlAccessP12s() bool {
	if o != nil && o.UrlAccessP12s != nil {
		return true
	}

	return false
}

// SetUrlAccessP12s gets a reference to the given []P12 and assigns it to the UrlAccessP12s field.
func (o *SiteAllOfVpn) SetUrlAccessP12s(v []P12) {
	o.UrlAccessP12s = v
}

// GetIpAccessLogIntervalSeconds returns the IpAccessLogIntervalSeconds field value if set, zero value otherwise.
func (o *SiteAllOfVpn) GetIpAccessLogIntervalSeconds() float32 {
	if o == nil || o.IpAccessLogIntervalSeconds == nil {
		var ret float32
		return ret
	}
	return *o.IpAccessLogIntervalSeconds
}

// GetIpAccessLogIntervalSecondsOk returns a tuple with the IpAccessLogIntervalSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteAllOfVpn) GetIpAccessLogIntervalSecondsOk() (*float32, bool) {
	if o == nil || o.IpAccessLogIntervalSeconds == nil {
		return nil, false
	}
	return o.IpAccessLogIntervalSeconds, true
}

// HasIpAccessLogIntervalSeconds returns a boolean if a field has been set.
func (o *SiteAllOfVpn) HasIpAccessLogIntervalSeconds() bool {
	if o != nil && o.IpAccessLogIntervalSeconds != nil {
		return true
	}

	return false
}

// SetIpAccessLogIntervalSeconds gets a reference to the given float32 and assigns it to the IpAccessLogIntervalSeconds field.
func (o *SiteAllOfVpn) SetIpAccessLogIntervalSeconds(v float32) {
	o.IpAccessLogIntervalSeconds = &v
}

// GetLogNatIpAndNatPort returns the LogNatIpAndNatPort field value if set, zero value otherwise.
func (o *SiteAllOfVpn) GetLogNatIpAndNatPort() bool {
	if o == nil || o.LogNatIpAndNatPort == nil {
		var ret bool
		return ret
	}
	return *o.LogNatIpAndNatPort
}

// GetLogNatIpAndNatPortOk returns a tuple with the LogNatIpAndNatPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteAllOfVpn) GetLogNatIpAndNatPortOk() (*bool, bool) {
	if o == nil || o.LogNatIpAndNatPort == nil {
		return nil, false
	}
	return o.LogNatIpAndNatPort, true
}

// HasLogNatIpAndNatPort returns a boolean if a field has been set.
func (o *SiteAllOfVpn) HasLogNatIpAndNatPort() bool {
	if o != nil && o.LogNatIpAndNatPort != nil {
		return true
	}

	return false
}

// SetLogNatIpAndNatPort gets a reference to the given bool and assigns it to the LogNatIpAndNatPort field.
func (o *SiteAllOfVpn) SetLogNatIpAndNatPort(v bool) {
	o.LogNatIpAndNatPort = &v
}

func (o SiteAllOfVpn) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Snat != nil {
		toSerialize["snat"] = o.Snat
	}
	if o.Tls != nil {
		toSerialize["tls"] = o.Tls
	}
	if o.Dtls != nil {
		toSerialize["dtls"] = o.Dtls
	}
	if o.RouteVia != nil {
		toSerialize["routeVia"] = o.RouteVia
	}
	if o.UrlAccessEnabled != nil {
		toSerialize["urlAccessEnabled"] = o.UrlAccessEnabled
	}
	if o.UrlAccessP12s != nil {
		toSerialize["urlAccessP12s"] = o.UrlAccessP12s
	}
	if o.IpAccessLogIntervalSeconds != nil {
		toSerialize["ipAccessLogIntervalSeconds"] = o.IpAccessLogIntervalSeconds
	}
	if o.LogNatIpAndNatPort != nil {
		toSerialize["logNatIpAndNatPort"] = o.LogNatIpAndNatPort
	}
	return json.Marshal(toSerialize)
}

type NullableSiteAllOfVpn struct {
	value *SiteAllOfVpn
	isSet bool
}

func (v NullableSiteAllOfVpn) Get() *SiteAllOfVpn {
	return v.value
}

func (v *NullableSiteAllOfVpn) Set(val *SiteAllOfVpn) {
	v.value = val
	v.isSet = true
}

func (v NullableSiteAllOfVpn) IsSet() bool {
	return v.isSet
}

func (v *NullableSiteAllOfVpn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSiteAllOfVpn(val *SiteAllOfVpn) *NullableSiteAllOfVpn {
	return &NullableSiteAllOfVpn{value: val, isSet: true}
}

func (v NullableSiteAllOfVpn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSiteAllOfVpn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
