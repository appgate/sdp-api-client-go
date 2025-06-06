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

// AdministrativePrivilege Administrative Privilege item. Use type-target-map API to get the details on which types are valid for which targets and their scopes.
type AdministrativePrivilege struct {
	// The type of the Privilege defines the possible administrator actions.
	Type string `json:"type"`
	// The target of the Privilege defines the possible target objects for that type.
	Target string                        `json:"target"`
	Scope  *AdministrativePrivilegeScope `json:"scope,omitempty"`
	// The items in this list would be added automatically to the newly created objects' tags. Only applicable on \"Create\" type and targets with tagging capability. This field must be omitted if not applicable.
	DefaultTags []string `json:"defaultTags,omitempty"`
	// Privilege for changing Appliance Functions. Only applicable on \"AssignFunction\" type with Appliance or All target. This field must be omitted if not applicable.
	Functions []ApplianceFunction `json:"functions,omitempty"`
}

// NewAdministrativePrivilege instantiates a new AdministrativePrivilege object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdministrativePrivilege(type_ string, target string) *AdministrativePrivilege {
	this := AdministrativePrivilege{}
	this.Type = type_
	this.Target = target
	return &this
}

// NewAdministrativePrivilegeWithDefaults instantiates a new AdministrativePrivilege object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdministrativePrivilegeWithDefaults() *AdministrativePrivilege {
	this := AdministrativePrivilege{}
	return &this
}

// GetType returns the Type field value
func (o *AdministrativePrivilege) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AdministrativePrivilege) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AdministrativePrivilege) SetType(v string) {
	o.Type = v
}

// GetTarget returns the Target field value
func (o *AdministrativePrivilege) GetTarget() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Target
}

// GetTargetOk returns a tuple with the Target field value
// and a boolean to check if the value has been set.
func (o *AdministrativePrivilege) GetTargetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Target, true
}

// SetTarget sets field value
func (o *AdministrativePrivilege) SetTarget(v string) {
	o.Target = v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *AdministrativePrivilege) GetScope() AdministrativePrivilegeScope {
	if o == nil || o.Scope == nil {
		var ret AdministrativePrivilegeScope
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdministrativePrivilege) GetScopeOk() (*AdministrativePrivilegeScope, bool) {
	if o == nil || o.Scope == nil {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *AdministrativePrivilege) HasScope() bool {
	if o != nil && o.Scope != nil {
		return true
	}

	return false
}

// SetScope gets a reference to the given AdministrativePrivilegeScope and assigns it to the Scope field.
func (o *AdministrativePrivilege) SetScope(v AdministrativePrivilegeScope) {
	o.Scope = &v
}

// GetDefaultTags returns the DefaultTags field value if set, zero value otherwise.
func (o *AdministrativePrivilege) GetDefaultTags() []string {
	if o == nil || o.DefaultTags == nil {
		var ret []string
		return ret
	}
	return o.DefaultTags
}

// GetDefaultTagsOk returns a tuple with the DefaultTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdministrativePrivilege) GetDefaultTagsOk() ([]string, bool) {
	if o == nil || o.DefaultTags == nil {
		return nil, false
	}
	return o.DefaultTags, true
}

// HasDefaultTags returns a boolean if a field has been set.
func (o *AdministrativePrivilege) HasDefaultTags() bool {
	if o != nil && o.DefaultTags != nil {
		return true
	}

	return false
}

// SetDefaultTags gets a reference to the given []string and assigns it to the DefaultTags field.
func (o *AdministrativePrivilege) SetDefaultTags(v []string) {
	o.DefaultTags = v
}

// GetFunctions returns the Functions field value if set, zero value otherwise.
func (o *AdministrativePrivilege) GetFunctions() []ApplianceFunction {
	if o == nil || o.Functions == nil {
		var ret []ApplianceFunction
		return ret
	}
	return o.Functions
}

// GetFunctionsOk returns a tuple with the Functions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdministrativePrivilege) GetFunctionsOk() ([]ApplianceFunction, bool) {
	if o == nil || o.Functions == nil {
		return nil, false
	}
	return o.Functions, true
}

// HasFunctions returns a boolean if a field has been set.
func (o *AdministrativePrivilege) HasFunctions() bool {
	if o != nil && o.Functions != nil {
		return true
	}

	return false
}

// SetFunctions gets a reference to the given []ApplianceFunction and assigns it to the Functions field.
func (o *AdministrativePrivilege) SetFunctions(v []ApplianceFunction) {
	o.Functions = v
}

func (o AdministrativePrivilege) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["target"] = o.Target
	}
	if o.Scope != nil {
		toSerialize["scope"] = o.Scope
	}
	if o.DefaultTags != nil {
		toSerialize["defaultTags"] = o.DefaultTags
	}
	if o.Functions != nil {
		toSerialize["functions"] = o.Functions
	}
	return json.Marshal(toSerialize)
}

type NullableAdministrativePrivilege struct {
	value *AdministrativePrivilege
	isSet bool
}

func (v NullableAdministrativePrivilege) Get() *AdministrativePrivilege {
	return v.value
}

func (v *NullableAdministrativePrivilege) Set(val *AdministrativePrivilege) {
	v.value = val
	v.isSet = true
}

func (v NullableAdministrativePrivilege) IsSet() bool {
	return v.isSet
}

func (v *NullableAdministrativePrivilege) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdministrativePrivilege(val *AdministrativePrivilege) *NullableAdministrativePrivilege {
	return &NullableAdministrativePrivilege{value: val, isSet: true}
}

func (v NullableAdministrativePrivilege) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdministrativePrivilege) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
