/*
Appgate SDP Controller REST API

# About   This specification documents the REST API calls for the Appgate SDP Controller.    Please refer to the REST API chapter in the manual or contact Appgate support with any questions about   this functionality. # Getting Started   Requirements for API scripting:   - Access to the Admin/API TLS Connection (default port 8443) of a Controller appliance.     (https://sdphelp.appgate.com/adminguide/appliance-function-configure.html?anchor=admin-api)   - An API user with relevant permissions.     (https://sdphelp.appgate.com/adminguide/administrative-roles-configure.html)   - In order to use the simple login API, Admin MFA must be disabled or the API user must be excluded.     (https://sdphelp.appgate.com/adminguide/mfa-for-admins.html) # Base path   HTTPS requests must be sent to the Admin Interface hostname and port, with **_/admin** path.    For example: **https://appgate.company.com:8443/admin**    All requests must have the **Accept** header as:    **application/vnd.appgate.peer-v17+json** # API Conventions   API conventions are  important to understand and follow strictly.    - While updating objects (via PUT), entire object must be sent with all fields.     - For example, in order to add a remedy method to the condition below:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": []       }       ```     - send the entire object with updated and non-updated fields:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": [{\"type\": \"DisplayMessage\", \"message\": \"test message\"}]       }       ```    - In case Controller returns an error (non-2xx HTTP status code), response body is JSON.     The \"message\" field contains information about the error.     HTTP 422 \"Unprocessable Entity\" has extra `errors` field to list all the issues with specific fields.    - Empty string (\"\") is considered a different value than \"null\" or field being omitted from JSON.     Omitting the field is recommended if no value is intended.     Empty string (\"\") will be almost always rejected as invalid value.    - There are common pattern between many objects:     - **Configuration Objects**: There are many objects with common fields, namely \"id\", \"name\", \"notes\", \"created\"       and \"updated\". These entities are listed, queried, created, updated and deleted in a similar fashion.     - **Distinguished Name**: Users and Devices are identified with what is called Distinguished Names, as used in        LDAP. The distinguished format that identifies a device and a user combination is        \"CN=\\<Device ID\\>,CN=\\<username\\>,OU=\\<Identity Provider Name\\>\". Some objects have the        \"userDistinguishedName\" field, which does not include the CN for Device ID.        This identifies a user on every device.

API version: API version 17.1
Contact: appgatesdp.support@appgate.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// ServiceUserAllOf Represents a Service User.
type ServiceUserAllOf struct {
	// Labels of the service. Will be used as claim.
	Labels *map[string]string `json:"labels,omitempty"`
	// Password for the user. Omit the field to keep the old password when updating a user.
	Password string `json:"password"`
	// If true, the user will not be able to login.
	Disabled *bool `json:"disabled,omitempty"`
	// Number of wrong password login attempts since last successiful login.
	FailedLoginAttempts *float32 `json:"failedLoginAttempts,omitempty"`
	// The date time when the user got locked out. A service user is locked out of the system after X amount of consecutive failed login attempts, where X is configurable in the service Identity Provider. The lock is in effect for Y minute(s), where Y is configurable in the service Identity Provider. When the user logs in successfully, this field becomes null.
	LockStart *time.Time `json:"lockStart,omitempty"`
}

// NewServiceUserAllOf instantiates a new ServiceUserAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceUserAllOf(password string) *ServiceUserAllOf {
	this := ServiceUserAllOf{}
	this.Password = password
	return &this
}

// NewServiceUserAllOfWithDefaults instantiates a new ServiceUserAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceUserAllOfWithDefaults() *ServiceUserAllOf {
	this := ServiceUserAllOf{}
	return &this
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *ServiceUserAllOf) GetLabels() map[string]string {
	if o == nil || o.Labels == nil {
		var ret map[string]string
		return ret
	}
	return *o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceUserAllOf) GetLabelsOk() (*map[string]string, bool) {
	if o == nil || o.Labels == nil {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *ServiceUserAllOf) HasLabels() bool {
	if o != nil && o.Labels != nil {
		return true
	}

	return false
}

// SetLabels gets a reference to the given map[string]string and assigns it to the Labels field.
func (o *ServiceUserAllOf) SetLabels(v map[string]string) {
	o.Labels = &v
}

// GetPassword returns the Password field value
func (o *ServiceUserAllOf) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *ServiceUserAllOf) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *ServiceUserAllOf) SetPassword(v string) {
	o.Password = v
}

// GetDisabled returns the Disabled field value if set, zero value otherwise.
func (o *ServiceUserAllOf) GetDisabled() bool {
	if o == nil || o.Disabled == nil {
		var ret bool
		return ret
	}
	return *o.Disabled
}

// GetDisabledOk returns a tuple with the Disabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceUserAllOf) GetDisabledOk() (*bool, bool) {
	if o == nil || o.Disabled == nil {
		return nil, false
	}
	return o.Disabled, true
}

// HasDisabled returns a boolean if a field has been set.
func (o *ServiceUserAllOf) HasDisabled() bool {
	if o != nil && o.Disabled != nil {
		return true
	}

	return false
}

// SetDisabled gets a reference to the given bool and assigns it to the Disabled field.
func (o *ServiceUserAllOf) SetDisabled(v bool) {
	o.Disabled = &v
}

// GetFailedLoginAttempts returns the FailedLoginAttempts field value if set, zero value otherwise.
func (o *ServiceUserAllOf) GetFailedLoginAttempts() float32 {
	if o == nil || o.FailedLoginAttempts == nil {
		var ret float32
		return ret
	}
	return *o.FailedLoginAttempts
}

// GetFailedLoginAttemptsOk returns a tuple with the FailedLoginAttempts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceUserAllOf) GetFailedLoginAttemptsOk() (*float32, bool) {
	if o == nil || o.FailedLoginAttempts == nil {
		return nil, false
	}
	return o.FailedLoginAttempts, true
}

// HasFailedLoginAttempts returns a boolean if a field has been set.
func (o *ServiceUserAllOf) HasFailedLoginAttempts() bool {
	if o != nil && o.FailedLoginAttempts != nil {
		return true
	}

	return false
}

// SetFailedLoginAttempts gets a reference to the given float32 and assigns it to the FailedLoginAttempts field.
func (o *ServiceUserAllOf) SetFailedLoginAttempts(v float32) {
	o.FailedLoginAttempts = &v
}

// GetLockStart returns the LockStart field value if set, zero value otherwise.
func (o *ServiceUserAllOf) GetLockStart() time.Time {
	if o == nil || o.LockStart == nil {
		var ret time.Time
		return ret
	}
	return *o.LockStart
}

// GetLockStartOk returns a tuple with the LockStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceUserAllOf) GetLockStartOk() (*time.Time, bool) {
	if o == nil || o.LockStart == nil {
		return nil, false
	}
	return o.LockStart, true
}

// HasLockStart returns a boolean if a field has been set.
func (o *ServiceUserAllOf) HasLockStart() bool {
	if o != nil && o.LockStart != nil {
		return true
	}

	return false
}

// SetLockStart gets a reference to the given time.Time and assigns it to the LockStart field.
func (o *ServiceUserAllOf) SetLockStart(v time.Time) {
	o.LockStart = &v
}

func (o ServiceUserAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Labels != nil {
		toSerialize["labels"] = o.Labels
	}
	if true {
		toSerialize["password"] = o.Password
	}
	if o.Disabled != nil {
		toSerialize["disabled"] = o.Disabled
	}
	if o.FailedLoginAttempts != nil {
		toSerialize["failedLoginAttempts"] = o.FailedLoginAttempts
	}
	if o.LockStart != nil {
		toSerialize["lockStart"] = o.LockStart
	}
	return json.Marshal(toSerialize)
}

type NullableServiceUserAllOf struct {
	value *ServiceUserAllOf
	isSet bool
}

func (v NullableServiceUserAllOf) Get() *ServiceUserAllOf {
	return v.value
}

func (v *NullableServiceUserAllOf) Set(val *ServiceUserAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceUserAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceUserAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceUserAllOf(val *ServiceUserAllOf) *NullableServiceUserAllOf {
	return &NullableServiceUserAllOf{value: val, isSet: true}
}

func (v NullableServiceUserAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceUserAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
