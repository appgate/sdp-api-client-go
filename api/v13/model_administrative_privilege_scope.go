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

// AdministrativePrivilegeScope The scope of the Privilege. Only applicable to certain type-target combinations. Some types depends on the IdP/MFA type, such as GetUserAttributes. This field must be omitted if not applicable.
type AdministrativePrivilegeScope struct {
	// 'If \"true\", all objects are accessible. For example, \"type: Edit - target: Condition - scope.all: true\" means the administrator can edit all Conditions in the system.'
	All *bool `json:"all,omitempty"`
	// Specific object IDs this Privilege would have access to.
	Ids *[]string `json:"ids,omitempty"`
	// Object tags this privilege would have access to.
	Tags *[]string `json:"tags,omitempty"`
}

// NewAdministrativePrivilegeScope instantiates a new AdministrativePrivilegeScope object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdministrativePrivilegeScope() *AdministrativePrivilegeScope {
	this := AdministrativePrivilegeScope{}
	return &this
}

// NewAdministrativePrivilegeScopeWithDefaults instantiates a new AdministrativePrivilegeScope object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdministrativePrivilegeScopeWithDefaults() *AdministrativePrivilegeScope {
	this := AdministrativePrivilegeScope{}
	return &this
}

// GetAll returns the All field value if set, zero value otherwise.
func (o *AdministrativePrivilegeScope) GetAll() bool {
	if o == nil || o.All == nil {
		var ret bool
		return ret
	}
	return *o.All
}

// GetAllOk returns a tuple with the All field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdministrativePrivilegeScope) GetAllOk() (*bool, bool) {
	if o == nil || o.All == nil {
		return nil, false
	}
	return o.All, true
}

// HasAll returns a boolean if a field has been set.
func (o *AdministrativePrivilegeScope) HasAll() bool {
	if o != nil && o.All != nil {
		return true
	}

	return false
}

// SetAll gets a reference to the given bool and assigns it to the All field.
func (o *AdministrativePrivilegeScope) SetAll(v bool) {
	o.All = &v
}

// GetIds returns the Ids field value if set, zero value otherwise.
func (o *AdministrativePrivilegeScope) GetIds() []string {
	if o == nil || o.Ids == nil {
		var ret []string
		return ret
	}
	return *o.Ids
}

// GetIdsOk returns a tuple with the Ids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdministrativePrivilegeScope) GetIdsOk() (*[]string, bool) {
	if o == nil || o.Ids == nil {
		return nil, false
	}
	return o.Ids, true
}

// HasIds returns a boolean if a field has been set.
func (o *AdministrativePrivilegeScope) HasIds() bool {
	if o != nil && o.Ids != nil {
		return true
	}

	return false
}

// SetIds gets a reference to the given []string and assigns it to the Ids field.
func (o *AdministrativePrivilegeScope) SetIds(v []string) {
	o.Ids = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *AdministrativePrivilegeScope) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdministrativePrivilegeScope) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *AdministrativePrivilegeScope) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *AdministrativePrivilegeScope) SetTags(v []string) {
	o.Tags = &v
}

func (o AdministrativePrivilegeScope) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.All != nil {
		toSerialize["all"] = o.All
	}
	if o.Ids != nil {
		toSerialize["ids"] = o.Ids
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	return json.Marshal(toSerialize)
}

type NullableAdministrativePrivilegeScope struct {
	value *AdministrativePrivilegeScope
	isSet bool
}

func (v NullableAdministrativePrivilegeScope) Get() *AdministrativePrivilegeScope {
	return v.value
}

func (v *NullableAdministrativePrivilegeScope) Set(val *AdministrativePrivilegeScope) {
	v.value = val
	v.isSet = true
}

func (v NullableAdministrativePrivilegeScope) IsSet() bool {
	return v.isSet
}

func (v *NullableAdministrativePrivilegeScope) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdministrativePrivilegeScope(val *AdministrativePrivilegeScope) *NullableAdministrativePrivilegeScope {
	return &NullableAdministrativePrivilegeScope{value: val, isSet: true}
}

func (v NullableAdministrativePrivilegeScope) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdministrativePrivilegeScope) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
