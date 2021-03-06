/*
 * AppGate SDP Controller REST API
 *
 * # About   This specification documents the REST API calls for the AppGate SDP Controller.    Please refer to the Integration chapter in the manual or contact AppGate support with any questions about   this functionality. # Getting Started   Requirements for API scripting:   - Access to the peer interface (default port 444) of a Controller appliance.     (https://sdphelp.appgate.com/adminguide/appliances-configure.html?anchor=peer)   - An API user with relevant permissions.     (https://sdphelp.appgate.com/adminguide/administrative-roles-configure.html)   - In order to use the simple login API, Admin MFA must be disabled or the API user must be excluded.     (https://sdphelp.appgate.com/adminguide/mfa-for-admins.html) # Base path   HTTPS requests must be sent to the Peer Interface hostname and port, with **_/admin** path.    For example: **https://appgate.company.com:444/admin**    All requests must have the **Accept** header as:    **application/vnd.appgate.peer-v12+json** # API Conventions   API conventions are  important to understand and follow strictly.    - While updating objects (via PUT), entire object must be sent with all fields.     - For example, if in order to add a remedy method to the condition below:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": []       }       ```     - send the entire object with updated and non-updated fields:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": [{\"type\": \"DisplayMessage\", \"message\": \"test message\"}]       }       ```    - In case Controller returns an error (non-2xx HTTP status code), response body is JSON.     The \"message\" field contains information abot the error.     HTTP 422 \"Unprocessable Entity\" has extra `errors` field to list all the issues with specific fields.    - Empty string (\"\") is considered a different value than \"null\" or field being omitted from JSON.     Omitting the field is recommend if no value is intended.     Empty string (\"\") will be almost always rejected as invalid value.    - There are common pattern between many objects:     - **Configuration Objects**: There are many objects with common fields, namely \"id\", \"name\", \"notes\", \"created\"       and \"updated\". These entities are listed, queried, created, updated and deleted in a similar fashion.     - **Distinguished Name**: Users and Devices are identified with what is called Distinguished Names, as used in        LDAP. The distinguished format that identifies a device and a user combination is        \"CN=\\<Device ID\\>,CN=\\<username\\>,OU=\\<Identity Provider Name\\>\". Some objects have the        \"userDistinguishedName\" field, which does not include the CN for Device ID.        This identifies a user on every device.
 *
 * API version: API version 12
 * Contact: appgatesdp.support@appgate.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"encoding/json"
)

// Condition struct for Condition
type Condition struct {
	BaseEntity
	// Boolean expression in JavaScript.
	Expression string `json:"expression"`
	// A list of schedules that decides when to reevaluate the Condition. All the scheduled times will be effective. One will not override the other. - It can be a time of the day, e.g. 13:00, 10:25, 2:10 etc. - It can be one of the predefined   intervals, e.g. 1m, 5m, 15m, 1h. These intervals   will be always rounded up, i.e. if it's 15m and the   time is 12:07 when the Condition is evaluated   first, then the next evaluation will occur at   12:15, and the next one will be at   12:30 and so on.
	RepeatSchedules *[]string `json:"repeatSchedules,omitempty"`
	// The remedy methods that will be triggered if the evaluation fails.
	RemedyMethods *[]ConditionAllOfRemedyMethods `json:"remedyMethods,omitempty"`
}

// NewCondition instantiates a new Condition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCondition(expression string) *Condition {
	this := Condition{}
	this.Expression = expression
	return &this
}

// NewConditionWithDefaults instantiates a new Condition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConditionWithDefaults() *Condition {
	this := Condition{}
	return &this
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
	return *o.RepeatSchedules
}

// GetRepeatSchedulesOk returns a tuple with the RepeatSchedules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Condition) GetRepeatSchedulesOk() (*[]string, bool) {
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
	o.RepeatSchedules = &v
}

// GetRemedyMethods returns the RemedyMethods field value if set, zero value otherwise.
func (o *Condition) GetRemedyMethods() []ConditionAllOfRemedyMethods {
	if o == nil || o.RemedyMethods == nil {
		var ret []ConditionAllOfRemedyMethods
		return ret
	}
	return *o.RemedyMethods
}

// GetRemedyMethodsOk returns a tuple with the RemedyMethods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Condition) GetRemedyMethodsOk() (*[]ConditionAllOfRemedyMethods, bool) {
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

// SetRemedyMethods gets a reference to the given []ConditionAllOfRemedyMethods and assigns it to the RemedyMethods field.
func (o *Condition) SetRemedyMethods(v []ConditionAllOfRemedyMethods) {
	o.RemedyMethods = &v
}

func (o Condition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedBaseEntity, errBaseEntity := json.Marshal(o.BaseEntity)
	if errBaseEntity != nil {
		return []byte{}, errBaseEntity
	}
	errBaseEntity = json.Unmarshal([]byte(serializedBaseEntity), &toSerialize)
	if errBaseEntity != nil {
		return []byte{}, errBaseEntity
	}
	if true {
		toSerialize["expression"] = o.Expression
	}
	if o.RepeatSchedules != nil {
		toSerialize["repeatSchedules"] = o.RepeatSchedules
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
