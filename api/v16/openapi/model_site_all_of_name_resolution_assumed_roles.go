/*
Appgate SDP Controller REST API

# About   This specification documents the REST API calls for the Appgate SDP Controller.    Please refer to the REST API chapter in the manual or contact Appgate support with any questions about   this functionality. # Getting Started   Requirements for API scripting:   - Access to the Admin/API TLS Connection (default port 8443) of a Controller appliance.     (https://sdphelp.appgate.com/adminguide/appliance-function-configure.html?anchor=admin-api)   - An API user with relevant permissions.     (https://sdphelp.appgate.com/adminguide/administrative-roles-configure.html)   - In order to use the simple login API, Admin MFA must be disabled or the API user must be excluded.     (https://sdphelp.appgate.com/adminguide/mfa-for-admins.html) # Base path   HTTPS requests must be sent to the Admin Interface hostname and port, with **_/admin** path.    For example: **https://appgate.company.com:8443/admin**    All requests must have the **Accept** header as:    **application/vnd.appgate.peer-v16+json** # API Conventions   API conventions are  important to understand and follow strictly.    - While updating objects (via PUT), entire object must be sent with all fields.     - For example, in order to add a remedy method to the condition below:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": []       }       ```     - send the entire object with updated and non-updated fields:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": [{\"type\": \"DisplayMessage\", \"message\": \"test message\"}]       }       ```    - In case Controller returns an error (non-2xx HTTP status code), response body is JSON.     The \"message\" field contains information about the error.     HTTP 422 \"Unprocessable Entity\" has extra `errors` field to list all the issues with specific fields.    - Empty string (\"\") is considered a different value than \"null\" or field being omitted from JSON.     Omitting the field is recommend if no value is intended.     Empty string (\"\") will be almost always rejected as invalid value.    - There are common pattern between many objects:     - **Configuration Objects**: There are many objects with common fields, namely \"id\", \"name\", \"notes\", \"created\"       and \"updated\". These entities are listed, queried, created, updated and deleted in a similar fashion.     - **Distinguished Name**: Users and Devices are identified with what is called Distinguished Names, as used in        LDAP. The distinguished format that identifies a device and a user combination is        \"CN=\\<Device ID\\>,CN=\\<username\\>,OU=\\<Identity Provider Name\\>\". Some objects have the        \"userDistinguishedName\" field, which does not include the CN for Device ID.        This identifies a user on every device.

API version: API version 16.2
Contact: appgatesdp.support@appgate.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// SiteAllOfNameResolutionAssumedRoles struct for SiteAllOfNameResolutionAssumedRoles
type SiteAllOfNameResolutionAssumedRoles struct {
	// AWS account ID.
	AccountId *string `json:"accountId,omitempty"`
	// AWS role name
	RoleName *string `json:"roleName,omitempty"`
	// AWS role external id.
	ExternalId *string `json:"externalId,omitempty"`
	// AWS regions.
	Regions *[]string `json:"regions,omitempty"`
}

// NewSiteAllOfNameResolutionAssumedRoles instantiates a new SiteAllOfNameResolutionAssumedRoles object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSiteAllOfNameResolutionAssumedRoles() *SiteAllOfNameResolutionAssumedRoles {
	this := SiteAllOfNameResolutionAssumedRoles{}
	return &this
}

// NewSiteAllOfNameResolutionAssumedRolesWithDefaults instantiates a new SiteAllOfNameResolutionAssumedRoles object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSiteAllOfNameResolutionAssumedRolesWithDefaults() *SiteAllOfNameResolutionAssumedRoles {
	this := SiteAllOfNameResolutionAssumedRoles{}
	return &this
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *SiteAllOfNameResolutionAssumedRoles) GetAccountId() string {
	if o == nil || o.AccountId == nil {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteAllOfNameResolutionAssumedRoles) GetAccountIdOk() (*string, bool) {
	if o == nil || o.AccountId == nil {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *SiteAllOfNameResolutionAssumedRoles) HasAccountId() bool {
	if o != nil && o.AccountId != nil {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *SiteAllOfNameResolutionAssumedRoles) SetAccountId(v string) {
	o.AccountId = &v
}

// GetRoleName returns the RoleName field value if set, zero value otherwise.
func (o *SiteAllOfNameResolutionAssumedRoles) GetRoleName() string {
	if o == nil || o.RoleName == nil {
		var ret string
		return ret
	}
	return *o.RoleName
}

// GetRoleNameOk returns a tuple with the RoleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteAllOfNameResolutionAssumedRoles) GetRoleNameOk() (*string, bool) {
	if o == nil || o.RoleName == nil {
		return nil, false
	}
	return o.RoleName, true
}

// HasRoleName returns a boolean if a field has been set.
func (o *SiteAllOfNameResolutionAssumedRoles) HasRoleName() bool {
	if o != nil && o.RoleName != nil {
		return true
	}

	return false
}

// SetRoleName gets a reference to the given string and assigns it to the RoleName field.
func (o *SiteAllOfNameResolutionAssumedRoles) SetRoleName(v string) {
	o.RoleName = &v
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise.
func (o *SiteAllOfNameResolutionAssumedRoles) GetExternalId() string {
	if o == nil || o.ExternalId == nil {
		var ret string
		return ret
	}
	return *o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteAllOfNameResolutionAssumedRoles) GetExternalIdOk() (*string, bool) {
	if o == nil || o.ExternalId == nil {
		return nil, false
	}
	return o.ExternalId, true
}

// HasExternalId returns a boolean if a field has been set.
func (o *SiteAllOfNameResolutionAssumedRoles) HasExternalId() bool {
	if o != nil && o.ExternalId != nil {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given string and assigns it to the ExternalId field.
func (o *SiteAllOfNameResolutionAssumedRoles) SetExternalId(v string) {
	o.ExternalId = &v
}

// GetRegions returns the Regions field value if set, zero value otherwise.
func (o *SiteAllOfNameResolutionAssumedRoles) GetRegions() []string {
	if o == nil || o.Regions == nil {
		var ret []string
		return ret
	}
	return *o.Regions
}

// GetRegionsOk returns a tuple with the Regions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteAllOfNameResolutionAssumedRoles) GetRegionsOk() (*[]string, bool) {
	if o == nil || o.Regions == nil {
		return nil, false
	}
	return o.Regions, true
}

// HasRegions returns a boolean if a field has been set.
func (o *SiteAllOfNameResolutionAssumedRoles) HasRegions() bool {
	if o != nil && o.Regions != nil {
		return true
	}

	return false
}

// SetRegions gets a reference to the given []string and assigns it to the Regions field.
func (o *SiteAllOfNameResolutionAssumedRoles) SetRegions(v []string) {
	o.Regions = &v
}

func (o SiteAllOfNameResolutionAssumedRoles) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccountId != nil {
		toSerialize["accountId"] = o.AccountId
	}
	if o.RoleName != nil {
		toSerialize["roleName"] = o.RoleName
	}
	if o.ExternalId != nil {
		toSerialize["externalId"] = o.ExternalId
	}
	if o.Regions != nil {
		toSerialize["regions"] = o.Regions
	}
	return json.Marshal(toSerialize)
}

type NullableSiteAllOfNameResolutionAssumedRoles struct {
	value *SiteAllOfNameResolutionAssumedRoles
	isSet bool
}

func (v NullableSiteAllOfNameResolutionAssumedRoles) Get() *SiteAllOfNameResolutionAssumedRoles {
	return v.value
}

func (v *NullableSiteAllOfNameResolutionAssumedRoles) Set(val *SiteAllOfNameResolutionAssumedRoles) {
	v.value = val
	v.isSet = true
}

func (v NullableSiteAllOfNameResolutionAssumedRoles) IsSet() bool {
	return v.isSet
}

func (v *NullableSiteAllOfNameResolutionAssumedRoles) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSiteAllOfNameResolutionAssumedRoles(val *SiteAllOfNameResolutionAssumedRoles) *NullableSiteAllOfNameResolutionAssumedRoles {
	return &NullableSiteAllOfNameResolutionAssumedRoles{value: val, isSet: true}
}

func (v NullableSiteAllOfNameResolutionAssumedRoles) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSiteAllOfNameResolutionAssumedRoles) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
