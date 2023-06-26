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

// PolicyAllOf Represents a Policy.
type PolicyAllOf struct {
	// If true, the Policy will be disregarded during authorization.
	Disabled *bool `json:"disabled,omitempty"`
	// A JavaScript expression that returns boolean. Criteria Scripts may be used by calling them as functions.
	Expression string `json:"expression"`
	// Type of the Policy. The assigned type will be enforced by not allowing enabling other types of features on the Policy.
	Type *string `json:"type,omitempty"`
	// List of Entitlement IDs in this Policy.
	Entitlements []string `json:"entitlements,omitempty"`
	// List of Entitlement tags in this Policy.
	EntitlementLinks []string `json:"entitlementLinks,omitempty"`
	// List of Ringfence Rule IDs in this Policy.
	RingfenceRules []string `json:"ringfenceRules,omitempty"`
	// List of Ringfence Rule tags in this Policy.
	RingfenceRuleLinks []string `json:"ringfenceRuleLinks,omitempty"`
	// Will enable Tamper Proofing on desktop clients which will make sure the routes and ringfence configurations are not changed. Will be set to true if it's not provided and the type is \"Device\".
	TamperProofing *bool `json:"tamperProofing,omitempty"`
	// Site ID where all the Entitlements of this Policy must be deployed. This overrides Entitlement's own Site and to be used only in specific network layouts. Otherwise the assigned site on individual Entitlements will be used.
	OverrideSite *string `json:"overrideSite,omitempty"`
	// The path of a claim that contains the UUID of an override site. It should be defined as \"claims.xxx.xxx\" or \"claims.xxx.xxx.xxx\".
	OverrideSiteClaim *string `json:"overrideSiteClaim,omitempty"`
	// Overrides the Entitlements Site according to location of the client and Sites where this feature is enabled.
	OverrideNearestSite *bool `json:"overrideNearestSite,omitempty"`
	// The Entitlements in this Policy will be available in the fallback Sites if the corresponding Sites are configured accordingly.
	ApplyFallbackSite   *bool                           `json:"applyFallbackSite,omitempty"`
	ProxyAutoConfig     *PolicyAllOfProxyAutoConfig     `json:"proxyAutoConfig,omitempty"`
	TrustedNetworkCheck *PolicyAllOfTrustedNetworkCheck `json:"trustedNetworkCheck,omitempty"`
	// List of domain names with DNS server IPs that the Client should be using.
	DnsSettings           []PolicyAllOfDnsSettings          `json:"dnsSettings,omitempty"`
	ClientSettings        *PolicyAllOfClientSettings        `json:"clientSettings,omitempty"`
	ClientProfileSettings *PolicyAllOfClientProfileSettings `json:"clientProfileSettings,omitempty"`
	// Set custom help URL on Client UI.
	CustomClientHelpUrl *string `json:"customClientHelpUrl,omitempty"`
	// List of Administrative Role IDs in this Policy.
	AdministrativeRoles []string `json:"administrativeRoles,omitempty"`
}

// NewPolicyAllOf instantiates a new PolicyAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicyAllOf(expression string) *PolicyAllOf {
	this := PolicyAllOf{}
	var disabled bool = false
	this.Disabled = &disabled
	this.Expression = expression
	var type_ string = "Mixed"
	this.Type = &type_
	return &this
}

// NewPolicyAllOfWithDefaults instantiates a new PolicyAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicyAllOfWithDefaults() *PolicyAllOf {
	this := PolicyAllOf{}
	var disabled bool = false
	this.Disabled = &disabled
	var type_ string = "Mixed"
	this.Type = &type_
	return &this
}

// GetDisabled returns the Disabled field value if set, zero value otherwise.
func (o *PolicyAllOf) GetDisabled() bool {
	if o == nil || o.Disabled == nil {
		var ret bool
		return ret
	}
	return *o.Disabled
}

// GetDisabledOk returns a tuple with the Disabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyAllOf) GetDisabledOk() (*bool, bool) {
	if o == nil || o.Disabled == nil {
		return nil, false
	}
	return o.Disabled, true
}

// HasDisabled returns a boolean if a field has been set.
func (o *PolicyAllOf) HasDisabled() bool {
	if o != nil && o.Disabled != nil {
		return true
	}

	return false
}

// SetDisabled gets a reference to the given bool and assigns it to the Disabled field.
func (o *PolicyAllOf) SetDisabled(v bool) {
	o.Disabled = &v
}

// GetExpression returns the Expression field value
func (o *PolicyAllOf) GetExpression() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value
// and a boolean to check if the value has been set.
func (o *PolicyAllOf) GetExpressionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Expression, true
}

// SetExpression sets field value
func (o *PolicyAllOf) SetExpression(v string) {
	o.Expression = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PolicyAllOf) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyAllOf) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PolicyAllOf) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *PolicyAllOf) SetType(v string) {
	o.Type = &v
}

// GetEntitlements returns the Entitlements field value if set, zero value otherwise.
func (o *PolicyAllOf) GetEntitlements() []string {
	if o == nil || o.Entitlements == nil {
		var ret []string
		return ret
	}
	return o.Entitlements
}

// GetEntitlementsOk returns a tuple with the Entitlements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyAllOf) GetEntitlementsOk() ([]string, bool) {
	if o == nil || o.Entitlements == nil {
		return nil, false
	}
	return o.Entitlements, true
}

// HasEntitlements returns a boolean if a field has been set.
func (o *PolicyAllOf) HasEntitlements() bool {
	if o != nil && o.Entitlements != nil {
		return true
	}

	return false
}

// SetEntitlements gets a reference to the given []string and assigns it to the Entitlements field.
func (o *PolicyAllOf) SetEntitlements(v []string) {
	o.Entitlements = v
}

// GetEntitlementLinks returns the EntitlementLinks field value if set, zero value otherwise.
func (o *PolicyAllOf) GetEntitlementLinks() []string {
	if o == nil || o.EntitlementLinks == nil {
		var ret []string
		return ret
	}
	return o.EntitlementLinks
}

// GetEntitlementLinksOk returns a tuple with the EntitlementLinks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyAllOf) GetEntitlementLinksOk() ([]string, bool) {
	if o == nil || o.EntitlementLinks == nil {
		return nil, false
	}
	return o.EntitlementLinks, true
}

// HasEntitlementLinks returns a boolean if a field has been set.
func (o *PolicyAllOf) HasEntitlementLinks() bool {
	if o != nil && o.EntitlementLinks != nil {
		return true
	}

	return false
}

// SetEntitlementLinks gets a reference to the given []string and assigns it to the EntitlementLinks field.
func (o *PolicyAllOf) SetEntitlementLinks(v []string) {
	o.EntitlementLinks = v
}

// GetRingfenceRules returns the RingfenceRules field value if set, zero value otherwise.
func (o *PolicyAllOf) GetRingfenceRules() []string {
	if o == nil || o.RingfenceRules == nil {
		var ret []string
		return ret
	}
	return o.RingfenceRules
}

// GetRingfenceRulesOk returns a tuple with the RingfenceRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyAllOf) GetRingfenceRulesOk() ([]string, bool) {
	if o == nil || o.RingfenceRules == nil {
		return nil, false
	}
	return o.RingfenceRules, true
}

// HasRingfenceRules returns a boolean if a field has been set.
func (o *PolicyAllOf) HasRingfenceRules() bool {
	if o != nil && o.RingfenceRules != nil {
		return true
	}

	return false
}

// SetRingfenceRules gets a reference to the given []string and assigns it to the RingfenceRules field.
func (o *PolicyAllOf) SetRingfenceRules(v []string) {
	o.RingfenceRules = v
}

// GetRingfenceRuleLinks returns the RingfenceRuleLinks field value if set, zero value otherwise.
func (o *PolicyAllOf) GetRingfenceRuleLinks() []string {
	if o == nil || o.RingfenceRuleLinks == nil {
		var ret []string
		return ret
	}
	return o.RingfenceRuleLinks
}

// GetRingfenceRuleLinksOk returns a tuple with the RingfenceRuleLinks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyAllOf) GetRingfenceRuleLinksOk() ([]string, bool) {
	if o == nil || o.RingfenceRuleLinks == nil {
		return nil, false
	}
	return o.RingfenceRuleLinks, true
}

// HasRingfenceRuleLinks returns a boolean if a field has been set.
func (o *PolicyAllOf) HasRingfenceRuleLinks() bool {
	if o != nil && o.RingfenceRuleLinks != nil {
		return true
	}

	return false
}

// SetRingfenceRuleLinks gets a reference to the given []string and assigns it to the RingfenceRuleLinks field.
func (o *PolicyAllOf) SetRingfenceRuleLinks(v []string) {
	o.RingfenceRuleLinks = v
}

// GetTamperProofing returns the TamperProofing field value if set, zero value otherwise.
func (o *PolicyAllOf) GetTamperProofing() bool {
	if o == nil || o.TamperProofing == nil {
		var ret bool
		return ret
	}
	return *o.TamperProofing
}

// GetTamperProofingOk returns a tuple with the TamperProofing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyAllOf) GetTamperProofingOk() (*bool, bool) {
	if o == nil || o.TamperProofing == nil {
		return nil, false
	}
	return o.TamperProofing, true
}

// HasTamperProofing returns a boolean if a field has been set.
func (o *PolicyAllOf) HasTamperProofing() bool {
	if o != nil && o.TamperProofing != nil {
		return true
	}

	return false
}

// SetTamperProofing gets a reference to the given bool and assigns it to the TamperProofing field.
func (o *PolicyAllOf) SetTamperProofing(v bool) {
	o.TamperProofing = &v
}

// GetOverrideSite returns the OverrideSite field value if set, zero value otherwise.
func (o *PolicyAllOf) GetOverrideSite() string {
	if o == nil || o.OverrideSite == nil {
		var ret string
		return ret
	}
	return *o.OverrideSite
}

// GetOverrideSiteOk returns a tuple with the OverrideSite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyAllOf) GetOverrideSiteOk() (*string, bool) {
	if o == nil || o.OverrideSite == nil {
		return nil, false
	}
	return o.OverrideSite, true
}

// HasOverrideSite returns a boolean if a field has been set.
func (o *PolicyAllOf) HasOverrideSite() bool {
	if o != nil && o.OverrideSite != nil {
		return true
	}

	return false
}

// SetOverrideSite gets a reference to the given string and assigns it to the OverrideSite field.
func (o *PolicyAllOf) SetOverrideSite(v string) {
	o.OverrideSite = &v
}

// GetOverrideSiteClaim returns the OverrideSiteClaim field value if set, zero value otherwise.
func (o *PolicyAllOf) GetOverrideSiteClaim() string {
	if o == nil || o.OverrideSiteClaim == nil {
		var ret string
		return ret
	}
	return *o.OverrideSiteClaim
}

// GetOverrideSiteClaimOk returns a tuple with the OverrideSiteClaim field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyAllOf) GetOverrideSiteClaimOk() (*string, bool) {
	if o == nil || o.OverrideSiteClaim == nil {
		return nil, false
	}
	return o.OverrideSiteClaim, true
}

// HasOverrideSiteClaim returns a boolean if a field has been set.
func (o *PolicyAllOf) HasOverrideSiteClaim() bool {
	if o != nil && o.OverrideSiteClaim != nil {
		return true
	}

	return false
}

// SetOverrideSiteClaim gets a reference to the given string and assigns it to the OverrideSiteClaim field.
func (o *PolicyAllOf) SetOverrideSiteClaim(v string) {
	o.OverrideSiteClaim = &v
}

// GetOverrideNearestSite returns the OverrideNearestSite field value if set, zero value otherwise.
func (o *PolicyAllOf) GetOverrideNearestSite() bool {
	if o == nil || o.OverrideNearestSite == nil {
		var ret bool
		return ret
	}
	return *o.OverrideNearestSite
}

// GetOverrideNearestSiteOk returns a tuple with the OverrideNearestSite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyAllOf) GetOverrideNearestSiteOk() (*bool, bool) {
	if o == nil || o.OverrideNearestSite == nil {
		return nil, false
	}
	return o.OverrideNearestSite, true
}

// HasOverrideNearestSite returns a boolean if a field has been set.
func (o *PolicyAllOf) HasOverrideNearestSite() bool {
	if o != nil && o.OverrideNearestSite != nil {
		return true
	}

	return false
}

// SetOverrideNearestSite gets a reference to the given bool and assigns it to the OverrideNearestSite field.
func (o *PolicyAllOf) SetOverrideNearestSite(v bool) {
	o.OverrideNearestSite = &v
}

// GetApplyFallbackSite returns the ApplyFallbackSite field value if set, zero value otherwise.
func (o *PolicyAllOf) GetApplyFallbackSite() bool {
	if o == nil || o.ApplyFallbackSite == nil {
		var ret bool
		return ret
	}
	return *o.ApplyFallbackSite
}

// GetApplyFallbackSiteOk returns a tuple with the ApplyFallbackSite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyAllOf) GetApplyFallbackSiteOk() (*bool, bool) {
	if o == nil || o.ApplyFallbackSite == nil {
		return nil, false
	}
	return o.ApplyFallbackSite, true
}

// HasApplyFallbackSite returns a boolean if a field has been set.
func (o *PolicyAllOf) HasApplyFallbackSite() bool {
	if o != nil && o.ApplyFallbackSite != nil {
		return true
	}

	return false
}

// SetApplyFallbackSite gets a reference to the given bool and assigns it to the ApplyFallbackSite field.
func (o *PolicyAllOf) SetApplyFallbackSite(v bool) {
	o.ApplyFallbackSite = &v
}

// GetProxyAutoConfig returns the ProxyAutoConfig field value if set, zero value otherwise.
func (o *PolicyAllOf) GetProxyAutoConfig() PolicyAllOfProxyAutoConfig {
	if o == nil || o.ProxyAutoConfig == nil {
		var ret PolicyAllOfProxyAutoConfig
		return ret
	}
	return *o.ProxyAutoConfig
}

// GetProxyAutoConfigOk returns a tuple with the ProxyAutoConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyAllOf) GetProxyAutoConfigOk() (*PolicyAllOfProxyAutoConfig, bool) {
	if o == nil || o.ProxyAutoConfig == nil {
		return nil, false
	}
	return o.ProxyAutoConfig, true
}

// HasProxyAutoConfig returns a boolean if a field has been set.
func (o *PolicyAllOf) HasProxyAutoConfig() bool {
	if o != nil && o.ProxyAutoConfig != nil {
		return true
	}

	return false
}

// SetProxyAutoConfig gets a reference to the given PolicyAllOfProxyAutoConfig and assigns it to the ProxyAutoConfig field.
func (o *PolicyAllOf) SetProxyAutoConfig(v PolicyAllOfProxyAutoConfig) {
	o.ProxyAutoConfig = &v
}

// GetTrustedNetworkCheck returns the TrustedNetworkCheck field value if set, zero value otherwise.
func (o *PolicyAllOf) GetTrustedNetworkCheck() PolicyAllOfTrustedNetworkCheck {
	if o == nil || o.TrustedNetworkCheck == nil {
		var ret PolicyAllOfTrustedNetworkCheck
		return ret
	}
	return *o.TrustedNetworkCheck
}

// GetTrustedNetworkCheckOk returns a tuple with the TrustedNetworkCheck field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyAllOf) GetTrustedNetworkCheckOk() (*PolicyAllOfTrustedNetworkCheck, bool) {
	if o == nil || o.TrustedNetworkCheck == nil {
		return nil, false
	}
	return o.TrustedNetworkCheck, true
}

// HasTrustedNetworkCheck returns a boolean if a field has been set.
func (o *PolicyAllOf) HasTrustedNetworkCheck() bool {
	if o != nil && o.TrustedNetworkCheck != nil {
		return true
	}

	return false
}

// SetTrustedNetworkCheck gets a reference to the given PolicyAllOfTrustedNetworkCheck and assigns it to the TrustedNetworkCheck field.
func (o *PolicyAllOf) SetTrustedNetworkCheck(v PolicyAllOfTrustedNetworkCheck) {
	o.TrustedNetworkCheck = &v
}

// GetDnsSettings returns the DnsSettings field value if set, zero value otherwise.
func (o *PolicyAllOf) GetDnsSettings() []PolicyAllOfDnsSettings {
	if o == nil || o.DnsSettings == nil {
		var ret []PolicyAllOfDnsSettings
		return ret
	}
	return o.DnsSettings
}

// GetDnsSettingsOk returns a tuple with the DnsSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyAllOf) GetDnsSettingsOk() ([]PolicyAllOfDnsSettings, bool) {
	if o == nil || o.DnsSettings == nil {
		return nil, false
	}
	return o.DnsSettings, true
}

// HasDnsSettings returns a boolean if a field has been set.
func (o *PolicyAllOf) HasDnsSettings() bool {
	if o != nil && o.DnsSettings != nil {
		return true
	}

	return false
}

// SetDnsSettings gets a reference to the given []PolicyAllOfDnsSettings and assigns it to the DnsSettings field.
func (o *PolicyAllOf) SetDnsSettings(v []PolicyAllOfDnsSettings) {
	o.DnsSettings = v
}

// GetClientSettings returns the ClientSettings field value if set, zero value otherwise.
func (o *PolicyAllOf) GetClientSettings() PolicyAllOfClientSettings {
	if o == nil || o.ClientSettings == nil {
		var ret PolicyAllOfClientSettings
		return ret
	}
	return *o.ClientSettings
}

// GetClientSettingsOk returns a tuple with the ClientSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyAllOf) GetClientSettingsOk() (*PolicyAllOfClientSettings, bool) {
	if o == nil || o.ClientSettings == nil {
		return nil, false
	}
	return o.ClientSettings, true
}

// HasClientSettings returns a boolean if a field has been set.
func (o *PolicyAllOf) HasClientSettings() bool {
	if o != nil && o.ClientSettings != nil {
		return true
	}

	return false
}

// SetClientSettings gets a reference to the given PolicyAllOfClientSettings and assigns it to the ClientSettings field.
func (o *PolicyAllOf) SetClientSettings(v PolicyAllOfClientSettings) {
	o.ClientSettings = &v
}

// GetClientProfileSettings returns the ClientProfileSettings field value if set, zero value otherwise.
func (o *PolicyAllOf) GetClientProfileSettings() PolicyAllOfClientProfileSettings {
	if o == nil || o.ClientProfileSettings == nil {
		var ret PolicyAllOfClientProfileSettings
		return ret
	}
	return *o.ClientProfileSettings
}

// GetClientProfileSettingsOk returns a tuple with the ClientProfileSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyAllOf) GetClientProfileSettingsOk() (*PolicyAllOfClientProfileSettings, bool) {
	if o == nil || o.ClientProfileSettings == nil {
		return nil, false
	}
	return o.ClientProfileSettings, true
}

// HasClientProfileSettings returns a boolean if a field has been set.
func (o *PolicyAllOf) HasClientProfileSettings() bool {
	if o != nil && o.ClientProfileSettings != nil {
		return true
	}

	return false
}

// SetClientProfileSettings gets a reference to the given PolicyAllOfClientProfileSettings and assigns it to the ClientProfileSettings field.
func (o *PolicyAllOf) SetClientProfileSettings(v PolicyAllOfClientProfileSettings) {
	o.ClientProfileSettings = &v
}

// GetCustomClientHelpUrl returns the CustomClientHelpUrl field value if set, zero value otherwise.
func (o *PolicyAllOf) GetCustomClientHelpUrl() string {
	if o == nil || o.CustomClientHelpUrl == nil {
		var ret string
		return ret
	}
	return *o.CustomClientHelpUrl
}

// GetCustomClientHelpUrlOk returns a tuple with the CustomClientHelpUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyAllOf) GetCustomClientHelpUrlOk() (*string, bool) {
	if o == nil || o.CustomClientHelpUrl == nil {
		return nil, false
	}
	return o.CustomClientHelpUrl, true
}

// HasCustomClientHelpUrl returns a boolean if a field has been set.
func (o *PolicyAllOf) HasCustomClientHelpUrl() bool {
	if o != nil && o.CustomClientHelpUrl != nil {
		return true
	}

	return false
}

// SetCustomClientHelpUrl gets a reference to the given string and assigns it to the CustomClientHelpUrl field.
func (o *PolicyAllOf) SetCustomClientHelpUrl(v string) {
	o.CustomClientHelpUrl = &v
}

// GetAdministrativeRoles returns the AdministrativeRoles field value if set, zero value otherwise.
func (o *PolicyAllOf) GetAdministrativeRoles() []string {
	if o == nil || o.AdministrativeRoles == nil {
		var ret []string
		return ret
	}
	return o.AdministrativeRoles
}

// GetAdministrativeRolesOk returns a tuple with the AdministrativeRoles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyAllOf) GetAdministrativeRolesOk() ([]string, bool) {
	if o == nil || o.AdministrativeRoles == nil {
		return nil, false
	}
	return o.AdministrativeRoles, true
}

// HasAdministrativeRoles returns a boolean if a field has been set.
func (o *PolicyAllOf) HasAdministrativeRoles() bool {
	if o != nil && o.AdministrativeRoles != nil {
		return true
	}

	return false
}

// SetAdministrativeRoles gets a reference to the given []string and assigns it to the AdministrativeRoles field.
func (o *PolicyAllOf) SetAdministrativeRoles(v []string) {
	o.AdministrativeRoles = v
}

func (o PolicyAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Disabled != nil {
		toSerialize["disabled"] = o.Disabled
	}
	if true {
		toSerialize["expression"] = o.Expression
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Entitlements != nil {
		toSerialize["entitlements"] = o.Entitlements
	}
	if o.EntitlementLinks != nil {
		toSerialize["entitlementLinks"] = o.EntitlementLinks
	}
	if o.RingfenceRules != nil {
		toSerialize["ringfenceRules"] = o.RingfenceRules
	}
	if o.RingfenceRuleLinks != nil {
		toSerialize["ringfenceRuleLinks"] = o.RingfenceRuleLinks
	}
	if o.TamperProofing != nil {
		toSerialize["tamperProofing"] = o.TamperProofing
	}
	if o.OverrideSite != nil {
		toSerialize["overrideSite"] = o.OverrideSite
	}
	if o.OverrideSiteClaim != nil {
		toSerialize["overrideSiteClaim"] = o.OverrideSiteClaim
	}
	if o.OverrideNearestSite != nil {
		toSerialize["overrideNearestSite"] = o.OverrideNearestSite
	}
	if o.ApplyFallbackSite != nil {
		toSerialize["applyFallbackSite"] = o.ApplyFallbackSite
	}
	if o.ProxyAutoConfig != nil {
		toSerialize["proxyAutoConfig"] = o.ProxyAutoConfig
	}
	if o.TrustedNetworkCheck != nil {
		toSerialize["trustedNetworkCheck"] = o.TrustedNetworkCheck
	}
	if o.DnsSettings != nil {
		toSerialize["dnsSettings"] = o.DnsSettings
	}
	if o.ClientSettings != nil {
		toSerialize["clientSettings"] = o.ClientSettings
	}
	if o.ClientProfileSettings != nil {
		toSerialize["clientProfileSettings"] = o.ClientProfileSettings
	}
	if o.CustomClientHelpUrl != nil {
		toSerialize["customClientHelpUrl"] = o.CustomClientHelpUrl
	}
	if o.AdministrativeRoles != nil {
		toSerialize["administrativeRoles"] = o.AdministrativeRoles
	}
	return json.Marshal(toSerialize)
}

type NullablePolicyAllOf struct {
	value *PolicyAllOf
	isSet bool
}

func (v NullablePolicyAllOf) Get() *PolicyAllOf {
	return v.value
}

func (v *NullablePolicyAllOf) Set(val *PolicyAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyAllOf(val *PolicyAllOf) *NullablePolicyAllOf {
	return &NullablePolicyAllOf{value: val, isSet: true}
}

func (v NullablePolicyAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
