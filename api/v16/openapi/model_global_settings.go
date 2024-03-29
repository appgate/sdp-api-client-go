/*
Appgate SDP Controller REST API

# About   This specification documents the REST API calls for the Appgate SDP Controller.    Please refer to the REST API chapter in the manual or contact Appgate support with any questions about   this functionality. # Getting Started   Requirements for API scripting:   - Access to the Admin/API TLS Connection (default port 8443) of a Controller appliance.     (https://sdphelp.appgate.com/adminguide/appliance-function-configure.html?anchor=admin-api)   - An API user with relevant permissions.     (https://sdphelp.appgate.com/adminguide/administrative-roles-configure.html)   - In order to use the simple login API, Admin MFA must be disabled or the API user must be excluded.     (https://sdphelp.appgate.com/adminguide/mfa-for-admins.html) # Base path   HTTPS requests must be sent to the Admin Interface hostname and port, with **_/admin** path.    For example: **https://appgate.company.com:8443/admin**    All requests must have the **Accept** header as:    **application/vnd.appgate.peer-v16+json** # API Conventions   API conventions are  important to understand and follow strictly.    - While updating objects (via PUT), entire object must be sent with all fields.     - For example, in order to add a remedy method to the condition below:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": []       }       ```     - send the entire object with updated and non-updated fields:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": [{\"type\": \"DisplayMessage\", \"message\": \"test message\"}]       }       ```    - In case Controller returns an error (non-2xx HTTP status code), response body is JSON.     The \"message\" field contains information about the error.     HTTP 422 \"Unprocessable Entity\" has extra `errors` field to list all the issues with specific fields.    - Empty string (\"\") is considered a different value than \"null\" or field being omitted from JSON.     Omitting the field is recommend if no value is intended.     Empty string (\"\") will be almost always rejected as invalid value.    - There are common pattern between many objects:     - **Configuration Objects**: There are many objects with common fields, namely \"id\", \"name\", \"notes\", \"created\"       and \"updated\". These entities are listed, queried, created, updated and deleted in a similar fashion.     - **Distinguished Name**: Users and Devices are identified with what is called Distinguished Names, as used in        LDAP. The distinguished format that identifies a device and a user combination is        \"CN=\\<Device ID\\>,CN=\\<username\\>,OU=\\<Identity Provider Name\\>\". Some objects have the        \"userDistinguishedName\" field, which does not include the CN for Device ID.        This identifies a user on every device.

API version: API version 16.5
Contact: appgatesdp.support@appgate.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// GlobalSettings struct for GlobalSettings
type GlobalSettings struct {
	// Number of minutes the Claims Token is valid both for administrators and clients.
	ClaimsTokenExpiration float32 `json:"claimsTokenExpiration"`
	// Number of minutes the Entitlement Token is valid for clients.
	EntitlementTokenExpiration float32 `json:"entitlementTokenExpiration"`
	// Number of minutes the administration Token is valid for administrators.
	AdministrationTokenExpiration float32 `json:"administrationTokenExpiration"`
	// Number of minutes the VPN certificates is valid for clients.
	VpnCertificateExpiration float32 `json:"vpnCertificateExpiration"`
	// SPA mode.
	SpaMode *string `json:"spaMode,omitempty"`
	// The configured message will be displayed on the login UI.
	LoginBannerMessage *string `json:"loginBannerMessage,omitempty"`
	// The configured message will be displayed after a successful login.
	MessageOfTheDay *string `json:"messageOfTheDay,omitempty"`
	// Whether the backup API is enabled or not.
	BackupApiEnabled *bool `json:"backupApiEnabled,omitempty"`
	// The passphrase to encrypt Appliance Backups when backup API is used.
	BackupPassphrase *string `json:"backupPassphrase,omitempty"`
	// FIPS 140-2 Compliant Tunneling.
	// Deprecated
	Fips *bool `json:"fips,omitempty"`
	// Whether the automatic GeoIp updates are enabled or not.
	GeoIpUpdates *bool `json:"geoIpUpdates,omitempty"`
	// Audit Log persistence mode.
	AuditLogPersistenceMode string `json:"auditLogPersistenceMode"`
	// Domains to monitor for for App Discovery feature.
	AppDiscoveryDomains *[]string `json:"appDiscoveryDomains,omitempty"`
	// A randomly generated ID during first installation to identify the Collective.
	CollectiveId *string `json:"collectiveId,omitempty"`
}

// NewGlobalSettings instantiates a new GlobalSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGlobalSettings(claimsTokenExpiration float32, entitlementTokenExpiration float32, administrationTokenExpiration float32, vpnCertificateExpiration float32, auditLogPersistenceMode string) *GlobalSettings {
	this := GlobalSettings{}
	this.ClaimsTokenExpiration = claimsTokenExpiration
	this.EntitlementTokenExpiration = entitlementTokenExpiration
	this.AdministrationTokenExpiration = administrationTokenExpiration
	this.VpnCertificateExpiration = vpnCertificateExpiration
	this.AuditLogPersistenceMode = auditLogPersistenceMode
	return &this
}

// NewGlobalSettingsWithDefaults instantiates a new GlobalSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGlobalSettingsWithDefaults() *GlobalSettings {
	this := GlobalSettings{}
	return &this
}

// GetClaimsTokenExpiration returns the ClaimsTokenExpiration field value
func (o *GlobalSettings) GetClaimsTokenExpiration() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.ClaimsTokenExpiration
}

// GetClaimsTokenExpirationOk returns a tuple with the ClaimsTokenExpiration field value
// and a boolean to check if the value has been set.
func (o *GlobalSettings) GetClaimsTokenExpirationOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClaimsTokenExpiration, true
}

// SetClaimsTokenExpiration sets field value
func (o *GlobalSettings) SetClaimsTokenExpiration(v float32) {
	o.ClaimsTokenExpiration = v
}

// GetEntitlementTokenExpiration returns the EntitlementTokenExpiration field value
func (o *GlobalSettings) GetEntitlementTokenExpiration() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.EntitlementTokenExpiration
}

// GetEntitlementTokenExpirationOk returns a tuple with the EntitlementTokenExpiration field value
// and a boolean to check if the value has been set.
func (o *GlobalSettings) GetEntitlementTokenExpirationOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntitlementTokenExpiration, true
}

// SetEntitlementTokenExpiration sets field value
func (o *GlobalSettings) SetEntitlementTokenExpiration(v float32) {
	o.EntitlementTokenExpiration = v
}

// GetAdministrationTokenExpiration returns the AdministrationTokenExpiration field value
func (o *GlobalSettings) GetAdministrationTokenExpiration() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.AdministrationTokenExpiration
}

// GetAdministrationTokenExpirationOk returns a tuple with the AdministrationTokenExpiration field value
// and a boolean to check if the value has been set.
func (o *GlobalSettings) GetAdministrationTokenExpirationOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdministrationTokenExpiration, true
}

// SetAdministrationTokenExpiration sets field value
func (o *GlobalSettings) SetAdministrationTokenExpiration(v float32) {
	o.AdministrationTokenExpiration = v
}

// GetVpnCertificateExpiration returns the VpnCertificateExpiration field value
func (o *GlobalSettings) GetVpnCertificateExpiration() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.VpnCertificateExpiration
}

// GetVpnCertificateExpirationOk returns a tuple with the VpnCertificateExpiration field value
// and a boolean to check if the value has been set.
func (o *GlobalSettings) GetVpnCertificateExpirationOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VpnCertificateExpiration, true
}

// SetVpnCertificateExpiration sets field value
func (o *GlobalSettings) SetVpnCertificateExpiration(v float32) {
	o.VpnCertificateExpiration = v
}

// GetSpaMode returns the SpaMode field value if set, zero value otherwise.
func (o *GlobalSettings) GetSpaMode() string {
	if o == nil || o.SpaMode == nil {
		var ret string
		return ret
	}
	return *o.SpaMode
}

// GetSpaModeOk returns a tuple with the SpaMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalSettings) GetSpaModeOk() (*string, bool) {
	if o == nil || o.SpaMode == nil {
		return nil, false
	}
	return o.SpaMode, true
}

// HasSpaMode returns a boolean if a field has been set.
func (o *GlobalSettings) HasSpaMode() bool {
	if o != nil && o.SpaMode != nil {
		return true
	}

	return false
}

// SetSpaMode gets a reference to the given string and assigns it to the SpaMode field.
func (o *GlobalSettings) SetSpaMode(v string) {
	o.SpaMode = &v
}

// GetLoginBannerMessage returns the LoginBannerMessage field value if set, zero value otherwise.
func (o *GlobalSettings) GetLoginBannerMessage() string {
	if o == nil || o.LoginBannerMessage == nil {
		var ret string
		return ret
	}
	return *o.LoginBannerMessage
}

// GetLoginBannerMessageOk returns a tuple with the LoginBannerMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalSettings) GetLoginBannerMessageOk() (*string, bool) {
	if o == nil || o.LoginBannerMessage == nil {
		return nil, false
	}
	return o.LoginBannerMessage, true
}

// HasLoginBannerMessage returns a boolean if a field has been set.
func (o *GlobalSettings) HasLoginBannerMessage() bool {
	if o != nil && o.LoginBannerMessage != nil {
		return true
	}

	return false
}

// SetLoginBannerMessage gets a reference to the given string and assigns it to the LoginBannerMessage field.
func (o *GlobalSettings) SetLoginBannerMessage(v string) {
	o.LoginBannerMessage = &v
}

// GetMessageOfTheDay returns the MessageOfTheDay field value if set, zero value otherwise.
func (o *GlobalSettings) GetMessageOfTheDay() string {
	if o == nil || o.MessageOfTheDay == nil {
		var ret string
		return ret
	}
	return *o.MessageOfTheDay
}

// GetMessageOfTheDayOk returns a tuple with the MessageOfTheDay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalSettings) GetMessageOfTheDayOk() (*string, bool) {
	if o == nil || o.MessageOfTheDay == nil {
		return nil, false
	}
	return o.MessageOfTheDay, true
}

// HasMessageOfTheDay returns a boolean if a field has been set.
func (o *GlobalSettings) HasMessageOfTheDay() bool {
	if o != nil && o.MessageOfTheDay != nil {
		return true
	}

	return false
}

// SetMessageOfTheDay gets a reference to the given string and assigns it to the MessageOfTheDay field.
func (o *GlobalSettings) SetMessageOfTheDay(v string) {
	o.MessageOfTheDay = &v
}

// GetBackupApiEnabled returns the BackupApiEnabled field value if set, zero value otherwise.
func (o *GlobalSettings) GetBackupApiEnabled() bool {
	if o == nil || o.BackupApiEnabled == nil {
		var ret bool
		return ret
	}
	return *o.BackupApiEnabled
}

// GetBackupApiEnabledOk returns a tuple with the BackupApiEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalSettings) GetBackupApiEnabledOk() (*bool, bool) {
	if o == nil || o.BackupApiEnabled == nil {
		return nil, false
	}
	return o.BackupApiEnabled, true
}

// HasBackupApiEnabled returns a boolean if a field has been set.
func (o *GlobalSettings) HasBackupApiEnabled() bool {
	if o != nil && o.BackupApiEnabled != nil {
		return true
	}

	return false
}

// SetBackupApiEnabled gets a reference to the given bool and assigns it to the BackupApiEnabled field.
func (o *GlobalSettings) SetBackupApiEnabled(v bool) {
	o.BackupApiEnabled = &v
}

// GetBackupPassphrase returns the BackupPassphrase field value if set, zero value otherwise.
func (o *GlobalSettings) GetBackupPassphrase() string {
	if o == nil || o.BackupPassphrase == nil {
		var ret string
		return ret
	}
	return *o.BackupPassphrase
}

// GetBackupPassphraseOk returns a tuple with the BackupPassphrase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalSettings) GetBackupPassphraseOk() (*string, bool) {
	if o == nil || o.BackupPassphrase == nil {
		return nil, false
	}
	return o.BackupPassphrase, true
}

// HasBackupPassphrase returns a boolean if a field has been set.
func (o *GlobalSettings) HasBackupPassphrase() bool {
	if o != nil && o.BackupPassphrase != nil {
		return true
	}

	return false
}

// SetBackupPassphrase gets a reference to the given string and assigns it to the BackupPassphrase field.
func (o *GlobalSettings) SetBackupPassphrase(v string) {
	o.BackupPassphrase = &v
}

// GetFips returns the Fips field value if set, zero value otherwise.
// Deprecated
func (o *GlobalSettings) GetFips() bool {
	if o == nil || o.Fips == nil {
		var ret bool
		return ret
	}
	return *o.Fips
}

// GetFipsOk returns a tuple with the Fips field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *GlobalSettings) GetFipsOk() (*bool, bool) {
	if o == nil || o.Fips == nil {
		return nil, false
	}
	return o.Fips, true
}

// HasFips returns a boolean if a field has been set.
func (o *GlobalSettings) HasFips() bool {
	if o != nil && o.Fips != nil {
		return true
	}

	return false
}

// SetFips gets a reference to the given bool and assigns it to the Fips field.
// Deprecated
func (o *GlobalSettings) SetFips(v bool) {
	o.Fips = &v
}

// GetGeoIpUpdates returns the GeoIpUpdates field value if set, zero value otherwise.
func (o *GlobalSettings) GetGeoIpUpdates() bool {
	if o == nil || o.GeoIpUpdates == nil {
		var ret bool
		return ret
	}
	return *o.GeoIpUpdates
}

// GetGeoIpUpdatesOk returns a tuple with the GeoIpUpdates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalSettings) GetGeoIpUpdatesOk() (*bool, bool) {
	if o == nil || o.GeoIpUpdates == nil {
		return nil, false
	}
	return o.GeoIpUpdates, true
}

// HasGeoIpUpdates returns a boolean if a field has been set.
func (o *GlobalSettings) HasGeoIpUpdates() bool {
	if o != nil && o.GeoIpUpdates != nil {
		return true
	}

	return false
}

// SetGeoIpUpdates gets a reference to the given bool and assigns it to the GeoIpUpdates field.
func (o *GlobalSettings) SetGeoIpUpdates(v bool) {
	o.GeoIpUpdates = &v
}

// GetAuditLogPersistenceMode returns the AuditLogPersistenceMode field value
func (o *GlobalSettings) GetAuditLogPersistenceMode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuditLogPersistenceMode
}

// GetAuditLogPersistenceModeOk returns a tuple with the AuditLogPersistenceMode field value
// and a boolean to check if the value has been set.
func (o *GlobalSettings) GetAuditLogPersistenceModeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuditLogPersistenceMode, true
}

// SetAuditLogPersistenceMode sets field value
func (o *GlobalSettings) SetAuditLogPersistenceMode(v string) {
	o.AuditLogPersistenceMode = v
}

// GetAppDiscoveryDomains returns the AppDiscoveryDomains field value if set, zero value otherwise.
func (o *GlobalSettings) GetAppDiscoveryDomains() []string {
	if o == nil || o.AppDiscoveryDomains == nil {
		var ret []string
		return ret
	}
	return *o.AppDiscoveryDomains
}

// GetAppDiscoveryDomainsOk returns a tuple with the AppDiscoveryDomains field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalSettings) GetAppDiscoveryDomainsOk() (*[]string, bool) {
	if o == nil || o.AppDiscoveryDomains == nil {
		return nil, false
	}
	return o.AppDiscoveryDomains, true
}

// HasAppDiscoveryDomains returns a boolean if a field has been set.
func (o *GlobalSettings) HasAppDiscoveryDomains() bool {
	if o != nil && o.AppDiscoveryDomains != nil {
		return true
	}

	return false
}

// SetAppDiscoveryDomains gets a reference to the given []string and assigns it to the AppDiscoveryDomains field.
func (o *GlobalSettings) SetAppDiscoveryDomains(v []string) {
	o.AppDiscoveryDomains = &v
}

// GetCollectiveId returns the CollectiveId field value if set, zero value otherwise.
func (o *GlobalSettings) GetCollectiveId() string {
	if o == nil || o.CollectiveId == nil {
		var ret string
		return ret
	}
	return *o.CollectiveId
}

// GetCollectiveIdOk returns a tuple with the CollectiveId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalSettings) GetCollectiveIdOk() (*string, bool) {
	if o == nil || o.CollectiveId == nil {
		return nil, false
	}
	return o.CollectiveId, true
}

// HasCollectiveId returns a boolean if a field has been set.
func (o *GlobalSettings) HasCollectiveId() bool {
	if o != nil && o.CollectiveId != nil {
		return true
	}

	return false
}

// SetCollectiveId gets a reference to the given string and assigns it to the CollectiveId field.
func (o *GlobalSettings) SetCollectiveId(v string) {
	o.CollectiveId = &v
}

func (o GlobalSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["claimsTokenExpiration"] = o.ClaimsTokenExpiration
	}
	if true {
		toSerialize["entitlementTokenExpiration"] = o.EntitlementTokenExpiration
	}
	if true {
		toSerialize["administrationTokenExpiration"] = o.AdministrationTokenExpiration
	}
	if true {
		toSerialize["vpnCertificateExpiration"] = o.VpnCertificateExpiration
	}
	if o.SpaMode != nil {
		toSerialize["spaMode"] = o.SpaMode
	}
	if o.LoginBannerMessage != nil {
		toSerialize["loginBannerMessage"] = o.LoginBannerMessage
	}
	if o.MessageOfTheDay != nil {
		toSerialize["messageOfTheDay"] = o.MessageOfTheDay
	}
	if o.BackupApiEnabled != nil {
		toSerialize["backupApiEnabled"] = o.BackupApiEnabled
	}
	if o.BackupPassphrase != nil {
		toSerialize["backupPassphrase"] = o.BackupPassphrase
	}
	if o.Fips != nil {
		toSerialize["fips"] = o.Fips
	}
	if o.GeoIpUpdates != nil {
		toSerialize["geoIpUpdates"] = o.GeoIpUpdates
	}
	if true {
		toSerialize["auditLogPersistenceMode"] = o.AuditLogPersistenceMode
	}
	if o.AppDiscoveryDomains != nil {
		toSerialize["appDiscoveryDomains"] = o.AppDiscoveryDomains
	}
	if o.CollectiveId != nil {
		toSerialize["collectiveId"] = o.CollectiveId
	}
	return json.Marshal(toSerialize)
}

type NullableGlobalSettings struct {
	value *GlobalSettings
	isSet bool
}

func (v NullableGlobalSettings) Get() *GlobalSettings {
	return v.value
}

func (v *NullableGlobalSettings) Set(val *GlobalSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableGlobalSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableGlobalSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGlobalSettings(val *GlobalSettings) *NullableGlobalSettings {
	return &NullableGlobalSettings{value: val, isSet: true}
}

func (v NullableGlobalSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGlobalSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
