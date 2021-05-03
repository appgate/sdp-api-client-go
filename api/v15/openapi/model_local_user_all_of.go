/*
 * Appgate SDP Controller REST API
 *
 * # About   This specification documents the REST API calls for the Appgate SDP Controller.    Please refer to the Integration chapter in the manual or contact Appgate support with any questions about   this functionality. # Getting Started   Requirements for API scripting:   - Access to the Admin/API TLS Connection (default port 8443) of a Controller appliance.     (https://sdphelp.appgate.com/adminguide/appliance-functions-configure.html?anchor=admin-api)   - An API user with relevant permissions.     (https://sdphelp.appgate.com/adminguide/administrative-roles-configure.html)   - In order to use the simple login API, Admin MFA must be disabled or the API user must be excluded.     (https://sdphelp.appgate.com/adminguide/mfa-for-admins.html) # Base path   HTTPS requests must be sent to the Admin Interface hostname and port, with **_/admin** path.    For example: **https://appgate.company.com:8443/admin**    All requests must have the **Accept** header as:    **application/vnd.appgate.peer-v15+json** # API Conventions   API conventions are  important to understand and follow strictly.    - While updating objects (via PUT), entire object must be sent with all fields.     - For example, in order to add a remedy method to the condition below:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": []       }       ```     - send the entire object with updated and non-updated fields:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": [{\"type\": \"DisplayMessage\", \"message\": \"test message\"}]       }       ```    - In case Controller returns an error (non-2xx HTTP status code), response body is JSON.     The \"message\" field contains information about the error.     HTTP 422 \"Unprocessable Entity\" has extra `errors` field to list all the issues with specific fields.    - Empty string (\"\") is considered a different value than \"null\" or field being omitted from JSON.     Omitting the field is recommend if no value is intended.     Empty string (\"\") will be almost always rejected as invalid value.    - There are common pattern between many objects:     - **Configuration Objects**: There are many objects with common fields, namely \"id\", \"name\", \"notes\", \"created\"       and \"updated\". These entities are listed, queried, created, updated and deleted in a similar fashion.     - **Distinguished Name**: Users and Devices are identified with what is called Distinguished Names, as used in        LDAP. The distinguished format that identifies a device and a user combination is        \"CN=\\<Device ID\\>,CN=\\<username\\>,OU=\\<Identity Provider Name\\>\". Some objects have the        \"userDistinguishedName\" field, which does not include the CN for Device ID.        This identifies a user on every device.
 *
 * API version: API version 15
 * Contact: appgatesdp.support@appgate.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// LocalUserAllOf Represents a Local User.
type LocalUserAllOf struct {
	// First name of the user. May be used as claim.
	FirstName string `json:"firstName"`
	// Last name of the user. May be used as claim.
	LastName string `json:"lastName"`
	// Password for the user. Omit the field to keep the old password when updating a user.
	Password string `json:"password"`
	// E-mail address for the user. May be used as claim.
	Email *string `json:"email,omitempty"`
	// Phone number for the user. May be used as claim.
	Phone *string `json:"phone,omitempty"`
	// Number of wrong password login attempts since last successiful login.
	FailedLoginAttempts *float32 `json:"failedLoginAttempts,omitempty"`
	// The date time when the user got locked out. A local user is locked out of the system after 5 consecutive failed login attempts. The lock is in effect for 1 minute. When the user logs in successfully, this field becomes null.
	LockStart *time.Time `json:"lockStart,omitempty"`
}

// NewLocalUserAllOf instantiates a new LocalUserAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLocalUserAllOf(firstName string, lastName string, password string) *LocalUserAllOf {
	this := LocalUserAllOf{}
	this.FirstName = firstName
	this.LastName = lastName
	this.Password = password
	return &this
}

// NewLocalUserAllOfWithDefaults instantiates a new LocalUserAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLocalUserAllOfWithDefaults() *LocalUserAllOf {
	this := LocalUserAllOf{}
	return &this
}

// GetFirstName returns the FirstName field value
func (o *LocalUserAllOf) GetFirstName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value
// and a boolean to check if the value has been set.
func (o *LocalUserAllOf) GetFirstNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FirstName, true
}

// SetFirstName sets field value
func (o *LocalUserAllOf) SetFirstName(v string) {
	o.FirstName = v
}

// GetLastName returns the LastName field value
func (o *LocalUserAllOf) GetLastName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value
// and a boolean to check if the value has been set.
func (o *LocalUserAllOf) GetLastNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastName, true
}

// SetLastName sets field value
func (o *LocalUserAllOf) SetLastName(v string) {
	o.LastName = v
}

// GetPassword returns the Password field value
func (o *LocalUserAllOf) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *LocalUserAllOf) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *LocalUserAllOf) SetPassword(v string) {
	o.Password = v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *LocalUserAllOf) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocalUserAllOf) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *LocalUserAllOf) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *LocalUserAllOf) SetEmail(v string) {
	o.Email = &v
}

// GetPhone returns the Phone field value if set, zero value otherwise.
func (o *LocalUserAllOf) GetPhone() string {
	if o == nil || o.Phone == nil {
		var ret string
		return ret
	}
	return *o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocalUserAllOf) GetPhoneOk() (*string, bool) {
	if o == nil || o.Phone == nil {
		return nil, false
	}
	return o.Phone, true
}

// HasPhone returns a boolean if a field has been set.
func (o *LocalUserAllOf) HasPhone() bool {
	if o != nil && o.Phone != nil {
		return true
	}

	return false
}

// SetPhone gets a reference to the given string and assigns it to the Phone field.
func (o *LocalUserAllOf) SetPhone(v string) {
	o.Phone = &v
}

// GetFailedLoginAttempts returns the FailedLoginAttempts field value if set, zero value otherwise.
func (o *LocalUserAllOf) GetFailedLoginAttempts() float32 {
	if o == nil || o.FailedLoginAttempts == nil {
		var ret float32
		return ret
	}
	return *o.FailedLoginAttempts
}

// GetFailedLoginAttemptsOk returns a tuple with the FailedLoginAttempts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocalUserAllOf) GetFailedLoginAttemptsOk() (*float32, bool) {
	if o == nil || o.FailedLoginAttempts == nil {
		return nil, false
	}
	return o.FailedLoginAttempts, true
}

// HasFailedLoginAttempts returns a boolean if a field has been set.
func (o *LocalUserAllOf) HasFailedLoginAttempts() bool {
	if o != nil && o.FailedLoginAttempts != nil {
		return true
	}

	return false
}

// SetFailedLoginAttempts gets a reference to the given float32 and assigns it to the FailedLoginAttempts field.
func (o *LocalUserAllOf) SetFailedLoginAttempts(v float32) {
	o.FailedLoginAttempts = &v
}

// GetLockStart returns the LockStart field value if set, zero value otherwise.
func (o *LocalUserAllOf) GetLockStart() time.Time {
	if o == nil || o.LockStart == nil {
		var ret time.Time
		return ret
	}
	return *o.LockStart
}

// GetLockStartOk returns a tuple with the LockStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocalUserAllOf) GetLockStartOk() (*time.Time, bool) {
	if o == nil || o.LockStart == nil {
		return nil, false
	}
	return o.LockStart, true
}

// HasLockStart returns a boolean if a field has been set.
func (o *LocalUserAllOf) HasLockStart() bool {
	if o != nil && o.LockStart != nil {
		return true
	}

	return false
}

// SetLockStart gets a reference to the given time.Time and assigns it to the LockStart field.
func (o *LocalUserAllOf) SetLockStart(v time.Time) {
	o.LockStart = &v
}

func (o LocalUserAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["firstName"] = o.FirstName
	}
	if true {
		toSerialize["lastName"] = o.LastName
	}
	if true {
		toSerialize["password"] = o.Password
	}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.Phone != nil {
		toSerialize["phone"] = o.Phone
	}
	if o.FailedLoginAttempts != nil {
		toSerialize["failedLoginAttempts"] = o.FailedLoginAttempts
	}
	if o.LockStart != nil {
		toSerialize["lockStart"] = o.LockStart
	}
	return json.Marshal(toSerialize)
}

type NullableLocalUserAllOf struct {
	value *LocalUserAllOf
	isSet bool
}

func (v NullableLocalUserAllOf) Get() *LocalUserAllOf {
	return v.value
}

func (v *NullableLocalUserAllOf) Set(val *LocalUserAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableLocalUserAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableLocalUserAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocalUserAllOf(val *LocalUserAllOf) *NullableLocalUserAllOf {
	return &NullableLocalUserAllOf{value: val, isSet: true}
}

func (v NullableLocalUserAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocalUserAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}