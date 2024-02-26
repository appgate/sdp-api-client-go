/*
Appgate SDP Controller REST API

# About   This specification documents the REST API calls for the Appgate SDP Controller.    Please refer to the REST API chapter in the manual or contact Appgate support with any questions about   this functionality. # Getting Started   Requirements for API scripting:   - Access to the Admin/API TLS Connection (default port 8443) of a Controller appliance.     (https://sdphelp.appgate.com/adminguide/appliance-function-configure.html?anchor=admin-api)   - An API user with relevant permissions.     (https://sdphelp.appgate.com/adminguide/administrative-roles-configure.html)   - In order to use the simple login API, Admin MFA must be disabled or the API user must be excluded.     (https://sdphelp.appgate.com/adminguide/mfa-for-admins.html) # Base path   HTTPS requests must be sent to the Admin Interface hostname and port, with **_/admin** path.    For example: **https://appgate.company.com:8443/admin**    All requests must have the **Accept** header as:    **application/vnd.appgate.peer-v20+json**    An exception is made for the **_/admin/version** endpoint which instead expects an **application/json** Accept header. # API Conventions   API conventions are  important to understand and follow strictly.    - While updating objects (via PUT), entire object must be sent with all fields.     - For example, in order to add a remedy method to the condition below:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": []       }       ```     - send the entire object with updated and non-updated fields:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": [{\"type\": \"DisplayMessage\", \"message\": \"test message\"}]       }       ```    - In case Controller returns an error (non-2xx HTTP status code), response body is JSON.     The \"message\" field contains information about the error.     HTTP 422 \"Unprocessable Entity\" has extra `errors` field to list all the issues with specific fields.    - Empty string (\"\") is considered a different value than \"null\" or field being omitted from JSON.     Omitting the field is recommended if no value is intended.     Empty string (\"\") will be almost always rejected as invalid value.    - There are common pattern between many objects:     - **Configuration Objects**: There are many objects with common fields, namely \"id\", \"name\", \"notes\", \"created\"       and \"updated\". These entities are listed, queried, created, updated and deleted in a similar fashion.     - **Distinguished Name**: Users and Devices are identified with what is called Distinguished Names, as used in        LDAP. The distinguished format that identifies a device and a user combination is        \"CN=\\<Device ID\\>,CN=\\<username\\>,OU=\\<Identity Provider Name\\>\". Some objects have the        \"userDistinguishedName\" field, which does not include the CN for Device ID.        This identifies a user on every device.

API version: API version 20.0
Contact: appgatesdp.support@appgate.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// LicenseEntitlement License entitlement details.
type LicenseEntitlement struct {
	// The maximum users allowed by the license.
	MaxUsers *float32 `json:"maxUsers,omitempty"`
	// The maximum Portal users allowed by the license.
	MaxPortalUsers *float32 `json:"maxPortalUsers,omitempty"`
	// The maximum Service users allowed by the license.
	MaxServiceUsers *float32 `json:"maxServiceUsers,omitempty"`
	// Number of hours User Licenses will be reserved for users before they are reclaimable by others.
	UsersLeaseTimeHours *float32 `json:"usersLeaseTimeHours,omitempty"`
	// Number of hours User Licenses will be reserved for Portal users before they are reclaimable by others.
	PortalUsersLeaseTimeHours *float32 `json:"portalUsersLeaseTimeHours,omitempty"`
	// Number of hours User Licenses will be reserved for Service users before they are reclaimable by others.
	ServiceUsersLeaseTimeHours *float32 `json:"serviceUsersLeaseTimeHours,omitempty"`
	// The maximum sites allowed by the license. If it's the usage details, then it's the amount of sites in the system currently.
	MaxSites *float32 `json:"maxSites,omitempty"`
	// The maximum Connector groups allowed by the license.
	MaxConnectorGroups *float32 `json:"maxConnectorGroups,omitempty"`
}

// NewLicenseEntitlement instantiates a new LicenseEntitlement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLicenseEntitlement() *LicenseEntitlement {
	this := LicenseEntitlement{}
	return &this
}

// NewLicenseEntitlementWithDefaults instantiates a new LicenseEntitlement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLicenseEntitlementWithDefaults() *LicenseEntitlement {
	this := LicenseEntitlement{}
	return &this
}

// GetMaxUsers returns the MaxUsers field value if set, zero value otherwise.
func (o *LicenseEntitlement) GetMaxUsers() float32 {
	if o == nil || o.MaxUsers == nil {
		var ret float32
		return ret
	}
	return *o.MaxUsers
}

// GetMaxUsersOk returns a tuple with the MaxUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseEntitlement) GetMaxUsersOk() (*float32, bool) {
	if o == nil || o.MaxUsers == nil {
		return nil, false
	}
	return o.MaxUsers, true
}

// HasMaxUsers returns a boolean if a field has been set.
func (o *LicenseEntitlement) HasMaxUsers() bool {
	if o != nil && o.MaxUsers != nil {
		return true
	}

	return false
}

// SetMaxUsers gets a reference to the given float32 and assigns it to the MaxUsers field.
func (o *LicenseEntitlement) SetMaxUsers(v float32) {
	o.MaxUsers = &v
}

// GetMaxPortalUsers returns the MaxPortalUsers field value if set, zero value otherwise.
func (o *LicenseEntitlement) GetMaxPortalUsers() float32 {
	if o == nil || o.MaxPortalUsers == nil {
		var ret float32
		return ret
	}
	return *o.MaxPortalUsers
}

// GetMaxPortalUsersOk returns a tuple with the MaxPortalUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseEntitlement) GetMaxPortalUsersOk() (*float32, bool) {
	if o == nil || o.MaxPortalUsers == nil {
		return nil, false
	}
	return o.MaxPortalUsers, true
}

// HasMaxPortalUsers returns a boolean if a field has been set.
func (o *LicenseEntitlement) HasMaxPortalUsers() bool {
	if o != nil && o.MaxPortalUsers != nil {
		return true
	}

	return false
}

// SetMaxPortalUsers gets a reference to the given float32 and assigns it to the MaxPortalUsers field.
func (o *LicenseEntitlement) SetMaxPortalUsers(v float32) {
	o.MaxPortalUsers = &v
}

// GetMaxServiceUsers returns the MaxServiceUsers field value if set, zero value otherwise.
func (o *LicenseEntitlement) GetMaxServiceUsers() float32 {
	if o == nil || o.MaxServiceUsers == nil {
		var ret float32
		return ret
	}
	return *o.MaxServiceUsers
}

// GetMaxServiceUsersOk returns a tuple with the MaxServiceUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseEntitlement) GetMaxServiceUsersOk() (*float32, bool) {
	if o == nil || o.MaxServiceUsers == nil {
		return nil, false
	}
	return o.MaxServiceUsers, true
}

// HasMaxServiceUsers returns a boolean if a field has been set.
func (o *LicenseEntitlement) HasMaxServiceUsers() bool {
	if o != nil && o.MaxServiceUsers != nil {
		return true
	}

	return false
}

// SetMaxServiceUsers gets a reference to the given float32 and assigns it to the MaxServiceUsers field.
func (o *LicenseEntitlement) SetMaxServiceUsers(v float32) {
	o.MaxServiceUsers = &v
}

// GetUsersLeaseTimeHours returns the UsersLeaseTimeHours field value if set, zero value otherwise.
func (o *LicenseEntitlement) GetUsersLeaseTimeHours() float32 {
	if o == nil || o.UsersLeaseTimeHours == nil {
		var ret float32
		return ret
	}
	return *o.UsersLeaseTimeHours
}

// GetUsersLeaseTimeHoursOk returns a tuple with the UsersLeaseTimeHours field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseEntitlement) GetUsersLeaseTimeHoursOk() (*float32, bool) {
	if o == nil || o.UsersLeaseTimeHours == nil {
		return nil, false
	}
	return o.UsersLeaseTimeHours, true
}

// HasUsersLeaseTimeHours returns a boolean if a field has been set.
func (o *LicenseEntitlement) HasUsersLeaseTimeHours() bool {
	if o != nil && o.UsersLeaseTimeHours != nil {
		return true
	}

	return false
}

// SetUsersLeaseTimeHours gets a reference to the given float32 and assigns it to the UsersLeaseTimeHours field.
func (o *LicenseEntitlement) SetUsersLeaseTimeHours(v float32) {
	o.UsersLeaseTimeHours = &v
}

// GetPortalUsersLeaseTimeHours returns the PortalUsersLeaseTimeHours field value if set, zero value otherwise.
func (o *LicenseEntitlement) GetPortalUsersLeaseTimeHours() float32 {
	if o == nil || o.PortalUsersLeaseTimeHours == nil {
		var ret float32
		return ret
	}
	return *o.PortalUsersLeaseTimeHours
}

// GetPortalUsersLeaseTimeHoursOk returns a tuple with the PortalUsersLeaseTimeHours field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseEntitlement) GetPortalUsersLeaseTimeHoursOk() (*float32, bool) {
	if o == nil || o.PortalUsersLeaseTimeHours == nil {
		return nil, false
	}
	return o.PortalUsersLeaseTimeHours, true
}

// HasPortalUsersLeaseTimeHours returns a boolean if a field has been set.
func (o *LicenseEntitlement) HasPortalUsersLeaseTimeHours() bool {
	if o != nil && o.PortalUsersLeaseTimeHours != nil {
		return true
	}

	return false
}

// SetPortalUsersLeaseTimeHours gets a reference to the given float32 and assigns it to the PortalUsersLeaseTimeHours field.
func (o *LicenseEntitlement) SetPortalUsersLeaseTimeHours(v float32) {
	o.PortalUsersLeaseTimeHours = &v
}

// GetServiceUsersLeaseTimeHours returns the ServiceUsersLeaseTimeHours field value if set, zero value otherwise.
func (o *LicenseEntitlement) GetServiceUsersLeaseTimeHours() float32 {
	if o == nil || o.ServiceUsersLeaseTimeHours == nil {
		var ret float32
		return ret
	}
	return *o.ServiceUsersLeaseTimeHours
}

// GetServiceUsersLeaseTimeHoursOk returns a tuple with the ServiceUsersLeaseTimeHours field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseEntitlement) GetServiceUsersLeaseTimeHoursOk() (*float32, bool) {
	if o == nil || o.ServiceUsersLeaseTimeHours == nil {
		return nil, false
	}
	return o.ServiceUsersLeaseTimeHours, true
}

// HasServiceUsersLeaseTimeHours returns a boolean if a field has been set.
func (o *LicenseEntitlement) HasServiceUsersLeaseTimeHours() bool {
	if o != nil && o.ServiceUsersLeaseTimeHours != nil {
		return true
	}

	return false
}

// SetServiceUsersLeaseTimeHours gets a reference to the given float32 and assigns it to the ServiceUsersLeaseTimeHours field.
func (o *LicenseEntitlement) SetServiceUsersLeaseTimeHours(v float32) {
	o.ServiceUsersLeaseTimeHours = &v
}

// GetMaxSites returns the MaxSites field value if set, zero value otherwise.
func (o *LicenseEntitlement) GetMaxSites() float32 {
	if o == nil || o.MaxSites == nil {
		var ret float32
		return ret
	}
	return *o.MaxSites
}

// GetMaxSitesOk returns a tuple with the MaxSites field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseEntitlement) GetMaxSitesOk() (*float32, bool) {
	if o == nil || o.MaxSites == nil {
		return nil, false
	}
	return o.MaxSites, true
}

// HasMaxSites returns a boolean if a field has been set.
func (o *LicenseEntitlement) HasMaxSites() bool {
	if o != nil && o.MaxSites != nil {
		return true
	}

	return false
}

// SetMaxSites gets a reference to the given float32 and assigns it to the MaxSites field.
func (o *LicenseEntitlement) SetMaxSites(v float32) {
	o.MaxSites = &v
}

// GetMaxConnectorGroups returns the MaxConnectorGroups field value if set, zero value otherwise.
func (o *LicenseEntitlement) GetMaxConnectorGroups() float32 {
	if o == nil || o.MaxConnectorGroups == nil {
		var ret float32
		return ret
	}
	return *o.MaxConnectorGroups
}

// GetMaxConnectorGroupsOk returns a tuple with the MaxConnectorGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseEntitlement) GetMaxConnectorGroupsOk() (*float32, bool) {
	if o == nil || o.MaxConnectorGroups == nil {
		return nil, false
	}
	return o.MaxConnectorGroups, true
}

// HasMaxConnectorGroups returns a boolean if a field has been set.
func (o *LicenseEntitlement) HasMaxConnectorGroups() bool {
	if o != nil && o.MaxConnectorGroups != nil {
		return true
	}

	return false
}

// SetMaxConnectorGroups gets a reference to the given float32 and assigns it to the MaxConnectorGroups field.
func (o *LicenseEntitlement) SetMaxConnectorGroups(v float32) {
	o.MaxConnectorGroups = &v
}

func (o LicenseEntitlement) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.MaxUsers != nil {
		toSerialize["maxUsers"] = o.MaxUsers
	}
	if o.MaxPortalUsers != nil {
		toSerialize["maxPortalUsers"] = o.MaxPortalUsers
	}
	if o.MaxServiceUsers != nil {
		toSerialize["maxServiceUsers"] = o.MaxServiceUsers
	}
	if o.UsersLeaseTimeHours != nil {
		toSerialize["usersLeaseTimeHours"] = o.UsersLeaseTimeHours
	}
	if o.PortalUsersLeaseTimeHours != nil {
		toSerialize["portalUsersLeaseTimeHours"] = o.PortalUsersLeaseTimeHours
	}
	if o.ServiceUsersLeaseTimeHours != nil {
		toSerialize["serviceUsersLeaseTimeHours"] = o.ServiceUsersLeaseTimeHours
	}
	if o.MaxSites != nil {
		toSerialize["maxSites"] = o.MaxSites
	}
	if o.MaxConnectorGroups != nil {
		toSerialize["maxConnectorGroups"] = o.MaxConnectorGroups
	}
	return json.Marshal(toSerialize)
}

type NullableLicenseEntitlement struct {
	value *LicenseEntitlement
	isSet bool
}

func (v NullableLicenseEntitlement) Get() *LicenseEntitlement {
	return v.value
}

func (v *NullableLicenseEntitlement) Set(val *LicenseEntitlement) {
	v.value = val
	v.isSet = true
}

func (v NullableLicenseEntitlement) IsSet() bool {
	return v.isSet
}

func (v *NullableLicenseEntitlement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLicenseEntitlement(val *LicenseEntitlement) *NullableLicenseEntitlement {
	return &NullableLicenseEntitlement{value: val, isSet: true}
}

func (v NullableLicenseEntitlement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLicenseEntitlement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
