/*
Appgate SDP Controller REST API

# About   This specification documents the REST API calls for the Appgate SDP Controller.    Please refer to the REST API chapter in the manual or contact Appgate support with any questions about   this functionality. # Getting Started   Requirements for API scripting:   - Access to the Admin/API TLS Connection (default port 8443) of a Controller appliance.     (https://sdphelp.appgate.com/adminguide/appliance-function-configure.html?anchor=admin-api)   - An API user with relevant permissions.     (https://sdphelp.appgate.com/adminguide/administrative-roles-configure.html)   - In order to use the simple login API, Admin MFA must be disabled or the API user must be excluded.     (https://sdphelp.appgate.com/adminguide/mfa-for-admins.html) # Base path   HTTPS requests must be sent to the Admin Interface hostname and port, with **_/admin** path.    For example: **https://appgate.company.com:8443/admin**    All requests must have the **Accept** header as:    **application/vnd.appgate.peer-v19+json** # API Conventions   API conventions are  important to understand and follow strictly.    - While updating objects (via PUT), entire object must be sent with all fields.     - For example, in order to add a remedy method to the condition below:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": []       }       ```     - send the entire object with updated and non-updated fields:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": [{\"type\": \"DisplayMessage\", \"message\": \"test message\"}]       }       ```    - In case Controller returns an error (non-2xx HTTP status code), response body is JSON.     The \"message\" field contains information about the error.     HTTP 422 \"Unprocessable Entity\" has extra `errors` field to list all the issues with specific fields.    - Empty string (\"\") is considered a different value than \"null\" or field being omitted from JSON.     Omitting the field is recommended if no value is intended.     Empty string (\"\") will be almost always rejected as invalid value.    - There are common pattern between many objects:     - **Configuration Objects**: There are many objects with common fields, namely \"id\", \"name\", \"notes\", \"created\"       and \"updated\". These entities are listed, queried, created, updated and deleted in a similar fashion.     - **Distinguished Name**: Users and Devices are identified with what is called Distinguished Names, as used in        LDAP. The distinguished format that identifies a device and a user combination is        \"CN=\\<Device ID\\>,CN=\\<username\\>,OU=\\<Identity Provider Name\\>\". Some objects have the        \"userDistinguishedName\" field, which does not include the CN for Device ID.        This identifies a user on every device.

API version: API version 19.1
Contact: appgatesdp.support@appgate.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// AdminMessagesListDataInner struct for AdminMessagesListDataInner
type AdminMessagesListDataInner struct {
	// Randomly generated UUID.
	Id *string `json:"id,omitempty"`
	// The severity of the message.
	Level *string `json:"level,omitempty"`
	// Message category.
	Category *string `json:"category,omitempty"`
	// The message.
	Message *string `json:"message,omitempty"`
	// The name of the Appliance that generated this message.
	Source *string `json:"source,omitempty"`
	// When the Admin Message generated.
	Created *time.Time `json:"created,omitempty"`
}

// NewAdminMessagesListDataInner instantiates a new AdminMessagesListDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdminMessagesListDataInner() *AdminMessagesListDataInner {
	this := AdminMessagesListDataInner{}
	return &this
}

// NewAdminMessagesListDataInnerWithDefaults instantiates a new AdminMessagesListDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdminMessagesListDataInnerWithDefaults() *AdminMessagesListDataInner {
	this := AdminMessagesListDataInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AdminMessagesListDataInner) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdminMessagesListDataInner) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AdminMessagesListDataInner) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AdminMessagesListDataInner) SetId(v string) {
	o.Id = &v
}

// GetLevel returns the Level field value if set, zero value otherwise.
func (o *AdminMessagesListDataInner) GetLevel() string {
	if o == nil || o.Level == nil {
		var ret string
		return ret
	}
	return *o.Level
}

// GetLevelOk returns a tuple with the Level field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdminMessagesListDataInner) GetLevelOk() (*string, bool) {
	if o == nil || o.Level == nil {
		return nil, false
	}
	return o.Level, true
}

// HasLevel returns a boolean if a field has been set.
func (o *AdminMessagesListDataInner) HasLevel() bool {
	if o != nil && o.Level != nil {
		return true
	}

	return false
}

// SetLevel gets a reference to the given string and assigns it to the Level field.
func (o *AdminMessagesListDataInner) SetLevel(v string) {
	o.Level = &v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *AdminMessagesListDataInner) GetCategory() string {
	if o == nil || o.Category == nil {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdminMessagesListDataInner) GetCategoryOk() (*string, bool) {
	if o == nil || o.Category == nil {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *AdminMessagesListDataInner) HasCategory() bool {
	if o != nil && o.Category != nil {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *AdminMessagesListDataInner) SetCategory(v string) {
	o.Category = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *AdminMessagesListDataInner) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdminMessagesListDataInner) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *AdminMessagesListDataInner) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *AdminMessagesListDataInner) SetMessage(v string) {
	o.Message = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *AdminMessagesListDataInner) GetSource() string {
	if o == nil || o.Source == nil {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdminMessagesListDataInner) GetSourceOk() (*string, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *AdminMessagesListDataInner) HasSource() bool {
	if o != nil && o.Source != nil {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *AdminMessagesListDataInner) SetSource(v string) {
	o.Source = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *AdminMessagesListDataInner) GetCreated() time.Time {
	if o == nil || o.Created == nil {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdminMessagesListDataInner) GetCreatedOk() (*time.Time, bool) {
	if o == nil || o.Created == nil {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *AdminMessagesListDataInner) HasCreated() bool {
	if o != nil && o.Created != nil {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *AdminMessagesListDataInner) SetCreated(v time.Time) {
	o.Created = &v
}

func (o AdminMessagesListDataInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Level != nil {
		toSerialize["level"] = o.Level
	}
	if o.Category != nil {
		toSerialize["category"] = o.Category
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.Source != nil {
		toSerialize["source"] = o.Source
	}
	if o.Created != nil {
		toSerialize["created"] = o.Created
	}
	return json.Marshal(toSerialize)
}

type NullableAdminMessagesListDataInner struct {
	value *AdminMessagesListDataInner
	isSet bool
}

func (v NullableAdminMessagesListDataInner) Get() *AdminMessagesListDataInner {
	return v.value
}

func (v *NullableAdminMessagesListDataInner) Set(val *AdminMessagesListDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableAdminMessagesListDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableAdminMessagesListDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdminMessagesListDataInner(val *AdminMessagesListDataInner) *NullableAdminMessagesListDataInner {
	return &NullableAdminMessagesListDataInner{value: val, isSet: true}
}

func (v NullableAdminMessagesListDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdminMessagesListDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
