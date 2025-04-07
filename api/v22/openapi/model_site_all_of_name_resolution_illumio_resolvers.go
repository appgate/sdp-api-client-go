/*
Appgate SDP Controller REST API

# About   This specification documents the REST API calls for the Appgate SDP Controller.    Please refer to the REST API chapter in the manual or contact Appgate support with any questions about   this functionality. # Getting Started   Requirements for API scripting:   - Access to the Admin/API TLS Connection (default port 8443) of a Controller appliance.     (https://sdphelp.appgate.com/adminguide/appliance-function-configure.html?anchor=admin-api)   - An API user with relevant permissions.     (https://sdphelp.appgate.com/adminguide/administrative-roles-configure.html)   - In order to use the simple login API, Admin MFA must be disabled or the API user must be excluded.     (https://sdphelp.appgate.com/adminguide/mfa-for-admins.html) # Base path   HTTPS requests must be sent to the Admin Interface hostname and port, with **_/admin** path.    For example: **https://appgate.company.com:8443/admin**    All requests must have the **Accept** header as:    **application/vnd.appgate.peer-v22+json**    An exception is made for the **_/admin/version** endpoint which instead expects an **application/json** Accept header. # API Conventions   API conventions are  important to understand and follow strictly.    - While updating objects (via PUT), entire object must be sent with all fields.     - For example, in order to add a remedy method to the condition below:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": []       }       ```     - send the entire object with updated and non-updated fields:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": [{\"type\": \"DisplayMessage\", \"message\": \"test message\"}]       }       ```    - In case Controller returns an error (non-2xx HTTP status code), response body is JSON.     The \"message\" field contains information about the error.     HTTP 422 \"Unprocessable Entity\" has extra `errors` field to list all the issues with specific fields.    - Empty string (\"\") is considered a different value than \"null\" or field being omitted from JSON.     Omitting the field is recommended if no value is intended.     Empty string (\"\") will be almost always rejected as invalid value.    - There are common pattern between many objects:     - **Configuration Objects**: There are many objects with common fields, namely \"id\", \"name\", \"notes\", \"created\"       and \"updated\". These entities are listed, queried, created, updated and deleted in a similar fashion.     - **Distinguished Name**: Users and Devices are identified with what is called Distinguished Names, as used in        LDAP. The distinguished format that identifies a device and a user combination is        \"CN=\\<Device ID\\>,CN=\\<username\\>,OU=\\<Identity Provider Name\\>\". Some objects have the        \"userDistinguishedName\" field, which does not include the CN for Device ID.        This identifies a user on every device.

API version: API version 22.1
Contact: appgatesdp.support@appgate.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// SiteAllOfNameResolutionIllumioResolvers struct for SiteAllOfNameResolutionIllumioResolvers
type SiteAllOfNameResolutionIllumioResolvers struct {
	// Identifier name. Has no functional effect.
	Name string `json:"name"`
	// How often will the resolver poll the server. In seconds.
	UpdateInterval *int32 `json:"updateInterval,omitempty"`
	// Organization ID of the Illumio Resolver.
	OrgId string `json:"orgId"`
	// Hostname of the Illumio Resolver.
	Hostname string `json:"hostname"`
	// Port number of the Illumio Resolver.
	Port int32 `json:"port"`
	// Username with access to the Illumio Resolver.
	Username string `json:"username"`
	// Password for the username.
	Password *string `json:"password,omitempty"`
}

// NewSiteAllOfNameResolutionIllumioResolvers instantiates a new SiteAllOfNameResolutionIllumioResolvers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSiteAllOfNameResolutionIllumioResolvers(name string, orgId string, hostname string, port int32, username string) *SiteAllOfNameResolutionIllumioResolvers {
	this := SiteAllOfNameResolutionIllumioResolvers{}
	this.Name = name
	var updateInterval int32 = 60
	this.UpdateInterval = &updateInterval
	this.OrgId = orgId
	this.Hostname = hostname
	this.Port = port
	this.Username = username
	return &this
}

// NewSiteAllOfNameResolutionIllumioResolversWithDefaults instantiates a new SiteAllOfNameResolutionIllumioResolvers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSiteAllOfNameResolutionIllumioResolversWithDefaults() *SiteAllOfNameResolutionIllumioResolvers {
	this := SiteAllOfNameResolutionIllumioResolvers{}
	var updateInterval int32 = 60
	this.UpdateInterval = &updateInterval
	return &this
}

// GetName returns the Name field value
func (o *SiteAllOfNameResolutionIllumioResolvers) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SiteAllOfNameResolutionIllumioResolvers) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SiteAllOfNameResolutionIllumioResolvers) SetName(v string) {
	o.Name = v
}

// GetUpdateInterval returns the UpdateInterval field value if set, zero value otherwise.
func (o *SiteAllOfNameResolutionIllumioResolvers) GetUpdateInterval() int32 {
	if o == nil || o.UpdateInterval == nil {
		var ret int32
		return ret
	}
	return *o.UpdateInterval
}

// GetUpdateIntervalOk returns a tuple with the UpdateInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteAllOfNameResolutionIllumioResolvers) GetUpdateIntervalOk() (*int32, bool) {
	if o == nil || o.UpdateInterval == nil {
		return nil, false
	}
	return o.UpdateInterval, true
}

// HasUpdateInterval returns a boolean if a field has been set.
func (o *SiteAllOfNameResolutionIllumioResolvers) HasUpdateInterval() bool {
	if o != nil && o.UpdateInterval != nil {
		return true
	}

	return false
}

// SetUpdateInterval gets a reference to the given int32 and assigns it to the UpdateInterval field.
func (o *SiteAllOfNameResolutionIllumioResolvers) SetUpdateInterval(v int32) {
	o.UpdateInterval = &v
}

// GetOrgId returns the OrgId field value
func (o *SiteAllOfNameResolutionIllumioResolvers) GetOrgId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value
// and a boolean to check if the value has been set.
func (o *SiteAllOfNameResolutionIllumioResolvers) GetOrgIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgId, true
}

// SetOrgId sets field value
func (o *SiteAllOfNameResolutionIllumioResolvers) SetOrgId(v string) {
	o.OrgId = v
}

// GetHostname returns the Hostname field value
func (o *SiteAllOfNameResolutionIllumioResolvers) GetHostname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value
// and a boolean to check if the value has been set.
func (o *SiteAllOfNameResolutionIllumioResolvers) GetHostnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hostname, true
}

// SetHostname sets field value
func (o *SiteAllOfNameResolutionIllumioResolvers) SetHostname(v string) {
	o.Hostname = v
}

// GetPort returns the Port field value
func (o *SiteAllOfNameResolutionIllumioResolvers) GetPort() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Port
}

// GetPortOk returns a tuple with the Port field value
// and a boolean to check if the value has been set.
func (o *SiteAllOfNameResolutionIllumioResolvers) GetPortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Port, true
}

// SetPort sets field value
func (o *SiteAllOfNameResolutionIllumioResolvers) SetPort(v int32) {
	o.Port = v
}

// GetUsername returns the Username field value
func (o *SiteAllOfNameResolutionIllumioResolvers) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *SiteAllOfNameResolutionIllumioResolvers) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *SiteAllOfNameResolutionIllumioResolvers) SetUsername(v string) {
	o.Username = v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *SiteAllOfNameResolutionIllumioResolvers) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteAllOfNameResolutionIllumioResolvers) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *SiteAllOfNameResolutionIllumioResolvers) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *SiteAllOfNameResolutionIllumioResolvers) SetPassword(v string) {
	o.Password = &v
}

func (o SiteAllOfNameResolutionIllumioResolvers) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.UpdateInterval != nil {
		toSerialize["updateInterval"] = o.UpdateInterval
	}
	if true {
		toSerialize["orgId"] = o.OrgId
	}
	if true {
		toSerialize["hostname"] = o.Hostname
	}
	if true {
		toSerialize["port"] = o.Port
	}
	if true {
		toSerialize["username"] = o.Username
	}
	if o.Password != nil {
		toSerialize["password"] = o.Password
	}
	return json.Marshal(toSerialize)
}

type NullableSiteAllOfNameResolutionIllumioResolvers struct {
	value *SiteAllOfNameResolutionIllumioResolvers
	isSet bool
}

func (v NullableSiteAllOfNameResolutionIllumioResolvers) Get() *SiteAllOfNameResolutionIllumioResolvers {
	return v.value
}

func (v *NullableSiteAllOfNameResolutionIllumioResolvers) Set(val *SiteAllOfNameResolutionIllumioResolvers) {
	v.value = val
	v.isSet = true
}

func (v NullableSiteAllOfNameResolutionIllumioResolvers) IsSet() bool {
	return v.isSet
}

func (v *NullableSiteAllOfNameResolutionIllumioResolvers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSiteAllOfNameResolutionIllumioResolvers(val *SiteAllOfNameResolutionIllumioResolvers) *NullableSiteAllOfNameResolutionIllumioResolvers {
	return &NullableSiteAllOfNameResolutionIllumioResolvers{value: val, isSet: true}
}

func (v NullableSiteAllOfNameResolutionIllumioResolvers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSiteAllOfNameResolutionIllumioResolvers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
