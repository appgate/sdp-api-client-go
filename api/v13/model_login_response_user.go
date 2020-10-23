/*
 * AppGate SDP Controller REST API
 *
 * # About   This specification documents the REST API calls for the AppGate SDP Controller.    Please refer to the Integration chapter in the manual or contact AppGate support with any questions about   this functionality. # Getting Started   Requirements for API scripting:   - Access to the peer interface (default port 444) of a Controller appliance.     (https://sdphelp.appgate.com/adminguide/appliances-configure.html?anchor=peer)   - An API user with relevant permissions.     (https://sdphelp.appgate.com/adminguide/administrative-roles-configure.html)   - In order to use the simple login API, Admin MFA must be disabled or the API user must be excluded.     (https://sdphelp.appgate.com/adminguide/mfa-for-admins.html) # Base path   HTTPS requests must be sent to the Peer Interface hostname and port, with **_/admin** path.    For example: **https://appgate.company.com:444/admin**    All requests must have the **Accept** header as:    **application/vnd.appgate.peer-v13+json** # API Conventions   API conventions are  important to understand and follow strictly.    - While updating objects (via PUT), entire object must be sent with all fields.     - For example, if in order to add a remedy method to the condition below:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": []       }       ```     - send the entire object with updated and non-updated fields:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": [{\"type\": \"DisplayMessage\", \"message\": \"test message\"}]       }       ```    - In case Controller returns an error (non-2xx HTTP status code), response body is JSON.     The \"message\" field contains information about the error.     HTTP 422 \"Unprocessable Entity\" has extra `errors` field to list all the issues with specific fields.    - Empty string (\"\") is considered a different value than \"null\" or field being omitted from JSON.     Omitting the field is recommend if no value is intended.     Empty string (\"\") will be almost always rejected as invalid value.    - There are common pattern between many objects:     - **Configuration Objects**: There are many objects with common fields, namely \"id\", \"name\", \"notes\", \"created\"       and \"updated\". These entities are listed, queried, created, updated and deleted in a similar fashion.     - **Distinguished Name**: Users and Devices are identified with what is called Distinguished Names, as used in        LDAP. The distinguished format that identifies a device and a user combination is        \"CN=\\<Device ID\\>,CN=\\<username\\>,OU=\\<Identity Provider Name\\>\". Some objects have the        \"userDistinguishedName\" field, which does not include the CN for Device ID.        This identifies a user on every device.
 *
 * API version: API version 13
 * Contact: appgatesdp.support@appgate.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"encoding/json"
)

// LoginResponseUser Information about logged in user, such as username and email address, if exists.
type LoginResponseUser struct {
	// Username.
	Name *string `json:"name,omitempty"`
	// If true, it is not possible to complete login process without providing MFA.
	NeedTwoFactorAuth *bool `json:"needTwoFactorAuth,omitempty"`
	// Whether there is a LogServer deployed and the user has privileges to access to it.
	CanAccessAuditLogs *bool `json:"canAccessAuditLogs,omitempty"`
	// The privileges the user has.
	Privileges *[]AdministrativePrivilege `json:"privileges,omitempty"`
}

// NewLoginResponseUser instantiates a new LoginResponseUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoginResponseUser() *LoginResponseUser {
	this := LoginResponseUser{}
	return &this
}

// NewLoginResponseUserWithDefaults instantiates a new LoginResponseUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoginResponseUserWithDefaults() *LoginResponseUser {
	this := LoginResponseUser{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *LoginResponseUser) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoginResponseUser) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *LoginResponseUser) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *LoginResponseUser) SetName(v string) {
	o.Name = &v
}

// GetNeedTwoFactorAuth returns the NeedTwoFactorAuth field value if set, zero value otherwise.
func (o *LoginResponseUser) GetNeedTwoFactorAuth() bool {
	if o == nil || o.NeedTwoFactorAuth == nil {
		var ret bool
		return ret
	}
	return *o.NeedTwoFactorAuth
}

// GetNeedTwoFactorAuthOk returns a tuple with the NeedTwoFactorAuth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoginResponseUser) GetNeedTwoFactorAuthOk() (*bool, bool) {
	if o == nil || o.NeedTwoFactorAuth == nil {
		return nil, false
	}
	return o.NeedTwoFactorAuth, true
}

// HasNeedTwoFactorAuth returns a boolean if a field has been set.
func (o *LoginResponseUser) HasNeedTwoFactorAuth() bool {
	if o != nil && o.NeedTwoFactorAuth != nil {
		return true
	}

	return false
}

// SetNeedTwoFactorAuth gets a reference to the given bool and assigns it to the NeedTwoFactorAuth field.
func (o *LoginResponseUser) SetNeedTwoFactorAuth(v bool) {
	o.NeedTwoFactorAuth = &v
}

// GetCanAccessAuditLogs returns the CanAccessAuditLogs field value if set, zero value otherwise.
func (o *LoginResponseUser) GetCanAccessAuditLogs() bool {
	if o == nil || o.CanAccessAuditLogs == nil {
		var ret bool
		return ret
	}
	return *o.CanAccessAuditLogs
}

// GetCanAccessAuditLogsOk returns a tuple with the CanAccessAuditLogs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoginResponseUser) GetCanAccessAuditLogsOk() (*bool, bool) {
	if o == nil || o.CanAccessAuditLogs == nil {
		return nil, false
	}
	return o.CanAccessAuditLogs, true
}

// HasCanAccessAuditLogs returns a boolean if a field has been set.
func (o *LoginResponseUser) HasCanAccessAuditLogs() bool {
	if o != nil && o.CanAccessAuditLogs != nil {
		return true
	}

	return false
}

// SetCanAccessAuditLogs gets a reference to the given bool and assigns it to the CanAccessAuditLogs field.
func (o *LoginResponseUser) SetCanAccessAuditLogs(v bool) {
	o.CanAccessAuditLogs = &v
}

// GetPrivileges returns the Privileges field value if set, zero value otherwise.
func (o *LoginResponseUser) GetPrivileges() []AdministrativePrivilege {
	if o == nil || o.Privileges == nil {
		var ret []AdministrativePrivilege
		return ret
	}
	return *o.Privileges
}

// GetPrivilegesOk returns a tuple with the Privileges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoginResponseUser) GetPrivilegesOk() (*[]AdministrativePrivilege, bool) {
	if o == nil || o.Privileges == nil {
		return nil, false
	}
	return o.Privileges, true
}

// HasPrivileges returns a boolean if a field has been set.
func (o *LoginResponseUser) HasPrivileges() bool {
	if o != nil && o.Privileges != nil {
		return true
	}

	return false
}

// SetPrivileges gets a reference to the given []AdministrativePrivilege and assigns it to the Privileges field.
func (o *LoginResponseUser) SetPrivileges(v []AdministrativePrivilege) {
	o.Privileges = &v
}

func (o LoginResponseUser) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.NeedTwoFactorAuth != nil {
		toSerialize["needTwoFactorAuth"] = o.NeedTwoFactorAuth
	}
	if o.CanAccessAuditLogs != nil {
		toSerialize["canAccessAuditLogs"] = o.CanAccessAuditLogs
	}
	if o.Privileges != nil {
		toSerialize["privileges"] = o.Privileges
	}
	return json.Marshal(toSerialize)
}

type NullableLoginResponseUser struct {
	value *LoginResponseUser
	isSet bool
}

func (v NullableLoginResponseUser) Get() *LoginResponseUser {
	return v.value
}

func (v *NullableLoginResponseUser) Set(val *LoginResponseUser) {
	v.value = val
	v.isSet = true
}

func (v NullableLoginResponseUser) IsSet() bool {
	return v.isSet
}

func (v *NullableLoginResponseUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoginResponseUser(val *LoginResponseUser) *NullableLoginResponseUser {
	return &NullableLoginResponseUser{value: val, isSet: true}
}

func (v NullableLoginResponseUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoginResponseUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
