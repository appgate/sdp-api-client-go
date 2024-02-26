/*
Appgate SDP Controller REST API

# About   This specification documents the REST API calls for the Appgate SDP Controller.    Please refer to the REST API chapter in the manual or contact Appgate support with any questions about   this functionality. # Getting Started   Requirements for API scripting:   - Access to the Admin/API TLS Connection (default port 8443) of a Controller appliance.     (https://sdphelp.appgate.com/adminguide/appliance-function-configure.html?anchor=admin-api)   - An API user with relevant permissions.     (https://sdphelp.appgate.com/adminguide/administrative-roles-configure.html)   - In order to use the simple login API, Admin MFA must be disabled or the API user must be excluded.     (https://sdphelp.appgate.com/adminguide/mfa-for-admins.html) # Base path   HTTPS requests must be sent to the Admin Interface hostname and port, with **_/admin** path.    For example: **https://appgate.company.com:8443/admin**    All requests must have the **Accept** header as:    **application/vnd.appgate.peer-v19+json** # API Conventions   API conventions are  important to understand and follow strictly.    - While updating objects (via PUT), entire object must be sent with all fields.     - For example, in order to add a remedy method to the condition below:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": []       }       ```     - send the entire object with updated and non-updated fields:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": [{\"type\": \"DisplayMessage\", \"message\": \"test message\"}]       }       ```    - In case Controller returns an error (non-2xx HTTP status code), response body is JSON.     The \"message\" field contains information about the error.     HTTP 422 \"Unprocessable Entity\" has extra `errors` field to list all the issues with specific fields.    - Empty string (\"\") is considered a different value than \"null\" or field being omitted from JSON.     Omitting the field is recommended if no value is intended.     Empty string (\"\") will be almost always rejected as invalid value.    - There are common pattern between many objects:     - **Configuration Objects**: There are many objects with common fields, namely \"id\", \"name\", \"notes\", \"created\"       and \"updated\". These entities are listed, queried, created, updated and deleted in a similar fashion.     - **Distinguished Name**: Users and Devices are identified with what is called Distinguished Names, as used in        LDAP. The distinguished format that identifies a device and a user combination is        \"CN=\\<Device ID\\>,CN=\\<username\\>,OU=\\<Identity Provider Name\\>\". Some objects have the        \"userDistinguishedName\" field, which does not include the CN for Device ID.        This identifies a user on every device.

API version: API version 19.2
Contact: appgatesdp.support@appgate.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// LdapProviderAllOf Represents an LDAP Identity Provider.
type LdapProviderAllOf struct {
	// Hostnames/IP addresses to connect.
	Hostnames []string `json:"hostnames"`
	// Port to connect.
	Port int32 `json:"port"`
	// Whether to use LDAPS protocol or not.
	SslEnabled *bool `json:"sslEnabled,omitempty"`
	// The Distinguished Name to login to LDAP and query users with.
	AdminDistinguishedName string `json:"adminDistinguishedName"`
	// The password to login to LDAP and query users with. Required on creation.
	AdminPassword *string `json:"adminPassword,omitempty"`
	// The subset of the LDAP server to search users from. If not set, root of the server is used.
	BaseDn *string `json:"baseDn,omitempty"`
	// The object class of the users to be authenticated and queried. This field is translated to '(objectClass=value)' as user filter. Deprecated as of 6.2. Use userFilter instead.
	// Deprecated
	ObjectClass *string `json:"objectClass,omitempty"`
	// The LDAP filter to apply when searching for user during authentication.
	UserFilter *string `json:"userFilter,omitempty"`
	// The name of the attribute to get the exact username from the LDAP server.
	UsernameAttribute *string `json:"usernameAttribute,omitempty"`
	// The filter to use while querying users' nested groups.
	MembershipFilter *string `json:"membershipFilter,omitempty"`
	// The subset of the LDAP server to search groups from. If not set, \"baseDn\" is used.
	MembershipBaseDn *string                           `json:"membershipBaseDn,omitempty"`
	PasswordWarning  *LdapProviderAllOfPasswordWarning `json:"passwordWarning,omitempty"`
}

// NewLdapProviderAllOf instantiates a new LdapProviderAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLdapProviderAllOf(hostnames []string, port int32, adminDistinguishedName string) *LdapProviderAllOf {
	this := LdapProviderAllOf{}
	this.Hostnames = hostnames
	this.Port = port
	var sslEnabled bool = false
	this.SslEnabled = &sslEnabled
	this.AdminDistinguishedName = adminDistinguishedName
	var usernameAttribute string = "sAMAccountName"
	this.UsernameAttribute = &usernameAttribute
	var membershipFilter string = "(objectCategory=group)"
	this.MembershipFilter = &membershipFilter
	return &this
}

// NewLdapProviderAllOfWithDefaults instantiates a new LdapProviderAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLdapProviderAllOfWithDefaults() *LdapProviderAllOf {
	this := LdapProviderAllOf{}
	var sslEnabled bool = false
	this.SslEnabled = &sslEnabled
	var usernameAttribute string = "sAMAccountName"
	this.UsernameAttribute = &usernameAttribute
	var membershipFilter string = "(objectCategory=group)"
	this.MembershipFilter = &membershipFilter
	return &this
}

// GetHostnames returns the Hostnames field value
func (o *LdapProviderAllOf) GetHostnames() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Hostnames
}

// GetHostnamesOk returns a tuple with the Hostnames field value
// and a boolean to check if the value has been set.
func (o *LdapProviderAllOf) GetHostnamesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Hostnames, true
}

// SetHostnames sets field value
func (o *LdapProviderAllOf) SetHostnames(v []string) {
	o.Hostnames = v
}

// GetPort returns the Port field value
func (o *LdapProviderAllOf) GetPort() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Port
}

// GetPortOk returns a tuple with the Port field value
// and a boolean to check if the value has been set.
func (o *LdapProviderAllOf) GetPortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Port, true
}

// SetPort sets field value
func (o *LdapProviderAllOf) SetPort(v int32) {
	o.Port = v
}

// GetSslEnabled returns the SslEnabled field value if set, zero value otherwise.
func (o *LdapProviderAllOf) GetSslEnabled() bool {
	if o == nil || o.SslEnabled == nil {
		var ret bool
		return ret
	}
	return *o.SslEnabled
}

// GetSslEnabledOk returns a tuple with the SslEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapProviderAllOf) GetSslEnabledOk() (*bool, bool) {
	if o == nil || o.SslEnabled == nil {
		return nil, false
	}
	return o.SslEnabled, true
}

// HasSslEnabled returns a boolean if a field has been set.
func (o *LdapProviderAllOf) HasSslEnabled() bool {
	if o != nil && o.SslEnabled != nil {
		return true
	}

	return false
}

// SetSslEnabled gets a reference to the given bool and assigns it to the SslEnabled field.
func (o *LdapProviderAllOf) SetSslEnabled(v bool) {
	o.SslEnabled = &v
}

// GetAdminDistinguishedName returns the AdminDistinguishedName field value
func (o *LdapProviderAllOf) GetAdminDistinguishedName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AdminDistinguishedName
}

// GetAdminDistinguishedNameOk returns a tuple with the AdminDistinguishedName field value
// and a boolean to check if the value has been set.
func (o *LdapProviderAllOf) GetAdminDistinguishedNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdminDistinguishedName, true
}

// SetAdminDistinguishedName sets field value
func (o *LdapProviderAllOf) SetAdminDistinguishedName(v string) {
	o.AdminDistinguishedName = v
}

// GetAdminPassword returns the AdminPassword field value if set, zero value otherwise.
func (o *LdapProviderAllOf) GetAdminPassword() string {
	if o == nil || o.AdminPassword == nil {
		var ret string
		return ret
	}
	return *o.AdminPassword
}

// GetAdminPasswordOk returns a tuple with the AdminPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapProviderAllOf) GetAdminPasswordOk() (*string, bool) {
	if o == nil || o.AdminPassword == nil {
		return nil, false
	}
	return o.AdminPassword, true
}

// HasAdminPassword returns a boolean if a field has been set.
func (o *LdapProviderAllOf) HasAdminPassword() bool {
	if o != nil && o.AdminPassword != nil {
		return true
	}

	return false
}

// SetAdminPassword gets a reference to the given string and assigns it to the AdminPassword field.
func (o *LdapProviderAllOf) SetAdminPassword(v string) {
	o.AdminPassword = &v
}

// GetBaseDn returns the BaseDn field value if set, zero value otherwise.
func (o *LdapProviderAllOf) GetBaseDn() string {
	if o == nil || o.BaseDn == nil {
		var ret string
		return ret
	}
	return *o.BaseDn
}

// GetBaseDnOk returns a tuple with the BaseDn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapProviderAllOf) GetBaseDnOk() (*string, bool) {
	if o == nil || o.BaseDn == nil {
		return nil, false
	}
	return o.BaseDn, true
}

// HasBaseDn returns a boolean if a field has been set.
func (o *LdapProviderAllOf) HasBaseDn() bool {
	if o != nil && o.BaseDn != nil {
		return true
	}

	return false
}

// SetBaseDn gets a reference to the given string and assigns it to the BaseDn field.
func (o *LdapProviderAllOf) SetBaseDn(v string) {
	o.BaseDn = &v
}

// GetObjectClass returns the ObjectClass field value if set, zero value otherwise.
// Deprecated
func (o *LdapProviderAllOf) GetObjectClass() string {
	if o == nil || o.ObjectClass == nil {
		var ret string
		return ret
	}
	return *o.ObjectClass
}

// GetObjectClassOk returns a tuple with the ObjectClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *LdapProviderAllOf) GetObjectClassOk() (*string, bool) {
	if o == nil || o.ObjectClass == nil {
		return nil, false
	}
	return o.ObjectClass, true
}

// HasObjectClass returns a boolean if a field has been set.
func (o *LdapProviderAllOf) HasObjectClass() bool {
	if o != nil && o.ObjectClass != nil {
		return true
	}

	return false
}

// SetObjectClass gets a reference to the given string and assigns it to the ObjectClass field.
// Deprecated
func (o *LdapProviderAllOf) SetObjectClass(v string) {
	o.ObjectClass = &v
}

// GetUserFilter returns the UserFilter field value if set, zero value otherwise.
func (o *LdapProviderAllOf) GetUserFilter() string {
	if o == nil || o.UserFilter == nil {
		var ret string
		return ret
	}
	return *o.UserFilter
}

// GetUserFilterOk returns a tuple with the UserFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapProviderAllOf) GetUserFilterOk() (*string, bool) {
	if o == nil || o.UserFilter == nil {
		return nil, false
	}
	return o.UserFilter, true
}

// HasUserFilter returns a boolean if a field has been set.
func (o *LdapProviderAllOf) HasUserFilter() bool {
	if o != nil && o.UserFilter != nil {
		return true
	}

	return false
}

// SetUserFilter gets a reference to the given string and assigns it to the UserFilter field.
func (o *LdapProviderAllOf) SetUserFilter(v string) {
	o.UserFilter = &v
}

// GetUsernameAttribute returns the UsernameAttribute field value if set, zero value otherwise.
func (o *LdapProviderAllOf) GetUsernameAttribute() string {
	if o == nil || o.UsernameAttribute == nil {
		var ret string
		return ret
	}
	return *o.UsernameAttribute
}

// GetUsernameAttributeOk returns a tuple with the UsernameAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapProviderAllOf) GetUsernameAttributeOk() (*string, bool) {
	if o == nil || o.UsernameAttribute == nil {
		return nil, false
	}
	return o.UsernameAttribute, true
}

// HasUsernameAttribute returns a boolean if a field has been set.
func (o *LdapProviderAllOf) HasUsernameAttribute() bool {
	if o != nil && o.UsernameAttribute != nil {
		return true
	}

	return false
}

// SetUsernameAttribute gets a reference to the given string and assigns it to the UsernameAttribute field.
func (o *LdapProviderAllOf) SetUsernameAttribute(v string) {
	o.UsernameAttribute = &v
}

// GetMembershipFilter returns the MembershipFilter field value if set, zero value otherwise.
func (o *LdapProviderAllOf) GetMembershipFilter() string {
	if o == nil || o.MembershipFilter == nil {
		var ret string
		return ret
	}
	return *o.MembershipFilter
}

// GetMembershipFilterOk returns a tuple with the MembershipFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapProviderAllOf) GetMembershipFilterOk() (*string, bool) {
	if o == nil || o.MembershipFilter == nil {
		return nil, false
	}
	return o.MembershipFilter, true
}

// HasMembershipFilter returns a boolean if a field has been set.
func (o *LdapProviderAllOf) HasMembershipFilter() bool {
	if o != nil && o.MembershipFilter != nil {
		return true
	}

	return false
}

// SetMembershipFilter gets a reference to the given string and assigns it to the MembershipFilter field.
func (o *LdapProviderAllOf) SetMembershipFilter(v string) {
	o.MembershipFilter = &v
}

// GetMembershipBaseDn returns the MembershipBaseDn field value if set, zero value otherwise.
func (o *LdapProviderAllOf) GetMembershipBaseDn() string {
	if o == nil || o.MembershipBaseDn == nil {
		var ret string
		return ret
	}
	return *o.MembershipBaseDn
}

// GetMembershipBaseDnOk returns a tuple with the MembershipBaseDn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapProviderAllOf) GetMembershipBaseDnOk() (*string, bool) {
	if o == nil || o.MembershipBaseDn == nil {
		return nil, false
	}
	return o.MembershipBaseDn, true
}

// HasMembershipBaseDn returns a boolean if a field has been set.
func (o *LdapProviderAllOf) HasMembershipBaseDn() bool {
	if o != nil && o.MembershipBaseDn != nil {
		return true
	}

	return false
}

// SetMembershipBaseDn gets a reference to the given string and assigns it to the MembershipBaseDn field.
func (o *LdapProviderAllOf) SetMembershipBaseDn(v string) {
	o.MembershipBaseDn = &v
}

// GetPasswordWarning returns the PasswordWarning field value if set, zero value otherwise.
func (o *LdapProviderAllOf) GetPasswordWarning() LdapProviderAllOfPasswordWarning {
	if o == nil || o.PasswordWarning == nil {
		var ret LdapProviderAllOfPasswordWarning
		return ret
	}
	return *o.PasswordWarning
}

// GetPasswordWarningOk returns a tuple with the PasswordWarning field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapProviderAllOf) GetPasswordWarningOk() (*LdapProviderAllOfPasswordWarning, bool) {
	if o == nil || o.PasswordWarning == nil {
		return nil, false
	}
	return o.PasswordWarning, true
}

// HasPasswordWarning returns a boolean if a field has been set.
func (o *LdapProviderAllOf) HasPasswordWarning() bool {
	if o != nil && o.PasswordWarning != nil {
		return true
	}

	return false
}

// SetPasswordWarning gets a reference to the given LdapProviderAllOfPasswordWarning and assigns it to the PasswordWarning field.
func (o *LdapProviderAllOf) SetPasswordWarning(v LdapProviderAllOfPasswordWarning) {
	o.PasswordWarning = &v
}

func (o LdapProviderAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["hostnames"] = o.Hostnames
	}
	if true {
		toSerialize["port"] = o.Port
	}
	if o.SslEnabled != nil {
		toSerialize["sslEnabled"] = o.SslEnabled
	}
	if true {
		toSerialize["adminDistinguishedName"] = o.AdminDistinguishedName
	}
	if o.AdminPassword != nil {
		toSerialize["adminPassword"] = o.AdminPassword
	}
	if o.BaseDn != nil {
		toSerialize["baseDn"] = o.BaseDn
	}
	if o.ObjectClass != nil {
		toSerialize["objectClass"] = o.ObjectClass
	}
	if o.UserFilter != nil {
		toSerialize["userFilter"] = o.UserFilter
	}
	if o.UsernameAttribute != nil {
		toSerialize["usernameAttribute"] = o.UsernameAttribute
	}
	if o.MembershipFilter != nil {
		toSerialize["membershipFilter"] = o.MembershipFilter
	}
	if o.MembershipBaseDn != nil {
		toSerialize["membershipBaseDn"] = o.MembershipBaseDn
	}
	if o.PasswordWarning != nil {
		toSerialize["passwordWarning"] = o.PasswordWarning
	}
	return json.Marshal(toSerialize)
}

type NullableLdapProviderAllOf struct {
	value *LdapProviderAllOf
	isSet bool
}

func (v NullableLdapProviderAllOf) Get() *LdapProviderAllOf {
	return v.value
}

func (v *NullableLdapProviderAllOf) Set(val *LdapProviderAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableLdapProviderAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableLdapProviderAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLdapProviderAllOf(val *LdapProviderAllOf) *NullableLdapProviderAllOf {
	return &NullableLdapProviderAllOf{value: val, isSet: true}
}

func (v NullableLdapProviderAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLdapProviderAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
