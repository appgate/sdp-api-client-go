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
	"time"
)

// AdminMessageMetadata struct for AdminMessageMetadata
type AdminMessageMetadata struct {
	// The uuid of the message metadata.
	Id *string `json:"id,omitempty"`
	// The message text.
	Message string `json:"message"`
	// The status of the message.
	Status string `json:"status"`
	// When the message metadata was updated.
	AppliedAt *time.Time `json:"appliedAt,omitempty"`
	// The distinguished name of who updated the message metadata.
	AppliedBy *string `json:"appliedBy,omitempty"`
	// The username of the user who updated the message metadata.
	AppliedByUsername *string `json:"appliedByUsername,omitempty"`
	// The provider name of the user who updated the message metadata.
	AppliedByProviderName *string `json:"appliedByProviderName,omitempty"`
	// The device ID of the user who updated the message metadata.
	AppliedByDeviceId *string `json:"appliedByDeviceId,omitempty"`
}

// NewAdminMessageMetadata instantiates a new AdminMessageMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdminMessageMetadata(message string, status string) *AdminMessageMetadata {
	this := AdminMessageMetadata{}
	this.Message = message
	this.Status = status
	return &this
}

// NewAdminMessageMetadataWithDefaults instantiates a new AdminMessageMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdminMessageMetadataWithDefaults() *AdminMessageMetadata {
	this := AdminMessageMetadata{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AdminMessageMetadata) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdminMessageMetadata) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AdminMessageMetadata) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AdminMessageMetadata) SetId(v string) {
	o.Id = &v
}

// GetMessage returns the Message field value
func (o *AdminMessageMetadata) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *AdminMessageMetadata) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *AdminMessageMetadata) SetMessage(v string) {
	o.Message = v
}

// GetStatus returns the Status field value
func (o *AdminMessageMetadata) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *AdminMessageMetadata) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *AdminMessageMetadata) SetStatus(v string) {
	o.Status = v
}

// GetAppliedAt returns the AppliedAt field value if set, zero value otherwise.
func (o *AdminMessageMetadata) GetAppliedAt() time.Time {
	if o == nil || o.AppliedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.AppliedAt
}

// GetAppliedAtOk returns a tuple with the AppliedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdminMessageMetadata) GetAppliedAtOk() (*time.Time, bool) {
	if o == nil || o.AppliedAt == nil {
		return nil, false
	}
	return o.AppliedAt, true
}

// HasAppliedAt returns a boolean if a field has been set.
func (o *AdminMessageMetadata) HasAppliedAt() bool {
	if o != nil && o.AppliedAt != nil {
		return true
	}

	return false
}

// SetAppliedAt gets a reference to the given time.Time and assigns it to the AppliedAt field.
func (o *AdminMessageMetadata) SetAppliedAt(v time.Time) {
	o.AppliedAt = &v
}

// GetAppliedBy returns the AppliedBy field value if set, zero value otherwise.
func (o *AdminMessageMetadata) GetAppliedBy() string {
	if o == nil || o.AppliedBy == nil {
		var ret string
		return ret
	}
	return *o.AppliedBy
}

// GetAppliedByOk returns a tuple with the AppliedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdminMessageMetadata) GetAppliedByOk() (*string, bool) {
	if o == nil || o.AppliedBy == nil {
		return nil, false
	}
	return o.AppliedBy, true
}

// HasAppliedBy returns a boolean if a field has been set.
func (o *AdminMessageMetadata) HasAppliedBy() bool {
	if o != nil && o.AppliedBy != nil {
		return true
	}

	return false
}

// SetAppliedBy gets a reference to the given string and assigns it to the AppliedBy field.
func (o *AdminMessageMetadata) SetAppliedBy(v string) {
	o.AppliedBy = &v
}

// GetAppliedByUsername returns the AppliedByUsername field value if set, zero value otherwise.
func (o *AdminMessageMetadata) GetAppliedByUsername() string {
	if o == nil || o.AppliedByUsername == nil {
		var ret string
		return ret
	}
	return *o.AppliedByUsername
}

// GetAppliedByUsernameOk returns a tuple with the AppliedByUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdminMessageMetadata) GetAppliedByUsernameOk() (*string, bool) {
	if o == nil || o.AppliedByUsername == nil {
		return nil, false
	}
	return o.AppliedByUsername, true
}

// HasAppliedByUsername returns a boolean if a field has been set.
func (o *AdminMessageMetadata) HasAppliedByUsername() bool {
	if o != nil && o.AppliedByUsername != nil {
		return true
	}

	return false
}

// SetAppliedByUsername gets a reference to the given string and assigns it to the AppliedByUsername field.
func (o *AdminMessageMetadata) SetAppliedByUsername(v string) {
	o.AppliedByUsername = &v
}

// GetAppliedByProviderName returns the AppliedByProviderName field value if set, zero value otherwise.
func (o *AdminMessageMetadata) GetAppliedByProviderName() string {
	if o == nil || o.AppliedByProviderName == nil {
		var ret string
		return ret
	}
	return *o.AppliedByProviderName
}

// GetAppliedByProviderNameOk returns a tuple with the AppliedByProviderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdminMessageMetadata) GetAppliedByProviderNameOk() (*string, bool) {
	if o == nil || o.AppliedByProviderName == nil {
		return nil, false
	}
	return o.AppliedByProviderName, true
}

// HasAppliedByProviderName returns a boolean if a field has been set.
func (o *AdminMessageMetadata) HasAppliedByProviderName() bool {
	if o != nil && o.AppliedByProviderName != nil {
		return true
	}

	return false
}

// SetAppliedByProviderName gets a reference to the given string and assigns it to the AppliedByProviderName field.
func (o *AdminMessageMetadata) SetAppliedByProviderName(v string) {
	o.AppliedByProviderName = &v
}

// GetAppliedByDeviceId returns the AppliedByDeviceId field value if set, zero value otherwise.
func (o *AdminMessageMetadata) GetAppliedByDeviceId() string {
	if o == nil || o.AppliedByDeviceId == nil {
		var ret string
		return ret
	}
	return *o.AppliedByDeviceId
}

// GetAppliedByDeviceIdOk returns a tuple with the AppliedByDeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdminMessageMetadata) GetAppliedByDeviceIdOk() (*string, bool) {
	if o == nil || o.AppliedByDeviceId == nil {
		return nil, false
	}
	return o.AppliedByDeviceId, true
}

// HasAppliedByDeviceId returns a boolean if a field has been set.
func (o *AdminMessageMetadata) HasAppliedByDeviceId() bool {
	if o != nil && o.AppliedByDeviceId != nil {
		return true
	}

	return false
}

// SetAppliedByDeviceId gets a reference to the given string and assigns it to the AppliedByDeviceId field.
func (o *AdminMessageMetadata) SetAppliedByDeviceId(v string) {
	o.AppliedByDeviceId = &v
}

func (o AdminMessageMetadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["message"] = o.Message
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if o.AppliedAt != nil {
		toSerialize["appliedAt"] = o.AppliedAt
	}
	if o.AppliedBy != nil {
		toSerialize["appliedBy"] = o.AppliedBy
	}
	if o.AppliedByUsername != nil {
		toSerialize["appliedByUsername"] = o.AppliedByUsername
	}
	if o.AppliedByProviderName != nil {
		toSerialize["appliedByProviderName"] = o.AppliedByProviderName
	}
	if o.AppliedByDeviceId != nil {
		toSerialize["appliedByDeviceId"] = o.AppliedByDeviceId
	}
	return json.Marshal(toSerialize)
}

type NullableAdminMessageMetadata struct {
	value *AdminMessageMetadata
	isSet bool
}

func (v NullableAdminMessageMetadata) Get() *AdminMessageMetadata {
	return v.value
}

func (v *NullableAdminMessageMetadata) Set(val *AdminMessageMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableAdminMessageMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableAdminMessageMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdminMessageMetadata(val *AdminMessageMetadata) *NullableAdminMessageMetadata {
	return &NullableAdminMessageMetadata{value: val, isSet: true}
}

func (v NullableAdminMessageMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdminMessageMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
