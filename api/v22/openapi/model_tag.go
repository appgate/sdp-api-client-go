/*
Appgate SDP Controller REST API

# About   This specification documents the REST API calls for the Appgate SDP Controller.    Please refer to the REST API chapter in the manual or contact Appgate support with any questions about   this functionality. # Getting Started   Requirements for API scripting:   - Access to the Admin/API TLS Connection (default port 8443) of a Controller appliance.     (https://sdphelp.appgate.com/adminguide/appliance-function-configure.html?anchor=admin-api)   - An API user with relevant permissions.     (https://sdphelp.appgate.com/adminguide/administrative-roles-configure.html)   - In order to use the simple login API, Admin MFA must be disabled or the API user must be excluded.     (https://sdphelp.appgate.com/adminguide/mfa-for-admins.html) # Base path   HTTPS requests must be sent to the Admin Interface hostname and port, with **_/admin** path.    For example: **https://appgate.company.com:8443/admin**    All requests must have the **Accept** header as:    **application/vnd.appgate.peer-v22+json**    An exception is made for the **_/admin/version** endpoint which instead expects an **application/json** Accept header. # API Conventions   API conventions are  important to understand and follow strictly.    - While updating objects (via PUT), entire object must be sent with all fields.     - For example, in order to add a remedy method to the condition below:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": []       }       ```     - send the entire object with updated and non-updated fields:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": [{\"type\": \"DisplayMessage\", \"message\": \"test message\"}]       }       ```    - In case Controller returns an error (non-2xx HTTP status code), response body is JSON.     The \"message\" field contains information about the error.     HTTP 422 \"Unprocessable Entity\" has extra `errors` field to list all the issues with specific fields.    - Empty string (\"\") is considered a different value than \"null\" or field being omitted from JSON.     Omitting the field is recommended if no value is intended.     Empty string (\"\") will be almost always rejected as invalid value.    - There are common pattern between many objects:     - **Configuration Objects**: There are many objects with common fields, namely \"id\", \"name\", \"notes\", \"created\"       and \"updated\". These entities are listed, queried, created, updated and deleted in a similar fashion.     - **Distinguished Name**: Users and Devices are identified with what is called Distinguished Names, as used in        LDAP. The distinguished format that identifies a device and a user combination is        \"CN=\\<Device ID\\>,CN=\\<username\\>,OU=\\<Identity Provider Name\\>\". Some objects have the        \"userDistinguishedName\" field, which does not include the CN for Device ID.        This identifies a user on every device.

API version: API version 22.2
Contact: appgatesdp.support@appgate.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// Tag Represents a Tag.
type Tag struct {
	// ID of the object.
	Id *string `json:"id,omitempty"`
	// Name of the object.
	Name *string `json:"name,omitempty"`
	// Notes for the object. Used for documentation purposes.
	Notes *string `json:"notes,omitempty"`
	// The color code of the tag. Default will be random color. - 1: Light Green - 2: Green - 3: Indigo - 4: Deep Purple - 5: Yellow - 6: Lime - 7: Light Blue - 8: Blue - 9: Amber - 10: Orange - 11: Cyan - 12: Teal - 13: Deep Orange - 14: Red - 15: Gray - 16: Brown - 17: Pink - 18: Purple - 19: Blue Gray - 20: Near Black
	ColorCode *int32 `json:"colorCode,omitempty"`
	// The name associated with the color code.
	ColorName *string `json:"colorName,omitempty"`
	// Create date.
	Created *time.Time `json:"created,omitempty"`
	// Last update date.
	Updated *time.Time `json:"updated,omitempty"`
	// The number of entities tagged with this Tag.
	TaggedOccurrences *int32 `json:"taggedOccurrences,omitempty"`
	// The number of entities linked via this Tag.
	LinkedOccurrences *int32 `json:"linkedOccurrences,omitempty"`
}

// NewTag instantiates a new Tag object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTag() *Tag {
	this := Tag{}
	return &this
}

// NewTagWithDefaults instantiates a new Tag object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTagWithDefaults() *Tag {
	this := Tag{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Tag) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tag) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Tag) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Tag) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Tag) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tag) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Tag) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Tag) SetName(v string) {
	o.Name = &v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *Tag) GetNotes() string {
	if o == nil || o.Notes == nil {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tag) GetNotesOk() (*string, bool) {
	if o == nil || o.Notes == nil {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *Tag) HasNotes() bool {
	if o != nil && o.Notes != nil {
		return true
	}

	return false
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *Tag) SetNotes(v string) {
	o.Notes = &v
}

// GetColorCode returns the ColorCode field value if set, zero value otherwise.
func (o *Tag) GetColorCode() int32 {
	if o == nil || o.ColorCode == nil {
		var ret int32
		return ret
	}
	return *o.ColorCode
}

// GetColorCodeOk returns a tuple with the ColorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tag) GetColorCodeOk() (*int32, bool) {
	if o == nil || o.ColorCode == nil {
		return nil, false
	}
	return o.ColorCode, true
}

// HasColorCode returns a boolean if a field has been set.
func (o *Tag) HasColorCode() bool {
	if o != nil && o.ColorCode != nil {
		return true
	}

	return false
}

// SetColorCode gets a reference to the given int32 and assigns it to the ColorCode field.
func (o *Tag) SetColorCode(v int32) {
	o.ColorCode = &v
}

// GetColorName returns the ColorName field value if set, zero value otherwise.
func (o *Tag) GetColorName() string {
	if o == nil || o.ColorName == nil {
		var ret string
		return ret
	}
	return *o.ColorName
}

// GetColorNameOk returns a tuple with the ColorName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tag) GetColorNameOk() (*string, bool) {
	if o == nil || o.ColorName == nil {
		return nil, false
	}
	return o.ColorName, true
}

// HasColorName returns a boolean if a field has been set.
func (o *Tag) HasColorName() bool {
	if o != nil && o.ColorName != nil {
		return true
	}

	return false
}

// SetColorName gets a reference to the given string and assigns it to the ColorName field.
func (o *Tag) SetColorName(v string) {
	o.ColorName = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *Tag) GetCreated() time.Time {
	if o == nil || o.Created == nil {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tag) GetCreatedOk() (*time.Time, bool) {
	if o == nil || o.Created == nil {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *Tag) HasCreated() bool {
	if o != nil && o.Created != nil {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *Tag) SetCreated(v time.Time) {
	o.Created = &v
}

// GetUpdated returns the Updated field value if set, zero value otherwise.
func (o *Tag) GetUpdated() time.Time {
	if o == nil || o.Updated == nil {
		var ret time.Time
		return ret
	}
	return *o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tag) GetUpdatedOk() (*time.Time, bool) {
	if o == nil || o.Updated == nil {
		return nil, false
	}
	return o.Updated, true
}

// HasUpdated returns a boolean if a field has been set.
func (o *Tag) HasUpdated() bool {
	if o != nil && o.Updated != nil {
		return true
	}

	return false
}

// SetUpdated gets a reference to the given time.Time and assigns it to the Updated field.
func (o *Tag) SetUpdated(v time.Time) {
	o.Updated = &v
}

// GetTaggedOccurrences returns the TaggedOccurrences field value if set, zero value otherwise.
func (o *Tag) GetTaggedOccurrences() int32 {
	if o == nil || o.TaggedOccurrences == nil {
		var ret int32
		return ret
	}
	return *o.TaggedOccurrences
}

// GetTaggedOccurrencesOk returns a tuple with the TaggedOccurrences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tag) GetTaggedOccurrencesOk() (*int32, bool) {
	if o == nil || o.TaggedOccurrences == nil {
		return nil, false
	}
	return o.TaggedOccurrences, true
}

// HasTaggedOccurrences returns a boolean if a field has been set.
func (o *Tag) HasTaggedOccurrences() bool {
	if o != nil && o.TaggedOccurrences != nil {
		return true
	}

	return false
}

// SetTaggedOccurrences gets a reference to the given int32 and assigns it to the TaggedOccurrences field.
func (o *Tag) SetTaggedOccurrences(v int32) {
	o.TaggedOccurrences = &v
}

// GetLinkedOccurrences returns the LinkedOccurrences field value if set, zero value otherwise.
func (o *Tag) GetLinkedOccurrences() int32 {
	if o == nil || o.LinkedOccurrences == nil {
		var ret int32
		return ret
	}
	return *o.LinkedOccurrences
}

// GetLinkedOccurrencesOk returns a tuple with the LinkedOccurrences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tag) GetLinkedOccurrencesOk() (*int32, bool) {
	if o == nil || o.LinkedOccurrences == nil {
		return nil, false
	}
	return o.LinkedOccurrences, true
}

// HasLinkedOccurrences returns a boolean if a field has been set.
func (o *Tag) HasLinkedOccurrences() bool {
	if o != nil && o.LinkedOccurrences != nil {
		return true
	}

	return false
}

// SetLinkedOccurrences gets a reference to the given int32 and assigns it to the LinkedOccurrences field.
func (o *Tag) SetLinkedOccurrences(v int32) {
	o.LinkedOccurrences = &v
}

func (o Tag) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Notes != nil {
		toSerialize["notes"] = o.Notes
	}
	if o.ColorCode != nil {
		toSerialize["colorCode"] = o.ColorCode
	}
	if o.ColorName != nil {
		toSerialize["colorName"] = o.ColorName
	}
	if o.Created != nil {
		toSerialize["created"] = o.Created
	}
	if o.Updated != nil {
		toSerialize["updated"] = o.Updated
	}
	if o.TaggedOccurrences != nil {
		toSerialize["taggedOccurrences"] = o.TaggedOccurrences
	}
	if o.LinkedOccurrences != nil {
		toSerialize["linkedOccurrences"] = o.LinkedOccurrences
	}
	return json.Marshal(toSerialize)
}

type NullableTag struct {
	value *Tag
	isSet bool
}

func (v NullableTag) Get() *Tag {
	return v.value
}

func (v *NullableTag) Set(val *Tag) {
	v.value = val
	v.isSet = true
}

func (v NullableTag) IsSet() bool {
	return v.isSet
}

func (v *NullableTag) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTag(val *Tag) *NullableTag {
	return &NullableTag{value: val, isSet: true}
}

func (v NullableTag) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTag) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
