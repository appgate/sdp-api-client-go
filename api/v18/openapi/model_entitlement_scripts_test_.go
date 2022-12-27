/*
Appgate SDP Controller REST API

# About   This specification documents the REST API calls for the Appgate SDP Controller.    Please refer to the REST API chapter in the manual or contact Appgate support with any questions about   this functionality. # Getting Started   Requirements for API scripting:   - Access to the Admin/API TLS Connection (default port 8443) of a Controller appliance.     (https://sdphelp.appgate.com/adminguide/appliance-function-configure.html?anchor=admin-api)   - An API user with relevant permissions.     (https://sdphelp.appgate.com/adminguide/administrative-roles-configure.html)   - In order to use the simple login API, Admin MFA must be disabled or the API user must be excluded.     (https://sdphelp.appgate.com/adminguide/mfa-for-admins.html) # Base path   HTTPS requests must be sent to the Admin Interface hostname and port, with **_/admin** path.    For example: **https://appgate.company.com:8443/admin**    All requests must have the **Accept** header as:    **application/vnd.appgate.peer-v18+json** # API Conventions   API conventions are  important to understand and follow strictly.    - While updating objects (via PUT), entire object must be sent with all fields.     - For example, in order to add a remedy method to the condition below:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": []       }       ```     - send the entire object with updated and non-updated fields:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": [{\"type\": \"DisplayMessage\", \"message\": \"test message\"}]       }       ```    - In case Controller returns an error (non-2xx HTTP status code), response body is JSON.     The \"message\" field contains information about the error.     HTTP 422 \"Unprocessable Entity\" has extra `errors` field to list all the issues with specific fields.    - Empty string (\"\") is considered a different value than \"null\" or field being omitted from JSON.     Omitting the field is recommended if no value is intended.     Empty string (\"\") will be almost always rejected as invalid value.    - There are common pattern between many objects:     - **Configuration Objects**: There are many objects with common fields, namely \"id\", \"name\", \"notes\", \"created\"       and \"updated\". These entities are listed, queried, created, updated and deleted in a similar fashion.     - **Distinguished Name**: Users and Devices are identified with what is called Distinguished Names, as used in        LDAP. The distinguished format that identifies a device and a user combination is        \"CN=\\<Device ID\\>,CN=\\<username\\>,OU=\\<Identity Provider Name\\>\". Some objects have the        \"userDistinguishedName\" field, which does not include the CN for Device ID.        This identifies a user on every device.

API version: API version 18.2
Contact: appgatesdp.support@appgate.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// EntitlementScriptsTest struct for EntitlementScriptsTest
type EntitlementScriptsTest struct {
	// The javascript expression to evaluate.
	Expression   string                 `json:"expression"`
	UserClaims   map[string]interface{} `json:"userClaims,omitempty"`
	DeviceClaims map[string]interface{} `json:"deviceClaims,omitempty"`
	SystemClaims map[string]interface{} `json:"systemClaims,omitempty"`
	Time         *time.Time             `json:"time,omitempty"`
	// The type of the Entitlement Script.
	Type string `json:"type"`
}

// NewEntitlementScriptsTest instantiates a new EntitlementScriptsTest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitlementScriptsTest(expression string, type_ string) *EntitlementScriptsTest {
	this := EntitlementScriptsTest{}
	this.Expression = expression
	this.Type = type_
	return &this
}

// NewEntitlementScriptsTestWithDefaults instantiates a new EntitlementScriptsTest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitlementScriptsTestWithDefaults() *EntitlementScriptsTest {
	this := EntitlementScriptsTest{}
	return &this
}

// GetExpression returns the Expression field value
func (o *EntitlementScriptsTest) GetExpression() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value
// and a boolean to check if the value has been set.
func (o *EntitlementScriptsTest) GetExpressionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Expression, true
}

// SetExpression sets field value
func (o *EntitlementScriptsTest) SetExpression(v string) {
	o.Expression = v
}

// GetUserClaims returns the UserClaims field value if set, zero value otherwise.
func (o *EntitlementScriptsTest) GetUserClaims() map[string]interface{} {
	if o == nil || o.UserClaims == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.UserClaims
}

// GetUserClaimsOk returns a tuple with the UserClaims field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementScriptsTest) GetUserClaimsOk() (map[string]interface{}, bool) {
	if o == nil || o.UserClaims == nil {
		return nil, false
	}
	return o.UserClaims, true
}

// HasUserClaims returns a boolean if a field has been set.
func (o *EntitlementScriptsTest) HasUserClaims() bool {
	if o != nil && o.UserClaims != nil {
		return true
	}

	return false
}

// SetUserClaims gets a reference to the given map[string]interface{} and assigns it to the UserClaims field.
func (o *EntitlementScriptsTest) SetUserClaims(v map[string]interface{}) {
	o.UserClaims = v
}

// GetDeviceClaims returns the DeviceClaims field value if set, zero value otherwise.
func (o *EntitlementScriptsTest) GetDeviceClaims() map[string]interface{} {
	if o == nil || o.DeviceClaims == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.DeviceClaims
}

// GetDeviceClaimsOk returns a tuple with the DeviceClaims field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementScriptsTest) GetDeviceClaimsOk() (map[string]interface{}, bool) {
	if o == nil || o.DeviceClaims == nil {
		return nil, false
	}
	return o.DeviceClaims, true
}

// HasDeviceClaims returns a boolean if a field has been set.
func (o *EntitlementScriptsTest) HasDeviceClaims() bool {
	if o != nil && o.DeviceClaims != nil {
		return true
	}

	return false
}

// SetDeviceClaims gets a reference to the given map[string]interface{} and assigns it to the DeviceClaims field.
func (o *EntitlementScriptsTest) SetDeviceClaims(v map[string]interface{}) {
	o.DeviceClaims = v
}

// GetSystemClaims returns the SystemClaims field value if set, zero value otherwise.
func (o *EntitlementScriptsTest) GetSystemClaims() map[string]interface{} {
	if o == nil || o.SystemClaims == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.SystemClaims
}

// GetSystemClaimsOk returns a tuple with the SystemClaims field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementScriptsTest) GetSystemClaimsOk() (map[string]interface{}, bool) {
	if o == nil || o.SystemClaims == nil {
		return nil, false
	}
	return o.SystemClaims, true
}

// HasSystemClaims returns a boolean if a field has been set.
func (o *EntitlementScriptsTest) HasSystemClaims() bool {
	if o != nil && o.SystemClaims != nil {
		return true
	}

	return false
}

// SetSystemClaims gets a reference to the given map[string]interface{} and assigns it to the SystemClaims field.
func (o *EntitlementScriptsTest) SetSystemClaims(v map[string]interface{}) {
	o.SystemClaims = v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *EntitlementScriptsTest) GetTime() time.Time {
	if o == nil || o.Time == nil {
		var ret time.Time
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementScriptsTest) GetTimeOk() (*time.Time, bool) {
	if o == nil || o.Time == nil {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *EntitlementScriptsTest) HasTime() bool {
	if o != nil && o.Time != nil {
		return true
	}

	return false
}

// SetTime gets a reference to the given time.Time and assigns it to the Time field.
func (o *EntitlementScriptsTest) SetTime(v time.Time) {
	o.Time = &v
}

// GetType returns the Type field value
func (o *EntitlementScriptsTest) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *EntitlementScriptsTest) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *EntitlementScriptsTest) SetType(v string) {
	o.Type = v
}

func (o EntitlementScriptsTest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["expression"] = o.Expression
	}
	if o.UserClaims != nil {
		toSerialize["userClaims"] = o.UserClaims
	}
	if o.DeviceClaims != nil {
		toSerialize["deviceClaims"] = o.DeviceClaims
	}
	if o.SystemClaims != nil {
		toSerialize["systemClaims"] = o.SystemClaims
	}
	if o.Time != nil {
		toSerialize["time"] = o.Time
	}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableEntitlementScriptsTest struct {
	value *EntitlementScriptsTest
	isSet bool
}

func (v NullableEntitlementScriptsTest) Get() *EntitlementScriptsTest {
	return v.value
}

func (v *NullableEntitlementScriptsTest) Set(val *EntitlementScriptsTest) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitlementScriptsTest) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitlementScriptsTest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitlementScriptsTest(val *EntitlementScriptsTest) *NullableEntitlementScriptsTest {
	return &NullableEntitlementScriptsTest{value: val, isSet: true}
}

func (v NullableEntitlementScriptsTest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitlementScriptsTest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
