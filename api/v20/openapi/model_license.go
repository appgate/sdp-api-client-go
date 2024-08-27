/*
Appgate SDP Controller REST API

# About   This specification documents the REST API calls for the Appgate SDP Controller.    Please refer to the REST API chapter in the manual or contact Appgate support with any questions about   this functionality. # Getting Started   Requirements for API scripting:   - Access to the Admin/API TLS Connection (default port 8443) of a Controller appliance.     (https://sdphelp.appgate.com/adminguide/appliance-function-configure.html?anchor=admin-api)   - An API user with relevant permissions.     (https://sdphelp.appgate.com/adminguide/administrative-roles-configure.html)   - In order to use the simple login API, Admin MFA must be disabled or the API user must be excluded.     (https://sdphelp.appgate.com/adminguide/mfa-for-admins.html) # Base path   HTTPS requests must be sent to the Admin Interface hostname and port, with **_/admin** path.    For example: **https://appgate.company.com:8443/admin**    All requests must have the **Accept** header as:    **application/vnd.appgate.peer-v20+json**    An exception is made for the **_/admin/version** endpoint which instead expects an **application/json** Accept header. # API Conventions   API conventions are  important to understand and follow strictly.    - While updating objects (via PUT), entire object must be sent with all fields.     - For example, in order to add a remedy method to the condition below:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": []       }       ```     - send the entire object with updated and non-updated fields:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": [{\"type\": \"DisplayMessage\", \"message\": \"test message\"}]       }       ```    - In case Controller returns an error (non-2xx HTTP status code), response body is JSON.     The \"message\" field contains information about the error.     HTTP 422 \"Unprocessable Entity\" has extra `errors` field to list all the issues with specific fields.    - Empty string (\"\") is considered a different value than \"null\" or field being omitted from JSON.     Omitting the field is recommended if no value is intended.     Empty string (\"\") will be almost always rejected as invalid value.    - There are common pattern between many objects:     - **Configuration Objects**: There are many objects with common fields, namely \"id\", \"name\", \"notes\", \"created\"       and \"updated\". These entities are listed, queried, created, updated and deleted in a similar fashion.     - **Distinguished Name**: Users and Devices are identified with what is called Distinguished Names, as used in        LDAP. The distinguished format that identifies a device and a user combination is        \"CN=\\<Device ID\\>,CN=\\<username\\>,OU=\\<Identity Provider Name\\>\". Some objects have the        \"userDistinguishedName\" field, which does not include the CN for Device ID.        This identifies a user on every device.

API version: API version 20.3
Contact: appgatesdp.support@appgate.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// License struct for License
type License struct {
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
	// Unique ID for the license.
	Id *string `json:"id,omitempty"`
	// License version. There can be only one v1 License in the system, while v2 Licenses can be added as extra.
	Version *float32 `json:"version,omitempty"`
	// Type of the license. 1: production, 2: installation, 3: test, 4: built-in, 5: aws built-in, 6: metered
	Type *float32 `json:"type,omitempty"`
	// Request code for the license. If built-in license is in place, use this code to get a license. It's based on the CA certificate.
	Request *string `json:"request,omitempty"`
	// The expiration date of the license.
	Expiration *time.Time `json:"expiration,omitempty"`
	// Error message if the License is not invalid.
	Error *string `json:"error,omitempty"`
}

// NewLicense instantiates a new License object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLicense() *License {
	this := License{}
	return &this
}

// NewLicenseWithDefaults instantiates a new License object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLicenseWithDefaults() *License {
	this := License{}
	return &this
}

// GetMaxUsers returns the MaxUsers field value if set, zero value otherwise.
func (o *License) GetMaxUsers() float32 {
	if o == nil || o.MaxUsers == nil {
		var ret float32
		return ret
	}
	return *o.MaxUsers
}

// GetMaxUsersOk returns a tuple with the MaxUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *License) GetMaxUsersOk() (*float32, bool) {
	if o == nil || o.MaxUsers == nil {
		return nil, false
	}
	return o.MaxUsers, true
}

// HasMaxUsers returns a boolean if a field has been set.
func (o *License) HasMaxUsers() bool {
	if o != nil && o.MaxUsers != nil {
		return true
	}

	return false
}

// SetMaxUsers gets a reference to the given float32 and assigns it to the MaxUsers field.
func (o *License) SetMaxUsers(v float32) {
	o.MaxUsers = &v
}

// GetMaxPortalUsers returns the MaxPortalUsers field value if set, zero value otherwise.
func (o *License) GetMaxPortalUsers() float32 {
	if o == nil || o.MaxPortalUsers == nil {
		var ret float32
		return ret
	}
	return *o.MaxPortalUsers
}

// GetMaxPortalUsersOk returns a tuple with the MaxPortalUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *License) GetMaxPortalUsersOk() (*float32, bool) {
	if o == nil || o.MaxPortalUsers == nil {
		return nil, false
	}
	return o.MaxPortalUsers, true
}

// HasMaxPortalUsers returns a boolean if a field has been set.
func (o *License) HasMaxPortalUsers() bool {
	if o != nil && o.MaxPortalUsers != nil {
		return true
	}

	return false
}

// SetMaxPortalUsers gets a reference to the given float32 and assigns it to the MaxPortalUsers field.
func (o *License) SetMaxPortalUsers(v float32) {
	o.MaxPortalUsers = &v
}

// GetMaxServiceUsers returns the MaxServiceUsers field value if set, zero value otherwise.
func (o *License) GetMaxServiceUsers() float32 {
	if o == nil || o.MaxServiceUsers == nil {
		var ret float32
		return ret
	}
	return *o.MaxServiceUsers
}

// GetMaxServiceUsersOk returns a tuple with the MaxServiceUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *License) GetMaxServiceUsersOk() (*float32, bool) {
	if o == nil || o.MaxServiceUsers == nil {
		return nil, false
	}
	return o.MaxServiceUsers, true
}

// HasMaxServiceUsers returns a boolean if a field has been set.
func (o *License) HasMaxServiceUsers() bool {
	if o != nil && o.MaxServiceUsers != nil {
		return true
	}

	return false
}

// SetMaxServiceUsers gets a reference to the given float32 and assigns it to the MaxServiceUsers field.
func (o *License) SetMaxServiceUsers(v float32) {
	o.MaxServiceUsers = &v
}

// GetUsersLeaseTimeHours returns the UsersLeaseTimeHours field value if set, zero value otherwise.
func (o *License) GetUsersLeaseTimeHours() float32 {
	if o == nil || o.UsersLeaseTimeHours == nil {
		var ret float32
		return ret
	}
	return *o.UsersLeaseTimeHours
}

// GetUsersLeaseTimeHoursOk returns a tuple with the UsersLeaseTimeHours field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *License) GetUsersLeaseTimeHoursOk() (*float32, bool) {
	if o == nil || o.UsersLeaseTimeHours == nil {
		return nil, false
	}
	return o.UsersLeaseTimeHours, true
}

// HasUsersLeaseTimeHours returns a boolean if a field has been set.
func (o *License) HasUsersLeaseTimeHours() bool {
	if o != nil && o.UsersLeaseTimeHours != nil {
		return true
	}

	return false
}

// SetUsersLeaseTimeHours gets a reference to the given float32 and assigns it to the UsersLeaseTimeHours field.
func (o *License) SetUsersLeaseTimeHours(v float32) {
	o.UsersLeaseTimeHours = &v
}

// GetPortalUsersLeaseTimeHours returns the PortalUsersLeaseTimeHours field value if set, zero value otherwise.
func (o *License) GetPortalUsersLeaseTimeHours() float32 {
	if o == nil || o.PortalUsersLeaseTimeHours == nil {
		var ret float32
		return ret
	}
	return *o.PortalUsersLeaseTimeHours
}

// GetPortalUsersLeaseTimeHoursOk returns a tuple with the PortalUsersLeaseTimeHours field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *License) GetPortalUsersLeaseTimeHoursOk() (*float32, bool) {
	if o == nil || o.PortalUsersLeaseTimeHours == nil {
		return nil, false
	}
	return o.PortalUsersLeaseTimeHours, true
}

// HasPortalUsersLeaseTimeHours returns a boolean if a field has been set.
func (o *License) HasPortalUsersLeaseTimeHours() bool {
	if o != nil && o.PortalUsersLeaseTimeHours != nil {
		return true
	}

	return false
}

// SetPortalUsersLeaseTimeHours gets a reference to the given float32 and assigns it to the PortalUsersLeaseTimeHours field.
func (o *License) SetPortalUsersLeaseTimeHours(v float32) {
	o.PortalUsersLeaseTimeHours = &v
}

// GetServiceUsersLeaseTimeHours returns the ServiceUsersLeaseTimeHours field value if set, zero value otherwise.
func (o *License) GetServiceUsersLeaseTimeHours() float32 {
	if o == nil || o.ServiceUsersLeaseTimeHours == nil {
		var ret float32
		return ret
	}
	return *o.ServiceUsersLeaseTimeHours
}

// GetServiceUsersLeaseTimeHoursOk returns a tuple with the ServiceUsersLeaseTimeHours field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *License) GetServiceUsersLeaseTimeHoursOk() (*float32, bool) {
	if o == nil || o.ServiceUsersLeaseTimeHours == nil {
		return nil, false
	}
	return o.ServiceUsersLeaseTimeHours, true
}

// HasServiceUsersLeaseTimeHours returns a boolean if a field has been set.
func (o *License) HasServiceUsersLeaseTimeHours() bool {
	if o != nil && o.ServiceUsersLeaseTimeHours != nil {
		return true
	}

	return false
}

// SetServiceUsersLeaseTimeHours gets a reference to the given float32 and assigns it to the ServiceUsersLeaseTimeHours field.
func (o *License) SetServiceUsersLeaseTimeHours(v float32) {
	o.ServiceUsersLeaseTimeHours = &v
}

// GetMaxSites returns the MaxSites field value if set, zero value otherwise.
func (o *License) GetMaxSites() float32 {
	if o == nil || o.MaxSites == nil {
		var ret float32
		return ret
	}
	return *o.MaxSites
}

// GetMaxSitesOk returns a tuple with the MaxSites field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *License) GetMaxSitesOk() (*float32, bool) {
	if o == nil || o.MaxSites == nil {
		return nil, false
	}
	return o.MaxSites, true
}

// HasMaxSites returns a boolean if a field has been set.
func (o *License) HasMaxSites() bool {
	if o != nil && o.MaxSites != nil {
		return true
	}

	return false
}

// SetMaxSites gets a reference to the given float32 and assigns it to the MaxSites field.
func (o *License) SetMaxSites(v float32) {
	o.MaxSites = &v
}

// GetMaxConnectorGroups returns the MaxConnectorGroups field value if set, zero value otherwise.
func (o *License) GetMaxConnectorGroups() float32 {
	if o == nil || o.MaxConnectorGroups == nil {
		var ret float32
		return ret
	}
	return *o.MaxConnectorGroups
}

// GetMaxConnectorGroupsOk returns a tuple with the MaxConnectorGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *License) GetMaxConnectorGroupsOk() (*float32, bool) {
	if o == nil || o.MaxConnectorGroups == nil {
		return nil, false
	}
	return o.MaxConnectorGroups, true
}

// HasMaxConnectorGroups returns a boolean if a field has been set.
func (o *License) HasMaxConnectorGroups() bool {
	if o != nil && o.MaxConnectorGroups != nil {
		return true
	}

	return false
}

// SetMaxConnectorGroups gets a reference to the given float32 and assigns it to the MaxConnectorGroups field.
func (o *License) SetMaxConnectorGroups(v float32) {
	o.MaxConnectorGroups = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *License) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *License) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *License) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *License) SetId(v string) {
	o.Id = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *License) GetVersion() float32 {
	if o == nil || o.Version == nil {
		var ret float32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *License) GetVersionOk() (*float32, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *License) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given float32 and assigns it to the Version field.
func (o *License) SetVersion(v float32) {
	o.Version = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *License) GetType() float32 {
	if o == nil || o.Type == nil {
		var ret float32
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *License) GetTypeOk() (*float32, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *License) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given float32 and assigns it to the Type field.
func (o *License) SetType(v float32) {
	o.Type = &v
}

// GetRequest returns the Request field value if set, zero value otherwise.
func (o *License) GetRequest() string {
	if o == nil || o.Request == nil {
		var ret string
		return ret
	}
	return *o.Request
}

// GetRequestOk returns a tuple with the Request field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *License) GetRequestOk() (*string, bool) {
	if o == nil || o.Request == nil {
		return nil, false
	}
	return o.Request, true
}

// HasRequest returns a boolean if a field has been set.
func (o *License) HasRequest() bool {
	if o != nil && o.Request != nil {
		return true
	}

	return false
}

// SetRequest gets a reference to the given string and assigns it to the Request field.
func (o *License) SetRequest(v string) {
	o.Request = &v
}

// GetExpiration returns the Expiration field value if set, zero value otherwise.
func (o *License) GetExpiration() time.Time {
	if o == nil || o.Expiration == nil {
		var ret time.Time
		return ret
	}
	return *o.Expiration
}

// GetExpirationOk returns a tuple with the Expiration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *License) GetExpirationOk() (*time.Time, bool) {
	if o == nil || o.Expiration == nil {
		return nil, false
	}
	return o.Expiration, true
}

// HasExpiration returns a boolean if a field has been set.
func (o *License) HasExpiration() bool {
	if o != nil && o.Expiration != nil {
		return true
	}

	return false
}

// SetExpiration gets a reference to the given time.Time and assigns it to the Expiration field.
func (o *License) SetExpiration(v time.Time) {
	o.Expiration = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *License) GetError() string {
	if o == nil || o.Error == nil {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *License) GetErrorOk() (*string, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *License) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *License) SetError(v string) {
	o.Error = &v
}

func (o License) MarshalJSON() ([]byte, error) {
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
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Request != nil {
		toSerialize["request"] = o.Request
	}
	if o.Expiration != nil {
		toSerialize["expiration"] = o.Expiration
	}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	return json.Marshal(toSerialize)
}

type NullableLicense struct {
	value *License
	isSet bool
}

func (v NullableLicense) Get() *License {
	return v.value
}

func (v *NullableLicense) Set(val *License) {
	v.value = val
	v.isSet = true
}

func (v NullableLicense) IsSet() bool {
	return v.isSet
}

func (v *NullableLicense) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLicense(val *License) *NullableLicense {
	return &NullableLicense{value: val, isSet: true}
}

func (v NullableLicense) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLicense) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
