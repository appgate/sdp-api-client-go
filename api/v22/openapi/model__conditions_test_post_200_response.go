/*
Appgate SDP Controller REST API

# About   This specification documents the REST API calls for the Appgate SDP Controller.    Please refer to the REST API chapter in the manual or contact Appgate support with any questions about   this functionality. # Getting Started   Requirements for API scripting:   - Access to the Admin/API TLS Connection (default port 8443) of a Controller appliance.     (https://sdphelp.appgate.com/adminguide/appliance-function-configure.html?anchor=admin-api)   - An API user with relevant permissions.     (https://sdphelp.appgate.com/adminguide/administrative-roles-configure.html)   - In order to use the simple login API, Admin MFA must be disabled or the API user must be excluded.     (https://sdphelp.appgate.com/adminguide/mfa-for-admins.html) # Base path   HTTPS requests must be sent to the Admin Interface hostname and port, with **_/admin** path.    For example: **https://appgate.company.com:8443/admin**    All requests must have the **Accept** header as:    **application/vnd.appgate.peer-v22+json**    An exception is made for the **_/admin/version** endpoint which instead expects an **application/json** Accept header. # API Conventions   API conventions are  important to understand and follow strictly.    - While updating objects (via PUT), entire object must be sent with all fields.     - For example, in order to add a remedy method to the condition below:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": []       }       ```     - send the entire object with updated and non-updated fields:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": [{\"type\": \"DisplayMessage\", \"message\": \"test message\"}]       }       ```    - In case Controller returns an error (non-2xx HTTP status code), response body is JSON.     The \"message\" field contains information about the error.     HTTP 422 \"Unprocessable Entity\" has extra `errors` field to list all the issues with specific fields.    - Empty string (\"\") is considered a different value than \"null\" or field being omitted from JSON.     Omitting the field is recommended if no value is intended.     Empty string (\"\") will be almost always rejected as invalid value.    - There are common pattern between many objects:     - **Configuration Objects**: There are many objects with common fields, namely \"id\", \"name\", \"notes\", \"created\"       and \"updated\". These entities are listed, queried, created, updated and deleted in a similar fashion.     - **Distinguished Name**: Users and Devices are identified with what is called Distinguished Names, as used in        LDAP. The distinguished format that identifies a device and a user combination is        \"CN=\\<Device ID\\>,CN=\\<username\\>,OU=\\<Identity Provider Name\\>\". Some objects have the        \"userDistinguishedName\" field, which does not include the CN for Device ID.        This identifies a user on every device.

API version: API version 22.0
Contact: appgatesdp.support@appgate.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// ConditionsTestPost200Response struct for ConditionsTestPost200Response
type ConditionsTestPost200Response struct {
	// Whether the evaluation succeeded or not.
	Result *bool `json:"result,omitempty"`
	// The output logs from the evaluation. Generated by \"console.log\" and \"print\" functions.
	Output *string `json:"output,omitempty"`
	// The error text. Available if the evaluation has an error.
	Error *string `json:"error,omitempty"`
	// How long it took to evaluate the expression.
	ExecutionMs *float32 `json:"executionMs,omitempty"`
}

// NewConditionsTestPost200Response instantiates a new ConditionsTestPost200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConditionsTestPost200Response() *ConditionsTestPost200Response {
	this := ConditionsTestPost200Response{}
	return &this
}

// NewConditionsTestPost200ResponseWithDefaults instantiates a new ConditionsTestPost200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConditionsTestPost200ResponseWithDefaults() *ConditionsTestPost200Response {
	this := ConditionsTestPost200Response{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *ConditionsTestPost200Response) GetResult() bool {
	if o == nil || o.Result == nil {
		var ret bool
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConditionsTestPost200Response) GetResultOk() (*bool, bool) {
	if o == nil || o.Result == nil {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *ConditionsTestPost200Response) HasResult() bool {
	if o != nil && o.Result != nil {
		return true
	}

	return false
}

// SetResult gets a reference to the given bool and assigns it to the Result field.
func (o *ConditionsTestPost200Response) SetResult(v bool) {
	o.Result = &v
}

// GetOutput returns the Output field value if set, zero value otherwise.
func (o *ConditionsTestPost200Response) GetOutput() string {
	if o == nil || o.Output == nil {
		var ret string
		return ret
	}
	return *o.Output
}

// GetOutputOk returns a tuple with the Output field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConditionsTestPost200Response) GetOutputOk() (*string, bool) {
	if o == nil || o.Output == nil {
		return nil, false
	}
	return o.Output, true
}

// HasOutput returns a boolean if a field has been set.
func (o *ConditionsTestPost200Response) HasOutput() bool {
	if o != nil && o.Output != nil {
		return true
	}

	return false
}

// SetOutput gets a reference to the given string and assigns it to the Output field.
func (o *ConditionsTestPost200Response) SetOutput(v string) {
	o.Output = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *ConditionsTestPost200Response) GetError() string {
	if o == nil || o.Error == nil {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConditionsTestPost200Response) GetErrorOk() (*string, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *ConditionsTestPost200Response) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *ConditionsTestPost200Response) SetError(v string) {
	o.Error = &v
}

// GetExecutionMs returns the ExecutionMs field value if set, zero value otherwise.
func (o *ConditionsTestPost200Response) GetExecutionMs() float32 {
	if o == nil || o.ExecutionMs == nil {
		var ret float32
		return ret
	}
	return *o.ExecutionMs
}

// GetExecutionMsOk returns a tuple with the ExecutionMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConditionsTestPost200Response) GetExecutionMsOk() (*float32, bool) {
	if o == nil || o.ExecutionMs == nil {
		return nil, false
	}
	return o.ExecutionMs, true
}

// HasExecutionMs returns a boolean if a field has been set.
func (o *ConditionsTestPost200Response) HasExecutionMs() bool {
	if o != nil && o.ExecutionMs != nil {
		return true
	}

	return false
}

// SetExecutionMs gets a reference to the given float32 and assigns it to the ExecutionMs field.
func (o *ConditionsTestPost200Response) SetExecutionMs(v float32) {
	o.ExecutionMs = &v
}

func (o ConditionsTestPost200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Result != nil {
		toSerialize["result"] = o.Result
	}
	if o.Output != nil {
		toSerialize["output"] = o.Output
	}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	if o.ExecutionMs != nil {
		toSerialize["executionMs"] = o.ExecutionMs
	}
	return json.Marshal(toSerialize)
}

type NullableConditionsTestPost200Response struct {
	value *ConditionsTestPost200Response
	isSet bool
}

func (v NullableConditionsTestPost200Response) Get() *ConditionsTestPost200Response {
	return v.value
}

func (v *NullableConditionsTestPost200Response) Set(val *ConditionsTestPost200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableConditionsTestPost200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableConditionsTestPost200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConditionsTestPost200Response(val *ConditionsTestPost200Response) *NullableConditionsTestPost200Response {
	return &NullableConditionsTestPost200Response{value: val, isSet: true}
}

func (v NullableConditionsTestPost200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConditionsTestPost200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
