/*
Appgate SDP Controller REST API

# About   This specification documents the REST API calls for the Appgate SDP Controller.    Please refer to the REST API chapter in the manual or contact Appgate support with any questions about   this functionality. # Getting Started   Requirements for API scripting:   - Access to the Admin/API TLS Connection (default port 8443) of a Controller appliance.     (https://sdphelp.appgate.com/adminguide/appliance-function-configure.html?anchor=admin-api)   - An API user with relevant permissions.     (https://sdphelp.appgate.com/adminguide/administrative-roles-configure.html)   - In order to use the simple login API, Admin MFA must be disabled or the API user must be excluded.     (https://sdphelp.appgate.com/adminguide/mfa-for-admins.html) # Base path   HTTPS requests must be sent to the Admin Interface hostname and port, with **_/admin** path.    For example: **https://appgate.company.com:8443/admin**    All requests must have the **Accept** header as:    **application/vnd.appgate.peer-v17+json** # API Conventions   API conventions are  important to understand and follow strictly.    - While updating objects (via PUT), entire object must be sent with all fields.     - For example, in order to add a remedy method to the condition below:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": []       }       ```     - send the entire object with updated and non-updated fields:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": [{\"type\": \"DisplayMessage\", \"message\": \"test message\"}]       }       ```    - In case Controller returns an error (non-2xx HTTP status code), response body is JSON.     The \"message\" field contains information about the error.     HTTP 422 \"Unprocessable Entity\" has extra `errors` field to list all the issues with specific fields.    - Empty string (\"\") is considered a different value than \"null\" or field being omitted from JSON.     Omitting the field is recommended if no value is intended.     Empty string (\"\") will be almost always rejected as invalid value.    - There are common pattern between many objects:     - **Configuration Objects**: There are many objects with common fields, namely \"id\", \"name\", \"notes\", \"created\"       and \"updated\". These entities are listed, queried, created, updated and deleted in a similar fashion.     - **Distinguished Name**: Users and Devices are identified with what is called Distinguished Names, as used in        LDAP. The distinguished format that identifies a device and a user combination is        \"CN=\\<Device ID\\>,CN=\\<username\\>,OU=\\<Identity Provider Name\\>\". Some objects have the        \"userDistinguishedName\" field, which does not include the CN for Device ID.        This identifies a user on every device.

API version: API version 17.6
Contact: appgatesdp.support@appgate.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// TestRequest struct for TestRequest
type TestRequest struct {
	// The javascript expression to evaluate.
	Expression   string                 `json:"expression"`
	UserClaims   map[string]interface{} `json:"userClaims,omitempty"`
	DeviceClaims map[string]interface{} `json:"deviceClaims,omitempty"`
	SystemClaims map[string]interface{} `json:"systemClaims,omitempty"`
	Time         *time.Time             `json:"time,omitempty"`
}

// NewTestRequest instantiates a new TestRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestRequest(expression string) *TestRequest {
	this := TestRequest{}
	this.Expression = expression
	return &this
}

// NewTestRequestWithDefaults instantiates a new TestRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestRequestWithDefaults() *TestRequest {
	this := TestRequest{}
	return &this
}

// GetExpression returns the Expression field value
func (o *TestRequest) GetExpression() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value
// and a boolean to check if the value has been set.
func (o *TestRequest) GetExpressionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Expression, true
}

// SetExpression sets field value
func (o *TestRequest) SetExpression(v string) {
	o.Expression = v
}

// GetUserClaims returns the UserClaims field value if set, zero value otherwise.
func (o *TestRequest) GetUserClaims() map[string]interface{} {
	if o == nil || o.UserClaims == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.UserClaims
}

// GetUserClaimsOk returns a tuple with the UserClaims field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestRequest) GetUserClaimsOk() (map[string]interface{}, bool) {
	if o == nil || o.UserClaims == nil {
		return nil, false
	}
	return o.UserClaims, true
}

// HasUserClaims returns a boolean if a field has been set.
func (o *TestRequest) HasUserClaims() bool {
	if o != nil && o.UserClaims != nil {
		return true
	}

	return false
}

// SetUserClaims gets a reference to the given map[string]interface{} and assigns it to the UserClaims field.
func (o *TestRequest) SetUserClaims(v map[string]interface{}) {
	o.UserClaims = v
}

// GetDeviceClaims returns the DeviceClaims field value if set, zero value otherwise.
func (o *TestRequest) GetDeviceClaims() map[string]interface{} {
	if o == nil || o.DeviceClaims == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.DeviceClaims
}

// GetDeviceClaimsOk returns a tuple with the DeviceClaims field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestRequest) GetDeviceClaimsOk() (map[string]interface{}, bool) {
	if o == nil || o.DeviceClaims == nil {
		return nil, false
	}
	return o.DeviceClaims, true
}

// HasDeviceClaims returns a boolean if a field has been set.
func (o *TestRequest) HasDeviceClaims() bool {
	if o != nil && o.DeviceClaims != nil {
		return true
	}

	return false
}

// SetDeviceClaims gets a reference to the given map[string]interface{} and assigns it to the DeviceClaims field.
func (o *TestRequest) SetDeviceClaims(v map[string]interface{}) {
	o.DeviceClaims = v
}

// GetSystemClaims returns the SystemClaims field value if set, zero value otherwise.
func (o *TestRequest) GetSystemClaims() map[string]interface{} {
	if o == nil || o.SystemClaims == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.SystemClaims
}

// GetSystemClaimsOk returns a tuple with the SystemClaims field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestRequest) GetSystemClaimsOk() (map[string]interface{}, bool) {
	if o == nil || o.SystemClaims == nil {
		return nil, false
	}
	return o.SystemClaims, true
}

// HasSystemClaims returns a boolean if a field has been set.
func (o *TestRequest) HasSystemClaims() bool {
	if o != nil && o.SystemClaims != nil {
		return true
	}

	return false
}

// SetSystemClaims gets a reference to the given map[string]interface{} and assigns it to the SystemClaims field.
func (o *TestRequest) SetSystemClaims(v map[string]interface{}) {
	o.SystemClaims = v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *TestRequest) GetTime() time.Time {
	if o == nil || o.Time == nil {
		var ret time.Time
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestRequest) GetTimeOk() (*time.Time, bool) {
	if o == nil || o.Time == nil {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *TestRequest) HasTime() bool {
	if o != nil && o.Time != nil {
		return true
	}

	return false
}

// SetTime gets a reference to the given time.Time and assigns it to the Time field.
func (o *TestRequest) SetTime(v time.Time) {
	o.Time = &v
}

func (o TestRequest) MarshalJSON() ([]byte, error) {
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
	return json.Marshal(toSerialize)
}

type NullableTestRequest struct {
	value *TestRequest
	isSet bool
}

func (v NullableTestRequest) Get() *TestRequest {
	return v.value
}

func (v *NullableTestRequest) Set(val *TestRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTestRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTestRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestRequest(val *TestRequest) *NullableTestRequest {
	return &NullableTestRequest{value: val, isSet: true}
}

func (v NullableTestRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
