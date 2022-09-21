/*
Appgate SDP Controller REST API

# About   This specification documents the REST API calls for the Appgate SDP Controller.    Please refer to the REST API chapter in the manual or contact Appgate support with any questions about   this functionality. # Getting Started   Requirements for API scripting:   - Access to the Admin/API TLS Connection (default port 8443) of a Controller appliance.     (https://sdphelp.appgate.com/adminguide/appliance-function-configure.html?anchor=admin-api)   - An API user with relevant permissions.     (https://sdphelp.appgate.com/adminguide/administrative-roles-configure.html)   - In order to use the simple login API, Admin MFA must be disabled or the API user must be excluded.     (https://sdphelp.appgate.com/adminguide/mfa-for-admins.html) # Base path   HTTPS requests must be sent to the Admin Interface hostname and port, with **_/admin** path.    For example: **https://appgate.company.com:8443/admin**    All requests must have the **Accept** header as:    **application/vnd.appgate.peer-v17+json** # API Conventions   API conventions are  important to understand and follow strictly.    - While updating objects (via PUT), entire object must be sent with all fields.     - For example, in order to add a remedy method to the condition below:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": []       }       ```     - send the entire object with updated and non-updated fields:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": [{\"type\": \"DisplayMessage\", \"message\": \"test message\"}]       }       ```    - In case Controller returns an error (non-2xx HTTP status code), response body is JSON.     The \"message\" field contains information about the error.     HTTP 422 \"Unprocessable Entity\" has extra `errors` field to list all the issues with specific fields.    - Empty string (\"\") is considered a different value than \"null\" or field being omitted from JSON.     Omitting the field is recommended if no value is intended.     Empty string (\"\") will be almost always rejected as invalid value.    - There are common pattern between many objects:     - **Configuration Objects**: There are many objects with common fields, namely \"id\", \"name\", \"notes\", \"created\"       and \"updated\". These entities are listed, queried, created, updated and deleted in a similar fashion.     - **Distinguished Name**: Users and Devices are identified with what is called Distinguished Names, as used in        LDAP. The distinguished format that identifies a device and a user combination is        \"CN=\\<Device ID\\>,CN=\\<username\\>,OU=\\<Identity Provider Name\\>\". Some objects have the        \"userDistinguishedName\" field, which does not include the CN for Device ID.        This identifies a user on every device.

API version: API version 17.4
Contact: appgatesdp.support@appgate.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// ConditionAllOf Represents a Condition.
type ConditionAllOf struct {
	// Boolean expression in JavaScript.
	Expression string `json:"expression"`
	// A list of schedules that decides when to reevaluate the Condition. All the scheduled times will be effective. One will not override the other. - It can be a time of the day, e.g. 13:00, 10:25, 2:10 etc. - It can be one of the predefined   intervals, e.g. 1m, 5m, 15m, 1h. These intervals   will be always rounded up, i.e. if it's 15m and the   time is 12:07 when the Condition is evaluated   first, then the next evaluation will occur at   12:15, and the next one will be at   12:30 and so on.
	RepeatSchedules []string `json:"repeatSchedules,omitempty"`
	// Whether all the Remedy Methods must succeed to pass this Condition or just one.
	RemedyLogic *string `json:"remedyLogic,omitempty"`
	// The remedy methods that will be triggered if the evaluation fails.
	RemedyMethods []RemedyMethod `json:"remedyMethods,omitempty"`
}

// NewConditionAllOf instantiates a new ConditionAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConditionAllOf(expression string) *ConditionAllOf {
	this := ConditionAllOf{}
	this.Expression = expression
	var remedyLogic string = "and"
	this.RemedyLogic = &remedyLogic
	return &this
}

// NewConditionAllOfWithDefaults instantiates a new ConditionAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConditionAllOfWithDefaults() *ConditionAllOf {
	this := ConditionAllOf{}
	var remedyLogic string = "and"
	this.RemedyLogic = &remedyLogic
	return &this
}

// GetExpression returns the Expression field value
func (o *ConditionAllOf) GetExpression() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value
// and a boolean to check if the value has been set.
func (o *ConditionAllOf) GetExpressionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Expression, true
}

// SetExpression sets field value
func (o *ConditionAllOf) SetExpression(v string) {
	o.Expression = v
}

// GetRepeatSchedules returns the RepeatSchedules field value if set, zero value otherwise.
func (o *ConditionAllOf) GetRepeatSchedules() []string {
	if o == nil || o.RepeatSchedules == nil {
		var ret []string
		return ret
	}
	return o.RepeatSchedules
}

// GetRepeatSchedulesOk returns a tuple with the RepeatSchedules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConditionAllOf) GetRepeatSchedulesOk() ([]string, bool) {
	if o == nil || o.RepeatSchedules == nil {
		return nil, false
	}
	return o.RepeatSchedules, true
}

// HasRepeatSchedules returns a boolean if a field has been set.
func (o *ConditionAllOf) HasRepeatSchedules() bool {
	if o != nil && o.RepeatSchedules != nil {
		return true
	}

	return false
}

// SetRepeatSchedules gets a reference to the given []string and assigns it to the RepeatSchedules field.
func (o *ConditionAllOf) SetRepeatSchedules(v []string) {
	o.RepeatSchedules = v
}

// GetRemedyLogic returns the RemedyLogic field value if set, zero value otherwise.
func (o *ConditionAllOf) GetRemedyLogic() string {
	if o == nil || o.RemedyLogic == nil {
		var ret string
		return ret
	}
	return *o.RemedyLogic
}

// GetRemedyLogicOk returns a tuple with the RemedyLogic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConditionAllOf) GetRemedyLogicOk() (*string, bool) {
	if o == nil || o.RemedyLogic == nil {
		return nil, false
	}
	return o.RemedyLogic, true
}

// HasRemedyLogic returns a boolean if a field has been set.
func (o *ConditionAllOf) HasRemedyLogic() bool {
	if o != nil && o.RemedyLogic != nil {
		return true
	}

	return false
}

// SetRemedyLogic gets a reference to the given string and assigns it to the RemedyLogic field.
func (o *ConditionAllOf) SetRemedyLogic(v string) {
	o.RemedyLogic = &v
}

// GetRemedyMethods returns the RemedyMethods field value if set, zero value otherwise.
func (o *ConditionAllOf) GetRemedyMethods() []RemedyMethod {
	if o == nil || o.RemedyMethods == nil {
		var ret []RemedyMethod
		return ret
	}
	return o.RemedyMethods
}

// GetRemedyMethodsOk returns a tuple with the RemedyMethods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConditionAllOf) GetRemedyMethodsOk() ([]RemedyMethod, bool) {
	if o == nil || o.RemedyMethods == nil {
		return nil, false
	}
	return o.RemedyMethods, true
}

// HasRemedyMethods returns a boolean if a field has been set.
func (o *ConditionAllOf) HasRemedyMethods() bool {
	if o != nil && o.RemedyMethods != nil {
		return true
	}

	return false
}

// SetRemedyMethods gets a reference to the given []RemedyMethod and assigns it to the RemedyMethods field.
func (o *ConditionAllOf) SetRemedyMethods(v []RemedyMethod) {
	o.RemedyMethods = v
}

func (o ConditionAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
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

type NullableConditionAllOf struct {
	value *ConditionAllOf
	isSet bool
}

func (v NullableConditionAllOf) Get() *ConditionAllOf {
	return v.value
}

func (v *NullableConditionAllOf) Set(val *ConditionAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableConditionAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableConditionAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConditionAllOf(val *ConditionAllOf) *NullableConditionAllOf {
	return &NullableConditionAllOf{value: val, isSet: true}
}

func (v NullableConditionAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConditionAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
