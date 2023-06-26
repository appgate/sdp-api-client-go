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
)

// Risk Risk definition.
type Risk struct {
	// The action to take for the risk level.
	Action *string `json:"action,omitempty"`
	// Enter to show a Display Message User Action to the user when the action is Deny.
	DenyMessage *string         `json:"denyMessage,omitempty"`
	UserAction  *RiskUserAction `json:"userAction,omitempty"`
}

// NewRisk instantiates a new Risk object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRisk() *Risk {
	this := Risk{}
	var action string = "Allow"
	this.Action = &action
	return &this
}

// NewRiskWithDefaults instantiates a new Risk object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRiskWithDefaults() *Risk {
	this := Risk{}
	var action string = "Allow"
	this.Action = &action
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *Risk) GetAction() string {
	if o == nil || o.Action == nil {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Risk) GetActionOk() (*string, bool) {
	if o == nil || o.Action == nil {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *Risk) HasAction() bool {
	if o != nil && o.Action != nil {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *Risk) SetAction(v string) {
	o.Action = &v
}

// GetDenyMessage returns the DenyMessage field value if set, zero value otherwise.
func (o *Risk) GetDenyMessage() string {
	if o == nil || o.DenyMessage == nil {
		var ret string
		return ret
	}
	return *o.DenyMessage
}

// GetDenyMessageOk returns a tuple with the DenyMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Risk) GetDenyMessageOk() (*string, bool) {
	if o == nil || o.DenyMessage == nil {
		return nil, false
	}
	return o.DenyMessage, true
}

// HasDenyMessage returns a boolean if a field has been set.
func (o *Risk) HasDenyMessage() bool {
	if o != nil && o.DenyMessage != nil {
		return true
	}

	return false
}

// SetDenyMessage gets a reference to the given string and assigns it to the DenyMessage field.
func (o *Risk) SetDenyMessage(v string) {
	o.DenyMessage = &v
}

// GetUserAction returns the UserAction field value if set, zero value otherwise.
func (o *Risk) GetUserAction() RiskUserAction {
	if o == nil || o.UserAction == nil {
		var ret RiskUserAction
		return ret
	}
	return *o.UserAction
}

// GetUserActionOk returns a tuple with the UserAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Risk) GetUserActionOk() (*RiskUserAction, bool) {
	if o == nil || o.UserAction == nil {
		return nil, false
	}
	return o.UserAction, true
}

// HasUserAction returns a boolean if a field has been set.
func (o *Risk) HasUserAction() bool {
	if o != nil && o.UserAction != nil {
		return true
	}

	return false
}

// SetUserAction gets a reference to the given RiskUserAction and assigns it to the UserAction field.
func (o *Risk) SetUserAction(v RiskUserAction) {
	o.UserAction = &v
}

func (o Risk) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Action != nil {
		toSerialize["action"] = o.Action
	}
	if o.DenyMessage != nil {
		toSerialize["denyMessage"] = o.DenyMessage
	}
	if o.UserAction != nil {
		toSerialize["userAction"] = o.UserAction
	}
	return json.Marshal(toSerialize)
}

type NullableRisk struct {
	value *Risk
	isSet bool
}

func (v NullableRisk) Get() *Risk {
	return v.value
}

func (v *NullableRisk) Set(val *Risk) {
	v.value = val
	v.isSet = true
}

func (v NullableRisk) IsSet() bool {
	return v.isSet
}

func (v *NullableRisk) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRisk(val *Risk) *NullableRisk {
	return &NullableRisk{value: val, isSet: true}
}

func (v NullableRisk) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRisk) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
