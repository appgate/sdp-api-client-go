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
	"time"
)

// SiteWithStatus struct for SiteWithStatus
type SiteWithStatus struct {
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
	// A short 4 letter name for the Site to be displayed on the Client.
	ShortName *string `json:"shortName,omitempty"`
	// Description of the Site to be displayed on the Client.
	Description *string               `json:"description,omitempty"`
	Geolocation *SiteAllOfGeolocation `json:"geolocation,omitempty"`
	// Network subnets in CIDR format to define the Site's boundaries. They are added as routes by the Client.
	NetworkSubnets []string `json:"networkSubnets,omitempty"`
	// When the Client fails to connect to the Site for a certain period of time, configured Entitlements (see Policy) will be moved to this \"Fallback\" Site.
	FallbackSite       *string                      `json:"fallbackSite,omitempty"`
	LocalSiteDetection *SiteAllOfLocalSiteDetection `json:"localSiteDetection,omitempty"`
	// If enabled, this Site will be included in the nearest Site override selection in Policies.
	UseForNearestSiteSelection *bool `json:"useForNearestSiteSelection,omitempty"`
	// List of IP Pool mappings for this specific Site. When IPs are allocated this Site, they will be mapped to a new one using this setting.
	IpPoolMappings []SiteAllOfIpPoolMappings `json:"ipPoolMappings,omitempty"`
	DefaultGateway *SiteAllOfDefaultGateway  `json:"defaultGateway,omitempty"`
	// When enabled, the routes are sent to the Client by the Gateways according to the user's Entitlements \"networkSubnets\" should be left be empty if it's enabled.
	EntitlementBasedRouting *bool                    `json:"entitlementBasedRouting,omitempty"`
	Vpn                     *SiteAllOfVpn            `json:"vpn,omitempty"`
	NameResolution          *SiteAllOfNameResolution `json:"nameResolution,omitempty"`
	// Health status of the Site. Depends on the status of the Appliances in this Site.
	Status *string `json:"status,omitempty"`
}

// NewSiteWithStatus instantiates a new SiteWithStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSiteWithStatus(name string) *SiteWithStatus {
	this := SiteWithStatus{}
	this.Name = name
	var entitlementBasedRouting bool = false
	this.EntitlementBasedRouting = &entitlementBasedRouting
	return &this
}

// NewSiteWithStatusWithDefaults instantiates a new SiteWithStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSiteWithStatusWithDefaults() *SiteWithStatus {
	this := SiteWithStatus{}
	var entitlementBasedRouting bool = false
	this.EntitlementBasedRouting = &entitlementBasedRouting
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SiteWithStatus) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteWithStatus) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SiteWithStatus) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SiteWithStatus) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *SiteWithStatus) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SiteWithStatus) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SiteWithStatus) SetName(v string) {
	o.Name = v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *SiteWithStatus) GetNotes() string {
	if o == nil || o.Notes == nil {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteWithStatus) GetNotesOk() (*string, bool) {
	if o == nil || o.Notes == nil {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *SiteWithStatus) HasNotes() bool {
	if o != nil && o.Notes != nil {
		return true
	}

	return false
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *SiteWithStatus) SetNotes(v string) {
	o.Notes = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *SiteWithStatus) GetCreated() time.Time {
	if o == nil || o.Created == nil {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteWithStatus) GetCreatedOk() (*time.Time, bool) {
	if o == nil || o.Created == nil {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *SiteWithStatus) HasCreated() bool {
	if o != nil && o.Created != nil {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *SiteWithStatus) SetCreated(v time.Time) {
	o.Created = &v
}

// GetUpdated returns the Updated field value if set, zero value otherwise.
func (o *SiteWithStatus) GetUpdated() time.Time {
	if o == nil || o.Updated == nil {
		var ret time.Time
		return ret
	}
	return *o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteWithStatus) GetUpdatedOk() (*time.Time, bool) {
	if o == nil || o.Updated == nil {
		return nil, false
	}
	return o.Updated, true
}

// HasUpdated returns a boolean if a field has been set.
func (o *SiteWithStatus) HasUpdated() bool {
	if o != nil && o.Updated != nil {
		return true
	}

	return false
}

// SetUpdated gets a reference to the given time.Time and assigns it to the Updated field.
func (o *SiteWithStatus) SetUpdated(v time.Time) {
	o.Updated = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *SiteWithStatus) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteWithStatus) GetTagsOk() ([]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *SiteWithStatus) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *SiteWithStatus) SetTags(v []string) {
	o.Tags = v
}

// GetShortName returns the ShortName field value if set, zero value otherwise.
func (o *SiteWithStatus) GetShortName() string {
	if o == nil || o.ShortName == nil {
		var ret string
		return ret
	}
	return *o.ShortName
}

// GetShortNameOk returns a tuple with the ShortName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteWithStatus) GetShortNameOk() (*string, bool) {
	if o == nil || o.ShortName == nil {
		return nil, false
	}
	return o.ShortName, true
}

// HasShortName returns a boolean if a field has been set.
func (o *SiteWithStatus) HasShortName() bool {
	if o != nil && o.ShortName != nil {
		return true
	}

	return false
}

// SetShortName gets a reference to the given string and assigns it to the ShortName field.
func (o *SiteWithStatus) SetShortName(v string) {
	o.ShortName = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SiteWithStatus) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteWithStatus) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SiteWithStatus) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SiteWithStatus) SetDescription(v string) {
	o.Description = &v
}

// GetGeolocation returns the Geolocation field value if set, zero value otherwise.
func (o *SiteWithStatus) GetGeolocation() SiteAllOfGeolocation {
	if o == nil || o.Geolocation == nil {
		var ret SiteAllOfGeolocation
		return ret
	}
	return *o.Geolocation
}

// GetGeolocationOk returns a tuple with the Geolocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteWithStatus) GetGeolocationOk() (*SiteAllOfGeolocation, bool) {
	if o == nil || o.Geolocation == nil {
		return nil, false
	}
	return o.Geolocation, true
}

// HasGeolocation returns a boolean if a field has been set.
func (o *SiteWithStatus) HasGeolocation() bool {
	if o != nil && o.Geolocation != nil {
		return true
	}

	return false
}

// SetGeolocation gets a reference to the given SiteAllOfGeolocation and assigns it to the Geolocation field.
func (o *SiteWithStatus) SetGeolocation(v SiteAllOfGeolocation) {
	o.Geolocation = &v
}

// GetNetworkSubnets returns the NetworkSubnets field value if set, zero value otherwise.
func (o *SiteWithStatus) GetNetworkSubnets() []string {
	if o == nil || o.NetworkSubnets == nil {
		var ret []string
		return ret
	}
	return o.NetworkSubnets
}

// GetNetworkSubnetsOk returns a tuple with the NetworkSubnets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteWithStatus) GetNetworkSubnetsOk() ([]string, bool) {
	if o == nil || o.NetworkSubnets == nil {
		return nil, false
	}
	return o.NetworkSubnets, true
}

// HasNetworkSubnets returns a boolean if a field has been set.
func (o *SiteWithStatus) HasNetworkSubnets() bool {
	if o != nil && o.NetworkSubnets != nil {
		return true
	}

	return false
}

// SetNetworkSubnets gets a reference to the given []string and assigns it to the NetworkSubnets field.
func (o *SiteWithStatus) SetNetworkSubnets(v []string) {
	o.NetworkSubnets = v
}

// GetFallbackSite returns the FallbackSite field value if set, zero value otherwise.
func (o *SiteWithStatus) GetFallbackSite() string {
	if o == nil || o.FallbackSite == nil {
		var ret string
		return ret
	}
	return *o.FallbackSite
}

// GetFallbackSiteOk returns a tuple with the FallbackSite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteWithStatus) GetFallbackSiteOk() (*string, bool) {
	if o == nil || o.FallbackSite == nil {
		return nil, false
	}
	return o.FallbackSite, true
}

// HasFallbackSite returns a boolean if a field has been set.
func (o *SiteWithStatus) HasFallbackSite() bool {
	if o != nil && o.FallbackSite != nil {
		return true
	}

	return false
}

// SetFallbackSite gets a reference to the given string and assigns it to the FallbackSite field.
func (o *SiteWithStatus) SetFallbackSite(v string) {
	o.FallbackSite = &v
}

// GetLocalSiteDetection returns the LocalSiteDetection field value if set, zero value otherwise.
func (o *SiteWithStatus) GetLocalSiteDetection() SiteAllOfLocalSiteDetection {
	if o == nil || o.LocalSiteDetection == nil {
		var ret SiteAllOfLocalSiteDetection
		return ret
	}
	return *o.LocalSiteDetection
}

// GetLocalSiteDetectionOk returns a tuple with the LocalSiteDetection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteWithStatus) GetLocalSiteDetectionOk() (*SiteAllOfLocalSiteDetection, bool) {
	if o == nil || o.LocalSiteDetection == nil {
		return nil, false
	}
	return o.LocalSiteDetection, true
}

// HasLocalSiteDetection returns a boolean if a field has been set.
func (o *SiteWithStatus) HasLocalSiteDetection() bool {
	if o != nil && o.LocalSiteDetection != nil {
		return true
	}

	return false
}

// SetLocalSiteDetection gets a reference to the given SiteAllOfLocalSiteDetection and assigns it to the LocalSiteDetection field.
func (o *SiteWithStatus) SetLocalSiteDetection(v SiteAllOfLocalSiteDetection) {
	o.LocalSiteDetection = &v
}

// GetUseForNearestSiteSelection returns the UseForNearestSiteSelection field value if set, zero value otherwise.
func (o *SiteWithStatus) GetUseForNearestSiteSelection() bool {
	if o == nil || o.UseForNearestSiteSelection == nil {
		var ret bool
		return ret
	}
	return *o.UseForNearestSiteSelection
}

// GetUseForNearestSiteSelectionOk returns a tuple with the UseForNearestSiteSelection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteWithStatus) GetUseForNearestSiteSelectionOk() (*bool, bool) {
	if o == nil || o.UseForNearestSiteSelection == nil {
		return nil, false
	}
	return o.UseForNearestSiteSelection, true
}

// HasUseForNearestSiteSelection returns a boolean if a field has been set.
func (o *SiteWithStatus) HasUseForNearestSiteSelection() bool {
	if o != nil && o.UseForNearestSiteSelection != nil {
		return true
	}

	return false
}

// SetUseForNearestSiteSelection gets a reference to the given bool and assigns it to the UseForNearestSiteSelection field.
func (o *SiteWithStatus) SetUseForNearestSiteSelection(v bool) {
	o.UseForNearestSiteSelection = &v
}

// GetIpPoolMappings returns the IpPoolMappings field value if set, zero value otherwise.
func (o *SiteWithStatus) GetIpPoolMappings() []SiteAllOfIpPoolMappings {
	if o == nil || o.IpPoolMappings == nil {
		var ret []SiteAllOfIpPoolMappings
		return ret
	}
	return o.IpPoolMappings
}

// GetIpPoolMappingsOk returns a tuple with the IpPoolMappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteWithStatus) GetIpPoolMappingsOk() ([]SiteAllOfIpPoolMappings, bool) {
	if o == nil || o.IpPoolMappings == nil {
		return nil, false
	}
	return o.IpPoolMappings, true
}

// HasIpPoolMappings returns a boolean if a field has been set.
func (o *SiteWithStatus) HasIpPoolMappings() bool {
	if o != nil && o.IpPoolMappings != nil {
		return true
	}

	return false
}

// SetIpPoolMappings gets a reference to the given []SiteAllOfIpPoolMappings and assigns it to the IpPoolMappings field.
func (o *SiteWithStatus) SetIpPoolMappings(v []SiteAllOfIpPoolMappings) {
	o.IpPoolMappings = v
}

// GetDefaultGateway returns the DefaultGateway field value if set, zero value otherwise.
func (o *SiteWithStatus) GetDefaultGateway() SiteAllOfDefaultGateway {
	if o == nil || o.DefaultGateway == nil {
		var ret SiteAllOfDefaultGateway
		return ret
	}
	return *o.DefaultGateway
}

// GetDefaultGatewayOk returns a tuple with the DefaultGateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteWithStatus) GetDefaultGatewayOk() (*SiteAllOfDefaultGateway, bool) {
	if o == nil || o.DefaultGateway == nil {
		return nil, false
	}
	return o.DefaultGateway, true
}

// HasDefaultGateway returns a boolean if a field has been set.
func (o *SiteWithStatus) HasDefaultGateway() bool {
	if o != nil && o.DefaultGateway != nil {
		return true
	}

	return false
}

// SetDefaultGateway gets a reference to the given SiteAllOfDefaultGateway and assigns it to the DefaultGateway field.
func (o *SiteWithStatus) SetDefaultGateway(v SiteAllOfDefaultGateway) {
	o.DefaultGateway = &v
}

// GetEntitlementBasedRouting returns the EntitlementBasedRouting field value if set, zero value otherwise.
func (o *SiteWithStatus) GetEntitlementBasedRouting() bool {
	if o == nil || o.EntitlementBasedRouting == nil {
		var ret bool
		return ret
	}
	return *o.EntitlementBasedRouting
}

// GetEntitlementBasedRoutingOk returns a tuple with the EntitlementBasedRouting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteWithStatus) GetEntitlementBasedRoutingOk() (*bool, bool) {
	if o == nil || o.EntitlementBasedRouting == nil {
		return nil, false
	}
	return o.EntitlementBasedRouting, true
}

// HasEntitlementBasedRouting returns a boolean if a field has been set.
func (o *SiteWithStatus) HasEntitlementBasedRouting() bool {
	if o != nil && o.EntitlementBasedRouting != nil {
		return true
	}

	return false
}

// SetEntitlementBasedRouting gets a reference to the given bool and assigns it to the EntitlementBasedRouting field.
func (o *SiteWithStatus) SetEntitlementBasedRouting(v bool) {
	o.EntitlementBasedRouting = &v
}

// GetVpn returns the Vpn field value if set, zero value otherwise.
func (o *SiteWithStatus) GetVpn() SiteAllOfVpn {
	if o == nil || o.Vpn == nil {
		var ret SiteAllOfVpn
		return ret
	}
	return *o.Vpn
}

// GetVpnOk returns a tuple with the Vpn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteWithStatus) GetVpnOk() (*SiteAllOfVpn, bool) {
	if o == nil || o.Vpn == nil {
		return nil, false
	}
	return o.Vpn, true
}

// HasVpn returns a boolean if a field has been set.
func (o *SiteWithStatus) HasVpn() bool {
	if o != nil && o.Vpn != nil {
		return true
	}

	return false
}

// SetVpn gets a reference to the given SiteAllOfVpn and assigns it to the Vpn field.
func (o *SiteWithStatus) SetVpn(v SiteAllOfVpn) {
	o.Vpn = &v
}

// GetNameResolution returns the NameResolution field value if set, zero value otherwise.
func (o *SiteWithStatus) GetNameResolution() SiteAllOfNameResolution {
	if o == nil || o.NameResolution == nil {
		var ret SiteAllOfNameResolution
		return ret
	}
	return *o.NameResolution
}

// GetNameResolutionOk returns a tuple with the NameResolution field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteWithStatus) GetNameResolutionOk() (*SiteAllOfNameResolution, bool) {
	if o == nil || o.NameResolution == nil {
		return nil, false
	}
	return o.NameResolution, true
}

// HasNameResolution returns a boolean if a field has been set.
func (o *SiteWithStatus) HasNameResolution() bool {
	if o != nil && o.NameResolution != nil {
		return true
	}

	return false
}

// SetNameResolution gets a reference to the given SiteAllOfNameResolution and assigns it to the NameResolution field.
func (o *SiteWithStatus) SetNameResolution(v SiteAllOfNameResolution) {
	o.NameResolution = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SiteWithStatus) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteWithStatus) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SiteWithStatus) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *SiteWithStatus) SetStatus(v string) {
	o.Status = &v
}

func (o SiteWithStatus) MarshalJSON() ([]byte, error) {
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
	if o.ShortName != nil {
		toSerialize["shortName"] = o.ShortName
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Geolocation != nil {
		toSerialize["geolocation"] = o.Geolocation
	}
	if o.NetworkSubnets != nil {
		toSerialize["networkSubnets"] = o.NetworkSubnets
	}
	if o.FallbackSite != nil {
		toSerialize["fallbackSite"] = o.FallbackSite
	}
	if o.LocalSiteDetection != nil {
		toSerialize["localSiteDetection"] = o.LocalSiteDetection
	}
	if o.UseForNearestSiteSelection != nil {
		toSerialize["useForNearestSiteSelection"] = o.UseForNearestSiteSelection
	}
	if o.IpPoolMappings != nil {
		toSerialize["ipPoolMappings"] = o.IpPoolMappings
	}
	if o.DefaultGateway != nil {
		toSerialize["defaultGateway"] = o.DefaultGateway
	}
	if o.EntitlementBasedRouting != nil {
		toSerialize["entitlementBasedRouting"] = o.EntitlementBasedRouting
	}
	if o.Vpn != nil {
		toSerialize["vpn"] = o.Vpn
	}
	if o.NameResolution != nil {
		toSerialize["nameResolution"] = o.NameResolution
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableSiteWithStatus struct {
	value *SiteWithStatus
	isSet bool
}

func (v NullableSiteWithStatus) Get() *SiteWithStatus {
	return v.value
}

func (v *NullableSiteWithStatus) Set(val *SiteWithStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableSiteWithStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableSiteWithStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSiteWithStatus(val *SiteWithStatus) *NullableSiteWithStatus {
	return &NullableSiteWithStatus{value: val, isSet: true}
}

func (v NullableSiteWithStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSiteWithStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
