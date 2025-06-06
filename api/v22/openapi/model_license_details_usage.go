/*
Appgate SDP Controller REST API

# About   This specification documents the REST API calls for the Appgate SDP Controller.    Please refer to the REST API chapter in the manual or contact Appgate support with any questions about   this functionality. # Getting Started   Requirements for API scripting:   - Access to the Admin/API TLS Connection (default port 8443) of a Controller appliance.     (https://sdphelp.appgate.com/adminguide/appliance-function-configure.html?anchor=admin-api)   - An API user with relevant permissions.     (https://sdphelp.appgate.com/adminguide/administrative-roles-configure.html)   - In order to use the simple login API, Admin MFA must be disabled or the API user must be excluded.     (https://sdphelp.appgate.com/adminguide/mfa-for-admins.html) # Base path   HTTPS requests must be sent to the Admin Interface hostname and port, with **_/admin** path.    For example: **https://appgate.company.com:8443/admin**    All requests must have the **Accept** header as:    **application/vnd.appgate.peer-v22+json**    An exception is made for the **_/admin/version** endpoint which instead expects an **application/json** Accept header. # API Conventions   API conventions are  important to understand and follow strictly.    - While updating objects (via PUT), entire object must be sent with all fields.     - For example, in order to add a remedy method to the condition below:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": []       }       ```     - send the entire object with updated and non-updated fields:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": [{\"type\": \"DisplayMessage\", \"message\": \"test message\"}]       }       ```    - In case Controller returns an error (non-2xx HTTP status code), response body is JSON.     The \"message\" field contains information about the error.     HTTP 422 \"Unprocessable Entity\" has extra `errors` field to list all the issues with specific fields.    - Empty string (\"\") is considered a different value than \"null\" or field being omitted from JSON.     Omitting the field is recommended if no value is intended.     Empty string (\"\") will be almost always rejected as invalid value.    - There are common pattern between many objects:     - **Configuration Objects**: There are many objects with common fields, namely \"id\", \"name\", \"notes\", \"created\"       and \"updated\". These entities are listed, queried, created, updated and deleted in a similar fashion.     - **Distinguished Name**: Users and Devices are identified with what is called Distinguished Names, as used in        LDAP. The distinguished format that identifies a device and a user combination is        \"CN=\\<Device ID\\>,CN=\\<username\\>,OU=\\<Identity Provider Name\\>\". Some objects have the        \"userDistinguishedName\" field, which does not include the CN for Device ID.        This identifies a user on every device.

API version: API version 22.4
Contact: appgatesdp.support@appgate.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// LicenseDetailsUsage License usage information.
type LicenseDetailsUsage struct {
	// The amount of licensed users in the system at present.
	Users *float32 `json:"users,omitempty"`
	// The amount of licensed portal users in the system at present.
	PortalUsers *float32 `json:"portalUsers,omitempty"`
	// The amount of licensed service users in the system at present.
	ServiceUsers *float32 `json:"serviceUsers,omitempty"`
	// The amount of Sites in the system at present.
	Sites *float32 `json:"sites,omitempty"`
	// The amount of grouped Connector clients in the system at present.
	ConnectorGroups *float32 `json:"connectorGroups,omitempty"`
	// The amount of licensed multi-user clients in the system at present.
	MultiUserClients *float32                        `json:"multiUserClients,omitempty"`
	ClientTypes      *LicenseDetailsUsageClientTypes `json:"clientTypes,omitempty"`
	// The amount of non-built-in MFA Providers in the system at present.
	MfaProviders *float32 `json:"mfaProviders,omitempty"`
	// The amount of non-built-in Device Claim Scripts in the system at present.
	DeviceScripts *float32 `json:"deviceScripts,omitempty"`
	// The amount of non-built-in Criteria Scripts in the system at present.
	CriteriaScripts *float32 `json:"criteriaScripts,omitempty"`
	// The amount of non-built-in EntitlementScripts in the system at present.
	EntitlementScripts *float32 `json:"entitlementScripts,omitempty"`
	// The amount of non-built-in User Claim Scripts in the system at present.
	UserScripts *float32 `json:"userScripts,omitempty"`
	// The amount of fallback Sites in the system at present.
	FallbackSites *float32 `json:"fallbackSites,omitempty"`
	// The amount of nearest Sites in the system at present.
	NearestSites *float32 `json:"nearestSites,omitempty"`
	// The amount of local Sites in the system at present.
	LocalSites *float32 `json:"localSites,omitempty"`
	// The amount of Profile Groups in the system at present.
	ProfileGroups *float32 `json:"profileGroups,omitempty"`
	// The amount of Appliance Customizations in the system at present.
	ApplianceCustomizations *float32 `json:"applianceCustomizations,omitempty"`
	// The amount of non-built-in Ringfence Rules in the system at present.
	RingfenceRules *float32 `json:"ringfenceRules,omitempty"`
	// The amount of Time-based OTP Seeds in the system at present.
	OtpSeeds *float32 `json:"otpSeeds,omitempty"`
	// The amount of FIDO2 devices in the system at present.
	Fido2Devices      *float32                              `json:"fido2Devices,omitempty"`
	IdentityProviders *LicenseDetailsUsageIdentityProviders `json:"identityProviders,omitempty"`
	Policies          *LicenseDetailsUsagePolicies          `json:"policies,omitempty"`
	// Whether the risk model is utilized or not.
	RiskModel *bool `json:"riskModel,omitempty"`
}

// NewLicenseDetailsUsage instantiates a new LicenseDetailsUsage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLicenseDetailsUsage() *LicenseDetailsUsage {
	this := LicenseDetailsUsage{}
	return &this
}

// NewLicenseDetailsUsageWithDefaults instantiates a new LicenseDetailsUsage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLicenseDetailsUsageWithDefaults() *LicenseDetailsUsage {
	this := LicenseDetailsUsage{}
	return &this
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *LicenseDetailsUsage) GetUsers() float32 {
	if o == nil || o.Users == nil {
		var ret float32
		return ret
	}
	return *o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseDetailsUsage) GetUsersOk() (*float32, bool) {
	if o == nil || o.Users == nil {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *LicenseDetailsUsage) HasUsers() bool {
	if o != nil && o.Users != nil {
		return true
	}

	return false
}

// SetUsers gets a reference to the given float32 and assigns it to the Users field.
func (o *LicenseDetailsUsage) SetUsers(v float32) {
	o.Users = &v
}

// GetPortalUsers returns the PortalUsers field value if set, zero value otherwise.
func (o *LicenseDetailsUsage) GetPortalUsers() float32 {
	if o == nil || o.PortalUsers == nil {
		var ret float32
		return ret
	}
	return *o.PortalUsers
}

// GetPortalUsersOk returns a tuple with the PortalUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseDetailsUsage) GetPortalUsersOk() (*float32, bool) {
	if o == nil || o.PortalUsers == nil {
		return nil, false
	}
	return o.PortalUsers, true
}

// HasPortalUsers returns a boolean if a field has been set.
func (o *LicenseDetailsUsage) HasPortalUsers() bool {
	if o != nil && o.PortalUsers != nil {
		return true
	}

	return false
}

// SetPortalUsers gets a reference to the given float32 and assigns it to the PortalUsers field.
func (o *LicenseDetailsUsage) SetPortalUsers(v float32) {
	o.PortalUsers = &v
}

// GetServiceUsers returns the ServiceUsers field value if set, zero value otherwise.
func (o *LicenseDetailsUsage) GetServiceUsers() float32 {
	if o == nil || o.ServiceUsers == nil {
		var ret float32
		return ret
	}
	return *o.ServiceUsers
}

// GetServiceUsersOk returns a tuple with the ServiceUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseDetailsUsage) GetServiceUsersOk() (*float32, bool) {
	if o == nil || o.ServiceUsers == nil {
		return nil, false
	}
	return o.ServiceUsers, true
}

// HasServiceUsers returns a boolean if a field has been set.
func (o *LicenseDetailsUsage) HasServiceUsers() bool {
	if o != nil && o.ServiceUsers != nil {
		return true
	}

	return false
}

// SetServiceUsers gets a reference to the given float32 and assigns it to the ServiceUsers field.
func (o *LicenseDetailsUsage) SetServiceUsers(v float32) {
	o.ServiceUsers = &v
}

// GetSites returns the Sites field value if set, zero value otherwise.
func (o *LicenseDetailsUsage) GetSites() float32 {
	if o == nil || o.Sites == nil {
		var ret float32
		return ret
	}
	return *o.Sites
}

// GetSitesOk returns a tuple with the Sites field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseDetailsUsage) GetSitesOk() (*float32, bool) {
	if o == nil || o.Sites == nil {
		return nil, false
	}
	return o.Sites, true
}

// HasSites returns a boolean if a field has been set.
func (o *LicenseDetailsUsage) HasSites() bool {
	if o != nil && o.Sites != nil {
		return true
	}

	return false
}

// SetSites gets a reference to the given float32 and assigns it to the Sites field.
func (o *LicenseDetailsUsage) SetSites(v float32) {
	o.Sites = &v
}

// GetConnectorGroups returns the ConnectorGroups field value if set, zero value otherwise.
func (o *LicenseDetailsUsage) GetConnectorGroups() float32 {
	if o == nil || o.ConnectorGroups == nil {
		var ret float32
		return ret
	}
	return *o.ConnectorGroups
}

// GetConnectorGroupsOk returns a tuple with the ConnectorGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseDetailsUsage) GetConnectorGroupsOk() (*float32, bool) {
	if o == nil || o.ConnectorGroups == nil {
		return nil, false
	}
	return o.ConnectorGroups, true
}

// HasConnectorGroups returns a boolean if a field has been set.
func (o *LicenseDetailsUsage) HasConnectorGroups() bool {
	if o != nil && o.ConnectorGroups != nil {
		return true
	}

	return false
}

// SetConnectorGroups gets a reference to the given float32 and assigns it to the ConnectorGroups field.
func (o *LicenseDetailsUsage) SetConnectorGroups(v float32) {
	o.ConnectorGroups = &v
}

// GetMultiUserClients returns the MultiUserClients field value if set, zero value otherwise.
func (o *LicenseDetailsUsage) GetMultiUserClients() float32 {
	if o == nil || o.MultiUserClients == nil {
		var ret float32
		return ret
	}
	return *o.MultiUserClients
}

// GetMultiUserClientsOk returns a tuple with the MultiUserClients field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseDetailsUsage) GetMultiUserClientsOk() (*float32, bool) {
	if o == nil || o.MultiUserClients == nil {
		return nil, false
	}
	return o.MultiUserClients, true
}

// HasMultiUserClients returns a boolean if a field has been set.
func (o *LicenseDetailsUsage) HasMultiUserClients() bool {
	if o != nil && o.MultiUserClients != nil {
		return true
	}

	return false
}

// SetMultiUserClients gets a reference to the given float32 and assigns it to the MultiUserClients field.
func (o *LicenseDetailsUsage) SetMultiUserClients(v float32) {
	o.MultiUserClients = &v
}

// GetClientTypes returns the ClientTypes field value if set, zero value otherwise.
func (o *LicenseDetailsUsage) GetClientTypes() LicenseDetailsUsageClientTypes {
	if o == nil || o.ClientTypes == nil {
		var ret LicenseDetailsUsageClientTypes
		return ret
	}
	return *o.ClientTypes
}

// GetClientTypesOk returns a tuple with the ClientTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseDetailsUsage) GetClientTypesOk() (*LicenseDetailsUsageClientTypes, bool) {
	if o == nil || o.ClientTypes == nil {
		return nil, false
	}
	return o.ClientTypes, true
}

// HasClientTypes returns a boolean if a field has been set.
func (o *LicenseDetailsUsage) HasClientTypes() bool {
	if o != nil && o.ClientTypes != nil {
		return true
	}

	return false
}

// SetClientTypes gets a reference to the given LicenseDetailsUsageClientTypes and assigns it to the ClientTypes field.
func (o *LicenseDetailsUsage) SetClientTypes(v LicenseDetailsUsageClientTypes) {
	o.ClientTypes = &v
}

// GetMfaProviders returns the MfaProviders field value if set, zero value otherwise.
func (o *LicenseDetailsUsage) GetMfaProviders() float32 {
	if o == nil || o.MfaProviders == nil {
		var ret float32
		return ret
	}
	return *o.MfaProviders
}

// GetMfaProvidersOk returns a tuple with the MfaProviders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseDetailsUsage) GetMfaProvidersOk() (*float32, bool) {
	if o == nil || o.MfaProviders == nil {
		return nil, false
	}
	return o.MfaProviders, true
}

// HasMfaProviders returns a boolean if a field has been set.
func (o *LicenseDetailsUsage) HasMfaProviders() bool {
	if o != nil && o.MfaProviders != nil {
		return true
	}

	return false
}

// SetMfaProviders gets a reference to the given float32 and assigns it to the MfaProviders field.
func (o *LicenseDetailsUsage) SetMfaProviders(v float32) {
	o.MfaProviders = &v
}

// GetDeviceScripts returns the DeviceScripts field value if set, zero value otherwise.
func (o *LicenseDetailsUsage) GetDeviceScripts() float32 {
	if o == nil || o.DeviceScripts == nil {
		var ret float32
		return ret
	}
	return *o.DeviceScripts
}

// GetDeviceScriptsOk returns a tuple with the DeviceScripts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseDetailsUsage) GetDeviceScriptsOk() (*float32, bool) {
	if o == nil || o.DeviceScripts == nil {
		return nil, false
	}
	return o.DeviceScripts, true
}

// HasDeviceScripts returns a boolean if a field has been set.
func (o *LicenseDetailsUsage) HasDeviceScripts() bool {
	if o != nil && o.DeviceScripts != nil {
		return true
	}

	return false
}

// SetDeviceScripts gets a reference to the given float32 and assigns it to the DeviceScripts field.
func (o *LicenseDetailsUsage) SetDeviceScripts(v float32) {
	o.DeviceScripts = &v
}

// GetCriteriaScripts returns the CriteriaScripts field value if set, zero value otherwise.
func (o *LicenseDetailsUsage) GetCriteriaScripts() float32 {
	if o == nil || o.CriteriaScripts == nil {
		var ret float32
		return ret
	}
	return *o.CriteriaScripts
}

// GetCriteriaScriptsOk returns a tuple with the CriteriaScripts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseDetailsUsage) GetCriteriaScriptsOk() (*float32, bool) {
	if o == nil || o.CriteriaScripts == nil {
		return nil, false
	}
	return o.CriteriaScripts, true
}

// HasCriteriaScripts returns a boolean if a field has been set.
func (o *LicenseDetailsUsage) HasCriteriaScripts() bool {
	if o != nil && o.CriteriaScripts != nil {
		return true
	}

	return false
}

// SetCriteriaScripts gets a reference to the given float32 and assigns it to the CriteriaScripts field.
func (o *LicenseDetailsUsage) SetCriteriaScripts(v float32) {
	o.CriteriaScripts = &v
}

// GetEntitlementScripts returns the EntitlementScripts field value if set, zero value otherwise.
func (o *LicenseDetailsUsage) GetEntitlementScripts() float32 {
	if o == nil || o.EntitlementScripts == nil {
		var ret float32
		return ret
	}
	return *o.EntitlementScripts
}

// GetEntitlementScriptsOk returns a tuple with the EntitlementScripts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseDetailsUsage) GetEntitlementScriptsOk() (*float32, bool) {
	if o == nil || o.EntitlementScripts == nil {
		return nil, false
	}
	return o.EntitlementScripts, true
}

// HasEntitlementScripts returns a boolean if a field has been set.
func (o *LicenseDetailsUsage) HasEntitlementScripts() bool {
	if o != nil && o.EntitlementScripts != nil {
		return true
	}

	return false
}

// SetEntitlementScripts gets a reference to the given float32 and assigns it to the EntitlementScripts field.
func (o *LicenseDetailsUsage) SetEntitlementScripts(v float32) {
	o.EntitlementScripts = &v
}

// GetUserScripts returns the UserScripts field value if set, zero value otherwise.
func (o *LicenseDetailsUsage) GetUserScripts() float32 {
	if o == nil || o.UserScripts == nil {
		var ret float32
		return ret
	}
	return *o.UserScripts
}

// GetUserScriptsOk returns a tuple with the UserScripts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseDetailsUsage) GetUserScriptsOk() (*float32, bool) {
	if o == nil || o.UserScripts == nil {
		return nil, false
	}
	return o.UserScripts, true
}

// HasUserScripts returns a boolean if a field has been set.
func (o *LicenseDetailsUsage) HasUserScripts() bool {
	if o != nil && o.UserScripts != nil {
		return true
	}

	return false
}

// SetUserScripts gets a reference to the given float32 and assigns it to the UserScripts field.
func (o *LicenseDetailsUsage) SetUserScripts(v float32) {
	o.UserScripts = &v
}

// GetFallbackSites returns the FallbackSites field value if set, zero value otherwise.
func (o *LicenseDetailsUsage) GetFallbackSites() float32 {
	if o == nil || o.FallbackSites == nil {
		var ret float32
		return ret
	}
	return *o.FallbackSites
}

// GetFallbackSitesOk returns a tuple with the FallbackSites field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseDetailsUsage) GetFallbackSitesOk() (*float32, bool) {
	if o == nil || o.FallbackSites == nil {
		return nil, false
	}
	return o.FallbackSites, true
}

// HasFallbackSites returns a boolean if a field has been set.
func (o *LicenseDetailsUsage) HasFallbackSites() bool {
	if o != nil && o.FallbackSites != nil {
		return true
	}

	return false
}

// SetFallbackSites gets a reference to the given float32 and assigns it to the FallbackSites field.
func (o *LicenseDetailsUsage) SetFallbackSites(v float32) {
	o.FallbackSites = &v
}

// GetNearestSites returns the NearestSites field value if set, zero value otherwise.
func (o *LicenseDetailsUsage) GetNearestSites() float32 {
	if o == nil || o.NearestSites == nil {
		var ret float32
		return ret
	}
	return *o.NearestSites
}

// GetNearestSitesOk returns a tuple with the NearestSites field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseDetailsUsage) GetNearestSitesOk() (*float32, bool) {
	if o == nil || o.NearestSites == nil {
		return nil, false
	}
	return o.NearestSites, true
}

// HasNearestSites returns a boolean if a field has been set.
func (o *LicenseDetailsUsage) HasNearestSites() bool {
	if o != nil && o.NearestSites != nil {
		return true
	}

	return false
}

// SetNearestSites gets a reference to the given float32 and assigns it to the NearestSites field.
func (o *LicenseDetailsUsage) SetNearestSites(v float32) {
	o.NearestSites = &v
}

// GetLocalSites returns the LocalSites field value if set, zero value otherwise.
func (o *LicenseDetailsUsage) GetLocalSites() float32 {
	if o == nil || o.LocalSites == nil {
		var ret float32
		return ret
	}
	return *o.LocalSites
}

// GetLocalSitesOk returns a tuple with the LocalSites field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseDetailsUsage) GetLocalSitesOk() (*float32, bool) {
	if o == nil || o.LocalSites == nil {
		return nil, false
	}
	return o.LocalSites, true
}

// HasLocalSites returns a boolean if a field has been set.
func (o *LicenseDetailsUsage) HasLocalSites() bool {
	if o != nil && o.LocalSites != nil {
		return true
	}

	return false
}

// SetLocalSites gets a reference to the given float32 and assigns it to the LocalSites field.
func (o *LicenseDetailsUsage) SetLocalSites(v float32) {
	o.LocalSites = &v
}

// GetProfileGroups returns the ProfileGroups field value if set, zero value otherwise.
func (o *LicenseDetailsUsage) GetProfileGroups() float32 {
	if o == nil || o.ProfileGroups == nil {
		var ret float32
		return ret
	}
	return *o.ProfileGroups
}

// GetProfileGroupsOk returns a tuple with the ProfileGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseDetailsUsage) GetProfileGroupsOk() (*float32, bool) {
	if o == nil || o.ProfileGroups == nil {
		return nil, false
	}
	return o.ProfileGroups, true
}

// HasProfileGroups returns a boolean if a field has been set.
func (o *LicenseDetailsUsage) HasProfileGroups() bool {
	if o != nil && o.ProfileGroups != nil {
		return true
	}

	return false
}

// SetProfileGroups gets a reference to the given float32 and assigns it to the ProfileGroups field.
func (o *LicenseDetailsUsage) SetProfileGroups(v float32) {
	o.ProfileGroups = &v
}

// GetApplianceCustomizations returns the ApplianceCustomizations field value if set, zero value otherwise.
func (o *LicenseDetailsUsage) GetApplianceCustomizations() float32 {
	if o == nil || o.ApplianceCustomizations == nil {
		var ret float32
		return ret
	}
	return *o.ApplianceCustomizations
}

// GetApplianceCustomizationsOk returns a tuple with the ApplianceCustomizations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseDetailsUsage) GetApplianceCustomizationsOk() (*float32, bool) {
	if o == nil || o.ApplianceCustomizations == nil {
		return nil, false
	}
	return o.ApplianceCustomizations, true
}

// HasApplianceCustomizations returns a boolean if a field has been set.
func (o *LicenseDetailsUsage) HasApplianceCustomizations() bool {
	if o != nil && o.ApplianceCustomizations != nil {
		return true
	}

	return false
}

// SetApplianceCustomizations gets a reference to the given float32 and assigns it to the ApplianceCustomizations field.
func (o *LicenseDetailsUsage) SetApplianceCustomizations(v float32) {
	o.ApplianceCustomizations = &v
}

// GetRingfenceRules returns the RingfenceRules field value if set, zero value otherwise.
func (o *LicenseDetailsUsage) GetRingfenceRules() float32 {
	if o == nil || o.RingfenceRules == nil {
		var ret float32
		return ret
	}
	return *o.RingfenceRules
}

// GetRingfenceRulesOk returns a tuple with the RingfenceRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseDetailsUsage) GetRingfenceRulesOk() (*float32, bool) {
	if o == nil || o.RingfenceRules == nil {
		return nil, false
	}
	return o.RingfenceRules, true
}

// HasRingfenceRules returns a boolean if a field has been set.
func (o *LicenseDetailsUsage) HasRingfenceRules() bool {
	if o != nil && o.RingfenceRules != nil {
		return true
	}

	return false
}

// SetRingfenceRules gets a reference to the given float32 and assigns it to the RingfenceRules field.
func (o *LicenseDetailsUsage) SetRingfenceRules(v float32) {
	o.RingfenceRules = &v
}

// GetOtpSeeds returns the OtpSeeds field value if set, zero value otherwise.
func (o *LicenseDetailsUsage) GetOtpSeeds() float32 {
	if o == nil || o.OtpSeeds == nil {
		var ret float32
		return ret
	}
	return *o.OtpSeeds
}

// GetOtpSeedsOk returns a tuple with the OtpSeeds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseDetailsUsage) GetOtpSeedsOk() (*float32, bool) {
	if o == nil || o.OtpSeeds == nil {
		return nil, false
	}
	return o.OtpSeeds, true
}

// HasOtpSeeds returns a boolean if a field has been set.
func (o *LicenseDetailsUsage) HasOtpSeeds() bool {
	if o != nil && o.OtpSeeds != nil {
		return true
	}

	return false
}

// SetOtpSeeds gets a reference to the given float32 and assigns it to the OtpSeeds field.
func (o *LicenseDetailsUsage) SetOtpSeeds(v float32) {
	o.OtpSeeds = &v
}

// GetFido2Devices returns the Fido2Devices field value if set, zero value otherwise.
func (o *LicenseDetailsUsage) GetFido2Devices() float32 {
	if o == nil || o.Fido2Devices == nil {
		var ret float32
		return ret
	}
	return *o.Fido2Devices
}

// GetFido2DevicesOk returns a tuple with the Fido2Devices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseDetailsUsage) GetFido2DevicesOk() (*float32, bool) {
	if o == nil || o.Fido2Devices == nil {
		return nil, false
	}
	return o.Fido2Devices, true
}

// HasFido2Devices returns a boolean if a field has been set.
func (o *LicenseDetailsUsage) HasFido2Devices() bool {
	if o != nil && o.Fido2Devices != nil {
		return true
	}

	return false
}

// SetFido2Devices gets a reference to the given float32 and assigns it to the Fido2Devices field.
func (o *LicenseDetailsUsage) SetFido2Devices(v float32) {
	o.Fido2Devices = &v
}

// GetIdentityProviders returns the IdentityProviders field value if set, zero value otherwise.
func (o *LicenseDetailsUsage) GetIdentityProviders() LicenseDetailsUsageIdentityProviders {
	if o == nil || o.IdentityProviders == nil {
		var ret LicenseDetailsUsageIdentityProviders
		return ret
	}
	return *o.IdentityProviders
}

// GetIdentityProvidersOk returns a tuple with the IdentityProviders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseDetailsUsage) GetIdentityProvidersOk() (*LicenseDetailsUsageIdentityProviders, bool) {
	if o == nil || o.IdentityProviders == nil {
		return nil, false
	}
	return o.IdentityProviders, true
}

// HasIdentityProviders returns a boolean if a field has been set.
func (o *LicenseDetailsUsage) HasIdentityProviders() bool {
	if o != nil && o.IdentityProviders != nil {
		return true
	}

	return false
}

// SetIdentityProviders gets a reference to the given LicenseDetailsUsageIdentityProviders and assigns it to the IdentityProviders field.
func (o *LicenseDetailsUsage) SetIdentityProviders(v LicenseDetailsUsageIdentityProviders) {
	o.IdentityProviders = &v
}

// GetPolicies returns the Policies field value if set, zero value otherwise.
func (o *LicenseDetailsUsage) GetPolicies() LicenseDetailsUsagePolicies {
	if o == nil || o.Policies == nil {
		var ret LicenseDetailsUsagePolicies
		return ret
	}
	return *o.Policies
}

// GetPoliciesOk returns a tuple with the Policies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseDetailsUsage) GetPoliciesOk() (*LicenseDetailsUsagePolicies, bool) {
	if o == nil || o.Policies == nil {
		return nil, false
	}
	return o.Policies, true
}

// HasPolicies returns a boolean if a field has been set.
func (o *LicenseDetailsUsage) HasPolicies() bool {
	if o != nil && o.Policies != nil {
		return true
	}

	return false
}

// SetPolicies gets a reference to the given LicenseDetailsUsagePolicies and assigns it to the Policies field.
func (o *LicenseDetailsUsage) SetPolicies(v LicenseDetailsUsagePolicies) {
	o.Policies = &v
}

// GetRiskModel returns the RiskModel field value if set, zero value otherwise.
func (o *LicenseDetailsUsage) GetRiskModel() bool {
	if o == nil || o.RiskModel == nil {
		var ret bool
		return ret
	}
	return *o.RiskModel
}

// GetRiskModelOk returns a tuple with the RiskModel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseDetailsUsage) GetRiskModelOk() (*bool, bool) {
	if o == nil || o.RiskModel == nil {
		return nil, false
	}
	return o.RiskModel, true
}

// HasRiskModel returns a boolean if a field has been set.
func (o *LicenseDetailsUsage) HasRiskModel() bool {
	if o != nil && o.RiskModel != nil {
		return true
	}

	return false
}

// SetRiskModel gets a reference to the given bool and assigns it to the RiskModel field.
func (o *LicenseDetailsUsage) SetRiskModel(v bool) {
	o.RiskModel = &v
}

func (o LicenseDetailsUsage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Users != nil {
		toSerialize["users"] = o.Users
	}
	if o.PortalUsers != nil {
		toSerialize["portalUsers"] = o.PortalUsers
	}
	if o.ServiceUsers != nil {
		toSerialize["serviceUsers"] = o.ServiceUsers
	}
	if o.Sites != nil {
		toSerialize["sites"] = o.Sites
	}
	if o.ConnectorGroups != nil {
		toSerialize["connectorGroups"] = o.ConnectorGroups
	}
	if o.MultiUserClients != nil {
		toSerialize["multiUserClients"] = o.MultiUserClients
	}
	if o.ClientTypes != nil {
		toSerialize["clientTypes"] = o.ClientTypes
	}
	if o.MfaProviders != nil {
		toSerialize["mfaProviders"] = o.MfaProviders
	}
	if o.DeviceScripts != nil {
		toSerialize["deviceScripts"] = o.DeviceScripts
	}
	if o.CriteriaScripts != nil {
		toSerialize["criteriaScripts"] = o.CriteriaScripts
	}
	if o.EntitlementScripts != nil {
		toSerialize["entitlementScripts"] = o.EntitlementScripts
	}
	if o.UserScripts != nil {
		toSerialize["userScripts"] = o.UserScripts
	}
	if o.FallbackSites != nil {
		toSerialize["fallbackSites"] = o.FallbackSites
	}
	if o.NearestSites != nil {
		toSerialize["nearestSites"] = o.NearestSites
	}
	if o.LocalSites != nil {
		toSerialize["localSites"] = o.LocalSites
	}
	if o.ProfileGroups != nil {
		toSerialize["profileGroups"] = o.ProfileGroups
	}
	if o.ApplianceCustomizations != nil {
		toSerialize["applianceCustomizations"] = o.ApplianceCustomizations
	}
	if o.RingfenceRules != nil {
		toSerialize["ringfenceRules"] = o.RingfenceRules
	}
	if o.OtpSeeds != nil {
		toSerialize["otpSeeds"] = o.OtpSeeds
	}
	if o.Fido2Devices != nil {
		toSerialize["fido2Devices"] = o.Fido2Devices
	}
	if o.IdentityProviders != nil {
		toSerialize["identityProviders"] = o.IdentityProviders
	}
	if o.Policies != nil {
		toSerialize["policies"] = o.Policies
	}
	if o.RiskModel != nil {
		toSerialize["riskModel"] = o.RiskModel
	}
	return json.Marshal(toSerialize)
}

type NullableLicenseDetailsUsage struct {
	value *LicenseDetailsUsage
	isSet bool
}

func (v NullableLicenseDetailsUsage) Get() *LicenseDetailsUsage {
	return v.value
}

func (v *NullableLicenseDetailsUsage) Set(val *LicenseDetailsUsage) {
	v.value = val
	v.isSet = true
}

func (v NullableLicenseDetailsUsage) IsSet() bool {
	return v.isSet
}

func (v *NullableLicenseDetailsUsage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLicenseDetailsUsage(val *LicenseDetailsUsage) *NullableLicenseDetailsUsage {
	return &NullableLicenseDetailsUsage{value: val, isSet: true}
}

func (v NullableLicenseDetailsUsage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLicenseDetailsUsage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
