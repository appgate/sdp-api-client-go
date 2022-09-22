/*
Appgate SDP Controller REST API

# About   This specification documents the REST API calls for the Appgate SDP Controller.    Please refer to the REST API chapter in the manual or contact Appgate support with any questions about   this functionality. # Getting Started   Requirements for API scripting:   - Access to the Admin/API TLS Connection (default port 8443) of a Controller appliance.     (https://sdphelp.appgate.com/adminguide/appliance-function-configure.html?anchor=admin-api)   - An API user with relevant permissions.     (https://sdphelp.appgate.com/adminguide/administrative-roles-configure.html)   - In order to use the simple login API, Admin MFA must be disabled or the API user must be excluded.     (https://sdphelp.appgate.com/adminguide/mfa-for-admins.html) # Base path   HTTPS requests must be sent to the Admin Interface hostname and port, with **_/admin** path.    For example: **https://appgate.company.com:8443/admin**    All requests must have the **Accept** header as:    **application/vnd.appgate.peer-v17+json** # API Conventions   API conventions are  important to understand and follow strictly.    - While updating objects (via PUT), entire object must be sent with all fields.     - For example, in order to add a remedy method to the condition below:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": []       }       ```     - send the entire object with updated and non-updated fields:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": [{\"type\": \"DisplayMessage\", \"message\": \"test message\"}]       }       ```    - In case Controller returns an error (non-2xx HTTP status code), response body is JSON.     The \"message\" field contains information about the error.     HTTP 422 \"Unprocessable Entity\" has extra `errors` field to list all the issues with specific fields.    - Empty string (\"\") is considered a different value than \"null\" or field being omitted from JSON.     Omitting the field is recommended if no value is intended.     Empty string (\"\") will be almost always rejected as invalid value.    - There are common pattern between many objects:     - **Configuration Objects**: There are many objects with common fields, namely \"id\", \"name\", \"notes\", \"created\"       and \"updated\". These entities are listed, queried, created, updated and deleted in a similar fashion.     - **Distinguished Name**: Users and Devices are identified with what is called Distinguished Names, as used in        LDAP. The distinguished format that identifies a device and a user combination is        \"CN=\\<Device ID\\>,CN=\\<username\\>,OU=\\<Identity Provider Name\\>\". Some objects have the        \"userDistinguishedName\" field, which does not include the CN for Device ID.        This identifies a user on every device.

API version: API version 17.4
Contact: appgatesdp.support@appgate.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// LicenseDetailsUsage License usage information.
type LicenseDetailsUsage struct {
	// The amount of licensed users in the system currently.
	Users *float32 `json:"users,omitempty"`
	// The amount of licensed portal users in the system currently.
	PortalUsers *float32 `json:"portalUsers,omitempty"`
	// The amount of licensed service users in the system currently.
	ServiceUsers *float32 `json:"serviceUsers,omitempty"`
	// The amount of Sites in the system currently.
	Sites *float32 `json:"sites,omitempty"`
	// The amount of grouped Connector clients in the system currently.
	ConnectorGroups *float32 `json:"connectorGroups,omitempty"`
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
