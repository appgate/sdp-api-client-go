/*
 * AppGate SDP Controller REST API
 *
 * # About   This specification documents the REST API calls for the AppGate SDP Controller.    Please refer to the Integration chapter in the manual or contact AppGate support with any questions about   this functionality. # Getting Started   Requirements for API scripting:   - Access to the peer interface (default port 444) of a Controller appliance.     (https://sdphelp.appgate.com/adminguide/appliances-configure.html?anchor=peer)   - An API user with relevant permissions.     (https://sdphelp.appgate.com/adminguide/administrative-roles-configure.html)   - In order to use the simple login API, Admin MFA must be disabled or the API user must be excluded.     (https://sdphelp.appgate.com/adminguide/mfa-for-admins.html) # Base path   HTTPS requests must be sent to the Peer Interface hostname and port, with **_/admin** path.    For example: **https://appgate.company.com:444/admin**    All requests must have the **Accept** header as:    **application/vnd.appgate.peer-v12+json** # API Conventions   API conventions are  important to understand and follow strictly.    - While updating objects (via PUT), entire object must be sent with all fields.     - For example, if in order to add a remedy method to the condition below:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": []       }       ```     - send the entire object with updated and non-updated fields:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": [{\"type\": \"DisplayMessage\", \"message\": \"test message\"}]       }       ```    - In case Controller returns an error (non-2xx HTTP status code), response body is JSON.     The \"message\" field contains information abot the error.     HTTP 422 \"Unprocessable Entity\" has extra `errors` field to list all the issues with specific fields.    - Empty string (\"\") is considered a different value than \"null\" or field being omitted from JSON.     Omitting the field is recommend if no value is intended.     Empty string (\"\") will be almost always rejected as invalid value.    - There are common pattern between many objects:     - **Configuration Objects**: There are many objects with common fields, namely \"id\", \"name\", \"notes\", \"created\"       and \"updated\". These entities are listed, queried, created, updated and deleted in a similar fashion.     - **Distinguished Name**: Users and Devices are identified with what is called Distinguished Names, as used in        LDAP. The distinguished format that identifies a device and a user combination is        \"CN=\\<Device ID\\>,CN=\\<username\\>,OU=\\<Identity Provider Name\\>\". Some objects have the        \"userDistinguishedName\" field, which does not include the CN for Device ID.        This identifies a user on every device.
 *
 * API version: API version 12
 * Contact: appgatesdp.support@appgate.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"encoding/json"
)

// IdentityProviderAllOf Represents an Identity Provider.
type IdentityProviderAllOf struct {
	// The type of the Identity Provider.
	Type string `json:"type"`
	// The name displayed to the user. \"name\" field is used for Distinguished Name generation. Deprecated as of 5.1 since the Client does not have the option to choose Identity Provider anymore.
	DisplayName *string `json:"displayName,omitempty"`
	// Whether the provider will be chosen by default in the Client UI. If enabled, it will remove the default flag of the current default Identity Provider.
	Default *bool `json:"default,omitempty"`
	// Whether the provider will be listed in the Client UI or not. Deprecated as of 5.1 since the Client does not have the option to choose Identity Provider anymore.
	ClientProvider *bool `json:"clientProvider,omitempty"`
	// Whether the provider will be listed in the Admin UI or not.
	AdminProvider *bool                               `json:"adminProvider,omitempty"`
	OnBoarding2FA *IdentityProviderAllOfOnBoarding2FA `json:"onBoarding2FA,omitempty"`
	// Client on-boarding type. Deprecated as of 5.0. Use onBoarding2FA object instead.
	OnBoardingType *string `json:"onBoardingType,omitempty"`
	// On-boarding MFA Provider ID if \"onBoardingType\" is Require2FA.  Deprecated as of 5.0. Use onBoarding2FA object instead.
	OnBoardingOtpProvider *string `json:"onBoardingOtpProvider,omitempty"`
	// On-boarding MFA message to be displayed on the Client UI if \"onBoardingType\" is Require2FA. Deprecated as of 5.0. Use onBoarding2FA object instead.
	OnBoardingOtpMessage *string `json:"onBoardingOtpMessage,omitempty"`
	// (Desktop) clients will sign out automatically after the user has been inactive on the device for the configured duration. Set it to 0 to disable.
	InactivityTimeoutMinutes *int32 `json:"inactivityTimeoutMinutes,omitempty"`
	// The IPv4 Pool ID the users in this Identity Provider are going to use to allocate IP addresses for the tunnels.
	IpPoolV4 *string `json:"ipPoolV4,omitempty"`
	// The IPv6 Pool ID the users in this Identity Provider are going to use to allocate IP addresses for the tunnels.
	IpPoolV6 *string `json:"ipPoolV6,omitempty"`
	// The dns servers to be assigned to the Clients of the users in this Identity Provider.
	DnsServers *[]string `json:"dnsServers,omitempty"`
	// The dns search domains to be assigned to Clients of the users in this Identity Provider.
	DnsSearchDomains *[]string `json:"dnsSearchDomains,omitempty"`
	// Whether the Windows Client will block local DNS requests or not.
	BlockLocalDnsRequests *bool `json:"blockLocalDnsRequests,omitempty"`
	// The mapping of Identity Provider attributes to claims.
	ClaimMappings *[]map[string]interface{} `json:"claimMappings,omitempty"`
	// The mapping of Identity Provider on demand attributes to claims.
	OnDemandClaimMappings *[]map[string]interface{} `json:"onDemandClaimMappings,omitempty"`
}

// NewIdentityProviderAllOf instantiates a new IdentityProviderAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityProviderAllOf(type_ string) *IdentityProviderAllOf {
	this := IdentityProviderAllOf{}
	this.Type = type_
	var default_ bool = false
	this.Default = &default_
	var clientProvider bool = false
	this.ClientProvider = &clientProvider
	var adminProvider bool = false
	this.AdminProvider = &adminProvider
	var inactivityTimeoutMinutes int32 = 0
	this.InactivityTimeoutMinutes = &inactivityTimeoutMinutes
	var blockLocalDnsRequests bool = false
	this.BlockLocalDnsRequests = &blockLocalDnsRequests
	return &this
}

// NewIdentityProviderAllOfWithDefaults instantiates a new IdentityProviderAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityProviderAllOfWithDefaults() *IdentityProviderAllOf {
	this := IdentityProviderAllOf{}
	var default_ bool = false
	this.Default = &default_
	var clientProvider bool = false
	this.ClientProvider = &clientProvider
	var adminProvider bool = false
	this.AdminProvider = &adminProvider
	var inactivityTimeoutMinutes int32 = 0
	this.InactivityTimeoutMinutes = &inactivityTimeoutMinutes
	var blockLocalDnsRequests bool = false
	this.BlockLocalDnsRequests = &blockLocalDnsRequests
	return &this
}

// GetType returns the Type field value
func (o *IdentityProviderAllOf) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *IdentityProviderAllOf) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *IdentityProviderAllOf) SetType(v string) {
	o.Type = v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *IdentityProviderAllOf) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderAllOf) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *IdentityProviderAllOf) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *IdentityProviderAllOf) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *IdentityProviderAllOf) GetDefault() bool {
	if o == nil || o.Default == nil {
		var ret bool
		return ret
	}
	return *o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderAllOf) GetDefaultOk() (*bool, bool) {
	if o == nil || o.Default == nil {
		return nil, false
	}
	return o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *IdentityProviderAllOf) HasDefault() bool {
	if o != nil && o.Default != nil {
		return true
	}

	return false
}

// SetDefault gets a reference to the given bool and assigns it to the Default field.
func (o *IdentityProviderAllOf) SetDefault(v bool) {
	o.Default = &v
}

// GetClientProvider returns the ClientProvider field value if set, zero value otherwise.
func (o *IdentityProviderAllOf) GetClientProvider() bool {
	if o == nil || o.ClientProvider == nil {
		var ret bool
		return ret
	}
	return *o.ClientProvider
}

// GetClientProviderOk returns a tuple with the ClientProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderAllOf) GetClientProviderOk() (*bool, bool) {
	if o == nil || o.ClientProvider == nil {
		return nil, false
	}
	return o.ClientProvider, true
}

// HasClientProvider returns a boolean if a field has been set.
func (o *IdentityProviderAllOf) HasClientProvider() bool {
	if o != nil && o.ClientProvider != nil {
		return true
	}

	return false
}

// SetClientProvider gets a reference to the given bool and assigns it to the ClientProvider field.
func (o *IdentityProviderAllOf) SetClientProvider(v bool) {
	o.ClientProvider = &v
}

// GetAdminProvider returns the AdminProvider field value if set, zero value otherwise.
func (o *IdentityProviderAllOf) GetAdminProvider() bool {
	if o == nil || o.AdminProvider == nil {
		var ret bool
		return ret
	}
	return *o.AdminProvider
}

// GetAdminProviderOk returns a tuple with the AdminProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderAllOf) GetAdminProviderOk() (*bool, bool) {
	if o == nil || o.AdminProvider == nil {
		return nil, false
	}
	return o.AdminProvider, true
}

// HasAdminProvider returns a boolean if a field has been set.
func (o *IdentityProviderAllOf) HasAdminProvider() bool {
	if o != nil && o.AdminProvider != nil {
		return true
	}

	return false
}

// SetAdminProvider gets a reference to the given bool and assigns it to the AdminProvider field.
func (o *IdentityProviderAllOf) SetAdminProvider(v bool) {
	o.AdminProvider = &v
}

// GetOnBoarding2FA returns the OnBoarding2FA field value if set, zero value otherwise.
func (o *IdentityProviderAllOf) GetOnBoarding2FA() IdentityProviderAllOfOnBoarding2FA {
	if o == nil || o.OnBoarding2FA == nil {
		var ret IdentityProviderAllOfOnBoarding2FA
		return ret
	}
	return *o.OnBoarding2FA
}

// GetOnBoarding2FAOk returns a tuple with the OnBoarding2FA field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderAllOf) GetOnBoarding2FAOk() (*IdentityProviderAllOfOnBoarding2FA, bool) {
	if o == nil || o.OnBoarding2FA == nil {
		return nil, false
	}
	return o.OnBoarding2FA, true
}

// HasOnBoarding2FA returns a boolean if a field has been set.
func (o *IdentityProviderAllOf) HasOnBoarding2FA() bool {
	if o != nil && o.OnBoarding2FA != nil {
		return true
	}

	return false
}

// SetOnBoarding2FA gets a reference to the given IdentityProviderAllOfOnBoarding2FA and assigns it to the OnBoarding2FA field.
func (o *IdentityProviderAllOf) SetOnBoarding2FA(v IdentityProviderAllOfOnBoarding2FA) {
	o.OnBoarding2FA = &v
}

// GetOnBoardingType returns the OnBoardingType field value if set, zero value otherwise.
func (o *IdentityProviderAllOf) GetOnBoardingType() string {
	if o == nil || o.OnBoardingType == nil {
		var ret string
		return ret
	}
	return *o.OnBoardingType
}

// GetOnBoardingTypeOk returns a tuple with the OnBoardingType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderAllOf) GetOnBoardingTypeOk() (*string, bool) {
	if o == nil || o.OnBoardingType == nil {
		return nil, false
	}
	return o.OnBoardingType, true
}

// HasOnBoardingType returns a boolean if a field has been set.
func (o *IdentityProviderAllOf) HasOnBoardingType() bool {
	if o != nil && o.OnBoardingType != nil {
		return true
	}

	return false
}

// SetOnBoardingType gets a reference to the given string and assigns it to the OnBoardingType field.
func (o *IdentityProviderAllOf) SetOnBoardingType(v string) {
	o.OnBoardingType = &v
}

// GetOnBoardingOtpProvider returns the OnBoardingOtpProvider field value if set, zero value otherwise.
func (o *IdentityProviderAllOf) GetOnBoardingOtpProvider() string {
	if o == nil || o.OnBoardingOtpProvider == nil {
		var ret string
		return ret
	}
	return *o.OnBoardingOtpProvider
}

// GetOnBoardingOtpProviderOk returns a tuple with the OnBoardingOtpProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderAllOf) GetOnBoardingOtpProviderOk() (*string, bool) {
	if o == nil || o.OnBoardingOtpProvider == nil {
		return nil, false
	}
	return o.OnBoardingOtpProvider, true
}

// HasOnBoardingOtpProvider returns a boolean if a field has been set.
func (o *IdentityProviderAllOf) HasOnBoardingOtpProvider() bool {
	if o != nil && o.OnBoardingOtpProvider != nil {
		return true
	}

	return false
}

// SetOnBoardingOtpProvider gets a reference to the given string and assigns it to the OnBoardingOtpProvider field.
func (o *IdentityProviderAllOf) SetOnBoardingOtpProvider(v string) {
	o.OnBoardingOtpProvider = &v
}

// GetOnBoardingOtpMessage returns the OnBoardingOtpMessage field value if set, zero value otherwise.
func (o *IdentityProviderAllOf) GetOnBoardingOtpMessage() string {
	if o == nil || o.OnBoardingOtpMessage == nil {
		var ret string
		return ret
	}
	return *o.OnBoardingOtpMessage
}

// GetOnBoardingOtpMessageOk returns a tuple with the OnBoardingOtpMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderAllOf) GetOnBoardingOtpMessageOk() (*string, bool) {
	if o == nil || o.OnBoardingOtpMessage == nil {
		return nil, false
	}
	return o.OnBoardingOtpMessage, true
}

// HasOnBoardingOtpMessage returns a boolean if a field has been set.
func (o *IdentityProviderAllOf) HasOnBoardingOtpMessage() bool {
	if o != nil && o.OnBoardingOtpMessage != nil {
		return true
	}

	return false
}

// SetOnBoardingOtpMessage gets a reference to the given string and assigns it to the OnBoardingOtpMessage field.
func (o *IdentityProviderAllOf) SetOnBoardingOtpMessage(v string) {
	o.OnBoardingOtpMessage = &v
}

// GetInactivityTimeoutMinutes returns the InactivityTimeoutMinutes field value if set, zero value otherwise.
func (o *IdentityProviderAllOf) GetInactivityTimeoutMinutes() int32 {
	if o == nil || o.InactivityTimeoutMinutes == nil {
		var ret int32
		return ret
	}
	return *o.InactivityTimeoutMinutes
}

// GetInactivityTimeoutMinutesOk returns a tuple with the InactivityTimeoutMinutes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderAllOf) GetInactivityTimeoutMinutesOk() (*int32, bool) {
	if o == nil || o.InactivityTimeoutMinutes == nil {
		return nil, false
	}
	return o.InactivityTimeoutMinutes, true
}

// HasInactivityTimeoutMinutes returns a boolean if a field has been set.
func (o *IdentityProviderAllOf) HasInactivityTimeoutMinutes() bool {
	if o != nil && o.InactivityTimeoutMinutes != nil {
		return true
	}

	return false
}

// SetInactivityTimeoutMinutes gets a reference to the given int32 and assigns it to the InactivityTimeoutMinutes field.
func (o *IdentityProviderAllOf) SetInactivityTimeoutMinutes(v int32) {
	o.InactivityTimeoutMinutes = &v
}

// GetIpPoolV4 returns the IpPoolV4 field value if set, zero value otherwise.
func (o *IdentityProviderAllOf) GetIpPoolV4() string {
	if o == nil || o.IpPoolV4 == nil {
		var ret string
		return ret
	}
	return *o.IpPoolV4
}

// GetIpPoolV4Ok returns a tuple with the IpPoolV4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderAllOf) GetIpPoolV4Ok() (*string, bool) {
	if o == nil || o.IpPoolV4 == nil {
		return nil, false
	}
	return o.IpPoolV4, true
}

// HasIpPoolV4 returns a boolean if a field has been set.
func (o *IdentityProviderAllOf) HasIpPoolV4() bool {
	if o != nil && o.IpPoolV4 != nil {
		return true
	}

	return false
}

// SetIpPoolV4 gets a reference to the given string and assigns it to the IpPoolV4 field.
func (o *IdentityProviderAllOf) SetIpPoolV4(v string) {
	o.IpPoolV4 = &v
}

// GetIpPoolV6 returns the IpPoolV6 field value if set, zero value otherwise.
func (o *IdentityProviderAllOf) GetIpPoolV6() string {
	if o == nil || o.IpPoolV6 == nil {
		var ret string
		return ret
	}
	return *o.IpPoolV6
}

// GetIpPoolV6Ok returns a tuple with the IpPoolV6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderAllOf) GetIpPoolV6Ok() (*string, bool) {
	if o == nil || o.IpPoolV6 == nil {
		return nil, false
	}
	return o.IpPoolV6, true
}

// HasIpPoolV6 returns a boolean if a field has been set.
func (o *IdentityProviderAllOf) HasIpPoolV6() bool {
	if o != nil && o.IpPoolV6 != nil {
		return true
	}

	return false
}

// SetIpPoolV6 gets a reference to the given string and assigns it to the IpPoolV6 field.
func (o *IdentityProviderAllOf) SetIpPoolV6(v string) {
	o.IpPoolV6 = &v
}

// GetDnsServers returns the DnsServers field value if set, zero value otherwise.
func (o *IdentityProviderAllOf) GetDnsServers() []string {
	if o == nil || o.DnsServers == nil {
		var ret []string
		return ret
	}
	return *o.DnsServers
}

// GetDnsServersOk returns a tuple with the DnsServers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderAllOf) GetDnsServersOk() (*[]string, bool) {
	if o == nil || o.DnsServers == nil {
		return nil, false
	}
	return o.DnsServers, true
}

// HasDnsServers returns a boolean if a field has been set.
func (o *IdentityProviderAllOf) HasDnsServers() bool {
	if o != nil && o.DnsServers != nil {
		return true
	}

	return false
}

// SetDnsServers gets a reference to the given []string and assigns it to the DnsServers field.
func (o *IdentityProviderAllOf) SetDnsServers(v []string) {
	o.DnsServers = &v
}

// GetDnsSearchDomains returns the DnsSearchDomains field value if set, zero value otherwise.
func (o *IdentityProviderAllOf) GetDnsSearchDomains() []string {
	if o == nil || o.DnsSearchDomains == nil {
		var ret []string
		return ret
	}
	return *o.DnsSearchDomains
}

// GetDnsSearchDomainsOk returns a tuple with the DnsSearchDomains field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderAllOf) GetDnsSearchDomainsOk() (*[]string, bool) {
	if o == nil || o.DnsSearchDomains == nil {
		return nil, false
	}
	return o.DnsSearchDomains, true
}

// HasDnsSearchDomains returns a boolean if a field has been set.
func (o *IdentityProviderAllOf) HasDnsSearchDomains() bool {
	if o != nil && o.DnsSearchDomains != nil {
		return true
	}

	return false
}

// SetDnsSearchDomains gets a reference to the given []string and assigns it to the DnsSearchDomains field.
func (o *IdentityProviderAllOf) SetDnsSearchDomains(v []string) {
	o.DnsSearchDomains = &v
}

// GetBlockLocalDnsRequests returns the BlockLocalDnsRequests field value if set, zero value otherwise.
func (o *IdentityProviderAllOf) GetBlockLocalDnsRequests() bool {
	if o == nil || o.BlockLocalDnsRequests == nil {
		var ret bool
		return ret
	}
	return *o.BlockLocalDnsRequests
}

// GetBlockLocalDnsRequestsOk returns a tuple with the BlockLocalDnsRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderAllOf) GetBlockLocalDnsRequestsOk() (*bool, bool) {
	if o == nil || o.BlockLocalDnsRequests == nil {
		return nil, false
	}
	return o.BlockLocalDnsRequests, true
}

// HasBlockLocalDnsRequests returns a boolean if a field has been set.
func (o *IdentityProviderAllOf) HasBlockLocalDnsRequests() bool {
	if o != nil && o.BlockLocalDnsRequests != nil {
		return true
	}

	return false
}

// SetBlockLocalDnsRequests gets a reference to the given bool and assigns it to the BlockLocalDnsRequests field.
func (o *IdentityProviderAllOf) SetBlockLocalDnsRequests(v bool) {
	o.BlockLocalDnsRequests = &v
}

// GetClaimMappings returns the ClaimMappings field value if set, zero value otherwise.
func (o *IdentityProviderAllOf) GetClaimMappings() []map[string]interface{} {
	if o == nil || o.ClaimMappings == nil {
		var ret []map[string]interface{}
		return ret
	}
	return *o.ClaimMappings
}

// GetClaimMappingsOk returns a tuple with the ClaimMappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderAllOf) GetClaimMappingsOk() (*[]map[string]interface{}, bool) {
	if o == nil || o.ClaimMappings == nil {
		return nil, false
	}
	return o.ClaimMappings, true
}

// HasClaimMappings returns a boolean if a field has been set.
func (o *IdentityProviderAllOf) HasClaimMappings() bool {
	if o != nil && o.ClaimMappings != nil {
		return true
	}

	return false
}

// SetClaimMappings gets a reference to the given []map[string]interface{} and assigns it to the ClaimMappings field.
func (o *IdentityProviderAllOf) SetClaimMappings(v []map[string]interface{}) {
	o.ClaimMappings = &v
}

// GetOnDemandClaimMappings returns the OnDemandClaimMappings field value if set, zero value otherwise.
func (o *IdentityProviderAllOf) GetOnDemandClaimMappings() []map[string]interface{} {
	if o == nil || o.OnDemandClaimMappings == nil {
		var ret []map[string]interface{}
		return ret
	}
	return *o.OnDemandClaimMappings
}

// GetOnDemandClaimMappingsOk returns a tuple with the OnDemandClaimMappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderAllOf) GetOnDemandClaimMappingsOk() (*[]map[string]interface{}, bool) {
	if o == nil || o.OnDemandClaimMappings == nil {
		return nil, false
	}
	return o.OnDemandClaimMappings, true
}

// HasOnDemandClaimMappings returns a boolean if a field has been set.
func (o *IdentityProviderAllOf) HasOnDemandClaimMappings() bool {
	if o != nil && o.OnDemandClaimMappings != nil {
		return true
	}

	return false
}

// SetOnDemandClaimMappings gets a reference to the given []map[string]interface{} and assigns it to the OnDemandClaimMappings field.
func (o *IdentityProviderAllOf) SetOnDemandClaimMappings(v []map[string]interface{}) {
	o.OnDemandClaimMappings = &v
}

func (o IdentityProviderAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}
	if o.Default != nil {
		toSerialize["default"] = o.Default
	}
	if o.ClientProvider != nil {
		toSerialize["clientProvider"] = o.ClientProvider
	}
	if o.AdminProvider != nil {
		toSerialize["adminProvider"] = o.AdminProvider
	}
	if o.OnBoarding2FA != nil {
		toSerialize["onBoarding2FA"] = o.OnBoarding2FA
	}
	if o.OnBoardingType != nil {
		toSerialize["onBoardingType"] = o.OnBoardingType
	}
	if o.OnBoardingOtpProvider != nil {
		toSerialize["onBoardingOtpProvider"] = o.OnBoardingOtpProvider
	}
	if o.OnBoardingOtpMessage != nil {
		toSerialize["onBoardingOtpMessage"] = o.OnBoardingOtpMessage
	}
	if o.InactivityTimeoutMinutes != nil {
		toSerialize["inactivityTimeoutMinutes"] = o.InactivityTimeoutMinutes
	}
	if o.IpPoolV4 != nil {
		toSerialize["ipPoolV4"] = o.IpPoolV4
	}
	if o.IpPoolV6 != nil {
		toSerialize["ipPoolV6"] = o.IpPoolV6
	}
	if o.DnsServers != nil {
		toSerialize["dnsServers"] = o.DnsServers
	}
	if o.DnsSearchDomains != nil {
		toSerialize["dnsSearchDomains"] = o.DnsSearchDomains
	}
	if o.BlockLocalDnsRequests != nil {
		toSerialize["blockLocalDnsRequests"] = o.BlockLocalDnsRequests
	}
	if o.ClaimMappings != nil {
		toSerialize["claimMappings"] = o.ClaimMappings
	}
	if o.OnDemandClaimMappings != nil {
		toSerialize["onDemandClaimMappings"] = o.OnDemandClaimMappings
	}
	return json.Marshal(toSerialize)
}

type NullableIdentityProviderAllOf struct {
	value *IdentityProviderAllOf
	isSet bool
}

func (v NullableIdentityProviderAllOf) Get() *IdentityProviderAllOf {
	return v.value
}

func (v *NullableIdentityProviderAllOf) Set(val *IdentityProviderAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityProviderAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityProviderAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityProviderAllOf(val *IdentityProviderAllOf) *NullableIdentityProviderAllOf {
	return &NullableIdentityProviderAllOf{value: val, isSet: true}
}

func (v NullableIdentityProviderAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityProviderAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
