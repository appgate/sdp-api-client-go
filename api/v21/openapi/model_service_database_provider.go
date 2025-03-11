/*
Appgate SDP Controller REST API

# About   This specification documents the REST API calls for the Appgate SDP Controller.    Please refer to the REST API chapter in the manual or contact Appgate support with any questions about   this functionality. # Getting Started   Requirements for API scripting:   - Access to the Admin/API TLS Connection (default port 8443) of a Controller appliance.     (https://sdphelp.appgate.com/adminguide/appliance-function-configure.html?anchor=admin-api)   - An API user with relevant permissions.     (https://sdphelp.appgate.com/adminguide/administrative-roles-configure.html)   - In order to use the simple login API, Admin MFA must be disabled or the API user must be excluded.     (https://sdphelp.appgate.com/adminguide/mfa-for-admins.html) # Base path   HTTPS requests must be sent to the Admin Interface hostname and port, with **_/admin** path.    For example: **https://appgate.company.com:8443/admin**    All requests must have the **Accept** header as:    **application/vnd.appgate.peer-v21+json**    An exception is made for the **_/admin/version** endpoint which instead expects an **application/json** Accept header. # API Conventions   API conventions are  important to understand and follow strictly.    - While updating objects (via PUT), entire object must be sent with all fields.     - For example, in order to add a remedy method to the condition below:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": []       }       ```     - send the entire object with updated and non-updated fields:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": [{\"type\": \"DisplayMessage\", \"message\": \"test message\"}]       }       ```    - In case Controller returns an error (non-2xx HTTP status code), response body is JSON.     The \"message\" field contains information about the error.     HTTP 422 \"Unprocessable Entity\" has extra `errors` field to list all the issues with specific fields.    - Empty string (\"\") is considered a different value than \"null\" or field being omitted from JSON.     Omitting the field is recommended if no value is intended.     Empty string (\"\") will be almost always rejected as invalid value.    - There are common pattern between many objects:     - **Configuration Objects**: There are many objects with common fields, namely \"id\", \"name\", \"notes\", \"created\"       and \"updated\". These entities are listed, queried, created, updated and deleted in a similar fashion.     - **Distinguished Name**: Users and Devices are identified with what is called Distinguished Names, as used in        LDAP. The distinguished format that identifies a device and a user combination is        \"CN=\\<Device ID\\>,CN=\\<username\\>,OU=\\<Identity Provider Name\\>\". Some objects have the        \"userDistinguishedName\" field, which does not include the CN for Device ID.        This identifies a user on every device.

API version: API version 21.1
Contact: appgatesdp.support@appgate.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// ServiceDatabaseProvider struct for ServiceDatabaseProvider
type ServiceDatabaseProvider struct {
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
	// The DNS servers to be assigned to the Clients of the users in this Identity Provider. Deprecated as of 6.4. Use either the Automatic Client DNS feature in Site DNS resolvers or DNS Policies.
	// Deprecated
	DnsServers []string `json:"dnsServers,omitempty"`
	// The DNS search domains to be assigned to Clients of the users in this Identity Provider. Deprecated as of 6.4. Use either the Automatic Client DNS feature in Site DNS resolvers or DNS Policies.
	// Deprecated
	DnsSearchDomains []string `json:"dnsSearchDomains,omitempty"`
	// The device limit per user. The existing on-boarded devices will still be able to sign in even if the limit is exceeded.
	DeviceLimitPerUser *int32 `json:"deviceLimitPerUser,omitempty"`
	// After how many failed authentication attempts will a local user be locked out from authenticating again.
	UserLockoutThreshold *int32 `json:"userLockoutThreshold,omitempty"`
	// For how long lockout will last for local users.
	UserLockoutDurationMinutes *int32 `json:"userLockoutDurationMinutes,omitempty"`
	// Minimum password length requirement for local users.
	MinPasswordLength *int32 `json:"minPasswordLength,omitempty"`
}

// NewServiceDatabaseProvider instantiates a new ServiceDatabaseProvider object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceDatabaseProvider(name string, type_ string) *ServiceDatabaseProvider {
	this := ServiceDatabaseProvider{}
	this.Name = name
	this.Type = type_
	var deviceLimitPerUser int32 = 100
	this.DeviceLimitPerUser = &deviceLimitPerUser
	var userLockoutThreshold int32 = 5
	this.UserLockoutThreshold = &userLockoutThreshold
	var userLockoutDurationMinutes int32 = 1
	this.UserLockoutDurationMinutes = &userLockoutDurationMinutes
	var minPasswordLength int32 = 0
	this.MinPasswordLength = &minPasswordLength
	return &this
}

// NewServiceDatabaseProviderWithDefaults instantiates a new ServiceDatabaseProvider object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceDatabaseProviderWithDefaults() *ServiceDatabaseProvider {
	this := ServiceDatabaseProvider{}
	var deviceLimitPerUser int32 = 100
	this.DeviceLimitPerUser = &deviceLimitPerUser
	var userLockoutThreshold int32 = 5
	this.UserLockoutThreshold = &userLockoutThreshold
	var userLockoutDurationMinutes int32 = 1
	this.UserLockoutDurationMinutes = &userLockoutDurationMinutes
	var minPasswordLength int32 = 0
	this.MinPasswordLength = &minPasswordLength
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ServiceDatabaseProvider) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDatabaseProvider) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ServiceDatabaseProvider) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ServiceDatabaseProvider) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *ServiceDatabaseProvider) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ServiceDatabaseProvider) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ServiceDatabaseProvider) SetName(v string) {
	o.Name = v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *ServiceDatabaseProvider) GetNotes() string {
	if o == nil || o.Notes == nil {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDatabaseProvider) GetNotesOk() (*string, bool) {
	if o == nil || o.Notes == nil {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *ServiceDatabaseProvider) HasNotes() bool {
	if o != nil && o.Notes != nil {
		return true
	}

	return false
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *ServiceDatabaseProvider) SetNotes(v string) {
	o.Notes = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *ServiceDatabaseProvider) GetCreated() time.Time {
	if o == nil || o.Created == nil {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDatabaseProvider) GetCreatedOk() (*time.Time, bool) {
	if o == nil || o.Created == nil {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *ServiceDatabaseProvider) HasCreated() bool {
	if o != nil && o.Created != nil {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *ServiceDatabaseProvider) SetCreated(v time.Time) {
	o.Created = &v
}

// GetUpdated returns the Updated field value if set, zero value otherwise.
func (o *ServiceDatabaseProvider) GetUpdated() time.Time {
	if o == nil || o.Updated == nil {
		var ret time.Time
		return ret
	}
	return *o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDatabaseProvider) GetUpdatedOk() (*time.Time, bool) {
	if o == nil || o.Updated == nil {
		return nil, false
	}
	return o.Updated, true
}

// HasUpdated returns a boolean if a field has been set.
func (o *ServiceDatabaseProvider) HasUpdated() bool {
	if o != nil && o.Updated != nil {
		return true
	}

	return false
}

// SetUpdated gets a reference to the given time.Time and assigns it to the Updated field.
func (o *ServiceDatabaseProvider) SetUpdated(v time.Time) {
	o.Updated = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ServiceDatabaseProvider) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDatabaseProvider) GetTagsOk() ([]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ServiceDatabaseProvider) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *ServiceDatabaseProvider) SetTags(v []string) {
	o.Tags = v
}

// GetType returns the Type field value
func (o *ServiceDatabaseProvider) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ServiceDatabaseProvider) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ServiceDatabaseProvider) SetType(v string) {
	o.Type = v
}

// GetIpPoolV4 returns the IpPoolV4 field value if set, zero value otherwise.
func (o *ServiceDatabaseProvider) GetIpPoolV4() string {
	if o == nil || o.IpPoolV4 == nil {
		var ret string
		return ret
	}
	return *o.IpPoolV4
}

// GetIpPoolV4Ok returns a tuple with the IpPoolV4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDatabaseProvider) GetIpPoolV4Ok() (*string, bool) {
	if o == nil || o.IpPoolV4 == nil {
		return nil, false
	}
	return o.IpPoolV4, true
}

// HasIpPoolV4 returns a boolean if a field has been set.
func (o *ServiceDatabaseProvider) HasIpPoolV4() bool {
	if o != nil && o.IpPoolV4 != nil {
		return true
	}

	return false
}

// SetIpPoolV4 gets a reference to the given string and assigns it to the IpPoolV4 field.
func (o *ServiceDatabaseProvider) SetIpPoolV4(v string) {
	o.IpPoolV4 = &v
}

// GetIpPoolV6 returns the IpPoolV6 field value if set, zero value otherwise.
func (o *ServiceDatabaseProvider) GetIpPoolV6() string {
	if o == nil || o.IpPoolV6 == nil {
		var ret string
		return ret
	}
	return *o.IpPoolV6
}

// GetIpPoolV6Ok returns a tuple with the IpPoolV6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDatabaseProvider) GetIpPoolV6Ok() (*string, bool) {
	if o == nil || o.IpPoolV6 == nil {
		return nil, false
	}
	return o.IpPoolV6, true
}

// HasIpPoolV6 returns a boolean if a field has been set.
func (o *ServiceDatabaseProvider) HasIpPoolV6() bool {
	if o != nil && o.IpPoolV6 != nil {
		return true
	}

	return false
}

// SetIpPoolV6 gets a reference to the given string and assigns it to the IpPoolV6 field.
func (o *ServiceDatabaseProvider) SetIpPoolV6(v string) {
	o.IpPoolV6 = &v
}

// GetClaimMappings returns the ClaimMappings field value if set, zero value otherwise.
func (o *ServiceDatabaseProvider) GetClaimMappings() []ClaimMappingsInner {
	if o == nil || o.ClaimMappings == nil {
		var ret []ClaimMappingsInner
		return ret
	}
	return o.ClaimMappings
}

// GetClaimMappingsOk returns a tuple with the ClaimMappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDatabaseProvider) GetClaimMappingsOk() ([]ClaimMappingsInner, bool) {
	if o == nil || o.ClaimMappings == nil {
		return nil, false
	}
	return o.ClaimMappings, true
}

// HasClaimMappings returns a boolean if a field has been set.
func (o *ServiceDatabaseProvider) HasClaimMappings() bool {
	if o != nil && o.ClaimMappings != nil {
		return true
	}

	return false
}

// SetClaimMappings gets a reference to the given []ClaimMappingsInner and assigns it to the ClaimMappings field.
func (o *ServiceDatabaseProvider) SetClaimMappings(v []ClaimMappingsInner) {
	o.ClaimMappings = v
}

// GetUserScripts returns the UserScripts field value if set, zero value otherwise.
func (o *ServiceDatabaseProvider) GetUserScripts() []string {
	if o == nil || o.UserScripts == nil {
		var ret []string
		return ret
	}
	return o.UserScripts
}

// GetUserScriptsOk returns a tuple with the UserScripts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDatabaseProvider) GetUserScriptsOk() ([]string, bool) {
	if o == nil || o.UserScripts == nil {
		return nil, false
	}
	return o.UserScripts, true
}

// HasUserScripts returns a boolean if a field has been set.
func (o *ServiceDatabaseProvider) HasUserScripts() bool {
	if o != nil && o.UserScripts != nil {
		return true
	}

	return false
}

// SetUserScripts gets a reference to the given []string and assigns it to the UserScripts field.
func (o *ServiceDatabaseProvider) SetUserScripts(v []string) {
	o.UserScripts = v
}

// GetDnsServers returns the DnsServers field value if set, zero value otherwise.
// Deprecated
func (o *ServiceDatabaseProvider) GetDnsServers() []string {
	if o == nil || o.DnsServers == nil {
		var ret []string
		return ret
	}
	return o.DnsServers
}

// GetDnsServersOk returns a tuple with the DnsServers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *ServiceDatabaseProvider) GetDnsServersOk() ([]string, bool) {
	if o == nil || o.DnsServers == nil {
		return nil, false
	}
	return o.DnsServers, true
}

// HasDnsServers returns a boolean if a field has been set.
func (o *ServiceDatabaseProvider) HasDnsServers() bool {
	if o != nil && o.DnsServers != nil {
		return true
	}

	return false
}

// SetDnsServers gets a reference to the given []string and assigns it to the DnsServers field.
// Deprecated
func (o *ServiceDatabaseProvider) SetDnsServers(v []string) {
	o.DnsServers = v
}

// GetDnsSearchDomains returns the DnsSearchDomains field value if set, zero value otherwise.
// Deprecated
func (o *ServiceDatabaseProvider) GetDnsSearchDomains() []string {
	if o == nil || o.DnsSearchDomains == nil {
		var ret []string
		return ret
	}
	return o.DnsSearchDomains
}

// GetDnsSearchDomainsOk returns a tuple with the DnsSearchDomains field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *ServiceDatabaseProvider) GetDnsSearchDomainsOk() ([]string, bool) {
	if o == nil || o.DnsSearchDomains == nil {
		return nil, false
	}
	return o.DnsSearchDomains, true
}

// HasDnsSearchDomains returns a boolean if a field has been set.
func (o *ServiceDatabaseProvider) HasDnsSearchDomains() bool {
	if o != nil && o.DnsSearchDomains != nil {
		return true
	}

	return false
}

// SetDnsSearchDomains gets a reference to the given []string and assigns it to the DnsSearchDomains field.
// Deprecated
func (o *ServiceDatabaseProvider) SetDnsSearchDomains(v []string) {
	o.DnsSearchDomains = v
}

// GetDeviceLimitPerUser returns the DeviceLimitPerUser field value if set, zero value otherwise.
func (o *ServiceDatabaseProvider) GetDeviceLimitPerUser() int32 {
	if o == nil || o.DeviceLimitPerUser == nil {
		var ret int32
		return ret
	}
	return *o.DeviceLimitPerUser
}

// GetDeviceLimitPerUserOk returns a tuple with the DeviceLimitPerUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDatabaseProvider) GetDeviceLimitPerUserOk() (*int32, bool) {
	if o == nil || o.DeviceLimitPerUser == nil {
		return nil, false
	}
	return o.DeviceLimitPerUser, true
}

// HasDeviceLimitPerUser returns a boolean if a field has been set.
func (o *ServiceDatabaseProvider) HasDeviceLimitPerUser() bool {
	if o != nil && o.DeviceLimitPerUser != nil {
		return true
	}

	return false
}

// SetDeviceLimitPerUser gets a reference to the given int32 and assigns it to the DeviceLimitPerUser field.
func (o *ServiceDatabaseProvider) SetDeviceLimitPerUser(v int32) {
	o.DeviceLimitPerUser = &v
}

// GetUserLockoutThreshold returns the UserLockoutThreshold field value if set, zero value otherwise.
func (o *ServiceDatabaseProvider) GetUserLockoutThreshold() int32 {
	if o == nil || o.UserLockoutThreshold == nil {
		var ret int32
		return ret
	}
	return *o.UserLockoutThreshold
}

// GetUserLockoutThresholdOk returns a tuple with the UserLockoutThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDatabaseProvider) GetUserLockoutThresholdOk() (*int32, bool) {
	if o == nil || o.UserLockoutThreshold == nil {
		return nil, false
	}
	return o.UserLockoutThreshold, true
}

// HasUserLockoutThreshold returns a boolean if a field has been set.
func (o *ServiceDatabaseProvider) HasUserLockoutThreshold() bool {
	if o != nil && o.UserLockoutThreshold != nil {
		return true
	}

	return false
}

// SetUserLockoutThreshold gets a reference to the given int32 and assigns it to the UserLockoutThreshold field.
func (o *ServiceDatabaseProvider) SetUserLockoutThreshold(v int32) {
	o.UserLockoutThreshold = &v
}

// GetUserLockoutDurationMinutes returns the UserLockoutDurationMinutes field value if set, zero value otherwise.
func (o *ServiceDatabaseProvider) GetUserLockoutDurationMinutes() int32 {
	if o == nil || o.UserLockoutDurationMinutes == nil {
		var ret int32
		return ret
	}
	return *o.UserLockoutDurationMinutes
}

// GetUserLockoutDurationMinutesOk returns a tuple with the UserLockoutDurationMinutes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDatabaseProvider) GetUserLockoutDurationMinutesOk() (*int32, bool) {
	if o == nil || o.UserLockoutDurationMinutes == nil {
		return nil, false
	}
	return o.UserLockoutDurationMinutes, true
}

// HasUserLockoutDurationMinutes returns a boolean if a field has been set.
func (o *ServiceDatabaseProvider) HasUserLockoutDurationMinutes() bool {
	if o != nil && o.UserLockoutDurationMinutes != nil {
		return true
	}

	return false
}

// SetUserLockoutDurationMinutes gets a reference to the given int32 and assigns it to the UserLockoutDurationMinutes field.
func (o *ServiceDatabaseProvider) SetUserLockoutDurationMinutes(v int32) {
	o.UserLockoutDurationMinutes = &v
}

// GetMinPasswordLength returns the MinPasswordLength field value if set, zero value otherwise.
func (o *ServiceDatabaseProvider) GetMinPasswordLength() int32 {
	if o == nil || o.MinPasswordLength == nil {
		var ret int32
		return ret
	}
	return *o.MinPasswordLength
}

// GetMinPasswordLengthOk returns a tuple with the MinPasswordLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDatabaseProvider) GetMinPasswordLengthOk() (*int32, bool) {
	if o == nil || o.MinPasswordLength == nil {
		return nil, false
	}
	return o.MinPasswordLength, true
}

// HasMinPasswordLength returns a boolean if a field has been set.
func (o *ServiceDatabaseProvider) HasMinPasswordLength() bool {
	if o != nil && o.MinPasswordLength != nil {
		return true
	}

	return false
}

// SetMinPasswordLength gets a reference to the given int32 and assigns it to the MinPasswordLength field.
func (o *ServiceDatabaseProvider) SetMinPasswordLength(v int32) {
	o.MinPasswordLength = &v
}

func (o ServiceDatabaseProvider) MarshalJSON() ([]byte, error) {
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
	if o.DeviceLimitPerUser != nil {
		toSerialize["deviceLimitPerUser"] = o.DeviceLimitPerUser
	}
	if o.UserLockoutThreshold != nil {
		toSerialize["userLockoutThreshold"] = o.UserLockoutThreshold
	}
	if o.UserLockoutDurationMinutes != nil {
		toSerialize["userLockoutDurationMinutes"] = o.UserLockoutDurationMinutes
	}
	if o.MinPasswordLength != nil {
		toSerialize["minPasswordLength"] = o.MinPasswordLength
	}
	return json.Marshal(toSerialize)
}

type NullableServiceDatabaseProvider struct {
	value *ServiceDatabaseProvider
	isSet bool
}

func (v NullableServiceDatabaseProvider) Get() *ServiceDatabaseProvider {
	return v.value
}

func (v *NullableServiceDatabaseProvider) Set(val *ServiceDatabaseProvider) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceDatabaseProvider) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceDatabaseProvider) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceDatabaseProvider(val *ServiceDatabaseProvider) *NullableServiceDatabaseProvider {
	return &NullableServiceDatabaseProvider{value: val, isSet: true}
}

func (v NullableServiceDatabaseProvider) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceDatabaseProvider) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
