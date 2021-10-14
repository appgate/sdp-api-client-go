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

// SessionInfoDistinguishedNameData Session Details reported by the Gateway.
type SessionInfoDistinguishedNameData struct {
	UserClaims   *map[string]interface{} `json:"userClaims,omitempty"`
	DeviceClaims *map[string]interface{} `json:"deviceClaims,omitempty"`
	SystemClaims *map[string]interface{} `json:"systemClaims,omitempty"`
	// Entitlement information reported by the Gateway. The key is the Entitlement name.
	EntitlementInfos *map[string]SessionInfoDistinguishedNameEntitlementInfos `json:"entitlementInfos,omitempty"`
	// The Discovered Apps by this Gateway. Note that this API does not filter out the existing Entitlements like the App Discovery API does.
	DiscoveredApps *[]string `json:"discoveredApps,omitempty"`
	// VPN details of the session
	Vpn *map[string]interface{} `json:"vpn,omitempty"`
	// The Site name for the Gateway.
	Site *string `json:"site,omitempty"`
}

// NewSessionInfoDistinguishedNameData instantiates a new SessionInfoDistinguishedNameData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSessionInfoDistinguishedNameData() *SessionInfoDistinguishedNameData {
	this := SessionInfoDistinguishedNameData{}
	return &this
}

// NewSessionInfoDistinguishedNameDataWithDefaults instantiates a new SessionInfoDistinguishedNameData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSessionInfoDistinguishedNameDataWithDefaults() *SessionInfoDistinguishedNameData {
	this := SessionInfoDistinguishedNameData{}
	return &this
}

// GetUserClaims returns the UserClaims field value if set, zero value otherwise.
func (o *SessionInfoDistinguishedNameData) GetUserClaims() map[string]interface{} {
	if o == nil || o.UserClaims == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.UserClaims
}

// GetUserClaimsOk returns a tuple with the UserClaims field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionInfoDistinguishedNameData) GetUserClaimsOk() (*map[string]interface{}, bool) {
	if o == nil || o.UserClaims == nil {
		return nil, false
	}
	return o.UserClaims, true
}

// HasUserClaims returns a boolean if a field has been set.
func (o *SessionInfoDistinguishedNameData) HasUserClaims() bool {
	if o != nil && o.UserClaims != nil {
		return true
	}

	return false
}

// SetUserClaims gets a reference to the given map[string]interface{} and assigns it to the UserClaims field.
func (o *SessionInfoDistinguishedNameData) SetUserClaims(v map[string]interface{}) {
	o.UserClaims = &v
}

// GetDeviceClaims returns the DeviceClaims field value if set, zero value otherwise.
func (o *SessionInfoDistinguishedNameData) GetDeviceClaims() map[string]interface{} {
	if o == nil || o.DeviceClaims == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.DeviceClaims
}

// GetDeviceClaimsOk returns a tuple with the DeviceClaims field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionInfoDistinguishedNameData) GetDeviceClaimsOk() (*map[string]interface{}, bool) {
	if o == nil || o.DeviceClaims == nil {
		return nil, false
	}
	return o.DeviceClaims, true
}

// HasDeviceClaims returns a boolean if a field has been set.
func (o *SessionInfoDistinguishedNameData) HasDeviceClaims() bool {
	if o != nil && o.DeviceClaims != nil {
		return true
	}

	return false
}

// SetDeviceClaims gets a reference to the given map[string]interface{} and assigns it to the DeviceClaims field.
func (o *SessionInfoDistinguishedNameData) SetDeviceClaims(v map[string]interface{}) {
	o.DeviceClaims = &v
}

// GetSystemClaims returns the SystemClaims field value if set, zero value otherwise.
func (o *SessionInfoDistinguishedNameData) GetSystemClaims() map[string]interface{} {
	if o == nil || o.SystemClaims == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.SystemClaims
}

// GetSystemClaimsOk returns a tuple with the SystemClaims field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionInfoDistinguishedNameData) GetSystemClaimsOk() (*map[string]interface{}, bool) {
	if o == nil || o.SystemClaims == nil {
		return nil, false
	}
	return o.SystemClaims, true
}

// HasSystemClaims returns a boolean if a field has been set.
func (o *SessionInfoDistinguishedNameData) HasSystemClaims() bool {
	if o != nil && o.SystemClaims != nil {
		return true
	}

	return false
}

// SetSystemClaims gets a reference to the given map[string]interface{} and assigns it to the SystemClaims field.
func (o *SessionInfoDistinguishedNameData) SetSystemClaims(v map[string]interface{}) {
	o.SystemClaims = &v
}

// GetEntitlementInfos returns the EntitlementInfos field value if set, zero value otherwise.
func (o *SessionInfoDistinguishedNameData) GetEntitlementInfos() map[string]SessionInfoDistinguishedNameEntitlementInfos {
	if o == nil || o.EntitlementInfos == nil {
		var ret map[string]SessionInfoDistinguishedNameEntitlementInfos
		return ret
	}
	return *o.EntitlementInfos
}

// GetEntitlementInfosOk returns a tuple with the EntitlementInfos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionInfoDistinguishedNameData) GetEntitlementInfosOk() (*map[string]SessionInfoDistinguishedNameEntitlementInfos, bool) {
	if o == nil || o.EntitlementInfos == nil {
		return nil, false
	}
	return o.EntitlementInfos, true
}

// HasEntitlementInfos returns a boolean if a field has been set.
func (o *SessionInfoDistinguishedNameData) HasEntitlementInfos() bool {
	if o != nil && o.EntitlementInfos != nil {
		return true
	}

	return false
}

// SetEntitlementInfos gets a reference to the given map[string]SessionInfoDistinguishedNameEntitlementInfos and assigns it to the EntitlementInfos field.
func (o *SessionInfoDistinguishedNameData) SetEntitlementInfos(v map[string]SessionInfoDistinguishedNameEntitlementInfos) {
	o.EntitlementInfos = &v
}

// GetDiscoveredApps returns the DiscoveredApps field value if set, zero value otherwise.
func (o *SessionInfoDistinguishedNameData) GetDiscoveredApps() []string {
	if o == nil || o.DiscoveredApps == nil {
		var ret []string
		return ret
	}
	return *o.DiscoveredApps
}

// GetDiscoveredAppsOk returns a tuple with the DiscoveredApps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionInfoDistinguishedNameData) GetDiscoveredAppsOk() (*[]string, bool) {
	if o == nil || o.DiscoveredApps == nil {
		return nil, false
	}
	return o.DiscoveredApps, true
}

// HasDiscoveredApps returns a boolean if a field has been set.
func (o *SessionInfoDistinguishedNameData) HasDiscoveredApps() bool {
	if o != nil && o.DiscoveredApps != nil {
		return true
	}

	return false
}

// SetDiscoveredApps gets a reference to the given []string and assigns it to the DiscoveredApps field.
func (o *SessionInfoDistinguishedNameData) SetDiscoveredApps(v []string) {
	o.DiscoveredApps = &v
}

// GetVpn returns the Vpn field value if set, zero value otherwise.
func (o *SessionInfoDistinguishedNameData) GetVpn() map[string]interface{} {
	if o == nil || o.Vpn == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Vpn
}

// GetVpnOk returns a tuple with the Vpn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionInfoDistinguishedNameData) GetVpnOk() (*map[string]interface{}, bool) {
	if o == nil || o.Vpn == nil {
		return nil, false
	}
	return o.Vpn, true
}

// HasVpn returns a boolean if a field has been set.
func (o *SessionInfoDistinguishedNameData) HasVpn() bool {
	if o != nil && o.Vpn != nil {
		return true
	}

	return false
}

// SetVpn gets a reference to the given map[string]interface{} and assigns it to the Vpn field.
func (o *SessionInfoDistinguishedNameData) SetVpn(v map[string]interface{}) {
	o.Vpn = &v
}

// GetSite returns the Site field value if set, zero value otherwise.
func (o *SessionInfoDistinguishedNameData) GetSite() string {
	if o == nil || o.Site == nil {
		var ret string
		return ret
	}
	return *o.Site
}

// GetSiteOk returns a tuple with the Site field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionInfoDistinguishedNameData) GetSiteOk() (*string, bool) {
	if o == nil || o.Site == nil {
		return nil, false
	}
	return o.Site, true
}

// HasSite returns a boolean if a field has been set.
func (o *SessionInfoDistinguishedNameData) HasSite() bool {
	if o != nil && o.Site != nil {
		return true
	}

	return false
}

// SetSite gets a reference to the given string and assigns it to the Site field.
func (o *SessionInfoDistinguishedNameData) SetSite(v string) {
	o.Site = &v
}

func (o SessionInfoDistinguishedNameData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UserClaims != nil {
		toSerialize["userClaims"] = o.UserClaims
	}
	if o.DeviceClaims != nil {
		toSerialize["deviceClaims"] = o.DeviceClaims
	}
	if o.SystemClaims != nil {
		toSerialize["systemClaims"] = o.SystemClaims
	}
	if o.EntitlementInfos != nil {
		toSerialize["entitlementInfos"] = o.EntitlementInfos
	}
	if o.DiscoveredApps != nil {
		toSerialize["discoveredApps"] = o.DiscoveredApps
	}
	if o.Vpn != nil {
		toSerialize["vpn"] = o.Vpn
	}
	if o.Site != nil {
		toSerialize["site"] = o.Site
	}
	return json.Marshal(toSerialize)
}

type NullableSessionInfoDistinguishedNameData struct {
	value *SessionInfoDistinguishedNameData
	isSet bool
}

func (v NullableSessionInfoDistinguishedNameData) Get() *SessionInfoDistinguishedNameData {
	return v.value
}

func (v *NullableSessionInfoDistinguishedNameData) Set(val *SessionInfoDistinguishedNameData) {
	v.value = val
	v.isSet = true
}

func (v NullableSessionInfoDistinguishedNameData) IsSet() bool {
	return v.isSet
}

func (v *NullableSessionInfoDistinguishedNameData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSessionInfoDistinguishedNameData(val *SessionInfoDistinguishedNameData) *NullableSessionInfoDistinguishedNameData {
	return &NullableSessionInfoDistinguishedNameData{value: val, isSet: true}
}

func (v NullableSessionInfoDistinguishedNameData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSessionInfoDistinguishedNameData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
