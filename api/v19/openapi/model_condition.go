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

// Condition struct for Condition
type Condition struct {
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
	// Boolean expression in JavaScript.
	Expression string `json:"expression"`
	// A list of schedules that decides when to reevaluate the Condition. All the scheduled times will be effective. One will not override the other. - It can be a time of the day, e.g. 13:00, 10:25, 2:10 etc. - It can be one of the predefined   intervals, e.g. 1m, 5m, 15m, 1h. These intervals   will be always rounded up, i.e. if it's 15m and the   time is 12:07 when the Condition is evaluated   first, then the next evaluation will occur at   12:15, and the next one will be at   12:30 and so on.
	RepeatSchedules []string `json:"repeatSchedules,omitempty"`
	// Whether all the Remedy Methods must succeed to pass this Condition or just one.
	RemedyLogic *string `json:"remedyLogic,omitempty"`
	// The remedy methods that will be triggered if the evaluation fails.
	RemedyMethods []RemedyMethod `json:"remedyMethods,omitempty"`
}

// NewCondition instantiates a new Condition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCondition(name string, expression string) *Condition {
	this := Condition{}
	this.Name = name
	this.Expression = expression
	var remedyLogic string = "and"
	this.RemedyLogic = &remedyLogic
	return &this
}

// NewConditionWithDefaults instantiates a new Condition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConditionWithDefaults() *Condition {
	this := Condition{}
	var remedyLogic string = "and"
	this.RemedyLogic = &remedyLogic
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Condition) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Condition) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Condition) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Condition) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *Condition) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Condition) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Condition) SetName(v string) {
	o.Name = v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *Condition) GetNotes() string {
	if o == nil || o.Notes == nil {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Condition) GetNotesOk() (*string, bool) {
	if o == nil || o.Notes == nil {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *Condition) HasNotes() bool {
	if o != nil && o.Notes != nil {
		return true
	}

	return false
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *Condition) SetNotes(v string) {
	o.Notes = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *Condition) GetCreated() time.Time {
	if o == nil || o.Created == nil {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Condition) GetCreatedOk() (*time.Time, bool) {
	if o == nil || o.Created == nil {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *Condition) HasCreated() bool {
	if o != nil && o.Created != nil {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *Condition) SetCreated(v time.Time) {
	o.Created = &v
}

// GetUpdated returns the Updated field value if set, zero value otherwise.
func (o *Condition) GetUpdated() time.Time {
	if o == nil || o.Updated == nil {
		var ret time.Time
		return ret
	}
	return *o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Condition) GetUpdatedOk() (*time.Time, bool) {
	if o == nil || o.Updated == nil {
		return nil, false
	}
	return o.Updated, true
}

// HasUpdated returns a boolean if a field has been set.
func (o *Condition) HasUpdated() bool {
	if o != nil && o.Updated != nil {
		return true
	}

	return false
}

// SetUpdated gets a reference to the given time.Time and assigns it to the Updated field.
func (o *Condition) SetUpdated(v time.Time) {
	o.Updated = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *Condition) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Condition) GetTagsOk() ([]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *Condition) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *Condition) SetTags(v []string) {
	o.Tags = v
}

// GetExpression returns the Expression field value
func (o *Condition) GetExpression() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value
// and a boolean to check if the value has been set.
func (o *Condition) GetExpressionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Expression, true
}

// SetExpression sets field value
func (o *Condition) SetExpression(v string) {
	o.Expression = v
}

// GetRepeatSchedules returns the RepeatSchedules field value if set, zero value otherwise.
func (o *Condition) GetRepeatSchedules() []string {
	if o == nil || o.RepeatSchedules == nil {
		var ret []string
		return ret
	}
	return o.RepeatSchedules
}

// GetRepeatSchedulesOk returns a tuple with the RepeatSchedules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Condition) GetRepeatSchedulesOk() ([]string, bool) {
	if o == nil || o.RepeatSchedules == nil {
		return nil, false
	}
	return o.RepeatSchedules, true
}

// HasRepeatSchedules returns a boolean if a field has been set.
func (o *Condition) HasRepeatSchedules() bool {
	if o != nil && o.RepeatSchedules != nil {
		return true
	}

	return false
}

// SetRepeatSchedules gets a reference to the given []string and assigns it to the RepeatSchedules field.
func (o *Condition) SetRepeatSchedules(v []string) {
	o.RepeatSchedules = v
}

// GetRemedyLogic returns the RemedyLogic field value if set, zero value otherwise.
func (o *Condition) GetRemedyLogic() string {
	if o == nil || o.RemedyLogic == nil {
		var ret string
		return ret
	}
	return *o.RemedyLogic
}

// GetRemedyLogicOk returns a tuple with the RemedyLogic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Condition) GetRemedyLogicOk() (*string, bool) {
	if o == nil || o.RemedyLogic == nil {
		return nil, false
	}
	return o.RemedyLogic, true
}

// HasRemedyLogic returns a boolean if a field has been set.
func (o *Condition) HasRemedyLogic() bool {
	if o != nil && o.RemedyLogic != nil {
		return true
	}

	return false
}

// SetRemedyLogic gets a reference to the given string and assigns it to the RemedyLogic field.
func (o *Condition) SetRemedyLogic(v string) {
	o.RemedyLogic = &v
}

// GetRemedyMethods returns the RemedyMethods field value if set, zero value otherwise.
func (o *Condition) GetRemedyMethods() []RemedyMethod {
	if o == nil || o.RemedyMethods == nil {
		var ret []RemedyMethod
		return ret
	}
	return o.RemedyMethods
}

// GetRemedyMethodsOk returns a tuple with the RemedyMethods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Condition) GetRemedyMethodsOk() ([]RemedyMethod, bool) {
	if o == nil || o.RemedyMethods == nil {
		return nil, false
	}
	return o.RemedyMethods, true
}

// HasRemedyMethods returns a boolean if a field has been set.
func (o *Condition) HasRemedyMethods() bool {
	if o != nil && o.RemedyMethods != nil {
		return true
	}

	return false
}

// SetRemedyMethods gets a reference to the given []RemedyMethod and assigns it to the RemedyMethods field.
func (o *Condition) SetRemedyMethods(v []RemedyMethod) {
	o.RemedyMethods = v
}

func (o Condition) MarshalJSON() ([]byte, error) {
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
		toSerialize["expression"] = o.Expression
	}
	if o.RepeatSchedules != nil {
		toSerialize["repeatSchedules"] = o.RepeatSchedules
	}
	if o.RemedyLogic != nil {
		toSerialize["remedyLogic"] = o.RemedyLogic
	}
	if o.RemedyMethods != nil {
		toSerialize["remedyMethods"] = o.RemedyMethods
	}
	return json.Marshal(toSerialize)
}

type NullableCondition struct {
	value *Condition
	isSet bool
}

func (v NullableCondition) Get() *Condition {
	return v.value
}

func (v *NullableCondition) Set(val *Condition) {
	v.value = val
	v.isSet = true
}

func (v NullableCondition) IsSet() bool {
	return v.isSet
}

func (v *NullableCondition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCondition(val *Condition) *NullableCondition {
	return &NullableCondition{value: val, isSet: true}
}

func (v NullableCondition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCondition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
