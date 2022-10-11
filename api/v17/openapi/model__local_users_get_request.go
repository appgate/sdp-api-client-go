/*
Appgate SDP Controller REST API

# About   This specification documents the REST API calls for the Appgate SDP Controller.    Please refer to the REST API chapter in the manual or contact Appgate support with any questions about   this functionality. # Getting Started   Requirements for API scripting:   - Access to the Admin/API TLS Connection (default port 8443) of a Controller appliance.     (https://sdphelp.appgate.com/adminguide/appliance-function-configure.html?anchor=admin-api)   - An API user with relevant permissions.     (https://sdphelp.appgate.com/adminguide/administrative-roles-configure.html)   - In order to use the simple login API, Admin MFA must be disabled or the API user must be excluded.     (https://sdphelp.appgate.com/adminguide/mfa-for-admins.html) # Base path   HTTPS requests must be sent to the Admin Interface hostname and port, with **_/admin** path.    For example: **https://appgate.company.com:8443/admin**    All requests must have the **Accept** header as:    **application/vnd.appgate.peer-v17+json** # API Conventions   API conventions are  important to understand and follow strictly.    - While updating objects (via PUT), entire object must be sent with all fields.     - For example, in order to add a remedy method to the condition below:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": []       }       ```     - send the entire object with updated and non-updated fields:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": [{\"type\": \"DisplayMessage\", \"message\": \"test message\"}]       }       ```    - In case Controller returns an error (non-2xx HTTP status code), response body is JSON.     The \"message\" field contains information about the error.     HTTP 422 \"Unprocessable Entity\" has extra `errors` field to list all the issues with specific fields.    - Empty string (\"\") is considered a different value than \"null\" or field being omitted from JSON.     Omitting the field is recommended if no value is intended.     Empty string (\"\") will be almost always rejected as invalid value.    - There are common pattern between many objects:     - **Configuration Objects**: There are many objects with common fields, namely \"id\", \"name\", \"notes\", \"created\"       and \"updated\". These entities are listed, queried, created, updated and deleted in a similar fashion.     - **Distinguished Name**: Users and Devices are identified with what is called Distinguished Names, as used in        LDAP. The distinguished format that identifies a device and a user combination is        \"CN=\\<Device ID\\>,CN=\\<username\\>,OU=\\<Identity Provider Name\\>\". Some objects have the        \"userDistinguishedName\" field, which does not include the CN for Device ID.        This identifies a user on every device.

API version: API version 17.5
Contact: appgatesdp.support@appgate.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// LocalUsersGetRequest struct for LocalUsersGetRequest
type LocalUsersGetRequest struct {
	// ID of the object.
	Id *string `json:"id,omitempty"`
	// Name of the object.
	Name string `json:"name"`
	// Notes for the object. Used for documentation purposes.
	Notes *string `json:"notes,omitempty"`
	// Create date.
	Created *time.Time `json:"created,omitempty"`
	// Last update date.
	Updated *time.Time `json:"updated,omitempty"`
	// Array of tags.
	Tags []string `json:"tags,omitempty"`
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
	// If true, the user will not be able to login.
	Disabled *bool `json:"disabled,omitempty"`
	// Number of wrong password login attempts since last successiful login.
	FailedLoginAttempts *float32 `json:"failedLoginAttempts,omitempty"`
	// The date time when the user got locked out. A local user is locked out of the system after X amount of consecutive failed login attempts, where X is configurable in the local Identity Provider. The lock is in effect for Y minute(s), where Y is configurable in the local Identity Provider. When the user logs in successfully, this field becomes null.
	LockStart *time.Time `json:"lockStart,omitempty"`
}

// NewLocalUsersGetRequest instantiates a new LocalUsersGetRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLocalUsersGetRequest(name string, firstName string, lastName string, password string) *LocalUsersGetRequest {
	this := LocalUsersGetRequest{}
	this.Name = name
	this.FirstName = firstName
	this.LastName = lastName
	this.Password = password
	return &this
}

// NewLocalUsersGetRequestWithDefaults instantiates a new LocalUsersGetRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLocalUsersGetRequestWithDefaults() *LocalUsersGetRequest {
	this := LocalUsersGetRequest{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *LocalUsersGetRequest) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocalUsersGetRequest) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *LocalUsersGetRequest) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *LocalUsersGetRequest) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *LocalUsersGetRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *LocalUsersGetRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *LocalUsersGetRequest) SetName(v string) {
	o.Name = v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *LocalUsersGetRequest) GetNotes() string {
	if o == nil || o.Notes == nil {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocalUsersGetRequest) GetNotesOk() (*string, bool) {
	if o == nil || o.Notes == nil {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *LocalUsersGetRequest) HasNotes() bool {
	if o != nil && o.Notes != nil {
		return true
	}

	return false
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *LocalUsersGetRequest) SetNotes(v string) {
	o.Notes = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *LocalUsersGetRequest) GetCreated() time.Time {
	if o == nil || o.Created == nil {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocalUsersGetRequest) GetCreatedOk() (*time.Time, bool) {
	if o == nil || o.Created == nil {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *LocalUsersGetRequest) HasCreated() bool {
	if o != nil && o.Created != nil {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *LocalUsersGetRequest) SetCreated(v time.Time) {
	o.Created = &v
}

// GetUpdated returns the Updated field value if set, zero value otherwise.
func (o *LocalUsersGetRequest) GetUpdated() time.Time {
	if o == nil || o.Updated == nil {
		var ret time.Time
		return ret
	}
	return *o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocalUsersGetRequest) GetUpdatedOk() (*time.Time, bool) {
	if o == nil || o.Updated == nil {
		return nil, false
	}
	return o.Updated, true
}

// HasUpdated returns a boolean if a field has been set.
func (o *LocalUsersGetRequest) HasUpdated() bool {
	if o != nil && o.Updated != nil {
		return true
	}

	return false
}

// SetUpdated gets a reference to the given time.Time and assigns it to the Updated field.
func (o *LocalUsersGetRequest) SetUpdated(v time.Time) {
	o.Updated = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *LocalUsersGetRequest) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocalUsersGetRequest) GetTagsOk() ([]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *LocalUsersGetRequest) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *LocalUsersGetRequest) SetTags(v []string) {
	o.Tags = v
}

// GetFirstName returns the FirstName field value
func (o *LocalUsersGetRequest) GetFirstName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value
// and a boolean to check if the value has been set.
func (o *LocalUsersGetRequest) GetFirstNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FirstName, true
}

// SetFirstName sets field value
func (o *LocalUsersGetRequest) SetFirstName(v string) {
	o.FirstName = v
}

// GetLastName returns the LastName field value
func (o *LocalUsersGetRequest) GetLastName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value
// and a boolean to check if the value has been set.
func (o *LocalUsersGetRequest) GetLastNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastName, true
}

// SetLastName sets field value
func (o *LocalUsersGetRequest) SetLastName(v string) {
	o.LastName = v
}

// GetPassword returns the Password field value
func (o *LocalUsersGetRequest) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *LocalUsersGetRequest) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *LocalUsersGetRequest) SetPassword(v string) {
	o.Password = v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *LocalUsersGetRequest) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocalUsersGetRequest) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *LocalUsersGetRequest) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *LocalUsersGetRequest) SetEmail(v string) {
	o.Email = &v
}

// GetPhone returns the Phone field value if set, zero value otherwise.
func (o *LocalUsersGetRequest) GetPhone() string {
	if o == nil || o.Phone == nil {
		var ret string
		return ret
	}
	return *o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocalUsersGetRequest) GetPhoneOk() (*string, bool) {
	if o == nil || o.Phone == nil {
		return nil, false
	}
	return o.Phone, true
}

// HasPhone returns a boolean if a field has been set.
func (o *LocalUsersGetRequest) HasPhone() bool {
	if o != nil && o.Phone != nil {
		return true
	}

	return false
}

// SetPhone gets a reference to the given string and assigns it to the Phone field.
func (o *LocalUsersGetRequest) SetPhone(v string) {
	o.Phone = &v
}

// GetDisabled returns the Disabled field value if set, zero value otherwise.
func (o *LocalUsersGetRequest) GetDisabled() bool {
	if o == nil || o.Disabled == nil {
		var ret bool
		return ret
	}
	return *o.Disabled
}

// GetDisabledOk returns a tuple with the Disabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocalUsersGetRequest) GetDisabledOk() (*bool, bool) {
	if o == nil || o.Disabled == nil {
		return nil, false
	}
	return o.Disabled, true
}

// HasDisabled returns a boolean if a field has been set.
func (o *LocalUsersGetRequest) HasDisabled() bool {
	if o != nil && o.Disabled != nil {
		return true
	}

	return false
}

// SetDisabled gets a reference to the given bool and assigns it to the Disabled field.
func (o *LocalUsersGetRequest) SetDisabled(v bool) {
	o.Disabled = &v
}

// GetFailedLoginAttempts returns the FailedLoginAttempts field value if set, zero value otherwise.
func (o *LocalUsersGetRequest) GetFailedLoginAttempts() float32 {
	if o == nil || o.FailedLoginAttempts == nil {
		var ret float32
		return ret
	}
	return *o.FailedLoginAttempts
}

// GetFailedLoginAttemptsOk returns a tuple with the FailedLoginAttempts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocalUsersGetRequest) GetFailedLoginAttemptsOk() (*float32, bool) {
	if o == nil || o.FailedLoginAttempts == nil {
		return nil, false
	}
	return o.FailedLoginAttempts, true
}

// HasFailedLoginAttempts returns a boolean if a field has been set.
func (o *LocalUsersGetRequest) HasFailedLoginAttempts() bool {
	if o != nil && o.FailedLoginAttempts != nil {
		return true
	}

	return false
}

// SetFailedLoginAttempts gets a reference to the given float32 and assigns it to the FailedLoginAttempts field.
func (o *LocalUsersGetRequest) SetFailedLoginAttempts(v float32) {
	o.FailedLoginAttempts = &v
}

// GetLockStart returns the LockStart field value if set, zero value otherwise.
func (o *LocalUsersGetRequest) GetLockStart() time.Time {
	if o == nil || o.LockStart == nil {
		var ret time.Time
		return ret
	}
	return *o.LockStart
}

// GetLockStartOk returns a tuple with the LockStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocalUsersGetRequest) GetLockStartOk() (*time.Time, bool) {
	if o == nil || o.LockStart == nil {
		return nil, false
	}
	return o.LockStart, true
}

// HasLockStart returns a boolean if a field has been set.
func (o *LocalUsersGetRequest) HasLockStart() bool {
	if o != nil && o.LockStart != nil {
		return true
	}

	return false
}

// SetLockStart gets a reference to the given time.Time and assigns it to the LockStart field.
func (o *LocalUsersGetRequest) SetLockStart(v time.Time) {
	o.LockStart = &v
}

func (o LocalUsersGetRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Notes != nil {
		toSerialize["notes"] = o.Notes
	}
	if o.Created != nil {
		toSerialize["created"] = o.Created
	}
	if o.Updated != nil {
		toSerialize["updated"] = o.Updated
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
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

type NullableLocalUsersGetRequest struct {
	value *LocalUsersGetRequest
	isSet bool
}

func (v NullableLocalUsersGetRequest) Get() *LocalUsersGetRequest {
	return v.value
}

func (v *NullableLocalUsersGetRequest) Set(val *LocalUsersGetRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableLocalUsersGetRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableLocalUsersGetRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocalUsersGetRequest(val *LocalUsersGetRequest) *NullableLocalUsersGetRequest {
	return &NullableLocalUsersGetRequest{value: val, isSet: true}
}

func (v NullableLocalUsersGetRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocalUsersGetRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
