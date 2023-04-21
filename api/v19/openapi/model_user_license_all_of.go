/*
Appgate SDP Controller REST API

# About   This specification documents the REST API calls for the Appgate SDP Controller.    Please refer to the REST API chapter in the manual or contact Appgate support with any questions about   this functionality. # Getting Started   Requirements for API scripting:   - Access to the Admin/API TLS Connection (default port 8443) of a Controller appliance.     (https://sdphelp.appgate.com/adminguide/appliance-function-configure.html?anchor=admin-api)   - An API user with relevant permissions.     (https://sdphelp.appgate.com/adminguide/administrative-roles-configure.html)   - In order to use the simple login API, Admin MFA must be disabled or the API user must be excluded.     (https://sdphelp.appgate.com/adminguide/mfa-for-admins.html) # Base path   HTTPS requests must be sent to the Admin Interface hostname and port, with **_/admin** path.    For example: **https://appgate.company.com:8443/admin**    All requests must have the **Accept** header as:    **application/vnd.appgate.peer-v19+json** # API Conventions   API conventions are  important to understand and follow strictly.    - While updating objects (via PUT), entire object must be sent with all fields.     - For example, in order to add a remedy method to the condition below:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": []       }       ```     - send the entire object with updated and non-updated fields:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": [{\"type\": \"DisplayMessage\", \"message\": \"test message\"}]       }       ```    - In case Controller returns an error (non-2xx HTTP status code), response body is JSON.     The \"message\" field contains information about the error.     HTTP 422 \"Unprocessable Entity\" has extra `errors` field to list all the issues with specific fields.    - Empty string (\"\") is considered a different value than \"null\" or field being omitted from JSON.     Omitting the field is recommended if no value is intended.     Empty string (\"\") will be almost always rejected as invalid value.    - There are common pattern between many objects:     - **Configuration Objects**: There are many objects with common fields, namely \"id\", \"name\", \"notes\", \"created\"       and \"updated\". These entities are listed, queried, created, updated and deleted in a similar fashion.     - **Distinguished Name**: Users and Devices are identified with what is called Distinguished Names, as used in        LDAP. The distinguished format that identifies a device and a user combination is        \"CN=\\<Device ID\\>,CN=\\<username\\>,OU=\\<Identity Provider Name\\>\". Some objects have the        \"userDistinguishedName\" field, which does not include the CN for Device ID.        This identifies a user on every device.

API version: API version 19.0
Contact: appgatesdp.support@appgate.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// UserLicenseAllOf struct for UserLicenseAllOf
type UserLicenseAllOf struct {
	// Type of User License.
	Type *string `json:"type,omitempty"`
	// Creation date.
	Created *time.Time `json:"created,omitempty"`
	// The time when the user last signed in.
	LastSeenAt *time.Time `json:"lastSeenAt,omitempty"`
}

// NewUserLicenseAllOf instantiates a new UserLicenseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserLicenseAllOf() *UserLicenseAllOf {
	this := UserLicenseAllOf{}
	return &this
}

// NewUserLicenseAllOfWithDefaults instantiates a new UserLicenseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserLicenseAllOfWithDefaults() *UserLicenseAllOf {
	this := UserLicenseAllOf{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *UserLicenseAllOf) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserLicenseAllOf) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *UserLicenseAllOf) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *UserLicenseAllOf) SetType(v string) {
	o.Type = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *UserLicenseAllOf) GetCreated() time.Time {
	if o == nil || o.Created == nil {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserLicenseAllOf) GetCreatedOk() (*time.Time, bool) {
	if o == nil || o.Created == nil {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *UserLicenseAllOf) HasCreated() bool {
	if o != nil && o.Created != nil {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *UserLicenseAllOf) SetCreated(v time.Time) {
	o.Created = &v
}

// GetLastSeenAt returns the LastSeenAt field value if set, zero value otherwise.
func (o *UserLicenseAllOf) GetLastSeenAt() time.Time {
	if o == nil || o.LastSeenAt == nil {
		var ret time.Time
		return ret
	}
	return *o.LastSeenAt
}

// GetLastSeenAtOk returns a tuple with the LastSeenAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserLicenseAllOf) GetLastSeenAtOk() (*time.Time, bool) {
	if o == nil || o.LastSeenAt == nil {
		return nil, false
	}
	return o.LastSeenAt, true
}

// HasLastSeenAt returns a boolean if a field has been set.
func (o *UserLicenseAllOf) HasLastSeenAt() bool {
	if o != nil && o.LastSeenAt != nil {
		return true
	}

	return false
}

// SetLastSeenAt gets a reference to the given time.Time and assigns it to the LastSeenAt field.
func (o *UserLicenseAllOf) SetLastSeenAt(v time.Time) {
	o.LastSeenAt = &v
}

func (o UserLicenseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Created != nil {
		toSerialize["created"] = o.Created
	}
	if o.LastSeenAt != nil {
		toSerialize["lastSeenAt"] = o.LastSeenAt
	}
	return json.Marshal(toSerialize)
}

type NullableUserLicenseAllOf struct {
	value *UserLicenseAllOf
	isSet bool
}

func (v NullableUserLicenseAllOf) Get() *UserLicenseAllOf {
	return v.value
}

func (v *NullableUserLicenseAllOf) Set(val *UserLicenseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableUserLicenseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableUserLicenseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserLicenseAllOf(val *UserLicenseAllOf) *NullableUserLicenseAllOf {
	return &NullableUserLicenseAllOf{value: val, isSet: true}
}

func (v NullableUserLicenseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserLicenseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
