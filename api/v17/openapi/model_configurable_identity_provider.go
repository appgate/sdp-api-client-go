/*
Appgate SDP Controller REST API

# About   This specification documents the REST API calls for the Appgate SDP Controller.    Please refer to the REST API chapter in the manual or contact Appgate support with any questions about   this functionality. # Getting Started   Requirements for API scripting:   - Access to the Admin/API TLS Connection (default port 8443) of a Controller appliance.     (https://sdphelp.appgate.com/adminguide/appliance-function-configure.html?anchor=admin-api)   - An API user with relevant permissions.     (https://sdphelp.appgate.com/adminguide/administrative-roles-configure.html)   - In order to use the simple login API, Admin MFA must be disabled or the API user must be excluded.     (https://sdphelp.appgate.com/adminguide/mfa-for-admins.html) # Base path   HTTPS requests must be sent to the Admin Interface hostname and port, with **_/admin** path.    For example: **https://appgate.company.com:8443/admin**    All requests must have the **Accept** header as:    **application/vnd.appgate.peer-v17+json** # API Conventions   API conventions are  important to understand and follow strictly.    - While updating objects (via PUT), entire object must be sent with all fields.     - For example, in order to add a remedy method to the condition below:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": []       }       ```     - send the entire object with updated and non-updated fields:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": [{\"type\": \"DisplayMessage\", \"message\": \"test message\"}]       }       ```    - In case Controller returns an error (non-2xx HTTP status code), response body is JSON.     The \"message\" field contains information about the error.     HTTP 422 \"Unprocessable Entity\" has extra `errors` field to list all the issues with specific fields.    - Empty string (\"\") is considered a different value than \"null\" or field being omitted from JSON.     Omitting the field is recommended if no value is intended.     Empty string (\"\") will be almost always rejected as invalid value.    - There are common pattern between many objects:     - **Configuration Objects**: There are many objects with common fields, namely \"id\", \"name\", \"notes\", \"created\"       and \"updated\". These entities are listed, queried, created, updated and deleted in a similar fashion.     - **Distinguished Name**: Users and Devices are identified with what is called Distinguished Names, as used in        LDAP. The distinguished format that identifies a device and a user combination is        \"CN=\\<Device ID\\>,CN=\\<username\\>,OU=\\<Identity Provider Name\\>\". Some objects have the        \"userDistinguishedName\" field, which does not include the CN for Device ID.        This identifies a user on every device.

API version: API version 17.4
Contact: appgatesdp.support@appgate.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// ConfigurableIdentityProvider struct for ConfigurableIdentityProvider
type ConfigurableIdentityProvider struct {
	// ID of the object.
	Id *string `json:"id,omitempty"`
	// Name of the object.
	Name string `json:"name"`
	// Notes for the object. Used for documentation purposes.
	Notes *string `json:"notes,omitempty"`
	// Create date.
	Created *time.Time `json:"created,omitempty"`
	// Last update date.
	Updated *time.Time `json:"updated,omitempty"`
	// Array of tags.
	Tags []string `json:"tags,omitempty"`
	// The type of the Identity Provider.
	Type string `json:"type"`
	// The IPv4 Pool ID the users in this Identity Provider are going to use to allocate IP addresses for the tunnels.
	IpPoolV4 *string `json:"ipPoolV4,omitempty"`
	// The IPv6 Pool ID the users in this Identity Provider are going to use to allocate IP addresses for the tunnels.
	IpPoolV6 *string `json:"ipPoolV6,omitempty"`
	// The mapping of Identity Provider attributes to claims.
	ClaimMappings []ClaimMappingsInner `json:"claimMappings,omitempty"`
	// ID of the User Claim Scripts to run during authorization.
	UserScripts []string `json:"userScripts,omitempty"`
	// The DNS servers to be assigned to the Clients of the users in this Identity Provider.
	DnsServers []string `json:"dnsServers,omitempty"`
	// The DNS search domains to be assigned to Clients of the users in this Identity Provider.
	DnsSearchDomains []string `json:"dnsSearchDomains,omitempty"`
	// Whether the provider will be listed in the Admin UI or not.
	AdminProvider *bool `json:"adminProvider,omitempty"`
	// The device limit per user. The existing on-boarded devices will still be able to sign in even if the limit is exceeded.
	DeviceLimitPerUser *int32                                          `json:"deviceLimitPerUser,omitempty"`
	OnBoarding2FA      *ConfigurableIdentityProviderAllOfOnBoarding2FA `json:"onBoarding2FA,omitempty"`
	// (Desktop) clients will sign out automatically after the user has been inactive on the device for the configured duration. Set it to 0 to disable.
	InactivityTimeoutMinutes *int32 `json:"inactivityTimeoutMinutes,omitempty"`
	// If enabled, Windows Client will configure the network profile as \"DomainAuthenticated\".
	EnforceWindowsNetworkProfileAsDomain *bool `json:"enforceWindowsNetworkProfileAsDomain,omitempty"`
	// Whether the Windows Client will block local DNS requests or not.
	BlockLocalDnsRequests *bool `json:"blockLocalDnsRequests,omitempty"`
	// The mapping of Identity Provider on demand attributes to claims.
	OnDemandClaimMappings []OnDemandClaimMappingsInner `json:"onDemandClaimMappings,omitempty"`
}

// NewConfigurableIdentityProvider instantiates a new ConfigurableIdentityProvider object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigurableIdentityProvider(name string, type_ string) *ConfigurableIdentityProvider {
	this := ConfigurableIdentityProvider{}
	this.Name = name
	this.Type = type_
	var adminProvider bool = false
	this.AdminProvider = &adminProvider
	var deviceLimitPerUser int32 = 100
	this.DeviceLimitPerUser = &deviceLimitPerUser
	var inactivityTimeoutMinutes int32 = 0
	this.InactivityTimeoutMinutes = &inactivityTimeoutMinutes
	var blockLocalDnsRequests bool = false
	this.BlockLocalDnsRequests = &blockLocalDnsRequests
	return &this
}

// NewConfigurableIdentityProviderWithDefaults instantiates a new ConfigurableIdentityProvider object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigurableIdentityProviderWithDefaults() *ConfigurableIdentityProvider {
	this := ConfigurableIdentityProvider{}
	var adminProvider bool = false
	this.AdminProvider = &adminProvider
	var deviceLimitPerUser int32 = 100
	this.DeviceLimitPerUser = &deviceLimitPerUser
	var inactivityTimeoutMinutes int32 = 0
	this.InactivityTimeoutMinutes = &inactivityTimeoutMinutes
	var blockLocalDnsRequests bool = false
	this.BlockLocalDnsRequests = &blockLocalDnsRequests
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ConfigurableIdentityProvider) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurableIdentityProvider) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ConfigurableIdentityProvider) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ConfigurableIdentityProvider) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *ConfigurableIdentityProvider) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ConfigurableIdentityProvider) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ConfigurableIdentityProvider) SetName(v string) {
	o.Name = v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *ConfigurableIdentityProvider) GetNotes() string {
	if o == nil || o.Notes == nil {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurableIdentityProvider) GetNotesOk() (*string, bool) {
	if o == nil || o.Notes == nil {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *ConfigurableIdentityProvider) HasNotes() bool {
	if o != nil && o.Notes != nil {
		return true
	}

	return false
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *ConfigurableIdentityProvider) SetNotes(v string) {
	o.Notes = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *ConfigurableIdentityProvider) GetCreated() time.Time {
	if o == nil || o.Created == nil {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurableIdentityProvider) GetCreatedOk() (*time.Time, bool) {
	if o == nil || o.Created == nil {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *ConfigurableIdentityProvider) HasCreated() bool {
	if o != nil && o.Created != nil {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *ConfigurableIdentityProvider) SetCreated(v time.Time) {
	o.Created = &v
}

// GetUpdated returns the Updated field value if set, zero value otherwise.
func (o *ConfigurableIdentityProvider) GetUpdated() time.Time {
	if o == nil || o.Updated == nil {
		var ret time.Time
		return ret
	}
	return *o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurableIdentityProvider) GetUpdatedOk() (*time.Time, bool) {
	if o == nil || o.Updated == nil {
		return nil, false
	}
	return o.Updated, true
}

// HasUpdated returns a boolean if a field has been set.
func (o *ConfigurableIdentityProvider) HasUpdated() bool {
	if o != nil && o.Updated != nil {
		return true
	}

	return false
}

// SetUpdated gets a reference to the given time.Time and assigns it to the Updated field.
func (o *ConfigurableIdentityProvider) SetUpdated(v time.Time) {
	o.Updated = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ConfigurableIdentityProvider) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurableIdentityProvider) GetTagsOk() ([]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ConfigurableIdentityProvider) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *ConfigurableIdentityProvider) SetTags(v []string) {
	o.Tags = v
}

// GetType returns the Type field value
func (o *ConfigurableIdentityProvider) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ConfigurableIdentityProvider) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ConfigurableIdentityProvider) SetType(v string) {
	o.Type = v
}

// GetIpPoolV4 returns the IpPoolV4 field value if set, zero value otherwise.
func (o *ConfigurableIdentityProvider) GetIpPoolV4() string {
	if o == nil || o.IpPoolV4 == nil {
		var ret string
		return ret
	}
	return *o.IpPoolV4
}

// GetIpPoolV4Ok returns a tuple with the IpPoolV4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurableIdentityProvider) GetIpPoolV4Ok() (*string, bool) {
	if o == nil || o.IpPoolV4 == nil {
		return nil, false
	}
	return o.IpPoolV4, true
}

// HasIpPoolV4 returns a boolean if a field has been set.
func (o *ConfigurableIdentityProvider) HasIpPoolV4() bool {
	if o != nil && o.IpPoolV4 != nil {
		return true
	}

	return false
}

// SetIpPoolV4 gets a reference to the given string and assigns it to the IpPoolV4 field.
func (o *ConfigurableIdentityProvider) SetIpPoolV4(v string) {
	o.IpPoolV4 = &v
}

// GetIpPoolV6 returns the IpPoolV6 field value if set, zero value otherwise.
func (o *ConfigurableIdentityProvider) GetIpPoolV6() string {
	if o == nil || o.IpPoolV6 == nil {
		var ret string
		return ret
	}
	return *o.IpPoolV6
}

// GetIpPoolV6Ok returns a tuple with the IpPoolV6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurableIdentityProvider) GetIpPoolV6Ok() (*string, bool) {
	if o == nil || o.IpPoolV6 == nil {
		return nil, false
	}
	return o.IpPoolV6, true
}

// HasIpPoolV6 returns a boolean if a field has been set.
func (o *ConfigurableIdentityProvider) HasIpPoolV6() bool {
	if o != nil && o.IpPoolV6 != nil {
		return true
	}

	return false
}

// SetIpPoolV6 gets a reference to the given string and assigns it to the IpPoolV6 field.
func (o *ConfigurableIdentityProvider) SetIpPoolV6(v string) {
	o.IpPoolV6 = &v
}

// GetClaimMappings returns the ClaimMappings field value if set, zero value otherwise.
func (o *ConfigurableIdentityProvider) GetClaimMappings() []ClaimMappingsInner {
	if o == nil || o.ClaimMappings == nil {
		var ret []ClaimMappingsInner
		return ret
	}
	return o.ClaimMappings
}

// GetClaimMappingsOk returns a tuple with the ClaimMappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurableIdentityProvider) GetClaimMappingsOk() ([]ClaimMappingsInner, bool) {
	if o == nil || o.ClaimMappings == nil {
		return nil, false
	}
	return o.ClaimMappings, true
}

// HasClaimMappings returns a boolean if a field has been set.
func (o *ConfigurableIdentityProvider) HasClaimMappings() bool {
	if o != nil && o.ClaimMappings != nil {
		return true
	}

	return false
}

// SetClaimMappings gets a reference to the given []ClaimMappingsInner and assigns it to the ClaimMappings field.
func (o *ConfigurableIdentityProvider) SetClaimMappings(v []ClaimMappingsInner) {
	o.ClaimMappings = v
}

// GetUserScripts returns the UserScripts field value if set, zero value otherwise.
func (o *ConfigurableIdentityProvider) GetUserScripts() []string {
	if o == nil || o.UserScripts == nil {
		var ret []string
		return ret
	}
	return o.UserScripts
}

// GetUserScriptsOk returns a tuple with the UserScripts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurableIdentityProvider) GetUserScriptsOk() ([]string, bool) {
	if o == nil || o.UserScripts == nil {
		return nil, false
	}
	return o.UserScripts, true
}

// HasUserScripts returns a boolean if a field has been set.
func (o *ConfigurableIdentityProvider) HasUserScripts() bool {
	if o != nil && o.UserScripts != nil {
		return true
	}

	return false
}

// SetUserScripts gets a reference to the given []string and assigns it to the UserScripts field.
func (o *ConfigurableIdentityProvider) SetUserScripts(v []string) {
	o.UserScripts = v
}

// GetDnsServers returns the DnsServers field value if set, zero value otherwise.
func (o *ConfigurableIdentityProvider) GetDnsServers() []string {
	if o == nil || o.DnsServers == nil {
		var ret []string
		return ret
	}
	return o.DnsServers
}

// GetDnsServersOk returns a tuple with the DnsServers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurableIdentityProvider) GetDnsServersOk() ([]string, bool) {
	if o == nil || o.DnsServers == nil {
		return nil, false
	}
	return o.DnsServers, true
}

// HasDnsServers returns a boolean if a field has been set.
func (o *ConfigurableIdentityProvider) HasDnsServers() bool {
	if o != nil && o.DnsServers != nil {
		return true
	}

	return false
}

// SetDnsServers gets a reference to the given []string and assigns it to the DnsServers field.
func (o *ConfigurableIdentityProvider) SetDnsServers(v []string) {
	o.DnsServers = v
}

// GetDnsSearchDomains returns the DnsSearchDomains field value if set, zero value otherwise.
func (o *ConfigurableIdentityProvider) GetDnsSearchDomains() []string {
	if o == nil || o.DnsSearchDomains == nil {
		var ret []string
		return ret
	}
	return o.DnsSearchDomains
}

// GetDnsSearchDomainsOk returns a tuple with the DnsSearchDomains field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurableIdentityProvider) GetDnsSearchDomainsOk() ([]string, bool) {
	if o == nil || o.DnsSearchDomains == nil {
		return nil, false
	}
	return o.DnsSearchDomains, true
}

// HasDnsSearchDomains returns a boolean if a field has been set.
func (o *ConfigurableIdentityProvider) HasDnsSearchDomains() bool {
	if o != nil && o.DnsSearchDomains != nil {
		return true
	}

	return false
}

// SetDnsSearchDomains gets a reference to the given []string and assigns it to the DnsSearchDomains field.
func (o *ConfigurableIdentityProvider) SetDnsSearchDomains(v []string) {
	o.DnsSearchDomains = v
}

// GetAdminProvider returns the AdminProvider field value if set, zero value otherwise.
func (o *ConfigurableIdentityProvider) GetAdminProvider() bool {
	if o == nil || o.AdminProvider == nil {
		var ret bool
		return ret
	}
	return *o.AdminProvider
}

// GetAdminProviderOk returns a tuple with the AdminProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurableIdentityProvider) GetAdminProviderOk() (*bool, bool) {
	if o == nil || o.AdminProvider == nil {
		return nil, false
	}
	return o.AdminProvider, true
}

// HasAdminProvider returns a boolean if a field has been set.
func (o *ConfigurableIdentityProvider) HasAdminProvider() bool {
	if o != nil && o.AdminProvider != nil {
		return true
	}

	return false
}

// SetAdminProvider gets a reference to the given bool and assigns it to the AdminProvider field.
func (o *ConfigurableIdentityProvider) SetAdminProvider(v bool) {
	o.AdminProvider = &v
}

// GetDeviceLimitPerUser returns the DeviceLimitPerUser field value if set, zero value otherwise.
func (o *ConfigurableIdentityProvider) GetDeviceLimitPerUser() int32 {
	if o == nil || o.DeviceLimitPerUser == nil {
		var ret int32
		return ret
	}
	return *o.DeviceLimitPerUser
}

// GetDeviceLimitPerUserOk returns a tuple with the DeviceLimitPerUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurableIdentityProvider) GetDeviceLimitPerUserOk() (*int32, bool) {
	if o == nil || o.DeviceLimitPerUser == nil {
		return nil, false
	}
	return o.DeviceLimitPerUser, true
}

// HasDeviceLimitPerUser returns a boolean if a field has been set.
func (o *ConfigurableIdentityProvider) HasDeviceLimitPerUser() bool {
	if o != nil && o.DeviceLimitPerUser != nil {
		return true
	}

	return false
}

// SetDeviceLimitPerUser gets a reference to the given int32 and assigns it to the DeviceLimitPerUser field.
func (o *ConfigurableIdentityProvider) SetDeviceLimitPerUser(v int32) {
	o.DeviceLimitPerUser = &v
}

// GetOnBoarding2FA returns the OnBoarding2FA field value if set, zero value otherwise.
func (o *ConfigurableIdentityProvider) GetOnBoarding2FA() ConfigurableIdentityProviderAllOfOnBoarding2FA {
	if o == nil || o.OnBoarding2FA == nil {
		var ret ConfigurableIdentityProviderAllOfOnBoarding2FA
		return ret
	}
	return *o.OnBoarding2FA
}

// GetOnBoarding2FAOk returns a tuple with the OnBoarding2FA field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurableIdentityProvider) GetOnBoarding2FAOk() (*ConfigurableIdentityProviderAllOfOnBoarding2FA, bool) {
	if o == nil || o.OnBoarding2FA == nil {
		return nil, false
	}
	return o.OnBoarding2FA, true
}

// HasOnBoarding2FA returns a boolean if a field has been set.
func (o *ConfigurableIdentityProvider) HasOnBoarding2FA() bool {
	if o != nil && o.OnBoarding2FA != nil {
		return true
	}

	return false
}

// SetOnBoarding2FA gets a reference to the given ConfigurableIdentityProviderAllOfOnBoarding2FA and assigns it to the OnBoarding2FA field.
func (o *ConfigurableIdentityProvider) SetOnBoarding2FA(v ConfigurableIdentityProviderAllOfOnBoarding2FA) {
	o.OnBoarding2FA = &v
}

// GetInactivityTimeoutMinutes returns the InactivityTimeoutMinutes field value if set, zero value otherwise.
func (o *ConfigurableIdentityProvider) GetInactivityTimeoutMinutes() int32 {
	if o == nil || o.InactivityTimeoutMinutes == nil {
		var ret int32
		return ret
	}
	return *o.InactivityTimeoutMinutes
}

// GetInactivityTimeoutMinutesOk returns a tuple with the InactivityTimeoutMinutes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurableIdentityProvider) GetInactivityTimeoutMinutesOk() (*int32, bool) {
	if o == nil || o.InactivityTimeoutMinutes == nil {
		return nil, false
	}
	return o.InactivityTimeoutMinutes, true
}

// HasInactivityTimeoutMinutes returns a boolean if a field has been set.
func (o *ConfigurableIdentityProvider) HasInactivityTimeoutMinutes() bool {
	if o != nil && o.InactivityTimeoutMinutes != nil {
		return true
	}

	return false
}

// SetInactivityTimeoutMinutes gets a reference to the given int32 and assigns it to the InactivityTimeoutMinutes field.
func (o *ConfigurableIdentityProvider) SetInactivityTimeoutMinutes(v int32) {
	o.InactivityTimeoutMinutes = &v
}

// GetEnforceWindowsNetworkProfileAsDomain returns the EnforceWindowsNetworkProfileAsDomain field value if set, zero value otherwise.
func (o *ConfigurableIdentityProvider) GetEnforceWindowsNetworkProfileAsDomain() bool {
	if o == nil || o.EnforceWindowsNetworkProfileAsDomain == nil {
		var ret bool
		return ret
	}
	return *o.EnforceWindowsNetworkProfileAsDomain
}

// GetEnforceWindowsNetworkProfileAsDomainOk returns a tuple with the EnforceWindowsNetworkProfileAsDomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurableIdentityProvider) GetEnforceWindowsNetworkProfileAsDomainOk() (*bool, bool) {
	if o == nil || o.EnforceWindowsNetworkProfileAsDomain == nil {
		return nil, false
	}
	return o.EnforceWindowsNetworkProfileAsDomain, true
}

// HasEnforceWindowsNetworkProfileAsDomain returns a boolean if a field has been set.
func (o *ConfigurableIdentityProvider) HasEnforceWindowsNetworkProfileAsDomain() bool {
	if o != nil && o.EnforceWindowsNetworkProfileAsDomain != nil {
		return true
	}

	return false
}

// SetEnforceWindowsNetworkProfileAsDomain gets a reference to the given bool and assigns it to the EnforceWindowsNetworkProfileAsDomain field.
func (o *ConfigurableIdentityProvider) SetEnforceWindowsNetworkProfileAsDomain(v bool) {
	o.EnforceWindowsNetworkProfileAsDomain = &v
}

// GetBlockLocalDnsRequests returns the BlockLocalDnsRequests field value if set, zero value otherwise.
func (o *ConfigurableIdentityProvider) GetBlockLocalDnsRequests() bool {
	if o == nil || o.BlockLocalDnsRequests == nil {
		var ret bool
		return ret
	}
	return *o.BlockLocalDnsRequests
}

// GetBlockLocalDnsRequestsOk returns a tuple with the BlockLocalDnsRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurableIdentityProvider) GetBlockLocalDnsRequestsOk() (*bool, bool) {
	if o == nil || o.BlockLocalDnsRequests == nil {
		return nil, false
	}
	return o.BlockLocalDnsRequests, true
}

// HasBlockLocalDnsRequests returns a boolean if a field has been set.
func (o *ConfigurableIdentityProvider) HasBlockLocalDnsRequests() bool {
	if o != nil && o.BlockLocalDnsRequests != nil {
		return true
	}

	return false
}

// SetBlockLocalDnsRequests gets a reference to the given bool and assigns it to the BlockLocalDnsRequests field.
func (o *ConfigurableIdentityProvider) SetBlockLocalDnsRequests(v bool) {
	o.BlockLocalDnsRequests = &v
}

// GetOnDemandClaimMappings returns the OnDemandClaimMappings field value if set, zero value otherwise.
func (o *ConfigurableIdentityProvider) GetOnDemandClaimMappings() []OnDemandClaimMappingsInner {
	if o == nil || o.OnDemandClaimMappings == nil {
		var ret []OnDemandClaimMappingsInner
		return ret
	}
	return o.OnDemandClaimMappings
}

// GetOnDemandClaimMappingsOk returns a tuple with the OnDemandClaimMappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurableIdentityProvider) GetOnDemandClaimMappingsOk() ([]OnDemandClaimMappingsInner, bool) {
	if o == nil || o.OnDemandClaimMappings == nil {
		return nil, false
	}
	return o.OnDemandClaimMappings, true
}

// HasOnDemandClaimMappings returns a boolean if a field has been set.
func (o *ConfigurableIdentityProvider) HasOnDemandClaimMappings() bool {
	if o != nil && o.OnDemandClaimMappings != nil {
		return true
	}

	return false
}

// SetOnDemandClaimMappings gets a reference to the given []OnDemandClaimMappingsInner and assigns it to the OnDemandClaimMappings field.
func (o *ConfigurableIdentityProvider) SetOnDemandClaimMappings(v []OnDemandClaimMappingsInner) {
	o.OnDemandClaimMappings = v
}

func (o ConfigurableIdentityProvider) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Notes != nil {
		toSerialize["notes"] = o.Notes
	}
	if o.Created != nil {
		toSerialize["created"] = o.Created
	}
	if o.Updated != nil {
		toSerialize["updated"] = o.Updated
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if o.IpPoolV4 != nil {
		toSerialize["ipPoolV4"] = o.IpPoolV4
	}
	if o.IpPoolV6 != nil {
		toSerialize["ipPoolV6"] = o.IpPoolV6
	}
	if o.ClaimMappings != nil {
		toSerialize["claimMappings"] = o.ClaimMappings
	}
	if o.UserScripts != nil {
		toSerialize["userScripts"] = o.UserScripts
	}
	if o.DnsServers != nil {
		toSerialize["dnsServers"] = o.DnsServers
	}
	if o.DnsSearchDomains != nil {
		toSerialize["dnsSearchDomains"] = o.DnsSearchDomains
	}
	if o.AdminProvider != nil {
		toSerialize["adminProvider"] = o.AdminProvider
	}
	if o.DeviceLimitPerUser != nil {
		toSerialize["deviceLimitPerUser"] = o.DeviceLimitPerUser
	}
	if o.OnBoarding2FA != nil {
		toSerialize["onBoarding2FA"] = o.OnBoarding2FA
	}
	if o.InactivityTimeoutMinutes != nil {
		toSerialize["inactivityTimeoutMinutes"] = o.InactivityTimeoutMinutes
	}
	if o.EnforceWindowsNetworkProfileAsDomain != nil {
		toSerialize["enforceWindowsNetworkProfileAsDomain"] = o.EnforceWindowsNetworkProfileAsDomain
	}
	if o.BlockLocalDnsRequests != nil {
		toSerialize["blockLocalDnsRequests"] = o.BlockLocalDnsRequests
	}
	if o.OnDemandClaimMappings != nil {
		toSerialize["onDemandClaimMappings"] = o.OnDemandClaimMappings
	}
	return json.Marshal(toSerialize)
}

type NullableConfigurableIdentityProvider struct {
	value *ConfigurableIdentityProvider
	isSet bool
}

func (v NullableConfigurableIdentityProvider) Get() *ConfigurableIdentityProvider {
	return v.value
}

func (v *NullableConfigurableIdentityProvider) Set(val *ConfigurableIdentityProvider) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigurableIdentityProvider) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigurableIdentityProvider) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigurableIdentityProvider(val *ConfigurableIdentityProvider) *NullableConfigurableIdentityProvider {
	return &NullableConfigurableIdentityProvider{value: val, isSet: true}
}

func (v NullableConfigurableIdentityProvider) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigurableIdentityProvider) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
